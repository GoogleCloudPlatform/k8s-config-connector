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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/migrationcenter/apiv1/migrationcenterpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/migrationcenter/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func MigrationcenterReportConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ReportConfig) *krm.MigrationcenterReportConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.MigrationcenterReportConfigObservedState{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: GroupPreferencesetAssignments
	return out
}
func MigrationcenterReportConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.MigrationcenterReportConfigObservedState) *pb.ReportConfig {
	if in == nil {
		return nil
	}
	out := &pb.ReportConfig{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: GroupPreferencesetAssignments
	return out
}
func MigrationcenterReportConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.ReportConfig) *krm.MigrationcenterReportConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.MigrationcenterReportConfigSpec{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: GroupPreferencesetAssignments
	return out
}
func MigrationcenterReportConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.MigrationcenterReportConfigSpec) *pb.ReportConfig {
	if in == nil {
		return nil
	}
	out := &pb.ReportConfig{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: GroupPreferencesetAssignments
	return out
}
func ReportConfig_FromProto(mapCtx *direct.MapContext, in *pb.ReportConfig) *krm.ReportConfig {
	if in == nil {
		return nil
	}
	out := &krm.ReportConfig{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.GroupPreferencesetAssignments = direct.Slice_FromProto(mapCtx, in.GroupPreferencesetAssignments, ReportConfig_GroupPreferenceSetAssignment_FromProto)
	return out
}
func ReportConfig_ToProto(mapCtx *direct.MapContext, in *krm.ReportConfig) *pb.ReportConfig {
	if in == nil {
		return nil
	}
	out := &pb.ReportConfig{}
	// MISSING: Name
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.GroupPreferencesetAssignments = direct.Slice_ToProto(mapCtx, in.GroupPreferencesetAssignments, ReportConfig_GroupPreferenceSetAssignment_ToProto)
	return out
}
func ReportConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ReportConfig) *krm.ReportConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ReportConfigObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: GroupPreferencesetAssignments
	return out
}
func ReportConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ReportConfigObservedState) *pb.ReportConfig {
	if in == nil {
		return nil
	}
	out := &pb.ReportConfig{}
	out.Name = direct.ValueOf(in.Name)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: GroupPreferencesetAssignments
	return out
}
func ReportConfig_GroupPreferenceSetAssignment_FromProto(mapCtx *direct.MapContext, in *pb.ReportConfig_GroupPreferenceSetAssignment) *krm.ReportConfig_GroupPreferenceSetAssignment {
	if in == nil {
		return nil
	}
	out := &krm.ReportConfig_GroupPreferenceSetAssignment{}
	out.Group = direct.LazyPtr(in.GetGroup())
	out.PreferenceSet = direct.LazyPtr(in.GetPreferenceSet())
	return out
}
func ReportConfig_GroupPreferenceSetAssignment_ToProto(mapCtx *direct.MapContext, in *krm.ReportConfig_GroupPreferenceSetAssignment) *pb.ReportConfig_GroupPreferenceSetAssignment {
	if in == nil {
		return nil
	}
	out := &pb.ReportConfig_GroupPreferenceSetAssignment{}
	out.Group = direct.ValueOf(in.Group)
	out.PreferenceSet = direct.ValueOf(in.PreferenceSet)
	return out
}
