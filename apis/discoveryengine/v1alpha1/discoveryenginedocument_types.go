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

// API sources for DiscoveryengineDocument, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/discoveryengine/v1/document.proto
// +kcc:source:service-docs=https://cloud.google.com/generative-ai-app-builder/docs/
// +kcc:source:resource-docs=https://docs.cloud.google.com/generative-ai-app-builder/docs/reference/rest/v1/projects.locations.dataStores.branches.documents

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DiscoveryengineDocumentGVK = GroupVersion.WithKind("DiscoveryengineDocument")

// DiscoveryengineDocumentSpec defines the desired state of DiscoveryengineDocument
// +kcc:spec:proto=google.cloud.discoveryengine.v1.Document
// +kcc:required-from-proto
type DiscoveryengineDocumentSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The DiscoveryengineDocument name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// The structured JSON data for the document. It should conform to the
	//  registered [Schema][google.cloud.discoveryengine.v1.Schema] or an
	//  `INVALID_ARGUMENT` error is thrown.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Document.struct_data
	StructData apiextensionsv1.JSON `json:"structData,omitempty"`

	// The JSON string representation of the document. It should conform to the
	//  registered [Schema][google.cloud.discoveryengine.v1.Schema] or an
	//  `INVALID_ARGUMENT` error is thrown.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Document.json_data
	JsonData *string `json:"jsonData,omitempty"`

	// Immutable. The identifier of the document.
	//
	//  Id should conform to [RFC-1034](https://tools.ietf.org/html/rfc1034)
	//  standard with a length limit of 128 characters.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Document.id
	ID *string `json:"id,omitempty"`

	// The identifier of the schema located in the same data store.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Document.schema_id
	SchemaID *string `json:"schemaID,omitempty"`

	// The unstructured data linked to this document. Content can only be set
	//  and must be set if this document is under a `CONTENT_REQUIRED` data store.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Document.content
	Content *Document_Content `json:"content,omitempty"`

	// The identifier of the parent document. Currently supports at most two level
	//  document hierarchy.
	//
	//  Id should conform to [RFC-1034](https://tools.ietf.org/html/rfc1034)
	//  standard with a length limit of 63 characters.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Document.parent_document_id
	ParentDocumentID *string `json:"parentDocumentID,omitempty"`

	// Access control information for the document.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Document.acl_info
	AclInfo *Document_AclInfo `json:"aclInfo,omitempty"`
}

// DiscoveryengineDocumentStatus defines the config connector machine state of DiscoveryengineDocument
type DiscoveryengineDocumentStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DiscoveryengineDocument resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DiscoveryengineDocumentObservedState `json:"observedState,omitempty"`
}

// DiscoveryengineDocumentObservedState is the state of the DiscoveryengineDocument resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.discoveryengine.v1.Document
type DiscoveryengineDocumentObservedState struct {
	// Output only. This field is OUTPUT_ONLY.
	//  It contains derived data that are not in the original input document.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Document.derived_struct_data
	DerivedStructData apiextensionsv1.JSON `json:"derivedStructData,omitempty"`

	// Output only. The last time the document was indexed. If this field is set,
	//  the document could be returned in search results.
	//
	//  This field is OUTPUT_ONLY. If this field is not populated, it means the
	//  document has never been indexed.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Document.index_time
	IndexTime *string `json:"indexTime,omitempty"`

	// Output only. The index status of the document.
	//
	//  * If document is indexed successfully, the index_time field is populated.
	//  * Otherwise, if document is not indexed due to errors, the error_samples
	//    field is populated.
	//  * Otherwise, if document's index is in progress, the pending_message field
	//    is populated.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.Document.index_status
	IndexStatus *Document_IndexStatus `json:"indexStatus,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdiscoveryenginedocument;gcpdiscoveryenginedocuments
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DiscoveryengineDocument is the Schema for the DiscoveryengineDocument API
// +k8s:openapi-gen=true
type DiscoveryengineDocument struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DiscoveryengineDocumentSpec   `json:"spec,omitempty"`
	Status DiscoveryengineDocumentStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DiscoveryengineDocumentList contains a list of DiscoveryengineDocument
type DiscoveryengineDocumentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DiscoveryengineDocument `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DiscoveryengineDocument{}, &DiscoveryengineDocumentList{})
}
