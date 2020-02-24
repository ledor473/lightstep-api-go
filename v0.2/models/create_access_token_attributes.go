// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateAccessTokenAttributes create access token attributes
// swagger:model createAccessTokenAttributes
type CreateAccessTokenAttributes struct {

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this create access token attributes
func (m *CreateAccessTokenAttributes) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateAccessTokenAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateAccessTokenAttributes) UnmarshalBinary(b []byte) error {
	var res CreateAccessTokenAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
