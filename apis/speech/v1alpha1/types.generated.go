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


// +kcc:proto=google.cloud.speech.v1.CustomClass
type CustomClass struct {
	// The resource name of the custom class.
	// +kcc:proto:field=google.cloud.speech.v1.CustomClass.name
	Name *string `json:"name,omitempty"`

	// If this custom class is a resource, the custom_class_id is the resource id
	//  of the CustomClass. Case sensitive.
	// +kcc:proto:field=google.cloud.speech.v1.CustomClass.custom_class_id
	CustomClassID *string `json:"customClassID,omitempty"`

	// A collection of class items.
	// +kcc:proto:field=google.cloud.speech.v1.CustomClass.items
	Items []CustomClass_ClassItem `json:"items,omitempty"`
}

// +kcc:proto=google.cloud.speech.v1.CustomClass.ClassItem
type CustomClass_ClassItem struct {
	// The class item's value.
	// +kcc:proto:field=google.cloud.speech.v1.CustomClass.ClassItem.value
	Value *string `json:"value,omitempty"`
}
