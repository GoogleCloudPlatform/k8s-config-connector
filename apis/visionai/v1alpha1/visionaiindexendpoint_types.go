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

// API sources for VisionAIIndexEndpoint, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/visionai/v1/warehouse.proto
// +kcc:source:service-docs=https://cloud.google.com/vision-ai/docs
// +kcc:source:resource-docs=https://docs.cloud.google.com/vision-ai/docs/reference/rest/v1/projects.locations.indexEndpoints

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var VisionAIIndexEndpointGVK = GroupVersion.WithKind("VisionAIIndexEndpoint")

// VisionAIIndexEndpointSpec defines the desired state of VisionAIIndexEndpoint
// +kcc:spec:proto=google.cloud.visionai.v1.IndexEndpoint
// +kcc:required-from-proto
type VisionAIIndexEndpointSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The VisionAIIndexEndpoint name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. Display name of the IndexEndpoint. Can be up to 32 characters
	//  long.
	// +kcc:proto:field=google.cloud.visionai.v1.IndexEndpoint.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Description of the IndexEndpoint. Can be up to 25000 characters
	//  long.
	// +kcc:proto:field=google.cloud.visionai.v1.IndexEndpoint.description
	Description *string `json:"description,omitempty"`

	// Optional. The labels applied to a resource must meet the following
	//  requirements:
	//
	//  * Each resource can have multiple labels, up to a maximum of 64.
	//  * Each label must be a key-value pair.
	//  * Keys have a minimum length of 1 character and a maximum length of 63
	//    characters and cannot be empty. Values can be empty and have a maximum
	//    length of 63 characters.
	//  * Keys and values can contain only lowercase letters, numeric characters,
	//    underscores, and dashes. All characters must use UTF-8 encoding, and
	//    international characters are allowed.
	//  * The key portion of a label must be unique. However, you can use the same
	//    key with multiple resources.
	//  * Keys must start with a lowercase letter or international character.
	//
	//  See [Google Cloud
	//  Document](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements)
	//  for more details.
	// +kcc:proto:field=google.cloud.visionai.v1.IndexEndpoint.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// VisionAIIndexEndpointStatus defines the config connector machine state of VisionAIIndexEndpoint
type VisionAIIndexEndpointStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VisionAIIndexEndpoint resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VisionAIIndexEndpointObservedState `json:"observedState,omitempty"`
}

// VisionAIIndexEndpointObservedState is the state of the VisionAIIndexEndpoint resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.visionai.v1.IndexEndpoint
type VisionAIIndexEndpointObservedState struct {
	// Output only. The Index deployed in this IndexEndpoint.
	// +kcc:proto:field=google.cloud.visionai.v1.IndexEndpoint.deployed_index
	DeployedIndex *DeployedIndexObservedState `json:"deployedIndex,omitempty"`

	// Output only. IndexEndpoint state.
	// +kcc:proto:field=google.cloud.visionai.v1.IndexEndpoint.state
	State *string `json:"state,omitempty"`

	// Output only. Create timestamp.
	// +kcc:proto:field=google.cloud.visionai.v1.IndexEndpoint.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Update timestamp.
	// +kcc:proto:field=google.cloud.visionai.v1.IndexEndpoint.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. This boolean field is only set for projects that have Physical
	//  Zone Separation enabled via an Org Policy constraint. It is set to true
	//  when the index endpoint is a valid zone separated index endpoint and false
	//  if it isn't.
	// +kcc:proto:field=google.cloud.visionai.v1.IndexEndpoint.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. This boolean field is only set for projects that have Physical
	//  Zone Isolation enabled via an Org Policy constraint. It is set to true when
	//  the index endpoint is a valid zone isolated index endpoint and false if it
	//  isn't.
	// +kcc:proto:field=google.cloud.visionai.v1.IndexEndpoint.satisfies_pzi
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvisionaiindexendpoint;gcpvisionaiindexendpoints
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VisionAIIndexEndpoint is the Schema for the VisionAIIndexEndpoint API
// +k8s:openapi-gen=true
type VisionAIIndexEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VisionAIIndexEndpointSpec   `json:"spec,omitempty"`
	Status VisionAIIndexEndpointStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VisionAIIndexEndpointList contains a list of VisionAIIndexEndpoint
type VisionAIIndexEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VisionAIIndexEndpoint `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VisionAIIndexEndpoint{}, &VisionAIIndexEndpointList{})
}
