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

var NetworkSecuritySecurityProfileGroupGVK = GroupVersion.WithKind("NetworkSecuritySecurityProfileGroup")

// NetworkSecuritySecurityProfileGroupSpec defines the desired state of NetworkSecuritySecurityProfileGroup
// +kcc:spec:proto=google.cloud.networksecurity.v1.SecurityProfileGroup
type NetworkSecuritySecurityProfileGroupSpec struct {
	// The project that this resource belongs to.
	// +optional
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef,omitempty"`

	// The organization that this resource belongs to.
	// +optional
	OrganizationRef *refsv1beta1.OrganizationRef `json:"organizationRef,omitempty"`

	// The location of this resource.
	// +required
	Location *string `json:"location"`

	// The NetworkSecuritySecurityProfileGroup name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. An optional description of the profile group. Max length 2048
	// characters.
	Description *string `json:"description,omitempty"`

	// Optional. Labels as key value pairs.
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Reference to a SecurityProfile with the ThreatPrevention
	// configuration.
	ThreatPreventionProfileRef *refsv1beta1.NetworkSecuritySecurityProfileRef `json:"threatPreventionProfileRef,omitempty"`

	// Optional. Reference to a SecurityProfile with the CustomMirroring
	// configuration.
	CustomMirroringProfileRef *refsv1beta1.NetworkSecuritySecurityProfileRef `json:"customMirroringProfileRef,omitempty"`

	// Optional. Reference to a SecurityProfile with the CustomIntercept
	// configuration.
	CustomInterceptProfileRef *refsv1beta1.NetworkSecuritySecurityProfileRef `json:"customInterceptProfileRef,omitempty"`

	// Optional. Reference to a SecurityProfile with the UrlFiltering
	// configuration.
	URLFilteringProfileRef *refsv1beta1.NetworkSecuritySecurityProfileRef `json:"urlFilteringProfileRef,omitempty"`
}

// NetworkSecuritySecurityProfileGroupStatus defines the config connector machine state of NetworkSecuritySecurityProfileGroup
type NetworkSecuritySecurityProfileGroupStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NetworkSecuritySecurityProfileGroup resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *NetworkSecuritySecurityProfileGroupObservedState `json:"observedState,omitempty"`
}

// NetworkSecuritySecurityProfileGroupObservedState is the state of the NetworkSecuritySecurityProfileGroup resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.networksecurity.v1.SecurityProfileGroup
type NetworkSecuritySecurityProfileGroupObservedState struct {
	// Output only. Resource creation timestamp.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Last resource update timestamp.
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. This checksum is computed by the server based on the value of
	// other fields, and may be sent on update and delete requests to ensure the
	// client has an up-to-date value before proceeding.
	Etag *string `json:"etag,omitempty"`

	// Output only. Identifier used by the data-path. Unique within {container,
	// location}.
	DataPathID *int64 `json:"dataPathID,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetworksecuritysecurityprofilegroup;gcpnetworksecuritysecurityprofilegroups
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetworkSecuritySecurityProfileGroup is the Schema for the NetworkSecuritySecurityProfileGroup API
// +k8s:openapi-gen=true
type NetworkSecuritySecurityProfileGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NetworkSecuritySecurityProfileGroupSpec   `json:"spec,omitempty"`
	Status NetworkSecuritySecurityProfileGroupStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NetworkSecuritySecurityProfileGroupList contains a list of NetworkSecuritySecurityProfileGroup
type NetworkSecuritySecurityProfileGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkSecuritySecurityProfileGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkSecuritySecurityProfileGroup{}, &NetworkSecuritySecurityProfileGroupList{})
}
