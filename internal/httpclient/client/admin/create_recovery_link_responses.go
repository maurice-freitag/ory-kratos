// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ory/kratos-client-go/models"
)

// CreateRecoveryLinkReader is a Reader for the CreateRecoveryLink structure.
type CreateRecoveryLinkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRecoveryLinkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateRecoveryLinkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateRecoveryLinkBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateRecoveryLinkNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateRecoveryLinkInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateRecoveryLinkOK creates a CreateRecoveryLinkOK with default headers values
func NewCreateRecoveryLinkOK() *CreateRecoveryLinkOK {
	return &CreateRecoveryLinkOK{}
}

/*CreateRecoveryLinkOK handles this case with default header values.

recoveryLink
*/
type CreateRecoveryLinkOK struct {
	Payload *models.RecoveryLink
}

func (o *CreateRecoveryLinkOK) Error() string {
	return fmt.Sprintf("[POST /recovery/link][%d] createRecoveryLinkOK  %+v", 200, o.Payload)
}

func (o *CreateRecoveryLinkOK) GetPayload() *models.RecoveryLink {
	return o.Payload
}

func (o *CreateRecoveryLinkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RecoveryLink)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRecoveryLinkBadRequest creates a CreateRecoveryLinkBadRequest with default headers values
func NewCreateRecoveryLinkBadRequest() *CreateRecoveryLinkBadRequest {
	return &CreateRecoveryLinkBadRequest{}
}

/*CreateRecoveryLinkBadRequest handles this case with default header values.

genericError
*/
type CreateRecoveryLinkBadRequest struct {
	Payload *models.GenericError
}

func (o *CreateRecoveryLinkBadRequest) Error() string {
	return fmt.Sprintf("[POST /recovery/link][%d] createRecoveryLinkBadRequest  %+v", 400, o.Payload)
}

func (o *CreateRecoveryLinkBadRequest) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *CreateRecoveryLinkBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRecoveryLinkNotFound creates a CreateRecoveryLinkNotFound with default headers values
func NewCreateRecoveryLinkNotFound() *CreateRecoveryLinkNotFound {
	return &CreateRecoveryLinkNotFound{}
}

/*CreateRecoveryLinkNotFound handles this case with default header values.

genericError
*/
type CreateRecoveryLinkNotFound struct {
	Payload *models.GenericError
}

func (o *CreateRecoveryLinkNotFound) Error() string {
	return fmt.Sprintf("[POST /recovery/link][%d] createRecoveryLinkNotFound  %+v", 404, o.Payload)
}

func (o *CreateRecoveryLinkNotFound) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *CreateRecoveryLinkNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRecoveryLinkInternalServerError creates a CreateRecoveryLinkInternalServerError with default headers values
func NewCreateRecoveryLinkInternalServerError() *CreateRecoveryLinkInternalServerError {
	return &CreateRecoveryLinkInternalServerError{}
}

/*CreateRecoveryLinkInternalServerError handles this case with default header values.

genericError
*/
type CreateRecoveryLinkInternalServerError struct {
	Payload *models.GenericError
}

func (o *CreateRecoveryLinkInternalServerError) Error() string {
	return fmt.Sprintf("[POST /recovery/link][%d] createRecoveryLinkInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateRecoveryLinkInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *CreateRecoveryLinkInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CreateRecoveryLinkBody create recovery link body
swagger:model CreateRecoveryLinkBody
*/
type CreateRecoveryLinkBody struct {

	// Link Expires In
	//
	// The recovery link will expire at that point in time. Defaults to the configuration value of
	// `selfservice.flows.recovery.lifespan`.
	// Pattern: ^[0-9]+(ns|us|ms|s|m|h)$
	ExpiresIn string `json:"expires_in,omitempty"`

	// identity id
	// Required: true
	// Format: uuid4
	IdentityID models.UUID `json:"identity_id"`
}

// Validate validates this create recovery link body
func (o *CreateRecoveryLinkBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateExpiresIn(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateIdentityID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateRecoveryLinkBody) validateExpiresIn(formats strfmt.Registry) error {

	if swag.IsZero(o.ExpiresIn) { // not required
		return nil
	}

	if err := validate.Pattern("Body"+"."+"expires_in", "body", string(o.ExpiresIn), `^[0-9]+(ns|us|ms|s|m|h)$`); err != nil {
		return err
	}

	return nil
}

func (o *CreateRecoveryLinkBody) validateIdentityID(formats strfmt.Registry) error {

	if err := o.IdentityID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Body" + "." + "identity_id")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateRecoveryLinkBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateRecoveryLinkBody) UnmarshalBinary(b []byte) error {
	var res CreateRecoveryLinkBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
