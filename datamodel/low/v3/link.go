// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package v3

import (
	"github.com/pb33f/libopenapi/datamodel/low"
	"github.com/pb33f/libopenapi/index"
	"gopkg.in/yaml.v3"
)

// Link represents a low-level OpenAPI 3+ Link object.
//
// The Link object represents a possible design-time link for a response. The presence of a link does not guarantee the
// caller’s ability to successfully invoke it, rather it provides a known relationship and traversal mechanism between
// responses and other operations.
//
// Unlike dynamic links (i.e. links provided in the response payload), the OAS linking mechanism does not require
// link information in the runtime response.
//
// For computing links, and providing instructions to execute them, a runtime expression is used for accessing values
// in an operation and using them as parameters while invoking the linked operation.
//  - https://spec.openapis.org/oas/v3.1.0#link-object
type Link struct {
	OperationRef low.NodeReference[string]
	OperationId  low.NodeReference[string]
	Parameters   low.KeyReference[map[low.KeyReference[string]]low.ValueReference[string]]
	RequestBody  low.NodeReference[string]
	Description  low.NodeReference[string]
	Server       low.NodeReference[*Server]
	Extensions   map[low.KeyReference[string]]low.ValueReference[any]
}

// FindParameter will attempt to locate a parameter string value, using a parameter name input.
func (l *Link) FindParameter(pName string) *low.ValueReference[string] {
	return low.FindItemInMap[string](pName, l.Parameters.Value)
}

// FindExtension will attempt to locate an extension with a specific key
func (l *Link) FindExtension(ext string) *low.ValueReference[any] {
	return low.FindItemInMap[any](ext, l.Extensions)
}

// Build will extract extensions and servers from the node.
func (l *Link) Build(root *yaml.Node, idx *index.SpecIndex) error {
	l.Extensions = low.ExtractExtensions(root)

	// extract server.
	ser, sErr := low.ExtractObject[*Server](ServerLabel, root, idx)
	if sErr != nil {
		return sErr
	}
	l.Server = ser
	return nil
}
