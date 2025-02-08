// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1


// +kcc:proto=google.cloud.networkservices.v1beta1.ExtensionChain
type ExtensionChain struct {
	// Required. The name for this extension chain.
	//  The name is logged as part of the HTTP request logs.
	//  The name must conform with RFC-1034, is restricted to lower-cased letters,
	//  numbers and hyphens, and can have a maximum length of 63 characters.
	//  Additionally, the first character must be a letter and the last a letter or
	//  a number.
	// +kcc:proto:field=google.cloud.networkservices.v1beta1.ExtensionChain.name
	Name *string `json:"name,omitempty"`

	// Required. Conditions under which this chain is invoked for a request.
	// +kcc:proto:field=google.cloud.networkservices.v1beta1.ExtensionChain.match_condition
	MatchCondition *ExtensionChain_MatchCondition `json:"matchCondition,omitempty"`

	// Required. A set of extensions to execute for the matching request.
	//  At least one extension is required.
	//  Up to 3 extensions can be defined for each extension chain
	//  for `LbTrafficExtension` resource.
	//  `LbRouteExtension` chains are limited to 1 extension per extension chain.
	// +kcc:proto:field=google.cloud.networkservices.v1beta1.ExtensionChain.extensions
	Extensions []ExtensionChain_Extension `json:"extensions,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1beta1.ExtensionChain.Extension
type ExtensionChain_Extension struct {
	// Required. The name for this extension.
	//  The name is logged as part of the HTTP request logs.
	//  The name must conform with RFC-1034, is restricted to lower-cased
	//  letters, numbers and hyphens, and can have a maximum length of 63
	//  characters. Additionally, the first character must be a letter and the
	//  last a letter or a number.
	// +kcc:proto:field=google.cloud.networkservices.v1beta1.ExtensionChain.Extension.name
	Name *string `json:"name,omitempty"`

	// Optional. The `:authority` header in the gRPC request sent from Envoy
	//  to the extension service.
	//  Required for Callout extensions.
	// +kcc:proto:field=google.cloud.networkservices.v1beta1.ExtensionChain.Extension.authority
	Authority *string `json:"authority,omitempty"`

	// Required. The reference to the service that runs the extension.
	//
	//  Currently only callout extensions are supported here.
	//
	//  To configure a callout extension, `service` must be a fully-qualified
	//  reference
	//  to a [backend
	//  service](https://cloud.google.com/compute/docs/reference/rest/v1/backendServices)
	//  in the format:
	//  `https://www.googleapis.com/compute/v1/projects/{project}/regions/{region}/backendServices/{backendService}`
	//  or
	//  `https://www.googleapis.com/compute/v1/projects/{project}/global/backendServices/{backendService}`.
	// +kcc:proto:field=google.cloud.networkservices.v1beta1.ExtensionChain.Extension.service
	Service *string `json:"service,omitempty"`

	// Optional. A set of events during request or response processing for which
	//  this extension is called. This field is required for the
	//  `LbTrafficExtension` resource. It's not relevant for the
	//  `LbRouteExtension` resource.
	// +kcc:proto:field=google.cloud.networkservices.v1beta1.ExtensionChain.Extension.supported_events
	SupportedEvents []string `json:"supportedEvents,omitempty"`

	// Optional. Specifies the timeout for each individual message on the
	//  stream. The timeout must be between 10-1000 milliseconds. Required for
	//  Callout extensions.
	// +kcc:proto:field=google.cloud.networkservices.v1beta1.ExtensionChain.Extension.timeout
	Timeout *string `json:"timeout,omitempty"`

	// Optional. Determines how the proxy behaves if the call to the extension
	//  fails or times out.
	//
	//  When set to `TRUE`, request or response processing continues without
	//  error. Any subsequent extensions in the extension chain are also
	//  executed. When set to `FALSE` or the default setting of `FALSE` is used,
	//  one of the following happens:
	//  * If response headers have not been delivered to the downstream client,
	//  a generic 500 error is returned to the client. The error response can be
	//  tailored by configuring a custom error response in the load balancer.
	//
	//  * If response headers have been delivered, then the HTTP stream to the
	//  downstream client is reset.
	// +kcc:proto:field=google.cloud.networkservices.v1beta1.ExtensionChain.Extension.fail_open
	FailOpen *bool `json:"failOpen,omitempty"`

	// Optional. List of the HTTP headers to forward to the extension
	//  (from the client or backend). If omitted, all headers are sent.
	//  Each element is a string indicating the header name.
	// +kcc:proto:field=google.cloud.networkservices.v1beta1.ExtensionChain.Extension.forward_headers
	ForwardHeaders []string `json:"forwardHeaders,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1beta1.ExtensionChain.MatchCondition
type ExtensionChain_MatchCondition struct {
	// Required. A Common Expression Language (CEL) expression that is used to
	//  match requests for which the extension chain is executed.
	//
	//  For more information, see [CEL matcher language
	//  reference](https://cloud.google.com/service-extensions/docs/cel-matcher-language-reference).
	// +kcc:proto:field=google.cloud.networkservices.v1beta1.ExtensionChain.MatchCondition.cel_expression
	CelExpression *string `json:"celExpression,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1beta1.LbRouteExtension
type LbRouteExtension struct {
	// Required. Identifier. Name of the `LbRouteExtension` resource in the
	//  following format:
	//  `projects/{project}/locations/{location}/lbRouteExtensions/{lb_route_extension}`.
	// +kcc:proto:field=google.cloud.networkservices.v1beta1.LbRouteExtension.name
	Name *string `json:"name,omitempty"`

	// Optional. A human-readable description of the resource.
	// +kcc:proto:field=google.cloud.networkservices.v1beta1.LbRouteExtension.description
	Description *string `json:"description,omitempty"`

	// Optional. Set of labels associated with the `LbRouteExtension` resource.
	//
	//  The format must comply with [the requirements for
	//  labels](https://cloud.google.com/compute/docs/labeling-resources#requirements)
	//  for Google Cloud resources.
	// +kcc:proto:field=google.cloud.networkservices.v1beta1.LbRouteExtension.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. A list of references to the forwarding rules to which this
	//  service extension is attached to. At least one forwarding rule is required.
	//  There can be only one `LbRouteExtension` resource per forwarding rule.
	// +kcc:proto:field=google.cloud.networkservices.v1beta1.LbRouteExtension.forwarding_rules
	ForwardingRules []string `json:"forwardingRules,omitempty"`

	// Required. A set of ordered extension chains that contain the match
	//  conditions and extensions to execute. Match conditions for each extension
	//  chain are evaluated in sequence for a given request. The first extension
	//  chain that has a condition that matches the request is executed.
	//  Any subsequent extension chains do not execute.
	//  Limited to 5 extension chains per resource.
	// +kcc:proto:field=google.cloud.networkservices.v1beta1.LbRouteExtension.extension_chains
	ExtensionChains []ExtensionChain `json:"extensionChains,omitempty"`

	// Required. All backend services and forwarding rules referenced by this
	//  extension must share the same load balancing scheme. Supported values:
	//  `INTERNAL_MANAGED`, `EXTERNAL_MANAGED`. For more information, refer to
	//  [Choosing a load
	//  balancer](https://cloud.google.com/load-balancing/docs/backend-service).
	// +kcc:proto:field=google.cloud.networkservices.v1beta1.LbRouteExtension.load_balancing_scheme
	LoadBalancingScheme *string `json:"loadBalancingScheme,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1beta1.LbRouteExtension
type LbRouteExtensionObservedState struct {
	// Output only. The timestamp when the resource was created.
	// +kcc:proto:field=google.cloud.networkservices.v1beta1.LbRouteExtension.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the resource was updated.
	// +kcc:proto:field=google.cloud.networkservices.v1beta1.LbRouteExtension.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
