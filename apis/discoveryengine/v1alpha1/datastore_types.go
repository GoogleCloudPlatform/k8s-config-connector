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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DiscoveryEngineDataStoreGVK = GroupVersion.WithKind("DiscoveryEngineDataStore")

type DiscoveryEngineDataStoreParent struct {
	// Required. The location of the application.
	// +required
	Location *string `json:"location,omitempty"`

	// Required. The host project of the application.
	// +required
	ProjectRef *v1beta1.ProjectRef `json:"projectRef,omitempty"`

	// TODO: The collection is a parent but it does not show up in the proto.
	// Collection is optional.

	// Immutable. The collection for the DataStore.
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Collection field is immutable"
	Collection *string `json:"collection"`
}

// DiscoveryEngineDataStoreSpec defines the desired state of DiscoveryEngineDataStore
// +kcc:proto=google.cloud.discoveryengine.v1alpha.DataStore
type DiscoveryEngineDataStoreSpec struct {
	// The DiscoveryEngineDataStore name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. Defines the parent path of the resource.
	DiscoveryEngineDataStoreParent `json:",inline"`

	// Required. The data store display name.
	//
	//  This field must be a UTF-8 encoded string with a length limit of 128
	//  characters. Otherwise, an INVALID_ARGUMENT error is returned.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.DataStore.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="the field is immutable"
	// Immutable. The industry vertical that the data store registers.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.DataStore.industry_vertical
	IndustryVertical *string `json:"industryVertical,omitempty"`

	// Immutable.
	// The solutions that the data store enrolls. Available solutions for each
	//  [industry_vertical][google.cloud.discoveryengine.v1alpha.DataStore.industry_vertical]:
	//
	//  * `MEDIA`: `SOLUTION_TYPE_RECOMMENDATION` and `SOLUTION_TYPE_SEARCH`.
	//  * `SITE_SEARCH`: `SOLUTION_TYPE_SEARCH` is automatically enrolled. Other
	//    solutions cannot be enrolled.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.DataStore.solution_types
	SolutionTypes []string `json:"solutionTypes,omitempty"`

	// +kubebuilder:validation:Enum=CONTENT_CONFIG_UNSPECIFIED;NO_CONTENT;CONTENT_REQUIRED;PUBLIC_WEBSITE;GOOGLE_WORKSPACE
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="the field is immutable"
	// Immutable. The content config of the data store. If this field is unset,
	//  the server behavior defaults to
	//  [ContentConfig.NO_CONTENT][google.cloud.discoveryengine.v1alpha.DataStore.ContentConfig.NO_CONTENT].
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.DataStore.content_config
	ContentConfig *string `json:"contentConfig,omitempty"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="the field is immutable"
	// Data store level identity provider config.
	// This needs to be set up separately in the Vertex AI "Authentication settings"
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.DataStore.idp_config
	IdpConfig *IdpConfig `json:"idpConfig,omitempty"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="the field is immutable"
	// Immutable. Whether data in the
	//  [DataStore][google.cloud.discoveryengine.v1alpha.DataStore] has ACL
	//  information. If set to `true`, the source data must have ACL. ACL will be
	//  devingested when data is ingested by
	//  [DocumentService.ImportDocuments][google.cloud.discoveryengine.v1alpha.DocumentService.ImportDocuments]
	//  methods.
	//
	//  When ACL is enabled for the
	//  [DataStore][google.cloud.discoveryengine.v1alpha.DataStore],
	//  [Document][google.cloud.discoveryengine.v1alpha.Document] can't be accessed
	//  by calling
	//  [DocumentService.GetDocument][google.cloud.discoveryengine.v1alpha.DocumentService.GetDocument]
	//  or
	//  [DocumentService.ListDocuments][google.cloud.discoveryengine.v1alpha.DocumentService.ListDocuments].
	//
	//  Currently ACL is only supported in `GENERIC` industry vertical with
	//  non-`PUBLIC_WEBSITE` content config.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.DataStore.acl_enabled
	AclEnabled *bool `json:"aclEnabled,omitempty"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="the field is immutable"
	// Config to store data store type configuration for workspace data. This
	//  must be set when
	//  [DataStore.content_config][google.cloud.discoveryengine.v1alpha.DataStore.content_config]
	//  is set as
	//  [DataStore.ContentConfig.GOOGLE_WORKSPACE][google.cloud.discoveryengine.v1alpha.DataStore.ContentConfig.GOOGLE_WORKSPACE].
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.DataStore.workspace_config
	WorkspaceConfig *WorkspaceConfig `json:"workspaceConfig,omitempty"`

	// Language info for DataStore.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.DataStore.language_info
	LanguageInfo *LanguageInfo `json:"languageInfo,omitempty"`

	// Configuration for Document understanding and enrichment.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.DataStore.document_processing_config
	DocumentProcessingConfig *DocumentProcessingConfig `json:"documentProcessingConfig,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1alpha.IdpConfig
type IdpConfig struct {
	// Identity provider type configured.
	// +kubebuilder:validation:Enum=THIRD_PARTY;GSUITE;IDP_TYPE_UNSPECIFIED
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.IdpConfig.idp_type
	IdpType *string `json:"idpType,omitempty"`

	// External Identity provider config.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.IdpConfig.external_idp_config
	ExternalIdpConfig *IdpConfig_ExternalIdpConfig `json:"externalIdpConfig,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1alpha.DocumentProcessingConfig
type DocumentProcessingConfig struct {
	// Whether chunking mode is enabled.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.DocumentProcessingConfig.chunking_config
	ChunkingConfig *DocumentProcessingConfig_ChunkingConfig `json:"chunkingConfig,omitempty"`

	// Configurations for default Document parser.
	//  If not specified, we will configure it as default DigitalParsingConfig, and
	//  the default parsing config will be applied to all file types for Document
	//  parsing.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.DocumentProcessingConfig.default_parsing_config
	DefaultParsingConfig *DocumentProcessingConfig_ParsingConfig `json:"defaultParsingConfig,omitempty"`

	// TODO: unsupported map type with key string and value message
}

// +kcc:proto=google.cloud.discoveryengine.v1alpha.LanguageInfo
type LanguageInfo struct {
	// The language code for the DataStore. See https://cloud.google.com/vertex-ai/docs/general/supported-languages
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.LanguageInfo.language_code
	LanguageCode *string `json:"languageCode,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1alpha.DataStore
type DataStoreObservedState struct {
	// Output only. The id of the default
	//  [Schema][google.cloud.discoveryengine.v1alpha.Schema] asscociated to this
	//  data store.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.DataStore.default_schema_id
	DefaultSchemaID *string `json:"defaultSchemaID,omitempty"`

	// Output only. Timestamp the
	//  [DataStore][google.cloud.discoveryengine.v1alpha.DataStore] was created at.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.DataStore.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Language info for DataStore.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.DataStore.language_info
	LanguageInfo *LanguageInfoObservedState `json:"languageInfo,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1alpha.WorkspaceConfig
type WorkspaceConfig struct {
	// +kubebuilder:validation:Enum=TYPE_UNSPECIFIED;GOOGLE_DRIVE;GOOGLE_MAIL;GOOGLE_SITES;GOOGLE_CALENDAR;GOOGLE_CHAT;GOOGLE_GROUPS;GOOGLE_KEEP
	// The Google Workspace data source. Valid values are TYPE_UNSPECIFIED, GOOGLE_DRIVE, GOOGLE_MAIL, GOOGLE_SITES, GOOGLE_CALENDAR,
	// GOOGLE_CHAT, GOOGLE_GROUPS, GOOGLE_KEEP
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.WorkspaceConfig.type
	Type *string `json:"type,omitempty"`

	// Obfuscated Dasher customer ID.
	// +kcc:proto:field=google.cloud.discoveryengine.v1alpha.WorkspaceConfig.dasher_customer_id
	DasherCustomerID *string `json:"dasherCustomerID,omitempty"`
}

// DiscoveryEngineDataStoreStatus defines the config connector machine state of DiscoveryEngineDataStore
type DiscoveryEngineDataStoreStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DiscoveryEngineDataStore resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DataStoreObservedState `json:"observedState,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdiscoveryenginedatastore;gcpdiscoveryenginedatastores
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DiscoveryEngineDataStore is the Schema for the DiscoveryEngineDataStore API
// +k8s:openapi-gen=true
type DiscoveryEngineDataStore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DiscoveryEngineDataStoreSpec   `json:"spec,omitempty"`
	Status DiscoveryEngineDataStoreStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DiscoveryEngineDataStoreList contains a list of DiscoveryEngineDataStore
type DiscoveryEngineDataStoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DiscoveryEngineDataStore `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DiscoveryEngineDataStore{}, &DiscoveryEngineDataStoreList{})
}
