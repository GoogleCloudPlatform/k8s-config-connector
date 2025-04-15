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
	pb "cloud.google.com/go/asset/apiv1/assetpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/asset/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// AssetSavedQueryStatus_FromProto converts the AssetSavedQueryStatus field from its Protobuf representation.
func AssetSavedQueryStatus_FromProto(mapCtx *direct.MapContext, in *pb.SavedQuery) *krm.AssetSavedQueryStatus {
	if in == nil {
		return nil
	}
	out := &krm.AssetSavedQueryStatus{}
	out.ObservedState = &krm.AssetSavedQueryObservedState{}
	out.ObservedState.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.ObservedState.Creator = direct.LazyPtr(in.GetCreator())
	out.ObservedState.LastUpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLastUpdateTime())
	out.ObservedState.LastUpdater = direct.LazyPtr(in.GetLastUpdater())
	out.ExternalRef = direct.LazyPtr(in.GetName()) // Set the external ref from the 'name' field
	return out
}
