package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// AddUserToGroupDetails add user to group details
// swagger:model AddUserToGroupDetails
type AddUserToGroupDetails struct {

	// The OCID of the group.
	// Required: true
	GroupID *string `json:"groupId"`

	// The OCID of the user.
	// Required: true
	UserID *string `json:"userId"`
}

// Validate validates this add user to group details
func (m *AddUserToGroupDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroupID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AddUserToGroupDetails) validateGroupID(formats strfmt.Registry) error {

	if err := validate.Required("groupId", "body", m.GroupID); err != nil {
		return err
	}

	return nil
}

func (m *AddUserToGroupDetails) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("userId", "body", m.UserID); err != nil {
		return err
	}

	return nil
}
