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

// +tool:fuzz-gen
// proto.message: google.cloud.run.v2.Job
// api.group: run.cnrm.cloud.google.com

package run

import (
	pb "cloud.google.com/go/run/apiv2/runpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/run/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func RunJobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Job) *krm.RunJobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.RunJobObservedState{}
	// MISSING: Name
	out.Uid = direct.LazyPtr(in.GetUid())
	// MISSING: Generation
	// MISSING: Labels
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.DeleteTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDeleteTime())
	out.ExpireTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpireTime())
	out.Creator = direct.LazyPtr(in.GetCreator())
	out.LastModifier = direct.LazyPtr(in.GetLastModifier())
	// MISSING: ObservedGeneration
	if v := in.GetTerminalCondition(); v != nil {
		out.TerminalCondition = []*krm.Condition{Condition_FromProto(mapCtx, v)}
	}
	// MISSING: Conditions
	out.ExecutionCount = direct.LazyPtr(in.GetExecutionCount())
	if v := in.GetLatestCreatedExecution(); v != nil {
		out.LatestCreatedExecution = []*krm.ExecutionReference{ExecutionReference_FromProto(mapCtx, v)}
		out.LatestCreatedExecution = []*krm.
			ExecutionReference{ExecutionReference_FromProto(mapCtx, v)}
	}
	out.Reconciling = direct.LazyPtr(in.GetReconciling())
	// MISSING: SatisfiesPzs
	// MISSING: StartExecutionToken
	// MISSING: RunExecutionToken
	out.Etag = direct.LazyPtr(in.GetEtag())
	return out
}
func RunJobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.RunJobObservedState) *pb.Job {
	if in == nil {
		return nil
	}
	out := &pb.Job{}
	// MISSING: Name
	out.Uid = direct.ValueOf(in.Uid)
	// MISSING: Generation
	// MISSING: Labels
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.DeleteTime = direct.StringTimestamp_ToProto(mapCtx, in.DeleteTime)
	out.ExpireTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpireTime)
	out.Creator = direct.ValueOf(in.Creator)
	out.LastModifier = direct.ValueOf(in.LastModifier)
	// MISSING: ObservedGeneration
	if len(in.TerminalCondition) > 0 && in.TerminalCondition[0] != nil {
		out.TerminalCondition = Condition_ToProto(mapCtx, in.TerminalCondition[0])
	}
	// MISSING: Conditions
	out.ExecutionCount = direct.ValueOf(in.ExecutionCount)
	if len(in.LatestCreatedExecution) > 0 && in.LatestCreatedExecution[0] != nil {
		out.LatestCreatedExecution = ExecutionReference_ToProto(mapCtx, in.LatestCreatedExecution[0])
	}
	out.Reconciling = direct.ValueOf(in.Reconciling)
	// MISSING: SatisfiesPzs
	// MISSING: StartExecutionToken
	// MISSING: RunExecutionToken
	out.Etag = direct.ValueOf(in.Etag)
	return out
}
