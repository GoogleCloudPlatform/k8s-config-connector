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

// +generated:types
// krm.group: clouddms.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.clouddms.v1
// resource: CloudDMSConversionWorkspace:ConversionWorkspace

package v1alpha1

// +kcc:proto=google.cloud.clouddms.v1.ConversionWorkspace
type ConversionWorkspace struct {
	// Full name of the workspace resource, in the form of:
	//  projects/{project}/locations/{location}/conversionWorkspaces/{conversion_workspace}.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConversionWorkspace.name
	Name *string `json:"name,omitempty"`

	// Required. The source engine details.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConversionWorkspace.source
	Source *DatabaseEngineInfo `json:"source,omitempty"`

	// Required. The destination engine details.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConversionWorkspace.destination
	Destination *DatabaseEngineInfo `json:"destination,omitempty"`

	// Optional. A generic list of settings for the workspace.
	//  The settings are database pair dependant and can indicate default behavior
	//  for the mapping rules engine or turn on or off specific features.
	//  Such examples can be: convert_foreign_key_to_interleave=true,
	//  skip_triggers=false, ignore_non_table_synonyms=true
	// +kcc:proto:field=google.cloud.clouddms.v1.ConversionWorkspace.global_settings
	GlobalSettings map[string]string `json:"globalSettings,omitempty"`

	// Optional. The display name for the workspace.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConversionWorkspace.display_name
	DisplayName *string `json:"displayName,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.DatabaseEngineInfo
type DatabaseEngineInfo struct {
	// Required. Engine type.
	// +kcc:proto:field=google.cloud.clouddms.v1.DatabaseEngineInfo.engine
	Engine *string `json:"engine,omitempty"`

	// Required. Engine named version, for example 12.c.1.
	// +kcc:proto:field=google.cloud.clouddms.v1.DatabaseEngineInfo.version
	Version *string `json:"version,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.ConversionWorkspace
type ConversionWorkspaceObservedState struct {
	// Output only. Whether the workspace has uncommitted changes (changes which
	//  were made after the workspace was committed).
	// +kcc:proto:field=google.cloud.clouddms.v1.ConversionWorkspace.has_uncommitted_changes
	HasUncommittedChanges *bool `json:"hasUncommittedChanges,omitempty"`

	// Output only. The latest commit ID.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConversionWorkspace.latest_commit_id
	LatestCommitID *string `json:"latestCommitID,omitempty"`

	// Output only. The timestamp when the workspace was committed.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConversionWorkspace.latest_commit_time
	LatestCommitTime *string `json:"latestCommitTime,omitempty"`

	// Output only. The timestamp when the workspace resource was created.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConversionWorkspace.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the workspace resource was last updated.
	// +kcc:proto:field=google.cloud.clouddms.v1.ConversionWorkspace.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
