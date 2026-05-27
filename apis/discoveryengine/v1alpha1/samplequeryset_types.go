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

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DiscoveryEngineSampleQuerySetGVK = GroupVersion.WithKind("DiscoveryEngineSampleQuerySet")

// DiscoveryEngineSampleQuerySetSpec defines the desired state of DiscoveryEngineSampleQuerySet
// +kcc:spec:proto=google.cloud.discoveryengine.v1.SampleQuerySet
type DiscoveryEngineSampleQuerySetSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// Immutable. The location of this resource.
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Location field is immutable"
	// +required
	Location string `json:"location"`

	// The DiscoveryEngineSampleQuerySet name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The sample query set display name.
	//
	// This field must be a UTF-8 encoded string with a length limit of 128
	// characters.
	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName"`

	// The description of the
	// [SampleQuerySet][google.cloud.discoveryengine.v1.SampleQuerySet].
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty"`
}

// DiscoveryEngineSampleQuerySetStatus defines the config connector machine state of DiscoveryEngineSampleQuerySet
type DiscoveryEngineSampleQuerySetStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DiscoveryEngineSampleQuerySet resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DiscoveryEngineSampleQuerySetObservedState `json:"observedState,omitempty"`
}

// DiscoveryEngineSampleQuerySetObservedState is the state of the DiscoveryEngineSampleQuerySet resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.discoveryengine.v1.SampleQuerySet
type DiscoveryEngineSampleQuerySetObservedState struct {
	// Output only. Timestamp the
	// [SampleQuerySet][google.cloud.discoveryengine.v1.SampleQuerySet] was
	// created at.
	CreateTime *string `json:"createTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdiscoveryenginesamplequeryset;gcpdiscoveryenginesamplequerysets
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DiscoveryEngineSampleQuerySet is the Schema for the DiscoveryEngineSampleQuerySet API
// +k8s:openapi-gen=true
type DiscoveryEngineSampleQuerySet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DiscoveryEngineSampleQuerySetSpec   `json:"spec,omitempty"`
	Status DiscoveryEngineSampleQuerySetStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DiscoveryEngineSampleQuerySetList contains a list of DiscoveryEngineSampleQuerySet
type DiscoveryEngineSampleQuerySetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DiscoveryEngineSampleQuerySet `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DiscoveryEngineSampleQuerySet{}, &DiscoveryEngineSampleQuerySetList{})
}
