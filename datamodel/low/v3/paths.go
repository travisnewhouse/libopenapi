// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package v3

import (
	"fmt"
	"github.com/pb33f/libopenapi/datamodel/low"
	"github.com/pb33f/libopenapi/index"
	"github.com/pb33f/libopenapi/utils"
	"gopkg.in/yaml.v3"
	"strings"
)

// Paths represents a high-level OpenAPI 3+ Paths object, that is backed by a low-level one.
//
// Holds the relative paths to the individual endpoints and their operations. The path is appended to the URL from the
// Server Object in order to construct the full URL. The Paths MAY be empty, due to Access Control List (ACL)
// constraints.
//  - https://spec.openapis.org/oas/v3.1.0#paths-object
type Paths struct {
	PathItems  map[low.KeyReference[string]]low.ValueReference[*PathItem]
	Extensions map[low.KeyReference[string]]low.ValueReference[any]
}

// FindPath will attempt to locate a PathItem using the provided path string.
func (p *Paths) FindPath(path string) *low.ValueReference[*PathItem] {
	for k, j := range p.PathItems {
		if k.Value == path {
			return &j
		}
	}
	return nil
}

// FindExtension will attempt to locate an extension using the specified string.
func (p *Paths) FindExtension(ext string) *low.ValueReference[any] {
	return low.FindItemInMap[any](ext, p.Extensions)
}

// Build will extract extensions and all PathItems. This happens asynchronously for speed.
func (p *Paths) Build(root *yaml.Node, idx *index.SpecIndex) error {
	p.Extensions = low.ExtractExtensions(root)
	skip := false
	var currentNode *yaml.Node

	pathsMap := make(map[low.KeyReference[string]]low.ValueReference[*PathItem])

	// build each new path, in a new thread.
	type pathBuildResult struct {
		k low.KeyReference[string]
		v low.ValueReference[*PathItem]
	}

	bChan := make(chan pathBuildResult)
	eChan := make(chan error)
	var buildPathItem = func(cNode, pNode *yaml.Node, b chan<- pathBuildResult, e chan<- error) {
		if ok, _, _ := utils.IsNodeRefValue(pNode); ok {
			r, err := low.LocateRefNode(pNode, idx)
			if r != nil {
				pNode = r
				if err != nil {
					if !idx.AllowCircularReferenceResolving() {
						e <- fmt.Errorf("path item build failed: %s", err.Error())
						return
					}
				}
			} else {
				e <- fmt.Errorf("path item build failed: cannot find reference: %s at line %d, col %d",
					pNode.Content[1].Value, pNode.Content[1].Line, pNode.Content[1].Column)
				return
			}
		}

		path := new(PathItem)
		_ = low.BuildModel(pNode, path)
		err := path.Build(pNode, idx)
		if err != nil {
			e <- err
			return
		}
		b <- pathBuildResult{
			k: low.KeyReference[string]{
				Value:   cNode.Value,
				KeyNode: cNode,
			},
			v: low.ValueReference[*PathItem]{
				Value:     path,
				ValueNode: pNode,
			},
		}
	}

	pathCount := 0
	for i, pathNode := range root.Content {
		if strings.HasPrefix(strings.ToLower(pathNode.Value), "x-") {
			skip = true
			continue
		}
		if skip {
			skip = false
			continue
		}
		if i%2 == 0 {
			currentNode = pathNode
			continue
		}
		pathCount++
		go buildPathItem(currentNode, pathNode, bChan, eChan)
	}

	completedItems := 0
	for completedItems < pathCount {
		select {
		case err := <-eChan:
			return err
		case res := <-bChan:
			completedItems++
			pathsMap[res.k] = res.v
		}
	}
	p.PathItems = pathsMap
	return nil
}
