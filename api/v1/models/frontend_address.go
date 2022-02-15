// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FrontendAddress Layer 4 address. The protocol is currently ignored, all services will
// behave as if protocol any is specified. To restrict to a particular
// protocol, use policy.
//
//
// swagger:model FrontendAddress
type FrontendAddress struct {

	// Layer 3 address
	IP string `json:"ip,omitempty"`

	// Layer 4 port number
	Port uint16 `json:"port,omitempty"`

	// Layer 4 protocol
	// Enum: [tcp udp any]
	Protocol string `json:"protocol,omitempty"`

	// Load balancing scope for frontend address
	// Enum: [external internal]
	Scope string `json:"scope,omitempty"`
}

// Validate validates this frontend address
func (m *FrontendAddress) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var frontendAddressTypeProtocolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["tcp","udp","any"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		frontendAddressTypeProtocolPropEnum = append(frontendAddressTypeProtocolPropEnum, v)
	}
}

const (

	// FrontendAddressProtocolTCP captures enum value "tcp"
	FrontendAddressProtocolTCP string = "tcp"

	// FrontendAddressProtocolUDP captures enum value "udp"
	FrontendAddressProtocolUDP string = "udp"

	// FrontendAddressProtocolAny captures enum value "any"
	FrontendAddressProtocolAny string = "any"
)

// prop value enum
func (m *FrontendAddress) validateProtocolEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, frontendAddressTypeProtocolPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FrontendAddress) validateProtocol(formats strfmt.Registry) error {

	if swag.IsZero(m.Protocol) { // not required
		return nil
	}

	// value enum
	if err := m.validateProtocolEnum("protocol", "body", m.Protocol); err != nil {
		return err
	}

	return nil
}

var frontendAddressTypeScopePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["external","internal"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		frontendAddressTypeScopePropEnum = append(frontendAddressTypeScopePropEnum, v)
	}
}

const (

	// FrontendAddressScopeExternal captures enum value "external"
	FrontendAddressScopeExternal string = "external"

	// FrontendAddressScopeInternal captures enum value "internal"
	FrontendAddressScopeInternal string = "internal"
)

// prop value enum
func (m *FrontendAddress) validateScopeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, frontendAddressTypeScopePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FrontendAddress) validateScope(formats strfmt.Registry) error {

	if swag.IsZero(m.Scope) { // not required
		return nil
	}

	// value enum
	if err := m.validateScopeEnum("scope", "body", m.Scope); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FrontendAddress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FrontendAddress) UnmarshalBinary(b []byte) error {
	var res FrontendAddress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
