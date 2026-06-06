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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
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

func ArtifactRegistryRepositorySpec_FromProto(mapCtx *direct.MapContext, in *pb.Repository) *krm.ArtifactRegistryRepositorySpec {
	if in == nil {
		return nil
	}
	out := &krm.ArtifactRegistryRepositorySpec{}
	out.MavenConfig = Repository_MavenRepositoryConfig_FromProto(mapCtx, in.GetMavenConfig())
	out.DockerConfig = Repository_DockerRepositoryConfig_FromProto(mapCtx, in.GetDockerConfig())
	out.VirtualRepositoryConfig = ArtifactRegistryRepositoryVirtualRepositoryConfig_FromProto(mapCtx, in.GetVirtualRepositoryConfig())
	out.RemoteRepositoryConfig = ArtifactRegistryRepositoryRemoteRepositoryConfig_FromProto(mapCtx, in.GetRemoteRepositoryConfig())
	out.Format = direct.Enum_FromProto(mapCtx, in.GetFormat())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.Mode = direct.Enum_FromProto(mapCtx, in.GetMode())
	out.CleanupPolicies = CleanupPolicies_FromProto(mapCtx, in.GetCleanupPolicies())
	out.CleanupPolicyDryRun = direct.LazyPtr(in.GetCleanupPolicyDryRun())
	if in.GetKmsKeyName() != "" {
		out.KmsKeyRef = &refs.KMSCryptoKeyRef{External: in.GetKmsKeyName()}
	}
	return out
}

func ArtifactRegistryRepositorySpec_ToProto(mapCtx *direct.MapContext, in *krm.ArtifactRegistryRepositorySpec) *pb.Repository {
	if in == nil {
		return nil
	}
	out := &pb.Repository{}
	if oneof := Repository_MavenRepositoryConfig_ToProto(mapCtx, in.MavenConfig); oneof != nil {
		out.FormatConfig = &pb.Repository_MavenConfig{MavenConfig: oneof}
	}
	if oneof := Repository_DockerRepositoryConfig_ToProto(mapCtx, in.DockerConfig); oneof != nil {
		out.FormatConfig = &pb.Repository_DockerConfig{DockerConfig: oneof}
	}
	if oneof := ArtifactRegistryRepositoryVirtualRepositoryConfig_ToProto(mapCtx, in.VirtualRepositoryConfig); oneof != nil {
		out.ModeConfig = &pb.Repository_VirtualRepositoryConfig{VirtualRepositoryConfig: oneof}
	}
	if oneof := ArtifactRegistryRepositoryRemoteRepositoryConfig_ToProto(mapCtx, in.RemoteRepositoryConfig); oneof != nil {
		out.ModeConfig = &pb.Repository_RemoteRepositoryConfig{RemoteRepositoryConfig: oneof}
	}
	out.Format = direct.Enum_ToProto[pb.Repository_Format](mapCtx, in.Format)
	out.Description = direct.ValueOf(in.Description)
	out.Mode = direct.Enum_ToProto[pb.Repository_Mode](mapCtx, in.Mode)
	out.CleanupPolicies = CleanupPolicies_ToProto(mapCtx, in.CleanupPolicies)
	out.CleanupPolicyDryRun = direct.ValueOf(in.CleanupPolicyDryRun)
	if in.KmsKeyRef != nil {
		out.KmsKeyName = in.KmsKeyRef.External
	}
	return out
}
