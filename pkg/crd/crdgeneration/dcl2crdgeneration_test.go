// Copyright 2022 Google LLC
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

package crdgeneration

import (
	"testing"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/crdgeneration/crdboilerplate"
	dclmetadata "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gvks/supportedgvks"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"
	testdclschemaloader "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/dclschemaloader"
	testservicemetadataloader "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test/servicemetadataloader"

	"github.com/google/go-cmp/cmp"
	"github.com/nasa9084/go-openapi"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

func TestDCLSchemaToJSONSchema(t *testing.T) {
	tests := []struct {
		name                     string
		dclSchema                *openapi.Schema
		specSchema               *apiextensions.JSONSchemaProps
		statusSchema             *apiextensions.JSONSchemaProps
		resource                 dclmetadata.Resource
		hasErrorOnSpecGeneration bool
	}{
		{
			name: "primitive fields",
			dclSchema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"foo": {
						Type: "string",
					},
					"bar": {
						Type: "integer",
					},
					"baz": {
						Type: "boolean",
					},
					"quz": {
						Type: "number",
					},
				},
				Required: []string{"foo"},
			},
			specSchema: &apiextensions.JSONSchemaProps{
				Type: "object",
				Properties: map[string]apiextensions.JSONSchemaProps{
					"foo": {
						Type: "string",
					},
					"bar": {
						Type: "integer",
					},
					"baz": {
						Type: "boolean",
					},
					"quz": {
						Type: "number",
					},
				},
				Required: []string{"foo"},
			},
			statusSchema: nil,
		},
		{
			name: "primitive fields with read-only fields",
			dclSchema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"foo": {
						Type: "string",
					},
					"bar": {
						Type:     "integer",
						ReadOnly: true,
					},
				},
				Required: []string{"foo"},
			},
			specSchema: &apiextensions.JSONSchemaProps{
				Type: "object",
				Properties: map[string]apiextensions.JSONSchemaProps{
					"foo": {
						Type: "string",
					},
				},
				Required: []string{"foo"},
			},
			statusSchema: &apiextensions.JSONSchemaProps{
				Type: "object",
				Properties: map[string]apiextensions.JSONSchemaProps{
					"bar": {
						Type: "integer",
					},
				},
			},
		},
		{
			name: "nested fields with read-only fields",
			dclSchema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"foo": {
						Type: "object",
						Properties: map[string]*openapi.Schema{
							"nestedField1": {
								Type: "boolean",
							},
							"nestedField2": {
								Type:     "string",
								ReadOnly: true,
							},
						},
						Required: []string{"nestedField1"},
					},
					"bar": {
						Type: "integer",
					},
				},
				Required: []string{"foo"},
			},
			specSchema: &apiextensions.JSONSchemaProps{
				Type: "object",
				Properties: map[string]apiextensions.JSONSchemaProps{
					"foo": {
						Type: "object",
						Properties: map[string]apiextensions.JSONSchemaProps{
							"nestedField1": {
								Type: "boolean",
							},
						},
						Required: []string{"nestedField1"},
					},
					"bar": {
						Type: "integer",
					},
				},
				Required: []string{"foo"},
			},
			statusSchema: &apiextensions.JSONSchemaProps{
				Type: "object",
				Properties: map[string]apiextensions.JSONSchemaProps{
					"foo": {
						Type: "object",
						Properties: map[string]apiextensions.JSONSchemaProps{
							"nestedField2": {
								Type: "string",
							},
						},
					},
				},
			},
		},
		{
			name: "sensitive field",
			dclSchema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"foo": {
						Type: "string",
						Extension: map[string]interface{}{
							"x-dcl-sensitive": true,
						},
					},
					"bar": {
						Type: "integer",
					},
					// read-only sensitive fields will NOT be converted to secret references
					"baz": {
						Type:     "string",
						ReadOnly: true,
						Extension: map[string]interface{}{
							"x-dcl-sensitive": true,
						},
					},
				},
				Required: []string{"foo"},
			},
			specSchema: &apiextensions.JSONSchemaProps{
				Type: "object",
				Properties: map[string]apiextensions.JSONSchemaProps{
					"foo": crdboilerplate.GetSensitiveFieldSchemaBoilerplate(),
					"bar": {
						Type: "integer",
					},
				},
				Required: []string{"foo"},
			},
			statusSchema: &apiextensions.JSONSchemaProps{
				Type: "object",
				Properties: map[string]apiextensions.JSONSchemaProps{
					"baz": {
						Type: "string",
					},
				},
			},
		},
		{
			name: "reference field",
			dclSchema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"foo": {
						Type: "string",
						Extension: map[string]interface{}{
							"x-dcl-references": []interface{}{
								map[interface{}]interface{}{
									"resource": "Test1/Foo",
									"field":    "name",
								},
							},
						},
					},
					"bar": {
						Type: "integer",
					},
					// read-only reference fields will NOT be converted to resource references
					"baz": {
						Type:     "string",
						ReadOnly: true,
						Extension: map[string]interface{}{
							"x-dcl-references": []interface{}{
								map[interface{}]interface{}{
									"resource": "FakeService/FakeKind",
									"field":    "name",
								},
							},
						},
					},
				},
				Required: []string{"foo"},
			},
			specSchema: &apiextensions.JSONSchemaProps{
				Type: "object",
				Properties: map[string]apiextensions.JSONSchemaProps{
					"fooRef": *crdboilerplate.GetResourceReferenceSchemaBoilerplate(
						"Allowed value: The Google Cloud resource name of a `Test1Foo` resource (format: `projects/{{project}}/foo/{{name}}`).",
					),
					"bar": {
						Type: "integer",
					},
				},
				Required: []string{"fooRef"},
			},
			statusSchema: &apiextensions.JSONSchemaProps{
				Type: "object",
				Properties: map[string]apiextensions.JSONSchemaProps{
					"baz": {
						Type: "string",
					},
				},
			},
		},
		{
			name: "reference nested in object",
			dclSchema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"foo": {
						Type: "object",
						Properties: map[string]*openapi.Schema{
							"bar": {
								Type: "string",
								Extension: map[string]interface{}{
									"x-dcl-references": []interface{}{
										map[interface{}]interface{}{
											"resource": "Test1/Bar",
											"field":    "name",
										},
									},
								},
							},
							"baz": {
								Type: "integer",
							},
						},
					},
				},
			},
			specSchema: &apiextensions.JSONSchemaProps{
				Type: "object",
				Properties: map[string]apiextensions.JSONSchemaProps{
					"foo": {
						Type: "object",
						Properties: map[string]apiextensions.JSONSchemaProps{
							"barRef": *crdboilerplate.GetResourceReferenceSchemaBoilerplate(
								"Allowed value: The Google Cloud resource name of a `Test1Bar` resource (format: `projects/{{project}}/bar/{{name}}`).",
							),
							"baz": {
								Type: "integer",
							},
						},
					},
				},
			},
			statusSchema: nil,
		},
		{
			name: "a list of reference",
			dclSchema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"foos": {
						Type: "array",
						Items: &openapi.Schema{
							Type: "string",
							Extension: map[string]interface{}{
								"x-dcl-references": []interface{}{
									map[interface{}]interface{}{
										"resource": "Test1/Foo",
										"field":    "name",
									},
								},
							},
						},
					},
				},
			},
			specSchema: &apiextensions.JSONSchemaProps{
				Type: "object",
				Properties: map[string]apiextensions.JSONSchemaProps{
					"foos": {
						Type: "array",
						Items: &apiextensions.JSONSchemaPropsOrArray{
							Schema: crdboilerplate.GetResourceReferenceSchemaBoilerplate(
								"Allowed value: The Google Cloud resource name of a `Test1Foo` resource (format: `projects/{{project}}/foo/{{name}}`).",
							),
						},
					},
				},
			},
			statusSchema: nil,
		},
		{
			name: "a list of multi-kinds reference",
			dclSchema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"foos": {
						Type: "array",
						Items: &openapi.Schema{
							Type: "string",
							Extension: map[string]interface{}{
								"x-dcl-references": []interface{}{
									map[interface{}]interface{}{
										"resource": "Test1/Bar",
										"field":    "selfLink",
									},
									map[interface{}]interface{}{
										"resource": "Test2/Baz",
										"field":    "selfLink",
									},
								},
							},
						},
					},
				},
			},
			specSchema: &apiextensions.JSONSchemaProps{
				Type: "object",
				Properties: map[string]apiextensions.JSONSchemaProps{
					"foos": {
						Type: "array",
						Items: &apiextensions.JSONSchemaPropsOrArray{
							Schema: crdboilerplate.GetMultiKindResourceReferenceSchemaBoilerplate(
								"Allowed values:"+
									"\n* The `selfLink` field of a `Test1Bar` resource."+
									"\n* The `selfLink` field of a `Test2Baz` resource.",
								[]string{"Test1Bar", "Test2Baz"},
							),
						},
					},
				},
			},
			statusSchema: nil,
		},
		{
			name: "reference field for multiple kinds",
			dclSchema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"foo": {
						Type: "string",
						Extension: map[string]interface{}{
							"x-dcl-references": []interface{}{
								map[interface{}]interface{}{
									"resource": "Test1/Bar",
									"field":    "selfLink",
								},
								map[interface{}]interface{}{
									"resource": "Test2/Baz",
									"field":    "selfLink",
								},
							},
						},
					},
				},
				Required: []string{"foo"},
			},
			specSchema: &apiextensions.JSONSchemaProps{
				Type: "object",
				Properties: map[string]apiextensions.JSONSchemaProps{
					"fooRef": *crdboilerplate.GetMultiKindResourceReferenceSchemaBoilerplate(
						"Allowed values:"+
							"\n* The `selfLink` field of a `Test1Bar` resource."+
							"\n* The `selfLink` field of a `Test2Baz` resource.",
						[]string{"Test1Bar", "Test2Baz"},
					),
				},
				Required: []string{"fooRef"},
			},
			statusSchema: nil,
		},
		{
			name: "reference to not-yet-supported resources",
			dclSchema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"foo": {
						Type: "string",
						Extension: map[string]interface{}{
							"x-dcl-references": []interface{}{
								map[interface{}]interface{}{
									"resource": "Test1/NotYetSupportedKind",
									"field":    "name",
								},
							},
						},
					},
				},
				Required: []string{"foo"},
			},
			specSchema: &apiextensions.JSONSchemaProps{
				Type: "object",
				Properties: map[string]apiextensions.JSONSchemaProps{
					"fooRef": *markReferencedKindsNotSupported(
						crdboilerplate.GetResourceReferenceSchemaBoilerplate(""),
						[]string{"Test1NotYetSupportedKind"},
					),
				},
				Required: []string{"fooRef"},
			},
		},
		{
			name: "the service of referenced resource is not declared",
			dclSchema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"foo": {
						Type: "string",
						Extension: map[string]interface{}{
							"x-dcl-references": []interface{}{
								map[interface{}]interface{}{
									"resource": "SomeNotDeclaredService/Foo",
									"field":    "name",
								},
							},
						},
					},
				},
				Required: []string{"foo"},
			},
			hasErrorOnSpecGeneration: true,
		},
		{
			name: "referenced resource's (target) field is 'name' but its DCL schema (which contains 'x-dcl-id' extension) is not found",
			dclSchema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"foo": {
						Type: "string",
						Extension: map[string]interface{}{
							"x-dcl-references": []interface{}{
								map[interface{}]interface{}{
									"resource": "Test1/FakeTFBasedResource",
									"field":    "name",
								},
							},
						},
					},
				},
				Required: []string{"foo"},
			},
			hasErrorOnSpecGeneration: true,
		},
		{
			name: "one hierarchical reference: projectRef",
			dclSchema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"project": {
						Type: "string",
						Extension: map[string]interface{}{
							"x-dcl-references": []interface{}{
								map[interface{}]interface{}{
									"resource": "Cloudresourcemanager/Project",
									"field":    "name",
									"parent":   true,
								},
							},
						},
					},
				},
				Required: []string{"project"},
			},
			// TODO(b/186159460): Remove this field once all resources support
			// hierarchical references.
			resource: dclmetadata.Resource{
				SupportsHierarchicalReferences: true,
			},
			specSchema: &apiextensions.JSONSchemaProps{
				Type: "object",
				Properties: map[string]apiextensions.JSONSchemaProps{
					"projectRef": *resourceRefBoilerplateWithDescription(
						"The Project that this resource belongs to.",
						"Allowed value: The Google Cloud resource name of a `Project` resource (format: `projects/{{name}}`).",
					),
				},
				Required: []string{"projectRef"},
			},
			statusSchema: nil,
		},
		{
			name: "one hierarchical reference: folderRef",
			dclSchema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"folder": {
						Type: "string",
						Extension: map[string]interface{}{
							"x-dcl-references": []interface{}{
								map[interface{}]interface{}{
									"resource": "Cloudresourcemanager/Folder",
									"field":    "name",
									"parent":   true,
								},
							},
						},
					},
				},
				Required: []string{"folder"},
			},
			// TODO(b/186159460): Remove this field once all resources support
			// hierarchical references.
			resource: dclmetadata.Resource{
				SupportsHierarchicalReferences: true,
			},
			specSchema: &apiextensions.JSONSchemaProps{
				Type: "object",
				Properties: map[string]apiextensions.JSONSchemaProps{
					"folderRef": *resourceRefBoilerplateWithDescription(
						"The Folder that this resource belongs to.",
						"Allowed value: The Google Cloud resource name of a `Folder` resource (format: `folders/{{name}}`).",
					),
				},
				Required: []string{"folderRef"},
			},
			statusSchema: nil,
		},
		{
			name: "one hierarchical reference: organizationRef",
			dclSchema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"organization": {
						Type: "string",
						Extension: map[string]interface{}{
							"x-dcl-references": []interface{}{
								map[interface{}]interface{}{
									"resource": "Cloudresourcemanager/Organization",
									"field":    "name",
									"parent":   true,
								},
							},
						},
					},
				},
				Required: []string{"organization"},
			},
			// TODO(b/186159460): Remove this field once all resources support
			// hierarchical references.
			resource: dclmetadata.Resource{
				SupportsHierarchicalReferences: true,
			},
			specSchema: &apiextensions.JSONSchemaProps{
				Type: "object",
				Properties: map[string]apiextensions.JSONSchemaProps{
					"organizationRef": *markReferencedKindsNotSupported(
						resourceRefBoilerplateWithDescription(
							"The Organization that this resource belongs to.",
							"Allowed value: The Google Cloud resource name of a Google Cloud Organization (format: `organizations/{{name}}`).",
						),
						[]string{"Organization"},
					),
				},
				Required: []string{"organizationRef"},
			},
			statusSchema: nil,
		},
		{
			name: "multiple hierarchical references: projectRef, folderRef, and organizationRef",
			dclSchema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"parent": {
						Type: "string",
						Extension: map[string]interface{}{
							"x-dcl-references": []interface{}{
								map[interface{}]interface{}{
									"resource": "Cloudresourcemanager/Project",
									"field":    "name",
									"parent":   true,
								},
								map[interface{}]interface{}{
									"resource": "Cloudresourcemanager/Folder",
									"field":    "name",
									"parent":   true,
								},
								map[interface{}]interface{}{
									"resource": "Cloudresourcemanager/Organization",
									"field":    "name",
									"parent":   true,
								},
							},
						},
					},
				},
				Required: []string{"parent"},
			},
			// TODO(b/186159460): Remove this field once all resources support
			// hierarchical references.
			resource: dclmetadata.Resource{
				SupportsHierarchicalReferences: true,
			},
			specSchema: &apiextensions.JSONSchemaProps{
				Type: "object",
				Properties: map[string]apiextensions.JSONSchemaProps{
					"projectRef": *resourceRefBoilerplateWithDescription(
						"The Project that this resource belongs to. Only one of [projectRef, folderRef, organizationRef] may be specified.",
						"Allowed value: The Google Cloud resource name of a `Project` resource (format: `projects/{{name}}`).",
					),
					"folderRef": *resourceRefBoilerplateWithDescription(
						"The Folder that this resource belongs to. Only one of [projectRef, folderRef, organizationRef] may be specified.",
						"Allowed value: The Google Cloud resource name of a `Folder` resource (format: `folders/{{name}}`).",
					),
					"organizationRef": *markReferencedKindsNotSupported(
						resourceRefBoilerplateWithDescription(
							"The Organization that this resource belongs to. Only one of [projectRef, folderRef, organizationRef] may be specified.",
							"Allowed value: The Google Cloud resource name of a Google Cloud Organization (format: `organizations/{{name}}`).",
						),
						[]string{"Organization"},
					),
				},
				OneOf: []apiextensions.JSONSchemaProps{
					{Required: []string{"projectRef"}},
					{Required: []string{"folderRef"}},
					{Required: []string{"organizationRef"}},
				},
			},
			statusSchema: nil,
		},
		{
			name: "multiple hierarchical references: folderRef and organizationRef",
			dclSchema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"parent": {
						Type: "string",
						Extension: map[string]interface{}{
							"x-dcl-references": []interface{}{
								map[interface{}]interface{}{
									"resource": "Cloudresourcemanager/Folder",
									"field":    "name",
									"parent":   true,
								},
								map[interface{}]interface{}{
									"resource": "Cloudresourcemanager/Organization",
									"field":    "name",
									"parent":   true,
								},
							},
						},
					},
				},
				Required: []string{"parent"},
			},
			// TODO(b/186159460): Remove this field once all resources support
			// hierarchical references.
			resource: dclmetadata.Resource{
				SupportsHierarchicalReferences: true,
			},
			specSchema: &apiextensions.JSONSchemaProps{
				Type: "object",
				Properties: map[string]apiextensions.JSONSchemaProps{
					"folderRef": *resourceRefBoilerplateWithDescription(
						"The Folder that this resource belongs to. Only one of [folderRef, organizationRef] may be specified.",
						"Allowed value: The Google Cloud resource name of a `Folder` resource (format: `folders/{{name}}`).",
					),
					"organizationRef": *markReferencedKindsNotSupported(
						resourceRefBoilerplateWithDescription(
							"The Organization that this resource belongs to. Only one of [folderRef, organizationRef] may be specified.",
							"Allowed value: The Google Cloud resource name of a Google Cloud Organization (format: `organizations/{{name}}`).",
						),
						[]string{"Organization"},
					),
				},
				OneOf: []apiextensions.JSONSchemaProps{
					{Required: []string{"folderRef"}},
					{Required: []string{"organizationRef"}},
				},
			},
			statusSchema: nil,
		},
		{
			name: "resource that supports container annotations and has one hierarchical reference: projectRef",
			dclSchema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"project": {
						Type: "string",
						Extension: map[string]interface{}{
							"x-dcl-references": []interface{}{
								map[interface{}]interface{}{
									"resource": "Cloudresourcemanager/Project",
									"field":    "name",
									"parent":   true,
								},
							},
						},
					},
				},
				Required: []string{"project"},
			},
			resource: dclmetadata.Resource{
				SupportsContainerAnnotations: true,
				// TODO(b/186159460): Remove this field once all resources
				// support hierarchical references.
				SupportsHierarchicalReferences: true,
			},
			specSchema: &apiextensions.JSONSchemaProps{
				Type: "object",
				Properties: map[string]apiextensions.JSONSchemaProps{
					"projectRef": *resourceRefBoilerplateWithDescription(
						"The Project that this resource belongs to.",
						"Allowed value: The Google Cloud resource name of a `Project` resource (format: `projects/{{name}}`).",
					),
				},
			},
			statusSchema: nil,
		},
		{
			name: "resource that supports container annotations and has multiple hierarchical references: projectRef, folderRef, and organizationRef",
			dclSchema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"parent": {
						Type: "string",
						Extension: map[string]interface{}{
							"x-dcl-references": []interface{}{
								map[interface{}]interface{}{
									"resource": "Cloudresourcemanager/Project",
									"field":    "name",
									"parent":   true,
								},
								map[interface{}]interface{}{
									"resource": "Cloudresourcemanager/Folder",
									"field":    "name",
									"parent":   true,
								},
								map[interface{}]interface{}{
									"resource": "Cloudresourcemanager/Organization",
									"field":    "name",
									"parent":   true,
								},
							},
						},
					},
				},
				Required: []string{"parent"},
			},
			resource: dclmetadata.Resource{
				SupportsContainerAnnotations: true,
				// TODO(b/186159460): Remove this field once all resources
				// support hierarchical references.
				SupportsHierarchicalReferences: true,
			},
			specSchema: &apiextensions.JSONSchemaProps{
				Type: "object",
				Properties: map[string]apiextensions.JSONSchemaProps{
					"projectRef": *resourceRefBoilerplateWithDescription(
						"The Project that this resource belongs to. Only one of [projectRef, folderRef, organizationRef] may be specified.",
						"Allowed value: The Google Cloud resource name of a `Project` resource (format: `projects/{{name}}`).",
					),
					"folderRef": *resourceRefBoilerplateWithDescription(
						"The Folder that this resource belongs to. Only one of [projectRef, folderRef, organizationRef] may be specified.",
						"Allowed value: The Google Cloud resource name of a `Folder` resource (format: `folders/{{name}}`).",
					),
					"organizationRef": *markReferencedKindsNotSupported(
						resourceRefBoilerplateWithDescription(
							"The Organization that this resource belongs to. Only one of [projectRef, folderRef, organizationRef] may be specified.",
							"Allowed value: The Google Cloud resource name of a Google Cloud Organization (format: `organizations/{{name}}`).",
						),
						[]string{"Organization"},
					),
				},
				OneOf: []apiextensions.JSONSchemaProps{
					{Required: []string{"projectRef"}},
					{Required: []string{"folderRef"}},
					{Required: []string{"organizationRef"}},
					{
						Not: &apiextensions.JSONSchemaProps{
							AnyOf: []apiextensions.JSONSchemaProps{
								{Required: []string{"projectRef"}},
								{Required: []string{"folderRef"}},
								{Required: []string{"organizationRef"}},
							},
						},
					},
				},
			},
			statusSchema: nil,
		},
		{
			name: "container field will be ignored",
			dclSchema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"foo": {
						Type: "string",
					},
					"project": {
						Type: "string",
					},
				},
				Required: []string{"foo", "project"},
			},
			specSchema: &apiextensions.JSONSchemaProps{
				Type: "object",
				Properties: map[string]apiextensions.JSONSchemaProps{
					"foo": {
						Type: "string",
					},
				},
				Required: []string{"foo"},
			},
			statusSchema: nil,
		},
		{
			name: "name field will be converted to ResourceID",
			dclSchema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"foo": {
						Type: "string",
					},
					"name": {
						Type: "string",
					},
				},
				Required: []string{"foo", "name"},
			},
			specSchema: &apiextensions.JSONSchemaProps{
				Type: "object",
				Properties: map[string]apiextensions.JSONSchemaProps{
					"foo": {
						Type: "string",
					},
					"resourceID": {
						Description: GenerateResourceIDFieldDescription("name", false),
						Type:        "string",
					},
				},
				Required: []string{"foo"},
			},
			statusSchema: nil,
		},
		{
			name: "string-object maps",
			dclSchema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"baz": {
						Type: "object",
						AdditionalProperties: &openapi.Schema{
							Type: "object",
							Properties: map[string]*openapi.Schema{
								"objectField1": {
									Type: "string",
								},
								"objectField2": {
									ReadOnly: true,
									Type:     "integer",
								},
								"objectField3": {
									Type: "array",
									Items: &openapi.Schema{
										Type: "string",
									},
								},
								"foo": {
									Type: "string",
									Extension: map[string]interface{}{
										"x-dcl-references": []interface{}{
											map[interface{}]interface{}{
												"resource": "Test1/Foo",
												"field":    "name",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			specSchema: &apiextensions.JSONSchemaProps{
				Type: "object",
				Properties: map[string]apiextensions.JSONSchemaProps{
					"baz": {
						Type: "object",
						AdditionalProperties: &apiextensions.JSONSchemaPropsOrBool{
							Schema: &apiextensions.JSONSchemaProps{
								Type: "object",
								Properties: map[string]apiextensions.JSONSchemaProps{
									"objectField1": {
										Type: "string",
									},
									"objectField2": {
										Type: "integer",
									},
									"objectField3": {
										Type: "array",
										Items: &apiextensions.JSONSchemaPropsOrArray{
											Schema: &apiextensions.JSONSchemaProps{
												Type: "string",
											},
										},
									},
									"fooRef": *crdboilerplate.GetResourceReferenceSchemaBoilerplate(
										"Allowed value: The Google Cloud resource name of a `Test1Foo` resource (format: `projects/{{project}}/foo/{{name}}`).",
									),
								},
							},
						},
					},
				},
			},
			statusSchema: nil,
		},
		{
			name: "array type, additionalProperties",
			dclSchema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"foo": {
						Type: "array",
						Items: &openapi.Schema{
							Type: "string",
						},
					},
					"baz": {
						Type: "object",
						AdditionalProperties: &openapi.Schema{
							Type: "string",
						},
					},
					"qux": {
						Type: "array",
						Items: &openapi.Schema{
							Type: "string",
						},
						ReadOnly: true,
					},
				},
				Required: []string{"foo"},
			},
			specSchema: &apiextensions.JSONSchemaProps{
				Type: "object",
				Properties: map[string]apiextensions.JSONSchemaProps{
					"foo": {
						Type: "array",
						Items: &apiextensions.JSONSchemaPropsOrArray{
							Schema: &apiextensions.JSONSchemaProps{
								Type: "string",
							},
						},
					},
					"baz": {
						Type: "object",
						AdditionalProperties: &apiextensions.JSONSchemaPropsOrBool{
							Schema: &apiextensions.JSONSchemaProps{
								Type: "string",
							},
						},
					},
				},
				Required: []string{"foo"},
			},
			statusSchema: &apiextensions.JSONSchemaProps{
				Type: "object",
				Properties: map[string]apiextensions.JSONSchemaProps{
					"qux": {
						Type: "array",
						Items: &apiextensions.JSONSchemaPropsOrArray{
							Schema: &apiextensions.JSONSchemaProps{
								Type: "string",
							},
						},
					},
				},
			},
		},
		{
			name: "Enum is not exposed by design in CRD schema",
			dclSchema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"foo": {
						Type: "string",
						Enum: []string{"VAL1", "VAL2"},
					},
					"baz": {
						Type:     "string",
						ReadOnly: true,
						Enum:     []string{"VAL1", "VAL2"},
					},
				},
				Required: []string{"foo"},
			},
			specSchema: &apiextensions.JSONSchemaProps{
				Type: "object",
				Properties: map[string]apiextensions.JSONSchemaProps{
					"foo": {
						Type: "string",
					},
				},
				Required: []string{"foo"},
			},
			statusSchema: &apiextensions.JSONSchemaProps{
				Type: "object",
				Properties: map[string]apiextensions.JSONSchemaProps{
					"baz": {
						Type: "string",
					},
				},
			},
		},
		{
			name: "read-only fields in arrays will be preserved in spec",
			dclSchema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"foo": {
						Type: "array",
						Items: &openapi.Schema{
							Type: "object",
							Properties: map[string]*openapi.Schema{
								"nestedField1": {
									Type: "boolean",
								},
								"nestedField2": {
									Type:     "string",
									ReadOnly: true,
								},
								"nestedObject": {
									Type: "object",
									Properties: map[string]*openapi.Schema{
										"state": {
											Type:     "string",
											ReadOnly: true,
										},
										"name": {
											Type: "string",
										},
									},
								},
								"readOnlySensitiveField": {
									Type:     "string",
									ReadOnly: true,
									Extension: map[string]interface{}{
										"x-dcl-sensitive": true,
									},
								},
								"readOnlyReferenceField": {
									Type:     "string",
									ReadOnly: true,
									Extension: map[string]interface{}{
										"x-dcl-references": []interface{}{
											map[interface{}]interface{}{
												"resource": "FakeService/FakeKind",
												"field":    "name",
											},
										},
									},
								},
							},
							Required: []string{"nestedField1"},
						},
					},
					"bar": {
						Type: "integer",
					},
				},
				Required: []string{"foo"},
			},
			specSchema: &apiextensions.JSONSchemaProps{
				Type: "object",
				Properties: map[string]apiextensions.JSONSchemaProps{
					"foo": {
						Type: "array",
						Items: &apiextensions.JSONSchemaPropsOrArray{
							Schema: &apiextensions.JSONSchemaProps{
								Type: "object",
								Properties: map[string]apiextensions.JSONSchemaProps{
									"nestedField1": {
										Type: "boolean",
									},
									"nestedField2": {
										Type: "string",
									},
									"nestedObject": {
										Type: "object",
										Properties: map[string]apiextensions.JSONSchemaProps{
											"name": {
												Type: "string",
											},
											"state": {
												Type: "string",
											},
										},
									},
									"readOnlySensitiveField": {
										Type: "string",
									},
									"readOnlyReferenceField": {
										Type: "string",
									},
								},
								Required: []string{"nestedField1"},
							},
						},
					},
					"bar": {
						Type: "integer",
					},
				},
				Required: []string{"foo"},
			},
			statusSchema: nil,
		},
		{
			name: "empty spec",
			dclSchema: &openapi.Schema{
				Type: "object",
			},
			specSchema: &apiextensions.JSONSchemaProps{
				Type: "object",
			},
			statusSchema: nil,
		},
		{
			// in this case, the labels field should be removed from the
			// spec, as the values will be sourced from Kubernetes labels.
			name: "field specified as x-dcl-labels",
			dclSchema: &openapi.Schema{
				Type: "object",
				Extension: map[string]interface{}{
					"x-dcl-labels": "labels",
				},
				Properties: map[string]*openapi.Schema{
					"labels": {
						Type: "object",
						AdditionalProperties: &openapi.Schema{
							Type: "string",
						},
					},
				},
			},
			specSchema: &apiextensions.JSONSchemaProps{
				Type: "object",
			},
			statusSchema: nil,
		},
		{
			// labels field should not be removed from the spec. This was
			// the previous no longer desired behavior, so verifying there is
			// no regression.
			name: "labels field exists, but not specified as x-dcl-labels",
			dclSchema: &openapi.Schema{
				Type: "object",
				Properties: map[string]*openapi.Schema{
					"labels": {
						Type: "object",
						AdditionalProperties: &openapi.Schema{
							Type: "string",
						},
					},
				},
			},
			specSchema: &apiextensions.JSONSchemaProps{
				Type: "object",
				Properties: map[string]apiextensions.JSONSchemaProps{
					"labels": {
						Type: "object",
						AdditionalProperties: &apiextensions.JSONSchemaPropsOrBool{
							Schema: &apiextensions.JSONSchemaProps{
								Type: "string",
							},
						},
					},
				},
			},
			statusSchema: nil,
		},
	}

	smLoader := servicemappingloader.NewFromServiceMappings(test.FakeServiceMappingsWithHierarchicalResources())
	serviceMetadataLoader := dclmetadata.NewFromServiceList(testservicemetadataloader.FakeServiceMetadataWithHierarchicalResources())
	dclSchemaLoader := testdclschemaloader.New(dclSchemaMap())
	allSupportedGVKs, err := supportedgvks.All(smLoader, serviceMetadataLoader)
	if err != nil {
		t.Fatalf("error loading all supported GVKs: %v", err)
	}
	a := New(serviceMetadataLoader, dclSchemaLoader, allSupportedGVKs)
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			spec, err := a.generateSpecJSONSchema(tc.dclSchema, tc.resource)
			if tc.hasErrorOnSpecGeneration {
				if err == nil {
					t.Fatalf("got nil, but expect to get error on generating the spec")
				}
				return
			}
			if err != nil {
				t.Fatalf("error generating spec schema: %v", err)
			}
			if !test.Equals(t, tc.specSchema, spec) {
				t.Fatalf("unexpected spec diff (-want +got): \n%v", cmp.Diff(tc.specSchema, spec))
			}
			status, err := generateStatusJSONSchema(tc.dclSchema)
			if err != nil {
				t.Fatalf("error generating status schema: %v", err)
			}
			if !test.Equals(t, tc.statusSchema, status) {
				t.Fatalf("unexpected status diff (-want +got): \n%v", cmp.Diff(tc.statusSchema, status))
			}
		})
	}
}

func dclSchemaMap() map[string]*openapi.Schema {
	return map[string]*openapi.Schema{
		"test1_beta_foo": &openapi.Schema{
			Extension: map[string]interface{}{
				"x-dcl-id": "projects/{{project}}/foo/{{name}}",
			},
		},
		"test1_beta_bar": &openapi.Schema{
			Extension: map[string]interface{}{
				"x-dcl-id": "projects/{{project}}/bar/{{name}}",
			},
		},

		// Add the following to the list of fake DCL schemas to allow for our
		// test to test resources that reference hierarchical resources
		// (e.g. "Cloudresourcemanager/Project").
		"cloudresourcemanager_ga_project": &openapi.Schema{
			Extension: map[string]interface{}{
				"x-dcl-id": "projects/{{name}}",
			},
		},
		"cloudresourcemanager_ga_folder": &openapi.Schema{
			Extension: map[string]interface{}{
				"x-dcl-id": "folders/{{name}}",
			},
		},
	}
}

func resourceRefBoilerplateWithDescription(description, externalRefDescription string) *apiextensions.JSONSchemaProps {
	schema := crdboilerplate.GetResourceReferenceSchemaBoilerplate(externalRefDescription)
	schema.Description = description
	return schema
}

func markReferencedKindsNotSupported(schema *apiextensions.JSONSchemaProps, kinds []string) *apiextensions.JSONSchemaProps {
	s := schema.DeepCopy()
	MarkReferencedKindsNotSupported(s, kinds)
	return s
}
