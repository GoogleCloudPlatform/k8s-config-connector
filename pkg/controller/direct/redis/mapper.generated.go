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

package redis

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/redis/cluster/apiv1beta1/clusterpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/redis/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func BackupCollection_FromProto(mapCtx *direct.MapContext, in *pb.BackupCollection) *krm.BackupCollection {
	if in == nil {
		return nil
	}
	out := &krm.BackupCollection{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: ClusterUid
	// MISSING: Cluster
	// MISSING: KMSKey
	// MISSING: Uid
	return out
}
func BackupCollection_ToProto(mapCtx *direct.MapContext, in *krm.BackupCollection) *pb.BackupCollection {
	if in == nil {
		return nil
	}
	out := &pb.BackupCollection{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: ClusterUid
	// MISSING: Cluster
	// MISSING: KMSKey
	// MISSING: Uid
	return out
}
func BackupCollectionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BackupCollection) *krm.BackupCollectionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BackupCollectionObservedState{}
	// MISSING: Name
	out.ClusterUid = direct.LazyPtr(in.GetClusterUid())
	out.Cluster = direct.LazyPtr(in.GetCluster())
	out.KMSKey = direct.LazyPtr(in.GetKmsKey())
	out.Uid = direct.LazyPtr(in.GetUid())
	return out
}
func BackupCollectionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BackupCollectionObservedState) *pb.BackupCollection {
	if in == nil {
		return nil
	}
	out := &pb.BackupCollection{}
	// MISSING: Name
	out.ClusterUid = direct.ValueOf(in.ClusterUid)
	out.Cluster = direct.ValueOf(in.Cluster)
	out.KmsKey = direct.ValueOf(in.KMSKey)
	out.Uid = direct.ValueOf(in.Uid)
	return out
}
func RedisBackupCollectionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BackupCollection) *krm.RedisBackupCollectionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.RedisBackupCollectionObservedState{}
	// MISSING: Name
	// MISSING: ClusterUid
	// MISSING: Cluster
	// MISSING: KMSKey
	// MISSING: Uid
	return out
}
func RedisBackupCollectionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.RedisBackupCollectionObservedState) *pb.BackupCollection {
	if in == nil {
		return nil
	}
	out := &pb.BackupCollection{}
	// MISSING: Name
	// MISSING: ClusterUid
	// MISSING: Cluster
	// MISSING: KMSKey
	// MISSING: Uid
	return out
}
func RedisBackupCollectionSpec_FromProto(mapCtx *direct.MapContext, in *pb.BackupCollection) *krm.RedisBackupCollectionSpec {
	if in == nil {
		return nil
	}
	out := &krm.RedisBackupCollectionSpec{}
	// MISSING: Name
	// MISSING: ClusterUid
	// MISSING: Cluster
	// MISSING: KMSKey
	// MISSING: Uid
	return out
}
func RedisBackupCollectionSpec_ToProto(mapCtx *direct.MapContext, in *krm.RedisBackupCollectionSpec) *pb.BackupCollection {
	if in == nil {
		return nil
	}
	out := &pb.BackupCollection{}
	// MISSING: Name
	// MISSING: ClusterUid
	// MISSING: Cluster
	// MISSING: KMSKey
	// MISSING: Uid
	return out
}
