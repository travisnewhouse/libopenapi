package v3

import (
    "github.com/pb33f/libopenapi/datamodel/low"
    "gopkg.in/yaml.v3"
)

type Responses struct {
    Node    *yaml.Node
    Codes   map[string]Response
    Default Response
}

type Response struct {
    Node        *yaml.Node
    Description low.NodeReference[string]
    Headers     map[string]Parameter
    Content     map[string]MediaType
    Extensions  map[string]low.ObjectReference
}
