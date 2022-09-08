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

func DCLPatchDeploymentSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "OSConfig/PatchDeployment",
			Description: "Patch deployments are configurations that individual patch jobs use to complete a patch. These configurations include instance filter, package repository settings, and a schedule.",
			StructName:  "PatchDeployment",
			Reference: &dcl.Link{
				Text: "API documentation",
				URL:  "https://cloud.google.com/compute/docs/osconfig/rest",
			},
			Guides: []*dcl.Link{
				&dcl.Link{
					Text: "Official Documentation",
					URL:  "https://cloud.google.com/compute/docs/os-patch-management",
				},
			},
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a PatchDeployment",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "patchDeployment",
						Required:    true,
						Description: "A full instance of a PatchDeployment",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a PatchDeployment",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "patchDeployment",
						Required:    true,
						Description: "A full instance of a PatchDeployment",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a PatchDeployment",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "patchDeployment",
						Required:    true,
						Description: "A full instance of a PatchDeployment",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all PatchDeployment",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:     "project",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
				},
			},
			List: &dcl.Path{
				Description: "The function used to list information about many PatchDeployment",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:     "project",
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
				"PatchDeployment": &dcl.Component{
					Title:           "PatchDeployment",
					ID:              "projects/{{project}}/patchDeployments/{{name}}",
					ParentContainer: "project",
					HasCreate:       true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"name",
							"instanceFilter",
							"project",
						},
						Properties: map[string]*dcl.Property{
							"createTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "CreateTime",
								ReadOnly:    true,
								Description: "Output only. Time the patch deployment was created. Timestamp is in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.",
								Immutable:   true,
							},
							"description": &dcl.Property{
								Type:        "string",
								GoName:      "Description",
								Description: "Optional. Description of the patch deployment. Length of the description is limited to 1024 characters.",
							},
							"duration": &dcl.Property{
								Type:        "string",
								GoName:      "Duration",
								Description: "Optional. Duration of the patch. After the duration ends, the patch times out.",
							},
							"instanceFilter": &dcl.Property{
								Type:        "object",
								GoName:      "InstanceFilter",
								GoType:      "PatchDeploymentInstanceFilter",
								Description: "Required. VM instances to patch.",
								Properties: map[string]*dcl.Property{
									"all": &dcl.Property{
										Type:        "boolean",
										GoName:      "All",
										Description: "Target all VM instances in the project. If true, no other criteria is permitted.",
										Conflicts: []string{
											"instanceNamePrefixes",
										},
									},
									"groupLabels": &dcl.Property{
										Type:        "array",
										GoName:      "GroupLabels",
										Description: "Targets VM instances matching ANY of these GroupLabels. This allows targeting of disparate groups of VM instances.",
										Conflicts: []string{
											"all",
										},
										SendEmpty: true,
										ListType:  "list",
										Items: &dcl.Property{
											Type:   "object",
											GoType: "PatchDeploymentInstanceFilterGroupLabels",
											Properties: map[string]*dcl.Property{
												"labels": &dcl.Property{
													Type: "object",
													AdditionalProperties: &dcl.Property{
														Type: "string",
													},
													GoName:      "Labels",
													Description: "Compute Engine instance labels that must be present for a VM instance to be targeted by this filter.",
												},
											},
										},
									},
									"instanceNamePrefixes": &dcl.Property{
										Type:        "array",
										GoName:      "InstanceNamePrefixes",
										Description: "Targets VMs whose name starts with one of these prefixes. Similar to labels, this is another way to group VMs when targeting configs, for example prefix=\"prod-\".",
										Conflicts: []string{
											"all",
										},
										SendEmpty: true,
										ListType:  "list",
										Items: &dcl.Property{
											Type:   "string",
											GoType: "string",
										},
									},
									"instances": &dcl.Property{
										Type:        "array",
										GoName:      "Instances",
										Description: "Targets any of the VM instances specified. Instances are specified by their URI in the form `zones/[ZONE]/instances/[INSTANCE_NAME]`, `projects/[PROJECT_ID]/zones/[ZONE]/instances/[INSTANCE_NAME]`, or `https://www.googleapis.com/compute/v1/projects/[PROJECT_ID]/zones/[ZONE]/instances/[INSTANCE_NAME]`",
										Conflicts: []string{
											"all",
										},
										SendEmpty: true,
										ListType:  "list",
										Items: &dcl.Property{
											Type:   "string",
											GoType: "string",
											ResourceReferences: []*dcl.PropertyResourceReference{
												&dcl.PropertyResourceReference{
													Resource: "Compute/Instance",
													Field:    "name",
												},
											},
										},
									},
									"zones": &dcl.Property{
										Type:        "array",
										GoName:      "Zones",
										Description: "Targets VM instances in ANY of these zones. Leave empty to target VM instances in any zone.",
										Conflicts: []string{
											"all",
										},
										SendEmpty: true,
										ListType:  "list",
										Items: &dcl.Property{
											Type:   "string",
											GoType: "string",
										},
									},
								},
							},
							"lastExecuteTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "LastExecuteTime",
								ReadOnly:    true,
								Description: "Output only. The last time a patch job was started by this deployment. Timestamp is in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.",
								Immutable:   true,
							},
							"name": &dcl.Property{
								Type:        "string",
								GoName:      "Name",
								Description: "Unique name for the patch deployment resource in a project. The patch deployment name is in the form: `projects/{project_id}/patchDeployments/{patch_deployment_id}`. This field is ignored when you create a new patch deployment.",
								Immutable:   true,
							},
							"oneTimeSchedule": &dcl.Property{
								Type:        "object",
								GoName:      "OneTimeSchedule",
								GoType:      "PatchDeploymentOneTimeSchedule",
								Description: "Required. Schedule a one-time execution.",
								Conflicts: []string{
									"recurringSchedule",
								},
								Required: []string{
									"executeTime",
								},
								Properties: map[string]*dcl.Property{
									"executeTime": &dcl.Property{
										Type:        "string",
										Format:      "date-time",
										GoName:      "ExecuteTime",
										Description: "Required. The desired patch job execution time.",
									},
								},
							},
							"patchConfig": &dcl.Property{
								Type:        "object",
								GoName:      "PatchConfig",
								GoType:      "PatchDeploymentPatchConfig",
								Description: "Optional. Patch configuration that is applied.",
								Properties: map[string]*dcl.Property{
									"apt": &dcl.Property{
										Type:        "object",
										GoName:      "Apt",
										GoType:      "PatchDeploymentPatchConfigApt",
										Description: "Apt update settings. Use this setting to override the default `apt` patch rules.",
										Properties: map[string]*dcl.Property{
											"excludes": &dcl.Property{
												Type:        "array",
												GoName:      "Excludes",
												Description: "List of packages to exclude from update. These packages will be excluded",
												Conflicts: []string{
													"exclusivePackages",
												},
												SendEmpty: true,
												ListType:  "list",
												Items: &dcl.Property{
													Type:   "string",
													GoType: "string",
												},
											},
											"exclusivePackages": &dcl.Property{
												Type:        "array",
												GoName:      "ExclusivePackages",
												Description: "An exclusive list of packages to be updated. These are the only packages that will be updated. If these packages are not installed, they will be ignored. This field cannot be specified with any other patch configuration fields.",
												Conflicts: []string{
													"excludes",
												},
												SendEmpty: true,
												ListType:  "list",
												Items: &dcl.Property{
													Type:   "string",
													GoType: "string",
												},
											},
											"type": &dcl.Property{
												Type:        "string",
												GoName:      "Type",
												GoType:      "PatchDeploymentPatchConfigAptTypeEnum",
												Description: "By changing the type to DIST, the patching is performed using `apt-get dist-upgrade` instead. Possible values: TYPE_UNSPECIFIED, DIST, UPGRADE",
												Enum: []string{
													"TYPE_UNSPECIFIED",
													"DIST",
													"UPGRADE",
												},
											},
										},
									},
									"goo": &dcl.Property{
										Type:        "object",
										GoName:      "Goo",
										GoType:      "PatchDeploymentPatchConfigGoo",
										Description: "Goo update settings. Use this setting to override the default `goo` patch rules.",
										Properties:  map[string]*dcl.Property{},
									},
									"postStep": &dcl.Property{
										Type:        "object",
										GoName:      "PostStep",
										GoType:      "PatchDeploymentPatchConfigPostStep",
										Description: "The `ExecStep` to run after the patch update.",
										Properties: map[string]*dcl.Property{
											"linuxExecStepConfig": &dcl.Property{
												Type:        "object",
												GoName:      "LinuxExecStepConfig",
												GoType:      "PatchDeploymentPatchConfigPostStepLinuxExecStepConfig",
												Description: "The ExecStepConfig for all Linux VMs targeted by the PatchJob.",
												Properties: map[string]*dcl.Property{
													"allowedSuccessCodes": &dcl.Property{
														Type:        "array",
														GoName:      "AllowedSuccessCodes",
														Description: "Defaults to [0]. A list of possible return values that the execution can return to indicate a success.",
														SendEmpty:   true,
														ListType:    "list",
														Items: &dcl.Property{
															Type:   "integer",
															Format: "int64",
															GoType: "int64",
														},
													},
													"gcsObject": &dcl.Property{
														Type:        "object",
														GoName:      "GcsObject",
														GoType:      "PatchDeploymentPatchConfigPostStepLinuxExecStepConfigGcsObject",
														Description: "A Cloud Storage object containing the executable.",
														Conflicts: []string{
															"localPath",
														},
														Required: []string{
															"bucket",
															"object",
															"generationNumber",
														},
														Properties: map[string]*dcl.Property{
															"bucket": &dcl.Property{
																Type:        "string",
																GoName:      "Bucket",
																Description: "Required. Bucket of the Cloud Storage object.",
															},
															"generationNumber": &dcl.Property{
																Type:        "integer",
																Format:      "int64",
																GoName:      "GenerationNumber",
																Description: "Required. Generation number of the Cloud Storage object. This is used to ensure that the ExecStep specified by this PatchJob does not change.",
															},
															"object": &dcl.Property{
																Type:        "string",
																GoName:      "Object",
																Description: "Required. Name of the Cloud Storage object.",
															},
														},
													},
													"interpreter": &dcl.Property{
														Type:        "string",
														GoName:      "Interpreter",
														GoType:      "PatchDeploymentPatchConfigPostStepLinuxExecStepConfigInterpreterEnum",
														Description: "The script interpreter to use to run the script. If no interpreter is specified the script will be executed directly, which will likely only succeed for scripts with [shebang lines] (https://en.wikipedia.org/wiki/Shebang_(Unix)). Possible values: INTERPRETER_UNSPECIFIED, NONE, SHELL, POWERSHELL",
														Enum: []string{
															"INTERPRETER_UNSPECIFIED",
															"NONE",
															"SHELL",
															"POWERSHELL",
														},
													},
													"localPath": &dcl.Property{
														Type:        "string",
														GoName:      "LocalPath",
														Description: "An absolute path to the executable on the VM.",
														Conflicts: []string{
															"gcsObject",
														},
													},
												},
											},
											"windowsExecStepConfig": &dcl.Property{
												Type:        "object",
												GoName:      "WindowsExecStepConfig",
												GoType:      "PatchDeploymentPatchConfigPostStepWindowsExecStepConfig",
												Description: "The ExecStepConfig for all Windows VMs targeted by the PatchJob.",
												Required: []string{
													"interpreter",
												},
												Properties: map[string]*dcl.Property{
													"allowedSuccessCodes": &dcl.Property{
														Type:        "array",
														GoName:      "AllowedSuccessCodes",
														Description: "Defaults to [0]. A list of possible return values that the execution can return to indicate a success.",
														SendEmpty:   true,
														ListType:    "list",
														Items: &dcl.Property{
															Type:   "integer",
															Format: "int64",
															GoType: "int64",
														},
													},
													"gcsObject": &dcl.Property{
														Type:        "object",
														GoName:      "GcsObject",
														GoType:      "PatchDeploymentPatchConfigPostStepWindowsExecStepConfigGcsObject",
														Description: "A Cloud Storage object containing the executable.",
														Conflicts: []string{
															"localPath",
														},
														Required: []string{
															"bucket",
															"object",
															"generationNumber",
														},
														Properties: map[string]*dcl.Property{
															"bucket": &dcl.Property{
																Type:        "string",
																GoName:      "Bucket",
																Description: "Required. Bucket of the Cloud Storage object.",
															},
															"generationNumber": &dcl.Property{
																Type:        "integer",
																Format:      "int64",
																GoName:      "GenerationNumber",
																Description: "Required. Generation number of the Cloud Storage object. This is used to ensure that the ExecStep specified by this PatchJob does not change.",
															},
															"object": &dcl.Property{
																Type:        "string",
																GoName:      "Object",
																Description: "Required. Name of the Cloud Storage object.",
															},
														},
													},
													"interpreter": &dcl.Property{
														Type:        "string",
														GoName:      "Interpreter",
														GoType:      "PatchDeploymentPatchConfigPostStepWindowsExecStepConfigInterpreterEnum",
														Description: "The script interpreter to use to run the script. If no interpreter is specified the script will be executed directly, which will likely only succeed for scripts with [shebang lines] (https://en.wikipedia.org/wiki/Shebang_(Unix)). Possible values: INTERPRETER_UNSPECIFIED, NONE, SHELL, POWERSHELL",
														Enum: []string{
															"INTERPRETER_UNSPECIFIED",
															"NONE",
															"SHELL",
															"POWERSHELL",
														},
													},
													"localPath": &dcl.Property{
														Type:        "string",
														GoName:      "LocalPath",
														Description: "An absolute path to the executable on the VM.",
														Conflicts: []string{
															"gcsObject",
														},
													},
												},
											},
										},
									},
									"preStep": &dcl.Property{
										Type:        "object",
										GoName:      "PreStep",
										GoType:      "PatchDeploymentPatchConfigPreStep",
										Description: "The `ExecStep` to run before the patch update.",
										Properties: map[string]*dcl.Property{
											"linuxExecStepConfig": &dcl.Property{
												Type:        "object",
												GoName:      "LinuxExecStepConfig",
												GoType:      "PatchDeploymentPatchConfigPreStepLinuxExecStepConfig",
												Description: "The ExecStepConfig for all Linux VMs targeted by the PatchJob.",
												Properties: map[string]*dcl.Property{
													"allowedSuccessCodes": &dcl.Property{
														Type:        "array",
														GoName:      "AllowedSuccessCodes",
														Description: "Defaults to [0]. A list of possible return values that the execution can return to indicate a success.",
														SendEmpty:   true,
														ListType:    "list",
														Items: &dcl.Property{
															Type:   "integer",
															Format: "int64",
															GoType: "int64",
														},
													},
													"gcsObject": &dcl.Property{
														Type:        "object",
														GoName:      "GcsObject",
														GoType:      "PatchDeploymentPatchConfigPreStepLinuxExecStepConfigGcsObject",
														Description: "A Cloud Storage object containing the executable.",
														Conflicts: []string{
															"localPath",
														},
														Required: []string{
															"bucket",
															"object",
															"generationNumber",
														},
														Properties: map[string]*dcl.Property{
															"bucket": &dcl.Property{
																Type:        "string",
																GoName:      "Bucket",
																Description: "Required. Bucket of the Cloud Storage object.",
															},
															"generationNumber": &dcl.Property{
																Type:        "integer",
																Format:      "int64",
																GoName:      "GenerationNumber",
																Description: "Required. Generation number of the Cloud Storage object. This is used to ensure that the ExecStep specified by this PatchJob does not change.",
															},
															"object": &dcl.Property{
																Type:        "string",
																GoName:      "Object",
																Description: "Required. Name of the Cloud Storage object.",
															},
														},
													},
													"interpreter": &dcl.Property{
														Type:        "string",
														GoName:      "Interpreter",
														GoType:      "PatchDeploymentPatchConfigPreStepLinuxExecStepConfigInterpreterEnum",
														Description: "The script interpreter to use to run the script. If no interpreter is specified the script will be executed directly, which will likely only succeed for scripts with [shebang lines] (https://en.wikipedia.org/wiki/Shebang_(Unix)). Possible values: INTERPRETER_UNSPECIFIED, NONE, SHELL, POWERSHELL",
														Enum: []string{
															"INTERPRETER_UNSPECIFIED",
															"NONE",
															"SHELL",
															"POWERSHELL",
														},
													},
													"localPath": &dcl.Property{
														Type:        "string",
														GoName:      "LocalPath",
														Description: "An absolute path to the executable on the VM.",
														Conflicts: []string{
															"gcsObject",
														},
													},
												},
											},
											"windowsExecStepConfig": &dcl.Property{
												Type:        "object",
												GoName:      "WindowsExecStepConfig",
												GoType:      "PatchDeploymentPatchConfigPreStepWindowsExecStepConfig",
												Description: "The ExecStepConfig for all Windows VMs targeted by the PatchJob.",
												Required: []string{
													"interpreter",
												},
												Properties: map[string]*dcl.Property{
													"allowedSuccessCodes": &dcl.Property{
														Type:        "array",
														GoName:      "AllowedSuccessCodes",
														Description: "Defaults to [0]. A list of possible return values that the execution can return to indicate a success.",
														SendEmpty:   true,
														ListType:    "list",
														Items: &dcl.Property{
															Type:   "integer",
															Format: "int64",
															GoType: "int64",
														},
													},
													"gcsObject": &dcl.Property{
														Type:        "object",
														GoName:      "GcsObject",
														GoType:      "PatchDeploymentPatchConfigPreStepWindowsExecStepConfigGcsObject",
														Description: "A Cloud Storage object containing the executable.",
														Conflicts: []string{
															"localPath",
														},
														Required: []string{
															"bucket",
															"object",
															"generationNumber",
														},
														Properties: map[string]*dcl.Property{
															"bucket": &dcl.Property{
																Type:        "string",
																GoName:      "Bucket",
																Description: "Required. Bucket of the Cloud Storage object.",
															},
															"generationNumber": &dcl.Property{
																Type:        "integer",
																Format:      "int64",
																GoName:      "GenerationNumber",
																Description: "Required. Generation number of the Cloud Storage object. This is used to ensure that the ExecStep specified by this PatchJob does not change.",
															},
															"object": &dcl.Property{
																Type:        "string",
																GoName:      "Object",
																Description: "Required. Name of the Cloud Storage object.",
															},
														},
													},
													"interpreter": &dcl.Property{
														Type:        "string",
														GoName:      "Interpreter",
														GoType:      "PatchDeploymentPatchConfigPreStepWindowsExecStepConfigInterpreterEnum",
														Description: "The script interpreter to use to run the script. If no interpreter is specified the script will be executed directly, which will likely only succeed for scripts with [shebang lines] (https://en.wikipedia.org/wiki/Shebang_(Unix)). Possible values: INTERPRETER_UNSPECIFIED, NONE, SHELL, POWERSHELL",
														Enum: []string{
															"INTERPRETER_UNSPECIFIED",
															"NONE",
															"SHELL",
															"POWERSHELL",
														},
													},
													"localPath": &dcl.Property{
														Type:        "string",
														GoName:      "LocalPath",
														Description: "An absolute path to the executable on the VM.",
														Conflicts: []string{
															"gcsObject",
														},
													},
												},
											},
										},
									},
									"rebootConfig": &dcl.Property{
										Type:        "string",
										GoName:      "RebootConfig",
										GoType:      "PatchDeploymentPatchConfigRebootConfigEnum",
										Description: "Post-patch reboot settings. Possible values: REBOOT_CONFIG_UNSPECIFIED, DEFAULT, ALWAYS, NEVER",
										Enum: []string{
											"REBOOT_CONFIG_UNSPECIFIED",
											"DEFAULT",
											"ALWAYS",
											"NEVER",
										},
									},
									"windowsUpdate": &dcl.Property{
										Type:        "object",
										GoName:      "WindowsUpdate",
										GoType:      "PatchDeploymentPatchConfigWindowsUpdate",
										Description: "Windows update settings. Use this override the default windows patch rules.",
										Properties: map[string]*dcl.Property{
											"classifications": &dcl.Property{
												Type:        "array",
												GoName:      "Classifications",
												Description: "Only apply updates of these windows update classifications. If empty, all updates are applied.",
												Conflicts: []string{
													"exclusivePatches",
												},
												SendEmpty: true,
												ListType:  "list",
												Items: &dcl.Property{
													Type:   "string",
													GoType: "PatchDeploymentPatchConfigWindowsUpdateClassificationsEnum",
													Enum: []string{
														"CLASSIFICATION_UNSPECIFIED",
														"CRITICAL",
														"SECURITY",
														"DEFINITION",
														"DRIVER",
														"FEATURE_PACK",
														"SERVICE_PACK",
														"TOOL",
														"UPDATE_ROLLUP",
														"UPDATE",
													},
												},
											},
											"excludes": &dcl.Property{
												Type:        "array",
												GoName:      "Excludes",
												Description: "List of KBs to exclude from update.",
												Conflicts: []string{
													"exclusivePatches",
												},
												SendEmpty: true,
												ListType:  "list",
												Items: &dcl.Property{
													Type:   "string",
													GoType: "string",
												},
											},
											"exclusivePatches": &dcl.Property{
												Type:        "array",
												GoName:      "ExclusivePatches",
												Description: "An exclusive list of kbs to be updated. These are the only patches that will be updated. This field must not be used with other patch configurations.",
												Conflicts: []string{
													"excludes",
												},
												SendEmpty: true,
												ListType:  "list",
												Items: &dcl.Property{
													Type:   "string",
													GoType: "string",
												},
											},
										},
									},
									"yum": &dcl.Property{
										Type:        "object",
										GoName:      "Yum",
										GoType:      "PatchDeploymentPatchConfigYum",
										Description: "Yum update settings. Use this setting to override the default `yum` patch rules.",
										Properties: map[string]*dcl.Property{
											"excludes": &dcl.Property{
												Type:        "array",
												GoName:      "Excludes",
												Description: "List of packages to exclude from update. These packages are excluded by using the yum `--exclude` flag.",
												Conflicts: []string{
													"exclusivePackages",
												},
												SendEmpty: true,
												ListType:  "list",
												Items: &dcl.Property{
													Type:   "string",
													GoType: "string",
												},
											},
											"exclusivePackages": &dcl.Property{
												Type:        "array",
												GoName:      "ExclusivePackages",
												Description: "An exclusive list of packages to be updated. These are the only packages that will be updated. If these packages are not installed, they will be ignored. This field must not be specified with any other patch configuration fields.",
												Conflicts: []string{
													"excludes",
												},
												SendEmpty: true,
												ListType:  "list",
												Items: &dcl.Property{
													Type:   "string",
													GoType: "string",
												},
											},
											"minimal": &dcl.Property{
												Type:        "boolean",
												GoName:      "Minimal",
												Description: "Will cause patch to run `yum update-minimal` instead.",
											},
											"security": &dcl.Property{
												Type:        "boolean",
												GoName:      "Security",
												Description: "Adds the `--security` flag to `yum update`. Not supported on all platforms.",
											},
										},
									},
									"zypper": &dcl.Property{
										Type:        "object",
										GoName:      "Zypper",
										GoType:      "PatchDeploymentPatchConfigZypper",
										Description: "Zypper update settings. Use this setting to override the default `zypper` patch rules.",
										Properties: map[string]*dcl.Property{
											"categories": &dcl.Property{
												Type:        "array",
												GoName:      "Categories",
												Description: "Install only patches with these categories. Common categories include security, recommended, and feature.",
												Conflicts: []string{
													"exclusivePatches",
												},
												SendEmpty: true,
												ListType:  "list",
												Items: &dcl.Property{
													Type:   "string",
													GoType: "string",
												},
											},
											"excludes": &dcl.Property{
												Type:        "array",
												GoName:      "Excludes",
												Description: "List of patches to exclude from update.",
												Conflicts: []string{
													"exclusivePatches",
												},
												SendEmpty: true,
												ListType:  "list",
												Items: &dcl.Property{
													Type:   "string",
													GoType: "string",
												},
											},
											"exclusivePatches": &dcl.Property{
												Type:        "array",
												GoName:      "ExclusivePatches",
												Description: "An exclusive list of patches to be updated. These are the only patches that will be installed using 'zypper patch patch:' command. This field must not be used with any other patch configuration fields.",
												Conflicts: []string{
													"excludes",
												},
												SendEmpty: true,
												ListType:  "list",
												Items: &dcl.Property{
													Type:   "string",
													GoType: "string",
												},
											},
											"severities": &dcl.Property{
												Type:        "array",
												GoName:      "Severities",
												Description: "Install only patches with these severities. Common severities include critical, important, moderate, and low.",
												Conflicts: []string{
													"exclusivePatches",
												},
												SendEmpty: true,
												ListType:  "list",
												Items: &dcl.Property{
													Type:   "string",
													GoType: "string",
												},
											},
											"withOptional": &dcl.Property{
												Type:        "boolean",
												GoName:      "WithOptional",
												Description: "Adds the `--with-optional` flag to `zypper patch`.",
												Conflicts: []string{
													"exclusivePatches",
												},
											},
											"withUpdate": &dcl.Property{
												Type:        "boolean",
												GoName:      "WithUpdate",
												Description: "Adds the `--with-update` flag, to `zypper patch`.",
												Conflicts: []string{
													"exclusivePatches",
												},
											},
										},
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
							"recurringSchedule": &dcl.Property{
								Type:        "object",
								GoName:      "RecurringSchedule",
								GoType:      "PatchDeploymentRecurringSchedule",
								Description: "Required. Schedule recurring executions.",
								Conflicts: []string{
									"oneTimeSchedule",
								},
								Required: []string{
									"timeZone",
									"timeOfDay",
									"frequency",
								},
								Properties: map[string]*dcl.Property{
									"endTime": &dcl.Property{
										Type:        "string",
										Format:      "date-time",
										GoName:      "EndTime",
										Description: "Optional. The end time at which a recurring patch deployment schedule is no longer active.",
									},
									"frequency": &dcl.Property{
										Type:        "string",
										GoName:      "Frequency",
										GoType:      "PatchDeploymentRecurringScheduleFrequencyEnum",
										Description: "Required. The frequency unit of this recurring schedule. Possible values: FREQUENCY_UNSPECIFIED, WEEKLY, MONTHLY, DAILY",
										Enum: []string{
											"FREQUENCY_UNSPECIFIED",
											"WEEKLY",
											"MONTHLY",
											"DAILY",
										},
									},
									"lastExecuteTime": &dcl.Property{
										Type:        "string",
										Format:      "date-time",
										GoName:      "LastExecuteTime",
										ReadOnly:    true,
										Description: "Output only. The time the last patch job ran successfully.",
									},
									"monthly": &dcl.Property{
										Type:        "object",
										GoName:      "Monthly",
										GoType:      "PatchDeploymentRecurringScheduleMonthly",
										Description: "Required. Schedule with monthly executions.",
										Conflicts: []string{
											"weekly",
										},
										Properties: map[string]*dcl.Property{
											"monthDay": &dcl.Property{
												Type:        "integer",
												Format:      "int64",
												GoName:      "MonthDay",
												Description: "Required. One day of the month. 1-31 indicates the 1st to the 31st day. -1 indicates the last day of the month. Months without the target day will be skipped. For example, a schedule to run \"every month on the 31st\" will not run in February, April, June, etc.",
												Conflicts: []string{
													"weekDayOfMonth",
												},
											},
											"weekDayOfMonth": &dcl.Property{
												Type:        "object",
												GoName:      "WeekDayOfMonth",
												GoType:      "PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonth",
												Description: "Required. Week day in a month.",
												Conflicts: []string{
													"monthDay",
												},
												Required: []string{
													"weekOrdinal",
													"dayOfWeek",
												},
												Properties: map[string]*dcl.Property{
													"dayOfWeek": &dcl.Property{
														Type:        "string",
														GoName:      "DayOfWeek",
														GoType:      "PatchDeploymentRecurringScheduleMonthlyWeekDayOfMonthDayOfWeekEnum",
														Description: "Required. A day of the week. Possible values: DAY_OF_WEEK_UNSPECIFIED, MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY, SUNDAY",
														Enum: []string{
															"DAY_OF_WEEK_UNSPECIFIED",
															"MONDAY",
															"TUESDAY",
															"WEDNESDAY",
															"THURSDAY",
															"FRIDAY",
															"SATURDAY",
															"SUNDAY",
														},
													},
													"weekOrdinal": &dcl.Property{
														Type:        "integer",
														Format:      "int64",
														GoName:      "WeekOrdinal",
														Description: "Required. Week number in a month. 1-4 indicates the 1st to 4th week of the month. -1 indicates the last week of the month.",
													},
												},
											},
										},
									},
									"nextExecuteTime": &dcl.Property{
										Type:        "string",
										Format:      "date-time",
										GoName:      "NextExecuteTime",
										ReadOnly:    true,
										Description: "Output only. The time the next patch job is scheduled to run.",
									},
									"startTime": &dcl.Property{
										Type:        "string",
										Format:      "date-time",
										GoName:      "StartTime",
										Description: "Optional. The time that the recurring schedule becomes effective. Defaults to `create_time` of the patch deployment.",
									},
									"timeOfDay": &dcl.Property{
										Type:        "object",
										GoName:      "TimeOfDay",
										GoType:      "PatchDeploymentRecurringScheduleTimeOfDay",
										Description: "Required. Time of the day to run a recurring deployment.",
										SendEmpty:   true,
										Properties: map[string]*dcl.Property{
											"hours": &dcl.Property{
												Type:        "integer",
												Format:      "int64",
												GoName:      "Hours",
												Description: "Hours of day in 24 hour format. Should be from 0 to 23. An API may choose to allow the value \"24:00:00\" for scenarios like business closing time.",
											},
											"minutes": &dcl.Property{
												Type:        "integer",
												Format:      "int64",
												GoName:      "Minutes",
												Description: "Minutes of hour of day. Must be from 0 to 59.",
											},
											"nanos": &dcl.Property{
												Type:        "integer",
												Format:      "int64",
												GoName:      "Nanos",
												Description: "Fractions of seconds in nanoseconds. Must be from 0 to 999,999,999.",
											},
											"seconds": &dcl.Property{
												Type:        "integer",
												Format:      "int64",
												GoName:      "Seconds",
												Description: "Seconds of minutes of the time. Must normally be from 0 to 59. An API may allow the value 60 if it allows leap-seconds.",
											},
										},
									},
									"timeZone": &dcl.Property{
										Type:        "object",
										GoName:      "TimeZone",
										GoType:      "PatchDeploymentRecurringScheduleTimeZone",
										Description: "Required. Defines the time zone that `time_of_day` is relative to. The rules for daylight saving time are determined by the chosen time zone.",
										Properties: map[string]*dcl.Property{
											"id": &dcl.Property{
												Type:        "string",
												GoName:      "Id",
												Description: "IANA Time Zone Database time zone, e.g. \"America/New_York\".",
											},
											"version": &dcl.Property{
												Type:        "string",
												GoName:      "Version",
												Description: "Optional. IANA Time Zone Database version number, e.g. \"2019a\".",
											},
										},
									},
									"weekly": &dcl.Property{
										Type:        "object",
										GoName:      "Weekly",
										GoType:      "PatchDeploymentRecurringScheduleWeekly",
										Description: "Required. Schedule with weekly executions.",
										Conflicts: []string{
											"monthly",
										},
										Required: []string{
											"dayOfWeek",
										},
										Properties: map[string]*dcl.Property{
											"dayOfWeek": &dcl.Property{
												Type:        "string",
												GoName:      "DayOfWeek",
												GoType:      "PatchDeploymentRecurringScheduleWeeklyDayOfWeekEnum",
												Description: "Required. Day of the week. Possible values: DAY_OF_WEEK_UNSPECIFIED, MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY, SUNDAY",
												Enum: []string{
													"DAY_OF_WEEK_UNSPECIFIED",
													"MONDAY",
													"TUESDAY",
													"WEDNESDAY",
													"THURSDAY",
													"FRIDAY",
													"SATURDAY",
													"SUNDAY",
												},
											},
										},
									},
								},
							},
							"rollout": &dcl.Property{
								Type:        "object",
								GoName:      "Rollout",
								GoType:      "PatchDeploymentRollout",
								Description: "Optional. Rollout strategy of the patch job.",
								Required: []string{
									"mode",
									"disruptionBudget",
								},
								Properties: map[string]*dcl.Property{
									"disruptionBudget": &dcl.Property{
										Type:        "object",
										GoName:      "DisruptionBudget",
										GoType:      "PatchDeploymentRolloutDisruptionBudget",
										Description: "The maximum number (or percentage) of VMs per zone to disrupt at any given moment. The number of VMs calculated from multiplying the percentage by the total number of VMs in a zone is rounded up. During patching, a VM is considered disrupted from the time the agent is notified to begin until patching has completed. This disruption time includes the time to complete reboot and any post-patch steps. A VM contributes to the disruption budget if its patching operation fails either when applying the patches, running pre or post patch steps, or if it fails to respond with a success notification before timing out. VMs that are not running or do not have an active agent do not count toward this disruption budget. For zone-by-zone rollouts, if the disruption budget in a zone is exceeded, the patch job stops, because continuing to the next zone requires completion of the patch process in the previous zone. For example, if the disruption budget has a fixed value of `10`, and 8 VMs fail to patch in the current zone, the patch job continues to patch 2 VMs at a time until the zone is completed. When that zone is completed successfully, patching begins with 10 VMs at a time in the next zone. If 10 VMs in the next zone fail to patch, the patch job stops.",
										Properties: map[string]*dcl.Property{
											"fixed": &dcl.Property{
												Type:        "integer",
												Format:      "int64",
												GoName:      "Fixed",
												Description: "Specifies a fixed value.",
												Conflicts: []string{
													"percent",
												},
											},
											"percent": &dcl.Property{
												Type:        "integer",
												Format:      "int64",
												GoName:      "Percent",
												Description: "Specifies the relative value defined as a percentage, which will be multiplied by a reference value.",
												Conflicts: []string{
													"fixed",
												},
											},
										},
									},
									"mode": &dcl.Property{
										Type:        "string",
										GoName:      "Mode",
										GoType:      "PatchDeploymentRolloutModeEnum",
										Description: "Mode of the patch rollout. Possible values: MODE_UNSPECIFIED, VALIDATION, ENFORCEMENT",
										Enum: []string{
											"MODE_UNSPECIFIED",
											"VALIDATION",
											"ENFORCEMENT",
										},
									},
								},
							},
							"updateTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "UpdateTime",
								ReadOnly:    true,
								Description: "Output only. Time the patch deployment was last updated. Timestamp is in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.",
								Immutable:   true,
							},
						},
					},
				},
			},
		},
	}
}
