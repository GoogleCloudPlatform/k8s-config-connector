// Copyright 2026 Google LLC
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

package compute

import (
	pb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func NodeGroupAutoscalingPolicy_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.NodeGroupAutoscalingPolicy) *krm.NodeGroupAutoscalingPolicy {
	if in == nil {
		return nil
	}
	out := &krm.NodeGroupAutoscalingPolicy{}
	out.MaxNodes = direct.PtrInt32ToPtrInt64(in.MaxNodes)
	out.MinNodes = direct.PtrInt32ToPtrInt64(in.MinNodes)
	out.Mode = in.Mode
	return out
}

func NodeGroupAutoscalingPolicy_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.NodeGroupAutoscalingPolicy) *pb.NodeGroupAutoscalingPolicy {
	if in == nil {
		return nil
	}
	out := &pb.NodeGroupAutoscalingPolicy{}
	out.MaxNodes = direct.PtrInt64ToPtrInt32(in.MaxNodes)
	out.MinNodes = direct.PtrInt64ToPtrInt32(in.MinNodes)
	out.Mode = in.Mode
	return out
}

func NodeGroupMaintenanceWindow_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.NodeGroupMaintenanceWindow) *krm.NodeGroupMaintenanceWindow {
	if in == nil {
		return nil
	}
	out := &krm.NodeGroupMaintenanceWindow{}
	out.StartTime = in.GetStartTime()
	return out
}

func NodeGroupMaintenanceWindow_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.NodeGroupMaintenanceWindow) *pb.NodeGroupMaintenanceWindow {
	if in == nil {
		return nil
	}
	out := &pb.NodeGroupMaintenanceWindow{}
	if in.StartTime != "" {
		out.StartTime = &in.StartTime
	}
	return out
}

func NodeGroupShareSettings_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.ShareSettings) *krm.NodeGroupShareSettings {
	if in == nil {
		return nil
	}
	out := &krm.NodeGroupShareSettings{}
	out.ShareType = in.GetShareType()
	for k, v := range in.ProjectMap {
		out.ProjectMap = append(out.ProjectMap, krm.NodeGroupShareSettingsProjectMap{
			IDRef: refs.ProjectRef{
				External: k,
			},
			ProjectIDRef: refs.ProjectRef{
				External: v.GetProjectId(),
			},
		})
	}
	return out
}

func NodeGroupShareSettings_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.NodeGroupShareSettings) *pb.ShareSettings {
	if in == nil {
		return nil
	}
	out := &pb.ShareSettings{}
	if in.ShareType != "" {
		out.ShareType = &in.ShareType
	}
	if len(in.ProjectMap) > 0 {
		out.ProjectMap = make(map[string]*pb.ShareSettingsProjectConfig)
		for _, entry := range in.ProjectMap {
			// Create local variables because we need pointers to strings
			externalID := entry.IDRef.External
			projectID := entry.ProjectIDRef.External
			var projectIDPtr *string
			if projectID != "" {
				projectIDPtr = &projectID
			}
			out.ProjectMap[externalID] = &pb.ShareSettingsProjectConfig{
				ProjectId: projectIDPtr,
			}
		}
	}
	return out
}

func ComputeNodeGroupSpec_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.NodeGroup) *krm.ComputeNodeGroupSpec {
	if in == nil {
		return nil
	}
	out := &krm.ComputeNodeGroupSpec{}
	out.AutoscalingPolicy = NodeGroupAutoscalingPolicy_v1beta1_FromProto(mapCtx, in.GetAutoscalingPolicy())
	out.Description = in.Description
	out.MaintenancePolicy = in.MaintenancePolicy
	out.MaintenanceWindow = NodeGroupMaintenanceWindow_v1beta1_FromProto(mapCtx, in.GetMaintenanceWindow())
	if in.GetNodeTemplate() != "" {
		out.NodeTemplateRef = &krm.ComputeNodeTemplateRef{External: in.GetNodeTemplate()}
	}
	out.ShareSettings = NodeGroupShareSettings_v1beta1_FromProto(mapCtx, in.GetShareSettings())
	out.Size = direct.PtrInt32ToPtrInt64(in.Size)
	out.Zone = in.GetZone()
	return out
}

func ComputeNodeGroupSpec_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeNodeGroupSpec) *pb.NodeGroup {
	if in == nil {
		return nil
	}
	out := &pb.NodeGroup{}
	out.AutoscalingPolicy = NodeGroupAutoscalingPolicy_v1beta1_ToProto(mapCtx, in.AutoscalingPolicy)
	out.Description = in.Description
	out.MaintenancePolicy = in.MaintenancePolicy
	out.MaintenanceWindow = NodeGroupMaintenanceWindow_v1beta1_ToProto(mapCtx, in.MaintenanceWindow)
	if in.NodeTemplateRef != nil {
		out.NodeTemplate = &in.NodeTemplateRef.External
	}
	out.ShareSettings = NodeGroupShareSettings_v1beta1_ToProto(mapCtx, in.ShareSettings)
	out.Size = direct.PtrInt64ToPtrInt32(in.Size)
	if in.Zone != "" {
		out.Zone = &in.Zone
	}
	return out
}

func ComputeNodeGroupStatus_v1beta1_FromProto(mapCtx *direct.MapContext, in *pb.NodeGroup) *krm.ComputeNodeGroupStatus {
	if in == nil {
		return nil
	}
	out := &krm.ComputeNodeGroupStatus{}
	out.CreationTimestamp = in.CreationTimestamp
	out.SelfLink = in.SelfLink
	return out
}

func ComputeNodeGroupStatus_v1beta1_ToProto(mapCtx *direct.MapContext, in *krm.ComputeNodeGroupStatus) *pb.NodeGroup {
	if in == nil {
		return nil
	}
	out := &pb.NodeGroup{}
	out.CreationTimestamp = in.CreationTimestamp
	out.SelfLink = in.SelfLink
	return out
}
