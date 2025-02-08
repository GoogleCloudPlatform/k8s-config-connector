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


// +kcc:proto=google.cloud.networkmanagement.v1beta1.VpcFlowLogsConfig
type VpcFlowLogsConfig struct {
	// Identifier. Unique name of the configuration using the form:
	//      `projects/{project_id}/locations/global/vpcFlowLogsConfigs/{vpc_flow_logs_config_id}`
	// +kcc:proto:field=google.cloud.networkmanagement.v1beta1.VpcFlowLogsConfig.name
	Name *string `json:"name,omitempty"`

	// Optional. The user-supplied description of the VPC Flow Logs configuration.
	//  Maximum of 512 characters.
	// +kcc:proto:field=google.cloud.networkmanagement.v1beta1.VpcFlowLogsConfig.description
	Description *string `json:"description,omitempty"`

	// Optional. The state of the VPC Flow Log configuration. Default value is
	//  ENABLED. When creating a new configuration, it must be enabled.
	// +kcc:proto:field=google.cloud.networkmanagement.v1beta1.VpcFlowLogsConfig.state
	State *string `json:"state,omitempty"`

	// Optional. The aggregation interval for the logs. Default value is
	//  INTERVAL_5_SEC.
	// +kcc:proto:field=google.cloud.networkmanagement.v1beta1.VpcFlowLogsConfig.aggregation_interval
	AggregationInterval *string `json:"aggregationInterval,omitempty"`

	// Optional. The value of the field must be in (0, 1]. The sampling rate of
	//  VPC Flow Logs where 1.0 means all collected logs are reported. Setting the
	//  sampling rate to 0.0 is not allowed. If you want to disable VPC Flow Logs,
	//  use the state field instead. Default value is 1.0.
	// +kcc:proto:field=google.cloud.networkmanagement.v1beta1.VpcFlowLogsConfig.flow_sampling
	FlowSampling *float32 `json:"flowSampling,omitempty"`

	// Optional. Configures whether all, none or a subset of metadata fields
	//  should be added to the reported VPC flow logs. Default value is
	//  INCLUDE_ALL_METADATA.
	// +kcc:proto:field=google.cloud.networkmanagement.v1beta1.VpcFlowLogsConfig.metadata
	Metadata *string `json:"metadata,omitempty"`

	// Optional. Custom metadata fields to include in the reported VPC flow logs.
	//  Can only be specified if "metadata" was set to CUSTOM_METADATA.
	// +kcc:proto:field=google.cloud.networkmanagement.v1beta1.VpcFlowLogsConfig.metadata_fields
	MetadataFields []string `json:"metadataFields,omitempty"`

	// Optional. Export filter used to define which VPC Flow Logs should be
	//  logged.
	// +kcc:proto:field=google.cloud.networkmanagement.v1beta1.VpcFlowLogsConfig.filter_expr
	FilterExpr *string `json:"filterExpr,omitempty"`

	// Traffic will be logged from the Interconnect Attachment.
	//  Format:
	//  projects/{project_id}/regions/{region}/interconnectAttachments/{name}
	// +kcc:proto:field=google.cloud.networkmanagement.v1beta1.VpcFlowLogsConfig.interconnect_attachment
	InterconnectAttachment *string `json:"interconnectAttachment,omitempty"`

	// Traffic will be logged from the VPN Tunnel.
	//  Format: projects/{project_id}/regions/{region}/vpnTunnels/{name}
	// +kcc:proto:field=google.cloud.networkmanagement.v1beta1.VpcFlowLogsConfig.vpn_tunnel
	VpnTunnel *string `json:"vpnTunnel,omitempty"`

	// Optional. Resource labels to represent user-provided metadata.
	// +kcc:proto:field=google.cloud.networkmanagement.v1beta1.VpcFlowLogsConfig.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// +kcc:proto=google.cloud.networkmanagement.v1beta1.VpcFlowLogsConfig
type VpcFlowLogsConfigObservedState struct {
	// Output only. The time the config was created.
	// +kcc:proto:field=google.cloud.networkmanagement.v1beta1.VpcFlowLogsConfig.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time the config was updated.
	// +kcc:proto:field=google.cloud.networkmanagement.v1beta1.VpcFlowLogsConfig.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
