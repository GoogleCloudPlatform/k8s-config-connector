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

var NetworkSecurityUrlListGVK = GroupVersion.WithKind("NetworkSecurityUrlList")

// NetworkSecurityUrlListSpec defines the desired state of NetworkSecurityUrlList
// +kcc:spec:proto=google.cloud.networksecurity.v1.UrlList
type NetworkSecurityUrlListSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The NetworkSecurityUrlList name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Free-text description of the resource.
	// +kcc:proto:field=google.cloud.networksecurity.v1.UrlList.description
	Description *string `json:"description,omitempty"`

	// Required. FQDNs and URLs.
	// +kubebuilder:validation:Required
	// +kcc:proto:field=google.cloud.networksecurity.v1.UrlList.values
	Values []string `json:"values"`
}

// NetworkSecurityUrlListStatus defines the config connector machine state of NetworkSecurityUrlList
type NetworkSecurityUrlListStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NetworkSecurityUrlList resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *NetworkSecurityUrlListObservedState `json:"observedState,omitempty"`
}

// NetworkSecurityUrlListObservedState is the state of the NetworkSecurityUrlList resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.networksecurity.v1.UrlList
type NetworkSecurityUrlListObservedState struct {
	// Output only. Time when the security policy was created.
	// +kcc:proto:field=google.cloud.networksecurity.v1.UrlList.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time when the security policy was updated.
	// +kcc:proto:field=google.cloud.networksecurity.v1.UrlList.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetworksecurityurllist;gcpnetworksecurityurllists
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetworkSecurityUrlList is the Schema for the NetworkSecurityUrlList API
// +k8s:openapi-gen=true
type NetworkSecurityUrlList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NetworkSecurityUrlListSpec   `json:"spec,omitempty"`
	Status NetworkSecurityUrlListStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NetworkSecurityUrlListList contains a list of NetworkSecurityUrlList
type NetworkSecurityUrlListList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkSecurityUrlList `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkSecurityUrlList{}, &NetworkSecurityUrlListList{})
}
