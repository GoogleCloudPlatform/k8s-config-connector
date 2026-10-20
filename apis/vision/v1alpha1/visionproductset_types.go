// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// API sources for VisionProductSet, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/vision/v1/product_search_service.proto
// +kcc:source:service-docs=https://cloud.google.com/vision/
// +kcc:source:resource-docs=https://docs.cloud.google.com/vision/docs/reference/rest/v1/projects.locations.productSets

package v1alpha1

import (
	common "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var VisionProductSetGVK = GroupVersion.WithKind("VisionProductSet")

// VisionProductSetSpec defines the desired state of VisionProductSet
// +kcc:spec:proto=google.cloud.vision.v1.ProductSet
// +kcc:required-from-proto
type VisionProductSetSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The VisionProductSet name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// The user-provided name for this ProductSet. Must not be empty. Must be at
	//  most 4096 characters long.
	// +kcc:proto:field=google.cloud.vision.v1.ProductSet.display_name
	DisplayName *string `json:"displayName,omitempty"`
}

// VisionProductSetStatus defines the config connector machine state of VisionProductSet
type VisionProductSetStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VisionProductSet resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VisionProductSetObservedState `json:"observedState,omitempty"`
}

// VisionProductSetObservedState is the state of the VisionProductSet resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.vision.v1.ProductSet
type VisionProductSetObservedState struct {
	// Output only. The time at which this ProductSet was last indexed. Query
	//  results will reflect all updates before this time. If this ProductSet has
	//  never been indexed, this timestamp is the default value
	//  "1970-01-01T00:00:00Z".
	//
	//  This field is ignored when creating a ProductSet.
	// +kcc:proto:field=google.cloud.vision.v1.ProductSet.index_time
	IndexTime *string `json:"indexTime,omitempty"`

	// Output only. If there was an error with indexing the product set, the field
	//  is populated.
	//
	//  This field is ignored when creating a ProductSet.
	// +kcc:proto:field=google.cloud.vision.v1.ProductSet.index_error
	IndexError *common.Status `json:"indexError,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvisionproductset;gcpvisionproductsets
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VisionProductSet is the Schema for the VisionProductSet API
// +k8s:openapi-gen=true
type VisionProductSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VisionProductSetSpec   `json:"spec,omitempty"`
	Status VisionProductSetStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VisionProductSetList contains a list of VisionProductSet
type VisionProductSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VisionProductSet `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VisionProductSet{}, &VisionProductSetList{})
}
