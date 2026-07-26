// Copyright 2026 Google LLC
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
	pb "cloud.google.com/go/speech/apiv2/speechpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/speech/v1beta1"
	krmspeechv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/speech/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// --- Unversioned delegating forwarders / manual overrides version wrappers ---

func SpeechCustomClassSpec_FromProto(mapCtx *direct.MapContext, in *pb.CustomClass) *krmspeechv1beta1.SpeechCustomClassSpec {
	return SpeechCustomClassSpec_v1beta1_FromProto(mapCtx, in)
}

func SpeechCustomClassSpec_ToProto(mapCtx *direct.MapContext, in *krmspeechv1beta1.SpeechCustomClassSpec) *pb.CustomClass {
	return SpeechCustomClassSpec_v1beta1_ToProto(mapCtx, in)
}

func SpeechRecognizerSpec_FromProto(mapCtx *direct.MapContext, in *pb.Recognizer) *krmspeechv1beta1.SpeechRecognizerSpec {
	return SpeechRecognizerSpec_v1beta1_FromProto(mapCtx, in)
}

func SpeechRecognizerSpec_ToProto(mapCtx *direct.MapContext, in *krmspeechv1beta1.SpeechRecognizerSpec) *pb.Recognizer {
	return SpeechRecognizerSpec_v1beta1_ToProto(mapCtx, in)
}

func SpeechRecognizerObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Recognizer) *krm.SpeechRecognizerObservedState {
	return SpeechRecognizerObservedState_v1beta1_FromProto(mapCtx, in)
}

// --- Unversioned delegating forwarders / manual overrides version wrappers ---

func SpeechRecognizerObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SpeechRecognizerObservedState) *pb.Recognizer {
	return SpeechRecognizerObservedState_v1beta1_ToProto(mapCtx, in)
}
