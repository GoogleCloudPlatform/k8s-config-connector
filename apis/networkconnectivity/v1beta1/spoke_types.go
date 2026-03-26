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

package v1beta1

import (
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var NetworkConnectivitySpokeGVK = GroupVersion.WithKind("NetworkConnectivitySpoke")

type SpokeHubRef struct {
	/* Immutable. The URI of the hub that this spoke is attached to. Allowed value: The Google Cloud resource name of a `NetworkConnectivityHub` resource (format: `projects/{{project}}/locations/global/hubs/{{name}}`). */
	External string `json:"external,omitempty"`
	/* Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names */
	Name string `json:"name,omitempty"`
	/* Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ */
	Namespace string `json:"namespace,omitempty"`
}

// +kcc:proto=google.cloud.networkconnectivity.v1.LinkedVpnTunnels
type LinkedVPNTunnels struct {
	/* Immutable. A value that controls whether site-to-site data transfer is enabled for these resources. Note that data transfer is available only in supported locations. */
	// +required
	SiteToSiteDataTransfer *bool `json:"siteToSiteDataTransfer,omitempty"`
	/* Immutable. */
	Uris []computev1beta1.ComputeVPNTunnelRef `json:"uris"`
}

// +kcc:proto=google.cloud.networkconnectivity.v1.LinkedInterconnectAttachments
type LinkedInterconnectAttachments struct {
	/* Immutable. A collection of VLAN attachment resources. These resources should be redundant attachments that all advertise the same prefixes to Google Cloud. Alternatively, in active/passive configurations, all attachments should be capable of advertising the same prefixes. */
	// +required
	SiteToSiteDataTransfer *bool `json:"siteToSiteDataTransfer,omitempty"`
	/* Immutable. */
	Uris []computev1beta1.ComputeInterconnectAttachmentRef `json:"uris"`
}

// +kcc:proto=google.cloud.networkconnectivity.v1.RouterApplianceInstance
type RouterApplianceInstance struct {
	/* Immutable. The IP address on the VM to use for peering. */
	IpAddress *string `json:"ipAddress,omitempty"`
	/* Immutable. */
	VirtualMachineRef *computev1beta1.InstanceRef `json:"virtualMachineRef,omitempty"`
}

// +kcc:proto=google.cloud.networkconnectivity.v1.LinkedRouterApplianceInstances
type LinkedRouterApplianceInstances struct {
	/* Immutable. The list of router appliance instances */
	Instances []RouterApplianceInstance `json:"instances"`
	/* Immutable. A value that controls whether site-to-site data transfer is enabled for these resources. Note that data transfer is available only in supported locations. */
	// +required
	SiteToSiteDataTransfer *bool `json:"siteToSiteDataTransfer,omitempty"`
}

// +kcc:proto=google.cloud.networkconnectivity.v1.LinkedVpcNetwork
type LinkedVPCNetwork struct {
	/* Immutable. IP ranges encompassing the subnets to be excluded from peering. */
	ExcludeExportRanges []string `json:"excludeExportRanges,omitempty"`
	/* Immutable. */
	UriRef computev1beta1.ComputeNetworkRef `json:"uriRef"`
}

// NetworkConnectivitySpokeSpec defines the desired state of NetworkConnectivitySpoke
// +kcc:spec:proto=google.cloud.networkconnectivity.v1.Spoke
type NetworkConnectivitySpokeSpec struct {
	/* Immutable. An optional description of the spoke. */
	Description *string `json:"description,omitempty"`

	/* Immutable. */
	// +required
	HubRef *SpokeHubRef `json:"hubRef"`

	/* Immutable. A collection of VLAN attachment resources. These resources should be redundant attachments that all advertise the same prefixes to Google Cloud. Alternatively, in active/passive configurations, all attachments should be capable of advertising the same prefixes. */
	LinkedInterconnectAttachments *LinkedInterconnectAttachments `json:"linkedInterconnectAttachments,omitempty"`

	/* Immutable. The URIs of linked Router appliance resources */
	LinkedRouterApplianceInstances *LinkedRouterApplianceInstances `json:"linkedRouterApplianceInstances,omitempty"`

	/* Immutable. VPC network that is associated with the spoke. */
	LinkedVPCNetwork *LinkedVPCNetwork `json:"linkedVPCNetwork,omitempty"`

	/* Immutable. The URIs of linked VPN tunnel resources */
	LinkedVpnTunnels *LinkedVPNTunnels `json:"linkedVpnTunnels,omitempty"`

	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +required
	Location string `json:"location"`

	// The NetworkConnectivitySpoke name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
}

// NetworkConnectivitySpokeStatus defines the config connector machine state of NetworkConnectivitySpoke
type NetworkConnectivitySpokeStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* Output only. The time the spoke was created. */
	// +kubebuilder:validation:Format=date-time
	CreateTime *string `json:"createTime,omitempty"`

	/* Output only. The current lifecycle state of this spoke. Possible values: STATE_UNSPECIFIED, CREATING, ACTIVE, DELETING */
	State *string `json:"state,omitempty"`

	/* Output only. The Google-generated UUID for the spoke. This value is unique across all spoke resources. If a spoke is deleted and another with the same name is created, the new spoke is assigned a different unique_id. */
	UniqueId *string `json:"uniqueId,omitempty"`

	/* Output only. The time the spoke was last updated. */
	// +kubebuilder:validation:Format=date-time
	UpdateTime *string `json:"updateTime,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *SpokeObservedState `json:"observedState,omitempty"`

	// A unique specifier for the NetworkConnectivitySpoke resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetworkconnectivityspoke;gcpnetworkconnectivityspokes
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/dcl2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetworkConnectivitySpoke is the Schema for the NetworkConnectivitySpoke API
// +k8s:openapi-gen=true
type NetworkConnectivitySpoke struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NetworkConnectivitySpokeSpec   `json:"spec,omitempty"`
	Status NetworkConnectivitySpokeStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NetworkConnectivitySpokeList contains a list of NetworkConnectivitySpoke
type NetworkConnectivitySpokeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkConnectivitySpoke `json:"items"`
}

// SpokeObservedState is the state of the Spoke resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.networkconnectivity.v1.Spoke
type SpokeObservedState struct {
	// Output only. The time the spoke was created.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Spoke.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time the spoke was last updated.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Spoke.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The Google-generated UUID for the spoke. This value is unique
	//  across all spoke resources. If a spoke is deleted and another with the same
	//  name is created, the new spoke is assigned a different `unique_id`.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Spoke.unique_id
	UniqueId *string `json:"uniqueId,omitempty"`

	// Output only. The current lifecycle state of this spoke.
	// +kcc:proto:field=google.cloud.networkconnectivity.v1.Spoke.state
	State *string `json:"state,omitempty"`
}

func init() {
	SchemeBuilder.Register(&NetworkConnectivitySpoke{}, &NetworkConnectivitySpokeList{})
}
