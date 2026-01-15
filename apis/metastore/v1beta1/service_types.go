// Copyright 2025 Google LLC
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

package v1beta1

// +kcc:proto=google.cloud.metastore.v1.HiveMetastoreConfig
type HiveMetastoreConfig struct {
	// Immutable. The Hive metastore schema version.
	// +kcc:proto:field=google.cloud.metastore.v1.HiveMetastoreConfig.version
	Version *string `json:"version,omitempty"`

	// A mapping of Hive metastore configuration key-value pairs to apply to the
	//  Hive metastore (configured in `hive-site.xml`). The mappings
	//  override system defaults (some keys cannot be overridden). These
	//  overrides are also applied to auxiliary versions and can be further
	//  customized in the auxiliary version's `AuxiliaryVersionConfig`.
	// +kcc:proto:field=google.cloud.metastore.v1.HiveMetastoreConfig.config_overrides
	ConfigOverrides map[string]string `json:"configOverrides,omitempty"`

	// Information used to configure the Hive metastore service as a service
	//  principal in a Kerberos realm. To disable Kerberos, use the `UpdateService`
	//  method and specify this field's path
	//  (`hive_metastore_config.kerberos_config`) in the request's `update_mask`
	//  while omitting this field from the request's `service`.
	// +kcc:proto:field=google.cloud.metastore.v1.HiveMetastoreConfig.kerberos_config
	KerberosConfig *KerberosConfig `json:"kerberosConfig,omitempty"`

	// The protocol to use for the metastore service endpoint. If unspecified,
	//  defaults to `THRIFT`.
	// +kcc:proto:field=google.cloud.metastore.v1.HiveMetastoreConfig.endpoint_protocol
	EndpointProtocol *string `json:"endpointProtocol,omitempty"`

	// A mapping of Hive metastore version to the auxiliary version
	// configuration. When specified, a secondary Hive metastore service is
	// created along with the primary service. All auxiliary versions must be less
	// than the service's primary version. The key is the auxiliary service name
	// and it must match the regular expression [a-z]([-a-z0-9]*[a-z0-9])?. This
	// means that the first character must be a lowercase letter, and all the
	// following characters must be hyphens, lowercase letters, or digits, except
	// the last character, which cannot be a hyphen.
	// +kcc:proto:field=google.cloud.metastore.v1.HiveMetastoreConfig.auxiliary_versions
	AuxiliaryVersions map[string]AuxiliaryVersionConfig `json:"auxiliaryVersions,omitempty"`
}
