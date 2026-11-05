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

// +generated:mapper
// krm.group: eventarc.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.eventarc.v1

package eventarc

import (
	pb "cloud.google.com/go/eventarc/apiv1/eventarcpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/eventarc/v1alpha1"
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func EventarcGoogleAPISourceSpec_FromProto(mapCtx *direct.MapContext, in *pb.GoogleApiSource) *krm.EventarcGoogleAPISourceSpec {
	if in == nil {
		return nil
	}
	out := &krm.EventarcGoogleAPISourceSpec{}
	out.Labels = in.Labels
	out.Annotations = in.Annotations
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	if in.GetDestination() != "" {
		out.MessageBusRef = &krm.EventarcMessageBusRef{
			External: in.GetDestination(),
		}
	}
	if in.GetCryptoKeyName() != "" {
		out.CryptoKeyRef = &kmsv1beta1.KMSCryptoKeyRef{
			External: in.GetCryptoKeyName(),
		}
	}
	out.LoggingConfig = LoggingConfig_FromProto(mapCtx, in.GetLoggingConfig())
	return out
}

func EventarcGoogleAPISourceSpec_ToProto(mapCtx *direct.MapContext, in *krm.EventarcGoogleAPISourceSpec) *pb.GoogleApiSource {
	if in == nil {
		return nil
	}
	out := &pb.GoogleApiSource{}
	out.Labels = in.Labels
	out.Annotations = in.Annotations
	out.DisplayName = direct.ValueOf(in.DisplayName)
	if in.MessageBusRef != nil {
		out.Destination = in.MessageBusRef.External
	}
	if in.CryptoKeyRef != nil {
		out.CryptoKeyName = in.CryptoKeyRef.External
	}
	out.LoggingConfig = LoggingConfig_ToProto(mapCtx, in.LoggingConfig)
	return out
}
