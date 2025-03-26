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
	ComputeInstanceRef *string `json:"computeInstanceRef,omitempty"`

	// A forwarding rule and its corresponding IP address represent the frontend
	//  configuration of a Google Cloud load balancer. Forwarding rules are also
	//  used for protocol forwarding, Private Service Connect and other network
	//  services to provide forwarding information in the control plane. Format:
	//   projects/{project}/global/forwardingRules/{id} or
	//   projects/{project}/regions/{region}/forwardingRules/{id}
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Endpoint.forwarding_rule
	ComputeForwardingRuleRef *string `json:"computeForwardingRuleRef,omitempty"`

	// A cluster URI for [Google Kubernetes Engine cluster control
	//  plane](https://cloud.google.com/kubernetes-engine/docs/concepts/cluster-architecture).
	// +kcc:proto:field=google.cloud.networkmanagement.v1.Endpoint.gke_master_cluster
	ContainerClusterRef *string `json:"containerClusterRef,omitempty"`

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
