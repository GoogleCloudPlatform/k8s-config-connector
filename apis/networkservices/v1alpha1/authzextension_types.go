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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var NetworkServicesAuthzExtensionGVK = GroupVersion.WithKind("NetworkServicesAuthzExtension")

// NetworkServicesAuthzExtensionSpec defines the desired state of NetworkServicesAuthzExtension
// +kcc:spec:proto=google.cloud.networkservices.v1.AuthzExtension
type NetworkServicesAuthzExtensionSpec struct {
	// Required. Defines the parent path of the resource.
	*parent.ProjectAndLocationRef `json:",inline"`

	// The NetworkServicesAuthzExtension name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. A human-readable description of the resource.
	// +kcc:proto:field=google.cloud.networkservices.v1.AuthzExtension.description
	Description *string `json:"description,omitempty"`

	// Optional. Set of labels associated with the `AuthzExtension` resource.
	// +kcc:proto:field=google.cloud.networkservices.v1.AuthzExtension.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. All backend services and forwarding rules referenced by this
	// extension must share the same load balancing scheme. Supported values:
	// `INTERNAL_MANAGED`, `EXTERNAL_MANAGED`.
	// +kcc:proto:field=google.cloud.networkservices.v1.AuthzExtension.load_balancing_scheme
	// +required
	LoadBalancingScheme *string `json:"loadBalancingScheme,omitempty"`

	// Required. The `:authority` header in the gRPC request sent from Envoy
	// to the extension service.
	// +kcc:proto:field=google.cloud.networkservices.v1.AuthzExtension.authority
	// +required
	Authority *string `json:"authority,omitempty"`

	// Required. The reference to the service that runs the extension.
	// To configure a callout extension, `service` must be a fully-qualified reference
	// to a [backend service](https://cloud.google.com/compute/docs/reference/rest/v1/backendServices).
	// +kcc:proto:field=google.cloud.networkservices.v1.AuthzExtension.service
	// +required
	ServiceRef *computev1beta1.ComputeBackendServiceRef `json:"serviceRef,omitempty"`

	// Required. Specifies the timeout for each individual message on the stream.
	// The timeout must be between 10-10000 milliseconds.
	// +kcc:proto:field=google.cloud.networkservices.v1.AuthzExtension.timeout
	// +required
	Timeout *string `json:"timeout,omitempty"`

	// Optional. Determines how the proxy behaves if the call to the extension
	// fails or times out.
	// When set to `TRUE`, request or response processing continues without
	// error. Any subsequent extensions in the extension chain are also
	// executed. When set to `FALSE` or the default setting of `FALSE` is used,
	// one of the following happens:
	// * If response headers have not been delivered to the downstream client,
	// a generic 500 error is returned to the client. The error response can be
	// tailored by configuring a custom error response in the load balancer.
	// * If response headers have been delivered, then the HTTP stream to the
	// downstream client is reset.
	// +kcc:proto:field=google.cloud.networkservices.v1.AuthzExtension.fail_open
	FailOpen *bool `json:"failOpen,omitempty"`

	// Optional. The metadata provided here is included as part of the
	// `metadata_context` (of type `google.protobuf.Struct`) in the
	// `ProcessingRequest` message sent to the extension
	// server. The metadata is available under the namespace
	// `com.google.authz_extension.<resource_name>`.
	// The following variables are supported in the metadata Struct:
	// `{forwarding_rule_id}` - substituted with the forwarding rule's fully
	// qualified resource name.
	// +kcc:proto:field=google.cloud.networkservices.v1.AuthzExtension.metadata
	Metadata *apiextensionsv1.JSON `json:"metadata,omitempty"`

	// Optional. List of the HTTP headers to forward to the extension
	// (from the client). If omitted, all headers are sent.
	// Each element is a string indicating the header name.
	// +kcc:proto:field=google.cloud.networkservices.v1.AuthzExtension.forward_headers
	ForwardHeaders []string `json:"forwardHeaders,omitempty"`

	// Optional. The format of communication supported by the callout extension.
	// If not specified, the default value `EXT_PROC_GRPC` is used.
	// +kcc:proto:field=google.cloud.networkservices.v1.AuthzExtension.wire_format
	WireFormat *string `json:"wireFormat,omitempty"`
}

// NetworkServicesAuthzExtensionStatus defines the config connector machine state of NetworkServicesAuthzExtension
type NetworkServicesAuthzExtensionStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NetworkServicesAuthzExtension resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *NetworkServicesAuthzExtensionObservedState `json:"observedState,omitempty"`
}

// NetworkServicesAuthzExtensionObservedState is the state of the NetworkServicesAuthzExtension resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.networkservices.v1.AuthzExtension
type NetworkServicesAuthzExtensionObservedState struct {
	// Output only. The timestamp when the resource was created.
	// +kcc:proto:field=google.cloud.networkservices.v1.AuthzExtension.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the resource was updated.
	// +kcc:proto:field=google.cloud.networkservices.v1.AuthzExtension.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetworkservicesauthzextension;gcpnetworkservicesauthzextensions
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetworkServicesAuthzExtension is the Schema for the NetworkServicesAuthzExtension API
// +k8s:openapi-gen=true
type NetworkServicesAuthzExtension struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NetworkServicesAuthzExtensionSpec   `json:"spec,omitempty"`
	Status NetworkServicesAuthzExtensionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NetworkServicesAuthzExtensionList contains a list of NetworkServicesAuthzExtension
type NetworkServicesAuthzExtensionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkServicesAuthzExtension `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkServicesAuthzExtension{}, &NetworkServicesAuthzExtensionList{})
}
