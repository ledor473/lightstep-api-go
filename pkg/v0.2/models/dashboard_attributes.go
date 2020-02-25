// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DashboardAttributes dashboard attributes
// swagger:model dashboardAttributes
type DashboardAttributes struct {

	// Name for the dashboard (free-form string)
	// Required: true
	Name *string `json:"name"`

	// Streams to be added to the dashboard. If the stream exists already, include the ID.
	// If it does not, do not include an ID. A new stream will be created
	Streams []*StreamResponse `json:"streams"`
}

// Validate validates this dashboard attributes
func (m *DashboardAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStreams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DashboardAttributes) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *DashboardAttributes) validateStreams(formats strfmt.Registry) error {

	if swag.IsZero(m.Streams) { // not required
		return nil
	}

	for i := 0; i < len(m.Streams); i++ {
		if swag.IsZero(m.Streams[i]) { // not required
			continue
		}

		if m.Streams[i] != nil {
			if err := m.Streams[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("streams" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DashboardAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DashboardAttributes) UnmarshalBinary(b []byte) error {
	var res DashboardAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}