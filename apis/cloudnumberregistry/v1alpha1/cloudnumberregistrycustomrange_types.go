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

var CloudNumberRegistryCustomRangeGVK = GroupVersion.WithKind("CloudNumberRegistryCustomRange")

// +kcc:proto=google.cloud.numberregistry.v1alpha.Attribute
type Attribute struct {
	// Required. The key of the attribute.
	// +kcc:proto:field=google.cloud.numberregistry.v1alpha.Attribute.key
	Key *string `json:"key,omitempty"`

	// Required. The value of the attribute.
	// +kcc:proto:field=google.cloud.numberregistry.v1alpha.Attribute.value
	Value *string `json:"value,omitempty"`
}

// CloudNumberRegistryCustomRangeSpec defines the desired state of CloudNumberRegistryCustomRange
// +kcc:spec:proto=google.cloud.numberregistry.v1alpha.CustomRange
type CloudNumberRegistryCustomRangeSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The CloudNumberRegistryCustomRange name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. The IPv4 CIDR range of the CustomRange.
	// +kcc:proto:field=google.cloud.numberregistry.v1alpha.CustomRange.ipv4_cidr_range
	IPV4CIDRRange *string `json:"ipv4CIDRRange,omitempty"`

	// Optional. The IPv6 CIDR range of the CustomRange.
	// +kcc:proto:field=google.cloud.numberregistry.v1alpha.CustomRange.ipv6_cidr_range
	IPV6CIDRRange *string `json:"ipv6CIDRRange,omitempty"`

	// Optional. The resource name of the Realm associated with the CustomRange,
	//  in the format `projects/{project}/locations/{location}/realms/{realm}`. The
	//  Realm must be in the same project as the CustomRange. This field must not
	//  be set if the `parent_range` field is set, as the Realm will be inherited
	//  from the parent CustomRange.
	// +kcc:proto:field=google.cloud.numberregistry.v1alpha.CustomRange.realm
	Realm *string `json:"realm,omitempty"`

	// Optional. The resource name of the parent CustomRange, in the format
	//  `projects/{project}/locations/{location}/customRanges/{custom_range}`.
	//  If specified, the parent CustomRange must be in the same RegistryBook.
	//  This field is mutually exclusive with the `realm` field, as the Realm is
	//  inherited from the parent CustomRange.
	// +kcc:proto:field=google.cloud.numberregistry.v1alpha.CustomRange.parent_range
	ParentRange *string `json:"parentRange,omitempty"`

	// Optional. The attributes of the CustomRange.
	// +kcc:proto:field=google.cloud.numberregistry.v1alpha.CustomRange.attributes
	Attributes []Attribute `json:"attributes,omitempty"`

	// Optional. The description of the CustomRange.
	// +kcc:proto:field=google.cloud.numberregistry.v1alpha.CustomRange.description
	Description *string `json:"description,omitempty"`

	// Optional. User-defined labels.
	// +kcc:proto:field=google.cloud.numberregistry.v1alpha.CustomRange.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// CloudNumberRegistryCustomRangeStatus defines the config connector machine state of CloudNumberRegistryCustomRange
type CloudNumberRegistryCustomRangeStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the CloudNumberRegistryCustomRange resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *CloudNumberRegistryCustomRangeObservedState `json:"observedState,omitempty"`
}

// CloudNumberRegistryCustomRangeObservedState is the state of the CloudNumberRegistryCustomRange resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.numberregistry.v1alpha.CustomRange
type CloudNumberRegistryCustomRangeObservedState struct {
	// Output only. The RegistryBook of the CustomRange. This field is inherited
	//  from the Realm or parent CustomRange depending on which one is specified.
	// +kcc:proto:field=google.cloud.numberregistry.v1alpha.CustomRange.registry_book
	RegistryBook *string `json:"registryBook,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcloudnumberregistrycustomrange;gcpcloudnumberregistrycustomranges
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CloudNumberRegistryCustomRange is the Schema for the CloudNumberRegistryCustomRange API
// +k8s:openapi-gen=true
type CloudNumberRegistryCustomRange struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CloudNumberRegistryCustomRangeSpec   `json:"spec,omitempty"`
	Status CloudNumberRegistryCustomRangeStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CloudNumberRegistryCustomRangeList contains a list of CloudNumberRegistryCustomRange
type CloudNumberRegistryCustomRangeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudNumberRegistryCustomRange `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudNumberRegistryCustomRange{}, &CloudNumberRegistryCustomRangeList{})
}
