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

var ComputeTargetHTTPProxyGVK = GroupVersion.WithKind("ComputeTargetHTTPProxy")

type ComputeTargetHTTPProxyParent struct {
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// +required
	Location string `json:"location"`
}

// ComputeTargetHTTPProxySpec defines the desired state of ComputeTargetHTTPProxy
// +kcc:spec:proto=google.cloud.compute.v1.TargetHttpProxy
type ComputeTargetHTTPProxySpec struct {
	// Parent reference.
	ComputeTargetHTTPProxyParent `json:",inline"`

	// The ComputeTargetHTTPProxy name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// An optional description of this resource. Provide this property when you create the resource.
	// +kcc:proto:field=google.cloud.compute.v1.TargetHttpProxy.description
	Description *string `json:"description,omitempty"`

	// Specifies how long to keep a connection open, after completing a response, while there is no matching traffic (in seconds). If an HTTP keep-alive is not specified, a default value (610 seconds) will be used. For global external Application Load Balancers, the minimum allowed value is 5 seconds and the maximum allowed value is 1200 seconds. For classic Application Load Balancers, this option is not supported.
	// +kcc:proto:field=google.cloud.compute.v1.TargetHttpProxy.http_keep_alive_timeout_sec
	HTTPKeepAliveTimeoutSec *int32 `json:"httpKeepAliveTimeoutSec,omitempty"`

	// This field only applies when the forwarding rule that references this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED. When this field is set to true, Envoy proxies set up inbound traffic interception and bind to the IP address and port specified in the forwarding rule. This is generally useful when using Traffic Director to configure Envoy as a gateway or middle proxy (in other words, not a sidecar proxy). The Envoy proxy listens for inbound requests and handles requests when it receives them. The default is false.
	// +kcc:proto:field=google.cloud.compute.v1.TargetHttpProxy.proxy_bind
	ProxyBind *bool `json:"proxyBind,omitempty"`

	// URL to the UrlMap resource that defines the mapping from URL to the BackendService.
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.TargetHttpProxy.url_map
	UrlMapRef *computev1beta1.ComputeURLMapRef `json:"urlMapRef"`
}

// ComputeTargetHTTPProxyStatus defines the config connector machine state of ComputeTargetHTTPProxy
type ComputeTargetHTTPProxyStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeTargetHTTPProxy resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ComputeTargetHTTPProxyObservedState `json:"observedState,omitempty"`
}

// ComputeTargetHTTPProxyObservedState is the state of the ComputeTargetHTTPProxy resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.TargetHttpProxy
type ComputeTargetHTTPProxyObservedState struct {
	// [Output Only] Creation timestamp in RFC3339 text format.
	// +kcc:proto:field=google.cloud.compute.v1.TargetHttpProxy.creation_timestamp
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a TargetHttpProxy. An up-to-date fingerprint must be provided in order to patch/update the TargetHttpProxy; otherwise, the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve the TargetHttpProxy.
	// +kcc:proto:field=google.cloud.compute.v1.TargetHttpProxy.fingerprint
	Fingerprint *string `json:"fingerprint,omitempty"`

	// [Output Only] The unique identifier for the resource. This identifier is defined by the server.
	// +kcc:proto:field=google.cloud.compute.v1.TargetHttpProxy.id
	ID *uint64 `json:"id,omitempty"`

	// [Output Only] Type of resource. Always compute#targetHttpProxy for target HTTP proxies.
	// +kcc:proto:field=google.cloud.compute.v1.TargetHttpProxy.kind
	Kind *string `json:"kind,omitempty"`

	// [Output Only] URL of the region where the regional Target HTTP Proxy resides. This field is not applicable to global Target HTTP Proxies.
	// +kcc:proto:field=google.cloud.compute.v1.TargetHttpProxy.region
	Region *string `json:"region,omitempty"`

	// [Output Only] Server-defined URL for the resource.
	// +kcc:proto:field=google.cloud.compute.v1.TargetHttpProxy.self_link
	SelfLink *string `json:"selfLink,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputetargethttpproxy;gcpcomputetargethttpproxys
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeTargetHTTPProxy is the Schema for the ComputeTargetHTTPProxy API
// +k8s:openapi-gen=true
type ComputeTargetHTTPProxy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeTargetHTTPProxySpec   `json:"spec,omitempty"`
	Status ComputeTargetHTTPProxyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeTargetHTTPProxyList contains a list of ComputeTargetHTTPProxy
type ComputeTargetHTTPProxyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeTargetHTTPProxy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeTargetHTTPProxy{}, &ComputeTargetHTTPProxyList{})
}
