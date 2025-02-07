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
	pb "cloud.google.com/go/bigquery/dataexchange/apiv1beta1/dataexchangepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/bigquery/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
)
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
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.PrimaryContact = direct.LazyPtr(in.GetPrimaryContact())
	out.Documentation = direct.LazyPtr(in.GetDocumentation())
	// MISSING: State
	out.Icon = in.GetIcon()
	out.DataProvider = DataProvider_FromProto(mapCtx, in.GetDataProvider())
	out.Categories = direct.EnumSlice_FromProto(mapCtx, in.Categories)
	out.Publisher = Publisher_FromProto(mapCtx, in.GetPublisher())
	out.RequestAccess = direct.LazyPtr(in.GetRequestAccess())
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
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.PrimaryContact = direct.ValueOf(in.PrimaryContact)
	out.Documentation = direct.ValueOf(in.Documentation)
	// MISSING: State
	out.Icon = in.Icon
	out.DataProvider = DataProvider_ToProto(mapCtx, in.DataProvider)
	out.Categories = direct.EnumSlice_ToProto[pb.Listing_Category](mapCtx, in.Categories)
	out.Publisher = Publisher_ToProto(mapCtx, in.Publisher)
	out.RequestAccess = direct.ValueOf(in.RequestAccess)
	return out
}
func ListingObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Listing) *krm.ListingObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ListingObservedState{}
	// MISSING: BigqueryDataset
	out.Name = direct.LazyPtr(in.GetName())
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: PrimaryContact
	// MISSING: Documentation
	out.State = direct.Enum_FromProto(mapCtx, in.GetState())
	// MISSING: Icon
	// MISSING: DataProvider
	// MISSING: Categories
	// MISSING: Publisher
	// MISSING: RequestAccess
	return out
}
func ListingObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ListingObservedState) *pb.Listing {
	if in == nil {
		return nil
	}
	out := &pb.Listing{}
	// MISSING: BigqueryDataset
	out.Name = direct.ValueOf(in.Name)
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: PrimaryContact
	// MISSING: Documentation
	out.State = direct.Enum_ToProto[pb.Listing_State](mapCtx, in.State)
	// MISSING: Icon
	// MISSING: DataProvider
	// MISSING: Categories
	// MISSING: Publisher
	// MISSING: RequestAccess
	return out
}
func Listing_BigQueryDatasetSource_FromProto(mapCtx *direct.MapContext, in *pb.Listing_BigQueryDatasetSource) *krm.Listing_BigQueryDatasetSource {
	if in == nil {
		return nil
	}
	out := &krm.Listing_BigQueryDatasetSource{}
	out.Dataset = direct.LazyPtr(in.GetDataset())
	return out
}
func Listing_BigQueryDatasetSource_ToProto(mapCtx *direct.MapContext, in *krm.Listing_BigQueryDatasetSource) *pb.Listing_BigQueryDatasetSource {
	if in == nil {
		return nil
	}
	out := &pb.Listing_BigQueryDatasetSource{}
	out.Dataset = direct.ValueOf(in.Dataset)
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
