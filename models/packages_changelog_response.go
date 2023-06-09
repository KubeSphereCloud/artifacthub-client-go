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
	"github.com/go-openapi/validate"
)

// PackagesChangelogResponse PackagesChangelogResponse
//
// swagger:model PackagesChangelogResponse
type PackagesChangelogResponse struct {

	// changes
	Changes []*Change `json:"changes"`

	// contains security updates
	// Required: true
	ContainsSecurityUpdates *bool `json:"contains_security_updates"`

	// prerelease
	// Required: true
	Prerelease *bool `json:"prerelease"`

	// ts
	// Required: true
	Ts *int32 `json:"ts"`

	// version
	// Required: true
	Version *string `json:"version"`
}

// Validate validates this packages changelog response
func (m *PackagesChangelogResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChanges(formats); err != nil {
		res = append(res, err)
	}

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

func (m *PackagesChangelogResponse) validateChanges(formats strfmt.Registry) error {
	if swag.IsZero(m.Changes) { // not required
		return nil
	}

	for i := 0; i < len(m.Changes); i++ {
		if swag.IsZero(m.Changes[i]) { // not required
			continue
		}

		if m.Changes[i] != nil {
			if err := m.Changes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("changes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("changes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PackagesChangelogResponse) validateContainsSecurityUpdates(formats strfmt.Registry) error {

	if err := validate.Required("contains_security_updates", "body", m.ContainsSecurityUpdates); err != nil {
		return err
	}

	return nil
}

func (m *PackagesChangelogResponse) validatePrerelease(formats strfmt.Registry) error {

	if err := validate.Required("prerelease", "body", m.Prerelease); err != nil {
		return err
	}

	return nil
}

func (m *PackagesChangelogResponse) validateTs(formats strfmt.Registry) error {

	if err := validate.Required("ts", "body", m.Ts); err != nil {
		return err
	}

	return nil
}

func (m *PackagesChangelogResponse) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this packages changelog response based on the context it is used
func (m *PackagesChangelogResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateChanges(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PackagesChangelogResponse) contextValidateChanges(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Changes); i++ {

		if m.Changes[i] != nil {
			if err := m.Changes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("changes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("changes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PackagesChangelogResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PackagesChangelogResponse) UnmarshalBinary(b []byte) error {
	var res PackagesChangelogResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
