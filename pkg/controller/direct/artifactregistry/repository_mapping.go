// Copyright 2024 Google LLC
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
	"strings"
	
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/devtools/artifactregistry/v1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/artifactregistry/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// FromProto functions for reading from GCP

func ArtifactRegistryRepositorySpec_FromProto(mapCtx *direct.MapContext, in *pb.Repository) *krm.ArtifactRegistryRepositorySpec {
	if in == nil {
		return nil
	}
	out := &krm.ArtifactRegistryRepositorySpec{}
	
	// Convert format
	out.Format = Repository_Format_FromProto(mapCtx, in.Format)
	
	// Convert description
	if in.Description != "" {
		out.Description = &in.Description
	}
	
	// Convert mode
	out.Mode = Repository_Mode_FromProto(mapCtx, in.Mode)
	
	// Convert cleanup policy dry run
	if in.CleanupPolicyDryRun {
		out.CleanupPolicyDryRun = &in.CleanupPolicyDryRun
	}
	
	// Convert KMS key
	if in.KmsKeyName != "" {
		out.KmsKeyRef = &refs.KMSCryptoKeyRef{External: in.KmsKeyName}
	}
	
	return out
}

func Repository_Format_FromProto(mapCtx *direct.MapContext, format pb.Repository_Format) string {
	switch format {
	case pb.Repository_DOCKER:
		return "DOCKER"
	case pb.Repository_MAVEN:
		return "MAVEN"
	case pb.Repository_NPM:
		return "NPM"
	case pb.Repository_APT:
		return "APT"
	case pb.Repository_YUM:
		return "YUM"
	case pb.Repository_PYTHON:
		return "PYTHON"
	case pb.Repository_KFP:
		return "KFP"
	case pb.Repository_GO:
		return "GO"
	case pb.Repository_GENERIC:
		return "GENERIC"
	default:
		return "FORMAT_UNSPECIFIED"
	}
}

func Repository_Mode_FromProto(mapCtx *direct.MapContext, mode pb.Repository_Mode) *string {
	var result string
	switch mode {
	case pb.Repository_STANDARD_REPOSITORY:
		result = "STANDARD_REPOSITORY"
	case pb.Repository_VIRTUAL_REPOSITORY:
		result = "VIRTUAL_REPOSITORY"
	case pb.Repository_REMOTE_REPOSITORY:
		result = "REMOTE_REPOSITORY"
	default:
		result = "STANDARD_REPOSITORY"
	}
	return &result
}

// ToProto functions for writing to GCP

func ArtifactRegistryRepositoryObservedState_FromProto(mapCtx *direct.MapContext, in *pb.Repository) *krm.ArtifactRegistryRepositoryObservedState {
	if in == nil {
		return nil
	}
	out := &krm.ArtifactRegistryRepositoryObservedState{}
	
	if in.CreateTime != nil {
		out.CreateTime = direct.StringTimestamp_FromProto(mapCtx, in.CreateTime)
	}
	if in.UpdateTime != nil {
		out.UpdateTime = direct.StringTimestamp_FromProto(mapCtx, in.UpdateTime)
	}
	// Extract repository name from the full resource name
	// Format: projects/[project]/locations/[location]/repositories/[repository_id]
	if in.Name != "" {
		// Simple extraction from the resource path
		parts := strings.Split(in.Name, "/")
		if len(parts) >= 6 && parts[4] == "repositories" {
			repositoryName := parts[5]
			out.Name = &repositoryName
		}
	}
	
	return out
}

func ArtifactRegistryRepositorySpec_ToProto(mapCtx *direct.MapContext, in *krm.ArtifactRegistryRepositorySpec) *pb.Repository {
	if in == nil {
		return nil
	}
	out := &pb.Repository{}
	
	// Convert format
	out.Format = Repository_Format_ToProto(mapCtx, in.Format)
	
	// Convert description
	out.Description = direct.ValueOf(in.Description)
	
	// Convert mode
	out.Mode = Repository_Mode_ToProto(mapCtx, in.Mode)
	
	// Convert cleanup policies
	if len(in.CleanupPolicies) > 0 {
		out.CleanupPolicies = make(map[string]*pb.CleanupPolicy)
		for _, cleanupPolicy := range in.CleanupPolicies {
			out.CleanupPolicies[cleanupPolicy.Id] = CleanupPolicy_ToProto(mapCtx, &cleanupPolicy)
		}
	}
	
	// Convert cleanup policy dry run
	out.CleanupPolicyDryRun = direct.ValueOf(in.CleanupPolicyDryRun)
	
	// Convert KMS key reference
	if in.KmsKeyRef != nil {
		out.KmsKeyName = in.KmsKeyRef.External
	}
	
	// Convert Docker config
	if in.DockerConfig != nil {
		out.FormatConfig = &pb.Repository_DockerConfig{
			DockerConfig: DockerConfig_ToProto(mapCtx, in.DockerConfig),
		}
	}
	
	// Convert Maven config  
	if in.MavenConfig != nil {
		out.FormatConfig = &pb.Repository_MavenConfig{
			MavenConfig: MavenConfig_ToProto(mapCtx, in.MavenConfig),
		}
	}
	
	// Convert virtual repository config
	if in.VirtualRepositoryConfig != nil {
		out.ModeConfig = &pb.Repository_VirtualRepositoryConfig{
			VirtualRepositoryConfig: VirtualRepositoryConfig_ToProto(mapCtx, in.VirtualRepositoryConfig),
		}
	}
	
	// Convert remote repository config
	if in.RemoteRepositoryConfig != nil {
		out.ModeConfig = &pb.Repository_RemoteRepositoryConfig{
			RemoteRepositoryConfig: RemoteRepositoryConfig_ToProto(mapCtx, in.RemoteRepositoryConfig),
		}
	}
	
	return out
}

func Repository_Format_ToProto(mapCtx *direct.MapContext, format string) pb.Repository_Format {
	switch format {
	case "DOCKER":
		return pb.Repository_DOCKER
	case "MAVEN":
		return pb.Repository_MAVEN
	case "NPM":
		return pb.Repository_NPM
	case "APT":
		return pb.Repository_APT
	case "YUM":
		return pb.Repository_YUM
	case "PYTHON":
		return pb.Repository_PYTHON
	case "KFP":
		return pb.Repository_KFP
	case "GO":
		return pb.Repository_GO
	case "GENERIC":
		return pb.Repository_GENERIC
	default:
		mapCtx.Errorf("unknown format: %s", format)
		return pb.Repository_FORMAT_UNSPECIFIED
	}
}

func Repository_Mode_ToProto(mapCtx *direct.MapContext, mode *string) pb.Repository_Mode {
	if mode == nil {
		return pb.Repository_STANDARD_REPOSITORY // Default mode
	}
	switch *mode {
	case "STANDARD_REPOSITORY":
		return pb.Repository_STANDARD_REPOSITORY
	case "VIRTUAL_REPOSITORY":
		return pb.Repository_VIRTUAL_REPOSITORY
	case "REMOTE_REPOSITORY":
		return pb.Repository_REMOTE_REPOSITORY
	default:
		mapCtx.Errorf("unknown mode: %s", *mode)
		return pb.Repository_STANDARD_REPOSITORY
	}
}

func CleanupPolicy_ToProto(mapCtx *direct.MapContext, in *krm.CleanupPolicies) *pb.CleanupPolicy {
	if in == nil {
		return nil
	}
	out := &pb.CleanupPolicy{}
	
	// Set the ID
	out.Id = in.Id
	
	// Convert action
	if in.Action != nil {
		switch *in.Action {
		case "DELETE":
			out.Action = pb.CleanupPolicy_DELETE
		case "KEEP":
			out.Action = pb.CleanupPolicy_KEEP
		default:
			mapCtx.Errorf("unknown cleanup policy action: %s", *in.Action)
		}
	}
	
	// Convert condition (oneof)
	if in.Condition != nil {
		out.ConditionType = &pb.CleanupPolicy_Condition{
			Condition: CleanupPolicyCondition_ToProto(mapCtx, in.Condition),
		}
	}
	
	// Convert most recent versions (oneof)
	if in.MostRecentVersions != nil {
		out.ConditionType = &pb.CleanupPolicy_MostRecentVersions{
			MostRecentVersions: CleanupPolicyMostRecentVersions_ToProto(mapCtx, in.MostRecentVersions),
		}
	}
	
	return out
}

func CleanupPolicyCondition_ToProto(mapCtx *direct.MapContext, in *krm.CleanupPolicyCondition) *pb.CleanupPolicyCondition {
	if in == nil {
		return nil
	}
	out := &pb.CleanupPolicyCondition{}
	
	// Convert tag state
	if in.TagState != nil {
		var tagState pb.CleanupPolicyCondition_TagState
		switch *in.TagState {
		case "TAGGED":
			tagState = pb.CleanupPolicyCondition_TAGGED
		case "UNTAGGED":
			tagState = pb.CleanupPolicyCondition_UNTAGGED
		case "ANY":
			tagState = pb.CleanupPolicyCondition_TAG_STATE_UNSPECIFIED
		default:
			mapCtx.Errorf("unknown tag state: %s", *in.TagState)
			tagState = pb.CleanupPolicyCondition_TAG_STATE_UNSPECIFIED
		}
		out.TagState = &tagState
	}
	
	// Convert durations
	if in.NewerThan != nil {
		out.NewerThan = direct.Duration_ToProto(mapCtx, in.NewerThan)
	}
	if in.OlderThan != nil {
		out.OlderThan = direct.Duration_ToProto(mapCtx, in.OlderThan)
	}
	
	// Convert string slices
	out.PackageNamePrefixes = in.PackageNamePrefixes
	out.TagPrefixes = in.TagPrefixes
	out.VersionNamePrefixes = in.VersionNamePrefixes
	
	return out
}

func CleanupPolicyMostRecentVersions_ToProto(mapCtx *direct.MapContext, in *krm.CleanupPolicyMostRecentVersions) *pb.CleanupPolicyMostRecentVersions {
	if in == nil {
		return nil
	}
	out := &pb.CleanupPolicyMostRecentVersions{}
	
	if in.KeepCount != nil {
		keepCount := int32(*in.KeepCount)
		out.KeepCount = &keepCount
	}
	out.PackageNamePrefixes = in.PackageNamePrefixes
	
	return out
}

func DockerConfig_ToProto(mapCtx *direct.MapContext, in *krm.DockerConfig) *pb.Repository_DockerRepositoryConfig {
	if in == nil {
		return nil
	}
	out := &pb.Repository_DockerRepositoryConfig{}
	
	if in.ImmutableTags != nil {
		out.ImmutableTags = *in.ImmutableTags
	}
	
	return out
}

func MavenConfig_ToProto(mapCtx *direct.MapContext, in *krm.MavenConfig) *pb.Repository_MavenRepositoryConfig {
	if in == nil {
		return nil
	}
	out := &pb.Repository_MavenRepositoryConfig{}
	
	if in.AllowSnapshotOverwrites != nil {
		out.AllowSnapshotOverwrites = *in.AllowSnapshotOverwrites
	}
	
	if in.VersionPolicy != nil {
		switch *in.VersionPolicy {
		case "VERSION_POLICY_UNSPECIFIED":
			out.VersionPolicy = pb.Repository_MavenRepositoryConfig_VERSION_POLICY_UNSPECIFIED
		case "RELEASE":
			out.VersionPolicy = pb.Repository_MavenRepositoryConfig_RELEASE
		case "SNAPSHOT":
			out.VersionPolicy = pb.Repository_MavenRepositoryConfig_SNAPSHOT
		default:
			mapCtx.Errorf("unknown version policy: %s", *in.VersionPolicy)
		}
	}
	
	return out
}

func VirtualRepositoryConfig_ToProto(mapCtx *direct.MapContext, in *krm.VirtualRepositoryConfig) *pb.VirtualRepositoryConfig {
	if in == nil {
		return nil
	}
	out := &pb.VirtualRepositoryConfig{}
	
	for _, upstreamPolicy := range in.UpstreamPolicies {
		out.UpstreamPolicies = append(out.UpstreamPolicies, UpstreamPolicy_ToProto(mapCtx, &upstreamPolicy))
	}
	
	return out
}

func UpstreamPolicy_ToProto(mapCtx *direct.MapContext, in *krm.UpstreamPolicy) *pb.UpstreamPolicy {
	if in == nil {
		return nil
	}
	out := &pb.UpstreamPolicy{}
	
	if in.Id != nil {
		out.Id = *in.Id
	}
	if in.Priority != nil {
		out.Priority = int32(*in.Priority)
	}
	if in.RepositoryRef != nil {
		out.Repository = in.RepositoryRef.External
	}
	
	return out
}

func RemoteRepositoryConfig_ToProto(mapCtx *direct.MapContext, in *krm.RemoteRepositoryConfig) *pb.RemoteRepositoryConfig {
	if in == nil {
		return nil
	}
	out := &pb.RemoteRepositoryConfig{}
	
	if in.Description != nil {
		out.Description = *in.Description
	}
	
	// Convert specific repository configurations
	if in.DockerRepository != nil {
		out.RemoteSource = &pb.RemoteRepositoryConfig_DockerRepository_{
			DockerRepository: DockerRepository_ToProto(mapCtx, in.DockerRepository),
		}
	}
	if in.MavenRepository != nil {
		out.RemoteSource = &pb.RemoteRepositoryConfig_MavenRepository_{
			MavenRepository: MavenRepository_ToProto(mapCtx, in.MavenRepository),
		}
	}
	if in.NpmRepository != nil {
		out.RemoteSource = &pb.RemoteRepositoryConfig_NpmRepository_{
			NpmRepository: NpmRepository_ToProto(mapCtx, in.NpmRepository),
		}
	}
	if in.PythonRepository != nil {
		out.RemoteSource = &pb.RemoteRepositoryConfig_PythonRepository_{
			PythonRepository: PythonRepository_ToProto(mapCtx, in.PythonRepository),
		}
	}
	
	return out
}

func DockerRepository_ToProto(mapCtx *direct.MapContext, in *krm.DockerRepository) *pb.RemoteRepositoryConfig_DockerRepository {
	if in == nil {
		return nil
	}
	out := &pb.RemoteRepositoryConfig_DockerRepository{}
	
	if in.PublicRepository != nil {
		var publicRepo pb.RemoteRepositoryConfig_DockerRepository_PublicRepository
		switch *in.PublicRepository {
		case "DOCKER_HUB":
			publicRepo = pb.RemoteRepositoryConfig_DockerRepository_DOCKER_HUB
		default:
			mapCtx.Errorf("unknown Docker public repository: %s", *in.PublicRepository)
			publicRepo = pb.RemoteRepositoryConfig_DockerRepository_PUBLIC_REPOSITORY_UNSPECIFIED
		}
		out.Upstream = &pb.RemoteRepositoryConfig_DockerRepository_PublicRepository_{
			PublicRepository: publicRepo,
		}
	}
	
	return out
}

func MavenRepository_ToProto(mapCtx *direct.MapContext, in *krm.MavenRepository) *pb.RemoteRepositoryConfig_MavenRepository {
	if in == nil {
		return nil
	}
	out := &pb.RemoteRepositoryConfig_MavenRepository{}
	
	if in.PublicRepository != nil {
		var publicRepo pb.RemoteRepositoryConfig_MavenRepository_PublicRepository
		switch *in.PublicRepository {
		case "MAVEN_CENTRAL":
			publicRepo = pb.RemoteRepositoryConfig_MavenRepository_MAVEN_CENTRAL
		default:
			mapCtx.Errorf("unknown Maven public repository: %s", *in.PublicRepository)
			publicRepo = pb.RemoteRepositoryConfig_MavenRepository_PUBLIC_REPOSITORY_UNSPECIFIED
		}
		out.Upstream = &pb.RemoteRepositoryConfig_MavenRepository_PublicRepository_{
			PublicRepository: publicRepo,
		}
	}
	
	return out
}

func NpmRepository_ToProto(mapCtx *direct.MapContext, in *krm.NpmRepository) *pb.RemoteRepositoryConfig_NpmRepository {
	if in == nil {
		return nil
	}
	out := &pb.RemoteRepositoryConfig_NpmRepository{}
	
	if in.PublicRepository != nil {
		var publicRepo pb.RemoteRepositoryConfig_NpmRepository_PublicRepository
		switch *in.PublicRepository {
		case "NPMJS":
			publicRepo = pb.RemoteRepositoryConfig_NpmRepository_NPMJS
		default:
			mapCtx.Errorf("unknown NPM public repository: %s", *in.PublicRepository)
			publicRepo = pb.RemoteRepositoryConfig_NpmRepository_PUBLIC_REPOSITORY_UNSPECIFIED
		}
		out.Upstream = &pb.RemoteRepositoryConfig_NpmRepository_PublicRepository_{
			PublicRepository: publicRepo,
		}
	}
	
	return out
}

func PythonRepository_ToProto(mapCtx *direct.MapContext, in *krm.PythonRepository) *pb.RemoteRepositoryConfig_PythonRepository {
	if in == nil {
		return nil
	}
	out := &pb.RemoteRepositoryConfig_PythonRepository{}
	
	if in.PublicRepository != nil {
		var publicRepo pb.RemoteRepositoryConfig_PythonRepository_PublicRepository
		switch *in.PublicRepository {
		case "PYPI":
			publicRepo = pb.RemoteRepositoryConfig_PythonRepository_PYPI
		default:
			mapCtx.Errorf("unknown Python public repository: %s", *in.PublicRepository)
			publicRepo = pb.RemoteRepositoryConfig_PythonRepository_PUBLIC_REPOSITORY_UNSPECIFIED
		}
		out.Upstream = &pb.RemoteRepositoryConfig_PythonRepository_PublicRepository_{
			PublicRepository: publicRepo,
		}
	}
	
	return out
}