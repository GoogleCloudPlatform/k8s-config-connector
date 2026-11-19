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

// API sources for CloudDeployRelease, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/deploy/v1/cloud_deploy.proto
// +kcc:source:service-docs=https://cloud.google.com/deploy/
// +kcc:source:resource-docs=https://cloud.google.com/deploy/docs/reference/rest/v1/projects.locations.deliveryPipelines.releases

package v1alpha1

import (
	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/clouddeploy/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CloudDeployReleaseGVK = GroupVersion.WithKind("CloudDeployRelease")

// CloudDeployReleaseSpec defines the desired state of CloudDeployRelease
// +kcc:spec:proto=google.cloud.deploy.v1.Release
// +kcc:required-from-proto
type CloudDeployReleaseSpec struct {
	// The DeliveryPipeline this Release is created under.
	// +required
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="parentRef is immutable"
	ParentRef *v1beta1.DeliveryPipelineRef `json:"parentRef"`

	// The CloudDeployRelease name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Description of the `Release`. Max length is 255 characters.
	// +kcc:proto:field=google.cloud.deploy.v1.Release.description
	Description *string `json:"description,omitempty"`

	// Optional. User annotations. These attributes can only be set and used by
	//  the user, and not by Cloud Deploy. See
	//  https://google.aip.dev/128#annotations for more details such as format and
	//  size limitations.
	// +kcc:proto:field=google.cloud.deploy.v1.Release.annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// Optional. Cloud Storage URI of tar.gz archive containing Skaffold
	//  configuration.
	// +kcc:proto:field=google.cloud.deploy.v1.Release.skaffold_config_uri
	SkaffoldConfigURI *string `json:"skaffoldConfigURI,omitempty"`

	// Optional. Filepath of the Skaffold config inside of the config URI.
	// +kcc:proto:field=google.cloud.deploy.v1.Release.skaffold_config_path
	SkaffoldConfigPath *string `json:"skaffoldConfigPath,omitempty"`

	// Optional. List of artifacts to pass through to Skaffold command.
	// +kcc:proto:field=google.cloud.deploy.v1.Release.build_artifacts
	BuildArtifacts []BuildArtifact `json:"buildArtifacts,omitempty"`

	// Optional. The Skaffold version to use when operating on this release, such
	//  as "1.20.0". Not all versions are valid; Cloud Deploy supports a specific
	//  set of versions.
	//
	//  If unset, the most recent supported Skaffold version will be used.
	// +kcc:proto:field=google.cloud.deploy.v1.Release.skaffold_version
	SkaffoldVersion *string `json:"skaffoldVersion,omitempty"`

	// Optional. The deploy parameters to use for all targets in this release.
	// +kcc:proto:field=google.cloud.deploy.v1.Release.deploy_parameters
	DeployParameters map[string]string `json:"deployParameters,omitempty"`
}

// CloudDeployReleaseStatus defines the config connector machine state of CloudDeployRelease
type CloudDeployReleaseStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the CloudDeployRelease resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *CloudDeployReleaseObservedState `json:"observedState,omitempty"`
}

// CloudDeployReleaseObservedState is the state of the CloudDeployRelease resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.deploy.v1.Release
type CloudDeployReleaseObservedState struct {
	// Output only. Unique identifier of the `Release`.
	// +kcc:proto:field=google.cloud.deploy.v1.Release.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. Indicates whether this is an abandoned release.
	// +kcc:proto:field=google.cloud.deploy.v1.Release.abandoned
	Abandoned *bool `json:"abandoned,omitempty"`

	// Output only. Time at which the `Release` was created.
	// +kcc:proto:field=google.cloud.deploy.v1.Release.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time at which the render began.
	// +kcc:proto:field=google.cloud.deploy.v1.Release.render_start_time
	RenderStartTime *string `json:"renderStartTime,omitempty"`

	// Output only. Time at which the render completed.
	// +kcc:proto:field=google.cloud.deploy.v1.Release.render_end_time
	RenderEndTime *string `json:"renderEndTime,omitempty"`

	// Output only. Snapshot of the parent pipeline taken at release creation
	//  time.
	// +kcc:proto:field=google.cloud.deploy.v1.Release.delivery_pipeline_snapshot
	DeliveryPipelineSnapshot *DeliveryPipelineObservedState `json:"deliveryPipelineSnapshot,omitempty"`

	// Output only. Snapshot of the targets taken at release creation time.
	// +kcc:proto:field=google.cloud.deploy.v1.Release.target_snapshots
	TargetSnapshots []TargetSnapshotObservedState `json:"targetSnapshots,omitempty"`

	// Output only. Snapshot of the custom target types referenced by the targets
	//  taken at release creation time.
	// +kcc:proto:field=google.cloud.deploy.v1.Release.custom_target_type_snapshots
	CustomTargetTypeSnapshots []CustomTargetTypeObservedState `json:"customTargetTypeSnapshots,omitempty"`

	// Output only. Current state of the render operation.
	// +kcc:proto:field=google.cloud.deploy.v1.Release.render_state
	RenderState *string `json:"renderState,omitempty"`

	// Output only. Map from target ID to the target artifacts created
	//  during the render operation.
	// +kcc:proto:field=google.cloud.deploy.v1.Release.target_artifacts
	TargetArtifacts map[string]TargetArtifactObservedState `json:"targetArtifacts,omitempty"`

	// Output only. Map from target ID to details of the render operation for that
	//  target.
	// +kcc:proto:field=google.cloud.deploy.v1.Release.target_renders
	TargetRenders map[string]TargetRenderObservedState `json:"targetRenders,omitempty"`

	// Output only. Information around the state of the Release.
	// +kcc:proto:field=google.cloud.deploy.v1.Release.condition
	Condition *Release_ReleaseCondition `json:"condition,omitempty"`

	// This checksum is computed by the server based on the value of other
	//  fields, and may be sent on update and delete requests to ensure the
	//  client has an up-to-date value before proceeding.
	// +kcc:proto:field=google.cloud.deploy.v1.Release.etag
	Etag *string `json:"etag,omitempty"`
}

// TargetArtifactObservedState represents the output-only artifacts produced for a target during render.
// +kcc:observedstate:proto=google.cloud.deploy.v1.TargetArtifact
type TargetArtifactObservedState struct {
	// Output only. File path of the resolved Skaffold configuration relative to the URI.
	// +kcc:proto:field=google.cloud.deploy.v1.TargetArtifact.skaffold_config_path
	SkaffoldConfigPath *string `json:"skaffoldConfigPath,omitempty"`

	// Output only. File path of the rendered manifest relative to the URI.
	// +kcc:proto:field=google.cloud.deploy.v1.TargetArtifact.manifest_path
	ManifestPath *string `json:"manifestPath,omitempty"`

	// Output only. Map from the phase ID to the phase artifacts for the phase.
	// +kcc:proto:field=google.cloud.deploy.v1.TargetArtifact.phase_artifacts
	PhaseArtifacts map[string]TargetArtifact_PhaseArtifactObservedState `json:"phaseArtifacts,omitempty"`

	// Output only. URI of a directory containing the artifacts. All paths are relative to this location.
	// +kcc:proto:field=google.cloud.deploy.v1.TargetArtifact.artifact_uri
	ArtifactURI *string `json:"artifactURI,omitempty"`
}

// TargetArtifact_PhaseArtifactObservedState contains paths to artifacts for a phase.
// +kcc:observedstate:proto=google.cloud.deploy.v1.TargetArtifact.PhaseArtifact
type TargetArtifact_PhaseArtifactObservedState struct {
	// Output only. File path of the resolved Skaffold configuration relative to the URI.
	// +kcc:proto:field=google.cloud.deploy.v1.TargetArtifact.PhaseArtifact.skaffold_config_path
	SkaffoldConfigPath *string `json:"skaffoldConfigPath,omitempty"`

	// Output only. File path of the rendered manifest relative to the URI.
	// +kcc:proto:field=google.cloud.deploy.v1.TargetArtifact.PhaseArtifact.manifest_path
	ManifestPath *string `json:"manifestPath,omitempty"`

	// Output only. File path of the resolved job manifests relative to the URI.
	// +kcc:proto:field=google.cloud.deploy.v1.TargetArtifact.PhaseArtifact.job_manifests_path
	JobManifestsPath *string `json:"jobManifestsPath,omitempty"`
}

// TargetRenderObservedState contains details of rendering for a single target.
// +kcc:observedstate:proto=google.cloud.deploy.v1.Release.TargetRender
type TargetRenderObservedState struct {
	// Output only. The resource name of the Cloud Build `Build` object that is
	// used to render the manifest for this target. Format is
	// `projects/{project}/locations/{location}/builds/{build}`.
	// +kcc:proto:field=google.cloud.deploy.v1.Release.TargetRender.rendering_build
	RenderingBuild *string `json:"renderingBuild,omitempty"`

	// Output only. Current state of the render operation for this Target.
	// +kcc:proto:field=google.cloud.deploy.v1.Release.TargetRender.rendering_state
	RenderingState *string `json:"renderingState,omitempty"`

	// Output only. Metadata related to the `Release` render for this Target.
	// +kcc:proto:field=google.cloud.deploy.v1.Release.TargetRender.metadata
	Metadata *RenderMetadataObservedState `json:"metadata,omitempty"`

	// Output only. Reason this render failed.
	// +kcc:proto:field=google.cloud.deploy.v1.Release.TargetRender.failure_cause
	FailureCause *string `json:"failureCause,omitempty"`

	// Output only. Additional information about the render failure, if available.
	// +kcc:proto:field=google.cloud.deploy.v1.Release.TargetRender.failure_message
	FailureMessage *string `json:"failureMessage,omitempty"`
}

// RenderMetadataObservedState includes information associated with a `Release` render.
// +kcc:observedstate:proto=google.cloud.deploy.v1.RenderMetadata
type RenderMetadataObservedState struct {
	// Output only. Metadata associated with rendering for Cloud Run.
	// +kcc:proto:field=google.cloud.deploy.v1.RenderMetadata.cloud_run
	CloudRun *CloudRunRenderMetadataObservedState `json:"cloudRun,omitempty"`

	// Output only. Custom metadata provided by user-defined render operation.
	// +kcc:proto:field=google.cloud.deploy.v1.RenderMetadata.custom
	Custom *CustomMetadataObservedState `json:"custom,omitempty"`
}

// CloudRunRenderMetadataObservedState contains Cloud Run information associated with a `Release` render.
// +kcc:observedstate:proto=google.cloud.deploy.v1.CloudRunRenderMetadata
type CloudRunRenderMetadataObservedState struct {
	// Output only. The name of the Cloud Run Service in the rendered manifest.
	// Format is `projects/{project}/locations/{location}/services/{service}`.
	// +kcc:proto:field=google.cloud.deploy.v1.CloudRunRenderMetadata.service
	Service *string `json:"service,omitempty"`
}

// CustomMetadataObservedState contains information from a user-defined operation.
// +kcc:observedstate:proto=google.cloud.deploy.v1.CustomMetadata
type CustomMetadataObservedState struct {
	// Output only. Key-value pairs provided by the user-defined operation.
	// +kcc:proto:field=google.cloud.deploy.v1.CustomMetadata.values
	Values map[string]string `json:"values,omitempty"`
}

// TargetSnapshotObservedState models the target snapshot in Release status using string identifiers
// rather than KRM references, matching GCP API responses.
// +kcc:observedstate:proto=google.cloud.deploy.v1.Target
type TargetSnapshotObservedState struct {
	// Identifier. Name of the `Target`. Format is
	// `projects/{project}/locations/{location}/targets/{target}`.
	// +kcc:proto:field=google.cloud.deploy.v1.Target.name
	Name *string `json:"name,omitempty"`

	// Output only. Resource id of the `Target`.
	// +kcc:proto:field=google.cloud.deploy.v1.Target.target_id
	TargetID *string `json:"targetID,omitempty"`

	// Output only. Unique identifier of the `Target`.
	// +kcc:proto:field=google.cloud.deploy.v1.Target.uid
	Uid *string `json:"uid,omitempty"`

	// Optional. Description of the `Target`. Max length is 255 characters.
	// +kcc:proto:field=google.cloud.deploy.v1.Target.description
	Description *string `json:"description,omitempty"`

	// Optional. User annotations.
	// +kcc:proto:field=google.cloud.deploy.v1.Target.annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// Optional. Labels are attributes that can be set and used by both the user and by Cloud Deploy.
	// +kcc:proto:field=google.cloud.deploy.v1.Target.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Whether or not the `Target` requires approval.
	// +kcc:proto:field=google.cloud.deploy.v1.Target.require_approval
	RequireApproval *bool `json:"requireApproval,omitempty"`

	// Output only. Time at which the `Target` was created.
	// +kcc:proto:field=google.cloud.deploy.v1.Target.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Most recent time at which the `Target` was updated.
	// +kcc:proto:field=google.cloud.deploy.v1.Target.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Optional. Information specifying a GKE Cluster.
	// +kcc:proto:field=google.cloud.deploy.v1.Target.gke
	GKE *GkeClusterObservedState `json:"gke,omitempty"`

	// Optional. Information specifying an Anthos Cluster.
	// +kcc:proto:field=google.cloud.deploy.v1.Target.anthos_cluster
	AnthosCluster *AnthosClusterObservedState `json:"anthosCluster,omitempty"`

	// Optional. Information specifying a Cloud Run deployment target.
	// +kcc:proto:field=google.cloud.deploy.v1.Target.run
	Run *CloudRunLocationObservedState `json:"run,omitempty"`

	// Optional. Information specifying a multiTarget.
	// +kcc:proto:field=google.cloud.deploy.v1.Target.multi_target
	MultiTarget *MultiTargetObservedState `json:"multiTarget,omitempty"`

	// Optional. Information specifying a Custom Target.
	// +kcc:proto:field=google.cloud.deploy.v1.Target.custom_target
	CustomTarget *CustomTargetObservedState `json:"customTarget,omitempty"`

	// Optional. Map of entity IDs to their associated entities.
	// +kcc:proto:field=google.cloud.deploy.v1.Target.associated_entities
	AssociatedEntities map[string]AssociatedEntitiesObservedState `json:"associatedEntities,omitempty"`

	// Optional. This checksum is computed by the server based on the value of other fields.
	// +kcc:proto:field=google.cloud.deploy.v1.Target.etag
	Etag *string `json:"etag,omitempty"`

	// Optional. Configurations for all execution that relates to this `Target`.
	// +kcc:proto:field=google.cloud.deploy.v1.Target.execution_configs
	ExecutionConfigs []ExecutionConfigObservedState `json:"executionConfigs,omitempty"`

	// Optional. The deploy parameters to use for this target.
	// +kcc:proto:field=google.cloud.deploy.v1.Target.deploy_parameters
	DeployParameters map[string]string `json:"deployParameters,omitempty"`
}

// GkeClusterObservedState models a GKE cluster snapshot in status with a plain cluster string.
// +kcc:observedstate:proto=google.cloud.deploy.v1.GkeCluster
type GkeClusterObservedState struct {
	// Optional. Information specifying a GKE Cluster. Format is `projects/{project_id}/locations/{location_id}/clusters/{cluster_id}`.
	// +kcc:proto:field=google.cloud.deploy.v1.GkeCluster.cluster
	Cluster *string `json:"cluster,omitempty"`

	// Optional. If true, cluster is accessed using private IP.
	// +kcc:proto:field=google.cloud.deploy.v1.GkeCluster.internal_ip
	InternalIP *bool `json:"internalIP,omitempty"`

	// Optional. If set, used to configure a proxy to the Kubernetes server.
	// +kcc:proto:field=google.cloud.deploy.v1.GkeCluster.proxy_url
	ProxyURL *string `json:"proxyURL,omitempty"`

	// Optional. If set, the cluster will be accessed using the DNS endpoint.
	// +kcc:proto:field=google.cloud.deploy.v1.GkeCluster.dns_endpoint
	DNSEndpoint *bool `json:"dnsEndpoint,omitempty"`
}

// AnthosClusterObservedState models an Anthos cluster snapshot in status with a plain membership string.
// +kcc:observedstate:proto=google.cloud.deploy.v1.AnthosCluster
type AnthosClusterObservedState struct {
	// Optional. Membership of the GKE Hub-registered cluster.
	// Format is `projects/{project}/locations/{location}/memberships/{membership_name}`.
	// +kcc:proto:field=google.cloud.deploy.v1.AnthosCluster.membership
	Membership *string `json:"membership,omitempty"`
}

// MultiTargetObservedState models a multiTarget snapshot in status with plain target ID strings.
// +kcc:observedstate:proto=google.cloud.deploy.v1.MultiTarget
type MultiTargetObservedState struct {
	// Required. The target_ids of this multiTarget.
	// +kcc:proto:field=google.cloud.deploy.v1.MultiTarget.target_ids
	TargetIDs []string `json:"targetIDs,omitempty"`
}

// CustomTargetObservedState models a CustomTarget snapshot in status with a plain CustomTargetType string.
// +kcc:observedstate:proto=google.cloud.deploy.v1.CustomTarget
type CustomTargetObservedState struct {
	// Required. The name of the CustomTargetType. Format must be
	// `projects/{project}/locations/{location}/customTargetTypes/{custom_target_type}`.
	// +kcc:proto:field=google.cloud.deploy.v1.CustomTarget.custom_target_type
	CustomTargetType *string `json:"customTargetType,omitempty"`
}

// AssociatedEntitiesObservedState models associated entities in status.
// +kcc:observedstate:proto=google.cloud.deploy.v1.AssociatedEntities
type AssociatedEntitiesObservedState struct {
	// Optional. Information specifying GKE clusters as associated entities.
	// +kcc:proto:field=google.cloud.deploy.v1.AssociatedEntities.gke_clusters
	GKEClusters []GkeClusterObservedState `json:"gkeClusters,omitempty"`

	// Optional. Information specifying Anthos clusters as associated entities.
	// +kcc:proto:field=google.cloud.deploy.v1.AssociatedEntities.anthos_clusters
	AnthosClusters []AnthosClusterObservedState `json:"anthosClusters,omitempty"`
}

// DefaultPoolObservedState models default execution pool in status with a plain serviceAccount string.
// +kcc:observedstate:proto=google.cloud.deploy.v1.DefaultPool
type DefaultPoolObservedState struct {
	// Optional. Google service account to use for execution.
	// +kcc:proto:field=google.cloud.deploy.v1.DefaultPool.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// Optional. Cloud Storage location where execution outputs should be stored.
	// +kcc:proto:field=google.cloud.deploy.v1.DefaultPool.artifact_storage
	ArtifactStorage *string `json:"artifactStorage,omitempty"`
}

// PrivatePoolObservedState models private execution pool in status with plain workerPool and serviceAccount strings.
// +kcc:observedstate:proto=google.cloud.deploy.v1.PrivatePool
type PrivatePoolObservedState struct {
	// Required. Resource name of the Cloud Build worker pool to use.
	// +kcc:proto:field=google.cloud.deploy.v1.PrivatePool.worker_pool
	WorkerPool *string `json:"workerPool,omitempty"`

	// Optional. Google service account to use for execution.
	// +kcc:proto:field=google.cloud.deploy.v1.PrivatePool.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// Optional. Cloud Storage location where execution outputs should be stored.
	// +kcc:proto:field=google.cloud.deploy.v1.PrivatePool.artifact_storage
	ArtifactStorage *string `json:"artifactStorage,omitempty"`
}

// ExecutionConfigObservedState models execution configuration in status with plain string references.
// +kcc:observedstate:proto=google.cloud.deploy.v1.ExecutionConfig
type ExecutionConfigObservedState struct {
	// Required. Usages when this configuration should be applied.
	// +kcc:proto:field=google.cloud.deploy.v1.ExecutionConfig.usages
	Usages []string `json:"usages,omitempty"`

	// Optional. Use default Cloud Build pool.
	// +kcc:proto:field=google.cloud.deploy.v1.ExecutionConfig.default_pool
	DefaultPool *DefaultPoolObservedState `json:"defaultPool,omitempty"`

	// Optional. Use private Cloud Build pool.
	// +kcc:proto:field=google.cloud.deploy.v1.ExecutionConfig.private_pool
	PrivatePool *PrivatePoolObservedState `json:"privatePool,omitempty"`

	// Optional. The resource name of the `WorkerPool`.
	// +kcc:proto:field=google.cloud.deploy.v1.ExecutionConfig.worker_pool
	WorkerPool *string `json:"workerPool,omitempty"`

	// Optional. Google service account to use for execution.
	// +kcc:proto:field=google.cloud.deploy.v1.ExecutionConfig.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// Optional. Cloud Storage location in which to store execution outputs.
	// +kcc:proto:field=google.cloud.deploy.v1.ExecutionConfig.artifact_storage
	ArtifactStorage *string `json:"artifactStorage,omitempty"`

	// Optional. Execution timeout for a Cloud Build Execution.
	// +kcc:proto:field=google.cloud.deploy.v1.ExecutionConfig.execution_timeout
	ExecutionTimeout *string `json:"executionTimeout,omitempty"`

	// Optional. If true, additional logging will be enabled.
	// +kcc:proto:field=google.cloud.deploy.v1.ExecutionConfig.verbose
	Verbose *bool `json:"verbose,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpclouddeployrelease;gcpclouddeployreleases
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CloudDeployRelease is the Schema for the CloudDeployRelease API
// +k8s:openapi-gen=true
type CloudDeployRelease struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CloudDeployReleaseSpec   `json:"spec,omitempty"`
	Status CloudDeployReleaseStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CloudDeployReleaseList contains a list of CloudDeployRelease
type CloudDeployReleaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudDeployRelease `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudDeployRelease{}, &CloudDeployReleaseList{})
}
