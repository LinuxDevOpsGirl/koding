class SidebarMachineBox extends KDView

  constructor: (options = {}, data) ->

    options.cssClass = "sidebar-machine-box #{data.machine.label}"

    super options, data

    { machine } = data

    @machine = new Machine machine: KD.remote.revive machine

    @addSubView @machineItem = new NavigationMachineItem {}, @machine

    @createWorkspacesLabel()
    @createWorkspacesList()


  createWorkspacesList: ->

    { machine, workspaces } = @getData()

    @listController = new KDListViewController
      itemClass     : NavigationWorkspaceItem
      itemOptions   : { machine }

    @addWorkspace ws  for ws in workspaces

    @listView = @listController.getView()
    @addSubView @listView

    @collapseList()


  addWorkspace: (wsData, storeData = no) ->

    @listController.addItem wsData

    return  unless storeData

    { workspaces } = @getData()
    workspaces.push wsData


  createWorkspacesLabel: ->

    @addSubView @workspacesLabel = new KDCustomHTMLView
      cssClass : 'workspaces-link'
      partial  : 'Workspaces'
      click    : =>
        modal = new MoreWorkspacesModal {}, @getData().workspaces
        modal.once 'NewWorkspaceRequested', @bound 'createAddWorkspaceInput'


  createAddWorkspaceInput: ->

    if @addWorkspaceView
      @addWorkspaceView.input.setFocus()
      return no

    { machine } = @getData()

    data =
      machineUId   : machine.uid
      machineLabel : machine.label

    @addWorkspaceView = new AddWorkspaceView {}, data

    @addWorkspaceView.once 'KDObjectWillBeDestroyed', @bound 'removeAddWorkspaceInput'
    @addWorkspaceView.once 'WorkspaceCreateFailed',   @bound 'removeAddWorkspaceInput'
    @addWorkspaceView.once 'WorkspaceCreated', (ws) =>
      @addWorkspace ws, yes
      @removeAddWorkspaceInput()

    @listView.addSubView @addWorkspaceView

    KD.utils.wait 177, => @addWorkspaceView.input.setFocus()


  removeAddWorkspaceInput: ->

    @addWorkspaceView.destroy()
    @addWorkspaceView = null


  collapseList: ->

    return  if @isListCollapsed

    @listView.setClass 'hidden'
    @workspacesLabel.setClass 'hidden'
    @isListCollapsed = yes


  expandList: ->

    @listView.unsetClass 'hidden'
    @workspacesLabel.unsetClass 'hidden'
    @isListCollapsed = no


  selectWorkspace: (slug) ->

    { machine } = @getData()

    if machine.status.state is Machine.State.Running
      @expandList()

    @deselectWorkspaces()
    @forEachWorkspaceItem (item) ->
      if item.getData().slug is slug
        item.setClass 'selected'


  deselectWorkspaces: ->

    @forEachWorkspaceItem (item) -> item.unsetClass 'selected'


  forEachWorkspaceItem: (callback) ->

    callback item  for item in @listController.getItemsOrdered()
