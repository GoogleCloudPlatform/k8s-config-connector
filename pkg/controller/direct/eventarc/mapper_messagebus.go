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

package eventarc

import (
	pb "cloud.google.com/go/eventarc/apiv1/eventarcpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/eventarc/v1alpha1"
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func EventarcMessageBusSpec_FromProto(mapCtx *direct.MapContext, in *pb.MessageBus) *krm.EventarcMessageBusSpec {
	if in == nil {
		return nil
	}
	out := &krm.EventarcMessageBusSpec{}
	out.Labels = in.Labels
	out.Annotations = in.Annotations
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	if in.GetCryptoKeyName() != "" {
		out.CryptoKeyRef = &kmsv1beta1.KMSCryptoKeyRef{
			External: in.GetCryptoKeyName(),
		}
	}
	out.LoggingConfig = LoggingConfig_FromProto(mapCtx, in.GetLoggingConfig())
	return out
}

func EventarcMessageBusSpec_ToProto(mapCtx *direct.MapContext, in *krm.EventarcMessageBusSpec) *pb.MessageBus {
	if in == nil {
		return nil
	}
	out := &pb.MessageBus{}
	out.Labels = in.Labels
	out.Annotations = in.Annotations
	out.DisplayName = direct.ValueOf(in.DisplayName)
	if in.CryptoKeyRef != nil {
		out.CryptoKeyName = in.CryptoKeyRef.External
	}
	out.LoggingConfig = LoggingConfig_ToProto(mapCtx, in.LoggingConfig)
	return out
}
