// Copyright 2025 Google LLC
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

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/tags/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var TagsLocationTagBindingGVK = GroupVersion.WithKind("TagsLocationTagBinding")

// TagsLocationTagBindingSpec defines the desired state of TagsLocationTagBinding
// +kcc:spec:proto=google.cloud.resourcemanager.v3.TagBinding
type TagsLocationTagBindingSpec struct {
	// +kcc:ref=Project
	ParentRef *v1beta1.TagsTagBindingParentRef `json:"parentRef"`

	// The location for the resource being tagged.
	// +required
	Location *string `json:"location"`

	// +kcc:ref=TagsTagValue
	TagValueRef *v1beta1.TagsTagValueRef `json:"tagValueRef"`

	// The service-generated name of the resource. Used for acquisition only. Leave unset to create a new
	// resource.
	ResourceID *string `json:"resourceID,omitempty"`
}

// TagsLocationTagBindingStatus defines the config connector machine state of TagsLocationTagBinding
type TagsLocationTagBindingStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// The generated id for the TagBinding. This is a string of the form: tagBindings/{full-resource-name}/{tag-value-name}.
	Name *string `json:"name,omitempty"`

	// A unique specifier for the TagsLocationTagBinding resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	// ObservedState *TagsLocationTagBindingObservedState `json:"observedState,omitempty"`
}

// TagsLocationTagBindingObservedState is the state of the TagsLocationTagBinding resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.resourcemanager.v3.TagBinding
type TagsLocationTagBindingObservedState struct {
	// Output only. The name of the TagBinding. This is a String of the form:
	// `tagBindings/{full-resource-name}/{tag-value-name}`
	// (e.g. `tagBindings/%2F%2Fcloudresourcemanager.googleapis.com%2Fprojects%2F123/tagValues/456`).
	Name *string `json:"name,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcptagslocationtagbinding;gcptagslocationtagbindings
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// TagsLocationTagBinding is the Schema for the TagsLocationTagBinding API
// +k8s:openapi-gen=true
type TagsLocationTagBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   TagsLocationTagBindingSpec   `json:"spec,omitempty"`
	Status TagsLocationTagBindingStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TagsLocationTagBindingList contains a list of TagsLocationTagBinding
type TagsLocationTagBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TagsLocationTagBinding `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TagsLocationTagBinding{}, &TagsLocationTagBindingList{})
}
