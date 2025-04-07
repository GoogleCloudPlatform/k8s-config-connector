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

// +generated:mapper
// krm.group: dataplex.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.dataplex.v1

package dataplex

import (
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	krmv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataplex/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DataplexZoneObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Zone) *krmv1alpha1.DataplexZoneObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.DataplexZoneObservedState{}
	// MISSING: Name
	out.Uid = direct.LazyPtr(in.GetUid())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.AssetStatus = AssetStatus_FromProto(mapCtx, in.GetAssetStatus())
	return out
}
func DataplexZoneObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.DataplexZoneObservedState) *pb.Zone {
	if in == nil {
		return nil
	}
	out := &pb.Zone{}
	// MISSING: Name
	out.Uid = direct.ValueOf(in.Uid)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.State = direct.Enum_ToProto[pb.State](mapCtx, in.State)
	out.AssetStatus = AssetStatus_ToProto(mapCtx, in.AssetStatus)
	return out
}
func DataplexZoneSpec_FromProto(mapCtx *direct.MapContext, in *pb.Zone) *krmv1alpha1.DataplexZoneSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.DataplexZoneSpec{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Labels = in.Labels
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.DiscoverySpec = Zone_DiscoverySpec_FromProto(mapCtx, in.GetDiscoverySpec())
	out.ResourceSpec = Zone_ResourceSpec_FromProto(mapCtx, in.GetResourceSpec())
	return out
}
func DataplexZoneSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.DataplexZoneSpec) *pb.Zone {
	if in == nil {
		return nil
	}
	out := &pb.Zone{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Labels = in.Labels
	out.Description = direct.ValueOf(in.Description)
	out.Type = direct.Enum_ToProto[pb.Zone_Type](mapCtx, in.Type)
	out.DiscoverySpec = Zone_DiscoverySpec_ToProto(mapCtx, in.DiscoverySpec)
	out.ResourceSpec = Zone_ResourceSpec_ToProto(mapCtx, in.ResourceSpec)
	return out
}
func Zone_DiscoverySpec_FromProto(mapCtx *direct.MapContext, in *pb.Zone_DiscoverySpec) *krmv1alpha1.Zone_DiscoverySpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Zone_DiscoverySpec{}
	out.Enabled = direct.LazyPtr(in.GetEnabled())
	out.IncludePatterns = in.IncludePatterns
	out.ExcludePatterns = in.ExcludePatterns
	out.CsvOptions = Zone_DiscoverySpec_CsvOptions_FromProto(mapCtx, in.GetCsvOptions())
	out.JsonOptions = Zone_DiscoverySpec_JsonOptions_FromProto(mapCtx, in.GetJsonOptions())
	out.Schedule = direct.LazyPtr(in.GetSchedule())
	return out
}
func Zone_DiscoverySpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Zone_DiscoverySpec) *pb.Zone_DiscoverySpec {
	if in == nil {
		return nil
	}
	out := &pb.Zone_DiscoverySpec{}
	out.Enabled = direct.ValueOf(in.Enabled)
	out.IncludePatterns = in.IncludePatterns
	out.ExcludePatterns = in.ExcludePatterns
	out.CsvOptions = Zone_DiscoverySpec_CsvOptions_ToProto(mapCtx, in.CsvOptions)
	out.JsonOptions = Zone_DiscoverySpec_JsonOptions_ToProto(mapCtx, in.JsonOptions)
	if oneof := Zone_DiscoverySpec_Schedule_ToProto(mapCtx, in.Schedule); oneof != nil {
		out.Trigger = oneof
	}
	return out
}
func Zone_DiscoverySpec_CsvOptions_FromProto(mapCtx *direct.MapContext, in *pb.Zone_DiscoverySpec_CsvOptions) *krmv1alpha1.Zone_DiscoverySpec_CsvOptions {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Zone_DiscoverySpec_CsvOptions{}
	out.HeaderRows = direct.LazyPtr(in.GetHeaderRows())
	out.Delimiter = direct.LazyPtr(in.GetDelimiter())
	out.Encoding = direct.LazyPtr(in.GetEncoding())
	out.DisableTypeInference = direct.LazyPtr(in.GetDisableTypeInference())
	return out
}
func Zone_DiscoverySpec_CsvOptions_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Zone_DiscoverySpec_CsvOptions) *pb.Zone_DiscoverySpec_CsvOptions {
	if in == nil {
		return nil
	}
	out := &pb.Zone_DiscoverySpec_CsvOptions{}
	out.HeaderRows = direct.ValueOf(in.HeaderRows)
	out.Delimiter = direct.ValueOf(in.Delimiter)
	out.Encoding = direct.ValueOf(in.Encoding)
	out.DisableTypeInference = direct.ValueOf(in.DisableTypeInference)
	return out
}
func Zone_DiscoverySpec_JsonOptions_FromProto(mapCtx *direct.MapContext, in *pb.Zone_DiscoverySpec_JsonOptions) *krmv1alpha1.Zone_DiscoverySpec_JsonOptions {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Zone_DiscoverySpec_JsonOptions{}
	out.Encoding = direct.LazyPtr(in.GetEncoding())
	out.DisableTypeInference = direct.LazyPtr(in.GetDisableTypeInference())
	return out
}
func Zone_DiscoverySpec_JsonOptions_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Zone_DiscoverySpec_JsonOptions) *pb.Zone_DiscoverySpec_JsonOptions {
	if in == nil {
		return nil
	}
	out := &pb.Zone_DiscoverySpec_JsonOptions{}
	out.Encoding = direct.ValueOf(in.Encoding)
	out.DisableTypeInference = direct.ValueOf(in.DisableTypeInference)
	return out
}
func Zone_DiscoverySpec_Schedule_ToProto(mapCtx *direct.MapContext, in *string) *pb.Zone_DiscoverySpec_Schedule {
	if in == nil {
		return nil
	}
	out := &pb.Zone_DiscoverySpec_Schedule{}
	out.Schedule = direct.ValueOf(in)
	return out
}
func Zone_ResourceSpec_FromProto(mapCtx *direct.MapContext, in *pb.Zone_ResourceSpec) *krmv1alpha1.Zone_ResourceSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Zone_ResourceSpec{}
	out.LocationType = direct.Enum_FromProto(mapCtx, in.GetLocationType())
	return out
}
func Zone_ResourceSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Zone_ResourceSpec) *pb.Zone_ResourceSpec {
	if in == nil {
		return nil
	}
	out := &pb.Zone_ResourceSpec{}
	out.LocationType = direct.Enum_ToProto[pb.Zone_ResourceSpec_LocationType](mapCtx, in.LocationType)
	return out
}
