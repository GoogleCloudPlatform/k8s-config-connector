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
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/retail/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/retail/apiv2beta/retailpb"
)
func Catalog_FromProto(mapCtx *direct.MapContext, in *pb.Catalog) *krm.Catalog {
	if in == nil {
		return nil
	}
	out := &krm.Catalog{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.ProductLevelConfig = ProductLevelConfig_FromProto(mapCtx, in.GetProductLevelConfig())
	out.MerchantCenterLinkingConfig = MerchantCenterLinkingConfig_FromProto(mapCtx, in.GetMerchantCenterLinkingConfig())
	return out
}
func Catalog_ToProto(mapCtx *direct.MapContext, in *krm.Catalog) *pb.Catalog {
	if in == nil {
		return nil
	}
	out := &pb.Catalog{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.ProductLevelConfig = ProductLevelConfig_ToProto(mapCtx, in.ProductLevelConfig)
	out.MerchantCenterLinkingConfig = MerchantCenterLinkingConfig_ToProto(mapCtx, in.MerchantCenterLinkingConfig)
	return out
}
func MerchantCenterFeedFilter_FromProto(mapCtx *direct.MapContext, in *pb.MerchantCenterFeedFilter) *krm.MerchantCenterFeedFilter {
	if in == nil {
		return nil
	}
	out := &krm.MerchantCenterFeedFilter{}
	out.PrimaryFeedID = direct.LazyPtr(in.GetPrimaryFeedId())
	out.PrimaryFeedName = direct.LazyPtr(in.GetPrimaryFeedName())
	return out
}
func MerchantCenterFeedFilter_ToProto(mapCtx *direct.MapContext, in *krm.MerchantCenterFeedFilter) *pb.MerchantCenterFeedFilter {
	if in == nil {
		return nil
	}
	out := &pb.MerchantCenterFeedFilter{}
	out.PrimaryFeedId = direct.ValueOf(in.PrimaryFeedID)
	out.PrimaryFeedName = direct.ValueOf(in.PrimaryFeedName)
	return out
}
func MerchantCenterLink_FromProto(mapCtx *direct.MapContext, in *pb.MerchantCenterLink) *krm.MerchantCenterLink {
	if in == nil {
		return nil
	}
	out := &krm.MerchantCenterLink{}
	out.MerchantCenterAccountID = direct.LazyPtr(in.GetMerchantCenterAccountId())
	out.BranchID = direct.LazyPtr(in.GetBranchId())
	out.Destinations = in.Destinations
	out.RegionCode = direct.LazyPtr(in.GetRegionCode())
	out.LanguageCode = direct.LazyPtr(in.GetLanguageCode())
	out.Feeds = direct.Slice_FromProto(mapCtx, in.Feeds, MerchantCenterFeedFilter_FromProto)
	return out
}
func MerchantCenterLink_ToProto(mapCtx *direct.MapContext, in *krm.MerchantCenterLink) *pb.MerchantCenterLink {
	if in == nil {
		return nil
	}
	out := &pb.MerchantCenterLink{}
	out.MerchantCenterAccountId = direct.ValueOf(in.MerchantCenterAccountID)
	out.BranchId = direct.ValueOf(in.BranchID)
	out.Destinations = in.Destinations
	out.RegionCode = direct.ValueOf(in.RegionCode)
	out.LanguageCode = direct.ValueOf(in.LanguageCode)
	out.Feeds = direct.Slice_ToProto(mapCtx, in.Feeds, MerchantCenterFeedFilter_ToProto)
	return out
}
func MerchantCenterLinkingConfig_FromProto(mapCtx *direct.MapContext, in *pb.MerchantCenterLinkingConfig) *krm.MerchantCenterLinkingConfig {
	if in == nil {
		return nil
	}
	out := &krm.MerchantCenterLinkingConfig{}
	out.Links = direct.Slice_FromProto(mapCtx, in.Links, MerchantCenterLink_FromProto)
	return out
}
func MerchantCenterLinkingConfig_ToProto(mapCtx *direct.MapContext, in *krm.MerchantCenterLinkingConfig) *pb.MerchantCenterLinkingConfig {
	if in == nil {
		return nil
	}
	out := &pb.MerchantCenterLinkingConfig{}
	out.Links = direct.Slice_ToProto(mapCtx, in.Links, MerchantCenterLink_ToProto)
	return out
}
func ProductLevelConfig_FromProto(mapCtx *direct.MapContext, in *pb.ProductLevelConfig) *krm.ProductLevelConfig {
	if in == nil {
		return nil
	}
	out := &krm.ProductLevelConfig{}
	out.IngestionProductType = direct.LazyPtr(in.GetIngestionProductType())
	out.MerchantCenterProductIDField = direct.LazyPtr(in.GetMerchantCenterProductIdField())
	return out
}
func ProductLevelConfig_ToProto(mapCtx *direct.MapContext, in *krm.ProductLevelConfig) *pb.ProductLevelConfig {
	if in == nil {
		return nil
	}
	out := &pb.ProductLevelConfig{}
	out.IngestionProductType = direct.ValueOf(in.IngestionProductType)
	out.MerchantCenterProductIdField = direct.ValueOf(in.MerchantCenterProductIDField)
	return out
}
