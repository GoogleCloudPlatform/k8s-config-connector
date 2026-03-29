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

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	ComputeTargetHTTPSProxyGVK = schema.GroupVersionKind{
		Group:   "compute.cnrm.cloud.google.com",
		Version: "v1beta1",
		Kind:    "ComputeTargetHTTPSProxy",
	}
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputetargethttpsproxy;gcpcomputetargethttpsproxies
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"
// +k8s:openapi-gen=true
type ComputeTargetHTTPSProxy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeTargetHTTPSProxySpec   `json:"spec,omitempty"`
	Status ComputeTargetHTTPSProxyStatus `json:"status,omitempty"`
}

// +kcc:spec:proto=google.cloud.compute.v1.TargetHttpsProxy
type ComputeTargetHTTPSProxySpec struct {
	// URLs to certificate manager certificate resources that are used to authenticate connections between users and the load balancer.
	// Currently, you may specify up to 15 certificates. Certificate manager certificates do not apply when the load balancing scheme is set to INTERNAL_SELF_MANAGED.
	// sslCertificates and certificateManagerCertificates fields cannot be defined together.
	// +kcc:proto:field=google.cloud.compute.v1.TargetHttpsProxy.ssl_certificates
	CertificateManagerCertificates []CertificateManagerCertificateRef `json:"certificateManagerCertificates,omitempty"`

	// A reference to the CertificateMap resource uri that identifies a
	// certificate map associated with the given target proxy. This field
	// can only be set for global target proxies. This field is only supported
	// for EXTERNAL and EXTERNAL_MANAGED load balancing schemes.
	// For INTERNAL_MANAGED, use certificateManagerCertificates instead.
	// sslCertificates and certificateMap fields cannot be defined together.
	// +kcc:proto:field=google.cloud.compute.v1.TargetHttpsProxy.certificate_map
	CertificateMapRef *CertificateManagerCertificateMapRef `json:"certificateMapRef,omitempty"`

	// Immutable. An optional description of this resource.
	// +kcc:proto:field=google.cloud.compute.v1.TargetHttpsProxy.description
	Description *string `json:"description,omitempty"`

	// Immutable. Specifies how long to keep a connection open, after completing a response,
	// while there is no matching traffic (in seconds). If an HTTP keepalive is
	// not specified, a default value (610 seconds) will be used. For Global
	// external HTTP(S) load balancer, the minimum allowed value is 5 seconds and
	// the maximum allowed value is 1200 seconds. For Global external HTTP(S)
	// load balancer (classic), this option is not available publicly.
	// +kcc:proto:field=google.cloud.compute.v1.TargetHttpsProxy.http_keep_alive_timeout_sec
	HttpKeepAliveTimeoutSec *int `json:"httpKeepAliveTimeoutSec,omitempty"`

	// Location represents the geographical location of the
	// ComputeTargetHTTPSProxy. Specify a region name or "global" for global
	// resources. Reference: GCP definition of regions/zones (https://cloud.google.com/compute/docs/regions-zones/)
	// +required
	Location *string `json:"location"`

	// Immutable. This field only applies when the forwarding rule that references
	// this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED.
	// +kcc:proto:field=google.cloud.compute.v1.TargetHttpsProxy.proxy_bind
	ProxyBind *bool `json:"proxyBind,omitempty"`

	// Specifies the QUIC override policy for this resource. This determines
	// whether the load balancer will attempt to negotiate QUIC with clients
	// or not. Can specify one of NONE, ENABLE, or DISABLE. If NONE is
	// specified, Google manages whether QUIC is used. Default value: "NONE" Possible values: ["NONE", "ENABLE", "DISABLE"].
	// +kcc:proto:field=google.cloud.compute.v1.TargetHttpsProxy.quic_override
	QuicOverride *string `json:"quicOverride,omitempty"`

	// Immutable. Optional. The name of the resource. Used for
	// creation and acquisition. When unset, the value of `metadata.name`
	// is used as the default.
	ResourceID *string `json:"resourceID,omitempty"`

	// Immutable. A URL referring to a networksecurity.ServerTlsPolicy
	// resource that describes how the proxy should authenticate inbound
	// traffic. serverTlsPolicy only applies to a global TargetHttpsProxy
	// attached to globalForwardingRules with the loadBalancingScheme
	// set to INTERNAL_SELF_MANAGED or EXTERNAL or EXTERNAL_MANAGED.
	// For details which ServerTlsPolicy resources are accepted with
	// INTERNAL_SELF_MANAGED and which with EXTERNAL, EXTERNAL_MANAGED
	// loadBalancingScheme consult ServerTlsPolicy documentation.
	// If left blank, communications are not encrypted.
	// +kcc:proto:field=google.cloud.compute.v1.TargetHttpsProxy.server_tls_policy
	ServerTlsPolicyRef *NetworkSecurityServerTLSPolicyRef `json:"serverTlsPolicyRef,omitempty"`

	// A list of ComputeSSLCertificate resources that are used to
	// authenticate connections between users and the load balancer. At
	// least one SSL certificate must be specified.
	// +kcc:proto:field=google.cloud.compute.v1.TargetHttpsProxy.ssl_certificates
	SslCertificates []ComputeSSLCertificateRef `json:"sslCertificates,omitempty"`

	// A reference to the ComputeSSLPolicy resource that will be
	// associated with the ComputeTargetHTTPSProxy resource. If not set,
	// the ComputeTargetHTTPSProxy resource will not have any SSL policy
	// configured.
	// +kcc:proto:field=google.cloud.compute.v1.TargetHttpsProxy.ssl_policy
	SslPolicyRef *ComputeSSLPolicyRef `json:"sslPolicyRef,omitempty"`

	// A reference to the ComputeURLMap resource that defines the mapping
	// from URL to the BackendService.
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.TargetHttpsProxy.url_map
	UrlMapRef *ComputeURLMapRef `json:"urlMapRef"`
}

// ComputeTargetHTTPSProxyStatus defines the config connector machine state of ComputeTargetHTTPSProxy
type ComputeTargetHTTPSProxyStatus struct {
	// Conditions represent the latest available observations of the object's current state.
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeTargetHTTPSProxy resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// Creation timestamp in RFC3339 text format.
	// +kcc:proto:field=google.cloud.compute.v1.TargetHttpsProxy.creation_timestamp
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// The unique identifier for the resource.
	// +kcc:proto:field=google.cloud.compute.v1.TargetHttpsProxy.id
	ProxyId *int `json:"proxyId,omitempty"`

	// The SelfLink for the resource.
	// +kcc:proto:field=google.cloud.compute.v1.TargetHttpsProxy.self_link
	SelfLink *string `json:"selfLink,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ComputeTargetHTTPSProxyObservedState `json:"observedState,omitempty"`
}

// ComputeTargetHTTPSProxyObservedState is the state of the ComputeTargetHTTPSProxy resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.TargetHttpsProxy
type ComputeTargetHTTPSProxyObservedState struct {
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.
	// +kcc:proto:field=google.cloud.compute.v1.TargetHttpsProxy.fingerprint
	Fingerprint *string `json:"fingerprint,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeTargetHTTPSProxyList contains a list of ComputeTargetHTTPSProxy
type ComputeTargetHTTPSProxyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeTargetHTTPSProxy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeTargetHTTPSProxy{}, &ComputeTargetHTTPSProxyList{})
}
