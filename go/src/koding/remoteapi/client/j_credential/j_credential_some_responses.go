package j_credential

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// JCredentialSomeReader is a Reader for the JCredentialSome structure.
type JCredentialSomeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JCredentialSomeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewJCredentialSomeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewJCredentialSomeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJCredentialSomeOK creates a JCredentialSomeOK with default headers values
func NewJCredentialSomeOK() *JCredentialSomeOK {
	return &JCredentialSomeOK{}
}

/*JCredentialSomeOK handles this case with default header values.

Request processed successfully
*/
type JCredentialSomeOK struct {
	Payload *models.DefaultResponse
}

func (o *JCredentialSomeOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JCredential.some][%d] jCredentialSomeOK  %+v", 200, o.Payload)
}

func (o *JCredentialSomeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewJCredentialSomeUnauthorized creates a JCredentialSomeUnauthorized with default headers values
func NewJCredentialSomeUnauthorized() *JCredentialSomeUnauthorized {
	return &JCredentialSomeUnauthorized{}
}

/*JCredentialSomeUnauthorized handles this case with default header values.

Unauthorized request
*/
type JCredentialSomeUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *JCredentialSomeUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/JCredential.some][%d] jCredentialSomeUnauthorized  %+v", 401, o.Payload)
}

func (o *JCredentialSomeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
