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

package v1beta1

/*
// +kcc:proto=google.cloud.workstations.v1.WorkstationCluster
type WorkstationCluster struct {
	// Full name of this workstation cluster.
	Name *string `json:"name,omitempty"`

	// Optional. Human-readable name for this workstation cluster.
	DisplayName *string `json:"displayName,omitempty"`

	// Output only. A system-assigned unique identifier for this workstation
	//  cluster.
	Uid *string `json:"uid,omitempty"`

	// Output only. Indicates whether this workstation cluster is currently being
	//  updated to match its intended state.
	Reconciling *bool `json:"reconciling,omitempty"`

	// Optional. Client-specified annotations.
	Annotations map[string]string `json:"annotations,omitempty"`

	// Optional.
	//  [Labels](https://cloud.google.com/workstations/docs/label-resources) that
	//  are applied to the workstation cluster and that are also propagated to the
	//  underlying Compute Engine resources.
	Labels map[string]string `json:"labels,omitempty"`

	// Output only. Time when this workstation cluster was created.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time when this workstation cluster was most recently updated.
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Time when this workstation cluster was soft-deleted.
	DeleteTime *string `json:"deleteTime,omitempty"`

	// Optional. Checksum computed by the server. May be sent on update and delete
	//  requests to make sure that the client has an up-to-date value before
	//  proceeding.
	Etag *string `json:"etag,omitempty"`

	// Immutable. Name of the Compute Engine network in which instances associated
	//  with this workstation cluster will be created.
	Network *string `json:"network,omitempty"`

	// Immutable. Name of the Compute Engine subnetwork in which instances
	//  associated with this workstation cluster will be created. Must be part of
	//  the subnetwork specified for this workstation cluster.
	Subnetwork *string `json:"subnetwork,omitempty"`

	// Output only. The private IP address of the control plane for this
	//  workstation cluster. Workstation VMs need access to this IP address to work
	//  with the service, so make sure that your firewall rules allow egress from
	//  the workstation VMs to this address.
	ControlPlaneIp *string `json:"controlPlaneIp,omitempty"`

	// Optional. Configuration for private workstation cluster.
	PrivateClusterConfig *WorkstationCluster_PrivateClusterConfig `json:"privateClusterConfig,omitempty"`

	// Output only. Whether this workstation cluster is in degraded mode, in which
	//  case it may require user action to restore full functionality. Details can
	//  be found in
	//  [conditions][google.cloud.workstations.v1.WorkstationCluster.conditions].
	Degraded *bool `json:"degraded,omitempty"`

	// Output only. Status conditions describing the workstation cluster's current
	//  state.
	Conditions []Status `json:"conditions,omitempty"`
}

// +kcc:proto=google.cloud.workstations.v1.WorkstationCluster.PrivateClusterConfig
type WorkstationCluster_PrivateClusterConfig struct {
	// Immutable. Whether Workstations endpoint is private.
	EnablePrivateEndpoint *bool `json:"enablePrivateEndpoint,omitempty"`

	// Output only. Hostname for the workstation cluster. This field will be
	//  populated only when private endpoint is enabled. To access workstations
	//  in the workstation cluster, create a new DNS zone mapping this domain
	//  name to an internal IP address and a forwarding rule mapping that address
	//  to the service attachment.
	ClusterHostname *string `json:"clusterHostname,omitempty"`

	// Output only. Service attachment URI for the workstation cluster. The
	//  service attachment is created when private endpoint is enabled. To access
	//  workstations in the workstation cluster, configure access to the managed
	//  service using [Private Service
	//  Connect](https://cloud.google.com/vpc/docs/configure-private-service-connect-services).
	ServiceAttachmentUri *string `json:"serviceAttachmentUri,omitempty"`

	// Optional. Additional projects that are allowed to attach to the
	//  workstation cluster's service attachment. By default, the workstation
	//  cluster's project and the VPC host project (if different) are allowed.
	AllowedProjects []string `json:"allowedProjects,omitempty"`
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
	TypeURL *string `json:"typeURL,omitempty"`

	// Must be a valid serialized protocol buffer of the above specified type.
	Value []byte `json:"value,omitempty"`
}

// +kcc:proto=google.rpc.Status
type Status struct {
	// The status code, which should be an enum value of
	//  [google.rpc.Code][google.rpc.Code].
	Code *int32 `json:"code,omitempty"`

	// A developer-facing error message, which should be in English. Any
	//  user-facing error message should be localized and sent in the
	//  [google.rpc.Status.details][google.rpc.Status.details] field, or localized
	//  by the client.
	Message *string `json:"message,omitempty"`

	// A list of messages that carry the error details.  There is a common set of
	//  message types for APIs to use.
	Details []Any `json:"details,omitempty"`
}
*/
