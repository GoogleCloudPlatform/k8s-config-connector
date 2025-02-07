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

package asset

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	pb "cloud.google.com/go/asset/apiv1p7beta1/assetpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/asset/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)
func Asset_FromProto(mapCtx *direct.MapContext, in *pb.Asset) *krm.Asset {
	if in == nil {
		return nil
	}
	out := &krm.Asset{}
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	out.Name = direct.LazyPtr(in.GetName())
	out.AssetType = direct.LazyPtr(in.GetAssetType())
	out.Resource = Resource_FromProto(mapCtx, in.GetResource())
	out.IamPolicy = Policy_FromProto(mapCtx, in.GetIamPolicy())
	out.OrgPolicy = direct.Slice_FromProto(mapCtx, in.OrgPolicy, Policy_FromProto)
	out.AccessPolicy = AccessPolicy_FromProto(mapCtx, in.GetAccessPolicy())
	out.AccessLevel = AccessLevel_FromProto(mapCtx, in.GetAccessLevel())
	out.ServicePerimeter = ServicePerimeter_FromProto(mapCtx, in.GetServicePerimeter())
	out.RelatedAssets = RelatedAssets_FromProto(mapCtx, in.GetRelatedAssets())
	out.Ancestors = in.Ancestors
	return out
}
func Asset_ToProto(mapCtx *direct.MapContext, in *krm.Asset) *pb.Asset {
	if in == nil {
		return nil
	}
	out := &pb.Asset{}
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	out.Name = direct.ValueOf(in.Name)
	out.AssetType = direct.ValueOf(in.AssetType)
	out.Resource = Resource_ToProto(mapCtx, in.Resource)
	out.IamPolicy = Policy_ToProto(mapCtx, in.IamPolicy)
	out.OrgPolicy = direct.Slice_ToProto(mapCtx, in.OrgPolicy, Policy_ToProto)
	if oneof := AccessPolicy_ToProto(mapCtx, in.AccessPolicy); oneof != nil {
		out.AccessContextPolicy = &pb.Asset_AccessPolicy{AccessPolicy: oneof}
	}
	if oneof := AccessLevel_ToProto(mapCtx, in.AccessLevel); oneof != nil {
		out.AccessContextPolicy = &pb.Asset_AccessLevel{AccessLevel: oneof}
	}
	if oneof := ServicePerimeter_ToProto(mapCtx, in.ServicePerimeter); oneof != nil {
		out.AccessContextPolicy = &pb.Asset_ServicePerimeter{ServicePerimeter: oneof}
	}
	out.RelatedAssets = RelatedAssets_ToProto(mapCtx, in.RelatedAssets)
	out.Ancestors = in.Ancestors
	return out
}
func RelatedAsset_FromProto(mapCtx *direct.MapContext, in *pb.RelatedAsset) *krm.RelatedAsset {
	if in == nil {
		return nil
	}
	out := &krm.RelatedAsset{}
	out.Asset = direct.LazyPtr(in.GetAsset())
	out.AssetType = direct.LazyPtr(in.GetAssetType())
	out.Ancestors = in.Ancestors
	return out
}
func RelatedAsset_ToProto(mapCtx *direct.MapContext, in *krm.RelatedAsset) *pb.RelatedAsset {
	if in == nil {
		return nil
	}
	out := &pb.RelatedAsset{}
	out.Asset = direct.ValueOf(in.Asset)
	out.AssetType = direct.ValueOf(in.AssetType)
	out.Ancestors = in.Ancestors
	return out
}
func RelatedAssets_FromProto(mapCtx *direct.MapContext, in *pb.RelatedAssets) *krm.RelatedAssets {
	if in == nil {
		return nil
	}
	out := &krm.RelatedAssets{}
	out.RelationshipAttributes = RelationshipAttributes_FromProto(mapCtx, in.GetRelationshipAttributes())
	out.Assets = direct.Slice_FromProto(mapCtx, in.Assets, RelatedAsset_FromProto)
	return out
}
func RelatedAssets_ToProto(mapCtx *direct.MapContext, in *krm.RelatedAssets) *pb.RelatedAssets {
	if in == nil {
		return nil
	}
	out := &pb.RelatedAssets{}
	out.RelationshipAttributes = RelationshipAttributes_ToProto(mapCtx, in.RelationshipAttributes)
	out.Assets = direct.Slice_ToProto(mapCtx, in.Assets, RelatedAsset_ToProto)
	return out
}
func RelationshipAttributes_FromProto(mapCtx *direct.MapContext, in *pb.RelationshipAttributes) *krm.RelationshipAttributes {
	if in == nil {
		return nil
	}
	out := &krm.RelationshipAttributes{}
	out.Type = direct.LazyPtr(in.GetType())
	out.SourceResourceType = direct.LazyPtr(in.GetSourceResourceType())
	out.TargetResourceType = direct.LazyPtr(in.GetTargetResourceType())
	out.Action = direct.LazyPtr(in.GetAction())
	return out
}
func RelationshipAttributes_ToProto(mapCtx *direct.MapContext, in *krm.RelationshipAttributes) *pb.RelationshipAttributes {
	if in == nil {
		return nil
	}
	out := &pb.RelationshipAttributes{}
	out.Type = direct.ValueOf(in.Type)
	out.SourceResourceType = direct.ValueOf(in.SourceResourceType)
	out.TargetResourceType = direct.ValueOf(in.TargetResourceType)
	out.Action = direct.ValueOf(in.Action)
	return out
}
func Resource_FromProto(mapCtx *direct.MapContext, in *pb.Resource) *krm.Resource {
	if in == nil {
		return nil
	}
	out := &krm.Resource{}
	out.Version = direct.LazyPtr(in.GetVersion())
	out.DiscoveryDocumentURI = direct.LazyPtr(in.GetDiscoveryDocumentUri())
	out.DiscoveryName = direct.LazyPtr(in.GetDiscoveryName())
	out.ResourceURL = direct.LazyPtr(in.GetResourceUrl())
	out.Parent = direct.LazyPtr(in.GetParent())
	out.Data = Data_FromProto(mapCtx, in.GetData())
	out.Location = direct.LazyPtr(in.GetLocation())
	return out
}
func Resource_ToProto(mapCtx *direct.MapContext, in *krm.Resource) *pb.Resource {
	if in == nil {
		return nil
	}
	out := &pb.Resource{}
	out.Version = direct.ValueOf(in.Version)
	out.DiscoveryDocumentUri = direct.ValueOf(in.DiscoveryDocumentURI)
	out.DiscoveryName = direct.ValueOf(in.DiscoveryName)
	out.ResourceUrl = direct.ValueOf(in.ResourceURL)
	out.Parent = direct.ValueOf(in.Parent)
	out.Data = Data_ToProto(mapCtx, in.Data)
	out.Location = direct.ValueOf(in.Location)
	return out
}
