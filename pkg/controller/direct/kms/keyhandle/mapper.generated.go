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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func KMSKeyHandleSpec_FromProto(mapCtx *direct.MapContext, in *pb.KeyHandle) *krm.KMSKeyHandleSpec {
	if in == nil || in.Name == "" {
		return nil
	}
	out := &krm.KMSKeyHandleSpec{}
	parent, resourceID, err := krm.ParseKMSKeyHandleExternal(in.Name)
	if err != nil {
		return nil
	}
	out.ResourceID = &resourceID
	out.ProjectRef = &refs.ProjectRef{
		External: "projects/" + parent.ProjectID,
	}
	out.Location = &parent.Location
	out.ResourceTypeSelector = &in.ResourceTypeSelector
	return out
}

func KMSKeyHandleSpec_ToProto(mapCtx *direct.MapContext, in *krm.KMSKeyHandleSpec, id *krm.KMSKeyHandleRef) *pb.KeyHandle {
	if in == nil {
		return nil
	}
	out := &pb.KeyHandle{}
	out.Name = id.External
	out.ResourceTypeSelector = *in.ResourceTypeSelector
	return out
}
func KeyHandle_FromProto(mapCtx *direct.MapContext, in *pb.KeyHandle) *krm.KeyHandle {
	if in == nil {
		return nil
	}
	out := &krm.KeyHandle{}
	out.Name = direct.LazyPtr(in.GetName())
	if in.KmsKey != "" {
		out.KmsKey = direct.LazyPtr(in.GetKmsKey())
	}
	out.ResourceTypeSelector = direct.LazyPtr(in.GetResourceTypeSelector())
	return out
}
func KeyHandle_ToProto(mapCtx *direct.MapContext, in *krm.KeyHandle) *pb.KeyHandle {
	if in == nil {
		return nil
	}
	out := &pb.KeyHandle{}
	out.Name = direct.ValueOf(in.Name)
	out.KmsKey = direct.ValueOf(in.KmsKey)
	out.ResourceTypeSelector = direct.ValueOf(in.ResourceTypeSelector)
	return out
}

func KMSKeyHandleStatusObservedState_FromProto(mapCtx *direct.MapContext, in *pb.KeyHandle) *krm.KMSKeyHandleObservedState {
	out := &krm.KMSKeyHandleObservedState{}
	if in.KmsKey != "" {
		out.KMSKey = direct.LazyPtr(in.GetKmsKey())
	}
	return out
}
