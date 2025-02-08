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

package speech

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/speech/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/speech/apiv2/speechpb"
)
func Config_FromProto(mapCtx *direct.MapContext, in *pb.Config) *krm.Config {
	if in == nil {
		return nil
	}
	out := &krm.Config{}
	// MISSING: Name
	out.KMSKeyName = direct.LazyPtr(in.GetKmsKeyName())
	// MISSING: UpdateTime
	return out
}
func Config_ToProto(mapCtx *direct.MapContext, in *krm.Config) *pb.Config {
	if in == nil {
		return nil
	}
	out := &pb.Config{}
	// MISSING: Name
	out.KmsKeyName = direct.ValueOf(in.KMSKeyName)
	// MISSING: UpdateTime
	return out
}
func ConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Config) *krm.ConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ConfigObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: KMSKeyName
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func ConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ConfigObservedState) *pb.Config {
	if in == nil {
		return nil
	}
	out := &pb.Config{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: KMSKeyName
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func SpeechConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Config) *krm.SpeechConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SpeechConfigObservedState{}
	// MISSING: Name
	// MISSING: KMSKeyName
	// MISSING: UpdateTime
	return out
}
func SpeechConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SpeechConfigObservedState) *pb.Config {
	if in == nil {
		return nil
	}
	out := &pb.Config{}
	// MISSING: Name
	// MISSING: KMSKeyName
	// MISSING: UpdateTime
	return out
}
func SpeechConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.Config) *krm.SpeechConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.SpeechConfigSpec{}
	// MISSING: Name
	// MISSING: KMSKeyName
	// MISSING: UpdateTime
	return out
}
func SpeechConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.SpeechConfigSpec) *pb.Config {
	if in == nil {
		return nil
	}
	out := &pb.Config{}
	// MISSING: Name
	// MISSING: KMSKeyName
	// MISSING: UpdateTime
	return out
}
