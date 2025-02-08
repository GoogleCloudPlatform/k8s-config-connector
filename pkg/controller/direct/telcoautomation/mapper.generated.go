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

package telcoautomation

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/telcoautomation/apiv1/telcoautomationpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/telcoautomation/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func File_FromProto(mapCtx *direct.MapContext, in *pb.File) *krm.File {
	if in == nil {
		return nil
	}
	out := &krm.File{}
	out.Path = direct.LazyPtr(in.GetPath())
	out.Content = direct.LazyPtr(in.GetContent())
	out.Deleted = direct.LazyPtr(in.GetDeleted())
	out.Editable = direct.LazyPtr(in.GetEditable())
	return out
}
func File_ToProto(mapCtx *direct.MapContext, in *krm.File) *pb.File {
	if in == nil {
		return nil
	}
	out := &pb.File{}
	out.Path = direct.ValueOf(in.Path)
	out.Content = direct.ValueOf(in.Content)
	out.Deleted = direct.ValueOf(in.Deleted)
	out.Editable = direct.ValueOf(in.Editable)
	return out
}
func HydratedDeployment_FromProto(mapCtx *direct.MapContext, in *pb.HydratedDeployment) *krm.HydratedDeployment {
	if in == nil {
		return nil
	}
	out := &krm.HydratedDeployment{}
	// MISSING: Name
	// MISSING: State
	out.Files = direct.Slice_FromProto(mapCtx, in.Files, File_FromProto)
	// MISSING: WorkloadCluster
	return out
}
func HydratedDeployment_ToProto(mapCtx *direct.MapContext, in *krm.HydratedDeployment) *pb.HydratedDeployment {
	if in == nil {
		return nil
	}
	out := &pb.HydratedDeployment{}
	// MISSING: Name
	// MISSING: State
	out.Files = direct.Slice_ToProto(mapCtx, in.Files, File_ToProto)
	// MISSING: WorkloadCluster
	return out
}
func HydratedDeploymentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.HydratedDeployment) *krm.HydratedDeploymentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.HydratedDeploymentObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: Files
	out.WorkloadCluster = direct.LazyPtr(in.GetWorkloadCluster())
	return out
}
func HydratedDeploymentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.HydratedDeploymentObservedState) *pb.HydratedDeployment {
	if in == nil {
		return nil
	}
	out := &pb.HydratedDeployment{}
	out.Name = direct.ValueOf(in.Name)
	out.State = direct.Enum_ToProto[pb.HydratedDeployment_State](mapCtx, in.State)
	// MISSING: Files
	out.WorkloadCluster = direct.ValueOf(in.WorkloadCluster)
	return out
}
func TelcoautomationHydratedDeploymentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.HydratedDeployment) *krm.TelcoautomationHydratedDeploymentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TelcoautomationHydratedDeploymentObservedState{}
	// MISSING: Name
	// MISSING: State
	// MISSING: Files
	// MISSING: WorkloadCluster
	return out
}
func TelcoautomationHydratedDeploymentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TelcoautomationHydratedDeploymentObservedState) *pb.HydratedDeployment {
	if in == nil {
		return nil
	}
	out := &pb.HydratedDeployment{}
	// MISSING: Name
	// MISSING: State
	// MISSING: Files
	// MISSING: WorkloadCluster
	return out
}
func TelcoautomationHydratedDeploymentSpec_FromProto(mapCtx *direct.MapContext, in *pb.HydratedDeployment) *krm.TelcoautomationHydratedDeploymentSpec {
	if in == nil {
		return nil
	}
	out := &krm.TelcoautomationHydratedDeploymentSpec{}
	// MISSING: Name
	// MISSING: State
	// MISSING: Files
	// MISSING: WorkloadCluster
	return out
}
func TelcoautomationHydratedDeploymentSpec_ToProto(mapCtx *direct.MapContext, in *krm.TelcoautomationHydratedDeploymentSpec) *pb.HydratedDeployment {
	if in == nil {
		return nil
	}
	out := &pb.HydratedDeployment{}
	// MISSING: Name
	// MISSING: State
	// MISSING: Files
	// MISSING: WorkloadCluster
	return out
}
