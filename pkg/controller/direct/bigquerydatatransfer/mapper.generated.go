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
// krm.group: bigquerydatatransfer.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.cloud.bigquery.datatransfer.v1

package bigquerydatatransfer

import (
	pb "cloud.google.com/go/bigquery/datatransfer/apiv1/datatransferpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquerydatatransfer/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func EmailPreferences_FromProto(mapCtx *direct.MapContext, in *pb.EmailPreferences) *krm.EmailPreferences {
	if in == nil {
		return nil
	}
	out := &krm.EmailPreferences{}
	out.EnableFailureEmail = direct.LazyPtr(in.GetEnableFailureEmail())
	return out
}
func EmailPreferences_ToProto(mapCtx *direct.MapContext, in *krm.EmailPreferences) *pb.EmailPreferences {
	if in == nil {
		return nil
	}
	out := &pb.EmailPreferences{}
	out.EnableFailureEmail = direct.ValueOf(in.EnableFailureEmail)
	return out
}
func ManualSchedule_FromProto(mapCtx *direct.MapContext, in *pb.ManualSchedule) *krm.ManualSchedule {
	if in == nil {
		return nil
	}
	out := &krm.ManualSchedule{}
	return out
}
func ManualSchedule_ToProto(mapCtx *direct.MapContext, in *krm.ManualSchedule) *pb.ManualSchedule {
	if in == nil {
		return nil
	}
	out := &pb.ManualSchedule{}
	return out
}
func ScheduleOptions_FromProto(mapCtx *direct.MapContext, in *pb.ScheduleOptions) *krm.ScheduleOptions {
	if in == nil {
		return nil
	}
	out := &krm.ScheduleOptions{}
	out.DisableAutoScheduling = direct.LazyPtr(in.GetDisableAutoScheduling())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	return out
}
func ScheduleOptions_ToProto(mapCtx *direct.MapContext, in *krm.ScheduleOptions) *pb.ScheduleOptions {
	if in == nil {
		return nil
	}
	out := &pb.ScheduleOptions{}
	out.DisableAutoScheduling = direct.ValueOf(in.DisableAutoScheduling)
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	return out
}
func ScheduleOptionsV2_FromProto(mapCtx *direct.MapContext, in *pb.ScheduleOptionsV2) *krm.ScheduleOptionsV2 {
	if in == nil {
		return nil
	}
	out := &krm.ScheduleOptionsV2{}
	out.TimeBasedSchedule = TimeBasedSchedule_FromProto(mapCtx, in.GetTimeBasedSchedule())
	out.ManualSchedule = ManualSchedule_FromProto(mapCtx, in.GetManualSchedule())
	out.EventDrivenSchedule = EventDrivenSchedule_FromProto(mapCtx, in.GetEventDrivenSchedule())
	return out
}
func ScheduleOptionsV2_ToProto(mapCtx *direct.MapContext, in *krm.ScheduleOptionsV2) *pb.ScheduleOptionsV2 {
	if in == nil {
		return nil
	}
	out := &pb.ScheduleOptionsV2{}
	if oneof := TimeBasedSchedule_ToProto(mapCtx, in.TimeBasedSchedule); oneof != nil {
		out.Schedule = &pb.ScheduleOptionsV2_TimeBasedSchedule{TimeBasedSchedule: oneof}
	}
	if oneof := ManualSchedule_ToProto(mapCtx, in.ManualSchedule); oneof != nil {
		out.Schedule = &pb.ScheduleOptionsV2_ManualSchedule{ManualSchedule: oneof}
	}
	if oneof := EventDrivenSchedule_ToProto(mapCtx, in.EventDrivenSchedule); oneof != nil {
		out.Schedule = &pb.ScheduleOptionsV2_EventDrivenSchedule{EventDrivenSchedule: oneof}
	}
	return out
}
func TimeBasedSchedule_FromProto(mapCtx *direct.MapContext, in *pb.TimeBasedSchedule) *krm.TimeBasedSchedule {
	if in == nil {
		return nil
	}
	out := &krm.TimeBasedSchedule{}
	out.Schedule = direct.LazyPtr(in.GetSchedule())
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	return out
}
func TimeBasedSchedule_ToProto(mapCtx *direct.MapContext, in *krm.TimeBasedSchedule) *pb.TimeBasedSchedule {
	if in == nil {
		return nil
	}
	out := &pb.TimeBasedSchedule{}
	out.Schedule = direct.ValueOf(in.Schedule)
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	return out
}
func UserInfo_FromProto(mapCtx *direct.MapContext, in *pb.UserInfo) *krm.UserInfo {
	if in == nil {
		return nil
	}
	out := &krm.UserInfo{}
	out.Email = in.Email
	return out
}
func UserInfo_ToProto(mapCtx *direct.MapContext, in *krm.UserInfo) *pb.UserInfo {
	if in == nil {
		return nil
	}
	out := &pb.UserInfo{}
	out.Email = in.Email
	return out
}
