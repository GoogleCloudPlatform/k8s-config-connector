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

package networkservices

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/networkservices/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/networkservices/apiv1/networkservicespb"
)
func Mesh_FromProto(mapCtx *direct.MapContext, in *pb.Mesh) *krm.Mesh {
	if in == nil {
		return nil
	}
	out := &krm.Mesh{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: SelfLink
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.LazyPtr(in.GetDescription())
	out.InterceptionPort = direct.LazyPtr(in.GetInterceptionPort())
	return out
}
func Mesh_ToProto(mapCtx *direct.MapContext, in *krm.Mesh) *pb.Mesh {
	if in == nil {
		return nil
	}
	out := &pb.Mesh{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: SelfLink
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.ValueOf(in.Description)
	out.InterceptionPort = direct.ValueOf(in.InterceptionPort)
	return out
}
func MeshObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Mesh) *krm.MeshObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MeshObservedState{}
	// MISSING: Name
	out.SelfLink = direct.LazyPtr(in.GetSelfLink())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: Description
	// MISSING: InterceptionPort
	return out
}
func MeshObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MeshObservedState) *pb.Mesh {
	if in == nil {
		return nil
	}
	out := &pb.Mesh{}
	// MISSING: Name
	out.SelfLink = direct.ValueOf(in.SelfLink)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: Description
	// MISSING: InterceptionPort
	return out
}
func NetworkservicesMeshObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Mesh) *krm.NetworkservicesMeshObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetworkservicesMeshObservedState{}
	// MISSING: Name
	// MISSING: SelfLink
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: InterceptionPort
	return out
}
func NetworkservicesMeshObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetworkservicesMeshObservedState) *pb.Mesh {
	if in == nil {
		return nil
	}
	out := &pb.Mesh{}
	// MISSING: Name
	// MISSING: SelfLink
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: InterceptionPort
	return out
}
func NetworkservicesMeshSpec_FromProto(mapCtx *direct.MapContext, in *pb.Mesh) *krm.NetworkservicesMeshSpec {
	if in == nil {
		return nil
	}
	out := &krm.NetworkservicesMeshSpec{}
	// MISSING: Name
	// MISSING: SelfLink
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: InterceptionPort
	return out
}
func NetworkservicesMeshSpec_ToProto(mapCtx *direct.MapContext, in *krm.NetworkservicesMeshSpec) *pb.Mesh {
	if in == nil {
		return nil
	}
	out := &pb.Mesh{}
	// MISSING: Name
	// MISSING: SelfLink
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: InterceptionPort
	return out
}
