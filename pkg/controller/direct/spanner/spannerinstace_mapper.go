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
	pb "cloud.google.com/go/spanner/admin/instance/apiv1/instancepb"
	// krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/spanner/v1beta1"
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
