package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// CrossConnectPortSpeedShape An individual port speed level for cross-connects.
//
// swagger:model CrossConnectPortSpeedShape
type CrossConnectPortSpeedShape struct {

	// The name of the port speed shape.
	//
	// Example: `10 Gbps`
	//
	// Required: true
	// Max Length: 255
	// Min Length: 1
	Name *string `json:"name"`

	// The port speed in Gbps.
	//
	// Example: `10`
	//
	// Required: true
	// Max Length: 255
	// Min Length: 1
	PortSpeedInGbps *int64 `json:"portSpeedInGbps"`
}

// Validate validates this cross connect port speed shape
func (m *CrossConnectPortSpeedShape) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePortSpeedInGbps(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CrossConnectPortSpeedShape) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 255); err != nil {
		return err
	}

	return nil
}

func (m *CrossConnectPortSpeedShape) validatePortSpeedInGbps(formats strfmt.Registry) error {

	if err := validate.Required("portSpeedInGbps", "body", m.PortSpeedInGbps); err != nil {
		return err
	}

	if err := validate.MinLength("portSpeedInGbps", "body", string(*m.PortSpeedInGbps), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("portSpeedInGbps", "body", string(*m.PortSpeedInGbps), 255); err != nil {
		return err
	}

	return nil
}
