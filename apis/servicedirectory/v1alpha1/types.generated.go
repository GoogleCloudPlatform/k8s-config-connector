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


// +kcc:proto=google.cloud.servicedirectory.v1.Namespace
type Namespace struct {
	// Immutable. The resource name for the namespace in the format
	//  `projects/*/locations/*/namespaces/*`.
	// +kcc:proto:field=google.cloud.servicedirectory.v1.Namespace.name
	Name *string `json:"name,omitempty"`

	// Optional. Resource labels associated with this namespace.
	//  No more than 64 user labels can be associated with a given resource. Label
	//  keys and values can be no longer than 63 characters.
	// +kcc:proto:field=google.cloud.servicedirectory.v1.Namespace.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// +kcc:proto=google.cloud.servicedirectory.v1.Namespace
type NamespaceObservedState struct {
	// Output only. The globally unique identifier of the namespace in the UUID4
	//  format.
	// +kcc:proto:field=google.cloud.servicedirectory.v1.Namespace.uid
	Uid *string `json:"uid,omitempty"`
}
