// Copyright 2026 Google LLC
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

package sql

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/sql/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gcpclients/generated/google/cloud/sql/v1beta4"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func Int32Value_FromProto(mapCtx *direct.MapContext, in *wrapperspb.Int32Value) *krm.Int32Value {
	if in == nil {
		return nil
	}
	return &krm.Int32Value{
		Value: &in.Value,
	}
}

func Int32Value_v1beta1_FromProto(mapCtx *direct.MapContext, in *wrapperspb.Int32Value) *krm.Int32Value {
	return Int32Value_FromProto(mapCtx, in)
}

func Int32Value_ToProto(mapCtx *direct.MapContext, in *krm.Int32Value) *wrapperspb.Int32Value {
	if in == nil || in.Value == nil {
		return nil
	}
	return &wrapperspb.Int32Value{
		Value: *in.Value,
	}
}

func Int32Value_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.Int32Value) *wrapperspb.Int32Value {
	return Int32Value_ToProto(mapCtx, in)
}

func ReplicationCluster_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.ReplicationCluster) *krm.ReplicationCluster {
	return ReplicationCluster_FromProto(mapCtx, in)
}

func ReplicationCluster_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ReplicationCluster) *pb.ReplicationCluster {
	return ReplicationCluster_ToProto(mapCtx, in)
}

func ReplicationClusterObservedState_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.ReplicationCluster) *krm.ReplicationClusterObservedState {
	return ReplicationClusterObservedState_FromProto(mapCtx, in)
}

func ReplicationClusterObservedState_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ReplicationClusterObservedState) *pb.ReplicationCluster {
	return ReplicationClusterObservedState_ToProto(mapCtx, in)
}

var (
	DatabaseInstance_FromProto              = DatabaseInstance_v1beta1_FromProto
	DatabaseInstance_ToProto                = DatabaseInstance_v1beta1_ToProto
	DatabaseInstanceObservedState_FromProto = DatabaseInstanceObservedState_v1beta1_FromProto
	DatabaseInstanceObservedState_ToProto   = DatabaseInstanceObservedState_v1beta1_ToProto
)

func BackupConfigurationObservedState_TransactionalLogStorageState_ToProto(mapCtx *direct.MapContext, in *string) *pb.BackupConfiguration_TransactionalLogStorageState {
	if in == nil {
		return nil
	}
	mapCtx.NotImplemented()
	return nil
}

func DatabaseInstance_SqlNetworkArchitecture_ToProto(mapCtx *direct.MapContext, in *string) *pb.DatabaseInstance_SqlNetworkArchitecture {
	if in == nil {
		return nil
	}
	mapCtx.NotImplemented()
	return nil
}

func DatabaseInstance_SQLOutOfDiskReport_SqlOutOfDiskState_ToProto(mapCtx *direct.MapContext, in *string) *pb.DatabaseInstance_SqlOutOfDiskReport_SqlOutOfDiskState {
	if in == nil {
		return nil
	}
	out := direct.Enum_ToProto[pb.DatabaseInstance_SqlOutOfDiskReport_SqlOutOfDiskState](mapCtx, in)
	return &out
}
