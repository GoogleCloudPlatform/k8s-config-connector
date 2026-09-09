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
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AIPlatformPersistentResourceGVK = GroupVersion.WithKind("AIPlatformPersistentResource")

// AIPlatformPersistentResourceSpec defines the desired state of AIPlatformPersistentResource
// +kcc:spec:proto=google.cloud.aiplatform.v1.PersistentResource
type AIPlatformPersistentResourceSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The AIPlatformPersistentResource name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. The display name of the PersistentResource.
	// The name can be up to 128 characters long and can consist of any UTF-8
	// characters.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PersistentResource.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. The spec of the pools of different resources.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PersistentResource.resource_pools
	ResourcePools []ResourcePool `json:"resourcePools,omitempty"`

	// Optional. The labels with user-defined metadata to organize PersistentResource.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PersistentResource.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. The full name of the Compute Engine network to peered with Vertex AI to host the persistent resources.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PersistentResource.network
	NetworkRef *computev1beta1.ComputeNetworkRef `json:"networkRef,omitempty"`

	// Optional. Configuration for PSC-I for PersistentResource.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PersistentResource.psc_interface_config
	PSCInterfaceConfig *PSCInterfaceConfig `json:"pscInterfaceConfig,omitempty"`

	// Optional. Customer-managed encryption key spec for a PersistentResource.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PersistentResource.encryption_spec
	EncryptionSpec *EncryptionSpec `json:"encryptionSpec,omitempty"`

	// Optional. Persistent Resource runtime spec.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PersistentResource.resource_runtime_spec
	ResourceRuntimeSpec *ResourceRuntimeSpec `json:"resourceRuntimeSpec,omitempty"`

	// Optional. A list of names for the reserved IP ranges under the VPC network that can be used for this persistent resource.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PersistentResource.reserved_ip_ranges
	ReservedIPRanges []string `json:"reservedIPRanges,omitempty"`
}

// AIPlatformPersistentResourceStatus defines the config connector machine state of AIPlatformPersistentResource
type AIPlatformPersistentResourceStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AIPlatformPersistentResource resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AIPlatformPersistentResourceObservedState `json:"observedState,omitempty"`
}

// AIPlatformPersistentResourceObservedState is the state of the AIPlatformPersistentResource resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.aiplatform.v1.PersistentResource
type AIPlatformPersistentResourceObservedState struct {
	// Required. The spec of the pools of different resources.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PersistentResource.resource_pools
	ResourcePools []ResourcePoolObservedState `json:"resourcePools,omitempty"`

	// Output only. The detailed state of a Study.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PersistentResource.state
	State *string `json:"state,omitempty"`

	// Output only. Only populated when persistent resource's state is `STOPPING`
	//  or `ERROR`.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PersistentResource.error
	Error *common.Status `json:"error,omitempty"`

	// Output only. Time when the PersistentResource was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PersistentResource.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time when the PersistentResource for the first time entered
	//  the `RUNNING` state.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PersistentResource.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Output only. Time when the PersistentResource was most recently updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PersistentResource.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Runtime information of the Persistent Resource.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PersistentResource.resource_runtime
	ResourceRuntime *ResourceRuntimeObservedState `json:"resourceRuntime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpaiplatformpersistentresource;gcpaiplatformpersistentresources
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AIPlatformPersistentResource is the Schema for the AIPlatformPersistentResource API
// +k8s:openapi-gen=true
type AIPlatformPersistentResource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AIPlatformPersistentResourceSpec   `json:"spec,omitempty"`
	Status AIPlatformPersistentResourceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AIPlatformPersistentResourceList contains a list of AIPlatformPersistentResource
type AIPlatformPersistentResourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AIPlatformPersistentResource `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AIPlatformPersistentResource{}, &AIPlatformPersistentResourceList{})
}
