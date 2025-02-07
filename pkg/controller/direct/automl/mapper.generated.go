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

package automl

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/automl/apiv1beta1/automlpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/automl/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func AutomlTableSpecObservedState_FromProto(mapCtx *direct.MapContext, in *pb.TableSpec) *krm.AutomlTableSpecObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AutomlTableSpecObservedState{}
	// MISSING: Name
	// MISSING: TimeColumnSpecID
	// MISSING: RowCount
	// MISSING: ValidRowCount
	// MISSING: ColumnCount
	// MISSING: InputConfigs
	// MISSING: Etag
	return out
}
func AutomlTableSpecObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AutomlTableSpecObservedState) *pb.TableSpec {
	if in == nil {
		return nil
	}
	out := &pb.TableSpec{}
	// MISSING: Name
	// MISSING: TimeColumnSpecID
	// MISSING: RowCount
	// MISSING: ValidRowCount
	// MISSING: ColumnCount
	// MISSING: InputConfigs
	// MISSING: Etag
	return out
}
func AutomlTableSpecSpec_FromProto(mapCtx *direct.MapContext, in *pb.TableSpec) *krm.AutomlTableSpecSpec {
	if in == nil {
		return nil
	}
	out := &krm.AutomlTableSpecSpec{}
	// MISSING: Name
	// MISSING: TimeColumnSpecID
	// MISSING: RowCount
	// MISSING: ValidRowCount
	// MISSING: ColumnCount
	// MISSING: InputConfigs
	// MISSING: Etag
	return out
}
func AutomlTableSpecSpec_ToProto(mapCtx *direct.MapContext, in *krm.AutomlTableSpecSpec) *pb.TableSpec {
	if in == nil {
		return nil
	}
	out := &pb.TableSpec{}
	// MISSING: Name
	// MISSING: TimeColumnSpecID
	// MISSING: RowCount
	// MISSING: ValidRowCount
	// MISSING: ColumnCount
	// MISSING: InputConfigs
	// MISSING: Etag
	return out
}
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
func GcsSource_FromProto(mapCtx *direct.MapContext, in *pb.GcsSource) *krm.GcsSource {
	if in == nil {
		return nil
	}
	out := &krm.GcsSource{}
	out.InputUris = in.InputUris
	return out
}
func GcsSource_ToProto(mapCtx *direct.MapContext, in *krm.GcsSource) *pb.GcsSource {
	if in == nil {
		return nil
	}
	out := &pb.GcsSource{}
	out.InputUris = in.InputUris
	return out
}
func InputConfig_FromProto(mapCtx *direct.MapContext, in *pb.InputConfig) *krm.InputConfig {
	if in == nil {
		return nil
	}
	out := &krm.InputConfig{}
	out.GcsSource = GcsSource_FromProto(mapCtx, in.GetGcsSource())
	out.BigquerySource = BigQuerySource_FromProto(mapCtx, in.GetBigquerySource())
	out.Params = in.Params
	return out
}
func InputConfig_ToProto(mapCtx *direct.MapContext, in *krm.InputConfig) *pb.InputConfig {
	if in == nil {
		return nil
	}
	out := &pb.InputConfig{}
	if oneof := GcsSource_ToProto(mapCtx, in.GcsSource); oneof != nil {
		out.Source = &pb.InputConfig_GcsSource{GcsSource: oneof}
	}
	if oneof := BigQuerySource_ToProto(mapCtx, in.BigquerySource); oneof != nil {
		out.Source = &pb.InputConfig_BigquerySource{BigquerySource: oneof}
	}
	out.Params = in.Params
	return out
}
func TableSpec_FromProto(mapCtx *direct.MapContext, in *pb.TableSpec) *krm.TableSpec {
	if in == nil {
		return nil
	}
	out := &krm.TableSpec{}
	out.Name = direct.LazyPtr(in.GetName())
	out.TimeColumnSpecID = direct.LazyPtr(in.GetTimeColumnSpecId())
	out.RowCount = direct.LazyPtr(in.GetRowCount())
	out.ValidRowCount = direct.LazyPtr(in.GetValidRowCount())
	out.ColumnCount = direct.LazyPtr(in.GetColumnCount())
	out.InputConfigs = direct.Slice_FromProto(mapCtx, in.InputConfigs, InputConfig_FromProto)
	out.Etag = direct.LazyPtr(in.GetEtag())
	return out
}
func TableSpec_ToProto(mapCtx *direct.MapContext, in *krm.TableSpec) *pb.TableSpec {
	if in == nil {
		return nil
	}
	out := &pb.TableSpec{}
	out.Name = direct.ValueOf(in.Name)
	out.TimeColumnSpecId = direct.ValueOf(in.TimeColumnSpecID)
	out.RowCount = direct.ValueOf(in.RowCount)
	out.ValidRowCount = direct.ValueOf(in.ValidRowCount)
	out.ColumnCount = direct.ValueOf(in.ColumnCount)
	out.InputConfigs = direct.Slice_ToProto(mapCtx, in.InputConfigs, InputConfig_ToProto)
	out.Etag = direct.ValueOf(in.Etag)
	return out
}
