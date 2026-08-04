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

var VertexAITensorboardGVK = GroupVersion.WithKind("VertexAITensorboard")

// VertexAITensorboardSpec defines the desired state of VertexAITensorboard
// +kcc:spec:proto=google.cloud.aiplatform.v1beta1.Tensorboard
type VertexAITensorboardSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The region of this resource.
	Region string `json:"region"`

	// The VertexAITensorboard ID (which is server-generated). If not given, Config Connector will create a new Tensorboard. If given, Config Connector will acquire the existing Tensorboard with this ID.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. User provided name of this Tensorboard.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Tensorboard.display_name
	DisplayName string `json:"displayName"`

	// Description of this Tensorboard.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Tensorboard.description
	Description *string `json:"description,omitempty"`

	// Customer-managed encryption key spec for a Tensorboard. If set, this
	//  Tensorboard and all sub-resources of this Tensorboard will be secured by
	//  this key.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Tensorboard.encryption_spec
	EncryptionSpec *EncryptionSpec `json:"encryptionSpec,omitempty"`

	// Used to indicate if the TensorBoard instance is the default one.
	//  Each project & region can have at most one default TensorBoard instance.
	//  Creation of a default TensorBoard instance and updating an existing
	//  TensorBoard instance to be default will mark all other TensorBoard
	//  instances (if any) as non default.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Tensorboard.is_default
	IsDefault *bool `json:"isDefault,omitempty"`
}

// VertexAITensorboardStatus defines the config connector machine state of VertexAITensorboard
type VertexAITensorboardStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VertexAITensorboard resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VertexAITensorboardObservedState `json:"observedState,omitempty"`
}

// VertexAITensorboardObservedState is the state of the VertexAITensorboard resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.aiplatform.v1beta1.Tensorboard
type VertexAITensorboardObservedState struct {
	// Output only. Name of the Tensorboard.
	//  Format:
	//  `projects/{project}/locations/{location}/tensorboards/{tensorboard}`
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Tensorboard.name
	Name *string `json:"name,omitempty"`

	// Output only. Consumer project Cloud Storage path prefix used to store blob
	//  data, which can either be a bucket or directory. Does not end with a '/'.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Tensorboard.blob_storage_path_prefix
	BlobStoragePathPrefix *string `json:"blobStoragePathPrefix,omitempty"`

	// Output only. The number of Runs stored in this Tensorboard.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Tensorboard.run_count
	RunCount *int32 `json:"runCount,omitempty"`

	// Output only. Timestamp when this Tensorboard was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Tensorboard.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this Tensorboard was last updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Tensorboard.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Tensorboard.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.Tensorboard.satisfies_pzi
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvertexaitensorboard;gcpvertexaitensorboards
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VertexAITensorboard is the Schema for the VertexAITensorboard API
// +k8s:openapi-gen=true
type VertexAITensorboard struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VertexAITensorboardSpec   `json:"spec,omitempty"`
	Status VertexAITensorboardStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VertexAITensorboardList contains a list of VertexAITensorboard
type VertexAITensorboardList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VertexAITensorboard `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VertexAITensorboard{}, &VertexAITensorboardList{})
}
