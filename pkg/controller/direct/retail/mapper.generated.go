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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/retail/apiv2beta/retailpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/retail/v1alpha1"
)
func AttributesConfig_FromProto(mapCtx *direct.MapContext, in *pb.AttributesConfig) *krm.AttributesConfig {
	if in == nil {
		return nil
	}
	out := &krm.AttributesConfig{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: CatalogAttributes
	// MISSING: AttributeConfigLevel
	return out
}
func AttributesConfig_ToProto(mapCtx *direct.MapContext, in *krm.AttributesConfig) *pb.AttributesConfig {
	if in == nil {
		return nil
	}
	out := &pb.AttributesConfig{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: CatalogAttributes
	// MISSING: AttributeConfigLevel
	return out
}
func AttributesConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.AttributesConfig) *krm.AttributesConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AttributesConfigObservedState{}
	// MISSING: Name
	// MISSING: CatalogAttributes
	out.AttributeConfigLevel = direct.Enum_FromProto(mapCtx, in.GetAttributeConfigLevel())
	return out
}
func AttributesConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AttributesConfigObservedState) *pb.AttributesConfig {
	if in == nil {
		return nil
	}
	out := &pb.AttributesConfig{}
	// MISSING: Name
	// MISSING: CatalogAttributes
	out.AttributeConfigLevel = direct.Enum_ToProto[pb.AttributeConfigLevel](mapCtx, in.AttributeConfigLevel)
	return out
}
func CatalogAttribute_FromProto(mapCtx *direct.MapContext, in *pb.CatalogAttribute) *krm.CatalogAttribute {
	if in == nil {
		return nil
	}
	out := &krm.CatalogAttribute{}
	out.Key = direct.LazyPtr(in.GetKey())
	// MISSING: InUse
	// MISSING: Type
	out.IndexableOption = direct.Enum_FromProto(mapCtx, in.GetIndexableOption())
	out.DynamicFacetableOption = direct.Enum_FromProto(mapCtx, in.GetDynamicFacetableOption())
	out.SearchableOption = direct.Enum_FromProto(mapCtx, in.GetSearchableOption())
	out.RecommendationsFilteringOption = direct.Enum_FromProto(mapCtx, in.GetRecommendationsFilteringOption())
	out.ExactSearchableOption = direct.Enum_FromProto(mapCtx, in.GetExactSearchableOption())
	out.RetrievableOption = direct.Enum_FromProto(mapCtx, in.GetRetrievableOption())
	out.FacetConfig = CatalogAttribute_FacetConfig_FromProto(mapCtx, in.GetFacetConfig())
	return out
}
func CatalogAttribute_ToProto(mapCtx *direct.MapContext, in *krm.CatalogAttribute) *pb.CatalogAttribute {
	if in == nil {
		return nil
	}
	out := &pb.CatalogAttribute{}
	out.Key = direct.ValueOf(in.Key)
	// MISSING: InUse
	// MISSING: Type
	out.IndexableOption = direct.Enum_ToProto[pb.CatalogAttribute_IndexableOption](mapCtx, in.IndexableOption)
	out.DynamicFacetableOption = direct.Enum_ToProto[pb.CatalogAttribute_DynamicFacetableOption](mapCtx, in.DynamicFacetableOption)
	out.SearchableOption = direct.Enum_ToProto[pb.CatalogAttribute_SearchableOption](mapCtx, in.SearchableOption)
	out.RecommendationsFilteringOption = direct.Enum_ToProto[pb.RecommendationsFilteringOption](mapCtx, in.RecommendationsFilteringOption)
	out.ExactSearchableOption = direct.Enum_ToProto[pb.CatalogAttribute_ExactSearchableOption](mapCtx, in.ExactSearchableOption)
	out.RetrievableOption = direct.Enum_ToProto[pb.CatalogAttribute_RetrievableOption](mapCtx, in.RetrievableOption)
	out.FacetConfig = CatalogAttribute_FacetConfig_ToProto(mapCtx, in.FacetConfig)
	return out
}
func CatalogAttribute_FacetConfig_FromProto(mapCtx *direct.MapContext, in *pb.CatalogAttribute_FacetConfig) *krm.CatalogAttribute_FacetConfig {
	if in == nil {
		return nil
	}
	out := &krm.CatalogAttribute_FacetConfig{}
	out.FacetIntervals = direct.Slice_FromProto(mapCtx, in.FacetIntervals, Interval_FromProto)
	out.IgnoredFacetValues = direct.Slice_FromProto(mapCtx, in.IgnoredFacetValues, CatalogAttribute_FacetConfig_IgnoredFacetValues_FromProto)
	out.MergedFacetValues = direct.Slice_FromProto(mapCtx, in.MergedFacetValues, CatalogAttribute_FacetConfig_MergedFacetValue_FromProto)
	out.MergedFacet = CatalogAttribute_FacetConfig_MergedFacet_FromProto(mapCtx, in.GetMergedFacet())
	out.RerankConfig = CatalogAttribute_FacetConfig_RerankConfig_FromProto(mapCtx, in.GetRerankConfig())
	return out
}
func CatalogAttribute_FacetConfig_ToProto(mapCtx *direct.MapContext, in *krm.CatalogAttribute_FacetConfig) *pb.CatalogAttribute_FacetConfig {
	if in == nil {
		return nil
	}
	out := &pb.CatalogAttribute_FacetConfig{}
	out.FacetIntervals = direct.Slice_ToProto(mapCtx, in.FacetIntervals, Interval_ToProto)
	out.IgnoredFacetValues = direct.Slice_ToProto(mapCtx, in.IgnoredFacetValues, CatalogAttribute_FacetConfig_IgnoredFacetValues_ToProto)
	out.MergedFacetValues = direct.Slice_ToProto(mapCtx, in.MergedFacetValues, CatalogAttribute_FacetConfig_MergedFacetValue_ToProto)
	out.MergedFacet = CatalogAttribute_FacetConfig_MergedFacet_ToProto(mapCtx, in.MergedFacet)
	out.RerankConfig = CatalogAttribute_FacetConfig_RerankConfig_ToProto(mapCtx, in.RerankConfig)
	return out
}
func CatalogAttribute_FacetConfig_IgnoredFacetValues_FromProto(mapCtx *direct.MapContext, in *pb.CatalogAttribute_FacetConfig_IgnoredFacetValues) *krm.CatalogAttribute_FacetConfig_IgnoredFacetValues {
	if in == nil {
		return nil
	}
	out := &krm.CatalogAttribute_FacetConfig_IgnoredFacetValues{}
	out.Values = in.Values
	out.StartTime = direct.StringTimestamp_FromProto(mapCtx, in.GetStartTime())
	out.EndTime = direct.StringTimestamp_FromProto(mapCtx, in.GetEndTime())
	return out
}
func CatalogAttribute_FacetConfig_IgnoredFacetValues_ToProto(mapCtx *direct.MapContext, in *krm.CatalogAttribute_FacetConfig_IgnoredFacetValues) *pb.CatalogAttribute_FacetConfig_IgnoredFacetValues {
	if in == nil {
		return nil
	}
	out := &pb.CatalogAttribute_FacetConfig_IgnoredFacetValues{}
	out.Values = in.Values
	out.StartTime = direct.StringTimestamp_ToProto(mapCtx, in.StartTime)
	out.EndTime = direct.StringTimestamp_ToProto(mapCtx, in.EndTime)
	return out
}
func CatalogAttribute_FacetConfig_MergedFacet_FromProto(mapCtx *direct.MapContext, in *pb.CatalogAttribute_FacetConfig_MergedFacet) *krm.CatalogAttribute_FacetConfig_MergedFacet {
	if in == nil {
		return nil
	}
	out := &krm.CatalogAttribute_FacetConfig_MergedFacet{}
	out.MergedFacetKey = direct.LazyPtr(in.GetMergedFacetKey())
	return out
}
func CatalogAttribute_FacetConfig_MergedFacet_ToProto(mapCtx *direct.MapContext, in *krm.CatalogAttribute_FacetConfig_MergedFacet) *pb.CatalogAttribute_FacetConfig_MergedFacet {
	if in == nil {
		return nil
	}
	out := &pb.CatalogAttribute_FacetConfig_MergedFacet{}
	out.MergedFacetKey = direct.ValueOf(in.MergedFacetKey)
	return out
}
func CatalogAttribute_FacetConfig_MergedFacetValue_FromProto(mapCtx *direct.MapContext, in *pb.CatalogAttribute_FacetConfig_MergedFacetValue) *krm.CatalogAttribute_FacetConfig_MergedFacetValue {
	if in == nil {
		return nil
	}
	out := &krm.CatalogAttribute_FacetConfig_MergedFacetValue{}
	out.Values = in.Values
	out.MergedValue = direct.LazyPtr(in.GetMergedValue())
	return out
}
func CatalogAttribute_FacetConfig_MergedFacetValue_ToProto(mapCtx *direct.MapContext, in *krm.CatalogAttribute_FacetConfig_MergedFacetValue) *pb.CatalogAttribute_FacetConfig_MergedFacetValue {
	if in == nil {
		return nil
	}
	out := &pb.CatalogAttribute_FacetConfig_MergedFacetValue{}
	out.Values = in.Values
	out.MergedValue = direct.ValueOf(in.MergedValue)
	return out
}
func CatalogAttribute_FacetConfig_RerankConfig_FromProto(mapCtx *direct.MapContext, in *pb.CatalogAttribute_FacetConfig_RerankConfig) *krm.CatalogAttribute_FacetConfig_RerankConfig {
	if in == nil {
		return nil
	}
	out := &krm.CatalogAttribute_FacetConfig_RerankConfig{}
	out.RerankFacet = direct.LazyPtr(in.GetRerankFacet())
	out.FacetValues = in.FacetValues
	return out
}
func CatalogAttribute_FacetConfig_RerankConfig_ToProto(mapCtx *direct.MapContext, in *krm.CatalogAttribute_FacetConfig_RerankConfig) *pb.CatalogAttribute_FacetConfig_RerankConfig {
	if in == nil {
		return nil
	}
	out := &pb.CatalogAttribute_FacetConfig_RerankConfig{}
	out.RerankFacet = direct.ValueOf(in.RerankFacet)
	out.FacetValues = in.FacetValues
	return out
}
func Interval_FromProto(mapCtx *direct.MapContext, in *pb.Interval) *krm.Interval {
	if in == nil {
		return nil
	}
	out := &krm.Interval{}
	out.Minimum = direct.LazyPtr(in.GetMinimum())
	out.ExclusiveMinimum = direct.LazyPtr(in.GetExclusiveMinimum())
	out.Maximum = direct.LazyPtr(in.GetMaximum())
	out.ExclusiveMaximum = direct.LazyPtr(in.GetExclusiveMaximum())
	return out
}
func Interval_ToProto(mapCtx *direct.MapContext, in *krm.Interval) *pb.Interval {
	if in == nil {
		return nil
	}
	out := &pb.Interval{}
	if oneof := Interval_Minimum_ToProto(mapCtx, in.Minimum); oneof != nil {
		out.Min = oneof
	}
	if oneof := Interval_ExclusiveMinimum_ToProto(mapCtx, in.ExclusiveMinimum); oneof != nil {
		out.Min = oneof
	}
	if oneof := Interval_Maximum_ToProto(mapCtx, in.Maximum); oneof != nil {
		out.Max = oneof
	}
	if oneof := Interval_ExclusiveMaximum_ToProto(mapCtx, in.ExclusiveMaximum); oneof != nil {
		out.Max = oneof
	}
	return out
}
