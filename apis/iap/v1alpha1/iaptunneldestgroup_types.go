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

// API sources for IAPTunnelDestGroup, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/iap/v1/service.proto
// +kcc:source:service-docs=https://cloud.google.com/iap
// +kcc:source:resource-docs=https://docs.cloud.google.com/iap/docs/reference/rest/v1/projects.iap_tunnel.locations.destGroups

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var IAPTunnelDestGroupGVK = GroupVersion.WithKind("IAPTunnelDestGroup")

// IAPTunnelDestGroupSpec defines the desired state of IAPTunnelDestGroup
// +kcc:spec:proto=google.cloud.iap.v1.TunnelDestGroup
// +kcc:required-from-proto
type IAPTunnelDestGroupSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The IAPTunnelDestGroup name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. Unordered list. List of CIDRs that this group applies to.
	// +kcc:proto:field=google.cloud.iap.v1.TunnelDestGroup.cidrs
	CIDRs []string `json:"cidrs,omitempty"`

	// Optional. Unordered list. List of FQDNs that this group applies to.
	// +kcc:proto:field=google.cloud.iap.v1.TunnelDestGroup.fqdns
	FQDNs []string `json:"fqdns,omitempty"`
}

// IAPTunnelDestGroupStatus defines the config connector machine state of IAPTunnelDestGroup
type IAPTunnelDestGroupStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the IAPTunnelDestGroup resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	// ObservedState *IAPTunnelDestGroupObservedState `json:"observedState,omitempty"`
}

// IAPTunnelDestGroupObservedState is the state of the IAPTunnelDestGroup resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.iap.v1.TunnelDestGroup
// type IAPTunnelDestGroupObservedState struct {
// }

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpiaptunneldestgroup;gcpiaptunneldestgroups
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// IAPTunnelDestGroup is the Schema for the IAPTunnelDestGroup API
// +k8s:openapi-gen=true
type IAPTunnelDestGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   IAPTunnelDestGroupSpec   `json:"spec,omitempty"`
	Status IAPTunnelDestGroupStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// IAPTunnelDestGroupList contains a list of IAPTunnelDestGroup
type IAPTunnelDestGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IAPTunnelDestGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&IAPTunnelDestGroup{}, &IAPTunnelDestGroupList{})
}
