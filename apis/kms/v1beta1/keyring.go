// Copyright 2024 Google LLC
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

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	KMSKeyRingGVK = schema.GroupVersionKind{
		Group:   SchemeGroupVersion.Group,
		Version: SchemeGroupVersion.Version,
		Kind:    "KMSKeyRing",
	}
)

// +kcc:spec:proto=google.cloud.kms.v1.KeyRing
type KMSKeyRingSpec struct {
	// /* Immutable. The Project that this resource belongs to. */
	// ProjectRef refs.ProjectRef `json:"projectRef"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// Immutable. The location for the KeyRing.
	//                  A full list of valid locations can be found by running 'gcloud kms locations list'.
	// +required
	Location *string `json:"location,omitempty"`
}

// +kcc:status:proto=google.cloud.kms.v1.KeyRing
type KMSKeyRingStatus struct {
	/* Conditions represent the latest available observations of the
	   KMSKeyRing's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	//   // Output only. The resource name for the
	// //  [KeyRing][google.cloud.kms.v1.KeyRing] in the format
	// //  `projects/*/locations/*/keyRings/*`.
	// Name *string `json:"name,omitempty"`

	// The self link of the created KeyRing in the format projects/{project}/locations/{location}/keyRings/{name}.
	SelfLink *string `json:"selfLink,omitempty"`

	// // Output only. The time at which this [KeyRing][google.cloud.kms.v1.KeyRing]
	// //  was created.
	// CreateTime *string `json:"createTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpkmskeyring;gcpkmskeyrings
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true";"cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=stable";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// KMSKeyRing represents a KMS KeyRing.
// +k8s:openapi-gen=true
type KMSKeyRing struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec KMSKeyRingSpec `json:"spec,omitempty"`

	Status KMSKeyRingStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KMSKeyRingList contains a list of KMSKeyRing
type KMSKeyRingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KMSKeyRing `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KMSKeyRing{}, &KMSKeyRingList{})
}
