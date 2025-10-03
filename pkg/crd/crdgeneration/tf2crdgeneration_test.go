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

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/crdgeneration/crdboilerplate"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/test"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	k8sschema "k8s.io/apimachinery/pkg/runtime/schema"

	// Register all direct controllers.
	_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/register"
)

var sensitiveSchemaBoilerplate = crdboilerplate.GetSensitiveFieldSchemaBoilerplate()

func TestTFSchemaToJSONSchema(t *testing.T) {
	tests := []struct {
		name string
		tf   schema.Schema
		json apiextensions.JSONSchemaProps
	}{
		{
			name: "primitive field",
			tf: schema.Schema{
				Type: schema.TypeString,
			},
			json: apiextensions.JSONSchemaProps{
				Type: "string",
			},
		},
		{
			name: "map with no value schema",
			tf: schema.Schema{
				Type: schema.TypeMap,
			},
			json: apiextensions.JSONSchemaProps{
				Type: "object",
			},
		},
		{
			name: "sensitive, configurable (required) string field",
			tf: schema.Schema{
				Type:      schema.TypeString,
				Sensitive: true,
				Required:  true,
			},
			json: sensitiveSchemaBoilerplate,
		},
		{
			name: "sensitive, configurable (optional) string field",
			tf: schema.Schema{
				Type:      schema.TypeString,
				Sensitive: true,
				Optional:  true,
			},
			json: sensitiveSchemaBoilerplate,
		},
		{
			name: "sensitive, non-configurable string field",
			tf: schema.Schema{
				Type:      schema.TypeString,
				Sensitive: true,
			},
			json: apiextensions.JSONSchemaProps{
				Type: "string",
			},
		},
		{
			name: "nested sensitive, configurable (required) string field",
			tf: schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sensitive_field": {
							Type:      schema.TypeString,
							Sensitive: true,
							Required:  true,
						},
					},
				},
			},
			json: apiextensions.JSONSchemaProps{
				Type: "object",
				Properties: map[string]apiextensions.JSONSchemaProps{
					"sensitiveField": sensitiveSchemaBoilerplate,
				},
				Required: []string{"sensitiveField"},
			},
		},
		{
			name: "nested sensitive, configurable (optional) string field",
			tf: schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sensitive_field": {
							Type:      schema.TypeString,
							Sensitive: true,
							Optional:  true,
						},
					},
				},
			},
			json: apiextensions.JSONSchemaProps{
				Type: "object",
				Properties: map[string]apiextensions.JSONSchemaProps{
					"sensitiveField": sensitiveSchemaBoilerplate,
				},
			},
		},
		{
			name: "nested sensitive, non-configurable string field",
			tf: schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sensitive_field": {
							Type:      schema.TypeString,
							Sensitive: true,
						},
					},
				},
			},
			json: apiextensions.JSONSchemaProps{
				Type: "object",
				Properties: map[string]apiextensions.JSONSchemaProps{
					"sensitiveField": {
						Type: "string",
					},
				},
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			out := tfSchemaToJSONSchema(&tc.tf)
			if !test.Equals(t, tc.json, out) {
				t.Fatalf("\n\nexpected: %v,\n\nactual: %v", tc.json, *out)
			}
		})
	}
}

func TestHandleResourceReference(t *testing.T) {
	tests := []struct {
		name      string
		refConfig v1alpha1.ReferenceConfig
		in        apiextensions.JSONSchemaProps
		out       apiextensions.JSONSchemaProps
	}{
		{
			name: "embedded reference",
			refConfig: v1alpha1.ReferenceConfig{
				TFField: "bar",
				TypeConfig: v1alpha1.TypeConfig{
					Key: "barRef",
					GVK: k8sschema.GroupVersionKind{
						Kind: "Bar",
					},
				},
			},
			in: apiextensions.JSONSchemaProps{
				Properties: map[string]apiextensions.JSONSchemaProps{
					"bar": {Type: "string"},
				},
			},
			out: apiextensions.JSONSchemaProps{
				Properties: map[string]apiextensions.JSONSchemaProps{
					"barRef": *GetResourceReferenceSchemaFromTypeConfig(v1alpha1.TypeConfig{
						Key: "barRef",
						GVK: k8sschema.GroupVersionKind{
							Kind: "Bar",
						},
					}),
				},
			},
		},
		{
			name: "embedded reference, required",
			refConfig: v1alpha1.ReferenceConfig{
				TFField: "bar",
				TypeConfig: v1alpha1.TypeConfig{
					Key: "barRef",
					GVK: k8sschema.GroupVersionKind{
						Kind: "Bar",
					},
				},
			},
			in: apiextensions.JSONSchemaProps{
				Properties: map[string]apiextensions.JSONSchemaProps{
					"bar": {Type: "string"},
				},
				Required: []string{"bar"},
			},
			out: apiextensions.JSONSchemaProps{
				Properties: map[string]apiextensions.JSONSchemaProps{
					"barRef": *GetResourceReferenceSchemaFromTypeConfig(v1alpha1.TypeConfig{
						Key: "barRef",
						GVK: k8sschema.GroupVersionKind{
							Kind: "Bar",
						},
					}),
				},
				Required: []string{"barRef"},
			},
		},
		{
			name: "complex reference",
			refConfig: v1alpha1.ReferenceConfig{
				TFField: "bar",
				Types: []v1alpha1.TypeConfig{
					{
						Key: "barRef",
						GVK: k8sschema.GroupVersionKind{
							Kind: "Bar",
						},
					},
					{
						Key:            "value",
						JSONSchemaType: "string",
					},
				},
			},
			in: apiextensions.JSONSchemaProps{
				Properties: map[string]apiextensions.JSONSchemaProps{
					"bar": {Type: "string"},
				},
			},
			out: apiextensions.JSONSchemaProps{
				Properties: map[string]apiextensions.JSONSchemaProps{
					"bar": {
						Type: "object",
						Properties: map[string]apiextensions.JSONSchemaProps{
							"barRef": *GetResourceReferenceSchemaFromTypeConfig(v1alpha1.TypeConfig{
								Key: "barRef",
								GVK: k8sschema.GroupVersionKind{
									Kind: "Bar",
								},
							}),
							"value": {Type: "string"},
						},
					},
				},
			},
		},
		{
			name: "reference nested in object",
			refConfig: v1alpha1.ReferenceConfig{
				TFField: "bar.name",
				TypeConfig: v1alpha1.TypeConfig{
					Key: "barRef",
					GVK: k8sschema.GroupVersionKind{
						Kind: "Bar",
					},
				},
			},
			in: apiextensions.JSONSchemaProps{
				Properties: map[string]apiextensions.JSONSchemaProps{
					"bar": {
						Type: "object",
						Properties: map[string]apiextensions.JSONSchemaProps{
							"name": {Type: "string"},
						},
					},
				},
			},
			out: apiextensions.JSONSchemaProps{
				Properties: map[string]apiextensions.JSONSchemaProps{
					"bar": {
						Type: "object",
						Properties: map[string]apiextensions.JSONSchemaProps{
							"barRef": *GetResourceReferenceSchemaFromTypeConfig(v1alpha1.TypeConfig{
								Key: "barRef",
								GVK: k8sschema.GroupVersionKind{
									Kind: "Bar",
								},
							}),
						},
					},
				},
			},
		},
		{
			name: "reference nested in list of objects",
			refConfig: v1alpha1.ReferenceConfig{
				TFField: "bar.name",
				TypeConfig: v1alpha1.TypeConfig{
					Key: "barRef",
					GVK: k8sschema.GroupVersionKind{
						Kind: "Bar",
					},
				},
			},
			in: apiextensions.JSONSchemaProps{
				Properties: map[string]apiextensions.JSONSchemaProps{
					"bar": {
						Type: "array",
						Items: &apiextensions.JSONSchemaPropsOrArray{
							Schema: &apiextensions.JSONSchemaProps{
								Properties: map[string]apiextensions.JSONSchemaProps{
									"name": {Type: "string"},
								},
							},
						},
					},
				},
			},
			out: apiextensions.JSONSchemaProps{
				Properties: map[string]apiextensions.JSONSchemaProps{
					"bar": {
						Type: "array",
						Items: &apiextensions.JSONSchemaPropsOrArray{
							Schema: &apiextensions.JSONSchemaProps{
								Properties: map[string]apiextensions.JSONSchemaProps{
									"barRef": *GetResourceReferenceSchemaFromTypeConfig(v1alpha1.TypeConfig{
										Key: "barRef",
										GVK: k8sschema.GroupVersionKind{
											Kind: "Bar",
										},
									}),
								},
							},
						},
					},
				},
			},
		},
		{
			name: "list of references",
			refConfig: v1alpha1.ReferenceConfig{
				TFField: "bars",
				TypeConfig: v1alpha1.TypeConfig{
					GVK: k8sschema.GroupVersionKind{
						Kind: "Bar",
					},
				},
			},
			in: apiextensions.JSONSchemaProps{
				Properties: map[string]apiextensions.JSONSchemaProps{
					"bars": {
						Type: "array",
						Items: &apiextensions.JSONSchemaPropsOrArray{
							Schema: &apiextensions.JSONSchemaProps{
								Type: "string",
							},
						},
					},
				},
			},
			out: apiextensions.JSONSchemaProps{
				Properties: map[string]apiextensions.JSONSchemaProps{
					"bars": {
						Type: "array",
						Items: &apiextensions.JSONSchemaPropsOrArray{
							Schema: GetResourceReferenceSchemaFromTypeConfig(v1alpha1.TypeConfig{
								Key: "barRef",
								GVK: k8sschema.GroupVersionKind{
									Kind: "Bar",
								},
							}),
						},
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			handleResourceReference(tc.refConfig, &tc.in)
			if !test.Equals(t, tc.in, tc.out) {
				t.Fatalf("expected: %v, actual: %v", tc.out, tc.in)
			}
		})
	}
}

func TestGetDescriptionForExternalRef(t *testing.T) {
	tests := []struct {
		name       string
		typeConfig v1alpha1.TypeConfig
		want       string
	}{
		{
			name: "refWithValueTemplate",
			typeConfig: v1alpha1.TypeConfig{
				TargetField:   "email",
				ValueTemplate: "projects/{{project}}/serviceAccounts/{{value}}",
				GVK: k8sschema.GroupVersionKind{
					Kind: "IAMServiceAccount",
				},
			},
			want: "Allowed value: string of the format `projects/{{project}}/serviceAccounts/{{value}}`, where {{value}} is the `email` field of an `IAMServiceAccount` resource.",
		},
		{
			name: "refWithoutValueTemplate",
			typeConfig: v1alpha1.TypeConfig{
				TargetField: "email",
				GVK: k8sschema.GroupVersionKind{
					Kind: "IAMServiceAccount",
				},
			},
			want: "Allowed value: The `email` field of an `IAMServiceAccount` resource.",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := getDescriptionForExternalRef(tc.typeConfig)
			if got != tc.want {
				t.Errorf("getDescriptionForExternalRef(%v) = '%v', want: '%v'", tc.typeConfig, got, tc.want)
			}
		})
	}
}
