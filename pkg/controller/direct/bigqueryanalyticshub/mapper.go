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

package bigqueryanalyticshub

import (
	pb "cloud.google.com/go/bigquery/analyticshub/apiv1/analyticshubpb"
	krmv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigqueryanalyticshub/v1alpha1"
	krmv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigqueryanalyticshub/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func BigQueryAnalyticsHubDataExchangeObservedState_FromProto(mapCtx *direct.MapContext, in *pb.DataExchange) *krmv1beta1.BigQueryAnalyticsHubDataExchangeObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.BigQueryAnalyticsHubDataExchangeObservedState{}
	out.ListingCount = direct.LazyPtr(int64(in.GetListingCount()))
	// MISSING: SharingEnvironmentConfig // not yet
	return out
}

func BigQueryAnalyticsHubDataExchangeSpec_FromProto(mapCtx *direct.MapContext, in *pb.DataExchange) *krmv1beta1.BigQueryAnalyticsHubDataExchangeSpec {
	if in == nil {
		return nil
	}
	out := &krmv1beta1.BigQueryAnalyticsHubDataExchangeSpec{}
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.PrimaryContact = direct.LazyPtr(in.GetPrimaryContact())
	out.Documentation = direct.LazyPtr(in.GetDocumentation())
	// s := string(in.GetIcon())
	// out.Icon = &s // not yet
	// MISSING: SharingEnvironmentConfig // not yet
	out.DiscoveryType = direct.Enum_FromProto(mapCtx, in.GetDiscoveryType())
	return out
}
func BigQueryAnalyticsHubDataExchangeSpec_ToProto(mapCtx *direct.MapContext, in *krmv1beta1.BigQueryAnalyticsHubDataExchangeSpec) *pb.DataExchange {
	if in == nil {
		return nil
	}

	out := &pb.DataExchange{}
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.PrimaryContact = direct.ValueOf(in.PrimaryContact)
	out.Documentation = direct.ValueOf(in.Documentation)
	// out.Icon = []byte(direct.ValueOf(in.Icon)) // not yet
	// MISSING: SharingEnvironmentConfig // not yet
	dtype := direct.Enum_ToProto[pb.DiscoveryType](mapCtx, in.DiscoveryType)
	out.DiscoveryType = &dtype

	return out
}

func Categories_FromProto(mapCtx *direct.MapContext, in []pb.Listing_Category) []string {
	if in == nil {
		return nil
	}
	ret := []string{}
	for _, v := range in {
		toProto := direct.Enum_FromProto(mapCtx, v)
		if toProto != nil {
			ret = append(ret, *toProto)
		}
	}

	return ret
}

func BigQueryAnalyticsHubListingSpec_FromProto(mapCtx *direct.MapContext, in *pb.Listing) *krmv1alpha1.BigQueryAnalyticsHubListingSpec {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.BigQueryAnalyticsHubListingSpec{}

	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.PrimaryContact = direct.LazyPtr(in.GetPrimaryContact())
	out.Documentation = direct.LazyPtr(in.GetDocumentation())
	// MISSING: Icon // not yet
	out.DataProvider = DataProvider_FromProto(mapCtx, in.GetDataProvider())
	out.Categories = Categories_FromProto(mapCtx, in.Categories)
	out.Publisher = Publisher_FromProto(mapCtx, in.GetPublisher())
	out.RequestAccess = direct.LazyPtr(in.GetRequestAccess())
	// MISSING: RestrictedExportConfig // not yet
	out.DiscoveryType = direct.Enum_FromProto(mapCtx, in.GetDiscoveryType())

	out.Source = &krmv1alpha1.Source{
		// TODO(KCC): in the future enforce mutual exclusion / one of b/w BigQueryDatasetSource and PubSubTopicSource
		BigQueryDatasetSource: Listing_BigQueryDatasetSource_FromProto(mapCtx, in.GetBigqueryDataset()),
	}
	return out
}

func Listing_BigQueryDatasetSource_FromProto(mapCtx *direct.MapContext, in *pb.Listing_BigQueryDatasetSource) *krmv1alpha1.BigQueryDatasetSource {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.BigQueryDatasetSource{}
	if out.Dataset != nil {
		out.Dataset = &refs.BigQueryDatasetRef{
			External: in.Dataset,
		}
	}

	out.SelectedResources = direct.Slice_FromProto(mapCtx, in.SelectedResources, Listing_BigQueryDatasetSource_SelectedResource_FromProto)
	out.RestrictedExportPolicy = Listing_BigQueryDatasetSource_RestrictedExportPolicy_FromProto(mapCtx, in.GetRestrictedExportPolicy())
	return out
}

func Listing_BigQueryDatasetSource_SelectedResource_FromProto(mapCtx *direct.MapContext, in *pb.Listing_BigQueryDatasetSource_SelectedResource) *krmv1alpha1.SelectedResource {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.SelectedResource{}
	if in.GetTable() != "" {
		out.TableRef = &refs.BigQueryTableRef{
			External: in.GetTable(),
		}
	}

	return out
}
func Listing_BigQueryDatasetSource_SelectedResource_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.SelectedResource) *pb.Listing_BigQueryDatasetSource_SelectedResource {
	if in == nil {
		return nil
	}
	out := &pb.Listing_BigQueryDatasetSource_SelectedResource{}
	if in.TableRef != nil {
		out.Resource = &pb.Listing_BigQueryDatasetSource_SelectedResource_Table{
			Table: in.TableRef.External,
		}
	}

	return out
}

func Listing_BigQueryDatasetSource_RestrictedExportPolicy_FromProto(mapCtx *direct.MapContext, in *pb.Listing_BigQueryDatasetSource_RestrictedExportPolicy) *krmv1alpha1.RestrictedExportPolicy {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.RestrictedExportPolicy{}
	if in.GetEnabled() != nil {
		out.Enabled = direct.LazyPtr(in.GetEnabled().GetValue())
	}
	if in.GetRestrictDirectTableAccess() != nil {
		out.RestrictDirectTableAccess = direct.LazyPtr(in.GetRestrictDirectTableAccess().GetValue())
	}
	if in.GetRestrictQueryResult() != nil {
		out.RestrictQueryResult = direct.LazyPtr(in.GetRestrictQueryResult().GetValue())
	}

	return out
}
func Listing_BigQueryDatasetSource_RestrictedExportPolicy_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.RestrictedExportPolicy) *pb.Listing_BigQueryDatasetSource_RestrictedExportPolicy {
	if in == nil {
		return nil
	}
	out := &pb.Listing_BigQueryDatasetSource_RestrictedExportPolicy{}
	if in.Enabled != nil {
		out.Enabled = &wrapperspb.BoolValue{Value: *in.Enabled}
	}
	if in.RestrictDirectTableAccess != nil {
		out.RestrictDirectTableAccess = &wrapperspb.BoolValue{Value: *in.RestrictDirectTableAccess}
	}
	if in.RestrictQueryResult != nil {
		out.RestrictQueryResult = &wrapperspb.BoolValue{Value: *in.RestrictQueryResult}
	}

	return out
}

func BigQueryAnalyticsHubListingObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Listing) *krmv1alpha1.BigQueryAnalyticsHubListingObservedState {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.BigQueryAnalyticsHubListingObservedState{}
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: Icon // not yet
	// MISSING: RestrictedExportConfig // not yet
	return out
}

func BigQueryAnalyticsHubListingObservedState_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.BigQueryAnalyticsHubListingObservedState) *pb.Listing {
	if in == nil {
		return nil
	}
	out := &pb.Listing{}
	out.State = direct.Enum_ToProto[pb.Listing_State](mapCtx, in.State)
	// MISSING: Icon // not yet
	// MISSING: RestrictedExportConfig // not yet
	return out
}

func Categories_ToProto(mapCtx *direct.MapContext, in []string) []pb.Listing_Category {
	if in == nil {
		return nil
	}

	ret := []pb.Listing_Category{}
	for _, v := range in {
		ret = append(ret, direct.Enum_ToProto[pb.Listing_Category](mapCtx, &v))
	}

	return ret
}

func BigQueryAnalyticsHubListingSpec_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.BigQueryAnalyticsHubListingSpec) *pb.Listing {
	if in == nil {
		return nil
	}
	out := &pb.Listing{}

	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.PrimaryContact = direct.ValueOf(in.PrimaryContact)
	out.Documentation = direct.ValueOf(in.Documentation)
	// MISSING: Icon // not yet
	out.DataProvider = DataProvider_ToProto(mapCtx, in.DataProvider)
	out.Categories = Categories_ToProto(mapCtx, in.Categories)
	out.Publisher = Publisher_ToProto(mapCtx, in.Publisher)
	out.RequestAccess = direct.ValueOf(in.RequestAccess)
	// MISSING: RestrictedExportConfig // not yet

	dtype := direct.Enum_ToProto[pb.DiscoveryType](mapCtx, in.DiscoveryType)
	out.DiscoveryType = &dtype

	// todo acpana do source
	return out
}

func DataProvider_FromProto(mapCtx *direct.MapContext, in *pb.DataProvider) *krmv1alpha1.DataProvider {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.DataProvider{}
	out.Name = direct.LazyPtr(in.GetName())
	out.PrimaryContact = direct.LazyPtr(in.GetPrimaryContact())
	return out
}

func DataProvider_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.DataProvider) *pb.DataProvider {
	if in == nil {
		return nil
	}
	out := &pb.DataProvider{}
	out.Name = direct.ValueOf(in.Name)
	out.PrimaryContact = direct.ValueOf(in.PrimaryContact)
	return out
}

func Publisher_FromProto(mapCtx *direct.MapContext, in *pb.Publisher) *krmv1alpha1.Publisher {
	if in == nil {
		return nil
	}
	out := &krmv1alpha1.Publisher{}
	out.Name = direct.LazyPtr(in.GetName())
	out.PrimaryContact = direct.LazyPtr(in.GetPrimaryContact())
	return out
}

func Publisher_ToProto(mapCtx *direct.MapContext, in *krmv1alpha1.Publisher) *pb.Publisher {
	if in == nil {
		return nil
	}
	out := &pb.Publisher{}
	out.Name = direct.ValueOf(in.Name)
	out.PrimaryContact = direct.ValueOf(in.PrimaryContact)
	return out
}
