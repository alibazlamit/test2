package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CrossConnectGroup For use with Oracle Bare Metal Cloud Services FastConnect. A cross-connect group
// is a link aggregation group (LAG), which can contain one or more
// [CrossConnects](#/en/iaas/20160918/CrossConnect). Customers who are colocated with
// Oracle in a FastConnect location create and use cross-connect groups. For more
// information, see [FastConnect Overview](/Content/Network/Concepts/fastconnect.htm).
//
// **Note:** If you're a provider who is setting up a physical connection to Oracle so customers
// can use FastConnect over the connection, be aware that your connection is modeled the
// same way as a colocated customer's (with `CrossConnect` and `CrossConnectGroup` objects, etc.).
//
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// [Getting Started with Policies](/Content/Identity/Concepts/policygetstarted.htm).
//
// swagger:model CrossConnectGroup
type CrossConnectGroup struct {

	// The OCID of the compartment containing the cross-connect group.
	// Max Length: 255
	// Min Length: 1
	CompartmentID string `json:"compartmentId,omitempty"`

	// The display name of A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	//
	// Max Length: 255
	// Min Length: 1
	DisplayName string `json:"displayName,omitempty"`

	// The cross-connect group's Oracle ID (OCID).
	// Max Length: 255
	// Min Length: 1
	ID string `json:"id,omitempty"`

	// The cross-connect group's current state.
	LifecycleState string `json:"lifecycleState,omitempty"`

	// The date and time the cross-connect group was created, in the format defined by RFC3339.
	//
	// Example: `2016-08-25T21:10:29.600Z`
	//
	TimeCreated strfmt.DateTime `json:"timeCreated,omitempty"`
}

// Validate validates this cross connect group
func (m *CrossConnectGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompartmentID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDisplayName(formats); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CrossConnectGroup) validateCompartmentID(formats strfmt.Registry) error {

	if swag.IsZero(m.CompartmentID) { // not required
		return nil
	}

	if err := validate.MinLength("compartmentId", "body", string(m.CompartmentID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("compartmentId", "body", string(m.CompartmentID), 255); err != nil {
		return err
	}

	return nil
}

func (m *CrossConnectGroup) validateDisplayName(formats strfmt.Registry) error {

	if swag.IsZero(m.DisplayName) { // not required
		return nil
	}

	if err := validate.MinLength("displayName", "body", string(m.DisplayName), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("displayName", "body", string(m.DisplayName), 255); err != nil {
		return err
	}

	return nil
}

func (m *CrossConnectGroup) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.MinLength("id", "body", string(m.ID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("id", "body", string(m.ID), 255); err != nil {
		return err
	}

	return nil
}

var crossConnectGroupTypeLifecycleStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PROVISIONING","PROVISIONED","INACTIVE","TERMINATING","TERMINATED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		crossConnectGroupTypeLifecycleStatePropEnum = append(crossConnectGroupTypeLifecycleStatePropEnum, v)
	}
}

const (
	// CrossConnectGroupLifecycleStatePROVISIONING captures enum value "PROVISIONING"
	CrossConnectGroupLifecycleStatePROVISIONING string = "PROVISIONING"
	// CrossConnectGroupLifecycleStatePROVISIONED captures enum value "PROVISIONED"
	CrossConnectGroupLifecycleStatePROVISIONED string = "PROVISIONED"
	// CrossConnectGroupLifecycleStateINACTIVE captures enum value "INACTIVE"
	CrossConnectGroupLifecycleStateINACTIVE string = "INACTIVE"
	// CrossConnectGroupLifecycleStateTERMINATING captures enum value "TERMINATING"
	CrossConnectGroupLifecycleStateTERMINATING string = "TERMINATING"
	// CrossConnectGroupLifecycleStateTERMINATED captures enum value "TERMINATED"
	CrossConnectGroupLifecycleStateTERMINATED string = "TERMINATED"
)

// prop value enum
func (m *CrossConnectGroup) validateLifecycleStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, crossConnectGroupTypeLifecycleStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CrossConnectGroup) validateLifecycleState(formats strfmt.Registry) error {

	if swag.IsZero(m.LifecycleState) { // not required
		return nil
	}

	// value enum
	if err := m.validateLifecycleStateEnum("lifecycleState", "body", m.LifecycleState); err != nil {
		return err
	}

	return nil
}
