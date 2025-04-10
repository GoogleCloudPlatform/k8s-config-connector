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

package v1alpha1

import (
	kmsv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DocumentAIProcessorVersionGVK = GroupVersion.WithKind("DocumentAIProcessorVersion")

// DocumentAIProcessorVersionSpec defines the desired state of DocumentAIProcessorVersion
// +kcc:spec:proto=google.cloud.documentai.v1.ProcessorVersion
type DocumentAIProcessorVersionSpec struct {
	Parent `json:",inline"`

	// The DocumentAIProcessorVersion name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// If set, information about the eventual deprecation of this version.
	// +optional
	DeprecationInfo *ProcessorVersion_DeprecationInfo `json:"deprecationInfo,omitempty"`

	// The display name of the processor version.
	// +optional
	DisplayName *string `json:"displayName,omitempty"`

	// The KMS key name used for encryption.
	// +optional
	KMSKeyNameRef *refs.KMSCryptoKeyRef `json:"kmsKeyNameRef,omitempty"`

	// The KMS key version with which data is encrypted.
	// +optional
	KMSKeyVersionNameRef *kmsv1alpha1.KMSCryptoKeyVersionRef `json:"kmsKeyVersionNameRef,omitempty"`
}

type Parent struct {
	// +required
	ProcessorRef *ProcessorRef `json:"processorRef"`
}

// DocumentAIProcessorVersionStatus defines the config connector machine state of DocumentAIProcessorVersion
type DocumentAIProcessorVersionStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DocumentAI resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DocumentAIProcessorVersionObservedState `json:"observedState,omitempty"`
}

// DocumentAIProcessorVersionObservedState is the state of the DocumentAIProcessorVersion resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.documentai.v1.ProcessorVersion
type DocumentAIProcessorVersionObservedState struct {
	// Output only. The state of the processor version.
	State *string `json:"state,omitempty"`

	// The time the processor version was created.
	CreateTime *string `json:"create_time,omitempty"`

	// The most recently invoked evaluation for the processor version.
	LatestEvaluation *EvaluationReference `json:"latest_evaluation,omitempty"`

	// Output only. Denotes that this `ProcessorVersion` is managed by Google.
	GoogleManaged *bool `json:"google_managed,omitempty"`

	// Output only. The model type of this processor version.
	ModelType *string `json:"model_type,omitempty"`

	// Output only. Reserved for future use.
	SatisfiesPzs *bool `json:"satisfies_pzs,omitempty"`

	// Output only. Reserved for future use.
	SatisfiesPzi *bool `json:"satisfies_pzi,omitempty"`

	// Output only. Information about Generative AI model-based processor versions.
	GenAiModelInfo *ProcessorVersion_GenAiModelInfo `json:"gen_ai_model_info,omitempty"`

	// The schema of the processor version. Describes the output.
	DocumentSchema *DocumentSchema `json:"document_schema,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdocumentaiprocessorversion;gcpdocumentaiprocessorversions
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DocumentAIProcessorVersion is the Schema for the DocumentAIProcessorVersion API
// +k8s:openapi-gen=true
type DocumentAIProcessorVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DocumentAIProcessorVersionSpec   `json:"spec,omitempty"`
	Status DocumentAIProcessorVersionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DocumentAIProcessorVersionList contains a list of DocumentAIProcessorVersion
type DocumentAIProcessorVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DocumentAIProcessorVersion `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DocumentAIProcessorVersion{}, &DocumentAIProcessorVersionList{})
}
