// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NodeStatus Connectivity status of a remote cilium-health instance
//
// swagger:model NodeStatus
type NodeStatus struct {

	// DEPRECATED: Please use the health-endpoint field instead, which
	// supports reporting the status of different addresses for the endpoint
	//
	Endpoint *PathStatus `json:"endpoint,omitempty"`

	// Connectivity status to simulated endpoint on the node
	HealthEndpoint *EndpointStatus `json:"health-endpoint,omitempty"`

	// Connectivity status to cilium-health instance on node IP
	Host *HostStatus `json:"host,omitempty"`

	// Identifying name for the node
	Name string `json:"name,omitempty"`
}

// Validate validates this node status
func (m *NodeStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndpoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHealthEndpoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHost(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NodeStatus) validateEndpoint(formats strfmt.Registry) error {

	if swag.IsZero(m.Endpoint) { // not required
		return nil
	}

	if m.Endpoint != nil {
		if err := m.Endpoint.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("endpoint")
			}
			return err
		}
	}

	return nil
}

func (m *NodeStatus) validateHealthEndpoint(formats strfmt.Registry) error {

	if swag.IsZero(m.HealthEndpoint) { // not required
		return nil
	}

	if m.HealthEndpoint != nil {
		if err := m.HealthEndpoint.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("health-endpoint")
			}
			return err
		}
	}

	return nil
}

func (m *NodeStatus) validateHost(formats strfmt.Registry) error {

	if swag.IsZero(m.Host) { // not required
		return nil
	}

	if m.Host != nil {
		if err := m.Host.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("host")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NodeStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NodeStatus) UnmarshalBinary(b []byte) error {
	var res NodeStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
