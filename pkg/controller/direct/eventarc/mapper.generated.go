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

package eventarc

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/eventarc/apiv1/eventarcpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/eventarc/v1alpha1"
)
func EventarcMessageBusObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MessageBus) *krm.EventarcMessageBusObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EventarcMessageBusObservedState{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Etag
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Annotations
	// MISSING: DisplayName
	// MISSING: CryptoKeyName
	// MISSING: LoggingConfig
	return out
}
func EventarcMessageBusObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EventarcMessageBusObservedState) *pb.MessageBus {
	if in == nil {
		return nil
	}
	out := &pb.MessageBus{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Etag
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Annotations
	// MISSING: DisplayName
	// MISSING: CryptoKeyName
	// MISSING: LoggingConfig
	return out
}
func EventarcMessageBusSpec_FromProto(mapCtx *direct.MapContext, in *pb.MessageBus) *krm.EventarcMessageBusSpec {
	if in == nil {
		return nil
	}
	out := &krm.EventarcMessageBusSpec{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Etag
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Annotations
	// MISSING: DisplayName
	// MISSING: CryptoKeyName
	// MISSING: LoggingConfig
	return out
}
func EventarcMessageBusSpec_ToProto(mapCtx *direct.MapContext, in *krm.EventarcMessageBusSpec) *pb.MessageBus {
	if in == nil {
		return nil
	}
	out := &pb.MessageBus{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Etag
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Annotations
	// MISSING: DisplayName
	// MISSING: CryptoKeyName
	// MISSING: LoggingConfig
	return out
}
func LoggingConfig_FromProto(mapCtx *direct.MapContext, in *pb.LoggingConfig) *krm.LoggingConfig {
	if in == nil {
		return nil
	}
	out := &krm.LoggingConfig{}
	out.LogSeverity = direct.Enum_FromProto(mapCtx, in.GetLogSeverity())
	return out
}
func LoggingConfig_ToProto(mapCtx *direct.MapContext, in *krm.LoggingConfig) *pb.LoggingConfig {
	if in == nil {
		return nil
	}
	out := &pb.LoggingConfig{}
	out.LogSeverity = direct.Enum_ToProto[pb.LoggingConfig_LogSeverity](mapCtx, in.LogSeverity)
	return out
}
func MessageBus_FromProto(mapCtx *direct.MapContext, in *pb.MessageBus) *krm.MessageBus {
	if in == nil {
		return nil
	}
	out := &krm.MessageBus{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: Uid
	// MISSING: Etag
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Annotations = in.Annotations
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.CryptoKeyName = direct.LazyPtr(in.GetCryptoKeyName())
	out.LoggingConfig = LoggingConfig_FromProto(mapCtx, in.GetLoggingConfig())
	return out
}
func MessageBus_ToProto(mapCtx *direct.MapContext, in *krm.MessageBus) *pb.MessageBus {
	if in == nil {
		return nil
	}
	out := &pb.MessageBus{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: Uid
	// MISSING: Etag
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Annotations = in.Annotations
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.CryptoKeyName = direct.ValueOf(in.CryptoKeyName)
	out.LoggingConfig = LoggingConfig_ToProto(mapCtx, in.LoggingConfig)
	return out
}
func MessageBusObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MessageBus) *krm.MessageBusObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MessageBusObservedState{}
	// MISSING: Name
	out.Uid = direct.LazyPtr(in.GetUid())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: Annotations
	// MISSING: DisplayName
	// MISSING: CryptoKeyName
	// MISSING: LoggingConfig
	return out
}
func MessageBusObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MessageBusObservedState) *pb.MessageBus {
	if in == nil {
		return nil
	}
	out := &pb.MessageBus{}
	// MISSING: Name
	out.Uid = direct.ValueOf(in.Uid)
	out.Etag = direct.ValueOf(in.Etag)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: Annotations
	// MISSING: DisplayName
	// MISSING: CryptoKeyName
	// MISSING: LoggingConfig
	return out
}
