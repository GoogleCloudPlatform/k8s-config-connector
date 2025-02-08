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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/retail/apiv2/retailpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/retail/v1alpha1"
)
func Catalog_FromProto(mapCtx *direct.MapContext, in *pb.Catalog) *krm.Catalog {
	if in == nil {
		return nil
	}
	out := &krm.Catalog{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.ProductLevelConfig = ProductLevelConfig_FromProto(mapCtx, in.GetProductLevelConfig())
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
func RetailCatalogObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Catalog) *krm.RetailCatalogObservedState {
	if in == nil {
		return nil
	}
	out := &krm.RetailCatalogObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: ProductLevelConfig
	return out
}
func RetailCatalogObservedState_ToProto(mapCtx *direct.MapContext, in *krm.RetailCatalogObservedState) *pb.Catalog {
	if in == nil {
		return nil
	}
	out := &pb.Catalog{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: ProductLevelConfig
	return out
}
func RetailCatalogSpec_FromProto(mapCtx *direct.MapContext, in *pb.Catalog) *krm.RetailCatalogSpec {
	if in == nil {
		return nil
	}
	out := &krm.RetailCatalogSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: ProductLevelConfig
	return out
}
func RetailCatalogSpec_ToProto(mapCtx *direct.MapContext, in *krm.RetailCatalogSpec) *pb.Catalog {
	if in == nil {
		return nil
	}
	out := &pb.Catalog{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: ProductLevelConfig
	return out
}
