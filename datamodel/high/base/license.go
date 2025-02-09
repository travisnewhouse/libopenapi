// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package base

import (
	low "github.com/pb33f/libopenapi/datamodel/low/base"
)

// License is a high-level representation of a License object as defined by OpenAPI 2 and OpenAPI 3
//  v2 - https://swagger.io/specification/v2/#licenseObject
//  v3 - https://spec.openapis.org/oas/v3.1.0#license-object
type License struct {
	Name string
	URL  string
	low  *low.License
}

// NewLicense will create a new high-level License instance from a low-level one.
func NewLicense(license *low.License) *License {
	l := new(License)
	l.low = license
	if !license.URL.IsEmpty() {
		l.URL = license.URL.Value
	}
	if !license.Name.IsEmpty() {
		l.Name = license.Name.Value
	}
	return l
}

// GoLow will return the low-level License used to create the high-level one.
func (l *License) GoLow() *low.License {
	return l.low
}
