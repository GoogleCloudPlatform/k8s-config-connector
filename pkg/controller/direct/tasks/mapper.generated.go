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
	pb "cloud.google.com/go/cloudtasks/apiv2beta3/cloudtaskspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/tasks/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func AppEngineHttpQueue_FromProto(mapCtx *direct.MapContext, in *pb.AppEngineHttpQueue) *krm.AppEngineHttpQueue {
	if in == nil {
		return nil
	}
	out := &krm.AppEngineHttpQueue{}
	out.AppEngineRoutingOverride = AppEngineRouting_FromProto(mapCtx, in.GetAppEngineRoutingOverride())
	return out
}
func AppEngineHttpQueue_ToProto(mapCtx *direct.MapContext, in *krm.AppEngineHttpQueue) *pb.AppEngineHttpQueue {
	if in == nil {
		return nil
	}
	out := &pb.AppEngineHttpQueue{}
	out.AppEngineRoutingOverride = AppEngineRouting_ToProto(mapCtx, in.AppEngineRoutingOverride)
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
	out.URIOverride = UriOverride_FromProto(mapCtx, in.GetUriOverride())
	out.HTTPMethod = direct.Enum_FromProto(mapCtx, in.GetHttpMethod())
	out.HeaderOverrides = direct.Slice_FromProto(mapCtx, in.HeaderOverrides, HttpTarget_HeaderOverride_FromProto)
	out.OauthToken = OAuthToken_FromProto(mapCtx, in.GetOauthToken())
	out.OidcToken = OidcToken_FromProto(mapCtx, in.GetOidcToken())
	return out
}
func HttpTarget_ToProto(mapCtx *direct.MapContext, in *krm.HttpTarget) *pb.HttpTarget {
	if in == nil {
		return nil
	}
	out := &pb.HttpTarget{}
	out.UriOverride = UriOverride_ToProto(mapCtx, in.URIOverride)
	out.HttpMethod = direct.Enum_ToProto[pb.HttpMethod](mapCtx, in.HTTPMethod)
	out.HeaderOverrides = direct.Slice_ToProto(mapCtx, in.HeaderOverrides, HttpTarget_HeaderOverride_ToProto)
	if oneof := OAuthToken_ToProto(mapCtx, in.OauthToken); oneof != nil {
		out.AuthorizationHeader = &pb.HttpTarget_OauthToken{OauthToken: oneof}
	}
	if oneof := OidcToken_ToProto(mapCtx, in.OidcToken); oneof != nil {
		out.AuthorizationHeader = &pb.HttpTarget_OidcToken{OidcToken: oneof}
	}
	return out
}
func HttpTarget_Header_FromProto(mapCtx *direct.MapContext, in *pb.HttpTarget_Header) *krm.HttpTarget_Header {
	if in == nil {
		return nil
	}
	out := &krm.HttpTarget_Header{}
	out.Key = direct.LazyPtr(in.GetKey())
	out.Value = direct.LazyPtr(in.GetValue())
	return out
}
func HttpTarget_Header_ToProto(mapCtx *direct.MapContext, in *krm.HttpTarget_Header) *pb.HttpTarget_Header {
	if in == nil {
		return nil
	}
	out := &pb.HttpTarget_Header{}
	out.Key = direct.ValueOf(in.Key)
	out.Value = direct.ValueOf(in.Value)
	return out
}
func HttpTarget_HeaderOverride_FromProto(mapCtx *direct.MapContext, in *pb.HttpTarget_HeaderOverride) *krm.HttpTarget_HeaderOverride {
	if in == nil {
		return nil
	}
	out := &krm.HttpTarget_HeaderOverride{}
	out.Header = HttpTarget_Header_FromProto(mapCtx, in.GetHeader())
	return out
}
func HttpTarget_HeaderOverride_ToProto(mapCtx *direct.MapContext, in *krm.HttpTarget_HeaderOverride) *pb.HttpTarget_HeaderOverride {
	if in == nil {
		return nil
	}
	out := &pb.HttpTarget_HeaderOverride{}
	out.Header = HttpTarget_Header_ToProto(mapCtx, in.Header)
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
func PathOverride_FromProto(mapCtx *direct.MapContext, in *pb.PathOverride) *krm.PathOverride {
	if in == nil {
		return nil
	}
	out := &krm.PathOverride{}
	out.Path = direct.LazyPtr(in.GetPath())
	return out
}
func PathOverride_ToProto(mapCtx *direct.MapContext, in *krm.PathOverride) *pb.PathOverride {
	if in == nil {
		return nil
	}
	out := &pb.PathOverride{}
	out.Path = direct.ValueOf(in.Path)
	return out
}
func QueryOverride_FromProto(mapCtx *direct.MapContext, in *pb.QueryOverride) *krm.QueryOverride {
	if in == nil {
		return nil
	}
	out := &krm.QueryOverride{}
	out.QueryParams = direct.LazyPtr(in.GetQueryParams())
	return out
}
func QueryOverride_ToProto(mapCtx *direct.MapContext, in *krm.QueryOverride) *pb.QueryOverride {
	if in == nil {
		return nil
	}
	out := &pb.QueryOverride{}
	out.QueryParams = direct.ValueOf(in.QueryParams)
	return out
}
func Queue_FromProto(mapCtx *direct.MapContext, in *pb.Queue) *krm.Queue {
	if in == nil {
		return nil
	}
	out := &krm.Queue{}
	out.Name = direct.LazyPtr(in.GetName())
	out.AppEngineHTTPQueue = AppEngineHttpQueue_FromProto(mapCtx, in.GetAppEngineHttpQueue())
	out.HTTPTarget = HttpTarget_FromProto(mapCtx, in.GetHttpTarget())
	out.RateLimits = RateLimits_FromProto(mapCtx, in.GetRateLimits())
	out.RetryConfig = RetryConfig_FromProto(mapCtx, in.GetRetryConfig())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.PurgeTime = direct.StringTimestamp_FromProto(mapCtx, in.GetPurgeTime())
	out.TaskTtl = direct.StringDuration_FromProto(mapCtx, in.GetTaskTtl())
	out.TombstoneTtl = direct.StringDuration_FromProto(mapCtx, in.GetTombstoneTtl())
	out.StackdriverLoggingConfig = StackdriverLoggingConfig_FromProto(mapCtx, in.GetStackdriverLoggingConfig())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	// MISSING: Stats
	return out
}
func Queue_ToProto(mapCtx *direct.MapContext, in *krm.Queue) *pb.Queue {
	if in == nil {
		return nil
	}
	out := &pb.Queue{}
	out.Name = direct.ValueOf(in.Name)
	if oneof := AppEngineHttpQueue_ToProto(mapCtx, in.AppEngineHTTPQueue); oneof != nil {
		out.QueueType = &pb.Queue_AppEngineHttpQueue{AppEngineHttpQueue: oneof}
	}
	out.HttpTarget = HttpTarget_ToProto(mapCtx, in.HTTPTarget)
	out.RateLimits = RateLimits_ToProto(mapCtx, in.RateLimits)
	out.RetryConfig = RetryConfig_ToProto(mapCtx, in.RetryConfig)
	out.State = direct.Enum_ToProto[pb.Queue_State](mapCtx, in.State)
	out.PurgeTime = direct.StringTimestamp_ToProto(mapCtx, in.PurgeTime)
	out.TaskTtl = direct.StringDuration_ToProto(mapCtx, in.TaskTtl)
	out.TombstoneTtl = direct.StringDuration_ToProto(mapCtx, in.TombstoneTtl)
	out.StackdriverLoggingConfig = StackdriverLoggingConfig_ToProto(mapCtx, in.StackdriverLoggingConfig)
	out.Type = direct.Enum_ToProto[pb.Queue_Type](mapCtx, in.Type)
	// MISSING: Stats
	return out
}
func QueueObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Queue) *krm.QueueObservedState {
	if in == nil {
		return nil
	}
	out := &krm.QueueObservedState{}
	// MISSING: Name
	// MISSING: AppEngineHTTPQueue
	// MISSING: HTTPTarget
	// MISSING: RateLimits
	// MISSING: RetryConfig
	// MISSING: State
	// MISSING: PurgeTime
	// MISSING: TaskTtl
	// MISSING: TombstoneTtl
	// MISSING: StackdriverLoggingConfig
	// MISSING: Type
	out.Stats = QueueStats_FromProto(mapCtx, in.GetStats())
	return out
}
func QueueObservedState_ToProto(mapCtx *direct.MapContext, in *krm.QueueObservedState) *pb.Queue {
	if in == nil {
		return nil
	}
	out := &pb.Queue{}
	// MISSING: Name
	// MISSING: AppEngineHTTPQueue
	// MISSING: HTTPTarget
	// MISSING: RateLimits
	// MISSING: RetryConfig
	// MISSING: State
	// MISSING: PurgeTime
	// MISSING: TaskTtl
	// MISSING: TombstoneTtl
	// MISSING: StackdriverLoggingConfig
	// MISSING: Type
	out.Stats = QueueStats_ToProto(mapCtx, in.Stats)
	return out
}
func QueueStats_FromProto(mapCtx *direct.MapContext, in *pb.QueueStats) *krm.QueueStats {
	if in == nil {
		return nil
	}
	out := &krm.QueueStats{}
	// MISSING: TasksCount
	// MISSING: OldestEstimatedArrivalTime
	// MISSING: ExecutedLastMinuteCount
	// MISSING: ConcurrentDispatchesCount
	// MISSING: EffectiveExecutionRate
	return out
}
func QueueStats_ToProto(mapCtx *direct.MapContext, in *krm.QueueStats) *pb.QueueStats {
	if in == nil {
		return nil
	}
	out := &pb.QueueStats{}
	// MISSING: TasksCount
	// MISSING: OldestEstimatedArrivalTime
	// MISSING: ExecutedLastMinuteCount
	// MISSING: ConcurrentDispatchesCount
	// MISSING: EffectiveExecutionRate
	return out
}
func QueueStatsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.QueueStats) *krm.QueueStatsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.QueueStatsObservedState{}
	out.TasksCount = direct.LazyPtr(in.GetTasksCount())
	out.OldestEstimatedArrivalTime = direct.StringTimestamp_FromProto(mapCtx, in.GetOldestEstimatedArrivalTime())
	out.ExecutedLastMinuteCount = direct.LazyPtr(in.GetExecutedLastMinuteCount())
	out.ConcurrentDispatchesCount = direct.LazyPtr(in.GetConcurrentDispatchesCount())
	out.EffectiveExecutionRate = direct.LazyPtr(in.GetEffectiveExecutionRate())
	return out
}
func QueueStatsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.QueueStatsObservedState) *pb.QueueStats {
	if in == nil {
		return nil
	}
	out := &pb.QueueStats{}
	out.TasksCount = direct.ValueOf(in.TasksCount)
	out.OldestEstimatedArrivalTime = direct.StringTimestamp_ToProto(mapCtx, in.OldestEstimatedArrivalTime)
	out.ExecutedLastMinuteCount = direct.ValueOf(in.ExecutedLastMinuteCount)
	out.ConcurrentDispatchesCount = direct.ValueOf(in.ConcurrentDispatchesCount)
	out.EffectiveExecutionRate = direct.ValueOf(in.EffectiveExecutionRate)
	return out
}
func RateLimits_FromProto(mapCtx *direct.MapContext, in *pb.RateLimits) *krm.RateLimits {
	if in == nil {
		return nil
	}
	out := &krm.RateLimits{}
	out.MaxDispatchesPerSecond = direct.LazyPtr(in.GetMaxDispatchesPerSecond())
	out.MaxBurstSize = direct.LazyPtr(in.GetMaxBurstSize())
	out.MaxConcurrentDispatches = direct.LazyPtr(in.GetMaxConcurrentDispatches())
	return out
}
func RateLimits_ToProto(mapCtx *direct.MapContext, in *krm.RateLimits) *pb.RateLimits {
	if in == nil {
		return nil
	}
	out := &pb.RateLimits{}
	out.MaxDispatchesPerSecond = direct.ValueOf(in.MaxDispatchesPerSecond)
	out.MaxBurstSize = direct.ValueOf(in.MaxBurstSize)
	out.MaxConcurrentDispatches = direct.ValueOf(in.MaxConcurrentDispatches)
	return out
}
func RetryConfig_FromProto(mapCtx *direct.MapContext, in *pb.RetryConfig) *krm.RetryConfig {
	if in == nil {
		return nil
	}
	out := &krm.RetryConfig{}
	out.MaxAttempts = direct.LazyPtr(in.GetMaxAttempts())
	out.MaxRetryDuration = direct.StringDuration_FromProto(mapCtx, in.GetMaxRetryDuration())
	out.MinBackoff = direct.StringDuration_FromProto(mapCtx, in.GetMinBackoff())
	out.MaxBackoff = direct.StringDuration_FromProto(mapCtx, in.GetMaxBackoff())
	out.MaxDoublings = direct.LazyPtr(in.GetMaxDoublings())
	return out
}
func RetryConfig_ToProto(mapCtx *direct.MapContext, in *krm.RetryConfig) *pb.RetryConfig {
	if in == nil {
		return nil
	}
	out := &pb.RetryConfig{}
	out.MaxAttempts = direct.ValueOf(in.MaxAttempts)
	out.MaxRetryDuration = direct.StringDuration_ToProto(mapCtx, in.MaxRetryDuration)
	out.MinBackoff = direct.StringDuration_ToProto(mapCtx, in.MinBackoff)
	out.MaxBackoff = direct.StringDuration_ToProto(mapCtx, in.MaxBackoff)
	out.MaxDoublings = direct.ValueOf(in.MaxDoublings)
	return out
}
func StackdriverLoggingConfig_FromProto(mapCtx *direct.MapContext, in *pb.StackdriverLoggingConfig) *krm.StackdriverLoggingConfig {
	if in == nil {
		return nil
	}
	out := &krm.StackdriverLoggingConfig{}
	out.SamplingRatio = direct.LazyPtr(in.GetSamplingRatio())
	return out
}
func StackdriverLoggingConfig_ToProto(mapCtx *direct.MapContext, in *krm.StackdriverLoggingConfig) *pb.StackdriverLoggingConfig {
	if in == nil {
		return nil
	}
	out := &pb.StackdriverLoggingConfig{}
	out.SamplingRatio = direct.ValueOf(in.SamplingRatio)
	return out
}
func UriOverride_FromProto(mapCtx *direct.MapContext, in *pb.UriOverride) *krm.UriOverride {
	if in == nil {
		return nil
	}
	out := &krm.UriOverride{}
	out.Scheme = direct.Enum_FromProto(mapCtx, in.GetScheme())
	out.Host = in.Host
	out.Port = in.Port
	out.PathOverride = PathOverride_FromProto(mapCtx, in.GetPathOverride())
	out.QueryOverride = QueryOverride_FromProto(mapCtx, in.GetQueryOverride())
	out.URIOverrideEnforceMode = direct.Enum_FromProto(mapCtx, in.GetUriOverrideEnforceMode())
	return out
}
func UriOverride_ToProto(mapCtx *direct.MapContext, in *krm.UriOverride) *pb.UriOverride {
	if in == nil {
		return nil
	}
	out := &pb.UriOverride{}
	if oneof := UriOverride_Scheme_ToProto(mapCtx, in.Scheme); oneof != nil {
		out.Scheme = oneof
	}
	out.Host = in.Host
	out.Port = in.Port
	out.PathOverride = PathOverride_ToProto(mapCtx, in.PathOverride)
	out.QueryOverride = QueryOverride_ToProto(mapCtx, in.QueryOverride)
	out.UriOverrideEnforceMode = direct.Enum_ToProto[pb.UriOverride_UriOverrideEnforceMode](mapCtx, in.URIOverrideEnforceMode)
	return out
}
