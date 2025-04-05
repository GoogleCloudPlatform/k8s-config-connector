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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var NetworkConnectivityRegionalEndpointGVK = GroupVersion.WithKind("NetworkConnectivityRegionalEndpoint")

// NetworkConnectivityRegionalEndpointSpec defines the desired state of NetworkConnectivityRegionalEndpoint
// +kcc:proto=mockgcp.cloud.networkconnectivity.v1.RegionalEndpoint
type NetworkConnectivityRegionalEndpointSpec struct {
	// The NetworkConnectivityRegionalEndpoint name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. Defines the parent path of the resource.
	*Parent `json:",inline"`

	// Optional. A description of this resource.
	// +kcc:proto:field=mockgcp.cloud.networkconnectivity.v1.RegionalEndpoint.description
	Description *string `json:"description,omitempty"`

	// Required. The access type of this regional endpoint. This field is reflected in the PSC Forwarding Rule configuration to enable global access.
	// +kcc:proto:field=mockgcp.cloud.networkconnectivity.v1.RegionalEndpoint.access_type
	AccessType *string `json:"accessType,omitempty"`

	// Optional. The IP Address of the Regional Endpoint. When no address is provided, an IP from the subnetwork is allocated. Use one of the following formats: * IPv4 address as in `10.0.0.1` * Address resource URI as in `projects/{project}/regions/{region}/addresses/{address_name}` for an IPv4 or IPv6 address.
	// +kcc:proto:field=mockgcp.cloud.networkconnectivity.v1.RegionalEndpoint.address
	Address *string `json:"address,omitempty"`

	// User-defined labels.
	// +kcc:proto:field=mockgcp.cloud.networkconnectivity.v1.RegionalEndpoint.labels
	Labels map[string]string `json:"labels,omitempty"`

	// The name of the VPC network for this private regional endpoint. Format: `projects/{project}/global/networks/{network}`
	// +kcc:proto:field=mockgcp.cloud.networkconnectivity.v1.RegionalEndpoint.network
	Network *string `json:"network,omitempty"`

	// The name of the subnetwork from which the IP address will be allocated. Format: `projects/{project}/regions/{region}/subnetworks/{subnetwork}`
	// +kcc:proto:field=mockgcp.cloud.networkconnectivity.v1.RegionalEndpoint.subnetwork
	Subnetwork *string `json:"subnetwork,omitempty"`

	// Required. The service endpoint this private regional endpoint connects to. Format: `{apiname}.{region}.p.rep.googleapis.com` Example: "cloudkms.us-central1.p.rep.googleapis.com".
	// +kcc:proto:field=mockgcp.cloud.networkconnectivity.v1.RegionalEndpoint.target_google_api
	TargetGoogleAPI *string `json:"targetGoogleAPI,omitempty"`
}

// NetworkConnectivityRegionalEndpointStatus defines the config connector machine state of NetworkConnectivityRegionalEndpoint
type NetworkConnectivityRegionalEndpointStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NetworkConnectivityRegionalEndpoint resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *NetworkConnectivityRegionalEndpointObservedState `json:"observedState,omitempty"`
}

// NetworkConnectivityRegionalEndpointObservedState is the state of the NetworkConnectivityRegionalEndpoint resource as most recently observed in GCP.
// +kcc:proto=mockgcp.cloud.networkconnectivity.v1.RegionalEndpoint
type NetworkConnectivityRegionalEndpointObservedState struct {
	// Output only. Time when the RegionalEndpoint was created.
	// +kcc:proto:field=mockgcp.cloud.networkconnectivity.v1.RegionalEndpoint.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The literal IP address of the PSC Forwarding Rule created on behalf of the customer. This field is deprecated. Use address instead.
	// +kcc:proto:field=mockgcp.cloud.networkconnectivity.v1.RegionalEndpoint.ip_address
	IPAddress *string `json:"ipAddress,omitempty"`

	// Output only. The resource reference of the PSC Forwarding Rule created on behalf of the customer. Format: `//compute.googleapis.com/projects/{project}/regions/{region}/forwardingRules/{forwarding_rule_name}`
	// +kcc:proto:field=mockgcp.cloud.networkconnectivity.v1.RegionalEndpoint.psc_forwarding_rule
	PSCForwardingRule *string `json:"pscForwardingRule,omitempty"`

	// Output only. Time when the RegionalEndpoint was updated.
	// +kcc:proto:field=mockgcp.cloud.networkconnectivity.v1.RegionalEndpoint.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetworkconnectivityregionalendpoint;gcpnetworkconnectivityregionalendpoints
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetworkConnectivityRegionalEndpoint is the Schema for the NetworkConnectivityRegionalEndpoint API
// +k8s:openapi-gen=true
type NetworkConnectivityRegionalEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NetworkConnectivityRegionalEndpointSpec   `json:"spec,omitempty"`
	Status NetworkConnectivityRegionalEndpointStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NetworkConnectivityRegionalEndpointList contains a list of NetworkConnectivityRegionalEndpoint
type NetworkConnectivityRegionalEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkConnectivityRegionalEndpoint `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkConnectivityRegionalEndpoint{}, &NetworkConnectivityRegionalEndpointList{})
}
