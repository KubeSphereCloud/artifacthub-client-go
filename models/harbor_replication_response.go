// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HarborReplicationResponse HarborReplicationResponse
//
// swagger:model HarborReplicationResponse
type HarborReplicationResponse struct {

	// package
	Package string `json:"package,omitempty"`

	// repository
	Repository string `json:"repository,omitempty"`

	// url
	URL string `json:"url,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this harbor replication response
func (m *HarborReplicationResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this harbor replication response based on context it is used
func (m *HarborReplicationResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HarborReplicationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HarborReplicationResponse) UnmarshalBinary(b []byte) error {
	var res HarborReplicationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}