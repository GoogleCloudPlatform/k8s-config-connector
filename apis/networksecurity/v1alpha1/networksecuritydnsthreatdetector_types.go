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
	computerefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/refs"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var NetworkSecurityDNSThreatDetectorGVK = GroupVersion.WithKind("NetworkSecurityDNSThreatDetector")

// NetworkSecurityDNSThreatDetectorSpec defines the desired state of NetworkSecurityDNSThreatDetector
// +kcc:spec:proto=google.cloud.networksecurity.v1.DnsThreatDetector
type NetworkSecurityDNSThreatDetectorSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The NetworkSecurityDNSThreatDetector name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. A human-readable description of the resource.
	// +kcc:proto:field=google.cloud.networksecurity.v1.DnsThreatDetector.description
	Description *string `json:"description,omitempty"`

	// Optional. Labels as key value pairs.
	// +kcc:proto:field=google.cloud.networksecurity.v1.DnsThreatDetector.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. The provider of the DNS threat detection service.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=INFOBLOX
	// +kcc:proto:field=google.cloud.networksecurity.v1.DnsThreatDetector.provider
	Provider *string `json:"provider"`

	// Optional. The list of VPC networks to exclude from DNS threat detection.
	// If empty, all networks in the project are included.
	// +kcc:proto:field=google.cloud.networksecurity.v1.DnsThreatDetector.excluded_networks
	ExcludedNetworkRefs []computerefs.ComputeNetworkRef `json:"excludedNetworkRefs,omitempty"`
}

// NetworkSecurityDNSThreatDetectorStatus defines the config connector machine state of NetworkSecurityDNSThreatDetector
type NetworkSecurityDNSThreatDetectorStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NetworkSecurityDNSThreatDetector resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *NetworkSecurityDNSThreatDetectorObservedState `json:"observedState,omitempty"`
}

// NetworkSecurityDNSThreatDetectorObservedState is the state of the NetworkSecurityDNSThreatDetector resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.networksecurity.v1.DnsThreatDetector
type NetworkSecurityDNSThreatDetectorObservedState struct {
	// Output only. The timestamp when the resource was created.
	// +kcc:proto:field=google.cloud.networksecurity.v1.DnsThreatDetector.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the resource was updated.
	// +kcc:proto:field=google.cloud.networksecurity.v1.DnsThreatDetector.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetworksecuritydnsthreatdetector;gcpnetworksecuritydnsthreatdetectors
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetworkSecurityDNSThreatDetector is the Schema for the NetworkSecurityDNSThreatDetector API
// +k8s:openapi-gen=true
type NetworkSecurityDNSThreatDetector struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NetworkSecurityDNSThreatDetectorSpec   `json:"spec,omitempty"`
	Status NetworkSecurityDNSThreatDetectorStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NetworkSecurityDNSThreatDetectorList contains a list of NetworkSecurityDNSThreatDetector
type NetworkSecurityDNSThreatDetectorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkSecurityDNSThreatDetector `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkSecurityDNSThreatDetector{}, &NetworkSecurityDNSThreatDetectorList{})
}
