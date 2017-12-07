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

// Instance A compute host. The image used to launch the instance determines its operating system and other
// software. The shape specified during the launch process determines the number of CPUs and memory
// allocated to the instance. For more information, see
// [Overview of the Compute Service](/Content/Compute/Concepts/computeoverview.htm).
//
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// [Getting Started with Policies](/Content/Identity/Concepts/policygetstarted.htm).
//
// swagger:model Instance
type Instance struct {

	// The Availability Domain the instance is running in.
	//
	// Example: `Uocm:PHX-AD-1`
	//
	// Required: true
	// Max Length: 255
	// Min Length: 1
	AvailabilityDomain *string `json:"availabilityDomain"`

	// The OCID of the compartment that contains the instance.
	// Required: true
	// Max Length: 255
	// Min Length: 1
	CompartmentID *string `json:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	//
	// Example: `My bare metal instance`
	//
	// Max Length: 255
	// Min Length: 1
	DisplayName string `json:"displayName,omitempty"`

	// Additional metadata key/value pairs that you provide.  They serve a similar purpose and functionality from fields in the 'metadata' object.
	//
	// They are distinguished from 'metadata' fields in that these can be nested JSON objects (whereas 'metadata' fields are string/string maps only).
	//
	// If you don't need nested metadata values, it is strongly advised to avoid using this object and use the Metadata object instead.
	//
	ExtendedMetadata map[string]interface{} `json:"extendedMetadata,omitempty"`

	// The OCID of the instance.
	// Required: true
	// Max Length: 255
	// Min Length: 1
	ID *string `json:"id"`

	// The image used to boot the instance. You can enumerate all available images by calling
	// [ListImages](#/en/iaas/20160918/Image/ListImages).
	//
	// Max Length: 255
	// Min Length: 1
	ImageID string `json:"imageId,omitempty"`

	// When an Oracle Bare Metal Cloud Services or virtual machine
	// instance boots, the iPXE firmware that runs on the instance is
	// configured to run an iPXE script to continue the boot process.
	//
	// If you want more control over the boot process, you can provide
	// your own custom iPXE script that will run when the instance boots;
	// however, you should be aware that the same iPXE script will run
	// every time an instance boots; not only after the initial
	// LaunchInstance call.
	//
	// The default iPXE script connects to the instance's local boot
	// volume over iSCSI and performs a network boot. If you use a custom iPXE
	// script and want to network-boot from the instance's local boot volume
	// over iSCSI the same way as the default iPXE script, you should use the
	// following iSCSI IP address: 169.254.0.2, and boot volume IQN:
	// iqn.2015-02.oracle.boot.
	//
	// For more information about the Bring Your Own Image feature of
	// Oracle Bare Metal Cloud Services, see
	// [Bring Your Own Image](/Content/Compute/References/bringyourownimage.htm).
	//
	// For more information about iPXE, see http://ipxe.org.
	//
	// Max Length: 10240
	// Min Length: 1
	IPXEScript string `json:"ipxeScript,omitempty"`

	// The current state of the instance.
	// Required: true
	LifecycleState *string `json:"lifecycleState"`

	// Custom metadata that you provide.
	Metadata map[string]string `json:"metadata,omitempty"`

	// The region that contains the Availability Domain the instance is running in.
	//
	// Example: `phx`
	//
	// Required: true
	// Max Length: 255
	// Min Length: 1
	Region *string `json:"region"`

	// The shape of the instance. The shape determines the number of CPUs and the amount of memory
	// allocated to the instance. You can enumerate all available shapes by calling
	// [ListShapes](#/en/iaas/20160918/Shape/ListShapes).
	//
	// Required: true
	// Max Length: 255
	// Min Length: 1
	Shape *string `json:"shape"`

	// The date and time the instance was created, in the format defined by RFC3339.
	//
	// Example: `2016-08-25T21:10:29.600Z`
	//
	// Required: true
	TimeCreated *strfmt.DateTime `json:"timeCreated"`
}

// Validate validates this instance
func (m *Instance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailabilityDomain(formats); err != nil {
		// prop
		res = append(res, err)
	}

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

	if err := m.validateImageID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIPXEScript(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLifecycleState(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRegion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateShape(formats); err != nil {
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

func (m *Instance) validateAvailabilityDomain(formats strfmt.Registry) error {

	if err := validate.Required("availabilityDomain", "body", m.AvailabilityDomain); err != nil {
		return err
	}

	if err := validate.MinLength("availabilityDomain", "body", string(*m.AvailabilityDomain), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("availabilityDomain", "body", string(*m.AvailabilityDomain), 255); err != nil {
		return err
	}

	return nil
}

func (m *Instance) validateCompartmentID(formats strfmt.Registry) error {

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

func (m *Instance) validateDisplayName(formats strfmt.Registry) error {

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

func (m *Instance) validateID(formats strfmt.Registry) error {

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

func (m *Instance) validateImageID(formats strfmt.Registry) error {

	if swag.IsZero(m.ImageID) { // not required
		return nil
	}

	if err := validate.MinLength("imageId", "body", string(m.ImageID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("imageId", "body", string(m.ImageID), 255); err != nil {
		return err
	}

	return nil
}

func (m *Instance) validateIPXEScript(formats strfmt.Registry) error {

	if swag.IsZero(m.IPXEScript) { // not required
		return nil
	}

	if err := validate.MinLength("ipxeScript", "body", string(m.IPXEScript), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("ipxeScript", "body", string(m.IPXEScript), 10240); err != nil {
		return err
	}

	return nil
}

var instanceTypeLifecycleStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PROVISIONING","RUNNING","STARTING","STOPPING","STOPPED","CREATING_IMAGE","TERMINATING","TERMINATED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		instanceTypeLifecycleStatePropEnum = append(instanceTypeLifecycleStatePropEnum, v)
	}
}

const (
	// InstanceLifecycleStatePROVISIONING captures enum value "PROVISIONING"
	InstanceLifecycleStatePROVISIONING string = "PROVISIONING"
	// InstanceLifecycleStateRUNNING captures enum value "RUNNING"
	InstanceLifecycleStateRUNNING string = "RUNNING"
	// InstanceLifecycleStateSTARTING captures enum value "STARTING"
	InstanceLifecycleStateSTARTING string = "STARTING"
	// InstanceLifecycleStateSTOPPING captures enum value "STOPPING"
	InstanceLifecycleStateSTOPPING string = "STOPPING"
	// InstanceLifecycleStateSTOPPED captures enum value "STOPPED"
	InstanceLifecycleStateSTOPPED string = "STOPPED"
	// InstanceLifecycleStateCREATINGIMAGE captures enum value "CREATING_IMAGE"
	InstanceLifecycleStateCREATINGIMAGE string = "CREATING_IMAGE"
	// InstanceLifecycleStateTERMINATING captures enum value "TERMINATING"
	InstanceLifecycleStateTERMINATING string = "TERMINATING"
	// InstanceLifecycleStateTERMINATED captures enum value "TERMINATED"
	InstanceLifecycleStateTERMINATED string = "TERMINATED"
)

// prop value enum
func (m *Instance) validateLifecycleStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, instanceTypeLifecycleStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Instance) validateLifecycleState(formats strfmt.Registry) error {

	if err := validate.Required("lifecycleState", "body", m.LifecycleState); err != nil {
		return err
	}

	// value enum
	if err := m.validateLifecycleStateEnum("lifecycleState", "body", *m.LifecycleState); err != nil {
		return err
	}

	return nil
}

func (m *Instance) validateRegion(formats strfmt.Registry) error {

	if err := validate.Required("region", "body", m.Region); err != nil {
		return err
	}

	if err := validate.MinLength("region", "body", string(*m.Region), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("region", "body", string(*m.Region), 255); err != nil {
		return err
	}

	return nil
}

func (m *Instance) validateShape(formats strfmt.Registry) error {

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

func (m *Instance) validateTimeCreated(formats strfmt.Registry) error {

	if err := validate.Required("timeCreated", "body", m.TimeCreated); err != nil {
		return err
	}

	return nil
}
