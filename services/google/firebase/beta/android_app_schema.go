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

func DCLAndroidAppSchema() *dcl.Schema {
	return &dcl.Schema{
		Info: &dcl.Info{
			Title:      "Firebase/AndroidApp",
			StructName: "AndroidApp",
			Reference: &dcl.Link{
				Text: "Firebase AndroidApp API Documentation",
				URL:  "https://firebase.google.com/docs/projects/api/reference/rest#rest-resource:-v1beta1.projects.androidapps",
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
				Description: "The function used to get information about a AndroidApp",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "androidApp",
						Required:    true,
						Description: "A full instance of a AndroidApp",
					},
				},
			},
			Apply: &dcl.Path{
				Description: "The function used to apply information about a AndroidApp",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "androidApp",
						Required:    true,
						Description: "A full instance of a AndroidApp",
					},
				},
			},
			Delete: &dcl.Path{
				Description: "The function used to delete a AndroidApp",
				Parameters: []dcl.PathParameters{
					dcl.PathParameters{
						Name:        "androidApp",
						Required:    true,
						Description: "A full instance of a AndroidApp",
					},
				},
			},
			DeleteAll: &dcl.Path{
				Description: "The function used to delete all AndroidApp",
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
				Description: "The function used to list information about many AndroidApp",
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
				"AndroidApp": &dcl.Component{
					Title:           "AndroidApp",
					ID:              "projects/{{project}}/androidApps/{{name}}",
					ParentContainer: "project",
					HasCreate:       true,
					SchemaProperty: dcl.Property{
						Type: "object",
						Required: []string{
							"packageName",
							"project",
						},
						Properties: map[string]*dcl.Property{
							"apiKeyId": &dcl.Property{
								Type:        "string",
								GoName:      "ApiKeyId",
								Description: "The key_id of the GCP ApiKey associated with this App. If set must have no restrictions, or only have restrictions that are valid for the associated Firebase App. Cannot be set in create requests, instead an existing valid API Key will be chosen, or if no valid API Keys exist, one will be provisioned for you. Cannot be set to an empty value in update requests.",
								ResourceReferences: []*dcl.PropertyResourceReference{
									&dcl.PropertyResourceReference{
										Resource: "Apikeys/Key",
										Field:    "name",
									},
								},
							},
							"appId": &dcl.Property{
								Type:        "string",
								GoName:      "AppId",
								ReadOnly:    true,
								Description: "Output only. Immutable. The globally unique, Firebase-assigned identifier for the `AndroidApp`. This identifier should be treated as an opaque token, as the data format is not specified.",
								Immutable:   true,
							},
							"displayName": &dcl.Property{
								Type:        "string",
								GoName:      "DisplayName",
								Description: "The user-assigned display name for the `AndroidApp`.",
							},
							"name": &dcl.Property{
								Type:                     "string",
								GoName:                   "Name",
								Description:              "The resource name of the AndroidApp, in the format: `projects/PROJECT_IDENTIFIER/androidApps/APP_ID` * PROJECT_IDENTIFIER: the parent Project's [`ProjectNumber`](../projects#FirebaseProject.FIELDS.project_number) ***(recommended)*** or its [`ProjectId`](../projects#FirebaseProject.FIELDS.project_id). Learn more about using project identifiers in Google's [AIP 2510 standard](https://google.aip.dev/cloud/2510). Note that the value for PROJECT_IDENTIFIER in any response body will be the `ProjectId`. * APP_ID: the globally unique, Firebase-assigned identifier for the App (see [`appId`](../projects.androidApps#AndroidApp.FIELDS.app_id)).",
								Immutable:                true,
								ServerGeneratedParameter: true,
							},
							"packageName": &dcl.Property{
								Type:        "string",
								GoName:      "PackageName",
								Description: "Immutable. The canonical package name of the Android app as would appear in the Google Play Developer Console.",
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
							"projectId": &dcl.Property{
								Type:        "string",
								GoName:      "ProjectId",
								ReadOnly:    true,
								Description: "Output only. Immutable. A user-assigned unique identifier of the parent FirebaseProject for the `AndroidApp`.",
								Immutable:   true,
							},
						},
					},
				},
			},
		},
	}
}
