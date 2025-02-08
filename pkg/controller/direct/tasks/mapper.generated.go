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

package tasks

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/cloudtasks/apiv2beta2/cloudtaskspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/tasks/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func AppEngineRouting_FromProto(mapCtx *direct.MapContext, in *pb.AppEngineRouting) *krm.AppEngineRouting {
	if in == nil {
		return nil
	}
	out := &krm.AppEngineRouting{}
	out.Service = direct.LazyPtr(in.GetService())
	out.Version = direct.LazyPtr(in.GetVersion())
	out.Instance = direct.LazyPtr(in.GetInstance())
	out.Host = direct.LazyPtr(in.GetHost())
	return out
}
func AppEngineRouting_ToProto(mapCtx *direct.MapContext, in *krm.AppEngineRouting) *pb.AppEngineRouting {
	if in == nil {
		return nil
	}
	out := &pb.AppEngineRouting{}
	out.Service = direct.ValueOf(in.Service)
	out.Version = direct.ValueOf(in.Version)
	out.Instance = direct.ValueOf(in.Instance)
	out.Host = direct.ValueOf(in.Host)
	return out
}
func AttemptStatus_FromProto(mapCtx *direct.MapContext, in *pb.AttemptStatus) *krm.AttemptStatus {
	if in == nil {
		return nil
	}
	out := &krm.AttemptStatus{}
	out.ScheduleTime = direct.StringTimestamp_FromProto(mapCtx, in.GetScheduleTime())
	out.DispatchTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDispatchTime())
	out.ResponseTime = direct.StringTimestamp_FromProto(mapCtx, in.GetResponseTime())
	out.ResponseStatus = Status_FromProto(mapCtx, in.GetResponseStatus())
	return out
}
func AttemptStatus_ToProto(mapCtx *direct.MapContext, in *krm.AttemptStatus) *pb.AttemptStatus {
	if in == nil {
		return nil
	}
	out := &pb.AttemptStatus{}
	out.ScheduleTime = direct.StringTimestamp_ToProto(mapCtx, in.ScheduleTime)
	out.DispatchTime = direct.StringTimestamp_ToProto(mapCtx, in.DispatchTime)
	out.ResponseTime = direct.StringTimestamp_ToProto(mapCtx, in.ResponseTime)
	out.ResponseStatus = Status_ToProto(mapCtx, in.ResponseStatus)
	return out
}
func OAuthToken_FromProto(mapCtx *direct.MapContext, in *pb.OAuthToken) *krm.OAuthToken {
	if in == nil {
		return nil
	}
	out := &krm.OAuthToken{}
	out.ServiceAccountEmail = direct.LazyPtr(in.GetServiceAccountEmail())
	out.Scope = direct.LazyPtr(in.GetScope())
	return out
}
func OAuthToken_ToProto(mapCtx *direct.MapContext, in *krm.OAuthToken) *pb.OAuthToken {
	if in == nil {
		return nil
	}
	out := &pb.OAuthToken{}
	out.ServiceAccountEmail = direct.ValueOf(in.ServiceAccountEmail)
	out.Scope = direct.ValueOf(in.Scope)
	return out
}
func OidcToken_FromProto(mapCtx *direct.MapContext, in *pb.OidcToken) *krm.OidcToken {
	if in == nil {
		return nil
	}
	out := &krm.OidcToken{}
	out.ServiceAccountEmail = direct.LazyPtr(in.GetServiceAccountEmail())
	out.Audience = direct.LazyPtr(in.GetAudience())
	return out
}
func OidcToken_ToProto(mapCtx *direct.MapContext, in *krm.OidcToken) *pb.OidcToken {
	if in == nil {
		return nil
	}
	out := &pb.OidcToken{}
	out.ServiceAccountEmail = direct.ValueOf(in.ServiceAccountEmail)
	out.Audience = direct.ValueOf(in.Audience)
	return out
}
func PullMessage_FromProto(mapCtx *direct.MapContext, in *pb.PullMessage) *krm.PullMessage {
	if in == nil {
		return nil
	}
	out := &krm.PullMessage{}
	out.Payload = in.GetPayload()
	out.Tag = direct.LazyPtr(in.GetTag())
	return out
}
func PullMessage_ToProto(mapCtx *direct.MapContext, in *krm.PullMessage) *pb.PullMessage {
	if in == nil {
		return nil
	}
	out := &pb.PullMessage{}
	out.Payload = in.Payload
	out.Tag = direct.ValueOf(in.Tag)
	return out
}
func Task_FromProto(mapCtx *direct.MapContext, in *pb.Task) *krm.Task {
	if in == nil {
		return nil
	}
	out := &krm.Task{}
	out.Name = direct.LazyPtr(in.GetName())
	out.AppEngineHTTPRequest = AppEngineHttpRequest_FromProto(mapCtx, in.GetAppEngineHttpRequest())
	out.PullMessage = PullMessage_FromProto(mapCtx, in.GetPullMessage())
	out.HTTPRequest = HttpRequest_FromProto(mapCtx, in.GetHttpRequest())
	out.ScheduleTime = direct.StringTimestamp_FromProto(mapCtx, in.GetScheduleTime())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.Status = TaskStatus_FromProto(mapCtx, in.GetStatus())
	out.View = direct.Enum_FromProto(mapCtx, in.GetView())
	return out
}
func Task_ToProto(mapCtx *direct.MapContext, in *krm.Task) *pb.Task {
	if in == nil {
		return nil
	}
	out := &pb.Task{}
	out.Name = direct.ValueOf(in.Name)
	if oneof := AppEngineHttpRequest_ToProto(mapCtx, in.AppEngineHTTPRequest); oneof != nil {
		out.PayloadType = &pb.Task_AppEngineHttpRequest{AppEngineHttpRequest: oneof}
	}
	if oneof := PullMessage_ToProto(mapCtx, in.PullMessage); oneof != nil {
		out.PayloadType = &pb.Task_PullMessage{PullMessage: oneof}
	}
	if oneof := HttpRequest_ToProto(mapCtx, in.HTTPRequest); oneof != nil {
		out.PayloadType = &pb.Task_HttpRequest{HttpRequest: oneof}
	}
	out.ScheduleTime = direct.StringTimestamp_ToProto(mapCtx, in.ScheduleTime)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.Status = TaskStatus_ToProto(mapCtx, in.Status)
	out.View = direct.Enum_ToProto[pb.Task_View](mapCtx, in.View)
	return out
}
func TaskStatus_FromProto(mapCtx *direct.MapContext, in *pb.TaskStatus) *krm.TaskStatus {
	if in == nil {
		return nil
	}
	out := &krm.TaskStatus{}
	out.AttemptDispatchCount = direct.LazyPtr(in.GetAttemptDispatchCount())
	out.AttemptResponseCount = direct.LazyPtr(in.GetAttemptResponseCount())
	out.FirstAttemptStatus = AttemptStatus_FromProto(mapCtx, in.GetFirstAttemptStatus())
	out.LastAttemptStatus = AttemptStatus_FromProto(mapCtx, in.GetLastAttemptStatus())
	return out
}
func TaskStatus_ToProto(mapCtx *direct.MapContext, in *krm.TaskStatus) *pb.TaskStatus {
	if in == nil {
		return nil
	}
	out := &pb.TaskStatus{}
	out.AttemptDispatchCount = direct.ValueOf(in.AttemptDispatchCount)
	out.AttemptResponseCount = direct.ValueOf(in.AttemptResponseCount)
	out.FirstAttemptStatus = AttemptStatus_ToProto(mapCtx, in.FirstAttemptStatus)
	out.LastAttemptStatus = AttemptStatus_ToProto(mapCtx, in.LastAttemptStatus)
	return out
}
