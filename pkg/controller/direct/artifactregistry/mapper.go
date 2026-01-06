// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package artifactregistry

import (
	"strings"

	pb "cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/artifactregistry/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func CleanupPolicy_FromProto(mapCtx *direct.MapContext, in *pb.CleanupPolicy) *krm.CleanupPolicy {
	if in == nil {
		return nil
	}
	out := &krm.CleanupPolicy{}
	out.ID = direct.LazyPtr(in.GetId())
	out.Action = direct.Enum_FromProto(mapCtx, in.GetAction())
	out.Condition = CleanupPolicyCondition_FromProto(mapCtx, in.GetCondition())
	out.MostRecentVersions = CleanupPolicyMostRecentVersions_FromProto(mapCtx, in.GetMostRecentVersions())
	return out
}

func CleanupPolicy_ToProto(mapCtx *direct.MapContext, in *krm.CleanupPolicy) *pb.CleanupPolicy {
	if in == nil {
		return nil
	}
	out := &pb.CleanupPolicy{}
	out.Id = direct.ValueOf(in.ID)
	if in.Action != nil {
		out.Action = pb.CleanupPolicy_Action(pb.CleanupPolicy_Action_value[*in.Action])
	}
	if in.Condition != nil {
		out.ConditionType = &pb.CleanupPolicy_Condition{
			Condition: CleanupPolicyCondition_ToProto(mapCtx, in.Condition),
		}
	}
	if in.MostRecentVersions != nil {
		out.ConditionType = &pb.CleanupPolicy_MostRecentVersions{
			MostRecentVersions: CleanupPolicyMostRecentVersions_ToProto(mapCtx, in.MostRecentVersions),
		}
	}
	return out
}

func UpstreamPolicy_FromProto(mapCtx *direct.MapContext, in *pb.UpstreamPolicy) *krm.UpstreamPolicy {
	if in == nil {
		return nil
	}
	out := &krm.UpstreamPolicy{}
	out.ID = direct.LazyPtr(in.GetId())
	if in.GetRepository() != "" {
		out.RepositoryRef = &krm.ArtifactRegistryRepositoryRef{
			External: in.GetRepository(),
		}
	}
	out.Priority = direct.LazyPtr(in.GetPriority())
	return out
}

func UpstreamPolicy_ToProto(mapCtx *direct.MapContext, in *krm.UpstreamPolicy) *pb.UpstreamPolicy {
	if in == nil {
		return nil
	}
	out := &pb.UpstreamPolicy{}
	out.Id = direct.ValueOf(in.ID)
	if in.RepositoryRef != nil {
		out.Repository = in.RepositoryRef.External
	}
	out.Priority = direct.ValueOf(in.Priority)
	return out
}

func KMSCryptoKeyRef_ToProto(mapCtx *direct.MapContext, in *krm.KMSCryptoKeyRef) string {
	if in == nil {
		return ""
	}
	return in.External
}

func KMSCryptoKeyRef_FromProto(mapCtx *direct.MapContext, in string) *krm.KMSCryptoKeyRef {
	if in == "" {
		return nil
	}
	return &krm.KMSCryptoKeyRef{
		External: in,
	}
}

func ArtifactRegistryRepositoryRef_ToProto(mapCtx *direct.MapContext, in *krm.ArtifactRegistryRepositoryRef) string {
	if in == nil {
		return ""
	}
	return in.External
}

func ArtifactRegistryRepositoryRef_FromProto(mapCtx *direct.MapContext, in string) *krm.ArtifactRegistryRepositoryRef {
	if in == "" {
		return nil
	}
	return &krm.ArtifactRegistryRepositoryRef{
		External: in,
	}
}

func CleanupPolicyCondition_TagState_ToProto(mapCtx *direct.MapContext, in *string) pb.CleanupPolicyCondition_TagState {
	if in == nil {
		return pb.CleanupPolicyCondition_TAG_STATE_UNSPECIFIED
	}
	return pb.CleanupPolicyCondition_TagState(pb.CleanupPolicyCondition_TagState_value[*in])
}

func CleanupPolicyCondition_TagState_FromProto(mapCtx *direct.MapContext, in pb.CleanupPolicyCondition_TagState) *string {
	if in == pb.CleanupPolicyCondition_TAG_STATE_UNSPECIFIED {
		return nil
	}
	s := in.String()
	return &s
}

func CleanupPolicyCondition_ToProto(mapCtx *direct.MapContext, in *krm.CleanupPolicyCondition) *pb.CleanupPolicyCondition {
	if in == nil {
		return nil
	}
	out := &pb.CleanupPolicyCondition{}
	if in.TagState != nil {
		val := CleanupPolicyCondition_TagState_ToProto(mapCtx, in.TagState)
		out.TagState = &val
	}
	out.TagPrefixes = in.TagPrefixes
	out.VersionNamePrefixes = in.VersionNamePrefixes
	out.PackageNamePrefixes = in.PackageNamePrefixes
	out.OlderThan = direct.Duration_ToProto(mapCtx, in.OlderThan)
	out.NewerThan = direct.Duration_ToProto(mapCtx, in.NewerThan)
	return out
}

func CleanupPolicyCondition_FromProto(mapCtx *direct.MapContext, in *pb.CleanupPolicyCondition) *krm.CleanupPolicyCondition {
	if in == nil {
		return nil
	}
	out := &krm.CleanupPolicyCondition{}
	out.TagState = CleanupPolicyCondition_TagState_FromProto(mapCtx, in.GetTagState())
	out.TagPrefixes = in.TagPrefixes
	out.VersionNamePrefixes = in.VersionNamePrefixes
	out.PackageNamePrefixes = in.PackageNamePrefixes
	out.OlderThan = direct.Duration_FromProto(mapCtx, in.GetOlderThan())
	out.NewerThan = direct.Duration_FromProto(mapCtx, in.GetNewerThan())
	return out
}

func ArtifactRegistryRepositorySpec_ToProto(mapCtx *direct.MapContext, in *krm.ArtifactRegistryRepositorySpec) *pb.Repository {
	if in == nil {
		return nil
	}
	out := &pb.Repository{}
	out.Format = pb.Repository_Format(pb.Repository_Format_value[in.Format])
	out.Description = direct.ValueOf(in.Description)
	if in.KmsKeyRef != nil {
		out.KmsKeyName = in.KmsKeyRef.External
	}
	if in.Mode != nil {
		out.Mode = pb.Repository_Mode(pb.Repository_Mode_value[*in.Mode])
	}
	if in.CleanupPolicies != nil {
		out.CleanupPolicies = make(map[string]*pb.CleanupPolicy)
		for _, cp := range in.CleanupPolicies {
			out.CleanupPolicies[direct.ValueOf(cp.ID)] = CleanupPolicy_ToProto(mapCtx, &cp)
		}
	}
	out.CleanupPolicyDryRun = direct.ValueOf(in.CleanupPolicyDryRun)
	if in.MavenConfig != nil {
		out.FormatConfig = &pb.Repository_MavenConfig{
			MavenConfig: Repository_MavenRepositoryConfig_ToProto(mapCtx, in.MavenConfig),
		}
	}
	if in.DockerConfig != nil {
		out.FormatConfig = &pb.Repository_DockerConfig{
			DockerConfig: Repository_DockerRepositoryConfig_ToProto(mapCtx, in.DockerConfig),
		}
	}
	if in.VirtualRepositoryConfig != nil {
		out.ModeConfig = &pb.Repository_VirtualRepositoryConfig{
			VirtualRepositoryConfig: VirtualRepositoryConfig_ToProto(mapCtx, in.VirtualRepositoryConfig),
		}
	}
	if in.RemoteRepositoryConfig != nil {
		out.ModeConfig = &pb.Repository_RemoteRepositoryConfig{
			RemoteRepositoryConfig: RemoteRepositoryConfig_ToProto(mapCtx, in.RemoteRepositoryConfig),
		}
	}
	out.VulnerabilityScanningConfig = Repository_VulnerabilityScanningConfig_ToProto(mapCtx, in.VulnerabilityScanningConfig)
	return out
}

func ArtifactRegistryRepositorySpec_FromProto(mapCtx *direct.MapContext, in *pb.Repository) *krm.ArtifactRegistryRepositorySpec {
	if in == nil {
		return nil
	}
	out := &krm.ArtifactRegistryRepositorySpec{}
	out.Format = in.GetFormat().String()
	out.Description = direct.LazyPtr(in.GetDescription())
	if in.GetKmsKeyName() != "" {
		out.KmsKeyRef = &krm.KMSCryptoKeyRef{
			External: in.GetKmsKeyName(),
		}
	}
	if in.GetMode() != pb.Repository_MODE_UNSPECIFIED {
		out.Mode = direct.LazyPtr(in.GetMode().String())
	}
	if in.GetCleanupPolicies() != nil {
		for _, cp := range in.GetCleanupPolicies() {
			out.CleanupPolicies = append(out.CleanupPolicies, *CleanupPolicy_FromProto(mapCtx, cp))
		}
	}
	out.CleanupPolicyDryRun = direct.LazyPtr(in.GetCleanupPolicyDryRun())
	out.MavenConfig = Repository_MavenRepositoryConfig_FromProto(mapCtx, in.GetMavenConfig())
	out.DockerConfig = Repository_DockerRepositoryConfig_FromProto(mapCtx, in.GetDockerConfig())
	out.VirtualRepositoryConfig = VirtualRepositoryConfig_FromProto(mapCtx, in.GetVirtualRepositoryConfig())
	out.RemoteRepositoryConfig = RemoteRepositoryConfig_FromProto(mapCtx, in.GetRemoteRepositoryConfig())
	out.VulnerabilityScanningConfig = Repository_VulnerabilityScanningConfig_FromProto(mapCtx, in.GetVulnerabilityScanningConfig())
	return out
}

func ArtifactRegistryRepositoryStatus_FromProto(mapCtx *direct.MapContext, in *pb.Repository) *krm.ArtifactRegistryRepositoryStatus {
	if in == nil {
		return nil
	}
	out := &krm.ArtifactRegistryRepositoryStatus{}
	out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetCreateTime())
	out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.GetUpdateTime())
	if in.GetName() != "" {
		parts := strings.Split(in.GetName(), "/")
		out.Name = direct.LazyPtr(parts[len(parts)-1])
	}
	out.ObservedState = ArtifactRegistryRepositoryObservedState_FromProto(mapCtx, in)
	return out
}

func ArtifactRegistryRepositoryStatus_ToProto(mapCtx *direct.MapContext, in *krm.ArtifactRegistryRepositoryStatus) *pb.Repository {
	if in == nil {
		return nil
	}
	out := &pb.Repository{}
	out.CreateTime = direct.StringTimestamp_ToProto(mapCtx, in.CreateTime)
	out.UpdateTime = direct.StringTimestamp_ToProto(mapCtx, in.UpdateTime)
	// Name mapping depends on parent, usually not mapped back to proto for status
	return out
}

func ArtifactRegistryRepositoryObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Repository) *krm.ArtifactRegistryRepositoryObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ArtifactRegistryRepositoryObservedState{}
	return out
}

func ArtifactRegistryRepositoryObservedState_ToProto(mapCtx *direct.MapContext, in *krm.ArtifactRegistryRepositoryObservedState) *pb.Repository {
	if in == nil {
		return nil
	}
	out := &pb.Repository{}
	return out
}
