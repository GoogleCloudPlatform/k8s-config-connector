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
package bigquery

import (
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func DCLDatasetSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "Bigquery/Dataset",
			Description: "The Bigquery Dataset resource",
			StructName:  "Dataset",
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a Dataset",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "dataset",
						Required:    true,
						Description: "A full instance of a Dataset",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a Dataset",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "dataset",
						Required:    true,
						Description: "A full instance of a Dataset",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a Dataset",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "dataset",
						Required:    true,
						Description: "A full instance of a Dataset",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all Dataset",
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
				Description: "The function used to list information about many Dataset",
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
				"Dataset": &dcl.Component{
					Title:           "Dataset",
					ID:              "projects/{{project}}/datasets/{{name}}",
					ParentContainer: "project",
					LabelsField:     "labels",
					HasCreate:       true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"name",
							"project",
						},
						Properties: map[string]*dcl.Property{
							"access": &dcl.Property{
								Type:        "array",
								GoName:      "Access",
								Description: "Optional. An array of objects that define dataset access for one or more entities. You can set this property when inserting or updating a dataset in order to control who is allowed to access the data. If unspecified at dataset creation time, BigQuery adds default dataset access for the following entities: access.specialGroup: projectReaders; access.role: READER; access.specialGroup: projectWriters; access.role: WRITER; access.specialGroup: projectOwners; access.role: OWNER; access.userByEmail: ; access.role: OWNER;",
								SendEmpty:   true,
								ListType:    "set",
								Items: &dcl.Property{
									Type:   "object",
									GoType: "DatasetAccess",
									Required: []string{
										"role",
									},
									Properties: map[string]*dcl.Property{
										"domain": &dcl.Property{
											Type:        "string",
											GoName:      "Domain",
											Description: "A domain to grant access to. Any users signed in with the domain specified will be granted the specified access. Example: \"example.com\". Maps to IAM policy member \"domain:DOMAIN\".",
											Conflicts: []string{
												"userByEmail",
												"groupByEmail",
												"specialGroup",
												"iamMember",
												"view",
												"routine",
											},
										},
										"groupByEmail": &dcl.Property{
											Type:        "string",
											GoName:      "GroupByEmail",
											Description: "An email address of a Google Group to grant access to. Maps to IAM policy member \"group:GROUP\".",
											Conflicts: []string{
												"userByEmail",
												"domain",
												"specialGroup",
												"iamMember",
												"view",
												"routine",
											},
										},
										"iamMember": &dcl.Property{
											Type:        "string",
											GoName:      "IamMember",
											Description: "Some other type of member that appears in the IAM Policy but isn't a user, group, domain, or special group.",
											Conflicts: []string{
												"userByEmail",
												"groupByEmail",
												"domain",
												"specialGroup",
												"view",
												"routine",
											},
										},
										"role": &dcl.Property{
											Type:        "string",
											GoName:      "Role",
											Description: "Required. An IAM role ID that should be granted to the user, group, or domain specified in this access entry. The following legacy mappings will be applied: OWNER <=> roles/bigquery.dataOwner WRITER <=> roles/bigquery.dataEditor READER <=> roles/bigquery.dataViewer This field will accept any of the above formats, but will return only the legacy format. For example, if you set this field to \"roles/bigquery.dataOwner\", it will be returned back as \"OWNER\".",
										},
										"routine": &dcl.Property{
											Type:        "object",
											GoName:      "Routine",
											GoType:      "DatasetAccessRoutine",
											Description: "A routine from a different dataset to grant access to. Queries executed against that routine will have read access to views/tables/routines in this dataset. Only UDF is supported for now. The role field is not required when this field is set. If that routine is updated by any user, access to the routine needs to be granted again via an update operation.",
											Conflicts: []string{
												"userByEmail",
												"groupByEmail",
												"domain",
												"specialGroup",
												"iamMember",
												"view",
											},
											Required: []string{
												"projectId",
												"datasetId",
												"routineId",
											},
											Properties: map[string]*dcl.Property{
												"datasetId": &dcl.Property{
													Type:        "string",
													GoName:      "DatasetId",
													Description: "Required. The ID of the dataset containing this routine.",
													ResourceReferences: []*dcl.PropertyResourceReference{
														&dcl.PropertyResourceReference{
															Resource: "Bigquery/Dataset",
															Field:    "name",
														},
													},
												},
												"projectId": &dcl.Property{
													Type:        "string",
													GoName:      "ProjectId",
													Description: "Required. The ID of the project containing this routine.",
													ResourceReferences: []*dcl.PropertyResourceReference{
														&dcl.PropertyResourceReference{
															Resource: "Cloudresourcemanager/Project",
															Field:    "name",
														},
													},
												},
												"routineId": &dcl.Property{
													Type:        "string",
													GoName:      "RoutineId",
													Description: "Required. The ID of the routine. The ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_). The maximum length is 256 characters.",
													ResourceReferences: []*dcl.PropertyResourceReference{
														&dcl.PropertyResourceReference{
															Resource: "Bigquery/Routine",
															Field:    "name",
														},
													},
												},
											},
										},
										"specialGroup": &dcl.Property{
											Type:        "string",
											GoName:      "SpecialGroup",
											Description: "A special group to grant access to. Possible values include: projectOwners: Owners of the enclosing project. projectReaders: Readers of the enclosing project. projectWriters: Writers of the enclosing project. allAuthenticatedUsers: All authenticated BigQuery users. Maps to similarly-named IAM members.",
											Conflicts: []string{
												"userByEmail",
												"groupByEmail",
												"domain",
												"iamMember",
												"view",
												"routine",
											},
										},
										"userByEmail": &dcl.Property{
											Type:        "string",
											GoName:      "UserByEmail",
											Description: "An email address of a user to grant access to. For example: fred@example.com. Maps to IAM policy member \"user:EMAIL\" or \"serviceAccount:EMAIL\".",
											Conflicts: []string{
												"groupByEmail",
												"domain",
												"specialGroup",
												"iamMember",
												"view",
												"routine",
											},
										},
										"view": &dcl.Property{
											Type:        "object",
											GoName:      "View",
											GoType:      "DatasetAccessView",
											Description: "A view from a different dataset to grant access to. Queries executed against that view will have read access to views/tables/routines in this dataset. The role field is not required when this field is set. If that view is updated by any user, access to the view needs to be granted again via an update operation.",
											Conflicts: []string{
												"userByEmail",
												"groupByEmail",
												"domain",
												"specialGroup",
												"iamMember",
												"routine",
											},
											Required: []string{
												"projectId",
												"datasetId",
												"tableId",
											},
											Properties: map[string]*dcl.Property{
												"datasetId": &dcl.Property{
													Type:        "string",
													GoName:      "DatasetId",
													Description: "Required. The ID of the dataset containing this table.",
													ResourceReferences: []*dcl.PropertyResourceReference{
														&dcl.PropertyResourceReference{
															Resource: "Bigquery/Dataset",
															Field:    "name",
														},
													},
												},
												"projectId": &dcl.Property{
													Type:        "string",
													GoName:      "ProjectId",
													Description: "Required. The ID of the project containing this table.",
													ResourceReferences: []*dcl.PropertyResourceReference{
														&dcl.PropertyResourceReference{
															Resource: "Cloudresourcemanager/Project",
															Field:    "name",
														},
													},
												},
												"tableId": &dcl.Property{
													Type:        "string",
													GoName:      "TableId",
													Description: "Required. The ID of the table. The ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_). The maximum length is 1,024 characters. Certain operations allow suffixing of the table ID with a partition decorator, such as `sample_table$20190123`.",
													ResourceReferences: []*dcl.PropertyResourceReference{
														&dcl.PropertyResourceReference{
															Resource: "Bigquery/Table",
															Field:    "name",
														},
													},
												},
											},
										},
									},
								},
							},
							"creationTime": &dcl.Property{
								Type:        "integer",
								Format:      "int64",
								GoName:      "CreationTime",
								ReadOnly:    true,
								Description: "Output only. The time when this dataset was created, in milliseconds since the epoch.",
								Immutable:   true,
							},
							"defaultEncryptionConfiguration": &dcl.Property{
								Type:        "object",
								GoName:      "DefaultEncryptionConfiguration",
								GoType:      "DatasetDefaultEncryptionConfiguration",
								Description: "The default encryption key for all tables in the dataset. Once this property is set, all newly-created partitioned tables in the dataset will have encryption key set to this value, unless table creation request (or query) overrides the key.",
								Immutable:   true,
								Properties: map[string]*dcl.Property{
									"kmsKeyName": &dcl.Property{
										Type:        "string",
										GoName:      "KmsKeyName",
										Description: "Optional. Describes the Cloud KMS encryption key that will be used to protect destination BigQuery table. The BigQuery Service Account associated with your project requires access to this encryption key.",
										Immutable:   true,
										ResourceReferences: []*dcl.PropertyResourceReference{
											&dcl.PropertyResourceReference{
												Resource: "Cloudkms/CryptoKey",
												Field:    "name",
											},
										},
									},
								},
							},
							"defaultPartitionExpirationMs": &dcl.Property{
								Type:        "string",
								GoName:      "DefaultPartitionExpirationMs",
								Description: "This default partition expiration, expressed in milliseconds. When new time-partitioned tables are created in a dataset where this property is set, the table will inherit this value, propagated as the `TimePartitioning.expirationMs` property on the new table. If you set `TimePartitioning.expirationMs` explicitly when creating a table, the `defaultPartitionExpirationMs` of the containing dataset is ignored. When creating a partitioned table, if `defaultPartitionExpirationMs` is set, the `defaultTableExpirationMs` value is ignored and the table will not inherit a table expiration deadline.",
							},
							"defaultTableExpirationMs": &dcl.Property{
								Type:        "string",
								GoName:      "DefaultTableExpirationMs",
								Description: "Optional. The default lifetime of all tables in the dataset, in milliseconds. The minimum lifetime value is 3600000 milliseconds (one hour). To clear an existing default expiration with a PATCH request, set to 0. Once this property is set, all newly-created tables in the dataset will have an expirationTime property set to the creation time plus the value in this property, and changing the value will only affect new tables, not existing ones. When the expirationTime for a given table is reached, that table will be deleted automatically. If a table's expirationTime is modified or removed before the table expires, or if you provide an explicit expirationTime when creating a table, that value takes precedence over the default expiration time indicated by this property.",
							},
							"description": &dcl.Property{
								Type:        "string",
								GoName:      "Description",
								Description: "Optional. A user-friendly description of the dataset.",
							},
							"etag": &dcl.Property{
								Type:        "string",
								GoName:      "Etag",
								ReadOnly:    true,
								Description: "Output only. A hash of the resource.",
								Immutable:   true,
							},
							"friendlyName": &dcl.Property{
								Type:        "string",
								GoName:      "FriendlyName",
								Description: "Optional. A descriptive name for the dataset.",
							},
							"id": &dcl.Property{
								Type:        "string",
								GoName:      "Id",
								ReadOnly:    true,
								Description: "Output only. The fully-qualified unique name of the dataset in the format projectId:datasetId. The dataset name without the project name is given in the datasetId field. When creating a new dataset, leave this field blank, and instead specify the datasetId field.",
								Immutable:   true,
							},
							"labels": &dcl.Property{
								Type: "object",
								AdditionalProperties: &dcl.Property{
									Type: "string",
								},
								GoName:      "Labels",
								Description: "The labels associated with this dataset. You can use these to organize and group your datasets. You can set this property when inserting or updating a dataset. See (/bigquery/docs/creating-managing-labels#creating_and_updating_dataset_labels) for more information.",
							},
							"lastModifiedTime": &dcl.Property{
								Type:        "integer",
								Format:      "int64",
								GoName:      "LastModifiedTime",
								ReadOnly:    true,
								Description: "Output only. The date when this dataset was last modified, in milliseconds since the epoch.",
								Immutable:   true,
							},
							"location": &dcl.Property{
								Type:        "string",
								GoName:      "Location",
								Description: "The geographic location where the dataset should reside. See https://cloud.google.com/bigquery/docs/locations for supported locations.",
								Immutable:   true,
							},
							"name": &dcl.Property{
								Type:        "string",
								GoName:      "Name",
								Description: "Required. A unique ID for this dataset, without the project name. The ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_). The maximum length is 1,024 characters.",
								Immutable:   true,
							},
							"project": &dcl.Property{
								Type:        "string",
								GoName:      "Project",
								Description: "The ID of the project containing this dataset.",
								Immutable:   true,
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Cloudresourcemanager/Project",
										Field:    "name",
										Parent:   true,
									},
								},
							},
							"published": &dcl.Property{
								Type:        "boolean",
								GoName:      "Published",
								Description: "Whether this dataset is visible to all users in public search. This field can only be set to true if READER access is given to allAuthenticatedUsers in the access list. The default value is false.",
								Immutable:   true,
							},
							"selfLink": &dcl.Property{
								Type:        "string",
								GoName:      "SelfLink",
								ReadOnly:    true,
								Description: "Output only. A URL that can be used to access the resource again. You can use this URL in Get or Update requests to the resource.",
								Immutable:   true,
							},
						},
					},
				},
			},
		},
	}
}
