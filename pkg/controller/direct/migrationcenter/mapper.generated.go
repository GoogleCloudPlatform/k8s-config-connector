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

package migrationcenter

import (
	pb "cloud.google.com/go/migrationcenter/apiv1/migrationcenterpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/migrationcenter/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func MigrationcenterSettingsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Settings) *krm.MigrationcenterSettingsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MigrationcenterSettingsObservedState{}
	// MISSING: Name
	// MISSING: PreferenceSet
	return out
}
func MigrationcenterSettingsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MigrationcenterSettingsObservedState) *pb.Settings {
	if in == nil {
		return nil
	}
	out := &pb.Settings{}
	// MISSING: Name
	// MISSING: PreferenceSet
	return out
}
func MigrationcenterSettingsSpec_FromProto(mapCtx *direct.MapContext, in *pb.Settings) *krm.MigrationcenterSettingsSpec {
	if in == nil {
		return nil
	}
	out := &krm.MigrationcenterSettingsSpec{}
	// MISSING: Name
	// MISSING: PreferenceSet
	return out
}
func MigrationcenterSettingsSpec_ToProto(mapCtx *direct.MapContext, in *krm.MigrationcenterSettingsSpec) *pb.Settings {
	if in == nil {
		return nil
	}
	out := &pb.Settings{}
	// MISSING: Name
	// MISSING: PreferenceSet
	return out
}
func Settings_FromProto(mapCtx *direct.MapContext, in *pb.Settings) *krm.Settings {
	if in == nil {
		return nil
	}
	out := &krm.Settings{}
	// MISSING: Name
	out.PreferenceSet = direct.LazyPtr(in.GetPreferenceSet())
	return out
}
func Settings_ToProto(mapCtx *direct.MapContext, in *krm.Settings) *pb.Settings {
	if in == nil {
		return nil
	}
	out := &pb.Settings{}
	// MISSING: Name
	out.PreferenceSet = direct.ValueOf(in.PreferenceSet)
	return out
}
func SettingsObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Settings) *krm.SettingsObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SettingsObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: PreferenceSet
	return out
}
func SettingsObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SettingsObservedState) *pb.Settings {
	if in == nil {
		return nil
	}
	out := &pb.Settings{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: PreferenceSet
	return out
}
