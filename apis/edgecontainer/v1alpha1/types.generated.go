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


// +kcc:proto=google.cloud.edgecontainer.v1.NodePool
type NodePool struct {
	// Required. The resource name of the node pool.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.NodePool.name
	Name *string `json:"name,omitempty"`

	// Labels associated with this resource.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.NodePool.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Name of the Google Distributed Cloud Edge zone where this node pool will be
	//  created. For example: `us-central1-edge-customer-a`.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.NodePool.node_location
	NodeLocation *string `json:"nodeLocation,omitempty"`

	// Required. The number of nodes in the pool.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.NodePool.node_count
	NodeCount *int32 `json:"nodeCount,omitempty"`

	// Only machines matching this filter will be allowed to join the node pool.
	//  The filtering language accepts strings like "name=<name>", and is
	//  documented in more detail in [AIP-160](https://google.aip.dev/160).
	// +kcc:proto:field=google.cloud.edgecontainer.v1.NodePool.machine_filter
	MachineFilter *string `json:"machineFilter,omitempty"`

	// Optional. Local disk encryption options. This field is only used when
	//  enabling CMEK support.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.NodePool.local_disk_encryption
	LocalDiskEncryption *NodePool_LocalDiskEncryption `json:"localDiskEncryption,omitempty"`

	// Optional. Configuration for each node in the NodePool
	// +kcc:proto:field=google.cloud.edgecontainer.v1.NodePool.node_config
	NodeConfig *NodePool_NodeConfig `json:"nodeConfig,omitempty"`
}

// +kcc:proto=google.cloud.edgecontainer.v1.NodePool.LocalDiskEncryption
type NodePool_LocalDiskEncryption struct {
	// Optional. The Cloud KMS CryptoKey e.g.
	//  projects/{project}/locations/{location}/keyRings/{keyRing}/cryptoKeys/{cryptoKey}
	//  to use for protecting node local disks. If not specified, a
	//  Google-managed key will be used instead.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.NodePool.LocalDiskEncryption.kms_key
	KMSKey *string `json:"kmsKey,omitempty"`
}

// +kcc:proto=google.cloud.edgecontainer.v1.NodePool.NodeConfig
type NodePool_NodeConfig struct {
	// Optional. The Kubernetes node labels
	// +kcc:proto:field=google.cloud.edgecontainer.v1.NodePool.NodeConfig.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Name for the storage schema of worker nodes.
	//
	//  Warning: Configurable node local storage schema feature is an
	//  experimental feature, and is not recommended for general use
	//  in production clusters/nodepools.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.NodePool.NodeConfig.node_storage_schema
	NodeStorageSchema *string `json:"nodeStorageSchema,omitempty"`
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

// +kcc:proto=google.cloud.edgecontainer.v1.NodePool
type NodePoolObservedState struct {
	// Output only. The time when the node pool was created.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.NodePool.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the node pool was last updated.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.NodePool.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Optional. Local disk encryption options. This field is only used when
	//  enabling CMEK support.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.NodePool.local_disk_encryption
	LocalDiskEncryption *NodePool_LocalDiskEncryptionObservedState `json:"localDiskEncryption,omitempty"`

	// Output only. The lowest release version among all worker nodes.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.NodePool.node_version
	NodeVersion *string `json:"nodeVersion,omitempty"`
}

// +kcc:proto=google.cloud.edgecontainer.v1.NodePool.LocalDiskEncryption
type NodePool_LocalDiskEncryptionObservedState struct {
	// Output only. The Cloud KMS CryptoKeyVersion currently in use for
	//  protecting node local disks. Only applicable if kms_key is set.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.NodePool.LocalDiskEncryption.kms_key_active_version
	KMSKeyActiveVersion *string `json:"kmsKeyActiveVersion,omitempty"`

	// Output only. Availability of the Cloud KMS CryptoKey. If not
	//  `KEY_AVAILABLE`, then nodes may go offline as they cannot access their
	//  local data. This can be caused by a lack of permissions to use the key,
	//  or if the key is disabled or deleted.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.NodePool.LocalDiskEncryption.kms_key_state
	KMSKeyState *string `json:"kmsKeyState,omitempty"`

	// Output only. Error status returned by Cloud KMS when using this key. This
	//  field may be populated only if `kms_key_state` is not
	//  `KMS_KEY_STATE_KEY_AVAILABLE`. If populated, this field contains the
	//  error status reported by Cloud KMS.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.NodePool.LocalDiskEncryption.kms_status
	KMSStatus *Status `json:"kmsStatus,omitempty"`

	// Output only. The current resource state associated with the cmek.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.NodePool.LocalDiskEncryption.resource_state
	ResourceState *string `json:"resourceState,omitempty"`
}
