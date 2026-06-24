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
	refsv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ArtifactRegistryVPCSCConfigGVK = GroupVersion.WithKind("ArtifactRegistryVPCSCConfig")

// ArtifactRegistryVPCSCConfigSpec defines the desired state of ArtifactRegistryVPCSCConfig
// +kcc:spec:proto=google.devtools.artifactregistry.v1.VPCSCConfig
type ArtifactRegistryVPCSCConfigSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1alpha1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The ArtifactRegistryVPCSCConfig name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// The VPC SC policy for project and location.
	// Possible values:
	//   "VPCSC_POLICY_UNSPECIFIED" - the VPS SC policy is not defined.
	//   "DENY" - repository will block the requests to the Upstreams for the Remote Repositories if the resource is in the perimeter.
	//   "ALLOW" - repository will allow the requests to the Upstreams for the Remote Repositories if the resource is in the perimeter.
	VpcscPolicy *string `json:"vpcscPolicy,omitempty"`
}

// ArtifactRegistryVPCSCConfigStatus defines the config connector machine state of ArtifactRegistryVPCSCConfig
type ArtifactRegistryVPCSCConfigStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ArtifactRegistryVPCSCConfig resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ArtifactRegistryVPCSCConfigObservedState `json:"observedState,omitempty"`
}

// ArtifactRegistryVPCSCConfigObservedState is the state of the ArtifactRegistryVPCSCConfig resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.devtools.artifactregistry.v1.VPCSCConfig
type ArtifactRegistryVPCSCConfigObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpartifactregistryvpcscconfig;gcpartifactregistryvpcscconfigs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ArtifactRegistryVPCSCConfig is the Schema for the ArtifactRegistryVPCSCConfig API
// +k8s:openapi-gen=true
type ArtifactRegistryVPCSCConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ArtifactRegistryVPCSCConfigSpec   `json:"spec,omitempty"`
	Status ArtifactRegistryVPCSCConfigStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ArtifactRegistryVPCSCConfigList contains a list of ArtifactRegistryVPCSCConfig
type ArtifactRegistryVPCSCConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ArtifactRegistryVPCSCConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ArtifactRegistryVPCSCConfig{}, &ArtifactRegistryVPCSCConfigList{})
}
