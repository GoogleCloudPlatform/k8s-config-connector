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

package securitycenter

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/securitycenter/apiv2/securitycenterpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/securitycenter/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func AttackPath_FromProto(mapCtx *direct.MapContext, in *pb.AttackPath) *krm.AttackPath {
	if in == nil {
		return nil
	}
	out := &krm.AttackPath{}
	out.Name = direct.LazyPtr(in.GetName())
	out.PathNodes = direct.Slice_FromProto(mapCtx, in.PathNodes, AttackPath_AttackPathNode_FromProto)
	out.Edges = direct.Slice_FromProto(mapCtx, in.Edges, AttackPath_AttackPathEdge_FromProto)
	return out
}
func AttackPath_ToProto(mapCtx *direct.MapContext, in *krm.AttackPath) *pb.AttackPath {
	if in == nil {
		return nil
	}
	out := &pb.AttackPath{}
	out.Name = direct.ValueOf(in.Name)
	out.PathNodes = direct.Slice_ToProto(mapCtx, in.PathNodes, AttackPath_AttackPathNode_ToProto)
	out.Edges = direct.Slice_ToProto(mapCtx, in.Edges, AttackPath_AttackPathEdge_ToProto)
	return out
}
func AttackPath_AttackPathEdge_FromProto(mapCtx *direct.MapContext, in *pb.AttackPath_AttackPathEdge) *krm.AttackPath_AttackPathEdge {
	if in == nil {
		return nil
	}
	out := &krm.AttackPath_AttackPathEdge{}
	out.Source = direct.LazyPtr(in.GetSource())
	out.Destination = direct.LazyPtr(in.GetDestination())
	return out
}
func AttackPath_AttackPathEdge_ToProto(mapCtx *direct.MapContext, in *krm.AttackPath_AttackPathEdge) *pb.AttackPath_AttackPathEdge {
	if in == nil {
		return nil
	}
	out := &pb.AttackPath_AttackPathEdge{}
	out.Source = direct.ValueOf(in.Source)
	out.Destination = direct.ValueOf(in.Destination)
	return out
}
func AttackPath_AttackPathNode_FromProto(mapCtx *direct.MapContext, in *pb.AttackPath_AttackPathNode) *krm.AttackPath_AttackPathNode {
	if in == nil {
		return nil
	}
	out := &krm.AttackPath_AttackPathNode{}
	out.Resource = direct.LazyPtr(in.GetResource())
	out.ResourceType = direct.LazyPtr(in.GetResourceType())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.AssociatedFindings = direct.Slice_FromProto(mapCtx, in.AssociatedFindings, AttackPath_AttackPathNode_PathNodeAssociatedFinding_FromProto)
	out.Uuid = direct.LazyPtr(in.GetUuid())
	out.AttackSteps = direct.Slice_FromProto(mapCtx, in.AttackSteps, AttackPath_AttackPathNode_AttackStepNode_FromProto)
	return out
}
func AttackPath_AttackPathNode_ToProto(mapCtx *direct.MapContext, in *krm.AttackPath_AttackPathNode) *pb.AttackPath_AttackPathNode {
	if in == nil {
		return nil
	}
	out := &pb.AttackPath_AttackPathNode{}
	out.Resource = direct.ValueOf(in.Resource)
	out.ResourceType = direct.ValueOf(in.ResourceType)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.AssociatedFindings = direct.Slice_ToProto(mapCtx, in.AssociatedFindings, AttackPath_AttackPathNode_PathNodeAssociatedFinding_ToProto)
	out.Uuid = direct.ValueOf(in.Uuid)
	out.AttackSteps = direct.Slice_ToProto(mapCtx, in.AttackSteps, AttackPath_AttackPathNode_AttackStepNode_ToProto)
	return out
}
func AttackPath_AttackPathNode_AttackStepNode_FromProto(mapCtx *direct.MapContext, in *pb.AttackPath_AttackPathNode_AttackStepNode) *krm.AttackPath_AttackPathNode_AttackStepNode {
	if in == nil {
		return nil
	}
	out := &krm.AttackPath_AttackPathNode_AttackStepNode{}
	out.Uuid = direct.LazyPtr(in.GetUuid())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Labels = in.Labels
	out.Description = direct.LazyPtr(in.GetDescription())
	return out
}
func AttackPath_AttackPathNode_AttackStepNode_ToProto(mapCtx *direct.MapContext, in *krm.AttackPath_AttackPathNode_AttackStepNode) *pb.AttackPath_AttackPathNode_AttackStepNode {
	if in == nil {
		return nil
	}
	out := &pb.AttackPath_AttackPathNode_AttackStepNode{}
	out.Uuid = direct.ValueOf(in.Uuid)
	out.Type = direct.Enum_ToProto[pb.AttackPath_AttackPathNode_NodeType](mapCtx, in.Type)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Labels = in.Labels
	out.Description = direct.ValueOf(in.Description)
	return out
}
func AttackPath_AttackPathNode_PathNodeAssociatedFinding_FromProto(mapCtx *direct.MapContext, in *pb.AttackPath_AttackPathNode_PathNodeAssociatedFinding) *krm.AttackPath_AttackPathNode_PathNodeAssociatedFinding {
	if in == nil {
		return nil
	}
	out := &krm.AttackPath_AttackPathNode_PathNodeAssociatedFinding{}
	out.CanonicalFinding = direct.LazyPtr(in.GetCanonicalFinding())
	out.FindingCategory = direct.LazyPtr(in.GetFindingCategory())
	out.Name = direct.LazyPtr(in.GetName())
	return out
}
func AttackPath_AttackPathNode_PathNodeAssociatedFinding_ToProto(mapCtx *direct.MapContext, in *krm.AttackPath_AttackPathNode_PathNodeAssociatedFinding) *pb.AttackPath_AttackPathNode_PathNodeAssociatedFinding {
	if in == nil {
		return nil
	}
	out := &pb.AttackPath_AttackPathNode_PathNodeAssociatedFinding{}
	out.CanonicalFinding = direct.ValueOf(in.CanonicalFinding)
	out.FindingCategory = direct.ValueOf(in.FindingCategory)
	out.Name = direct.ValueOf(in.Name)
	return out
}
