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
// krm.group: networkmanagement.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.networkmanagement.v1
// resource: NetworkManagementConnectivityTest:ConnectivityTest

package v1alpha1

// +kcc:proto=google.cloud.networkmanagement.v1.AbortInfo
type AbortInfo struct {
	// Causes that the analysis is aborted.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.AbortInfo.cause
	Cause *string `json:"cause,omitempty"`

	// URI of the resource that caused the abort.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.AbortInfo.resource_uri
	ResourceURI *string `json:"resourceURI,omitempty"`

	// IP address that caused the abort.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.AbortInfo.ip_address
	IPAddress *string `json:"ipAddress,omitempty"`

	// List of project IDs the user specified in the request but lacks access to.
	//  In this case, analysis is aborted with the PERMISSION_DENIED cause.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.AbortInfo.projects_missing_permission
	ProjectsMissingPermission []string `json:"projectsMissingPermission,omitempty"`
}

// +kcc:proto=google.cloud.networkmanagement.v1.AppEngineVersionInfo
type AppEngineVersionInfo struct {
	// Name of an App Engine version.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.AppEngineVersionInfo.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// URI of an App Engine version.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.AppEngineVersionInfo.uri
	URI *string `json:"uri,omitempty"`

	// Runtime of the App Engine version.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.AppEngineVersionInfo.runtime
	Runtime *string `json:"runtime,omitempty"`

	// App Engine execution environment for a version.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.AppEngineVersionInfo.environment
	Environment *string `json:"environment,omitempty"`
}

// +kcc:proto=google.cloud.networkmanagement.v1.CloudFunctionInfo
type CloudFunctionInfo struct {
	// Name of a Cloud Function.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.CloudFunctionInfo.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// URI of a Cloud Function.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.CloudFunctionInfo.uri
	URI *string `json:"uri,omitempty"`

	// Location in which the Cloud Function is deployed.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.CloudFunctionInfo.location
	Location *string `json:"location,omitempty"`

	// Latest successfully deployed version id of the Cloud Function.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.CloudFunctionInfo.version_id
	VersionID *int64 `json:"versionID,omitempty"`
}

// +kcc:proto=google.cloud.networkmanagement.v1.CloudRunRevisionInfo
type CloudRunRevisionInfo struct {
	// Name of a Cloud Run revision.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.CloudRunRevisionInfo.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// URI of a Cloud Run revision.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.CloudRunRevisionInfo.uri
	URI *string `json:"uri,omitempty"`

	// Location in which this revision is deployed.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.CloudRunRevisionInfo.location
	Location *string `json:"location,omitempty"`

	// URI of Cloud Run service this revision belongs to.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.CloudRunRevisionInfo.service_uri
	ServiceURI *string `json:"serviceURI,omitempty"`
}

// +kcc:proto=google.cloud.networkmanagement.v1.CloudSQLInstanceInfo
type CloudSQLInstanceInfo struct {
	// Name of a Cloud SQL instance.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.CloudSQLInstanceInfo.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// URI of a Cloud SQL instance.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.CloudSQLInstanceInfo.uri
	URI *string `json:"uri,omitempty"`

	// URI of a Cloud SQL instance network or empty string if the instance does
	//  not have one.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.CloudSQLInstanceInfo.network_uri
	NetworkURI *string `json:"networkURI,omitempty"`

	// Internal IP address of a Cloud SQL instance.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.CloudSQLInstanceInfo.internal_ip
	InternalIP *string `json:"internalIP,omitempty"`

	// External IP address of a Cloud SQL instance.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.CloudSQLInstanceInfo.external_ip
	ExternalIP *string `json:"externalIP,omitempty"`

	// Region in which the Cloud SQL instance is running.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.CloudSQLInstanceInfo.region
	Region *string `json:"region,omitempty"`
}

// +kcc:proto=google.cloud.networkmanagement.v1.DeliverInfo
type DeliverInfo struct {
	// Target type where the packet is delivered to.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.DeliverInfo.target
	Target *string `json:"target,omitempty"`

	// URI of the resource that the packet is delivered to.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.DeliverInfo.resource_uri
	ResourceURI *string `json:"resourceURI,omitempty"`

	// IP address of the target (if applicable).
	// +kcc:proto:field=google.cloud.networkmanagement.v1.DeliverInfo.ip_address
	IPAddress *string `json:"ipAddress,omitempty"`

	// Name of the Cloud Storage Bucket the packet is delivered to (if
	//  applicable).
	// +kcc:proto:field=google.cloud.networkmanagement.v1.DeliverInfo.storage_bucket
	StorageBucket *string `json:"storageBucket,omitempty"`

	// PSC Google API target the packet is delivered to (if applicable).
	// +kcc:proto:field=google.cloud.networkmanagement.v1.DeliverInfo.psc_google_api_target
	PSCGoogleAPITarget *string `json:"pscGoogleAPITarget,omitempty"`
}

// +kcc:proto=google.cloud.networkmanagement.v1.DropInfo
type DropInfo struct {
	// Cause that the packet is dropped.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.DropInfo.cause
	Cause *string `json:"cause,omitempty"`

	// URI of the resource that caused the drop.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.DropInfo.resource_uri
	ResourceURI *string `json:"resourceURI,omitempty"`

	// Source IP address of the dropped packet (if relevant).
	// +kcc:proto:field=google.cloud.networkmanagement.v1.DropInfo.source_ip
	SourceIP *string `json:"sourceIP,omitempty"`

	// Destination IP address of the dropped packet (if relevant).
	// +kcc:proto:field=google.cloud.networkmanagement.v1.DropInfo.destination_ip
	DestinationIP *string `json:"destinationIP,omitempty"`

	// Region of the dropped packet (if relevant).
	// +kcc:proto:field=google.cloud.networkmanagement.v1.DropInfo.region
	Region *string `json:"region,omitempty"`
}

// +kcc:proto=google.cloud.networkmanagement.v1.Endpoint
type Endpoint struct {
	// The IP address of the endpoint, which can be an external or internal IP.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Endpoint.ip_address
	IPAddress *string `json:"ipAddress,omitempty"`

	// The IP protocol port of the endpoint.
	//  Only applicable when protocol is TCP or UDP.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Endpoint.port
	Port *int32 `json:"port,omitempty"`

	// A Compute Engine instance URI.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Endpoint.instance
	Instance *string `json:"instance,omitempty"`

	// A forwarding rule and its corresponding IP address represent the frontend
	//  configuration of a Google Cloud load balancer. Forwarding rules are also
	//  used for protocol forwarding, Private Service Connect and other network
	//  services to provide forwarding information in the control plane. Format:
	//   projects/{project}/global/forwardingRules/{id} or
	//   projects/{project}/regions/{region}/forwardingRules/{id}
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Endpoint.forwarding_rule
	ForwardingRule *string `json:"forwardingRule,omitempty"`

	// A cluster URI for [Google Kubernetes Engine cluster control
	//  plane](https://cloud.google.com/kubernetes-engine/docs/concepts/cluster-architecture).
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Endpoint.gke_master_cluster
	GkeMasterCluster *string `json:"gkeMasterCluster,omitempty"`

	// DNS endpoint of [Google Kubernetes Engine cluster control
	//  plane](https://cloud.google.com/kubernetes-engine/docs/concepts/cluster-architecture).
	//  Requires gke_master_cluster to be set, can't be used simultaneoulsly with
	//  ip_address or network. Applicable only to destination endpoint.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Endpoint.fqdn
	Fqdn *string `json:"fqdn,omitempty"`

	// A [Cloud SQL](https://cloud.google.com/sql) instance URI.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Endpoint.cloud_sql_instance
	CloudSQLInstance *string `json:"cloudSQLInstance,omitempty"`

	// A [Redis Instance](https://cloud.google.com/memorystore/docs/redis)
	//  URI.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Endpoint.redis_instance
	RedisInstance *string `json:"redisInstance,omitempty"`

	// A [Redis Cluster](https://cloud.google.com/memorystore/docs/cluster)
	//  URI.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Endpoint.redis_cluster
	RedisCluster *string `json:"redisCluster,omitempty"`

	// A [Cloud Function](https://cloud.google.com/functions).
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Endpoint.cloud_function
	CloudFunction *Endpoint_CloudFunctionEndpoint `json:"cloudFunction,omitempty"`

	// An [App Engine](https://cloud.google.com/appengine) [service
	//  version](https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions).
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Endpoint.app_engine_version
	AppEngineVersion *Endpoint_AppEngineVersionEndpoint `json:"appEngineVersion,omitempty"`

	// A [Cloud Run](https://cloud.google.com/run)
	//  [revision](https://cloud.google.com/run/docs/reference/rest/v1/namespaces.revisions/get)
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Endpoint.cloud_run_revision
	CloudRunRevision *Endpoint_CloudRunRevisionEndpoint `json:"cloudRunRevision,omitempty"`

	// A Compute Engine network URI.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Endpoint.network
	Network *string `json:"network,omitempty"`

	// Type of the network where the endpoint is located.
	//  Applicable only to source endpoint, as destination network type can be
	//  inferred from the source.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Endpoint.network_type
	NetworkType *string `json:"networkType,omitempty"`

	// Project ID where the endpoint is located.
	//  The Project ID can be derived from the URI if you provide a VM instance or
	//  network URI.
	//  The following are two cases where you must provide the project ID:
	//  1. Only the IP address is specified, and the IP address is within a Google
	//  Cloud project.
	//  2. When you are using Shared VPC and the IP address that you provide is
	//  from the service project. In this case, the network that the IP address
	//  resides in is defined in the host project.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Endpoint.project_id
	ProjectID *string `json:"projectID,omitempty"`
}

// +kcc:proto=google.cloud.networkmanagement.v1.Endpoint.AppEngineVersionEndpoint
type Endpoint_AppEngineVersionEndpoint struct {
	// An [App Engine](https://cloud.google.com/appengine) [service
	//  version](https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions)
	//  name.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Endpoint.AppEngineVersionEndpoint.uri
	URI *string `json:"uri,omitempty"`
}

// +kcc:proto=google.cloud.networkmanagement.v1.Endpoint.CloudFunctionEndpoint
type Endpoint_CloudFunctionEndpoint struct {
	// A [Cloud Function](https://cloud.google.com/functions) name.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Endpoint.CloudFunctionEndpoint.uri
	URI *string `json:"uri,omitempty"`
}

// +kcc:proto=google.cloud.networkmanagement.v1.Endpoint.CloudRunRevisionEndpoint
type Endpoint_CloudRunRevisionEndpoint struct {
	// A [Cloud Run](https://cloud.google.com/run)
	//  [revision](https://cloud.google.com/run/docs/reference/rest/v1/namespaces.revisions/get)
	//  URI. The format is:
	//  projects/{project}/locations/{location}/revisions/{revision}
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Endpoint.CloudRunRevisionEndpoint.uri
	URI *string `json:"uri,omitempty"`
}

// +kcc:proto=google.cloud.networkmanagement.v1.EndpointInfo
type EndpointInfo struct {
	// Source IP address.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.EndpointInfo.source_ip
	SourceIP *string `json:"sourceIP,omitempty"`

	// Destination IP address.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.EndpointInfo.destination_ip
	DestinationIP *string `json:"destinationIP,omitempty"`

	// IP protocol in string format, for example: "TCP", "UDP", "ICMP".
	// +kcc:proto:field=google.cloud.networkmanagement.v1.EndpointInfo.protocol
	Protocol *string `json:"protocol,omitempty"`

	// Source port. Only valid when protocol is TCP or UDP.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.EndpointInfo.source_port
	SourcePort *int32 `json:"sourcePort,omitempty"`

	// Destination port. Only valid when protocol is TCP or UDP.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.EndpointInfo.destination_port
	DestinationPort *int32 `json:"destinationPort,omitempty"`

	// URI of the network where this packet originates from.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.EndpointInfo.source_network_uri
	SourceNetworkURI *string `json:"sourceNetworkURI,omitempty"`

	// URI of the network where this packet is sent to.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.EndpointInfo.destination_network_uri
	DestinationNetworkURI *string `json:"destinationNetworkURI,omitempty"`

	// URI of the source telemetry agent this packet originates from.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.EndpointInfo.source_agent_uri
	SourceAgentURI *string `json:"sourceAgentURI,omitempty"`
}

// +kcc:proto=google.cloud.networkmanagement.v1.FirewallInfo
type FirewallInfo struct {
	// The display name of the firewall rule. This field might be empty for
	//  firewall policy rules.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.FirewallInfo.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The URI of the firewall rule. This field is not applicable to implied
	//  VPC firewall rules.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.FirewallInfo.uri
	URI *string `json:"uri,omitempty"`

	// Possible values: INGRESS, EGRESS
	// +kcc:proto:field=google.cloud.networkmanagement.v1.FirewallInfo.direction
	Direction *string `json:"direction,omitempty"`

	// Possible values: ALLOW, DENY, APPLY_SECURITY_PROFILE_GROUP
	// +kcc:proto:field=google.cloud.networkmanagement.v1.FirewallInfo.action
	Action *string `json:"action,omitempty"`

	// The priority of the firewall rule.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.FirewallInfo.priority
	Priority *int32 `json:"priority,omitempty"`

	// The URI of the VPC network that the firewall rule is associated with.
	//  This field is not applicable to hierarchical firewall policy rules.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.FirewallInfo.network_uri
	NetworkURI *string `json:"networkURI,omitempty"`

	// The target tags defined by the VPC firewall rule. This field is not
	//  applicable to firewall policy rules.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.FirewallInfo.target_tags
	TargetTags []string `json:"targetTags,omitempty"`

	// The target service accounts specified by the firewall rule.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.FirewallInfo.target_service_accounts
	TargetServiceAccounts []string `json:"targetServiceAccounts,omitempty"`

	// The name of the firewall policy that this rule is associated with.
	//  This field is not applicable to VPC firewall rules and implied VPC firewall
	//  rules.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.FirewallInfo.policy
	Policy *string `json:"policy,omitempty"`

	// The URI of the firewall policy that this rule is associated with.
	//  This field is not applicable to VPC firewall rules and implied VPC firewall
	//  rules.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.FirewallInfo.policy_uri
	PolicyURI *string `json:"policyURI,omitempty"`

	// The firewall rule's type.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.FirewallInfo.firewall_rule_type
	FirewallRuleType *string `json:"firewallRuleType,omitempty"`
}

// +kcc:proto=google.cloud.networkmanagement.v1.ForwardInfo
type ForwardInfo struct {
	// Target type where this packet is forwarded to.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ForwardInfo.target
	Target *string `json:"target,omitempty"`

	// URI of the resource that the packet is forwarded to.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ForwardInfo.resource_uri
	ResourceURI *string `json:"resourceURI,omitempty"`

	// IP address of the target (if applicable).
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ForwardInfo.ip_address
	IPAddress *string `json:"ipAddress,omitempty"`
}

// +kcc:proto=google.cloud.networkmanagement.v1.ForwardingRuleInfo
type ForwardingRuleInfo struct {
	// Name of the forwarding rule.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ForwardingRuleInfo.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// URI of the forwarding rule.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ForwardingRuleInfo.uri
	URI *string `json:"uri,omitempty"`

	// Protocol defined in the forwarding rule that matches the packet.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ForwardingRuleInfo.matched_protocol
	MatchedProtocol *string `json:"matchedProtocol,omitempty"`

	// Port range defined in the forwarding rule that matches the packet.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ForwardingRuleInfo.matched_port_range
	MatchedPortRange *string `json:"matchedPortRange,omitempty"`

	// VIP of the forwarding rule.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ForwardingRuleInfo.vip
	Vip *string `json:"vip,omitempty"`

	// Target type of the forwarding rule.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ForwardingRuleInfo.target
	Target *string `json:"target,omitempty"`

	// Network URI.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ForwardingRuleInfo.network_uri
	NetworkURI *string `json:"networkURI,omitempty"`

	// Region of the forwarding rule. Set only for regional forwarding rules.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ForwardingRuleInfo.region
	Region *string `json:"region,omitempty"`

	// Name of the load balancer the forwarding rule belongs to. Empty for
	//  forwarding rules not related to load balancers (like PSC forwarding rules).
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ForwardingRuleInfo.load_balancer_name
	LoadBalancerName *string `json:"loadBalancerName,omitempty"`

	// URI of the PSC service attachment this forwarding rule targets (if
	//  applicable).
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ForwardingRuleInfo.psc_service_attachment_uri
	PSCServiceAttachmentURI *string `json:"pscServiceAttachmentURI,omitempty"`

	// PSC Google API target this forwarding rule targets (if applicable).
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ForwardingRuleInfo.psc_google_api_target
	PSCGoogleAPITarget *string `json:"pscGoogleAPITarget,omitempty"`
}

// +kcc:proto=google.cloud.networkmanagement.v1.GKEMasterInfo
type GkeMasterInfo struct {
	// URI of a GKE cluster.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.GKEMasterInfo.cluster_uri
	ClusterURI *string `json:"clusterURI,omitempty"`

	// URI of a GKE cluster network.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.GKEMasterInfo.cluster_network_uri
	ClusterNetworkURI *string `json:"clusterNetworkURI,omitempty"`

	// Internal IP address of a GKE cluster control plane.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.GKEMasterInfo.internal_ip
	InternalIP *string `json:"internalIP,omitempty"`

	// External IP address of a GKE cluster control plane.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.GKEMasterInfo.external_ip
	ExternalIP *string `json:"externalIP,omitempty"`

	// DNS endpoint of a GKE cluster control plane.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.GKEMasterInfo.dns_endpoint
	DNSEndpoint *string `json:"dnsEndpoint,omitempty"`
}

// +kcc:proto=google.cloud.networkmanagement.v1.GoogleServiceInfo
type GoogleServiceInfo struct {
	// Source IP address.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.GoogleServiceInfo.source_ip
	SourceIP *string `json:"sourceIP,omitempty"`

	// Recognized type of a Google Service.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.GoogleServiceInfo.google_service_type
	GoogleServiceType *string `json:"googleServiceType,omitempty"`
}

// +kcc:proto=google.cloud.networkmanagement.v1.InstanceInfo
type InstanceInfo struct {
	// Name of a Compute Engine instance.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.InstanceInfo.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// URI of a Compute Engine instance.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.InstanceInfo.uri
	URI *string `json:"uri,omitempty"`

	// Name of the network interface of a Compute Engine instance.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.InstanceInfo.interface
	Interface *string `json:"interface,omitempty"`

	// URI of a Compute Engine network.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.InstanceInfo.network_uri
	NetworkURI *string `json:"networkURI,omitempty"`

	// Internal IP address of the network interface.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.InstanceInfo.internal_ip
	InternalIP *string `json:"internalIP,omitempty"`

	// External IP address of the network interface.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.InstanceInfo.external_ip
	ExternalIP *string `json:"externalIP,omitempty"`

	// Network tags configured on the instance.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.InstanceInfo.network_tags
	NetworkTags []string `json:"networkTags,omitempty"`

	// Service account authorized for the instance.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.InstanceInfo.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// URI of the PSC network attachment the NIC is attached to (if relevant).
	// +kcc:proto:field=google.cloud.networkmanagement.v1.InstanceInfo.psc_network_attachment_uri
	PSCNetworkAttachmentURI *string `json:"pscNetworkAttachmentURI,omitempty"`
}

// +kcc:proto=google.cloud.networkmanagement.v1.LatencyDistribution
type LatencyDistribution struct {
	// Representative latency percentiles.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.LatencyDistribution.latency_percentiles
	LatencyPercentiles []LatencyPercentile `json:"latencyPercentiles,omitempty"`
}

// +kcc:proto=google.cloud.networkmanagement.v1.LatencyPercentile
type LatencyPercentile struct {
	// Percentage of samples this data point applies to.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.LatencyPercentile.percent
	Percent *int32 `json:"percent,omitempty"`

	// percent-th percentile of latency observed, in microseconds.
	//  Fraction of percent/100 of samples have latency lower or
	//  equal to the value of this field.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.LatencyPercentile.latency_micros
	LatencyMicros *int64 `json:"latencyMicros,omitempty"`
}

// +kcc:proto=google.cloud.networkmanagement.v1.LoadBalancerBackend
type LoadBalancerBackend struct {
	// Name of a Compute Engine instance or network endpoint.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.LoadBalancerBackend.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// URI of a Compute Engine instance or network endpoint.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.LoadBalancerBackend.uri
	URI *string `json:"uri,omitempty"`

	// State of the health check firewall configuration.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.LoadBalancerBackend.health_check_firewall_state
	HealthCheckFirewallState *string `json:"healthCheckFirewallState,omitempty"`

	// A list of firewall rule URIs allowing probes from health check IP ranges.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.LoadBalancerBackend.health_check_allowing_firewall_rules
	HealthCheckAllowingFirewallRules []string `json:"healthCheckAllowingFirewallRules,omitempty"`

	// A list of firewall rule URIs blocking probes from health check IP ranges.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.LoadBalancerBackend.health_check_blocking_firewall_rules
	HealthCheckBlockingFirewallRules []string `json:"healthCheckBlockingFirewallRules,omitempty"`
}

// +kcc:proto=google.cloud.networkmanagement.v1.LoadBalancerBackendInfo
type LoadBalancerBackendInfo struct {
	// Display name of the backend. For example, it might be an instance name for
	//  the instance group backends, or an IP address and port for zonal network
	//  endpoint group backends.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.LoadBalancerBackendInfo.name
	Name *string `json:"name,omitempty"`

	// URI of the backend instance (if applicable). Populated for instance group
	//  backends, and zonal NEG backends.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.LoadBalancerBackendInfo.instance_uri
	InstanceURI *string `json:"instanceURI,omitempty"`

	// URI of the backend service this backend belongs to (if applicable).
	// +kcc:proto:field=google.cloud.networkmanagement.v1.LoadBalancerBackendInfo.backend_service_uri
	BackendServiceURI *string `json:"backendServiceURI,omitempty"`

	// URI of the instance group this backend belongs to (if applicable).
	// +kcc:proto:field=google.cloud.networkmanagement.v1.LoadBalancerBackendInfo.instance_group_uri
	InstanceGroupURI *string `json:"instanceGroupURI,omitempty"`

	// URI of the network endpoint group this backend belongs to (if applicable).
	// +kcc:proto:field=google.cloud.networkmanagement.v1.LoadBalancerBackendInfo.network_endpoint_group_uri
	NetworkEndpointGroupURI *string `json:"networkEndpointGroupURI,omitempty"`

	// URI of the backend bucket this backend targets (if applicable).
	// +kcc:proto:field=google.cloud.networkmanagement.v1.LoadBalancerBackendInfo.backend_bucket_uri
	BackendBucketURI *string `json:"backendBucketURI,omitempty"`

	// URI of the PSC service attachment this PSC NEG backend targets (if
	//  applicable).
	// +kcc:proto:field=google.cloud.networkmanagement.v1.LoadBalancerBackendInfo.psc_service_attachment_uri
	PSCServiceAttachmentURI *string `json:"pscServiceAttachmentURI,omitempty"`

	// PSC Google API target this PSC NEG backend targets (if applicable).
	// +kcc:proto:field=google.cloud.networkmanagement.v1.LoadBalancerBackendInfo.psc_google_api_target
	PSCGoogleAPITarget *string `json:"pscGoogleAPITarget,omitempty"`

	// URI of the health check attached to this backend (if applicable).
	// +kcc:proto:field=google.cloud.networkmanagement.v1.LoadBalancerBackendInfo.health_check_uri
	HealthCheckURI *string `json:"healthCheckURI,omitempty"`
}

// +kcc:proto=google.cloud.networkmanagement.v1.LoadBalancerInfo
type LoadBalancerInfo struct {
	// Type of the load balancer.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.LoadBalancerInfo.load_balancer_type
	LoadBalancerType *string `json:"loadBalancerType,omitempty"`

	// URI of the health check for the load balancer. Deprecated and no longer
	//  populated as different load balancer backends might have different health
	//  checks.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.LoadBalancerInfo.health_check_uri
	HealthCheckURI *string `json:"healthCheckURI,omitempty"`

	// Information for the loadbalancer backends.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.LoadBalancerInfo.backends
	Backends []LoadBalancerBackend `json:"backends,omitempty"`

	// Type of load balancer's backend configuration.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.LoadBalancerInfo.backend_type
	BackendType *string `json:"backendType,omitempty"`

	// Backend configuration URI.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.LoadBalancerInfo.backend_uri
	BackendURI *string `json:"backendURI,omitempty"`
}

// +kcc:proto=google.cloud.networkmanagement.v1.NatInfo
type NATInfo struct {
	// Type of NAT.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.NatInfo.type
	Type *string `json:"type,omitempty"`

	// IP protocol in string format, for example: "TCP", "UDP", "ICMP".
	// +kcc:proto:field=google.cloud.networkmanagement.v1.NatInfo.protocol
	Protocol *string `json:"protocol,omitempty"`

	// URI of the network where NAT translation takes place.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.NatInfo.network_uri
	NetworkURI *string `json:"networkURI,omitempty"`

	// Source IP address before NAT translation.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.NatInfo.old_source_ip
	OldSourceIP *string `json:"oldSourceIP,omitempty"`

	// Source IP address after NAT translation.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.NatInfo.new_source_ip
	NewSourceIP *string `json:"newSourceIP,omitempty"`

	// Destination IP address before NAT translation.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.NatInfo.old_destination_ip
	OldDestinationIP *string `json:"oldDestinationIP,omitempty"`

	// Destination IP address after NAT translation.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.NatInfo.new_destination_ip
	NewDestinationIP *string `json:"newDestinationIP,omitempty"`

	// Source port before NAT translation. Only valid when protocol is TCP or UDP.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.NatInfo.old_source_port
	OldSourcePort *int32 `json:"oldSourcePort,omitempty"`

	// Source port after NAT translation. Only valid when protocol is TCP or UDP.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.NatInfo.new_source_port
	NewSourcePort *int32 `json:"newSourcePort,omitempty"`

	// Destination port before NAT translation. Only valid when protocol is TCP or
	//  UDP.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.NatInfo.old_destination_port
	OldDestinationPort *int32 `json:"oldDestinationPort,omitempty"`

	// Destination port after NAT translation. Only valid when protocol is TCP or
	//  UDP.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.NatInfo.new_destination_port
	NewDestinationPort *int32 `json:"newDestinationPort,omitempty"`

	// Uri of the Cloud Router. Only valid when type is CLOUD_NAT.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.NatInfo.router_uri
	RouterURI *string `json:"routerURI,omitempty"`

	// The name of Cloud NAT Gateway. Only valid when type is CLOUD_NAT.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.NatInfo.nat_gateway_name
	NATGatewayName *string `json:"natGatewayName,omitempty"`
}

// +kcc:proto=google.cloud.networkmanagement.v1.NetworkInfo
type NetworkInfo struct {
	// Name of a Compute Engine network.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.NetworkInfo.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// URI of a Compute Engine network.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.NetworkInfo.uri
	URI *string `json:"uri,omitempty"`

	// URI of the subnet matching the source IP address of the test.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.NetworkInfo.matched_subnet_uri
	MatchedSubnetURI *string `json:"matchedSubnetURI,omitempty"`

	// The IP range of the subnet matching the source IP address of the test.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.NetworkInfo.matched_ip_range
	MatchedIPRange *string `json:"matchedIPRange,omitempty"`

	// The region of the subnet matching the source IP address of the test.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.NetworkInfo.region
	Region *string `json:"region,omitempty"`
}

// +kcc:proto=google.cloud.networkmanagement.v1.ProbingDetails
type ProbingDetails struct {
	// The overall result of active probing.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ProbingDetails.result
	Result *string `json:"result,omitempty"`

	// The time that reachability was assessed through active probing.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ProbingDetails.verify_time
	VerifyTime *string `json:"verifyTime,omitempty"`

	// Details about an internal failure or the cancellation of active probing.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ProbingDetails.error
	Error *Status `json:"error,omitempty"`

	// The reason probing was aborted.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ProbingDetails.abort_cause
	AbortCause *string `json:"abortCause,omitempty"`

	// Number of probes sent.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ProbingDetails.sent_probe_count
	SentProbeCount *int32 `json:"sentProbeCount,omitempty"`

	// Number of probes that reached the destination.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ProbingDetails.successful_probe_count
	SuccessfulProbeCount *int32 `json:"successfulProbeCount,omitempty"`

	// The source and destination endpoints derived from the test input and used
	//  for active probing.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ProbingDetails.endpoint_info
	EndpointInfo *EndpointInfo `json:"endpointInfo,omitempty"`

	// Latency as measured by active probing in one direction:
	//  from the source to the destination endpoint.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ProbingDetails.probing_latency
	ProbingLatency *LatencyDistribution `json:"probingLatency,omitempty"`

	// The EdgeLocation from which a packet destined for/originating from the
	//  internet will egress/ingress the Google network.
	//  This will only be populated for a connectivity test which has an internet
	//  destination/source address.
	//  The absence of this field *must not* be used as an indication that the
	//  destination/source is part of the Google network.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ProbingDetails.destination_egress_location
	DestinationEgressLocation *ProbingDetails_EdgeLocation `json:"destinationEgressLocation,omitempty"`
}

// +kcc:proto=google.cloud.networkmanagement.v1.ProbingDetails.EdgeLocation
type ProbingDetails_EdgeLocation struct {
	// Name of the metropolitan area.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ProbingDetails.EdgeLocation.metropolitan_area
	MetropolitanArea *string `json:"metropolitanArea,omitempty"`
}

// +kcc:proto=google.cloud.networkmanagement.v1.ProxyConnectionInfo
type ProxyConnectionInfo struct {
	// IP protocol in string format, for example: "TCP", "UDP", "ICMP".
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ProxyConnectionInfo.protocol
	Protocol *string `json:"protocol,omitempty"`

	// Source IP address of an original connection.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ProxyConnectionInfo.old_source_ip
	OldSourceIP *string `json:"oldSourceIP,omitempty"`

	// Source IP address of a new connection.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ProxyConnectionInfo.new_source_ip
	NewSourceIP *string `json:"newSourceIP,omitempty"`

	// Destination IP address of an original connection
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ProxyConnectionInfo.old_destination_ip
	OldDestinationIP *string `json:"oldDestinationIP,omitempty"`

	// Destination IP address of a new connection.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ProxyConnectionInfo.new_destination_ip
	NewDestinationIP *string `json:"newDestinationIP,omitempty"`

	// Source port of an original connection. Only valid when protocol is TCP or
	//  UDP.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ProxyConnectionInfo.old_source_port
	OldSourcePort *int32 `json:"oldSourcePort,omitempty"`

	// Source port of a new connection. Only valid when protocol is TCP or UDP.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ProxyConnectionInfo.new_source_port
	NewSourcePort *int32 `json:"newSourcePort,omitempty"`

	// Destination port of an original connection. Only valid when protocol is TCP
	//  or UDP.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ProxyConnectionInfo.old_destination_port
	OldDestinationPort *int32 `json:"oldDestinationPort,omitempty"`

	// Destination port of a new connection. Only valid when protocol is TCP or
	//  UDP.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ProxyConnectionInfo.new_destination_port
	NewDestinationPort *int32 `json:"newDestinationPort,omitempty"`

	// Uri of proxy subnet.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ProxyConnectionInfo.subnet_uri
	SubnetURI *string `json:"subnetURI,omitempty"`

	// URI of the network where connection is proxied.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ProxyConnectionInfo.network_uri
	NetworkURI *string `json:"networkURI,omitempty"`
}

// +kcc:proto=google.cloud.networkmanagement.v1.ReachabilityDetails
type ReachabilityDetails struct {
	// The overall result of the test's configuration analysis.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ReachabilityDetails.result
	Result *string `json:"result,omitempty"`

	// The time of the configuration analysis.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ReachabilityDetails.verify_time
	VerifyTime *string `json:"verifyTime,omitempty"`

	// The details of a failure or a cancellation of reachability analysis.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ReachabilityDetails.error
	Error *Status `json:"error,omitempty"`

	// Result may contain a list of traces if a test has multiple possible
	//  paths in the network, such as when destination endpoint is a load balancer
	//  with multiple backends.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ReachabilityDetails.traces
	Traces []Trace `json:"traces,omitempty"`
}

// +kcc:proto=google.cloud.networkmanagement.v1.RedisClusterInfo
type RedisClusterInfo struct {
	// Name of a Redis Cluster.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.RedisClusterInfo.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// URI of a Redis Cluster in format
	//  "projects/{project_id}/locations/{location}/clusters/{cluster_id}"
	// +kcc:proto:field=google.cloud.networkmanagement.v1.RedisClusterInfo.uri
	URI *string `json:"uri,omitempty"`

	// URI of a Redis Cluster network in format
	//  "projects/{project_id}/global/networks/{network_id}".
	// +kcc:proto:field=google.cloud.networkmanagement.v1.RedisClusterInfo.network_uri
	NetworkURI *string `json:"networkURI,omitempty"`

	// Discovery endpoint IP address of a Redis Cluster.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.RedisClusterInfo.discovery_endpoint_ip_address
	DiscoveryEndpointIPAddress *string `json:"discoveryEndpointIPAddress,omitempty"`

	// Secondary endpoint IP address of a Redis Cluster.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.RedisClusterInfo.secondary_endpoint_ip_address
	SecondaryEndpointIPAddress *string `json:"secondaryEndpointIPAddress,omitempty"`

	// Name of the region in which the Redis Cluster is defined. For example,
	//  "us-central1".
	// +kcc:proto:field=google.cloud.networkmanagement.v1.RedisClusterInfo.location
	Location *string `json:"location,omitempty"`
}

// +kcc:proto=google.cloud.networkmanagement.v1.RedisInstanceInfo
type RedisInstanceInfo struct {
	// Name of a Cloud Redis Instance.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.RedisInstanceInfo.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// URI of a Cloud Redis Instance.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.RedisInstanceInfo.uri
	URI *string `json:"uri,omitempty"`

	// URI of a Cloud Redis Instance network.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.RedisInstanceInfo.network_uri
	NetworkURI *string `json:"networkURI,omitempty"`

	// Primary endpoint IP address of a Cloud Redis Instance.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.RedisInstanceInfo.primary_endpoint_ip
	PrimaryEndpointIP *string `json:"primaryEndpointIP,omitempty"`

	// Read endpoint IP address of a Cloud Redis Instance (if applicable).
	// +kcc:proto:field=google.cloud.networkmanagement.v1.RedisInstanceInfo.read_endpoint_ip
	ReadEndpointIP *string `json:"readEndpointIP,omitempty"`

	// Region in which the Cloud Redis Instance is defined.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.RedisInstanceInfo.region
	Region *string `json:"region,omitempty"`
}

// +kcc:proto=google.cloud.networkmanagement.v1.RouteInfo
type RouteInfo struct {
	// Type of route.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.RouteInfo.route_type
	RouteType *string `json:"routeType,omitempty"`

	// Type of next hop.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.RouteInfo.next_hop_type
	NextHopType *string `json:"nextHopType,omitempty"`

	// Indicates where route is applicable.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.RouteInfo.route_scope
	RouteScope *string `json:"routeScope,omitempty"`

	// Name of a route.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.RouteInfo.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// URI of a route (if applicable).
	// +kcc:proto:field=google.cloud.networkmanagement.v1.RouteInfo.uri
	URI *string `json:"uri,omitempty"`

	// Region of the route (if applicable).
	// +kcc:proto:field=google.cloud.networkmanagement.v1.RouteInfo.region
	Region *string `json:"region,omitempty"`

	// Destination IP range of the route.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.RouteInfo.dest_ip_range
	DestIPRange *string `json:"destIPRange,omitempty"`

	// Next hop of the route.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.RouteInfo.next_hop
	NextHop *string `json:"nextHop,omitempty"`

	// URI of a Compute Engine network. NETWORK routes only.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.RouteInfo.network_uri
	NetworkURI *string `json:"networkURI,omitempty"`

	// Priority of the route.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.RouteInfo.priority
	Priority *int32 `json:"priority,omitempty"`

	// Instance tags of the route.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.RouteInfo.instance_tags
	InstanceTags []string `json:"instanceTags,omitempty"`

	// Source IP address range of the route. Policy based routes only.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.RouteInfo.src_ip_range
	SrcIPRange *string `json:"srcIPRange,omitempty"`

	// Destination port ranges of the route. Policy based routes only.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.RouteInfo.dest_port_ranges
	DestPortRanges []string `json:"destPortRanges,omitempty"`

	// Source port ranges of the route. Policy based routes only.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.RouteInfo.src_port_ranges
	SrcPortRanges []string `json:"srcPortRanges,omitempty"`

	// Protocols of the route. Policy based routes only.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.RouteInfo.protocols
	Protocols []string `json:"protocols,omitempty"`

	// URI of a NCC Hub. NCC_HUB routes only.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.RouteInfo.ncc_hub_uri
	NccHubURI *string `json:"nccHubURI,omitempty"`

	// URI of a NCC Spoke. NCC_HUB routes only.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.RouteInfo.ncc_spoke_uri
	NccSpokeURI *string `json:"nccSpokeURI,omitempty"`

	// For advertised dynamic routes, the URI of the Cloud Router that advertised
	//  the corresponding IP prefix.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.RouteInfo.advertised_route_source_router_uri
	AdvertisedRouteSourceRouterURI *string `json:"advertisedRouteSourceRouterURI,omitempty"`

	// For advertised routes, the URI of their next hop, i.e. the URI of the
	//  hybrid endpoint (VPN tunnel, Interconnect attachment, NCC router appliance)
	//  the advertised prefix is advertised through, or URI of the source peered
	//  network.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.RouteInfo.advertised_route_next_hop_uri
	AdvertisedRouteNextHopURI *string `json:"advertisedRouteNextHopURI,omitempty"`
}

// +kcc:proto=google.cloud.networkmanagement.v1.ServerlessNegInfo
type ServerlessNegInfo struct {
	// URI of the serverless network endpoint group.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ServerlessNegInfo.neg_uri
	NegURI *string `json:"negURI,omitempty"`
}

// +kcc:proto=google.cloud.networkmanagement.v1.Step
type Step struct {
	// A description of the step. Usually this is a summary of the state.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Step.description
	Description *string `json:"description,omitempty"`

	// Each step is in one of the pre-defined states.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Step.state
	State *string `json:"state,omitempty"`

	// This is a step that leads to the final state Drop.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Step.causes_drop
	CausesDrop *bool `json:"causesDrop,omitempty"`

	// Project ID that contains the configuration this step is validating.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Step.project_id
	ProjectID *string `json:"projectID,omitempty"`

	// Display information of a Compute Engine instance.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Step.instance
	Instance *InstanceInfo `json:"instance,omitempty"`

	// Display information of a Compute Engine firewall rule.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Step.firewall
	Firewall *FirewallInfo `json:"firewall,omitempty"`

	// Display information of a Compute Engine route.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Step.route
	Route *RouteInfo `json:"route,omitempty"`

	// Display information of the source and destination under analysis.
	//  The endpoint information in an intermediate state may differ with the
	//  initial input, as it might be modified by state like NAT,
	//  or Connection Proxy.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Step.endpoint
	Endpoint *EndpointInfo `json:"endpoint,omitempty"`

	// Display information of a Google service
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Step.google_service
	GoogleService *GoogleServiceInfo `json:"googleService,omitempty"`

	// Display information of a Compute Engine forwarding rule.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Step.forwarding_rule
	ForwardingRule *ForwardingRuleInfo `json:"forwardingRule,omitempty"`

	// Display information of a Compute Engine VPN gateway.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Step.vpn_gateway
	VpnGateway *VpnGatewayInfo `json:"vpnGateway,omitempty"`

	// Display information of a Compute Engine VPN tunnel.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Step.vpn_tunnel
	VpnTunnel *VpnTunnelInfo `json:"vpnTunnel,omitempty"`

	// Display information of a VPC connector.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Step.vpc_connector
	VpcConnector *VpcConnectorInfo `json:"vpcConnector,omitempty"`

	// Display information of the final state "deliver" and reason.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Step.deliver
	Deliver *DeliverInfo `json:"deliver,omitempty"`

	// Display information of the final state "forward" and reason.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Step.forward
	Forward *ForwardInfo `json:"forward,omitempty"`

	// Display information of the final state "abort" and reason.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Step.abort
	Abort *AbortInfo `json:"abort,omitempty"`

	// Display information of the final state "drop" and reason.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Step.drop
	Drop *DropInfo `json:"drop,omitempty"`

	// Display information of the load balancers. Deprecated in favor of the
	//  `load_balancer_backend_info` field, not used in new tests.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Step.load_balancer
	LoadBalancer *LoadBalancerInfo `json:"loadBalancer,omitempty"`

	// Display information of a Google Cloud network.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Step.network
	Network *NetworkInfo `json:"network,omitempty"`

	// Display information of a Google Kubernetes Engine cluster master.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Step.gke_master
	GkeMaster *GkeMasterInfo `json:"gkeMaster,omitempty"`

	// Display information of a Cloud SQL instance.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Step.cloud_sql_instance
	CloudSQLInstance *CloudSQLInstanceInfo `json:"cloudSQLInstance,omitempty"`

	// Display information of a Redis Instance.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Step.redis_instance
	RedisInstance *RedisInstanceInfo `json:"redisInstance,omitempty"`

	// Display information of a Redis Cluster.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Step.redis_cluster
	RedisCluster *RedisClusterInfo `json:"redisCluster,omitempty"`

	// Display information of a Cloud Function.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Step.cloud_function
	CloudFunction *CloudFunctionInfo `json:"cloudFunction,omitempty"`

	// Display information of an App Engine service version.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Step.app_engine_version
	AppEngineVersion *AppEngineVersionInfo `json:"appEngineVersion,omitempty"`

	// Display information of a Cloud Run revision.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Step.cloud_run_revision
	CloudRunRevision *CloudRunRevisionInfo `json:"cloudRunRevision,omitempty"`

	// Display information of a NAT.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Step.nat
	NAT *NATInfo `json:"nat,omitempty"`

	// Display information of a ProxyConnection.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Step.proxy_connection
	ProxyConnection *ProxyConnectionInfo `json:"proxyConnection,omitempty"`

	// Display information of a specific load balancer backend.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Step.load_balancer_backend_info
	LoadBalancerBackendInfo *LoadBalancerBackendInfo `json:"loadBalancerBackendInfo,omitempty"`

	// Display information of a Storage Bucket. Used only for return traces.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Step.storage_bucket
	StorageBucket *StorageBucketInfo `json:"storageBucket,omitempty"`

	// Display information of a Serverless network endpoint group backend. Used
	//  only for return traces.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Step.serverless_neg
	ServerlessNeg *ServerlessNegInfo `json:"serverlessNeg,omitempty"`
}

// +kcc:proto=google.cloud.networkmanagement.v1.StorageBucketInfo
type StorageBucketInfo struct {
	// Cloud Storage Bucket name.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.StorageBucketInfo.bucket
	Bucket *string `json:"bucket,omitempty"`
}

// +kcc:proto=google.cloud.networkmanagement.v1.Trace
type Trace struct {
	// Derived from the source and destination endpoints definition specified by
	//  user request, and validated by the data plane model.
	//  If there are multiple traces starting from different source locations, then
	//  the endpoint_info may be different between traces.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Trace.endpoint_info
	EndpointInfo *EndpointInfo `json:"endpointInfo,omitempty"`

	// A trace of a test contains multiple steps from the initial state to the
	//  final state (delivered, dropped, forwarded, or aborted).
	//
	//  The steps are ordered by the processing sequence within the simulated
	//  network state machine. It is critical to preserve the order of the steps
	//  and avoid reordering or sorting them.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Trace.steps
	Steps []Step `json:"steps,omitempty"`

	// ID of trace. For forward traces, this ID is unique for each trace. For
	//  return traces, it matches ID of associated forward trace. A single forward
	//  trace can be associated with none, one or more than one return trace.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Trace.forward_trace_id
	ForwardTraceID *int32 `json:"forwardTraceID,omitempty"`
}

// +kcc:proto=google.cloud.networkmanagement.v1.VpcConnectorInfo
type VpcConnectorInfo struct {
	// Name of a VPC connector.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.VpcConnectorInfo.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// URI of a VPC connector.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.VpcConnectorInfo.uri
	URI *string `json:"uri,omitempty"`

	// Location in which the VPC connector is deployed.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.VpcConnectorInfo.location
	Location *string `json:"location,omitempty"`
}

// +kcc:proto=google.cloud.networkmanagement.v1.VpnGatewayInfo
type VpnGatewayInfo struct {
	// Name of a VPN gateway.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.VpnGatewayInfo.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// URI of a VPN gateway.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.VpnGatewayInfo.uri
	URI *string `json:"uri,omitempty"`

	// URI of a Compute Engine network where the VPN gateway is configured.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.VpnGatewayInfo.network_uri
	NetworkURI *string `json:"networkURI,omitempty"`

	// IP address of the VPN gateway.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.VpnGatewayInfo.ip_address
	IPAddress *string `json:"ipAddress,omitempty"`

	// A VPN tunnel that is associated with this VPN gateway.
	//  There may be multiple VPN tunnels configured on a VPN gateway, and only
	//  the one relevant to the test is displayed.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.VpnGatewayInfo.vpn_tunnel_uri
	VpnTunnelURI *string `json:"vpnTunnelURI,omitempty"`

	// Name of a Google Cloud region where this VPN gateway is configured.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.VpnGatewayInfo.region
	Region *string `json:"region,omitempty"`
}

// +kcc:proto=google.cloud.networkmanagement.v1.VpnTunnelInfo
type VpnTunnelInfo struct {
	// Name of a VPN tunnel.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.VpnTunnelInfo.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// URI of a VPN tunnel.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.VpnTunnelInfo.uri
	URI *string `json:"uri,omitempty"`

	// URI of the VPN gateway at local end of the tunnel.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.VpnTunnelInfo.source_gateway
	SourceGateway *string `json:"sourceGateway,omitempty"`

	// URI of a VPN gateway at remote end of the tunnel.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.VpnTunnelInfo.remote_gateway
	RemoteGateway *string `json:"remoteGateway,omitempty"`

	// Remote VPN gateway's IP address.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.VpnTunnelInfo.remote_gateway_ip
	RemoteGatewayIP *string `json:"remoteGatewayIP,omitempty"`

	// Local VPN gateway's IP address.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.VpnTunnelInfo.source_gateway_ip
	SourceGatewayIP *string `json:"sourceGatewayIP,omitempty"`

	// URI of a Compute Engine network where the VPN tunnel is configured.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.VpnTunnelInfo.network_uri
	NetworkURI *string `json:"networkURI,omitempty"`

	// Name of a Google Cloud region where this VPN tunnel is configured.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.VpnTunnelInfo.region
	Region *string `json:"region,omitempty"`

	// Type of the routing policy.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.VpnTunnelInfo.routing_type
	RoutingType *string `json:"routingType,omitempty"`
}

// +kcc:proto=google.protobuf.Any
type Any struct {
	// A URL/resource name that uniquely identifies the type of the serialized
	//  protocol buffer message. This string must contain at least
	//  one "/" character. The last segment of the URL's path must represent
	//  the fully qualified name of the type (as in
	//  `path/google.protobuf.Duration`). The name should be in a canonical form
	//  (e.g., leading "." is not accepted).
	//
	//  In practice, teams usually precompile into the binary all types that they
	//  expect it to use in the context of Any. However, for URLs which use the
	//  scheme `http`, `https`, or no scheme, one can optionally set up a type
	//  server that maps type URLs to message definitions as follows:
	//
	//  * If no scheme is provided, `https` is assumed.
	//  * An HTTP GET on the URL must yield a [google.protobuf.Type][]
	//    value in binary format, or produce an error.
	//  * Applications are allowed to cache lookup results based on the
	//    URL, or have them precompiled into a binary to avoid any
	//    lookup. Therefore, binary compatibility needs to be preserved
	//    on changes to types. (Use versioned type names to manage
	//    breaking changes.)
	//
	//  Note: this functionality is not currently available in the official
	//  protobuf release, and it is not used for type URLs beginning with
	//  type.googleapis.com.
	//
	//  Schemes other than `http`, `https` (or the empty scheme) might be
	//  used with implementation specific semantics.
	// +kcc:proto:field=google.protobuf.Any.type_url
	TypeURL *string `json:"typeURL,omitempty"`

	// Must be a valid serialized protocol buffer of the above specified type.
	// +kcc:proto:field=google.protobuf.Any.value
	Value []byte `json:"value,omitempty"`
}

// +kcc:proto=google.rpc.Status
type Status struct {
	// The status code, which should be an enum value of
	//  [google.rpc.Code][google.rpc.Code].
	// +kcc:proto:field=google.rpc.Status.code
	Code *int32 `json:"code,omitempty"`

	// A developer-facing error message, which should be in English. Any
	//  user-facing error message should be localized and sent in the
	//  [google.rpc.Status.details][google.rpc.Status.details] field, or localized
	//  by the client.
	// +kcc:proto:field=google.rpc.Status.message
	Message *string `json:"message,omitempty"`

	// A list of messages that carry the error details.  There is a common set of
	//  message types for APIs to use.
	// +kcc:proto:field=google.rpc.Status.details
	Details []Any `json:"details,omitempty"`
}

// +kcc:proto=google.cloud.networkmanagement.v1.Endpoint
type EndpointObservedState struct {
	// Output only. Specifies the type of the target of the forwarding rule.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Endpoint.forwarding_rule_target
	ForwardingRuleTarget *string `json:"forwardingRuleTarget,omitempty"`

	// Output only. ID of the load balancer the forwarding rule points to. Empty
	//  for forwarding rules not related to load balancers.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Endpoint.load_balancer_id
	LoadBalancerID *string `json:"loadBalancerID,omitempty"`

	// Output only. Type of the load balancer the forwarding rule points to.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Endpoint.load_balancer_type
	LoadBalancerType *string `json:"loadBalancerType,omitempty"`
}

// +kcc:proto=google.cloud.networkmanagement.v1.LoadBalancerBackendInfo
type LoadBalancerBackendInfoObservedState struct {
	// Output only. Health check firewalls configuration state for the backend.
	//  This is a result of the static firewall analysis (verifying that health
	//  check traffic from required IP ranges to the backend is allowed or not).
	//  The backend might still be unhealthy even if these firewalls are
	//  configured. Please refer to the documentation for more information:
	//  https://cloud.google.com/load-balancing/docs/firewall-rules
	// +kcc:proto:field=google.cloud.networkmanagement.v1.LoadBalancerBackendInfo.health_check_firewalls_config_state
	HealthCheckFirewallsConfigState *string `json:"healthCheckFirewallsConfigState,omitempty"`
}

// +kcc:proto=google.cloud.networkmanagement.v1.ReachabilityDetails
type ReachabilityDetailsObservedState struct {
	// Result may contain a list of traces if a test has multiple possible
	//  paths in the network, such as when destination endpoint is a load balancer
	//  with multiple backends.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.ReachabilityDetails.traces
	Traces []TraceObservedState `json:"traces,omitempty"`
}

// +kcc:proto=google.cloud.networkmanagement.v1.Step
type StepObservedState struct {
	// Display information of a specific load balancer backend.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Step.load_balancer_backend_info
	LoadBalancerBackendInfo *LoadBalancerBackendInfoObservedState `json:"loadBalancerBackendInfo,omitempty"`
}

// +kcc:proto=google.cloud.networkmanagement.v1.Trace
type TraceObservedState struct {
	// A trace of a test contains multiple steps from the initial state to the
	//  final state (delivered, dropped, forwarded, or aborted).
	//
	//  The steps are ordered by the processing sequence within the simulated
	//  network state machine. It is critical to preserve the order of the steps
	//  and avoid reordering or sorting them.
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Trace.steps
	Steps []StepObservedState `json:"steps,omitempty"`
}
