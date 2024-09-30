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

package kms

import (
	pb "cloud.google.com/go/kms/apiv1/kmspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/kms/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func AutokeyConfig_FromProto(mapCtx *direct.MapContext, in *pb.AutokeyConfig) *krm.AutokeyConfig {
	if in == nil {
		return nil
	}
	out := &krm.AutokeyConfig{}
	out.Name = direct.LazyPtr(in.GetName())
	out.KeyProject = direct.LazyPtr(in.GetKeyProject())
	// MISSING: State
	return out
}
func AutokeyConfig_ToProto(mapCtx *direct.MapContext, in *krm.AutokeyConfig) *pb.AutokeyConfig {
	if in == nil {
		return nil
	}
	out := &pb.AutokeyConfig{}
	out.Name = direct.ValueOf(in.Name)
	out.KeyProject = direct.ValueOf(in.KeyProject)
	// MISSING: State
	return out
}
func KMSAutokeyConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AutokeyConfig) *krm.KMSAutokeyConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.KMSAutokeyConfigObservedState{
		//State:  direct.Enum_FromProto(mapCtx, in.GetState()),
	}
	return out
}
func KMSAutokeyConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.KMSAutokeyConfigObservedState) *pb.AutokeyConfig {
	if in == nil {
		return nil
	}
	out := &pb.AutokeyConfig{
		//State:  direct.Enum_ToProto[pb.AutokeyConfig_State](mapCtx, in.State),
	}
	return out
}
func KMSAutokeyConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.AutokeyConfig) *krm.KMSAutokeyConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.KMSAutokeyConfigSpec{}
	// MISSING: Name
	// MISSING: KeyProject
	//out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	return out
}
func KMSAutokeyConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.KMSAutokeyConfigSpec) *pb.AutokeyConfig {
	if in == nil {
		return nil
	}
	out := &pb.AutokeyConfig{}
	// MISSING: Name
	// MISSING: KeyProject
	//out.State = direct.Enum_ToProto[pb.AutokeyConfig_State](mapCtx, in.State)
	return out
}
/*
func KeyHandle_FromProto(mapCtx *direct.MapContext, in *pb.KeyHandle) *krm.KeyHandle {
	if in == nil {
		return nil
	}
	out := &krm.KeyHandle{}
	out.Name = direct.LazyPtr(in.GetName())
	out.KmsKey = direct.LazyPtr(in.GetKmsKey())
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
*/
