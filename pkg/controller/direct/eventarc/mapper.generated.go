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
	pb "cloud.google.com/go/eventarc/apiv1/eventarcpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/eventarc/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func EventarcGoogleApiSourceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.GoogleApiSource) *krm.EventarcGoogleApiSourceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EventarcGoogleApiSourceObservedState{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Etag
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Annotations
	// MISSING: DisplayName
	// MISSING: Destination
	// MISSING: CryptoKeyName
	// MISSING: LoggingConfig
	return out
}
func EventarcGoogleApiSourceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EventarcGoogleApiSourceObservedState) *pb.GoogleApiSource {
	if in == nil {
		return nil
	}
	out := &pb.GoogleApiSource{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Etag
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Annotations
	// MISSING: DisplayName
	// MISSING: Destination
	// MISSING: CryptoKeyName
	// MISSING: LoggingConfig
	return out
}
func EventarcGoogleApiSourceSpec_FromProto(mapCtx *direct.MapContext, in *pb.GoogleApiSource) *krm.EventarcGoogleApiSourceSpec {
	if in == nil {
		return nil
	}
	out := &krm.EventarcGoogleApiSourceSpec{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Etag
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Annotations
	// MISSING: DisplayName
	// MISSING: Destination
	// MISSING: CryptoKeyName
	// MISSING: LoggingConfig
	return out
}
func EventarcGoogleApiSourceSpec_ToProto(mapCtx *direct.MapContext, in *krm.EventarcGoogleApiSourceSpec) *pb.GoogleApiSource {
	if in == nil {
		return nil
	}
	out := &pb.GoogleApiSource{}
	// MISSING: Name
	// MISSING: Uid
	// MISSING: Etag
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Annotations
	// MISSING: DisplayName
	// MISSING: Destination
	// MISSING: CryptoKeyName
	// MISSING: LoggingConfig
	return out
}
func GoogleApiSource_FromProto(mapCtx *direct.MapContext, in *pb.GoogleApiSource) *krm.GoogleApiSource {
	if in == nil {
		return nil
	}
	out := &krm.GoogleApiSource{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: Uid
	// MISSING: Etag
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Annotations = in.Annotations
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Destination = direct.LazyPtr(in.GetDestination())
	out.CryptoKeyName = direct.LazyPtr(in.GetCryptoKeyName())
	out.LoggingConfig = LoggingConfig_FromProto(mapCtx, in.GetLoggingConfig())
	return out
}
func GoogleApiSource_ToProto(mapCtx *direct.MapContext, in *krm.GoogleApiSource) *pb.GoogleApiSource {
	if in == nil {
		return nil
	}
	out := &pb.GoogleApiSource{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: Uid
	// MISSING: Etag
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Annotations = in.Annotations
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Destination = direct.ValueOf(in.Destination)
	out.CryptoKeyName = direct.ValueOf(in.CryptoKeyName)
	out.LoggingConfig = LoggingConfig_ToProto(mapCtx, in.LoggingConfig)
	return out
}
func GoogleApiSourceObservedState_FromProto(mapCtx *direct.MapContext, in *pb.GoogleApiSource) *krm.GoogleApiSourceObservedState {
	if in == nil {
		return nil
	}
	out := &krm.GoogleApiSourceObservedState{}
	// MISSING: Name
	out.Uid = direct.LazyPtr(in.GetUid())
	out.Etag = direct.LazyPtr(in.GetEtag())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: Annotations
	// MISSING: DisplayName
	// MISSING: Destination
	// MISSING: CryptoKeyName
	// MISSING: LoggingConfig
	return out
}
func GoogleApiSourceObservedState_ToProto(mapCtx *direct.MapContext, in *krm.GoogleApiSourceObservedState) *pb.GoogleApiSource {
	if in == nil {
		return nil
	}
	out := &pb.GoogleApiSource{}
	// MISSING: Name
	out.Uid = direct.ValueOf(in.Uid)
	out.Etag = direct.ValueOf(in.Etag)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: Annotations
	// MISSING: DisplayName
	// MISSING: Destination
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
