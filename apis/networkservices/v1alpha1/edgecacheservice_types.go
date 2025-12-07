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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var NetworkServicesEdgeCacheServiceGVK = GroupVersion.WithKind("NetworkServicesEdgeCacheService")

// NetworkServicesEdgeCacheServiceSpec defines the desired state of NetworkServicesEdgeCacheService
// +kcc:spec:proto=mockgcp.cloud.edgecacheservice.v1.EdgeCacheService
type NetworkServicesEdgeCacheServiceSpec struct {
	// Required. Defines the parent path of the resource.
	*parent.ProjectAndLocationRef `json:",inline"`

	// The NetworkServicesEdgeCacheService name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.EdgeCacheService.labels
	Labels map[string]string `json:"labels,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.EdgeCacheService.description
	Description *string `json:"description,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.EdgeCacheService.routing
	Routing *Routing `json:"routing,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.EdgeCacheService.require_tls
	RequireTLS *bool `json:"requireTLS,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.EdgeCacheService.edge_ssl_certificates
	EdgeSSLCertificates []string `json:"edgeSSLCertificates,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.EdgeCacheService.edge_security_policy
	EdgeSecurityPolicy *string `json:"edgeSecurityPolicy,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.EdgeCacheService.log_config
	LogConfig *LogConfig `json:"logConfig,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.EdgeCacheService.disable_quic
	DisableQuic *bool `json:"disableQuic,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.EdgeCacheService.disable_http2
	DisableHttp2 *bool `json:"disableHttp2,omitempty"`
}

// NetworkServicesEdgeCacheServiceStatus defines the config connector machine state of NetworkServicesEdgeCacheService
type NetworkServicesEdgeCacheServiceStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NetworkServicesEdgeCacheService resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *EdgeCacheServiceObservedState `json:"observedState,omitempty"`
}

// +kcc:observedstate:proto=mockgcp.cloud.edgecacheservice.v1.EdgeCacheService
type EdgeCacheServiceObservedState struct {
	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.EdgeCacheService.name
	Name *string `json:"name,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.EdgeCacheService.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.EdgeCacheService.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.EdgeCacheService.ipv4_addresses
	IPv4Addresses []string `json:"ipv4Addresses,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.EdgeCacheService.ipv6_addresses
	IPv6Addresses []string `json:"ipv6Addresses,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetworkservicesedgecacheservice;gcpnetworkservicesedgecacheservices
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetworkServicesEdgeCacheService is the Schema for the NetworkServicesEdgeCacheService API
// +k8s:openapi-gen=true
type NetworkServicesEdgeCacheService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NetworkServicesEdgeCacheServiceSpec   `json:"spec,omitempty"`
	Status NetworkServicesEdgeCacheServiceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NetworkServicesEdgeCacheServiceList contains a list of NetworkServicesEdgeCacheService
type NetworkServicesEdgeCacheServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkServicesEdgeCacheService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkServicesEdgeCacheService{}, &NetworkServicesEdgeCacheServiceList{})
}
