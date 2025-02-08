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

package netapp

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/netapp/apiv1/netapppb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/netapp/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func DestinationVolumeParameters_FromProto(mapCtx *direct.MapContext, in *pb.DestinationVolumeParameters) *krm.DestinationVolumeParameters {
	if in == nil {
		return nil
	}
	out := &krm.DestinationVolumeParameters{}
	out.StoragePool = direct.LazyPtr(in.GetStoragePool())
	out.VolumeID = direct.LazyPtr(in.GetVolumeId())
	out.ShareName = direct.LazyPtr(in.GetShareName())
	out.Description = in.Description
	out.TieringPolicy = TieringPolicy_FromProto(mapCtx, in.GetTieringPolicy())
	return out
}
func DestinationVolumeParameters_ToProto(mapCtx *direct.MapContext, in *krm.DestinationVolumeParameters) *pb.DestinationVolumeParameters {
	if in == nil {
		return nil
	}
	out := &pb.DestinationVolumeParameters{}
	out.StoragePool = direct.ValueOf(in.StoragePool)
	out.VolumeId = direct.ValueOf(in.VolumeID)
	out.ShareName = direct.ValueOf(in.ShareName)
	out.Description = in.Description
	if oneof := TieringPolicy_ToProto(mapCtx, in.TieringPolicy); oneof != nil {
		out.TieringPolicy = &pb.DestinationVolumeParameters_TieringPolicy{TieringPolicy: oneof}
	}
	return out
}
func HybridPeeringDetails_FromProto(mapCtx *direct.MapContext, in *pb.HybridPeeringDetails) *krm.HybridPeeringDetails {
	if in == nil {
		return nil
	}
	out := &krm.HybridPeeringDetails{}
	out.SubnetIP = direct.LazyPtr(in.GetSubnetIp())
	out.Command = direct.LazyPtr(in.GetCommand())
	out.CommandExpiryTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCommandExpiryTime())
	out.Passphrase = direct.LazyPtr(in.GetPassphrase())
	return out
}
func HybridPeeringDetails_ToProto(mapCtx *direct.MapContext, in *krm.HybridPeeringDetails) *pb.HybridPeeringDetails {
	if in == nil {
		return nil
	}
	out := &pb.HybridPeeringDetails{}
	out.SubnetIp = direct.ValueOf(in.SubnetIP)
	out.Command = direct.ValueOf(in.Command)
	out.CommandExpiryTime = direct.StringTimestamp_ToProto(mapCtx, in.CommandExpiryTime)
	out.Passphrase = direct.ValueOf(in.Passphrase)
	return out
}
func NetappReplicationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Replication) *krm.NetappReplicationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetappReplicationObservedState{}
	// MISSING: Name
	// MISSING: State
	// MISSING: StateDetails
	// MISSING: Role
	// MISSING: ReplicationSchedule
	// MISSING: MirrorState
	// MISSING: Healthy
	// MISSING: CreateTime
	// MISSING: DestinationVolume
	// MISSING: TransferStats
	// MISSING: Labels
	// MISSING: Description
	// MISSING: DestinationVolumeParameters
	// MISSING: SourceVolume
	// MISSING: HybridPeeringDetails
	// MISSING: ClusterLocation
	// MISSING: HybridReplicationType
	return out
}
func NetappReplicationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetappReplicationObservedState) *pb.Replication {
	if in == nil {
		return nil
	}
	out := &pb.Replication{}
	// MISSING: Name
	// MISSING: State
	// MISSING: StateDetails
	// MISSING: Role
	// MISSING: ReplicationSchedule
	// MISSING: MirrorState
	// MISSING: Healthy
	// MISSING: CreateTime
	// MISSING: DestinationVolume
	// MISSING: TransferStats
	// MISSING: Labels
	// MISSING: Description
	// MISSING: DestinationVolumeParameters
	// MISSING: SourceVolume
	// MISSING: HybridPeeringDetails
	// MISSING: ClusterLocation
	// MISSING: HybridReplicationType
	return out
}
func NetappReplicationSpec_FromProto(mapCtx *direct.MapContext, in *pb.Replication) *krm.NetappReplicationSpec {
	if in == nil {
		return nil
	}
	out := &krm.NetappReplicationSpec{}
	// MISSING: Name
	// MISSING: State
	// MISSING: StateDetails
	// MISSING: Role
	// MISSING: ReplicationSchedule
	// MISSING: MirrorState
	// MISSING: Healthy
	// MISSING: CreateTime
	// MISSING: DestinationVolume
	// MISSING: TransferStats
	// MISSING: Labels
	// MISSING: Description
	// MISSING: DestinationVolumeParameters
	// MISSING: SourceVolume
	// MISSING: HybridPeeringDetails
	// MISSING: ClusterLocation
	// MISSING: HybridReplicationType
	return out
}
func NetappReplicationSpec_ToProto(mapCtx *direct.MapContext, in *krm.NetappReplicationSpec) *pb.Replication {
	if in == nil {
		return nil
	}
	out := &pb.Replication{}
	// MISSING: Name
	// MISSING: State
	// MISSING: StateDetails
	// MISSING: Role
	// MISSING: ReplicationSchedule
	// MISSING: MirrorState
	// MISSING: Healthy
	// MISSING: CreateTime
	// MISSING: DestinationVolume
	// MISSING: TransferStats
	// MISSING: Labels
	// MISSING: Description
	// MISSING: DestinationVolumeParameters
	// MISSING: SourceVolume
	// MISSING: HybridPeeringDetails
	// MISSING: ClusterLocation
	// MISSING: HybridReplicationType
	return out
}
func Replication_FromProto(mapCtx *direct.MapContext, in *pb.Replication) *krm.Replication {
	if in == nil {
		return nil
	}
	out := &krm.Replication{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: State
	// MISSING: StateDetails
	// MISSING: Role
	out.ReplicationSchedule = direct.Enum_FromProto(mapCtx, in.GetReplicationSchedule())
	// MISSING: MirrorState
	// MISSING: Healthy
	// MISSING: CreateTime
	// MISSING: DestinationVolume
	// MISSING: TransferStats
	out.Labels = in.Labels
	out.Description = in.Description
	out.DestinationVolumeParameters = DestinationVolumeParameters_FromProto(mapCtx, in.GetDestinationVolumeParameters())
	// MISSING: SourceVolume
	// MISSING: HybridPeeringDetails
	out.ClusterLocation = direct.LazyPtr(in.GetClusterLocation())
	// MISSING: HybridReplicationType
	return out
}
func Replication_ToProto(mapCtx *direct.MapContext, in *krm.Replication) *pb.Replication {
	if in == nil {
		return nil
	}
	out := &pb.Replication{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: State
	// MISSING: StateDetails
	// MISSING: Role
	out.ReplicationSchedule = direct.Enum_ToProto[pb.Replication_ReplicationSchedule](mapCtx, in.ReplicationSchedule)
	// MISSING: MirrorState
	// MISSING: Healthy
	// MISSING: CreateTime
	// MISSING: DestinationVolume
	// MISSING: TransferStats
	out.Labels = in.Labels
	out.Description = in.Description
	out.DestinationVolumeParameters = DestinationVolumeParameters_ToProto(mapCtx, in.DestinationVolumeParameters)
	// MISSING: SourceVolume
	// MISSING: HybridPeeringDetails
	out.ClusterLocation = direct.ValueOf(in.ClusterLocation)
	// MISSING: HybridReplicationType
	return out
}
func ReplicationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Replication) *krm.ReplicationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ReplicationObservedState{}
	// MISSING: Name
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.StateDetails = direct.LazyPtr(in.GetStateDetails())
	out.Role = direct.Enum_FromProto(mapCtx, in.GetRole())
	// MISSING: ReplicationSchedule
	out.MirrorState = direct.Enum_FromProto(mapCtx, in.GetMirrorState())
	out.Healthy = in.Healthy
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.DestinationVolume = direct.LazyPtr(in.GetDestinationVolume())
	out.TransferStats = TransferStats_FromProto(mapCtx, in.GetTransferStats())
	// MISSING: Labels
	// MISSING: Description
	// MISSING: DestinationVolumeParameters
	out.SourceVolume = direct.LazyPtr(in.GetSourceVolume())
	out.HybridPeeringDetails = HybridPeeringDetails_FromProto(mapCtx, in.GetHybridPeeringDetails())
	// MISSING: ClusterLocation
	out.HybridReplicationType = direct.Enum_FromProto(mapCtx, in.GetHybridReplicationType())
	return out
}
func ReplicationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ReplicationObservedState) *pb.Replication {
	if in == nil {
		return nil
	}
	out := &pb.Replication{}
	// MISSING: Name
	out.State = direct.Enum_ToProto[pb.Replication_State](mapCtx, in.State)
	out.StateDetails = direct.ValueOf(in.StateDetails)
	out.Role = direct.Enum_ToProto[pb.Replication_ReplicationRole](mapCtx, in.Role)
	// MISSING: ReplicationSchedule
	out.MirrorState = direct.Enum_ToProto[pb.Replication_MirrorState](mapCtx, in.MirrorState)
	out.Healthy = in.Healthy
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.DestinationVolume = direct.ValueOf(in.DestinationVolume)
	out.TransferStats = TransferStats_ToProto(mapCtx, in.TransferStats)
	// MISSING: Labels
	// MISSING: Description
	// MISSING: DestinationVolumeParameters
	out.SourceVolume = direct.ValueOf(in.SourceVolume)
	out.HybridPeeringDetails = HybridPeeringDetails_ToProto(mapCtx, in.HybridPeeringDetails)
	// MISSING: ClusterLocation
	out.HybridReplicationType = direct.Enum_ToProto[pb.Replication_HybridReplicationType](mapCtx, in.HybridReplicationType)
	return out
}
func TieringPolicy_FromProto(mapCtx *direct.MapContext, in *pb.TieringPolicy) *krm.TieringPolicy {
	if in == nil {
		return nil
	}
	out := &krm.TieringPolicy{}
	out.TierAction = direct.Enum_FromProto(mapCtx, in.GetTierAction())
	out.CoolingThresholdDays = in.CoolingThresholdDays
	return out
}
func TieringPolicy_ToProto(mapCtx *direct.MapContext, in *krm.TieringPolicy) *pb.TieringPolicy {
	if in == nil {
		return nil
	}
	out := &pb.TieringPolicy{}
	if oneof := TieringPolicy_TierAction_ToProto(mapCtx, in.TierAction); oneof != nil {
		out.TierAction = oneof
	}
	out.CoolingThresholdDays = in.CoolingThresholdDays
	return out
}
func TransferStats_FromProto(mapCtx *direct.MapContext, in *pb.TransferStats) *krm.TransferStats {
	if in == nil {
		return nil
	}
	out := &krm.TransferStats{}
	out.TransferBytes = in.TransferBytes
	out.TotalTransferDuration = direct.StringDuration_FromProto(mapCtx, in.GetTotalTransferDuration())
	out.LastTransferBytes = in.LastTransferBytes
	out.LastTransferDuration = direct.StringDuration_FromProto(mapCtx, in.GetLastTransferDuration())
	out.LagDuration = direct.StringDuration_FromProto(mapCtx, in.GetLagDuration())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.LastTransferEndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLastTransferEndTime())
	out.LastTransferError = in.LastTransferError
	return out
}
func TransferStats_ToProto(mapCtx *direct.MapContext, in *krm.TransferStats) *pb.TransferStats {
	if in == nil {
		return nil
	}
	out := &pb.TransferStats{}
	out.TransferBytes = in.TransferBytes
	if oneof := direct.StringDuration_ToProto(mapCtx, in.TotalTransferDuration); oneof != nil {
		out.TotalTransferDuration = &pb.TransferStats_TotalTransferDuration{TotalTransferDuration: oneof}
	}
	out.LastTransferBytes = in.LastTransferBytes
	if oneof := direct.StringDuration_ToProto(mapCtx, in.LastTransferDuration); oneof != nil {
		out.LastTransferDuration = &pb.TransferStats_LastTransferDuration{LastTransferDuration: oneof}
	}
	if oneof := direct.StringDuration_ToProto(mapCtx, in.LagDuration); oneof != nil {
		out.LagDuration = &pb.TransferStats_LagDuration{LagDuration: oneof}
	}
	if oneof := direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime); oneof != nil {
		out.UpdateTime = &pb.TransferStats_UpdateTime{UpdateTime: oneof}
	}
	if oneof := direct.StringTimestamp_ToProto(mapCtx, in.LastTransferEndTime); oneof != nil {
		out.LastTransferEndTime = &pb.TransferStats_LastTransferEndTime{LastTransferEndTime: oneof}
	}
	out.LastTransferError = in.LastTransferError
	return out
}
