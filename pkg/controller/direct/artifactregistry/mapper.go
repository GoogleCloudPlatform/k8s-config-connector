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

package artifactregistry

import (
	"sort"

	pb "cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/artifactregistry/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func CleanupPolicyCondition_TagState_ToProto(mapCtx *direct.MapContext, in *string) *pb.CleanupPolicyCondition_TagState {
	if in == nil {
		return nil
	}
	val := direct.Enum_ToProto[pb.CleanupPolicyCondition_TagState](mapCtx, in)
	return &val
}

func ArtifactRegistryRepositoryRef_FromProto(mapCtx *direct.MapContext, in *pb.Repository) *krm.ArtifactRegistryRepositoryRef {
	if in == nil {
		return nil
	}
	out := &krm.ArtifactRegistryRepositoryRef{}
	out.External = in.GetName()
	return out
}

func ArtifactRegistryRepositoryRef_ToProto(mapCtx *direct.MapContext, in *krm.ArtifactRegistryRepositoryRef) *pb.Repository {
	if in == nil {
		return nil
	}
	out := &pb.Repository{}
	out.Name = in.External
	return out
}

func CleanupPolicies_FromProto(mapCtx *direct.MapContext, in map[string]*pb.CleanupPolicy) []krm.CleanupPolicy {
	if in == nil {
		return nil
	}
	var out []krm.CleanupPolicy
	for id, policy := range in {
		p := CleanupPolicy_FromProto(mapCtx, policy)
		if p != nil {
			p.ID = direct.LazyPtr(id)
			out = append(out, *p)
		}
	}
	// Sort by ID to ensure deterministic order
	sort.Slice(out, func(i, j int) bool {
		return direct.ValueOf(out[i].ID) < direct.ValueOf(out[j].ID)
	})
	return out
}

func CleanupPolicies_ToProto(mapCtx *direct.MapContext, in []krm.CleanupPolicy) map[string]*pb.CleanupPolicy {
	if in == nil {
		return nil
	}
	out := make(map[string]*pb.CleanupPolicy)
	for _, policy := range in {
		id := direct.ValueOf(policy.ID)
		p := CleanupPolicy_ToProto(mapCtx, &policy)
		if p != nil {
			out[id] = p
		}
	}
	return out
}
