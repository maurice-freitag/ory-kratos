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

// CompleteSelfServiceLoginFlowWithPasswordMethodReader is a Reader for the CompleteSelfServiceLoginFlowWithPasswordMethod structure.
type CompleteSelfServiceLoginFlowWithPasswordMethodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CompleteSelfServiceLoginFlowWithPasswordMethodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCompleteSelfServiceLoginFlowWithPasswordMethodOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 302:
		result := NewCompleteSelfServiceLoginFlowWithPasswordMethodFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 400:
		result := NewCompleteSelfServiceLoginFlowWithPasswordMethodBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCompleteSelfServiceLoginFlowWithPasswordMethodInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCompleteSelfServiceLoginFlowWithPasswordMethodOK creates a CompleteSelfServiceLoginFlowWithPasswordMethodOK with default headers values
func NewCompleteSelfServiceLoginFlowWithPasswordMethodOK() *CompleteSelfServiceLoginFlowWithPasswordMethodOK {
	return &CompleteSelfServiceLoginFlowWithPasswordMethodOK{}
}

/*CompleteSelfServiceLoginFlowWithPasswordMethodOK handles this case with default header values.

sessionTokenContainer
*/
type CompleteSelfServiceLoginFlowWithPasswordMethodOK struct {
	Payload *models.SessionTokenContainer
}

func (o *CompleteSelfServiceLoginFlowWithPasswordMethodOK) Error() string {
	return fmt.Sprintf("[GET /self-service/login/methods/password][%d] completeSelfServiceLoginFlowWithPasswordMethodOK  %+v", 200, o.Payload)
}

func (o *CompleteSelfServiceLoginFlowWithPasswordMethodOK) GetPayload() *models.SessionTokenContainer {
	return o.Payload
}

func (o *CompleteSelfServiceLoginFlowWithPasswordMethodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SessionTokenContainer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompleteSelfServiceLoginFlowWithPasswordMethodFound creates a CompleteSelfServiceLoginFlowWithPasswordMethodFound with default headers values
func NewCompleteSelfServiceLoginFlowWithPasswordMethodFound() *CompleteSelfServiceLoginFlowWithPasswordMethodFound {
	return &CompleteSelfServiceLoginFlowWithPasswordMethodFound{}
}

/*CompleteSelfServiceLoginFlowWithPasswordMethodFound handles this case with default header values.

Empty responses are sent when, for example, resources are deleted. The HTTP status code for empty responses is
typically 201.
*/
type CompleteSelfServiceLoginFlowWithPasswordMethodFound struct {
}

func (o *CompleteSelfServiceLoginFlowWithPasswordMethodFound) Error() string {
	return fmt.Sprintf("[GET /self-service/login/methods/password][%d] completeSelfServiceLoginFlowWithPasswordMethodFound ", 302)
}

func (o *CompleteSelfServiceLoginFlowWithPasswordMethodFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCompleteSelfServiceLoginFlowWithPasswordMethodBadRequest creates a CompleteSelfServiceLoginFlowWithPasswordMethodBadRequest with default headers values
func NewCompleteSelfServiceLoginFlowWithPasswordMethodBadRequest() *CompleteSelfServiceLoginFlowWithPasswordMethodBadRequest {
	return &CompleteSelfServiceLoginFlowWithPasswordMethodBadRequest{}
}

/*CompleteSelfServiceLoginFlowWithPasswordMethodBadRequest handles this case with default header values.

genericError
*/
type CompleteSelfServiceLoginFlowWithPasswordMethodBadRequest struct {
	Payload *models.GenericError
}

func (o *CompleteSelfServiceLoginFlowWithPasswordMethodBadRequest) Error() string {
	return fmt.Sprintf("[GET /self-service/login/methods/password][%d] completeSelfServiceLoginFlowWithPasswordMethodBadRequest  %+v", 400, o.Payload)
}

func (o *CompleteSelfServiceLoginFlowWithPasswordMethodBadRequest) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *CompleteSelfServiceLoginFlowWithPasswordMethodBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCompleteSelfServiceLoginFlowWithPasswordMethodInternalServerError creates a CompleteSelfServiceLoginFlowWithPasswordMethodInternalServerError with default headers values
func NewCompleteSelfServiceLoginFlowWithPasswordMethodInternalServerError() *CompleteSelfServiceLoginFlowWithPasswordMethodInternalServerError {
	return &CompleteSelfServiceLoginFlowWithPasswordMethodInternalServerError{}
}

/*CompleteSelfServiceLoginFlowWithPasswordMethodInternalServerError handles this case with default header values.

genericError
*/
type CompleteSelfServiceLoginFlowWithPasswordMethodInternalServerError struct {
	Payload *models.GenericError
}

func (o *CompleteSelfServiceLoginFlowWithPasswordMethodInternalServerError) Error() string {
	return fmt.Sprintf("[GET /self-service/login/methods/password][%d] completeSelfServiceLoginFlowWithPasswordMethodInternalServerError  %+v", 500, o.Payload)
}

func (o *CompleteSelfServiceLoginFlowWithPasswordMethodInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *CompleteSelfServiceLoginFlowWithPasswordMethodInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
