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

var VertexAIFeatureGroupGVK = GroupVersion.WithKind("VertexAIFeatureGroup")

// VertexAIFeatureGroupSpec defines the desired state of VertexAIFeatureGroup
// +kcc:spec:proto=google.cloud.aiplatform.v1beta1.FeatureGroup
type VertexAIFeatureGroupSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location *string `json:"location"`

	// The VertexAIFeatureGroup name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Description of the FeatureGroup.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.FeatureGroup.description
	Description *string `json:"description,omitempty"`

	// Optional. The labels with user-defined metadata to organize your FeatureGroup.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.FeatureGroup.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Indicates that that the source of this FeatureGroup is BigQuery.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.FeatureGroup.big_query
	BigQuery *FeatureGroup_BigQuery `json:"bigQuery,omitempty"`

	// Optional. A service account type for BigQuery data reading.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.FeatureGroup.service_agent_type
	ServiceAgentType *string `json:"serviceAgentType,omitempty"`
}

// VertexAIFeatureGroupStatus defines the config connector machine state of VertexAIFeatureGroup
type VertexAIFeatureGroupStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VertexAIFeatureGroup resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VertexAIFeatureGroupObservedState `json:"observedState,omitempty"`
}

// VertexAIFeatureGroupObservedState is the state of the VertexAIFeatureGroup resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.aiplatform.v1beta1.FeatureGroup
type VertexAIFeatureGroupObservedState struct {
	// Output only. Timestamp when this FeatureGroup was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.FeatureGroup.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this FeatureGroup was last updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.FeatureGroup.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The email of the service account used to read the BigQuery data.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.FeatureGroup.service_account_email
	ServiceAccountEmail *string `json:"serviceAccountEmail,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvertexaifeaturegroup;gcpvertexaifeaturegroups
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VertexAIFeatureGroup is the Schema for the VertexAIFeatureGroup API
// +k8s:openapi-gen=true
type VertexAIFeatureGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VertexAIFeatureGroupSpec   `json:"spec,omitempty"`
	Status VertexAIFeatureGroupStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VertexAIFeatureGroupList contains a list of VertexAIFeatureGroup
type VertexAIFeatureGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VertexAIFeatureGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VertexAIFeatureGroup{}, &VertexAIFeatureGroupList{})
}
