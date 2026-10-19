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

// API sources for DialogflowDocument, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/dialogflow/v2/document.proto
// +kcc:source:service-docs=https://cloud.google.com/dialogflow/
// +kcc:source:resource-docs=https://docs.cloud.google.com/dialogflow/docs/reference/rest/v2/projects.knowledgeBases.documents

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DialogflowDocumentGVK = GroupVersion.WithKind("DialogflowDocument")

// DialogflowDocumentSpec defines the desired state of DialogflowDocument
// +kcc:spec:proto=google.cloud.dialogflow.v2.Document
// +kcc:required-from-proto
type DialogflowDocumentSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// A reference to the projects/{project}/knowledgeBases/{knowledge_base} this resource belongs to.
	// +kcc:guess
	// KnowledgeBaseRef *DialogflowKnowledgeBaseRef `json:"knowledgeBaseRef,omitempty"`

	// The DialogflowDocument name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. The display name of the document. The name must be 1024 bytes or
	//  less; otherwise, the creation request fails.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Document.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Required. The MIME type of this document.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Document.mime_type
	// +required
	MimeType *string `json:"mimeType,omitempty"`

	// Required. The knowledge type of document content.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Document.knowledge_types
	// +required
	KnowledgeTypes []string `json:"knowledgeTypes,omitempty"`

	// The URI where the file content is located.
	//
	//  For documents stored in Google Cloud Storage, these URIs must have
	//  the form `gs://<bucket-name>/<object-name>`.
	//
	//  NOTE: External URLs must correspond to public webpages, i.e., they must
	//  be indexed by Google Search. In particular, URLs for showing documents in
	//  Google Cloud Storage (i.e. the URL in your browser) are not supported.
	//  Instead use the `gs://` format URI described above.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Document.content_uri
	ContentURI *string `json:"contentURI,omitempty"`

	// The raw content of the document. This field is only permitted for
	//  EXTRACTIVE_QA and FAQ knowledge types.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Document.raw_content
	RawContent []byte `json:"rawContent,omitempty"`

	// Optional. If true, we try to automatically reload the document every day
	//  (at a time picked by the system). If false or unspecified, we don't try
	//  to automatically reload the document.
	//
	//  Currently you can only enable automatic reload for documents sourced from
	//  a public url, see `source` field for the source types.
	//
	//  Reload status can be tracked in `latest_reload_status`. If a reload
	//  fails, we will keep the document unchanged.
	//
	//  If a reload fails with internal errors, the system will try to reload the
	//  document on the next day.
	//  If a reload fails with non-retriable errors (e.g. PERMISSION_DENIED), the
	//  system will not try to reload the document anymore. You need to manually
	//  reload the document successfully by calling `ReloadDocument` and clear the
	//  errors.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Document.enable_auto_reload
	EnableAutoReload *bool `json:"enableAutoReload,omitempty"`

	// Optional. Metadata for the document. The metadata supports arbitrary
	//  key-value pairs. Suggested use cases include storing a document's title,
	//  an external URL distinct from the document's content_uri, etc.
	//  The max size of a `key` or a `value` of the metadata is 1024 bytes.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Document.metadata
	Metadata map[string]string `json:"metadata,omitempty"`
}

// DialogflowDocumentStatus defines the config connector machine state of DialogflowDocument
type DialogflowDocumentStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DialogflowDocument resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DialogflowDocumentObservedState `json:"observedState,omitempty"`
}

// DialogflowDocumentObservedState is the state of the DialogflowDocument resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.dialogflow.v2.Document
type DialogflowDocumentObservedState struct {
	// Output only. The time and status of the latest reload.
	//  This reload may have been triggered automatically or manually
	//  and may not have succeeded.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Document.latest_reload_status
	LatestReloadStatus *Document_ReloadStatus `json:"latestReloadStatus,omitempty"`

	// Output only. The current state of the document.
	// +kcc:proto:field=google.cloud.dialogflow.v2.Document.state
	State *string `json:"state,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdialogflowdocument;gcpdialogflowdocuments
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DialogflowDocument is the Schema for the DialogflowDocument API
// +k8s:openapi-gen=true
type DialogflowDocument struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DialogflowDocumentSpec   `json:"spec,omitempty"`
	Status DialogflowDocumentStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DialogflowDocumentList contains a list of DialogflowDocument
type DialogflowDocumentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DialogflowDocument `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DialogflowDocument{}, &DialogflowDocumentList{})
}
