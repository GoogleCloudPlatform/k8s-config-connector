// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1


// +kcc:proto=google.cloud.connectors.v1.Connector
type Connector struct {
}

// +kcc:proto=google.cloud.connectors.v1.Connector
type ConnectorObservedState struct {
	// Output only. Resource name of the Connector.
	//  Format:
	//  projects/{project}/locations/{location}/providers/{provider}/connectors/{connector}
	//  Only global location is supported for Connector resource.
	// +kcc:proto:field=google.cloud.connectors.v1.Connector.name
	Name *string `json:"name,omitempty"`

	// Output only. Created time.
	// +kcc:proto:field=google.cloud.connectors.v1.Connector.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Updated time.
	// +kcc:proto:field=google.cloud.connectors.v1.Connector.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Resource labels to represent user-provided metadata.
	//  Refer to cloud documentation on labels for more details.
	//  https://cloud.google.com/compute/docs/labeling-resources
	// +kcc:proto:field=google.cloud.connectors.v1.Connector.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Output only. Link to documentation page.
	// +kcc:proto:field=google.cloud.connectors.v1.Connector.documentation_uri
	DocumentationURI *string `json:"documentationURI,omitempty"`

	// Output only. Link to external page.
	// +kcc:proto:field=google.cloud.connectors.v1.Connector.external_uri
	ExternalURI *string `json:"externalURI,omitempty"`

	// Output only. Description of the resource.
	// +kcc:proto:field=google.cloud.connectors.v1.Connector.description
	Description *string `json:"description,omitempty"`

	// Output only. Cloud storage location of icons etc consumed by UI.
	// +kcc:proto:field=google.cloud.connectors.v1.Connector.web_assets_location
	WebAssetsLocation *string `json:"webAssetsLocation,omitempty"`

	// Output only. Display name.
	// +kcc:proto:field=google.cloud.connectors.v1.Connector.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Output only. Flag to mark the version indicating the launch stage.
	// +kcc:proto:field=google.cloud.connectors.v1.Connector.launch_stage
	LaunchStage *string `json:"launchStage,omitempty"`
}
