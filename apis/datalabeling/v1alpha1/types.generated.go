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


// +kcc:proto=google.cloud.datalabeling.v1beta1.AnnotationSpec
type AnnotationSpec struct {
	// Required. The display name of the AnnotationSpec. Maximum of 64 characters.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.AnnotationSpec.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. User-provided description of the annotation specification.
	//  The description can be up to 10,000 characters long.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.AnnotationSpec.description
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.AnnotationSpecSet
type AnnotationSpecSet struct {
	// Output only. The AnnotationSpecSet resource name in the following format:
	//
	//  "projects/<var>{project_id}</var>/annotationSpecSets/<var>{annotation_spec_set_id}</var>"
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.AnnotationSpecSet.name
	Name *string `json:"name,omitempty"`

	// Required. The display name for AnnotationSpecSet that you define when you
	//  create it. Maximum of 64 characters.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.AnnotationSpecSet.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. User-provided description of the annotation specification set.
	//  The description can be up to 10,000 characters long.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.AnnotationSpecSet.description
	Description *string `json:"description,omitempty"`

	// Required. The array of AnnotationSpecs that you define when you create the
	//  AnnotationSpecSet. These are the possible labels for the labeling task.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.AnnotationSpecSet.annotation_specs
	AnnotationSpecs []AnnotationSpec `json:"annotationSpecs,omitempty"`

	// Output only. The names of any related resources that are blocking changes
	//  to the annotation spec set.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.AnnotationSpecSet.blocking_resources
	BlockingResources []string `json:"blockingResources,omitempty"`
}
