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
	pb "cloud.google.com/go/cloudtasks/apiv2/cloudtaskspb"
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
func Attempt_FromProto(mapCtx *direct.MapContext, in *pb.Attempt) *krm.Attempt {
	if in == nil {
		return nil
	}
	out := &krm.Attempt{}
	out.ScheduleTime = direct.StringTimestamp_FromProto(mapCtx, in.GetScheduleTime())
	out.DispatchTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDispatchTime())
	out.ResponseTime = direct.StringTimestamp_FromProto(mapCtx, in.GetResponseTime())
	out.ResponseStatus = Status_FromProto(mapCtx, in.GetResponseStatus())
	return out
}
func Attempt_ToProto(mapCtx *direct.MapContext, in *krm.Attempt) *pb.Attempt {
	if in == nil {
		return nil
	}
	out := &pb.Attempt{}
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
func Task_FromProto(mapCtx *direct.MapContext, in *pb.Task) *krm.Task {
	if in == nil {
		return nil
	}
	out := &krm.Task{}
	out.Name = direct.LazyPtr(in.GetName())
	out.AppEngineHTTPRequest = AppEngineHttpRequest_FromProto(mapCtx, in.GetAppEngineHttpRequest())
	out.HTTPRequest = HttpRequest_FromProto(mapCtx, in.GetHttpRequest())
	out.ScheduleTime = direct.StringTimestamp_FromProto(mapCtx, in.GetScheduleTime())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.DispatchDeadline = direct.StringDuration_FromProto(mapCtx, in.GetDispatchDeadline())
	out.DispatchCount = direct.LazyPtr(in.GetDispatchCount())
	out.ResponseCount = direct.LazyPtr(in.GetResponseCount())
	out.FirstAttempt = Attempt_FromProto(mapCtx, in.GetFirstAttempt())
	out.LastAttempt = Attempt_FromProto(mapCtx, in.GetLastAttempt())
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
		out.MessageType = &pb.Task_AppEngineHttpRequest{AppEngineHttpRequest: oneof}
	}
	if oneof := HttpRequest_ToProto(mapCtx, in.HTTPRequest); oneof != nil {
		out.MessageType = &pb.Task_HttpRequest{HttpRequest: oneof}
	}
	out.ScheduleTime = direct.StringTimestamp_ToProto(mapCtx, in.ScheduleTime)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.DispatchDeadline = direct.StringDuration_ToProto(mapCtx, in.DispatchDeadline)
	out.DispatchCount = direct.ValueOf(in.DispatchCount)
	out.ResponseCount = direct.ValueOf(in.ResponseCount)
	out.FirstAttempt = Attempt_ToProto(mapCtx, in.FirstAttempt)
	out.LastAttempt = Attempt_ToProto(mapCtx, in.LastAttempt)
	out.View = direct.Enum_ToProto[pb.Task_View](mapCtx, in.View)
	return out
}
func TasksTaskObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Task) *krm.TasksTaskObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TasksTaskObservedState{}
	// MISSING: Name
	// MISSING: AppEngineHTTPRequest
	// MISSING: HTTPRequest
	// MISSING: ScheduleTime
	// MISSING: CreateTime
	// MISSING: DispatchDeadline
	// MISSING: DispatchCount
	// MISSING: ResponseCount
	// MISSING: FirstAttempt
	// MISSING: LastAttempt
	// MISSING: View
	return out
}
func TasksTaskObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TasksTaskObservedState) *pb.Task {
	if in == nil {
		return nil
	}
	out := &pb.Task{}
	// MISSING: Name
	// MISSING: AppEngineHTTPRequest
	// MISSING: HTTPRequest
	// MISSING: ScheduleTime
	// MISSING: CreateTime
	// MISSING: DispatchDeadline
	// MISSING: DispatchCount
	// MISSING: ResponseCount
	// MISSING: FirstAttempt
	// MISSING: LastAttempt
	// MISSING: View
	return out
}
func TasksTaskSpec_FromProto(mapCtx *direct.MapContext, in *pb.Task) *krm.TasksTaskSpec {
	if in == nil {
		return nil
	}
	out := &krm.TasksTaskSpec{}
	// MISSING: Name
	// MISSING: AppEngineHTTPRequest
	// MISSING: HTTPRequest
	// MISSING: ScheduleTime
	// MISSING: CreateTime
	// MISSING: DispatchDeadline
	// MISSING: DispatchCount
	// MISSING: ResponseCount
	// MISSING: FirstAttempt
	// MISSING: LastAttempt
	// MISSING: View
	return out
}
func TasksTaskSpec_ToProto(mapCtx *direct.MapContext, in *krm.TasksTaskSpec) *pb.Task {
	if in == nil {
		return nil
	}
	out := &pb.Task{}
	// MISSING: Name
	// MISSING: AppEngineHTTPRequest
	// MISSING: HTTPRequest
	// MISSING: ScheduleTime
	// MISSING: CreateTime
	// MISSING: DispatchDeadline
	// MISSING: DispatchCount
	// MISSING: ResponseCount
	// MISSING: FirstAttempt
	// MISSING: LastAttempt
	// MISSING: View
	return out
}
