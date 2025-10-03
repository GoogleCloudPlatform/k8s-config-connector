// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package openapi

import (
	"encoding/json"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/tools/gapic/pkg/sorted"
)

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
	Endpoints                    json.RawMessage       `json:"endpoints"`
	Etag                         json.RawMessage       `json:"etag"`
	FullyEncodeReservedExpansion bool                  `json:"fullyEncodeReservedExpansion"`
	Icons                        map[string]string     `json:"icons"`
	ID                           string                `json:"id"`
	Kind                         string                `json:"kind"`
	Labels                       json.RawMessage       `json:"labels"`
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
	VersionModule                bool                  `json:"version_module"`
}

// Resource describes a REST resource.
type Resource struct {
	Methods   map[string]*Method   `json:"methods"`
	Resources map[string]*Resource `json:"resources"`
}

// Method is a REST endpoint; a specific method against a resource.
type Method struct {
	Deprecated  bool   `json:"deprecated"`
	Description string `json:"description"`
	FlatPath    string `json:"flatPath"`

	HTTPMethod              string                         `json:"httpMethod"`
	ID                      string                         `json:"id"`
	ParameterOrder          []string                       `json:"parameterOrder"`
	Parameters              sorted.Map[string, *Parameter] `json:"parameters"`
	Path                    string                         `json:"path"`
	Request                 *Property
	Response                *Property
	Scopes                  []string        `json:"scopes"`
	MediaUpload             json.RawMessage `json:"mediaUpload"`
	SupportsMediaUpload     bool            `json:"supportsMediaUpload"`
	SupportsMediaDownload   bool            `json:"supportsMediaDownload"`
	UseMediaDownloadService bool            `json:"useMediaDownloadService"`
	SupportsSubscription    bool            `json:"supportsSubscription"`
}

// Parameter is a parameter as passed to a REST Method.
// TODO: Is this actually the same as Property?
type Parameter struct {
	Description      string   `json:"description"`
	Location         string   `json:"location"`
	Minimum          string   `json:"minimum"`
	Maximum          string   `json:"maximum"`
	Required         bool     `json:"required"`
	Type             string   `json:"type"`
	Default          string   `json:"default"`
	Format           string   `json:"format"`
	Pattern          string   `json:"pattern"`
	Enum             []string `json:"enum"`
	EnumDescriptions []string `json:"enumDescriptions"`
	Repeated         bool     `json:"repeated"`
}

// Property is a type for a resource or the type of a field of an object type.
type Property struct {
	Ref           string `json:"$ref"`
	ParameterName string `json:"parameterName"`

	Description          string                        `json:"description"`
	ID                   string                        `json:"id"`
	Properties           sorted.Map[string, *Property] `json:"properties"`
	Type                 string                        `json:"type"`
	Items                *Property                     `json:"items"`
	Default              string                        `json:"default"`
	Enum                 []string                      `json:"enum"`
	EnumDesc             []string                      `json:"enumDescriptions"`
	Format               string                        `json:"format"`
	Deprecated           bool                          `json:"deprecated"`
	AdditionalProperties *Property                     `json:"additionalProperties"`
}
