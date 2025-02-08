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

package deploy

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/deploy/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/deploy/apiv1/deploypb"
)
func Config_FromProto(mapCtx *direct.MapContext, in *pb.Config) *krm.Config {
	if in == nil {
		return nil
	}
	out := &krm.Config{}
	out.Name = direct.LazyPtr(in.GetName())
	out.SupportedVersions = direct.Slice_FromProto(mapCtx, in.SupportedVersions, SkaffoldVersion_FromProto)
	out.DefaultSkaffoldVersion = direct.LazyPtr(in.GetDefaultSkaffoldVersion())
	return out
}
func Config_ToProto(mapCtx *direct.MapContext, in *krm.Config) *pb.Config {
	if in == nil {
		return nil
	}
	out := &pb.Config{}
	out.Name = direct.ValueOf(in.Name)
	out.SupportedVersions = direct.Slice_ToProto(mapCtx, in.SupportedVersions, SkaffoldVersion_ToProto)
	out.DefaultSkaffoldVersion = direct.ValueOf(in.DefaultSkaffoldVersion)
	return out
}
func DeployConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Config) *krm.DeployConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DeployConfigObservedState{}
	// MISSING: Name
	// MISSING: SupportedVersions
	// MISSING: DefaultSkaffoldVersion
	return out
}
func DeployConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DeployConfigObservedState) *pb.Config {
	if in == nil {
		return nil
	}
	out := &pb.Config{}
	// MISSING: Name
	// MISSING: SupportedVersions
	// MISSING: DefaultSkaffoldVersion
	return out
}
func DeployConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.Config) *krm.DeployConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.DeployConfigSpec{}
	// MISSING: Name
	// MISSING: SupportedVersions
	// MISSING: DefaultSkaffoldVersion
	return out
}
func DeployConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.DeployConfigSpec) *pb.Config {
	if in == nil {
		return nil
	}
	out := &pb.Config{}
	// MISSING: Name
	// MISSING: SupportedVersions
	// MISSING: DefaultSkaffoldVersion
	return out
}
func SkaffoldVersion_FromProto(mapCtx *direct.MapContext, in *pb.SkaffoldVersion) *krm.SkaffoldVersion {
	if in == nil {
		return nil
	}
	out := &krm.SkaffoldVersion{}
	out.Version = direct.LazyPtr(in.GetVersion())
	out.MaintenanceModeTime = direct.StringTimestamp_FromProto(mapCtx, in.GetMaintenanceModeTime())
	out.SupportExpirationTime = direct.StringTimestamp_FromProto(mapCtx, in.GetSupportExpirationTime())
	out.SupportEndDate = Date_FromProto(mapCtx, in.GetSupportEndDate())
	return out
}
func SkaffoldVersion_ToProto(mapCtx *direct.MapContext, in *krm.SkaffoldVersion) *pb.SkaffoldVersion {
	if in == nil {
		return nil
	}
	out := &pb.SkaffoldVersion{}
	out.Version = direct.ValueOf(in.Version)
	out.MaintenanceModeTime = direct.StringTimestamp_ToProto(mapCtx, in.MaintenanceModeTime)
	out.SupportExpirationTime = direct.StringTimestamp_ToProto(mapCtx, in.SupportExpirationTime)
	out.SupportEndDate = Date_ToProto(mapCtx, in.SupportEndDate)
	return out
}
