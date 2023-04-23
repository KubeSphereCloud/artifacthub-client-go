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

// SubscriptionsResponse SubscriptionsResponse
//
// swagger:model SubscriptionsResponse
type SubscriptionsResponse struct {

	// event kinds
	// Required: true
	EventKinds []EventKindID `json:"event_kinds"`

	// logo image id
	// Example: 12345abcde
	LogoImageID string `json:"logo_image_id,omitempty"`

	// name
	// Example: pkg1
	// Required: true
	Name *string `json:"name"`

	// normalized name
	// Example: pkg1
	// Required: true
	NormalizedName *string `json:"normalized_name"`

	// package id
	// Required: true
	// Format: uuid
	PackageID *strfmt.UUID `json:"package_id"`

	// repository
	// Required: true
	Repository *RepositorySummary `json:"repository"`
}

// Validate validates this subscriptions response
func (m *SubscriptionsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEventKinds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNormalizedName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePackageID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRepository(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubscriptionsResponse) validateEventKinds(formats strfmt.Registry) error {

	if err := validate.Required("event_kinds", "body", m.EventKinds); err != nil {
		return err
	}

	for i := 0; i < len(m.EventKinds); i++ {

		if err := m.EventKinds[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("event_kinds" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("event_kinds" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *SubscriptionsResponse) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *SubscriptionsResponse) validateNormalizedName(formats strfmt.Registry) error {

	if err := validate.Required("normalized_name", "body", m.NormalizedName); err != nil {
		return err
	}

	return nil
}

func (m *SubscriptionsResponse) validatePackageID(formats strfmt.Registry) error {

	if err := validate.Required("package_id", "body", m.PackageID); err != nil {
		return err
	}

	if err := validate.FormatOf("package_id", "body", "uuid", m.PackageID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SubscriptionsResponse) validateRepository(formats strfmt.Registry) error {

	if err := validate.Required("repository", "body", m.Repository); err != nil {
		return err
	}

	if m.Repository != nil {
		if err := m.Repository.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("repository")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("repository")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this subscriptions response based on the context it is used
func (m *SubscriptionsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEventKinds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRepository(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubscriptionsResponse) contextValidateEventKinds(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EventKinds); i++ {

		if err := m.EventKinds[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("event_kinds" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("event_kinds" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *SubscriptionsResponse) contextValidateRepository(ctx context.Context, formats strfmt.Registry) error {

	if m.Repository != nil {
		if err := m.Repository.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("repository")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("repository")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SubscriptionsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubscriptionsResponse) UnmarshalBinary(b []byte) error {
	var res SubscriptionsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
