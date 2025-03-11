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
func TasksQueueObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Queue) *krm.TasksQueueObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TasksQueueObservedState{}
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.PurgeTime = direct.StringTimestamp_FromProto(mapCtx, in.GetPurgeTime())
	return out
}
func TasksQueueObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TasksQueueObservedState) *pb.Queue {
	if in == nil {
		return nil
	}
	out := &pb.Queue{}
	out.State = direct.Enum_ToProto[pb.Queue_State](mapCtx, in.State)
	out.PurgeTime = direct.StringTimestamp_ToProto(mapCtx, in.PurgeTime)
	return out
}
func TasksQueueSpec_FromProto(mapCtx *direct.MapContext, in *pb.Queue) *krm.TasksQueueSpec {
	if in == nil {
		return nil
	}
	out := &krm.TasksQueueSpec{}
	out.AppEngineRoutingOverride = AppEngineRouting_FromProto(mapCtx, in.GetAppEngineRoutingOverride())
	out.RateLimits = RateLimits_FromProto(mapCtx, in.GetRateLimits())
	out.RetryConfig = RetryConfig_FromProto(mapCtx, in.GetRetryConfig())
	out.StackdriverLoggingConfig = StackdriverLoggingConfig_FromProto(mapCtx, in.GetStackdriverLoggingConfig())
	return out
}
func TasksQueueSpec_ToProto(mapCtx *direct.MapContext, in *krm.TasksQueueSpec) *pb.Queue {
	if in == nil {
		return nil
	}
	out := &pb.Queue{}
	out.AppEngineRoutingOverride = AppEngineRouting_ToProto(mapCtx, in.AppEngineRoutingOverride)
	out.RateLimits = RateLimits_ToProto(mapCtx, in.RateLimits)
	out.RetryConfig = RetryConfig_ToProto(mapCtx, in.RetryConfig)
	out.StackdriverLoggingConfig = StackdriverLoggingConfig_ToProto(mapCtx, in.StackdriverLoggingConfig)
	return out
}
