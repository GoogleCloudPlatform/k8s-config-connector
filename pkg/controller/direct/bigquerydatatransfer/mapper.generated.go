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
