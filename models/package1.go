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

// Package1 Package1
//
// swagger:model Package1
type Package1 struct {

	// package id
	// Required: true
	// Format: uuid
	PackageID *strfmt.UUID `json:"package_id"`
}

// Validate validates this package1
func (m *Package1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePackageID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Package1) validatePackageID(formats strfmt.Registry) error {

	if err := validate.Required("package_id", "body", m.PackageID); err != nil {
		return err
	}

	if err := validate.FormatOf("package_id", "body", "uuid", m.PackageID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this package1 based on context it is used
func (m *Package1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Package1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Package1) UnmarshalBinary(b []byte) error {
	var res Package1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
