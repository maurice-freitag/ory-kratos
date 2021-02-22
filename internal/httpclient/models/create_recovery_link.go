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

// CreateRecoveryLink CreateRecoveryLink CreateRecoveryLink CreateRecoveryLink CreateRecoveryLink CreateRecoveryLink CreateRecoveryLink CreateRecoveryLink CreateRecoveryLink CreateRecoveryLink CreateRecoveryLink CreateRecoveryLink CreateRecoveryLink CreateRecoveryLink CreateRecoveryLink CreateRecoveryLink CreateRecoveryLink CreateRecoveryLink create recovery link
//
// swagger:model CreateRecoveryLink
type CreateRecoveryLink struct {

	// Link Expires In
	//
	// The recovery link will expire at that point in time. Defaults to the configuration value of
	// `selfservice.flows.recovery.request_lifespan`.
	// Pattern: ^[0-9]+(ns|us|ms|s|m|h)$
	ExpiresIn string `json:"expires_in,omitempty"`

	// identity id
	// Required: true
	// Format: uuid4
	IdentityID *UUID `json:"identity_id"`
}

// Validate validates this create recovery link
func (m *CreateRecoveryLink) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExpiresIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentityID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateRecoveryLink) validateExpiresIn(formats strfmt.Registry) error {
	if swag.IsZero(m.ExpiresIn) { // not required
		return nil
	}

	if err := validate.Pattern("expires_in", "body", m.ExpiresIn, `^[0-9]+(ns|us|ms|s|m|h)$`); err != nil {
		return err
	}

	return nil
}

func (m *CreateRecoveryLink) validateIdentityID(formats strfmt.Registry) error {

	if err := validate.Required("identity_id", "body", m.IdentityID); err != nil {
		return err
	}

	if err := validate.Required("identity_id", "body", m.IdentityID); err != nil {
		return err
	}

	if m.IdentityID != nil {
		if err := m.IdentityID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("identity_id")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this create recovery link based on the context it is used
func (m *CreateRecoveryLink) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIdentityID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateRecoveryLink) contextValidateIdentityID(ctx context.Context, formats strfmt.Registry) error {

	if m.IdentityID != nil {
		if err := m.IdentityID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("identity_id")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateRecoveryLink) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateRecoveryLink) UnmarshalBinary(b []byte) error {
	var res CreateRecoveryLink
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
