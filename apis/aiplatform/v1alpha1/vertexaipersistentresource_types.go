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

	// Optional. The display name of the PersistentResource.
	//  The name can be up to 128 characters long and can consist of any UTF-8
	//  characters.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PersistentResource.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. The spec of the pools of different resources.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PersistentResource.resource_pools
	ResourcePools []ResourcePool `json:"resourcePools,omitempty"`

	// Optional. The labels with user-defined metadata to organize
	//  PersistentResource.
	//
	//  Label keys and values can be no longer than 64 characters
	//  (Unicode codepoints), can only contain lowercase letters, numeric
	//  characters, underscores and dashes. International characters are allowed.
	//
	//  See https://goo.gl/xmQnxf for more information and examples of labels.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PersistentResource.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. The full name of the Compute Engine network to peer with Vertex AI to host the persistent resources.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PersistentResource.network
	NetworkRef *computev1beta1.ComputeNetworkRef `json:"networkRef,omitempty"`

	// Optional. Configuration for PSC-I for PersistentResource.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PersistentResource.psc_interface_config
	PSCInterfaceConfig *PSCInterfaceConfig `json:"pscInterfaceConfig,omitempty"`

	// Optional. Customer-managed encryption key spec for a PersistentResource.
	//  If set, this PersistentResource and all sub-resources of this
	//  PersistentResource will be secured by this key.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PersistentResource.encryption_spec
	EncryptionSpec *EncryptionSpec `json:"encryptionSpec,omitempty"`

	// Optional. Persistent Resource runtime spec.
	//  For example, used for Ray cluster configuration.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PersistentResource.resource_runtime_spec
	ResourceRuntimeSpec *ResourceRuntimeSpec `json:"resourceRuntimeSpec,omitempty"`

	// Optional. A list of names for the reserved IP ranges under the VPC network
	//  that can be used for this persistent resource.
	//
	//  If set, we will deploy the persistent resource within the provided IP
	//  ranges. Otherwise, the persistent resource is deployed to any IP
	//  ranges under the provided VPC network.
	//
	//  Example: ['vertex-ai-ip-range'].
	// +kcc:proto:field=google.cloud.aiplatform.v1.PersistentResource.reserved_ip_ranges
	ReservedIPRanges []string `json:"reservedIPRanges,omitempty"`
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
	// Output only. Resource pools inside the PersistentResource.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PersistentResource.resource_pools
	ResourcePools []ResourcePoolObservedState `json:"resourcePools,omitempty"`

	// Output only. Only populated when state is `FREE`.
	// +kcc:proto:field=google.cloud.aiplatform.v1.PersistentResource.state
	State *string `json:"state,omitempty"`

	// Output only. Only populated when state is `ERROR`.
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

	// Output only. Runtime information of the PersistentResource.
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

// +kcc:proto=google.cloud.aiplatform.v1.ServiceAccountSpec
type ServiceAccountSpec struct {
	// Required. If true, custom user-managed service account is enforced to run
	//  any workloads (for example, Vertex Jobs) on the resource. Otherwise, uses
	//  the [Vertex AI Custom Code Service
	//  Agent](https://cloud.google.com/vertex-ai/docs/general/access-control#service-agents).
	// +kcc:proto:field=google.cloud.aiplatform.v1.ServiceAccountSpec.enable_custom_service_account
	EnableCustomServiceAccount *bool `json:"enableCustomServiceAccount,omitempty"`

	// Optional. Required when all below conditions are met
	//   * `enable_custom_service_account` is true;
	//   * any runtime is specified via `ResourceRuntimeSpec` on creation time,
	//     for example, Ray
	//
	//  The users must have `iam.serviceAccounts.actAs` permission on this service
	//  account and then the specified runtime containers will run as it.
	//
	//  Do not set this field if you want to submit jobs using custom service
	//  account to this PersistentResource after creation, but only specify the
	//  `service_account` inside the job.
	// +kcc:proto:field=google.cloud.aiplatform.v1.ServiceAccountSpec.service_account
	ServiceAccountRef *refsv1beta1.IAMServiceAccountRef `json:"serviceAccountRef,omitempty"`
}

func init() {
	SchemeBuilder.Register(&VertexAIPersistentResource{}, &VertexAIPersistentResourceList{})
}
