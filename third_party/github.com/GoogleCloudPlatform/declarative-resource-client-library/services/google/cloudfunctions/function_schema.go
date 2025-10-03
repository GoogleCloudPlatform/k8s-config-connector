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
package cloudfunctions

import (
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func DCLFunctionSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "CloudFunctions/Function",
			Description: "The CloudFunctions Function resource",
			StructName:  "Function",
			HasIAM:      true,
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a Function",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "function",
						Required:    true,
						Description: "A full instance of a Function",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a Function",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "function",
						Required:    true,
						Description: "A full instance of a Function",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a Function",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "function",
						Required:    true,
						Description: "A full instance of a Function",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all Function",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:     "project",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
					dcl.PathParameters{
						Name:     "region",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
				},
			},
			List: &dcl.Path{
				Description: "The function used to list information about many Function",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:     "project",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
					dcl.PathParameters{
						Name:     "region",
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
				"Function": &dcl.Component{
					Title: "Function",
					ID:    "projects/{{project}}/locations/{{region}}/functions/{{name}}",
					Locations: []string{
						"region",
					},
					ParentContainer: "project",
					LabelsField:     "labels",
					HasCreate:       true,
					HasIAM:          true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"name",
							"runtime",
							"region",
							"project",
						},
						Properties: map[string]*dcl.Property{
							"availableMemoryMb": &dcl.Property{
								Type:          "integer",
								Format:        "int64",
								GoName:        "AvailableMemoryMb",
								Description:   "Memory (in MB), available to the function. Default value is 256MB. Allowed values are: 128MB, 256MB, 512MB, 1024MB, and 2048MB.",
								ServerDefault: true,
							},
							"description": &dcl.Property{
								Type:        "string",
								GoName:      "Description",
								Description: "User-provided description of a function.",
							},
							"entryPoint": &dcl.Property{
								Type:        "string",
								GoName:      "EntryPoint",
								Description: "The name of the function (as defined in source code) that will be\nexecuted. Defaults to the resource name suffix, if not specified. For\nbackward compatibility, if function with given name is not found, then the\nsystem will try to use function named \"function\".\nFor Node.js this is name of a function exported by the module specified\nin `source_location`.",
								Immutable:   true,
							},
							"environmentVariables": &dcl.Property{
								Type: "object",
								AdditionalProperties: &dcl.Property{
									Type: "string",
								},
								GoName:      "EnvironmentVariables",
								Description: "Environment variables that shall be available during function execution.",
							},
							"eventTrigger": &dcl.Property{
								Type:        "object",
								GoName:      "EventTrigger",
								GoType:      "FunctionEventTrigger",
								Description: "A source that fires events in response to a condition in another service.",
								Immutable:   true,
								Conflicts: []string{
									"httpsTrigger",
								},
								Required: []string{
									"eventType",
									"resource",
								},
								Properties: map[string]*dcl.Property{
									"eventType": &dcl.Property{
										Type:        "string",
										GoName:      "EventType",
										Description: "Required. The type of event to observe. For example:\n`providers/cloud.storage/eventTypes/object.change` and\n`providers/cloud.pubsub/eventTypes/topic.publish`.\n\nEvent types match pattern `providers/*/eventTypes/*.*`.\nThe pattern contains:\n\n1. namespace: For example, `cloud.storage` and\n   `google.firebase.analytics`.\n2. resource type: The type of resource on which event occurs. For\n   example, the Google Cloud Storage API includes the type `object`.\n3. action: The action that generates the event. For example, action for\n   a Google Cloud Storage Object is 'change'.\nThese parts are lower case.",
										Immutable:   true,
									},
									"failurePolicy": &dcl.Property{
										Type:        "boolean",
										GoName:      "FailurePolicy",
										Description: "Specifies policy for failed executions.",
										Immutable:   true,
									},
									"resource": &dcl.Property{
										Type:        "string",
										GoName:      "Resource",
										Description: "Required. The resource(s) from which to observe events, for example,\n`projects/_/buckets/myBucket`.\n\nNot all syntactically correct values are accepted by all services. For\nexample:\n\n1. The authorization model must support it. Google Cloud Functions\n   only allows EventTriggers to be deployed that observe resources in the\n   same project as the `Function`.\n2. The resource type must match the pattern expected for an\n   `event_type`. For example, an `EventTrigger` that has an\n   `event_type` of \"google.pubsub.topic.publish\" should have a resource\n   that matches Google Cloud Pub/Sub topics.\n\nAdditionally, some services may support short names when creating an\n`EventTrigger`. These will always be returned in the normalized \"long\"\nformat.\n\nSee each *service's* documentation for supported formats.",
										Immutable:   true,
										ResourceReferences: []*dcl.PropertyResourceReference{
											&dcl.PropertyResourceReference{
												Resource: "Storage/Bucket",
												Field:    "name",
												Format:   "projects/{{project}}/buckets/{{name}}",
											},
											&dcl.PropertyResourceReference{
												Resource: "Pubsub/Topic",
												Field:    "name",
											},
										},
									},
									"service": &dcl.Property{
										Type:        "string",
										GoName:      "Service",
										Description: "The hostname of the service that should be observed.\n\nIf no string is provided, the default service implementing the API will\nbe used. For example, `storage.googleapis.com` is the default for all\nevent types in the `google.storage` namespace.\n",
										Immutable:   true,
									},
								},
							},
							"httpsTrigger": &dcl.Property{
								Type:        "object",
								GoName:      "HttpsTrigger",
								GoType:      "FunctionHttpsTrigger",
								Description: "An HTTPS endpoint type of source that can be triggered via URL.",
								Immutable:   true,
								Conflicts: []string{
									"eventTrigger",
								},
								Properties: map[string]*dcl.Property{
									"securityLevel": &dcl.Property{
										Type:        "string",
										GoName:      "SecurityLevel",
										GoType:      "FunctionHttpsTriggerSecurityLevelEnum",
										Description: "Both HTTP and HTTPS requests with URLs that match the handler succeed without redirects. The application can examine the request to determine which protocol was used and respond accordingly. Possible values: SECURITY_LEVEL_UNSPECIFIED, SECURE_ALWAYS, SECURE_OPTIONAL",
										Immutable:   true,
										Enum: []string{
											"SECURITY_LEVEL_UNSPECIFIED",
											"SECURE_ALWAYS",
											"SECURE_OPTIONAL",
										},
									},
									"url": &dcl.Property{
										Type:        "string",
										GoName:      "Url",
										ReadOnly:    true,
										Description: "Output only. The deployed url for the function.",
										Immutable:   true,
									},
								},
							},
							"ingressSettings": &dcl.Property{
								Type:          "string",
								GoName:        "IngressSettings",
								GoType:        "FunctionIngressSettingsEnum",
								Description:   "The ingress settings for the function, controlling what traffic can reach\nit. Possible values: INGRESS_SETTINGS_UNSPECIFIED, ALLOW_ALL, ALLOW_INTERNAL_ONLY, ALLOW_INTERNAL_AND_GCLB",
								ServerDefault: true,
								Enum: []string{
									"INGRESS_SETTINGS_UNSPECIFIED",
									"ALLOW_ALL",
									"ALLOW_INTERNAL_ONLY",
									"ALLOW_INTERNAL_AND_GCLB",
								},
							},
							"labels": &dcl.Property{
								Type: "object",
								AdditionalProperties: &dcl.Property{
									Type: "string",
								},
								GoName:      "Labels",
								Description: "Labels associated with this Cloud Function.",
							},
							"maxInstances": &dcl.Property{
								Type:        "integer",
								Format:      "int64",
								GoName:      "MaxInstances",
								Description: "The limit on the maximum number of function instances that may coexist at a\ngiven time.",
							},
							"name": &dcl.Property{
								Type:        "string",
								GoName:      "Name",
								Description: "A user-defined name of the function. Function names must be unique globally.",
								Immutable:   true,
							},
							"project": &dcl.Property{
								Type:        "string",
								GoName:      "Project",
								Description: "The project id of the function.",
								Immutable:   true,
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Cloudresourcemanager/Project",
										Field:    "name",
										Parent:   true,
									},
								},
							},
							"region": &dcl.Property{
								Type:        "string",
								GoName:      "Region",
								Description: "The name of the Cloud Functions region of the function.",
								Immutable:   true,
							},
							"runtime": &dcl.Property{
								Type:        "string",
								GoName:      "Runtime",
								Description: "The runtime in which to run the function. Required when deploying a new\nfunction, optional when updating an existing function. For a complete\nlist of possible choices, see the\n[`gcloud` command\nreference](/sdk/gcloud/reference/functions/deploy#--runtime).\n",
							},
							"serviceAccountEmail": &dcl.Property{
								Type:          "string",
								GoName:        "ServiceAccountEmail",
								Description:   "The email of the function's service account. If empty, defaults to\n`{project_id}@appspot.gserviceaccount.com`.",
								Immutable:     true,
								ServerDefault: true,
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Iam/ServiceAccount",
										Field:    "email",
									},
								},
							},
							"sourceArchiveUrl": &dcl.Property{
								Type:        "string",
								GoName:      "SourceArchiveUrl",
								Description: "The Google Cloud Storage URL, starting with gs://, pointing to the zip archive which contains the function.",
								Immutable:   true,
								Conflicts: []string{
									"sourceRepository",
								},
							},
							"sourceRepository": &dcl.Property{
								Type:        "object",
								GoName:      "SourceRepository",
								GoType:      "FunctionSourceRepository",
								Description: "Represents parameters related to source repository where a function is hosted.",
								Immutable:   true,
								Conflicts: []string{
									"sourceArchiveUrl",
								},
								Required: []string{
									"url",
								},
								Properties: map[string]*dcl.Property{
									"deployedUrl": &dcl.Property{
										Type:        "string",
										GoName:      "DeployedUrl",
										ReadOnly:    true,
										Description: "Output only. The URL pointing to the hosted repository where the function\nwere defined at the time of deployment. It always points to a specific\ncommit in the format described above.",
										Immutable:   true,
									},
									"url": &dcl.Property{
										Type:        "string",
										GoName:      "Url",
										Description: "The URL pointing to the hosted repository where the function is defined.\nThere are supported Cloud Source Repository URLs in the following\nformats:\n\nTo refer to a specific commit:\n`https://source.developers.google.com/projects/*/repos/*/revisions/*/paths/*`\nTo refer to a moveable alias (branch):\n`https://source.developers.google.com/projects/*/repos/*/moveable-aliases/*/paths/*`\nIn particular, to refer to HEAD use `master` moveable alias.\nTo refer to a specific fixed alias (tag):\n`https://source.developers.google.com/projects/*/repos/*/fixed-aliases/*/paths/*`\n\nYou may omit `paths/*` if you want to use the main directory.",
										Immutable:   true,
									},
								},
							},
							"status": &dcl.Property{
								Type:        "string",
								GoName:      "Status",
								GoType:      "FunctionStatusEnum",
								ReadOnly:    true,
								Description: "Output only. Status of the function deployment. Possible values: CLOUD_FUNCTION_STATUS_UNSPECIFIED, ACTIVE, OFFLINE, DEPLOY_IN_PROGRESS, DELETE_IN_PROGRESS, UNKNOWN",
								Immutable:   true,
								Enum: []string{
									"CLOUD_FUNCTION_STATUS_UNSPECIFIED",
									"ACTIVE",
									"OFFLINE",
									"DEPLOY_IN_PROGRESS",
									"DELETE_IN_PROGRESS",
									"UNKNOWN",
								},
							},
							"timeout": &dcl.Property{
								Type:          "string",
								GoName:        "Timeout",
								Description:   "The function execution timeout. Execution is considered failed and\ncan be terminated if the function is not completed at the end of the\ntimeout period. Defaults to 60 seconds.",
								ServerDefault: true,
							},
							"updateTime": &dcl.Property{
								Type:        "string",
								GoName:      "UpdateTime",
								ReadOnly:    true,
								Description: "Output only. The last update timestamp of a Cloud Function in RFC3339 UTC 'Zulu' format, with nanosecond resolution and up to nine fractional digits.",
								Immutable:   true,
							},
							"versionId": &dcl.Property{
								Type:        "integer",
								Format:      "int64",
								GoName:      "VersionId",
								ReadOnly:    true,
								Description: "Output only. The version identifier of the Cloud Function. Each deployment attempt\nresults in a new version of a function being created.",
								Immutable:   true,
							},
							"vpcConnector": &dcl.Property{
								Type:          "string",
								GoName:        "VPCConnector",
								Description:   "The VPC Network Connector that this cloud function can connect to. It can\nbe either the fully-qualified URI, or the short name of the network\nconnector resource. The format of this field is\n`projects/*/locations/*/connectors/*`",
								ServerDefault: true,
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Vpcaccess/Connector",
										Field:    "name",
									},
								},
							},
							"vpcConnectorEgressSettings": &dcl.Property{
								Type:        "string",
								GoName:      "VPCConnectorEgressSettings",
								GoType:      "FunctionVPCConnectorEgressSettingsEnum",
								Description: "The egress settings for the connector, controlling what traffic is diverted\nthrough it. Possible values: VPC_CONNECTOR_EGRESS_SETTINGS_UNSPECIFIED, PRIVATE_RANGES_ONLY, ALL_TRAFFIC",
								Enum: []string{
									"VPC_CONNECTOR_EGRESS_SETTINGS_UNSPECIFIED",
									"PRIVATE_RANGES_ONLY",
									"ALL_TRAFFIC",
								},
							},
						},
					},
				},
			},
		},
	}
}
