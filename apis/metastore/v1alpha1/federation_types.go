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
	refv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var MetastoreFederationGVK = GroupVersion.WithKind("MetastoreFederation")

type Parent struct {
	// +required
	Location string `json:"location"`

	// +optional
	ProjectRef *refv1beta1.ProjectRef `json:"projectRef,omitempty"`
}

// MetastoreFederationSpec defines the desired state of MetastoreFederation
// +kcc:proto=google.cloud.metastore.v1.Federation
type MetastoreFederationSpec struct {
	Parent Parent `json:",inline"`
	// The MetastoreFederation name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Immutable. The relative resource name of the federation, of the
	//  form:
	//  projects/{project_number}/locations/{location_id}/federations/{federation_id}`.
	// +kcc:proto:field=google.cloud.metastore.v1.Federation.name
	Name *string `json:"name,omitempty"`

	// User-defined labels for the metastore federation.
	// +kcc:proto:field=google.cloud.metastore.v1.Federation.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Immutable. The Apache Hive metastore version of the federation. All backend
	//  metastore versions must be compatible with the federation version.
	// +kcc:proto:field=google.cloud.metastore.v1.Federation.version
	Version *string `json:"version,omitempty"`

	// TODO: unsupported map type with key int32 and value message
}

// MetastoreFederationStatus defines the config connector machine state of MetastoreFederation
type MetastoreFederationStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the MetastoreFederation resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *MetastoreFederationObservedState `json:"observedState,omitempty"`
}

// MetastoreFederationObservedState is the state of the MetastoreFederation resource as most recently observed in GCP.
// +kcc:proto=google.cloud.metastore.v1.Federation
type MetastoreFederationObservedState struct {
	// Output only. The time when the metastore federation was created.
	// +kcc:proto:field=google.cloud.metastore.v1.Federation.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the metastore federation was last updated.
	// +kcc:proto:field=google.cloud.metastore.v1.Federation.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The federation endpoint.
	// +kcc:proto:field=google.cloud.metastore.v1.Federation.endpoint_uri
	EndpointURI *string `json:"endpointURI,omitempty"`

	// Output only. The current state of the federation.
	// +kcc:proto:field=google.cloud.metastore.v1.Federation.state
	State *string `json:"state,omitempty"`

	// Output only. Additional information about the current state of the
	//  metastore federation, if available.
	// +kcc:proto:field=google.cloud.metastore.v1.Federation.state_message
	StateMessage *string `json:"stateMessage,omitempty"`

	// Output only. The globally unique resource identifier of the metastore
	//  federation.
	// +kcc:proto:field=google.cloud.metastore.v1.Federation.uid
	Uid *string `json:"uid,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TODO(user): make sure the pluralizaiton below is correct
// +kubebuilder:resource:categories=gcp,shortName=gcpmetastorefederation;gcpmetastorefederations
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// MetastoreFederation is the Schema for the MetastoreFederation API
// +k8s:openapi-gen=true
type MetastoreFederation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   MetastoreFederationSpec   `json:"spec,omitempty"`
	Status MetastoreFederationStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// MetastoreFederationList contains a list of MetastoreFederation
type MetastoreFederationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MetastoreFederation `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MetastoreFederation{}, &MetastoreFederationList{})
}
