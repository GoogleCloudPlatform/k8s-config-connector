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


// +kcc:proto=google.cloud.vmwareengine.v1.DnsForwarding
type DnsForwarding struct {

	// Required. List of domain mappings to configure
	// +kcc:proto:field=google.cloud.vmwareengine.v1.DnsForwarding.forwarding_rules
	ForwardingRules []DnsForwarding_ForwardingRule `json:"forwardingRules,omitempty"`
}

// +kcc:proto=google.cloud.vmwareengine.v1.DnsForwarding.ForwardingRule
type DnsForwarding_ForwardingRule struct {
	// Required. Domain used to resolve a `name_servers` list.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.DnsForwarding.ForwardingRule.domain
	Domain *string `json:"domain,omitempty"`

	// Required. List of DNS servers to use for domain resolution
	// +kcc:proto:field=google.cloud.vmwareengine.v1.DnsForwarding.ForwardingRule.name_servers
	NameServers []string `json:"nameServers,omitempty"`
}

// +kcc:proto=google.cloud.vmwareengine.v1.DnsForwarding
type DnsForwardingObservedState struct {
	// Output only. The resource name of this DNS profile.
	//  Resource names are schemeless URIs that follow the conventions in
	//  https://cloud.google.com/apis/design/resource_names.
	//  For example:
	//  `projects/my-project/locations/us-central1-a/privateClouds/my-cloud/dnsForwarding`
	// +kcc:proto:field=google.cloud.vmwareengine.v1.DnsForwarding.name
	Name *string `json:"name,omitempty"`

	// Output only. Creation time of this resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.DnsForwarding.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Last update time of this resource.
	// +kcc:proto:field=google.cloud.vmwareengine.v1.DnsForwarding.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
