// Copyright 2024 Google LLC
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

package v1beta1

import (
	refv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ArtifactRegistryRepositoryGVK = GroupVersion.WithKind("ArtifactRegistryRepository")

// CleanupPolicies defines cleanup policies for package versions
// +kcc:proto=google.devtools.artifactregistry.v1.CleanupPolicy
type CleanupPolicies struct {
	/* Policy action. Possible values: ["DELETE", "KEEP"]. */
	// +optional
	Action *string `json:"action,omitempty"`

	/* Policy condition for matching versions. */
	// +optional
	Condition *CleanupPolicyCondition `json:"condition,omitempty"`

	Id string `json:"id"`

	/* Policy condition for retaining a minimum number of versions. May only be
	specified with a Keep action. */
	// +optional
	MostRecentVersions *CleanupPolicyMostRecentVersions `json:"mostRecentVersions,omitempty"`
}

// CleanupPolicyCondition defines conditions for cleanup policies
// +kcc:proto=google.devtools.artifactregistry.v1.CleanupPolicyCondition
type CleanupPolicyCondition struct {
	/* Match versions newer than a duration. */
	// +optional
	NewerThan *string `json:"newerThan,omitempty"`

	/* Match versions older than a duration. */
	// +optional
	OlderThan *string `json:"olderThan,omitempty"`

	/* Match versions by package prefix. Applied on any prefix match. */
	// +optional
	PackageNamePrefixes []string `json:"packageNamePrefixes,omitempty"`

	/* Match versions by tag prefix. Applied on any prefix match. */
	// +optional
	TagPrefixes []string `json:"tagPrefixes,omitempty"`

	/* Match versions by tag status. Default value: "ANY" Possible values: ["TAGGED", "UNTAGGED", "ANY"]. */
	// +optional
	TagState *string `json:"tagState,omitempty"`

	/* Match versions by version name prefix. Applied on any prefix match. */
	// +optional
	VersionNamePrefixes []string `json:"versionNamePrefixes,omitempty"`
}

// CleanupPolicyMostRecentVersions defines policy for retaining recent versions
// +kcc:proto=google.devtools.artifactregistry.v1.CleanupPolicyMostRecentVersions
type CleanupPolicyMostRecentVersions struct {
	/* Minimum number of versions to keep. */
	// +optional
	KeepCount *int64 `json:"keepCount,omitempty"`

	/* Match versions by package prefix. Applied on any prefix match. */
	// +optional
	PackageNamePrefixes []string `json:"packageNamePrefixes,omitempty"`
}

// DockerConfig defines Docker repository configuration
// +kcc:proto=google.devtools.artifactregistry.v1.DockerRepositoryConfig
type DockerConfig struct {
	/* The repository which enabled this flag prevents all tags from being modified, moved or deleted. This does not prevent tags from being created. */
	// +optional
	ImmutableTags *bool `json:"immutableTags,omitempty"`
}

// MavenConfig defines Maven repository configuration
// +kcc:proto=google.devtools.artifactregistry.v1.MavenRepositoryConfig
type MavenConfig struct {
	/* Immutable. The repository with this flag will allow publishing the same
	snapshot versions. */
	// +optional
	AllowSnapshotOverwrites *bool `json:"allowSnapshotOverwrites,omitempty"`

	/* Immutable. Version policy defines the versions that the registry will accept. Default value: "VERSION_POLICY_UNSPECIFIED" Possible values: ["VERSION_POLICY_UNSPECIFIED", "RELEASE", "SNAPSHOT"]. */
	// +optional
	VersionPolicy *string `json:"versionPolicy,omitempty"`
}

// DockerRepository defines Docker remote repository settings
// +kcc:proto=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.DockerRepository
type DockerRepository struct {
	/* Immutable. Address of the remote repository. Default value: "DOCKER_HUB" Possible values: ["DOCKER_HUB"]. */
	// +optional
	PublicRepository *string `json:"publicRepository,omitempty"`
}

// MavenRepository defines Maven remote repository settings
// +kcc:proto=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.MavenRepository
type MavenRepository struct {
	/* Immutable. Address of the remote repository. Default value: "MAVEN_CENTRAL" Possible values: ["MAVEN_CENTRAL"]. */
	// +optional
	PublicRepository *string `json:"publicRepository,omitempty"`
}

// NpmRepository defines NPM remote repository settings
// +kcc:proto=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.NpmRepository
type NpmRepository struct {
	/* Immutable. Address of the remote repository. Default value: "NPMJS" Possible values: ["NPMJS"]. */
	// +optional
	PublicRepository *string `json:"publicRepository,omitempty"`
}

// PythonRepository defines Python remote repository settings
// +kcc:proto=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.PythonRepository
type PythonRepository struct {
	/* Immutable. Address of the remote repository. Default value: "PYPI" Possible values: ["PYPI"]. */
	// +optional
	PublicRepository *string `json:"publicRepository,omitempty"`
}

// RemoteRepositoryConfig defines configuration for remote repositories
// +kcc:proto=google.devtools.artifactregistry.v1.RemoteRepositoryConfig
type RemoteRepositoryConfig struct {
	/* Immutable. The description of the remote source. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Immutable. Specific settings for a Docker remote repository. */
	// +optional
	DockerRepository *DockerRepository `json:"dockerRepository,omitempty"`

	/* Immutable. Specific settings for a Maven remote repository. */
	// +optional
	MavenRepository *MavenRepository `json:"mavenRepository,omitempty"`

	/* Immutable. Specific settings for an Npm remote repository. */
	// +optional
	NpmRepository *NpmRepository `json:"npmRepository,omitempty"`

	/* Immutable. Specific settings for a Python remote repository. */
	// +optional
	PythonRepository *PythonRepository `json:"pythonRepository,omitempty"`
}

// UpstreamPolicy defines upstream policies for virtual repositories
// +kcc:proto=google.devtools.artifactregistry.v1.UpstreamPolicy
type UpstreamPolicy struct {
	/* The user-provided ID of the upstream policy. */
	// +optional
	Id *string `json:"id,omitempty"`

	/* Entries with a greater priority value take precedence in the pull order. */
	// +optional
	Priority *int64 `json:"priority,omitempty"`

	/* A reference to the repository resource, for example:
	"projects/p1/locations/us-central1/repositories/repo1". */
	// +optional
	RepositoryRef *refv1beta1.ArtifactRegistryRepositoryRef `json:"repositoryRef,omitempty"`
}

// VirtualRepositoryConfig defines configuration for virtual repositories
// +kcc:proto=google.devtools.artifactregistry.v1.VirtualRepositoryConfig
type VirtualRepositoryConfig struct {
	/* Policies that configure the upstream artifacts distributed by the Virtual
	Repository. Upstream policies cannot be set on a standard repository. */
	// +optional
	UpstreamPolicies []UpstreamPolicy `json:"upstreamPolicies,omitempty"`
}

// ArtifactRegistryRepositorySpec defines the desired state of ArtifactRegistryRepository
// +kcc:proto=google.devtools.artifactregistry.v1.Repository
type ArtifactRegistryRepositorySpec struct {
	/* Cleanup policies for this repository. Cleanup policies indicate when
	certain package versions can be automatically deleted.
	Map keys are policy IDs supplied by users during policy creation. They must
	unique within a repository and be under 128 characters in length. */
	// +optional
	CleanupPolicies []CleanupPolicies `json:"cleanupPolicies,omitempty"`

	/* If true, the cleanup pipeline is prevented from deleting versions in this
	repository. */
	// +optional
	CleanupPolicyDryRun *bool `json:"cleanupPolicyDryRun,omitempty"`

	/* The user-provided description of the repository. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Docker repository config contains repository level configuration for the repositories of docker type. */
	// +optional
	DockerConfig *DockerConfig `json:"dockerConfig,omitempty"`

	/* Immutable. The format of packages that are stored in the repository. Supported formats
	can be found [here](https://cloud.google.com/artifact-registry/docs/supported-formats).
	You can only create alpha formats if you are a member of the
	[alpha user group](https://cloud.google.com/artifact-registry/docs/supported-formats#alpha-access). */
	// +required
	Format string `json:"format"`

	/* The customer managed encryption key that's used to encrypt the
	contents of the Repository. */
	// +optional
	KmsKeyRef *refv1beta1.KMSCryptoKeyRef `json:"kmsKeyRef,omitempty"`

	/* Immutable. The name of the location this repository is located in. */
	// +required
	Location string `json:"location"`

	/* MavenRepositoryConfig is maven related repository details.
	Provides additional configuration details for repositories of the maven
	format type. */
	// +optional
	MavenConfig *MavenConfig `json:"mavenConfig,omitempty"`

	/* Immutable. The mode configures the repository to serve artifacts from different sources. Default value: "STANDARD_REPOSITORY" Possible values: ["STANDARD_REPOSITORY", "VIRTUAL_REPOSITORY", "REMOTE_REPOSITORY"]. */
	// +optional
	Mode *string `json:"mode,omitempty"`

	/* Immutable. Configuration specific for a Remote Repository. */
	// +optional
	RemoteRepositoryConfig *RemoteRepositoryConfig `json:"remoteRepositoryConfig,omitempty"`

	/* Immutable. Optional. The repositoryId of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Configuration specific for a Virtual Repository. */
	// +optional
	VirtualRepositoryConfig *VirtualRepositoryConfig `json:"virtualRepositoryConfig,omitempty"`
}

// ArtifactRegistryRepositoryStatus defines the config connector machine state of ArtifactRegistryRepository
type ArtifactRegistryRepositoryStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// The time when the repository was created.
	CreateTime *string `json:"createTime,omitempty"`

	// The name of the repository, for example: "repo1".
	Name *string `json:"name,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// The time when the repository was last updated.
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpartifactregistryrepository;gcpartifactregistryrepositories
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ArtifactRegistryRepository is the Schema for the artifactregistry API
// +k8s:openapi-gen=true
// +kubebuilder:storageversion
type ArtifactRegistryRepository struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ArtifactRegistryRepositorySpec   `json:"spec,omitempty"`
	Status ArtifactRegistryRepositoryStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ArtifactRegistryRepositoryList contains a list of ArtifactRegistryRepository
type ArtifactRegistryRepositoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ArtifactRegistryRepository `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ArtifactRegistryRepository{}, &ArtifactRegistryRepositoryList{})
}
