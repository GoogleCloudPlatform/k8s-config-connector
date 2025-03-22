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

// +kcc:proto=google.cloud.dataplex.v1.AssetStatus
type AssetStatus struct {
	// Last update time of the status.
	// +kcc:proto:field=google.cloud.dataplex.v1.AssetStatus.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Number of active assets.
	// +kcc:proto:field=google.cloud.dataplex.v1.AssetStatus.active_assets
	ActiveAssets *int32 `json:"activeAssets,omitempty"`

	// Number of assets that are in process of updating the security policy on
	//  attached resources.
	// +kcc:proto:field=google.cloud.dataplex.v1.AssetStatus.security_policy_applying_assets
	SecurityPolicyApplyingAssets *int32 `json:"securityPolicyApplyingAssets,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Lake
type Lake struct {

	// Optional. User friendly display name.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. User-defined labels for the lake.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Description of the lake.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.description
	Description *string `json:"description,omitempty"`

	// Optional. Settings to manage lake and Dataproc Metastore service instance
	//  association.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.metastore
	Metastore *Lake_Metastore `json:"metastore,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Lake.MetastoreStatus
type Lake_MetastoreStatus struct {
	// Current state of association.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.MetastoreStatus.state
	State *string `json:"state,omitempty"`

	// Additional information about the current status.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.MetastoreStatus.message
	Message *string `json:"message,omitempty"`

	// Last update time of the metastore status of the lake.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.MetastoreStatus.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// The URI of the endpoint used to access the Metastore service.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.MetastoreStatus.endpoint
	Endpoint *string `json:"endpoint,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Lake
type LakeObservedState struct {
	// Output only. The relative resource name of the lake, of the form:
	//  `projects/{project_number}/locations/{location_id}/lakes/{lake_id}`.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.name
	Name *string `json:"name,omitempty"`

	// Output only. System generated globally unique ID for the lake. This ID will
	//  be different if the lake is deleted and re-created with the same name.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The time when the lake was created.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the lake was last updated.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Current state of the lake.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.state
	State *string `json:"state,omitempty"`

	// Output only. Service account associated with this lake. This service
	//  account must be authorized to access or operate on resources managed by the
	//  lake.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// Output only. Aggregated status of the underlying assets of the lake.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.asset_status
	AssetStatus *AssetStatus `json:"assetStatus,omitempty"`

	// Output only. Metastore status of the lake.
	// +kcc:proto:field=google.cloud.dataplex.v1.Lake.metastore_status
	MetastoreStatus *Lake_MetastoreStatus `json:"metastoreStatus,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Content
type Content struct {

	// Required. The path for the Content file, represented as directory
	//  structure. Unique within a lake. Limited to alphanumerics, hyphens,
	//  underscores, dots and slashes.
	// +kcc:proto:field=google.cloud.dataplex.v1.Content.path
	Path *string `json:"path,omitempty"`

	// Optional. User defined labels for the content.
	// +kcc:proto:field=google.cloud.dataplex.v1.Content.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Description of the content.
	// +kcc:proto:field=google.cloud.dataplex.v1.Content.description
	Description *string `json:"description,omitempty"`

	// Required. Content data in string format.
	// +kcc:proto:field=google.cloud.dataplex.v1.Content.data_text
	DataText *string `json:"dataText,omitempty"`

	// Sql Script related configurations.
	// +kcc:proto:field=google.cloud.dataplex.v1.Content.sql_script
	SQLScript *Content_SQLScript `json:"sqlScript,omitempty"`

	// Notebook related configurations.
	// +kcc:proto:field=google.cloud.dataplex.v1.Content.notebook
	Notebook *Content_Notebook `json:"notebook,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Content.Notebook
type Content_Notebook struct {
	// Required. Kernel Type of the notebook.
	// +kcc:proto:field=google.cloud.dataplex.v1.Content.Notebook.kernel_type
	KernelType *string `json:"kernelType,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Content.SqlScript
type Content_SQLScript struct {
	// Required. Query Engine to be used for the Sql Query.
	// +kcc:proto:field=google.cloud.dataplex.v1.Content.SqlScript.engine
	Engine *string `json:"engine,omitempty"`
}

// +kcc:proto=google.cloud.dataplex.v1.Content
type ContentObservedState struct {
	// Output only. The relative resource name of the content, of the form:
	//  projects/{project_id}/locations/{location_id}/lakes/{lake_id}/content/{content_id}
	// +kcc:proto:field=google.cloud.dataplex.v1.Content.name
	Name *string `json:"name,omitempty"`

	// Output only. System generated globally unique ID for the content. This ID
	//  will be different if the content is deleted and re-created with the same
	//  name.
	// +kcc:proto:field=google.cloud.dataplex.v1.Content.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. Content creation time.
	// +kcc:proto:field=google.cloud.dataplex.v1.Content.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the content was last updated.
	// +kcc:proto:field=google.cloud.dataplex.v1.Content.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
