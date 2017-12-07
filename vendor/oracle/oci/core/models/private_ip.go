package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PrivateIP A *private IP* is a conceptual term that refers to a private IP address and related properties.
// The `privateIp` object is the API representation of a private IP.
//
// Each instance has a *primary private IP* that is automatically created and
// assigned to the primary VNIC during instance launch. If you add a secondary
// VNIC to the instance, it also automatically gets a primary private IP. You
// can't remove a primary private IP from its VNIC. The primary private IP is
// automatically deleted when the VNIC is terminated.
//
// You can add *secondary private IPs* to a VNIC after it's created. For more
// information, see the `privateIp` operations and also
// [Managing IP Addresses](/Content/Network/Tasks/managingIPaddresses.htm).
//
// **Note:** Only
// [ListPrivateIps](#/en/iaas/20160918/PrivateIp/ListPrivateIps) and
// [GetPrivateIp](#/en/iaas/20160918/PrivateIp/GetPrivateIp) work with
// *primary* private IPs. To create and update primary private IPs, you instead
// work with instance and VNIC operations. For example, a primary private IP's
// properties come from the values you specify in
// [CreateVnicDetails](#/en/iaas/20160918/CreateVnicDetails/) when calling either
// [LaunchInstance](#/en/iaas/20160918/Instance/LaunchInstance) or
// [AttachVnic](#/en/iaas/20160918/VnicAttachment/AttachVnic). To update the hostname
// for a primary private IP, you use [UpdateVnic](#/en/iaas/20160918/Vnic/UpdateVnic).
//
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// [Getting Started with Policies](/Content/Identity/Concepts/policygetstarted.htm).
//
// swagger:model PrivateIp
type PrivateIP struct {

	// The private IP's Availability Domain.
	//
	// Example: `Uocm:PHX-AD-1`
	//
	// Max Length: 255
	// Min Length: 1
	AvailabilityDomain string `json:"availabilityDomain,omitempty"`

	// The OCID of the compartment containing the private IP.
	// Max Length: 255
	// Min Length: 1
	CompartmentID string `json:"compartmentId,omitempty"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid
	// entering confidential information.
	//
	// Max Length: 255
	// Min Length: 1
	DisplayName string `json:"displayName,omitempty"`

	// The hostname for the private IP. Used for DNS. The value is the hostname
	// portion of the private IP's fully qualified domain name (FQDN)
	// (for example, `bminstance-1` in FQDN `bminstance-1.subnet123.vcn1.oraclevcn.com`).
	// Must be unique across all VNICs in the subnet and comply with
	// [RFC 952](https://tools.ietf.org/html/rfc952) and
	// [RFC 1123](https://tools.ietf.org/html/rfc1123).
	//
	// For more information, see
	// [DNS in Your Virtual Cloud Network](/Content/Network/Concepts/dns.htm).
	//
	// Example: `bminstance-1`
	//
	// Max Length: 63
	// Min Length: 1
	HostnameLabel string `json:"hostnameLabel,omitempty"`

	// The private IP's Oracle ID (OCID).
	// Max Length: 255
	// Min Length: 1
	ID string `json:"id,omitempty"`

	// The private IP address of the `privateIp` object. The address is within the CIDR
	// of the VNIC's subnet.
	//
	// Example: `10.0.3.3`
	//
	IPAddress string `json:"ipAddress,omitempty"`

	// Whether this private IP is the primary one on the VNIC. Primary private IPs
	// are unassigned and deleted automatically when the VNIC is terminated.
	//
	// Example: `true`
	//
	IsPrimary bool `json:"isPrimary,omitempty"`

	// The OCID of the subnet the VNIC is in.
	// Max Length: 255
	// Min Length: 1
	SubnetID string `json:"subnetId,omitempty"`

	// The date and time the private IP was created, in the format defined by RFC3339.
	//
	// Example: `2016-08-25T21:10:29.600Z`
	//
	TimeCreated strfmt.DateTime `json:"timeCreated,omitempty"`

	// The OCID of the VNIC the private IP is assigned to. The VNIC and private IP
	// must be in the same subnet.
	//
	// Max Length: 255
	// Min Length: 1
	VnicID string `json:"vnicId,omitempty"`
}

// Validate validates this private Ip
func (m *PrivateIP) Validate(formats strfmt.Registry) error {
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

	if err := m.validateHostnameLabel(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSubnetID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVnicID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PrivateIP) validateAvailabilityDomain(formats strfmt.Registry) error {

	if swag.IsZero(m.AvailabilityDomain) { // not required
		return nil
	}

	if err := validate.MinLength("availabilityDomain", "body", string(m.AvailabilityDomain), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("availabilityDomain", "body", string(m.AvailabilityDomain), 255); err != nil {
		return err
	}

	return nil
}

func (m *PrivateIP) validateCompartmentID(formats strfmt.Registry) error {

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

func (m *PrivateIP) validateDisplayName(formats strfmt.Registry) error {

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

func (m *PrivateIP) validateHostnameLabel(formats strfmt.Registry) error {

	if swag.IsZero(m.HostnameLabel) { // not required
		return nil
	}

	if err := validate.MinLength("hostnameLabel", "body", string(m.HostnameLabel), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("hostnameLabel", "body", string(m.HostnameLabel), 63); err != nil {
		return err
	}

	return nil
}

func (m *PrivateIP) validateID(formats strfmt.Registry) error {

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

func (m *PrivateIP) validateSubnetID(formats strfmt.Registry) error {

	if swag.IsZero(m.SubnetID) { // not required
		return nil
	}

	if err := validate.MinLength("subnetId", "body", string(m.SubnetID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("subnetId", "body", string(m.SubnetID), 255); err != nil {
		return err
	}

	return nil
}

func (m *PrivateIP) validateVnicID(formats strfmt.Registry) error {

	if swag.IsZero(m.VnicID) { // not required
		return nil
	}

	if err := validate.MinLength("vnicId", "body", string(m.VnicID), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("vnicId", "body", string(m.VnicID), 255); err != nil {
		return err
	}

	return nil
}