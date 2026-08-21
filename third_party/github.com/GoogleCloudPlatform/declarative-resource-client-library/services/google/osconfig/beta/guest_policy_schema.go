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

func DCLGuestPolicySchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "OSConfig/GuestPolicy",
			Description: "The OSConfig GuestPolicy resource",
			StructName:  "GuestPolicy",
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a GuestPolicy",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "guestPolicy",
						Required:    true,
						Description: "A full instance of a GuestPolicy",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a GuestPolicy",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "guestPolicy",
						Required:    true,
						Description: "A full instance of a GuestPolicy",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a GuestPolicy",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "guestPolicy",
						Required:    true,
						Description: "A full instance of a GuestPolicy",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all GuestPolicy",
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
				Description: "The function used to list information about many GuestPolicy",
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
				"GuestPolicy": &dcl.Component{
					Title:           "GuestPolicy",
					ID:              "projects/{{project}}/guestPolicies/{{name}}",
					ParentContainer: "project",
					HasCreate:       true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"name",
							"project",
						},
						Properties: map[string]*dcl.Property{
							"assignment": &dcl.Property{
								Type:        "object",
								GoName:      "Assignment",
								GoType:      "GuestPolicyAssignment",
								Description: "Specifies the VMs that are assigned this policy. This allows you to target sets or groups of VMs by different parameters such as labels, names, OS, or zones. Empty assignments will target ALL VMs underneath this policy. Conflict Management Policies that exist higher up in the resource hierarchy (closer to the Org) will override those lower down if there is a conflict. At the same level in the resource hierarchy (ie. within a project), the service will prevent the creation of multiple policies that conflict with each other. If there are multiple policies that specify the same config (eg. package, software recipe, repository, etc.), the service will ensure that no VM could potentially receive instructions from both policies. To create multiple policies that specify different versions of a package or different configs for different Operating Systems, each policy must be mutually exclusive in their targeting according to labels, OS, or other criteria. Different configs are identified for conflicts in different ways. Packages are identified by their name and the package manager(s) they target. Package repositories are identified by their unique id where applicable. Some package managers don't have a unique identifier for repositories and where that's the case, no uniqueness is validated by the service. Note that if OS Inventory is disabled, a VM will not be assigned a policy that targets by OS because the service will see this VM's OS as unknown.",
								Properties: map[string]*dcl.Property{
									"groupLabels": &dcl.Property{
										Type:        "array",
										GoName:      "GroupLabels",
										Description: "Targets instances matching at least one of these label sets. This allows an assignment to target disparate groups, for example \"env=prod or env=staging\".",
										SendEmpty:   true,
										ListType:    "list",
										Items: &dcl.Property{
											Type:   "object",
											GoType: "GuestPolicyAssignmentGroupLabels",
											Properties: map[string]*dcl.Property{
												"labels": &dcl.Property{
													Type: "object",
													AdditionalProperties: &dcl.Property{
														Type: "string",
													},
													GoName:      "Labels",
													Description: "Google Compute Engine instance labels that must be present for an instance to be included in this assignment group.",
												},
											},
										},
									},
									"instanceNamePrefixes": &dcl.Property{
										Type:        "array",
										GoName:      "InstanceNamePrefixes",
										Description: "Targets VM instances whose name starts with one of these prefixes. Like labels, this is another way to group VM instances when targeting configs, for example prefix=\"prod-\". Only supported for project-level policies.",
										SendEmpty:   true,
										ListType:    "list",
										Items: &dcl.Property{
											Type:   "string",
											GoType: "string",
										},
									},
									"instances": &dcl.Property{
										Type:        "array",
										GoName:      "Instances",
										Description: "Targets any of the instances specified. Instances are specified by their URI in the form `zones/[ZONE]/instances/[INSTANCE_NAME]`. Instance targeting is uncommon and is supported to facilitate the management of changes by the instance or to target specific VM instances for development and testing. Only supported for project-level policies and must reference instances within this project.",
										SendEmpty:   true,
										ListType:    "list",
										Items: &dcl.Property{
											Type:   "string",
											GoType: "string",
											ResourceReferences: []*dcl.PropertyResourceReference{
												&dcl.PropertyResourceReference{
													Resource: "Compute/Instance",
													Field:    "selfLink",
												},
											},
										},
									},
									"osTypes": &dcl.Property{
										Type:        "array",
										GoName:      "OSTypes",
										Description: "Targets VM instances matching at least one of the following OS types. VM instances must match all supplied criteria for a given OsType to be included.",
										SendEmpty:   true,
										ListType:    "list",
										Items: &dcl.Property{
											Type:   "object",
											GoType: "GuestPolicyAssignmentOSTypes",
											Properties: map[string]*dcl.Property{
												"osArchitecture": &dcl.Property{
													Type:        "string",
													GoName:      "OSArchitecture",
													Description: "Targets VM instances with OS Inventory enabled and having the following OS architecture.",
												},
												"osShortName": &dcl.Property{
													Type:        "string",
													GoName:      "OSShortName",
													Description: "Targets VM instances with OS Inventory enabled and having the following OS short name, for example \"debian\" or \"windows\".",
												},
												"osVersion": &dcl.Property{
													Type:        "string",
													GoName:      "OSVersion",
													Description: "Targets VM instances with OS Inventory enabled and having the following following OS version.",
												},
											},
										},
									},
									"zones": &dcl.Property{
										Type:        "array",
										GoName:      "Zones",
										Description: "Targets instances in any of these zones. Leave empty to target instances in any zone. Zonal targeting is uncommon and is supported to facilitate the management of changes by zone.",
										SendEmpty:   true,
										ListType:    "list",
										Items: &dcl.Property{
											Type:   "string",
											GoType: "string",
										},
									},
								},
							},
							"createTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "CreateTime",
								ReadOnly:    true,
								Description: "Output only. Time this GuestPolicy was created.",
								Immutable:   true,
							},
							"description": &dcl.Property{
								Type:        "string",
								GoName:      "Description",
								Description: "Description of the GuestPolicy. Length of the description is limited to 1024 characters.",
							},
							"etag": &dcl.Property{
								Type:        "string",
								GoName:      "Etag",
								ReadOnly:    true,
								Description: "The etag for this GuestPolicy. If this is provided on update, it must match the server's etag.",
								Immutable:   true,
							},
							"name": &dcl.Property{
								Type:        "string",
								GoName:      "Name",
								Description: "Unique name of the resource in this project using the form: `projects/{project_id}/guestPolicies/{guest_policy_id}`.",
							},
							"packageRepositories": &dcl.Property{
								Type:        "array",
								GoName:      "PackageRepositories",
								Description: "List of package repository configurations assigned to the VM instance.",
								SendEmpty:   true,
								ListType:    "list",
								Items: &dcl.Property{
									Type:   "object",
									GoType: "GuestPolicyPackageRepositories",
									Properties: map[string]*dcl.Property{
										"apt": &dcl.Property{
											Type:        "object",
											GoName:      "Apt",
											GoType:      "GuestPolicyPackageRepositoriesApt",
											Description: "An Apt Repository.",
											Conflicts: []string{
												"goo",
												"yum",
												"zypper",
											},
											Required: []string{
												"uri",
												"distribution",
											},
											Properties: map[string]*dcl.Property{
												"archiveType": &dcl.Property{
													Type:        "string",
													GoName:      "ArchiveType",
													GoType:      "GuestPolicyPackageRepositoriesAptArchiveTypeEnum",
													Description: "Type of archive files in this repository. The default behavior is DEB. Possible values: ARCHIVE_TYPE_UNSPECIFIED, DEB, DEB_SRC",
													Enum: []string{
														"ARCHIVE_TYPE_UNSPECIFIED",
														"DEB",
														"DEB_SRC",
													},
												},
												"components": &dcl.Property{
													Type:        "array",
													GoName:      "Components",
													Description: "Required. List of components for this repository. Must contain at least one item.",
													SendEmpty:   true,
													ListType:    "list",
													Items: &dcl.Property{
														Type:   "string",
														GoType: "string",
													},
												},
												"distribution": &dcl.Property{
													Type:        "string",
													GoName:      "Distribution",
													Description: "Required. Distribution of this repository.",
												},
												"gpgKey": &dcl.Property{
													Type:        "string",
													GoName:      "GpgKey",
													Description: "URI of the key file for this repository. The agent maintains a keyring at `/etc/apt/trusted.gpg.d/osconfig_agent_managed.gpg` containing all the keys in any applied guest policy.",
												},
												"uri": &dcl.Property{
													Type:        "string",
													GoName:      "Uri",
													Description: "Required. URI for this repository.",
												},
											},
										},
										"goo": &dcl.Property{
											Type:        "object",
											GoName:      "Goo",
											GoType:      "GuestPolicyPackageRepositoriesGoo",
											Description: "A Goo Repository.",
											Conflicts: []string{
												"apt",
												"yum",
												"zypper",
											},
											Required: []string{
												"name",
												"url",
											},
											Properties: map[string]*dcl.Property{
												"name": &dcl.Property{
													Type:        "string",
													GoName:      "Name",
													Description: "Required. The name of the repository.",
												},
												"url": &dcl.Property{
													Type:        "string",
													GoName:      "Url",
													Description: "Required. The url of the repository.",
												},
											},
										},
										"yum": &dcl.Property{
											Type:        "object",
											GoName:      "Yum",
											GoType:      "GuestPolicyPackageRepositoriesYum",
											Description: "A Yum Repository.",
											Conflicts: []string{
												"apt",
												"goo",
												"zypper",
											},
											Required: []string{
												"id",
												"baseUrl",
											},
											Properties: map[string]*dcl.Property{
												"baseUrl": &dcl.Property{
													Type:        "string",
													GoName:      "BaseUrl",
													Description: "Required. The location of the repository directory.",
												},
												"displayName": &dcl.Property{
													Type:        "string",
													GoName:      "DisplayName",
													Description: "The display name of the repository.",
												},
												"gpgKeys": &dcl.Property{
													Type:        "array",
													GoName:      "GpgKeys",
													Description: "URIs of GPG keys.",
													SendEmpty:   true,
													ListType:    "list",
													Items: &dcl.Property{
														Type:   "string",
														GoType: "string",
													},
												},
												"id": &dcl.Property{
													Type:        "string",
													GoName:      "Id",
													Description: "Required. A one word, unique name for this repository. This is the `repo id` in the Yum config file and also the `display_name` if `display_name` is omitted. This id is also used as the unique identifier when checking for guest policy conflicts.",
												},
											},
										},
										"zypper": &dcl.Property{
											Type:        "object",
											GoName:      "Zypper",
											GoType:      "GuestPolicyPackageRepositoriesZypper",
											Description: "A Zypper Repository.",
											Conflicts: []string{
												"apt",
												"goo",
												"yum",
											},
											Required: []string{
												"id",
												"baseUrl",
											},
											Properties: map[string]*dcl.Property{
												"baseUrl": &dcl.Property{
													Type:        "string",
													GoName:      "BaseUrl",
													Description: "Required. The location of the repository directory.",
												},
												"displayName": &dcl.Property{
													Type:        "string",
													GoName:      "DisplayName",
													Description: "The display name of the repository.",
												},
												"gpgKeys": &dcl.Property{
													Type:        "array",
													GoName:      "GpgKeys",
													Description: "URIs of GPG keys.",
													SendEmpty:   true,
													ListType:    "list",
													Items: &dcl.Property{
														Type:   "string",
														GoType: "string",
													},
												},
												"id": &dcl.Property{
													Type:        "string",
													GoName:      "Id",
													Description: "Required. A one word, unique name for this repository. This is the `repo id` in the zypper config file and also the `display_name` if `display_name` is omitted. This id is also used as the unique identifier when checking for guest policy conflicts.",
												},
											},
										},
									},
								},
							},
							"packages": &dcl.Property{
								Type:        "array",
								GoName:      "Packages",
								Description: "List of package configurations assigned to the VM instance.",
								SendEmpty:   true,
								ListType:    "list",
								Items: &dcl.Property{
									Type:   "object",
									GoType: "GuestPolicyPackages",
									Properties: map[string]*dcl.Property{
										"desiredState": &dcl.Property{
											Type:        "string",
											GoName:      "DesiredState",
											GoType:      "GuestPolicyPackagesDesiredStateEnum",
											Description: "The desired_state the agent should maintain for this package. The default is to ensure the package is installed. Possible values: DESIRED_STATE_UNSPECIFIED, INSTALLED, REMOVED",
											Enum: []string{
												"DESIRED_STATE_UNSPECIFIED",
												"INSTALLED",
												"REMOVED",
											},
										},
										"manager": &dcl.Property{
											Type:        "string",
											GoName:      "Manager",
											GoType:      "GuestPolicyPackagesManagerEnum",
											Description: "Type of package manager that can be used to install this package. If a system does not have the package manager, the package is not installed or removed no error message is returned. By default, or if you specify `ANY`, the agent attempts to install and remove this package using the default package manager. This is useful when creating a policy that applies to different types of systems. The default behavior is ANY. Possible values: MANAGER_UNSPECIFIED, ANY, APT, YUM, ZYPPER, GOO",
											Enum: []string{
												"MANAGER_UNSPECIFIED",
												"ANY",
												"APT",
												"YUM",
												"ZYPPER",
												"GOO",
											},
										},
										"name": &dcl.Property{
											Type:        "string",
											GoName:      "Name",
											Description: "Required. The name of the package. A package is uniquely identified for conflict validation by checking the package name and the manager(s) that the package targets.",
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
							"recipes": &dcl.Property{
								Type:        "array",
								GoName:      "Recipes",
								Description: "Optional. A list of Recipes to install on the VM.",
								SendEmpty:   true,
								ListType:    "list",
								Items: &dcl.Property{
									Type:   "object",
									GoType: "GuestPolicyRecipes",
									Properties: map[string]*dcl.Property{
										"artifacts": &dcl.Property{
											Type:        "array",
											GoName:      "Artifacts",
											Description: "Resources available to be used in the steps in the recipe.",
											SendEmpty:   true,
											ListType:    "list",
											Items: &dcl.Property{
												Type:   "object",
												GoType: "GuestPolicyRecipesArtifacts",
												Properties: map[string]*dcl.Property{
													"allowInsecure": &dcl.Property{
														Type:        "boolean",
														GoName:      "AllowInsecure",
														Description: "Defaults to false. When false, recipes are subject to validations based on the artifact type: Remote: A checksum must be specified, and only protocols with transport-layer security are permitted. GCS: An object generation number must be specified.",
													},
													"gcs": &dcl.Property{
														Type:        "object",
														GoName:      "Gcs",
														GoType:      "GuestPolicyRecipesArtifactsGcs",
														Description: "A Google Cloud Storage artifact.",
														Properties: map[string]*dcl.Property{
															"bucket": &dcl.Property{
																Type:        "string",
																GoName:      "Bucket",
																Description: "Bucket of the Google Cloud Storage object. Given an example URL: `https://storage.googleapis.com/my-bucket/foo/bar#1234567` this value would be `my-bucket`.",
																ResourceReferences: []*dcl.PropertyResourceReference{
																	&dcl.PropertyResourceReference{
																		Resource: "Storage/Bucket",
																		Field:    "name",
																	},
																},
															},
															"generation": &dcl.Property{
																Type:        "integer",
																Format:      "int64",
																GoName:      "Generation",
																Description: "Must be provided if allow_insecure is false. Generation number of the Google Cloud Storage object. `https://storage.googleapis.com/my-bucket/foo/bar#1234567` this value would be `1234567`.",
															},
															"object": &dcl.Property{
																Type:        "string",
																GoName:      "Object",
																Description: "Name of the Google Cloud Storage object. As specified [here] (https://cloud.google.com/storage/docs/naming#objectnames) Given an example URL: `https://storage.googleapis.com/my-bucket/foo/bar#1234567` this value would be `foo/bar`.",
															},
														},
													},
													"id": &dcl.Property{
														Type:        "string",
														GoName:      "Id",
														Description: "Required. Id of the artifact, which the installation and update steps of this recipe can reference. Artifacts in a recipe cannot have the same id.",
													},
													"remote": &dcl.Property{
														Type:        "object",
														GoName:      "Remote",
														GoType:      "GuestPolicyRecipesArtifactsRemote",
														Description: "A generic remote artifact.",
														Properties: map[string]*dcl.Property{
															"checksum": &dcl.Property{
																Type:        "string",
																GoName:      "Checksum",
																Description: "Must be provided if `allow_insecure` is `false`. SHA256 checksum in hex format, to compare to the checksum of the artifact. If the checksum is not empty and it doesn't match the artifact then the recipe installation fails before running any of the steps.",
															},
															"uri": &dcl.Property{
																Type:        "string",
																GoName:      "Uri",
																Description: "URI from which to fetch the object. It should contain both the protocol and path following the format: {protocol}://{location}.",
															},
														},
													},
												},
											},
										},
										"desiredState": &dcl.Property{
											Type:        "string",
											GoName:      "DesiredState",
											GoType:      "GuestPolicyRecipesDesiredStateEnum",
											Description: "Default is INSTALLED. The desired state the agent should maintain for this recipe. INSTALLED: The software recipe is installed on the instance but won't be updated to new versions. UPDATED: The software recipe is installed on the instance. The recipe is updated to a higher version, if a higher version of the recipe is assigned to this instance. REMOVE: Remove is unsupported for software recipes and attempts to create or update a recipe to the REMOVE state is rejected. Possible values: DESIRED_STATE_UNSPECIFIED, INSTALLED, REMOVED",
											Enum: []string{
												"DESIRED_STATE_UNSPECIFIED",
												"INSTALLED",
												"REMOVED",
											},
										},
										"installSteps": &dcl.Property{
											Type:        "array",
											GoName:      "InstallSteps",
											Description: "Actions to be taken for installing this recipe. On failure it stops executing steps and does not attempt another installation. Any steps taken (including partially completed steps) are not rolled back.",
											Conflicts: []string{
												"updateSteps",
											},
											SendEmpty: true,
											ListType:  "list",
											Items: &dcl.Property{
												Type:   "object",
												GoType: "GuestPolicyRecipesInstallSteps",
												Properties: map[string]*dcl.Property{
													"archiveExtraction": &dcl.Property{
														Type:        "object",
														GoName:      "ArchiveExtraction",
														GoType:      "GuestPolicyRecipesInstallStepsArchiveExtraction",
														Description: "Extracts an archive into the specified directory.",
														Properties: map[string]*dcl.Property{
															"artifactId": &dcl.Property{
																Type:        "string",
																GoName:      "ArtifactId",
																Description: "Required. The id of the relevant artifact in the recipe.",
															},
															"destination": &dcl.Property{
																Type:        "string",
																GoName:      "Destination",
																Description: "Directory to extract archive to. Defaults to `/` on Linux or `C:` on Windows.",
															},
															"type": &dcl.Property{
																Type:        "string",
																GoName:      "Type",
																GoType:      "GuestPolicyRecipesInstallStepsArchiveExtractionTypeEnum",
																Description: "Required. The type of the archive to extract. Possible values: TYPE_UNSPECIFIED, VALIDATION, DESIRED_STATE_CHECK, DESIRED_STATE_ENFORCEMENT, DESIRED_STATE_CHECK_POST_ENFORCEMENT",
																Enum: []string{
																	"TYPE_UNSPECIFIED",
																	"VALIDATION",
																	"DESIRED_STATE_CHECK",
																	"DESIRED_STATE_ENFORCEMENT",
																	"DESIRED_STATE_CHECK_POST_ENFORCEMENT",
																},
															},
														},
													},
													"dpkgInstallation": &dcl.Property{
														Type:        "object",
														GoName:      "DpkgInstallation",
														GoType:      "GuestPolicyRecipesInstallStepsDpkgInstallation",
														Description: "Installs a deb file via dpkg.",
														Properties: map[string]*dcl.Property{
															"artifactId": &dcl.Property{
																Type:        "string",
																GoName:      "ArtifactId",
																Description: "Required. The id of the relevant artifact in the recipe.",
															},
														},
													},
													"fileCopy": &dcl.Property{
														Type:        "object",
														GoName:      "FileCopy",
														GoType:      "GuestPolicyRecipesInstallStepsFileCopy",
														Description: "Copies a file onto the instance.",
														Properties: map[string]*dcl.Property{
															"artifactId": &dcl.Property{
																Type:        "string",
																GoName:      "ArtifactId",
																Description: "Required. The id of the relevant artifact in the recipe.",
															},
															"destination": &dcl.Property{
																Type:        "string",
																GoName:      "Destination",
																Description: "Required. The absolute path on the instance to put the file.",
															},
															"overwrite": &dcl.Property{
																Type:        "boolean",
																GoName:      "Overwrite",
																Description: "Whether to allow this step to overwrite existing files. If this is false and the file already exists the file is not overwritten and the step is considered a success. Defaults to false.",
															},
															"permissions": &dcl.Property{
																Type:        "string",
																GoName:      "Permissions",
																Description: "Consists of three octal digits which represent, in order, the permissions of the owner, group, and other users for the file (similarly to the numeric mode used in the linux chmod utility). Each digit represents a three bit number with the 4 bit corresponding to the read permissions, the 2 bit corresponds to the write bit, and the one bit corresponds to the execute permission. Default behavior is 755. Below are some examples of permissions and their associated values: read, write, and execute: 7 read and execute: 5 read and write: 6 read only: 4",
															},
														},
													},
													"fileExec": &dcl.Property{
														Type:        "object",
														GoName:      "FileExec",
														GoType:      "GuestPolicyRecipesInstallStepsFileExec",
														Description: "Executes an artifact or local file.",
														Properties: map[string]*dcl.Property{
															"allowedExitCodes": &dcl.Property{
																Type:        "array",
																GoName:      "AllowedExitCodes",
																Description: "Defaults to [0]. A list of possible return values that the program can return to indicate a success.",
																SendEmpty:   true,
																ListType:    "list",
																Items: &dcl.Property{
																	Type:   "integer",
																	Format: "int64",
																	GoType: "int64",
																},
															},
															"args": &dcl.Property{
																Type:        "array",
																GoName:      "Args",
																Description: "Arguments to be passed to the provided executable.",
																SendEmpty:   true,
																ListType:    "list",
																Items: &dcl.Property{
																	Type:   "string",
																	GoType: "string",
																},
															},
															"artifactId": &dcl.Property{
																Type:        "string",
																GoName:      "ArtifactId",
																Description: "The id of the relevant artifact in the recipe.",
															},
															"localPath": &dcl.Property{
																Type:        "string",
																GoName:      "LocalPath",
																Description: "The absolute path of the file on the local filesystem.",
															},
														},
													},
													"msiInstallation": &dcl.Property{
														Type:        "object",
														GoName:      "MsiInstallation",
														GoType:      "GuestPolicyRecipesInstallStepsMsiInstallation",
														Description: "Installs an MSI file.",
														Properties: map[string]*dcl.Property{
															"allowedExitCodes": &dcl.Property{
																Type:        "array",
																GoName:      "AllowedExitCodes",
																Description: "Return codes that indicate that the software installed or updated successfully. Behaviour defaults to [0]",
																SendEmpty:   true,
																ListType:    "list",
																Items: &dcl.Property{
																	Type:   "integer",
																	Format: "int64",
																	GoType: "int64",
																},
															},
															"artifactId": &dcl.Property{
																Type:        "string",
																GoName:      "ArtifactId",
																Description: "Required. The id of the relevant artifact in the recipe.",
															},
															"flags": &dcl.Property{
																Type:        "array",
																GoName:      "Flags",
																Description: "The flags to use when installing the MSI defaults to [\"/i\"] (i.e. the install flag).",
																SendEmpty:   true,
																ListType:    "list",
																Items: &dcl.Property{
																	Type:   "string",
																	GoType: "string",
																},
															},
														},
													},
													"rpmInstallation": &dcl.Property{
														Type:        "object",
														GoName:      "RpmInstallation",
														GoType:      "GuestPolicyRecipesInstallStepsRpmInstallation",
														Description: "Installs an rpm file via the rpm utility.",
														Properties: map[string]*dcl.Property{
															"artifactId": &dcl.Property{
																Type:        "string",
																GoName:      "ArtifactId",
																Description: "Required. The id of the relevant artifact in the recipe.",
															},
														},
													},
													"scriptRun": &dcl.Property{
														Type:        "object",
														GoName:      "ScriptRun",
														GoType:      "GuestPolicyRecipesInstallStepsScriptRun",
														Description: "Runs commands in a shell.",
														Properties: map[string]*dcl.Property{
															"allowedExitCodes": &dcl.Property{
																Type:        "array",
																GoName:      "AllowedExitCodes",
																Description: "Return codes that indicate that the software installed or updated successfully. Behaviour defaults to [0]",
																SendEmpty:   true,
																ListType:    "list",
																Items: &dcl.Property{
																	Type:   "integer",
																	Format: "int64",
																	GoType: "int64",
																},
															},
															"interpreter": &dcl.Property{
																Type:        "string",
																GoName:      "Interpreter",
																GoType:      "GuestPolicyRecipesInstallStepsScriptRunInterpreterEnum",
																Description: "The script interpreter to use to run the script. If no interpreter is specified the script is executed directly, which likely only succeed for scripts with [shebang lines](https://en.wikipedia.org/wiki/Shebang_(Unix)). Possible values: INTERPRETER_UNSPECIFIED, NONE, SHELL, POWERSHELL",
																Enum: []string{
																	"INTERPRETER_UNSPECIFIED",
																	"NONE",
																	"SHELL",
																	"POWERSHELL",
																},
															},
															"script": &dcl.Property{
																Type:        "string",
																GoName:      "Script",
																Description: "Required. The shell script to be executed.",
															},
														},
													},
												},
											},
										},
										"name": &dcl.Property{
											Type:        "string",
											GoName:      "Name",
											Description: "Required. Unique identifier for the recipe. Only one recipe with a given name is installed on an instance. Names are also used to identify resources which helps to determine whether guest policies have conflicts. This means that requests to create multiple recipes with the same name and version are rejected since they could potentially have conflicting assignments.",
										},
										"updateSteps": &dcl.Property{
											Type:        "array",
											GoName:      "UpdateSteps",
											Description: "Actions to be taken for updating this recipe. On failure it stops executing steps and does not attempt another update for this recipe. Any steps taken (including partially completed steps) are not rolled back.",
											Conflicts: []string{
												"installSteps",
											},
											SendEmpty: true,
											ListType:  "list",
											Items: &dcl.Property{
												Type:   "object",
												GoType: "GuestPolicyRecipesUpdateSteps",
												Properties: map[string]*dcl.Property{
													"archiveExtraction": &dcl.Property{
														Type:        "object",
														GoName:      "ArchiveExtraction",
														GoType:      "GuestPolicyRecipesUpdateStepsArchiveExtraction",
														Description: "Extracts an archive into the specified directory.",
														Properties: map[string]*dcl.Property{
															"artifactId": &dcl.Property{
																Type:        "string",
																GoName:      "ArtifactId",
																Description: "Required. The id of the relevant artifact in the recipe.",
															},
															"destination": &dcl.Property{
																Type:        "string",
																GoName:      "Destination",
																Description: "Directory to extract archive to. Defaults to `/` on Linux or `C:` on Windows.",
															},
															"type": &dcl.Property{
																Type:        "string",
																GoName:      "Type",
																GoType:      "GuestPolicyRecipesUpdateStepsArchiveExtractionTypeEnum",
																Description: "Required. The type of the archive to extract. Possible values: TYPE_UNSPECIFIED, VALIDATION, DESIRED_STATE_CHECK, DESIRED_STATE_ENFORCEMENT, DESIRED_STATE_CHECK_POST_ENFORCEMENT",
																Enum: []string{
																	"TYPE_UNSPECIFIED",
																	"VALIDATION",
																	"DESIRED_STATE_CHECK",
																	"DESIRED_STATE_ENFORCEMENT",
																	"DESIRED_STATE_CHECK_POST_ENFORCEMENT",
																},
															},
														},
													},
													"dpkgInstallation": &dcl.Property{
														Type:        "object",
														GoName:      "DpkgInstallation",
														GoType:      "GuestPolicyRecipesUpdateStepsDpkgInstallation",
														Description: "Installs a deb file via dpkg.",
														Properties: map[string]*dcl.Property{
															"artifactId": &dcl.Property{
																Type:        "string",
																GoName:      "ArtifactId",
																Description: "Required. The id of the relevant artifact in the recipe.",
															},
														},
													},
													"fileCopy": &dcl.Property{
														Type:        "object",
														GoName:      "FileCopy",
														GoType:      "GuestPolicyRecipesUpdateStepsFileCopy",
														Description: "Copies a file onto the instance.",
														Properties: map[string]*dcl.Property{
															"artifactId": &dcl.Property{
																Type:        "string",
																GoName:      "ArtifactId",
																Description: "Required. The id of the relevant artifact in the recipe.",
															},
															"destination": &dcl.Property{
																Type:        "string",
																GoName:      "Destination",
																Description: "Required. The absolute path on the instance to put the file.",
															},
															"overwrite": &dcl.Property{
																Type:        "boolean",
																GoName:      "Overwrite",
																Description: "Whether to allow this step to overwrite existing files. If this is false and the file already exists the file is not overwritten and the step is considered a success. Defaults to false.",
															},
															"permissions": &dcl.Property{
																Type:        "string",
																GoName:      "Permissions",
																Description: "Consists of three octal digits which represent, in order, the permissions of the owner, group, and other users for the file (similarly to the numeric mode used in the linux chmod utility). Each digit represents a three bit number with the 4 bit corresponding to the read permissions, the 2 bit corresponds to the write bit, and the one bit corresponds to the execute permission. Default behavior is 755. Below are some examples of permissions and their associated values: read, write, and execute: 7 read and execute: 5 read and write: 6 read only: 4",
															},
														},
													},
													"fileExec": &dcl.Property{
														Type:        "object",
														GoName:      "FileExec",
														GoType:      "GuestPolicyRecipesUpdateStepsFileExec",
														Description: "Executes an artifact or local file.",
														Properties: map[string]*dcl.Property{
															"allowedExitCodes": &dcl.Property{
																Type:        "array",
																GoName:      "AllowedExitCodes",
																Description: "Defaults to [0]. A list of possible return values that the program can return to indicate a success.",
																SendEmpty:   true,
																ListType:    "list",
																Items: &dcl.Property{
																	Type:   "integer",
																	Format: "int64",
																	GoType: "int64",
																},
															},
															"args": &dcl.Property{
																Type:        "array",
																GoName:      "Args",
																Description: "Arguments to be passed to the provided executable.",
																SendEmpty:   true,
																ListType:    "list",
																Items: &dcl.Property{
																	Type:   "string",
																	GoType: "string",
																},
															},
															"artifactId": &dcl.Property{
																Type:        "string",
																GoName:      "ArtifactId",
																Description: "The id of the relevant artifact in the recipe.",
															},
															"localPath": &dcl.Property{
																Type:        "string",
																GoName:      "LocalPath",
																Description: "The absolute path of the file on the local filesystem.",
															},
														},
													},
													"msiInstallation": &dcl.Property{
														Type:        "object",
														GoName:      "MsiInstallation",
														GoType:      "GuestPolicyRecipesUpdateStepsMsiInstallation",
														Description: "Installs an MSI file.",
														Properties: map[string]*dcl.Property{
															"allowedExitCodes": &dcl.Property{
																Type:        "array",
																GoName:      "AllowedExitCodes",
																Description: "Return codes that indicate that the software installed or updated successfully. Behaviour defaults to [0]",
																SendEmpty:   true,
																ListType:    "list",
																Items: &dcl.Property{
																	Type:   "integer",
																	Format: "int64",
																	GoType: "int64",
																},
															},
															"artifactId": &dcl.Property{
																Type:        "string",
																GoName:      "ArtifactId",
																Description: "Required. The id of the relevant artifact in the recipe.",
															},
															"flags": &dcl.Property{
																Type:        "array",
																GoName:      "Flags",
																Description: "The flags to use when installing the MSI defaults to [\"/i\"] (i.e. the install flag).",
																SendEmpty:   true,
																ListType:    "list",
																Items: &dcl.Property{
																	Type:   "string",
																	GoType: "string",
																},
															},
														},
													},
													"rpmInstallation": &dcl.Property{
														Type:        "object",
														GoName:      "RpmInstallation",
														GoType:      "GuestPolicyRecipesUpdateStepsRpmInstallation",
														Description: "Installs an rpm file via the rpm utility.",
														Properties: map[string]*dcl.Property{
															"artifactId": &dcl.Property{
																Type:        "string",
																GoName:      "ArtifactId",
																Description: "Required. The id of the relevant artifact in the recipe.",
															},
														},
													},
													"scriptRun": &dcl.Property{
														Type:        "object",
														GoName:      "ScriptRun",
														GoType:      "GuestPolicyRecipesUpdateStepsScriptRun",
														Description: "Runs commands in a shell.",
														Properties: map[string]*dcl.Property{
															"allowedExitCodes": &dcl.Property{
																Type:        "array",
																GoName:      "AllowedExitCodes",
																Description: "Return codes that indicate that the software installed or updated successfully. Behaviour defaults to [0]",
																SendEmpty:   true,
																ListType:    "list",
																Items: &dcl.Property{
																	Type:   "integer",
																	Format: "int64",
																	GoType: "int64",
																},
															},
															"interpreter": &dcl.Property{
																Type:        "string",
																GoName:      "Interpreter",
																GoType:      "GuestPolicyRecipesUpdateStepsScriptRunInterpreterEnum",
																Description: "The script interpreter to use to run the script. If no interpreter is specified the script is executed directly, which likely only succeed for scripts with [shebang lines](https://en.wikipedia.org/wiki/Shebang_(Unix)). Possible values: INTERPRETER_UNSPECIFIED, NONE, SHELL, POWERSHELL",
																Enum: []string{
																	"INTERPRETER_UNSPECIFIED",
																	"NONE",
																	"SHELL",
																	"POWERSHELL",
																},
															},
															"script": &dcl.Property{
																Type:        "string",
																GoName:      "Script",
																Description: "Required. The shell script to be executed.",
															},
														},
													},
												},
											},
										},
										"version": &dcl.Property{
											Type:        "string",
											GoName:      "Version",
											Description: "The version of this software recipe. Version can be up to 4 period separated numbers (e.g. 12.34.56.78).",
										},
									},
								},
							},
							"updateTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "UpdateTime",
								ReadOnly:    true,
								Description: "Output only. Last time this GuestPolicy was updated.",
								Immutable:   true,
							},
						},
					},
				},
			},
		},
	}
}
