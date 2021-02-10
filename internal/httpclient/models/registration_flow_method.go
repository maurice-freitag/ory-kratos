// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RegistrationFlowMethod RegistrationFlowMethod registration flow method
//
// swagger:model registrationFlowMethod
type RegistrationFlowMethod struct {

	// config
	// Required: true
	Config *RegistrationFlowMethodConfig `json:"config"`

	// method
	// Required: true
	Method *CredentialsType `json:"method"`
}

// Validate validates this registration flow method
func (m *RegistrationFlowMethod) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMethod(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegistrationFlowMethod) validateConfig(formats strfmt.Registry) error {

	if err := validate.Required("config", "body", m.Config); err != nil {
		return err
	}

	if m.Config != nil {
		if err := m.Config.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

func (m *RegistrationFlowMethod) validateMethod(formats strfmt.Registry) error {

	if err := validate.Required("method", "body", m.Method); err != nil {
		return err
	}

	if err := validate.Required("method", "body", m.Method); err != nil {
		return err
	}

	if m.Method != nil {
		if err := m.Method.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("method")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this registration flow method based on the context it is used
func (m *RegistrationFlowMethod) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMethod(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RegistrationFlowMethod) contextValidateConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.Config != nil {
		if err := m.Config.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

func (m *RegistrationFlowMethod) contextValidateMethod(ctx context.Context, formats strfmt.Registry) error {

	if m.Method != nil {
		if err := m.Method.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("method")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RegistrationFlowMethod) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RegistrationFlowMethod) UnmarshalBinary(b []byte) error {
	var res RegistrationFlowMethod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
