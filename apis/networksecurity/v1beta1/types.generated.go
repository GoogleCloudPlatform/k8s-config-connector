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

package v1beta1

// +kcc:proto=google.cloud.networksecurity.v1beta1.AuthorizationPolicy.Rule
type AuthorizationPolicy_Rule struct {
	// Optional. List of attributes for the traffic source. All of the sources
	//  must match. A source is a match if both principals and ip_blocks match.
	//  If not set, the action specified in the 'action' field will be applied
	//  without any rule checks for the source.
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.AuthorizationPolicy.Rule.sources
	Sources []AuthorizationPolicy_Rule_Source `json:"sources,omitempty"`

	// Optional. List of attributes for the traffic destination. All of the
	//  destinations must match. A destination is a match if a request matches
	//  all the specified hosts, ports, methods and headers. If not set, the
	//  action specified in the 'action' field will be applied without any rule
	//  checks for the destination.
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.AuthorizationPolicy.Rule.destinations
	Destinations []AuthorizationPolicy_Rule_Destination `json:"destinations,omitempty"`
}

// +kcc:proto=google.cloud.networksecurity.v1beta1.AuthorizationPolicy.Rule.Source
type AuthorizationPolicy_Rule_Source struct {
	// Optional. List of peer identities to match for authorization. At least
	//  one principal should match. Each peer can be an exact match, or a
	//  prefix match (example, "namespace/*") or a suffix match (example,
	//  "*/service-account") or a presence match "*". Authorization based on
	//  the principal name without certificate validation (configured by
	//  ServerTlsPolicy resource) is considered insecure.
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.AuthorizationPolicy.Rule.Source.principals
	Principals []string `json:"principals,omitempty"`

	// Optional. List of CIDR ranges to match based on source IP address. At
	//  least one IP block should match. Single IP (e.g., "1.2.3.4") and CIDR
	//  (e.g., "1.2.3.0/24") are supported. Authorization based on source IP
	//  alone should be avoided. The IP addresses of any load balancers or
	//  proxies should be considered untrusted.
	// +kcc:proto:field=google.cloud.networksecurity.v1beta1.AuthorizationPolicy.Rule.Source.ip_blocks
	IPBlocks []string `json:"ipBlocks,omitempty"`
}
