// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ory/kratos/internal/httpclient/models"
)

// InitializeSelfServiceBrowserRegistrationFlowReader is a Reader for the InitializeSelfServiceBrowserRegistrationFlow structure.
type InitializeSelfServiceBrowserRegistrationFlowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InitializeSelfServiceBrowserRegistrationFlowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 302:
		result := NewInitializeSelfServiceBrowserRegistrationFlowFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewInitializeSelfServiceBrowserRegistrationFlowInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewInitializeSelfServiceBrowserRegistrationFlowFound creates a InitializeSelfServiceBrowserRegistrationFlowFound with default headers values
func NewInitializeSelfServiceBrowserRegistrationFlowFound() *InitializeSelfServiceBrowserRegistrationFlowFound {
	return &InitializeSelfServiceBrowserRegistrationFlowFound{}
}

/*InitializeSelfServiceBrowserRegistrationFlowFound handles this case with default header values.

Empty responses are sent when, for example, resources are deleted. The HTTP status code for empty responses is
typically 201.
*/
type InitializeSelfServiceBrowserRegistrationFlowFound struct {
}

func (o *InitializeSelfServiceBrowserRegistrationFlowFound) Error() string {
	return fmt.Sprintf("[GET /self-service/browser/flows/registration][%d] initializeSelfServiceBrowserRegistrationFlowFound ", 302)
}

func (o *InitializeSelfServiceBrowserRegistrationFlowFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInitializeSelfServiceBrowserRegistrationFlowInternalServerError creates a InitializeSelfServiceBrowserRegistrationFlowInternalServerError with default headers values
func NewInitializeSelfServiceBrowserRegistrationFlowInternalServerError() *InitializeSelfServiceBrowserRegistrationFlowInternalServerError {
	return &InitializeSelfServiceBrowserRegistrationFlowInternalServerError{}
}

/*InitializeSelfServiceBrowserRegistrationFlowInternalServerError handles this case with default header values.

genericError
*/
type InitializeSelfServiceBrowserRegistrationFlowInternalServerError struct {
	Payload *models.GenericError
}

func (o *InitializeSelfServiceBrowserRegistrationFlowInternalServerError) Error() string {
	return fmt.Sprintf("[GET /self-service/browser/flows/registration][%d] initializeSelfServiceBrowserRegistrationFlowInternalServerError  %+v", 500, o.Payload)
}

func (o *InitializeSelfServiceBrowserRegistrationFlowInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *InitializeSelfServiceBrowserRegistrationFlowInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
