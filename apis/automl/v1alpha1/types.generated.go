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


// +kcc:proto=google.cloud.automl.v1.AnnotationSpec
type AnnotationSpec struct {
	// Output only. Resource name of the annotation spec.
	//  Form:
	//  'projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/annotationSpecs/{annotation_spec_id}'
	// +kcc:proto:field=google.cloud.automl.v1.AnnotationSpec.name
	Name *string `json:"name,omitempty"`

	// Required. The name of the annotation spec to show in the interface. The name can be
	//  up to 32 characters long and must match the regexp `[a-zA-Z0-9_]+`.
	// +kcc:proto:field=google.cloud.automl.v1.AnnotationSpec.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Output only. The number of examples in the parent dataset
	//  labeled by the annotation spec.
	// +kcc:proto:field=google.cloud.automl.v1.AnnotationSpec.example_count
	ExampleCount *int32 `json:"exampleCount,omitempty"`
}
