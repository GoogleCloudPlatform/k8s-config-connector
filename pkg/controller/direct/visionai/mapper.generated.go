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

package visionai

import (
	pb "cloud.google.com/go/visionai/apiv1/visionaipb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/visionai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func Cluster_FromProto(mapCtx *direct.MapContext, in *pb.Cluster) *krm.Cluster {
	if in == nil {
		return nil
	}
	out := &krm.Cluster{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Annotations = in.Annotations
	// MISSING: DataplaneServiceEndpoint
	// MISSING: State
	// MISSING: PscTarget
	return out
}
func Cluster_ToProto(mapCtx *direct.MapContext, in *krm.Cluster) *pb.Cluster {
	if in == nil {
		return nil
	}
	out := &pb.Cluster{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Annotations = in.Annotations
	// MISSING: DataplaneServiceEndpoint
	// MISSING: State
	// MISSING: PscTarget
	return out
}
func ClusterObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Cluster) *krm.ClusterObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ClusterObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: Annotations
	out.DataplaneServiceEndpoint = direct.LazyPtr(in.GetDataplaneServiceEndpoint())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.PscTarget = direct.LazyPtr(in.GetPscTarget())
	return out
}
func ClusterObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ClusterObservedState) *pb.Cluster {
	if in == nil {
		return nil
	}
	out := &pb.Cluster{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: Annotations
	out.DataplaneServiceEndpoint = direct.ValueOf(in.DataplaneServiceEndpoint)
	out.State = direct.Enum_ToProto[pb.Cluster_State](mapCtx, in.State)
	out.PscTarget = direct.ValueOf(in.PscTarget)
	return out
}
func VisionaiClusterObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Cluster) *krm.VisionaiClusterObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VisionaiClusterObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Annotations
	// MISSING: DataplaneServiceEndpoint
	// MISSING: State
	// MISSING: PscTarget
	return out
}
func VisionaiClusterObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VisionaiClusterObservedState) *pb.Cluster {
	if in == nil {
		return nil
	}
	out := &pb.Cluster{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Annotations
	// MISSING: DataplaneServiceEndpoint
	// MISSING: State
	// MISSING: PscTarget
	return out
}
func VisionaiClusterSpec_FromProto(mapCtx *direct.MapContext, in *pb.Cluster) *krm.VisionaiClusterSpec {
	if in == nil {
		return nil
	}
	out := &krm.VisionaiClusterSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Annotations
	// MISSING: DataplaneServiceEndpoint
	// MISSING: State
	// MISSING: PscTarget
	return out
}
func VisionaiClusterSpec_ToProto(mapCtx *direct.MapContext, in *krm.VisionaiClusterSpec) *pb.Cluster {
	if in == nil {
		return nil
	}
	out := &pb.Cluster{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Annotations
	// MISSING: DataplaneServiceEndpoint
	// MISSING: State
	// MISSING: PscTarget
	return out
}
