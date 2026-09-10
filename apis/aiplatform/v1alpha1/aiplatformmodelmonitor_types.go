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

var AIPlatformModelMonitorGVK = GroupVersion.WithKind("AIPlatformModelMonitor")

// AIPlatformModelMonitorSpec defines the desired state of AIPlatformModelMonitor
// +kcc:spec:proto=google.cloud.aiplatform.v1beta1.ModelMonitor
type AIPlatformModelMonitorSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location *string `json:"location"`

	// The AIPlatformModelMonitor name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// The display name of the ModelMonitor.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitor.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The entity that is subject to analysis.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitor.model_monitoring_target
	ModelMonitoringTarget *ModelMonitor_ModelMonitoringTarget `json:"modelMonitoringTarget,omitempty"`

	// Optional default tabular model monitoring objective.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitor.tabular_objective
	TabularObjective *ModelMonitoringObjectiveSpec_TabularObjective `json:"tabularObjective,omitempty"`

	// Optional training dataset used to train the model.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitor.training_dataset
	TrainingDataset *ModelMonitoringInput `json:"trainingDataset,omitempty"`

	// Optional default notification spec, it can be overridden in the ModelMonitoringJob notification spec.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitor.notification_spec
	NotificationSpec *ModelMonitoringNotificationSpec `json:"notificationSpec,omitempty"`

	// Optional default monitoring metrics/logs export spec, it can be overridden in the ModelMonitoringJob output spec.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitor.output_spec
	OutputSpec *ModelMonitoringOutputSpec `json:"outputSpec,omitempty"`

	// Monitoring Schema is to specify the model's features, prediction outputs and ground truth properties.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitor.model_monitoring_schema
	ModelMonitoringSchema *ModelMonitoringSchema `json:"modelMonitoringSchema,omitempty"`

	// Optional. Describes the virtual explanation spec that can be overridden in the ModelMonitoringJob.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitor.explanation_spec
	ExplanationSpec *ExplanationSpec `json:"explanationSpec,omitempty"`

	// Optional. Customer-managed encryption key spec for a ModelMonitor. If this is set, then all resources created by the ModelMonitor will be encrypted with the provided encryption key.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitor.encryption_spec
	EncryptionSpec *EncryptionSpec `json:"encryptionSpec,omitempty"`
}

// AIPlatformModelMonitorStatus defines the config connector machine state of AIPlatformModelMonitor
type AIPlatformModelMonitorStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AIPlatformModelMonitor resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AIPlatformModelMonitorObservedState `json:"observedState,omitempty"`
}

// AIPlatformModelMonitorObservedState is the state of the AIPlatformModelMonitor resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.aiplatform.v1beta1.ModelMonitor
type AIPlatformModelMonitorObservedState struct {
	// Output only. Timestamp when this ModelMonitor was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitor.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this ModelMonitor was updated most recently.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitor.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitor.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitor.satisfies_pzi
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpaiplatformmodelmonitor;gcpaiplatformmodelmonitors
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AIPlatformModelMonitor is the Schema for the AIPlatformModelMonitor API
// +k8s:openapi-gen=true
type AIPlatformModelMonitor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AIPlatformModelMonitorSpec   `json:"spec,omitempty"`
	Status AIPlatformModelMonitorStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AIPlatformModelMonitorList contains a list of AIPlatformModelMonitor
type AIPlatformModelMonitorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AIPlatformModelMonitor `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AIPlatformModelMonitor{}, &AIPlatformModelMonitorList{})
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.ModelMonitoringObjectiveSpec.TabularObjective
type ModelMonitoringObjectiveSpec_TabularObjective struct {
	// Input feature distribution drift monitoring spec.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringObjectiveSpec.TabularObjective.feature_drift_spec
	FeatureDriftSpec *ModelMonitoringObjectiveSpec_DataDriftSpec `json:"featureDriftSpec,omitempty"`

	// Prediction output distribution drift monitoring spec.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringObjectiveSpec.TabularObjective.prediction_output_drift_spec
	PredictionOutputDriftSpec *ModelMonitoringObjectiveSpec_DataDriftSpec `json:"predictionOutputDriftSpec,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.ModelMonitoringOutputSpec
type ModelMonitoringOutputSpec struct {
	// Google Cloud Storage base folder path for metrics, error logs, etc.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitoringOutputSpec.gcs_base_directory
	GCSBaseDirectory *ModelMonitorGcsDestination `json:"gcsBaseDirectory,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.GcsDestination
type ModelMonitorGcsDestination struct {
	// Required. Google Cloud Storage URI to write xml files to.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.GcsDestination.output_uri_prefix
	OutputUriPrefix *string `json:"outputUriPrefix,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1beta1.ModelMonitor.ModelMonitoringTarget.VertexModelSource
type ModelMonitor_ModelMonitoringTarget_VertexModelSource struct {
	// Model resource name. Format:
	//  projects/{project}/locations/{location}/models/{model}.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitor.ModelMonitoringTarget.VertexModelSource.model
	ModelRef *AIPlatformModelRef `json:"modelRef,omitempty"`

	// Model version id.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.ModelMonitor.ModelMonitoringTarget.VertexModelSource.model_version_id
	ModelVersionID *string `json:"modelVersionID,omitempty"`
}
