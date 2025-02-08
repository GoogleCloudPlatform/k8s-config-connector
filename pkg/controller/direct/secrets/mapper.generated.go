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

package secrets

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/secrets/apiv1beta1/secretspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/secrets/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func SecretVersion_FromProto(mapCtx *direct.MapContext, in *pb.SecretVersion) *krm.SecretVersion {
	if in == nil {
		return nil
	}
	out := &krm.SecretVersion{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: DestroyTime
	// MISSING: State
	return out
}
func SecretVersion_ToProto(mapCtx *direct.MapContext, in *krm.SecretVersion) *pb.SecretVersion {
	if in == nil {
		return nil
	}
	out := &pb.SecretVersion{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: DestroyTime
	// MISSING: State
	return out
}
func SecretVersionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SecretVersion) *krm.SecretVersionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SecretVersionObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.DestroyTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDestroyTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}
func SecretVersionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SecretVersionObservedState) *pb.SecretVersion {
	if in == nil {
		return nil
	}
	out := &pb.SecretVersion{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.DestroyTime = direct.StringTimestamp_ToProto(mapCtx, in.DestroyTime)
	out.State = direct.Enum_ToProto[pb.SecretVersion_State](mapCtx, in.State)
	return out
}
func SecretsSecretVersionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SecretVersion) *krm.SecretsSecretVersionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SecretsSecretVersionObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: DestroyTime
	// MISSING: State
	return out
}
func SecretsSecretVersionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SecretsSecretVersionObservedState) *pb.SecretVersion {
	if in == nil {
		return nil
	}
	out := &pb.SecretVersion{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: DestroyTime
	// MISSING: State
	return out
}
func SecretsSecretVersionSpec_FromProto(mapCtx *direct.MapContext, in *pb.SecretVersion) *krm.SecretsSecretVersionSpec {
	if in == nil {
		return nil
	}
	out := &krm.SecretsSecretVersionSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: DestroyTime
	// MISSING: State
	return out
}
func SecretsSecretVersionSpec_ToProto(mapCtx *direct.MapContext, in *krm.SecretsSecretVersionSpec) *pb.SecretVersion {
	if in == nil {
		return nil
	}
	out := &pb.SecretVersion{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: DestroyTime
	// MISSING: State
	return out
}
