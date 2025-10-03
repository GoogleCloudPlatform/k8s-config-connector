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
package pubsub

import (
	"github.com/GoogleCloudPlatform/declarative-resource-client-library/dcl"
)

func DCLTopicSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:       "Pubsub/Topic",
			Description: "The Pubsub Topic resource",
			StructName:  "Topic",
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a Topic",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "topic",
						Required:    true,
						Description: "A full instance of a Topic",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a Topic",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "topic",
						Required:    true,
						Description: "A full instance of a Topic",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a Topic",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "topic",
						Required:    true,
						Description: "A full instance of a Topic",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all Topic",
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
				Description: "The function used to list information about many Topic",
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
				"Topic": &dcl.Component{
					Title:           "Topic",
					ID:              "projects/{{project}}/topics/{{name}}",
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
							"kmsKeyName": &dcl.Property{
								Type:        "string",
								GoName:      "KmsKeyName",
								Description: "The resource name of the Cloud KMS CryptoKey to be used to protect access to messages published on this topic. Your project's Pub/Sub service account (`service-{{PROJECT_NUMBER}}@gcp-sa-pubsub.iam.gserviceaccount.com`) must have `roles/cloudkms.cryptoKeyEncrypterDecrypter` to use this feature.  The expected format is `projects/*/locations/*/keyRings/*/cryptoKeys/*` ",
							},
							"labels": &dcl.Property{
								Type: "object",
								AdditionalProperties: &dcl.Property{
									Type: "string",
								},
								GoName:      "Labels",
								Description: "A set of key/value label pairs to assign to this Topic. ",
							},
							"messageStoragePolicy": &dcl.Property{
								Type:          "object",
								GoName:        "MessageStoragePolicy",
								GoType:        "TopicMessageStoragePolicy",
								Description:   "Policy constraining the set of Google Cloud Platform regions where messages published to the topic may be stored. If not present, then no constraints are in effect. ",
								ServerDefault: true,
								Properties: map[string]*dcl.Property{
									"allowedPersistenceRegions": &dcl.Property{
										Type:        "array",
										GoName:      "AllowedPersistenceRegions",
										Description: "A list of IDs of GCP regions where messages that are published to the topic may be persisted in storage. Messages published by publishers running in non-allowed GCP regions (or running outside of GCP altogether) will be routed for storage in one of the allowed regions. An empty list means that no regions are allowed, and is not a valid configuration. ",
										SendEmpty:   true,
										ListType:    "set",
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
								Description: "Name of the topic.",
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
						},
					},
				},
			},
		},
	}
}
