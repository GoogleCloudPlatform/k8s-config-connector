// Copyright 2024 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package beta

import (
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func DCLWorkloadIdentityPoolProviderSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "Iam/WorkloadIdentityPoolProvider",
			Description: "The Iam WorkloadIdentityPoolProvider resource",
			StructName:  "WorkloadIdentityPoolProvider",
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a WorkloadIdentityPoolProvider",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "workloadIdentityPoolProvider",
						Required:    true,
						Description: "A full instance of a WorkloadIdentityPoolProvider",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a WorkloadIdentityPoolProvider",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "workloadIdentityPoolProvider",
						Required:    true,
						Description: "A full instance of a WorkloadIdentityPoolProvider",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a WorkloadIdentityPoolProvider",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "workloadIdentityPoolProvider",
						Required:    true,
						Description: "A full instance of a WorkloadIdentityPoolProvider",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all WorkloadIdentityPoolProvider",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:     "project",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
					dcl.PathParameters{
						Name:     "location",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
					dcl.PathParameters{
						Name:     "workloadIdentityPool",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
				},
			},
			List: &dcl.Path{
				Description: "The function used to list information about many WorkloadIdentityPoolProvider",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:     "project",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
					dcl.PathParameters{
						Name:     "location",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
					dcl.PathParameters{
						Name:     "workloadIdentityPool",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
				},
			},
		},
		Components: &dcl.Components{
			Schemas: map[string]*dcl.Component{
				"WorkloadIdentityPoolProvider": &dcl.Component{
					Title:           "WorkloadIdentityPoolProvider",
					ID:              "projects/{{project}}/locations/{{location}}/workloadIdentityPools/{{workload_identity_pool}}/providers/{{name}}",
					UsesStateHint:   true,
					ParentContainer: "project",
					HasCreate:       true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"name",
							"project",
							"location",
							"workloadIdentityPool",
						},
						Properties: map[string]*dcl.Property{
							"attributeCondition": &dcl.Property{
								Type:        "string",
								GoName:      "AttributeCondition",
								Description: "[A Common Expression Language](https://opensource.google/projects/cel) expression, in plain text, to restrict what otherwise valid authentication credentials issued by the provider should not be accepted. The expression must output a boolean representing whether to allow the federation. The following keywords may be referenced in the expressions: * `assertion`: JSON representing the authentication credential issued by the provider. * `google`: The Google attributes mapped from the assertion in the `attribute_mappings`. * `attribute`: The custom attributes mapped from the assertion in the `attribute_mappings`. The maximum length of the attribute condition expression is 4096 characters. If unspecified, all valid authentication credential are accepted. The following example shows how to only allow credentials with a mapped `google.groups` value of `admins`: ``` \"'admins' in google.groups\" ```",
							},
							"attributeMapping": &dcl.Property{
								Type: "object",
								AdditionalProperties: &dcl.Property{
									Type: "string",
								},
								GoName:      "AttributeMapping",
								Description: "Maps attributes from authentication credentials issued by an external identity provider to Google Cloud attributes, such as `subject` and `segment`. Each key must be a string specifying the Google Cloud IAM attribute to map to. The following keys are supported: * `google.subject`: The principal IAM is authenticating. You can reference this value in IAM bindings. This is also the subject that appears in Cloud Logging logs. Cannot exceed 127 characters. * `google.groups`: Groups the external identity belongs to. You can grant groups access to resources using an IAM `principalSet` binding; access applies to all members of the group. You can also provide custom attributes by specifying `attribute.{custom_attribute}`, where `{custom_attribute}` is the name of the custom attribute to be mapped. You can define a maximum of 50 custom attributes. The maximum length of a mapped attribute key is 100 characters, and the key may only contain the characters [a-z0-9_]. You can reference these attributes in IAM policies to define fine-grained access for a workload to Google Cloud resources. For example: * `google.subject`: `principal://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/subject/{value}` * `google.groups`: `principalSet://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/group/{value}` * `attribute.{custom_attribute}`: `principalSet://iam.googleapis.com/projects/{project}/locations/{location}/workloadIdentityPools/{pool}/attribute.{custom_attribute}/{value}` Each value must be a [Common Expression Language] (https://opensource.google/projects/cel) function that maps an identity provider credential to the normalized attribute specified by the corresponding map key. You can use the `assertion` keyword in the expression to access a JSON representation of the authentication credential issued by the provider. The maximum length of an attribute mapping expression is 2048 characters. When evaluated, the total size of all mapped attributes must not exceed 8KB. For AWS providers, if no attribute mapping is defined, the following default mapping applies: ``` { \"google.subject\":\"assertion.arn\", \"attribute.aws_role\": \"assertion.arn.contains('assumed-role')\" \" ? assertion.arn.extract('{account_arn}assumed-role/')\" \" + 'assumed-role/'\" \" + assertion.arn.extract('assumed-role/{role_name}/')\" \" : assertion.arn\", } ``` If any custom attribute mappings are defined, they must include a mapping to the `google.subject` attribute. For OIDC providers, you must supply a custom mapping, which must include the `google.subject` attribute. For example, the following maps the `sub` claim of the incoming credential to the `subject` attribute on a Google token: ``` {\"google.subject\": \"assertion.sub\"} ```",
							},
							"aws": &dcl.Property{
								Type:        "object",
								GoName:      "Aws",
								GoType:      "WorkloadIdentityPoolProviderAws",
								Description: "An Amazon Web Services identity provider.",
								Conflicts: []string{
									"oidc",
								},
								Required: []string{
									"accountId",
								},
								Properties: map[string]*dcl.Property{
									"accountId": &dcl.Property{
										Type:        "string",
										GoName:      "AccountId",
										Description: "Required. The AWS account ID.",
									},
									"stsUri": &dcl.Property{
										Type:        "array",
										GoName:      "StsUri",
										Description: "A list of AWS STS URIs that can be used when exchanging credentials. If not provided, any valid AWS STS URI is allowed. URIs must use the form `https://sts.amazonaws.com` or `https://sts.{region}.amazonaws.com`, where {region} is a valid AWS region. You can specify a maximum of 25 URIs.",
										SendEmpty:   true,
										ListType:    "list",
										Items: &dcl.Property{
											Type:   "string",
											GoType: "string",
										},
										Unreadable: true,
									},
								},
							},
							"description": &dcl.Property{
								Type:        "string",
								GoName:      "Description",
								Description: "A description for the provider. Cannot exceed 256 characters.",
							},
							"disabled": &dcl.Property{
								Type:        "boolean",
								GoName:      "Disabled",
								Description: "Whether the provider is disabled. You cannot use a disabled provider to exchange tokens. However, existing tokens still grant access.",
							},
							"displayName": &dcl.Property{
								Type:        "string",
								GoName:      "DisplayName",
								Description: "A display name for the provider. Cannot exceed 32 characters.",
							},
							"location": &dcl.Property{
								Type:        "string",
								GoName:      "Location",
								Description: "The location for the resource",
								Immutable:   true,
							},
							"name": &dcl.Property{
								Type:        "string",
								GoName:      "Name",
								Description: "Output only. The resource name of the provider.",
								Immutable:   true,
							},
							"oidc": &dcl.Property{
								Type:        "object",
								GoName:      "Oidc",
								GoType:      "WorkloadIdentityPoolProviderOidc",
								Description: "An OpenId Connect 1.0 identity provider.",
								Conflicts: []string{
									"aws",
								},
								Required: []string{
									"issuerUri",
								},
								Properties: map[string]*dcl.Property{
									"allowedAudiences": &dcl.Property{
										Type:        "array",
										GoName:      "AllowedAudiences",
										Description: "Acceptable values for the `aud` field (audience) in the OIDC token. Token exchange requests are rejected if the token audience does not match one of the configured values. Each audience may be at most 256 characters. A maximum of 10 audiences may be configured. If this list is empty, the OIDC token audience must be equal to the full canonical resource name of the WorkloadIdentityPoolProvider, with or without the HTTPS prefix. For example: ``` //iam.googleapis.com/projects//locations//workloadIdentityPools//providers/ https://iam.googleapis.com/projects//locations//workloadIdentityPools//providers/ ```",
										SendEmpty:   true,
										ListType:    "list",
										Items: &dcl.Property{
											Type:   "string",
											GoType: "string",
										},
									},
									"issuerUri": &dcl.Property{
										Type:        "string",
										GoName:      "IssuerUri",
										Description: "Required. The OIDC issuer URL. Must be an HTTPS endpoint.",
									},
								},
							},
							"project": &dcl.Property{
								Type:        "string",
								GoName:      "Project",
								Description: "The project for the resource",
								Immutable:   true,
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Cloudresourcemanager/Project",
										Field:    "name",
										Parent:   true,
									},
								},
							},
							"state": &dcl.Property{
								Type:        "string",
								GoName:      "State",
								GoType:      "WorkloadIdentityPoolProviderStateEnum",
								ReadOnly:    true,
								Description: "Output only. The state of the provider. Possible values: STATE_UNSPECIFIED, ACTIVE, DELETED",
								Immutable:   true,
								Enum: []string{
									"STATE_UNSPECIFIED",
									"ACTIVE",
									"DELETED",
								},
							},
							"workloadIdentityPool": &dcl.Property{
								Type:        "string",
								GoName:      "WorkloadIdentityPool",
								Description: "The workloadIdentityPool for the resource",
								Immutable:   true,
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Iam/WorkloadIdentityPool",
										Field:    "name",
										Parent:   true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
