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

var NetworkServicesLBRouteExtensionGVK = GroupVersion.WithKind("NetworkServicesLBRouteExtension")

// NetworkServicesLBRouteExtensionSpec defines the desired state of NetworkServicesLBRouteExtension
// +kcc:spec:proto=google.cloud.networkservices.v1.LbRouteExtension
type NetworkServicesLBRouteExtensionSpec struct {
	// Required. Defines the parent path of the resource.
	*parent.ProjectAndLocationRef `json:",inline"`

	// The NetworkServicesLBRouteExtension name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. A human-readable description of the resource.
	// +kcc:proto:field=google.cloud.networkservices.v1.LbRouteExtension.description
	Description *string `json:"description,omitempty"`

	// Optional. Set of labels associated with the LBRouteExtension resource.
	// +kcc:proto:field=google.cloud.networkservices.v1.LbRouteExtension.labels
	// Labels map[string]string `json:"labels,omitempty"`

	// Required. A list of references to the forwarding rules to which this service extension is attached. At least one forwarding rule is required. Only one LBRouteExtension resource can be associated with a forwarding rule.
	// +kcc:proto:field=google.cloud.networkservices.v1.LbRouteExtension.forwarding_rules
	// +required
	ForwardingRuleRefs []*computev1beta1.ForwardingRuleRef `json:"forwardingRuleRefs,omitempty"`

	// Required. A set of ordered extension chains that contain the match conditions and extensions to execute. Match conditions for each extension chain are evaluated in sequence for a given request. The first extension chain that has a condition that matches the request is executed. Any subsequent extension chains do not execute. Limited to 5 extension chains per resource.
	// +kcc:proto:field=google.cloud.networkservices.v1.LbRouteExtension.extension_chains
	// +required
	ExtensionChains []ExtensionChain `json:"extensionChains,omitempty"`

	// Required. All backend services and forwarding rules referenced by this extension must share the same load balancing scheme. Supported values: INTERNAL_MANAGED, EXTERNAL_MANAGED.
	// +kcc:proto:field=google.cloud.networkservices.v1.LbRouteExtension.load_balancing_scheme
	// +required
	LoadBalancingScheme *string `json:"loadBalancingScheme,omitempty"`

	// Optional. The metadata provided here is included as part of the metadata_context (of type google.protobuf.Struct) in the ProcessingRequest message sent to the extension server.
	// +kcc:proto:field=google.cloud.networkservices.v1.LbRouteExtension.metadata
	Metadata *apiextensionsv1.JSON `json:"metadata,omitempty"`
}

// NetworkServicesLBRouteExtensionStatus defines the config connector machine state of NetworkServicesLBRouteExtension
type NetworkServicesLBRouteExtensionStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the NetworkServicesLBRouteExtension resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *NetworkServicesLBRouteExtensionObservedState `json:"observedState,omitempty"`
}

// NetworkServicesLBRouteExtensionObservedState is the state of the NetworkServicesLBRouteExtension resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.networkservices.v1.LbRouteExtension
type NetworkServicesLBRouteExtensionObservedState struct {
	// Output only. The timestamp when the resource was created.
	// +kcc:proto:field=google.cloud.networkservices.v1.LbRouteExtension.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the resource was updated.
	// +kcc:proto:field=google.cloud.networkservices.v1.LbRouteExtension.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.ExtensionChain
type ExtensionChain struct {
	// Required. The name for this extension chain.
	// The name is logged as part of the HTTP request logs.
	// The name must conform with RFC-1034, is restricted to lower-cased letters,
	// numbers and hyphens, and can have a maximum length of 63 characters.
	// Additionally, the first character must be a letter and the last a letter or
	// a number.
	// +kcc:proto:field=google.cloud.networkservices.v1.ExtensionChain.name
	// +required
	Name *string `json:"name,omitempty"`

	// Required. Conditions under which this chain is invoked for a request.
	// +kcc:proto:field=google.cloud.networkservices.v1.ExtensionChain.match_condition
	// +required
	MatchCondition *ExtensionChain_MatchCondition `json:"matchCondition,omitempty"`

	// Required. A set of extensions to execute for the matching request.
	// At least one extension is required.
	// Up to 3 extensions can be defined for each extension chain
	// for `LBTrafficExtension` resource.
	// `LBRouteExtension` and `LBEdgeExtension` chains are limited to 1 extension
	// per extension chain.
	// +kcc:proto:field=google.cloud.networkservices.v1.ExtensionChain.extensions
	// +required
	Extensions []ExtensionChain_Extension `json:"extensions,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.ExtensionChain.MatchCondition
type ExtensionChain_MatchCondition struct {
	// Required. A Common Expression Language (CEL) expression that is used to
	// match requests for which the extension chain is executed.
	//
	// For more information, see [CEL matcher language
	// reference](https://cloud.google.com/service-extensions/docs/cel-matcher-language-reference).
	// +kcc:proto:field=google.cloud.networkservices.v1.ExtensionChain.MatchCondition.cel_expression
	// +required
	CelExpression *string `json:"celExpression,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.ExtensionChain.Extension
// +kubebuilder:validation:XValidation:rule="has(self.backendServiceRef) != has(self.wasmPluginRef)",message="either backendServiceRef or wasmPluginRef must be set"
type ExtensionChain_Extension struct {
	// Required. The name for this extension.
	// The name is logged as part of the HTTP request logs.
	// The name must conform with RFC-1034, is restricted to lower-cased
	// letters, numbers and hyphens, and can have a maximum length of 63
	// characters. Additionally, the first character must be a letter and the
	// last a letter or a number.
	// +kcc:proto:field=google.cloud.networkservices.v1.ExtensionChain.Extension.name
	// +required
	Name *string `json:"name,omitempty"`

	// Optional. The `:authority` header in the gRPC request sent from Envoy
	// to the extension service.
	// Required for Callout extensions.
	//
	// This field is not supported for plugin extensions. Setting it results in
	// a validation error.
	// +kcc:proto:field=google.cloud.networkservices.v1.ExtensionChain.Extension.authority
	Authority *string `json:"authority,omitempty"`

	// Optional. The reference to the service that runs the extension.
	//
	// To configure a callout extension, `service` must be a fully-qualified reference
	// to a [backend service](https://cloud.google.com/compute/docs/reference/rest/v1/backendServices)
	// in the format:
	// `https://www.googleapis.com/compute/v1/projects/{project}/regions/{region}/backendServices/{backendService}`
	// or
	// `https://www.googleapis.com/compute/v1/projects/{project}/global/backendServices/{backendService}`.
	// +kcc:proto:field=google.cloud.networkservices.v1.ExtensionChain.Extension.service
	BackendServiceRef *computev1beta1.ComputeBackendServiceRef `json:"backendServiceRef,omitempty"`

	// Optional. The reference to the WasmPlugin resource.
	//
	// To configure a plugin extension, `service` must be a reference to a
	// [`WasmPlugin` resource](https://cloud.google.com/service-extensions/docs/reference/rest/v1beta1/projects.locations.wasmPlugins)
	// in the format:
	// `projects/{project}/locations/{location}/wasmPlugins/{plugin}`
	// or
	// `//networkservices.googleapis.com/projects/{project}/locations/{location}/wasmPlugins/{wasmPlugin}`.
	//
	// Plugin extensions are currently supported for the `LBTrafficExtension`, the
	// `LBRouteExtension`, and the `LBEdgeExtension` resources.
	//
	// Note: Only the `external` subfield is supported as the
	// `NetworkServicesWasmPlugin` resource is not yet supported in Config Connector.
	// +kcc:proto:field=google.cloud.networkservices.v1.ExtensionChain.Extension.service
	WasmPluginRef *NetworkServicesWasmPluginRef `json:"wasmPluginRef,omitempty"`

	// Optional. A set of events during request or response processing for which
	// this extension is called.
	//
	// For the `LBTrafficExtension` resource, this field is required.
	//
	// For the `LBRouteExtension` resource, this field is optional. If
	// unspecified, `REQUEST_HEADERS` event is assumed as supported.
	//
	// For the `LBEdgeExtension` resource, this field is required and must only
	// contain `REQUEST_HEADERS` event.
	// +kcc:proto:field=google.cloud.networkservices.v1.ExtensionChain.Extension.supported_events
	SupportedEvents []string `json:"supportedEvents,omitempty"`

	// Optional. Specifies the timeout for each individual message on the
	// stream. The timeout must be between `10`-`10000` milliseconds. Required
	// for callout extensions.
	//
	// This field is not supported for plugin extensions. Setting it results in
	// a validation error.
	// +kcc:proto:field=google.cloud.networkservices.v1.ExtensionChain.Extension.timeout
	Timeout *string `json:"timeout,omitempty"`

	// Optional. Determines how the proxy behaves if the call to the extension
	// fails or times out.
	//
	// When set to `TRUE`, request or response processing continues without
	// error. Any subsequent extensions in the extension chain are also
	// executed. When set to `FALSE` or the default setting of `FALSE` is used,
	// one of the following happens:
	//
	// * If response headers have not been delivered to the downstream client,
	// a generic 500 error is returned to the client. The error response can be
	// tailored by configuring a custom error response in the load balancer.
	//
	// * If response headers have been delivered, then the HTTP stream to the
	// downstream client is reset.
	// +kcc:proto:field=google.cloud.networkservices.v1.ExtensionChain.Extension.fail_open
	FailOpen *bool `json:"failOpen,omitempty"`

	// Optional. List of the HTTP headers to forward to the extension
	// (from the client or backend). If omitted, all headers are sent.
	// Each element is a string indicating the header name.
	// +kcc:proto:field=google.cloud.networkservices.v1.ExtensionChain.Extension.forward_headers
	ForwardHeaders []string `json:"forwardHeaders,omitempty"`

	// Optional. The metadata provided here is included as part of the
	// `metadata_context` (of type `google.protobuf.Struct`) in the
	// `ProcessingRequest` message sent to the extension server.
	//
	// The metadata is available under the namespace
	// `com.google.<extension_type>.<resource_name>.<extension_chain_name>.<extension_name>`.
	// For example:
	// `com.google.lb_traffic_extension.lbtrafficextension1.chain1.ext1`.
	//
	// The following variables are supported in the metadata:
	//
	// `{forwarding_rule_id}` - substituted with the forwarding rule's fully
	// qualified resource name.
	//
	// This field must not be set for plugin extensions. Setting it results in
	// a validation error.
	//
	// You can set metadata at either the resource level or the extension level.
	// The extension level metadata is recommended because you can pass a
	// different set of metadata through each extension to the backend.
	//
	// This field is subject to following limitations:
	//
	// * The total size of the metadata must be less than 1KiB.
	// * The total number of keys in the metadata must be less than 16.
	// * The length of each key must be less than 64 characters.
	// * The length of each value must be less than 1024 characters.
	// * All values must be strings.
	// +kcc:proto:field=google.cloud.networkservices.v1.ExtensionChain.Extension.metadata
	Metadata *apiextensionsv1.JSON `json:"metadata,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetworkserviceslbrouteextension;gcpnetworkserviceslbrouteextensions
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetworkServicesLBRouteExtension is the Schema for the NetworkServicesLBRouteExtension API
// +k8s:openapi-gen=true
type NetworkServicesLBRouteExtension struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   NetworkServicesLBRouteExtensionSpec   `json:"spec,omitempty"`
	Status NetworkServicesLBRouteExtensionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NetworkServicesLBRouteExtensionList contains a list of NetworkServicesLBRouteExtension
type NetworkServicesLBRouteExtensionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkServicesLBRouteExtension `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkServicesLBRouteExtension{}, &NetworkServicesLBRouteExtensionList{})
}
