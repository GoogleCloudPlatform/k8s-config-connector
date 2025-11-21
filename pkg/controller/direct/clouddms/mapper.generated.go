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

// +generated:mapper
// krm.group: clouddms.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.clouddms.v1

package clouddms

import (
	pb "cloud.google.com/go/clouddms/apiv1/clouddmspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/clouddms/v1alpha1"
	krmcomputev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func CloudDMSConversionWorkspaceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ConversionWorkspace) *krm.CloudDMSConversionWorkspaceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CloudDMSConversionWorkspaceObservedState{}
	// MISSING: Name
	out.HasUncommittedChanges = direct.LazyPtr(in.GetHasUncommittedChanges())
	out.LatestCommitID = direct.LazyPtr(in.GetLatestCommitId())
	out.LatestCommitTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLatestCommitTime())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func CloudDMSConversionWorkspaceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CloudDMSConversionWorkspaceObservedState) *pb.ConversionWorkspace {
	if in == nil {
		return nil
	}
	out := &pb.ConversionWorkspace{}
	// MISSING: Name
	out.HasUncommittedChanges = direct.ValueOf(in.HasUncommittedChanges)
	out.LatestCommitId = direct.ValueOf(in.LatestCommitID)
	out.LatestCommitTime = direct.StringTimestamp_ToProto(mapCtx, in.LatestCommitTime)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func CloudDMSConversionWorkspaceSpec_FromProto(mapCtx *direct.MapContext, in *pb.ConversionWorkspace) *krm.CloudDMSConversionWorkspaceSpec {
	if in == nil {
		return nil
	}
	out := &krm.CloudDMSConversionWorkspaceSpec{}
	// MISSING: Name
	out.Source = DatabaseEngineInfo_FromProto(mapCtx, in.GetSource())
	out.Destination = DatabaseEngineInfo_FromProto(mapCtx, in.GetDestination())
	out.GlobalSettings = in.GlobalSettings
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	return out
}
func CloudDMSConversionWorkspaceSpec_ToProto(mapCtx *direct.MapContext, in *krm.CloudDMSConversionWorkspaceSpec) *pb.ConversionWorkspace {
	if in == nil {
		return nil
	}
	out := &pb.ConversionWorkspace{}
	// MISSING: Name
	out.Source = DatabaseEngineInfo_ToProto(mapCtx, in.Source)
	out.Destination = DatabaseEngineInfo_ToProto(mapCtx, in.Destination)
	out.GlobalSettings = in.GlobalSettings
	out.DisplayName = direct.ValueOf(in.DisplayName)
	return out
}
func CloudDMSMigrationJobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MigrationJob) *krm.CloudDMSMigrationJobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CloudDMSMigrationJobObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: Phase
	// MISSING: Duration
	// MISSING: Error
	// MISSING: EndTime
	return out
}
func CloudDMSMigrationJobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CloudDMSMigrationJobObservedState) *pb.MigrationJob {
	if in == nil {
		return nil
	}
	out := &pb.MigrationJob{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: Phase
	// MISSING: Duration
	// MISSING: Error
	// MISSING: EndTime
	return out
}
func CloudDMSMigrationJobSpec_FromProto(mapCtx *direct.MapContext, in *pb.MigrationJob) *krm.CloudDMSMigrationJobSpec {
	if in == nil {
		return nil
	}
	out := &krm.CloudDMSMigrationJobSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: State
	// MISSING: Phase
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.DumpPath = direct.LazyPtr(in.GetDumpPath())
	out.DumpFlags = MigrationJob_DumpFlags_FromProto(mapCtx, in.GetDumpFlags())
	if in.GetSource() != "" {
		out.SourceRef = &krm.CloudDMSConnectionProfileRef{External: in.GetSource()}
	}
	if in.GetDestination() != "" {
		out.DestinationRef = &krm.CloudDMSConnectionProfileRef{External: in.GetDestination()}
	}
	out.ReverseSSHConnectivity = ReverseSSHConnectivity_FromProto(mapCtx, in.GetReverseSshConnectivity())
	out.VPCPeeringConnectivity = VPCPeeringConnectivity_FromProto(mapCtx, in.GetVpcPeeringConnectivity())
	out.StaticIPConnectivity = StaticIPConnectivity_FromProto(mapCtx, in.GetStaticIpConnectivity())
	// MISSING: Duration
	// MISSING: Error
	out.SourceDatabase = DatabaseType_FromProto(mapCtx, in.GetSourceDatabase())
	out.DestinationDatabase = DatabaseType_FromProto(mapCtx, in.GetDestinationDatabase())
	// MISSING: EndTime
	out.ConversionWorkspace = ConversionWorkspaceInfo_FromProto(mapCtx, in.GetConversionWorkspace())
	out.Filter = direct.LazyPtr(in.GetFilter())
	if in.GetCmekKeyName() != "" {
		out.CmekKeyNameRef = &refsv1beta1.KMSCryptoKeyRef{External: in.GetCmekKeyName()}
	}
	out.PerformanceConfig = MigrationJob_PerformanceConfig_FromProto(mapCtx, in.GetPerformanceConfig())
	return out
}
func CloudDMSMigrationJobSpec_ToProto(mapCtx *direct.MapContext, in *krm.CloudDMSMigrationJobSpec) *pb.MigrationJob {
	if in == nil {
		return nil
	}
	out := &pb.MigrationJob{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: State
	// MISSING: Phase
	out.Type = direct.Enum_ToProto[pb.MigrationJob_Type](mapCtx, in.Type)
	out.DumpPath = direct.ValueOf(in.DumpPath)
	out.DumpFlags = MigrationJob_DumpFlags_ToProto(mapCtx, in.DumpFlags)
	if in.SourceRef != nil {
		out.Source = in.SourceRef.External
	}
	if in.DestinationRef != nil {
		out.Destination = in.DestinationRef.External
	}
	if oneof := ReverseSSHConnectivity_ToProto(mapCtx, in.ReverseSSHConnectivity); oneof != nil {
		out.Connectivity = &pb.MigrationJob_ReverseSshConnectivity{ReverseSshConnectivity: oneof}
	}
	if oneof := VPCPeeringConnectivity_ToProto(mapCtx, in.VPCPeeringConnectivity); oneof != nil {
		out.Connectivity = &pb.MigrationJob_VpcPeeringConnectivity{VpcPeeringConnectivity: oneof}
	}
	if oneof := StaticIPConnectivity_ToProto(mapCtx, in.StaticIPConnectivity); oneof != nil {
		out.Connectivity = &pb.MigrationJob_StaticIpConnectivity{StaticIpConnectivity: oneof}
	}
	// MISSING: Duration
	// MISSING: Error
	out.SourceDatabase = DatabaseType_ToProto(mapCtx, in.SourceDatabase)
	out.DestinationDatabase = DatabaseType_ToProto(mapCtx, in.DestinationDatabase)
	// MISSING: EndTime
	out.ConversionWorkspace = ConversionWorkspaceInfo_ToProto(mapCtx, in.ConversionWorkspace)
	out.Filter = direct.ValueOf(in.Filter)
	if in.CmekKeyNameRef != nil {
		out.CmekKeyName = in.CmekKeyNameRef.External
	}
	out.PerformanceConfig = MigrationJob_PerformanceConfig_ToProto(mapCtx, in.PerformanceConfig)
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
func DatabaseEngineInfo_FromProto(mapCtx *direct.MapContext, in *pb.DatabaseEngineInfo) *krm.DatabaseEngineInfo {
	if in == nil {
		return nil
	}
	out := &krm.DatabaseEngineInfo{}
	out.Engine = direct.Enum_FromProto(mapCtx, in.GetEngine())
	out.Version = direct.LazyPtr(in.GetVersion())
	return out
}
func DatabaseEngineInfo_ToProto(mapCtx *direct.MapContext, in *krm.DatabaseEngineInfo) *pb.DatabaseEngineInfo {
	if in == nil {
		return nil
	}
	out := &pb.DatabaseEngineInfo{}
	out.Engine = direct.Enum_ToProto[pb.DatabaseEngine](mapCtx, in.Engine)
	out.Version = direct.ValueOf(in.Version)
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
func ReverseSSHConnectivity_FromProto(mapCtx *direct.MapContext, in *pb.ReverseSshConnectivity) *krm.ReverseSSHConnectivity {
	if in == nil {
		return nil
	}
	out := &krm.ReverseSSHConnectivity{}
	out.VMIP = direct.LazyPtr(in.GetVmIp())
	out.VMPort = direct.LazyPtr(in.GetVmPort())
	if in.GetVm() != "" {
		out.VMRef = &krmcomputev1beta1.InstanceRef{External: in.GetVm()}
	}
	if in.GetVpc() != "" {
		out.VPCRef = &refsv1beta1.ComputeNetworkRef{External: in.GetVpc()}
	}
	return out
}
func ReverseSSHConnectivity_ToProto(mapCtx *direct.MapContext, in *krm.ReverseSSHConnectivity) *pb.ReverseSshConnectivity {
	if in == nil {
		return nil
	}
	out := &pb.ReverseSshConnectivity{}
	out.VmIp = direct.ValueOf(in.VMIP)
	out.VmPort = direct.ValueOf(in.VMPort)
	if in.VMRef != nil {
		out.Vm = in.VMRef.External
	}
	if in.VPCRef != nil {
		out.Vpc = in.VPCRef.External
	}
	return out
}
func StaticIPConnectivity_FromProto(mapCtx *direct.MapContext, in *pb.StaticIpConnectivity) *krm.StaticIPConnectivity {
	if in == nil {
		return nil
	}
	out := &krm.StaticIPConnectivity{}
	return out
}
func StaticIPConnectivity_ToProto(mapCtx *direct.MapContext, in *krm.StaticIPConnectivity) *pb.StaticIpConnectivity {
	if in == nil {
		return nil
	}
	out := &pb.StaticIpConnectivity{}
	return out
}
func VPCPeeringConnectivity_FromProto(mapCtx *direct.MapContext, in *pb.VpcPeeringConnectivity) *krm.VPCPeeringConnectivity {
	if in == nil {
		return nil
	}
	out := &krm.VPCPeeringConnectivity{}
	if in.GetVpc() != "" {
		out.VPCRef = &refsv1beta1.ComputeNetworkRef{External: in.GetVpc()}
	}
	return out
}
func VPCPeeringConnectivity_ToProto(mapCtx *direct.MapContext, in *krm.VPCPeeringConnectivity) *pb.VpcPeeringConnectivity {
	if in == nil {
		return nil
	}
	out := &pb.VpcPeeringConnectivity{}
	if in.VPCRef != nil {
		out.Vpc = in.VPCRef.External
	}
	return out
}
