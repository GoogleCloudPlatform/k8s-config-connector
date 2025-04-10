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

package keyhandle

import (
	pb "cloud.google.com/go/kms/apiv1/kmspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func KMSKeyHandleSpec_FromProto(mapCtx *direct.MapContext, in *pb.KeyHandle) *krm.KMSKeyHandleSpec {
	if in == nil {
		return nil
	}
	out := &krm.KMSKeyHandleSpec{}
	out.ResourceTypeSelector = &in.ResourceTypeSelector
	return out
}

func KMSKeyHandleSpec_ToProto(mapCtx *direct.MapContext, in *krm.KMSKeyHandleSpec) *pb.KeyHandle {
	if in == nil {
		return nil
	}
	out := &pb.KeyHandle{}
	out.ResourceTypeSelector = *in.ResourceTypeSelector
	return out
}

func KMSKeyHandleStatusObservedState_FromProto(mapCtx *direct.MapContext, in *pb.KeyHandle) *krm.KMSKeyHandleObservedState {
	out := &krm.KMSKeyHandleObservedState{}
	if in.KmsKey != "" {
		out.KMSKey = direct.LazyPtr(in.GetKmsKey())
	}
	return out
}
