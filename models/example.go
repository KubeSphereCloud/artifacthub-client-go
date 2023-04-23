// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Example Example
//
// swagger:model Example
type Example struct {

	// cases
	Cases []*Case `json:"cases"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this example
func (m *Example) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCases(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Example) validateCases(formats strfmt.Registry) error {
	if swag.IsZero(m.Cases) { // not required
		return nil
	}

	for i := 0; i < len(m.Cases); i++ {
		if swag.IsZero(m.Cases[i]) { // not required
			continue
		}

		if m.Cases[i] != nil {
			if err := m.Cases[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cases" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cases" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this example based on the context it is used
func (m *Example) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCases(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Example) contextValidateCases(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Cases); i++ {

		if m.Cases[i] != nil {
			if err := m.Cases[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cases" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cases" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Example) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Example) UnmarshalBinary(b []byte) error {
	var res Example
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
