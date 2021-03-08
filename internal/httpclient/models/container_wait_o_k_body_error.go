// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ContainerWaitOKBodyError ContainerWaitOKBodyError ContainerWaitOKBodyError container waiting error, if any
//
// swagger:model ContainerWaitOKBodyError
type ContainerWaitOKBodyError struct {

	// Details of an error
	Message string `json:"Message,omitempty"`
}

// Validate validates this container wait o k body error
func (m *ContainerWaitOKBodyError) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container wait o k body error based on context it is used
func (m *ContainerWaitOKBodyError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ContainerWaitOKBodyError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContainerWaitOKBodyError) UnmarshalBinary(b []byte) error {
	var res ContainerWaitOKBodyError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
