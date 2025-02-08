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


// +kcc:proto=google.cloud.vmmigration.v1.TargetProject
type TargetProject struct {

	// The target project ID (number) or project name.
	// +kcc:proto:field=google.cloud.vmmigration.v1.TargetProject.project
	Project *string `json:"project,omitempty"`

	// The target project's description.
	// +kcc:proto:field=google.cloud.vmmigration.v1.TargetProject.description
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.cloud.vmmigration.v1.TargetProject
type TargetProjectObservedState struct {
	// Output only. The name of the target project.
	// +kcc:proto:field=google.cloud.vmmigration.v1.TargetProject.name
	Name *string `json:"name,omitempty"`

	// Output only. The time this target project resource was created (not related
	//  to when the Compute Engine project it points to was created).
	// +kcc:proto:field=google.cloud.vmmigration.v1.TargetProject.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The last time the target project resource was updated.
	// +kcc:proto:field=google.cloud.vmmigration.v1.TargetProject.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
