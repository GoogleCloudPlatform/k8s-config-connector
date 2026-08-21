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

func DCLBackupSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "Filestore/Backup",
			Description: "The Filestore Backup resource",
			StructName:  "Backup",
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a Backup",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "backup",
						Required:    true,
						Description: "A full instance of a Backup",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a Backup",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "backup",
						Required:    true,
						Description: "A full instance of a Backup",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a Backup",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "backup",
						Required:    true,
						Description: "A full instance of a Backup",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all Backup",
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
				Description: "The function used to list information about many Backup",
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
				"Backup": &dcl.Component{
					Title:           "Backup",
					ID:              "projects/{{project}}/locations/{{location}}/backups/{{name}}",
					ParentContainer: "project",
					LabelsField:     "labels",
					HasCreate:       true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"name",
							"sourceInstance",
							"sourceFileShare",
							"project",
							"location",
						},
						Properties: map[string]*dcl.Property{
							"capacityGb": &dcl.Property{
								Type:        "integer",
								Format:      "int64",
								GoName:      "CapacityGb",
								ReadOnly:    true,
								Description: "Output only. Capacity of the source file share when the backup was created.",
								Immutable:   true,
							},
							"createTime": &dcl.Property{
								Type:        "string",
								Format:      "date-time",
								GoName:      "CreateTime",
								ReadOnly:    true,
								Description: "Output only. The time when the backup was created.",
								Immutable:   true,
							},
							"description": &dcl.Property{
								Type:        "string",
								GoName:      "Description",
								Description: "A description of the backup with 2048 characters or less. Requests with longer descriptions will be rejected.",
							},
							"downloadBytes": &dcl.Property{
								Type:        "integer",
								Format:      "int64",
								GoName:      "DownloadBytes",
								ReadOnly:    true,
								Description: "Output only. Amount of bytes that will be downloaded if the backup is restored. This may be different than storage bytes, since sequential backups of the same disk will share storage.",
								Immutable:   true,
							},
							"labels": &dcl.Property{
								Type: "object",
								AdditionalProperties: &dcl.Property{
									Type: "string",
								},
								GoName:      "Labels",
								Description: "Resource labels to represent user provided metadata.",
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
								Description: "The resource name of the backup.",
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
							"sourceFileShare": &dcl.Property{
								Type:        "string",
								GoName:      "SourceFileShare",
								Description: "Name of the file share in the source Cloud Filestore instance that the backup is created from.",
								Immutable:   true,
							},
							"sourceInstance": &dcl.Property{
								Type:        "string",
								GoName:      "SourceInstance",
								Description: "The resource name of the source Cloud Filestore instance, in the format projects/{project_number}/locations/{location_id}/instances/{instance_id}, used to create this backup.",
								Immutable:   true,
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Filestore/Instance",
										Field:    "name",
									},
								},
							},
							"sourceInstanceTier": &dcl.Property{
								Type:        "string",
								GoName:      "SourceInstanceTier",
								GoType:      "BackupSourceInstanceTierEnum",
								ReadOnly:    true,
								Description: "Output only. The service tier of the source Cloud Filestore instance that this backup is created from. Possible values: TIER_UNSPECIFIED, STANDARD, PREMIUM, BASIC_HDD, BASIC_SSD, HIGH_SCALE_SSD",
								Immutable:   true,
								Enum: []string{
									"TIER_UNSPECIFIED",
									"STANDARD",
									"PREMIUM",
									"BASIC_HDD",
									"BASIC_SSD",
									"HIGH_SCALE_SSD",
								},
							},
							"state": &dcl.Property{
								Type:        "string",
								GoName:      "State",
								GoType:      "BackupStateEnum",
								ReadOnly:    true,
								Description: "Output only. The backup state. Possible values: STATE_UNSPECIFIED, CREATING, READY, REPAIRING, DELETING, ERROR, RESTORING",
								Immutable:   true,
								Enum: []string{
									"STATE_UNSPECIFIED",
									"CREATING",
									"READY",
									"REPAIRING",
									"DELETING",
									"ERROR",
									"RESTORING",
								},
							},
							"storageBytes": &dcl.Property{
								Type:        "integer",
								Format:      "int64",
								GoName:      "StorageBytes",
								ReadOnly:    true,
								Description: "Output only. The size of the storage used by the backup. As backups share storage, this number is expected to change with backup creation/deletion.",
								Immutable:   true,
							},
						},
					},
				},
			},
		},
	}
}
