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

package privatecatalog

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/privatecatalog/apiv1beta1/privatecatalogpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/privatecatalog/v1alpha1"
)
func Catalog_FromProto(mapCtx *direct.MapContext, in *pb.Catalog) *krm.Catalog {
	if in == nil {
		return nil
	}
	out := &krm.Catalog{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func Catalog_ToProto(mapCtx *direct.MapContext, in *krm.Catalog) *pb.Catalog {
	if in == nil {
		return nil
	}
	out := &pb.Catalog{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func CatalogObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Catalog) *krm.CatalogObservedState {
	if in == nil {
		return nil
	}
	out := &krm.CatalogObservedState{}
	out.Name = direct.LazyPtr(in.GetName())
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	return out
}
func CatalogObservedState_ToProto(mapCtx *direct.MapContext, in *krm.CatalogObservedState) *pb.Catalog {
	if in == nil {
		return nil
	}
	out := &pb.Catalog{}
	out.Name = direct.ValueOf(in.Name)
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	return out
}
func PrivatecatalogCatalogObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Catalog) *krm.PrivatecatalogCatalogObservedState {
	if in == nil {
		return nil
	}
	out := &krm.PrivatecatalogCatalogObservedState{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func PrivatecatalogCatalogObservedState_ToProto(mapCtx *direct.MapContext, in *krm.PrivatecatalogCatalogObservedState) *pb.Catalog {
	if in == nil {
		return nil
	}
	out := &pb.Catalog{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func PrivatecatalogCatalogSpec_FromProto(mapCtx *direct.MapContext, in *pb.Catalog) *krm.PrivatecatalogCatalogSpec {
	if in == nil {
		return nil
	}
	out := &krm.PrivatecatalogCatalogSpec{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
func PrivatecatalogCatalogSpec_ToProto(mapCtx *direct.MapContext, in *krm.PrivatecatalogCatalogSpec) *pb.Catalog {
	if in == nil {
		return nil
	}
	out := &pb.Catalog{}
	// MISSING: Name
	// MISSING: DisplayName
	// MISSING: Description
	// MISSING: CreateTime
	// MISSING: UpdateTime
	return out
}
