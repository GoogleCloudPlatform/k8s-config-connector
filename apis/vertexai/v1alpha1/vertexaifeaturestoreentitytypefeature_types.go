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

var VertexAIFeaturestoreEntityTypeFeatureGVK = GroupVersion.WithKind("VertexAIFeaturestoreEntityTypeFeature")

// +kcc:spec:proto=google.cloud.aiplatform.v1beta1.Feature
type VertexAIFeaturestoreEntityTypeFeatureSpec struct {
	// Description of the feature.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Feature.description
	Description *string `json:"description,omitempty"`

	// Immutable. The name of the Featurestore to use, in the format projects/{project}/locations/{location}/featurestores/{featurestore}/entityTypes/{entitytype}.
	// +required
	Entitytype string `json:"entitytype"`

	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	ResourceID *string `json:"resourceID,omitempty"`

	// Immutable. Type of Feature value. Immutable. https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.featurestores.entityTypes.features#ValueType.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Feature.value_type
	// +required
	ValueType string `json:"valueType"`
}

// +kcc:status:proto=google.cloud.aiplatform.v1beta1.Feature
type VertexAIFeaturestoreEntityTypeFeatureStatus struct {
	// Conditions represent the latest available observation of the resource's current state.
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// The timestamp of when the entity type was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Feature.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Used to perform consistent read-modify-write updates.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Feature.etag
	Etag *string `json:"etag,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// The region of the feature.
	Region *string `json:"region,omitempty"`

	// The timestamp when the entity type was most recently updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Feature.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvertexaifeaturestoreentitytypefeature;gcpvertexaifeaturestoreentitytypefeatures
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VertexAIFeaturestoreEntityTypeFeature is the Schema for the VertexAIFeaturestoreEntityTypeFeature API
// +k8s:openapi-gen=true
type VertexAIFeaturestoreEntityTypeFeature struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VertexAIFeaturestoreEntityTypeFeatureSpec   `json:"spec,omitempty"`
	Status VertexAIFeaturestoreEntityTypeFeatureStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VertexAIFeaturestoreEntityTypeFeatureList contains a list of VertexAIFeaturestoreEntityTypeFeature
type VertexAIFeaturestoreEntityTypeFeatureList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VertexAIFeaturestoreEntityTypeFeature `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VertexAIFeaturestoreEntityTypeFeature{}, &VertexAIFeaturestoreEntityTypeFeatureList{})
}
