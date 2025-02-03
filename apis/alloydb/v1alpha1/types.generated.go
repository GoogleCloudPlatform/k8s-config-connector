// Copyright 2024 Google LLC
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

// +kcc:proto=google.cloud.alloydb.v1beta.GeminiInstanceConfig
type GeminiInstanceConfig struct {
	// Output only. Whether the Gemini in Databases add-on is enabled for the
	//  instance. It will be true only if the add-on has been enabled for the
	//  billing account corresponding to the instance. Its status is toggled from
	//  the Admin Control Center (ACC) and cannot be toggled using AlloyDBInstance's APIs.
	Entitled *bool `json:"entitled,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1beta.Instance.ClientConnectionConfig
type Instance_ClientConnectionConfig struct {
	// Optional. Configuration to enforce connectors only (ex: AuthProxy)
	//  connections to the database.
	RequireConnectors *bool `json:"requireConnectors,omitempty"`

	// Optional. SSL configuration option for this instance.
	SSLConfig *SSLConfig `json:"sslConfig,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1beta.Instance.InstanceNetworkConfig
type Instance_InstanceNetworkConfig struct {
	// Optional. A list of external network authorized to access this instance.
	// This field is only allowed to be set when 'enablePublicIp' is set to true.
	AuthorizedExternalNetworks []Instance_InstanceNetworkConfig_AuthorizedNetwork `json:"authorizedExternalNetworks,omitempty"`

	// Optional. Enabling public ip for the instance. If a user wishes
	// to disable this, please also clear the list of the authorized
	// external networks set on the same instance.
	EnablePublicIP *bool `json:"enablePublicIp,omitempty"`

	// Optional. Enabling an outbound public IP address to support a database
	//  server sending requests out into the internet.
	EnableOutboundPublicIP *bool `json:"enableOutboundPublicIp,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1beta.Instance.InstanceNetworkConfig.AuthorizedNetwork
type Instance_InstanceNetworkConfig_AuthorizedNetwork struct {
	// CIDR range for one authorzied network of the instance.
	CidrRange *string `json:"cidrRange,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1beta.Instance.MachineConfig
type Instance_MachineConfig struct {
	// The number of CPU's in the VM instance.
	CPUCount *int32 `json:"cpuCount,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1beta.Instance.Node
type Instance_Node struct {
	// The Compute Engine zone of the VM e.g. "us-central1-b".
	ZoneID *string `json:"zoneID,omitempty"`

	// The identifier of the VM e.g. "test-read-0601-407e52be-ms3l".
	ID *string `json:"id,omitempty"`

	// The private IP address of the VM e.g. "10.57.0.34".
	IP *string `json:"ip,omitempty"`

	// Determined by state of the compute VM and postgres-service health.
	//  Compute VM state can have values listed in
	//  https://cloud.google.com/compute/docs/instances/instance-life-cycle and
	//  postgres-service health can have values: HEALTHY and UNHEALTHY.
	State *string `json:"state,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1beta.Instance.ObservabilityInstanceConfig
type Instance_ObservabilityInstanceConfig struct {
	// Observability feature status for an instance.
	//  This flag is turned "off" by default.
	Enabled *bool `json:"enabled,omitempty"`

	// Preserve comments in query string for an instance.
	//  This flag is turned "off" by default.
	PreserveComments *bool `json:"preserveComments,omitempty"`

	// Track wait events during query execution for an instance.
	//  This flag is turned "on" by default but tracking is enabled only after
	//  observability enabled flag is also turned on.
	TrackWaitEvents *bool `json:"trackWaitEvents,omitempty"`

	// Output only. Track wait event types during query execution for an
	//  instance. This flag is turned "on" by default but tracking is enabled
	//  only after observability enabled flag is also turned on. This is
	//  read-only flag and only modifiable by producer API.
	TrackWaitEventTypes *bool `json:"trackWaitEventTypes,omitempty"`

	// Query string length. The default value is 10k.
	MaxQueryStringLength *int32 `json:"maxQueryStringLength,omitempty"`

	// Record application tags for an instance.
	//  This flag is turned "off" by default.
	RecordApplicationTags *bool `json:"recordApplicationTags,omitempty"`

	// Number of query execution plans captured by Insights per minute
	//  for all queries combined. The default value is 200.
	//  Any integer between 0 to 200 is considered valid.
	QueryPlansPerMinute *int32 `json:"queryPlansPerMinute,omitempty"`

	// Track actively running queries on the instance.
	//  If not set, this flag is "off" by default.
	TrackActiveQueries *bool `json:"trackActiveQueries,omitempty"`

	// Track client address for an instance.
	//  If not set, default value is "off".
	TrackClientAddress *bool `json:"trackClientAddress,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1beta.Instance.PscInstanceConfig
type Instance_PscInstanceConfig struct {
	// Output only. The service attachment created when Private
	//  Service Connect (PSC) is enabled for the instance.
	//  The name of the resource will be in the format of
	//  `projects/<alloydb-tenant-project-number>/regions/<region-name>/serviceAttachments/<service-attachment-name>`
	ServiceAttachmentLink *string `json:"serviceAttachmentLink,omitempty"`

	// Optional. List of consumer projects that are allowed to create
	//  PSC endpoints to service-attachments to this instance.
	AllowedConsumerProjects []string `json:"allowedConsumerProjects,omitempty"`

	// Output only. The DNS name of the instance for PSC connectivity.
	//  Name convention: <uid>.<uid>.<region>.alloydb-psc.goog
	PSCDNSName *string `json:"pscDnsName,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1beta.Instance.QueryInsightsInstanceConfig
type Instance_QueryInsightsInstanceConfig struct {
	// Record application tags for an instance.
	//  This flag is turned "on" by default.
	RecordApplicationTags *bool `json:"recordApplicationTags,omitempty"`

	// Record client address for an instance. Client address is PII information.
	//  This flag is turned "on" by default.
	RecordClientAddress *bool `json:"recordClientAddress,omitempty"`

	// Query string length. The default value is 1024.
	//  Any integer between 256 and 4500 is considered valid.
	QueryStringLength *uint32 `json:"queryStringLength,omitempty"`

	// Number of query execution plans captured by Insights per minute
	//  for all queries combined. The default value is 5.
	//  Any integer between 0 and 20 is considered valid.
	QueryPlansPerMinute *uint32 `json:"queryPlansPerMinute,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1beta.Instance.ReadPoolConfig
type Instance_ReadPoolConfig struct {
	// Read capacity, i.e. number of nodes in a read pool instance.
	NodeCount *int32 `json:"nodeCount,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1beta.Instance.UpdatePolicy
type Instance_UpdatePolicy struct {
	// Mode for updating the instance.
	Mode *string `json:"mode,omitempty"`
}

// +kcc:proto=google.cloud.alloydb.v1beta.SslConfig
type SSLConfig struct {
	// Optional. SSL mode. Specifies client-server SSL/TLS connection behavior.
	SSLMode *string `json:"sslMode,omitempty"`

	// Optional. Certificate Authority (CA) source. Only CA_SOURCE_MANAGED is
	//  supported currently, and is the default value.
	CASource *string `json:"caSource,omitempty"`
}
