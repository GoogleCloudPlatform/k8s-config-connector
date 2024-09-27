// Copyright 2024 Google LLC
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

package dataexchangelisting

import (
	krm "/usr/local/google/home/acpana/glinux_oss_work/kcc/8/k8s-config-connector/apis/bigqueryanalyticshub/dataexchangelisting/v1alpha1"

	pb "cloud.google.com/go/bigquery/analyticshub/apiv1/analyticshubpb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func BigQueryAnalyticsHubDataExchangeListingObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Listing) *krm.BigQueryAnalyticsHubDataExchangeListingObservedState {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryAnalyticsHubDataExchangeListingObservedState{}
	// MISSING: BigqueryDataset
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: PrimaryContact
	// MISSING: Documentation
	// MISSING: State
	// MISSING: Icon
	// MISSING: DataProvider
	// MISSING: Categories
	// MISSING: Publisher
	// MISSING: RequestAccess
	// MISSING: RestrictedExportConfig
	// MISSING: DiscoveryType
	return out
}
func BigQueryAnalyticsHubDataExchangeListingObservedState_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryAnalyticsHubDataExchangeListingObservedState) *pb.Listing {
	if in == nil {
		return nil
	}
	out := &pb.Listing{}
	// MISSING: BigqueryDataset
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: PrimaryContact
	// MISSING: Documentation
	// MISSING: State
	// MISSING: Icon
	// MISSING: DataProvider
	// MISSING: Categories
	// MISSING: Publisher
	// MISSING: RequestAccess
	// MISSING: RestrictedExportConfig
	// MISSING: DiscoveryType
	return out
}
func BigQueryAnalyticsHubDataExchangeListingSpec_FromProto(mapCtx *direct.MapContext, in *pb.Listing) *krm.BigQueryAnalyticsHubDataExchangeListingSpec {
	if in == nil {
		return nil
	}
	out := &krm.BigQueryAnalyticsHubDataExchangeListingSpec{}
	// MISSING: BigqueryDataset
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: PrimaryContact
	// MISSING: Documentation
	// MISSING: State
	// MISSING: Icon
	// MISSING: DataProvider
	// MISSING: Categories
	// MISSING: Publisher
	// MISSING: RequestAccess
	// MISSING: RestrictedExportConfig
	// MISSING: DiscoveryType
	return out
}
func BigQueryAnalyticsHubDataExchangeListingSpec_ToProto(mapCtx *direct.MapContext, in *krm.BigQueryAnalyticsHubDataExchangeListingSpec) *pb.Listing {
	if in == nil {
		return nil
	}
	out := &pb.Listing{}
	// MISSING: BigqueryDataset
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: PrimaryContact
	// MISSING: Documentation
	// MISSING: State
	// MISSING: Icon
	// MISSING: DataProvider
	// MISSING: Categories
	// MISSING: Publisher
	// MISSING: RequestAccess
	// MISSING: RestrictedExportConfig
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
	out.Enabled = BoolValue_FromProto(mapCtx, in.GetEnabled())
	out.RestrictDirectTableAccess = BoolValue_FromProto(mapCtx, in.GetRestrictDirectTableAccess())
	out.RestrictQueryResult = BoolValue_FromProto(mapCtx, in.GetRestrictQueryResult())
	return out
}
func Listing_BigQueryDatasetSource_RestrictedExportPolicy_ToProto(mapCtx *direct.MapContext, in *krm.Listing_BigQueryDatasetSource_RestrictedExportPolicy) *pb.Listing_BigQueryDatasetSource_RestrictedExportPolicy {
	if in == nil {
		return nil
	}
	out := &pb.Listing_BigQueryDatasetSource_RestrictedExportPolicy{}
	out.Enabled = BoolValue_ToProto(mapCtx, in.Enabled)
	out.RestrictDirectTableAccess = BoolValue_ToProto(mapCtx, in.RestrictDirectTableAccess)
	out.RestrictQueryResult = BoolValue_ToProto(mapCtx, in.RestrictQueryResult)
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
