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


// +kcc:proto=google.cloud.documentai.v1beta3.ProcessorType
type ProcessorType struct {
	// The resource name of the processor type.
	//  Format: `projects/{project}/processorTypes/{processor_type}`
	// +kcc:proto:field=google.cloud.documentai.v1beta3.ProcessorType.name
	Name *string `json:"name,omitempty"`

	// The processor type, such as: `OCR_PROCESSOR`, `INVOICE_PROCESSOR`.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.ProcessorType.type
	Type *string `json:"type,omitempty"`

	// The processor category, used by UI to group processor types.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.ProcessorType.category
	Category *string `json:"category,omitempty"`

	// The locations in which this processor is available.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.ProcessorType.available_locations
	AvailableLocations []ProcessorType_LocationInfo `json:"availableLocations,omitempty"`

	// Whether the processor type allows creation. If true, users can create a
	//  processor of this processor type. Otherwise, users need to request access.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.ProcessorType.allow_creation
	AllowCreation *bool `json:"allowCreation,omitempty"`

	// Launch stage of the processor type
	// +kcc:proto:field=google.cloud.documentai.v1beta3.ProcessorType.launch_stage
	LaunchStage *string `json:"launchStage,omitempty"`

	// A set of Cloud Storage URIs of sample documents for this processor.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.ProcessorType.sample_document_uris
	SampleDocumentUris []string `json:"sampleDocumentUris,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1beta3.ProcessorType.LocationInfo
type ProcessorType_LocationInfo struct {
	// The location ID. For supported locations, refer to [regional and
	//  multi-regional support](/document-ai/docs/regions).
	// +kcc:proto:field=google.cloud.documentai.v1beta3.ProcessorType.LocationInfo.location_id
	LocationID *string `json:"locationID,omitempty"`
}
