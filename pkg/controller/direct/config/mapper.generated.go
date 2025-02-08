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

package config

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/config/apiv1/configpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/config/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func ConfigTerraformVersionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TerraformVersion) *krm.ConfigTerraformVersionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ConfigTerraformVersionObservedState{}
	// MISSING: Name
	// MISSING: State
	// MISSING: SupportTime
	// MISSING: DeprecateTime
	// MISSING: ObsoleteTime
	return out
}
func ConfigTerraformVersionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ConfigTerraformVersionObservedState) *pb.TerraformVersion {
	if in == nil {
		return nil
	}
	out := &pb.TerraformVersion{}
	// MISSING: Name
	// MISSING: State
	// MISSING: SupportTime
	// MISSING: DeprecateTime
	// MISSING: ObsoleteTime
	return out
}
func ConfigTerraformVersionSpec_FromProto(mapCtx *direct.MapContext, in *pb.TerraformVersion) *krm.ConfigTerraformVersionSpec {
	if in == nil {
		return nil
	}
	out := &krm.ConfigTerraformVersionSpec{}
	// MISSING: Name
	// MISSING: State
	// MISSING: SupportTime
	// MISSING: DeprecateTime
	// MISSING: ObsoleteTime
	return out
}
func ConfigTerraformVersionSpec_ToProto(mapCtx *direct.MapContext, in *krm.ConfigTerraformVersionSpec) *pb.TerraformVersion {
	if in == nil {
		return nil
	}
	out := &pb.TerraformVersion{}
	// MISSING: Name
	// MISSING: State
	// MISSING: SupportTime
	// MISSING: DeprecateTime
	// MISSING: ObsoleteTime
	return out
}
func TerraformVersion_FromProto(mapCtx *direct.MapContext, in *pb.TerraformVersion) *krm.TerraformVersion {
	if in == nil {
		return nil
	}
	out := &krm.TerraformVersion{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: State
	// MISSING: SupportTime
	// MISSING: DeprecateTime
	// MISSING: ObsoleteTime
	return out
}
func TerraformVersion_ToProto(mapCtx *direct.MapContext, in *krm.TerraformVersion) *pb.TerraformVersion {
	if in == nil {
		return nil
	}
	out := &pb.TerraformVersion{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: State
	// MISSING: SupportTime
	// MISSING: DeprecateTime
	// MISSING: ObsoleteTime
	return out
}
func TerraformVersionObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TerraformVersion) *krm.TerraformVersionObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TerraformVersionObservedState{}
	// MISSING: Name
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.SupportTime = direct.StringTimestamp_FromProto(mapCtx, in.GetSupportTime())
	out.DeprecateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetDeprecateTime())
	out.ObsoleteTime = direct.StringTimestamp_FromProto(mapCtx, in.GetObsoleteTime())
	return out
}
func TerraformVersionObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TerraformVersionObservedState) *pb.TerraformVersion {
	if in == nil {
		return nil
	}
	out := &pb.TerraformVersion{}
	// MISSING: Name
	out.State = direct.Enum_ToProto[pb.TerraformVersion_State](mapCtx, in.State)
	out.SupportTime = direct.StringTimestamp_ToProto(mapCtx, in.SupportTime)
	if oneof := direct.StringTimestamp_ToProto(mapCtx, in.DeprecateTime); oneof != nil {
		out.DeprecateTime = &pb.TerraformVersion_DeprecateTime{DeprecateTime: oneof}
	}
	if oneof := direct.StringTimestamp_ToProto(mapCtx, in.ObsoleteTime); oneof != nil {
		out.ObsoleteTime = &pb.TerraformVersion_ObsoleteTime{ObsoleteTime: oneof}
	}
	return out
}
