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
	"sort"

	pb "cloud.google.com/go/dataplex/apiv1/dataplexpb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/dataplex/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AccessGroups_FromProto(mapCtx *direct.MapContext, in map[string]*pb.DataProduct_AccessGroup) []krm.DataProduct_AccessGroup {
	if len(in) == 0 {
		return nil
	}

	var out []krm.DataProduct_AccessGroup
	for k, v := range in {
		outItem := DataProduct_AccessGroup_FromProto(mapCtx, v)
		if outItem == nil {
			outItem = &krm.DataProduct_AccessGroup{}
		}
		// Capture the map key as the ID
		kCopy := k
		outItem.ID = &kCopy
		out = append(out, *outItem)
	}

	// Sort by ID to ensure deterministic ordering
	sort.Slice(out, func(i, j int) bool {
		idI := ""
		if out[i].ID != nil {
			idI = *out[i].ID
		}
		idJ := ""
		if out[j].ID != nil {
			idJ = *out[j].ID
		}
		return idI < idJ
	})

	return out
}

func AccessGroups_ToProto(mapCtx *direct.MapContext, in []krm.DataProduct_AccessGroup) map[string]*pb.DataProduct_AccessGroup {
	if len(in) == 0 {
		return nil
	}

	out := make(map[string]*pb.DataProduct_AccessGroup)
	for _, item := range in {
		v := DataProduct_AccessGroup_ToProto(mapCtx, &item)
		if v == nil {
			continue
		}
		key := ""
		if item.ID != nil {
			key = *item.ID
		}
		out[key] = v
	}

	return out
}
