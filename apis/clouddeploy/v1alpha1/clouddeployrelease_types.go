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
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/deploy/docs/reference/rest/v1/projects.locations.deliveryPipelines.releases

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CloudDeployReleaseGVK = GroupVersion.WithKind("CloudDeployRelease")

// CloudDeployReleaseSpec defines the desired state of CloudDeployRelease
// +kcc:spec:proto=google.cloud.deploy.v1.Release
// +kcc:required-from-proto
type CloudDeployReleaseSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

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

	// Labels are attributes that can be set and used by both the
	//  user and by Cloud Deploy. Labels must meet the following constraints:
	//
	//  * Keys and values can contain only lowercase letters, numeric characters,
	//  underscores, and dashes.
	//  * All characters must use UTF-8 encoding, and international characters are
	//  allowed.
	//  * Keys must start with a lowercase letter or international character.
	//  * Each resource is limited to a maximum of 64 labels.
	//
	//  Both keys and values are additionally constrained to be <= 128 bytes.
	// +kcc:proto:field=google.cloud.deploy.v1.Release.labels
	Labels map[string]string `json:"labels,omitempty"`

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

	// This checksum is computed by the server based on the value of other
	//  fields, and may be sent on update and delete requests to ensure the
	//  client has an up-to-date value before proceeding.
	// +kcc:proto:field=google.cloud.deploy.v1.Release.etag
	Etag *string `json:"etag,omitempty"`

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
	TargetSnapshots []TargetObservedState `json:"targetSnapshots,omitempty"`

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
	TargetArtifacts map[string]TargetArtifact `json:"targetArtifacts,omitempty"`

	// Output only. Map from target ID to details of the render operation for that
	//  target.
	// +kcc:proto:field=google.cloud.deploy.v1.Release.target_renders
	TargetRenders map[string]Release_TargetRender `json:"targetRenders,omitempty"`

	// Output only. Information around the state of the Release.
	// +kcc:proto:field=google.cloud.deploy.v1.Release.condition
	Condition *Release_ReleaseCondition `json:"condition,omitempty"`
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
