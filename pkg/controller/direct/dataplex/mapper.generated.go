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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataplex/v1alpha1"
)
func Asset_FromProto(mapCtx *direct.MapContext, in *pb.Asset) *krm.Asset {
	if in == nil {
		return nil
	}
	out := &krm.Asset{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: State
	out.ResourceSpec = Asset_ResourceSpec_FromProto(mapCtx, in.GetResourceSpec())
	// MISSING: ResourceStatus
	// MISSING: SecurityStatus
	out.DiscoverySpec = Asset_DiscoverySpec_FromProto(mapCtx, in.GetDiscoverySpec())
	// MISSING: DiscoveryStatus
	return out
}
func Asset_ToProto(mapCtx *direct.MapContext, in *krm.Asset) *pb.Asset {
	if in == nil {
		return nil
	}
	out := &pb.Asset{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.ValueOf(in.Description)
	// MISSING: State
	out.ResourceSpec = Asset_ResourceSpec_ToProto(mapCtx, in.ResourceSpec)
	// MISSING: ResourceStatus
	// MISSING: SecurityStatus
	out.DiscoverySpec = Asset_DiscoverySpec_ToProto(mapCtx, in.DiscoverySpec)
	// MISSING: DiscoveryStatus
	return out
}
func AssetObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Asset) *krm.AssetObservedState {
	if in == nil {
		return nil
	}
	out := &krm.AssetObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: DisplayName
	out.Uid = direct.LazyPtr(in.GetUid())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: Description
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: ResourceSpec
	out.ResourceStatus = Asset_ResourceStatus_FromProto(mapCtx, in.GetResourceStatus())
	out.SecurityStatus = Asset_SecurityStatus_FromProto(mapCtx, in.GetSecurityStatus())
	// MISSING: DiscoverySpec
	out.DiscoveryStatus = Asset_DiscoveryStatus_FromProto(mapCtx, in.GetDiscoveryStatus())
	return out
}
func AssetObservedState_ToProto(mapCtx *direct.MapContext, in *krm.AssetObservedState) *pb.Asset {
	if in == nil {
		return nil
	}
	out := &pb.Asset{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: DisplayName
	out.Uid = direct.ValueOf(in.Uid)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: Description
	out.State = direct.Enum_ToProto[pb.State](mapCtx, in.State)
	// MISSING: ResourceSpec
	out.ResourceStatus = Asset_ResourceStatus_ToProto(mapCtx, in.ResourceStatus)
	out.SecurityStatus = Asset_SecurityStatus_ToProto(mapCtx, in.SecurityStatus)
	// MISSING: DiscoverySpec
	out.DiscoveryStatus = Asset_DiscoveryStatus_ToProto(mapCtx, in.DiscoveryStatus)
	return out
}
func Asset_DiscoverySpec_FromProto(mapCtx *direct.MapContext, in *pb.Asset_DiscoverySpec) *krm.Asset_DiscoverySpec {
	if in == nil {
		return nil
	}
	out := &krm.Asset_DiscoverySpec{}
	out.Enabled = direct.LazyPtr(in.GetEnabled())
	out.IncludePatterns = in.IncludePatterns
	out.ExcludePatterns = in.ExcludePatterns
	out.CsvOptions = Asset_DiscoverySpec_CsvOptions_FromProto(mapCtx, in.GetCsvOptions())
	out.JsonOptions = Asset_DiscoverySpec_JsonOptions_FromProto(mapCtx, in.GetJsonOptions())
	out.Schedule = direct.LazyPtr(in.GetSchedule())
	return out
}
func Asset_DiscoverySpec_ToProto(mapCtx *direct.MapContext, in *krm.Asset_DiscoverySpec) *pb.Asset_DiscoverySpec {
	if in == nil {
		return nil
	}
	out := &pb.Asset_DiscoverySpec{}
	out.Enabled = direct.ValueOf(in.Enabled)
	out.IncludePatterns = in.IncludePatterns
	out.ExcludePatterns = in.ExcludePatterns
	out.CsvOptions = Asset_DiscoverySpec_CsvOptions_ToProto(mapCtx, in.CsvOptions)
	out.JsonOptions = Asset_DiscoverySpec_JsonOptions_ToProto(mapCtx, in.JsonOptions)
	if oneof := Asset_DiscoverySpec_Schedule_ToProto(mapCtx, in.Schedule); oneof != nil {
		out.Trigger = oneof
	}
	return out
}
func Asset_DiscoverySpec_CsvOptions_FromProto(mapCtx *direct.MapContext, in *pb.Asset_DiscoverySpec_CsvOptions) *krm.Asset_DiscoverySpec_CsvOptions {
	if in == nil {
		return nil
	}
	out := &krm.Asset_DiscoverySpec_CsvOptions{}
	out.HeaderRows = direct.LazyPtr(in.GetHeaderRows())
	out.Delimiter = direct.LazyPtr(in.GetDelimiter())
	out.Encoding = direct.LazyPtr(in.GetEncoding())
	out.DisableTypeInference = direct.LazyPtr(in.GetDisableTypeInference())
	return out
}
func Asset_DiscoverySpec_CsvOptions_ToProto(mapCtx *direct.MapContext, in *krm.Asset_DiscoverySpec_CsvOptions) *pb.Asset_DiscoverySpec_CsvOptions {
	if in == nil {
		return nil
	}
	out := &pb.Asset_DiscoverySpec_CsvOptions{}
	out.HeaderRows = direct.ValueOf(in.HeaderRows)
	out.Delimiter = direct.ValueOf(in.Delimiter)
	out.Encoding = direct.ValueOf(in.Encoding)
	out.DisableTypeInference = direct.ValueOf(in.DisableTypeInference)
	return out
}
func Asset_DiscoverySpec_JsonOptions_FromProto(mapCtx *direct.MapContext, in *pb.Asset_DiscoverySpec_JsonOptions) *krm.Asset_DiscoverySpec_JsonOptions {
	if in == nil {
		return nil
	}
	out := &krm.Asset_DiscoverySpec_JsonOptions{}
	out.Encoding = direct.LazyPtr(in.GetEncoding())
	out.DisableTypeInference = direct.LazyPtr(in.GetDisableTypeInference())
	return out
}
func Asset_DiscoverySpec_JsonOptions_ToProto(mapCtx *direct.MapContext, in *krm.Asset_DiscoverySpec_JsonOptions) *pb.Asset_DiscoverySpec_JsonOptions {
	if in == nil {
		return nil
	}
	out := &pb.Asset_DiscoverySpec_JsonOptions{}
	out.Encoding = direct.ValueOf(in.Encoding)
	out.DisableTypeInference = direct.ValueOf(in.DisableTypeInference)
	return out
}
func Asset_DiscoveryStatus_FromProto(mapCtx *direct.MapContext, in *pb.Asset_DiscoveryStatus) *krm.Asset_DiscoveryStatus {
	if in == nil {
		return nil
	}
	out := &krm.Asset_DiscoveryStatus{}
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Message = direct.LazyPtr(in.GetMessage())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.LastRunTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLastRunTime())
	out.Stats = Asset_DiscoveryStatus_Stats_FromProto(mapCtx, in.GetStats())
	out.LastRunDuration = direct.StringDuration_FromProto(mapCtx, in.GetLastRunDuration())
	return out
}
func Asset_DiscoveryStatus_ToProto(mapCtx *direct.MapContext, in *krm.Asset_DiscoveryStatus) *pb.Asset_DiscoveryStatus {
	if in == nil {
		return nil
	}
	out := &pb.Asset_DiscoveryStatus{}
	out.State = direct.Enum_ToProto[pb.Asset_DiscoveryStatus_State](mapCtx, in.State)
	out.Message = direct.ValueOf(in.Message)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.LastRunTime = direct.StringTimestamp_ToProto(mapCtx, in.LastRunTime)
	out.Stats = Asset_DiscoveryStatus_Stats_ToProto(mapCtx, in.Stats)
	out.LastRunDuration = direct.StringDuration_ToProto(mapCtx, in.LastRunDuration)
	return out
}
func Asset_DiscoveryStatus_Stats_FromProto(mapCtx *direct.MapContext, in *pb.Asset_DiscoveryStatus_Stats) *krm.Asset_DiscoveryStatus_Stats {
	if in == nil {
		return nil
	}
	out := &krm.Asset_DiscoveryStatus_Stats{}
	out.DataItems = direct.LazyPtr(in.GetDataItems())
	out.DataSize = direct.LazyPtr(in.GetDataSize())
	out.Tables = direct.LazyPtr(in.GetTables())
	out.Filesets = direct.LazyPtr(in.GetFilesets())
	return out
}
func Asset_DiscoveryStatus_Stats_ToProto(mapCtx *direct.MapContext, in *krm.Asset_DiscoveryStatus_Stats) *pb.Asset_DiscoveryStatus_Stats {
	if in == nil {
		return nil
	}
	out := &pb.Asset_DiscoveryStatus_Stats{}
	out.DataItems = direct.ValueOf(in.DataItems)
	out.DataSize = direct.ValueOf(in.DataSize)
	out.Tables = direct.ValueOf(in.Tables)
	out.Filesets = direct.ValueOf(in.Filesets)
	return out
}
func Asset_ResourceSpec_FromProto(mapCtx *direct.MapContext, in *pb.Asset_ResourceSpec) *krm.Asset_ResourceSpec {
	if in == nil {
		return nil
	}
	out := &krm.Asset_ResourceSpec{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.ReadAccessMode = direct.Enum_FromProto(mapCtx, in.GetReadAccessMode())
	return out
}
func Asset_ResourceSpec_ToProto(mapCtx *direct.MapContext, in *krm.Asset_ResourceSpec) *pb.Asset_ResourceSpec {
	if in == nil {
		return nil
	}
	out := &pb.Asset_ResourceSpec{}
	out.Name = direct.ValueOf(in.Name)
	out.Type = direct.Enum_ToProto[pb.Asset_ResourceSpec_Type](mapCtx, in.Type)
	out.ReadAccessMode = direct.Enum_ToProto[pb.Asset_ResourceSpec_AccessMode](mapCtx, in.ReadAccessMode)
	return out
}
func Asset_ResourceStatus_FromProto(mapCtx *direct.MapContext, in *pb.Asset_ResourceStatus) *krm.Asset_ResourceStatus {
	if in == nil {
		return nil
	}
	out := &krm.Asset_ResourceStatus{}
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Message = direct.LazyPtr(in.GetMessage())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: ManagedAccessIdentity
	return out
}
func Asset_ResourceStatus_ToProto(mapCtx *direct.MapContext, in *krm.Asset_ResourceStatus) *pb.Asset_ResourceStatus {
	if in == nil {
		return nil
	}
	out := &pb.Asset_ResourceStatus{}
	out.State = direct.Enum_ToProto[pb.Asset_ResourceStatus_State](mapCtx, in.State)
	out.Message = direct.ValueOf(in.Message)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: ManagedAccessIdentity
	return out
}
func Asset_ResourceStatusObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Asset_ResourceStatus) *krm.Asset_ResourceStatusObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Asset_ResourceStatusObservedState{}
	// MISSING: State
	// MISSING: Message
	// MISSING: UpdateTime
	out.ManagedAccessIdentity = direct.LazyPtr(in.GetManagedAccessIdentity())
	return out
}
func Asset_ResourceStatusObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Asset_ResourceStatusObservedState) *pb.Asset_ResourceStatus {
	if in == nil {
		return nil
	}
	out := &pb.Asset_ResourceStatus{}
	// MISSING: State
	// MISSING: Message
	// MISSING: UpdateTime
	out.ManagedAccessIdentity = direct.ValueOf(in.ManagedAccessIdentity)
	return out
}
func Asset_SecurityStatus_FromProto(mapCtx *direct.MapContext, in *pb.Asset_SecurityStatus) *krm.Asset_SecurityStatus {
	if in == nil {
		return nil
	}
	out := &krm.Asset_SecurityStatus{}
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Message = direct.LazyPtr(in.GetMessage())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func Asset_SecurityStatus_ToProto(mapCtx *direct.MapContext, in *krm.Asset_SecurityStatus) *pb.Asset_SecurityStatus {
	if in == nil {
		return nil
	}
	out := &pb.Asset_SecurityStatus{}
	out.State = direct.Enum_ToProto[pb.Asset_SecurityStatus_State](mapCtx, in.State)
	out.Message = direct.ValueOf(in.Message)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func DataplexAssetObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Asset) *krm.DataplexAssetObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataplexAssetObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: State
	// MISSING: ResourceSpec
	// MISSING: ResourceStatus
	// MISSING: SecurityStatus
	// MISSING: DiscoverySpec
	// MISSING: DiscoveryStatus
	return out
}
func DataplexAssetObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataplexAssetObservedState) *pb.Asset {
	if in == nil {
		return nil
	}
	out := &pb.Asset{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: State
	// MISSING: ResourceSpec
	// MISSING: ResourceStatus
	// MISSING: SecurityStatus
	// MISSING: DiscoverySpec
	// MISSING: DiscoveryStatus
	return out
}
func DataplexAssetSpec_FromProto(mapCtx *direct.MapContext, in *pb.Asset) *krm.DataplexAssetSpec {
	if in == nil {
		return nil
	}
	out := &krm.DataplexAssetSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: State
	// MISSING: ResourceSpec
	// MISSING: ResourceStatus
	// MISSING: SecurityStatus
	// MISSING: DiscoverySpec
	// MISSING: DiscoveryStatus
	return out
}
func DataplexAssetSpec_ToProto(mapCtx *direct.MapContext, in *krm.DataplexAssetSpec) *pb.Asset {
	if in == nil {
		return nil
	}
	out := &pb.Asset{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: Labels
	// MISSING: Description
	// MISSING: State
	// MISSING: ResourceSpec
	// MISSING: ResourceStatus
	// MISSING: SecurityStatus
	// MISSING: DiscoverySpec
	// MISSING: DiscoveryStatus
	return out
}
