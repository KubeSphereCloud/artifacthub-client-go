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

// WebhookTest WebhookTest
//
// swagger:model WebhookTest
type WebhookTest struct {

	// content type
	// Example: application/json
	ContentType string `json:"content_type,omitempty"`

	// event kinds
	// Example: [0]
	// Required: true
	EventKinds []EventKindID `json:"event_kinds"`

	// template
	// Example: {"text":"Package {{ .Package.Name }} version {{ .Package.Version }} released! {{ .Package.URL }}"}
	Template string `json:"template,omitempty"`

	// url
	// Example: http://url
	// Required: true
	URL *string `json:"url"`
}

// Validate validates this webhook test
func (m *WebhookTest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEventKinds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebhookTest) validateEventKinds(formats strfmt.Registry) error {

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

func (m *WebhookTest) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this webhook test based on the context it is used
func (m *WebhookTest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEventKinds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebhookTest) contextValidateEventKinds(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *WebhookTest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebhookTest) UnmarshalBinary(b []byte) error {
	var res WebhookTest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
