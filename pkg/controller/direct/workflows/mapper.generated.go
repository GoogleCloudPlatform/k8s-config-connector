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

package workflows

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/workflows/apiv1beta/workflowspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/workflows/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Workflow_FromProto(mapCtx *direct.MapContext, in *pb.Workflow) *krm.Workflow {
	if in == nil {
		return nil
	}
	out := &krm.Workflow{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: State
	// MISSING: RevisionID
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: RevisionCreateTime
	out.Labels = in.Labels
	out.ServiceAccount = direct.LazyPtr(in.GetServiceAccount())
	out.SourceContents = direct.LazyPtr(in.GetSourceContents())
	return out
}
func Workflow_ToProto(mapCtx *direct.MapContext, in *krm.Workflow) *pb.Workflow {
	if in == nil {
		return nil
	}
	out := &pb.Workflow{}
	out.Name = direct.ValueOf(in.Name)
	out.Description = direct.ValueOf(in.Description)
	// MISSING: State
	// MISSING: RevisionID
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: RevisionCreateTime
	out.Labels = in.Labels
	out.ServiceAccount = direct.ValueOf(in.ServiceAccount)
	if oneof := Workflow_SourceContents_ToProto(mapCtx, in.SourceContents); oneof != nil {
		out.SourceCode = oneof
	}
	return out
}
func WorkflowObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Workflow) *krm.WorkflowObservedState {
	if in == nil {
		return nil
	}
	out := &krm.WorkflowObservedState{}
	// MISSING: Name
	// MISSING: Description
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.RevisionID = direct.LazyPtr(in.GetRevisionId())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.RevisionCreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetRevisionCreateTime())
	// MISSING: Labels
	// MISSING: ServiceAccount
	// MISSING: SourceContents
	return out
}
func WorkflowObservedState_ToProto(mapCtx *direct.MapContext, in *krm.WorkflowObservedState) *pb.Workflow {
	if in == nil {
		return nil
	}
	out := &pb.Workflow{}
	// MISSING: Name
	// MISSING: Description
	out.State = direct.Enum_ToProto[pb.Workflow_State](mapCtx, in.State)
	out.RevisionId = direct.ValueOf(in.RevisionID)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.RevisionCreateTime = direct.StringTimestamp_ToProto(mapCtx, in.RevisionCreateTime)
	// MISSING: Labels
	// MISSING: ServiceAccount
	// MISSING: SourceContents
	return out
}
