// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package base

import (
	"fmt"
	lowmodel "github.com/pb33f/libopenapi/datamodel/low"
	lowbase "github.com/pb33f/libopenapi/datamodel/low/base"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
	"testing"
)

func TestNewInfo(t *testing.T) {

	var cNode yaml.Node

	yml := `title: chicken
description: nugget
termsOfService: chicken soup
contact:
  name: buckaroo
license:
  name: pb33f
  url: https://pb33f.io
version: 99.99`

	_ = yaml.Unmarshal([]byte(yml), &cNode)

	var lowInfo lowbase.Info
	_ = lowmodel.BuildModel(&cNode, &lowInfo)
	_ = lowInfo.Build(cNode.Content[0], nil)

	highInfo := NewInfo(&lowInfo)

	assert.Equal(t, "chicken", highInfo.Title)
	assert.Equal(t, "nugget", highInfo.Description)
	assert.Equal(t, "chicken soup", highInfo.TermsOfService)
	assert.Equal(t, "buckaroo", highInfo.Contact.Name)
	assert.Equal(t, "pb33f", highInfo.License.Name)
	assert.Equal(t, "https://pb33f.io", highInfo.License.URL)
	assert.Equal(t, "99.99", highInfo.Version)

	wentLow := highInfo.GoLow()
	assert.Equal(t, 9, wentLow.Version.ValueNode.Line)

	wentLower := highInfo.License.GoLow()
	assert.Equal(t, 8, wentLower.URL.ValueNode.Line)

}

func ExampleNewInfo() {

	// create an example info object (including contact and license)
	// this can be either JSON or YAML.
	yml := `title: some spec by some company
description: this is a specification, for an API, by a company.
termsOfService: https://pb33f.io/tos
contact:
  name: buckaroo
license:
  name: MIT
  url: https://opensource.org/licenses/MIT
version: 1.2.3`

	// unmarshal raw bytes
	var node yaml.Node
	_ = yaml.Unmarshal([]byte(yml), &node)

	// build out the low-level model
	var lowInfo lowbase.Info
	_ = lowmodel.BuildModel(&node, &lowInfo)
	_ = lowInfo.Build(node.Content[0], nil)

	// build the high level model
	highInfo := NewInfo(&lowInfo)

	// print out the contact name.
	fmt.Print(highInfo.Contact.Name)
	// Output: buckaroo
}

func ExampleNewLicense() {

	// create an example license object
	// this can be either JSON or YAML.
	yml := `name: MIT
url: https://opensource.org/licenses/MIT`

	// unmarshal raw bytes
	var node yaml.Node
	_ = yaml.Unmarshal([]byte(yml), &node)

	// build out the low-level model
	var lowLicense lowbase.License
	_ = lowmodel.BuildModel(&node, &lowLicense)
	_ = lowLicense.Build(node.Content[0], nil)

	// build the high level model
	highLicense := NewLicense(&lowLicense)

	// print out the contact name.
	fmt.Print(highLicense.Name)
	// Output: MIT
}
