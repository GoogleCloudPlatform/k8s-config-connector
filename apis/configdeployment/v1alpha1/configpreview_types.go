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

// API sources for ConfigPreview, recorded by generate-types:
// +kcc:source:proto=https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/master/mockgcp/apis/google/cloud/config/v1/config.proto
// +kcc:source:service-docs=https://cloud.google.com/infrastructure-manager/docs
// +kcc:source:resource-docs=https://docs.cloud.google.com/infrastructure-manager/docs/reference/rest/v1/projects.locations.previews

package v1alpha1

import (
	common "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ConfigPreviewGVK = GroupVersion.WithKind("ConfigPreview")

// ConfigPreviewSpec defines the desired state of ConfigPreview
// +kcc:spec:proto=google.cloud.config.v1.Preview
// +kcc:required-from-proto
type ConfigPreviewSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The ConfigPreview name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// The terraform blueprint to preview.
	// +kcc:proto:field=google.cloud.config.v1.Preview.terraform_blueprint
	TerraformBlueprint *TerraformBlueprint `json:"terraformBlueprint,omitempty"`

	// Optional. User-defined labels for the preview.
	// +kcc:proto:field=google.cloud.config.v1.Preview.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Optional deployment reference. If specified, the preview will be
	//  performed using the provided deployment's current state and use any
	//  relevant fields from the deployment unless explicitly specified in the
	//  preview create request.
	// +kcc:proto:field=google.cloud.config.v1.Preview.deployment
	Deployment *string `json:"deployment,omitempty"`

	// Optional. Current mode of preview.
	// +kcc:proto:field=google.cloud.config.v1.Preview.preview_mode
	PreviewMode *string `json:"previewMode,omitempty"`

	// Required. User-specified Service Account (SA) credentials to be used when
	//  previewing resources.
	//  Format: `projects/{projectID}/serviceAccounts/{serviceAccount}`
	// +kcc:proto:field=google.cloud.config.v1.Preview.service_account
	// +required
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// Optional. User-defined location of Cloud Build logs, artifacts, and
	//  in Google Cloud Storage.
	//  Format: `gs://{bucket}/{folder}`
	//  A default bucket will be bootstrapped if the field is not set or empty
	//  Default Bucket Format: `gs://<project number>-<region>-blueprint-config`
	//  Constraints:
	//  - The bucket needs to be in the same project as the deployment
	//  - The path cannot be within the path of `gcs_source`
	//  If omitted and deployment resource ref provided has artifacts_gcs_bucket
	//  defined, that artifact bucket is used.
	// +kcc:proto:field=google.cloud.config.v1.Preview.artifacts_gcs_bucket
	ArtifactsGCSBucket *string `json:"artifactsGCSBucket,omitempty"`

	// Optional. The user-specified Worker Pool resource in which the Cloud Build
	//  job will execute. Format
	//  projects/{project}/locations/{location}/workerPools/{workerPoolId} If this
	//  field is unspecified, the default Cloud Build worker pool will be used. If
	//  omitted and deployment resource ref provided has worker_pool defined, that
	//  worker pool is used.
	// +kcc:proto:field=google.cloud.config.v1.Preview.worker_pool
	WorkerPool *string `json:"workerPool,omitempty"`

	// Optional. The user-specified Terraform version constraint.
	//  Example: "=1.3.10".
	// +kcc:proto:field=google.cloud.config.v1.Preview.tf_version_constraint
	TfVersionConstraint *string `json:"tfVersionConstraint,omitempty"`

	// Optional. Arbitrary key-value metadata storage e.g. to help client tools
	//  identify preview during automation. See
	//  https://google.aip.dev/148#annotations for details on format and size
	//  limitations.
	// +kcc:proto:field=google.cloud.config.v1.Preview.annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// Optional. This field specifies the provider configurations.
	// +kcc:proto:field=google.cloud.config.v1.Preview.provider_config
	ProviderConfig *ProviderConfig `json:"providerConfig,omitempty"`
}

// ConfigPreviewStatus defines the config connector machine state of ConfigPreview
type ConfigPreviewStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ConfigPreview resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ConfigPreviewObservedState `json:"observedState,omitempty"`
}

// ConfigPreviewObservedState is the state of the ConfigPreview resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.config.v1.Preview
type ConfigPreviewObservedState struct {
	// Output only. Time the preview was created.
	// +kcc:proto:field=google.cloud.config.v1.Preview.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Current state of the preview.
	// +kcc:proto:field=google.cloud.config.v1.Preview.state
	State *string `json:"state,omitempty"`

	// Output only. Code describing any errors that may have occurred.
	// +kcc:proto:field=google.cloud.config.v1.Preview.error_code
	ErrorCode *string `json:"errorCode,omitempty"`

	// Output only. Additional information regarding the current state.
	// +kcc:proto:field=google.cloud.config.v1.Preview.error_status
	ErrorStatus *common.Status `json:"errorStatus,omitempty"`

	// Output only. Cloud Build instance UUID associated with this preview.
	// +kcc:proto:field=google.cloud.config.v1.Preview.build
	Build *string `json:"build,omitempty"`

	// Output only. Summary of errors encountered during Terraform preview.
	//  It has a size limit of 10, i.e. only top 10 errors will be summarized here.
	// +kcc:proto:field=google.cloud.config.v1.Preview.tf_errors
	TfErrors []TerraformErrorObservedState `json:"tfErrors,omitempty"`

	// Output only. Link to tf-error.ndjson file, which contains the full list of
	//  the errors encountered during a Terraform preview.
	//  Format: `gs://{bucket}/{object}`.
	// +kcc:proto:field=google.cloud.config.v1.Preview.error_logs
	ErrorLogs *string `json:"errorLogs,omitempty"`

	// Output only. Artifacts from preview.
	// +kcc:proto:field=google.cloud.config.v1.Preview.preview_artifacts
	PreviewArtifacts *PreviewArtifactsObservedState `json:"previewArtifacts,omitempty"`

	// Output only. Location of preview logs in `gs://{bucket}/{object}` format.
	// +kcc:proto:field=google.cloud.config.v1.Preview.logs
	Logs *string `json:"logs,omitempty"`

	// Output only. The current Terraform version set on the preview.
	//  It is in the format of "Major.Minor.Patch", for example, "1.3.10".
	// +kcc:proto:field=google.cloud.config.v1.Preview.tf_version
	TfVersion *string `json:"tfVersion,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpconfigpreview;gcpconfigpreviews
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ConfigPreview is the Schema for the ConfigPreview API
// +k8s:openapi-gen=true
type ConfigPreview struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ConfigPreviewSpec   `json:"spec,omitempty"`
	Status ConfigPreviewStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ConfigPreviewList contains a list of ConfigPreview
type ConfigPreviewList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConfigPreview `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ConfigPreview{}, &ConfigPreviewList{})
}
