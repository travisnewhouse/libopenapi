// Copyright 2022 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package v2

import (
	"github.com/pb33f/libopenapi/datamodel/low"
	"github.com/pb33f/libopenapi/index"
	"gopkg.in/yaml.v3"
)

// SecurityScheme is a low-level representation of a Swagger / OpenAPI 2 SecurityScheme object.
//
// SecurityScheme allows the definition of a security scheme that can be used by the operations. Supported schemes are
// basic authentication, an API key (either as a header or as a query parameter) and OAuth2's common flows
// (implicit, password, application and access code)
//  - https://swagger.io/specification/v2/#securityDefinitionsObject
type SecurityScheme struct {
	Type             low.NodeReference[string]
	Description      low.NodeReference[string]
	Name             low.NodeReference[string]
	In               low.NodeReference[string]
	Flow             low.NodeReference[string]
	AuthorizationUrl low.NodeReference[string]
	TokenUrl         low.NodeReference[string]
	Scopes           low.NodeReference[*Scopes]
	Extensions       map[low.KeyReference[string]]low.ValueReference[any]
}

// Build will extract extensions and scopes from the node.
func (ss *SecurityScheme) Build(root *yaml.Node, idx *index.SpecIndex) error {
	ss.Extensions = low.ExtractExtensions(root)

	scopes, sErr := low.ExtractObject[*Scopes](ScopesLabel, root, idx)
	if sErr != nil {
		return sErr
	}
	ss.Scopes = scopes
	return nil
}
