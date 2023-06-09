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

// AvailableVersion AvailableVersion
//
// swagger:model AvailableVersion
type AvailableVersion struct {

	// contains security updates
	// Required: true
	ContainsSecurityUpdates *bool `json:"contains_security_updates"`

	// prerelease
	// Required: true
	Prerelease *bool `json:"prerelease"`

	// ts
	// Example: 1618431211
	// Required: true
	Ts *int32 `json:"ts"`

	// version
	// Example: 1.0.0
	// Required: true
	Version *string `json:"version"`
}

// Validate validates this available version
func (m *AvailableVersion) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContainsSecurityUpdates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrerelease(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AvailableVersion) validateContainsSecurityUpdates(formats strfmt.Registry) error {

	if err := validate.Required("contains_security_updates", "body", m.ContainsSecurityUpdates); err != nil {
		return err
	}

	return nil
}

func (m *AvailableVersion) validatePrerelease(formats strfmt.Registry) error {

	if err := validate.Required("prerelease", "body", m.Prerelease); err != nil {
		return err
	}

	return nil
}

func (m *AvailableVersion) validateTs(formats strfmt.Registry) error {

	if err := validate.Required("ts", "body", m.Ts); err != nil {
		return err
	}

	return nil
}

func (m *AvailableVersion) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this available version based on context it is used
func (m *AvailableVersion) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AvailableVersion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AvailableVersion) UnmarshalBinary(b []byte) error {
	var res AvailableVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
