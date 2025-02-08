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


// +kcc:proto=google.cloud.contentwarehouse.v1.DocumentLink
type DocumentLink struct {
	// Name of this document-link.
	//  It is required that the parent derived form the name to be consistent with
	//  the source document reference. Otherwise an exception will be thrown.
	//  Format:
	//  projects/{project_number}/locations/{location}/documents/{source_document_id}/documentLinks/{document_link_id}.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.DocumentLink.name
	Name *string `json:"name,omitempty"`

	// Document references of the source document.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.DocumentLink.source_document_reference
	SourceDocumentReference *DocumentReference `json:"sourceDocumentReference,omitempty"`

	// Document references of the target document.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.DocumentLink.target_document_reference
	TargetDocumentReference *DocumentReference `json:"targetDocumentReference,omitempty"`

	// Description of this document-link.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.DocumentLink.description
	Description *string `json:"description,omitempty"`

	// The state of the documentlink. If target node has been deleted, the
	//  link is marked as invalid. Removing a source node will result in removal
	//  of all associated links.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.DocumentLink.state
	State *string `json:"state,omitempty"`
}

// +kcc:proto=google.cloud.contentwarehouse.v1.DocumentReference
type DocumentReference struct {
	// Required. Name of the referenced document.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.DocumentReference.document_name
	DocumentName *string `json:"documentName,omitempty"`

	// display_name of the referenced document; this name does not need to be
	//  consistent to the display_name in the Document proto, depending on the ACL
	//  constraint.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.DocumentReference.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Stores the subset of the referenced document's content.
	//  This is useful to allow user peek the information of the referenced
	//  document.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.DocumentReference.snippet
	Snippet *string `json:"snippet,omitempty"`

	// The document type of the document being referenced.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.DocumentReference.document_is_folder
	DocumentIsFolder *bool `json:"documentIsFolder,omitempty"`

	// Document is a folder with retention policy.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.DocumentReference.document_is_retention_folder
	DocumentIsRetentionFolder *bool `json:"documentIsRetentionFolder,omitempty"`

	// Document is a folder with legal hold.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.DocumentReference.document_is_legal_hold_folder
	DocumentIsLegalHoldFolder *bool `json:"documentIsLegalHoldFolder,omitempty"`
}

// +kcc:proto=google.cloud.contentwarehouse.v1.DocumentLink
type DocumentLinkObservedState struct {
	// Document references of the source document.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.DocumentLink.source_document_reference
	SourceDocumentReference *DocumentReferenceObservedState `json:"sourceDocumentReference,omitempty"`

	// Output only. The time when the documentLink is last updated.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.DocumentLink.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The time when the documentLink is created.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.DocumentLink.create_time
	CreateTime *string `json:"createTime,omitempty"`
}

// +kcc:proto=google.cloud.contentwarehouse.v1.DocumentReference
type DocumentReferenceObservedState struct {
	// Output only. The time when the document is last updated.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.DocumentReference.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The time when the document is created.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.DocumentReference.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the document is deleted.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.DocumentReference.delete_time
	DeleteTime *string `json:"deleteTime,omitempty"`
}
