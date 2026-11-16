// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ResourceRule defines criteria for what to sync.
// +kubebuilder:object:generate=true
type ResourceRule struct {
	// Group is the API group of the resource to be synchronized.
	Group string `json:"group"`
	// Version is the API version of the resource to be synchronized.
	Version string `json:"version"`
	// Kind is the Kind of the resource to be synchronized.
	Kind string `json:"kind"`
	// Namespaces is an optional list of namespaces to watch. If not provided, all namespaces are synchronized.
	// +optional
	Namespaces []string `json:"namespaces,omitempty"`
	// SyncFields is an optional list of fields to synchronize. If not provided, only the "status" field is synchronized.
	// Examples: "spec", "status", "spec.resourceID".
	// +optional
	// +kubebuilder:validation:items:Enum=spec;status;spec.resourceID
	// +kubebuilder:default={"status"}
	SyncFields []string `json:"syncFields,omitempty"`
}

// Mode defines the synchronization mode.
type Mode string

const (
	// ModePush means the syncer watches resources on the local cluster and pushes them to the remote cluster.
	ModePush Mode = "push"
	// ModePull means the syncer watches resources on the remote cluster and pulls them to the local cluster.
	ModePull Mode = "pull"
)

// RemoteConfig defines the remote cluster configuration.
// Exactly one remote cluster type must be set.
// +kubebuilder:object:generate=true
// +kubebuilder:validation:XValidation:rule="has(self.gkeCluster)",message="a remote cluster must be specified (gkeCluster)"
type RemoteConfig struct {
	// GKECluster references a remote GKE cluster. The controller authenticates
	// to it using its own Google identity (Workload Identity or Application
	// Default Credentials), so no credentials need to be stored in the cluster.
	// +optional
	GKECluster *GKECluster `json:"gkeCluster,omitempty"`
}

// GKEEndpoint selects which control plane endpoint of a GKE cluster to use.
type GKEEndpoint string

const (
	// GKEEndpointDefault uses the cluster's default control plane endpoint,
	// the same one `gcloud container clusters get-credentials` uses.
	GKEEndpointDefault GKEEndpoint = "Default"
	// GKEEndpointDNS uses the cluster's DNS-based control plane endpoint.
	// The DNS endpoint must be enabled on the cluster.
	GKEEndpointDNS GKEEndpoint = "DNS"
	// GKEEndpointPrivateIP uses the cluster's private control plane IP. The
	// controller must have network connectivity to the cluster's VPC.
	GKEEndpointPrivateIP GKEEndpoint = "PrivateIP"
)

// GKECluster identifies a GKE cluster.
// +kubebuilder:object:generate=true
type GKECluster struct {
	// Project is the ID of the Google Cloud project that hosts the cluster.
	// +kubebuilder:validation:MinLength=1
	Project string `json:"project"`
	// Location is the region or zone of the cluster, e.g. "us-central1".
	// +kubebuilder:validation:MinLength=1
	Location string `json:"location"`
	// Name is the name of the cluster.
	// +kubebuilder:validation:MinLength=1
	Name string `json:"name"`
	// Endpoint selects the control plane endpoint used to reach the cluster:
	// Default, DNS or PrivateIP.
	// +optional
	// +kubebuilder:validation:Enum=Default;DNS;PrivateIP
	// +kubebuilder:default=Default
	Endpoint GKEEndpoint `json:"endpoint,omitempty"`
}

// KRMSyncerSpec defines the desired state.
// +kubebuilder:object:generate=true
type KRMSyncerSpec struct {
	// Suspend tells the controller to suspend the sync operations.
	// +optional
	Suspend bool `json:"suspend,omitempty"`

	// Mode defines the sync mode: push or pull.
	// +optional
	// +kubebuilder:validation:Enum=push;pull
	// +kubebuilder:default=pull
	Mode Mode `json:"mode,omitempty"`

	// Remote defines the remote cluster for the sync.
	Remote *RemoteConfig `json:"remote"`

	// Rules defines which resources to watch and sync. If unset, sync all resources by default.
	Rules []ResourceRule `json:"rules"`
}

// KRMSyncerStatus defines the observed state.
// +kubebuilder:object:generate=true
type KRMSyncerStatus struct {
	// Conditions of the Syncer.
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// KRMSyncer is the Schema for the krmsyncers API
type KRMSyncer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KRMSyncerSpec   `json:"spec,omitempty"`
	Status KRMSyncerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KRMSyncerList contains a list of KRMSyncer
type KRMSyncerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KRMSyncer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KRMSyncer{}, &KRMSyncerList{})
}
