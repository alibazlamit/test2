package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// PortRange port range
// swagger:model PortRange
type PortRange struct {

	// The maximum port number. Must not be lower than the minimum port number. To specify
	// a single port number, set both the min and max to the same value.
	//
	// Required: true
	// Maximum: 65535
	// Minimum: 1
	Max *int64 `json:"max"`

	// The minimum port number. Must not be greater than the maximum port number.
	// Required: true
	// Maximum: 65535
	// Minimum: 1
	Min *int64 `json:"min"`
}

// Validate validates this port range
func (m *PortRange) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMax(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMin(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PortRange) validateMax(formats strfmt.Registry) error {

	if err := validate.Required("max", "body", m.Max); err != nil {
		return err
	}

	if err := validate.MinimumInt("max", "body", int64(*m.Max), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("max", "body", int64(*m.Max), 65535, false); err != nil {
		return err
	}

	return nil
}

func (m *PortRange) validateMin(formats strfmt.Registry) error {

	if err := validate.Required("min", "body", m.Min); err != nil {
		return err
	}

	if err := validate.MinimumInt("min", "body", int64(*m.Min), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("min", "body", int64(*m.Min), 65535, false); err != nil {
		return err
	}

	return nil
}
