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


// +kcc:proto=google.cloud.redis.cluster.v1beta1.BackupCollection
type BackupCollection struct {
	// Identifier. Full resource path of the backup collection.
	// +kcc:proto:field=google.cloud.redis.cluster.v1beta1.BackupCollection.name
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.cloud.redis.cluster.v1beta1.BackupCollection
type BackupCollectionObservedState struct {
	// Output only. The cluster uid of the backup collection.
	// +kcc:proto:field=google.cloud.redis.cluster.v1beta1.BackupCollection.cluster_uid
	ClusterUid *string `json:"clusterUid,omitempty"`

	// Output only. The full resource path of the cluster the backup collection
	//  belongs to. Example:
	//  projects/{project}/locations/{location}/clusters/{cluster}
	// +kcc:proto:field=google.cloud.redis.cluster.v1beta1.BackupCollection.cluster
	Cluster *string `json:"cluster,omitempty"`

	// Output only. The KMS key used to encrypt the backups under this backup
	//  collection.
	// +kcc:proto:field=google.cloud.redis.cluster.v1beta1.BackupCollection.kms_key
	KMSKey *string `json:"kmsKey,omitempty"`

	// Output only. System assigned unique identifier of the backup collection.
	// +kcc:proto:field=google.cloud.redis.cluster.v1beta1.BackupCollection.uid
	Uid *string `json:"uid,omitempty"`
}
