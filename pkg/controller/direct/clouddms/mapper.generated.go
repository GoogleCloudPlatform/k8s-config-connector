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
func PrivateConnection_FromProto(mapCtx *direct.MapContext, in *pb.PrivateConnection) *krm.PrivateConnection {
	if in == nil {
		return nil
	}
	out := &krm.PrivateConnection{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: State
	// MISSING: Error
	out.VpcPeeringConfig = VpcPeeringConfig_FromProto(mapCtx, in.GetVpcPeeringConfig())
	return out
}
func PrivateConnection_ToProto(mapCtx *direct.MapContext, in *krm.PrivateConnection) *pb.PrivateConnection {
	if in == nil {
		return nil
	}
	out := &pb.PrivateConnection{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: State
	// MISSING: Error
	if oneof := VpcPeeringConfig_ToProto(mapCtx, in.VpcPeeringConfig); oneof != nil {
		out.Connectivity = &pb.PrivateConnection_VpcPeeringConfig{VpcPeeringConfig: oneof}
	}
	return out
}
func PrivateConnectionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PrivateConnection) *krm.PrivateConnectionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PrivateConnectionObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: DisplayName
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Error = Status_FromProto(mapCtx, in.GetError())
	// MISSING: VpcPeeringConfig
	return out
}
func PrivateConnectionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PrivateConnectionObservedState) *pb.PrivateConnection {
	if in == nil {
		return nil
	}
	out := &pb.PrivateConnection{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: DisplayName
	out.State = direct.Enum_ToProto[pb.PrivateConnection_State](mapCtx, in.State)
	out.Error = Status_ToProto(mapCtx, in.Error)
	// MISSING: VpcPeeringConfig
	return out
}
func VpcPeeringConfig_FromProto(mapCtx *direct.MapContext, in *pb.VpcPeeringConfig) *krm.VpcPeeringConfig {
	if in == nil {
		return nil
	}
	out := &krm.VpcPeeringConfig{}
	out.VpcName = direct.LazyPtr(in.GetVpcName())
	out.Subnet = direct.LazyPtr(in.GetSubnet())
	return out
}
func VpcPeeringConfig_ToProto(mapCtx *direct.MapContext, in *krm.VpcPeeringConfig) *pb.VpcPeeringConfig {
	if in == nil {
		return nil
	}
	out := &pb.VpcPeeringConfig{}
	out.VpcName = direct.ValueOf(in.VpcName)
	out.Subnet = direct.ValueOf(in.Subnet)
	return out
}
