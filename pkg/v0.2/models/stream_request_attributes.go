// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StreamRequestAttributes stream request attributes
//
// swagger:model streamRequestAttributes
type StreamRequestAttributes struct {

	// Custom JSON data that can be set by an end user and will be included in notifications.
	CustomData map[string]interface{} `json:"custom_data,omitempty"`

	// Name for the stream (free-form string)
	// Required: true
	Name *string `json:"name"`

	// The query string itself (see Query Syntax section for details).
	// Once a stream has been created, this string cannot be modified.
	// <b>Required when creating new streams.</b>
	Query string `json:"query,omitempty"`
}

// Validate validates this stream request attributes
func (m *StreamRequestAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StreamRequestAttributes) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StreamRequestAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StreamRequestAttributes) UnmarshalBinary(b []byte) error {
	var res StreamRequestAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
