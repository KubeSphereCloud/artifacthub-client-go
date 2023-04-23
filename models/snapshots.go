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

// Snapshots Snapshots
//
// swagger:model Snapshots
type Snapshots struct {

	// created monthly
	CreatedMonthly [][]int32 `json:"created_monthly"`

	// running total
	RunningTotal [][]int32 `json:"running_total"`

	// total
	// Required: true
	Total *int32 `json:"total"`
}

// Validate validates this snapshots
func (m *Snapshots) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTotal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Snapshots) validateTotal(formats strfmt.Registry) error {

	if err := validate.Required("total", "body", m.Total); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this snapshots based on context it is used
func (m *Snapshots) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Snapshots) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Snapshots) UnmarshalBinary(b []byte) error {
	var res Snapshots
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
