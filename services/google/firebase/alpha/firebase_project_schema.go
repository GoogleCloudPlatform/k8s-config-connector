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

func DCLFirebaseProjectSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:      "Firebase/FirebaseProject",
			StructName: "FirebaseProject",
			Reference: &dcl.Link{
				Text: "Firebase Project API Documentation",
				URL:  "https://firebase.google.com/docs/projects/api/reference/rest#rest-resource:-v1beta1.projects",
			},
			Guides: []*dcl.Link{
				&dcl.Link{
					Text: "Get started with Firebase Projects and Apps",
					URL:  "https://firebase.google.com/docs/projects/api/workflow_set-up-and-manage-project",
				},
			},
		},
		Paths: &dcl.Paths{
			Get: &dcl.Path{
				Description: "The function used to get information about a FirebaseProject",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "firebaseProject",
						Required:    true,
						Description: "A full instance of a FirebaseProject",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a FirebaseProject",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "firebaseProject",
						Required:    true,
						Description: "A full instance of a FirebaseProject",
					},
				},
			},
			List: &dcl.Path{
				Description: "The function used to list information about many FirebaseProject",
			},
		},
		Components: &dcl.Components{
			Schemas: map[string]*dcl.Component{
				"FirebaseProject": &dcl.Component{
					Title:           "FirebaseProject",
					ID:              "projects/{{project}}",
					ParentContainer: "project",
					LabelsField:     "annotations",
					HasCreate:       true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"project",
						},
						Properties: map[string]*dcl.Property{
							"annotations": &dcl.Property{
								Type: "object",
								AdditionalProperties: &dcl.Property{
									Type: "string",
								},
								GoName:      "Annotations",
								Description: "  // Set of user-defined annotations for the [FirebaseProject][] as per [AIP-128](https://google.aip.dev/128#annotations).  These annotations are intended solely for developers and client-side tools Firebase services will not mutate this annotation set.  This field may only be assigned on Update",
							},
							"displayName": &dcl.Property{
								Type:        "string",
								GoName:      "DisplayName",
								Description: "The user-assigned display name of the Project.  This field may only be assigned on Update",
							},
							"project": &dcl.Property{
								Type:        "string",
								GoName:      "Project",
								Description: "The project for the resource.  FirebaseProjects are generally referneced by the GCP Project they augment.",
								Immutable:   true,
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Cloudresourcemanager/Project",
										Field:    "name",
										Parent:   true,
									},
								},
							},
							"projectId": &dcl.Property{
								Type:        "string",
								GoName:      "ProjectId",
								ReadOnly:    true,
								Description: "Immutable. A user-assigned unique identifier for the Project. This identifier may appear in URLs or names for some Firebase resources associated with the Project, but it should generally be treated as a convenience alias to reference the Project.",
								Immutable:   true,
							},
							"projectNumber": &dcl.Property{
								Type:        "integer",
								Format:      "int64",
								GoName:      "ProjectNumber",
								ReadOnly:    true,
								Description: "Immutable. The globally unique, Google-assigned canonical identifier for the Project. Use this identifier when configuring integrations and/or making API calls to Firebase or third-party services.",
								Immutable:   true,
							},
							"resources": &dcl.Property{
								Type:        "object",
								GoName:      "Resources",
								GoType:      "FirebaseProjectResources",
								ReadOnly:    true,
								Description: "The default Firebase resources associated with the Project.",
								Immutable:   true,
								Properties: map[string]*dcl.Property{
									"hostingSite": &dcl.Property{
										Type:        "string",
										GoName:      "HostingSite",
										ReadOnly:    true,
										Description: "The default Firebase Hosting site name, in the format: `PROJECT_ID` Though rare, your `projectId` might already be used as the name for an existing Hosting site in another project (learn more about creating non-default, [additional sites](https://firebase.google.com/docs/hosting/multisites)). In these cases, your `projectId` is appended with a hyphen then five alphanumeric characters to create your default Hosting site name. For example, if your `projectId` is `myproject123`, your default Hosting site name might be: `myproject123-a5c16`",
										Immutable:   true,
									},
									"locationId": &dcl.Property{
										Type:        "string",
										GoName:      "LocationId",
										ReadOnly:    true,
										Description: "The ID of the Project's default GCP resource location. The location is one of the available [GCP resource locations](https://firebase.google.com/docs/projects/locations). This field is omitted if the default GCP resource location has not been finalized yet. To set a Project's default GCP resource location, call [`FinalizeDefaultLocation`](../projects.defaultLocation/finalize) after you add Firebase resources to the Project.",
										Immutable:   true,
									},
									"realtimeDatabaseInstance": &dcl.Property{
										Type:        "string",
										GoName:      "RealtimeDatabaseInstance",
										ReadOnly:    true,
										Description: "The default Firebase Realtime Database instance name, in the format: `PROJECT_ID` Though rare, your `projectId` might already be used as the name for an existing Realtime Database instance in another project (learn more about [database sharding](https://firebase.google.com/docs/database/usage/sharding)). In these cases, your `projectId` is appended with a hyphen then five alphanumeric characters to create your default Realtime Database instance name. For example, if your `projectId` is `myproject123`, your default database instance name might be: `myproject123-a5c16`",
										Immutable:   true,
									},
									"storageBucket": &dcl.Property{
										Type:        "string",
										GoName:      "StorageBucket",
										ReadOnly:    true,
										Description: "The default Cloud Storage for Firebase storage bucket, in the format: `PROJECT_ID.appspot.com`",
										Immutable:   true,
									},
								},
							},
							"state": &dcl.Property{
								Type:        "string",
								GoName:      "State",
								GoType:      "FirebaseProjectStateEnum",
								ReadOnly:    true,
								Description: "Output only. The lifecycle state of the Project. Updates to the state must be performed via com.google.cloudresourcemanager.v1.Projects.DeleteProject and com.google.cloudresourcemanager.v1.Projects.UndeleteProject Possible values: STATE_UNSPECIFIED, ACTIVE, DELETED",
								Immutable:   true,
								Enum: []string{
									"STATE_UNSPECIFIED",
									"ACTIVE",
									"DELETED",
								},
							},
						},
					},
				},
			},
		},
	}
}
