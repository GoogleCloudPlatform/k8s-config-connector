// Copyright 2024 Google LLC
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
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	commonv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/common/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DocumentAIProcessorGVK = GroupVersion.WithKind("DocumentAIProcessor")

// DocumentAIProcessorSpec defines the desired state of DocumentAIProcessor
// +kcc:spec:proto=google.cloud.documentai.v1.Processor
type DocumentAIProcessorSpec struct {
	commonv1alpha1.CommonSpec `json:",inline"`

	// +required
	Location string `json:"location"`

	// The DocumentAIProcessor name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// The processor type, such as: `OCR_PROCESSOR`, `INVOICE_PROCESSOR`.
	//  To get a list of processor types, see
	//  [FetchProcessorTypes][google.cloud.documentai.v1.DocumentProcessorService.FetchProcessorTypes].
	// +kcc:proto:field=google.cloud.documentai.v1.Processor.type
	Type *string `json:"type,omitempty"`

	// The display name of the processor.
	// +kcc:proto:field=google.cloud.documentai.v1.Processor.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The default processor version.
	// +kcc:proto:field=google.cloud.documentai.v1.Processor.default_processor_version
	// NOTYET
	// DefaultProcessorVersion *string `json:"defaultProcessorVersion,omitempty"`

	// The [KMS key](https://cloud.google.com/security-key-management) used for
	//  encryption and decryption in CMEK scenarios.
	// +kcc:proto:field=google.cloud.documentai.v1.Processor.kms_key_name
	KMSKeyRef *kmsv1beta1.KMSKeyRef_OneOf `json:"kmsKeyRef,omitempty"`
}

// DocumentAIProcessorStatus defines the config connector machine state of DocumentAIProcessor
type DocumentAIProcessorStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DocumentAIProcessor resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DocumentAIProcessorObservedState `json:"observedState,omitempty"`
}

// DocumentAIProcessorSpec defines the desired state of DocumentAIProcessor
// +kcc:proto=google.cloud.documentai.v1.Processor
// DocumentAIProcessorObservedState is the state of the DocumentAIProcessor resource as most recently observed in GCP.
type DocumentAIProcessorObservedState struct {
	// The time the processor was created.
	// +kcc:proto:field=google.cloud.documentai.v1.Processor.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Immutable. The resource name of the processor.
	//  Format: `projects/{project}/locations/{location}/processors/{processor}`
	// +kcc:proto:field=google.cloud.documentai.v1.Processor.name
	Name *string `json:"name,omitempty"`

	// Output only. The state of the processor.
	// +kcc:proto:field=google.cloud.documentai.v1.Processor.state
	State *string `json:"state,omitempty"`

	// Output only. The processor version aliases.
	// +kcc:proto:field=google.cloud.documentai.v1.Processor.processor_version_aliases
	ProcessorVersionAliases []ProcessorVersionAlias `json:"processorVersionAliases,omitempty"`

	// Output only. Immutable. The http endpoint that can be called to invoke
	//  processing.
	// +kcc:proto:field=google.cloud.documentai.v1.Processor.process_endpoint
	ProcessEndpoint *string `json:"processEndpoint,omitempty"`

	// The default processor version.
	// +kcc:proto:field=google.cloud.documentai.v1.Processor.default_processor_version
	DefaultProcessorVersion *string `json:"defaultProcessorVersion,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.documentai.v1.Processor.satisfies_pzs
	// NOTYET
	// SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.documentai.v1.Processor.satisfies_pzi
	// NOTYET
	// SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdocumentaiprocessor;gcpdocumentaiprocessors
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DocumentAIProcessor is the Schema for the DocumentAIProcessor API
// +k8s:openapi-gen=true
type DocumentAIProcessor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DocumentAIProcessorSpec   `json:"spec,omitempty"`
	Status DocumentAIProcessorStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DocumentAIProcessorList contains a list of DocumentAIProcessor
type DocumentAIProcessorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DocumentAIProcessor `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DocumentAIProcessor{}, &DocumentAIProcessorList{})
}
