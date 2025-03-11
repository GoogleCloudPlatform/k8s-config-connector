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

// +kcc:proto=google.cloud.networksecurity.v1.AuthorizationPolicy
type AuthorizationPolicy struct {
	// Required. Name of the AuthorizationPolicy resource. It matches pattern
	//  `projects/{project}/locations/{location}/authorizationPolicies/<authorization_policy>`.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AuthorizationPolicy.name
	Name *string `json:"name,omitempty"`

	// Optional. Free-text description of the resource.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AuthorizationPolicy.description
	Description *string `json:"description,omitempty"`

	// Optional. Set of label tags associated with the AuthorizationPolicy resource.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AuthorizationPolicy.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. The action to take when a rule match is found. Possible values
	//  are "ALLOW" or "DENY".
	// +kcc:proto:field=google.cloud.networksecurity.v1.AuthorizationPolicy.action
	Action *string `json:"action,omitempty"`

	// Optional. List of rules to match. Note that at least one of the rules must match in
	//  order for the action specified in the 'action' field to be taken. A rule is
	//  a match if there is a matching source and destination. If left blank, the
	//  action specified in the `action` field will be applied on every request.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AuthorizationPolicy.rules
	Rules []AuthorizationPolicy_Rule `json:"rules,omitempty"`
}

// +kcc:proto=google.cloud.networksecurity.v1.AuthorizationPolicy.Rule
type AuthorizationPolicy_Rule struct {
	// Optional. List of attributes for the traffic source. All of the sources must match.
	//  A source is a match if both principals and ip_blocks match. If not set,
	//  the action specified in the 'action' field will be applied without any
	//  rule checks for the source.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AuthorizationPolicy.Rule.sources
	Sources []AuthorizationPolicy_Rule_Source `json:"sources,omitempty"`

	// Optional. List of attributes for the traffic destination. All of the destinations
	//  must match. A destination is a match if a request matches all the
	//  specified hosts, ports, methods and headers. If not set, the
	//  action specified in the 'action' field will be applied without any rule
	//  checks for the destination.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AuthorizationPolicy.Rule.destinations
	Destinations []AuthorizationPolicy_Rule_Destination `json:"destinations,omitempty"`
}

// +kcc:proto=google.cloud.networksecurity.v1.AuthorizationPolicy.Rule.Destination
type AuthorizationPolicy_Rule_Destination struct {
	// Required. List of host names to match. Matched against the ":authority" header in
	//  http requests. At least one host should match. Each host can be an
	//  exact match, or a prefix match (example "mydomain.*") or a suffix
	//  match (example "*.myorg.com") or a presence (any) match "*".
	// +kcc:proto:field=google.cloud.networksecurity.v1.AuthorizationPolicy.Rule.Destination.hosts
	Hosts []string `json:"hosts,omitempty"`

	// Required. List of destination ports to match. At least one port should match.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AuthorizationPolicy.Rule.Destination.ports
	Ports []uint32 `json:"ports,omitempty"`

	// Optional. A list of HTTP methods to match. At least one method should
	//  match. Should not be set for gRPC services.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AuthorizationPolicy.Rule.Destination.methods
	Methods []string `json:"methods,omitempty"`

	// Optional. Match against key:value pair in http header. Provides a flexible match
	//  based on HTTP headers, for potentially advanced use cases. At least one
	//  header should match. Avoid using header matches to make authorization
	//  decisions unless there is a strong guarantee that requests arrive
	//  through a trusted client or proxy.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AuthorizationPolicy.Rule.Destination.http_header_match
	HTTPHeaderMatch *AuthorizationPolicy_Rule_Destination_HTTPHeaderMatch `json:"httpHeaderMatch,omitempty"`
}

// +kcc:proto=google.cloud.networksecurity.v1.AuthorizationPolicy.Rule.Destination.HttpHeaderMatch
type AuthorizationPolicy_Rule_Destination_HTTPHeaderMatch struct {
	// Required. The value of the header must match the regular expression
	//  specified in regexMatch. For regular expression grammar,
	//  please see: en.cppreference.com/w/cpp/regex/ecmascript
	//  For matching against a port specified in the HTTP
	//  request, use a headerMatch with headerName set to Host
	//  and a regular expression that satisfies the RFC2616 Host
	//  header's port specifier.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AuthorizationPolicy.Rule.Destination.HttpHeaderMatch.regex_match
	RegexMatch *string `json:"regexMatch,omitempty"`

	// Required. The name of the HTTP header to match. For matching
	//  against the HTTP request's authority, use a headerMatch
	//  with the header name ":authority". For matching a
	//  request's method, use the headerName ":method".
	// +kcc:proto:field=google.cloud.networksecurity.v1.AuthorizationPolicy.Rule.Destination.HttpHeaderMatch.header_name
	HeaderName *string `json:"headerName,omitempty"`
}

// +kcc:proto=google.cloud.networksecurity.v1.AuthorizationPolicy.Rule.Source
type AuthorizationPolicy_Rule_Source struct {
	// Optional. List of peer identities to match for authorization. At least one
	//  principal should match. Each peer can be an exact match, or a prefix
	//  match (example, "namespace/*") or a suffix match (example,
	//  "*/service-account") or a presence match "*". Authorization based on
	//  the principal name without certificate validation (configured by
	//  ServerTlsPolicy resource) is considered insecure.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AuthorizationPolicy.Rule.Source.principals
	Principals []string `json:"principals,omitempty"`

	// Optional. List of CIDR ranges to match based on source IP address. At least one
	//  IP block should match. Single IP (e.g., "1.2.3.4") and CIDR (e.g.,
	//  "1.2.3.0/24") are supported. Authorization based on source IP alone
	//  should be avoided. The IP addresses of any load balancers or proxies
	//  should be considered untrusted.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AuthorizationPolicy.Rule.Source.ip_blocks
	IPBlocks []string `json:"ipBlocks,omitempty"`
}

// +kcc:proto=google.cloud.networksecurity.v1.AuthorizationPolicy
type AuthorizationPolicyObservedState struct {
	// Output only. The timestamp when the resource was created.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AuthorizationPolicy.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the resource was updated.
	// +kcc:proto:field=google.cloud.networksecurity.v1.AuthorizationPolicy.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
