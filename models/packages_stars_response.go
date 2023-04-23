// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PackagesStarsResponse PackagesStarsResponse
//
// swagger:model PackagesStarsResponse
type PackagesStarsResponse struct {

	// starred by user
	StarredByUser bool `json:"starred_by_user,omitempty"`

	// stars
	Stars int32 `json:"stars,omitempty"`
}

// Validate validates this packages stars response
func (m *PackagesStarsResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this packages stars response based on context it is used
func (m *PackagesStarsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PackagesStarsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PackagesStarsResponse) UnmarshalBinary(b []byte) error {
	var res PackagesStarsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
