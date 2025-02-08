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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/securitycenter/apiv1/securitycenterpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/securitycenter/v1alpha1"
)
func BigQueryExport_FromProto(mapCtx *direct.MapContext, in *pb.BigQueryExport) *krm.BigQueryExport {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryExport{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Filter = direct.LazyPtr(in.GetFilter())
	out.Dataset = direct.LazyPtr(in.GetDataset())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: MostRecentEditor
	// MISSING: Principal
	return out
}
func BigQueryExport_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryExport) *pb.BigQueryExport {
	if in == nil {
		return nil
	}
	out := &pb.BigQueryExport{}
	out.Name = direct.ValueOf(in.Name)
	out.Description = direct.ValueOf(in.Description)
	out.Filter = direct.ValueOf(in.Filter)
	out.Dataset = direct.ValueOf(in.Dataset)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: MostRecentEditor
	// MISSING: Principal
	return out
}
func BigQueryExportObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BigQueryExport) *krm.BigQueryExportObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryExportObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Filter
	// MISSING: Dataset
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.MostRecentEditor = direct.LazyPtr(in.GetMostRecentEditor())
	out.Principal = direct.LazyPtr(in.GetPrincipal())
	return out
}
func BigQueryExportObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryExportObservedState) *pb.BigQueryExport {
	if in == nil {
		return nil
	}
	out := &pb.BigQueryExport{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Filter
	// MISSING: Dataset
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.MostRecentEditor = direct.ValueOf(in.MostRecentEditor)
	out.Principal = direct.ValueOf(in.Principal)
	return out
}
func SecuritycenterBigQueryExportObservedState_FromProto(mapCtx *direct.MapContext, in *pb.BigQueryExport) *krm.SecuritycenterBigQueryExportObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SecuritycenterBigQueryExportObservedState{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Filter
	// MISSING: Dataset
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: MostRecentEditor
	// MISSING: Principal
	return out
}
func SecuritycenterBigQueryExportObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SecuritycenterBigQueryExportObservedState) *pb.BigQueryExport {
	if in == nil {
		return nil
	}
	out := &pb.BigQueryExport{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Filter
	// MISSING: Dataset
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: MostRecentEditor
	// MISSING: Principal
	return out
}
func SecuritycenterBigQueryExportSpec_FromProto(mapCtx *direct.MapContext, in *pb.BigQueryExport) *krm.SecuritycenterBigQueryExportSpec {
	if in == nil {
		return nil
	}
	out := &krm.SecuritycenterBigQueryExportSpec{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Filter
	// MISSING: Dataset
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: MostRecentEditor
	// MISSING: Principal
	return out
}
func SecuritycenterBigQueryExportSpec_ToProto(mapCtx *direct.MapContext, in *krm.SecuritycenterBigQueryExportSpec) *pb.BigQueryExport {
	if in == nil {
		return nil
	}
	out := &pb.BigQueryExport{}
	// MISSING: Name
	// MISSING: Description
	// MISSING: Filter
	// MISSING: Dataset
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: MostRecentEditor
	// MISSING: Principal
	return out
}
