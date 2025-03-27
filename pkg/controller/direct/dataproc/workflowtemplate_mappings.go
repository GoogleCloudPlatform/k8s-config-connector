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

package dataproc

import (
	pb "cloud.google.com/go/dataproc/v2/apiv1/dataprocpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataproc/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DataprocWorkflowTemplateObservedState_FromProto(mapCtx *direct.MapContext, in *pb.WorkflowTemplate) *krm.DataprocWorkflowTemplateObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataprocWorkflowTemplateObservedState{}
	out.Name = direct.LazyPtr(in.Name)
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.Placement = WorkflowTemplatePlacementObservedState_FromProto(mapCtx, in.GetPlacement())
	return out
}
func DataprocWorkflowTemplateObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataprocWorkflowTemplateObservedState) *pb.WorkflowTemplate {
	if in == nil {
		return nil
	}
	out := &pb.WorkflowTemplate{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.Placement = WorkflowTemplatePlacementObservedState_ToProto(mapCtx, in.Placement)
	return out
}
func DataprocWorkflowTemplateSpec_FromProto(mapCtx *direct.MapContext, in *pb.WorkflowTemplate) *krm.DataprocWorkflowTemplateSpec {
	if in == nil {
		return nil
	}
	out := &krm.DataprocWorkflowTemplateSpec{}
	out.ID = direct.LazyPtr(in.Id)
	out.Version = direct.LazyPtr(in.GetVersion())
	out.Labels = in.Labels
	out.Placement = WorkflowTemplatePlacement_FromProto(mapCtx, in.GetPlacement())
	out.Jobs = direct.Slice_FromProto(mapCtx, in.GetJobs(), OrderedJob_FromProto)
	out.Parameters = direct.Slice_FromProto(mapCtx, in.GetParameters(), TemplateParameter_FromProto)
	out.DagTimeout = direct.Duration_FromProto(mapCtx, in.GetDagTimeout())
	out.EncryptionConfig = WorkflowTemplate_EncryptionConfig_FromProto(mapCtx, in.GetEncryptionConfig())
	return out
}
func DataprocWorkflowTemplateSpec_ToProto(mapCtx *direct.MapContext, in *krm.DataprocWorkflowTemplateSpec) *pb.WorkflowTemplate {
	if in == nil {
		return nil
	}
	out := &pb.WorkflowTemplate{}
	out.Id = direct.ValueOf(in.ID)
	out.Version = direct.ValueOf(in.Version)
	out.Labels = in.Labels
	out.Placement = WorkflowTemplatePlacement_ToProto(mapCtx, in.Placement)
	out.Jobs = direct.Slice_ToProto(mapCtx, in.Jobs, OrderedJob_ToProto)
	out.Parameters = direct.Slice_ToProto(mapCtx, in.Parameters, TemplateParameter_ToProto)
	out.DagTimeout = direct.Duration_ToProto(mapCtx, in.DagTimeout)
	out.EncryptionConfig = WorkflowTemplate_EncryptionConfig_ToProto(mapCtx, in.EncryptionConfig)
	return out
}
