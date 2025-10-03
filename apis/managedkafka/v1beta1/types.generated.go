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
// krm.group: managedkafka.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.cloud.managedkafka.v1
// resource: ManagedKafkaCluster:Cluster
// resource: ManagedKafkaTopic:Topic

package v1beta1

// +kcc:proto=google.cloud.managedkafka.v1.RebalanceConfig
type RebalanceConfig struct {
	// Optional. The rebalance behavior for the cluster.
	//  When not specified, defaults to `NO_REBALANCE`.
	// +kcc:proto:field=google.cloud.managedkafka.v1.RebalanceConfig.mode
	Mode *string `json:"mode,omitempty"`
}

// +kcc:proto=google.cloud.managedkafka.v1.TlsConfig
type TLSConfig struct {
	// Optional. The configuration of the broker truststore. If specified, clients
	//  can use mTLS for authentication.
	// +kcc:proto:field=google.cloud.managedkafka.v1.TlsConfig.trust_config
	TrustConfig *TrustConfig `json:"trustConfig,omitempty"`

	// Optional. A list of rules for mapping from SSL principal names to
	//  short names. These are applied in order by Kafka.
	//  Refer to the Apache Kafka documentation for `ssl.principal.mapping.rules`
	//  for the precise formatting details and syntax.
	//  Example: "RULE:^CN=(.*?),OU=ServiceUsers.*$/$1@example.com/,DEFAULT"
	//
	//  This is a static Kafka broker configuration. Setting or modifying this
	//  field will trigger a rolling restart of the Kafka brokers to apply
	//  the change. An empty string means no rules are applied (Kafka default).
	// +kcc:proto:field=google.cloud.managedkafka.v1.TlsConfig.ssl_principal_mapping_rules
	SSLPrincipalMappingRules *string `json:"sslPrincipalMappingRules,omitempty"`
}

// +kcc:proto=google.cloud.managedkafka.v1.TrustConfig
type TrustConfig struct {
	// Optional. Configuration for the Google Certificate Authority Service.
	//  Maximum 10.
	// +kcc:proto:field=google.cloud.managedkafka.v1.TrustConfig.cas_configs
	CasConfigs []TrustConfig_CertificateAuthorityServiceConfig `json:"casConfigs,omitempty"`
}

// +kcc:proto=google.cloud.managedkafka.v1.TrustConfig.CertificateAuthorityServiceConfig
type TrustConfig_CertificateAuthorityServiceConfig struct {
	// Required. The name of the CA pool to pull CA certificates from.
	//  Structured like:
	//  projects/{project}/locations/{location}/caPools/{ca_pool}.
	//  The CA pool does not need to be in the same project or location as the
	//  Kafka cluster.
	// +kcc:proto:field=google.cloud.managedkafka.v1.TrustConfig.CertificateAuthorityServiceConfig.ca_pool
	CAPool *string `json:"caPool,omitempty"`
}
