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
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var MemorystoreInstanceEndpointGVK = GroupVersion.WithKind("MemorystoreInstanceEndpoint")

// MemorystoreInstanceEndpointSpec defines the desired state of MemorystoreInstanceEndpoint
// +kcc:spec:proto=google.cloud.memorystore.v1.Instance
type MemorystoreInstanceEndpointSpec struct {
	// Required. The Memorystore instance reference of the endpoint.
	InstanceRef *refsv1beta1.MemorystoreInstanceRef `json:"instanceRef"`

	// The MemorystoreInstanceEndpoint name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Endpoints for the instance.
	// +kcc:proto:field=google.cloud.memorystore.v1.Instance.endpoints
	Endpoints []Endpoint `json:"endpoints,omitempty"`
}

// MemorystoreInstanceEndpointStatus defines the config connector machine state of MemorystoreInstanceEndpoint
type MemorystoreInstanceEndpointStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *MemorystoreInstanceEndpointObservedState `json:"observedState,omitempty"`
}

// MemorystoreInstanceEndpointObservedState is the state of the MemorystoreInstanceEndpoint resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.memorystore.v1.Instance
type MemorystoreInstanceEndpointObservedState struct {
	// Optional. Endpoints for the instance.
	// +kcc:proto:field=google.cloud.memorystore.v1.Instance.endpoints
	Endpoints []EndpointObservedState `json:"endpoints,omitempty"`
}

// +kcc:proto=google.cloud.memorystore.v1.Instance.InstanceEndpoint
type Endpoint struct {
	// Optional. A group of PSC connections. They are created in the same VPC
	//  network, one for each service attachment in the cluster.
	// +kcc:proto:field=google.cloud.memorystore.v1.Instance.InstanceEndpoint.connections
	Connections []Endpoint_ConnectionDetail `json:"connections,omitempty"`
}

// +kcc:proto=google.cloud.memorystore.v1.Instance.ConnectionDetail
type Endpoint_ConnectionDetail struct {
	// Detailed information of a PSC connection that is created through
	//  service connectivity automation.
	// +kcc:proto:field=google.cloud.memorystore.v1.Instance.ConnectionDetail.psc_connection
	PscConnection *PscConnection `json:"pscConnection,omitempty"`
}

// +kcc:proto=google.cloud.memorystore.v1.Instance.PscConnection
type PscConnection struct {
	// The PSC connection id.
	// +kcc:proto:field=google.cloud.memorystore.v1.Instance.PscConnection.forwarding_rule
	ForwardingRuleRef *computev1beta1.ForwardingRuleRef `json:"forwardingRuleRef,omitempty"`
	// Optional. The port number of the PSC connection.
	//  Port will only be set for Primary/Reader or Discovery endpoint.
	// +kcc:proto:field=google.cloud.memorystore.v1.PscConnection.port
	Port *int32 `json:"port,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.memorystore.v1.Instance.InstanceEndpoint
type EndpointObservedState struct {
	// Optional. A group of PSC connections. They are created in the same VPC
	//  network, one for each service attachment in the cluster.
	// +kcc:proto:field=google.cloud.memorystore.v1.Instance.InstanceEndpoint.connections
	Connections []Endpoint_ConnectionDetailObservedState `json:"connections,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.memorystore.v1.Instance.ConnectionDetail
type Endpoint_ConnectionDetailObservedState struct {
	// Detailed information of a PSC connection that is created through
	//  service connectivity automation.
	// +kcc:proto:field=google.cloud.memorystore.v1.Instance.ConnectionDetail.psc_connection
	PscConnection *PscConnectionObservedState `json:"pscConnection,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.memorystore.v1.Instance.PscConnection
type PscConnectionObservedState struct {
	// Output only. The consumer project_id where the forwarding rule is created from.
	// +kcc:proto:field=google.cloud.memorystore.v1.PscConnection.project_id
	ProjectID *string `json:"projectID,omitempty"`

	// Output only. The status of the PSC connection: whether a connection exists
	//  and ACTIVE or it no longer exists. Please note that this value is updated
	//  periodically. Please use Private Service Connect APIs for the latest
	//  status. For valid values,
	//  see https://docs.cloud.google.com/memorystore/docs/valkey/reference/rest/v1/projects.locations.instances#pscconnectionstatus
	// +kcc:proto:field=google.cloud.memorystore.v1.PscAutoConnection.psc_connection_status
	PscConnectionStatus *string `json:"pscConnectionStatus,omitempty"`

	// Output only. Type of the PSC connection. For valid values,
	//  see https://docs.cloud.google.com/memorystore/docs/valkey/reference/rest/v1/projects.locations.instances#connectiontype
	// +kcc:proto:field=google.cloud.memorystore.v1.PscAutoConnection.connection_type
	ConnectionType *string `json:"connectionType,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpmemorystoreinstanceendpoint;gcpmemorystoreinstanceendpoints
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// MemorystoreInstanceEndpoint is the Schema for the MemorystoreInstanceEndpoint API
// +k8s:openapi-gen=true
type MemorystoreInstanceEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   MemorystoreInstanceEndpointSpec   `json:"spec,omitempty"`
	Status MemorystoreInstanceEndpointStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// MemorystoreInstanceEndpointList contains a list of MemorystoreInstanceEndpoint
type MemorystoreInstanceEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MemorystoreInstanceEndpoint `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MemorystoreInstanceEndpoint{}, &MemorystoreInstanceEndpointList{})
}
