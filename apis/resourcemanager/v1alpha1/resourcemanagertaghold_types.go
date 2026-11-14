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

// API sources for ResourceManagerTagHold, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/resourcemanager/v3/tag_holds.proto
// +kcc:source:service-docs=https://cloud.google.com/resource-manager
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/resource-manager/docs/reference/rest/v3/tagValues.tagHolds

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ResourceManagerTagHoldGVK = GroupVersion.WithKind("ResourceManagerTagHold")

// ResourceManagerTagHoldSpec defines the desired state of ResourceManagerTagHold
// +kcc:spec:proto=google.cloud.resourcemanager.v3.TagHold
// +kcc:required-from-proto
type ResourceManagerTagHoldSpec struct {

	// The ResourceManagerTagHold name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. The name of the resource where the TagValue is being used. Must
	//  be less than 200 characters. E.g.
	//  `//compute.googleapis.com/compute/projects/myproject/regions/us-east-1/instanceGroupManagers/instance-group`
	// +kcc:proto:field=google.cloud.resourcemanager.v3.TagHold.holder
	// +required
	Holder *string `json:"holder,omitempty"`

	// Optional. An optional string representing the origin of this request. This
	//  field should include human-understandable information to distinguish
	//  origins from each other. Must be less than 200 characters. E.g.
	//  `migs-35678234`
	// +kcc:proto:field=google.cloud.resourcemanager.v3.TagHold.origin
	Origin *string `json:"origin,omitempty"`

	// Optional. A URL where an end user can learn more about removing this hold.
	//  E.g.
	//  `https://cloud.google.com/resource-manager/docs/tags/tags-creating-and-managing`
	// +kcc:proto:field=google.cloud.resourcemanager.v3.TagHold.help_link
	HelpLink *string `json:"helpLink,omitempty"`
}

// ResourceManagerTagHoldStatus defines the config connector machine state of ResourceManagerTagHold
type ResourceManagerTagHoldStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ResourceManagerTagHold resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ResourceManagerTagHoldObservedState `json:"observedState,omitempty"`
}

// ResourceManagerTagHoldObservedState is the state of the ResourceManagerTagHold resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.resourcemanager.v3.TagHold
type ResourceManagerTagHoldObservedState struct {
	// Output only. The time this TagHold was created.
	// +kcc:proto:field=google.cloud.resourcemanager.v3.TagHold.create_time
	CreateTime *string `json:"createTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpresourcemanagertaghold;gcpresourcemanagertagholds
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ResourceManagerTagHold is the Schema for the ResourceManagerTagHold API
// +k8s:openapi-gen=true
type ResourceManagerTagHold struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ResourceManagerTagHoldSpec   `json:"spec,omitempty"`
	Status ResourceManagerTagHoldStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ResourceManagerTagHoldList contains a list of ResourceManagerTagHold
type ResourceManagerTagHoldList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourceManagerTagHold `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ResourceManagerTagHold{}, &ResourceManagerTagHoldList{})
}
