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

// +generated:mapper
// krm.group: eventarc.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.eventarc.v1

package eventarc

import (
	pb "cloud.google.com/go/eventarc/apiv1/eventarcpb"
	krmv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/eventarc/v1alpha1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func EventarcGoogleChannelConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.GoogleChannelConfig) *krmv1alpha1.EventarcGoogleChannelConfigSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.EventarcGoogleChannelConfigSpec{}
	// MISSING: Name
	if in.GetCryptoKeyName() != "" {
		out.CryptoKeyRef = &refs.KMSCryptoKeyRef{
			External: in.GetCryptoKeyName(),
		}
	}
	return out
}
func EventarcGoogleChannelConfigSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.EventarcGoogleChannelConfigSpec) *pb.GoogleChannelConfig {
	if in == nil {
		return nil
	}
	out := &pb.GoogleChannelConfig{}
	// MISSING: Name
	if in.CryptoKeyRef != nil && in.CryptoKeyRef.External != "" {
		out.CryptoKeyName = in.CryptoKeyRef.External
	}
	return out
}
