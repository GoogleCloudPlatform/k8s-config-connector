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

// +generated:mapper
// krm.group: run.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.cloud.run.v2

package run

import (
	pb "cloud.google.com/go/run/apiv2/runpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/run/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func BinaryAuthorization_UseDefault_FromProto(mapCtx *direct.MapContext, in *pb.BinaryAuthorization_UseDefault) *bool {
	if in == nil {
		return nil
	}
	return &in.UseDefault
}
func BinaryAuthorization_UseDefault_ToProto(mapCtx *direct.MapContext, in *bool) *pb.BinaryAuthorization_UseDefault {
	if in == nil {
		return nil
	}
	out := &pb.BinaryAuthorization_UseDefault{}
	out.UseDefault = direct.ValueOf(in)
	return out
}

func EnvVar_Value_FromProto(mapCtx *direct.MapContext, in *pb.EnvVar_Value) *string {
	if in == nil {
		return nil
	}
	return &in.Value
}
func EnvVar_Value_ToProto(mapCtx *direct.MapContext, in *string) *pb.EnvVar_Value {
	if in == nil {
		return nil
	}
	out := &pb.EnvVar_Value{}
	out.Value = direct.ValueOf(in)
	return out
}

func TaskTemplate_MaxRetries_ToProto(mapCtx *direct.MapContext, in *int32) *pb.TaskTemplate_MaxRetries {
	if in == nil {
		return nil
	}
	out := &pb.TaskTemplate_MaxRetries{}
	out.MaxRetries = direct.ValueOf(in)
	return out
}

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
	out.TerminalCondition = Conditions_FromProto(mapCtx, in.GetTerminalCondition())
	// MISSING: Conditions
	out.ExecutionCount = direct.LazyPtr(in.GetExecutionCount())
	out.LatestCreatedExecution = ExecutionReferences_FromProto(mapCtx, in.GetLatestCreatedExecution())
	out.Reconciling = direct.LazyPtr(in.GetReconciling())
	// MISSING: SatisfiesPzs
	// MISSING: StartExecutionToken
	// MISSING: RunExecutionToken
	out.Etag = direct.LazyPtr(in.GetEtag())
	return out
}

func Conditions_FromProto(mapCtx *direct.MapContext, in *pb.Condition) []*krm.Condition {
	if in == nil {
		return nil
	}
	out := []*krm.Condition{}
	item := Condition_FromProto(mapCtx, in)
	if item != nil {
		out = append(out, item)
	}
	return out
}

func ExecutionReferences_FromProto(mapCtx *direct.MapContext, in *pb.ExecutionReference) []*krm.ExecutionReference {
	if in == nil {
		return nil
	}
	out := []*krm.ExecutionReference{}
	item := ExecutionReference_FromProto(mapCtx, in)
	if item != nil {
		out = append(out, item)
	}
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
	out.TerminalCondition = Conditions_ToProto(mapCtx, in.TerminalCondition)
	// MISSING: Conditions
	out.ExecutionCount = direct.ValueOf(in.ExecutionCount)
	out.LatestCreatedExecution = ExecutionReferences_ToProto(mapCtx, in.LatestCreatedExecution)
	out.Reconciling = direct.ValueOf(in.Reconciling)
	// MISSING: SatisfiesPzs
	// MISSING: StartExecutionToken
	// MISSING: RunExecutionToken
	out.Etag = direct.ValueOf(in.Etag)
	return out
}

func Conditions_ToProto(mapCtx *direct.MapContext, in []*krm.Condition) *pb.Condition {
	if in == nil {
		return nil
	}
	out := &pb.Condition{}
	if len(in) > 0 {
		item := Condition_ToProto(mapCtx, in[0])
		if item != nil {
			out = item
		}
	}
	return out
}

func ExecutionReferences_ToProto(mapCtx *direct.MapContext, in []*krm.ExecutionReference) *pb.ExecutionReference {
	if in == nil {
		return nil
	}
	out := &pb.ExecutionReference{}
	if len(in) > 0 {
		item := ExecutionReference_ToProto(mapCtx, in[0])
		if item != nil {
			out = item
		}
	}
	return out
}
func ConditionObservedState_Reason_ToProto(mapCtx *direct.MapContext, in *string) *pb.Condition_Reason {
	if in == nil {
		return nil
	}
	out := &pb.Condition_Reason{
		Reason: pb.Condition_CommonReason(direct.StringPtrToInt32(mapCtx, in)),
	}
	return out
}

func ConditionObservedState_RevisionReason_ToProto(mapCtx *direct.MapContext, in *string) *pb.Condition_RevisionReason_ {
	if in == nil {
		return nil
	}
	out := &pb.Condition_RevisionReason_{
		RevisionReason: pb.Condition_RevisionReason(direct.StringPtrToInt32(mapCtx, in)),
	}
	return out
}
func ConditionObservedState_ExecutionReason_ToProto(mapCtx *direct.MapContext, in *string) *pb.Condition_ExecutionReason_ {
	if in == nil {
		return nil
	}
	out := &pb.Condition_ExecutionReason_{
		ExecutionReason: pb.Condition_ExecutionReason(direct.StringPtrToInt32(mapCtx, in)),
	}
	return out
}
