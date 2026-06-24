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

package kms

import (
	pb "cloud.google.com/go/kms/apiv1/kmspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func KMSKeyRingSpec_FromProto(mapCtx *direct.MapContext, in *pb.KeyRing) *krm.KMSKeyRingSpec {
	if in == nil {
		return nil
	}
	out := &krm.KMSKeyRingSpec{}
	return out
}

func KMSKeyRingSpec_ToProto(mapCtx *direct.MapContext, in *krm.KMSKeyRingSpec) *pb.KeyRing {
	if in == nil {
		return nil
	}
	out := &pb.KeyRing{}
	return out
}

func KMSKeyRingStatus_FromProto(mapCtx *direct.MapContext, in *pb.KeyRing) *krm.KMSKeyRingStatus {
	if in == nil {
		return nil
	}
	out := &krm.KMSKeyRingStatus{}
	if in.Name != "" {
		out.SelfLink = &in.Name
	}
	return out
}

func KMSKeyRingStatus_ToProto(mapCtx *direct.MapContext, in *krm.KMSKeyRingStatus) *pb.KeyRing {
	if in == nil {
		return nil
	}
	out := &pb.KeyRing{}
	if in.SelfLink != nil {
		out.Name = *in.SelfLink
	}
	return out
}
