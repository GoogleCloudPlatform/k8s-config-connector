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

// +tool:fuzz-gen
// proto.message: google.devtools.artifactregistry.v1.Repository

package artifactregistry

import (
	pb "cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/fuzztesting"
)

func init() {
	fuzztesting.RegisterKRMFuzzer(ArtifactRegistryRepositoryFuzzer())
}

func ArtifactRegistryRepositoryFuzzer() fuzztesting.KRMFuzzer {
	f := fuzztesting.NewKRMTypedFuzzer(&pb.Repository{},
		ArtifactRegistryRepositorySpec_FromProto, ArtifactRegistryRepositorySpec_ToProto,
		ArtifactRegistryRepositoryStatus_FromProto, ArtifactRegistryRepositoryStatus_ToProto,
	)

	f.FilterSpec = func(in *pb.Repository) {
		for id, policy := range in.CleanupPolicies {
			if policy != nil {
				policy.Id = id
				if cond := policy.GetCondition(); cond != nil {
					if cond.TagState != nil && *cond.TagState == pb.CleanupPolicyCondition_TAG_STATE_UNSPECIFIED {
						cond.TagState = nil
					}
				}
			}
		}

		if remoteConfig := in.GetRemoteRepositoryConfig(); remoteConfig != nil {
			if r := remoteConfig.GetDockerRepository(); r != nil {
				if r.GetPublicRepository() == pb.RemoteRepositoryConfig_DockerRepository_PUBLIC_REPOSITORY_UNSPECIFIED {
					r.Upstream = nil
				}
			}
			if r := remoteConfig.GetMavenRepository(); r != nil {
				if r.GetPublicRepository() == pb.RemoteRepositoryConfig_MavenRepository_PUBLIC_REPOSITORY_UNSPECIFIED {
					r.Upstream = nil
				}
			}
			if r := remoteConfig.GetNpmRepository(); r != nil {
				if r.GetPublicRepository() == pb.RemoteRepositoryConfig_NpmRepository_PUBLIC_REPOSITORY_UNSPECIFIED {
					r.Upstream = nil
				}
			}
			if r := remoteConfig.GetPythonRepository(); r != nil {
				if r.GetPublicRepository() == pb.RemoteRepositoryConfig_PythonRepository_PUBLIC_REPOSITORY_UNSPECIFIED {
					r.Upstream = nil
				}
			}
		}
	}

	f.UnimplementedFields.Insert(".annotations")
	f.UnimplementedFields.Insert(".etag")
	f.UnimplementedFields.Insert(".labels")
	f.UnimplementedFields.Insert(".size_bytes")
	f.UnimplementedFields.Insert(".satisfies_pzs")
	f.UnimplementedFields.Insert(".vulnerability_scanning_config")
	f.UnimplementedFields.Insert(".disallow_unspecified_mode")
	f.UnimplementedFields.Insert(".satisfies_pzi")
	f.UnimplementedFields.Insert(".registry_uri")

	// Sub-fields of remote_repository_config that are not mapped in KRM spec
	f.UnimplementedFields.Insert(".remote_repository_config.upstream_credentials")
	f.UnimplementedFields.Insert(".remote_repository_config.disable_upstream_validation")
	f.UnimplementedFields.Insert(".remote_repository_config.apt_repository")
	f.UnimplementedFields.Insert(".remote_repository_config.yum_repository")
	f.UnimplementedFields.Insert(".remote_repository_config.common_repository")
	f.UnimplementedFields.Insert(".remote_repository_config.docker_repository.custom_repository")
	f.UnimplementedFields.Insert(".remote_repository_config.maven_repository.custom_repository")
	f.UnimplementedFields.Insert(".remote_repository_config.npm_repository.custom_repository")
	f.UnimplementedFields.Insert(".remote_repository_config.python_repository.custom_repository")

	f.SpecFields.Insert(".maven_config")
	f.SpecFields.Insert(".docker_config")
	f.SpecFields.Insert(".virtual_repository_config")
	f.SpecFields.Insert(".remote_repository_config")
	f.SpecFields.Insert(".format")
	f.SpecFields.Insert(".description")
	f.SpecFields.Insert(".mode")
	f.SpecFields.Insert(".cleanup_policies")
	f.SpecFields.Insert(".cleanup_policy_dry_run")
	f.SpecFields.Insert(".kms_key_name")

	f.StatusFields.Insert(".name")
	f.StatusFields.Insert(".create_time")
	f.StatusFields.Insert(".update_time")

	return f
}
