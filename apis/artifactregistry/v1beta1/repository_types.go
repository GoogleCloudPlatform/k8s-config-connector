// Copyright 2025 Google LLC
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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ArtifactRegistryRepositorySpec defines the desired state of ArtifactRegistryRepository
// +kcc:spec:proto=google.devtools.artifactregistry.v1.Repository
type ArtifactRegistryRepositorySpec struct {
	// Immutable. The name of the location this repository is located in.
	Location string `json:"location"`

	// The ArtifactRegistryRepository name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Cleanup policies for this repository. Cleanup policies indicate when
	//  certain package versions can be automatically deleted.
	//  Map keys are policy IDs supplied by users during policy creation. They must
	//  unique within a repository and be under 128 characters in length.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.Repository.cleanup_policies
	CleanupPolicies []CleanupPolicy `json:"cleanupPolicies,omitempty"`

	// Optional. If true, the cleanup pipeline is prevented from deleting versions
	//  in this repository.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.Repository.cleanup_policy_dry_run
	CleanupPolicyDryRun *bool `json:"cleanupPolicyDryRun,omitempty"`

	// The user-provided description of the repository.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.Repository.description
	Description *string `json:"description,omitempty"`

	// Docker repository config contains repository level configuration
	//  for the repositories of docker type.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.Repository.docker_config
	DockerConfig *Repository_DockerRepositoryConfig `json:"dockerConfig,omitempty"`

	// Optional. The format of packages that are stored in the repository.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.Repository.format
	Format *string `json:"format,omitempty"`

	// The Cloud KMS resource name of the customer managed encryption key that's
	//  used to encrypt the contents of the Repository. Has the form:
	//  `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`.
	//  This value may not be changed after the Repository has been created.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.Repository.kms_key_name
	KmsKeyRef *refs.KMSCryptoKeyRef `json:"kmsKeyRef,omitempty"`

	// Maven repository config contains repository level configuration
	//  for the repositories of maven type.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.Repository.maven_config
	MavenConfig *Repository_MavenRepositoryConfig `json:"mavenConfig,omitempty"`

	// Optional. The mode of the repository.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.Repository.mode
	Mode *string `json:"mode,omitempty"`

	// Configuration specific for a Remote Repository.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.Repository.remote_repository_config
	RemoteRepositoryConfig *ArtifactRegistryRepositoryRemoteRepositoryConfig `json:"remoteRepositoryConfig,omitempty"`

	// Configuration specific for a Virtual Repository.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.Repository.virtual_repository_config
	VirtualRepositoryConfig *ArtifactRegistryRepositoryVirtualRepositoryConfig `json:"virtualRepositoryConfig,omitempty"`
}

// ArtifactRegistryRepositoryStatus defines the config connector machine state of ArtifactRegistryRepository
type ArtifactRegistryRepositoryStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// Output only. The time when the repository was created.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.Repository.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time when the repository was last updated.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.Repository.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// The name of the repository, for example: "repo1".
	// +kcc:proto:field=google.devtools.artifactregistry.v1.Repository.name
	Name *string `json:"name,omitempty"`
}

// +kcc:proto=google.devtools.artifactregistry.v1.UpstreamPolicy
type ArtifactRegistryRepositoryUpstreamPolicy struct {
	// The user-provided ID of the upstream policy.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.UpstreamPolicy.id
	ID *string `json:"id,omitempty"`

	// Entries with a greater priority value take precedence in the pull order.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.UpstreamPolicy.priority
	Priority *int32 `json:"priority,omitempty"`

	// A reference to the repository resource, for example:
	//  `projects/p1/locations/us-central1/repositories/repo1`.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.UpstreamPolicy.repository
	RepositoryRef *ArtifactRegistryRepositoryRef `json:"repositoryRef,omitempty"`
}

// +kcc:proto=google.devtools.artifactregistry.v1.VirtualRepositoryConfig
type ArtifactRegistryRepositoryVirtualRepositoryConfig struct {
	// Policies that configure the upstream artifacts distributed by the Virtual
	//  Repository. Upstream policies cannot be set on a standard repository.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.VirtualRepositoryConfig.upstream_policies
	UpstreamPolicies []ArtifactRegistryRepositoryUpstreamPolicy `json:"upstreamPolicies,omitempty"`
}

// +kcc:proto=google.devtools.artifactregistry.v1.RemoteRepositoryConfig
type ArtifactRegistryRepositoryRemoteRepositoryConfig struct {
	// Specific settings for a Docker remote repository.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.docker_repository
	DockerRepository *ArtifactRegistryRepositoryDockerRepository `json:"dockerRepository,omitempty"`

	// Specific settings for a Maven remote repository.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.maven_repository
	MavenRepository *ArtifactRegistryRepositoryMavenRepository `json:"mavenRepository,omitempty"`

	// Specific settings for an Npm remote repository.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.npm_repository
	NpmRepository *ArtifactRegistryRepositoryNpmRepository `json:"npmRepository,omitempty"`

	// Specific settings for a Python remote repository.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.python_repository
	PythonRepository *ArtifactRegistryRepositoryPythonRepository `json:"pythonRepository,omitempty"`

	// The description of the remote source.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.description
	Description *string `json:"description,omitempty"`
}

// +kcc:proto=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.DockerRepository
type ArtifactRegistryRepositoryDockerRepository struct {
	// One of the publicly available Docker repositories supported by Artifact
	//  Registry.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.DockerRepository.public_repository
	PublicRepository *string `json:"publicRepository,omitempty"`
}

// +kcc:proto=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.MavenRepository
type ArtifactRegistryRepositoryMavenRepository struct {
	// One of the publicly available Maven repositories supported by Artifact
	//  Registry.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.MavenRepository.public_repository
	PublicRepository *string `json:"publicRepository,omitempty"`
}

// +kcc:proto=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.NpmRepository
type ArtifactRegistryRepositoryNpmRepository struct {
	// One of the publicly available Npm repositories supported by Artifact
	//  Registry.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.NpmRepository.public_repository
	PublicRepository *string `json:"publicRepository,omitempty"`
}

// +kcc:proto=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.PythonRepository
type ArtifactRegistryRepositoryPythonRepository struct {
	// One of the publicly available Python repositories supported by Artifact
	//  Registry.
	// +kcc:proto:field=google.devtools.artifactregistry.v1.RemoteRepositoryConfig.PythonRepository.public_repository
	PublicRepository *string `json:"publicRepository,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpartifactregistryrepository;gcpartifactregistryrepositories
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ArtifactRegistryRepository is the Schema for the ArtifactRegistryRepository API
// +k8s:openapi-gen=true
type ArtifactRegistryRepository struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
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
