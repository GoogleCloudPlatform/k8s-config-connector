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

// +kcc:proto=mockgcp.cloud.edgecacheservice.v1.EdgeCacheService
type EdgeCacheService struct {
	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.EdgeCacheService.name
	Name *string `json:"name,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.EdgeCacheService.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.EdgeCacheService.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.EdgeCacheService.labels
	Labels map[string]string `json:"labels,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.EdgeCacheService.description
	Description *string `json:"description,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.EdgeCacheService.routing
	Routing *Routing `json:"routing,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.EdgeCacheService.require_tls
	RequireTLS *bool `json:"requireTLS,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.EdgeCacheService.edge_ssl_certificates
	EdgeSSLCertificates []string `json:"edgeSSLCertificates,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.EdgeCacheService.edge_security_policy
	EdgeSecurityPolicy *string `json:"edgeSecurityPolicy,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.EdgeCacheService.log_config
	LogConfig *LogConfig `json:"logConfig,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.EdgeCacheService.disable_quic
	DisableQuic *bool `json:"disableQuic,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.EdgeCacheService.disable_http2
	DisableHttp2 *bool `json:"disableHttp2,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.EdgeCacheService.ipv4_addresses
	IPv4Addresses []string `json:"ipv4Addresses,omitempty"`

	// +kcc:proto:field=mockgcp.cloud.edgecacheservice.v1.EdgeCacheService.ipv6_addresses
	IPv6Addresses []string `json:"ipv6Addresses,omitempty"`
}

type EdgeCacheServiceObservedState EdgeCacheService
