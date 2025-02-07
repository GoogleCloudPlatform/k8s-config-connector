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

package bigquery

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/bigquery/analyticshub/apiv1/analyticshubpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigqueryanalyticshub/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquery/v1alpha1"
)
func BigQueryAnalyticsHubDataExchangeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataExchange) *krm.BigQueryAnalyticsHubDataExchangeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryAnalyticsHubDataExchangeObservedState{}
	// MISSING: Name
	out.ListingCount = direct.LazyPtr(in.GetListingCount())
	// MISSING: Icon
	// MISSING: SharingEnvironmentConfig
	return out
}
func BigQueryAnalyticsHubDataExchangeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryAnalyticsHubDataExchangeObservedState) *pb.DataExchange {
	if in == nil {
		return nil
	}
	out := &pb.DataExchange{}
	// MISSING: Name
	out.ListingCount = direct.ValueOf(in.ListingCount)
	// MISSING: Icon
	// MISSING: SharingEnvironmentConfig
	return out
}
func BigQueryAnalyticsHubDataExchangeSpec_FromProto(mapCtx *direct.MapContext, in *pb.DataExchange) *krm.BigQueryAnalyticsHubDataExchangeSpec {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryAnalyticsHubDataExchangeSpec{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.PrimaryContact = direct.LazyPtr(in.GetPrimaryContact())
	out.Documentation = direct.LazyPtr(in.GetDocumentation())
	// MISSING: Icon
	// MISSING: SharingEnvironmentConfig
	out.DiscoveryType = direct.Enum_FromProto(mapCtx, in.GetDiscoveryType())
	return out
}
func BigQueryAnalyticsHubDataExchangeSpec_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryAnalyticsHubDataExchangeSpec) *pb.DataExchange {
	if in == nil {
		return nil
	}
	out := &pb.DataExchange{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.PrimaryContact = direct.ValueOf(in.PrimaryContact)
	out.Documentation = direct.ValueOf(in.Documentation)
	// MISSING: Icon
	// MISSING: SharingEnvironmentConfig
	if oneof := BigQueryAnalyticsHubDataExchangeSpec_DiscoveryType_ToProto(mapCtx, in.DiscoveryType); oneof != nil {
		out.DiscoveryType = oneof
	}
	return out
}
func BigQueryAnalyticsHubListingObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Listing) *krm.BigQueryAnalyticsHubListingObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryAnalyticsHubListingObservedState{}
	// MISSING: BigqueryDataset
	// MISSING: Name
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: Icon
	// MISSING: RestrictedExportConfig
	return out
}
func BigQueryAnalyticsHubListingObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryAnalyticsHubListingObservedState) *pb.Listing {
	if in == nil {
		return nil
	}
	out := &pb.Listing{}
	// MISSING: BigqueryDataset
	// MISSING: Name
	out.State = direct.Enum_ToProto[pb.Listing_State](mapCtx, in.State)
	// MISSING: Icon
	// MISSING: RestrictedExportConfig
	return out
}
func BigQueryAnalyticsHubListingSpec_FromProto(mapCtx *direct.MapContext, in *pb.Listing) *krm.BigQueryAnalyticsHubListingSpec {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryAnalyticsHubListingSpec{}
	// MISSING: BigqueryDataset
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.PrimaryContact = direct.LazyPtr(in.GetPrimaryContact())
	out.Documentation = direct.LazyPtr(in.GetDocumentation())
	// MISSING: Icon
	out.DataProvider = DataProvider_FromProto(mapCtx, in.GetDataProvider())
	out.Categories = direct.EnumSlice_FromProto(mapCtx, in.Categories)
	out.Publisher = Publisher_FromProto(mapCtx, in.GetPublisher())
	out.RequestAccess = direct.LazyPtr(in.GetRequestAccess())
	// MISSING: RestrictedExportConfig
	out.DiscoveryType = direct.Enum_FromProto(mapCtx, in.GetDiscoveryType())
	return out
}
func BigQueryAnalyticsHubListingSpec_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryAnalyticsHubListingSpec) *pb.Listing {
	if in == nil {
		return nil
	}
	out := &pb.Listing{}
	// MISSING: BigqueryDataset
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.PrimaryContact = direct.ValueOf(in.PrimaryContact)
	out.Documentation = direct.ValueOf(in.Documentation)
	// MISSING: Icon
	out.DataProvider = DataProvider_ToProto(mapCtx, in.DataProvider)
	out.Categories = direct.EnumSlice_ToProto[pb.Listing_Category](mapCtx, in.Categories)
	out.Publisher = Publisher_ToProto(mapCtx, in.Publisher)
	out.RequestAccess = direct.ValueOf(in.RequestAccess)
	// MISSING: RestrictedExportConfig
	if oneof := BigQueryAnalyticsHubListingSpec_DiscoveryType_ToProto(mapCtx, in.DiscoveryType); oneof != nil {
		out.DiscoveryType = oneof
	}
	return out
}
func BigqueryDataExchangeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataExchange) *krm.BigqueryDataExchangeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BigqueryDataExchangeObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: PrimaryContact
	// MISSING: Documentation
	// MISSING: ListingCount
	// MISSING: Icon
	// MISSING: SharingEnvironmentConfig
	// MISSING: DiscoveryType
	return out
}
func BigqueryDataExchangeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BigqueryDataExchangeObservedState) *pb.DataExchange {
	if in == nil {
		return nil
	}
	out := &pb.DataExchange{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: PrimaryContact
	// MISSING: Documentation
	// MISSING: ListingCount
	// MISSING: Icon
	// MISSING: SharingEnvironmentConfig
	// MISSING: DiscoveryType
	return out
}
func BigqueryDataExchangeSpec_FromProto(mapCtx *direct.MapContext, in *pb.DataExchange) *krm.BigqueryDataExchangeSpec {
	if in == nil {
		return nil
	}
	out := &krm.BigqueryDataExchangeSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: PrimaryContact
	// MISSING: Documentation
	// MISSING: ListingCount
	// MISSING: Icon
	// MISSING: SharingEnvironmentConfig
	// MISSING: DiscoveryType
	return out
}
func BigqueryDataExchangeSpec_ToProto(mapCtx *direct.MapContext, in *krm.BigqueryDataExchangeSpec) *pb.DataExchange {
	if in == nil {
		return nil
	}
	out := &pb.DataExchange{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: PrimaryContact
	// MISSING: Documentation
	// MISSING: ListingCount
	// MISSING: Icon
	// MISSING: SharingEnvironmentConfig
	// MISSING: DiscoveryType
	return out
}
func DataExchange_FromProto(mapCtx *direct.MapContext, in *pb.DataExchange) *krm.DataExchange {
	if in == nil {
		return nil
	}
	out := &krm.DataExchange{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.PrimaryContact = direct.LazyPtr(in.GetPrimaryContact())
	out.Documentation = direct.LazyPtr(in.GetDocumentation())
	// MISSING: ListingCount
	out.Icon = in.GetIcon()
	out.SharingEnvironmentConfig = SharingEnvironmentConfig_FromProto(mapCtx, in.GetSharingEnvironmentConfig())
	out.DiscoveryType = direct.Enum_FromProto(mapCtx, in.GetDiscoveryType())
	return out
}
func DataExchange_ToProto(mapCtx *direct.MapContext, in *krm.DataExchange) *pb.DataExchange {
	if in == nil {
		return nil
	}
	out := &pb.DataExchange{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.PrimaryContact = direct.ValueOf(in.PrimaryContact)
	out.Documentation = direct.ValueOf(in.Documentation)
	// MISSING: ListingCount
	out.Icon = in.Icon
	out.SharingEnvironmentConfig = SharingEnvironmentConfig_ToProto(mapCtx, in.SharingEnvironmentConfig)
	if oneof := DataExchange_DiscoveryType_ToProto(mapCtx, in.DiscoveryType); oneof != nil {
		out.DiscoveryType = oneof
	}
	return out
}
func DataExchangeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataExchange) *krm.DataExchangeObservedState {
	if in == nil {
		return nil
	}
	out := &krm.DataExchangeObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: PrimaryContact
	// MISSING: Documentation
	out.ListingCount = direct.LazyPtr(in.GetListingCount())
	// MISSING: Icon
	out.SharingEnvironmentConfig = SharingEnvironmentConfigObservedState_FromProto(mapCtx, in.GetSharingEnvironmentConfig())
	// MISSING: DiscoveryType
	return out
}
func DataExchangeObservedState_ToProto(mapCtx *direct.MapContext, in *krm.DataExchangeObservedState) *pb.DataExchange {
	if in == nil {
		return nil
	}
	out := &pb.DataExchange{}
	out.Name = direct.ValueOf(in.Name)
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: PrimaryContact
	// MISSING: Documentation
	out.ListingCount = direct.ValueOf(in.ListingCount)
	// MISSING: Icon
	out.SharingEnvironmentConfig = SharingEnvironmentConfigObservedState_ToProto(mapCtx, in.SharingEnvironmentConfig)
	// MISSING: DiscoveryType
	return out
}
func DataProvider_FromProto(mapCtx *direct.MapContext, in *pb.DataProvider) *krm.DataProvider {
	if in == nil {
		return nil
	}
	out := &krm.DataProvider{}
	out.Name = direct.LazyPtr(in.GetName())
	out.PrimaryContact = direct.LazyPtr(in.GetPrimaryContact())
	return out
}
func DataProvider_ToProto(mapCtx *direct.MapContext, in *krm.DataProvider) *pb.DataProvider {
	if in == nil {
		return nil
	}
	out := &pb.DataProvider{}
	out.Name = direct.ValueOf(in.Name)
	out.PrimaryContact = direct.ValueOf(in.PrimaryContact)
	return out
}
func Listing_FromProto(mapCtx *direct.MapContext, in *pb.Listing) *krm.Listing {
	if in == nil {
		return nil
	}
	out := &krm.Listing{}
	out.BigqueryDataset = Listing_BigQueryDatasetSource_FromProto(mapCtx, in.GetBigqueryDataset())
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.PrimaryContact = direct.LazyPtr(in.GetPrimaryContact())
	out.Documentation = direct.LazyPtr(in.GetDocumentation())
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	out.Icon = in.GetIcon()
	out.DataProvider = DataProvider_FromProto(mapCtx, in.GetDataProvider())
	out.Categories = direct.EnumSlice_FromProto(mapCtx, in.Categories)
	out.Publisher = Publisher_FromProto(mapCtx, in.GetPublisher())
	out.RequestAccess = direct.LazyPtr(in.GetRequestAccess())
	out.RestrictedExportConfig = Listing_RestrictedExportConfig_FromProto(mapCtx, in.GetRestrictedExportConfig())
	out.DiscoveryType = direct.Enum_FromProto(mapCtx, in.GetDiscoveryType())
	return out
}
func Listing_ToProto(mapCtx *direct.MapContext, in *krm.Listing) *pb.Listing {
	if in == nil {
		return nil
	}
	out := &pb.Listing{}
	if oneof := Listing_BigQueryDatasetSource_ToProto(mapCtx, in.BigqueryDataset); oneof != nil {
		out.Source = &pb.Listing_BigqueryDataset{BigqueryDataset: oneof}
	}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.PrimaryContact = direct.ValueOf(in.PrimaryContact)
	out.Documentation = direct.ValueOf(in.Documentation)
	out.State = direct.Enum_ToProto[pb.Listing_State](mapCtx, in.State)
	out.Icon = in.Icon
	out.DataProvider = DataProvider_ToProto(mapCtx, in.DataProvider)
	out.Categories = direct.EnumSlice_ToProto[pb.Listing_Category](mapCtx, in.Categories)
	out.Publisher = Publisher_ToProto(mapCtx, in.Publisher)
	out.RequestAccess = direct.ValueOf(in.RequestAccess)
	out.RestrictedExportConfig = Listing_RestrictedExportConfig_ToProto(mapCtx, in.RestrictedExportConfig)
	if oneof := Listing_DiscoveryType_ToProto(mapCtx, in.DiscoveryType); oneof != nil {
		out.DiscoveryType = oneof
	}
	return out
}
func Listing_BigQueryDatasetSource_FromProto(mapCtx *direct.MapContext, in *pb.Listing_BigQueryDatasetSource) *krm.Listing_BigQueryDatasetSource {
	if in == nil {
		return nil
	}
	out := &krm.Listing_BigQueryDatasetSource{}
	out.Dataset = direct.LazyPtr(in.GetDataset())
	out.SelectedResources = direct.Slice_FromProto(mapCtx, in.SelectedResources, Listing_BigQueryDatasetSource_SelectedResource_FromProto)
	out.RestrictedExportPolicy = Listing_BigQueryDatasetSource_RestrictedExportPolicy_FromProto(mapCtx, in.GetRestrictedExportPolicy())
	return out
}
func Listing_BigQueryDatasetSource_ToProto(mapCtx *direct.MapContext, in *krm.Listing_BigQueryDatasetSource) *pb.Listing_BigQueryDatasetSource {
	if in == nil {
		return nil
	}
	out := &pb.Listing_BigQueryDatasetSource{}
	out.Dataset = direct.ValueOf(in.Dataset)
	out.SelectedResources = direct.Slice_ToProto(mapCtx, in.SelectedResources, Listing_BigQueryDatasetSource_SelectedResource_ToProto)
	out.RestrictedExportPolicy = Listing_BigQueryDatasetSource_RestrictedExportPolicy_ToProto(mapCtx, in.RestrictedExportPolicy)
	return out
}
func Listing_BigQueryDatasetSource_RestrictedExportPolicy_FromProto(mapCtx *direct.MapContext, in *pb.Listing_BigQueryDatasetSource_RestrictedExportPolicy) *krm.Listing_BigQueryDatasetSource_RestrictedExportPolicy {
	if in == nil {
		return nil
	}
	out := &krm.Listing_BigQueryDatasetSource_RestrictedExportPolicy{}
	out.Enabled = direct.BoolValue_FromProto(mapCtx, in.GetEnabled())
	out.RestrictDirectTableAccess = direct.BoolValue_FromProto(mapCtx, in.GetRestrictDirectTableAccess())
	out.RestrictQueryResult = direct.BoolValue_FromProto(mapCtx, in.GetRestrictQueryResult())
	return out
}
func Listing_BigQueryDatasetSource_RestrictedExportPolicy_ToProto(mapCtx *direct.MapContext, in *krm.Listing_BigQueryDatasetSource_RestrictedExportPolicy) *pb.Listing_BigQueryDatasetSource_RestrictedExportPolicy {
	if in == nil {
		return nil
	}
	out := &pb.Listing_BigQueryDatasetSource_RestrictedExportPolicy{}
	out.Enabled = direct.BoolValue_ToProto(mapCtx, in.Enabled)
	out.RestrictDirectTableAccess = direct.BoolValue_ToProto(mapCtx, in.RestrictDirectTableAccess)
	out.RestrictQueryResult = direct.BoolValue_ToProto(mapCtx, in.RestrictQueryResult)
	return out
}
func Listing_BigQueryDatasetSource_SelectedResource_FromProto(mapCtx *direct.MapContext, in *pb.Listing_BigQueryDatasetSource_SelectedResource) *krm.Listing_BigQueryDatasetSource_SelectedResource {
	if in == nil {
		return nil
	}
	out := &krm.Listing_BigQueryDatasetSource_SelectedResource{}
	out.Table = direct.LazyPtr(in.GetTable())
	return out
}
func Listing_BigQueryDatasetSource_SelectedResource_ToProto(mapCtx *direct.MapContext, in *krm.Listing_BigQueryDatasetSource_SelectedResource) *pb.Listing_BigQueryDatasetSource_SelectedResource {
	if in == nil {
		return nil
	}
	out := &pb.Listing_BigQueryDatasetSource_SelectedResource{}
	if oneof := Listing_BigQueryDatasetSource_SelectedResource_Table_ToProto(mapCtx, in.Table); oneof != nil {
		out.Resource = oneof
	}
	return out
}
func Listing_RestrictedExportConfig_FromProto(mapCtx *direct.MapContext, in *pb.Listing_RestrictedExportConfig) *krm.Listing_RestrictedExportConfig {
	if in == nil {
		return nil
	}
	out := &krm.Listing_RestrictedExportConfig{}
	out.Enabled = direct.LazyPtr(in.GetEnabled())
	out.RestrictDirectTableAccess = direct.LazyPtr(in.GetRestrictDirectTableAccess())
	out.RestrictQueryResult = direct.LazyPtr(in.GetRestrictQueryResult())
	return out
}
func Listing_RestrictedExportConfig_ToProto(mapCtx *direct.MapContext, in *krm.Listing_RestrictedExportConfig) *pb.Listing_RestrictedExportConfig {
	if in == nil {
		return nil
	}
	out := &pb.Listing_RestrictedExportConfig{}
	out.Enabled = direct.ValueOf(in.Enabled)
	out.RestrictDirectTableAccess = direct.ValueOf(in.RestrictDirectTableAccess)
	out.RestrictQueryResult = direct.ValueOf(in.RestrictQueryResult)
	return out
}
func Publisher_FromProto(mapCtx *direct.MapContext, in *pb.Publisher) *krm.Publisher {
	if in == nil {
		return nil
	}
	out := &krm.Publisher{}
	out.Name = direct.LazyPtr(in.GetName())
	out.PrimaryContact = direct.LazyPtr(in.GetPrimaryContact())
	return out
}
func Publisher_ToProto(mapCtx *direct.MapContext, in *krm.Publisher) *pb.Publisher {
	if in == nil {
		return nil
	}
	out := &pb.Publisher{}
	out.Name = direct.ValueOf(in.Name)
	out.PrimaryContact = direct.ValueOf(in.PrimaryContact)
	return out
}
func RestrictedExportPolicy_FromProto(mapCtx *direct.MapContext, in *pb.Listing_BigQueryDatasetSource_RestrictedExportPolicy) *krm.RestrictedExportPolicy {
	if in == nil {
		return nil
	}
	out := &krm.RestrictedExportPolicy{}
	out.Enabled = direct.BoolValue_FromProto(mapCtx, in.GetEnabled())
	out.RestrictDirectTableAccess = direct.BoolValue_FromProto(mapCtx, in.GetRestrictDirectTableAccess())
	out.RestrictQueryResult = direct.BoolValue_FromProto(mapCtx, in.GetRestrictQueryResult())
	return out
}
func RestrictedExportPolicy_ToProto(mapCtx *direct.MapContext, in *krm.RestrictedExportPolicy) *pb.Listing_BigQueryDatasetSource_RestrictedExportPolicy {
	if in == nil {
		return nil
	}
	out := &pb.Listing_BigQueryDatasetSource_RestrictedExportPolicy{}
	out.Enabled = direct.BoolValue_ToProto(mapCtx, in.Enabled)
	out.RestrictDirectTableAccess = direct.BoolValue_ToProto(mapCtx, in.RestrictDirectTableAccess)
	out.RestrictQueryResult = direct.BoolValue_ToProto(mapCtx, in.RestrictQueryResult)
	return out
}
func SelectedResource_FromProto(mapCtx *direct.MapContext, in *pb.Listing_BigQueryDatasetSource_SelectedResource) *krm.SelectedResource {
	if in == nil {
		return nil
	}
	out := &krm.SelectedResource{}
	if in.GetTable() != "" {
		out.TableRef = &refs.*refv1beta1.BigQueryTableRef{External: in.GetTable()}
	}
	return out
}
func SelectedResource_ToProto(mapCtx *direct.MapContext, in *krm.SelectedResource) *pb.Listing_BigQueryDatasetSource_SelectedResource {
	if in == nil {
		return nil
	}
	out := &pb.Listing_BigQueryDatasetSource_SelectedResource{}
	if in.TableRef != nil {
		out.Table = in.TableRef.External
	}
	return out
}
func SharingEnvironmentConfig_FromProto(mapCtx *direct.MapContext, in *pb.SharingEnvironmentConfig) *krm.SharingEnvironmentConfig {
	if in == nil {
		return nil
	}
	out := &krm.SharingEnvironmentConfig{}
	out.DefaultExchangeConfig = SharingEnvironmentConfig_DefaultExchangeConfig_FromProto(mapCtx, in.GetDefaultExchangeConfig())
	out.DcrExchangeConfig = SharingEnvironmentConfig_DcrExchangeConfig_FromProto(mapCtx, in.GetDcrExchangeConfig())
	return out
}
func SharingEnvironmentConfig_ToProto(mapCtx *direct.MapContext, in *krm.SharingEnvironmentConfig) *pb.SharingEnvironmentConfig {
	if in == nil {
		return nil
	}
	out := &pb.SharingEnvironmentConfig{}
	if oneof := SharingEnvironmentConfig_DefaultExchangeConfig_ToProto(mapCtx, in.DefaultExchangeConfig); oneof != nil {
		out.Environment = &pb.SharingEnvironmentConfig_DefaultExchangeConfig_{DefaultExchangeConfig: oneof}
	}
	if oneof := SharingEnvironmentConfig_DcrExchangeConfig_ToProto(mapCtx, in.DcrExchangeConfig); oneof != nil {
		out.Environment = &pb.SharingEnvironmentConfig_DcrExchangeConfig_{DcrExchangeConfig: oneof}
	}
	return out
}
func SharingEnvironmentConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SharingEnvironmentConfig) *krm.SharingEnvironmentConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SharingEnvironmentConfigObservedState{}
	// MISSING: DefaultExchangeConfig
	out.DcrExchangeConfig = SharingEnvironmentConfig_DcrExchangeConfigObservedState_FromProto(mapCtx, in.GetDcrExchangeConfig())
	return out
}
func SharingEnvironmentConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SharingEnvironmentConfigObservedState) *pb.SharingEnvironmentConfig {
	if in == nil {
		return nil
	}
	out := &pb.SharingEnvironmentConfig{}
	// MISSING: DefaultExchangeConfig
	if oneof := SharingEnvironmentConfig_DcrExchangeConfigObservedState_ToProto(mapCtx, in.DcrExchangeConfig); oneof != nil {
		out.Environment = &pb.SharingEnvironmentConfig_DcrExchangeConfig_{DcrExchangeConfig: oneof}
	}
	return out
}
func SharingEnvironmentConfig_DcrExchangeConfig_FromProto(mapCtx *direct.MapContext, in *pb.SharingEnvironmentConfig_DcrExchangeConfig) *krm.SharingEnvironmentConfig_DcrExchangeConfig {
	if in == nil {
		return nil
	}
	out := &krm.SharingEnvironmentConfig_DcrExchangeConfig{}
	out.SingleSelectedResourceSharingRestriction = in.SingleSelectedResourceSharingRestriction
	out.SingleLinkedDatasetPerCleanroom = in.SingleLinkedDatasetPerCleanroom
	return out
}
func SharingEnvironmentConfig_DcrExchangeConfig_ToProto(mapCtx *direct.MapContext, in *krm.SharingEnvironmentConfig_DcrExchangeConfig) *pb.SharingEnvironmentConfig_DcrExchangeConfig {
	if in == nil {
		return nil
	}
	out := &pb.SharingEnvironmentConfig_DcrExchangeConfig{}
	out.SingleSelectedResourceSharingRestriction = in.SingleSelectedResourceSharingRestriction
	out.SingleLinkedDatasetPerCleanroom = in.SingleLinkedDatasetPerCleanroom
	return out
}
func SharingEnvironmentConfig_DcrExchangeConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.SharingEnvironmentConfig_DcrExchangeConfig) *krm.SharingEnvironmentConfig_DcrExchangeConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.SharingEnvironmentConfig_DcrExchangeConfigObservedState{}
	out.SingleSelectedResourceSharingRestriction = in.SingleSelectedResourceSharingRestriction
	out.SingleLinkedDatasetPerCleanroom = in.SingleLinkedDatasetPerCleanroom
	return out
}
func SharingEnvironmentConfig_DcrExchangeConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.SharingEnvironmentConfig_DcrExchangeConfigObservedState) *pb.SharingEnvironmentConfig_DcrExchangeConfig {
	if in == nil {
		return nil
	}
	out := &pb.SharingEnvironmentConfig_DcrExchangeConfig{}
	out.SingleSelectedResourceSharingRestriction = in.SingleSelectedResourceSharingRestriction
	out.SingleLinkedDatasetPerCleanroom = in.SingleLinkedDatasetPerCleanroom
	return out
}
func SharingEnvironmentConfig_DefaultExchangeConfig_FromProto(mapCtx *direct.MapContext, in *pb.SharingEnvironmentConfig_DefaultExchangeConfig) *krm.SharingEnvironmentConfig_DefaultExchangeConfig {
	if in == nil {
		return nil
	}
	out := &krm.SharingEnvironmentConfig_DefaultExchangeConfig{}
	return out
}
func SharingEnvironmentConfig_DefaultExchangeConfig_ToProto(mapCtx *direct.MapContext, in *krm.SharingEnvironmentConfig_DefaultExchangeConfig) *pb.SharingEnvironmentConfig_DefaultExchangeConfig {
	if in == nil {
		return nil
	}
	out := &pb.SharingEnvironmentConfig_DefaultExchangeConfig{}
	return out
}
