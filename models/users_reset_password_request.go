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

// UsersResetPasswordRequest UsersResetPasswordRequest
//
// swagger:model UsersResetPasswordRequest
type UsersResetPasswordRequest struct {

	// code
	// Required: true
	// Format: uuid
	Code *strfmt.UUID `json:"code"`

	// password
	// Required: true
	Password *string `json:"password"`
}

// Validate validates this users reset password request
func (m *UsersResetPasswordRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UsersResetPasswordRequest) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("code", "body", m.Code); err != nil {
		return err
	}

	if err := validate.FormatOf("code", "body", "uuid", m.Code.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UsersResetPasswordRequest) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this users reset password request based on context it is used
func (m *UsersResetPasswordRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UsersResetPasswordRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UsersResetPasswordRequest) UnmarshalBinary(b []byte) error {
	var res UsersResetPasswordRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
