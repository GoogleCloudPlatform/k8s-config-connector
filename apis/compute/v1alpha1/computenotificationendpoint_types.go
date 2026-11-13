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

// API sources for ComputeNotificationEndpoint, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/compute/v1/compute.proto
// +kcc:source:service-docs=https://developers.google.com/compute/docs/reference/latest/
// +kcc:source:resource-docs=https://cloud.google.com/compute/docs/reference/rest/v1/regionNotificationEndpoints

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ComputeNotificationEndpointGVK = GroupVersion.WithKind("ComputeNotificationEndpoint")

// ComputeNotificationEndpointSpec defines the desired state of ComputeNotificationEndpoint
// +kcc:spec:proto=google.cloud.compute.v1.NotificationEndpoint
// +kcc:required-from-proto
type ComputeNotificationEndpointSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// Location represents the geographical location of the ComputeNotificationEndpoint. Specify a region name. Reference: GCP definition of regions/zones (https://cloud.google.com/compute/docs/regions-zones/)
	// +required
	Location string `json:"location"`

	// The ComputeNotificationEndpoint name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// An optional description of this resource. Provide this property when you create the resource.
	// +kcc:proto:field=google.cloud.compute.v1.NotificationEndpoint.description
	Description *string `json:"description,omitempty"`

	// Settings of the gRPC notification endpoint including the endpoint URL and the retry duration.
	// +kcc:proto:field=google.cloud.compute.v1.NotificationEndpoint.grpc_settings
	GrpcSettings *NotificationEndpointGrpcSettings `json:"grpcSettings,omitempty"`
}

// ComputeNotificationEndpointStatus defines the config connector machine state of ComputeNotificationEndpoint
type ComputeNotificationEndpointStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeNotificationEndpoint resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ComputeNotificationEndpointObservedState `json:"observedState,omitempty"`
}

// ComputeNotificationEndpointObservedState is the state of the ComputeNotificationEndpoint resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.NotificationEndpoint
type ComputeNotificationEndpointObservedState struct {
	// [Output Only] Creation timestamp in RFC3339 text format.
	// +kcc:proto:field=google.cloud.compute.v1.NotificationEndpoint.creation_timestamp
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// [Output Only] A unique identifier for this resource type. The server generates this identifier.
	// +kcc:proto:field=google.cloud.compute.v1.NotificationEndpoint.id
	ID *uint64 `json:"id,omitempty"`

	// [Output Only] Type of the resource. Always compute#notificationEndpoint for notification endpoints.
	// +kcc:proto:field=google.cloud.compute.v1.NotificationEndpoint.kind
	Kind *string `json:"kind,omitempty"`

	// [Output Only] URL of the region where the notification endpoint resides.
	// +kcc:proto:field=google.cloud.compute.v1.NotificationEndpoint.region
	Region *string `json:"region,omitempty"`

	// [Output Only] Server-defined URL for the resource.
	// +kcc:proto:field=google.cloud.compute.v1.NotificationEndpoint.self_link
	SelfLink *string `json:"selfLink,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputenotificationendpoint;gcpcomputenotificationendpoints
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeNotificationEndpoint is the Schema for the ComputeNotificationEndpoint API
// +k8s:openapi-gen=true
type ComputeNotificationEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeNotificationEndpointSpec   `json:"spec,omitempty"`
	Status ComputeNotificationEndpointStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeNotificationEndpointList contains a list of ComputeNotificationEndpoint
type ComputeNotificationEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeNotificationEndpoint `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeNotificationEndpoint{}, &ComputeNotificationEndpointList{})
}
