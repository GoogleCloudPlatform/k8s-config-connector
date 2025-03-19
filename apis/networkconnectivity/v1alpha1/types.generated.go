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

// +kcc:proto=mockgcp.cloud.networkconnectivity.v1.AuditConfig
type AuditConfig struct {
	// The configuration for logging of each type of permission.
	AuditLogConfigs []AuditLogConfig `json:"auditLogConfigs,omitempty"`

	// Specifies a service that will be enabled for audit logging. For example, `storage.googleapis.com`, `cloudsql.googleapis.com`. `allServices` is a special value that covers all services.
	Service *string `json:"service,omitempty"`
}

// +kcc:proto=mockgcp.cloud.networkconnectivity.v1.AuditLogConfig
type AuditLogConfig struct {
	// Specifies the identities that do not cause logging for this type of permission. Follows the same format of Binding.members.
	ExemptedMembers []string `json:"exemptedMembers,omitempty"`

	// The log type that this config enables.
	LogType *string `json:"logType,omitempty"`
}

// +kcc:proto=mockgcp.cloud.networkconnectivity.v1.AutoAccept
type AutoAccept struct {
	// A list of project ids or project numbers for which you want to enable auto-accept. The auto-accept setting is applied to spokes being created or updated in these projects.
	AutoAcceptProjects []string `json:"autoAcceptProjects,omitempty"`
}

// +kcc:proto=mockgcp.cloud.networkconnectivity.v1.Binding
type Binding struct {
	// The condition that is associated with this binding. If the condition evaluates to `true`, then this binding applies to the current request. If the condition evaluates to `false`, then this binding does not apply to the current request. However, a different role binding might grant the same role to one or more of the principals in this binding. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Condition *Expr `json:"condition,omitempty"`

	// Specifies the principals requesting access for a Google Cloud resource. `members` can have the following values: * `allUsers`: A special identifier that represents anyone who is on the internet; with or without a Google account. * `allAuthenticatedUsers`: A special identifier that represents anyone who is authenticated with a Google account or a service account. Does not include identities that come from external identity providers (IdPs) through identity federation. * `user:{emailid}`: An email address that represents a specific Google account. For example, `alice@example.com` . * `serviceAccount:{emailid}`: An email address that represents a Google service account. For example, `my-other-app@appspot.gserviceaccount.com`. * `serviceAccount:{projectid}.svc.id.goog[{namespace}/{kubernetes-sa}]`: An identifier for a [Kubernetes service account](https://cloud.google.com/kubernetes-engine/docs/how-to/kubernetes-service-accounts). For example, `my-project.svc.id.goog[my-namespace/my-kubernetes-sa]`. * `group:{emailid}`: An email address that represents a Google group. For example, `admins@example.com`. * `domain:{domain}`: The G Suite domain (primary) that represents all the users of that domain. For example, `google.com` or `example.com`. * `principal://iam.googleapis.com/locations/global/workforcePools/{pool_id}/subject/{subject_attribute_value}`: A single identity in a workforce identity pool. * `principalSet://iam.googleapis.com/locations/global/workforcePools/{pool_id}/group/{group_id}`: All workforce identities in a group. * `principalSet://iam.googleapis.com/locations/global/workforcePools/{pool_id}/attribute.{attribute_name}/{attribute_value}`: All workforce identities with a specific attribute value. * `principalSet://iam.googleapis.com/locations/global/workforcePools/{pool_id}/*`: All identities in a workforce identity pool. * `principal://iam.googleapis.com/projects/{project_number}/locations/global/workloadIdentityPools/{pool_id}/subject/{subject_attribute_value}`: A single identity in a workload identity pool. * `principalSet://iam.googleapis.com/projects/{project_number}/locations/global/workloadIdentityPools/{pool_id}/group/{group_id}`: A workload identity pool group. * `principalSet://iam.googleapis.com/projects/{project_number}/locations/global/workloadIdentityPools/{pool_id}/attribute.{attribute_name}/{attribute_value}`: All identities in a workload identity pool with a certain attribute. * `principalSet://iam.googleapis.com/projects/{project_number}/locations/global/workloadIdentityPools/{pool_id}/*`: All identities in a workload identity pool. * `deleted:user:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a user that has been recently deleted. For example, `alice@example.com?uid=123456789012345678901`. If the user is recovered, this value reverts to `user:{emailid}` and the recovered user retains the role in the binding. * `deleted:serviceAccount:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a service account that has been recently deleted. For example, `my-other-app@appspot.gserviceaccount.com?uid=123456789012345678901`. If the service account is undeleted, this value reverts to `serviceAccount:{emailid}` and the undeleted service account retains the role in the binding. * `deleted:group:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a Google group that has been recently deleted. For example, `admins@example.com?uid=123456789012345678901`. If the group is recovered, this value reverts to `group:{emailid}` and the recovered group retains the role in the binding. * `deleted:principal://iam.googleapis.com/locations/global/workforcePools/{pool_id}/subject/{subject_attribute_value}`: Deleted single identity in a workforce identity pool. For example, `deleted:principal://iam.googleapis.com/locations/global/workforcePools/my-pool-id/subject/my-subject-attribute-value`.
	Members []string `json:"members,omitempty"`

	// Role that is assigned to the list of `members`, or principals. For example, `roles/viewer`, `roles/editor`, or `roles/owner`. For an overview of the IAM roles and permissions, see the [IAM documentation](https://cloud.google.com/iam/docs/roles-overview). For a list of the available pre-defined roles, see [understanding roles](https://cloud.google.com/iam/docs/understanding-roles).
	Role *string `json:"role,omitempty"`
}

// +kcc:proto=mockgcp.cloud.networkconnectivity.v1.ConsumerPscConfig
type ConsumerPscConfig struct {
	// Required. The project ID or project number of the consumer project. This project is the one that the consumer uses to interact with the producer instance. From the perspective of a consumer who's created a producer instance, this is the project of the producer instance. Format: 'projects/' Eg. 'projects/consumer-project' or 'projects/1234'
	ConsumerInstanceProject *string `json:"consumerInstanceProject,omitempty"`

	// This is used in PSC consumer ForwardingRule to control whether the PSC endpoint can be accessed from another region.
	DisableGlobalAccess *bool `json:"disableGlobalAccess,omitempty"`

	// The resource path of the consumer network where PSC connections are allowed to be created in. Note, this network does not need be in the ConsumerPscConfig.project in the case of SharedVPC. Example: projects/{projectNumOrId}/global/networks/{networkId}.
	Network *string `json:"network,omitempty"`

	// Immutable. An immutable identifier for the producer instance.
	ProducerInstanceID *string `json:"producerInstanceID,omitempty"`

	// The consumer project where PSC connections are allowed to be created in.
	Project *string `json:"project,omitempty"`

	// Output only. A map to store mapping between customer vip and target service attachment. Only service attachment with producer specified ip addresses are stored here.
	ServiceAttachmentIpAddressMap map[string]string `json:"serviceAttachmentIpAddressMap,omitempty"`

	// Output only. Overall state of PSC Connections management for this consumer psc config.
	State *string `json:"state,omitempty"`
}

// +kcc:proto=mockgcp.cloud.networkconnectivity.v1.ConsumerPscConnection
type ConsumerPscConnection struct {
	// The most recent error during operating this connection.
	Error *GoogleRpcStatus `json:"error,omitempty"`

	// Output only. The error info for the latest error during operating this connection.
	ErrorInfo *GoogleRpcErrorInfo `json:"errorInfo,omitempty"`

	// The error type indicates whether the error is consumer facing, producer facing or system internal.
	ErrorType *string `json:"errorType,omitempty"`

	// The URI of the consumer forwarding rule created. Example: projects/{projectNumOrId}/regions/us-east1/networks/{resourceId}.
	ForwardingRule *string `json:"forwardingRule,omitempty"`

	// The last Compute Engine operation to setup PSC connection.
	GceOperation *string `json:"gceOperation,omitempty"`

	// The IP literal allocated on the consumer network for the PSC forwarding rule that is created to connect to the producer service attachment in this service connection map.
	Ip *string `json:"ip,omitempty"`

	// The consumer network whose PSC forwarding rule is connected to the service attachments in this service connection map. Note that the network could be on a different project (shared VPC).
	Network *string `json:"network,omitempty"`

	// Immutable. An immutable identifier for the producer instance.
	ProducerInstanceID *string `json:"producerInstanceID,omitempty"`

	// The consumer project whose PSC forwarding rule is connected to the service attachments in this service connection map.
	Project *string `json:"project,omitempty"`

	// The PSC connection id of the PSC forwarding rule connected to the service attachments in this service connection map.
	PscConnectionID *string `json:"pscConnectionID,omitempty"`

	// Output only. The URI of the selected subnetwork selected to allocate IP address for this connection.
	SelectedSubnetwork *string `json:"selectedSubnetwork,omitempty"`

	// The URI of a service attachment which is the target of the PSC connection.
	ServiceAttachmentUri *string `json:"serviceAttachmentUri,omitempty"`

	// The state of the PSC connection.
	State *string `json:"state,omitempty"`
}

// +kcc:proto=mockgcp.cloud.networkconnectivity.v1.Empty
type Empty struct {
}

// +kcc:proto=mockgcp.cloud.networkconnectivity.v1.Expr
type Expr struct {
	// Optional. Description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
	Description *string `json:"description,omitempty"`

	// Textual representation of an expression in Common Expression Language syntax.
	Expression *string `json:"expression,omitempty"`

	// Optional. String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
	Location *string `json:"location,omitempty"`

	// Optional. Title for the expression, i.e. a short string describing its purpose. This can be used e.g. in UIs which allow to enter the expression.
	Title *string `json:"title,omitempty"`
}

// +kcc:proto=mockgcp.cloud.networkconnectivity.v1.Filter
type Filter struct {
	// Optional. The destination IP range of outgoing packets that this policy-based route applies to. Default is "0.0.0.0/0" if protocol version is IPv4.
	DestRange *string `json:"destRange,omitempty"`

	// Optional. The IP protocol that this policy-based route applies to. Valid values are 'TCP', 'UDP', and 'ALL'. Default is 'ALL'.
	IpProtocol *string `json:"ipProtocol,omitempty"`

	// Required. Internet protocol versions this policy-based route applies to. For this version, only IPV4 is supported. IPV6 is supported in preview.
	ProtocolVersion *string `json:"protocolVersion,omitempty"`

	// Optional. The source IP range of outgoing packets that this policy-based route applies to. Default is "0.0.0.0/0" if protocol version is IPv4.
	SrcRange *string `json:"srcRange,omitempty"`
}

// +kcc:proto=mockgcp.cloud.networkconnectivity.v1.GoogleRpcErrorInfo
type GoogleRpcErrorInfo struct {
	// The logical grouping to which the "reason" belongs. The error domain is typically the registered service name of the tool or product that generates the error. Example: "pubsub.googleapis.com". If the error is generated by some common infrastructure, the error domain must be a globally unique value that identifies the infrastructure. For Google API infrastructure, the error domain is "googleapis.com".
	Domain *string `json:"domain,omitempty"`

	// Additional structured details about this error. Keys must match /a-z+/ but should ideally be lowerCamelCase. Also they must be limited to 64 characters in length. When identifying the current value of an exceeded limit, the units should be contained in the key, not the value. For example, rather than {"instanceLimit": "100/request"}, should be returned as, {"instanceLimitPerRequest": "100"}, if the client exceeds the number of instances that can be created in a single (batch) request.
	Metadata map[string]string `json:"metadata,omitempty"`

	// The reason of the error. This is a constant value that identifies the proximate cause of the error. Error reasons are unique within a particular domain of errors. This should be at most 63 characters and match a regular expression of `A-Z+[A-Z0-9]`, which represents UPPER_SNAKE_CASE.
	Reason *string `json:"reason,omitempty"`
}

// +kcc:proto=mockgcp.cloud.networkconnectivity.v1.Group
type Group struct {
	// Optional. The auto-accept setting for this group.
	AutoAccept *AutoAccept `json:"autoAccept,omitempty"`

	// Output only. The time the group was created.
	CreateTime *string `json:"createTime,omitempty"`

	// Optional. The description of the group.
	Description *string `json:"description,omitempty"`

	// Optional. Labels in key-value pair format. For more information about labels, see [Requirements for labels](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements).
	Labels map[string]string `json:"labels,omitempty"`

	// Immutable. The name of the group. Group names must be unique. They use the following form: `projects/{project_number}/locations/global/hubs/{hub}/groups/{group_id}`
	Name *string `json:"name,omitempty"`

	// Output only. The name of the route table that corresponds to this group. They use the following form: `projects/{project_number}/locations/global/hubs/{hub_id}/routeTables/{route_table_id}`
	RouteTable *string `json:"routeTable,omitempty"`

	// Output only. The current lifecycle state of this group.
	State *string `json:"state,omitempty"`

	// Output only. The Google-generated UUID for the group. This value is unique across all group resources. If a group is deleted and another with the same name is created, the new route table is assigned a different unique_id.
	Uid *string `json:"uid,omitempty"`

	// Output only. The time the group was last updated.
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=mockgcp.cloud.networkconnectivity.v1.Hub
type Hub struct {
	// Output only. The time the hub was created.
	CreateTime *string `json:"createTime,omitempty"`

	// An optional description of the hub.
	Description *string `json:"description,omitempty"`

	// Optional. Whether Private Service Connect transitivity is enabled for the hub. If true, Private Service Connect endpoints in VPC spokes attached to the hub are made accessible to other VPC spokes attached to the hub. The default value is false.
	ExportPsc *bool `json:"exportPsc,omitempty"`

	// Optional labels in key-value pair format. For more information about labels, see [Requirements for labels](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements).
	Labels map[string]string `json:"labels,omitempty"`

	// Immutable. The name of the hub. Hub names must be unique. They use the following form: `projects/{project_number}/locations/global/hubs/{hub_id}`
	Name *string `json:"name,omitempty"`

	// Optional. The policy mode of this hub. This field can be either PRESET or CUSTOM. If unspecified, the policy_mode defaults to PRESET.
	PolicyMode *string `json:"policyMode,omitempty"`

	// Optional. The topology implemented in this hub. Currently, this field is only used when policy_mode = PRESET. The available preset topologies are MESH and STAR. If preset_topology is unspecified and policy_mode = PRESET, the preset_topology defaults to MESH. When policy_mode = CUSTOM, the preset_topology is set to PRESET_TOPOLOGY_UNSPECIFIED.
	PresetTopology *string `json:"presetTopology,omitempty"`

	// Output only. The route tables that belong to this hub. They use the following form: `projects/{project_number}/locations/global/hubs/{hub_id}/routeTables/{route_table_id}` This field is read-only. Network Connectivity Center automatically populates it based on the route tables nested under the hub.
	RouteTables []string `json:"routeTables,omitempty"`

	// The VPC networks associated with this hub's spokes. This field is read-only. Network Connectivity Center automatically populates it based on the set of spokes attached to the hub.
	RoutingVpcs []RoutingVPC `json:"routingVpcs,omitempty"`

	// Output only. A summary of the spokes associated with a hub. The summary includes a count of spokes according to type and according to state. If any spokes are inactive, the summary also lists the reasons they are inactive, including a count for each reason.
	SpokeSummary *SpokeSummary `json:"spokeSummary,omitempty"`

	// Output only. The current lifecycle state of this hub.
	State *string `json:"state,omitempty"`

	// Output only. The Google-generated UUID for the hub. This value is unique across all hub resources. If a hub is deleted and another with the same name is created, the new hub is assigned a different unique_id.
	UniqueID *string `json:"uniqueID,omitempty"`

	// Output only. The time the hub was last updated.
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=mockgcp.cloud.networkconnectivity.v1.InterconnectAttachment
type InterconnectAttachment struct {
	// Optional. Cloud region to install this policy-based route on interconnect attachment. Use `all` to install it on all interconnect attachments.
	Region *string `json:"region,omitempty"`
}

// +kcc:proto=mockgcp.cloud.networkconnectivity.v1.InternalRange
type InternalRange struct {
	// Time when the internal range was created.
	CreateTime *string `json:"createTime,omitempty"`

	// A description of this resource.
	Description *string `json:"description,omitempty"`

	// The IP range that this internal range defines.
	IpCidrRange *string `json:"ipCidrRange,omitempty"`

	// User-defined labels.
	Labels map[string]string `json:"labels,omitempty"`

	// Immutable. The name of an internal range. Format: projects/{project}/locations/{location}/internalRanges/{internal_range} See: https://google.aip.dev/122#fields-representing-resource-names
	Name *string `json:"name,omitempty"`

	// The URL or resource ID of the network in which to reserve the internal range. The network cannot be deleted if there are any reserved internal ranges referring to it. Legacy networks are not supported. For example: https://www.googleapis.com/compute/v1/projects/{project}/locations/global/networks/{network} projects/{project}/locations/global/networks/{network} {network}
	Network *string `json:"network,omitempty"`

	// Optional. Types of resources that are allowed to overlap with the current internal range.
	Overlaps []string `json:"overlaps,omitempty"`

	// The type of peering set for this internal range.
	Peering *string `json:"peering,omitempty"`

	// An alternate to ip_cidr_range. Can be set when trying to create a reservation that automatically finds a free range of the given size. If both ip_cidr_range and prefix_length are set, there is an error if the range sizes do not match. Can also be used during updates to change the range size.
	PrefixLength *int32 `json:"prefixLength,omitempty"`

	// Optional. Can be set to narrow down or pick a different address space while searching for a free range. If not set, defaults to the "10.0.0.0/8" address space. This can be used to search in other rfc-1918 address spaces like "172.16.0.0/12" and "192.168.0.0/16" or non-rfc-1918 address spaces used in the VPC.
	TargetCidrRange []string `json:"targetCidrRange,omitempty"`

	// Time when the internal range was updated.
	UpdateTime *string `json:"updateTime,omitempty"`

	// The type of usage set for this InternalRange.
	Usage *string `json:"usage,omitempty"`

	// Output only. The list of resources that refer to this internal range. Resources that use the internal range for their range allocation are referred to as users of the range. Other resources mark themselves as users while doing so by creating a reference to this internal range. Having a user, based on this reference, prevents deletion of the internal range referred to. Can be empty.
	Users []string `json:"users,omitempty"`
}

// +kcc:proto=mockgcp.cloud.networkconnectivity.v1.LinkedInterconnectAttachments
type LinkedInterconnectAttachments struct {
	// Optional. IP ranges allowed to be included during import from hub.(does not control transit connectivity) The only allowed value for now is "ALL_IPV4_RANGES".
	IncludeImportRanges []string `json:"includeImportRanges,omitempty"`

	// A value that controls whether site-to-site data transfer is enabled for these resources. Data transfer is available only in [supported locations](https://cloud.google.com/network-connectivity/docs/network-connectivity-center/concepts/locations).
	SiteToSiteDataTransfer *bool `json:"siteToSiteDataTransfer,omitempty"`

	// The URIs of linked interconnect attachment resources
	Uris []string `json:"uris,omitempty"`

	// Output only. The VPC network where these VLAN attachments are located.
	VpcNetwork *string `json:"vpcNetwork,omitempty"`
}

// +kcc:proto=mockgcp.cloud.networkconnectivity.v1.LinkedRouterApplianceInstances
type LinkedRouterApplianceInstances struct {
	// Optional. IP ranges allowed to be included during import from hub.(does not control transit connectivity) The only allowed value for now is "ALL_IPV4_RANGES".
	IncludeImportRanges []string `json:"includeImportRanges,omitempty"`

	// The list of router appliance instances.
	Instances []RouterApplianceInstance `json:"instances,omitempty"`

	// A value that controls whether site-to-site data transfer is enabled for these resources. Data transfer is available only in [supported locations](https://cloud.google.com/network-connectivity/docs/network-connectivity-center/concepts/locations).
	SiteToSiteDataTransfer *bool `json:"siteToSiteDataTransfer,omitempty"`

	// Output only. The VPC network where these router appliance instances are located.
	VpcNetwork *string `json:"vpcNetwork,omitempty"`
}

// +kcc:proto=mockgcp.cloud.networkconnectivity.v1.LinkedVpcNetwork
type LinkedVpcNetwork struct {
	// Optional. IP ranges encompassing the subnets to be excluded from peering.
	ExcludeExportRanges []string `json:"excludeExportRanges,omitempty"`

	// Optional. IP ranges allowed to be included from peering.
	IncludeExportRanges []string `json:"includeExportRanges,omitempty"`

	// Required. The URI of the VPC network resource.
	Uri *string `json:"uri,omitempty"`
}

// +kcc:proto=mockgcp.cloud.networkconnectivity.v1.LinkedVpnTunnels
type LinkedVpnTunnels struct {
	// Optional. IP ranges allowed to be included during import from hub.(does not control transit connectivity) The only allowed value for now is "ALL_IPV4_RANGES".
	IncludeImportRanges []string `json:"includeImportRanges,omitempty"`

	// A value that controls whether site-to-site data transfer is enabled for these resources. Data transfer is available only in [supported locations](https://cloud.google.com/network-connectivity/docs/network-connectivity-center/concepts/locations).
	SiteToSiteDataTransfer *bool `json:"siteToSiteDataTransfer,omitempty"`

	// The URIs of linked VPN tunnel resources.
	Uris []string `json:"uris,omitempty"`

	// Output only. The VPC network where these VPN tunnels are located.
	VpcNetwork *string `json:"vpcNetwork,omitempty"`
}

// +kcc:proto=mockgcp.cloud.networkconnectivity.v1.Location
type Location struct {
	// The friendly name for this location, typically a nearby city name. For example, "Tokyo".
	DisplayName *string `json:"displayName,omitempty"`

	// Cross-service attributes for the location. For example {"cloud.googleapis.com/region": "us-east1"}
	Labels map[string]string `json:"labels,omitempty"`

	// The canonical id for this location. For example: `"us-east1"`.
	LocationID *string `json:"locationID,omitempty"`

	// TODO: map type string message for metadata

	// Resource name for the location, which may vary between implementations. For example: `"projects/example-project/locations/us-east1"`
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=mockgcp.cloud.networkconnectivity.v1.NextHopInterconnectAttachment
type NextHopInterconnectAttachment struct {
	// Indicates whether site-to-site data transfer is allowed for this interconnect attachment resource. Data transfer is available only in [supported locations](https://cloud.google.com/network-connectivity/docs/network-connectivity-center/concepts/locations).
	SiteToSiteDataTransfer *bool `json:"siteToSiteDataTransfer,omitempty"`

	// The URI of the interconnect attachment resource.
	Uri *string `json:"uri,omitempty"`

	// The VPC network where this interconnect attachment is located.
	VpcNetwork *string `json:"vpcNetwork,omitempty"`
}

// +kcc:proto=mockgcp.cloud.networkconnectivity.v1.NextHopRouterApplianceInstance
type NextHopRouterApplianceInstance struct {
	// Indicates whether site-to-site data transfer is allowed for this Router appliance instance resource. Data transfer is available only in [supported locations](https://cloud.google.com/network-connectivity/docs/network-connectivity-center/concepts/locations).
	SiteToSiteDataTransfer *bool `json:"siteToSiteDataTransfer,omitempty"`

	// The URI of the Router appliance instance.
	Uri *string `json:"uri,omitempty"`

	// The VPC network where this VM is located.
	VpcNetwork *string `json:"vpcNetwork,omitempty"`
}

// +kcc:proto=mockgcp.cloud.networkconnectivity.v1.NextHopVPNTunnel
type NextHopVPNTunnel struct {
	// Indicates whether site-to-site data transfer is allowed for this VPN tunnel resource. Data transfer is available only in [supported locations](https://cloud.google.com/network-connectivity/docs/network-connectivity-center/concepts/locations).
	SiteToSiteDataTransfer *bool `json:"siteToSiteDataTransfer,omitempty"`

	// The URI of the VPN tunnel resource.
	Uri *string `json:"uri,omitempty"`

	// The VPC network where this VPN tunnel is located.
	VpcNetwork *string `json:"vpcNetwork,omitempty"`
}

// +kcc:proto=mockgcp.cloud.networkconnectivity.v1.NextHopVpcNetwork
type NextHopVpcNetwork struct {
	// The URI of the VPC network resource
	Uri *string `json:"uri,omitempty"`
}

// +kcc:proto=mockgcp.cloud.networkconnectivity.v1.Policy
type Policy struct {
	// Specifies cloud audit logging configuration for this policy.
	AuditConfigs []AuditConfig `json:"auditConfigs,omitempty"`

	// Associates a list of `members`, or principals, with a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one principal. The `bindings` in a `Policy` can refer to up to 1,500 principals; up to 250 of these principals can be Google groups. Each occurrence of a principal counts towards these limits. For example, if the `bindings` grant 50 different roles to `user:alice@example.com`, and not to any other principal, then you can add another 1,450 principals to the `bindings` in the `Policy`.
	Bindings []Binding `json:"bindings,omitempty"`

	// `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
	Etag []byte `json:"etag,omitempty"`

	// Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Version *int32 `json:"version,omitempty"`
}

// +kcc:proto=mockgcp.cloud.networkconnectivity.v1.PolicyBasedRoute
type PolicyBasedRoute struct {
	// Output only. Time when the policy-based route was created.
	CreateTime *string `json:"createTime,omitempty"`

	// Optional. An optional description of this resource. Provide this field when you create the resource.
	Description *string `json:"description,omitempty"`

	// Required. The filter to match L4 traffic.
	Filter *Filter `json:"filter,omitempty"`

	// Optional. The interconnect attachments that this policy-based route applies to.
	InterconnectAttachment *InterconnectAttachment `json:"interconnectAttachment,omitempty"`

	// Output only. Type of this resource. Always networkconnectivity#policyBasedRoute for policy-based Route resources.
	Kind *string `json:"kind,omitempty"`

	// User-defined labels.
	Labels map[string]string `json:"labels,omitempty"`

	// Immutable. A unique name of the resource in the form of `projects/{project_number}/locations/global/PolicyBasedRoutes/{policy_based_route_id}`
	Name *string `json:"name,omitempty"`

	// Required. Fully-qualified URL of the network that this route applies to, for example: projects/my-project/global/networks/my-network.
	Network *string `json:"network,omitempty"`

	// Optional. The IP address of a global-access-enabled L4 ILB that is the next hop for matching packets. For this version, only nextHopIlbIp is supported.
	NextHopIlbIp *string `json:"nextHopIlbIp,omitempty"`

	// Optional. Other routes that will be referenced to determine the next hop of the packet.
	NextHopOtherRoutes *string `json:"nextHopOtherRoutes,omitempty"`

	// Optional. The priority of this policy-based route. Priority is used to break ties in cases where there are more than one matching policy-based routes found. In cases where multiple policy-based routes are matched, the one with the lowest-numbered priority value wins. The default value is 1000. The priority value must be from 1 to 65535, inclusive.
	Priority *int32 `json:"priority,omitempty"`

	// Output only. Server-defined fully-qualified URL for this resource.
	SelfLink *string `json:"selfLink,omitempty"`

	// Output only. Time when the policy-based route was updated.
	UpdateTime *string `json:"updateTime,omitempty"`

	// Optional. VM instances that this policy-based route applies to.
	VirtualMachine *VirtualMachine `json:"virtualMachine,omitempty"`

	// Output only. If potential misconfigurations are detected for this route, this field will be populated with warning messages.
	Warnings []Warnings `json:"warnings,omitempty"`
}

// +kcc:proto=mockgcp.cloud.networkconnectivity.v1.ProducerPscConfig
type ProducerPscConfig struct {
	// The resource path of a service attachment. Example: projects/{projectNumOrId}/regions/{region}/serviceAttachments/{resourceId}.
	ServiceAttachmentUri *string `json:"serviceAttachmentUri,omitempty"`
}

// +kcc:proto=mockgcp.cloud.networkconnectivity.v1.PscConnection
type PscConnection struct {
	// The resource reference of the consumer address.
	ConsumerAddress *string `json:"consumerAddress,omitempty"`

	// The resource reference of the PSC Forwarding Rule within the consumer VPC.
	ConsumerForwardingRule *string `json:"consumerForwardingRule,omitempty"`

	// The project where the PSC connection is created.
	ConsumerTargetProject *string `json:"consumerTargetProject,omitempty"`

	// The most recent error during operating this connection.
	Error *GoogleRpcStatus `json:"error,omitempty"`

	// Output only. The error info for the latest error during operating this connection.
	ErrorInfo *GoogleRpcErrorInfo `json:"errorInfo,omitempty"`

	// The error type indicates whether the error is consumer facing, producer facing or system internal.
	ErrorType *string `json:"errorType,omitempty"`

	// The last Compute Engine operation to setup PSC connection.
	GceOperation *string `json:"gceOperation,omitempty"`

	// Immutable. An immutable identifier for the producer instance.
	ProducerInstanceID *string `json:"producerInstanceID,omitempty"`

	// The PSC connection id of the PSC forwarding rule.
	PscConnectionID *string `json:"pscConnectionID,omitempty"`

	// Output only. The URI of the subnetwork selected to allocate IP address for this connection.
	SelectedSubnetwork *string `json:"selectedSubnetwork,omitempty"`

	// State of the PSC Connection
	State *string `json:"state,omitempty"`
}

// +kcc:proto=mockgcp.cloud.networkconnectivity.v1.RegionalEndpoint
type RegionalEndpoint struct {
	// Required. The access type of this regional endpoint. This field is reflected in the PSC Forwarding Rule configuration to enable global access.
	AccessType *string `json:"accessType,omitempty"`

	// Optional. The IP Address of the Regional Endpoint. When no address is provided, an IP from the subnetwork is allocated. Use one of the following formats: * IPv4 address as in `10.0.0.1` * Address resource URI as in `projects/{project}/regions/{region}/addresses/{address_name}`
	Address *string `json:"address,omitempty"`

	// Output only. Time when the RegionalEndpoint was created.
	CreateTime *string `json:"createTime,omitempty"`

	// Optional. A description of this resource.
	Description *string `json:"description,omitempty"`

	// Output only. The literal IP address of the PSC Forwarding Rule created on behalf of the customer. This field is deprecated. Use address instead.
	IpAddress *string `json:"ipAddress,omitempty"`

	// User-defined labels.
	Labels map[string]string `json:"labels,omitempty"`

	// Output only. The name of a RegionalEndpoint. Format: `projects/{project}/locations/{location}/regionalEndpoints/{regional_endpoint}`.
	Name *string `json:"name,omitempty"`

	// The name of the VPC network for this private regional endpoint. Format: `projects/{project}/global/networks/{network}`
	Network *string `json:"network,omitempty"`

	// Output only. The resource reference of the PSC Forwarding Rule created on behalf of the customer. Format: `//compute.googleapis.com/projects/{project}/regions/{region}/forwardingRules/{forwarding_rule_name}`
	PscForwardingRule *string `json:"pscForwardingRule,omitempty"`

	// The name of the subnetwork from which the IP address will be allocated. Format: `projects/{project}/regions/{region}/subnetworks/{subnetwork}`
	Subnetwork *string `json:"subnetwork,omitempty"`

	// Required. The service endpoint this private regional endpoint connects to. Format: `{apiname}.{region}.p.rep.googleapis.com` Example: "cloudkms.us-central1.p.rep.googleapis.com".
	TargetGoogleApi *string `json:"targetGoogleApi,omitempty"`

	// Output only. Time when the RegionalEndpoint was updated.
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=mockgcp.cloud.networkconnectivity.v1.Route
type Route struct {
	// Output only. The time the route was created.
	CreateTime *string `json:"createTime,omitempty"`

	// An optional description of the route.
	Description *string `json:"description,omitempty"`

	// The destination IP address range.
	IpCidrRange *string `json:"ipCidrRange,omitempty"`

	// Optional labels in key-value pair format. For more information about labels, see [Requirements for labels](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements).
	Labels map[string]string `json:"labels,omitempty"`

	// Output only. The origin location of the route. Uses the following form: "projects/{project}/locations/{location}" Example: projects/1234/locations/us-central1
	Location *string `json:"location,omitempty"`

	// Immutable. The name of the route. Route names must be unique. Route names use the following form: `projects/{project_number}/locations/global/hubs/{hub}/routeTables/{route_table_id}/routes/{route_id}`
	Name *string `json:"name,omitempty"`

	// Immutable. The next-hop VLAN attachment for packets on this route.
	NextHopInterconnectAttachment *NextHopInterconnectAttachment `json:"nextHopInterconnectAttachment,omitempty"`

	// Immutable. The next-hop Router appliance instance for packets on this route.
	NextHopRouterApplianceInstance *NextHopRouterApplianceInstance `json:"nextHopRouterApplianceInstance,omitempty"`

	// Immutable. The destination VPC network for packets on this route.
	NextHopVpcNetwork *NextHopVpcNetwork `json:"nextHopVpcNetwork,omitempty"`

	// Immutable. The next-hop VPN tunnel for packets on this route.
	NextHopVpnTunnel *NextHopVPNTunnel `json:"nextHopVpnTunnel,omitempty"`

	// Output only. The priority of this route. Priority is used to break ties in cases where a destination matches more than one route. In these cases the route with the lowest-numbered priority value wins.
	Priority *int64 `json:"priority,omitempty"`

	// Immutable. The spoke that this route leads to. Example: projects/12345/locations/global/spokes/SPOKE
	Spoke *string `json:"spoke,omitempty"`

	// Output only. The current lifecycle state of the route.
	State *string `json:"state,omitempty"`

	// Output only. The route's type. Its type is determined by the properties of its IP address range.
	Type *string `json:"type,omitempty"`

	// Output only. The Google-generated UUID for the route. This value is unique across all Network Connectivity Center route resources. If a route is deleted and another with the same name is created, the new route is assigned a different `uid`.
	Uid *string `json:"uid,omitempty"`

	// Output only. The time the route was last updated.
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=mockgcp.cloud.networkconnectivity.v1.RouteTable
type RouteTable struct {
	// Output only. The time the route table was created.
	CreateTime *string `json:"createTime,omitempty"`

	// An optional description of the route table.
	Description *string `json:"description,omitempty"`

	// Optional labels in key-value pair format. For more information about labels, see [Requirements for labels](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements).
	Labels map[string]string `json:"labels,omitempty"`

	// Immutable. The name of the route table. Route table names must be unique. They use the following form: `projects/{project_number}/locations/global/hubs/{hub}/routeTables/{route_table_id}`
	Name *string `json:"name,omitempty"`

	// Output only. The current lifecycle state of this route table.
	State *string `json:"state,omitempty"`

	// Output only. The Google-generated UUID for the route table. This value is unique across all route table resources. If a route table is deleted and another with the same name is created, the new route table is assigned a different `uid`.
	Uid *string `json:"uid,omitempty"`

	// Output only. The time the route table was last updated.
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=mockgcp.cloud.networkconnectivity.v1.RouterApplianceInstance
type RouterApplianceInstance struct {
	// The IP address on the VM to use for peering.
	IpAddress *string `json:"ipAddress,omitempty"`

	// The URI of the VM.
	VirtualMachine *string `json:"virtualMachine,omitempty"`
}

// +kcc:proto=mockgcp.cloud.networkconnectivity.v1.RoutingVPC
type RoutingVPC struct {
	// Output only. If true, indicates that this VPC network is currently associated with spokes that use the data transfer feature (spokes where the site_to_site_data_transfer field is set to true). If you create new spokes that use data transfer, they must be associated with this VPC network. At most, one VPC network will have this field set to true.
	RequiredForNewSiteToSiteDataTransferSpokes *bool `json:"requiredForNewSiteToSiteDataTransferSpokes,omitempty"`

	// The URI of the VPC network.
	Uri *string `json:"uri,omitempty"`
}

// +kcc:proto=mockgcp.cloud.networkconnectivity.v1.ServiceClass
type ServiceClass struct {
	// Output only. Time when the ServiceClass was created.
	CreateTime *string `json:"createTime,omitempty"`

	// A description of this resource.
	Description *string `json:"description,omitempty"`

	// Optional. The etag is computed by the server, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
	Etag *string `json:"etag,omitempty"`

	// User-defined labels.
	Labels map[string]string `json:"labels,omitempty"`

	// Immutable. The name of a ServiceClass resource. Format: projects/{project}/locations/{location}/serviceClasses/{service_class} See: https://google.aip.dev/122#fields-representing-resource-names
	Name *string `json:"name,omitempty"`

	// Output only. The generated service class name. Use this name to refer to the Service class in Service Connection Maps and Service Connection Policies.
	ServiceClass *string `json:"serviceClass,omitempty"`

	// Output only. Time when the ServiceClass was updated.
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=mockgcp.cloud.networkconnectivity.v1.ServiceConnectionMap
type ServiceConnectionMap struct {
	// The PSC configurations on consumer side.
	ConsumerPscConfigs []ConsumerPscConfig `json:"consumerPscConfigs,omitempty"`

	// Output only. PSC connection details on consumer side.
	ConsumerPscConnections []ConsumerPscConnection `json:"consumerPscConnections,omitempty"`

	// Output only. Time when the ServiceConnectionMap was created.
	CreateTime *string `json:"createTime,omitempty"`

	// A description of this resource.
	Description *string `json:"description,omitempty"`

	// Optional. The etag is computed by the server, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
	Etag *string `json:"etag,omitempty"`

	// Output only. The infrastructure used for connections between consumers/producers.
	Infrastructure *string `json:"infrastructure,omitempty"`

	// User-defined labels.
	Labels map[string]string `json:"labels,omitempty"`

	// Immutable. The name of a ServiceConnectionMap. Format: projects/{project}/locations/{location}/serviceConnectionMaps/{service_connection_map} See: https://google.aip.dev/122#fields-representing-resource-names
	Name *string `json:"name,omitempty"`

	// The PSC configurations on producer side.
	ProducerPscConfigs []ProducerPscConfig `json:"producerPscConfigs,omitempty"`

	// The service class identifier this ServiceConnectionMap is for. The user of ServiceConnectionMap create API needs to have networkconnecitivty.serviceclasses.use iam permission for the service class.
	ServiceClass *string `json:"serviceClass,omitempty"`

	// Output only. The service class uri this ServiceConnectionMap is for.
	ServiceClassUri *string `json:"serviceClassUri,omitempty"`

	// The token provided by the consumer. This token authenticates that the consumer can create a connection within the specified project and network.
	Token *string `json:"token,omitempty"`

	// Output only. Time when the ServiceConnectionMap was updated.
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=mockgcp.cloud.networkconnectivity.v1.ServiceConnectionPolicy
type ServiceConnectionPolicy struct {
	// Output only. Time when the ServiceConnectionMap was created.
	CreateTime *string `json:"createTime,omitempty"`

	// A description of this resource.
	Description *string `json:"description,omitempty"`

	// Optional. The etag is computed by the server, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
	Etag *string `json:"etag,omitempty"`

	// Output only. The type of underlying resources used to create the connection.
	Infrastructure *string `json:"infrastructure,omitempty"`

	// User-defined labels.
	Labels map[string]string `json:"labels,omitempty"`

	// Immutable. The name of a ServiceConnectionPolicy. Format: projects/{project}/locations/{location}/serviceConnectionPolicies/{service_connection_policy} See: https://google.aip.dev/122#fields-representing-resource-names
	Name *string `json:"name,omitempty"`

	// The resource path of the consumer network. Example: - projects/{projectNumOrId}/global/networks/{resourceId}.
	Network *string `json:"network,omitempty"`

	// Configuration used for Private Service Connect connections. Used when Infrastructure is PSC.
	PscConfig *PscConfig `json:"pscConfig,omitempty"`

	// Output only. [Output only] Information about each Private Service Connect connection.
	PscConnections []PscConnection `json:"pscConnections,omitempty"`

	// The service class identifier for which this ServiceConnectionPolicy is for. The service class identifier is a unique, symbolic representation of a ServiceClass. It is provided by the Service Producer. Google services have a prefix of gcp. For example, gcp-cloud-sql. 3rd party services do not. For example, test-service-a3dfcx.
	ServiceClass *string `json:"serviceClass,omitempty"`

	// Output only. Time when the ServiceConnectionMap was updated.
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=mockgcp.cloud.networkconnectivity.v1.ServiceConnectionToken
type ServiceConnectionToken struct {
	// Output only. Time when the ServiceConnectionToken was created.
	CreateTime *string `json:"createTime,omitempty"`

	// A description of this resource.
	Description *string `json:"description,omitempty"`

	// Optional. The etag is computed by the server, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
	Etag *string `json:"etag,omitempty"`

	// Output only. The time to which this token is valid.
	ExpireTime *string `json:"expireTime,omitempty"`

	// User-defined labels.
	Labels map[string]string `json:"labels,omitempty"`

	// Immutable. The name of a ServiceConnectionToken. Format: projects/{project}/locations/{location}/ServiceConnectionTokens/{service_connection_token} See: https://google.aip.dev/122#fields-representing-resource-names
	Name *string `json:"name,omitempty"`

	// The resource path of the network associated with this token. Example: projects/{projectNumOrId}/global/networks/{resourceId}.
	Network *string `json:"network,omitempty"`

	// Output only. The token generated by Automation.
	Token *string `json:"token,omitempty"`

	// Output only. Time when the ServiceConnectionToken was updated.
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=mockgcp.cloud.networkconnectivity.v1.Spoke
type Spoke struct {
	// Output only. The time the spoke was created.
	CreateTime *string `json:"createTime,omitempty"`

	// An optional description of the spoke.
	Description *string `json:"description,omitempty"`

	// Optional. The name of the group that this spoke is associated with.
	Group *string `json:"group,omitempty"`

	// Immutable. The name of the hub that this spoke is attached to.
	Hub *string `json:"hub,omitempty"`

	// Optional labels in key-value pair format. For more information about labels, see [Requirements for labels](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements).
	Labels map[string]string `json:"labels,omitempty"`

	// VLAN attachments that are associated with the spoke.
	LinkedInterconnectAttachments *LinkedInterconnectAttachments `json:"linkedInterconnectAttachments,omitempty"`

	// Router appliance instances that are associated with the spoke.
	LinkedRouterApplianceInstances *LinkedRouterApplianceInstances `json:"linkedRouterApplianceInstances,omitempty"`

	// Optional. VPC network that is associated with the spoke.
	LinkedVpcNetwork *LinkedVpcNetwork `json:"linkedVpcNetwork,omitempty"`

	// VPN tunnels that are associated with the spoke.
	LinkedVpnTunnels *LinkedVpnTunnels `json:"linkedVpnTunnels,omitempty"`

	// Immutable. The name of the spoke. Spoke names must be unique. They use the following form: `projects/{project_number}/locations/{region}/spokes/{spoke_id}`
	Name *string `json:"name,omitempty"`

	// Output only. The reasons for current state of the spoke. Only present when the spoke is in the `INACTIVE` state.
	Reasons []StateReason `json:"reasons,omitempty"`

	// Output only. The type of resource associated with the spoke.
	SpokeType *string `json:"spokeType,omitempty"`

	// Output only. The current lifecycle state of this spoke.
	State *string `json:"state,omitempty"`

	// Output only. The Google-generated UUID for the spoke. This value is unique across all spoke resources. If a spoke is deleted and another with the same name is created, the new spoke is assigned a different `unique_id`.
	UniqueID *string `json:"uniqueID,omitempty"`

	// Output only. The time the spoke was last updated.
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=mockgcp.cloud.networkconnectivity.v1.SpokeStateCount
type SpokeStateCount struct {
	// Output only. The total number of spokes that are in this state and associated with a given hub.
	Count *int64 `json:"count,omitempty"`

	// Output only. The state of the spokes.
	State *string `json:"state,omitempty"`
}

// +kcc:proto=mockgcp.cloud.networkconnectivity.v1.SpokeStateReasonCount
type SpokeStateReasonCount struct {
	// Output only. The total number of spokes that are inactive for a particular reason and associated with a given hub.
	Count *int64 `json:"count,omitempty"`

	// Output only. The reason that a spoke is inactive.
	StateReasonCode *string `json:"stateReasonCode,omitempty"`
}

// +kcc:proto=mockgcp.cloud.networkconnectivity.v1.SpokeSummary
type SpokeSummary struct {
	// Output only. Counts the number of spokes that are in each state and associated with a given hub.
	SpokeStateCounts []SpokeStateCount `json:"spokeStateCounts,omitempty"`

	// Output only. Counts the number of spokes that are inactive for each possible reason and associated with a given hub.
	SpokeStateReasonCounts []SpokeStateReasonCount `json:"spokeStateReasonCounts,omitempty"`

	// Output only. Counts the number of spokes of each type that are associated with a specific hub.
	SpokeTypeCounts []SpokeTypeCount `json:"spokeTypeCounts,omitempty"`
}

// +kcc:proto=mockgcp.cloud.networkconnectivity.v1.SpokeTypeCount
type SpokeTypeCount struct {
	// Output only. The total number of spokes of this type that are associated with the hub.
	Count *int64 `json:"count,omitempty"`

	// Output only. The type of the spokes.
	SpokeType *string `json:"spokeType,omitempty"`
}

// +kcc:proto=mockgcp.cloud.networkconnectivity.v1.StateReason
type StateReason struct {
	// The code associated with this reason.
	Code *string `json:"code,omitempty"`

	// Human-readable details about this reason.
	Message *string `json:"message,omitempty"`

	// Additional information provided by the user in the RejectSpoke call.
	UserDetails *string `json:"userDetails,omitempty"`
}

// +kcc:proto=mockgcp.cloud.networkconnectivity.v1.VirtualMachine
type VirtualMachine struct {
	// Optional. A list of VM instance tags that this policy-based route applies to. VM instances that have ANY of tags specified here installs this PBR.
	Tags []string `json:"tags,omitempty"`
}

// +kcc:proto=mockgcp.cloud.networkconnectivity.v1.Warnings
type Warnings struct {
	// Output only. A warning code, if applicable.
	Code *string `json:"code,omitempty"`

	// Output only. Metadata about this warning in key: value format. The key should provides more detail on the warning being returned. For example, for warnings where there are no results in a list request for a particular zone, this key might be scope and the key value might be the zone name. Other examples might be a key indicating a deprecated resource and a suggested replacement.
	Data map[string]string `json:"data,omitempty"`

	// Output only. A human-readable description of the warning code.
	WarningMessage *string `json:"warningMessage,omitempty"`
}

// +kcc:proto=mockgcp.cloud.networkconnectivity.v1.Migration
type Migration struct {
	// Immutable. Resource path as an URI of the source resource, for example a subnet. The project for the source resource should match the project for the InternalRange. An example: /projects/{project}/regions/{region}/subnetworks/{subnet}
	// +kcc:proto:field=mockgcp.cloud.networkconnectivity.v1.Migration.source
	Source *string `json:"source,omitempty"`

	// Immutable. Resource path of the target resource. The target project can be different, as in the cases when migrating to peer networks. The resource For example: /projects/{project}/regions/{region}/subnetworks/{subnet}
	// +kcc:proto:field=mockgcp.cloud.networkconnectivity.v1.Migration.target
	Target *string `json:"target,omitempty"`
}
