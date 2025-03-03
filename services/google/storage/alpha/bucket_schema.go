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

func DCLBucketSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "Storage/Bucket",
			Description: "The Storage Bucket resource",
			StructName:  "Bucket",
			HasIAM:      true,
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a Bucket",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "bucket",
						Required:    true,
						Description: "A full instance of a Bucket",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a Bucket",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "bucket",
						Required:    true,
						Description: "A full instance of a Bucket",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a Bucket",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "bucket",
						Required:    true,
						Description: "A full instance of a Bucket",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all Bucket",
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
				Description: "The function used to list information about many Bucket",
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
				"Bucket": &dcl.Component{
					Title:           "Bucket",
					ID:              "b/{{name}}?userProject={{project}}",
					ParentContainer: "project",
					HasCreate:       true,
					HasIAM:          true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"project",
							"location",
							"name",
						},
						Properties: map[string]*dcl.Property{
							"cors": &dcl.Property{
								Type:        "array",
								GoName:      "Cors",
								Description: "The bucket's Cross-Origin Resource Sharing (CORS) configuration. ",
								SendEmpty:   true,
								ListType:    "list",
								Items: &dcl.Property{
									Type:   "object",
									GoType: "BucketCors",
									Properties: map[string]*dcl.Property{
										"maxAgeSeconds": &dcl.Property{
											Type:        "integer",
											Format:      "int64",
											GoName:      "MaxAgeSeconds",
											Description: "The value, in seconds, to return in the Access-Control-Max-Age header used in preflight responses. ",
										},
										"method": &dcl.Property{
											Type:        "array",
											GoName:      "Method",
											Description: "The list of HTTP methods on which to include CORS response headers, (GET, OPTIONS, POST, etc) Note: \"*\" is permitted in the list of methods, and means \"any method\". ",
											SendEmpty:   true,
											ListType:    "list",
											Items: &dcl.Property{
												Type:   "string",
												GoType: "string",
											},
										},
										"origin": &dcl.Property{
											Type:        "array",
											GoName:      "Origin",
											Description: "The list of Origins eligible to receive CORS response headers. Note: \"*\" is permitted in the list of origins, and means \"any Origin\". ",
											SendEmpty:   true,
											ListType:    "list",
											Items: &dcl.Property{
												Type:   "string",
												GoType: "string",
											},
										},
										"responseHeader": &dcl.Property{
											Type:        "array",
											GoName:      "ResponseHeader",
											Description: "The list of HTTP headers other than the simple response headers to give permission for the user-agent to share across domains. ",
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
							"lifecycle": &dcl.Property{
								Type:        "object",
								GoName:      "Lifecycle",
								GoType:      "BucketLifecycle",
								Description: "The bucket's lifecycle configuration.  See https://developers.google.com/storage/docs/lifecycle for more information. ",
								Properties: map[string]*dcl.Property{
									"rule": &dcl.Property{
										Type:        "array",
										GoName:      "Rule",
										Description: "A lifecycle management rule, which is made of an action to take and the condition(s) under which the action will be taken. ",
										SendEmpty:   true,
										ListType:    "list",
										Items: &dcl.Property{
											Type:   "object",
											GoType: "BucketLifecycleRule",
											Properties: map[string]*dcl.Property{
												"action": &dcl.Property{
													Type:        "object",
													GoName:      "Action",
													GoType:      "BucketLifecycleRuleAction",
													Description: "The action to take.",
													Properties: map[string]*dcl.Property{
														"storageClass": &dcl.Property{
															Type:        "string",
															GoName:      "StorageClass",
															Description: "Target storage class. Required if the type of the action is SetStorageClass. ",
														},
														"type": &dcl.Property{
															Type:        "string",
															GoName:      "Type",
															GoType:      "BucketLifecycleRuleActionTypeEnum",
															Description: "Type of the action. Currently, only Delete and SetStorageClass are supported. ",
															Enum: []string{
																"Delete",
																"SetStorageClass",
															},
														},
													},
												},
												"condition": &dcl.Property{
													Type:        "object",
													GoName:      "Condition",
													GoType:      "BucketLifecycleRuleCondition",
													Description: "The condition(s) under which the action will be taken. ",
													Properties: map[string]*dcl.Property{
														"age": &dcl.Property{
															Type:        "integer",
															Format:      "int64",
															GoName:      "Age",
															Description: "Age of an object (in days). This condition is satisfied when an object reaches the specified age. ",
														},
														"createdBefore": &dcl.Property{
															Type:        "string",
															Format:      "date-time",
															GoName:      "CreatedBefore",
															Description: "A date in RFC 3339 format with only the date part (for instance, \"2013-01-15\"). This condition is satisfied when an object is created before midnight of the specified date in UTC. ",
														},
														"matchesStorageClass": &dcl.Property{
															Type:        "array",
															GoName:      "MatchesStorageClass",
															Description: "Objects having any of the storage classes specified by this condition will be matched. Values include MULTI_REGIONAL, REGIONAL, NEARLINE, COLDLINE, STANDARD, and DURABLE_REDUCED_AVAILABILITY. ",
															SendEmpty:   true,
															ListType:    "list",
															Items: &dcl.Property{
																Type:   "string",
																GoType: "string",
															},
														},
														"numNewerVersions": &dcl.Property{
															Type:        "integer",
															Format:      "int64",
															GoName:      "NumNewerVersions",
															Description: "Relevant only for versioned objects. If the value is N, this condition is satisfied when there are at least N versions (including the live version) newer than this version of the object. ",
														},
														"withState": &dcl.Property{
															Type:        "string",
															GoName:      "WithState",
															GoType:      "BucketLifecycleRuleConditionWithStateEnum",
															Description: "Match to live and/or archived objects. Unversioned buckets have only live objects. Supported values include: 'LIVE', 'ARCHIVED', 'ANY'.",
															Enum: []string{
																"LIVE",
																"ARCHIVED",
																"ANY",
															},
														},
													},
												},
											},
										},
									},
								},
							},
							"location": &dcl.Property{
								Type:        "string",
								GoName:      "Location",
								Description: "The location of the bucket. Object data for objects in the bucket resides in physical storage within this region. Defaults to `US`. ",
								Immutable:   true,
							},
							"logging": &dcl.Property{
								Type:        "object",
								GoName:      "Logging",
								GoType:      "BucketLogging",
								Description: "The bucket's logging configuration, which defines the destination bucket and optional name prefix for the current bucket's logs. ",
								Properties: map[string]*dcl.Property{
									"logBucket": &dcl.Property{
										Type:        "string",
										GoName:      "LogBucket",
										Description: "The destination bucket where the current bucket's logs should be placed. ",
									},
									"logObjectPrefix": &dcl.Property{
										Type:        "string",
										GoName:      "LogObjectPrefix",
										Description: "The object prefix for log objects. If it's not provided, it defaults to the bucket's name.",
									},
								},
							},
							"name": &dcl.Property{
								Type:        "string",
								GoName:      "Name",
								Description: "The name of the bucket. ",
								Immutable:   true,
							},
							"project": &dcl.Property{
								Type:        "string",
								GoName:      "Project",
								Description: "The project id of the resource.",
								Immutable:   true,
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Cloudresourcemanager/Project",
										Field:    "name",
										Parent:   true,
									},
								},
							},
							"storageClass": &dcl.Property{
								Type:        "string",
								GoName:      "StorageClass",
								GoType:      "BucketStorageClassEnum",
								Description: "The bucket's default storage class, used whenever no storageClass is specified for a newly-created object. This defines how objects in the bucket are stored and determines the SLA and the cost of storage. Values include MULTI_REGIONAL, REGIONAL, STANDARD, NEARLINE, COLDLINE, ARCHIVE, and DURABLE_REDUCED_AVAILABILITY. If this value is not specified when the bucket is created, it will default to STANDARD. For more information, see storage classes. ",
								Enum: []string{
									"MULTI_REGIONAL",
									"REGIONAL",
									"STANDARD",
									"NEARLINE",
									"COLDLINE",
									"ARCHIVE",
									"DURABLE_REDUCED_AVAILABILITY",
								},
							},
							"versioning": &dcl.Property{
								Type:        "object",
								GoName:      "Versioning",
								GoType:      "BucketVersioning",
								Description: "The bucket's versioning configuration.",
								Properties: map[string]*dcl.Property{
									"enabled": &dcl.Property{
										Type:        "boolean",
										GoName:      "Enabled",
										Description: "While set to true, versioning is fully enabled for this bucket. ",
									},
								},
							},
							"website": &dcl.Property{
								Type:        "object",
								GoName:      "Website",
								GoType:      "BucketWebsite",
								Description: "The bucket's website configuration, controlling how the service behaves when accessing bucket contents as a web site. See the Static Website Examples for more information. ",
								Properties: map[string]*dcl.Property{
									"mainPageSuffix": &dcl.Property{
										Type:        "string",
										GoName:      "MainPageSuffix",
										Description: "If the requested object path is missing, the service will ensure the path has a trailing '/', append this suffix, and attempt to retrieve the resulting object. This allows the creation of index.html objects to represent directory pages. ",
									},
									"notFoundPage": &dcl.Property{
										Type:        "string",
										GoName:      "NotFoundPage",
										Description: "If the requested object path is missing, and any mainPageSuffix object is missing, if applicable, the service will return the named object from this bucket as the content for a 404 Not Found result. ",
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
