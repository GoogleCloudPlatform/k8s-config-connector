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

package netapp

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/netapp/apiv1/netapppb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/netapp/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func KmsConfig_FromProto(mapCtx *direct.MapContext, in *pb.KmsConfig) *krm.KmsConfig {
	if in == nil {
		return nil
	}
	out := &krm.KmsConfig{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CryptoKeyName = direct.LazyPtr(in.GetCryptoKeyName())
	// MISSING: State
	// MISSING: StateDetails
	// MISSING: CreateTime
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Labels = in.Labels
	// MISSING: Instructions
	// MISSING: ServiceAccount
	return out
}
func KmsConfig_ToProto(mapCtx *direct.MapContext, in *krm.KmsConfig) *pb.KmsConfig {
	if in == nil {
		return nil
	}
	out := &pb.KmsConfig{}
	out.Name = direct.ValueOf(in.Name)
	out.CryptoKeyName = direct.ValueOf(in.CryptoKeyName)
	// MISSING: State
	// MISSING: StateDetails
	// MISSING: CreateTime
	out.Description = direct.ValueOf(in.Description)
	out.Labels = in.Labels
	// MISSING: Instructions
	// MISSING: ServiceAccount
	return out
}
func KmsConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.KmsConfig) *krm.KmsConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.KmsConfigObservedState{}
	// MISSING: Name
	// MISSING: CryptoKeyName
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.StateDetails = direct.LazyPtr(in.GetStateDetails())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	// MISSING: Description
	// MISSING: Labels
	out.Instructions = direct.LazyPtr(in.GetInstructions())
	out.ServiceAccount = direct.LazyPtr(in.GetServiceAccount())
	return out
}
func KmsConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.KmsConfigObservedState) *pb.KmsConfig {
	if in == nil {
		return nil
	}
	out := &pb.KmsConfig{}
	// MISSING: Name
	// MISSING: CryptoKeyName
	out.State = direct.Enum_ToProto[pb.KmsConfig_State](mapCtx, in.State)
	out.StateDetails = direct.ValueOf(in.StateDetails)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	// MISSING: Description
	// MISSING: Labels
	out.Instructions = direct.ValueOf(in.Instructions)
	out.ServiceAccount = direct.ValueOf(in.ServiceAccount)
	return out
}
func NetappKmsConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.KmsConfig) *krm.NetappKmsConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.NetappKmsConfigObservedState{}
	// MISSING: Name
	// MISSING: CryptoKeyName
	// MISSING: State
	// MISSING: StateDetails
	// MISSING: CreateTime
	// MISSING: Description
	// MISSING: Labels
	// MISSING: Instructions
	// MISSING: ServiceAccount
	return out
}
func NetappKmsConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.NetappKmsConfigObservedState) *pb.KmsConfig {
	if in == nil {
		return nil
	}
	out := &pb.KmsConfig{}
	// MISSING: Name
	// MISSING: CryptoKeyName
	// MISSING: State
	// MISSING: StateDetails
	// MISSING: CreateTime
	// MISSING: Description
	// MISSING: Labels
	// MISSING: Instructions
	// MISSING: ServiceAccount
	return out
}
func NetappKmsConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.KmsConfig) *krm.NetappKmsConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.NetappKmsConfigSpec{}
	// MISSING: Name
	// MISSING: CryptoKeyName
	// MISSING: State
	// MISSING: StateDetails
	// MISSING: CreateTime
	// MISSING: Description
	// MISSING: Labels
	// MISSING: Instructions
	// MISSING: ServiceAccount
	return out
}
func NetappKmsConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.NetappKmsConfigSpec) *pb.KmsConfig {
	if in == nil {
		return nil
	}
	out := &pb.KmsConfig{}
	// MISSING: Name
	// MISSING: CryptoKeyName
	// MISSING: State
	// MISSING: StateDetails
	// MISSING: CreateTime
	// MISSING: Description
	// MISSING: Labels
	// MISSING: Instructions
	// MISSING: ServiceAccount
	return out
}
