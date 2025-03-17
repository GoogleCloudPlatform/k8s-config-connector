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

package notebooks

import (
	pb "cloud.google.com/go/notebooks/apiv1/notebookspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/notebooks/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ContainerImage_FromProto(mapCtx *direct.MapContext, in *pb.ContainerImage) *krm.ContainerImage {
	if in == nil {
		return nil
	}
	out := &krm.ContainerImage{}
	out.Repository = direct.LazyPtr(in.GetRepository())
	out.Tag = direct.LazyPtr(in.GetTag())
	return out
}
func ContainerImage_ToProto(mapCtx *direct.MapContext, in *krm.ContainerImage) *pb.ContainerImage {
	if in == nil {
		return nil
	}
	out := &pb.ContainerImage{}
	out.Repository = direct.ValueOf(in.Repository)
	out.Tag = direct.ValueOf(in.Tag)
	return out
}
func NotebooksEnvironmentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Environment) *krm.NotebooksEnvironmentObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NotebooksEnvironmentObservedState{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	return out
}
func NotebooksEnvironmentObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NotebooksEnvironmentObservedState) *pb.Environment {
	if in == nil {
		return nil
	}
	out := &pb.Environment{}
	// MISSING: Name
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	return out
}
func NotebooksEnvironmentSpec_FromProto(mapCtx *direct.MapContext, in *pb.Environment) *krm.NotebooksEnvironmentSpec {
	if in == nil {
		return nil
	}
	out := &krm.NotebooksEnvironmentSpec{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.VmImage = VmImage_FromProto(mapCtx, in.GetVmImage())
	out.ContainerImage = ContainerImage_FromProto(mapCtx, in.GetContainerImage())
	out.PostStartupScript = direct.LazyPtr(in.GetPostStartupScript())
	return out
}
func NotebooksEnvironmentSpec_ToProto(mapCtx *direct.MapContext, in *krm.NotebooksEnvironmentSpec) *pb.Environment {
	if in == nil {
		return nil
	}
	out := &pb.Environment{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	if oneof := VmImage_ToProto(mapCtx, in.VmImage); oneof != nil {
		out.ImageType = &pb.Environment_VmImage{VmImage: oneof}
	}
	if oneof := ContainerImage_ToProto(mapCtx, in.ContainerImage); oneof != nil {
		out.ImageType = &pb.Environment_ContainerImage{ContainerImage: oneof}
	}
	out.PostStartupScript = direct.ValueOf(in.PostStartupScript)
	return out
}
func VmImage_FromProto(mapCtx *direct.MapContext, in *pb.VmImage) *krm.VmImage {
	if in == nil {
		return nil
	}
	out := &krm.VmImage{}
	out.Project = direct.LazyPtr(in.GetProject())
	out.ImageName = direct.LazyPtr(in.GetImageName())
	out.ImageFamily = direct.LazyPtr(in.GetImageFamily())
	return out
}
func VmImage_ToProto(mapCtx *direct.MapContext, in *krm.VmImage) *pb.VmImage {
	if in == nil {
		return nil
	}
	out := &pb.VmImage{}
	out.Project = direct.ValueOf(in.Project)
	if oneof := VmImage_ImageName_ToProto(mapCtx, in.ImageName); oneof != nil {
		out.Image = oneof
	}
	if oneof := VmImage_ImageFamily_ToProto(mapCtx, in.ImageFamily); oneof != nil {
		out.Image = oneof
	}
	return out
}
