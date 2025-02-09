// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

// Package v3 represents all OpenAPI 3+ low-level models. Low-level models are more difficult to navigate
// than higher-level models, however they are packed with all the raw AST and node data required to perform
// any kind of analysis on the underlying data.
//
// Every property is wrapped in a NodeReference or a KeyReference or a ValueReference.
package v3

import (
	"github.com/pb33f/libopenapi/datamodel/low"
	"github.com/pb33f/libopenapi/datamodel/low/base"
	"github.com/pb33f/libopenapi/index"
)

type Document struct {

	// Version is the version of OpenAPI being used, extracted from the 'openapi: x.x.x' definition.
	// This is not a standard property of the OpenAPI model, it's a convenience mechanism only.
	Version low.ValueReference[string]

	// Info represents a specification Info definitions
	// Provides metadata about the API. The metadata MAY be used by tooling as required.
	// - https://spec.openapis.org/oas/v3.1.0#info-object
	Info low.NodeReference[*base.Info]

	// JsonSchemaDialect is a 3.1+ property that sets the dialect to use for validating *base.Schema definitions
	// The default value for the $schema keyword within Schema Objects contained within this OAS document.
	// This MUST be in the form of a URI.
	// - https://spec.openapis.org/oas/v3.1.0#schema-object
	JsonSchemaDialect low.NodeReference[string] // 3.1

	// Webhooks is a 3.1+ property that is similar to callbacks, except, this defines incoming webhooks.
	// The incoming webhooks that MAY be received as part of this API and that the API consumer MAY choose to implement.
	// Closely related to the callbacks feature, this section describes requests initiated other than by an API call,
	// for example by an out-of-band registration. The key name is a unique string to refer to each webhook,
	// while the (optionally referenced) Path Item Object describes a request that may be initiated by the API provider
	// and the expected responses. An example is available.
	Webhooks low.NodeReference[map[low.KeyReference[string]]low.ValueReference[*PathItem]] // 3.1

	// Servers is a slice of Server instances which provide connectivity information to a target server. If the servers
	// property is not provided, or is an empty array, the default value would be a Server Object with a url value of /.
	// - https://spec.openapis.org/oas/v3.1.0#server-object
	Servers low.NodeReference[[]low.ValueReference[*Server]]

	// Paths contains all the PathItem definitions for the specification.
	// The available paths and operations for the API, The most important part of ths spec.
	// - https://spec.openapis.org/oas/v3.1.0#paths-object
	Paths low.NodeReference[*Paths]

	// Components is an element to hold various schemas for the document.
	// - https://spec.openapis.org/oas/v3.1.0#components-object
	Components low.NodeReference[*Components]

	// Security contains global security requirements/roles for the specification
	// A declaration of which security mechanisms can be used across the API. The list of values includes alternative
	// security requirement objects that can be used. Only one of the security requirement objects need to be satisfied
	// to authorize a request. Individual operations can override this definition. To make security optional,
	// an empty security requirement ({}) can be included in the array.
	// - https://spec.openapis.org/oas/v3.1.0#security-requirement-object
	Security low.NodeReference[*SecurityRequirement]

	// Tags is a slice of base.Tag instances defined by the specification
	// A list of tags used by the document with additional metadata. The order of the tags can be used to reflect on
	// their order by the parsing tools. Not all tags that are used by the Operation Object must be declared.
	// The tags that are not declared MAY be organized randomly or based on the tools’ logic.
	// Each tag name in the list MUST be unique.
	// - https://spec.openapis.org/oas/v3.1.0#tag-object
	Tags low.NodeReference[[]low.ValueReference[*base.Tag]]

	// ExternalDocs is an instance of base.ExternalDoc for.. well, obvious really, innit.
	// - https://spec.openapis.org/oas/v3.1.0#external-documentation-object
	ExternalDocs low.NodeReference[*base.ExternalDoc]

	// Extensions contains all custom extensions defined for the top-level document.
	Extensions map[low.KeyReference[string]]low.ValueReference[any]

	// Index is a reference to the *index.SpecIndex that was created for the document and used
	// as a guide when building out the Document. Ideal if further processing is required on the model and
	// the original details are required to continue the work.
	//
	// This property is not a part of the OpenAPI schema, this is custom to libopenapi.
	Index *index.SpecIndex
}

// TODO: this is early prototype mutation/modification code, keeping it around for later.
//func (d *Document) AddTag() *base.Tag {
//	t := base.NewTag()
//	//d.Tags.KeyNode
//	t.Name.Value = "nice new tag"
//
//	dat, _ := yaml.Marshal(t)
//	var inject yaml.Node
//	_ = yaml.Unmarshal(dat, &inject)
//
//	d.Tags.ValueNode.Content = append(d.Tags.ValueNode.Content, inject.Content[0])
//
//	return t
//}
