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

// +generated:types
// krm.group: compute.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.cloud.compute.v1beta
// resource: ComputeTargetTcpProxy:TargetTcpProxy

package v1beta1

// +kcc:proto=google.cloud.compute.v1beta.TargetTcpProxy
type TargetTCPProxy struct {
	// [Output Only] Creation timestamp in RFC3339 text format.
	// +kcc:proto:field=google.cloud.compute.v1beta.TargetTcpProxy.creation_timestamp
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// An optional description of this resource. Provide this property when you create the resource.
	// +kcc:proto:field=google.cloud.compute.v1beta.TargetTcpProxy.description
	Description *string `json:"description,omitempty"`

	// [Output Only] The unique identifier for the resource. This identifier is defined by the server.
	// +kcc:proto:field=google.cloud.compute.v1beta.TargetTcpProxy.id
	ID *uint64 `json:"id,omitempty"`

	// [Output Only] Type of the resource. Always compute#targetTcpProxy for target TCP proxies.
	// +kcc:proto:field=google.cloud.compute.v1beta.TargetTcpProxy.kind
	Kind *string `json:"kind,omitempty"`

	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	// +kcc:proto:field=google.cloud.compute.v1beta.TargetTcpProxy.name
	Name *string `json:"name,omitempty"`

	// This field only applies when the forwarding rule that references this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED. When this field is set to true, Envoy proxies set up inbound traffic interception and bind to the IP address and port specified in the forwarding rule. This is generally useful when using Traffic Director to configure Envoy as a gateway or middle proxy (in other words, not a sidecar proxy). The Envoy proxy listens for inbound requests and handles requests when it receives them. The default is false.
	// +kcc:proto:field=google.cloud.compute.v1beta.TargetTcpProxy.proxy_bind
	ProxyBind *bool `json:"proxyBind,omitempty"`

	// Specifies the type of proxy header to append before sending data to the backend, either NONE or PROXY_V1. The default is NONE.
	//  Check the ProxyHeader enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1beta.TargetTcpProxy.proxy_header
	ProxyHeader *string `json:"proxyHeader,omitempty"`

	// [Output Only] URL of the region where the regional TCP proxy resides. This field is not applicable to global TCP proxy.
	// +kcc:proto:field=google.cloud.compute.v1beta.TargetTcpProxy.region
	Region *string `json:"region,omitempty"`

	// [Output Only] Server-defined URL for the resource.
	// +kcc:proto:field=google.cloud.compute.v1beta.TargetTcpProxy.self_link
	SelfLink *string `json:"selfLink,omitempty"`

	// URL to the BackendService resource.
	// +kcc:proto:field=google.cloud.compute.v1beta.TargetTcpProxy.service
	Service *string `json:"service,omitempty"`
}
