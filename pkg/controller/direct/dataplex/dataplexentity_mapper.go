// Copyright 2026 Google LLC
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
	"strings"

	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataplex/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DataplexEntitySpec_FromProto(mapCtx *direct.MapContext, in *pb.Entity) *krm.DataplexEntitySpec {
	if in == nil {
		return nil
	}
	out := &krm.DataplexEntitySpec{}
	// MISSING: Name
	out.DisplayName = direct.LazyPtr(in.GetDisplayName())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.ID = direct.LazyPtr(in.GetId())
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	if in.GetAsset() != "" {
		asset := in.GetAsset()
		if !strings.HasPrefix(asset, "projects/") && in.GetName() != "" {
			entityID := &krm.DataplexEntityIdentity{}
			if err := entityID.FromExternal(in.GetName()); err == nil {
				assetID := &krm.DataplexAssetIdentity{
					Project:  entityID.Project,
					Location: entityID.Location,
					Lake:     entityID.Lake,
					Zone:     entityID.Zone,
					Asset:    asset,
				}
				asset = assetID.String()
			}
		}
		out.AssetRef = &krm.DataplexAssetRef{External: asset}
	}
	out.DataPath = direct.LazyPtr(in.GetDataPath())
	out.DataPathPattern = direct.LazyPtr(in.GetDataPathPattern())
	out.System = direct.Enum_FromProto(mapCtx, in.GetSystem())
	out.Format = StorageFormat_FromProto(mapCtx, in.GetFormat())
	out.Schema = Schema_FromProto(mapCtx, in.GetSchema())
	return out
}

func DataplexEntitySpec_ToProto(mapCtx *direct.MapContext, in *krm.DataplexEntitySpec) *pb.Entity {
	if in == nil {
		return nil
	}
	out := &pb.Entity{}
	// MISSING: Name
	out.DisplayName = direct.ValueOf(in.DisplayName)
	out.Description = direct.ValueOf(in.Description)
	out.Id = direct.ValueOf(in.ID)
	out.Type = direct.Enum_ToProto[pb.Entity_Type](mapCtx, in.Type)
	if in.AssetRef != nil {
		asset := in.AssetRef.External
		id := &krm.DataplexAssetIdentity{}
		if err := id.FromExternal(asset); err == nil {
			out.Asset = id.Asset
		} else {
			out.Asset = asset
		}
	}
	out.DataPath = direct.ValueOf(in.DataPath)
	out.DataPathPattern = direct.ValueOf(in.DataPathPattern)
	out.System = direct.Enum_ToProto[pb.StorageSystem](mapCtx, in.System)
	out.Format = StorageFormat_ToProto(mapCtx, in.Format)
	out.Schema = Schema_ToProto(mapCtx, in.Schema)
	return out
}
