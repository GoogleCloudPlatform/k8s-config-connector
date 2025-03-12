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

package dataplex

import (
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataplex/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AssetStatus_FromProto(mapCtx *direct.MapContext, in *pb.AssetStatus) *krm.AssetStatus {
	if in == nil {
		return nil
	}
	out := &krm.AssetStatus{}
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.ActiveAssets = direct.LazyPtr(in.GetActiveAssets())
	out.SecurityPolicyApplyingAssets = direct.LazyPtr(in.GetSecurityPolicyApplyingAssets())
	return out
}
func AssetStatus_ToProto(mapCtx *direct.MapContext, in *krm.AssetStatus) *pb.AssetStatus {
	if in == nil {
		return nil
	}
	out := &pb.AssetStatus{}
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.ActiveAssets = direct.ValueOf(in.ActiveAssets)
	out.SecurityPolicyApplyingAssets = direct.ValueOf(in.SecurityPolicyApplyingAssets)
	return out
}
func DataplexZoneObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Zone) *krm.DataplexZoneObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataplexZoneObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: State
	// MISSING: Type
	// MISSING: DiscoverySpec
	// MISSING: ResourceSpec
	// MISSING: AssetStatus
	return out
}
func DataplexZoneObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataplexZoneObservedState) *pb.Zone {
	if in == nil {
		return nil
	}
	out := &pb.Zone{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: State
	// MISSING: Type
	// MISSING: DiscoverySpec
	// MISSING: ResourceSpec
	// MISSING: AssetStatus
	return out
}
func DataplexZoneSpec_FromProto(mapCtx *direct.MapContext, in *pb.Zone) *krm.DataplexZoneSpec {
	if in == nil {
		return nil
	}
	out := &krm.DataplexZoneSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: State
	// MISSING: Type
	// MISSING: DiscoverySpec
	// MISSING: ResourceSpec
	// MISSING: AssetStatus
	return out
}
func DataplexZoneSpec_ToProto(mapCtx *direct.MapContext, in *krm.DataplexZoneSpec) *pb.Zone {
	if in == nil {
		return nil
	}
	out := &pb.Zone{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: State
	// MISSING: Type
	// MISSING: DiscoverySpec
	// MISSING: ResourceSpec
	// MISSING: AssetStatus
	return out
}
func Zone_FromProto(mapCtx *direct.MapContext, in *pb.Zone) *krm.Zone {
	if in == nil {
		return nil
	}
	out := &krm.Zone{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: State
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.DiscoverySpec = Zone_DiscoverySpec_FromProto(mapCtx, in.GetDiscoverySpec())
	out.ResourceSpec = Zone_ResourceSpec_FromProto(mapCtx, in.GetResourceSpec())
	// MISSING: AssetStatus
	return out
}
func Zone_ToProto(mapCtx *direct.MapContext, in *krm.Zone) *pb.Zone {
	if in == nil {
		return nil
	}
	out := &pb.Zone{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.ValueOf(in.Description)
	// MISSING: State
	out.Type = direct.Enum_ToProto[pb.Zone_Type](mapCtx, in.Type)
	out.DiscoverySpec = Zone_DiscoverySpec_ToProto(mapCtx, in.DiscoverySpec)
	out.ResourceSpec = Zone_ResourceSpec_ToProto(mapCtx, in.ResourceSpec)
	// MISSING: AssetStatus
	return out
}
func ZoneObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Zone) *krm.ZoneObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ZoneObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: DisplayName
	out.Uid = direct.LazyPtr(in.GetUid())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: Description
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: Type
	// MISSING: DiscoverySpec
	// MISSING: ResourceSpec
	out.AssetStatus = AssetStatus_FromProto(mapCtx, in.GetAssetStatus())
	return out
}
func ZoneObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ZoneObservedState) *pb.Zone {
	if in == nil {
		return nil
	}
	out := &pb.Zone{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: DisplayName
	out.Uid = direct.ValueOf(in.Uid)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: Description
	out.State = direct.Enum_ToProto[pb.State](mapCtx, in.State)
	// MISSING: Type
	// MISSING: DiscoverySpec
	// MISSING: ResourceSpec
	out.AssetStatus = AssetStatus_ToProto(mapCtx, in.AssetStatus)
	return out
}
func Zone_DiscoverySpec_FromProto(mapCtx *direct.MapContext, in *pb.Zone_DiscoverySpec) *krm.Zone_DiscoverySpec {
	if in == nil {
		return nil
	}
	out := &krm.Zone_DiscoverySpec{}
	out.Enabled = direct.LazyPtr(in.GetEnabled())
	out.IncludePatterns = in.IncludePatterns
	out.ExcludePatterns = in.ExcludePatterns
	out.CsvOptions = Zone_DiscoverySpec_CsvOptions_FromProto(mapCtx, in.GetCsvOptions())
	out.JsonOptions = Zone_DiscoverySpec_JsonOptions_FromProto(mapCtx, in.GetJsonOptions())
	out.Schedule = direct.LazyPtr(in.GetSchedule())
	return out
}
func Zone_DiscoverySpec_ToProto(mapCtx *direct.MapContext, in *krm.Zone_DiscoverySpec) *pb.Zone_DiscoverySpec {
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
func Zone_DiscoverySpec_CsvOptions_FromProto(mapCtx *direct.MapContext, in *pb.Zone_DiscoverySpec_CsvOptions) *krm.Zone_DiscoverySpec_CsvOptions {
	if in == nil {
		return nil
	}
	out := &krm.Zone_DiscoverySpec_CsvOptions{}
	out.HeaderRows = direct.LazyPtr(in.GetHeaderRows())
	out.Delimiter = direct.LazyPtr(in.GetDelimiter())
	out.Encoding = direct.LazyPtr(in.GetEncoding())
	out.DisableTypeInference = direct.LazyPtr(in.GetDisableTypeInference())
	return out
}
func Zone_DiscoverySpec_CsvOptions_ToProto(mapCtx *direct.MapContext, in *krm.Zone_DiscoverySpec_CsvOptions) *pb.Zone_DiscoverySpec_CsvOptions {
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
func Zone_DiscoverySpec_JsonOptions_FromProto(mapCtx *direct.MapContext, in *pb.Zone_DiscoverySpec_JsonOptions) *krm.Zone_DiscoverySpec_JsonOptions {
	if in == nil {
		return nil
	}
	out := &krm.Zone_DiscoverySpec_JsonOptions{}
	out.Encoding = direct.LazyPtr(in.GetEncoding())
	out.DisableTypeInference = direct.LazyPtr(in.GetDisableTypeInference())
	return out
}
func Zone_DiscoverySpec_JsonOptions_ToProto(mapCtx *direct.MapContext, in *krm.Zone_DiscoverySpec_JsonOptions) *pb.Zone_DiscoverySpec_JsonOptions {
	if in == nil {
		return nil
	}
	out := &pb.Zone_DiscoverySpec_JsonOptions{}
	out.Encoding = direct.ValueOf(in.Encoding)
	out.DisableTypeInference = direct.ValueOf(in.DisableTypeInference)
	return out
}
func Zone_ResourceSpec_FromProto(mapCtx *direct.MapContext, in *pb.Zone_ResourceSpec) *krm.Zone_ResourceSpec {
	if in == nil {
		return nil
	}
	out := &krm.Zone_ResourceSpec{}
	out.LocationType = direct.Enum_FromProto(mapCtx, in.GetLocationType())
	return out
}
func Zone_ResourceSpec_ToProto(mapCtx *direct.MapContext, in *krm.Zone_ResourceSpec) *pb.Zone_ResourceSpec {
	if in == nil {
		return nil
	}
	out := &pb.Zone_ResourceSpec{}
	out.LocationType = direct.Enum_ToProto[pb.Zone_ResourceSpec_LocationType](mapCtx, in.LocationType)
	return out
}
