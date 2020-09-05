// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConditionRequestRelationships condition request relationships
//
// swagger:model conditionRequestRelationships
type ConditionRequestRelationships struct {

	// stream
	Stream *RelatedResourceObject `json:"stream,omitempty"`
}

// Validate validates this condition request relationships
func (m *ConditionRequestRelationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStream(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConditionRequestRelationships) validateStream(formats strfmt.Registry) error {

	if swag.IsZero(m.Stream) { // not required
		return nil
	}

	if m.Stream != nil {
		if err := m.Stream.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stream")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConditionRequestRelationships) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConditionRequestRelationships) UnmarshalBinary(b []byte) error {
	var res ConditionRequestRelationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
