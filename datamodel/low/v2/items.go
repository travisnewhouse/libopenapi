// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package v2

import (
	"github.com/pb33f/libopenapi/datamodel/low"
	"github.com/pb33f/libopenapi/index"
	"github.com/pb33f/libopenapi/utils"
	"gopkg.in/yaml.v3"
)

// Items is a low-level representation of a Swagger / OpenAPI 2 Items object.
//
// Items is a limited subset of JSON-Schema's items object. It is used by parameter definitions that are not
// located in "body"
//  - https://swagger.io/specification/v2/#itemsObject
type Items struct {
	Type             low.NodeReference[string]
	Format           low.NodeReference[string]
	CollectionFormat low.NodeReference[string]
	Items            low.NodeReference[*Items]
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
}

// Build will build out items and default value.
func (i *Items) Build(root *yaml.Node, idx *index.SpecIndex) error {
	items, iErr := low.ExtractObject[*Items](ItemsLabel, root, idx)
	if iErr != nil {
		return iErr
	}
	i.Items = items

	_, ln, vn := utils.FindKeyNodeFull(DefaultLabel, root.Content)
	if vn != nil {
		var n map[string]interface{}
		err := vn.Decode(&n)
		if err != nil {
			// if not a map, then try an array
			var k []interface{}
			err = vn.Decode(&k)
			if err != nil {
				// lets just default to interface
				var j interface{}
				_ = vn.Decode(&j)
				i.Default = low.NodeReference[any]{
					Value:     j,
					KeyNode:   ln,
					ValueNode: vn,
				}
				return nil
			}
			i.Default = low.NodeReference[any]{
				Value:     k,
				KeyNode:   ln,
				ValueNode: vn,
			}
			return nil
		}
		i.Default = low.NodeReference[any]{
			Value:     n,
			KeyNode:   ln,
			ValueNode: vn,
		}
		return nil
	}
	return nil
}
