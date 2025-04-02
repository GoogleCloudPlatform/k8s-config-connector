// Copyright 2024 Google LLC
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

// +tool:fuzz-gen
// proto.message: google.cloud.eventarc.v1.GoogleChannelConfig
// api.group: eventarc.cnrm.cloud.google.com

package eventarc

import (
	pb "cloud.google.com/go/eventarc/apiv1/eventarcpb"
	krmv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/eventarc/v1alpha1"
	refv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

// Needed for generated fuzzers
// Autogen overrides
func EventarcGoogleChannelConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.GoogleChannelConfig) *krmv1alpha1.EventarcGoogleChannelConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.EventarcGoogleChannelConfigObservedState{}
	// MISSING: Name
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: CryptoKeyName
	return out
}

// Needed for generated fuzzers
// Autogen overrides
func EventarcGoogleChannelConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.EventarcGoogleChannelConfigObservedState) *pb.GoogleChannelConfig {
	if in == nil {
		return nil
	}
	out := &pb.GoogleChannelConfig{}
	// MISSING: Name
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: CryptoKeyName
	return out
}

// Needed for generated fuzzers
// Autogen overrides
func EventarcGoogleChannelConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.GoogleChannelConfig) *krmv1alpha1.EventarcGoogleChannelConfigSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.EventarcGoogleChannelConfigSpec{}
	// MISSING: Name
	if in.GetCryptoKeyName() != "" {
		out.CryptoKeyRef = &refv1beta1.KMSCryptoKeyRef{External: in.GetCryptoKeyName()}
	}
	return out
}

// Needed for generated fuzzers
// Autogen overrides
func EventarcGoogleChannelConfigSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.EventarcGoogleChannelConfigSpec) *pb.GoogleChannelConfig {
	if in == nil {
		return nil
	}
	out := &pb.GoogleChannelConfig{}
	// MISSING: Name
	if in.CryptoKeyRef != nil {
		out.CryptoKeyName = in.CryptoKeyRef.External
	}
	return out
}

func init() {
	fuzztesting.RegisterKRMFuzzer(eventarcGoogleChannelConfigFuzzer())
}

func eventarcGoogleChannelConfigFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.GoogleChannelConfig{},
		EventarcGoogleChannelConfigSpec_FromProto, EventarcGoogleChannelConfigSpec_ToProto,
		EventarcGoogleChannelConfigObservedState_FromProto, EventarcGoogleChannelConfigObservedState_ToProto,
	)

	f.SpecFields.Insert(".crypto_key_name")

	f.StatusFields.Insert(".update_time")

	f.UnimplementedFields.Insert(".name") // special field

	return f
}
