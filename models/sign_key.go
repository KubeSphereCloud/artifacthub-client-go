// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SignKey SignKey
//
// swagger:model SignKey
type SignKey struct {

	// fingerprint
	// Example: 0011223344
	Fingerprint string `json:"fingerprint,omitempty"`

	// url
	// Example: https://key.url
	URL string `json:"url,omitempty"`
}

// Validate validates this sign key
func (m *SignKey) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this sign key based on context it is used
func (m *SignKey) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SignKey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SignKey) UnmarshalBinary(b []byte) error {
	var res SignKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}