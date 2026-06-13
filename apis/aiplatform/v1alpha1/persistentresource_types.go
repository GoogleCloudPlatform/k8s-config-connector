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
	common "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var VertexAIPersistentResourceGVK = GroupVersion.WithKind("VertexAIPersistentResource")

// VertexAIPersistentResourceSpec defines the desired state of VertexAIPersistentResource
// +kcc:spec:proto=google.cloud.aiplatform.v1.PersistentResource
type VertexAIPersistentResourceSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location *string `json:"location"`

	// The VertexAIPersistentResource name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The spec of the pools of different resources.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PersistentResource.resource_pools
	// +required
	ResourcePools []ResourcePool `json:"resourcePools"`

	// Optional. User-defined metadata labels for the PersistentResource.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PersistentResource.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. PersistentResource runtime spec.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PersistentResource.resource_runtime_spec
	ResourceRuntimeSpec *ResourceRuntimeSpec `json:"resourceRuntimeSpec,omitempty"`

	// Optional. Customer-managed encryption key spec for a PersistentResource. If set, this PersistentResource will be secured by this key.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PersistentResource.encryption_spec
	EncryptionSpec *EncryptionSpec `json:"encryptionSpec,omitempty"`
}

// VertexAIPersistentResourceStatus defines the config connector machine state of VertexAIPersistentResource
type VertexAIPersistentResourceStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VertexAIPersistentResource resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VertexAIPersistentResourceObservedState `json:"observedState,omitempty"`
}

// VertexAIPersistentResourceObservedState is the state of the VertexAIPersistentResource resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.aiplatform.v1.PersistentResource
type VertexAIPersistentResourceObservedState struct {
	// Output only. Time when the PersistentResource was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PersistentResource.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Only populated when the PersistentResource is RUNNING.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PersistentResource.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Output only. Only populated when the PersistentResource is RUNNING.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PersistentResource.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The detailed state of a Study.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PersistentResource.state
	State *string `json:"state,omitempty"`

	// Output only. Only populated when persistent resource's state is `STOPPING` or `ERROR`.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PersistentResource.error
	Error *common.Status `json:"error,omitempty"`

	// Output only. Runtime information of the Persistent Resource.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PersistentResource.resource_runtime
	ResourceRuntime *ResourceRuntimeObservedState `json:"resourceRuntime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvertexaipersistentresource;gcpvertexaipersistentresources
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VertexAIPersistentResource is the Schema for the VertexAIPersistentResource API
// +k8s:openapi-gen=true
type VertexAIPersistentResource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VertexAIPersistentResourceSpec   `json:"spec,omitempty"`
	Status VertexAIPersistentResourceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VertexAIPersistentResourceList contains a list of VertexAIPersistentResource
type VertexAIPersistentResourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VertexAIPersistentResource `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VertexAIPersistentResource{}, &VertexAIPersistentResourceList{})
}
