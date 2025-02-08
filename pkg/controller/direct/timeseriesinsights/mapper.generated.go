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

package timeseriesinsights

import (
	pb "cloud.google.com/go/timeseriesinsights/apiv1/timeseriesinsightspb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/timeseriesinsights/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func BigqueryMapping_FromProto(mapCtx *direct.MapContext, in *pb.BigqueryMapping) *krm.BigqueryMapping {
	if in == nil {
		return nil
	}
	out := &krm.BigqueryMapping{}
	out.TimestampColumn = direct.LazyPtr(in.GetTimestampColumn())
	out.GroupIDColumn = direct.LazyPtr(in.GetGroupIdColumn())
	out.DimensionColumn = in.DimensionColumn
	return out
}
func BigqueryMapping_ToProto(mapCtx *direct.MapContext, in *krm.BigqueryMapping) *pb.BigqueryMapping {
	if in == nil {
		return nil
	}
	out := &pb.BigqueryMapping{}
	out.TimestampColumn = direct.ValueOf(in.TimestampColumn)
	out.GroupIdColumn = direct.ValueOf(in.GroupIDColumn)
	out.DimensionColumn = in.DimensionColumn
	return out
}
func DataSet_FromProto(mapCtx *direct.MapContext, in *pb.DataSet) *krm.DataSet {
	if in == nil {
		return nil
	}
	out := &krm.DataSet{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DataNames = in.DataNames
	out.DataSources = direct.Slice_FromProto(mapCtx, in.DataSources, DataSource_FromProto)
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Status = Status_FromProto(mapCtx, in.GetStatus())
	out.Ttl = direct.StringDuration_FromProto(mapCtx, in.GetTtl())
	return out
}
func DataSet_ToProto(mapCtx *direct.MapContext, in *krm.DataSet) *pb.DataSet {
	if in == nil {
		return nil
	}
	out := &pb.DataSet{}
	out.Name = direct.ValueOf(in.Name)
	out.DataNames = in.DataNames
	out.DataSources = direct.Slice_ToProto(mapCtx, in.DataSources, DataSource_ToProto)
	out.State = direct.Enum_ToProto[pb.DataSet_State](mapCtx, in.State)
	out.Status = Status_ToProto(mapCtx, in.Status)
	out.Ttl = direct.StringDuration_ToProto(mapCtx, in.Ttl)
	return out
}
func DataSource_FromProto(mapCtx *direct.MapContext, in *pb.DataSource) *krm.DataSource {
	if in == nil {
		return nil
	}
	out := &krm.DataSource{}
	out.URI = direct.LazyPtr(in.GetUri())
	out.BqMapping = BigqueryMapping_FromProto(mapCtx, in.GetBqMapping())
	return out
}
func DataSource_ToProto(mapCtx *direct.MapContext, in *krm.DataSource) *pb.DataSource {
	if in == nil {
		return nil
	}
	out := &pb.DataSource{}
	out.Uri = direct.ValueOf(in.URI)
	out.BqMapping = BigqueryMapping_ToProto(mapCtx, in.BqMapping)
	return out
}
func TimeseriesinsightsDataSetObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataSet) *krm.TimeseriesinsightsDataSetObservedState {
	if in == nil {
		return nil
	}
	out := &krm.TimeseriesinsightsDataSetObservedState{}
	// MISSING: Name
	// MISSING: DataNames
	// MISSING: DataSources
	// MISSING: State
	// MISSING: Status
	// MISSING: Ttl
	return out
}
func TimeseriesinsightsDataSetObservedState_ToProto(mapCtx *direct.MapContext, in *krm.TimeseriesinsightsDataSetObservedState) *pb.DataSet {
	if in == nil {
		return nil
	}
	out := &pb.DataSet{}
	// MISSING: Name
	// MISSING: DataNames
	// MISSING: DataSources
	// MISSING: State
	// MISSING: Status
	// MISSING: Ttl
	return out
}
func TimeseriesinsightsDataSetSpec_FromProto(mapCtx *direct.MapContext, in *pb.DataSet) *krm.TimeseriesinsightsDataSetSpec {
	if in == nil {
		return nil
	}
	out := &krm.TimeseriesinsightsDataSetSpec{}
	// MISSING: Name
	// MISSING: DataNames
	// MISSING: DataSources
	// MISSING: State
	// MISSING: Status
	// MISSING: Ttl
	return out
}
func TimeseriesinsightsDataSetSpec_ToProto(mapCtx *direct.MapContext, in *krm.TimeseriesinsightsDataSetSpec) *pb.DataSet {
	if in == nil {
		return nil
	}
	out := &pb.DataSet{}
	// MISSING: Name
	// MISSING: DataNames
	// MISSING: DataSources
	// MISSING: State
	// MISSING: Status
	// MISSING: Ttl
	return out
}
