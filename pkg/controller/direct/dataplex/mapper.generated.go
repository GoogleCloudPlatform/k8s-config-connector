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

func Content_FromProto(mapCtx *direct.MapContext, in *pb.Content) *krmv1alpha1.Content {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Content{}
	// MISSING: Name
	// MISSING: Uid
	out.Path = direct.LazyPtr(in.GetPath())
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.LazyPtr(in.GetDescription())
	out.DataText = direct.LazyPtr(in.GetDataText())
	out.SQLScript = Content_SQLScript_FromProto(mapCtx, in.GetSqlScript())
	out.Notebook = Content_Notebook_FromProto(mapCtx, in.GetNotebook())
	return out
}
func Content_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Content) *pb.Content {
	if in == nil {
		return nil
	}
	out := &pb.Content{}
	// MISSING: Name
	// MISSING: Uid
	out.Path = direct.ValueOf(in.Path)
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.ValueOf(in.Description)
	if oneof := Content_DataText_ToProto(mapCtx, in.DataText); oneof != nil {
		out.Data = oneof
	}
	if oneof := Content_SQLScript_ToProto(mapCtx, in.SQLScript); oneof != nil {
		out.Content = &pb.Content_SqlScript_{SqlScript: oneof}
	}
	if oneof := Content_Notebook_ToProto(mapCtx, in.Notebook); oneof != nil {
		out.Content = &pb.Content_Notebook_{Notebook: oneof}
	}
	return out
}
func ContentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Content) *krmv1alpha1.ContentObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.ContentObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.Uid = direct.LazyPtr(in.GetUid())
	// MISSING: Path
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: Description
	// MISSING: DataText
	// MISSING: SQLScript
	// MISSING: Notebook
	return out
}
func ContentObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.ContentObservedState) *pb.Content {
	if in == nil {
		return nil
	}
	out := &pb.Content{}
	out.Name = direct.ValueOf(in.Name)
	out.Uid = direct.ValueOf(in.Uid)
	// MISSING: Path
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: Description
	// MISSING: DataText
	// MISSING: SQLScript
	// MISSING: Notebook
	return out
}
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
func Environment_FromProto(mapCtx *direct.MapContext, in *pb.Environment) *krmv1alpha1.Environment {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Environment{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: State
	out.InfrastructureSpec = Environment_InfrastructureSpec_FromProto(mapCtx, in.GetInfrastructureSpec())
	out.SessionSpec = Environment_SessionSpec_FromProto(mapCtx, in.GetSessionSpec())
	// MISSING: SessionStatus
	// MISSING: Endpoints
	return out
}
func Environment_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Environment) *pb.Environment {
	if in == nil {
		return nil
	}
	out := &pb.Environment{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.ValueOf(in.Description)
	// MISSING: State
	out.InfrastructureSpec = Environment_InfrastructureSpec_ToProto(mapCtx, in.InfrastructureSpec)
	out.SessionSpec = Environment_SessionSpec_ToProto(mapCtx, in.SessionSpec)
	// MISSING: SessionStatus
	// MISSING: Endpoints
	return out
}
func EnvironmentObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Environment) *krmv1alpha1.EnvironmentObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.EnvironmentObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: DisplayName
	out.Uid = direct.LazyPtr(in.GetUid())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: Description
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: InfrastructureSpec
	// MISSING: SessionSpec
	out.SessionStatus = Environment_SessionStatusObservedState_FromProto(mapCtx, in.GetSessionStatus())
	out.Endpoints = Environment_EndpointsObservedState_FromProto(mapCtx, in.GetEndpoints())
	return out
}
func EnvironmentObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.EnvironmentObservedState) *pb.Environment {
	if in == nil {
		return nil
	}
	out := &pb.Environment{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: DisplayName
	out.Uid = direct.ValueOf(in.Uid)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: Description
	out.State = direct.Enum_ToProto[pb.State](mapCtx, in.State)
	// MISSING: InfrastructureSpec
	// MISSING: SessionSpec
	out.SessionStatus = Environment_SessionStatusObservedState_ToProto(mapCtx, in.SessionStatus)
	out.Endpoints = Environment_EndpointsObservedState_ToProto(mapCtx, in.Endpoints)
	return out
}
func Environment_Endpoints_FromProto(mapCtx *direct.MapContext, in *pb.Environment_Endpoints) *krmv1alpha1.Environment_Endpoints {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Environment_Endpoints{}
	// MISSING: Notebooks
	// MISSING: SQL
	return out
}
func Environment_Endpoints_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Environment_Endpoints) *pb.Environment_Endpoints {
	if in == nil {
		return nil
	}
	out := &pb.Environment_Endpoints{}
	// MISSING: Notebooks
	// MISSING: SQL
	return out
}
func Environment_SessionStatus_FromProto(mapCtx *direct.MapContext, in *pb.Environment_SessionStatus) *krmv1alpha1.Environment_SessionStatus {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Environment_SessionStatus{}
	// MISSING: Active
	return out
}
func Environment_SessionStatus_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Environment_SessionStatus) *pb.Environment_SessionStatus {
	if in == nil {
		return nil
	}
	out := &pb.Environment_SessionStatus{}
	// MISSING: Active
	return out
}
func Lake_FromProto(mapCtx *direct.MapContext, in *pb.Lake) *krmv1alpha1.Lake {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Lake{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: State
	// MISSING: ServiceAccount
	out.Metastore = Lake_Metastore_FromProto(mapCtx, in.GetMetastore())
	// MISSING: AssetStatus
	// MISSING: MetastoreStatus
	return out
}
func Lake_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Lake) *pb.Lake {
	if in == nil {
		return nil
	}
	out := &pb.Lake{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	// MISSING: Uid
	// MISSING: CreateTime
	// MISSING: UpdateTime
	out.Labels = in.Labels
	out.Description = direct.ValueOf(in.Description)
	// MISSING: State
	// MISSING: ServiceAccount
	out.Metastore = Lake_Metastore_ToProto(mapCtx, in.Metastore)
	// MISSING: AssetStatus
	// MISSING: MetastoreStatus
	return out
}
func LakeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Lake) *krmv1alpha1.LakeObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.LakeObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: DisplayName
	out.Uid = direct.LazyPtr(in.GetUid())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	// MISSING: Labels
	// MISSING: Description
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.ServiceAccount = direct.LazyPtr(in.GetServiceAccount())
	// MISSING: Metastore
	out.AssetStatus = AssetStatus_FromProto(mapCtx, in.GetAssetStatus())
	out.MetastoreStatus = Lake_MetastoreStatus_FromProto(mapCtx, in.GetMetastoreStatus())
	return out
}
func LakeObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.LakeObservedState) *pb.Lake {
	if in == nil {
		return nil
	}
	out := &pb.Lake{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: DisplayName
	out.Uid = direct.ValueOf(in.Uid)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// MISSING: Labels
	// MISSING: Description
	out.State = direct.Enum_ToProto[pb.State](mapCtx, in.State)
	out.ServiceAccount = direct.ValueOf(in.ServiceAccount)
	// MISSING: Metastore
	out.AssetStatus = AssetStatus_ToProto(mapCtx, in.AssetStatus)
	out.MetastoreStatus = Lake_MetastoreStatus_ToProto(mapCtx, in.MetastoreStatus)
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
