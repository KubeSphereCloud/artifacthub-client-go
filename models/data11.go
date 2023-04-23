// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Data11 Data11
//
// swagger:model Data11
type Data11 struct {

	// examples
	Examples interface{} `json:"examples,omitempty"`

	// manifest raw
	ManifestRaw map[string]string `json:"manifestRaw,omitempty"`

	// pipelines min version
	// Example: 2.0.0
	PipelinesMinVersion string `json:"pipelines.minVersion,omitempty"`

	// platforms
	Platforms []string `json:"platforms"`
}

// Validate validates this data11
func (m *Data11) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this data11 based on context it is used
func (m *Data11) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Data11) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Data11) UnmarshalBinary(b []byte) error {
	var res Data11
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}