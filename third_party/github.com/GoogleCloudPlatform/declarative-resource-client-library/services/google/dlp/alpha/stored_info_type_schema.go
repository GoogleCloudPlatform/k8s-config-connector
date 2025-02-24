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

func DCLStoredInfoTypeSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "Dlp/StoredInfoType",
			Description: "The Dlp StoredInfoType resource",
			StructName:  "StoredInfoType",
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a StoredInfoType",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "storedInfoType",
						Required:    true,
						Description: "A full instance of a StoredInfoType",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a StoredInfoType",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "storedInfoType",
						Required:    true,
						Description: "A full instance of a StoredInfoType",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a StoredInfoType",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "storedInfoType",
						Required:    true,
						Description: "A full instance of a StoredInfoType",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all StoredInfoType",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:     "location",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
					dcl.PathParameters{
						Name:     "parent",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
				},
			},
			List: &dcl.Path{
				Description: "The function used to list information about many StoredInfoType",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:     "location",
						Required: true,
						Schema: &dcl.PathParametersSchema{
							Type: "string",
						},
					},
					dcl.PathParameters{
						Name:     "parent",
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
				"StoredInfoType": &dcl.Component{
					Title: "StoredInfoType",
					ID:    "{{parent}}/storedInfoTypes/{{name}}",
					Locations: []string{
						"region",
					},
					HasCreate: true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"parent",
						},
						Properties: map[string]*dcl.Property{
							"description": &dcl.Property{
								Type:        "string",
								GoName:      "Description",
								Description: "Description of the StoredInfoType (max 256 characters).",
							},
							"dictionary": &dcl.Property{
								Type:        "object",
								GoName:      "Dictionary",
								GoType:      "StoredInfoTypeDictionary",
								Description: "Store dictionary-based CustomInfoType.",
								Conflicts: []string{
									"largeCustomDictionary",
									"regex",
								},
								Properties: map[string]*dcl.Property{
									"cloudStoragePath": &dcl.Property{
										Type:        "object",
										GoName:      "CloudStoragePath",
										GoType:      "StoredInfoTypeDictionaryCloudStoragePath",
										Description: "Newline-delimited file of words in Cloud Storage. Only a single file is accepted.",
										Conflicts: []string{
											"wordList",
										},
										Required: []string{
											"path",
										},
										Properties: map[string]*dcl.Property{
											"path": &dcl.Property{
												Type:        "string",
												GoName:      "Path",
												Description: "A url representing a file or path (no wildcards) in Cloud Storage. Example: gs://[BUCKET_NAME]/dictionary.txt",
											},
										},
									},
									"wordList": &dcl.Property{
										Type:        "object",
										GoName:      "WordList",
										GoType:      "StoredInfoTypeDictionaryWordList",
										Description: "List of words or phrases to search for.",
										Conflicts: []string{
											"cloudStoragePath",
										},
										Required: []string{
											"words",
										},
										Properties: map[string]*dcl.Property{
											"words": &dcl.Property{
												Type:        "array",
												GoName:      "Words",
												Description: "Words or phrases defining the dictionary. The dictionary must contain at least one phrase and every phrase must contain at least 2 characters that are letters or digits. [required]",
												SendEmpty:   true,
												ListType:    "list",
												Items: &dcl.Property{
													Type:   "string",
													GoType: "string",
												},
											},
										},
									},
								},
							},
							"displayName": &dcl.Property{
								Type:        "string",
								GoName:      "DisplayName",
								Description: "Display name of the StoredInfoType (max 256 characters).",
							},
							"largeCustomDictionary": &dcl.Property{
								Type:        "object",
								GoName:      "LargeCustomDictionary",
								GoType:      "StoredInfoTypeLargeCustomDictionary",
								Description: "StoredInfoType where findings are defined by a dictionary of phrases.",
								Conflicts: []string{
									"dictionary",
									"regex",
								},
								Properties: map[string]*dcl.Property{
									"bigQueryField": &dcl.Property{
										Type:        "object",
										GoName:      "BigQueryField",
										GoType:      "StoredInfoTypeLargeCustomDictionaryBigQueryField",
										Description: "Field in a BigQuery table where each cell represents a dictionary phrase.",
										Conflicts: []string{
											"cloudStorageFileSet",
										},
										Properties: map[string]*dcl.Property{
											"field": &dcl.Property{
												Type:        "object",
												GoName:      "Field",
												GoType:      "StoredInfoTypeLargeCustomDictionaryBigQueryFieldField",
												Description: "Designated field in the BigQuery table.",
												Properties: map[string]*dcl.Property{
													"name": &dcl.Property{
														Type:        "string",
														GoName:      "Name",
														Description: "Name describing the field.",
													},
												},
											},
											"table": &dcl.Property{
												Type:        "object",
												GoName:      "Table",
												GoType:      "StoredInfoTypeLargeCustomDictionaryBigQueryFieldTable",
												Description: "Source table of the field.",
												Properties: map[string]*dcl.Property{
													"datasetId": &dcl.Property{
														Type:        "string",
														GoName:      "DatasetId",
														Description: "Dataset ID of the table.",
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
														Description: "The Google Cloud Platform project ID of the project containing the table. If omitted, project ID is inferred from the API call.",
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
														Description: "Name of the table.",
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
									"cloudStorageFileSet": &dcl.Property{
										Type:        "object",
										GoName:      "CloudStorageFileSet",
										GoType:      "StoredInfoTypeLargeCustomDictionaryCloudStorageFileSet",
										Description: "Set of files containing newline-delimited lists of dictionary phrases.",
										Conflicts: []string{
											"bigQueryField",
										},
										Required: []string{
											"url",
										},
										Properties: map[string]*dcl.Property{
											"url": &dcl.Property{
												Type:        "string",
												GoName:      "Url",
												Description: "The url, in the format `gs:///`. Trailing wildcard in the path is allowed.",
											},
										},
									},
									"outputPath": &dcl.Property{
										Type:        "object",
										GoName:      "OutputPath",
										GoType:      "StoredInfoTypeLargeCustomDictionaryOutputPath",
										Description: "Location to store dictionary artifacts in Google Cloud Storage. These files will only be accessible by project owners and the DLP API. If any of these artifacts are modified, the dictionary is considered invalid and can no longer be used.",
										Required: []string{
											"path",
										},
										Properties: map[string]*dcl.Property{
											"path": &dcl.Property{
												Type:        "string",
												GoName:      "Path",
												Description: "A url representing a file or path (no wildcards) in Cloud Storage. Example: gs://[BUCKET_NAME]/dictionary.txt",
											},
										},
									},
								},
							},
							"location": &dcl.Property{
								Type:        "string",
								GoName:      "Location",
								Description: "The location of the resource",
								Immutable:   true,
							},
							"name": &dcl.Property{
								Type:                     "string",
								GoName:                   "Name",
								Description:              "Resource name.",
								Immutable:                true,
								ServerGeneratedParameter: true,
							},
							"parent": &dcl.Property{
								Type:                "string",
								GoName:              "Parent",
								Description:         "The parent of the resource.",
								Immutable:           true,
								ForwardSlashAllowed: true,
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Cloudresourcemanager/Organization",
										Field:    "name",
										Parent:   true,
									},
									&dcl.PropertyResourceReference{
										Resource: "Cloudresourcemanager/Project",
										Field:    "name",
										Parent:   true,
									},
								},
							},
							"regex": &dcl.Property{
								Type:        "object",
								GoName:      "Regex",
								GoType:      "StoredInfoTypeRegex",
								Description: "Store regular expression-based StoredInfoType.",
								Conflicts: []string{
									"largeCustomDictionary",
									"dictionary",
								},
								Required: []string{
									"pattern",
								},
								Properties: map[string]*dcl.Property{
									"groupIndexes": &dcl.Property{
										Type:        "array",
										GoName:      "GroupIndexes",
										Description: "The index of the submatch to extract as findings. When not specified, the entire match is returned. No more than 3 may be included.",
										SendEmpty:   true,
										ListType:    "list",
										Items: &dcl.Property{
											Type:   "integer",
											Format: "int64",
											GoType: "int64",
										},
									},
									"pattern": &dcl.Property{
										Type:        "string",
										GoName:      "Pattern",
										Description: "Pattern defining the regular expression. Its syntax (https://github.com/google/re2/wiki/Syntax) can be found under the google/re2 repository on GitHub.",
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
