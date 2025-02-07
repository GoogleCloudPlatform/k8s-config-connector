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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/clouddms/apiv1/clouddmspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/clouddms/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func ClouddmsConnectionProfileObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ConnectionProfile) *krm.ClouddmsConnectionProfileObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ClouddmsConnectionProfileObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: DisplayName
	// MISSING: Mysql
	// MISSING: Postgresql
	// MISSING: Oracle
	// MISSING: Cloudsql
	// MISSING: Alloydb
	// MISSING: Error
	// MISSING: Provider
	return out
}
func ClouddmsConnectionProfileObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ClouddmsConnectionProfileObservedState) *pb.ConnectionProfile {
	if in == nil {
		return nil
	}
	out := &pb.ConnectionProfile{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: DisplayName
	// MISSING: Mysql
	// MISSING: Postgresql
	// MISSING: Oracle
	// MISSING: Cloudsql
	// MISSING: Alloydb
	// MISSING: Error
	// MISSING: Provider
	return out
}
func ClouddmsConnectionProfileSpec_FromProto(mapCtx *direct.MapContext, in *pb.ConnectionProfile) *krm.ClouddmsConnectionProfileSpec {
	if in == nil {
		return nil
	}
	out := &krm.ClouddmsConnectionProfileSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: DisplayName
	// MISSING: Mysql
	// MISSING: Postgresql
	// MISSING: Oracle
	// MISSING: Cloudsql
	// MISSING: Alloydb
	// MISSING: Error
	// MISSING: Provider
	return out
}
func ClouddmsConnectionProfileSpec_ToProto(mapCtx *direct.MapContext, in *krm.ClouddmsConnectionProfileSpec) *pb.ConnectionProfile {
	if in == nil {
		return nil
	}
	out := &pb.ConnectionProfile{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: State
	// MISSING: DisplayName
	// MISSING: Mysql
	// MISSING: Postgresql
	// MISSING: Oracle
	// MISSING: Cloudsql
	// MISSING: Alloydb
	// MISSING: Error
	// MISSING: Provider
	return out
}
func ClouddmsConversionWorkspaceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ConversionWorkspace) *krm.ClouddmsConversionWorkspaceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ClouddmsConversionWorkspaceObservedState{}
	// MISSING: Name
	// MISSING: Source
	// MISSING: Destination
	// MISSING: GlobalSettings
	// MISSING: HasUncommittedChanges
	// MISSING: LatestCommitID
	// MISSING: LatestCommitTime
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DisplayName
	return out
}
func ClouddmsConversionWorkspaceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ClouddmsConversionWorkspaceObservedState) *pb.ConversionWorkspace {
	if in == nil {
		return nil
	}
	out := &pb.ConversionWorkspace{}
	// MISSING: Name
	// MISSING: Source
	// MISSING: Destination
	// MISSING: GlobalSettings
	// MISSING: HasUncommittedChanges
	// MISSING: LatestCommitID
	// MISSING: LatestCommitTime
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DisplayName
	return out
}
func ClouddmsConversionWorkspaceSpec_FromProto(mapCtx *direct.MapContext, in *pb.ConversionWorkspace) *krm.ClouddmsConversionWorkspaceSpec {
	if in == nil {
		return nil
	}
	out := &krm.ClouddmsConversionWorkspaceSpec{}
	// MISSING: Name
	// MISSING: Source
	// MISSING: Destination
	// MISSING: GlobalSettings
	// MISSING: HasUncommittedChanges
	// MISSING: LatestCommitID
	// MISSING: LatestCommitTime
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DisplayName
	return out
}
func ClouddmsConversionWorkspaceSpec_ToProto(mapCtx *direct.MapContext, in *krm.ClouddmsConversionWorkspaceSpec) *pb.ConversionWorkspace {
	if in == nil {
		return nil
	}
	out := &pb.ConversionWorkspace{}
	// MISSING: Name
	// MISSING: Source
	// MISSING: Destination
	// MISSING: GlobalSettings
	// MISSING: HasUncommittedChanges
	// MISSING: LatestCommitID
	// MISSING: LatestCommitTime
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DisplayName
	return out
}
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
func ClouddmsPrivateConnectionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PrivateConnection) *krm.ClouddmsPrivateConnectionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ClouddmsPrivateConnectionObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: State
	// MISSING: Error
	// MISSING: VpcPeeringConfig
	return out
}
func ClouddmsPrivateConnectionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ClouddmsPrivateConnectionObservedState) *pb.PrivateConnection {
	if in == nil {
		return nil
	}
	out := &pb.PrivateConnection{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: State
	// MISSING: Error
	// MISSING: VpcPeeringConfig
	return out
}
func ClouddmsPrivateConnectionSpec_FromProto(mapCtx *direct.MapContext, in *pb.PrivateConnection) *krm.ClouddmsPrivateConnectionSpec {
	if in == nil {
		return nil
	}
	out := &krm.ClouddmsPrivateConnectionSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: State
	// MISSING: Error
	// MISSING: VpcPeeringConfig
	return out
}
func ClouddmsPrivateConnectionSpec_ToProto(mapCtx *direct.MapContext, in *krm.ClouddmsPrivateConnectionSpec) *pb.PrivateConnection {
	if in == nil {
		return nil
	}
	out := &pb.PrivateConnection{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: DisplayName
	// MISSING: State
	// MISSING: Error
	// MISSING: VpcPeeringConfig
	return out
}
func ConversionWorkspace_FromProto(mapCtx *direct.MapContext, in *pb.ConversionWorkspace) *krm.ConversionWorkspace {
	if in == nil {
		return nil
	}
	out := &krm.ConversionWorkspace{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Source = DatabaseEngineInfo_FromProto(mapCtx, in.GetSource())
	out.Destination = DatabaseEngineInfo_FromProto(mapCtx, in.GetDestination())
	out.GlobalSettings = in.GlobalSettings
	// MISSING: HasUncommittedChanges
	// MISSING: LatestCommitID
	// MISSING: LatestCommitTime
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	return out
}
func ConversionWorkspace_ToProto(mapCtx *direct.MapContext, in *krm.ConversionWorkspace) *pb.ConversionWorkspace {
	if in == nil {
		return nil
	}
	out := &pb.ConversionWorkspace{}
	out.Name = direct.ValueOf(in.Name)
	out.Source = DatabaseEngineInfo_ToProto(mapCtx, in.Source)
	out.Destination = DatabaseEngineInfo_ToProto(mapCtx, in.Destination)
	out.GlobalSettings = in.GlobalSettings
	// MISSING: HasUncommittedChanges
	// MISSING: LatestCommitID
	// MISSING: LatestCommitTime
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.DisplayName = direct.ValueOf(in.DisplayName)
	return out
}
func ConversionWorkspaceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ConversionWorkspace) *krm.ConversionWorkspaceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ConversionWorkspaceObservedState{}
	// MISSING: Name
	// MISSING: Source
	// MISSING: Destination
	// MISSING: GlobalSettings
	out.HasUncommittedChanges = direct.LazyPtr(in.GetHasUncommittedChanges())
	out.LatestCommitID = direct.LazyPtr(in.GetLatestCommitId())
	out.LatestCommitTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLatestCommitTime())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: DisplayName
	return out
}
func ConversionWorkspaceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ConversionWorkspaceObservedState) *pb.ConversionWorkspace {
	if in == nil {
		return nil
	}
	out := &pb.ConversionWorkspace{}
	// MISSING: Name
	// MISSING: Source
	// MISSING: Destination
	// MISSING: GlobalSettings
	out.HasUncommittedChanges = direct.ValueOf(in.HasUncommittedChanges)
	out.LatestCommitId = direct.ValueOf(in.LatestCommitID)
	out.LatestCommitTime = direct.StringTimestamp_ToProto(mapCtx, in.LatestCommitTime)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: DisplayName
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
