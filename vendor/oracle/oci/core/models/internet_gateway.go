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

// InternetGateway Represents a router that connects the edge of a VCN with the Internet. For an example scenario
// that uses an Internet Gateway, see
// [Typical Networking Service Scenarios](/Content/Network/Concepts/overview.htm#three).
//
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// [Getting Started with Policies](/Content/Identity/Concepts/policygetstarted.htm).
//
// swagger:model InternetGateway
type InternetGateway struct {

	// The OCID of the compartment containing the Internet Gateway.
	// Required: true
	// Max Length: 255
	// Min Length: 1
	CompartmentID *string `json:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	//
	// Max Length: 255
	// Min Length: 1
	DisplayName string `json:"displayName,omitempty"`

	// The Internet Gateway's Oracle ID (OCID).
	// Required: true
	// Max Length: 255
	// Min Length: 1
	ID *string `json:"id"`

	// Whether the gateway is enabled. When the gateway is disabled, traffic is not
	// routed to/from the Internet, regardless of route rules.
	//
	IsEnabled bool `json:"isEnabled,omitempty"`

	// The Internet Gateway's current state.
	// Required: true
	LifecycleState *string `json:"lifecycleState"`

	// The date and time the Internet Gateway was created, in the format defined by RFC3339.
	//
	// Example: `2016-08-25T21:10:29.600Z`
	//
	TimeCreated strfmt.DateTime `json:"timeCreated,omitempty"`

	// The OCID of the VCN the Internet Gateway belongs to.
	// Required: true
	// Max Length: 255
	// Min Length: 1
	VcnID *string `json:"vcnId"`
}

// Validate validates this internet gateway
func (m *InternetGateway) Validate(formats strfmt.Registry) error {
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

	if err := m.validateVcnID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InternetGateway) validateCompartmentID(formats strfmt.Registry) error {

	if err := validate.Required("compartmentId", "body", m.CompartmentID); err != nil {
		return err
	}

	if err := validate.MinLength("compartmentId", "body", string(*m.CompartmentID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("compartmentId", "body", string(*m.CompartmentID), 255); err != nil {
		return err
	}

	return nil
}

func (m *InternetGateway) validateDisplayName(formats strfmt.Registry) error {

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

func (m *InternetGateway) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.MinLength("id", "body", string(*m.ID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("id", "body", string(*m.ID), 255); err != nil {
		return err
	}

	return nil
}

var internetGatewayTypeLifecycleStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PROVISIONING","AVAILABLE","TERMINATING","TERMINATED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		internetGatewayTypeLifecycleStatePropEnum = append(internetGatewayTypeLifecycleStatePropEnum, v)
	}
}

const (
	// InternetGatewayLifecycleStatePROVISIONING captures enum value "PROVISIONING"
	InternetGatewayLifecycleStatePROVISIONING string = "PROVISIONING"
	// InternetGatewayLifecycleStateAVAILABLE captures enum value "AVAILABLE"
	InternetGatewayLifecycleStateAVAILABLE string = "AVAILABLE"
	// InternetGatewayLifecycleStateTERMINATING captures enum value "TERMINATING"
	InternetGatewayLifecycleStateTERMINATING string = "TERMINATING"
	// InternetGatewayLifecycleStateTERMINATED captures enum value "TERMINATED"
	InternetGatewayLifecycleStateTERMINATED string = "TERMINATED"
)

// prop value enum
func (m *InternetGateway) validateLifecycleStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, internetGatewayTypeLifecycleStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *InternetGateway) validateLifecycleState(formats strfmt.Registry) error {

	if err := validate.Required("lifecycleState", "body", m.LifecycleState); err != nil {
		return err
	}

	// value enum
	if err := m.validateLifecycleStateEnum("lifecycleState", "body", *m.LifecycleState); err != nil {
		return err
	}

	return nil
}

func (m *InternetGateway) validateVcnID(formats strfmt.Registry) error {

	if err := validate.Required("vcnId", "body", m.VcnID); err != nil {
		return err
	}

	if err := validate.MinLength("vcnId", "body", string(*m.VcnID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("vcnId", "body", string(*m.VcnID), 255); err != nil {
		return err
	}

	return nil
}
