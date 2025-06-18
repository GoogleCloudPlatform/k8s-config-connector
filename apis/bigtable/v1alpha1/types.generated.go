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
// krm.group: bigtable.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.bigtable.admin.v2
// resource: BigtableAuthorizedView:AuthorizedView
// resource: BigtableBackup:Backup
// resource: BigtableCluster:Cluster
// resource: BigtableLogicalView:LogicalView

package v1alpha1

// +kcc:proto=google.bigtable.admin.v2.AuthorizedView.FamilySubsets
type AuthorizedView_FamilySubsets struct {
	// Individual exact column qualifiers to be included in the AuthorizedView.
	// +kcc:proto:field=google.bigtable.admin.v2.AuthorizedView.FamilySubsets.qualifiers
	Qualifiers [][]byte `json:"qualifiers,omitempty"`

	// Prefixes for qualifiers to be included in the AuthorizedView. Every
	//  qualifier starting with one of these prefixes is included in the
	//  AuthorizedView. To provide access to all qualifiers, include the empty
	//  string as a prefix
	//  ("").
	// +kcc:proto:field=google.bigtable.admin.v2.AuthorizedView.FamilySubsets.qualifier_prefixes
	QualifierPrefixes [][]byte `json:"qualifierPrefixes,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.AuthorizedView.SubsetView
type AuthorizedView_SubsetView struct {
	// Row prefixes to be included in the AuthorizedView.
	//  To provide access to all rows, include the empty string as a prefix ("").
	// +kcc:proto:field=google.bigtable.admin.v2.AuthorizedView.SubsetView.row_prefixes
	RowPrefixes [][]byte `json:"rowPrefixes,omitempty"`

	// TODO: unsupported map type with key string and value message

}

// +kcc:proto=google.bigtable.admin.v2.AutoscalingLimits
type AutoscalingLimits struct {
	// Required. Minimum number of nodes to scale down to.
	// +kcc:proto:field=google.bigtable.admin.v2.AutoscalingLimits.min_serve_nodes
	MinServeNodes *int32 `json:"minServeNodes,omitempty"`

	// Required. Maximum number of nodes to scale up to.
	// +kcc:proto:field=google.bigtable.admin.v2.AutoscalingLimits.max_serve_nodes
	MaxServeNodes *int32 `json:"maxServeNodes,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.Cluster.ClusterAutoscalingConfig
type Cluster_ClusterAutoscalingConfig struct {
	// Required. Autoscaling limits for this cluster.
	// +kcc:proto:field=google.bigtable.admin.v2.Cluster.ClusterAutoscalingConfig.autoscaling_limits
	AutoscalingLimits *AutoscalingLimits `json:"autoscalingLimits,omitempty"`

	// Required. Autoscaling targets for this cluster.
	// +kcc:proto:field=google.bigtable.admin.v2.Cluster.ClusterAutoscalingConfig.autoscaling_targets
	AutoscalingTargets *AutoscalingTargets `json:"autoscalingTargets,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.Cluster.ClusterConfig
type Cluster_ClusterConfig struct {
	// Autoscaling configuration for this cluster.
	// +kcc:proto:field=google.bigtable.admin.v2.Cluster.ClusterConfig.cluster_autoscaling_config
	ClusterAutoscalingConfig *Cluster_ClusterAutoscalingConfig `json:"clusterAutoscalingConfig,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.EncryptionInfo
type EncryptionInfo struct {
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

// +kcc:observedstate:proto=google.bigtable.admin.v2.EncryptionInfo
type EncryptionInfoObservedState struct {
	// Output only. The type of encryption used to protect this resource.
	// +kcc:proto:field=google.bigtable.admin.v2.EncryptionInfo.encryption_type
	EncryptionType *string `json:"encryptionType,omitempty"`

	// Output only. The status of encrypt/decrypt calls on underlying data for
	//  this resource. Regardless of status, the existing data is always encrypted
	//  at rest.
	// +kcc:proto:field=google.bigtable.admin.v2.EncryptionInfo.encryption_status
	EncryptionStatus *Status `json:"encryptionStatus,omitempty"`

	// Output only. The version of the Cloud KMS key specified in the parent
	//  cluster that is in use for the data underlying this table.
	// +kcc:proto:field=google.bigtable.admin.v2.EncryptionInfo.kms_key_version
	KMSKeyVersion *string `json:"kmsKeyVersion,omitempty"`
}
