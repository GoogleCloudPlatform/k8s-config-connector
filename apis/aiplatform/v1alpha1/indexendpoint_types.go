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
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var VertexAIIndexEndpointGVK = GroupVersion.WithKind("VertexAIIndexEndpoint")

// VertexAIIndexEndpointSpec defines the desired state of VertexAIIndexEndpoint
// +kcc:spec:proto=google.cloud.aiplatform.v1.IndexEndpoint
type VertexAIIndexEndpointSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The VertexAIIndexEndpoint name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The display name of the IndexEndpoint.
	// The name can be up to 128 characters long and can consist of any UTF-8
	// characters.
	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName"`

	// The description of the IndexEndpoint.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty"`

	// The labels with user-defined metadata to organize your IndexEndpoints.
	// Label keys and values can be no longer than 64 characters
	// (Unicode codepoints), can only contain lowercase letters, numeric
	// characters, underscores and dashes. International characters are allowed.
	// +kubebuilder:validation:Optional
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. The full name of the Google Compute Engine
	// [network](https://cloud.google.com/compute/docs/networks-and-firewalls#networks)
	// to which the IndexEndpoint should be peered.
	// +kubebuilder:validation:Optional
	NetworkRef *computev1beta1.ComputeNetworkRef `json:"networkRef,omitempty"`

	// Optional. Configuration for private service connect.
	// +kubebuilder:validation:Optional
	PrivateServiceConnectConfig *PrivateServiceConnectConfig `json:"privateServiceConnectConfig,omitempty"`

	// Optional. If true, the deployed index will be accessible through public
	// endpoint.
	// +kubebuilder:validation:Optional
	PublicEndpointEnabled *bool `json:"publicEndpointEnabled,omitempty"`

	// Immutable. Customer-managed encryption key spec for an IndexEndpoint. If
	// set, this IndexEndpoint and all sub-resources of this IndexEndpoint will be
	// secured by this key.
	// +kubebuilder:validation:Optional
	EncryptionSpec *EncryptionSpec `json:"encryptionSpec,omitempty"`
}

// VertexAIIndexEndpointStatus defines the config connector machine state of VertexAIIndexEndpoint
type VertexAIIndexEndpointStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VertexAIIndexEndpoint resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VertexAIIndexEndpointObservedState `json:"observedState,omitempty"`
}

// VertexAIIndexEndpointObservedState is the state of the VertexAIIndexEndpoint resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.aiplatform.v1.IndexEndpoint
type VertexAIIndexEndpointObservedState struct {
	// Output only. The indexes deployed in this endpoint.
	DeployedIndexes []DeployedIndexObservedState `json:"deployedIndexes,omitempty"`

	// Output only. Timestamp when this IndexEndpoint was created.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this IndexEndpoint was last updated.
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. If
	// [public_endpoint_enabled][google.cloud.aiplatform.v1.IndexEndpoint.public_endpoint_enabled]
	// is true, this field will be populated with the domain name to use for this
	// index endpoint.
	PublicEndpointDomainName *string `json:"publicEndpointDomainName,omitempty"`

	// Output only. Reserved for future use.
	SatisfiesPZS *bool `json:"satisfiesPzs,omitempty"`

	// Output only. Reserved for future use.
	SatisfiesPZI *bool `json:"satisfiesPzi,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvertexaiindexendpoint;gcpvertexaiindexendpoints
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VertexAIIndexEndpoint is the Schema for the VertexAIIndexEndpoint API
// +k8s:openapi-gen=true
type VertexAIIndexEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VertexAIIndexEndpointSpec   `json:"spec,omitempty"`
	Status VertexAIIndexEndpointStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VertexAIIndexEndpointList contains a list of VertexAIIndexEndpoint
type VertexAIIndexEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VertexAIIndexEndpoint `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VertexAIIndexEndpoint{}, &VertexAIIndexEndpointList{})
}
