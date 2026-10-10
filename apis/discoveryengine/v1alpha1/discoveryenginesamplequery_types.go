// Copyright 2025 Google LLC
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

// API sources for DiscoveryEngineSampleQuery, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/discoveryengine/v1beta/sample_query.proto
// +kcc:source:service-docs=https://cloud.google.com/generative-ai-app-builder/docs/
// +kcc:source:resource-docs=https://docs.cloud.google.com/generative-ai-app-builder/docs/reference/rest/v1beta/projects.locations.sampleQuerySets.sampleQueries

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DiscoveryEngineSampleQueryGVK = GroupVersion.WithKind("DiscoveryEngineSampleQuery")

// DiscoveryEngineSampleQuerySpec defines the desired state of DiscoveryEngineSampleQuery
// +kcc:spec:proto=google.cloud.discoveryengine.v1beta.SampleQuery
// +kcc:required-from-proto
type DiscoveryEngineSampleQuerySpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The DiscoveryEngineSampleQuery name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// The query entry.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SampleQuery.query_entry
	QueryEntry *SampleQuery_QueryEntry `json:"queryEntry,omitempty"`
}

// DiscoveryEngineSampleQueryStatus defines the config connector machine state of DiscoveryEngineSampleQuery
type DiscoveryEngineSampleQueryStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DiscoveryEngineSampleQuery resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DiscoveryEngineSampleQueryObservedState `json:"observedState,omitempty"`
}

// DiscoveryEngineSampleQueryObservedState is the state of the DiscoveryEngineSampleQuery resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.discoveryengine.v1beta.SampleQuery
type DiscoveryEngineSampleQueryObservedState struct {
	// Output only. Timestamp the
	//  [SampleQuery][google.cloud.discoveryengine.v1beta.SampleQuery] was created
	//  at.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SampleQuery.create_time
	CreateTime *string `json:"createTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdiscoveryenginesamplequery;gcpdiscoveryenginesamplequerys
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DiscoveryEngineSampleQuery is the Schema for the DiscoveryEngineSampleQuery API
// +k8s:openapi-gen=true
type DiscoveryEngineSampleQuery struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DiscoveryEngineSampleQuerySpec   `json:"spec,omitempty"`
	Status DiscoveryEngineSampleQueryStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DiscoveryEngineSampleQueryList contains a list of DiscoveryEngineSampleQuery
type DiscoveryEngineSampleQueryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DiscoveryEngineSampleQuery `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DiscoveryEngineSampleQuery{}, &DiscoveryEngineSampleQueryList{})
}
