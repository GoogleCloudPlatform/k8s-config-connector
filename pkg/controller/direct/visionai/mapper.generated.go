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

package visionai

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/visionai/apiv1/visionaipb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/visionai/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func FacetProperty_FromProto(mapCtx *direct.MapContext, in *pb.FacetProperty) *krm.FacetProperty {
	if in == nil {
		return nil
	}
	out := &krm.FacetProperty{}
	out.FixedRangeBucketSpec = FacetProperty_FixedRangeBucketSpec_FromProto(mapCtx, in.GetFixedRangeBucketSpec())
	out.CustomRangeBucketSpec = FacetProperty_CustomRangeBucketSpec_FromProto(mapCtx, in.GetCustomRangeBucketSpec())
	out.DatetimeBucketSpec = FacetProperty_DateTimeBucketSpec_FromProto(mapCtx, in.GetDatetimeBucketSpec())
	out.MappedFields = in.MappedFields
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.ResultSize = direct.LazyPtr(in.GetResultSize())
	out.BucketType = direct.Enum_FromProto(mapCtx, in.GetBucketType())
	return out
}
func FacetProperty_ToProto(mapCtx *direct.MapContext, in *krm.FacetProperty) *pb.FacetProperty {
	if in == nil {
		return nil
	}
	out := &pb.FacetProperty{}
	if oneof := FacetProperty_FixedRangeBucketSpec_ToProto(mapCtx, in.FixedRangeBucketSpec); oneof != nil {
		out.RangeFacetConfig = &pb.FacetProperty_FixedRangeBucketSpec_{FixedRangeBucketSpec: oneof}
	}
	if oneof := FacetProperty_CustomRangeBucketSpec_ToProto(mapCtx, in.CustomRangeBucketSpec); oneof != nil {
		out.RangeFacetConfig = &pb.FacetProperty_CustomRangeBucketSpec_{CustomRangeBucketSpec: oneof}
	}
	if oneof := FacetProperty_DateTimeBucketSpec_ToProto(mapCtx, in.DatetimeBucketSpec); oneof != nil {
		out.RangeFacetConfig = &pb.FacetProperty_DatetimeBucketSpec{DatetimeBucketSpec: oneof}
	}
	out.MappedFields = in.MappedFields
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.ResultSize = direct.ValueOf(in.ResultSize)
	out.BucketType = direct.Enum_ToProto[pb.FacetBucketType](mapCtx, in.BucketType)
	return out
}
func FacetProperty_CustomRangeBucketSpec_FromProto(mapCtx *direct.MapContext, in *pb.FacetProperty_CustomRangeBucketSpec) *krm.FacetProperty_CustomRangeBucketSpec {
	if in == nil {
		return nil
	}
	out := &krm.FacetProperty_CustomRangeBucketSpec{}
	out.Endpoints = direct.Slice_FromProto(mapCtx, in.Endpoints, FacetValue_FromProto)
	return out
}
func FacetProperty_CustomRangeBucketSpec_ToProto(mapCtx *direct.MapContext, in *krm.FacetProperty_CustomRangeBucketSpec) *pb.FacetProperty_CustomRangeBucketSpec {
	if in == nil {
		return nil
	}
	out := &pb.FacetProperty_CustomRangeBucketSpec{}
	out.Endpoints = direct.Slice_ToProto(mapCtx, in.Endpoints, FacetValue_ToProto)
	return out
}
func FacetProperty_DateTimeBucketSpec_FromProto(mapCtx *direct.MapContext, in *pb.FacetProperty_DateTimeBucketSpec) *krm.FacetProperty_DateTimeBucketSpec {
	if in == nil {
		return nil
	}
	out := &krm.FacetProperty_DateTimeBucketSpec{}
	out.Granularity = direct.Enum_FromProto(mapCtx, in.GetGranularity())
	return out
}
func FacetProperty_DateTimeBucketSpec_ToProto(mapCtx *direct.MapContext, in *krm.FacetProperty_DateTimeBucketSpec) *pb.FacetProperty_DateTimeBucketSpec {
	if in == nil {
		return nil
	}
	out := &pb.FacetProperty_DateTimeBucketSpec{}
	out.Granularity = direct.Enum_ToProto[pb.FacetProperty_DateTimeBucketSpec_Granularity](mapCtx, in.Granularity)
	return out
}
func FacetProperty_FixedRangeBucketSpec_FromProto(mapCtx *direct.MapContext, in *pb.FacetProperty_FixedRangeBucketSpec) *krm.FacetProperty_FixedRangeBucketSpec {
	if in == nil {
		return nil
	}
	out := &krm.FacetProperty_FixedRangeBucketSpec{}
	out.BucketStart = FacetValue_FromProto(mapCtx, in.GetBucketStart())
	out.BucketGranularity = FacetValue_FromProto(mapCtx, in.GetBucketGranularity())
	out.BucketCount = direct.LazyPtr(in.GetBucketCount())
	return out
}
func FacetProperty_FixedRangeBucketSpec_ToProto(mapCtx *direct.MapContext, in *krm.FacetProperty_FixedRangeBucketSpec) *pb.FacetProperty_FixedRangeBucketSpec {
	if in == nil {
		return nil
	}
	out := &pb.FacetProperty_FixedRangeBucketSpec{}
	out.BucketStart = FacetValue_ToProto(mapCtx, in.BucketStart)
	out.BucketGranularity = FacetValue_ToProto(mapCtx, in.BucketGranularity)
	out.BucketCount = direct.ValueOf(in.BucketCount)
	return out
}
func FacetValue_FromProto(mapCtx *direct.MapContext, in *pb.FacetValue) *krm.FacetValue {
	if in == nil {
		return nil
	}
	out := &krm.FacetValue{}
	out.StringValue = direct.LazyPtr(in.GetStringValue())
	out.IntegerValue = direct.LazyPtr(in.GetIntegerValue())
	out.DatetimeValue = DateTime_FromProto(mapCtx, in.GetDatetimeValue())
	return out
}
func FacetValue_ToProto(mapCtx *direct.MapContext, in *krm.FacetValue) *pb.FacetValue {
	if in == nil {
		return nil
	}
	out := &pb.FacetValue{}
	if oneof := FacetValue_StringValue_ToProto(mapCtx, in.StringValue); oneof != nil {
		out.Value = oneof
	}
	if oneof := FacetValue_IntegerValue_ToProto(mapCtx, in.IntegerValue); oneof != nil {
		out.Value = oneof
	}
	if oneof := DateTime_ToProto(mapCtx, in.DatetimeValue); oneof != nil {
		out.Value = &pb.FacetValue_DatetimeValue{DatetimeValue: oneof}
	}
	return out
}
func SearchConfig_FromProto(mapCtx *direct.MapContext, in *pb.SearchConfig) *krm.SearchConfig {
	if in == nil {
		return nil
	}
	out := &krm.SearchConfig{}
	out.Name = direct.LazyPtr(in.GetName())
	out.FacetProperty = FacetProperty_FromProto(mapCtx, in.GetFacetProperty())
	out.SearchCriteriaProperty = SearchCriteriaProperty_FromProto(mapCtx, in.GetSearchCriteriaProperty())
	return out
}
func SearchConfig_ToProto(mapCtx *direct.MapContext, in *krm.SearchConfig) *pb.SearchConfig {
	if in == nil {
		return nil
	}
	out := &pb.SearchConfig{}
	out.Name = direct.ValueOf(in.Name)
	out.FacetProperty = FacetProperty_ToProto(mapCtx, in.FacetProperty)
	out.SearchCriteriaProperty = SearchCriteriaProperty_ToProto(mapCtx, in.SearchCriteriaProperty)
	return out
}
func SearchCriteriaProperty_FromProto(mapCtx *direct.MapContext, in *pb.SearchCriteriaProperty) *krm.SearchCriteriaProperty {
	if in == nil {
		return nil
	}
	out := &krm.SearchCriteriaProperty{}
	out.MappedFields = in.MappedFields
	return out
}
func SearchCriteriaProperty_ToProto(mapCtx *direct.MapContext, in *krm.SearchCriteriaProperty) *pb.SearchCriteriaProperty {
	if in == nil {
		return nil
	}
	out := &pb.SearchCriteriaProperty{}
	out.MappedFields = in.MappedFields
	return out
}
func VisionaiSearchConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SearchConfig) *krm.VisionaiSearchConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.VisionaiSearchConfigObservedState{}
	// MISSING: Name
	// MISSING: FacetProperty
	// MISSING: SearchCriteriaProperty
	return out
}
func VisionaiSearchConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.VisionaiSearchConfigObservedState) *pb.SearchConfig {
	if in == nil {
		return nil
	}
	out := &pb.SearchConfig{}
	// MISSING: Name
	// MISSING: FacetProperty
	// MISSING: SearchCriteriaProperty
	return out
}
func VisionaiSearchConfigSpec_FromProto(mapCtx *direct.MapContext, in *pb.SearchConfig) *krm.VisionaiSearchConfigSpec {
	if in == nil {
		return nil
	}
	out := &krm.VisionaiSearchConfigSpec{}
	// MISSING: Name
	// MISSING: FacetProperty
	// MISSING: SearchCriteriaProperty
	return out
}
func VisionaiSearchConfigSpec_ToProto(mapCtx *direct.MapContext, in *krm.VisionaiSearchConfigSpec) *pb.SearchConfig {
	if in == nil {
		return nil
	}
	out := &pb.SearchConfig{}
	// MISSING: Name
	// MISSING: FacetProperty
	// MISSING: SearchCriteriaProperty
	return out
}
