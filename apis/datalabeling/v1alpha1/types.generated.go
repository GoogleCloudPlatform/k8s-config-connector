// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1


// +kcc:proto=google.cloud.datalabeling.v1beta1.CsvInstruction
type CsvInstruction struct {
	// CSV file for the instruction. Only gcs path is allowed.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.CsvInstruction.gcs_file_uri
	GcsFileURI *string `json:"gcsFileURI,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.Instruction
type Instruction struct {
	// Output only. Instruction resource name, format:
	//  projects/{project_id}/instructions/{instruction_id}
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.Instruction.name
	Name *string `json:"name,omitempty"`

	// Required. The display name of the instruction. Maximum of 64 characters.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.Instruction.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. User-provided description of the instruction.
	//  The description can be up to 10000 characters long.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.Instruction.description
	Description *string `json:"description,omitempty"`

	// Output only. Creation time of instruction.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.Instruction.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Last update time of instruction.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.Instruction.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Required. The data type of this instruction.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.Instruction.data_type
	DataType *string `json:"dataType,omitempty"`

	// Deprecated: this instruction format is not supported any more.
	//  Instruction from a CSV file, such as for classification task.
	//  The CSV file should have exact two columns, in the following format:
	//
	//  * The first column is labeled data, such as an image reference, text.
	//  * The second column is comma separated labels associated with data.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.Instruction.csv_instruction
	CsvInstruction *CsvInstruction `json:"csvInstruction,omitempty"`

	// Instruction from a PDF document. The PDF should be in a Cloud Storage
	//  bucket.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.Instruction.pdf_instruction
	PdfInstruction *PdfInstruction `json:"pdfInstruction,omitempty"`

	// Output only. The names of any related resources that are blocking changes
	//  to the instruction.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.Instruction.blocking_resources
	BlockingResources []string `json:"blockingResources,omitempty"`
}

// +kcc:proto=google.cloud.datalabeling.v1beta1.PdfInstruction
type PdfInstruction struct {
	// PDF file for the instruction. Only gcs path is allowed.
	// +kcc:proto:field=google.cloud.datalabeling.v1beta1.PdfInstruction.gcs_file_uri
	GcsFileURI *string `json:"gcsFileURI,omitempty"`
}
