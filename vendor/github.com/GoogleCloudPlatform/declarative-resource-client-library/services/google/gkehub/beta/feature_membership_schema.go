// Copyright 2022 Google LLC. All Rights Reserved.
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

func DCLFeatureMembershipSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "GkeHub/FeatureMembership",
			Description: "The GkeHub FeatureMembership resource",
			StructName:  "FeatureMembership",
			Mutex:       "{{project}}/{{location}}/{{feature}}",
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a FeatureMembership",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "FeatureMembership",
						Required:    true,
						Description: "A full instance of a FeatureMembership",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a FeatureMembership",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "FeatureMembership",
						Required:    true,
						Description: "A full instance of a FeatureMembership",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a FeatureMembership",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "FeatureMembership",
						Required:    true,
						Description: "A full instance of a FeatureMembership",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all FeatureMembership",
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
						Name:     "feature",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
				},
			},
			List: &dcl.Path{
				Description: "The function used to list information about many FeatureMembership",
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
						Name:     "feature",
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
				"FeatureMembership": &dcl.Component{
					Title:           "FeatureMembership",
					ID:              "projects/{{project}}/locations/{{location}}/features/{{feature}}/memberships/{{membership}}",
					ParentContainer: "project",
					HasCreate:       true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"configmanagement",
							"project",
							"location",
							"feature",
							"membership",
						},
						Properties: map[string]*dcl.Property{
							"configmanagement": &dcl.Property{
								Type:        "object",
								GoName:      "Configmanagement",
								GoType:      "FeatureMembershipConfigmanagement",
								Description: "Config Management-specific spec.",
								Properties: map[string]*dcl.Property{
									"binauthz": &dcl.Property{
										Type:        "object",
										GoName:      "Binauthz",
										GoType:      "FeatureMembershipConfigmanagementBinauthz",
										Description: "Binauthz configuration for the cluster.",
										Properties: map[string]*dcl.Property{
											"enabled": &dcl.Property{
												Type:        "boolean",
												GoName:      "Enabled",
												Description: "Whether binauthz is enabled in this cluster.",
											},
										},
									},
									"configSync": &dcl.Property{
										Type:        "object",
										GoName:      "ConfigSync",
										GoType:      "FeatureMembershipConfigmanagementConfigSync",
										Description: "Config Sync configuration for the cluster.",
										SendEmpty:   true,
										Properties: map[string]*dcl.Property{
											"git": &dcl.Property{
												Type:   "object",
												GoName: "Git",
												GoType: "FeatureMembershipConfigmanagementConfigSyncGit",
												Properties: map[string]*dcl.Property{
													"gcpServiceAccountEmail": &dcl.Property{
														Type:        "string",
														GoName:      "GcpServiceAccountEmail",
														Description: "The GCP Service Account Email used for auth when secretType is gcpServiceAccount.",
														ResourceReferences: []*dcl.PropertyResourceReference{
															&dcl.PropertyResourceReference{
																Resource: "Iam/ServiceAccount",
																Field:    "email",
															},
														},
													},
													"httpsProxy": &dcl.Property{
														Type:        "string",
														GoName:      "HttpsProxy",
														Description: "URL for the HTTPS proxy to be used when communicating with the Git repo.",
													},
													"policyDir": &dcl.Property{
														Type:        "string",
														GoName:      "PolicyDir",
														Description: "The path within the Git repository that represents the top level of the repo to sync. Default: the root directory of the repository.",
													},
													"secretType": &dcl.Property{
														Type:        "string",
														GoName:      "SecretType",
														Description: "Type of secret configured for access to the Git repo. Must be one of ssh, cookiefile, gcenode, token, gcpserviceaccount or none. The validation of this is case-sensitive.",
													},
													"syncBranch": &dcl.Property{
														Type:        "string",
														GoName:      "SyncBranch",
														Description: "The branch of the repository to sync from. Default: master.",
													},
													"syncRepo": &dcl.Property{
														Type:        "string",
														GoName:      "SyncRepo",
														Description: "The URL of the Git repository to use as the source of truth.",
													},
													"syncRev": &dcl.Property{
														Type:        "string",
														GoName:      "SyncRev",
														Description: "Git revision (tag or hash) to check out. Default HEAD.",
													},
													"syncWaitSecs": &dcl.Property{
														Type:        "string",
														GoName:      "SyncWaitSecs",
														Description: "Period in seconds between consecutive syncs. Default: 15.",
													},
												},
											},
											"preventDrift": &dcl.Property{
												Type:          "boolean",
												GoName:        "PreventDrift",
												Description:   "Set to true to enable the Config Sync admission webhook to prevent drifts. If set to `false`, disables the Config Sync admission webhook and does not prevent drifts.",
												ServerDefault: true,
											},
											"sourceFormat": &dcl.Property{
												Type:        "string",
												GoName:      "SourceFormat",
												Description: "Specifies whether the Config Sync Repo is in \"hierarchical\" or \"unstructured\" mode.",
											},
										},
									},
									"hierarchyController": &dcl.Property{
										Type:        "object",
										GoName:      "HierarchyController",
										GoType:      "FeatureMembershipConfigmanagementHierarchyController",
										Description: "Hierarchy Controller configuration for the cluster.",
										Properties: map[string]*dcl.Property{
											"enableHierarchicalResourceQuota": &dcl.Property{
												Type:        "boolean",
												GoName:      "EnableHierarchicalResourceQuota",
												Description: "Whether hierarchical resource quota is enabled in this cluster.",
											},
											"enablePodTreeLabels": &dcl.Property{
												Type:        "boolean",
												GoName:      "EnablePodTreeLabels",
												Description: "Whether pod tree labels are enabled in this cluster.",
											},
											"enabled": &dcl.Property{
												Type:        "boolean",
												GoName:      "Enabled",
												Description: "Whether Hierarchy Controller is enabled in this cluster.",
											},
										},
									},
									"policyController": &dcl.Property{
										Type:        "object",
										GoName:      "PolicyController",
										GoType:      "FeatureMembershipConfigmanagementPolicyController",
										Description: "Policy Controller configuration for the cluster.",
										Properties: map[string]*dcl.Property{
											"auditIntervalSeconds": &dcl.Property{
												Type:        "string",
												GoName:      "AuditIntervalSeconds",
												Description: "Sets the interval for Policy Controller Audit Scans (in seconds). When set to 0, this disables audit functionality altogether.",
											},
											"enabled": &dcl.Property{
												Type:        "boolean",
												GoName:      "Enabled",
												Description: "Enables the installation of Policy Controller. If false, the rest of PolicyController fields take no effect.",
											},
											"exemptableNamespaces": &dcl.Property{
												Type:        "array",
												GoName:      "ExemptableNamespaces",
												Description: "The set of namespaces that are excluded from Policy Controller checks. Namespaces do not need to currently exist on the cluster.",
												SendEmpty:   true,
												ListType:    "list",
												Items: &dcl.Property{
													Type:   "string",
													GoType: "string",
												},
											},
											"logDeniesEnabled": &dcl.Property{
												Type:        "boolean",
												GoName:      "LogDeniesEnabled",
												Description: "Logs all denies and dry run failures.",
											},
											"monitoring": &dcl.Property{
												Type:          "object",
												GoName:        "Monitoring",
												GoType:        "FeatureMembershipConfigmanagementPolicyControllerMonitoring",
												Description:   "Specifies the backends Policy Controller should export metrics to. For example, to specify metrics should be exported to Cloud Monitoring and Prometheus, specify backends: [\"cloudmonitoring\", \"prometheus\"]. Default: [\"cloudmonitoring\", \"prometheus\"]",
												ServerDefault: true,
												Properties: map[string]*dcl.Property{
													"backends": &dcl.Property{
														Type:          "array",
														GoName:        "Backends",
														Description:   " Specifies the list of backends Policy Controller will export to. Specifying an empty value `[]` disables metrics export.",
														ServerDefault: true,
														SendEmpty:     true,
														ListType:      "list",
														Items: &dcl.Property{
															Type:   "string",
															GoType: "FeatureMembershipConfigmanagementPolicyControllerMonitoringBackendsEnum",
															Enum: []string{
																"MONITORING_BACKEND_UNSPECIFIED",
																"PROMETHEUS",
																"CLOUD_MONITORING",
															},
														},
													},
												},
											},
											"mutationEnabled": &dcl.Property{
												Type:        "boolean",
												GoName:      "MutationEnabled",
												Description: "Enable or disable mutation in policy controller. If true, mutation CRDs, webhook and controller deployment will be deployed to the cluster.",
											},
											"referentialRulesEnabled": &dcl.Property{
												Type:        "boolean",
												GoName:      "ReferentialRulesEnabled",
												Description: "Enables the ability to use Constraint Templates that reference to objects other than the object currently being evaluated.",
											},
											"templateLibraryInstalled": &dcl.Property{
												Type:        "boolean",
												GoName:      "TemplateLibraryInstalled",
												Description: "Installs the default template library along with Policy Controller.",
											},
										},
									},
									"version": &dcl.Property{
										Type:          "string",
										GoName:        "Version",
										Description:   "Optional. Version of ACM to install. Defaults to the latest version.",
										ServerDefault: true,
									},
								},
							},
							"feature": &dcl.Property{
								Type:        "string",
								GoName:      "Feature",
								Description: "The name of the feature",
								Immutable:   true,
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Gkehub/Feature",
										Field:    "name",
										Parent:   true,
									},
								},
							},
							"location": &dcl.Property{
								Type:        "string",
								GoName:      "Location",
								Description: "The location of the feature",
								Immutable:   true,
							},
							"membership": &dcl.Property{
								Type:        "string",
								GoName:      "Membership",
								Description: "The name of the membership",
								Immutable:   true,
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Gkehub/Membership",
										Field:    "name",
									},
								},
							},
							"project": &dcl.Property{
								Type:        "string",
								GoName:      "Project",
								Description: "The project of the feature",
								Immutable:   true,
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Cloudresourcemanager/Project",
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
