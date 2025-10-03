// Copyright 2023 Google LLC. All Rights Reserved.
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

func DCLJobSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "Run/Job",
			Description: "The Run Job resource",
			StructName:  "Job",
			HasIAM:      true,
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a Job",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "job",
						Required:    true,
						Description: "A full instance of a Job",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a Job",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "job",
						Required:    true,
						Description: "A full instance of a Job",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a Job",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "job",
						Required:    true,
						Description: "A full instance of a Job",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all Job",
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
				Description: "The function used to list information about many Job",
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
				"Job": &dcl.Component{
					Title:           "Job",
					ID:              "projects/{{project}}/locations/{{location}}/jobs/{{name}}",
					ParentContainer: "project",
					HasCreate:       true,
					HasIAM:          true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"name",
							"template",
							"project",
							"location",
						},
						Properties: map[string]*dcl.Property{
							"annotations": &dcl.Property{
								Type: "object",
								AdditionalProperties: &dcl.Property{
									Type: "string",
								},
								GoName:      "Annotations",
								Description: "KRM-style annotations for the resource. Unstructured key value map that may be set by external tools to store and arbitrary metadata. They are not queryable and should be preserved when modifying objects. Cloud Run will populate some annotations using 'run.googleapis.com' or 'serving.knative.dev' namespaces. This field follows Kubernetes annotations' namespacing, limits, and rules. More info: https://kubernetes.io/docs/user-guide/annotations",
							},
							"binaryAuthorization": &dcl.Property{
								Type:        "object",
								GoName:      "BinaryAuthorization",
								GoType:      "JobBinaryAuthorization",
								Description: "Settings for the Binary Authorization feature.",
								Properties: map[string]*dcl.Property{
									"breakglassJustification": &dcl.Property{
										Type:        "string",
										GoName:      "BreakglassJustification",
										Description: "If present, indicates to use Breakglass using this justification. If use_default is False, then it must be empty. For more information on breakglass, see https://cloud.google.com/binary-authorization/docs/using-breakglass",
									},
									"useDefault": &dcl.Property{
										Type:        "boolean",
										GoName:      "UseDefault",
										Description: "If True, indicates to use the default project's binary authorization policy. If False, binary authorization will be disabled.",
									},
								},
							},
							"client": &dcl.Property{
								Type:        "string",
								GoName:      "Client",
								Description: "Arbitrary identifier for the API client.",
							},
							"clientVersion": &dcl.Property{
								Type:        "string",
								GoName:      "ClientVersion",
								Description: "Arbitrary version identifier for the API client.",
							},
							"conditions": &dcl.Property{
								Type:        "array",
								GoName:      "Conditions",
								ReadOnly:    true,
								Description: "Output only. The Conditions of all other associated sub-resources. They contain additional diagnostics information in case the Job does not reach its desired state. See comments in `reconciling` for additional information on reconciliation process in Cloud Run.",
								Immutable:   true,
								ListType:    "list",
								Items: &dcl.Property{
									Type:   "object",
									GoType: "JobConditions",
									Properties: map[string]*dcl.Property{
										"executionReason": &dcl.Property{
											Type:        "string",
											GoName:      "ExecutionReason",
											GoType:      "JobConditionsExecutionReasonEnum",
											Description: "A reason for the execution condition. Possible values: EXECUTION_REASON_UNDEFINED, JOB_STATUS_SERVICE_POLLING_ERROR, NON_ZERO_EXIT_CODE",
											Conflicts: []string{
												"reason",
												"revisionReason",
											},
											Enum: []string{
												"EXECUTION_REASON_UNDEFINED",
												"JOB_STATUS_SERVICE_POLLING_ERROR",
												"NON_ZERO_EXIT_CODE",
											},
										},
										"lastTransitionTime": &dcl.Property{
											Type:        "string",
											Format:      "date-time",
											GoName:      "LastTransitionTime",
											Description: "Last time the condition transitioned from one status to another.",
										},
										"message": &dcl.Property{
											Type:        "string",
											GoName:      "Message",
											Description: "Human readable message indicating details about the current status.",
										},
										"reason": &dcl.Property{
											Type:        "string",
											GoName:      "Reason",
											GoType:      "JobConditionsReasonEnum",
											Description: "A common (service-level) reason for this condition. Possible values: COMMON_REASON_UNDEFINED, UNKNOWN, REVISION_FAILED, PROGRESS_DEADLINE_EXCEEDED, BUILD_STEP_FAILED, CONTAINER_MISSING, CONTAINER_PERMISSION_DENIED, CONTAINER_IMAGE_UNAUTHORIZED, CONTAINER_IMAGE_AUTHORIZATION_CHECK_FAILED, ENCRYPTION_KEY_PERMISSION_DENIED, ENCRYPTION_KEY_CHECK_FAILED, SECRETS_ACCESS_CHECK_FAILED, WAITING_FOR_OPERATION, IMMEDIATE_RETRY, POSTPONED_RETRY",
											Conflicts: []string{
												"revisionReason",
												"executionReason",
											},
											Enum: []string{
												"COMMON_REASON_UNDEFINED",
												"UNKNOWN",
												"REVISION_FAILED",
												"PROGRESS_DEADLINE_EXCEEDED",
												"BUILD_STEP_FAILED",
												"CONTAINER_MISSING",
												"CONTAINER_PERMISSION_DENIED",
												"CONTAINER_IMAGE_UNAUTHORIZED",
												"CONTAINER_IMAGE_AUTHORIZATION_CHECK_FAILED",
												"ENCRYPTION_KEY_PERMISSION_DENIED",
												"ENCRYPTION_KEY_CHECK_FAILED",
												"SECRETS_ACCESS_CHECK_FAILED",
												"WAITING_FOR_OPERATION",
												"IMMEDIATE_RETRY",
												"POSTPONED_RETRY",
											},
										},
										"revisionReason": &dcl.Property{
											Type:        "string",
											GoName:      "RevisionReason",
											GoType:      "JobConditionsRevisionReasonEnum",
											Description: "A reason for the revision condition. Possible values: REVISION_REASON_UNDEFINED, PENDING, RESERVE, RETIRED, RETIRING, RECREATING, HEALTH_CHECK_CONTAINER_ERROR, CUSTOMIZED_PATH_RESPONSE_PENDING, MIN_INSTANCES_NOT_PROVISIONED, ACTIVE_REVISION_LIMIT_REACHED, NO_DEPLOYMENT, HEALTH_CHECK_SKIPPED",
											Conflicts: []string{
												"reason",
												"executionReason",
											},
											Enum: []string{
												"REVISION_REASON_UNDEFINED",
												"PENDING",
												"RESERVE",
												"RETIRED",
												"RETIRING",
												"RECREATING",
												"HEALTH_CHECK_CONTAINER_ERROR",
												"CUSTOMIZED_PATH_RESPONSE_PENDING",
												"MIN_INSTANCES_NOT_PROVISIONED",
												"ACTIVE_REVISION_LIMIT_REACHED",
												"NO_DEPLOYMENT",
												"HEALTH_CHECK_SKIPPED",
											},
										},
										"severity": &dcl.Property{
											Type:        "string",
											GoName:      "Severity",
											GoType:      "JobConditionsSeverityEnum",
											Description: "How to interpret failures of this condition, one of Error, Warning, Info Possible values: SEVERITY_UNSPECIFIED, ERROR, WARNING, INFO",
											Enum: []string{
												"SEVERITY_UNSPECIFIED",
												"ERROR",
												"WARNING",
												"INFO",
											},
										},
										"state": &dcl.Property{
											Type:        "string",
											GoName:      "State",
											GoType:      "JobConditionsStateEnum",
											Description: "State of the condition. Possible values: STATE_UNSPECIFIED, CONDITION_PENDING, CONDITION_RECONCILING, CONDITION_FAILED, CONDITION_SUCCEEDED",
											Enum: []string{
												"STATE_UNSPECIFIED",
												"CONDITION_PENDING",
												"CONDITION_RECONCILING",
												"CONDITION_FAILED",
												"CONDITION_SUCCEEDED",
											},
										},
										"type": &dcl.Property{
											Type:        "string",
											GoName:      "Type",
											Description: "type is used to communicate the status of the reconciliation process. See also: https://github.com/knative/serving/blob/main/docs/spec/errors.md#error-conditions-and-reporting Types common to all resources include: * \"Ready\": True when the Resource is ready.",
										},
									},
								},
							},
							"createTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "CreateTime",
								ReadOnly:    true,
								Description: "Output only. The creation time.",
								Immutable:   true,
							},
							"creator": &dcl.Property{
								Type:        "string",
								GoName:      "Creator",
								ReadOnly:    true,
								Description: "Output only. Email address of the authenticated creator.",
								Immutable:   true,
							},
							"deleteTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "DeleteTime",
								ReadOnly:    true,
								Description: "Output only. The deletion time.",
								Immutable:   true,
							},
							"etag": &dcl.Property{
								Type:        "string",
								GoName:      "Etag",
								ReadOnly:    true,
								Description: "Output only. A system-generated fingerprint for this version of the resource. May be used to detect modification conflict during updates.",
								Immutable:   true,
							},
							"executionCount": &dcl.Property{
								Type:        "integer",
								Format:      "int64",
								GoName:      "ExecutionCount",
								ReadOnly:    true,
								Description: "Output only. Number of executions created for this job.",
								Immutable:   true,
							},
							"expireTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "ExpireTime",
								ReadOnly:    true,
								Description: "Output only. For a deleted resource, the time after which it will be permamently deleted.",
								Immutable:   true,
							},
							"generation": &dcl.Property{
								Type:        "integer",
								Format:      "int64",
								GoName:      "Generation",
								ReadOnly:    true,
								Description: "Output only. A number that monotonically increases every time the user modifies the desired state.",
								Immutable:   true,
							},
							"labels": &dcl.Property{
								Type: "object",
								AdditionalProperties: &dcl.Property{
									Type: "string",
								},
								GoName:      "Labels",
								ReadOnly:    true,
								Description: "KRM-style labels for the resource. User-provided labels are shared with Google's billing system, so they can be used to filter, or break down billing charges by team, component, environment, state, etc. For more information, visit https://cloud.google.com/resource-manager/docs/creating-managing-labels or https://cloud.google.com/run/docs/configuring/labels Cloud Run will populate some labels with 'run.googleapis.com' or 'serving.knative.dev' namespaces. Those labels are read-only, and user changes will not be preserved.",
								Immutable:   true,
							},
							"lastModifier": &dcl.Property{
								Type:        "string",
								GoName:      "LastModifier",
								ReadOnly:    true,
								Description: "Output only. Email address of the last authenticated modifier.",
								Immutable:   true,
							},
							"latestCreatedExecution": &dcl.Property{
								Type:        "object",
								GoName:      "LatestCreatedExecution",
								GoType:      "JobLatestCreatedExecution",
								ReadOnly:    true,
								Description: "Output only. Name of the last created execution.",
								Properties: map[string]*dcl.Property{
									"createTime": &dcl.Property{
										Type:        "string",
										Format:      "date-time",
										GoName:      "CreateTime",
										Description: "Creation timestamp of the execution.",
									},
									"name": &dcl.Property{
										Type:        "string",
										GoName:      "Name",
										Description: "Name of the execution.",
										ResourceReferences: []*dcl.PropertyResourceReference{
											&dcl.PropertyResourceReference{
												Resource: "Run/Execution",
												Field:    "selfLink",
												Parent:   true,
											},
										},
									},
								},
							},
							"latestSucceededExecution": &dcl.Property{
								Type:        "object",
								GoName:      "LatestSucceededExecution",
								GoType:      "JobLatestSucceededExecution",
								ReadOnly:    true,
								Description: "Output only. Name of the last succeeded execution.",
								Properties: map[string]*dcl.Property{
									"createTime": &dcl.Property{
										Type:        "string",
										Format:      "date-time",
										GoName:      "CreateTime",
										Description: "Creation timestamp of the execution.",
									},
									"name": &dcl.Property{
										Type:        "string",
										GoName:      "Name",
										Description: "Name of the execution.",
										ResourceReferences: []*dcl.PropertyResourceReference{
											&dcl.PropertyResourceReference{
												Resource: "Run/Execution",
												Field:    "selfLink",
												Parent:   true,
											},
										},
									},
								},
							},
							"launchStage": &dcl.Property{
								Type:        "string",
								GoName:      "LaunchStage",
								GoType:      "JobLaunchStageEnum",
								Description: "The launch stage as defined by [Google Cloud Platform Launch Stages](https://cloud.google.com/terms/launch-stages). Cloud Run supports `ALPHA`, `BETA`, and `GA`. If no value is specified, GA is assumed. Possible values: LAUNCH_STAGE_UNSPECIFIED, UNIMPLEMENTED, PRELAUNCH, EARLY_ACCESS, ALPHA, BETA, GA, DEPRECATED",
								Enum: []string{
									"LAUNCH_STAGE_UNSPECIFIED",
									"UNIMPLEMENTED",
									"PRELAUNCH",
									"EARLY_ACCESS",
									"ALPHA",
									"BETA",
									"GA",
									"DEPRECATED",
								},
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
								Description: "The fully qualified name of this Job. Format: projects/{project}/locations/{location}/jobs/{job}",
							},
							"observedGeneration": &dcl.Property{
								Type:        "integer",
								Format:      "int64",
								GoName:      "ObservedGeneration",
								ReadOnly:    true,
								Description: "Output only. The generation of this Job. See comments in `reconciling` for additional information on reconciliation process in Cloud Run.",
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
							"reconciling": &dcl.Property{
								Type:        "boolean",
								GoName:      "Reconciling",
								ReadOnly:    true,
								Description: "Output only. Returns true if the Job is currently being acted upon by the system to bring it into the desired state. When a new Job is created, or an existing one is updated, Cloud Run will asynchronously perform all necessary steps to bring the Job to the desired state. This process is called reconciliation. While reconciliation is in process, `observed_generation` and `latest_succeeded_execution`, will have transient values that might mismatch the intended state: Once reconciliation is over (and this field is false), there are two possible outcomes: reconciliation succeeded and the state matches the Job, or there was an error, and reconciliation failed. This state can be found in `terminal_condition.state`. If reconciliation succeeded, the following fields will match: `observed_generation` and `generation`, `latest_succeeded_execution` and `latest_created_execution`. If reconciliation failed, `observed_generation` and `latest_succeeded_execution` will have the state of the last succeeded execution or empty for newly created Job. Additional information on the failure can be found in `terminal_condition` and `conditions`.",
								Immutable:   true,
							},
							"template": &dcl.Property{
								Type:        "object",
								GoName:      "Template",
								GoType:      "JobTemplate",
								Description: "Required. The template used to create executions for this Job.",
								Required: []string{
									"template",
								},
								Properties: map[string]*dcl.Property{
									"annotations": &dcl.Property{
										Type: "object",
										AdditionalProperties: &dcl.Property{
											Type: "string",
										},
										GoName:      "Annotations",
										Description: "KRM-style annotations for the resource.",
									},
									"labels": &dcl.Property{
										Type: "object",
										AdditionalProperties: &dcl.Property{
											Type: "string",
										},
										GoName:      "Labels",
										Description: "KRM-style labels for the resource.",
									},
									"parallelism": &dcl.Property{
										Type:          "integer",
										Format:        "int64",
										GoName:        "Parallelism",
										Description:   "Specifies the maximum desired number of tasks the execution should run at any given time. Must be <= task_count. The actual number of tasks running in steady state will be less than this number when ((.spec.task_count - .status.successful) < .spec.parallelism), i.e. when the work left to do is less than max parallelism. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/",
										ServerDefault: true,
									},
									"taskCount": &dcl.Property{
										Type:          "integer",
										Format:        "int64",
										GoName:        "TaskCount",
										Description:   "Specifies the desired number of tasks the execution should run. Setting to 1 means that parallelism is limited to 1 and the success of that task signals the success of the execution. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/",
										ServerDefault: true,
									},
									"template": &dcl.Property{
										Type:        "object",
										GoName:      "Template",
										GoType:      "JobTemplateTemplate",
										Description: "Required. Describes the task(s) that will be created when executing an execution.",
										Properties: map[string]*dcl.Property{
											"containers": &dcl.Property{
												Type:        "array",
												GoName:      "Containers",
												Description: "Holds the single container that defines the unit of execution for this task.",
												SendEmpty:   true,
												ListType:    "list",
												Items: &dcl.Property{
													Type:   "object",
													GoType: "JobTemplateTemplateContainers",
													Required: []string{
														"image",
													},
													Properties: map[string]*dcl.Property{
														"args": &dcl.Property{
															Type:        "array",
															GoName:      "Args",
															Description: "Arguments to the entrypoint. The docker image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
															SendEmpty:   true,
															ListType:    "list",
															Items: &dcl.Property{
																Type:   "string",
																GoType: "string",
															},
														},
														"command": &dcl.Property{
															Type:        "array",
															GoName:      "Command",
															Description: "Entrypoint array. Not executed within a shell. The docker image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
															SendEmpty:   true,
															ListType:    "list",
															Items: &dcl.Property{
																Type:   "string",
																GoType: "string",
															},
														},
														"env": &dcl.Property{
															Type:        "array",
															GoName:      "Env",
															Description: "List of environment variables to set in the container.",
															SendEmpty:   true,
															ListType:    "list",
															Items: &dcl.Property{
																Type:   "object",
																GoType: "JobTemplateTemplateContainersEnv",
																Required: []string{
																	"name",
																},
																Properties: map[string]*dcl.Property{
																	"name": &dcl.Property{
																		Type:        "string",
																		GoName:      "Name",
																		Description: "Required. Name of the environment variable. Must be a C_IDENTIFIER, and mnay not exceed 32768 characters.",
																	},
																	"value": &dcl.Property{
																		Type:        "string",
																		GoName:      "Value",
																		Description: "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any route environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to \"\", and the maximum length is 32768 bytes.",
																		Conflicts: []string{
																			"valueSource",
																		},
																	},
																	"valueSource": &dcl.Property{
																		Type:        "object",
																		GoName:      "ValueSource",
																		GoType:      "JobTemplateTemplateContainersEnvValueSource",
																		Description: "Source for the environment variable's value.",
																		Conflicts: []string{
																			"value",
																		},
																		Properties: map[string]*dcl.Property{
																			"secretKeyRef": &dcl.Property{
																				Type:        "object",
																				GoName:      "SecretKeyRef",
																				GoType:      "JobTemplateTemplateContainersEnvValueSourceSecretKeyRef",
																				Description: "Selects a secret and a specific version from Cloud Secret Manager.",
																				Required: []string{
																					"secret",
																				},
																				Properties: map[string]*dcl.Property{
																					"secret": &dcl.Property{
																						Type:        "string",
																						GoName:      "Secret",
																						Description: "Required. The name of the secret in Cloud Secret Manager. Format: {secret_name} if the secret is in the same project. projects/{project}/secrets/{secret_name} if the secret is in a different project.",
																						ResourceReferences: []*dcl.PropertyResourceReference{
																							&dcl.PropertyResourceReference{
																								Resource: "Secretmanager/Secret",
																								Field:    "selfLink",
																							},
																						},
																					},
																					"version": &dcl.Property{
																						Type:        "string",
																						GoName:      "Version",
																						Description: "The Cloud Secret Manager secret version. Can be 'latest' for the latest value or an integer for a specific version.",
																						ResourceReferences: []*dcl.PropertyResourceReference{
																							&dcl.PropertyResourceReference{
																								Resource: "Secretmanager/SecretVersion",
																								Field:    "selfLink",
																							},
																						},
																					},
																				},
																			},
																		},
																	},
																},
															},
														},
														"image": &dcl.Property{
															Type:        "string",
															GoName:      "Image",
															Description: "Required. URL of the Container image in Google Container Registry or Docker More info: https://kubernetes.io/docs/concepts/containers/images",
														},
														"name": &dcl.Property{
															Type:        "string",
															GoName:      "Name",
															Description: "Name of the container specified as a DNS_LABEL.",
														},
														"ports": &dcl.Property{
															Type:        "array",
															GoName:      "Ports",
															Description: "List of ports to expose from the container. Only a single port can be specified. The specified ports must be listening on all interfaces (0.0.0.0) within the container to be accessible. If omitted, a port number will be chosen and passed to the container through the PORT environment variable for the container to listen on.",
															SendEmpty:   true,
															ListType:    "list",
															Items: &dcl.Property{
																Type:   "object",
																GoType: "JobTemplateTemplateContainersPorts",
																Properties: map[string]*dcl.Property{
																	"containerPort": &dcl.Property{
																		Type:        "integer",
																		Format:      "int64",
																		GoName:      "ContainerPort",
																		Description: "Port number the container listens on. This must be a valid TCP port number, 0 < container_port < 65536.",
																	},
																	"name": &dcl.Property{
																		Type:        "string",
																		GoName:      "Name",
																		Description: "If specified, used to specify which protocol to use. Allowed values are \"http1\" and \"h2c\".",
																	},
																},
															},
														},
														"resources": &dcl.Property{
															Type:          "object",
															GoName:        "Resources",
															GoType:        "JobTemplateTemplateContainersResources",
															Description:   "Compute Resource requirements by this container. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
															ServerDefault: true,
															Properties: map[string]*dcl.Property{
																"cpuIdle": &dcl.Property{
																	Type:          "boolean",
																	GoName:        "CpuIdle",
																	Description:   "Determines whether CPU should be throttled or not outside of requests.",
																	ServerDefault: true,
																},
																"limits": &dcl.Property{
																	Type: "object",
																	AdditionalProperties: &dcl.Property{
																		Type: "string",
																	},
																	GoName:      "Limits",
																	Description: "Only memory and CPU are supported. Note: The only supported values for CPU are '1', '2', and '4'. Setting 4 CPU requires at least 2Gi of memory. The values of the map is string form of the 'quantity' k8s type: https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apimachinery/pkg/api/resource/quantity.go",
																},
															},
														},
														"volumeMounts": &dcl.Property{
															Type:        "array",
															GoName:      "VolumeMounts",
															Description: "Volume to mount into the container's filesystem.",
															SendEmpty:   true,
															ListType:    "list",
															Items: &dcl.Property{
																Type:   "object",
																GoType: "JobTemplateTemplateContainersVolumeMounts",
																Required: []string{
																	"name",
																	"mountPath",
																},
																Properties: map[string]*dcl.Property{
																	"mountPath": &dcl.Property{
																		Type:        "string",
																		GoName:      "MountPath",
																		Description: "Required. Path within the container at which the volume should be mounted. Must not contain ':'. For Cloud SQL volumes, it can be left empty, or must otherwise be `/cloudsql`. All instances defined in the Volume will be available as `/cloudsql/[instance]`. For more information on Cloud SQL volumes, visit https://cloud.google.com/sql/docs/mysql/connect-run",
																	},
																	"name": &dcl.Property{
																		Type:        "string",
																		GoName:      "Name",
																		Description: "Required. This must match the Name of a Volume.",
																	},
																},
															},
														},
													},
												},
											},
											"encryptionKey": &dcl.Property{
												Type:        "string",
												GoName:      "EncryptionKey",
												Description: "A reference to a customer managed encryption key (CMEK) to use to encrypt this container image. For more information, go to https://cloud.google.com/run/docs/securing/using-cmek",
												ResourceReferences: []*dcl.PropertyResourceReference{
													&dcl.PropertyResourceReference{
														Resource: "Cloudkms/CryptoKey",
														Field:    "selfLink",
													},
												},
											},
											"executionEnvironment": &dcl.Property{
												Type:          "string",
												GoName:        "ExecutionEnvironment",
												GoType:        "JobTemplateTemplateExecutionEnvironmentEnum",
												Description:   "The execution environment being used to host this Task. Possible values: EXECUTION_ENVIRONMENT_UNSPECIFIED, EXECUTION_ENVIRONMENT_DEFAULT, EXECUTION_ENVIRONMENT_GEN2",
												ServerDefault: true,
												Enum: []string{
													"EXECUTION_ENVIRONMENT_UNSPECIFIED",
													"EXECUTION_ENVIRONMENT_DEFAULT",
													"EXECUTION_ENVIRONMENT_GEN2",
												},
											},
											"maxRetries": &dcl.Property{
												Type:          "integer",
												Format:        "int64",
												GoName:        "MaxRetries",
												Description:   "Number of retries allowed per Task, before marking this Task failed.",
												ServerDefault: true,
											},
											"serviceAccount": &dcl.Property{
												Type:          "string",
												GoName:        "ServiceAccount",
												Description:   "Email address of the IAM service account associated with the Task of a Job. The service account represents the identity of the running task, and determines what permissions the task has. If not provided, the task will use the project's default service account.",
												ServerDefault: true,
											},
											"timeout": &dcl.Property{
												Type:          "string",
												GoName:        "Timeout",
												Description:   "Max allowed time duration the Task may be active before the system will actively try to mark it failed and kill associated containers. This applies per attempt of a task, meaning each retry can run for the full timeout.",
												ServerDefault: true,
											},
											"volumes": &dcl.Property{
												Type:        "array",
												GoName:      "Volumes",
												Description: "A list of Volumes to make available to containers.",
												SendEmpty:   true,
												ListType:    "list",
												Items: &dcl.Property{
													Type:   "object",
													GoType: "JobTemplateTemplateVolumes",
													Required: []string{
														"name",
													},
													Properties: map[string]*dcl.Property{
														"cloudSqlInstance": &dcl.Property{
															Type:        "object",
															GoName:      "CloudSqlInstance",
															GoType:      "JobTemplateTemplateVolumesCloudSqlInstance",
															Description: "For Cloud SQL volumes, contains the specific instances that should be mounted. Visit https://cloud.google.com/sql/docs/mysql/connect-run for more information on how to connect Cloud SQL and Cloud Run.",
															Conflicts: []string{
																"secret",
															},
															Properties: map[string]*dcl.Property{
																"instances": &dcl.Property{
																	Type:        "array",
																	GoName:      "Instances",
																	Description: "The Cloud SQL instance connection names, as can be found in https://console.cloud.google.com/sql/instances. Visit https://cloud.google.com/sql/docs/mysql/connect-run for more information on how to connect Cloud SQL and Cloud Run. Format: {project}:{location}:{instance}",
																	SendEmpty:   true,
																	ListType:    "list",
																	Items: &dcl.Property{
																		Type:   "string",
																		GoType: "string",
																	},
																},
															},
														},
														"name": &dcl.Property{
															Type:        "string",
															GoName:      "Name",
															Description: "Required. Volume's name.",
														},
														"secret": &dcl.Property{
															Type:        "object",
															GoName:      "Secret",
															GoType:      "JobTemplateTemplateVolumesSecret",
															Description: "Secret represents a secret that should populate this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
															Conflicts: []string{
																"cloudSqlInstance",
															},
															Required: []string{
																"secret",
															},
															Properties: map[string]*dcl.Property{
																"defaultMode": &dcl.Property{
																	Type:        "integer",
																	Format:      "int64",
																	GoName:      "DefaultMode",
																	Description: "Integer representation of mode bits to use on created files by default. Must be a value between 0000 and 0777 (octal), defaulting to 0644. Directories within the path are not affected by this setting. Notes * Internally, a umask of 0222 will be applied to any non-zero value. * This is an integer representation of the mode bits. So, the octal integer value should look exactly as the chmod numeric notation with a leading zero. Some examples: for chmod 777 (a=rwx), set to 0777 (octal) or 511 (base-10). For chmod 640 (u=rw,g=r), set to 0640 (octal) or 416 (base-10). For chmod 755 (u=rwx,g=rx,o=rx), set to 0755 (octal) or 493 (base-10). * This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set. This might be in conflict with other options that affect the file mode, like fsGroup, and as a result, other mode bits could be set.",
																},
																"items": &dcl.Property{
																	Type:        "array",
																	GoName:      "Items",
																	Description: "If unspecified, the volume will expose a file whose name is the secret, relative to VolumeMount.mount_path. If specified, the key will be used as the version to fetch from Cloud Secret Manager and the path will be the name of the file exposed in the volume. When items are defined, they must specify a path and a version.",
																	SendEmpty:   true,
																	ListType:    "list",
																	Items: &dcl.Property{
																		Type:   "object",
																		GoType: "JobTemplateTemplateVolumesSecretItems",
																		Required: []string{
																			"path",
																		},
																		Properties: map[string]*dcl.Property{
																			"mode": &dcl.Property{
																				Type:        "integer",
																				Format:      "int64",
																				GoName:      "Mode",
																				Description: "Integer octal mode bits to use on this file, must be a value between 01 and 0777 (octal). If 0 or not set, the Volume's default mode will be used. Notes * Internally, a umask of 0222 will be applied to any non-zero value. * This is an integer representation of the mode bits. So, the octal integer value should look exactly as the chmod numeric notation with a leading zero. Some examples: for chmod 777 (a=rwx), set to 0777 (octal) or 511 (base-10). For chmod 640 (u=rw,g=r), set to 0640 (octal) or 416 (base-10). For chmod 755 (u=rwx,g=rx,o=rx), set to 0755 (octal) or 493 (base-10). * This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																			},
																			"path": &dcl.Property{
																				Type:        "string",
																				GoName:      "Path",
																				Description: "Required. The relative path of the secret in the container.",
																			},
																			"version": &dcl.Property{
																				Type:        "string",
																				GoName:      "Version",
																				Description: "The Cloud Secret Manager secret version. Can be 'latest' for the latest value or an integer for a specific version.",
																			},
																		},
																	},
																},
																"secret": &dcl.Property{
																	Type:        "string",
																	GoName:      "Secret",
																	Description: "Required. The name of the secret in Cloud Secret Manager. Format: {secret} if the secret is in the same project. projects/{project}/secrets/{secret} if the secret is in a different project.",
																},
															},
														},
													},
												},
											},
											"vpcAccess": &dcl.Property{
												Type:        "object",
												GoName:      "VPCAccess",
												GoType:      "JobTemplateTemplateVPCAccess",
												Description: "VPC Access configuration to use for this Task. For more information, visit https://cloud.google.com/run/docs/configuring/connecting-vpc.",
												Properties: map[string]*dcl.Property{
													"connector": &dcl.Property{
														Type:        "string",
														GoName:      "Connector",
														Description: "VPC Access connector name. Format: projects/{project}/locations/{location}/connectors/{connector}",
														ResourceReferences: []*dcl.PropertyResourceReference{
															&dcl.PropertyResourceReference{
																Resource: "Vpcaccess/Connector",
																Field:    "selfLink",
															},
														},
													},
													"egress": &dcl.Property{
														Type:        "string",
														GoName:      "Egress",
														GoType:      "JobTemplateTemplateVPCAccessEgressEnum",
														Description: "Traffic VPC egress settings. Possible values: VPC_EGRESS_UNSPECIFIED, ALL_TRAFFIC, PRIVATE_RANGES_ONLY",
														Enum: []string{
															"VPC_EGRESS_UNSPECIFIED",
															"ALL_TRAFFIC",
															"PRIVATE_RANGES_ONLY",
														},
													},
												},
											},
										},
									},
								},
							},
							"terminalCondition": &dcl.Property{
								Type:        "object",
								GoName:      "TerminalCondition",
								GoType:      "JobTerminalCondition",
								ReadOnly:    true,
								Description: "Output only. The Condition of this Job, containing its readiness status, and detailed error information in case it did not reach the desired state.",
								Properties: map[string]*dcl.Property{
									"domainMappingReason": &dcl.Property{
										Type:        "string",
										GoName:      "DomainMappingReason",
										GoType:      "JobTerminalConditionDomainMappingReasonEnum",
										Description: "A reason for the domain mapping condition. Possible values: DOMAIN_MAPPING_REASON_UNDEFINED, ROUTE_NOT_READY, PERMISSION_DENIED, CERTIFICATE_ALREADY_EXISTS, MAPPING_ALREADY_EXISTS, CERTIFICATE_PENDING, CERTIFICATE_FAILED",
										Conflicts: []string{
											"reason",
											"internalReason",
											"revisionReason",
											"executionReason",
										},
										Enum: []string{
											"DOMAIN_MAPPING_REASON_UNDEFINED",
											"ROUTE_NOT_READY",
											"PERMISSION_DENIED",
											"CERTIFICATE_ALREADY_EXISTS",
											"MAPPING_ALREADY_EXISTS",
											"CERTIFICATE_PENDING",
											"CERTIFICATE_FAILED",
										},
									},
									"executionReason": &dcl.Property{
										Type:        "string",
										GoName:      "ExecutionReason",
										GoType:      "JobTerminalConditionExecutionReasonEnum",
										Description: "A reason for the execution condition. Possible values: EXECUTION_REASON_UNDEFINED, JOB_STATUS_SERVICE_POLLING_ERROR, NON_ZERO_EXIT_CODE",
										Conflicts: []string{
											"reason",
											"internalReason",
											"domainMappingReason",
											"revisionReason",
										},
										Enum: []string{
											"EXECUTION_REASON_UNDEFINED",
											"JOB_STATUS_SERVICE_POLLING_ERROR",
											"NON_ZERO_EXIT_CODE",
										},
									},
									"internalReason": &dcl.Property{
										Type:        "string",
										GoName:      "InternalReason",
										GoType:      "JobTerminalConditionInternalReasonEnum",
										Description: "A reason for the internal condition. Possible values: INTERNAL_REASON_UNDEFINED, CONFLICTING_REVISION_NAME, REVISION_MISSING, CONFIGURATION_MISSING, ASSIGNING_TRAFFIC, UPDATING_INGRESS_TRAFFIC_ALLOWED, REVISION_ORG_POLICY_VIOLATION, ENABLING_GCFV2_URI_SUPPORT",
										Conflicts: []string{
											"reason",
											"domainMappingReason",
											"revisionReason",
											"executionReason",
										},
										Enum: []string{
											"INTERNAL_REASON_UNDEFINED",
											"CONFLICTING_REVISION_NAME",
											"REVISION_MISSING",
											"CONFIGURATION_MISSING",
											"ASSIGNING_TRAFFIC",
											"UPDATING_INGRESS_TRAFFIC_ALLOWED",
											"REVISION_ORG_POLICY_VIOLATION",
											"ENABLING_GCFV2_URI_SUPPORT",
										},
									},
									"lastTransitionTime": &dcl.Property{
										Type:        "string",
										Format:      "date-time",
										GoName:      "LastTransitionTime",
										Description: "Last time the condition transitioned from one status to another.",
									},
									"message": &dcl.Property{
										Type:        "string",
										GoName:      "Message",
										Description: "Human readable message indicating details about the current status.",
									},
									"reason": &dcl.Property{
										Type:        "string",
										GoName:      "Reason",
										GoType:      "JobTerminalConditionReasonEnum",
										Description: "A common (service-level) reason for this condition. Possible values: COMMON_REASON_UNDEFINED, UNKNOWN, ROUTE_MISSING, REVISION_FAILED, PROGRESS_DEADLINE_EXCEEDED, CONTAINER_MISSING, CONTAINER_PERMISSION_DENIED, CONTAINER_IMAGE_UNAUTHORIZED, CONTAINER_IMAGE_AUTHORIZATION_CHECK_FAILED, ENCRYPTION_KEY_PERMISSION_DENIED, ENCRYPTION_KEY_CHECK_FAILED, SECRETS_ACCESS_CHECK_FAILED, WAITING_FOR_OPERATION, IMMEDIATE_RETRY, POSTPONED_RETRY",
										Conflicts: []string{
											"internalReason",
											"domainMappingReason",
											"revisionReason",
											"executionReason",
										},
										Enum: []string{
											"COMMON_REASON_UNDEFINED",
											"UNKNOWN",
											"ROUTE_MISSING",
											"REVISION_FAILED",
											"PROGRESS_DEADLINE_EXCEEDED",
											"CONTAINER_MISSING",
											"CONTAINER_PERMISSION_DENIED",
											"CONTAINER_IMAGE_UNAUTHORIZED",
											"CONTAINER_IMAGE_AUTHORIZATION_CHECK_FAILED",
											"ENCRYPTION_KEY_PERMISSION_DENIED",
											"ENCRYPTION_KEY_CHECK_FAILED",
											"SECRETS_ACCESS_CHECK_FAILED",
											"WAITING_FOR_OPERATION",
											"IMMEDIATE_RETRY",
											"POSTPONED_RETRY",
										},
									},
									"revisionReason": &dcl.Property{
										Type:        "string",
										GoName:      "RevisionReason",
										GoType:      "JobTerminalConditionRevisionReasonEnum",
										Description: "A reason for the revision condition. Possible values: REVISION_REASON_UNDEFINED, PENDING, RESERVE, RETIRED, RETIRING, RECREATING, HEALTH_CHECK_CONTAINER_ERROR, CUSTOMIZED_PATH_RESPONSE_PENDING, MIN_INSTANCES_NOT_PROVISIONED, ACTIVE_REVISION_LIMIT_REACHED, NO_DEPLOYMENT, HEALTH_CHECK_SKIPPED",
										Conflicts: []string{
											"reason",
											"internalReason",
											"domainMappingReason",
											"executionReason",
										},
										Enum: []string{
											"REVISION_REASON_UNDEFINED",
											"PENDING",
											"RESERVE",
											"RETIRED",
											"RETIRING",
											"RECREATING",
											"HEALTH_CHECK_CONTAINER_ERROR",
											"CUSTOMIZED_PATH_RESPONSE_PENDING",
											"MIN_INSTANCES_NOT_PROVISIONED",
											"ACTIVE_REVISION_LIMIT_REACHED",
											"NO_DEPLOYMENT",
											"HEALTH_CHECK_SKIPPED",
										},
									},
									"severity": &dcl.Property{
										Type:        "string",
										GoName:      "Severity",
										GoType:      "JobTerminalConditionSeverityEnum",
										Description: "How to interpret failures of this condition, one of Error, Warning, Info Possible values: SEVERITY_UNSPECIFIED, ERROR, WARNING, INFO",
										Enum: []string{
											"SEVERITY_UNSPECIFIED",
											"ERROR",
											"WARNING",
											"INFO",
										},
									},
									"state": &dcl.Property{
										Type:        "string",
										GoName:      "State",
										GoType:      "JobTerminalConditionStateEnum",
										Description: "State of the condition. Possible values: STATE_UNSPECIFIED, CONDITION_PENDING, CONDITION_RECONCILING, CONDITION_FAILED, CONDITION_SUCCEEDED",
										Enum: []string{
											"STATE_UNSPECIFIED",
											"CONDITION_PENDING",
											"CONDITION_RECONCILING",
											"CONDITION_FAILED",
											"CONDITION_SUCCEEDED",
										},
									},
									"type": &dcl.Property{
										Type:        "string",
										GoName:      "Type",
										Description: "type is used to communicate the status of the reconciliation process. See also: https://github.com/knative/serving/blob/main/docs/spec/errors.md#error-conditions-and-reporting Types common to all resources include: * \"Ready\": True when the Resource is ready.",
									},
								},
							},
							"uid": &dcl.Property{
								Type:        "string",
								GoName:      "Uid",
								ReadOnly:    true,
								Description: "Output only. Server assigned unique identifier for the Execution. The value is a UUID4 string and guaranteed to remain unchanged until the resource is deleted.",
								Immutable:   true,
							},
							"updateTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "UpdateTime",
								ReadOnly:    true,
								Description: "Output only. The last-modified time.",
								Immutable:   true,
							},
						},
					},
				},
			},
		},
	}
}
