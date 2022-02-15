// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// KubeProxyReplacement Status of kube-proxy replacement
//
// +k8s:deepcopy-gen=true
//
// swagger:model KubeProxyReplacement
type KubeProxyReplacement struct {

	//
	//
	// +k8s:deepcopy-gen=true
	DeviceList []*KubeProxyReplacementDeviceListItems0 `json:"deviceList"`

	// devices
	Devices []string `json:"devices"`

	// direct routing device
	DirectRoutingDevice string `json:"directRoutingDevice,omitempty"`

	// features
	Features *KubeProxyReplacementFeatures `json:"features,omitempty"`

	// mode
	// Enum: [Disabled Strict Probe Partial]
	Mode string `json:"mode,omitempty"`
}

// Validate validates this kube proxy replacement
func (m *KubeProxyReplacement) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeviceList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFeatures(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KubeProxyReplacement) validateDeviceList(formats strfmt.Registry) error {

	if swag.IsZero(m.DeviceList) { // not required
		return nil
	}

	for i := 0; i < len(m.DeviceList); i++ {
		if swag.IsZero(m.DeviceList[i]) { // not required
			continue
		}

		if m.DeviceList[i] != nil {
			if err := m.DeviceList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("deviceList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *KubeProxyReplacement) validateFeatures(formats strfmt.Registry) error {

	if swag.IsZero(m.Features) { // not required
		return nil
	}

	if m.Features != nil {
		if err := m.Features.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("features")
			}
			return err
		}
	}

	return nil
}

var kubeProxyReplacementTypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Disabled","Strict","Probe","Partial"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		kubeProxyReplacementTypeModePropEnum = append(kubeProxyReplacementTypeModePropEnum, v)
	}
}

const (

	// KubeProxyReplacementModeDisabled captures enum value "Disabled"
	KubeProxyReplacementModeDisabled string = "Disabled"

	// KubeProxyReplacementModeStrict captures enum value "Strict"
	KubeProxyReplacementModeStrict string = "Strict"

	// KubeProxyReplacementModeProbe captures enum value "Probe"
	KubeProxyReplacementModeProbe string = "Probe"

	// KubeProxyReplacementModePartial captures enum value "Partial"
	KubeProxyReplacementModePartial string = "Partial"
)

// prop value enum
func (m *KubeProxyReplacement) validateModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, kubeProxyReplacementTypeModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *KubeProxyReplacement) validateMode(formats strfmt.Registry) error {

	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	// value enum
	if err := m.validateModeEnum("mode", "body", m.Mode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KubeProxyReplacement) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KubeProxyReplacement) UnmarshalBinary(b []byte) error {
	var res KubeProxyReplacement
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// KubeProxyReplacementDeviceListItems0
//
// +k8s:deepcopy-gen=true
//
// swagger:model KubeProxyReplacementDeviceListItems0
type KubeProxyReplacementDeviceListItems0 struct {

	//
	//
	// +k8s:deepcopy-gen=true
	IP []string `json:"ip"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this kube proxy replacement device list items0
func (m *KubeProxyReplacementDeviceListItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *KubeProxyReplacementDeviceListItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KubeProxyReplacementDeviceListItems0) UnmarshalBinary(b []byte) error {
	var res KubeProxyReplacementDeviceListItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// KubeProxyReplacementFeatures
//
// +k8s:deepcopy-gen=true
//
// swagger:model KubeProxyReplacementFeatures
type KubeProxyReplacementFeatures struct {

	// external i ps
	ExternalIPs *KubeProxyReplacementFeaturesExternalIPs `json:"externalIPs,omitempty"`

	// graceful termination
	GracefulTermination *KubeProxyReplacementFeaturesGracefulTermination `json:"gracefulTermination,omitempty"`

	// host port
	HostPort *KubeProxyReplacementFeaturesHostPort `json:"hostPort,omitempty"`

	// host reachable services
	HostReachableServices *KubeProxyReplacementFeaturesHostReachableServices `json:"hostReachableServices,omitempty"`

	// node port
	NodePort *KubeProxyReplacementFeaturesNodePort `json:"nodePort,omitempty"`

	// session affinity
	SessionAffinity *KubeProxyReplacementFeaturesSessionAffinity `json:"sessionAffinity,omitempty"`
}

// Validate validates this kube proxy replacement features
func (m *KubeProxyReplacementFeatures) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExternalIPs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGracefulTermination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostPort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostReachableServices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodePort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSessionAffinity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KubeProxyReplacementFeatures) validateExternalIPs(formats strfmt.Registry) error {

	if swag.IsZero(m.ExternalIPs) { // not required
		return nil
	}

	if m.ExternalIPs != nil {
		if err := m.ExternalIPs.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("features" + "." + "externalIPs")
			}
			return err
		}
	}

	return nil
}

func (m *KubeProxyReplacementFeatures) validateGracefulTermination(formats strfmt.Registry) error {

	if swag.IsZero(m.GracefulTermination) { // not required
		return nil
	}

	if m.GracefulTermination != nil {
		if err := m.GracefulTermination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("features" + "." + "gracefulTermination")
			}
			return err
		}
	}

	return nil
}

func (m *KubeProxyReplacementFeatures) validateHostPort(formats strfmt.Registry) error {

	if swag.IsZero(m.HostPort) { // not required
		return nil
	}

	if m.HostPort != nil {
		if err := m.HostPort.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("features" + "." + "hostPort")
			}
			return err
		}
	}

	return nil
}

func (m *KubeProxyReplacementFeatures) validateHostReachableServices(formats strfmt.Registry) error {

	if swag.IsZero(m.HostReachableServices) { // not required
		return nil
	}

	if m.HostReachableServices != nil {
		if err := m.HostReachableServices.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("features" + "." + "hostReachableServices")
			}
			return err
		}
	}

	return nil
}

func (m *KubeProxyReplacementFeatures) validateNodePort(formats strfmt.Registry) error {

	if swag.IsZero(m.NodePort) { // not required
		return nil
	}

	if m.NodePort != nil {
		if err := m.NodePort.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("features" + "." + "nodePort")
			}
			return err
		}
	}

	return nil
}

func (m *KubeProxyReplacementFeatures) validateSessionAffinity(formats strfmt.Registry) error {

	if swag.IsZero(m.SessionAffinity) { // not required
		return nil
	}

	if m.SessionAffinity != nil {
		if err := m.SessionAffinity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("features" + "." + "sessionAffinity")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KubeProxyReplacementFeatures) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KubeProxyReplacementFeatures) UnmarshalBinary(b []byte) error {
	var res KubeProxyReplacementFeatures
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// KubeProxyReplacementFeaturesExternalIPs
//
// +k8s:deepcopy-gen=true
//
// swagger:model KubeProxyReplacementFeaturesExternalIPs
type KubeProxyReplacementFeaturesExternalIPs struct {

	// enabled
	Enabled bool `json:"enabled,omitempty"`
}

// Validate validates this kube proxy replacement features external i ps
func (m *KubeProxyReplacementFeaturesExternalIPs) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *KubeProxyReplacementFeaturesExternalIPs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KubeProxyReplacementFeaturesExternalIPs) UnmarshalBinary(b []byte) error {
	var res KubeProxyReplacementFeaturesExternalIPs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// KubeProxyReplacementFeaturesGracefulTermination
//
// +k8s:deepcopy-gen=true
//
// swagger:model KubeProxyReplacementFeaturesGracefulTermination
type KubeProxyReplacementFeaturesGracefulTermination struct {

	// enabled
	Enabled bool `json:"enabled,omitempty"`
}

// Validate validates this kube proxy replacement features graceful termination
func (m *KubeProxyReplacementFeaturesGracefulTermination) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *KubeProxyReplacementFeaturesGracefulTermination) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KubeProxyReplacementFeaturesGracefulTermination) UnmarshalBinary(b []byte) error {
	var res KubeProxyReplacementFeaturesGracefulTermination
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// KubeProxyReplacementFeaturesHostPort
//
// +k8s:deepcopy-gen=true
//
// swagger:model KubeProxyReplacementFeaturesHostPort
type KubeProxyReplacementFeaturesHostPort struct {

	// enabled
	Enabled bool `json:"enabled,omitempty"`
}

// Validate validates this kube proxy replacement features host port
func (m *KubeProxyReplacementFeaturesHostPort) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *KubeProxyReplacementFeaturesHostPort) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KubeProxyReplacementFeaturesHostPort) UnmarshalBinary(b []byte) error {
	var res KubeProxyReplacementFeaturesHostPort
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// KubeProxyReplacementFeaturesHostReachableServices
//
// +k8s:deepcopy-gen=true
//
// swagger:model KubeProxyReplacementFeaturesHostReachableServices
type KubeProxyReplacementFeaturesHostReachableServices struct {

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// protocols
	Protocols []string `json:"protocols"`
}

// Validate validates this kube proxy replacement features host reachable services
func (m *KubeProxyReplacementFeaturesHostReachableServices) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *KubeProxyReplacementFeaturesHostReachableServices) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KubeProxyReplacementFeaturesHostReachableServices) UnmarshalBinary(b []byte) error {
	var res KubeProxyReplacementFeaturesHostReachableServices
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// KubeProxyReplacementFeaturesNodePort
//
// +k8s:deepcopy-gen=true
//
// swagger:model KubeProxyReplacementFeaturesNodePort
type KubeProxyReplacementFeaturesNodePort struct {

	// acceleration
	// Enum: [None Native Generic]
	Acceleration string `json:"acceleration,omitempty"`

	// algorithm
	// Enum: [Random Maglev]
	Algorithm string `json:"algorithm,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// lut size
	LutSize int64 `json:"lutSize,omitempty"`

	// mode
	// Enum: [SNAT DSR Hybrid]
	Mode string `json:"mode,omitempty"`

	// port max
	PortMax int64 `json:"portMax,omitempty"`

	// port min
	PortMin int64 `json:"portMin,omitempty"`
}

// Validate validates this kube proxy replacement features node port
func (m *KubeProxyReplacementFeaturesNodePort) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcceleration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAlgorithm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var kubeProxyReplacementFeaturesNodePortTypeAccelerationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["None","Native","Generic"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		kubeProxyReplacementFeaturesNodePortTypeAccelerationPropEnum = append(kubeProxyReplacementFeaturesNodePortTypeAccelerationPropEnum, v)
	}
}

const (

	// KubeProxyReplacementFeaturesNodePortAccelerationNone captures enum value "None"
	KubeProxyReplacementFeaturesNodePortAccelerationNone string = "None"

	// KubeProxyReplacementFeaturesNodePortAccelerationNative captures enum value "Native"
	KubeProxyReplacementFeaturesNodePortAccelerationNative string = "Native"

	// KubeProxyReplacementFeaturesNodePortAccelerationGeneric captures enum value "Generic"
	KubeProxyReplacementFeaturesNodePortAccelerationGeneric string = "Generic"
)

// prop value enum
func (m *KubeProxyReplacementFeaturesNodePort) validateAccelerationEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, kubeProxyReplacementFeaturesNodePortTypeAccelerationPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *KubeProxyReplacementFeaturesNodePort) validateAcceleration(formats strfmt.Registry) error {

	if swag.IsZero(m.Acceleration) { // not required
		return nil
	}

	// value enum
	if err := m.validateAccelerationEnum("features"+"."+"nodePort"+"."+"acceleration", "body", m.Acceleration); err != nil {
		return err
	}

	return nil
}

var kubeProxyReplacementFeaturesNodePortTypeAlgorithmPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Random","Maglev"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		kubeProxyReplacementFeaturesNodePortTypeAlgorithmPropEnum = append(kubeProxyReplacementFeaturesNodePortTypeAlgorithmPropEnum, v)
	}
}

const (

	// KubeProxyReplacementFeaturesNodePortAlgorithmRandom captures enum value "Random"
	KubeProxyReplacementFeaturesNodePortAlgorithmRandom string = "Random"

	// KubeProxyReplacementFeaturesNodePortAlgorithmMaglev captures enum value "Maglev"
	KubeProxyReplacementFeaturesNodePortAlgorithmMaglev string = "Maglev"
)

// prop value enum
func (m *KubeProxyReplacementFeaturesNodePort) validateAlgorithmEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, kubeProxyReplacementFeaturesNodePortTypeAlgorithmPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *KubeProxyReplacementFeaturesNodePort) validateAlgorithm(formats strfmt.Registry) error {

	if swag.IsZero(m.Algorithm) { // not required
		return nil
	}

	// value enum
	if err := m.validateAlgorithmEnum("features"+"."+"nodePort"+"."+"algorithm", "body", m.Algorithm); err != nil {
		return err
	}

	return nil
}

var kubeProxyReplacementFeaturesNodePortTypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SNAT","DSR","Hybrid"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		kubeProxyReplacementFeaturesNodePortTypeModePropEnum = append(kubeProxyReplacementFeaturesNodePortTypeModePropEnum, v)
	}
}

const (

	// KubeProxyReplacementFeaturesNodePortModeSNAT captures enum value "SNAT"
	KubeProxyReplacementFeaturesNodePortModeSNAT string = "SNAT"

	// KubeProxyReplacementFeaturesNodePortModeDSR captures enum value "DSR"
	KubeProxyReplacementFeaturesNodePortModeDSR string = "DSR"

	// KubeProxyReplacementFeaturesNodePortModeHybrid captures enum value "Hybrid"
	KubeProxyReplacementFeaturesNodePortModeHybrid string = "Hybrid"
)

// prop value enum
func (m *KubeProxyReplacementFeaturesNodePort) validateModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, kubeProxyReplacementFeaturesNodePortTypeModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *KubeProxyReplacementFeaturesNodePort) validateMode(formats strfmt.Registry) error {

	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	// value enum
	if err := m.validateModeEnum("features"+"."+"nodePort"+"."+"mode", "body", m.Mode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KubeProxyReplacementFeaturesNodePort) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KubeProxyReplacementFeaturesNodePort) UnmarshalBinary(b []byte) error {
	var res KubeProxyReplacementFeaturesNodePort
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// KubeProxyReplacementFeaturesSessionAffinity
//
// +k8s:deepcopy-gen=true
//
// swagger:model KubeProxyReplacementFeaturesSessionAffinity
type KubeProxyReplacementFeaturesSessionAffinity struct {

	// enabled
	Enabled bool `json:"enabled,omitempty"`
}

// Validate validates this kube proxy replacement features session affinity
func (m *KubeProxyReplacementFeaturesSessionAffinity) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *KubeProxyReplacementFeaturesSessionAffinity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KubeProxyReplacementFeaturesSessionAffinity) UnmarshalBinary(b []byte) error {
	var res KubeProxyReplacementFeaturesSessionAffinity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
