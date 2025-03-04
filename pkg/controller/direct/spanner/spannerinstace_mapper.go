// Copyright 2024 Google LLC
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

package spanner

import (
	"strings"

	pb "cloud.google.com/go/spanner/admin/instance/apiv1/instancepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/spanner/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AutoscalingConfig_AutoscalingLimits_MinNodes_ToProto(mapCtx *direct.MapContext, m *int32) *pb.AutoscalingConfig_AutoscalingLimits_MinNodes {
	if m == nil {
		return nil
	}
	return &pb.AutoscalingConfig_AutoscalingLimits_MinNodes{
		MinNodes: direct.ValueOf(m),
	}
}
func AutoscalingConfig_AutoscalingLimits_MinProcessingUnits_ToProto(mapCtx *direct.MapContext, m *int32) *pb.AutoscalingConfig_AutoscalingLimits_MinProcessingUnits {
	if m == nil {
		return nil
	}
	return &pb.AutoscalingConfig_AutoscalingLimits_MinProcessingUnits{
		MinProcessingUnits: direct.ValueOf(m),
	}
}
func AutoscalingConfig_AutoscalingLimits_MaxNodes_ToProto(mapCtx *direct.MapContext, m *int32) *pb.AutoscalingConfig_AutoscalingLimits_MaxNodes {
	if m == nil {
		return nil
	}
	return &pb.AutoscalingConfig_AutoscalingLimits_MaxNodes{
		MaxNodes: direct.ValueOf(m),
	}
}
func AutoscalingConfig_AutoscalingLimits_MaxProcessingUnits_ToProto(mapCtx *direct.MapContext, m *int32) *pb.AutoscalingConfig_AutoscalingLimits_MaxProcessingUnits {
	if m == nil {
		return nil
	}
	return &pb.AutoscalingConfig_AutoscalingLimits_MaxProcessingUnits{
		MaxProcessingUnits: direct.ValueOf(m),
	}
}
func ReplicaComputeCapacity_NodeCount_ToProto(mapCtx *direct.MapContext, m *int32) *pb.ReplicaComputeCapacity_NodeCount {
	if m == nil {
		return nil
	}
	return &pb.ReplicaComputeCapacity_NodeCount{NodeCount: direct.ValueOf(m)}
}
func ReplicaComputeCapacity_ProcessingUnits_ToProto(mapCtx *direct.MapContext, m *int32) *pb.ReplicaComputeCapacity_ProcessingUnits {
	if m == nil {
		return nil
	}
	return &pb.ReplicaComputeCapacity_ProcessingUnits{ProcessingUnits: direct.ValueOf(m)}
}
func State_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *string {
	return direct.Enum_FromProto(mapCtx, in.GetState())
}

func SpannerInstanceSpec_FromProto(mapCtx *direct.MapContext, in *pb.Instance, configPrefix string) *krm.SpannerInstanceSpec {
	if in == nil {
		return nil
	}
	out := &krm.SpannerInstanceSpec{}
	out.Config = strings.TrimPrefix(in.GetConfig(), configPrefix)
	out.DisplayName = in.GetDisplayName()
	out.ProcessingUnits = direct.LazyPtr(in.GetProcessingUnits())
	out.NumNodes = direct.LazyPtr(in.GetNodeCount())
	out.Edition = direct.LazyPtr(in.Edition.String())
	out.AutoscalingConfig = AutoscalingConfig_FromProto(mapCtx, in.GetAutoscalingConfig())
	return out
}

func SpannerInstanceSpec_ToProto(mapCtx *direct.MapContext, in *krm.SpannerInstanceSpec, configPrefix string) *pb.Instance {
	if in == nil {
		return nil
	}
	out := &pb.Instance{}
	out.Config = configPrefix + in.Config
	out.DisplayName = in.DisplayName
	out.NodeCount = direct.ValueOf(in.NumNodes)
	out.ProcessingUnits = direct.ValueOf(in.ProcessingUnits)
	out.Edition = direct.Enum_ToProto[pb.Instance_Edition](mapCtx, in.Edition)
	if out.Edition == pb.Instance_EDITION_UNSPECIFIED {
		out.Edition = pb.Instance_STANDARD
	}
	out.AutoscalingConfig = AutoscalingConfig_ToProto(mapCtx, in.AutoscalingConfig)
	return out
}

func SpannerInstanceStatus_FromProto(mapCtx *direct.MapContext, in *pb.Instance) *krm.SpannerInstanceStatus {
	if in == nil {
		return nil
	}
	out := &krm.SpannerInstanceStatus{}
	out.State = State_FromProto(mapCtx, in)
	out.ObservedState = &krm.SpannerInstanceObservedState{
		NumNodes:        direct.LazyPtr(in.NodeCount),
		ProcessingUnits: direct.LazyPtr(in.ProcessingUnits),
	}
	return out
}
