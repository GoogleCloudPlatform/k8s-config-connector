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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/eventarc/apiv1/eventarcpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/eventarc/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func EventarcGoogleChannelConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.GoogleChannelConfig) *krm.EventarcGoogleChannelConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.EventarcGoogleChannelConfigObservedState{}
	// MISSING: Name
	// MISSING: UpdateTime
	// MISSING: CryptoKeyName
	return out
}
func EventarcGoogleChannelConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.EventarcGoogleChannelConfigObservedState) *pb.GoogleChannelConfig {
	if in == nil {
		return nil
	}
	out := &pb.GoogleChannelConfig{}
	// MISSING: Name
	// MISSING: UpdateTime
	// MISSING: CryptoKeyName
	return out
}
func EventarcGoogleChannelConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.GoogleChannelConfig) *krm.EventarcGoogleChannelConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.EventarcGoogleChannelConfigSpec{}
	// MISSING: Name
	// MISSING: UpdateTime
	// MISSING: CryptoKeyName
	return out
}
func EventarcGoogleChannelConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.EventarcGoogleChannelConfigSpec) *pb.GoogleChannelConfig {
	if in == nil {
		return nil
	}
	out := &pb.GoogleChannelConfig{}
	// MISSING: Name
	// MISSING: UpdateTime
	// MISSING: CryptoKeyName
	return out
}
func GoogleChannelConfig_FromProto(mapCtx *direct.MapContext, in *pb.GoogleChannelConfig) *krm.GoogleChannelConfig {
	if in == nil {
		return nil
	}
	out := &krm.GoogleChannelConfig{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: UpdateTime
	out.CryptoKeyName = direct.LazyPtr(in.GetCryptoKeyName())
	return out
}
func GoogleChannelConfig_ToProto(mapCtx *direct.MapContext, in *krm.GoogleChannelConfig) *pb.GoogleChannelConfig {
	if in == nil {
		return nil
	}
	out := &pb.GoogleChannelConfig{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: UpdateTime
	out.CryptoKeyName = direct.ValueOf(in.CryptoKeyName)
	return out
}
func GoogleChannelConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.GoogleChannelConfig) *krm.GoogleChannelConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.GoogleChannelConfigObservedState{}
	// MISSING: Name
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: CryptoKeyName
	return out
}
func GoogleChannelConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.GoogleChannelConfigObservedState) *pb.GoogleChannelConfig {
	if in == nil {
		return nil
	}
	out := &pb.GoogleChannelConfig{}
	// MISSING: Name
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: CryptoKeyName
	return out
}
