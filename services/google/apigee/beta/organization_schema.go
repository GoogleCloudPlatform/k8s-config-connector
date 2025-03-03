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

func DCLOrganizationSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "Apigee/Organization",
			Description: "The Apigee Organization resource",
			StructName:  "Organization",
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a Organization",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "organization",
						Required:    true,
						Description: "A full instance of a Organization",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a Organization",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "organization",
						Required:    true,
						Description: "A full instance of a Organization",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a Organization",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "organization",
						Required:    true,
						Description: "A full instance of a Organization",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all Organization",
			},
			List: &dcl.Path{
				Description: "The function used to list information about many Organization",
			},
		},
		Components: &dcl.Components{
			Schemas: map[string]*dcl.Component{
				"Organization": &dcl.Component{
					Title:           "Organization",
					ID:              "organizations/{{name}}",
					ParentContainer: "project",
					HasCreate:       true,
					ApplyTimeout:    4800,
					DeleteTimeout:   4800,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"analyticsRegion",
							"runtimeType",
						},
						Properties: map[string]*dcl.Property{
							"addonsConfig": &dcl.Property{
								Type:        "object",
								GoName:      "AddonsConfig",
								GoType:      "OrganizationAddonsConfig",
								Description: "Addon configurations of the Apigee organization.",
								Properties: map[string]*dcl.Property{
									"advancedApiOpsConfig": &dcl.Property{
										Type:        "object",
										GoName:      "AdvancedApiOpsConfig",
										GoType:      "OrganizationAddonsConfigAdvancedApiOpsConfig",
										Description: "Configuration for the Advanced API Ops add-on.",
										Properties: map[string]*dcl.Property{
											"enabled": &dcl.Property{
												Type:        "boolean",
												GoName:      "Enabled",
												Description: "Flag that specifies whether the Advanced API Ops add-on is enabled.",
											},
										},
									},
									"monetizationConfig": &dcl.Property{
										Type:        "object",
										GoName:      "MonetizationConfig",
										GoType:      "OrganizationAddonsConfigMonetizationConfig",
										Description: "Configuration for the Monetization add-on.",
										Properties: map[string]*dcl.Property{
											"enabled": &dcl.Property{
												Type:        "boolean",
												GoName:      "Enabled",
												Description: "Flag that specifies whether the Monetization add-on is enabled.",
											},
										},
									},
								},
							},
							"analyticsRegion": &dcl.Property{
								Type:        "string",
								GoName:      "AnalyticsRegion",
								Description: "Required. Primary GCP region for analytics data storage. For valid values, see (https://cloud.google.com/apigee/docs/api-platform/get-started/create-org).",
								Immutable:   true,
							},
							"authorizedNetwork": &dcl.Property{
								Type:        "string",
								GoName:      "AuthorizedNetwork",
								Description: "Compute Engine network used for Service Networking to be peered with Apigee runtime instances. See (https://cloud.google.com/vpc/docs/shared-vpc). To use a shared VPC network, use the following format: `projects/{host-project-id}/{region}/networks/{network-name}`. For example: `projects/my-sharedvpc-host/global/networks/mynetwork` **Note:** Not supported for Apigee hybrid.",
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Compute/Network",
										Field:    "name",
									},
								},
							},
							"billingType": &dcl.Property{
								Type:        "string",
								GoName:      "BillingType",
								GoType:      "OrganizationBillingTypeEnum",
								ReadOnly:    true,
								Description: "Output only. Billing type of the Apigee organization. See (https://cloud.google.com/apigee/pricing). Possible values: BILLING_TYPE_UNSPECIFIED, SUBSCRIPTION, EVALUATION",
								Immutable:   true,
								Enum: []string{
									"BILLING_TYPE_UNSPECIFIED",
									"SUBSCRIPTION",
									"EVALUATION",
								},
							},
							"caCertificate": &dcl.Property{
								Type:        "string",
								GoName:      "CaCertificate",
								ReadOnly:    true,
								Description: "Output only. Base64-encoded public certificate for the root CA of the Apigee organization. Valid only when (#RuntimeType) is `CLOUD`.",
								Immutable:   true,
							},
							"createdAt": &dcl.Property{
								Type:        "integer",
								Format:      "int64",
								GoName:      "CreatedAt",
								ReadOnly:    true,
								Description: "Output only. Time that the Apigee organization was created in milliseconds since epoch.",
								Immutable:   true,
							},
							"description": &dcl.Property{
								Type:        "string",
								GoName:      "Description",
								Description: "Description of the Apigee organization.",
							},
							"displayName": &dcl.Property{
								Type:        "string",
								GoName:      "DisplayName",
								Description: "Display name for the Apigee organization.",
							},
							"environments": &dcl.Property{
								Type:        "array",
								GoName:      "Environments",
								ReadOnly:    true,
								Description: "Output only. List of environments in the Apigee organization.",
								Immutable:   true,
								ListType:    "list",
								Items: &dcl.Property{
									Type:   "string",
									GoType: "string",
								},
							},
							"expiresAt": &dcl.Property{
								Type:        "integer",
								Format:      "int64",
								GoName:      "ExpiresAt",
								ReadOnly:    true,
								Description: "Output only. Time that the Apigee organization is scheduled for deletion.",
								Immutable:   true,
							},
							"lastModifiedAt": &dcl.Property{
								Type:        "integer",
								Format:      "int64",
								GoName:      "LastModifiedAt",
								ReadOnly:    true,
								Description: "Output only. Time that the Apigee organization was last modified in milliseconds since epoch.",
								Immutable:   true,
							},
							"name": &dcl.Property{
								Type:                     "string",
								GoName:                   "Name",
								Description:              "Output only. Name of the Apigee organization.",
								Immutable:                true,
								ServerGeneratedParameter: true,
							},
							"project": &dcl.Property{
								Type:                "string",
								GoName:              "Project",
								Description:         "Required. Name of the GCP project in which to associate the Apigee organization. Pass the information as a query parameter using the following structure in your request: projects/<project> Authorization requires the following IAM permission on the specified resource parent: apigee.organizations.create",
								Immutable:           true,
								ForwardSlashAllowed: true,
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Cloudresourcemanager/Project",
										Field:    "name",
										Parent:   true,
									},
								},
							},
							"projectId": &dcl.Property{
								Type:        "string",
								GoName:      "ProjectId",
								ReadOnly:    true,
								Description: "Output only. Project ID associated with the Apigee organization.",
								Immutable:   true,
							},
							"properties": &dcl.Property{
								Type: "object",
								AdditionalProperties: &dcl.Property{
									Type: "string",
								},
								GoName:      "Properties",
								Description: "Properties defined in the Apigee organization profile.",
							},
							"runtimeDatabaseEncryptionKeyName": &dcl.Property{
								Type:        "string",
								GoName:      "RuntimeDatabaseEncryptionKeyName",
								Description: "Cloud KMS key name used for encrypting the data that is stored and replicated across runtime instances. Update is not allowed after the organization is created. Required when (#RuntimeType) is `TRIAL`, a Google-Managed encryption key will be used. For example: \"projects/foo/locations/us/keyRings/bar/cryptoKeys/baz\". **Note:** Not supported for Apigee hybrid.",
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Cloudkms/CryptoKey",
										Field:    "name",
									},
								},
							},
							"runtimeType": &dcl.Property{
								Type:        "string",
								GoName:      "RuntimeType",
								GoType:      "OrganizationRuntimeTypeEnum",
								Description: "Required. Runtime type of the Apigee organization based on the Apigee subscription purchased. Possible values: RUNTIME_TYPE_UNSPECIFIED, CLOUD, HYBRID",
								Immutable:   true,
								Enum: []string{
									"RUNTIME_TYPE_UNSPECIFIED",
									"CLOUD",
									"HYBRID",
								},
							},
							"state": &dcl.Property{
								Type:        "string",
								GoName:      "State",
								GoType:      "OrganizationStateEnum",
								ReadOnly:    true,
								Description: "Output only. State of the organization. Values other than ACTIVE means the resource is not ready to use. Possible values: SNAPSHOT_STATE_UNSPECIFIED, MISSING, OK_DOCSTORE, OK_SUBMITTED, OK_EXTERNAL, DELETED",
								Immutable:   true,
								Enum: []string{
									"SNAPSHOT_STATE_UNSPECIFIED",
									"MISSING",
									"OK_DOCSTORE",
									"OK_SUBMITTED",
									"OK_EXTERNAL",
									"DELETED",
								},
							},
							"subscriptionType": &dcl.Property{
								Type:        "string",
								GoName:      "SubscriptionType",
								GoType:      "OrganizationSubscriptionTypeEnum",
								ReadOnly:    true,
								Description: "Output only. DEPRECATED: This will eventually be replaced by BillingType. Subscription type of the Apigee organization. Valid values include trial (free, limited, and for evaluation purposes only) or paid (full subscription has been purchased). See (https://cloud.google.com/apigee/pricing/). Possible values: SUBSCRIPTION_TYPE_UNSPECIFIED, PAID, TRIAL",
								Immutable:   true,
								Enum: []string{
									"SUBSCRIPTION_TYPE_UNSPECIFIED",
									"PAID",
									"TRIAL",
								},
							},
						},
					},
				},
			},
		},
	}
}
