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

package v1beta1

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var VertexAIDatasetGVK = GroupVersion.WithKind("VertexAIDataset")

type DatasetEncryptionSpec struct {
	// Required. The Cloud KMS resource identifier of the customer managed encryption key used to protect a resource.
	// Has the form: projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key.
	// The key needs to be in the same region as where the compute resource is created.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.EncryptionSpec.kms_key_name
	KmsKeyNameRef *refsv1beta1.KMSCryptoKeyRef `json:"kmsKeyNameRef,omitempty"`
}

// VertexAIDatasetSpec defines the desired state of VertexAIDataset
// +kcc:spec:proto=google.cloud.aiplatform.v1beta1.Dataset
type VertexAIDatasetSpec struct {
	// The user-defined name of the Dataset. The name can be up to 128 characters long and can consist of any UTF-8 characters.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Dataset.display_name
	DisplayName *string `json:"displayName"`

	// Immutable. Customer-managed encryption key spec for a Dataset. If set, this Dataset and all sub-resources of this Dataset will be secured by this key.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Dataset.encryption_spec
	EncryptionSpec *DatasetEncryptionSpec `json:"encryptionSpec,omitempty"`

	// Immutable. Points to a YAML file stored on Google Cloud Storage describing additional information about the Dataset. The schema is defined as an OpenAPI 3.0.2 Schema Object. The schema files that can be used here are found in gs://google-cloud-aiplatform/schema/dataset/metadata/.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Dataset.metadata_schema_uri
	MetadataSchemaURI *string `json:"metadataSchemaUri"`

	// The project that this resource belongs to.
	ProjectRef refs.ProjectRef `json:"projectRef"`

	// Immutable. The region of the dataset. eg us-central1.
	Region *string `json:"region,omitempty"`

	// Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource.
	ResourceID *string `json:"resourceID,omitempty"`
}

// VertexAIDatasetStatus defines the config connector machine state of VertexAIDataset
type VertexAIDatasetStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VertexAIDatasetObservedState `json:"observedState,omitempty"`
}

// VertexAIDatasetObservedState is the state of the VertexAIDataset resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.aiplatform.v1beta1.Dataset
type VertexAIDatasetObservedState struct {
	// Output only. Timestamp when this Dataset was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Dataset.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The resource name of the Dataset instance.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Dataset.name
	Name *string `json:"name,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvertexaidataset;gcpvertexaidatasets
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:metadata:labels="internal.cloud.google.com/additional-versions=v1alpha1"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VertexAIDataset is the Schema for the VertexAIDataset API
// +k8s:openapi-gen=true
type VertexAIDataset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VertexAIDatasetSpec   `json:"spec,omitempty"`
	Status VertexAIDatasetStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VertexAIDatasetList contains a list of VertexAIDataset
type VertexAIDatasetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VertexAIDataset `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VertexAIDataset{}, &VertexAIDatasetList{})
}
