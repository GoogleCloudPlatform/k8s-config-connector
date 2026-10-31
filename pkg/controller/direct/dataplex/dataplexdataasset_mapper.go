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
	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataplex/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func DataplexDataAssetSpec_FromProto(mapCtx *direct.MapContext, in *pb.DataAsset) *krm.DataplexDataAssetSpec {
	if in == nil {
		return nil
	}
	out := &krm.DataplexDataAssetSpec{}
	// MISSING: Name
	// MISSING: Etag (placed in status.observedState)
	// MISSING: Labels (reconciled from metadata.labels)
	out.Resource = direct.LazyPtr(in.GetResource())
	if in.AccessGroupConfigs != nil {
		out.AccessGroupConfigs = make(map[string]krm.DataAsset_AccessGroupConfig, len(in.AccessGroupConfigs))
		for k, v := range in.AccessGroupConfigs {
			if val := DataAsset_AccessGroupConfig_FromProto(mapCtx, v); val != nil {
				out.AccessGroupConfigs[k] = *val
			}
		}
	}
	return out
}

func DataplexDataAssetSpec_ToProto(mapCtx *direct.MapContext, in *krm.DataplexDataAssetSpec) *pb.DataAsset {
	if in == nil {
		return nil
	}
	out := &pb.DataAsset{}
	// MISSING: Name
	// MISSING: Etag
	// MISSING: Labels
	out.Resource = direct.ValueOf(in.Resource)
	if in.AccessGroupConfigs != nil {
		out.AccessGroupConfigs = make(map[string]*pb.DataAsset_AccessGroupConfig, len(in.AccessGroupConfigs))
		for k, v := range in.AccessGroupConfigs {
			out.AccessGroupConfigs[k] = DataAsset_AccessGroupConfig_ToProto(mapCtx, &v)
		}
	}
	return out
}
