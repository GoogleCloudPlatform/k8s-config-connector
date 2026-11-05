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

// API sources for NetworkConnectivityDestination, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/cloud/networkconnectivity/v1/data_transfer.proto
// +kcc:source:service-docs=https://cloud.google.com/network-connectivity/docs/reference/networkconnectivity/rest
// +kcc:source:resource-docs=https://docs.cloud.google.com/network-connectivity/docs/reference/networkconnectivity/rest/v1/projects.locations.multicloudDataTransferConfigs.destinations

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var NetworkConnectivityDestinationGVK = GroupVersion.WithKind("NetworkConnectivityDestination")

// NetworkConnectivityDestinationSpec defines the desired state of NetworkConnectivityDestination
// +kcc:spec:proto=google.cloud.networkconnectivity.v1.Destination
// +kcc:required-from-proto
type NetworkConnectivityDestinationSpec struct {
	// The NetworkConnectivityMulticloudDataTransferConfig this NetworkConnectivityDestination is created under.
	// +required
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="parentRef is immutable"
	ParentRef *NetworkConnectivityMulticloudDataTransferConfigRef `json:"parentRef"`

	// The NetworkConnectivityDestination name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. A description of this resource.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Destination.description
	Description *string `json:"description,omitempty"`

	// Required. Immutable. The IP prefix that represents your workload on another
	//  CSP.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Destination.ip_prefix
	// +required
	IPPrefix *string `json:"ipPrefix,omitempty"`

	// Required. Unordered list. The list of `DestinationEndpoint` resources
	//  configured for the IP prefix.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Destination.endpoints
	// +required
	Endpoints []Destination_DestinationEndpoint `json:"endpoints,omitempty"`
}

// NetworkConnectivityDestinationStatus defines the config connector machine state of NetworkConnectivityDestination
type NetworkConnectivityDestinationStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NetworkConnectivityDestination resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *NetworkConnectivityDestinationObservedState `json:"observedState,omitempty"`
}

// NetworkConnectivityDestinationObservedState is the state of the NetworkConnectivityDestination resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.networkconnectivity.v1.Destination
type NetworkConnectivityDestinationObservedState struct {
	// The etag is computed by the server, and might be sent with update and
	//  delete requests so that the client has an up-to-date value before
	//  proceeding.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Destination.etag
	Etag *string `json:"etag,omitempty"`

	// Output only. Time when the `Destination` resource was created.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Destination.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time when the `Destination` resource was updated.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Destination.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Required. Unordered list. The list of `DestinationEndpoint` resources
	//  configured for the IP prefix.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Destination.endpoints
	Endpoints []Destination_DestinationEndpointObservedState `json:"endpoints,omitempty"`

	// Output only. The timeline of the expected `Destination` states or the
	//  current rest state. If a state change is expected, the value is `ADDING`,
	//  `DELETING` or `SUSPENDING`, depending on the action specified.
	//
	//  Example:
	//  "state_timeline": {
	//    "states": [
	//      {
	//        // The time when the `Destination` resource will be activated.
	//        "effectiveTime": "2024-12-01T08:00:00Z",
	//        "state": "ADDING"
	//      },
	//      {
	//        // The time when the `Destination` resource will be suspended.
	//        "effectiveTime": "2024-12-01T20:00:00Z",
	//        "state": "SUSPENDING"
	//      }
	//    ]
	//  }
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Destination.state_timeline
	StateTimeline *StateTimelineObservedState `json:"stateTimeline,omitempty"`

	// Output only. The Google-generated unique ID for the `Destination` resource.
	//  This value is unique across all `Destination` resources.
	//  If a resource is deleted and another with the same name is
	//  created, the new resource is assigned a different and unique ID.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Destination.uid
	Uid *string `json:"uid,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetworkconnectivitydestination;gcpnetworkconnectivitydestinations
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetworkConnectivityDestination is the Schema for the NetworkConnectivityDestination API
// +k8s:openapi-gen=true
type NetworkConnectivityDestination struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NetworkConnectivityDestinationSpec   `json:"spec,omitempty"`
	Status NetworkConnectivityDestinationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NetworkConnectivityDestinationList contains a list of NetworkConnectivityDestination
type NetworkConnectivityDestinationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkConnectivityDestination `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkConnectivityDestination{}, &NetworkConnectivityDestinationList{})
}
