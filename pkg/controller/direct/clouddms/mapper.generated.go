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

package clouddms

import (
	pb "cloud.google.com/go/clouddms/apiv1/clouddmspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/clouddms/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func ClouddmsMigrationJobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MigrationJob) *krm.ClouddmsMigrationJobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ClouddmsMigrationJobObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: State
	// MISSING: Phase
	// MISSING: Type
	// MISSING: DumpPath
	// MISSING: DumpFlags
	// MISSING: Source
	// MISSING: Destination
	// MISSING: ReverseSSHConnectivity
	// MISSING: VpcPeeringConnectivity
	// MISSING: StaticIPConnectivity
	// MISSING: Duration
	// MISSING: Error
	// MISSING: SourceDatabase
	// MISSING: DestinationDatabase
	// MISSING: EndTime
	// MISSING: ConversionWorkspace
	// MISSING: Filter
	// MISSING: CmekKeyName
	// MISSING: PerformanceConfig
	return out
}
func ClouddmsMigrationJobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ClouddmsMigrationJobObservedState) *pb.MigrationJob {
	if in == nil {
		return nil
	}
	out := &pb.MigrationJob{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: State
	// MISSING: Phase
	// MISSING: Type
	// MISSING: DumpPath
	// MISSING: DumpFlags
	// MISSING: Source
	// MISSING: Destination
	// MISSING: ReverseSSHConnectivity
	// MISSING: VpcPeeringConnectivity
	// MISSING: StaticIPConnectivity
	// MISSING: Duration
	// MISSING: Error
	// MISSING: SourceDatabase
	// MISSING: DestinationDatabase
	// MISSING: EndTime
	// MISSING: ConversionWorkspace
	// MISSING: Filter
	// MISSING: CmekKeyName
	// MISSING: PerformanceConfig
	return out
}
func ClouddmsMigrationJobSpec_FromProto(mapCtx *direct.MapContext, in *pb.MigrationJob) *krm.ClouddmsMigrationJobSpec {
	if in == nil {
		return nil
	}
	out := &krm.ClouddmsMigrationJobSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: State
	// MISSING: Phase
	// MISSING: Type
	// MISSING: DumpPath
	// MISSING: DumpFlags
	// MISSING: Source
	// MISSING: Destination
	// MISSING: ReverseSSHConnectivity
	// MISSING: VpcPeeringConnectivity
	// MISSING: StaticIPConnectivity
	// MISSING: Duration
	// MISSING: Error
	// MISSING: SourceDatabase
	// MISSING: DestinationDatabase
	// MISSING: EndTime
	// MISSING: ConversionWorkspace
	// MISSING: Filter
	// MISSING: CmekKeyName
	// MISSING: PerformanceConfig
	return out
}
func ClouddmsMigrationJobSpec_ToProto(mapCtx *direct.MapContext, in *krm.ClouddmsMigrationJobSpec) *pb.MigrationJob {
	if in == nil {
		return nil
	}
	out := &pb.MigrationJob{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: State
	// MISSING: Phase
	// MISSING: Type
	// MISSING: DumpPath
	// MISSING: DumpFlags
	// MISSING: Source
	// MISSING: Destination
	// MISSING: ReverseSSHConnectivity
	// MISSING: VpcPeeringConnectivity
	// MISSING: StaticIPConnectivity
	// MISSING: Duration
	// MISSING: Error
	// MISSING: SourceDatabase
	// MISSING: DestinationDatabase
	// MISSING: EndTime
	// MISSING: ConversionWorkspace
	// MISSING: Filter
	// MISSING: CmekKeyName
	// MISSING: PerformanceConfig
	return out
}
func ConversionWorkspaceInfo_FromProto(mapCtx *direct.MapContext, in *pb.ConversionWorkspaceInfo) *krm.ConversionWorkspaceInfo {
	if in == nil {
		return nil
	}
	out := &krm.ConversionWorkspaceInfo{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CommitID = direct.LazyPtr(in.GetCommitId())
	return out
}
func ConversionWorkspaceInfo_ToProto(mapCtx *direct.MapContext, in *krm.ConversionWorkspaceInfo) *pb.ConversionWorkspaceInfo {
	if in == nil {
		return nil
	}
	out := &pb.ConversionWorkspaceInfo{}
	out.Name = direct.ValueOf(in.Name)
	out.CommitId = direct.ValueOf(in.CommitID)
	return out
}
func DatabaseType_FromProto(mapCtx *direct.MapContext, in *pb.DatabaseType) *krm.DatabaseType {
	if in == nil {
		return nil
	}
	out := &krm.DatabaseType{}
	out.Provider = direct.Enum_FromProto(mapCtx, in.GetProvider())
	out.Engine = direct.Enum_FromProto(mapCtx, in.GetEngine())
	return out
}
func DatabaseType_ToProto(mapCtx *direct.MapContext, in *krm.DatabaseType) *pb.DatabaseType {
	if in == nil {
		return nil
	}
	out := &pb.DatabaseType{}
	out.Provider = direct.Enum_ToProto[pb.DatabaseProvider](mapCtx, in.Provider)
	out.Engine = direct.Enum_ToProto[pb.DatabaseEngine](mapCtx, in.Engine)
	return out
}
func MigrationJob_FromProto(mapCtx *direct.MapContext, in *pb.MigrationJob) *krm.MigrationJob {
	if in == nil {
		return nil
	}
	out := &krm.MigrationJob{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: Phase
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.DumpPath = direct.LazyPtr(in.GetDumpPath())
	out.DumpFlags = MigrationJob_DumpFlags_FromProto(mapCtx, in.GetDumpFlags())
	out.Source = direct.LazyPtr(in.GetSource())
	out.Destination = direct.LazyPtr(in.GetDestination())
	out.ReverseSSHConnectivity = ReverseSshConnectivity_FromProto(mapCtx, in.GetReverseSshConnectivity())
	out.VpcPeeringConnectivity = VpcPeeringConnectivity_FromProto(mapCtx, in.GetVpcPeeringConnectivity())
	out.StaticIPConnectivity = StaticIpConnectivity_FromProto(mapCtx, in.GetStaticIpConnectivity())
	// MISSING: Duration
	// MISSING: Error
	out.SourceDatabase = DatabaseType_FromProto(mapCtx, in.GetSourceDatabase())
	out.DestinationDatabase = DatabaseType_FromProto(mapCtx, in.GetDestinationDatabase())
	// MISSING: EndTime
	out.ConversionWorkspace = ConversionWorkspaceInfo_FromProto(mapCtx, in.GetConversionWorkspace())
	out.Filter = direct.LazyPtr(in.GetFilter())
	out.CmekKeyName = direct.LazyPtr(in.GetCmekKeyName())
	out.PerformanceConfig = MigrationJob_PerformanceConfig_FromProto(mapCtx, in.GetPerformanceConfig())
	return out
}
func MigrationJob_ToProto(mapCtx *direct.MapContext, in *krm.MigrationJob) *pb.MigrationJob {
	if in == nil {
		return nil
	}
	out := &pb.MigrationJob{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.State = direct.Enum_ToProto[pb.MigrationJob_State](mapCtx, in.State)
	// MISSING: Phase
	out.Type = direct.Enum_ToProto[pb.MigrationJob_Type](mapCtx, in.Type)
	out.DumpPath = direct.ValueOf(in.DumpPath)
	out.DumpFlags = MigrationJob_DumpFlags_ToProto(mapCtx, in.DumpFlags)
	out.Source = direct.ValueOf(in.Source)
	out.Destination = direct.ValueOf(in.Destination)
	if oneof := ReverseSshConnectivity_ToProto(mapCtx, in.ReverseSSHConnectivity); oneof != nil {
		out.Connectivity = &pb.MigrationJob_ReverseSshConnectivity{ReverseSshConnectivity: oneof}
	}
	if oneof := VpcPeeringConnectivity_ToProto(mapCtx, in.VpcPeeringConnectivity); oneof != nil {
		out.Connectivity = &pb.MigrationJob_VpcPeeringConnectivity{VpcPeeringConnectivity: oneof}
	}
	if oneof := StaticIpConnectivity_ToProto(mapCtx, in.StaticIPConnectivity); oneof != nil {
		out.Connectivity = &pb.MigrationJob_StaticIpConnectivity{StaticIpConnectivity: oneof}
	}
	// MISSING: Duration
	// MISSING: Error
	out.SourceDatabase = DatabaseType_ToProto(mapCtx, in.SourceDatabase)
	out.DestinationDatabase = DatabaseType_ToProto(mapCtx, in.DestinationDatabase)
	// MISSING: EndTime
	out.ConversionWorkspace = ConversionWorkspaceInfo_ToProto(mapCtx, in.ConversionWorkspace)
	out.Filter = direct.ValueOf(in.Filter)
	out.CmekKeyName = direct.ValueOf(in.CmekKeyName)
	out.PerformanceConfig = MigrationJob_PerformanceConfig_ToProto(mapCtx, in.PerformanceConfig)
	return out
}
func MigrationJobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MigrationJob) *krm.MigrationJobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MigrationJobObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: State
	out.Phase = direct.Enum_FromProto(mapCtx, in.GetPhase())
	// MISSING: Type
	// MISSING: DumpPath
	// MISSING: DumpFlags
	// MISSING: Source
	// MISSING: Destination
	// MISSING: ReverseSSHConnectivity
	// MISSING: VpcPeeringConnectivity
	// MISSING: StaticIPConnectivity
	out.Duration = direct.StringDuration_FromProto(mapCtx, in.GetDuration())
	out.Error = Status_FromProto(mapCtx, in.GetError())
	// MISSING: SourceDatabase
	// MISSING: DestinationDatabase
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	// MISSING: ConversionWorkspace
	// MISSING: Filter
	// MISSING: CmekKeyName
	// MISSING: PerformanceConfig
	return out
}
func MigrationJobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MigrationJobObservedState) *pb.MigrationJob {
	if in == nil {
		return nil
	}
	out := &pb.MigrationJob{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: State
	out.Phase = direct.Enum_ToProto[pb.MigrationJob_Phase](mapCtx, in.Phase)
	// MISSING: Type
	// MISSING: DumpPath
	// MISSING: DumpFlags
	// MISSING: Source
	// MISSING: Destination
	// MISSING: ReverseSSHConnectivity
	// MISSING: VpcPeeringConnectivity
	// MISSING: StaticIPConnectivity
	out.Duration = direct.StringDuration_ToProto(mapCtx, in.Duration)
	out.Error = Status_ToProto(mapCtx, in.Error)
	// MISSING: SourceDatabase
	// MISSING: DestinationDatabase
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	// MISSING: ConversionWorkspace
	// MISSING: Filter
	// MISSING: CmekKeyName
	// MISSING: PerformanceConfig
	return out
}
func MigrationJob_DumpFlag_FromProto(mapCtx *direct.MapContext, in *pb.MigrationJob_DumpFlag) *krm.MigrationJob_DumpFlag {
	if in == nil {
		return nil
	}
	out := &krm.MigrationJob_DumpFlag{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Value = direct.LazyPtr(in.GetValue())
	return out
}
func MigrationJob_DumpFlag_ToProto(mapCtx *direct.MapContext, in *krm.MigrationJob_DumpFlag) *pb.MigrationJob_DumpFlag {
	if in == nil {
		return nil
	}
	out := &pb.MigrationJob_DumpFlag{}
	out.Name = direct.ValueOf(in.Name)
	out.Value = direct.ValueOf(in.Value)
	return out
}
func MigrationJob_DumpFlags_FromProto(mapCtx *direct.MapContext, in *pb.MigrationJob_DumpFlags) *krm.MigrationJob_DumpFlags {
	if in == nil {
		return nil
	}
	out := &krm.MigrationJob_DumpFlags{}
	out.DumpFlags = direct.Slice_FromProto(mapCtx, in.DumpFlags, MigrationJob_DumpFlag_FromProto)
	return out
}
func MigrationJob_DumpFlags_ToProto(mapCtx *direct.MapContext, in *krm.MigrationJob_DumpFlags) *pb.MigrationJob_DumpFlags {
	if in == nil {
		return nil
	}
	out := &pb.MigrationJob_DumpFlags{}
	out.DumpFlags = direct.Slice_ToProto(mapCtx, in.DumpFlags, MigrationJob_DumpFlag_ToProto)
	return out
}
func MigrationJob_PerformanceConfig_FromProto(mapCtx *direct.MapContext, in *pb.MigrationJob_PerformanceConfig) *krm.MigrationJob_PerformanceConfig {
	if in == nil {
		return nil
	}
	out := &krm.MigrationJob_PerformanceConfig{}
	out.DumpParallelLevel = direct.Enum_FromProto(mapCtx, in.GetDumpParallelLevel())
	return out
}
func MigrationJob_PerformanceConfig_ToProto(mapCtx *direct.MapContext, in *krm.MigrationJob_PerformanceConfig) *pb.MigrationJob_PerformanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.MigrationJob_PerformanceConfig{}
	out.DumpParallelLevel = direct.Enum_ToProto[pb.MigrationJob_PerformanceConfig_DumpParallelLevel](mapCtx, in.DumpParallelLevel)
	return out
}
func ReverseSshConnectivity_FromProto(mapCtx *direct.MapContext, in *pb.ReverseSshConnectivity) *krm.ReverseSshConnectivity {
	if in == nil {
		return nil
	}
	out := &krm.ReverseSshConnectivity{}
	out.VmIP = direct.LazyPtr(in.GetVmIp())
	out.VmPort = direct.LazyPtr(in.GetVmPort())
	out.Vm = direct.LazyPtr(in.GetVm())
	out.Vpc = direct.LazyPtr(in.GetVpc())
	return out
}
func ReverseSshConnectivity_ToProto(mapCtx *direct.MapContext, in *krm.ReverseSshConnectivity) *pb.ReverseSshConnectivity {
	if in == nil {
		return nil
	}
	out := &pb.ReverseSshConnectivity{}
	out.VmIp = direct.ValueOf(in.VmIP)
	out.VmPort = direct.ValueOf(in.VmPort)
	out.Vm = direct.ValueOf(in.Vm)
	out.Vpc = direct.ValueOf(in.Vpc)
	return out
}
func StaticIpConnectivity_FromProto(mapCtx *direct.MapContext, in *pb.StaticIpConnectivity) *krm.StaticIpConnectivity {
	if in == nil {
		return nil
	}
	out := &krm.StaticIpConnectivity{}
	return out
}
func StaticIpConnectivity_ToProto(mapCtx *direct.MapContext, in *krm.StaticIpConnectivity) *pb.StaticIpConnectivity {
	if in == nil {
		return nil
	}
	out := &pb.StaticIpConnectivity{}
	return out
}
func VpcPeeringConnectivity_FromProto(mapCtx *direct.MapContext, in *pb.VpcPeeringConnectivity) *krm.VpcPeeringConnectivity {
	if in == nil {
		return nil
	}
	out := &krm.VpcPeeringConnectivity{}
	out.Vpc = direct.LazyPtr(in.GetVpc())
	return out
}
func VpcPeeringConnectivity_ToProto(mapCtx *direct.MapContext, in *krm.VpcPeeringConnectivity) *pb.VpcPeeringConnectivity {
	if in == nil {
		return nil
	}
	out := &pb.VpcPeeringConnectivity{}
	out.Vpc = direct.ValueOf(in.Vpc)
	return out
}
