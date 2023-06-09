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

// SubscriptionBody SubscriptionBody
//
// swagger:model SubscriptionBody
type SubscriptionBody struct {

	// event kind
	// Required: true
	EventKind *EventKindID `json:"event_kind"`

	// package id
	// Required: true
	// Format: uuid
	PackageID *strfmt.UUID `json:"package_id"`
}

// Validate validates this subscription body
func (m *SubscriptionBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEventKind(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePackageID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubscriptionBody) validateEventKind(formats strfmt.Registry) error {

	if err := validate.Required("event_kind", "body", m.EventKind); err != nil {
		return err
	}

	if err := validate.Required("event_kind", "body", m.EventKind); err != nil {
		return err
	}

	if m.EventKind != nil {
		if err := m.EventKind.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("event_kind")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("event_kind")
			}
			return err
		}
	}

	return nil
}

func (m *SubscriptionBody) validatePackageID(formats strfmt.Registry) error {

	if err := validate.Required("package_id", "body", m.PackageID); err != nil {
		return err
	}

	if err := validate.FormatOf("package_id", "body", "uuid", m.PackageID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this subscription body based on the context it is used
func (m *SubscriptionBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEventKind(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubscriptionBody) contextValidateEventKind(ctx context.Context, formats strfmt.Registry) error {

	if m.EventKind != nil {
		if err := m.EventKind.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("event_kind")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("event_kind")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SubscriptionBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubscriptionBody) UnmarshalBinary(b []byte) error {
	var res SubscriptionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
