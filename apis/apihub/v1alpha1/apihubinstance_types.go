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
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var APIHubInstanceGVK = GroupVersion.WithKind("APIHubInstance")

// APIHubInstanceSpec defines the desired state of APIHubInstance
// +kcc:spec:proto=google.cloud.apihub.v1.ApiHubInstance
type APIHubInstanceSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +kubebuilder:validation:Required
	Location *string `json:"location"`

	// The APIHubInstance name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. Config of the ApiHub instance.
	// +kubebuilder:validation:Required
	// +kcc:proto:field=google.cloud.apihub.v1.ApiHubInstance.config
	Config *APIHubInstance_Config `json:"config"`

	// Optional. Instance labels to represent user-provided metadata.
	//  Refer to cloud documentation on labels for more details.
	//  https://cloud.google.com/compute/docs/labeling-resources
	// +kubebuilder:validation:Optional
	// +kcc:proto:field=google.cloud.apihub.v1.ApiHubInstance.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Description of the ApiHub instance.
	// +kubebuilder:validation:Optional
	// +kcc:proto:field=google.cloud.apihub.v1.ApiHubInstance.description
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.cloud.apihub.v1.ApiHubInstance.Config
type APIHubInstance_Config struct {
	// Required. The Customer Managed Encryption Key (CMEK) used for data
	//  encryption. The CMEK name should follow the format of
	//  `projects/([^/]+)/locations/([^/]+)/keyRings/([^/]+)/cryptoKeys/([^/]+)`,
	//  where the location must match the instance location.
	// +kcc:proto:field=google.cloud.apihub.v1.ApiHubInstance.Config.cmek_key_name
	CmekKeyRef *kmsv1beta1.KMSCryptoKeyRef `json:"cmekKeyRef,omitempty"`

	// Optional. Disable ApiHub router.
	// +kcc:proto:field=google.cloud.apihub.v1.ApiHubInstance.Config.disable_search
	DisableSearch *bool `json:"disableSearch,omitempty"`

	// Optional. The Vertex AI location.
	// +kcc:proto:field=google.cloud.apihub.v1.ApiHubInstance.Config.vertex_location
	VertexLocation *string `json:"vertexLocation,omitempty"`

	// Optional. Encryption type for the region.
	// +kcc:proto:field=google.cloud.apihub.v1.ApiHubInstance.Config.encryption_type
	EncryptionType *string `json:"encryptionType,omitempty"`
}

// APIHubInstanceStatus defines the config connector machine state of APIHubInstance
type APIHubInstanceStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the APIHubInstance resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *APIHubInstanceObservedState `json:"observedState,omitempty"`
}

// APIHubInstanceObservedState is the state of the APIHubInstance resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.apihub.v1.ApiHubInstance
type APIHubInstanceObservedState struct {
	// Output only. Creation timestamp.
	// +kcc:proto:field=google.cloud.apihub.v1.ApiHubInstance.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Last update timestamp.
	// +kcc:proto:field=google.cloud.apihub.v1.ApiHubInstance.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The current state of the ApiHub instance.
	// +kcc:proto:field=google.cloud.apihub.v1.ApiHubInstance.state
	State *string `json:"state,omitempty"`

	// Output only. Extra information about ApiHub instance state. Currently the
	//  message would be populated when state is `FAILED`.
	// +kcc:proto:field=google.cloud.apihub.v1.ApiHubInstance.state_message
	StateMessage *string `json:"stateMessage,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpapihubinstance;gcpapihubinstances
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// APIHubInstance is the Schema for the APIHubInstance API
// +k8s:openapi-gen=true
type APIHubInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   APIHubInstanceSpec   `json:"spec,omitempty"`
	Status APIHubInstanceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// APIHubInstanceList contains a list of APIHubInstance
type APIHubInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []APIHubInstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&APIHubInstance{}, &APIHubInstanceList{})
}
