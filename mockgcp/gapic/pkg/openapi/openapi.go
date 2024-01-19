package openapi

import "encoding/json"

// Document is the root of an OpenAPI Document
type Document struct {
	Auth                         json.RawMessage       `json:"auth"`
	BasePath                     string                `json:"basePath"`
	BaseURL                      string                `json:"baseUrl"`
	BatchPath                    string                `json:"batchPath"`
	CanonicalName                string                `json:"canonicalName"`
	Description                  string                `json:"description"`
	DiscoveryVersion             string                `json:"discoveryVersion"`
	DocumentationLink            string                `json:"documentationLink"`
	FullyEncodeReservedExpansion bool                  `json:"fullyEncodeReservedExpansion"`
	Icons                        map[string]string     `json:"icons"`
	ID                           string                `json:"id"`
	Kind                         string                `json:"kind"`
	MtlsRootUrl                  string                `json:"mtlsRootUrl"`
	Name                         string                `json:"name"`
	OwnerDomain                  string                `json:"ownerDomain"`
	OwnerName                    string                `json:"ownerName"`
	Parameters                   map[string]*Parameter `json:"parameters"`
	Protocol                     string                `json:"protocol"`
	Revision                     string                `json:"revision"`
	RootUrl                      string                `json:"rootUrl"`
	Resources                    map[string]*Resource  `json:"resources"`
	ServicePath                  string                `json:"servicePath"`
	Schemas                      map[string]*Property  `json:"schemas"`
	Title                        string                `json:"title"`
	Version                      string                `json:"version"`
}

// Resource describes a REST resource.
type Resource struct {
	Methods map[string]*Method `json:"methods"`
}

// Method is a REST endpoint; a specific method against a resource.
type Method struct {
	Description    string                `json:"description"`
	FlatPath       string                `json:"flatPath"`
	HTTPMethod     string                `json:"httpMethod"`
	ID             string                `json:"id"`
	ParameterOrder []string              `json:"parameterOrder"`
	Parameters     map[string]*Parameter `json:"parameters"`
	Path           string                `json:"path"`
	Request        *Property
	Response       *Property
	Scopes         []string `json:"scopes"`
}

// Parameter is a parameter as passed to a REST Method.
// TODO: Is this actually the same as Property?
type Parameter struct {
	Description      string   `json:"description"`
	Location         string   `json:"location"`
	Required         bool     `json:"required"`
	Type             string   `json:"type"`
	Default          string   `json:"default"`
	Format           string   `json:"format"`
	Pattern          string   `json:"pattern"`
	Enum             []string `json:"enum"`
	EnumDescriptions []string `json:"enumDescriptions"`
}

// Property is a type for a resource or the type of a field of an object type.
type Property struct {
	Ref string `json:"$ref"`

	Description          string               `json:"description"`
	ID                   string               `json:"id"`
	Properties           map[string]*Property `json:"properties"`
	Type                 string               `json:"type"`
	Items                *Property            `json:"items"`
	Default              string               `json:"default"`
	Enum                 []string             `json:"enum"`
	EnumDesc             []string             `json:"enumDescriptions"`
	Format               string               `json:"format"`
	Deprecated           bool                 `json:"deprecated"`
	AdditionalProperties *Property            `json:"additionalProperties"`
}
