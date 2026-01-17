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

// +generated:types
// krm.group: artifactregistry.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.devtools.artifactregistry.v1
// resource: ArtifactRegistryRepository:Repository

package v1beta1

// +kcc:proto=google.devtools.artifactregistry.v1.CleanupPolicy
type CleanupPolicy struct {
	// Policy condition for matching versions.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.CleanupPolicy.condition
	Condition *CleanupPolicyCondition `json:"condition,omitempty"`

	// Policy condition for retaining a minimum number of versions. May only be
	//  specified with a Keep action.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.CleanupPolicy.most_recent_versions
	MostRecentVersions *CleanupPolicyMostRecentVersions `json:"mostRecentVersions,omitempty"`

	// The user-provided ID of the cleanup policy.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.CleanupPolicy.id
	ID *string `json:"id,omitempty"`

	// Policy action.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.CleanupPolicy.action
	Action *string `json:"action,omitempty"`
}

// +kcc:proto=google.devtools.artifactregistry.v1.CleanupPolicyCondition
type CleanupPolicyCondition struct {
	// Match versions by tag status.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.CleanupPolicyCondition.tag_state
	TagState *string `json:"tagState,omitempty"`

	// Match versions by tag prefix. Applied on any prefix match.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.CleanupPolicyCondition.tag_prefixes
	TagPrefixes []string `json:"tagPrefixes,omitempty"`

	// Match versions by version name prefix. Applied on any prefix match.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.CleanupPolicyCondition.version_name_prefixes
	VersionNamePrefixes []string `json:"versionNamePrefixes,omitempty"`

	// Match versions by package prefix. Applied on any prefix match.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.CleanupPolicyCondition.package_name_prefixes
	PackageNamePrefixes []string `json:"packageNamePrefixes,omitempty"`

	// Match versions older than a duration.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.CleanupPolicyCondition.older_than
	OlderThan *string `json:"olderThan,omitempty"`

	// Match versions newer than a duration.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.CleanupPolicyCondition.newer_than
	NewerThan *string `json:"newerThan,omitempty"`
}

// +kcc:proto=google.devtools.artifactregistry.v1.CleanupPolicyMostRecentVersions
type CleanupPolicyMostRecentVersions struct {
	// List of package name prefixes that will apply this rule.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.CleanupPolicyMostRecentVersions.package_name_prefixes
	PackageNamePrefixes []string `json:"packageNamePrefixes,omitempty"`

	// Minimum number of versions to keep.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.CleanupPolicyMostRecentVersions.keep_count
	KeepCount *int32 `json:"keepCount,omitempty"`
}

// +kcc:proto=google.devtools.artifactregistry.v1.RemoteRepositoryConfig
type RemoteRepositoryConfig struct {
	// Specific settings for a Docker remote repository.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.docker_repository
	DockerRepository *RemoteRepositoryConfig_DockerRepository `json:"dockerRepository,omitempty"`

	// Specific settings for a Maven remote repository.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.maven_repository
	MavenRepository *RemoteRepositoryConfig_MavenRepository `json:"mavenRepository,omitempty"`

	// Specific settings for an Npm remote repository.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.npm_repository
	NpmRepository *RemoteRepositoryConfig_NpmRepository `json:"npmRepository,omitempty"`

	// Specific settings for a Python remote repository.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.python_repository
	PythonRepository *RemoteRepositoryConfig_PythonRepository `json:"pythonRepository,omitempty"`

	// Specific settings for an Apt remote repository.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.apt_repository
	AptRepository *RemoteRepositoryConfig_AptRepository `json:"aptRepository,omitempty"`

	// Specific settings for a Yum remote repository.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.yum_repository
	YumRepository *RemoteRepositoryConfig_YumRepository `json:"yumRepository,omitempty"`

	// Common remote repository settings.
	//  Used as the remote repository upstream URL.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.common_repository
	CommonRepository *RemoteRepositoryConfig_CommonRemoteRepository `json:"commonRepository,omitempty"`

	// The description of the remote source.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.description
	Description *string `json:"description,omitempty"`

	// Optional. The credentials used to access the remote repository.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.upstream_credentials
	UpstreamCredentials *RemoteRepositoryConfig_UpstreamCredentials `json:"upstreamCredentials,omitempty"`

	// Input only. A create/update remote repo option to avoid making a HEAD/GET
	//  request to validate a remote repo and any supplied upstream credentials.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.disable_upstream_validation
	DisableUpstreamValidation *bool `json:"disableUpstreamValidation,omitempty"`
}

// +kcc:proto=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.AptRepository
type RemoteRepositoryConfig_AptRepository struct {
	// One of the publicly available Apt repositories supported by Artifact
	//  Registry.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.AptRepository.public_repository
	PublicRepository *RemoteRepositoryConfig_AptRepository_PublicRepository `json:"publicRepository,omitempty"`

	// Customer-specified remote repository.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.AptRepository.custom_repository
	CustomRepository *RemoteRepositoryConfig_AptRepository_CustomRepository `json:"customRepository,omitempty"`
}

// +kcc:proto=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.AptRepository.CustomRepository
type RemoteRepositoryConfig_AptRepository_CustomRepository struct {
	// An http/https uri reference to the upstream remote repository, for ex:
	//  "https://my.apt.registry/".
	// +kcc:proto:field=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.AptRepository.CustomRepository.uri
	URI *string `json:"uri,omitempty"`
}

// +kcc:proto=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.AptRepository.PublicRepository
type RemoteRepositoryConfig_AptRepository_PublicRepository struct {
	// A common public repository base for Apt.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.AptRepository.PublicRepository.repository_base
	RepositoryBase *string `json:"repositoryBase,omitempty"`

	// A custom field to define a path to a specific repository from the base.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.AptRepository.PublicRepository.repository_path
	RepositoryPath *string `json:"repositoryPath,omitempty"`
}

// +kcc:proto=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.CommonRemoteRepository
type RemoteRepositoryConfig_CommonRemoteRepository struct {
	// Required. A common public repository base for remote repository.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.CommonRemoteRepository.uri
	URI *string `json:"uri,omitempty"`
}

// +kcc:proto=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.DockerRepository
type RemoteRepositoryConfig_DockerRepository struct {
	// One of the publicly available Docker repositories supported by Artifact
	//  Registry.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.DockerRepository.public_repository
	PublicRepository *string `json:"publicRepository,omitempty"`

	// Customer-specified remote repository.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.DockerRepository.custom_repository
	CustomRepository *RemoteRepositoryConfig_DockerRepository_CustomRepository `json:"customRepository,omitempty"`
}

// +kcc:proto=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.DockerRepository.CustomRepository
type RemoteRepositoryConfig_DockerRepository_CustomRepository struct {
	// An http/https uri reference to the custom remote repository, for ex:
	//  "https://registry-1.docker.io".
	// +kcc:proto:field=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.DockerRepository.CustomRepository.uri
	URI *string `json:"uri,omitempty"`
}

// +kcc:proto=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.MavenRepository
type RemoteRepositoryConfig_MavenRepository struct {
	// One of the publicly available Maven repositories supported by Artifact
	//  Registry.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.MavenRepository.public_repository
	PublicRepository *string `json:"publicRepository,omitempty"`

	// Customer-specified remote repository.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.MavenRepository.custom_repository
	CustomRepository *RemoteRepositoryConfig_MavenRepository_CustomRepository `json:"customRepository,omitempty"`
}

// +kcc:proto=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.MavenRepository.CustomRepository
type RemoteRepositoryConfig_MavenRepository_CustomRepository struct {
	// An http/https uri reference to the upstream remote repository, for ex:
	//  "https://my.maven.registry/".
	// +kcc:proto:field=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.MavenRepository.CustomRepository.uri
	URI *string `json:"uri,omitempty"`
}

// +kcc:proto=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.NpmRepository
type RemoteRepositoryConfig_NpmRepository struct {
	// One of the publicly available Npm repositories supported by Artifact
	//  Registry.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.NpmRepository.public_repository
	PublicRepository *string `json:"publicRepository,omitempty"`

	// Customer-specified remote repository.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.NpmRepository.custom_repository
	CustomRepository *RemoteRepositoryConfig_NpmRepository_CustomRepository `json:"customRepository,omitempty"`
}

// +kcc:proto=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.NpmRepository.CustomRepository
type RemoteRepositoryConfig_NpmRepository_CustomRepository struct {
	// An http/https uri reference to the upstream remote repository, for ex:
	//  "https://my.npm.registry/".
	// +kcc:proto:field=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.NpmRepository.CustomRepository.uri
	URI *string `json:"uri,omitempty"`
}

// +kcc:proto=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.PythonRepository
type RemoteRepositoryConfig_PythonRepository struct {
	// One of the publicly available Python repositories supported by Artifact
	//  Registry.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.PythonRepository.public_repository
	PublicRepository *string `json:"publicRepository,omitempty"`

	// Customer-specified remote repository.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.PythonRepository.custom_repository
	CustomRepository *RemoteRepositoryConfig_PythonRepository_CustomRepository `json:"customRepository,omitempty"`
}

// +kcc:proto=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.PythonRepository.CustomRepository
type RemoteRepositoryConfig_PythonRepository_CustomRepository struct {
	// An http/https uri reference to the upstream remote repository, for ex:
	//  "https://my.python.registry/".
	// +kcc:proto:field=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.PythonRepository.CustomRepository.uri
	URI *string `json:"uri,omitempty"`
}

// +kcc:proto=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.UpstreamCredentials
type RemoteRepositoryConfig_UpstreamCredentials struct {
	// Use username and password to access the remote repository.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.UpstreamCredentials.username_password_credentials
	UsernamePasswordCredentials *RemoteRepositoryConfig_UpstreamCredentials_UsernamePasswordCredentials `json:"usernamePasswordCredentials,omitempty"`
}

// +kcc:proto=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.UpstreamCredentials.UsernamePasswordCredentials
type RemoteRepositoryConfig_UpstreamCredentials_UsernamePasswordCredentials struct {
	// The username to access the remote repository.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.UpstreamCredentials.UsernamePasswordCredentials.username
	Username *string `json:"username,omitempty"`

	// The Secret Manager key version that holds the password to access the
	//  remote repository. Must be in the format of
	//  `projects/{project}/secrets/{secret}/versions/{version}`.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.UpstreamCredentials.UsernamePasswordCredentials.password_secret_version
	PasswordSecretVersion *string `json:"passwordSecretVersion,omitempty"`
}

// +kcc:proto=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.YumRepository
type RemoteRepositoryConfig_YumRepository struct {
	// One of the publicly available Yum repositories supported by Artifact
	//  Registry.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.YumRepository.public_repository
	PublicRepository *RemoteRepositoryConfig_YumRepository_PublicRepository `json:"publicRepository,omitempty"`

	// Customer-specified remote repository.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.YumRepository.custom_repository
	CustomRepository *RemoteRepositoryConfig_YumRepository_CustomRepository `json:"customRepository,omitempty"`
}

// +kcc:proto=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.YumRepository.CustomRepository
type RemoteRepositoryConfig_YumRepository_CustomRepository struct {
	// An http/https uri reference to the upstream remote repository, for ex:
	//  "https://my.yum.registry/".
	// +kcc:proto:field=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.YumRepository.CustomRepository.uri
	URI *string `json:"uri,omitempty"`
}

// +kcc:proto=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.YumRepository.PublicRepository
type RemoteRepositoryConfig_YumRepository_PublicRepository struct {
	// A common public repository base for Yum.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.YumRepository.PublicRepository.repository_base
	RepositoryBase *string `json:"repositoryBase,omitempty"`

	// A custom field to define a path to a specific repository from the base.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.YumRepository.PublicRepository.repository_path
	RepositoryPath *string `json:"repositoryPath,omitempty"`
}

// +kcc:proto=google.devtools.artifactregistry.v1.Repository.DockerRepositoryConfig
type Repository_DockerRepositoryConfig struct {
	// The repository which enabled this flag prevents all tags from being
	//  modified, moved or deleted. This does not prevent tags from being
	//  created.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.Repository.DockerRepositoryConfig.immutable_tags
	ImmutableTags *bool `json:"immutableTags,omitempty"`
}

// +kcc:proto=google.devtools.artifactregistry.v1.Repository.MavenRepositoryConfig
type Repository_MavenRepositoryConfig struct {
	// The repository with this flag will allow publishing
	//  the same snapshot versions.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.Repository.MavenRepositoryConfig.allow_snapshot_overwrites
	AllowSnapshotOverwrites *bool `json:"allowSnapshotOverwrites,omitempty"`

	// Version policy defines the versions that the registry will accept.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.Repository.MavenRepositoryConfig.version_policy
	VersionPolicy *string `json:"versionPolicy,omitempty"`
}

// +kcc:proto=google.devtools.artifactregistry.v1.Repository.VulnerabilityScanningConfig
type Repository_VulnerabilityScanningConfig struct {
	// Optional. Config for whether this repository has vulnerability scanning
	//  disabled.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.Repository.VulnerabilityScanningConfig.enablement_config
	EnablementConfig *string `json:"enablementConfig,omitempty"`
}

// +kcc:proto=google.devtools.artifactregistry.v1.UpstreamPolicy
type UpstreamPolicy struct {
	// The user-provided ID of the upstream policy.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.UpstreamPolicy.id
	ID *string `json:"id,omitempty"`

	// A reference to the repository resource, for example:
	//  `projects/p1/locations/us-central1/repositories/repo1`.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.UpstreamPolicy.repository
	Repository *string `json:"repository,omitempty"`

	// Entries with a greater priority value take precedence in the pull order.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.UpstreamPolicy.priority
	Priority *int32 `json:"priority,omitempty"`
}

// +kcc:proto=google.devtools.artifactregistry.v1.VirtualRepositoryConfig
type VirtualRepositoryConfig struct {
	// Policies that configure the upstream artifacts distributed by the Virtual
	//  Repository. Upstream policies cannot be set on a standard repository.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.VirtualRepositoryConfig.upstream_policies
	UpstreamPolicies []UpstreamPolicy `json:"upstreamPolicies,omitempty"`
}

// +kcc:observedstate:proto=google.devtools.artifactregistry.v1.Repository.VulnerabilityScanningConfig
type Repository_VulnerabilityScanningConfigObservedState struct {
	// Output only. The last time this repository config was enabled.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.Repository.VulnerabilityScanningConfig.last_enable_time
	LastEnableTime *string `json:"lastEnableTime,omitempty"`

	// Output only. State of feature enablement, combining repository enablement
	//  config and API enablement state.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.Repository.VulnerabilityScanningConfig.enablement_state
	EnablementState *string `json:"enablementState,omitempty"`

	// Output only. Reason for the repository state.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.Repository.VulnerabilityScanningConfig.enablement_state_reason
	EnablementStateReason *string `json:"enablementStateReason,omitempty"`
}
