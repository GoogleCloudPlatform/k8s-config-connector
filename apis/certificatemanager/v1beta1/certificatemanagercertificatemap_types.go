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

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CertificateManagerCertificateMapGVK is already declared in certificatemanagercertificatemap_reference.go

// CertificateManagerCertificateMapSpec defines the desired state of CertificateManagerCertificateMap
// +kcc:spec:proto=google.cloud.certificatemanager.v1.CertificateMap
type CertificateManagerCertificateMapSpec struct {
	// A human-readable description of the resource.
	// +optional
	Description *string `json:"description,omitempty"`

	// The project that this resource belongs to.
	// +required
	ProjectRef refs.ProjectRef `json:"projectRef"`

	// The CertificateManagerCertificateMap name. If not given, the metadata.name will be used.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

// CertificateManagerCertificateMapStatus defines the config connector machine state of CertificateManagerCertificateMap
type CertificateManagerCertificateMapStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// Output only. The creation timestamp of a Certificate Map.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The update timestamp of a Certificate Map.
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. A list of GCLB targets which use this Certificate Map.
	GclbTargets []CertificateMap_GclbTargetObservedState `json:"gclbTargets,omitempty"`
}

// CertificateMap_GclbTargetObservedState defines the GclbTargets
// +kcc:observedstate:proto=google.cloud.certificatemanager.v1.CertificateMap.GclbTarget
type CertificateMap_GclbTargetObservedState struct {
	// IP configurations for this Target Stage.
	IPConfigs []CertificateMap_GclbTarget_IPConfigObservedState `json:"ipConfigs,omitempty"`

	// A HTTPS proxy serving as GCLB target.
	TargetHTTPSProxy *string `json:"targetHttpsProxy,omitempty"`

	// A SSL proxy serving as GCLB target.
	TargetSSLProxy *string `json:"targetSslProxy,omitempty"`
}

// CertificateMap_GclbTarget_IPConfigObservedState defines the IpConfigs
// +kcc:observedstate:proto=google.cloud.certificatemanager.v1.CertificateMap.GclbTarget.IpConfig
type CertificateMap_GclbTarget_IPConfigObservedState struct {
	// An external IP address.
	IPAddress *string `json:"ipAddress,omitempty"`

	// Ports.
	Ports []int32 `json:"ports,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcertificatemanagercertificatemap;gcpcertificatemanagercertificatemaps
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:metadata:labels="internal.cloud.google.com/additional-versions=v1alpha1"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"
// +kubebuilder:storageversion

// CertificateManagerCertificateMap is the Schema for the CertificateManagerCertificateMap API
// +k8s:openapi-gen=true
type CertificateManagerCertificateMap struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CertificateManagerCertificateMapSpec   `json:"spec,omitempty"`
	Status CertificateManagerCertificateMapStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CertificateManagerCertificateMapList contains a list of CertificateManagerCertificateMap
type CertificateManagerCertificateMapList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CertificateManagerCertificateMap `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CertificateManagerCertificateMap{}, &CertificateManagerCertificateMapList{})
}
