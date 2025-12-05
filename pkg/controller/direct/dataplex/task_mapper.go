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
// krm.group: dataplex.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.dataplex.v1

package dataplex

import (
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	krmv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataplex/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DataplexTaskObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Task) *krmv1alpha1.DataplexTaskObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.DataplexTaskObservedState{}
	out.UID = direct.LazyPtr(in.GetUid())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.ExecutionStatus = Task_ExecutionStatusObservedState_FromProto(mapCtx, in.GetExecutionStatus())
	return out
}

func DataplexTaskObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.DataplexTaskObservedState) *pb.Task {
	if in == nil {
		return nil
	}
	out := &pb.Task{}
	out.Uid = direct.ValueOf(in.UID)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.State = direct.Enum_ToProto[pb.State](mapCtx, in.State)
	out.ExecutionStatus = Task_ExecutionStatusObservedState_ToProto(mapCtx, in.ExecutionStatus)
	return out
}
