// Copyright 2026 Google LLC
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

package featuregroup

import (
	pb "cloud.google.com/go/aiplatform/apiv1beta1/aiplatformpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/vertexai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func BigQuerySource_FromProto(mapCtx *direct.MapContext, in *pb.BigQuerySource) *krm.BigQuerySource {
	if in == nil {
		return nil
	}
	out := &krm.BigQuerySource{}
	out.InputURI = direct.LazyPtr(in.GetInputUri())
	return out
}
func BigQuerySource_ToProto(mapCtx *direct.MapContext, in *krm.BigQuerySource) *pb.BigQuerySource {
	if in == nil {
		return nil
	}
	out := &pb.BigQuerySource{}
	out.InputUri = direct.ValueOf(in.InputURI)
	return out
}
func FeatureGroup_BigQuery_FromProto(mapCtx *direct.MapContext, in *pb.FeatureGroup_BigQuery) *krm.FeatureGroup_BigQuery {
	if in == nil {
		return nil
	}
	out := &krm.FeatureGroup_BigQuery{}
	out.BigQuerySource = BigQuerySource_FromProto(mapCtx, in.GetBigQuerySource())
	out.EntityIDColumns = in.GetEntityIdColumns()
	out.StaticDataSource = direct.LazyPtr(in.GetStaticDataSource())
	out.TimeSeries = FeatureGroup_BigQuery_TimeSeries_FromProto(mapCtx, in.GetTimeSeries())
	out.Dense = direct.LazyPtr(in.GetDense())
	return out
}
func FeatureGroup_BigQuery_ToProto(mapCtx *direct.MapContext, in *krm.FeatureGroup_BigQuery) *pb.FeatureGroup_BigQuery {
	if in == nil {
		return nil
	}
	out := &pb.FeatureGroup_BigQuery{}
	out.BigQuerySource = BigQuerySource_ToProto(mapCtx, in.BigQuerySource)
	out.EntityIdColumns = in.EntityIDColumns
	out.StaticDataSource = direct.ValueOf(in.StaticDataSource)
	out.TimeSeries = FeatureGroup_BigQuery_TimeSeries_ToProto(mapCtx, in.TimeSeries)
	out.Dense = direct.ValueOf(in.Dense)
	return out
}
func FeatureGroup_BigQuery_TimeSeries_FromProto(mapCtx *direct.MapContext, in *pb.FeatureGroup_BigQuery_TimeSeries) *krm.FeatureGroup_BigQuery_TimeSeries {
	if in == nil {
		return nil
	}
	out := &krm.FeatureGroup_BigQuery_TimeSeries{}
	out.TimestampColumn = direct.LazyPtr(in.GetTimestampColumn())
	return out
}
func FeatureGroup_BigQuery_TimeSeries_ToProto(mapCtx *direct.MapContext, in *krm.FeatureGroup_BigQuery_TimeSeries) *pb.FeatureGroup_BigQuery_TimeSeries {
	if in == nil {
		return nil
	}
	out := &pb.FeatureGroup_BigQuery_TimeSeries{}
	out.TimestampColumn = direct.ValueOf(in.TimestampColumn)
	return out
}
func VertexAIFeatureGroupObservedState_FromProto(mapCtx *direct.MapContext, in *pb.FeatureGroup) *krm.VertexAIFeatureGroupObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VertexAIFeatureGroupObservedState{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.ServiceAccountEmail = direct.LazyPtr(in.GetServiceAccountEmail())
	return out
}
func VertexAIFeatureGroupObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VertexAIFeatureGroupObservedState) *pb.FeatureGroup {
	if in == nil {
		return nil
	}
	out := &pb.FeatureGroup{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.ServiceAccountEmail = direct.ValueOf(in.ServiceAccountEmail)
	return out
}
func VertexAIFeatureGroupSpec_FromProto(mapCtx *direct.MapContext, in *pb.FeatureGroup) *krm.VertexAIFeatureGroupSpec {
	if in == nil {
		return nil
	}
	out := &krm.VertexAIFeatureGroupSpec{}
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Labels = in.GetLabels()
	out.BigQuery = FeatureGroup_BigQuery_FromProto(mapCtx, in.GetBigQuery())
	out.ServiceAgentType = direct.Enum_FromProto(mapCtx, in.GetServiceAgentType())
	return out
}
func VertexAIFeatureGroupSpec_ToProto(mapCtx *direct.MapContext, in *krm.VertexAIFeatureGroupSpec) *pb.FeatureGroup {
	if in == nil {
		return nil
	}
	out := &pb.FeatureGroup{}
	out.Description = direct.ValueOf(in.Description)
	out.Labels = in.Labels
	if oneof := FeatureGroup_BigQuery_ToProto(mapCtx, in.BigQuery); oneof != nil {
		out.Source = &pb.FeatureGroup_BigQuery_{BigQuery: oneof}
	}
	out.ServiceAgentType = direct.Enum_ToProto[pb.FeatureGroup_ServiceAgentType](mapCtx, in.ServiceAgentType)
	return out
}
