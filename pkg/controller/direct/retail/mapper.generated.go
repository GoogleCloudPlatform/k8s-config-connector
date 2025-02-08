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

package retail

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/retail/apiv2/retailpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/retail/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func BigQuerySource_FromProto(mapCtx *direct.MapContext, in *pb.BigQuerySource) *krm.BigQuerySource {
	if in == nil {
		return nil
	}
	out := &krm.BigQuerySource{}
	out.PartitionDate = Date_FromProto(mapCtx, in.GetPartitionDate())
	out.ProjectID = direct.LazyPtr(in.GetProjectId())
	out.DatasetID = direct.LazyPtr(in.GetDatasetId())
	out.TableID = direct.LazyPtr(in.GetTableId())
	out.GcsStagingDir = direct.LazyPtr(in.GetGcsStagingDir())
	out.DataSchema = direct.LazyPtr(in.GetDataSchema())
	return out
}
func BigQuerySource_ToProto(mapCtx *direct.MapContext, in *krm.BigQuerySource) *pb.BigQuerySource {
	if in == nil {
		return nil
	}
	out := &pb.BigQuerySource{}
	if oneof := Date_ToProto(mapCtx, in.PartitionDate); oneof != nil {
		out.Partition = &pb.BigQuerySource_PartitionDate{PartitionDate: oneof}
	}
	out.ProjectId = direct.ValueOf(in.ProjectID)
	out.DatasetId = direct.ValueOf(in.DatasetID)
	out.TableId = direct.ValueOf(in.TableID)
	out.GcsStagingDir = direct.ValueOf(in.GcsStagingDir)
	out.DataSchema = direct.ValueOf(in.DataSchema)
	return out
}
func CompletionConfig_FromProto(mapCtx *direct.MapContext, in *pb.CompletionConfig) *krm.CompletionConfig {
	if in == nil {
		return nil
	}
	out := &krm.CompletionConfig{}
	out.Name = direct.LazyPtr(in.GetName())
	out.MatchingOrder = direct.LazyPtr(in.GetMatchingOrder())
	out.MaxSuggestions = direct.LazyPtr(in.GetMaxSuggestions())
	out.MinPrefixLength = direct.LazyPtr(in.GetMinPrefixLength())
	out.AutoLearning = direct.LazyPtr(in.GetAutoLearning())
	// MISSING: SuggestionsInputConfig
	// MISSING: LastSuggestionsImportOperation
	// MISSING: DenylistInputConfig
	// MISSING: LastDenylistImportOperation
	// MISSING: AllowlistInputConfig
	// MISSING: LastAllowlistImportOperation
	return out
}
func CompletionConfig_ToProto(mapCtx *direct.MapContext, in *krm.CompletionConfig) *pb.CompletionConfig {
	if in == nil {
		return nil
	}
	out := &pb.CompletionConfig{}
	out.Name = direct.ValueOf(in.Name)
	out.MatchingOrder = direct.ValueOf(in.MatchingOrder)
	out.MaxSuggestions = direct.ValueOf(in.MaxSuggestions)
	out.MinPrefixLength = direct.ValueOf(in.MinPrefixLength)
	out.AutoLearning = direct.ValueOf(in.AutoLearning)
	// MISSING: SuggestionsInputConfig
	// MISSING: LastSuggestionsImportOperation
	// MISSING: DenylistInputConfig
	// MISSING: LastDenylistImportOperation
	// MISSING: AllowlistInputConfig
	// MISSING: LastAllowlistImportOperation
	return out
}
func CompletionConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CompletionConfig) *krm.CompletionConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CompletionConfigObservedState{}
	// MISSING: Name
	// MISSING: MatchingOrder
	// MISSING: MaxSuggestions
	// MISSING: MinPrefixLength
	// MISSING: AutoLearning
	out.SuggestionsInputConfig = CompletionDataInputConfig_FromProto(mapCtx, in.GetSuggestionsInputConfig())
	out.LastSuggestionsImportOperation = direct.LazyPtr(in.GetLastSuggestionsImportOperation())
	out.DenylistInputConfig = CompletionDataInputConfig_FromProto(mapCtx, in.GetDenylistInputConfig())
	out.LastDenylistImportOperation = direct.LazyPtr(in.GetLastDenylistImportOperation())
	out.AllowlistInputConfig = CompletionDataInputConfig_FromProto(mapCtx, in.GetAllowlistInputConfig())
	out.LastAllowlistImportOperation = direct.LazyPtr(in.GetLastAllowlistImportOperation())
	return out
}
func CompletionConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CompletionConfigObservedState) *pb.CompletionConfig {
	if in == nil {
		return nil
	}
	out := &pb.CompletionConfig{}
	// MISSING: Name
	// MISSING: MatchingOrder
	// MISSING: MaxSuggestions
	// MISSING: MinPrefixLength
	// MISSING: AutoLearning
	out.SuggestionsInputConfig = CompletionDataInputConfig_ToProto(mapCtx, in.SuggestionsInputConfig)
	out.LastSuggestionsImportOperation = direct.ValueOf(in.LastSuggestionsImportOperation)
	out.DenylistInputConfig = CompletionDataInputConfig_ToProto(mapCtx, in.DenylistInputConfig)
	out.LastDenylistImportOperation = direct.ValueOf(in.LastDenylistImportOperation)
	out.AllowlistInputConfig = CompletionDataInputConfig_ToProto(mapCtx, in.AllowlistInputConfig)
	out.LastAllowlistImportOperation = direct.ValueOf(in.LastAllowlistImportOperation)
	return out
}
func CompletionDataInputConfig_FromProto(mapCtx *direct.MapContext, in *pb.CompletionDataInputConfig) *krm.CompletionDataInputConfig {
	if in == nil {
		return nil
	}
	out := &krm.CompletionDataInputConfig{}
	out.BigQuerySource = BigQuerySource_FromProto(mapCtx, in.GetBigQuerySource())
	return out
}
func CompletionDataInputConfig_ToProto(mapCtx *direct.MapContext, in *krm.CompletionDataInputConfig) *pb.CompletionDataInputConfig {
	if in == nil {
		return nil
	}
	out := &pb.CompletionDataInputConfig{}
	if oneof := BigQuerySource_ToProto(mapCtx, in.BigQuerySource); oneof != nil {
		out.Source = &pb.CompletionDataInputConfig_BigQuerySource{BigQuerySource: oneof}
	}
	return out
}
func RetailCompletionConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.CompletionConfig) *krm.RetailCompletionConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.RetailCompletionConfigObservedState{}
	// MISSING: Name
	// MISSING: MatchingOrder
	// MISSING: MaxSuggestions
	// MISSING: MinPrefixLength
	// MISSING: AutoLearning
	// MISSING: SuggestionsInputConfig
	// MISSING: LastSuggestionsImportOperation
	// MISSING: DenylistInputConfig
	// MISSING: LastDenylistImportOperation
	// MISSING: AllowlistInputConfig
	// MISSING: LastAllowlistImportOperation
	return out
}
func RetailCompletionConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.RetailCompletionConfigObservedState) *pb.CompletionConfig {
	if in == nil {
		return nil
	}
	out := &pb.CompletionConfig{}
	// MISSING: Name
	// MISSING: MatchingOrder
	// MISSING: MaxSuggestions
	// MISSING: MinPrefixLength
	// MISSING: AutoLearning
	// MISSING: SuggestionsInputConfig
	// MISSING: LastSuggestionsImportOperation
	// MISSING: DenylistInputConfig
	// MISSING: LastDenylistImportOperation
	// MISSING: AllowlistInputConfig
	// MISSING: LastAllowlistImportOperation
	return out
}
func RetailCompletionConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.CompletionConfig) *krm.RetailCompletionConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.RetailCompletionConfigSpec{}
	// MISSING: Name
	// MISSING: MatchingOrder
	// MISSING: MaxSuggestions
	// MISSING: MinPrefixLength
	// MISSING: AutoLearning
	// MISSING: SuggestionsInputConfig
	// MISSING: LastSuggestionsImportOperation
	// MISSING: DenylistInputConfig
	// MISSING: LastDenylistImportOperation
	// MISSING: AllowlistInputConfig
	// MISSING: LastAllowlistImportOperation
	return out
}
func RetailCompletionConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.RetailCompletionConfigSpec) *pb.CompletionConfig {
	if in == nil {
		return nil
	}
	out := &pb.CompletionConfig{}
	// MISSING: Name
	// MISSING: MatchingOrder
	// MISSING: MaxSuggestions
	// MISSING: MinPrefixLength
	// MISSING: AutoLearning
	// MISSING: SuggestionsInputConfig
	// MISSING: LastSuggestionsImportOperation
	// MISSING: DenylistInputConfig
	// MISSING: LastDenylistImportOperation
	// MISSING: AllowlistInputConfig
	// MISSING: LastAllowlistImportOperation
	return out
}
