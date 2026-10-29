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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AppEngineServiceSplitTrafficGVK = GroupVersion.WithKind("AppEngineServiceSplitTraffic")

// +kcc:proto=google.appengine.v1.TrafficSplit
type AppEngineServiceSplitTrafficSplit struct {
	// Mapping from version IDs within the service to fractional (0.000, 1] allocations of traffic for that version. Each version can be specified only once, but some versions in the service may not have any traffic allocation. Services that have traffic allocated cannot be deleted until either the service is deleted or their traffic allocation is removed. Allocations must sum to 1. Up to two decimal place precision is supported for IP-based splits and up to three decimal places is supported for cookie-based splits.
	// +required
	Allocations map[string]string `json:"allocations"`

	// Mechanism used to determine which version a request is sent to. The traffic selection algorithm will be stable for either type until allocations are changed. Possible values: ["UNSPECIFIED", "COOKIE", "IP", "RANDOM"].
	// +optional
	// +kcc:proto:field=google.appengine.v1.TrafficSplit.shard_by
	ShardBy *string `json:"shardBy,omitempty"`
}

// +kcc:spec:proto=google.appengine.v1.Service
type AppEngineServiceSplitTrafficSpec struct {
	// If set to true traffic will be migrated to this version.
	// +optional
	MigrateTraffic *bool `json:"migrateTraffic,omitempty"`

	// Immutable.
	// +optional
	Project *string `json:"project,omitempty"`

	// Immutable. Optional. The service of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// Mapping that defines fractional HTTP traffic diversion to different versions within the service.
	// +required
	// +kcc:proto:field=google.appengine.v1.Service.split
	Split *AppEngineServiceSplitTrafficSplit `json:"split"`
}

type AppEngineServiceSplitTrafficStatus struct {
	/* Conditions represent the latest available observation of the resource's current state. */
	// +optional
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AppEngineServiceSplitTraffic resource in GCP.
	// +optional
	ExternalRef *string `json:"externalRef,omitempty"`
}

// AppEngineServiceSplitTrafficObservedState is the state of the AppEngineServiceSplitTraffic resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.appengine.v1.Service
type AppEngineServiceSplitTrafficObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpappengineservicesplittraffic;gcpappengineservicesplittraffics
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AppEngineServiceSplitTraffic is the Schema for the appengine API
// +k8s:openapi-gen=true
type AppEngineServiceSplitTraffic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AppEngineServiceSplitTrafficSpec   `json:"spec,omitempty"`
	Status AppEngineServiceSplitTrafficStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AppEngineServiceSplitTrafficList contains a list of AppEngineServiceSplitTraffic
type AppEngineServiceSplitTrafficList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppEngineServiceSplitTraffic `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AppEngineServiceSplitTraffic{}, &AppEngineServiceSplitTrafficList{})
}
