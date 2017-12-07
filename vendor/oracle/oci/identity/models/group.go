package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// Group A collection of users who all need the same type of access to a particular set of resources or compartment.
// For conceptual information about groups and other IAM Service components, see
// [Overview of the IAM Service](/Content/Identity/Concepts/overview.htm).
//
// If you're federating with an identity provider (IdP), you need to create mappings between the groups
// defined in the IdP and groups you define in the IAM service. For more information, see
// [Identity Providers and Federation](/Content/Identity/Concepts/federation.htm). Also see
// [IdentityProvider](#/en/identity/20160918/IdentityProvider/) and
// [IdpGroupMapping](#/en/identity/20160918/IdpGroupMapping/).
//
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access,
// see [Getting Started with Policies](/Content/Identity/Concepts/policygetstarted.htm).
//
// swagger:model Group
type Group struct {

	// The OCID of the tenancy containing the group.
	// Required: true
	CompartmentID *string `json:"compartmentId"`

	// The description you assign to the group. Does not have to be unique, and it's changeable.
	// Required: true
	// Max Length: 400
	// Min Length: 1
	Description *string `json:"description"`

	// The OCID of the group.
	// Required: true
	ID *string `json:"id"`

	// The detailed status of INACTIVE lifecycleState.
	InactiveStatus int64 `json:"inactiveStatus,omitempty"`

	// The group's current state. After creating a group, make sure its `lifecycleState` changes from CREATING to
	// ACTIVE before using it.
	//
	// Required: true
	// Max Length: 64
	// Min Length: 1
	LifecycleState *string `json:"lifecycleState"`

	// The name you assign to the group during creation. The name must be unique across all groups in
	// the tenancy and cannot be changed.
	//
	// Required: true
	// Max Length: 100
	// Min Length: 1
	Name *string `json:"name"`

	// Date and time the group was created, in the format defined by RFC3339.
	//
	// Example: `2016-08-25T21:10:29.600Z`
	//
	// Required: true
	TimeCreated *strfmt.DateTime `json:"timeCreated"`
}

// Validate validates this group
func (m *Group) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompartmentID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLifecycleState(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTimeCreated(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Group) validateCompartmentID(formats strfmt.Registry) error {

	if err := validate.Required("compartmentId", "body", m.CompartmentID); err != nil {
		return err
	}

	return nil
}

func (m *Group) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	if err := validate.MinLength("description", "body", string(*m.Description), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("description", "body", string(*m.Description), 400); err != nil {
		return err
	}

	return nil
}

func (m *Group) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

var groupTypeLifecycleStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CREATING","ACTIVE","INACTIVE","DELETING","DELETED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		groupTypeLifecycleStatePropEnum = append(groupTypeLifecycleStatePropEnum, v)
	}
}

const (
	// GroupLifecycleStateCREATING captures enum value "CREATING"
	GroupLifecycleStateCREATING string = "CREATING"
	// GroupLifecycleStateACTIVE captures enum value "ACTIVE"
	GroupLifecycleStateACTIVE string = "ACTIVE"
	// GroupLifecycleStateINACTIVE captures enum value "INACTIVE"
	GroupLifecycleStateINACTIVE string = "INACTIVE"
	// GroupLifecycleStateDELETING captures enum value "DELETING"
	GroupLifecycleStateDELETING string = "DELETING"
	// GroupLifecycleStateDELETED captures enum value "DELETED"
	GroupLifecycleStateDELETED string = "DELETED"
)

// prop value enum
func (m *Group) validateLifecycleStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, groupTypeLifecycleStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Group) validateLifecycleState(formats strfmt.Registry) error {

	if err := validate.Required("lifecycleState", "body", m.LifecycleState); err != nil {
		return err
	}

	if err := validate.MinLength("lifecycleState", "body", string(*m.LifecycleState), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("lifecycleState", "body", string(*m.LifecycleState), 64); err != nil {
		return err
	}

	// value enum
	if err := m.validateLifecycleStateEnum("lifecycleState", "body", *m.LifecycleState); err != nil {
		return err
	}

	return nil
}

func (m *Group) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 100); err != nil {
		return err
	}

	return nil
}

func (m *Group) validateTimeCreated(formats strfmt.Registry) error {

	if err := validate.Required("timeCreated", "body", m.TimeCreated); err != nil {
		return err
	}

	return nil
}