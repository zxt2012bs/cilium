// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
)

// NewGetPolicySelectorsParams creates a new GetPolicySelectorsParams object
// no default values defined in spec.
func NewGetPolicySelectorsParams() GetPolicySelectorsParams {

	return GetPolicySelectorsParams{}
}

// GetPolicySelectorsParams contains all the bound params for the get policy selectors operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetPolicySelectors
type GetPolicySelectorsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetPolicySelectorsParams() beforehand.
func (o *GetPolicySelectorsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
