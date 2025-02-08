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


// +kcc:proto=google.cloud.migrationcenter.v1.Source
type Source struct {

	// User-friendly display name.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.Source.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Free-text description.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.Source.description
	Description *string `json:"description,omitempty"`

	// Data source type.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.Source.type
	Type *string `json:"type,omitempty"`

	// The information confidence of the source.
	//  The higher the value, the higher the confidence.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.Source.priority
	Priority *int32 `json:"priority,omitempty"`

	// If `true`, the source is managed by other service(s).
	// +kcc:proto:field=google.cloud.migrationcenter.v1.Source.managed
	Managed *bool `json:"managed,omitempty"`
}

// +kcc:proto=google.cloud.migrationcenter.v1.Source
type SourceObservedState struct {
	// Output only. The full name of the source.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.Source.name
	Name *string `json:"name,omitempty"`

	// Output only. The timestamp when the source was created.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.Source.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the source was last updated.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.Source.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Number of frames that are still being processed.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.Source.pending_frame_count
	PendingFrameCount *int32 `json:"pendingFrameCount,omitempty"`

	// Output only. The number of frames that were reported by the source and
	//  contained errors.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.Source.error_frame_count
	ErrorFrameCount *int32 `json:"errorFrameCount,omitempty"`

	// Output only. The state of the source.
	// +kcc:proto:field=google.cloud.migrationcenter.v1.Source.state
	State *string `json:"state,omitempty"`
}
