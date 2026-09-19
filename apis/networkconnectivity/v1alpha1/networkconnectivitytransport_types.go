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
	computerefs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/refs"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var NetworkConnectivityTransportGVK = GroupVersion.WithKind("NetworkConnectivityTransport")

// NetworkConnectivityTransportSpec defines the desired state of NetworkConnectivityTransport
// +kcc:spec:proto=mockgcp.cloud.networkconnectivity.v1.Transport
type NetworkConnectivityTransportSpec struct {
	// The project that this resource belongs to.
	// +kubebuilder:validation:Required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +kubebuilder:validation:Required
	Location *string `json:"location"`

	// The NetworkConnectivityTransport name. If not given, the metadata.name will be used.
	// +kubebuilder:validation:Optional
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. List of IP Prefixes that will be advertised to the remote provider. Both IPv4 and IPv6 addresses are supported.
	// +kubebuilder:validation:Optional
	AdvertisedRoutes []string `json:"advertisedRoutes,omitempty"`

	// Optional. Bandwidth of the Transport. This must be one of the supported bandwidths for the remote profile, and must be set when no activation key is being provided.
	// +kubebuilder:validation:Optional
	Bandwidth *string `json:"bandwidth,omitempty"`

	// Optional. Description of the Transport.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty"`

	// Optional. Labels as key value pairs.
	// +kubebuilder:validation:Optional
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Immutable. Resource URI of the Network that will be peered with this Transport. This field must be provided during resource creation and cannot be changed.
	// +kubebuilder:validation:Optional
	NetworkRef *computerefs.ComputeNetworkRef `json:"networkRef,omitempty"`

	// Optional. Immutable. Key used for establishing a connection with the remote transport. This key can only be provided if the profile supports an INPUT key flow and the resource is in the PENDING_KEY state.
	// +kubebuilder:validation:Optional
	ProvidedActivationKey *string `json:"providedActivationKey,omitempty"`

	// Optional. Immutable. The user supplied account id for the CSP associated with the remote profile.
	// +kubebuilder:validation:Optional
	RemoteAccountID *string `json:"remoteAccountID,omitempty"`

	// Optional. Immutable. Name of the remoteTransportProfile that this Transport is connecting to.
	// +kubebuilder:validation:Optional
	RemoteProfileRef *NetworkConnectivityRemoteTransportProfileRef `json:"remoteProfileRef,omitempty"`

	// Optional. IP version stack for the established connectivity.
	// +kubebuilder:validation:Optional
	StackType *string `json:"stackType,omitempty"`
}

// NetworkConnectivityTransportStatus defines the config connector machine state of NetworkConnectivityTransport
type NetworkConnectivityTransportStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NetworkConnectivityTransport resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *NetworkConnectivityTransportObservedState `json:"observedState,omitempty"`
}

// NetworkConnectivityTransportObservedState is the state of the NetworkConnectivityTransport resource as most recently observed in GCP.
// +kcc:observedstate:proto=mockgcp.cloud.networkconnectivity.v1.Transport
type NetworkConnectivityTransportObservedState struct {
	// Output only. Create time stamp.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Google-generated activation key. This is only output if the selected profile supports an OUTPUT key flow. Inputting this to the provider is only valid while the resource is in a PENDING_KEY state. Once the provider has accepted the key, the resource will move to the CONFIGURING state.
	GeneratedActivationKey *string `json:"generatedActivationKey,omitempty"`

	// Output only. The maximum transmission unit (MTU) of a packet that can be sent over this transport.
	MtuLimit *int32 `json:"mtuLimit,omitempty"`

	// Output only. VPC Network URI that was created for the VPC Peering connection to the provided `network`. If VPC Peering is disconnected, this can be used to re-establish.
	PeeringNetwork *string `json:"peeringNetwork,omitempty"`

	// Output only. State of the underlying connectivity.
	State *string `json:"state,omitempty"`

	// Output only. Update time stamp.
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetworkconnectivitytransport;gcpnetworkconnectivitytransports
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetworkConnectivityTransport is the Schema for the NetworkConnectivityTransport API
// +k8s:openapi-gen=true
type NetworkConnectivityTransport struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NetworkConnectivityTransportSpec   `json:"spec,omitempty"`
	Status NetworkConnectivityTransportStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NetworkConnectivityTransportList contains a list of NetworkConnectivityTransport
type NetworkConnectivityTransportList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkConnectivityTransport `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkConnectivityTransport{}, &NetworkConnectivityTransportList{})
}
