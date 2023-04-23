// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// Sort sort
// Example: relevance
//
// swagger:model sort
type Sort string

func NewSort(value Sort) *Sort {
	return &value
}

// Pointer returns a pointer to a freshly-allocated Sort.
func (m Sort) Pointer() *Sort {
	return &m
}

const (

	// SortRelevance captures enum value "relevance"
	SortRelevance Sort = "relevance"

	// SortStars captures enum value "stars"
	SortStars Sort = "stars"
)

// for schema
var sortEnum []interface{}

func init() {
	var res []Sort
	if err := json.Unmarshal([]byte(`["relevance","stars"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sortEnum = append(sortEnum, v)
	}
}

func (m Sort) validateSortEnum(path, location string, value Sort) error {
	if err := validate.EnumCase(path, location, value, sortEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this sort
func (m Sort) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateSortEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this sort based on context it is used
func (m Sort) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
