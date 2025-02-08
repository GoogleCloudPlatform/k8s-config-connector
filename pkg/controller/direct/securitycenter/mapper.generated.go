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

package securitycenter

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/securitycenter/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/securitycenter/apiv1/securitycenterpb"
)
func MuteConfig_FromProto(mapCtx *direct.MapContext, in *pb.MuteConfig) *krm.MuteConfig {
	if in == nil {
		return nil
	}
	out := &krm.MuteConfig{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Filter = direct.LazyPtr(in.GetFilter())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: MostRecentEditor
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.ExpiryTime = direct.StringTimestamp_FromProto(mapCtx, in.GetExpiryTime())
	return out
}
func MuteConfig_ToProto(mapCtx *direct.MapContext, in *krm.MuteConfig) *pb.MuteConfig {
	if in == nil {
		return nil
	}
	out := &pb.MuteConfig{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.Filter = direct.ValueOf(in.Filter)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: MostRecentEditor
	out.Type = direct.Enum_ToProto[pb.MuteConfig_MuteConfigType](mapCtx, in.Type)
	out.ExpiryTime = direct.StringTimestamp_ToProto(mapCtx, in.ExpiryTime)
	return out
}
func MuteConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MuteConfig) *krm.MuteConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MuteConfigObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Filter
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.MostRecentEditor = direct.LazyPtr(in.GetMostRecentEditor())
	// MISSING: Type
	// MISSING: ExpiryTime
	return out
}
func MuteConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MuteConfigObservedState) *pb.MuteConfig {
	if in == nil {
		return nil
	}
	out := &pb.MuteConfig{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Filter
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.MostRecentEditor = direct.ValueOf(in.MostRecentEditor)
	// MISSING: Type
	// MISSING: ExpiryTime
	return out
}
func SecuritycenterMuteConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.MuteConfig) *krm.SecuritycenterMuteConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SecuritycenterMuteConfigObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Filter
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: MostRecentEditor
	// MISSING: Type
	// MISSING: ExpiryTime
	return out
}
func SecuritycenterMuteConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SecuritycenterMuteConfigObservedState) *pb.MuteConfig {
	if in == nil {
		return nil
	}
	out := &pb.MuteConfig{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Filter
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: MostRecentEditor
	// MISSING: Type
	// MISSING: ExpiryTime
	return out
}
func SecuritycenterMuteConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.MuteConfig) *krm.SecuritycenterMuteConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.SecuritycenterMuteConfigSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Filter
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: MostRecentEditor
	// MISSING: Type
	// MISSING: ExpiryTime
	return out
}
func SecuritycenterMuteConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.SecuritycenterMuteConfigSpec) *pb.MuteConfig {
	if in == nil {
		return nil
	}
	out := &pb.MuteConfig{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: Filter
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: MostRecentEditor
	// MISSING: Type
	// MISSING: ExpiryTime
	return out
}
