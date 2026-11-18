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

// API sources for ComputeInstanceGroupManagerResizeRequest, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/compute/v1/compute.proto
// +kcc:source:service-docs=https://developers.google.com/compute/docs/reference/latest/
// +kcc:source:resource-docs=https://cloud.google.com/compute/docs/reference/rest/v1/instanceGroupManagerResizeRequests

package v1alpha1

import (
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ComputeInstanceGroupManagerResizeRequestGVK = GroupVersion.WithKind("ComputeInstanceGroupManagerResizeRequest")

// ComputeInstanceGroupManagerResizeRequestSpec defines the desired state of ComputeInstanceGroupManagerResizeRequest
// +kcc:spec:proto=google.cloud.compute.v1.InstanceGroupManagerResizeRequest
// +kcc:required-from-proto
type ComputeInstanceGroupManagerResizeRequestSpec struct {
	// The ComputeInstanceGroupManager this ComputeInstanceGroupManagerResizeRequest is created under.
	// +required
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="parentRef is immutable"
	ParentRef *computev1beta1.ComputeInstanceGroupManagerRef `json:"parentRef"`

	// The ComputeInstanceGroupManagerResizeRequest name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// An optional description of this resource.
	// +kcc:proto:field=google.cloud.compute.v1.InstanceGroupManagerResizeRequest.description
	Description *string `json:"description,omitempty"`

	// Requested run duration for instances that will be created by this request. At the end of the run duration instance will be deleted.
	// +kcc:proto:field=google.cloud.compute.v1.InstanceGroupManagerResizeRequest.requested_run_duration
	RequestedRunDuration *Duration `json:"requestedRunDuration,omitempty"`

	// The number of instances to be created by this resize request. The group's target size will be increased by this number. This field cannot be used together with 'instances'.
	// +kcc:proto:field=google.cloud.compute.v1.InstanceGroupManagerResizeRequest.resize_by
	ResizeBy *int32 `json:"resizeBy,omitempty"`
}

// ComputeInstanceGroupManagerResizeRequestStatus defines the config connector machine state of ComputeInstanceGroupManagerResizeRequest
type ComputeInstanceGroupManagerResizeRequestStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeInstanceGroupManagerResizeRequest resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ComputeInstanceGroupManagerResizeRequestObservedState `json:"observedState,omitempty"`
}

// ComputeInstanceGroupManagerResizeRequestObservedState is the state of the ComputeInstanceGroupManagerResizeRequest resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.InstanceGroupManagerResizeRequest
type ComputeInstanceGroupManagerResizeRequestObservedState struct {
	// [Output Only] The creation timestamp for this resize request in RFC3339 text format.
	// +kcc:proto:field=google.cloud.compute.v1.InstanceGroupManagerResizeRequest.creation_timestamp
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// [Output Only] A unique identifier for this resource type. The server generates this identifier.
	// +kcc:proto:field=google.cloud.compute.v1.InstanceGroupManagerResizeRequest.id
	ID *uint64 `json:"id,omitempty"`

	// [Output Only] The resource type, which is always compute#instanceGroupManagerResizeRequest for resize requests.
	// +kcc:proto:field=google.cloud.compute.v1.InstanceGroupManagerResizeRequest.kind
	Kind *string `json:"kind,omitempty"`

	// [Output Only] The URL for this resize request. The server defines this URL.
	// +kcc:proto:field=google.cloud.compute.v1.InstanceGroupManagerResizeRequest.self_link
	SelfLink *string `json:"selfLink,omitempty"`

	// [Output Only] Server-defined URL for this resource with the resource id.
	// +kcc:proto:field=google.cloud.compute.v1.InstanceGroupManagerResizeRequest.self_link_with_id
	SelfLinkWithID *string `json:"selfLinkWithID,omitempty"`

	// [Output only] Current state of the request.
	// Check the State enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.InstanceGroupManagerResizeRequest.state
	State *string `json:"state,omitempty"`

	// [Output only] Status of the request.
	// +kcc:proto:field=google.cloud.compute.v1.InstanceGroupManagerResizeRequest.status
	Status *InstanceGroupManagerResizeRequestStatus `json:"status,omitempty"`

	// [Output Only] The URL of a zone where the resize request is located. Populated only for zonal resize requests.
	// +kcc:proto:field=google.cloud.compute.v1.InstanceGroupManagerResizeRequest.zone
	Zone *string `json:"zone,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputeinstancegroupmanagerresizerequest;gcpcomputeinstancegroupmanagerresizerequests
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeInstanceGroupManagerResizeRequest is the Schema for the ComputeInstanceGroupManagerResizeRequest API
// +k8s:openapi-gen=true
type ComputeInstanceGroupManagerResizeRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeInstanceGroupManagerResizeRequestSpec   `json:"spec,omitempty"`
	Status ComputeInstanceGroupManagerResizeRequestStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeInstanceGroupManagerResizeRequestList contains a list of ComputeInstanceGroupManagerResizeRequest
type ComputeInstanceGroupManagerResizeRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeInstanceGroupManagerResizeRequest `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeInstanceGroupManagerResizeRequest{}, &ComputeInstanceGroupManagerResizeRequestList{})
}
