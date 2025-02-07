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

package analytics

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "google.golang.org/genproto/googleapis/analytics/data/v1beta"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/analytics/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func AnalyticsAudienceExportObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AudienceExport) *krm.AnalyticsAudienceExportObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AnalyticsAudienceExportObservedState{}
	// MISSING: Name
	// MISSING: Audience
	// MISSING: AudienceDisplayName
	// MISSING: Dimensions
	// MISSING: State
	// MISSING: BeginCreatingTime
	// MISSING: CreationQuotaTokensCharged
	// MISSING: RowCount
	// MISSING: ErrorMessage
	// MISSING: PercentageCompleted
	return out
}
func AnalyticsAudienceExportObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AnalyticsAudienceExportObservedState) *pb.AudienceExport {
	if in == nil {
		return nil
	}
	out := &pb.AudienceExport{}
	// MISSING: Name
	// MISSING: Audience
	// MISSING: AudienceDisplayName
	// MISSING: Dimensions
	// MISSING: State
	// MISSING: BeginCreatingTime
	// MISSING: CreationQuotaTokensCharged
	// MISSING: RowCount
	// MISSING: ErrorMessage
	// MISSING: PercentageCompleted
	return out
}
func AnalyticsAudienceExportSpec_FromProto(mapCtx *direct.MapContext, in *pb.AudienceExport) *krm.AnalyticsAudienceExportSpec {
	if in == nil {
		return nil
	}
	out := &krm.AnalyticsAudienceExportSpec{}
	// MISSING: Name
	// MISSING: Audience
	// MISSING: AudienceDisplayName
	// MISSING: Dimensions
	// MISSING: State
	// MISSING: BeginCreatingTime
	// MISSING: CreationQuotaTokensCharged
	// MISSING: RowCount
	// MISSING: ErrorMessage
	// MISSING: PercentageCompleted
	return out
}
func AnalyticsAudienceExportSpec_ToProto(mapCtx *direct.MapContext, in *krm.AnalyticsAudienceExportSpec) *pb.AudienceExport {
	if in == nil {
		return nil
	}
	out := &pb.AudienceExport{}
	// MISSING: Name
	// MISSING: Audience
	// MISSING: AudienceDisplayName
	// MISSING: Dimensions
	// MISSING: State
	// MISSING: BeginCreatingTime
	// MISSING: CreationQuotaTokensCharged
	// MISSING: RowCount
	// MISSING: ErrorMessage
	// MISSING: PercentageCompleted
	return out
}
func AudienceDimension_FromProto(mapCtx *direct.MapContext, in *pb.AudienceDimension) *krm.AudienceDimension {
	if in == nil {
		return nil
	}
	out := &krm.AudienceDimension{}
	out.DimensionName = direct.LazyPtr(in.GetDimensionName())
	return out
}
func AudienceDimension_ToProto(mapCtx *direct.MapContext, in *krm.AudienceDimension) *pb.AudienceDimension {
	if in == nil {
		return nil
	}
	out := &pb.AudienceDimension{}
	out.DimensionName = direct.ValueOf(in.DimensionName)
	return out
}
func AudienceExport_FromProto(mapCtx *direct.MapContext, in *pb.AudienceExport) *krm.AudienceExport {
	if in == nil {
		return nil
	}
	out := &krm.AudienceExport{}
	// MISSING: Name
	out.Audience = direct.LazyPtr(in.GetAudience())
	// MISSING: AudienceDisplayName
	out.Dimensions = direct.Slice_FromProto(mapCtx, in.Dimensions, AudienceDimension_FromProto)
	// MISSING: State
	// MISSING: BeginCreatingTime
	// MISSING: CreationQuotaTokensCharged
	// MISSING: RowCount
	// MISSING: ErrorMessage
	// MISSING: PercentageCompleted
	return out
}
func AudienceExport_ToProto(mapCtx *direct.MapContext, in *krm.AudienceExport) *pb.AudienceExport {
	if in == nil {
		return nil
	}
	out := &pb.AudienceExport{}
	// MISSING: Name
	out.Audience = direct.ValueOf(in.Audience)
	// MISSING: AudienceDisplayName
	out.Dimensions = direct.Slice_ToProto(mapCtx, in.Dimensions, AudienceDimension_ToProto)
	// MISSING: State
	// MISSING: BeginCreatingTime
	// MISSING: CreationQuotaTokensCharged
	// MISSING: RowCount
	// MISSING: ErrorMessage
	// MISSING: PercentageCompleted
	return out
}
func AudienceExportObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AudienceExport) *krm.AudienceExportObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AudienceExportObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: Audience
	out.AudienceDisplayName = direct.LazyPtr(in.GetAudienceDisplayName())
	// MISSING: Dimensions
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.BeginCreatingTime = direct.StringTimestamp_FromProto(mapCtx, in.GetBeginCreatingTime())
	out.CreationQuotaTokensCharged = direct.LazyPtr(in.GetCreationQuotaTokensCharged())
	out.RowCount = in.RowCount
	out.ErrorMessage = in.ErrorMessage
	out.PercentageCompleted = in.PercentageCompleted
	return out
}
func AudienceExportObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AudienceExportObservedState) *pb.AudienceExport {
	if in == nil {
		return nil
	}
	out := &pb.AudienceExport{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: Audience
	out.AudienceDisplayName = direct.ValueOf(in.AudienceDisplayName)
	// MISSING: Dimensions
	if oneof := AudienceExportObservedState_State_ToProto(mapCtx, in.State); oneof != nil {
		out.State = oneof
	}
	if oneof := direct.StringTimestamp_ToProto(mapCtx, in.BeginCreatingTime); oneof != nil {
		out.BeginCreatingTime = &pb.AudienceExport_BeginCreatingTime{BeginCreatingTime: oneof}
	}
	out.CreationQuotaTokensCharged = direct.ValueOf(in.CreationQuotaTokensCharged)
	out.RowCount = in.RowCount
	out.ErrorMessage = in.ErrorMessage
	out.PercentageCompleted = in.PercentageCompleted
	return out
}
