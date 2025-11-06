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

	"cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/artifactregistry/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

// FromProto functions for reading from GCP

func ArtifactRegistryRepositorySpec_FromProto(mapCtx *direct.MapContext, in *artifactregistrypb.Repository) *krm.ArtifactRegistryRepositorySpec {
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

func Repository_Format_FromProto(mapCtx *direct.MapContext, format artifactregistrypb.Repository_Format) string {
	switch format {
	case artifactregistrypb.Repository_DOCKER:
		return "DOCKER"
	case artifactregistrypb.Repository_MAVEN:
		return "MAVEN"
	case artifactregistrypb.Repository_NPM:
		return "NPM"
	case artifactregistrypb.Repository_APT:
		return "APT"
	case artifactregistrypb.Repository_YUM:
		return "YUM"
	case artifactregistrypb.Repository_PYTHON:
		return "PYTHON"
	case artifactregistrypb.Repository_KFP:
		return "KFP"
	case artifactregistrypb.Repository_GO:
		return "GO"
	case artifactregistrypb.Repository_GENERIC:
		return "GENERIC"
	default:
		return "FORMAT_UNSPECIFIED"
	}
}

func Repository_Mode_FromProto(mapCtx *direct.MapContext, mode artifactregistrypb.Repository_Mode) *string {
	var result string
	switch mode {
	case artifactregistrypb.Repository_STANDARD_REPOSITORY:
		result = "STANDARD_REPOSITORY"
	case artifactregistrypb.Repository_VIRTUAL_REPOSITORY:
		result = "VIRTUAL_REPOSITORY"
	case artifactregistrypb.Repository_REMOTE_REPOSITORY:
		result = "REMOTE_REPOSITORY"
	default:
		result = "STANDARD_REPOSITORY"
	}
	return &result
}

// ToProto functions for writing to GCP

func ArtifactRegistryRepositoryObservedState_FromProto(mapCtx *direct.MapContext, in *artifactregistrypb.Repository) *krm.ArtifactRegistryRepositoryObservedState {
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

func ArtifactRegistryRepositorySpec_ToProto(mapCtx *direct.MapContext, in *krm.ArtifactRegistryRepositorySpec) *artifactregistrypb.Repository {
	if in == nil {
		return nil
	}
	out := &artifactregistrypb.Repository{}

	// Convert format
	out.Format = Repository_Format_ToProto(mapCtx, in.Format)

	// Convert description
	out.Description = direct.ValueOf(in.Description)

	// Convert mode
	out.Mode = Repository_Mode_ToProto(mapCtx, in.Mode)

	// Convert cleanup policies
	if len(in.CleanupPolicies) > 0 {
		out.CleanupPolicies = make(map[string]*artifactregistrypb.CleanupPolicy)
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
		out.FormatConfig = &artifactregistrypb.Repository_DockerConfig{
			DockerConfig: DockerConfig_ToProto(mapCtx, in.DockerConfig),
		}
	}

	// Convert Maven config
	if in.MavenConfig != nil {
		out.FormatConfig = &artifactregistrypb.Repository_MavenConfig{
			MavenConfig: MavenConfig_ToProto(mapCtx, in.MavenConfig),
		}
	}

	// Convert virtual repository config
	if in.VirtualRepositoryConfig != nil {
		out.ModeConfig = &artifactregistrypb.Repository_VirtualRepositoryConfig{
			VirtualRepositoryConfig: VirtualRepositoryConfig_ToProto(mapCtx, in.VirtualRepositoryConfig),
		}
	}

	// Convert remote repository config
	if in.RemoteRepositoryConfig != nil {
		out.ModeConfig = &artifactregistrypb.Repository_RemoteRepositoryConfig{
			RemoteRepositoryConfig: RemoteRepositoryConfig_ToProto(mapCtx, in.RemoteRepositoryConfig),
		}
	}

	return out
}

func Repository_Format_ToProto(mapCtx *direct.MapContext, format string) artifactregistrypb.Repository_Format {
	switch format {
	case "DOCKER":
		return artifactregistrypb.Repository_DOCKER
	case "MAVEN":
		return artifactregistrypb.Repository_MAVEN
	case "NPM":
		return artifactregistrypb.Repository_NPM
	case "APT":
		return artifactregistrypb.Repository_APT
	case "YUM":
		return artifactregistrypb.Repository_YUM
	case "PYTHON":
		return artifactregistrypb.Repository_PYTHON
	case "KFP":
		return artifactregistrypb.Repository_KFP
	case "GO":
		return artifactregistrypb.Repository_GO
	case "GENERIC":
		return artifactregistrypb.Repository_GENERIC
	default:
		mapCtx.Errorf("unknown format: %s", format)
		return artifactregistrypb.Repository_FORMAT_UNSPECIFIED
	}
}

func Repository_Mode_ToProto(mapCtx *direct.MapContext, mode *string) artifactregistrypb.Repository_Mode {
	if mode == nil {
		return artifactregistrypb.Repository_STANDARD_REPOSITORY // Default mode
	}
	switch *mode {
	case "STANDARD_REPOSITORY":
		return artifactregistrypb.Repository_STANDARD_REPOSITORY
	case "VIRTUAL_REPOSITORY":
		return artifactregistrypb.Repository_VIRTUAL_REPOSITORY
	case "REMOTE_REPOSITORY":
		return artifactregistrypb.Repository_REMOTE_REPOSITORY
	default:
		mapCtx.Errorf("unknown mode: %s", *mode)
		return artifactregistrypb.Repository_STANDARD_REPOSITORY
	}
}

func CleanupPolicy_ToProto(mapCtx *direct.MapContext, in *krm.CleanupPolicies) *artifactregistrypb.CleanupPolicy {
	if in == nil {
		return nil
	}
	out := &artifactregistrypb.CleanupPolicy{}

	// Set the ID
	out.Id = in.Id

	// Convert action
	if in.Action != nil {
		switch *in.Action {
		case "DELETE":
			out.Action = artifactregistrypb.CleanupPolicy_DELETE
		case "KEEP":
			out.Action = artifactregistrypb.CleanupPolicy_KEEP
		default:
			mapCtx.Errorf("unknown cleanup policy action: %s", *in.Action)
		}
	}

	// Convert condition (oneof)
	if in.Condition != nil {
		out.ConditionType = &artifactregistrypb.CleanupPolicy_Condition{
			Condition: CleanupPolicyCondition_ToProto(mapCtx, in.Condition),
		}
	}

	// Convert most recent versions (oneof)
	if in.MostRecentVersions != nil {
		out.ConditionType = &artifactregistrypb.CleanupPolicy_MostRecentVersions{
			MostRecentVersions: CleanupPolicyMostRecentVersions_ToProto(mapCtx, in.MostRecentVersions),
		}
	}

	return out
}

func CleanupPolicyCondition_ToProto(mapCtx *direct.MapContext, in *krm.CleanupPolicyCondition) *artifactregistrypb.CleanupPolicyCondition {
	if in == nil {
		return nil
	}
	out := &artifactregistrypb.CleanupPolicyCondition{}

	// Convert tag state
	if in.TagState != nil {
		var tagState artifactregistrypb.CleanupPolicyCondition_TagState
		switch *in.TagState {
		case "TAGGED":
			tagState = artifactregistrypb.CleanupPolicyCondition_TAGGED
		case "UNTAGGED":
			tagState = artifactregistrypb.CleanupPolicyCondition_UNTAGGED
		case "ANY":
			tagState = artifactregistrypb.CleanupPolicyCondition_TAG_STATE_UNSPECIFIED
		default:
			mapCtx.Errorf("unknown tag state: %s", *in.TagState)
			tagState = artifactregistrypb.CleanupPolicyCondition_TAG_STATE_UNSPECIFIED
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

func CleanupPolicyMostRecentVersions_ToProto(mapCtx *direct.MapContext, in *krm.CleanupPolicyMostRecentVersions) *artifactregistrypb.CleanupPolicyMostRecentVersions {
	if in == nil {
		return nil
	}
	out := &artifactregistrypb.CleanupPolicyMostRecentVersions{}

	if in.KeepCount != nil {
		keepCount := int32(*in.KeepCount)
		out.KeepCount = &keepCount
	}
	out.PackageNamePrefixes = in.PackageNamePrefixes

	return out
}

func DockerConfig_ToProto(mapCtx *direct.MapContext, in *krm.DockerConfig) *artifactregistrypb.Repository_DockerRepositoryConfig {
	if in == nil {
		return nil
	}
	out := &artifactregistrypb.Repository_DockerRepositoryConfig{}

	if in.ImmutableTags != nil {
		out.ImmutableTags = *in.ImmutableTags
	}

	return out
}

func MavenConfig_ToProto(mapCtx *direct.MapContext, in *krm.MavenConfig) *artifactregistrypb.Repository_MavenRepositoryConfig {
	if in == nil {
		return nil
	}
	out := &artifactregistrypb.Repository_MavenRepositoryConfig{}

	if in.AllowSnapshotOverwrites != nil {
		out.AllowSnapshotOverwrites = *in.AllowSnapshotOverwrites
	}

	if in.VersionPolicy != nil {
		switch *in.VersionPolicy {
		case "VERSION_POLICY_UNSPECIFIED":
			out.VersionPolicy = artifactregistrypb.Repository_MavenRepositoryConfig_VERSION_POLICY_UNSPECIFIED
		case "RELEASE":
			out.VersionPolicy = artifactregistrypb.Repository_MavenRepositoryConfig_RELEASE
		case "SNAPSHOT":
			out.VersionPolicy = artifactregistrypb.Repository_MavenRepositoryConfig_SNAPSHOT
		default:
			mapCtx.Errorf("unknown version policy: %s", *in.VersionPolicy)
		}
	}

	return out
}

func VirtualRepositoryConfig_ToProto(mapCtx *direct.MapContext, in *krm.VirtualRepositoryConfig) *artifactregistrypb.VirtualRepositoryConfig {
	if in == nil {
		return nil
	}
	out := &artifactregistrypb.VirtualRepositoryConfig{}

	for _, upstreamPolicy := range in.UpstreamPolicies {
		out.UpstreamPolicies = append(out.UpstreamPolicies, UpstreamPolicy_ToProto(mapCtx, &upstreamPolicy))
	}

	return out
}

func UpstreamPolicy_ToProto(mapCtx *direct.MapContext, in *krm.UpstreamPolicy) *artifactregistrypb.UpstreamPolicy {
	if in == nil {
		return nil
	}
	out := &artifactregistrypb.UpstreamPolicy{}

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

func RemoteRepositoryConfig_ToProto(mapCtx *direct.MapContext, in *krm.RemoteRepositoryConfig) *artifactregistrypb.RemoteRepositoryConfig {
	if in == nil {
		return nil
	}
	out := &artifactregistrypb.RemoteRepositoryConfig{}

	if in.Description != nil {
		out.Description = *in.Description
	}

	// Convert specific repository configurations
	if in.DockerRepository != nil {
		out.RemoteSource = &artifactregistrypb.RemoteRepositoryConfig_DockerRepository_{
			DockerRepository: DockerRepository_ToProto(mapCtx, in.DockerRepository),
		}
	}
	if in.MavenRepository != nil {
		out.RemoteSource = &artifactregistrypb.RemoteRepositoryConfig_MavenRepository_{
			MavenRepository: MavenRepository_ToProto(mapCtx, in.MavenRepository),
		}
	}
	if in.NpmRepository != nil {
		out.RemoteSource = &artifactregistrypb.RemoteRepositoryConfig_NpmRepository_{
			NpmRepository: NpmRepository_ToProto(mapCtx, in.NpmRepository),
		}
	}
	if in.PythonRepository != nil {
		out.RemoteSource = &artifactregistrypb.RemoteRepositoryConfig_PythonRepository_{
			PythonRepository: PythonRepository_ToProto(mapCtx, in.PythonRepository),
		}
	}

	return out
}

func DockerRepository_ToProto(mapCtx *direct.MapContext, in *krm.DockerRepository) *artifactregistrypb.RemoteRepositoryConfig_DockerRepository {
	if in == nil {
		return nil
	}
	out := &artifactregistrypb.RemoteRepositoryConfig_DockerRepository{}

	if in.PublicRepository != nil {
		var publicRepo artifactregistrypb.RemoteRepositoryConfig_DockerRepository_PublicRepository
		switch *in.PublicRepository {
		case "DOCKER_HUB":
			publicRepo = artifactregistrypb.RemoteRepositoryConfig_DockerRepository_DOCKER_HUB
		default:
			mapCtx.Errorf("unknown Docker public repository: %s", *in.PublicRepository)
			publicRepo = artifactregistrypb.RemoteRepositoryConfig_DockerRepository_PUBLIC_REPOSITORY_UNSPECIFIED
		}
		out.Upstream = &artifactregistrypb.RemoteRepositoryConfig_DockerRepository_PublicRepository_{
			PublicRepository: publicRepo,
		}
	}

	return out
}

func MavenRepository_ToProto(mapCtx *direct.MapContext, in *krm.MavenRepository) *artifactregistrypb.RemoteRepositoryConfig_MavenRepository {
	if in == nil {
		return nil
	}
	out := &artifactregistrypb.RemoteRepositoryConfig_MavenRepository{}

	if in.PublicRepository != nil {
		var publicRepo artifactregistrypb.RemoteRepositoryConfig_MavenRepository_PublicRepository
		switch *in.PublicRepository {
		case "MAVEN_CENTRAL":
			publicRepo = artifactregistrypb.RemoteRepositoryConfig_MavenRepository_MAVEN_CENTRAL
		default:
			mapCtx.Errorf("unknown Maven public repository: %s", *in.PublicRepository)
			publicRepo = artifactregistrypb.RemoteRepositoryConfig_MavenRepository_PUBLIC_REPOSITORY_UNSPECIFIED
		}
		out.Upstream = &artifactregistrypb.RemoteRepositoryConfig_MavenRepository_PublicRepository_{
			PublicRepository: publicRepo,
		}
	}

	return out
}

func NpmRepository_ToProto(mapCtx *direct.MapContext, in *krm.NpmRepository) *artifactregistrypb.RemoteRepositoryConfig_NpmRepository {
	if in == nil {
		return nil
	}
	out := &artifactregistrypb.RemoteRepositoryConfig_NpmRepository{}

	if in.PublicRepository != nil {
		var publicRepo artifactregistrypb.RemoteRepositoryConfig_NpmRepository_PublicRepository
		switch *in.PublicRepository {
		case "NPMJS":
			publicRepo = artifactregistrypb.RemoteRepositoryConfig_NpmRepository_NPMJS
		default:
			mapCtx.Errorf("unknown NPM public repository: %s", *in.PublicRepository)
			publicRepo = artifactregistrypb.RemoteRepositoryConfig_NpmRepository_PUBLIC_REPOSITORY_UNSPECIFIED
		}
		out.Upstream = &artifactregistrypb.RemoteRepositoryConfig_NpmRepository_PublicRepository_{
			PublicRepository: publicRepo,
		}
	}

	return out
}

func PythonRepository_ToProto(mapCtx *direct.MapContext, in *krm.PythonRepository) *artifactregistrypb.RemoteRepositoryConfig_PythonRepository {
	if in == nil {
		return nil
	}
	out := &artifactregistrypb.RemoteRepositoryConfig_PythonRepository{}

	if in.PublicRepository != nil {
		var publicRepo artifactregistrypb.RemoteRepositoryConfig_PythonRepository_PublicRepository
		switch *in.PublicRepository {
		case "PYPI":
			publicRepo = artifactregistrypb.RemoteRepositoryConfig_PythonRepository_PYPI
		default:
			mapCtx.Errorf("unknown Python public repository: %s", *in.PublicRepository)
			publicRepo = artifactregistrypb.RemoteRepositoryConfig_PythonRepository_PUBLIC_REPOSITORY_UNSPECIFIED
		}
		out.Upstream = &artifactregistrypb.RemoteRepositoryConfig_PythonRepository_PublicRepository_{
			PublicRepository: publicRepo,
		}
	}

	return out
}
