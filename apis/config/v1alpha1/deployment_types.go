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

package v1alpha1

import (
	cloudbuildv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/cloudbuild/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var InfraManagerDeploymentGVK = GroupVersion.WithKind("InfraManagerDeployment")

// InfraManagerDeploymentSpec defines the desired state of InfraManagerDeployment
// +kcc:spec:proto=google.cloud.config.v1.Deployment
type InfraManagerDeploymentSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The InfraManagerDeployment name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// A blueprint described using Terraform's HashiCorp Configuration Language
	//  as a root module.
	// +kcc:proto:field=google.cloud.config.v1.Deployment.terraform_blueprint
	TerraformBlueprint *TerraformBlueprint `json:"terraformBlueprint,omitempty"`

	// Optional. User-defined metadata for the deployment.
	// +kcc:proto:field=google.cloud.config.v1.Deployment.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. User-defined location of Cloud Build logs and artifacts in Google
	//  Cloud Storage. Format: `gs://{bucket}/{folder}`
	//
	//  A default bucket will be bootstrapped if the field is not set or empty.
	//  Default bucket format: `gs://<project number>-<region>-blueprint-config`
	//  Constraints:
	//  - The bucket needs to be in the same project as the deployment
	//  - The path cannot be within the path of `gcs_source`
	//  - The field cannot be updated, including changing its presence
	// +kcc:proto:field=google.cloud.config.v1.Deployment.artifacts_gcs_bucket
	ArtifactsGCSBucket *string `json:"artifactsGCSBucket,omitempty"`

	// Required. User-specified Service Account (SA) credentials to be used when
	//  actuating resources.
	//  Format: `projects/{projectID}/serviceAccounts/{serviceAccount}`
	// +kcc:proto:field=google.cloud.config.v1.Deployment.service_account
	ServiceAccountRef *refsv1beta1.IAMServiceAccountRef `json:"serviceAccountRef"`

	// By default, Infra Manager will return a failure when
	//  Terraform encounters a 409 code (resource conflict error) during actuation.
	//  If this flag is set to true, Infra Manager will instead
	//  attempt to automatically import the resource into the Terraform state (for
	//  supported resource types) and continue actuation.
	//
	//  Not all resource types are supported, refer to documentation.
	// +kcc:proto:field=google.cloud.config.v1.Deployment.import_existing_resources
	ImportExistingResources *bool `json:"importExistingResources,omitempty"`

	// Optional. The user-specified Cloud Build worker pool resource in which the
	//  Cloud Build job will execute. Format:
	//  `projects/{project}/locations/{location}/workerPools/{workerPoolId}`.
	//  If this field is unspecified, the default Cloud Build worker pool will be
	//  used.
	// +kcc:proto:field=google.cloud.config.v1.Deployment.worker_pool
	WorkerPoolRef *cloudbuildv1beta1.CloudBuildWorkerPoolRef `json:"workerPoolRef,omitempty"`

	// Optional. The user-specified Terraform version constraint.
	//  Example: "=1.3.10".
	// +kcc:proto:field=google.cloud.config.v1.Deployment.tf_version_constraint
	TfVersionConstraint *string `json:"tfVersionConstraint,omitempty"`

	// Optional. Input to control quota checks for resources in terraform
	//  configuration files. There are limited resources on which quota validation
	//  applies.
	// +kubebuilder:validation:Enum=QUOTA_VALIDATION_UNSPECIFIED;ENABLED;ENFORCED
	// +kcc:proto:field=google.cloud.config.v1.Deployment.quota_validation
	QuotaValidation *string `json:"quotaValidation,omitempty"`

	// Optional. Arbitrary key-value metadata storage e.g. to help client tools
	//  identify deployments during automation. See
	//  https://google.aip.dev/148#annotations for details on format and size
	//  limitations.
	// +kcc:proto:field=google.cloud.config.v1.Deployment.annotations
	Annotations map[string]string `json:"annotations,omitempty"`
}

// InfraManagerDeploymentStatus defines the config connector machine state of InfraManagerDeployment
type InfraManagerDeploymentStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the InfraManagerDeployment resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *InfraManagerDeploymentObservedState `json:"observedState,omitempty"`
}

// InfraManagerDeploymentObservedState is the state of the InfraManagerDeployment resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.config.v1.Deployment
type InfraManagerDeploymentObservedState struct {
	// Output only. Time when the deployment was created.
	// +kcc:proto:field=google.cloud.config.v1.Deployment.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time when the deployment was last modified.
	// +kcc:proto:field=google.cloud.config.v1.Deployment.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Current state of the deployment.
	// +kcc:proto:field=google.cloud.config.v1.Deployment.state
	State *string `json:"state,omitempty"`

	// Output only. Revision name that was most recently applied.
	//  Format: `projects/{project}/locations/{location}/deployments/{deployment}/
	//  revisions/{revision}`
	// +kcc:proto:field=google.cloud.config.v1.Deployment.latest_revision
	LatestRevision *string `json:"latestRevision,omitempty"`

	// Output only. Additional information regarding the current state.
	// +kcc:proto:field=google.cloud.config.v1.Deployment.state_detail
	StateDetail *string `json:"stateDetail,omitempty"`

	// Output only. Error code describing errors that may have occurred.
	// +kcc:proto:field=google.cloud.config.v1.Deployment.error_code
	ErrorCode *string `json:"errorCode,omitempty"`

	// Output only. Location of artifacts from a DeleteDeployment operation.
	// +kcc:proto:field=google.cloud.config.v1.Deployment.delete_results
	DeleteResults *ApplyResults `json:"deleteResults,omitempty"`

	// Output only. Cloud Build instance UUID associated with deleting this
	//  deployment.
	// +kcc:proto:field=google.cloud.config.v1.Deployment.delete_build
	DeleteBuild *string `json:"deleteBuild,omitempty"`

	// Output only. Location of Cloud Build logs in Google Cloud Storage,
	//  populated when deleting this deployment. Format: `gs://{bucket}/{object}`.
	// +kcc:proto:field=google.cloud.config.v1.Deployment.delete_logs
	DeleteLogs *string `json:"deleteLogs,omitempty"`

	// Output only. Errors encountered when deleting this deployment.
	//  Errors are truncated to 10 entries, see `delete_results` and `error_logs`
	//  for full details.
	// +kcc:proto:field=google.cloud.config.v1.Deployment.tf_errors
	TfErrors []TerraformError `json:"tfErrors,omitempty"`

	// Output only. Location of Terraform error logs in Google Cloud Storage.
	//  Format: `gs://{bucket}/{object}`.
	// +kcc:proto:field=google.cloud.config.v1.Deployment.error_logs
	ErrorLogs *string `json:"errorLogs,omitempty"`

	// Output only. Current lock state of the deployment.
	// +kcc:proto:field=google.cloud.config.v1.Deployment.lock_state
	LockState *string `json:"lockState,omitempty"`

	// Output only. The current Terraform version set on the deployment.
	//  It is in the format of "Major.Minor.Patch", for example, "1.3.10".
	// +kcc:proto:field=google.cloud.config.v1.Deployment.tf_version
	TfVersion *string `json:"tfVersion,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpinframanagerdeployment;gcpinframanagerdeployments
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// InfraManagerDeployment is the Schema for the InfraManagerDeployment API
// +k8s:openapi-gen=true
type InfraManagerDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   InfraManagerDeploymentSpec   `json:"spec,omitempty"`
	Status InfraManagerDeploymentStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// InfraManagerDeploymentList contains a list of InfraManagerDeployment
type InfraManagerDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InfraManagerDeployment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&InfraManagerDeployment{}, &InfraManagerDeploymentList{})
}
