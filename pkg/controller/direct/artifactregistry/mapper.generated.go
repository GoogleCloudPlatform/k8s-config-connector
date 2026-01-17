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

// +generated:mapper
// krm.group: artifactregistry.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.devtools.artifactregistry.v1

package artifactregistry

import (
	pb "cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/artifactregistry/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func ArtifactRegistryRepositorySpec_FromProto(mapCtx *direct.MapContext, in *pb.Repository) *krm.ArtifactRegistryRepositorySpec {
	if in == nil {
		return nil
	}
	out := &krm.ArtifactRegistryRepositorySpec{}
	out.MavenConfig = Repository_MavenRepositoryConfig_FromProto(mapCtx, in.GetMavenConfig())
	out.DockerConfig = Repository_DockerRepositoryConfig_FromProto(mapCtx, in.GetDockerConfig())
	out.VirtualRepositoryConfig = ArtifactRegistryRepositoryVirtualRepositoryConfig_FromProto(mapCtx, in.GetVirtualRepositoryConfig())
	out.RemoteRepositoryConfig = RemoteRepositoryConfig_FromProto(mapCtx, in.GetRemoteRepositoryConfig())
	// MISSING: Name
	out.Format = direct.Enum_FromProto(mapCtx, in.GetFormat())
	out.Description = direct.LazyPtr(in.GetDescription())
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: KMSKeyName
	out.Mode = direct.Enum_FromProto(mapCtx, in.GetMode())
	// TODO: map type string message for field CleanupPolicies
	// MISSING: SizeBytes
	// MISSING: SatisfiesPzs
	out.CleanupPolicyDryRun = direct.LazyPtr(in.GetCleanupPolicyDryRun())
	// MISSING: VulnerabilityScanningConfig
	// MISSING: DisallowUnspecifiedMode
	// MISSING: SatisfiesPzi
	// MISSING: RegistryURI
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
	if oneof := RemoteRepositoryConfig_ToProto(mapCtx, in.RemoteRepositoryConfig); oneof != nil {
		out.ModeConfig = &pb.Repository_RemoteRepositoryConfig{RemoteRepositoryConfig: oneof}
	}
	// MISSING: Name
	out.Format = direct.Enum_ToProto[pb.Repository_Format](mapCtx, in.Format)
	out.Description = direct.ValueOf(in.Description)
	// MISSING: Labels
	// MISSING: CreateTime
	// MISSING: UpdateTime
	// MISSING: KMSKeyName
	out.Mode = direct.Enum_ToProto[pb.Repository_Mode](mapCtx, in.Mode)
	// TODO: map type string message for field CleanupPolicies
	// MISSING: SizeBytes
	// MISSING: SatisfiesPzs
	out.CleanupPolicyDryRun = direct.ValueOf(in.CleanupPolicyDryRun)
	// MISSING: VulnerabilityScanningConfig
	// MISSING: DisallowUnspecifiedMode
	// MISSING: SatisfiesPzi
	// MISSING: RegistryURI
	return out
}
func CleanupPolicy_FromProto(mapCtx *direct.MapContext, in *pb.CleanupPolicy) *krm.CleanupPolicy {
	if in == nil {
		return nil
	}
	out := &krm.CleanupPolicy{}
	out.Condition = CleanupPolicyCondition_FromProto(mapCtx, in.GetCondition())
	out.MostRecentVersions = CleanupPolicyMostRecentVersions_FromProto(mapCtx, in.GetMostRecentVersions())
	out.ID = direct.LazyPtr(in.GetId())
	out.Action = direct.Enum_FromProto(mapCtx, in.GetAction())
	return out
}
func CleanupPolicy_ToProto(mapCtx *direct.MapContext, in *krm.CleanupPolicy) *pb.CleanupPolicy {
	if in == nil {
		return nil
	}
	out := &pb.CleanupPolicy{}
	if oneof := CleanupPolicyCondition_ToProto(mapCtx, in.Condition); oneof != nil {
		out.ConditionType = &pb.CleanupPolicy_Condition{Condition: oneof}
	}
	if oneof := CleanupPolicyMostRecentVersions_ToProto(mapCtx, in.MostRecentVersions); oneof != nil {
		out.ConditionType = &pb.CleanupPolicy_MostRecentVersions{MostRecentVersions: oneof}
	}
	out.Id = direct.ValueOf(in.ID)
	out.Action = direct.Enum_ToProto[pb.CleanupPolicy_Action](mapCtx, in.Action)
	return out
}
func CleanupPolicyCondition_FromProto(mapCtx *direct.MapContext, in *pb.CleanupPolicyCondition) *krm.CleanupPolicyCondition {
	if in == nil {
		return nil
	}
	out := &krm.CleanupPolicyCondition{}
	out.TagState = direct.Enum_FromProto(mapCtx, in.GetTagState())
	out.TagPrefixes = in.TagPrefixes
	out.VersionNamePrefixes = in.VersionNamePrefixes
	out.PackageNamePrefixes = in.PackageNamePrefixes
	out.OlderThan = direct.StringDuration_FromProto(mapCtx, in.GetOlderThan())
	out.NewerThan = direct.StringDuration_FromProto(mapCtx, in.GetNewerThan())
	return out
}
func CleanupPolicyCondition_ToProto(mapCtx *direct.MapContext, in *krm.CleanupPolicyCondition) *pb.CleanupPolicyCondition {
	if in == nil {
		return nil
	}
	out := &pb.CleanupPolicyCondition{}
	if oneof := CleanupPolicyCondition_TagState_ToProto(mapCtx, in.TagState); oneof != nil {
		out.TagState = oneof
	}
	out.TagPrefixes = in.TagPrefixes
	out.VersionNamePrefixes = in.VersionNamePrefixes
	out.PackageNamePrefixes = in.PackageNamePrefixes
	out.OlderThan = direct.StringDuration_ToProto(mapCtx, in.OlderThan)
	out.NewerThan = direct.StringDuration_ToProto(mapCtx, in.NewerThan)
	return out
}
func CleanupPolicyMostRecentVersions_FromProto(mapCtx *direct.MapContext, in *pb.CleanupPolicyMostRecentVersions) *krm.CleanupPolicyMostRecentVersions {
	if in == nil {
		return nil
	}
	out := &krm.CleanupPolicyMostRecentVersions{}
	out.PackageNamePrefixes = in.PackageNamePrefixes
	out.KeepCount = in.KeepCount
	return out
}
func CleanupPolicyMostRecentVersions_ToProto(mapCtx *direct.MapContext, in *krm.CleanupPolicyMostRecentVersions) *pb.CleanupPolicyMostRecentVersions {
	if in == nil {
		return nil
	}
	out := &pb.CleanupPolicyMostRecentVersions{}
	out.PackageNamePrefixes = in.PackageNamePrefixes
	out.KeepCount = in.KeepCount
	return out
}
func RemoteRepositoryConfig_FromProto(mapCtx *direct.MapContext, in *pb.RemoteRepositoryConfig) *krm.RemoteRepositoryConfig {
	if in == nil {
		return nil
	}
	out := &krm.RemoteRepositoryConfig{}
	out.DockerRepository = RemoteRepositoryConfig_DockerRepository_FromProto(mapCtx, in.GetDockerRepository())
	out.MavenRepository = RemoteRepositoryConfig_MavenRepository_FromProto(mapCtx, in.GetMavenRepository())
	out.NpmRepository = RemoteRepositoryConfig_NpmRepository_FromProto(mapCtx, in.GetNpmRepository())
	out.PythonRepository = RemoteRepositoryConfig_PythonRepository_FromProto(mapCtx, in.GetPythonRepository())
	out.AptRepository = RemoteRepositoryConfig_AptRepository_FromProto(mapCtx, in.GetAptRepository())
	out.YumRepository = RemoteRepositoryConfig_YumRepository_FromProto(mapCtx, in.GetYumRepository())
	out.CommonRepository = RemoteRepositoryConfig_CommonRemoteRepository_FromProto(mapCtx, in.GetCommonRepository())
	out.Description = direct.LazyPtr(in.GetDescription())
	out.UpstreamCredentials = RemoteRepositoryConfig_UpstreamCredentials_FromProto(mapCtx, in.GetUpstreamCredentials())
	out.DisableUpstreamValidation = direct.LazyPtr(in.GetDisableUpstreamValidation())
	return out
}
func RemoteRepositoryConfig_ToProto(mapCtx *direct.MapContext, in *krm.RemoteRepositoryConfig) *pb.RemoteRepositoryConfig {
	if in == nil {
		return nil
	}
	out := &pb.RemoteRepositoryConfig{}
	if oneof := RemoteRepositoryConfig_DockerRepository_ToProto(mapCtx, in.DockerRepository); oneof != nil {
		out.RemoteSource = &pb.RemoteRepositoryConfig_DockerRepository_{DockerRepository: oneof}
	}
	if oneof := RemoteRepositoryConfig_MavenRepository_ToProto(mapCtx, in.MavenRepository); oneof != nil {
		out.RemoteSource = &pb.RemoteRepositoryConfig_MavenRepository_{MavenRepository: oneof}
	}
	if oneof := RemoteRepositoryConfig_NpmRepository_ToProto(mapCtx, in.NpmRepository); oneof != nil {
		out.RemoteSource = &pb.RemoteRepositoryConfig_NpmRepository_{NpmRepository: oneof}
	}
	if oneof := RemoteRepositoryConfig_PythonRepository_ToProto(mapCtx, in.PythonRepository); oneof != nil {
		out.RemoteSource = &pb.RemoteRepositoryConfig_PythonRepository_{PythonRepository: oneof}
	}
	if oneof := RemoteRepositoryConfig_AptRepository_ToProto(mapCtx, in.AptRepository); oneof != nil {
		out.RemoteSource = &pb.RemoteRepositoryConfig_AptRepository_{AptRepository: oneof}
	}
	if oneof := RemoteRepositoryConfig_YumRepository_ToProto(mapCtx, in.YumRepository); oneof != nil {
		out.RemoteSource = &pb.RemoteRepositoryConfig_YumRepository_{YumRepository: oneof}
	}
	if oneof := RemoteRepositoryConfig_CommonRemoteRepository_ToProto(mapCtx, in.CommonRepository); oneof != nil {
		out.RemoteSource = &pb.RemoteRepositoryConfig_CommonRepository{CommonRepository: oneof}
	}
	out.Description = direct.ValueOf(in.Description)
	out.UpstreamCredentials = RemoteRepositoryConfig_UpstreamCredentials_ToProto(mapCtx, in.UpstreamCredentials)
	out.DisableUpstreamValidation = direct.ValueOf(in.DisableUpstreamValidation)
	return out
}
func RemoteRepositoryConfig_AptRepository_FromProto(mapCtx *direct.MapContext, in *pb.RemoteRepositoryConfig_AptRepository) *krm.RemoteRepositoryConfig_AptRepository {
	if in == nil {
		return nil
	}
	out := &krm.RemoteRepositoryConfig_AptRepository{}
	out.PublicRepository = RemoteRepositoryConfig_AptRepository_PublicRepository_FromProto(mapCtx, in.GetPublicRepository())
	out.CustomRepository = RemoteRepositoryConfig_AptRepository_CustomRepository_FromProto(mapCtx, in.GetCustomRepository())
	return out
}
func RemoteRepositoryConfig_AptRepository_ToProto(mapCtx *direct.MapContext, in *krm.RemoteRepositoryConfig_AptRepository) *pb.RemoteRepositoryConfig_AptRepository {
	if in == nil {
		return nil
	}
	out := &pb.RemoteRepositoryConfig_AptRepository{}
	if oneof := RemoteRepositoryConfig_AptRepository_PublicRepository_ToProto(mapCtx, in.PublicRepository); oneof != nil {
		out.Upstream = &pb.RemoteRepositoryConfig_AptRepository_PublicRepository_{PublicRepository: oneof}
	}
	if oneof := RemoteRepositoryConfig_AptRepository_CustomRepository_ToProto(mapCtx, in.CustomRepository); oneof != nil {
		out.Upstream = &pb.RemoteRepositoryConfig_AptRepository_CustomRepository_{CustomRepository: oneof}
	}
	return out
}
func RemoteRepositoryConfig_AptRepository_CustomRepository_FromProto(mapCtx *direct.MapContext, in *pb.RemoteRepositoryConfig_AptRepository_CustomRepository) *krm.RemoteRepositoryConfig_AptRepository_CustomRepository {
	if in == nil {
		return nil
	}
	out := &krm.RemoteRepositoryConfig_AptRepository_CustomRepository{}
	out.URI = direct.LazyPtr(in.GetUri())
	return out
}
func RemoteRepositoryConfig_AptRepository_CustomRepository_ToProto(mapCtx *direct.MapContext, in *krm.RemoteRepositoryConfig_AptRepository_CustomRepository) *pb.RemoteRepositoryConfig_AptRepository_CustomRepository {
	if in == nil {
		return nil
	}
	out := &pb.RemoteRepositoryConfig_AptRepository_CustomRepository{}
	out.Uri = direct.ValueOf(in.URI)
	return out
}
func RemoteRepositoryConfig_AptRepository_PublicRepository_FromProto(mapCtx *direct.MapContext, in *pb.RemoteRepositoryConfig_AptRepository_PublicRepository) *krm.RemoteRepositoryConfig_AptRepository_PublicRepository {
	if in == nil {
		return nil
	}
	out := &krm.RemoteRepositoryConfig_AptRepository_PublicRepository{}
	out.RepositoryBase = direct.Enum_FromProto(mapCtx, in.GetRepositoryBase())
	out.RepositoryPath = direct.LazyPtr(in.GetRepositoryPath())
	return out
}
func RemoteRepositoryConfig_AptRepository_PublicRepository_ToProto(mapCtx *direct.MapContext, in *krm.RemoteRepositoryConfig_AptRepository_PublicRepository) *pb.RemoteRepositoryConfig_AptRepository_PublicRepository {
	if in == nil {
		return nil
	}
	out := &pb.RemoteRepositoryConfig_AptRepository_PublicRepository{}
	out.RepositoryBase = direct.Enum_ToProto[pb.RemoteRepositoryConfig_AptRepository_PublicRepository_RepositoryBase](mapCtx, in.RepositoryBase)
	out.RepositoryPath = direct.ValueOf(in.RepositoryPath)
	return out
}
func RemoteRepositoryConfig_CommonRemoteRepository_FromProto(mapCtx *direct.MapContext, in *pb.RemoteRepositoryConfig_CommonRemoteRepository) *krm.RemoteRepositoryConfig_CommonRemoteRepository {
	if in == nil {
		return nil
	}
	out := &krm.RemoteRepositoryConfig_CommonRemoteRepository{}
	out.URI = direct.LazyPtr(in.GetUri())
	return out
}
func RemoteRepositoryConfig_CommonRemoteRepository_ToProto(mapCtx *direct.MapContext, in *krm.RemoteRepositoryConfig_CommonRemoteRepository) *pb.RemoteRepositoryConfig_CommonRemoteRepository {
	if in == nil {
		return nil
	}
	out := &pb.RemoteRepositoryConfig_CommonRemoteRepository{}
	out.Uri = direct.ValueOf(in.URI)
	return out
}
func RemoteRepositoryConfig_DockerRepository_FromProto(mapCtx *direct.MapContext, in *pb.RemoteRepositoryConfig_DockerRepository) *krm.RemoteRepositoryConfig_DockerRepository {
	if in == nil {
		return nil
	}
	out := &krm.RemoteRepositoryConfig_DockerRepository{}
	out.PublicRepository = direct.Enum_FromProto(mapCtx, in.GetPublicRepository())
	out.CustomRepository = RemoteRepositoryConfig_DockerRepository_CustomRepository_FromProto(mapCtx, in.GetCustomRepository())
	return out
}
func RemoteRepositoryConfig_DockerRepository_ToProto(mapCtx *direct.MapContext, in *krm.RemoteRepositoryConfig_DockerRepository) *pb.RemoteRepositoryConfig_DockerRepository {
	if in == nil {
		return nil
	}
	out := &pb.RemoteRepositoryConfig_DockerRepository{}
	if oneof := RemoteRepositoryConfig_DockerRepository_PublicRepository_ToProto(mapCtx, in.PublicRepository); oneof != nil {
		out.Upstream = oneof
	}
	if oneof := RemoteRepositoryConfig_DockerRepository_CustomRepository_ToProto(mapCtx, in.CustomRepository); oneof != nil {
		out.Upstream = &pb.RemoteRepositoryConfig_DockerRepository_CustomRepository_{CustomRepository: oneof}
	}
	return out
}
func RemoteRepositoryConfig_DockerRepository_PublicRepository_ToProto(mapCtx *direct.MapContext, in *string) *pb.RemoteRepositoryConfig_DockerRepository_PublicRepository_ {
	if in == nil {
		return nil
	}
	return &pb.RemoteRepositoryConfig_DockerRepository_PublicRepository_{PublicRepository: direct.Enum_ToProto[pb.RemoteRepositoryConfig_DockerRepository_PublicRepository](mapCtx, in)}
}
func RemoteRepositoryConfig_DockerRepository_CustomRepository_FromProto(mapCtx *direct.MapContext, in *pb.RemoteRepositoryConfig_DockerRepository_CustomRepository) *krm.RemoteRepositoryConfig_DockerRepository_CustomRepository {
	if in == nil {
		return nil
	}
	out := &krm.RemoteRepositoryConfig_DockerRepository_CustomRepository{}
	out.URI = direct.LazyPtr(in.GetUri())
	return out
}
func RemoteRepositoryConfig_DockerRepository_CustomRepository_ToProto(mapCtx *direct.MapContext, in *krm.RemoteRepositoryConfig_DockerRepository_CustomRepository) *pb.RemoteRepositoryConfig_DockerRepository_CustomRepository {
	if in == nil {
		return nil
	}
	out := &pb.RemoteRepositoryConfig_DockerRepository_CustomRepository{}
	out.Uri = direct.ValueOf(in.URI)
	return out
}
func RemoteRepositoryConfig_MavenRepository_FromProto(mapCtx *direct.MapContext, in *pb.RemoteRepositoryConfig_MavenRepository) *krm.RemoteRepositoryConfig_MavenRepository {
	if in == nil {
		return nil
	}
	out := &krm.RemoteRepositoryConfig_MavenRepository{}
	out.PublicRepository = direct.Enum_FromProto(mapCtx, in.GetPublicRepository())
	out.CustomRepository = RemoteRepositoryConfig_MavenRepository_CustomRepository_FromProto(mapCtx, in.GetCustomRepository())
	return out
}
func RemoteRepositoryConfig_MavenRepository_ToProto(mapCtx *direct.MapContext, in *krm.RemoteRepositoryConfig_MavenRepository) *pb.RemoteRepositoryConfig_MavenRepository {
	if in == nil {
		return nil
	}
	out := &pb.RemoteRepositoryConfig_MavenRepository{}
	if oneof := RemoteRepositoryConfig_MavenRepository_PublicRepository_ToProto(mapCtx, in.PublicRepository); oneof != nil {
		out.Upstream = oneof
	}
	if oneof := RemoteRepositoryConfig_MavenRepository_CustomRepository_ToProto(mapCtx, in.CustomRepository); oneof != nil {
		out.Upstream = &pb.RemoteRepositoryConfig_MavenRepository_CustomRepository_{CustomRepository: oneof}
	}
	return out
}
func RemoteRepositoryConfig_MavenRepository_PublicRepository_ToProto(mapCtx *direct.MapContext, in *string) *pb.RemoteRepositoryConfig_MavenRepository_PublicRepository_ {
	if in == nil {
		return nil
	}
	return &pb.RemoteRepositoryConfig_MavenRepository_PublicRepository_{PublicRepository: direct.Enum_ToProto[pb.RemoteRepositoryConfig_MavenRepository_PublicRepository](mapCtx, in)}
}
func RemoteRepositoryConfig_MavenRepository_CustomRepository_FromProto(mapCtx *direct.MapContext, in *pb.RemoteRepositoryConfig_MavenRepository_CustomRepository) *krm.RemoteRepositoryConfig_MavenRepository_CustomRepository {
	if in == nil {
		return nil
	}
	out := &krm.RemoteRepositoryConfig_MavenRepository_CustomRepository{}
	out.URI = direct.LazyPtr(in.GetUri())
	return out
}
func RemoteRepositoryConfig_MavenRepository_CustomRepository_ToProto(mapCtx *direct.MapContext, in *krm.RemoteRepositoryConfig_MavenRepository_CustomRepository) *pb.RemoteRepositoryConfig_MavenRepository_CustomRepository {
	if in == nil {
		return nil
	}
	out := &pb.RemoteRepositoryConfig_MavenRepository_CustomRepository{}
	out.Uri = direct.ValueOf(in.URI)
	return out
}
func RemoteRepositoryConfig_NpmRepository_FromProto(mapCtx *direct.MapContext, in *pb.RemoteRepositoryConfig_NpmRepository) *krm.RemoteRepositoryConfig_NpmRepository {
	if in == nil {
		return nil
	}
	out := &krm.RemoteRepositoryConfig_NpmRepository{}
	out.PublicRepository = direct.Enum_FromProto(mapCtx, in.GetPublicRepository())
	out.CustomRepository = RemoteRepositoryConfig_NpmRepository_CustomRepository_FromProto(mapCtx, in.GetCustomRepository())
	return out
}
func RemoteRepositoryConfig_NpmRepository_ToProto(mapCtx *direct.MapContext, in *krm.RemoteRepositoryConfig_NpmRepository) *pb.RemoteRepositoryConfig_NpmRepository {
	if in == nil {
		return nil
	}
	out := &pb.RemoteRepositoryConfig_NpmRepository{}
	if oneof := RemoteRepositoryConfig_NpmRepository_PublicRepository_ToProto(mapCtx, in.PublicRepository); oneof != nil {
		out.Upstream = oneof
	}
	if oneof := RemoteRepositoryConfig_NpmRepository_CustomRepository_ToProto(mapCtx, in.CustomRepository); oneof != nil {
		out.Upstream = &pb.RemoteRepositoryConfig_NpmRepository_CustomRepository_{CustomRepository: oneof}
	}
	return out
}
func RemoteRepositoryConfig_NpmRepository_PublicRepository_ToProto(mapCtx *direct.MapContext, in *string) *pb.RemoteRepositoryConfig_NpmRepository_PublicRepository_ {
	if in == nil {
		return nil
	}
	return &pb.RemoteRepositoryConfig_NpmRepository_PublicRepository_{PublicRepository: direct.Enum_ToProto[pb.RemoteRepositoryConfig_NpmRepository_PublicRepository](mapCtx, in)}
}
func RemoteRepositoryConfig_NpmRepository_CustomRepository_FromProto(mapCtx *direct.MapContext, in *pb.RemoteRepositoryConfig_NpmRepository_CustomRepository) *krm.RemoteRepositoryConfig_NpmRepository_CustomRepository {
	if in == nil {
		return nil
	}
	out := &krm.RemoteRepositoryConfig_NpmRepository_CustomRepository{}
	out.URI = direct.LazyPtr(in.GetUri())
	return out
}
func RemoteRepositoryConfig_NpmRepository_CustomRepository_ToProto(mapCtx *direct.MapContext, in *krm.RemoteRepositoryConfig_NpmRepository_CustomRepository) *pb.RemoteRepositoryConfig_NpmRepository_CustomRepository {
	if in == nil {
		return nil
	}
	out := &pb.RemoteRepositoryConfig_NpmRepository_CustomRepository{}
	out.Uri = direct.ValueOf(in.URI)
	return out
}
func RemoteRepositoryConfig_PythonRepository_FromProto(mapCtx *direct.MapContext, in *pb.RemoteRepositoryConfig_PythonRepository) *krm.RemoteRepositoryConfig_PythonRepository {
	if in == nil {
		return nil
	}
	out := &krm.RemoteRepositoryConfig_PythonRepository{}
	out.PublicRepository = direct.Enum_FromProto(mapCtx, in.GetPublicRepository())
	out.CustomRepository = RemoteRepositoryConfig_PythonRepository_CustomRepository_FromProto(mapCtx, in.GetCustomRepository())
	return out
}
func RemoteRepositoryConfig_PythonRepository_ToProto(mapCtx *direct.MapContext, in *krm.RemoteRepositoryConfig_PythonRepository) *pb.RemoteRepositoryConfig_PythonRepository {
	if in == nil {
		return nil
	}
	out := &pb.RemoteRepositoryConfig_PythonRepository{}
	if oneof := RemoteRepositoryConfig_PythonRepository_PublicRepository_ToProto(mapCtx, in.PublicRepository); oneof != nil {
		out.Upstream = oneof
	}
	if oneof := RemoteRepositoryConfig_PythonRepository_CustomRepository_ToProto(mapCtx, in.CustomRepository); oneof != nil {
		out.Upstream = &pb.RemoteRepositoryConfig_PythonRepository_CustomRepository_{CustomRepository: oneof}
	}
	return out
}
func RemoteRepositoryConfig_PythonRepository_PublicRepository_ToProto(mapCtx *direct.MapContext, in *string) *pb.RemoteRepositoryConfig_PythonRepository_PublicRepository_ {
	if in == nil {
		return nil
	}
	return &pb.RemoteRepositoryConfig_PythonRepository_PublicRepository_{PublicRepository: direct.Enum_ToProto[pb.RemoteRepositoryConfig_PythonRepository_PublicRepository](mapCtx, in)}
}
func RemoteRepositoryConfig_PythonRepository_CustomRepository_FromProto(mapCtx *direct.MapContext, in *pb.RemoteRepositoryConfig_PythonRepository_CustomRepository) *krm.RemoteRepositoryConfig_PythonRepository_CustomRepository {
	if in == nil {
		return nil
	}
	out := &krm.RemoteRepositoryConfig_PythonRepository_CustomRepository{}
	out.URI = direct.LazyPtr(in.GetUri())
	return out
}
func RemoteRepositoryConfig_PythonRepository_CustomRepository_ToProto(mapCtx *direct.MapContext, in *krm.RemoteRepositoryConfig_PythonRepository_CustomRepository) *pb.RemoteRepositoryConfig_PythonRepository_CustomRepository {
	if in == nil {
		return nil
	}
	out := &pb.RemoteRepositoryConfig_PythonRepository_CustomRepository{}
	out.Uri = direct.ValueOf(in.URI)
	return out
}
func RemoteRepositoryConfig_UpstreamCredentials_FromProto(mapCtx *direct.MapContext, in *pb.RemoteRepositoryConfig_UpstreamCredentials) *krm.RemoteRepositoryConfig_UpstreamCredentials {
	if in == nil {
		return nil
	}
	out := &krm.RemoteRepositoryConfig_UpstreamCredentials{}
	out.UsernamePasswordCredentials = RemoteRepositoryConfig_UpstreamCredentials_UsernamePasswordCredentials_FromProto(mapCtx, in.GetUsernamePasswordCredentials())
	return out
}
func RemoteRepositoryConfig_UpstreamCredentials_ToProto(mapCtx *direct.MapContext, in *krm.RemoteRepositoryConfig_UpstreamCredentials) *pb.RemoteRepositoryConfig_UpstreamCredentials {
	if in == nil {
		return nil
	}
	out := &pb.RemoteRepositoryConfig_UpstreamCredentials{}
	if oneof := RemoteRepositoryConfig_UpstreamCredentials_UsernamePasswordCredentials_ToProto(mapCtx, in.UsernamePasswordCredentials); oneof != nil {
		out.Credentials = &pb.RemoteRepositoryConfig_UpstreamCredentials_UsernamePasswordCredentials_{UsernamePasswordCredentials: oneof}
	}
	return out
}
func RemoteRepositoryConfig_UpstreamCredentials_UsernamePasswordCredentials_FromProto(mapCtx *direct.MapContext, in *pb.RemoteRepositoryConfig_UpstreamCredentials_UsernamePasswordCredentials) *krm.RemoteRepositoryConfig_UpstreamCredentials_UsernamePasswordCredentials {
	if in == nil {
		return nil
	}
	out := &krm.RemoteRepositoryConfig_UpstreamCredentials_UsernamePasswordCredentials{}
	out.Username = direct.LazyPtr(in.GetUsername())
	out.PasswordSecretVersion = direct.LazyPtr(in.GetPasswordSecretVersion())
	return out
}
func RemoteRepositoryConfig_UpstreamCredentials_UsernamePasswordCredentials_ToProto(mapCtx *direct.MapContext, in *krm.RemoteRepositoryConfig_UpstreamCredentials_UsernamePasswordCredentials) *pb.RemoteRepositoryConfig_UpstreamCredentials_UsernamePasswordCredentials {
	if in == nil {
		return nil
	}
	out := &pb.RemoteRepositoryConfig_UpstreamCredentials_UsernamePasswordCredentials{}
	out.Username = direct.ValueOf(in.Username)
	out.PasswordSecretVersion = direct.ValueOf(in.PasswordSecretVersion)
	return out
}
func RemoteRepositoryConfig_YumRepository_FromProto(mapCtx *direct.MapContext, in *pb.RemoteRepositoryConfig_YumRepository) *krm.RemoteRepositoryConfig_YumRepository {
	if in == nil {
		return nil
	}
	out := &krm.RemoteRepositoryConfig_YumRepository{}
	out.PublicRepository = RemoteRepositoryConfig_YumRepository_PublicRepository_FromProto(mapCtx, in.GetPublicRepository())
	out.CustomRepository = RemoteRepositoryConfig_YumRepository_CustomRepository_FromProto(mapCtx, in.GetCustomRepository())
	return out
}
func RemoteRepositoryConfig_YumRepository_ToProto(mapCtx *direct.MapContext, in *krm.RemoteRepositoryConfig_YumRepository) *pb.RemoteRepositoryConfig_YumRepository {
	if in == nil {
		return nil
	}
	out := &pb.RemoteRepositoryConfig_YumRepository{}
	if oneof := RemoteRepositoryConfig_YumRepository_PublicRepository_ToProto(mapCtx, in.PublicRepository); oneof != nil {
		out.Upstream = &pb.RemoteRepositoryConfig_YumRepository_PublicRepository_{PublicRepository: oneof}
	}
	if oneof := RemoteRepositoryConfig_YumRepository_CustomRepository_ToProto(mapCtx, in.CustomRepository); oneof != nil {
		out.Upstream = &pb.RemoteRepositoryConfig_YumRepository_CustomRepository_{CustomRepository: oneof}
	}
	return out
}
func RemoteRepositoryConfig_YumRepository_CustomRepository_FromProto(mapCtx *direct.MapContext, in *pb.RemoteRepositoryConfig_YumRepository_CustomRepository) *krm.RemoteRepositoryConfig_YumRepository_CustomRepository {
	if in == nil {
		return nil
	}
	out := &krm.RemoteRepositoryConfig_YumRepository_CustomRepository{}
	out.URI = direct.LazyPtr(in.GetUri())
	return out
}
func RemoteRepositoryConfig_YumRepository_CustomRepository_ToProto(mapCtx *direct.MapContext, in *krm.RemoteRepositoryConfig_YumRepository_CustomRepository) *pb.RemoteRepositoryConfig_YumRepository_CustomRepository {
	if in == nil {
		return nil
	}
	out := &pb.RemoteRepositoryConfig_YumRepository_CustomRepository{}
	out.Uri = direct.ValueOf(in.URI)
	return out
}
func RemoteRepositoryConfig_YumRepository_PublicRepository_FromProto(mapCtx *direct.MapContext, in *pb.RemoteRepositoryConfig_YumRepository_PublicRepository) *krm.RemoteRepositoryConfig_YumRepository_PublicRepository {
	if in == nil {
		return nil
	}
	out := &krm.RemoteRepositoryConfig_YumRepository_PublicRepository{}
	out.RepositoryBase = direct.Enum_FromProto(mapCtx, in.GetRepositoryBase())
	out.RepositoryPath = direct.LazyPtr(in.GetRepositoryPath())
	return out
}
func RemoteRepositoryConfig_YumRepository_PublicRepository_ToProto(mapCtx *direct.MapContext, in *krm.RemoteRepositoryConfig_YumRepository_PublicRepository) *pb.RemoteRepositoryConfig_YumRepository_PublicRepository {
	if in == nil {
		return nil
	}
	out := &pb.RemoteRepositoryConfig_YumRepository_PublicRepository{}
	out.RepositoryBase = direct.Enum_ToProto[pb.RemoteRepositoryConfig_YumRepository_PublicRepository_RepositoryBase](mapCtx, in.RepositoryBase)
	out.RepositoryPath = direct.ValueOf(in.RepositoryPath)
	return out
}
func Repository_DockerRepositoryConfig_FromProto(mapCtx *direct.MapContext, in *pb.Repository_DockerRepositoryConfig) *krm.Repository_DockerRepositoryConfig {
	if in == nil {
		return nil
	}
	out := &krm.Repository_DockerRepositoryConfig{}
	out.ImmutableTags = direct.LazyPtr(in.GetImmutableTags())
	return out
}
func Repository_DockerRepositoryConfig_ToProto(mapCtx *direct.MapContext, in *krm.Repository_DockerRepositoryConfig) *pb.Repository_DockerRepositoryConfig {
	if in == nil {
		return nil
	}
	out := &pb.Repository_DockerRepositoryConfig{}
	out.ImmutableTags = direct.ValueOf(in.ImmutableTags)
	return out
}
func Repository_MavenRepositoryConfig_FromProto(mapCtx *direct.MapContext, in *pb.Repository_MavenRepositoryConfig) *krm.Repository_MavenRepositoryConfig {
	if in == nil {
		return nil
	}
	out := &krm.Repository_MavenRepositoryConfig{}
	out.AllowSnapshotOverwrites = direct.LazyPtr(in.GetAllowSnapshotOverwrites())
	out.VersionPolicy = direct.Enum_FromProto(mapCtx, in.GetVersionPolicy())
	return out
}
func Repository_MavenRepositoryConfig_ToProto(mapCtx *direct.MapContext, in *krm.Repository_MavenRepositoryConfig) *pb.Repository_MavenRepositoryConfig {
	if in == nil {
		return nil
	}
	out := &pb.Repository_MavenRepositoryConfig{}
	out.AllowSnapshotOverwrites = direct.ValueOf(in.AllowSnapshotOverwrites)
	out.VersionPolicy = direct.Enum_ToProto[pb.Repository_MavenRepositoryConfig_VersionPolicy](mapCtx, in.VersionPolicy)
	return out
}
func Repository_VulnerabilityScanningConfig_FromProto(mapCtx *direct.MapContext, in *pb.Repository_VulnerabilityScanningConfig) *krm.Repository_VulnerabilityScanningConfig {
	if in == nil {
		return nil
	}
	out := &krm.Repository_VulnerabilityScanningConfig{}
	out.EnablementConfig = direct.Enum_FromProto(mapCtx, in.GetEnablementConfig())
	// MISSING: LastEnableTime
	// MISSING: EnablementState
	// MISSING: EnablementStateReason
	return out
}
func Repository_VulnerabilityScanningConfig_ToProto(mapCtx *direct.MapContext, in *krm.Repository_VulnerabilityScanningConfig) *pb.Repository_VulnerabilityScanningConfig {
	if in == nil {
		return nil
	}
	out := &pb.Repository_VulnerabilityScanningConfig{}
	out.EnablementConfig = direct.Enum_ToProto[pb.Repository_VulnerabilityScanningConfig_EnablementConfig](mapCtx, in.EnablementConfig)
	// MISSING: LastEnableTime
	// MISSING: EnablementState
	// MISSING: EnablementStateReason
	return out
}
func Repository_VulnerabilityScanningConfigObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Repository_VulnerabilityScanningConfig) *krm.Repository_VulnerabilityScanningConfigObservedState {
	if in == nil {
		return nil
	}
	out := &krm.Repository_VulnerabilityScanningConfigObservedState{}
	// MISSING: EnablementConfig
	out.LastEnableTime = direct.StringTimestamp_FromProto(mapCtx, in.GetLastEnableTime())
	out.EnablementState = direct.Enum_FromProto(mapCtx, in.GetEnablementState())
	out.EnablementStateReason = direct.LazyPtr(in.GetEnablementStateReason())
	return out
}
func Repository_VulnerabilityScanningConfigObservedState_ToProto(mapCtx *direct.MapContext, in *krm.Repository_VulnerabilityScanningConfigObservedState) *pb.Repository_VulnerabilityScanningConfig {
	if in == nil {
		return nil
	}
	out := &pb.Repository_VulnerabilityScanningConfig{}
	// MISSING: EnablementConfig
	out.LastEnableTime = direct.StringTimestamp_ToProto(mapCtx, in.LastEnableTime)
	out.EnablementState = direct.Enum_ToProto[pb.Repository_VulnerabilityScanningConfig_EnablementState](mapCtx, in.EnablementState)
	out.EnablementStateReason = direct.ValueOf(in.EnablementStateReason)
	return out
}
func UpstreamPolicy_FromProto(mapCtx *direct.MapContext, in *pb.UpstreamPolicy) *krm.UpstreamPolicy {
	if in == nil {
		return nil
	}
	out := &krm.UpstreamPolicy{}
	out.ID = direct.LazyPtr(in.GetId())
	out.Repository = direct.LazyPtr(in.GetRepository())
	out.Priority = direct.LazyPtr(in.GetPriority())
	return out
}
func UpstreamPolicy_ToProto(mapCtx *direct.MapContext, in *krm.UpstreamPolicy) *pb.UpstreamPolicy {
	if in == nil {
		return nil
	}
	out := &pb.UpstreamPolicy{}
	out.Id = direct.ValueOf(in.ID)
	out.Repository = direct.ValueOf(in.Repository)
	out.Priority = direct.ValueOf(in.Priority)
	return out
}
func VirtualRepositoryConfig_FromProto(mapCtx *direct.MapContext, in *pb.VirtualRepositoryConfig) *krm.VirtualRepositoryConfig {
	if in == nil {
		return nil
	}
	out := &krm.VirtualRepositoryConfig{}
	out.UpstreamPolicies = direct.Slice_FromProto(mapCtx, in.UpstreamPolicies, UpstreamPolicy_FromProto)
	return out
}
func VirtualRepositoryConfig_ToProto(mapCtx *direct.MapContext, in *krm.VirtualRepositoryConfig) *pb.VirtualRepositoryConfig {
	if in == nil {
		return nil
	}
	out := &pb.VirtualRepositoryConfig{}
	out.UpstreamPolicies = direct.Slice_ToProto(mapCtx, in.UpstreamPolicies, UpstreamPolicy_ToProto)
	return out
}
