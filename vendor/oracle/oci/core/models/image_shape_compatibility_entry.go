package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// ImageShapeCompatibilityEntry An image and shape that are compatible.
// swagger:model ImageShapeCompatibilityEntry
type ImageShapeCompatibilityEntry struct {

	// The image OCID.
	// Required: true
	// Max Length: 255
	// Min Length: 1
	ImageID *string `json:"imageId"`

	// The shape name.
	// Required: true
	// Max Length: 255
	// Min Length: 1
	Shape *string `json:"shape"`
}

// Validate validates this image shape compatibility entry
func (m *ImageShapeCompatibilityEntry) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImageID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateShape(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ImageShapeCompatibilityEntry) validateImageID(formats strfmt.Registry) error {

	if err := validate.Required("imageId", "body", m.ImageID); err != nil {
		return err
	}

	if err := validate.MinLength("imageId", "body", string(*m.ImageID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("imageId", "body", string(*m.ImageID), 255); err != nil {
		return err
	}

	return nil
}

func (m *ImageShapeCompatibilityEntry) validateShape(formats strfmt.Registry) error {

	if err := validate.Required("shape", "body", m.Shape); err != nil {
		return err
	}

	if err := validate.MinLength("shape", "body", string(*m.Shape), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("shape", "body", string(*m.Shape), 255); err != nil {
		return err
	}

	return nil
}
