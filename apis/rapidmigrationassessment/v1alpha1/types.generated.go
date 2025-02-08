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


// +kcc:proto=google.cloud.rapidmigrationassessment.v1.Annotation
type Annotation struct {
	// name of resource.
	// +kcc:proto:field=google.cloud.rapidmigrationassessment.v1.Annotation.name
	Name *string `json:"name,omitempty"`

	// Labels as key value pairs.
	// +kcc:proto:field=google.cloud.rapidmigrationassessment.v1.Annotation.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Type of an annotation.
	// +kcc:proto:field=google.cloud.rapidmigrationassessment.v1.Annotation.type
	Type *string `json:"type,omitempty"`
}

// +kcc:proto=google.cloud.rapidmigrationassessment.v1.Annotation
type AnnotationObservedState struct {
	// Output only. Create time stamp.
	// +kcc:proto:field=google.cloud.rapidmigrationassessment.v1.Annotation.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Update time stamp.
	// +kcc:proto:field=google.cloud.rapidmigrationassessment.v1.Annotation.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
