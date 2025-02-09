// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package v2

import (
	"github.com/pb33f/libopenapi/datamodel/low"
	"github.com/pb33f/libopenapi/index"
	"github.com/pb33f/libopenapi/utils"
	"gopkg.in/yaml.v3"
)

// Header Represents a low-level Swagger / OpenAPI 2 Header object.
//
// A Header is essentially identical to a Parameter, except it does not contain 'name' or 'in' properties.
//  - https://swagger.io/specification/v2/#headerObject
type Header struct {
	Type             low.NodeReference[string]
	Format           low.NodeReference[string]
	Description      low.NodeReference[string]
	Items            low.NodeReference[*Items]
	CollectionFormat low.NodeReference[string]
	Default          low.NodeReference[any]
	Maximum          low.NodeReference[int]
	ExclusiveMaximum low.NodeReference[bool]
	Minimum          low.NodeReference[int]
	ExclusiveMinimum low.NodeReference[bool]
	MaxLength        low.NodeReference[int]
	MinLength        low.NodeReference[int]
	Pattern          low.NodeReference[string]
	MaxItems         low.NodeReference[int]
	MinItems         low.NodeReference[int]
	UniqueItems      low.NodeReference[bool]
	Enum             low.NodeReference[[]low.ValueReference[string]]
	MultipleOf       low.NodeReference[int]
	Extensions       map[low.KeyReference[string]]low.ValueReference[any]
}

// FindExtension will attempt to locate an extension value using a name lookup.
func (h *Header) FindExtension(ext string) *low.ValueReference[any] {
	return low.FindItemInMap[any](ext, h.Extensions)
}

// Build will build out items, extensions and default value from the supplied node.
func (h *Header) Build(root *yaml.Node, idx *index.SpecIndex) error {
	h.Extensions = low.ExtractExtensions(root)
	items, err := low.ExtractObject[*Items](ItemsLabel, root, idx)
	if err != nil {
		return err
	}
	h.Items = items

	_, ln, vn := utils.FindKeyNodeFull(DefaultLabel, root.Content)
	if vn != nil {
		var n map[string]interface{}
		err = vn.Decode(&n)
		if err != nil {
			// if not a map, then try an array
			var k []interface{}
			err = vn.Decode(&k)
			if err != nil {
				// lets just default to interface
				var j interface{}
				_ = vn.Decode(&j)
				h.Default = low.NodeReference[any]{
					Value:     j,
					KeyNode:   ln,
					ValueNode: vn,
				}
				return nil
			}
			h.Default = low.NodeReference[any]{
				Value:     k,
				KeyNode:   ln,
				ValueNode: vn,
			}
			return nil
		}
		h.Default = low.NodeReference[any]{
			Value:     n,
			KeyNode:   ln,
			ValueNode: vn,
		}
		return nil
	}
	return nil
}
