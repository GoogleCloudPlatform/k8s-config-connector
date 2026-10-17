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

// API sources for DataplexEncryptionConfig, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/e7513c828b8db536fa12fc5493d541f323aedf90/google/cloud/dataplex/v1/cmek.proto
// +kcc:source:service-docs=https://cloud.google.com/dataplex/docs
// +kcc:source:resource-docs=https://docs.cloud.google.com/dataplex/docs/reference/rest/v1/organizations.locations.encryptionConfigs

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DataplexEncryptionConfigGVK = GroupVersion.WithKind("DataplexEncryptionConfig")

// DataplexEncryptionConfigSpec defines the desired state of DataplexEncryptionConfig
// +kcc:spec:proto=google.cloud.dataplex.v1.EncryptionConfig
// +kcc:required-from-proto
type DataplexEncryptionConfigSpec struct {
	// The organization that this resource belongs to.
	OrganizationRef *refsv1beta1.OrganizationRef `json:"organizationRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The DataplexEncryptionConfig name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. If a key is chosen, it means that the customer is using CMEK.
	//  If a key is not chosen, it means that the customer is using Google managed
	//  encryption.
	// +kcc:proto:field=google.cloud.dataplex.v1.EncryptionConfig.key
	Key *string `json:"key,omitempty"`

	// Etag of the EncryptionConfig. This is a strong etag.
	// +kcc:proto:field=google.cloud.dataplex.v1.EncryptionConfig.etag
	Etag *string `json:"etag,omitempty"`

	// Optional. Represent the state of CMEK opt-in for metastore.
	// +kcc:proto:field=google.cloud.dataplex.v1.EncryptionConfig.enable_metastore_encryption
	EnableMetastoreEncryption *bool `json:"enableMetastoreEncryption,omitempty"`
}

// DataplexEncryptionConfigStatus defines the config connector machine state of DataplexEncryptionConfig
type DataplexEncryptionConfigStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DataplexEncryptionConfig resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DataplexEncryptionConfigObservedState `json:"observedState,omitempty"`
}

// DataplexEncryptionConfigObservedState is the state of the DataplexEncryptionConfig resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.dataplex.v1.EncryptionConfig
type DataplexEncryptionConfigObservedState struct {
	// Output only. The time when the Encryption configuration was created.
	// +kcc:proto:field=google.cloud.dataplex.v1.EncryptionConfig.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the Encryption configuration was last updated.
	// +kcc:proto:field=google.cloud.dataplex.v1.EncryptionConfig.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The state of encryption of the databases.
	// +kcc:proto:field=google.cloud.dataplex.v1.EncryptionConfig.encryption_state
	EncryptionState *string `json:"encryptionState,omitempty"`

	// Output only. Details of the failure if anything related to Cmek db fails.
	// +kcc:proto:field=google.cloud.dataplex.v1.EncryptionConfig.failure_details
	FailureDetails *EncryptionConfig_FailureDetailsObservedState `json:"failureDetails,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdataplexencryptionconfig;gcpdataplexencryptionconfigs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataplexEncryptionConfig is the Schema for the DataplexEncryptionConfig API
// +k8s:openapi-gen=true
type DataplexEncryptionConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DataplexEncryptionConfigSpec   `json:"spec,omitempty"`
	Status DataplexEncryptionConfigStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DataplexEncryptionConfigList contains a list of DataplexEncryptionConfig
type DataplexEncryptionConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataplexEncryptionConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataplexEncryptionConfig{}, &DataplexEncryptionConfigList{})
}
