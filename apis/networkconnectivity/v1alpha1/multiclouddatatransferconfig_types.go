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

var NetworkConnectivityMulticloudDataTransferConfigGVK = GroupVersion.WithKind("NetworkConnectivityMulticloudDataTransferConfig")

// NetworkConnectivityMulticloudDataTransferConfigSpec defines the desired state of NetworkConnectivityMulticloudDataTransferConfig
// +kcc:spec:proto=mockgcp.cloud.networkconnectivity.v1.MulticloudDataTransferConfig
type NetworkConnectivityMulticloudDataTransferConfigSpec struct {
	// The project that this resource belongs to.
	// +kubebuilder:validation:Required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +kubebuilder:validation:Required
	Location string `json:"location"`

	// The NetworkConnectivityMulticloudDataTransferConfig name. If not given, the metadata.name will be used.
	// +kubebuilder:validation:Optional
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. A description of this resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty"`

	// Optional. User-defined labels.
	// +kubebuilder:validation:Optional
	Labels map[string]string `json:"labels,omitempty"`
}

// NetworkConnectivityMulticloudDataTransferConfigStatus defines the config connector machine state of NetworkConnectivityMulticloudDataTransferConfig
type NetworkConnectivityMulticloudDataTransferConfigStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NetworkConnectivityMulticloudDataTransferConfig resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *NetworkConnectivityMulticloudDataTransferConfigObservedState `json:"observedState,omitempty"`
}

// NetworkConnectivityMulticloudDataTransferConfigObservedState is the state of the NetworkConnectivityMulticloudDataTransferConfig resource as most recently observed in GCP.
// +kcc:observedstate:proto=mockgcp.cloud.networkconnectivity.v1.MulticloudDataTransferConfig
type NetworkConnectivityMulticloudDataTransferConfigObservedState struct {
	// Output only. Time when the MulticloudDataTransferConfig was created.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The number of Destination resources in use with the MulticloudDataTransferConfig resource.
	DestinationsActiveCount *int32 `json:"destinationsActiveCount,omitempty"`

	// Output only. The number of Destination resources configured for the MulticloudDataTransferConfig resource.
	DestinationsCount *int32 `json:"destinationsCount,omitempty"`

	// Output only. The etag is computed by the server, and might be sent with update and delete requests.
	Etag *string `json:"etag,omitempty"`

	// Output only. The Google-generated unique ID for the MulticloudDataTransferConfig resource.
	Uid *string `json:"uid,omitempty"`

	// Output only. Time when the MulticloudDataTransferConfig was updated.
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetworkconnectivitymulticlouddatatransferconfig;gcpnetworkconnectivitymulticlouddatatransferconfigs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetworkConnectivityMulticloudDataTransferConfig is the Schema for the NetworkConnectivityMulticloudDataTransferConfig API
// +k8s:openapi-gen=true
type NetworkConnectivityMulticloudDataTransferConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NetworkConnectivityMulticloudDataTransferConfigSpec   `json:"spec,omitempty"`
	Status NetworkConnectivityMulticloudDataTransferConfigStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NetworkConnectivityMulticloudDataTransferConfigList contains a list of NetworkConnectivityMulticloudDataTransferConfig
type NetworkConnectivityMulticloudDataTransferConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkConnectivityMulticloudDataTransferConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkConnectivityMulticloudDataTransferConfig{}, &NetworkConnectivityMulticloudDataTransferConfigList{})
}
