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

package documentai

import (
	pb "cloud.google.com/go/documentai/apiv1beta3/documentaipb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/documentai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
func Dataset_FromProto(mapCtx *direct.MapContext, in *pb.Dataset) *krm.Dataset {
	if in == nil {
		return nil
	}
	out := &krm.Dataset{}
	out.GcsManagedConfig = Dataset_GCSManagedConfig_FromProto(mapCtx, in.GetGcsManagedConfig())
	out.DocumentWarehouseConfig = Dataset_DocumentWarehouseConfig_FromProto(mapCtx, in.GetDocumentWarehouseConfig())
	out.UnmanagedDatasetConfig = Dataset_UnmanagedDatasetConfig_FromProto(mapCtx, in.GetUnmanagedDatasetConfig())
	out.SpannerIndexingConfig = Dataset_SpannerIndexingConfig_FromProto(mapCtx, in.GetSpannerIndexingConfig())
	out.Name = direct.LazyPtr(in.GetName())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func Dataset_ToProto(mapCtx *direct.MapContext, in *krm.Dataset) *pb.Dataset {
	if in == nil {
		return nil
	}
	out := &pb.Dataset{}
	if oneof := Dataset_GCSManagedConfig_ToProto(mapCtx, in.GcsManagedConfig); oneof != nil {
		out.StorageSource = &pb.Dataset_GcsManagedConfig{GcsManagedConfig: oneof}
	}
	if oneof := Dataset_DocumentWarehouseConfig_ToProto(mapCtx, in.DocumentWarehouseConfig); oneof != nil {
		out.StorageSource = &pb.Dataset_DocumentWarehouseConfig_{DocumentWarehouseConfig: oneof}
	}
	if oneof := Dataset_UnmanagedDatasetConfig_ToProto(mapCtx, in.UnmanagedDatasetConfig); oneof != nil {
		out.StorageSource = &pb.Dataset_UnmanagedDatasetConfig_{UnmanagedDatasetConfig: oneof}
	}
	if oneof := Dataset_SpannerIndexingConfig_ToProto(mapCtx, in.SpannerIndexingConfig); oneof != nil {
		out.IndexingSource = &pb.Dataset_SpannerIndexingConfig_{SpannerIndexingConfig: oneof}
	}
	out.Name = direct.ValueOf(in.Name)
	out.State = direct.Enum_ToProto[pb.Dataset_State](mapCtx, in.State)
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func DatasetObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Dataset) *krm.DatasetObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DatasetObservedState{}
	// MISSING: GcsManagedConfig
	out.DocumentWarehouseConfig = Dataset_DocumentWarehouseConfigObservedState_FromProto(mapCtx, in.GetDocumentWarehouseConfig())
	// MISSING: UnmanagedDatasetConfig
	// MISSING: SpannerIndexingConfig
	// MISSING: Name
	// MISSING: State
	out.SatisfiesPzs = direct.LazyPtr(in.GetSatisfiesPzs())
	out.SatisfiesPzi = direct.LazyPtr(in.GetSatisfiesPzi())
	return out
}
func DatasetObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DatasetObservedState) *pb.Dataset {
	if in == nil {
		return nil
	}
	out := &pb.Dataset{}
	// MISSING: GcsManagedConfig
	if oneof := Dataset_DocumentWarehouseConfigObservedState_ToProto(mapCtx, in.DocumentWarehouseConfig); oneof != nil {
		out.StorageSource = &pb.Dataset_DocumentWarehouseConfig_{DocumentWarehouseConfig: oneof}
	}
	// MISSING: UnmanagedDatasetConfig
	// MISSING: SpannerIndexingConfig
	// MISSING: Name
	// MISSING: State
	out.SatisfiesPzs = direct.ValueOf(in.SatisfiesPzs)
	out.SatisfiesPzi = direct.ValueOf(in.SatisfiesPzi)
	return out
}
func Dataset_DocumentWarehouseConfig_FromProto(mapCtx *direct.MapContext, in *pb.Dataset_DocumentWarehouseConfig) *krm.Dataset_DocumentWarehouseConfig {
	if in == nil {
		return nil
	}
	out := &krm.Dataset_DocumentWarehouseConfig{}
	// MISSING: Collection
	// MISSING: Schema
	return out
}
func Dataset_DocumentWarehouseConfig_ToProto(mapCtx *direct.MapContext, in *krm.Dataset_DocumentWarehouseConfig) *pb.Dataset_DocumentWarehouseConfig {
	if in == nil {
		return nil
	}
	out := &pb.Dataset_DocumentWarehouseConfig{}
	// MISSING: Collection
	// MISSING: Schema
	return out
}
func Dataset_DocumentWarehouseConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Dataset_DocumentWarehouseConfig) *krm.Dataset_DocumentWarehouseConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Dataset_DocumentWarehouseConfigObservedState{}
	out.Collection = direct.LazyPtr(in.GetCollection())
	out.Schema = direct.LazyPtr(in.GetSchema())
	return out
}
func Dataset_DocumentWarehouseConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Dataset_DocumentWarehouseConfigObservedState) *pb.Dataset_DocumentWarehouseConfig {
	if in == nil {
		return nil
	}
	out := &pb.Dataset_DocumentWarehouseConfig{}
	out.Collection = direct.ValueOf(in.Collection)
	out.Schema = direct.ValueOf(in.Schema)
	return out
}
func Dataset_GCSManagedConfig_FromProto(mapCtx *direct.MapContext, in *pb.Dataset_GCSManagedConfig) *krm.Dataset_GCSManagedConfig {
	if in == nil {
		return nil
	}
	out := &krm.Dataset_GCSManagedConfig{}
	out.GcsPrefix = GcsPrefix_FromProto(mapCtx, in.GetGcsPrefix())
	return out
}
func Dataset_GCSManagedConfig_ToProto(mapCtx *direct.MapContext, in *krm.Dataset_GCSManagedConfig) *pb.Dataset_GCSManagedConfig {
	if in == nil {
		return nil
	}
	out := &pb.Dataset_GCSManagedConfig{}
	out.GcsPrefix = GcsPrefix_ToProto(mapCtx, in.GcsPrefix)
	return out
}
func Dataset_SpannerIndexingConfig_FromProto(mapCtx *direct.MapContext, in *pb.Dataset_SpannerIndexingConfig) *krm.Dataset_SpannerIndexingConfig {
	if in == nil {
		return nil
	}
	out := &krm.Dataset_SpannerIndexingConfig{}
	return out
}
func Dataset_SpannerIndexingConfig_ToProto(mapCtx *direct.MapContext, in *krm.Dataset_SpannerIndexingConfig) *pb.Dataset_SpannerIndexingConfig {
	if in == nil {
		return nil
	}
	out := &pb.Dataset_SpannerIndexingConfig{}
	return out
}
func Dataset_UnmanagedDatasetConfig_FromProto(mapCtx *direct.MapContext, in *pb.Dataset_UnmanagedDatasetConfig) *krm.Dataset_UnmanagedDatasetConfig {
	if in == nil {
		return nil
	}
	out := &krm.Dataset_UnmanagedDatasetConfig{}
	return out
}
func Dataset_UnmanagedDatasetConfig_ToProto(mapCtx *direct.MapContext, in *krm.Dataset_UnmanagedDatasetConfig) *pb.Dataset_UnmanagedDatasetConfig {
	if in == nil {
		return nil
	}
	out := &pb.Dataset_UnmanagedDatasetConfig{}
	return out
}
func DocumentaiDatasetObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Dataset) *krm.DocumentaiDatasetObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DocumentaiDatasetObservedState{}
	// MISSING: GcsManagedConfig
	// MISSING: DocumentWarehouseConfig
	// MISSING: UnmanagedDatasetConfig
	// MISSING: SpannerIndexingConfig
	// MISSING: Name
	// MISSING: State
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func DocumentaiDatasetObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DocumentaiDatasetObservedState) *pb.Dataset {
	if in == nil {
		return nil
	}
	out := &pb.Dataset{}
	// MISSING: GcsManagedConfig
	// MISSING: DocumentWarehouseConfig
	// MISSING: UnmanagedDatasetConfig
	// MISSING: SpannerIndexingConfig
	// MISSING: Name
	// MISSING: State
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func DocumentaiDatasetSpec_FromProto(mapCtx *direct.MapContext, in *pb.Dataset) *krm.DocumentaiDatasetSpec {
	if in == nil {
		return nil
	}
	out := &krm.DocumentaiDatasetSpec{}
	// MISSING: GcsManagedConfig
	// MISSING: DocumentWarehouseConfig
	// MISSING: UnmanagedDatasetConfig
	// MISSING: SpannerIndexingConfig
	// MISSING: Name
	// MISSING: State
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func DocumentaiDatasetSpec_ToProto(mapCtx *direct.MapContext, in *krm.DocumentaiDatasetSpec) *pb.Dataset {
	if in == nil {
		return nil
	}
	out := &pb.Dataset{}
	// MISSING: GcsManagedConfig
	// MISSING: DocumentWarehouseConfig
	// MISSING: UnmanagedDatasetConfig
	// MISSING: SpannerIndexingConfig
	// MISSING: Name
	// MISSING: State
	// MISSING: SatisfiesPzs
	// MISSING: SatisfiesPzi
	return out
}
func GcsPrefix_FromProto(mapCtx *direct.MapContext, in *pb.GcsPrefix) *krm.GcsPrefix {
	if in == nil {
		return nil
	}
	out := &krm.GcsPrefix{}
	out.GcsURIPrefix = direct.LazyPtr(in.GetGcsUriPrefix())
	return out
}
func GcsPrefix_ToProto(mapCtx *direct.MapContext, in *krm.GcsPrefix) *pb.GcsPrefix {
	if in == nil {
		return nil
	}
	out := &pb.GcsPrefix{}
	out.GcsUriPrefix = direct.ValueOf(in.GcsURIPrefix)
	return out
}
