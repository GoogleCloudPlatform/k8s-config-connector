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

package crdboilerplate

import (
	"fmt"
	"strings"

	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

func GetSensitiveFieldSchemaBoilerplate() apiextensions.JSONSchemaProps {
	return apiextensions.JSONSchemaProps{
		Type: "object",
		Properties: map[string]apiextensions.JSONSchemaProps{
			"valueFrom": {
				Type:        "object",
				Description: "Source for the field's value. Cannot be used if 'value' is specified.",
				Properties: map[string]apiextensions.JSONSchemaProps{
					"secretKeyRef": {
						Type:        "object",
						Description: "Reference to a value with the given key in the given Secret in the resource's namespace.",
						Properties: map[string]apiextensions.JSONSchemaProps{
							"name": {
								Type:        "string",
								Description: "Name of the Secret to extract a value from.",
							},
							"key": {
								Type:        "string",
								Description: "Key that identifies the value to be extracted.",
							},
						},
						Required: []string{"name", "key"},
					},
				},
			},
			"value": {
				Description: "Value of the field. Cannot be used if 'valueFrom' is specified.",
				Type:        "string",
			},
		},

		// Enforces that 'value' or 'valueFrom' must be specified, but not both
		OneOf: []apiextensions.JSONSchemaProps{
			{
				Required: []string{"value"},
				Not: &apiextensions.JSONSchemaProps{
					Required: []string{"valueFrom"},
				},
			},
			{
				Required: []string{"valueFrom"},
				Not: &apiextensions.JSONSchemaProps{
					Required: []string{"value"},
				},
			},
		},
	}
}

func GetOpenAPIV3SchemaSkeleton() *apiextensions.JSONSchemaProps {
	//TODO See how CRD generation in the k8s "crd" tool works, so you don't have to create this scaffolding by hand
	return &apiextensions.JSONSchemaProps{
		Type: "object",
		Properties: map[string]apiextensions.JSONSchemaProps{
			"apiVersion": {
				Description: "apiVersion defines the versioned schema of this representation of an object. Servers " +
					"should convert recognized schemas to the latest internal value, and may reject unrecognized " +
					"values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
				Type: "string",
			},
			"kind": {
				Description: "kind is a string value representing the REST resource this object represents. Servers " +
					"may infer this from the endpoint the client submits requests to. Cannot be updated. In " +
					"CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#" +
					"types-kinds",
				Type: "string",
			},
			"metadata": {
				Type: "object",
			},
			"status": {
				Properties: map[string]apiextensions.JSONSchemaProps{
					"observedGeneration": {
						Description: "ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. " +
							"If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.",
						Type: "integer",
					},
					"conditions": {
						Description: "Conditions represent the latest available observation of the resource's " +
							"current state.",
						Items: &apiextensions.JSONSchemaPropsOrArray{
							Schema: &apiextensions.JSONSchemaProps{
								Type: "object",
								Properties: map[string]apiextensions.JSONSchemaProps{
									"lastTransitionTime": {
										Description: "Last time the condition transitioned from one status to another.",
										Type:        "string",
									},
									"message": {
										Description: "Human-readable message indicating details about last transition.",
										Type:        "string",
									},
									"reason": {
										Description: "Unique, one-word, CamelCase reason for the condition's last " +
											"transition.",
										Type: "string",
									},
									"status": {
										Description: "Status is the status of the condition. Can be True, False, " +
											"Unknown.",
										Type: "string",
									},
									"type": {
										Description: "Type is the type of the condition.",
										Type:        "string",
									},
								},
							},
						},
						Type: "array",
					},
				},
				Type: "object",
			},
		},
	}
}

func GetMultiKindResourceReferenceSchemaBoilerplate(externalRefDescription string, kinds []string) *apiextensions.JSONSchemaProps {
	return &apiextensions.JSONSchemaProps{
		Type: "object",
		Properties: map[string]apiextensions.JSONSchemaProps{
			"name": {
				Description: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
				Type:        "string",
			},
			"namespace": {
				Description: "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
				Type:        "string",
			},
			"kind": {
				Description: fmt.Sprintf("Kind of the referent. Allowed values: %v", strings.Join(kinds, ",")),
				Type:        "string",
			},
			"external": {
				Description: externalRefDescription,
				Type:        "string",
			},
		},

		// Enforces the following rules:
		// * either 'name' + 'kind' or 'external' (but not both) must be specified
		// * 'namespace' can only be specified if 'name' + 'kind' are specified
		OneOf: []apiextensions.JSONSchemaProps{
			{
				Required: []string{"name", "kind"},
				Not: &apiextensions.JSONSchemaProps{
					Required: []string{"external"},
				},
			},
			{
				Required: []string{"external"},
				Not: &apiextensions.JSONSchemaProps{
					AnyOf: []apiextensions.JSONSchemaProps{
						{
							Required: []string{"name"},
						},
						{
							Required: []string{"namespace"},
						},
						{
							Required: []string{"kind"},
						},
					},
				},
			},
		},
	}
}

func GetResourceReferenceSchemaBoilerplate(externalRefDescription string) *apiextensions.JSONSchemaProps {
	return &apiextensions.JSONSchemaProps{
		Type: "object",
		Properties: map[string]apiextensions.JSONSchemaProps{
			"name": {
				Description: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
				Type:        "string",
			},
			"namespace": {
				Description: "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
				Type:        "string",
			},
			"external": {
				Description: externalRefDescription,
				Type:        "string",
			},
		},

		// Enforces the following rules:
		// * either 'name' or 'external' (but not both) must be specified
		// * 'namespace' can only be specified if 'name' is specified
		OneOf: []apiextensions.JSONSchemaProps{
			{
				Required: []string{"name"},
				Not: &apiextensions.JSONSchemaProps{
					Required: []string{"external"},
				},
			},
			{
				Required: []string{"external"},
				Not: &apiextensions.JSONSchemaProps{
					AnyOf: []apiextensions.JSONSchemaProps{
						{
							Required: []string{"name"},
						},
						{
							Required: []string{"namespace"},
						},
					},
				},
			},
		},
	}
}

func GetAdditionalPrinterColumns() []apiextensions.CustomResourceColumnDefinition {
	// IMPORTANT: any changes to these definitions NEEDS to be copied into the static IAM resources (IAMPolicy, IAMAuditConfig, etc).
	return []apiextensions.CustomResourceColumnDefinition{
		{
			Name:     "Age",
			Type:     "date",
			JSONPath: ".metadata.creationTimestamp",
		},
		{
			Name:        "Ready",
			Type:        "string",
			Description: "When 'True', the most recent reconcile of the resource succeeded",
			JSONPath:    ".status.conditions[?(@.type=='Ready')].status",
		},
		{
			Name:        "Status",
			Type:        "string",
			Description: "The reason for the value in 'Ready'",
			JSONPath:    ".status.conditions[?(@.type=='Ready')].reason",
		},
		{
			Name:        "Status Age",
			Type:        "date",
			Description: "The last transition time for the value in 'Status'",
			JSONPath:    ".status.conditions[?(@.type=='Ready')].lastTransitionTime",
		},
	}
}
