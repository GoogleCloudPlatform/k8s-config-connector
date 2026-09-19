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

var NetworkConnectivityServiceConnectionMapGVK = GroupVersion.WithKind("NetworkConnectivityServiceConnectionMap")

// NetworkConnectivityServiceConnectionMapSpec defines the desired state of NetworkConnectivityServiceConnectionMap
// +kcc:spec:proto=mockgcp.cloud.networkconnectivity.v1.ServiceConnectionMap
type NetworkConnectivityServiceConnectionMapSpec struct {
	// The project that this resource belongs to.
	ProjectRef refsv1beta1.ProjectRef `json:"projectRef"`

	// Immutable. The location of this resource.
	Location *string `json:"location"`

	// The NetworkConnectivityServiceConnectionMap name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// A description of this resource.
	// +optional
	Description *string `json:"description,omitempty"`

	// User-defined labels.
	// +optional
	Labels map[string]string `json:"labels,omitempty"`

	// The service class identifier this ServiceConnectionMap is for. The user of ServiceConnectionMap create API needs to have networkconnectivity.serviceClasses.use IAM permission for the service class.
	// +optional
	ServiceClass *string `json:"serviceClass,omitempty"`

	// The token provided by the consumer. This token authenticates that the consumer can create a connection within the specified project and network.
	// +optional
	Token *string `json:"token,omitempty"`

	// The PSC configurations on consumer side.
	// +optional
	ConsumerPSCConfigs []ConsumerPSCConfig `json:"consumerPSCConfigs,omitempty"`

	// The PSC configurations on producer side.
	// +optional
	ProducerPSCConfigs []ProducerPSCConfig `json:"producerPSCConfigs,omitempty"`
}

// NetworkConnectivityServiceConnectionMapStatus defines the config connector machine state of NetworkConnectivityServiceConnectionMap
type NetworkConnectivityServiceConnectionMapStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NetworkConnectivityServiceConnectionMap resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *NetworkConnectivityServiceConnectionMapObservedState `json:"observedState,omitempty"`
}

// NetworkConnectivityServiceConnectionMapObservedState is the state of the NetworkConnectivityServiceConnectionMap resource as most recently observed in GCP.
// +kcc:observedstate:proto=mockgcp.cloud.networkconnectivity.v1.ServiceConnectionMap
type NetworkConnectivityServiceConnectionMapObservedState struct {
	// Output only. PSC connection details on consumer side.
	// +optional
	ConsumerPSCConnections []ConsumerPSCConnection `json:"consumerPSCConnections,omitempty"`

	// Output only. Time when the ServiceConnectionMap was created.
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	// Optional. The etag is computed by the server, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
	// +optional
	Etag *string `json:"etag,omitempty"`

	// Output only. The infrastructure used for connections between consumers/producers.
	// +optional
	Infrastructure *string `json:"infrastructure,omitempty"`

	// Output only. The service class uri this ServiceConnectionMap is for.
	// +optional
	ServiceClassURI *string `json:"serviceClassURI,omitempty"`

	// Output only. Time when the ServiceConnectionMap was updated.
	// +optional
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetworkconnectivityserviceconnectionmap;gcpnetworkconnectivityserviceconnectionmaps
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetworkConnectivityServiceConnectionMap is the Schema for the NetworkConnectivityServiceConnectionMap API
// +k8s:openapi-gen=true
type NetworkConnectivityServiceConnectionMap struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NetworkConnectivityServiceConnectionMapSpec   `json:"spec,omitempty"`
	Status NetworkConnectivityServiceConnectionMapStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NetworkConnectivityServiceConnectionMapList contains a list of NetworkConnectivityServiceConnectionMap
type NetworkConnectivityServiceConnectionMapList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkConnectivityServiceConnectionMap `json:"items"`
}

// +kcc:proto=mockgcp.cloud.networkconnectivity.v1.ConsumerPscConfig
type ConsumerPSCConfig struct {
	// Required. The project ID or project number of the consumer project. This project is the one that the consumer uses to interact with the producer instance. From the perspective of a consumer who's created a producer instance, this is the project of the producer instance. Format: 'projects/' Eg. 'projects/consumer-project' or 'projects/1234'
	// +kcc:proto:field=mockgcp.cloud.networkconnectivity.v1.ConsumerPscConfig.consumer_instance_project
	ConsumerInstanceProject *string `json:"consumerInstanceProject,omitempty"`

	// This is used in PSC consumer ForwardingRule to control whether the PSC endpoint can be accessed from another region.
	// +kcc:proto:field=mockgcp.cloud.networkconnectivity.v1.ConsumerPscConfig.disable_global_access
	DisableGlobalAccess *bool `json:"disableGlobalAccess,omitempty"`

	// The requested IP version for the PSC connection.
	// +kcc:proto:field=mockgcp.cloud.networkconnectivity.v1.ConsumerPscConfig.ip_version
	IPVersion *string `json:"ipVersion,omitempty"`

	// The resource path of the consumer network where PSC connections are allowed to be created in. Note, this network does not need be in the ConsumerPscConfig.project in the case of SharedVPC. Example: projects/{projectNumOrId}/global/networks/{networkId}.
	// +kcc:proto:field=mockgcp.cloud.networkconnectivity.v1.ConsumerPscConfig.network
	NetworkRef *computerefs.ComputeNetworkRef `json:"networkRef,omitempty"`

	// Immutable. Deprecated. Use producer_instance_metadata instead. An immutable identifier for the producer instance.
	// +kcc:proto:field=mockgcp.cloud.networkconnectivity.v1.ConsumerPscConfig.producer_instance_id
	ProducerInstanceID *string `json:"producerInstanceID,omitempty"`

	// Immutable. An immutable map for the producer instance metadata.
	// +kcc:proto:field=mockgcp.cloud.networkconnectivity.v1.ConsumerPscConfig.producer_instance_metadata
	ProducerInstanceMetadata map[string]string `json:"producerInstanceMetadata,omitempty"`

	// The consumer project where PSC connections are allowed to be created in.
	// +kcc:proto:field=mockgcp.cloud.networkconnectivity.v1.ConsumerPscConfig.project
	Project *string `json:"project,omitempty"`

	// Optional. A map to store mapping between customer vip and target service attachment. This field can be used to specify a static IP address for a PSC connection.
	// +kcc:proto:field=mockgcp.cloud.networkconnectivity.v1.ConsumerPscConfig.service_attachment_ip_address_map
	ServiceAttachmentIPAddressMap map[string]string `json:"serviceAttachmentIPAddressMap,omitempty"`

	// Output only. Overall state of PSC Connections management for this consumer psc config.
	// +kcc:proto:field=mockgcp.cloud.networkconnectivity.v1.ConsumerPscConfig.state
	State *string `json:"state,omitempty"`
}

// +kcc:proto=mockgcp.cloud.networkconnectivity.v1.ProducerPscConfig
type ProducerPSCConfig struct {
	// Optional. The specification for automatically creating a DNS record for this PSC connection.
	// +kcc:proto:field=mockgcp.cloud.networkconnectivity.v1.ProducerPscConfig.automated_dns_creation_spec
	AutomatedDNSCreationSpec *AutomatedDNSCreationSpec `json:"automatedDNSCreationSpec,omitempty"`

	// The resource path of a service attachment. Example: projects/{projectNumOrId}/regions/{region}/serviceAttachments/{resourceId}.
	// +kcc:proto:field=mockgcp.cloud.networkconnectivity.v1.ProducerPscConfig.service_attachment_uri
	ServiceAttachmentRef *refsv1beta1.ComputeServiceAttachmentRef `json:"serviceAttachmentRef,omitempty"`
}

func init() {
	SchemeBuilder.Register(&NetworkConnectivityServiceConnectionMap{}, &NetworkConnectivityServiceConnectionMapList{})
}
