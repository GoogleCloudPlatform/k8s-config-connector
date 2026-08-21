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

var DataLabelingInstructionGVK = GroupVersion.WithKind("DataLabelingInstruction")

// DataLabelingInstructionSpec defines the desired state of DataLabelingInstruction
// +kcc:spec:proto=google.cloud.datalabeling.v1beta1.Instruction
type DataLabelingInstructionSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The DataLabelingInstruction name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The display name of the instruction. Maximum of 64 characters.
	// +kubebuilder:validation:Required
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.Instruction.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. User-provided description of the instruction.
	//  The description can be up to 10000 characters long.
	// +kubebuilder:validation:Optional
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.Instruction.description
	Description *string `json:"description,omitempty"`

	// Required. The data type of this instruction.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=DATA_TYPE_UNSPECIFIED;IMAGE;VIDEO;TEXT;GENERAL_DATA
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.Instruction.data_type
	DataType *string `json:"dataType,omitempty"`

	// Deprecated: this instruction format is not supported any more.
	//  Instruction from a CSV file, such as for classification task.
	//  The CSV file should have exact two columns, in the following format:
	//
	//  * The first column is labeled data, such as an image reference, text.
	//  * The second column is comma separated labels associated with data.
	// +kubebuilder:validation:Optional
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.Instruction.csv_instruction
	CsvInstruction *InstructionCsvInstruction `json:"csvInstruction,omitempty"`

	// Instruction from a PDF document. The PDF should be in a Cloud Storage bucket.
	// +kubebuilder:validation:Optional
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.Instruction.pdf_instruction
	PdfInstruction *InstructionPdfInstruction `json:"pdfInstruction,omitempty"`
}

type InstructionCsvInstruction struct {
	// CSV file for the instruction. Only gcs path is allowed.
	// +kubebuilder:validation:Required
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.CsvInstruction.gcs_file_uri
	GcsFileURI *string `json:"gcsFileURI,omitempty"`
}

type InstructionPdfInstruction struct {
	// PDF file for the instruction. Only gcs path is allowed.
	// +kubebuilder:validation:Required
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.PdfInstruction.gcs_file_uri
	GcsFileURI *string `json:"gcsFileURI,omitempty"`
}

// DataLabelingInstructionStatus defines the config connector machine state of DataLabelingInstruction
type DataLabelingInstructionStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DataLabelingInstruction resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DataLabelingInstructionObservedState `json:"observedState,omitempty"`
}

// DataLabelingInstructionObservedState is the state of the DataLabelingInstruction resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.datalabeling.v1beta1.Instruction
type DataLabelingInstructionObservedState struct {
	// Output only. Creation time of instruction.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.Instruction.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Last update time of instruction.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.Instruction.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The names of any related resources that are blocking changes to the instruction.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.Instruction.blocking_resources
	BlockingResources []string `json:"blockingResources,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdatalabelinginstruction;gcpdatalabelinginstructions
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataLabelingInstruction is the Schema for the DataLabelingInstruction API
// +k8s:openapi-gen=true
type DataLabelingInstruction struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DataLabelingInstructionSpec   `json:"spec,omitempty"`
	Status DataLabelingInstructionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DataLabelingInstructionList contains a list of DataLabelingInstruction
type DataLabelingInstructionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataLabelingInstruction `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataLabelingInstruction{}, &DataLabelingInstructionList{})
}
