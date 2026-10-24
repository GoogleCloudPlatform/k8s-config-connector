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

package netapp

import (
	pb "cloud.google.com/go/netapp/apiv1/netapppb"
	kmsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/netapp/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func NetAppKMSConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.KmsConfig) *krm.NetAppKMSConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.NetAppKMSConfigSpec{}
	if in.GetCryptoKeyName() != "" {
		out.CryptoKeyRef = &kmsv1beta1.KMSCryptoKeyRef{External: in.GetCryptoKeyName()}
	}
	out.Description = direct.LazyPtr(in.GetDescription())
	return out
}

func NetAppKMSConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.NetAppKMSConfigSpec) *pb.KmsConfig {
	if in == nil {
		return nil
	}
	out := &pb.KmsConfig{}
	if in.CryptoKeyRef != nil {
		out.CryptoKeyName = in.CryptoKeyRef.External
	}
	out.Description = direct.ValueOf(in.Description)
	return out
}
