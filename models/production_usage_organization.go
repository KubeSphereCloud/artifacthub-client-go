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

// ProductionUsageOrganization ProductionUsageOrganization
//
// swagger:model ProductionUsageOrganization
type ProductionUsageOrganization struct {

	// display name
	// Example: Organization 1
	DisplayName string `json:"display_name,omitempty"`

	// home url
	// Example: http://url
	HomeURL string `json:"home_url,omitempty"`

	// logo image id
	// Example: 12345abcde
	LogoImageID string `json:"logo_image_id,omitempty"`

	// name
	// Example: org1
	// Required: true
	Name *string `json:"name"`

	// used in production
	UsedInProduction bool `json:"used_in_production,omitempty"`
}

// Validate validates this production usage organization
func (m *ProductionUsageOrganization) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProductionUsageOrganization) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this production usage organization based on context it is used
func (m *ProductionUsageOrganization) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProductionUsageOrganization) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProductionUsageOrganization) UnmarshalBinary(b []byte) error {
	var res ProductionUsageOrganization
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}