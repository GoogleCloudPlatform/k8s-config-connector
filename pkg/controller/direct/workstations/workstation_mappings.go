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

package workstations

import (
	pb "cloud.google.com/go/workstations/apiv1/workstationspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/workstations/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func WorkstationObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Workstation) *krm.WorkstationObservedState {
	if in == nil {
		return nil
	}
	out := &krm.WorkstationObservedState{}
	out.UID = direct.LazyPtr(in.GetUid())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.DeleteTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDeleteTime())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Host = direct.LazyPtr(in.GetHost())
	return out
}

func WorkstationObservedState_ToProto(mapCtx *direct.MapContext, in *krm.WorkstationObservedState) *pb.Workstation {
	if in == nil {
		return nil
	}
	out := &pb.Workstation{}
	out.Uid = direct.ValueOf(in.UID)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.DeleteTime = direct.StringTimestamp_ToProto(mapCtx, in.DeleteTime)
	out.Etag = direct.ValueOf(in.Etag)
	out.State = direct.Enum_ToProto[pb.Workstation_State](mapCtx, in.State)
	out.Host = direct.ValueOf(in.Host)
	return out
}

func WorkstationSpec_FromProto(mapCtx *direct.MapContext, in *pb.Workstation) *krm.WorkstationSpec {
	if in == nil {
		return nil
	}
	out := &krm.WorkstationSpec{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Annotations = WorkstationAnnotations_FromProto_Alpha(mapCtx, in.Annotations)
	out.Labels = WorkstationLabels_FromProto_Alpha(mapCtx, in.Labels)
	return out
}

func WorkstationSpec_ToProto(mapCtx *direct.MapContext, in *krm.WorkstationSpec) *pb.Workstation {
	if in == nil {
		return nil
	}
	out := &pb.Workstation{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Annotations = WorkstationAnnotations_ToProto_Alpha(mapCtx, in.Annotations)
	out.Labels = WorkstationLabels_ToProto_Alpha(mapCtx, in.Labels)
	return out
}
