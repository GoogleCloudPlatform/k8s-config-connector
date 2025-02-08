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

package storageinsights

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/storageinsights/apiv1/storageinsightspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/storageinsights/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func ReportDetail_FromProto(mapCtx *direct.MapContext, in *pb.ReportDetail) *krm.ReportDetail {
	if in == nil {
		return nil
	}
	out := &krm.ReportDetail{}
	out.Name = direct.LazyPtr(in.GetName())
	out.SnapshotTime = direct.StringTimestamp_FromProto(mapCtx, in.GetSnapshotTime())
	out.ReportPathPrefix = direct.LazyPtr(in.GetReportPathPrefix())
	out.ShardsCount = direct.LazyPtr(in.GetShardsCount())
	out.Status = Status_FromProto(mapCtx, in.GetStatus())
	out.Labels = in.Labels
	out.TargetDatetime = DateTime_FromProto(mapCtx, in.GetTargetDatetime())
	out.ReportMetrics = ReportDetail_Metrics_FromProto(mapCtx, in.GetReportMetrics())
	return out
}
func ReportDetail_ToProto(mapCtx *direct.MapContext, in *krm.ReportDetail) *pb.ReportDetail {
	if in == nil {
		return nil
	}
	out := &pb.ReportDetail{}
	out.Name = direct.ValueOf(in.Name)
	out.SnapshotTime = direct.StringTimestamp_ToProto(mapCtx, in.SnapshotTime)
	out.ReportPathPrefix = direct.ValueOf(in.ReportPathPrefix)
	out.ShardsCount = direct.ValueOf(in.ShardsCount)
	out.Status = Status_ToProto(mapCtx, in.Status)
	out.Labels = in.Labels
	out.TargetDatetime = DateTime_ToProto(mapCtx, in.TargetDatetime)
	out.ReportMetrics = ReportDetail_Metrics_ToProto(mapCtx, in.ReportMetrics)
	return out
}
func ReportDetail_Metrics_FromProto(mapCtx *direct.MapContext, in *pb.ReportDetail_Metrics) *krm.ReportDetail_Metrics {
	if in == nil {
		return nil
	}
	out := &krm.ReportDetail_Metrics{}
	out.ProcessedRecordsCount = direct.LazyPtr(in.GetProcessedRecordsCount())
	return out
}
func ReportDetail_Metrics_ToProto(mapCtx *direct.MapContext, in *krm.ReportDetail_Metrics) *pb.ReportDetail_Metrics {
	if in == nil {
		return nil
	}
	out := &pb.ReportDetail_Metrics{}
	out.ProcessedRecordsCount = direct.ValueOf(in.ProcessedRecordsCount)
	return out
}
func StorageinsightsReportDetailObservedState_FromProto(mapCtx *direct.MapContext, in *pb.ReportDetail) *krm.StorageinsightsReportDetailObservedState {
	if in == nil {
		return nil
	}
	out := &krm.StorageinsightsReportDetailObservedState{}
	// MISSING: Name
	// MISSING: SnapshotTime
	// MISSING: ReportPathPrefix
	// MISSING: ShardsCount
	// MISSING: Status
	// MISSING: Labels
	// MISSING: TargetDatetime
	// MISSING: ReportMetrics
	return out
}
func StorageinsightsReportDetailObservedState_ToProto(mapCtx *direct.MapContext, in *krm.StorageinsightsReportDetailObservedState) *pb.ReportDetail {
	if in == nil {
		return nil
	}
	out := &pb.ReportDetail{}
	// MISSING: Name
	// MISSING: SnapshotTime
	// MISSING: ReportPathPrefix
	// MISSING: ShardsCount
	// MISSING: Status
	// MISSING: Labels
	// MISSING: TargetDatetime
	// MISSING: ReportMetrics
	return out
}
func StorageinsightsReportDetailSpec_FromProto(mapCtx *direct.MapContext, in *pb.ReportDetail) *krm.StorageinsightsReportDetailSpec {
	if in == nil {
		return nil
	}
	out := &krm.StorageinsightsReportDetailSpec{}
	// MISSING: Name
	// MISSING: SnapshotTime
	// MISSING: ReportPathPrefix
	// MISSING: ShardsCount
	// MISSING: Status
	// MISSING: Labels
	// MISSING: TargetDatetime
	// MISSING: ReportMetrics
	return out
}
func StorageinsightsReportDetailSpec_ToProto(mapCtx *direct.MapContext, in *krm.StorageinsightsReportDetailSpec) *pb.ReportDetail {
	if in == nil {
		return nil
	}
	out := &pb.ReportDetail{}
	// MISSING: Name
	// MISSING: SnapshotTime
	// MISSING: ReportPathPrefix
	// MISSING: ShardsCount
	// MISSING: Status
	// MISSING: Labels
	// MISSING: TargetDatetime
	// MISSING: ReportMetrics
	return out
}
