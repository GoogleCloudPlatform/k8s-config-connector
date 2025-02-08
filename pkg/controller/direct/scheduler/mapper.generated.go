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

package scheduler

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/scheduler/apiv1/schedulerpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/scheduler/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func AppEngineHttpTarget_FromProto(mapCtx *direct.MapContext, in *pb.AppEngineHttpTarget) *krm.AppEngineHttpTarget {
	if in == nil {
		return nil
	}
	out := &krm.AppEngineHttpTarget{}
	out.HTTPMethod = direct.Enum_FromProto(mapCtx, in.GetHttpMethod())
	out.AppEngineRouting = AppEngineRouting_FromProto(mapCtx, in.GetAppEngineRouting())
	out.RelativeURI = direct.LazyPtr(in.GetRelativeUri())
	out.Headers = in.Headers
	out.Body = in.GetBody()
	return out
}
func AppEngineHttpTarget_ToProto(mapCtx *direct.MapContext, in *krm.AppEngineHttpTarget) *pb.AppEngineHttpTarget {
	if in == nil {
		return nil
	}
	out := &pb.AppEngineHttpTarget{}
	out.HttpMethod = direct.Enum_ToProto[pb.HttpMethod](mapCtx, in.HTTPMethod)
	out.AppEngineRouting = AppEngineRouting_ToProto(mapCtx, in.AppEngineRouting)
	out.RelativeUri = direct.ValueOf(in.RelativeURI)
	out.Headers = in.Headers
	out.Body = in.Body
	return out
}
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
func HttpTarget_FromProto(mapCtx *direct.MapContext, in *pb.HttpTarget) *krm.HttpTarget {
	if in == nil {
		return nil
	}
	out := &krm.HttpTarget{}
	out.URI = direct.LazyPtr(in.GetUri())
	out.HTTPMethod = direct.Enum_FromProto(mapCtx, in.GetHttpMethod())
	out.Headers = in.Headers
	out.Body = in.GetBody()
	out.OauthToken = OAuthToken_FromProto(mapCtx, in.GetOauthToken())
	out.OidcToken = OidcToken_FromProto(mapCtx, in.GetOidcToken())
	return out
}
func HttpTarget_ToProto(mapCtx *direct.MapContext, in *krm.HttpTarget) *pb.HttpTarget {
	if in == nil {
		return nil
	}
	out := &pb.HttpTarget{}
	out.Uri = direct.ValueOf(in.URI)
	out.HttpMethod = direct.Enum_ToProto[pb.HttpMethod](mapCtx, in.HTTPMethod)
	out.Headers = in.Headers
	out.Body = in.Body
	if oneof := OAuthToken_ToProto(mapCtx, in.OauthToken); oneof != nil {
		out.AuthorizationHeader = &pb.HttpTarget_OauthToken{OauthToken: oneof}
	}
	if oneof := OidcToken_ToProto(mapCtx, in.OidcToken); oneof != nil {
		out.AuthorizationHeader = &pb.HttpTarget_OidcToken{OidcToken: oneof}
	}
	return out
}
func Job_FromProto(mapCtx *direct.MapContext, in *pb.Job) *krm.Job {
	if in == nil {
		return nil
	}
	out := &krm.Job{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.PubsubTarget = PubsubTarget_FromProto(mapCtx, in.GetPubsubTarget())
	out.AppEngineHTTPTarget = AppEngineHttpTarget_FromProto(mapCtx, in.GetAppEngineHttpTarget())
	out.HTTPTarget = HttpTarget_FromProto(mapCtx, in.GetHttpTarget())
	out.Schedule = direct.LazyPtr(in.GetSchedule())
	out.TimeZone = direct.LazyPtr(in.GetTimeZone())
	out.UserUpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUserUpdateTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Status = Status_FromProto(mapCtx, in.GetStatus())
	out.ScheduleTime = direct.StringTimestamp_FromProto(mapCtx, in.GetScheduleTime())
	out.LastAttemptTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLastAttemptTime())
	out.RetryConfig = RetryConfig_FromProto(mapCtx, in.GetRetryConfig())
	out.AttemptDeadline = direct.StringDuration_FromProto(mapCtx, in.GetAttemptDeadline())
	return out
}
func Job_ToProto(mapCtx *direct.MapContext, in *krm.Job) *pb.Job {
	if in == nil {
		return nil
	}
	out := &pb.Job{}
	out.Name = direct.ValueOf(in.Name)
	out.Description = direct.ValueOf(in.Description)
	if oneof := PubsubTarget_ToProto(mapCtx, in.PubsubTarget); oneof != nil {
		out.Target = &pb.Job_PubsubTarget{PubsubTarget: oneof}
	}
	if oneof := AppEngineHttpTarget_ToProto(mapCtx, in.AppEngineHTTPTarget); oneof != nil {
		out.Target = &pb.Job_AppEngineHttpTarget{AppEngineHttpTarget: oneof}
	}
	if oneof := HttpTarget_ToProto(mapCtx, in.HTTPTarget); oneof != nil {
		out.Target = &pb.Job_HttpTarget{HttpTarget: oneof}
	}
	out.Schedule = direct.ValueOf(in.Schedule)
	out.TimeZone = direct.ValueOf(in.TimeZone)
	out.UserUpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UserUpdateTime)
	out.State = direct.Enum_ToProto[pb.Job_State](mapCtx, in.State)
	out.Status = Status_ToProto(mapCtx, in.Status)
	out.ScheduleTime = direct.StringTimestamp_ToProto(mapCtx, in.ScheduleTime)
	out.LastAttemptTime = direct.StringTimestamp_ToProto(mapCtx, in.LastAttemptTime)
	out.RetryConfig = RetryConfig_ToProto(mapCtx, in.RetryConfig)
	out.AttemptDeadline = direct.StringDuration_ToProto(mapCtx, in.AttemptDeadline)
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
func PubsubTarget_FromProto(mapCtx *direct.MapContext, in *pb.PubsubTarget) *krm.PubsubTarget {
	if in == nil {
		return nil
	}
	out := &krm.PubsubTarget{}
	out.TopicName = direct.LazyPtr(in.GetTopicName())
	out.Data = in.GetData()
	out.Attributes = in.Attributes
	return out
}
func PubsubTarget_ToProto(mapCtx *direct.MapContext, in *krm.PubsubTarget) *pb.PubsubTarget {
	if in == nil {
		return nil
	}
	out := &pb.PubsubTarget{}
	out.TopicName = direct.ValueOf(in.TopicName)
	out.Data = in.Data
	out.Attributes = in.Attributes
	return out
}
func RetryConfig_FromProto(mapCtx *direct.MapContext, in *pb.RetryConfig) *krm.RetryConfig {
	if in == nil {
		return nil
	}
	out := &krm.RetryConfig{}
	out.RetryCount = direct.LazyPtr(in.GetRetryCount())
	out.MaxRetryDuration = direct.StringDuration_FromProto(mapCtx, in.GetMaxRetryDuration())
	out.MinBackoffDuration = direct.StringDuration_FromProto(mapCtx, in.GetMinBackoffDuration())
	out.MaxBackoffDuration = direct.StringDuration_FromProto(mapCtx, in.GetMaxBackoffDuration())
	out.MaxDoublings = direct.LazyPtr(in.GetMaxDoublings())
	return out
}
func RetryConfig_ToProto(mapCtx *direct.MapContext, in *krm.RetryConfig) *pb.RetryConfig {
	if in == nil {
		return nil
	}
	out := &pb.RetryConfig{}
	out.RetryCount = direct.ValueOf(in.RetryCount)
	out.MaxRetryDuration = direct.StringDuration_ToProto(mapCtx, in.MaxRetryDuration)
	out.MinBackoffDuration = direct.StringDuration_ToProto(mapCtx, in.MinBackoffDuration)
	out.MaxBackoffDuration = direct.StringDuration_ToProto(mapCtx, in.MaxBackoffDuration)
	out.MaxDoublings = direct.ValueOf(in.MaxDoublings)
	return out
}
func SchedulerJobObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Job) *krm.SchedulerJobObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SchedulerJobObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: PubsubTarget
	// MISSING: AppEngineHTTPTarget
	// MISSING: HTTPTarget
	// MISSING: Schedule
	// MISSING: TimeZone
	// MISSING: UserUpdateTime
	// MISSING: State
	// MISSING: Status
	// MISSING: ScheduleTime
	// MISSING: LastAttemptTime
	// MISSING: RetryConfig
	// MISSING: AttemptDeadline
	return out
}
func SchedulerJobObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SchedulerJobObservedState) *pb.Job {
	if in == nil {
		return nil
	}
	out := &pb.Job{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: PubsubTarget
	// MISSING: AppEngineHTTPTarget
	// MISSING: HTTPTarget
	// MISSING: Schedule
	// MISSING: TimeZone
	// MISSING: UserUpdateTime
	// MISSING: State
	// MISSING: Status
	// MISSING: ScheduleTime
	// MISSING: LastAttemptTime
	// MISSING: RetryConfig
	// MISSING: AttemptDeadline
	return out
}
func SchedulerJobSpec_FromProto(mapCtx *direct.MapContext, in *pb.Job) *krm.SchedulerJobSpec {
	if in == nil {
		return nil
	}
	out := &krm.SchedulerJobSpec{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: PubsubTarget
	// MISSING: AppEngineHTTPTarget
	// MISSING: HTTPTarget
	// MISSING: Schedule
	// MISSING: TimeZone
	// MISSING: UserUpdateTime
	// MISSING: State
	// MISSING: Status
	// MISSING: ScheduleTime
	// MISSING: LastAttemptTime
	// MISSING: RetryConfig
	// MISSING: AttemptDeadline
	return out
}
func SchedulerJobSpec_ToProto(mapCtx *direct.MapContext, in *krm.SchedulerJobSpec) *pb.Job {
	if in == nil {
		return nil
	}
	out := &pb.Job{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: PubsubTarget
	// MISSING: AppEngineHTTPTarget
	// MISSING: HTTPTarget
	// MISSING: Schedule
	// MISSING: TimeZone
	// MISSING: UserUpdateTime
	// MISSING: State
	// MISSING: Status
	// MISSING: ScheduleTime
	// MISSING: LastAttemptTime
	// MISSING: RetryConfig
	// MISSING: AttemptDeadline
	return out
}
