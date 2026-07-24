// Copyright 2026 Google LLC
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

package v1alpha1

import (
	common "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var EdgeContainerNodePoolGVK = GroupVersion.WithKind("EdgeContainerNodePool")

// EdgeContainerNodePoolSpec defines the desired state of EdgeContainerNodePool
// +kcc:spec:proto=google.cloud.edgecontainer.v1.NodePool
type EdgeContainerNodePoolSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The EdgeContainerNodePool name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// The cluster that this node pool belongs to.
	ClusterRef *ClusterRef `json:"clusterRef,omitempty"`

	// Name of the Google Distributed Cloud Edge zone where this node pool will be
	//  created. For example: `us-central1-edge-customer-a`.
	NodeLocation *string `json:"nodeLocation,omitempty"`

	// Required. The number of nodes in the pool.
	NodeCount *int32 `json:"nodeCount,omitempty"`

	// Only machines matching this filter will be allowed to join the node pool.
	//  The filtering language accepts strings like "name=<name>", and is
	//  documented in more detail in [AIP-160](https://google.aip.dev/160).
	MachineFilter *string `json:"machineFilter,omitempty"`

	// Optional. Local disk encryption options. This field is only used when
	//  enabling CMEK support.
	LocalDiskEncryption *NodePool_LocalDiskEncryption `json:"localDiskEncryption,omitempty"`

	// Optional. Configuration for each node in the NodePool
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

// EdgeContainerNodePoolStatus defines the config connector machine state of EdgeContainerNodePool
type EdgeContainerNodePoolStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the EdgeContainerNodePool resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *NodePoolObservedState `json:"observedState,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.edgecontainer.v1.NodePool
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

// +kcc:observedstate:proto=google.cloud.edgecontainer.v1.NodePool.LocalDiskEncryption
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
	KMSStatus *common.Status `json:"kmsStatus,omitempty"`

	// Output only. The current resource state associated with the cmek.
	// +kcc:proto:field=google.cloud.edgecontainer.v1.NodePool.LocalDiskEncryption.resource_state
	ResourceState *string `json:"resourceState,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpedgecontainernodepool;gcpedgecontainernodepools
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// EdgeContainerNodePool is the Schema for the EdgeContainerNodePool API
// +k8s:openapi-gen=true
type EdgeContainerNodePool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   EdgeContainerNodePoolSpec   `json:"spec,omitempty"`
	Status EdgeContainerNodePoolStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// EdgeContainerNodePoolList contains a list of EdgeContainerNodePool
type EdgeContainerNodePoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EdgeContainerNodePool `json:"items"`
}

func init() {
	SchemeBuilder.Register(&EdgeContainerNodePool{}, &EdgeContainerNodePoolList{})
}
