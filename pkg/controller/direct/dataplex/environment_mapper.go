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

package dataplex

import (
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataplex/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DataplexEnvironmentSpec_FromProto(mapCtx *direct.MapContext, in *pb.Environment) *krm.Environment {
	if in == nil {
		return nil
	}
	out := &krm.Environment{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: State
	out.InfrastructureSpec = Environment_InfrastructureSpec_FromProto(mapCtx, in.GetInfrastructureSpec())
	out.SessionSpec = Environment_SessionSpec_FromProto(mapCtx, in.GetSessionSpec())
	// MISSING: SessionStatus
	// MISSING: Endpoints
	return out
}
func DataplexEnvironmentSpec_ToProto(mapCtx *direct.MapContext, in *krm.Environment) *pb.Environment {
	if in == nil {
		return nil
	}
	out := &pb.Environment{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.ValueOf(in.Description)
	// MISSING: State
	out.InfrastructureSpec = Environment_InfrastructureSpec_ToProto(mapCtx, in.InfrastructureSpec)
	out.SessionSpec = Environment_SessionSpec_ToProto(mapCtx, in.SessionSpec)
	// MISSING: SessionStatus
	// MISSING: Endpoints
	return out
}
func DataplexEnvironmentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Environment) *krm.EnvironmentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EnvironmentObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: DisplayName
	out.Uid = direct.LazyPtr(in.GetUid())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: Description
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: InfrastructureSpec
	// MISSING: SessionSpec
	out.SessionStatus = Environment_SessionStatus_FromProto(mapCtx, in.GetSessionStatus())
	out.Endpoints = Environment_Endpoints_FromProto(mapCtx, in.GetEndpoints())
	return out
}
func DataplexEnvironmentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EnvironmentObservedState) *pb.Environment {
	if in == nil {
		return nil
	}
	out := &pb.Environment{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: DisplayName
	out.Uid = direct.ValueOf(in.Uid)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: Description
	out.State = direct.Enum_ToProto[pb.State](mapCtx, in.State)
	// MISSING: InfrastructureSpec
	// MISSING: SessionSpec
	out.SessionStatus = Environment_SessionStatus_ToProto(mapCtx, in.SessionStatus)
	out.Endpoints = Environment_Endpoints_ToProto(mapCtx, in.Endpoints)
	return out
}
func Environment_Endpoints_FromProto(mapCtx *direct.MapContext, in *pb.Environment_Endpoints) *krm.Environment_Endpoints {
	if in == nil {
		return nil
	}
	out := &krm.Environment_Endpoints{}
	// MISSING: Notebooks
	// MISSING: SQL
	return out
}
func Environment_Endpoints_ToProto(mapCtx *direct.MapContext, in *krm.Environment_Endpoints) *pb.Environment_Endpoints {
	if in == nil {
		return nil
	}
	out := &pb.Environment_Endpoints{}
	// MISSING: Notebooks
	// MISSING: SQL
	return out
}
func Environment_EndpointsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Environment_Endpoints) *krm.Environment_EndpointsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Environment_EndpointsObservedState{}
	out.Notebooks = direct.LazyPtr(in.GetNotebooks())
	out.SQL = direct.LazyPtr(in.GetSql())
	return out
}
func Environment_EndpointsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Environment_EndpointsObservedState) *pb.Environment_Endpoints {
	if in == nil {
		return nil
	}
	out := &pb.Environment_Endpoints{}
	out.Notebooks = direct.ValueOf(in.Notebooks)
	out.Sql = direct.ValueOf(in.SQL)
	return out
}
func Environment_InfrastructureSpec_FromProto(mapCtx *direct.MapContext, in *pb.Environment_InfrastructureSpec) *krm.Environment_InfrastructureSpec {
	if in == nil {
		return nil
	}
	out := &krm.Environment_InfrastructureSpec{}
	out.Compute = Environment_InfrastructureSpec_ComputeResources_FromProto(mapCtx, in.GetCompute())
	out.OSImage = Environment_InfrastructureSpec_OSImageRuntime_FromProto(mapCtx, in.GetOsImage())
	return out
}
func Environment_InfrastructureSpec_ToProto(mapCtx *direct.MapContext, in *krm.Environment_InfrastructureSpec) *pb.Environment_InfrastructureSpec {
	if in == nil {
		return nil
	}
	out := &pb.Environment_InfrastructureSpec{}
	if oneof := Environment_InfrastructureSpec_ComputeResources_ToProto(mapCtx, in.Compute); oneof != nil {
		out.Resources = &pb.Environment_InfrastructureSpec_Compute{Compute: oneof}
	}
	if oneof := Environment_InfrastructureSpec_OSImageRuntime_ToProto(mapCtx, in.OSImage); oneof != nil {
		out.Runtime = &pb.Environment_InfrastructureSpec_OsImage{OsImage: oneof}
	}
	return out
}
func Environment_InfrastructureSpec_ComputeResources_FromProto(mapCtx *direct.MapContext, in *pb.Environment_InfrastructureSpec_ComputeResources) *krm.Environment_InfrastructureSpec_ComputeResources {
	if in == nil {
		return nil
	}
	out := &krm.Environment_InfrastructureSpec_ComputeResources{}
	out.DiskSizeGB = direct.LazyPtr(in.GetDiskSizeGb())
	out.NodeCount = direct.LazyPtr(in.GetNodeCount())
	out.MaxNodeCount = direct.LazyPtr(in.GetMaxNodeCount())
	return out
}
func Environment_InfrastructureSpec_ComputeResources_ToProto(mapCtx *direct.MapContext, in *krm.Environment_InfrastructureSpec_ComputeResources) *pb.Environment_InfrastructureSpec_ComputeResources {
	if in == nil {
		return nil
	}
	out := &pb.Environment_InfrastructureSpec_ComputeResources{}
	out.DiskSizeGb = direct.ValueOf(in.DiskSizeGB)
	out.NodeCount = direct.ValueOf(in.NodeCount)
	out.MaxNodeCount = direct.ValueOf(in.MaxNodeCount)
	return out
}
func Environment_InfrastructureSpec_OSImageRuntime_FromProto(mapCtx *direct.MapContext, in *pb.Environment_InfrastructureSpec_OsImageRuntime) *krm.Environment_InfrastructureSpec_OSImageRuntime {
	if in == nil {
		return nil
	}
	out := &krm.Environment_InfrastructureSpec_OSImageRuntime{}
	out.ImageVersion = direct.LazyPtr(in.GetImageVersion())
	out.JavaLibraries = in.JavaLibraries
	out.PythonPackages = in.PythonPackages
	out.Properties = in.Properties
	return out
}
func Environment_InfrastructureSpec_OSImageRuntime_ToProto(mapCtx *direct.MapContext, in *krm.Environment_InfrastructureSpec_OSImageRuntime) *pb.Environment_InfrastructureSpec_OsImageRuntime {
	if in == nil {
		return nil
	}
	out := &pb.Environment_InfrastructureSpec_OsImageRuntime{}
	out.ImageVersion = direct.ValueOf(in.ImageVersion)
	out.JavaLibraries = in.JavaLibraries
	out.PythonPackages = in.PythonPackages
	out.Properties = in.Properties
	return out
}
func Environment_SessionSpec_FromProto(mapCtx *direct.MapContext, in *pb.Environment_SessionSpec) *krm.Environment_SessionSpec {
	if in == nil {
		return nil
	}
	out := &krm.Environment_SessionSpec{}
	out.MaxIdleDuration = direct.StringDuration_FromProto(mapCtx, in.GetMaxIdleDuration())
	out.EnableFastStartup = direct.LazyPtr(in.GetEnableFastStartup())
	return out
}
func Environment_SessionSpec_ToProto(mapCtx *direct.MapContext, in *krm.Environment_SessionSpec) *pb.Environment_SessionSpec {
	if in == nil {
		return nil
	}
	out := &pb.Environment_SessionSpec{}
	out.MaxIdleDuration = direct.StringDuration_ToProto(mapCtx, in.MaxIdleDuration)
	out.EnableFastStartup = direct.ValueOf(in.EnableFastStartup)
	return out
}
func Environment_SessionStatus_FromProto(mapCtx *direct.MapContext, in *pb.Environment_SessionStatus) *krm.Environment_SessionStatus {
	if in == nil {
		return nil
	}
	out := &krm.Environment_SessionStatus{}
	// MISSING: Active
	return out
}
func Environment_SessionStatus_ToProto(mapCtx *direct.MapContext, in *krm.Environment_SessionStatus) *pb.Environment_SessionStatus {
	if in == nil {
		return nil
	}
	out := &pb.Environment_SessionStatus{}
	// MISSING: Active
	return out
}
func Environment_SessionStatusObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Environment_SessionStatus) *krm.Environment_SessionStatusObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Environment_SessionStatusObservedState{}
	out.Active = direct.LazyPtr(in.GetActive())
	return out
}
func Environment_SessionStatusObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Environment_SessionStatusObservedState) *pb.Environment_SessionStatus {
	if in == nil {
		return nil
	}
	out := &pb.Environment_SessionStatus{}
	out.Active = direct.ValueOf(in.Active)
	return out
}
