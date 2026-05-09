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

var VertexAIDeploymentResourcePoolGVK = GroupVersion.WithKind("VertexAIDeploymentResourcePool")

// VertexAIDeploymentResourcePoolSpec defines the desired state of VertexAIDeploymentResourcePool
// +kcc:spec:proto=google.cloud.aiplatform.v1beta1.DeploymentResourcePool
type VertexAIDeploymentResourcePoolSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
<<<<<<< ours
	Location *string `json:"location"`
=======
	Location string `json:"location"`

	// The VertexAIDeploymentResourcePool name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The underlying DedicatedResources that the DeploymentResourcePool
	//  uses.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.DeploymentResourcePool.dedicated_resources
	DedicatedResources *DedicatedResources `json:"dedicatedResources,omitempty"`

	// Customer-managed encryption key spec for a DeploymentResourcePool. If set,
	//  this DeploymentResourcePool will be secured by this key. Endpoints and the
	//  DeploymentResourcePool they deploy in need to have the same EncryptionSpec.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.DeploymentResourcePool.encryption_spec
	EncryptionSpec *EncryptionSpec `json:"encryptionSpec,omitempty"`

	// The service account that the DeploymentResourcePool's container(s) run as.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.DeploymentResourcePool.service_account
	ServiceAccountRef *refsv1beta1.IAMServiceAccountRef `json:"serviceAccountRef,omitempty"`

	// If the DeploymentResourcePool is deployed with custom-trained Models or
	//  AutoML Tabular Models, the container(s) of the DeploymentResourcePool will
	//  send `stderr` and `stdout` streams to Cloud Logging by default.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.DeploymentResourcePool.disable_container_logging
	DisableContainerLogging *bool `json:"disableContainerLogging,omitempty"`
}

// VertexAIDeploymentResourcePoolStatus defines the config connector machine state of VertexAIDeploymentResourcePool
type VertexAIDeploymentResourcePoolStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VertexAIDeploymentResourcePool resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VertexAIDeploymentResourcePoolObservedState `json:"observedState,omitempty"`
}

// VertexAIDeploymentResourcePoolObservedState is the state of the VertexAIDeploymentResourcePool resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.aiplatform.v1beta1.DeploymentResourcePool
type VertexAIDeploymentResourcePoolObservedState struct {
	// Output only. Timestamp when this DeploymentResourcePool was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.DeploymentResourcePool.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.DeploymentResourcePool.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1beta1.DeploymentResourcePool.satisfies_pzi
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvertexaideploymentresourcepool;gcpvertexaideploymentresourcepools
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VertexAIDeploymentResourcePool is the Schema for the VertexAIDeploymentResourcePool API
// +k8s:openapi-gen=true
type VertexAIDeploymentResourcePool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VertexAIDeploymentResourcePoolSpec   `json:"spec,omitempty"`
	Status VertexAIDeploymentResourcePoolStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VertexAIDeploymentResourcePoolList contains a list of VertexAIDeploymentResourcePool
type VertexAIDeploymentResourcePoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VertexAIDeploymentResourcePool `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VertexAIDeploymentResourcePool{}, &VertexAIDeploymentResourcePoolList{})
}
