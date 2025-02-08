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


// +kcc:proto=google.cloud.networkservices.v1.Gateway
type Gateway struct {
	// Required. Name of the Gateway resource. It matches pattern
	//  `projects/*/locations/*/gateways/<gateway_name>`.
	// +kcc:proto:field=google.cloud.networkservices.v1.Gateway.name
	Name *string `json:"name,omitempty"`

	// Optional. Set of label tags associated with the Gateway resource.
	// +kcc:proto:field=google.cloud.networkservices.v1.Gateway.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. A free-text description of the resource. Max length 1024
	//  characters.
	// +kcc:proto:field=google.cloud.networkservices.v1.Gateway.description
	Description *string `json:"description,omitempty"`

	// Immutable. The type of the customer managed gateway.
	//  This field is required. If unspecified, an error is returned.
	// +kcc:proto:field=google.cloud.networkservices.v1.Gateway.type
	Type *string `json:"type,omitempty"`

	// Required. One or more ports that the Gateway must receive traffic on. The
	//  proxy binds to the ports specified. Gateway listen on 0.0.0.0 on the ports
	//  specified below.
	// +kcc:proto:field=google.cloud.networkservices.v1.Gateway.ports
	Ports []int32 `json:"ports,omitempty"`

	// Required. Immutable. Scope determines how configuration across multiple
	//  Gateway instances are merged. The configuration for multiple Gateway
	//  instances with the same scope will be merged as presented as a single
	//  coniguration to the proxy/load balancer.
	//
	//  Max length 64 characters.
	//  Scope should start with a letter and can only have letters, numbers,
	//  hyphens.
	// +kcc:proto:field=google.cloud.networkservices.v1.Gateway.scope
	Scope *string `json:"scope,omitempty"`

	// Optional. A fully-qualified ServerTLSPolicy URL reference. Specifies how
	//  TLS traffic is terminated. If empty, TLS termination is disabled.
	// +kcc:proto:field=google.cloud.networkservices.v1.Gateway.server_tls_policy
	ServerTlsPolicy *string `json:"serverTlsPolicy,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.Gateway
type GatewayObservedState struct {
	// Output only. Server-defined URL of this resource
	// +kcc:proto:field=google.cloud.networkservices.v1.Gateway.self_link
	SelfLink *string `json:"selfLink,omitempty"`

	// Output only. The timestamp when the resource was created.
	// +kcc:proto:field=google.cloud.networkservices.v1.Gateway.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the resource was updated.
	// +kcc:proto:field=google.cloud.networkservices.v1.Gateway.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
