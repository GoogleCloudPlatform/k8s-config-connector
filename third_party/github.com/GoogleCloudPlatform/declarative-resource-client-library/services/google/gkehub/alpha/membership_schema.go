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
package alpha

import (
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func DCLMembershipSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "GkeHub/Membership",
			Description: "The GkeHub Membership resource",
			StructName:  "Membership",
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a Membership",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "membership",
						Required:    true,
						Description: "A full instance of a Membership",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a Membership",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "membership",
						Required:    true,
						Description: "A full instance of a Membership",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a Membership",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "membership",
						Required:    true,
						Description: "A full instance of a Membership",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all Membership",
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
				},
			},
			List: &dcl.Path{
				Description: "The function used to list information about many Membership",
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
				},
			},
		},
		Components: &dcl.Components{
			Schemas: map[string]*dcl.Component{
				"Membership": &dcl.Component{
					Title:           "Membership",
					ID:              "projects/{{project}}/locations/{{location}}/memberships/{{name}}",
					UsesStateHint:   true,
					ParentContainer: "project",
					LabelsField:     "labels",
					HasCreate:       true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"name",
							"project",
							"location",
						},
						Properties: map[string]*dcl.Property{
							"authority": &dcl.Property{
								Type:        "object",
								GoName:      "Authority",
								GoType:      "MembershipAuthority",
								Description: "Optional. How to identify workloads from this Membership. See the documentation on Workload Identity for more details: https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity",
								Properties: map[string]*dcl.Property{
									"identityProvider": &dcl.Property{
										Type:        "string",
										GoName:      "IdentityProvider",
										ReadOnly:    true,
										Description: "Output only. An identity provider that reflects the `issuer` in the workload identity pool.",
									},
									"issuer": &dcl.Property{
										Type:        "string",
										GoName:      "Issuer",
										Description: "Optional. A JSON Web Token (JWT) issuer URI. `issuer` must start with `https://` and be a valid URL with length <2000 characters. If set, then Google will allow valid OIDC tokens from this issuer to authenticate within the workload_identity_pool. OIDC discovery will be performed on this URI to validate tokens from the issuer. Clearing `issuer` disables Workload Identity. `issuer` cannot be directly modified; it must be cleared (and Workload Identity disabled) before using a new issuer (and re-enabling Workload Identity).",
									},
									"workloadIdentityPool": &dcl.Property{
										Type:        "string",
										GoName:      "WorkloadIdentityPool",
										ReadOnly:    true,
										Description: "Output only. The name of the workload identity pool in which `issuer` will be recognized. There is a single Workload Identity Pool per Hub that is shared between all Memberships that belong to that Hub. For a Hub hosted in: {PROJECT_ID}, the workload pool format is `{PROJECT_ID}.hub.id.goog`, although this is subject to change in newer versions of this API.",
									},
								},
							},
							"createTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "CreateTime",
								ReadOnly:    true,
								Description: "Output only. When the Membership was created.",
								Immutable:   true,
							},
							"deleteTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "DeleteTime",
								ReadOnly:    true,
								Description: "Output only. When the Membership was deleted.",
								Immutable:   true,
							},
							"description": &dcl.Property{
								Type:        "string",
								GoName:      "Description",
								Description: "Description of this membership, limited to 63 characters. Must match the regex: `*` This field is present for legacy purposes.",
							},
							"endpoint": &dcl.Property{
								Type:        "object",
								GoName:      "Endpoint",
								GoType:      "MembershipEndpoint",
								Description: "Optional. Endpoint information to reach this member.",
								Properties: map[string]*dcl.Property{
									"gkeCluster": &dcl.Property{
										Type:        "object",
										GoName:      "GkeCluster",
										GoType:      "MembershipEndpointGkeCluster",
										Description: "Optional. GKE-specific information. Only present if this Membership is a GKE cluster.",
										Properties: map[string]*dcl.Property{
											"resourceLink": &dcl.Property{
												Type:        "string",
												GoName:      "ResourceLink",
												Description: "Immutable. Self-link of the GCP resource for the GKE cluster. For example: //container.googleapis.com/projects/my-project/locations/us-west1-a/clusters/my-cluster Zonal clusters are also supported.",
												ResourceReferences: []*dcl.PropertyResourceReference{
													&dcl.PropertyResourceReference{
														Resource: "Container/Cluster",
														Field:    "selfLink",
													},
												},
											},
										},
									},
									"kubernetesMetadata": &dcl.Property{
										Type:        "object",
										GoName:      "KubernetesMetadata",
										GoType:      "MembershipEndpointKubernetesMetadata",
										ReadOnly:    true,
										Description: "Output only. Useful Kubernetes-specific metadata.",
										Properties: map[string]*dcl.Property{
											"kubernetesApiServerVersion": &dcl.Property{
												Type:        "string",
												GoName:      "KubernetesApiServerVersion",
												ReadOnly:    true,
												Description: "Output only. Kubernetes API server version string as reported by `/version`.",
											},
											"memoryMb": &dcl.Property{
												Type:        "integer",
												Format:      "int64",
												GoName:      "MemoryMb",
												ReadOnly:    true,
												Description: "Output only. The total memory capacity as reported by the sum of all Kubernetes nodes resources, defined in MB.",
											},
											"nodeCount": &dcl.Property{
												Type:        "integer",
												Format:      "int64",
												GoName:      "NodeCount",
												ReadOnly:    true,
												Description: "Output only. Node count as reported by Kubernetes nodes resources.",
											},
											"nodeProviderId": &dcl.Property{
												Type:        "string",
												GoName:      "NodeProviderId",
												ReadOnly:    true,
												Description: "Output only. Node providerID as reported by the first node in the list of nodes on the Kubernetes endpoint. On Kubernetes platforms that support zero-node clusters (like GKE-on-GCP), the node_count will be zero and the node_provider_id will be empty.",
											},
											"updateTime": &dcl.Property{
												Type:        "string",
												Format:      "date-time",
												GoName:      "UpdateTime",
												ReadOnly:    true,
												Description: "Output only. The time at which these details were last updated. This update_time is different from the Membership-level update_time since EndpointDetails are updated internally for API consumers.",
											},
											"vcpuCount": &dcl.Property{
												Type:        "integer",
												Format:      "int64",
												GoName:      "VcpuCount",
												ReadOnly:    true,
												Description: "Output only. vCPU count as reported by Kubernetes nodes resources.",
											},
										},
									},
									"kubernetesResource": &dcl.Property{
										Type:        "object",
										GoName:      "KubernetesResource",
										GoType:      "MembershipEndpointKubernetesResource",
										Description: "Optional. The in-cluster Kubernetes Resources that should be applied for a correctly registered cluster, in the steady state. These resources: * Ensure that the cluster is exclusively registered to one and only one Hub Membership. * Propagate Workload Pool Information available in the Membership Authority field. * Ensure proper initial configuration of default Hub Features.",
										Properties: map[string]*dcl.Property{
											"connectResources": &dcl.Property{
												Type:        "array",
												GoName:      "ConnectResources",
												ReadOnly:    true,
												Description: "Output only. The Kubernetes resources for installing the GKE Connect agent This field is only populated in the Membership returned from a successful long-running operation from CreateMembership or UpdateMembership. It is not populated during normal GetMembership or ListMemberships requests. To get the resource manifest after the initial registration, the caller should make a UpdateMembership call with an empty field mask.",
												ListType:    "list",
												Items: &dcl.Property{
													Type:   "object",
													GoType: "MembershipEndpointKubernetesResourceConnectResources",
													Properties: map[string]*dcl.Property{
														"clusterScoped": &dcl.Property{
															Type:        "boolean",
															GoName:      "ClusterScoped",
															Description: "Whether the resource provided in the manifest is `cluster_scoped`. If unset, the manifest is assumed to be namespace scoped. This field is used for REST mapping when applying the resource in a cluster.",
														},
														"manifest": &dcl.Property{
															Type:        "string",
															GoName:      "Manifest",
															Description: "YAML manifest of the resource.",
														},
													},
												},
											},
											"membershipCrManifest": &dcl.Property{
												Type:        "string",
												GoName:      "MembershipCrManifest",
												Description: "Input only. The YAML representation of the Membership CR. This field is ignored for GKE clusters where Hub can read the CR directly. Callers should provide the CR that is currently present in the cluster during CreateMembership or UpdateMembership, or leave this field empty if none exists. The CR manifest is used to validate the cluster has not been registered with another Membership.",
												Unreadable:  true,
											},
											"membershipResources": &dcl.Property{
												Type:        "array",
												GoName:      "MembershipResources",
												ReadOnly:    true,
												Description: "Output only. Additional Kubernetes resources that need to be applied to the cluster after Membership creation, and after every update. This field is only populated in the Membership returned from a successful long-running operation from CreateMembership or UpdateMembership. It is not populated during normal GetMembership or ListMemberships requests. To get the resource manifest after the initial registration, the caller should make a UpdateMembership call with an empty field mask.",
												ListType:    "list",
												Items: &dcl.Property{
													Type:   "object",
													GoType: "MembershipEndpointKubernetesResourceMembershipResources",
													Properties: map[string]*dcl.Property{
														"clusterScoped": &dcl.Property{
															Type:        "boolean",
															GoName:      "ClusterScoped",
															Description: "Whether the resource provided in the manifest is `cluster_scoped`. If unset, the manifest is assumed to be namespace scoped. This field is used for REST mapping when applying the resource in a cluster.",
														},
														"manifest": &dcl.Property{
															Type:        "string",
															GoName:      "Manifest",
															Description: "YAML manifest of the resource.",
														},
													},
												},
											},
											"resourceOptions": &dcl.Property{
												Type:        "object",
												GoName:      "ResourceOptions",
												GoType:      "MembershipEndpointKubernetesResourceResourceOptions",
												Description: "Optional. Options for Kubernetes resource generation.",
												Properties: map[string]*dcl.Property{
													"connectVersion": &dcl.Property{
														Type:        "string",
														GoName:      "ConnectVersion",
														Description: "Optional. The Connect agent version to use for connect_resources. Defaults to the latest GKE Connect version. The version must be a currently supported version, obsolete versions will be rejected.",
													},
													"v1beta1Crd": &dcl.Property{
														Type:        "boolean",
														GoName:      "V1Beta1Crd",
														Description: "Optional. Use `apiextensions/v1beta1` instead of `apiextensions/v1` for CustomResourceDefinition resources. This option should be set for clusters with Kubernetes apiserver versions <1.16.",
													},
												},
											},
										},
									},
								},
							},
							"externalId": &dcl.Property{
								Type:          "string",
								GoName:        "ExternalId",
								Description:   "Optional. An externally-generated and managed ID for this Membership. This ID may be modified after creation, but this is not recommended. The ID must match the regex: `*` If this Membership represents a Kubernetes cluster, this value should be set to the UID of the `kube-system` namespace object.",
								ServerDefault: true,
							},
							"infrastructureType": &dcl.Property{
								Type:          "string",
								GoName:        "InfrastructureType",
								GoType:        "MembershipInfrastructureTypeEnum",
								Description:   "Optional. The infrastructure type this Membership is running on. Possible values: INFRASTRUCTURE_TYPE_UNSPECIFIED, ON_PREM, MULTI_CLOUD",
								ServerDefault: true,
								Enum: []string{
									"INFRASTRUCTURE_TYPE_UNSPECIFIED",
									"ON_PREM",
									"MULTI_CLOUD",
								},
							},
							"labels": &dcl.Property{
								Type: "object",
								AdditionalProperties: &dcl.Property{
									Type: "string",
								},
								GoName:      "Labels",
								Description: "Optional. GCP labels for this membership.",
							},
							"lastConnectionTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "LastConnectionTime",
								ReadOnly:    true,
								Description: "Output only. For clusters using Connect, the timestamp of the most recent connection established with Google Cloud. This time is updated every several minutes, not continuously. For clusters that do not use GKE Connect, or that have never connected successfully, this field will be unset.",
								Immutable:   true,
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
								Description: "Output only. The full, unique name of this Membership resource in the format `projects/*/locations/*/memberships/{membership_id}`, set during creation. `membership_id` must be a valid RFC 1123 compliant DNS label: 1. At most 63 characters in length 2. It must consist of lower case alphanumeric characters or `-` 3. It must start and end with an alphanumeric character Which can be expressed as the regex: `)?`, with a maximum length of 63 characters.",
								Immutable:   true,
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
								Type:        "object",
								GoName:      "State",
								GoType:      "MembershipState",
								ReadOnly:    true,
								Description: "Output only. State of the Membership resource.",
								Immutable:   true,
								Properties: map[string]*dcl.Property{
									"code": &dcl.Property{
										Type:        "string",
										GoName:      "Code",
										GoType:      "MembershipStateCodeEnum",
										ReadOnly:    true,
										Description: "Output only. The current state of the Membership resource. Possible values: CODE_UNSPECIFIED, CREATING, READY, DELETING, UPDATING, SERVICE_UPDATING",
										Immutable:   true,
										Enum: []string{
											"CODE_UNSPECIFIED",
											"CREATING",
											"READY",
											"DELETING",
											"UPDATING",
											"SERVICE_UPDATING",
										},
									},
								},
							},
							"uniqueId": &dcl.Property{
								Type:        "string",
								GoName:      "UniqueId",
								ReadOnly:    true,
								Description: "Output only. Google-generated UUID for this resource. This is unique across all Membership resources. If a Membership resource is deleted and another resource with the same name is created, it gets a different unique_id.",
								Immutable:   true,
							},
							"updateTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "UpdateTime",
								ReadOnly:    true,
								Description: "Output only. When the Membership was last updated.",
								Immutable:   true,
							},
						},
					},
				},
			},
		},
	}
}
