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


// +kcc:proto=google.cloud.aiplatform.v1.AnnotationSpec
type AnnotationSpec struct {

	// Required. The user-defined name of the AnnotationSpec.
	//  The name can be up to 128 characters long and can consist of any UTF-8
	//  characters.
	// +kcc:proto:field=google.cloud.aiplatform.v1.AnnotationSpec.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Used to perform consistent read-modify-write updates. If not set,
	//  a blind "overwrite" update happens.
	// +kcc:proto:field=google.cloud.aiplatform.v1.AnnotationSpec.etag
	Etag *string `json:"etag,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.AnnotationSpec
type AnnotationSpecObservedState struct {
	// Output only. Resource name of the AnnotationSpec.
	// +kcc:proto:field=google.cloud.aiplatform.v1.AnnotationSpec.name
	Name *string `json:"name,omitempty"`

	// Output only. Timestamp when this AnnotationSpec was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.AnnotationSpec.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when AnnotationSpec was last updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1.AnnotationSpec.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
