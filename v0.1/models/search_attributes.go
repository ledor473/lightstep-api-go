// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SearchAttributes search attributes
// swagger:model searchAttributes
type SearchAttributes struct {

	// Custom JSON data that can be set by an end user and will be included in notifications.
	CustomData map[string]interface{} `json:"custom-data,omitempty"`

	// Name for the stream (free-form string)
	// Required: true
	Name *string `json:"name"`

	// The query string itself (see Query Syntax section for details).
	// Once a search has been created, this string cannot be modified.
	// <b>Required when creating new streams.</b>
	Query string `json:"query,omitempty"`
}

// Validate validates this search attributes
func (m *SearchAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SearchAttributes) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SearchAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SearchAttributes) UnmarshalBinary(b []byte) error {
	var res SearchAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
