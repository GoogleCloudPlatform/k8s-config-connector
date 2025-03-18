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

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)

// +kcc:proto=google.bigtable.admin.v2.AutoscalingLimits
type AutoscalingLimits struct {
	// Required. Minimum number of nodes to scale down to.
	// +kcc:proto:field=google.bigtable.admin.v2.AutoscalingLimits.min_serve_nodes
	MinServeNodes *int32 `json:"minServeNodes,omitempty"`

	// Required. Maximum number of nodes to scale up to.
	// +kcc:proto:field=google.bigtable.admin.v2.AutoscalingLimits.max_serve_nodes
	MaxServeNodes *int32 `json:"maxServeNodes,omitempty"`
}

// +kcc:proto=google.bigtable.admin.v2.AutoscalingTargets
type AutoscalingTargets struct {
	// The cpu utilization that the Autoscaler should be trying to achieve.
	//  This number is on a scale from 0 (no utilization) to
	//  100 (total utilization), and is limited between 10 and 80, otherwise it
	//  will return INVALID_ARGUMENT error.
	// +kcc:proto:field=google.bigtable.admin.v2.AutoscalingTargets.cpu_utilization_percent
	CPUUtilizationPercent *int32 `json:"cpuUtilizationPercent,omitempty"`

	// The storage utilization that the Autoscaler should be trying to achieve.
	//  This number is limited between 2560 (2.5TiB) and 5120 (5TiB) for a SSD
	//  cluster and between 8192 (8TiB) and 16384 (16TiB) for an HDD cluster,
	//  otherwise it will return INVALID_ARGUMENT error. If this value is set to 0,
	//  it will be treated as if it were set to the default value: 2560 for SSD,
	//  8192 for HDD.
	// +kcc:proto:field=google.bigtable.admin.v2.AutoscalingTargets.storage_utilization_gib_per_node
	StorageUtilizationGiBPerNode *int32 `json:"storageUtilizationGiBPerNode,omitempty"`
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

// +kcc:proto=google.bigtable.admin.v2.Cluster.EncryptionConfig
type Cluster_EncryptionConfig struct {
	// Describes the Cloud KMS encryption key that will be used to protect the
	//  destination Bigtable cluster. The requirements for this key are:
	//   1) The Cloud Bigtable service account associated with the project that
	//   contains this cluster must be granted the
	//   `cloudkms.cryptoKeyEncrypterDecrypter` role on the CMEK key.
	//   2) Only regional keys can be used and the region of the CMEK key must
	//   match the region of the cluster.
	//   3) All clusters within an instance must use the same CMEK key.
	//  Values are of the form
	//  `projects/{project}/locations/{location}/keyRings/{keyring}/cryptoKeys/{key}`
	// +kcc:proto:field=google.bigtable.admin.v2.Cluster.EncryptionConfig.kms_key_name
	KMSKeyRef *refs.KMSCryptoKeyRef `json:"kmsKeyRef,omitempty"`
}
