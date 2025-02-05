// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package v1alpha1

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/spanner/v1beta1"
)

type InstanceParent struct {
	// Immutable. The Project that this resource belongs to.
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="ResourceID field is immutable"
	ProjectRef *refs.ProjectRef `json:"projectRef"`

	// The reference to the parent instance
	InstanceRef *v1beta1.SpannerInstanceRef `json:"instanceRef,omitempty"`
}

type InstanceDatabaseParent struct {
	// Immutable. The Project that this resource belongs to.
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="ResourceID field is immutable"
	ProjectRef *refs.ProjectRef `json:"projectRef"`

	// The reference to the parent instance
	InstanceRef *v1beta1.SpannerInstanceRef `json:"instanceRef,omitempty"`

	// The reference to the parent database
	DatabaseRef *DatabaseRef `json:"databaseRef,omitempty"`
}
