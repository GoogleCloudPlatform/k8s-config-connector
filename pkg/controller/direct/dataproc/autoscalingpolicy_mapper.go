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

func DataprocAutoscalingPolicyStatus_FromProto(mapCtx *direct.MapContext, in *pb.AutoscalingPolicy) *krm.DataprocAutoscalingPolicyStatus {
	if in == nil {
		return nil
	}
	out := &krm.DataprocAutoscalingPolicyStatus{}
	out.ObservedState = DataprocAutoscalingPolicyObservedState_FromProto(mapCtx, in)
	return out
}
func DataprocAutoscalingPolicyStatus_ToProto(mapCtx *direct.MapContext, in *krm.DataprocAutoscalingPolicyStatus) *pb.AutoscalingPolicy {
	if in == nil {
		return nil
	}
	out := &pb.AutoscalingPolicy{}
	out = DataprocAutoscalingPolicyObservedState_ToProto(mapCtx, in.ObservedState)
	return out
}
func DataprocAutoscalingPolicyObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AutoscalingPolicy) *krm.DataprocAutoscalingPolicyObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataprocAutoscalingPolicyObservedState{}
	out.Name = direct.LazyPtr(in.Name)
	return out
}
func DataprocAutoscalingPolicyObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataprocAutoscalingPolicyObservedState) *pb.AutoscalingPolicy {
	if in == nil {
		return nil
	}
	out := &pb.AutoscalingPolicy{}
	out.Name = direct.ValueOf(in.Name)
	return out
}
func DataprocAutoscalingPolicySpec_FromProto(mapCtx *direct.MapContext, in *pb.AutoscalingPolicy) *krm.DataprocAutoscalingPolicySpec {
	if in == nil {
		return nil
	}
	out := &krm.DataprocAutoscalingPolicySpec{}
	out.ID = direct.LazyPtr(in.GetId())
	out.BasicAlgorithm = BasicAutoscalingAlgorithm_FromProto(mapCtx, in.GetBasicAlgorithm())
	out.WorkerConfig = InstanceGroupAutoscalingPolicyConfig_FromProto(mapCtx, in.GetWorkerConfig())
	out.SecondaryWorkerConfig = InstanceGroupAutoscalingPolicyConfig_FromProto(mapCtx, in.GetSecondaryWorkerConfig())
	out.Labels = in.Labels
	return out
}
func DataprocAutoscalingPolicySpec_ToProto(mapCtx *direct.MapContext, in *krm.DataprocAutoscalingPolicySpec) *pb.AutoscalingPolicy {
	if in == nil {
		return nil
	}
	out := &pb.AutoscalingPolicy{}
	out.Id = direct.ValueOf(in.ID)
	if oneof := BasicAutoscalingAlgorithm_ToProto(mapCtx, in.BasicAlgorithm); oneof != nil {
		out.Algorithm = &pb.AutoscalingPolicy_BasicAlgorithm{BasicAlgorithm: oneof}
	}
	out.WorkerConfig = InstanceGroupAutoscalingPolicyConfig_ToProto(mapCtx, in.WorkerConfig)
	out.SecondaryWorkerConfig = InstanceGroupAutoscalingPolicyConfig_ToProto(mapCtx, in.SecondaryWorkerConfig)
	out.Labels = in.Labels
	return out
}
