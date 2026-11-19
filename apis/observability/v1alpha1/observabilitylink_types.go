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

// API sources for ObservabilityLink, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/9f84face29bd2b2333f1927a24512b9ff2d195bc/google/cloud/observability/v1/observability.proto
// +kcc:source:service-docs=https://cloud.google.com/stackdriver/docs/
// +kcc:source:resource-docs=https://docs.cloud.google.com/stackdriver/docs/reference/observability/api/rest/v1/projects.locations.buckets.datasets.links

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ObservabilityLinkGVK = GroupVersion.WithKind("ObservabilityLink")

// ObservabilityLinkSpec defines the desired state of ObservabilityLink
// +kcc:spec:proto=google.cloud.observability.v1.Link
// +kcc:required-from-proto
type ObservabilityLinkSpec struct {
	// The ObservabilityDataset this ObservabilityLink is created under.
	// +required
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="parentRef is immutable"
	ParentRef *ObservabilityDatasetRef `json:"parentRef"`

	// The ObservabilityLink name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. A user friendly display name.
	// +kcc:proto:field=google.cloud.observability.v1.Link.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Description of the link.
	// +kcc:proto:field=google.cloud.observability.v1.Link.description
	Description *string `json:"description,omitempty"`
}

// ObservabilityLinkStatus defines the config connector machine state of ObservabilityLink
type ObservabilityLinkStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ObservabilityLink resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ObservabilityLinkObservedState `json:"observedState,omitempty"`
}

// ObservabilityLinkObservedState is the state of the ObservabilityLink resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.observability.v1.Link
type ObservabilityLinkObservedState struct {
	// Output only. Create timestamp.
	// +kcc:proto:field=google.cloud.observability.v1.Link.create_time
	CreateTime *string `json:"createTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpobservabilitylink;gcpobservabilitylinks
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ObservabilityLink is the Schema for the ObservabilityLink API
// +k8s:openapi-gen=true
type ObservabilityLink struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ObservabilityLinkSpec   `json:"spec,omitempty"`
	Status ObservabilityLinkStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ObservabilityLinkList contains a list of ObservabilityLink
type ObservabilityLinkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ObservabilityLink `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ObservabilityLink{}, &ObservabilityLinkList{})
}
