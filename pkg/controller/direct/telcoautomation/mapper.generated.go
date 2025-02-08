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

package telcoautomation

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/telcoautomation/apiv1/telcoautomationpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/telcoautomation/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func PublicBlueprint_FromProto(mapCtx *direct.MapContext, in *pb.PublicBlueprint) *krm.PublicBlueprint {
	if in == nil {
		return nil
	}
	out := &krm.PublicBlueprint{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.DeploymentLevel = direct.Enum_FromProto(mapCtx, in.GetDeploymentLevel())
	out.SourceProvider = direct.LazyPtr(in.GetSourceProvider())
	// MISSING: RollbackSupport
	return out
}
func PublicBlueprint_ToProto(mapCtx *direct.MapContext, in *krm.PublicBlueprint) *pb.PublicBlueprint {
	if in == nil {
		return nil
	}
	out := &pb.PublicBlueprint{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.DeploymentLevel = direct.Enum_ToProto[pb.DeploymentLevel](mapCtx, in.DeploymentLevel)
	out.SourceProvider = direct.ValueOf(in.SourceProvider)
	// MISSING: RollbackSupport
	return out
}
func PublicBlueprintObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PublicBlueprint) *krm.PublicBlueprintObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PublicBlueprintObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: DeploymentLevel
	// MISSING: SourceProvider
	out.RollbackSupport = direct.LazyPtr(in.GetRollbackSupport())
	return out
}
func PublicBlueprintObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PublicBlueprintObservedState) *pb.PublicBlueprint {
	if in == nil {
		return nil
	}
	out := &pb.PublicBlueprint{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: DeploymentLevel
	// MISSING: SourceProvider
	out.RollbackSupport = direct.ValueOf(in.RollbackSupport)
	return out
}
func TelcoautomationPublicBlueprintObservedState_FromProto(mapCtx *direct.MapContext, in *pb.PublicBlueprint) *krm.TelcoautomationPublicBlueprintObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TelcoautomationPublicBlueprintObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: DeploymentLevel
	// MISSING: SourceProvider
	// MISSING: RollbackSupport
	return out
}
func TelcoautomationPublicBlueprintObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TelcoautomationPublicBlueprintObservedState) *pb.PublicBlueprint {
	if in == nil {
		return nil
	}
	out := &pb.PublicBlueprint{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: DeploymentLevel
	// MISSING: SourceProvider
	// MISSING: RollbackSupport
	return out
}
func TelcoautomationPublicBlueprintSpec_FromProto(mapCtx *direct.MapContext, in *pb.PublicBlueprint) *krm.TelcoautomationPublicBlueprintSpec {
	if in == nil {
		return nil
	}
	out := &krm.TelcoautomationPublicBlueprintSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: DeploymentLevel
	// MISSING: SourceProvider
	// MISSING: RollbackSupport
	return out
}
func TelcoautomationPublicBlueprintSpec_ToProto(mapCtx *direct.MapContext, in *krm.TelcoautomationPublicBlueprintSpec) *pb.PublicBlueprint {
	if in == nil {
		return nil
	}
	out := &pb.PublicBlueprint{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: DeploymentLevel
	// MISSING: SourceProvider
	// MISSING: RollbackSupport
	return out
}
