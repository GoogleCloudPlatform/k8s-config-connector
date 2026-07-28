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

var VertexAIModelMonitorGVK = GroupVersion.WithKind("VertexAIModelMonitor")

// VertexAIModelMonitorSpec defines the desired state of VertexAIModelMonitor
type VertexAIModelMonitorSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location *string `json:"location"`

	// The VertexAIModelMonitor name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional default tabular model monitoring objective.
	TabularObjective *ModelMonitoringObjectiveSpec_TabularObjective `json:"tabularObjective,omitempty"`

	// The display name of the ModelMonitor.
	//  The name can be up to 128 characters long and can consist of any UTF-8.
	DisplayName *string `json:"displayName,omitempty"`

	// The entity that is subject to analysis.
	//  Currently only models in Vertex AI Model Registry are supported. If you
	//  want to analyze the model which is outside the Vertex AI, you could
	//  register a model in Vertex AI Model Registry using just a display name.
	ModelMonitoringTarget *ModelMonitor_ModelMonitoringTarget `json:"modelMonitoringTarget,omitempty"`

	// Optional training dataset used to train the model.
	//  It can serve as a reference dataset to identify changes in production.
	TrainingDataset *ModelMonitoringInput `json:"trainingDataset,omitempty"`

	// Optional default notification spec, it can be overridden in the
	//  ModelMonitoringJob notification spec.
	NotificationSpec *ModelMonitoringNotificationSpec `json:"notificationSpec,omitempty"`

	// Optional default monitoring metrics/logs export spec, it can be overridden
	//  in the ModelMonitoringJob output spec.
	//  If not specified, a default Google Cloud Storage bucket will be created
	//  under your project.
	OutputSpec *ModelMonitoringOutputSpec `json:"outputSpec,omitempty"`

	// Optional model explanation spec. It is used for feature attribution
	//  monitoring.
	ExplanationSpec *ExplanationSpec `json:"explanationSpec,omitempty"`

	// Monitoring Schema is to specify the model's features, prediction outputs
	//  and ground truth properties. It is used to extract pertinent data from the
	//  dataset and to process features based on their properties.
	//  Make sure that the schema aligns with your dataset, if it does not, we will
	//  be unable to extract data from the dataset.
	//  It is required for most models, but optional for Vertex AI AutoML Tables
	//  unless the schem information is not available.
	ModelMonitoringSchema *ModelMonitoringSchema `json:"modelMonitoringSchema,omitempty"`

	// Customer-managed encryption key spec for a ModelMonitor. If
	//  set, this ModelMonitor and all sub-resources of this
	//  ModelMonitor will be secured by this key.
	EncryptionSpec *EncryptionSpec `json:"encryptionSpec,omitempty"`
}

// VertexAIModelMonitorStatus defines the config connector machine state of VertexAIModelMonitor
type VertexAIModelMonitorStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VertexAIModelMonitor resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VertexAIModelMonitorObservedState `json:"observedState,omitempty"`
}

// VertexAIModelMonitorObservedState is the state of the VertexAIModelMonitor resource as most recently observed in GCP.
type VertexAIModelMonitorObservedState struct {
	// Output only. Timestamp when this ModelMonitor was created.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this ModelMonitor was updated most recently.
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Reserved for future use.
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. Reserved for future use.
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvertexaimodelmonitor;gcpvertexaimodelmonitors
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VertexAIModelMonitor is the Schema for the VertexAIModelMonitor API
// +k8s:openapi-gen=true
type VertexAIModelMonitor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VertexAIModelMonitorSpec   `json:"spec,omitempty"`
	Status VertexAIModelMonitorStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VertexAIModelMonitorList contains a list of VertexAIModelMonitor
type VertexAIModelMonitorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VertexAIModelMonitor `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VertexAIModelMonitor{}, &VertexAIModelMonitorList{})
}
