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

// +kcc:proto=google.bigtable.admin.v2.AuthorizedView
type AuthorizedView struct {
	// Identifier. The name of this AuthorizedView.
	//  Values are of the form
	//  `projects/{project}/instances/{instance}/tables/{table}/authorizedViews/{authorized_view}`
	// +kcc:proto:field=google.bigtable.admin.v2.AuthorizedView.name
	Name *string `json:"name,omitempty"`

	// An AuthorizedView permitting access to an explicit subset of a Table.
	// +kcc:proto:field=google.bigtable.admin.v2.AuthorizedView.subset_view
	SubsetView *AuthorizedView_SubsetView `json:"subsetView,omitempty"`

	// The etag for this AuthorizedView.
	//  If this is provided on update, it must match the server's etag. The server
	//  returns ABORTED error on a mismatched etag.
	// +kcc:proto:field=google.bigtable.admin.v2.AuthorizedView.etag
	Etag *string `json:"etag,omitempty"`

	// Set to true to make the AuthorizedView protected against deletion.
	//  The parent Table and containing Instance cannot be deleted if an
	//  AuthorizedView has this bit set.
	// +kcc:proto:field=google.bigtable.admin.v2.AuthorizedView.deletion_protection
	DeletionProtection *bool `json:"deletionProtection,omitempty"`
}

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
