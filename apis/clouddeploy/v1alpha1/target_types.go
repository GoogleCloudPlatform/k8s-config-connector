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

package v1alpha1

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var TargetGVK = GroupVersion.WithKind("CloudDeployTarget")

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpclouddeploytarget;gcpclouddeploytargets
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DeployTarget is the Schema for the DeployTarget API
// +k8s:openapi-gen=true
type DeployTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DeployTargetSpec `json:",inline"`
	Status TargetStatus     `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DeployTargetList contains a list of DeployTarget
type DeployTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DeployTarget `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DeployTarget{}, &DeployTargetList{})
}

// DeployTargetSpec defines the desired state of DeployTarget
// +kcc:proto=google.cloud.deploy.v1.Target
type DeployTargetSpec struct {
	// +required
	ProjectRef *refs.ProjectRef `json:"projectRef"`

	//ProjectId:"!{{project}}"

	// The DeployTarget name. If not given, the metadata.name will be used.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// Immutable. The name of the location this data CloudDeployTarget.
	Location string `json:"location"`

	// Optional. Description of the `Target`. Max length is 255 characters.
	// +kcc:proto:field=google.cloud.deploy.v1.Target.description
	Description *string `json:"description,omitempty"`

	// NOT YET
	// // Optional. User annotations. These attributes can only be set and used by
	// //  the user, and not by Cloud Deploy. See
	// //  https://google.aip.dev/128#annotations for more details such as format and
	// //  size limitations.
	// // +kcc:proto:field=google.cloud.deploy.v1.Target.annotations
	// Annotations map[string]string `json:"annotations,omitempty"`

	// // Optional. Labels are attributes that can be set and used by both the
	// //  user and by Cloud Deploy. Labels must meet the following constraints:
	// //
	// //  * Keys and values can contain only lowercase letters, numeric characters,
	// //  underscores, and dashes.
	// //  * All characters must use UTF-8 encoding, and international characters are
	// //  allowed.
	// //  * Keys must start with a lowercase letter or international character.
	// //  * Each resource is limited to a maximum of 64 labels.
	// //
	// //  Both keys and values are additionally constrained to be <= 128 bytes.
	// // +kcc:proto:field=google.cloud.deploy.v1.Target.labels
	// Labels map[string]string `json:"labels,omitempty"`

	// Optional. Whether or not the `Target` requires approval.
	// +kcc:proto:field=google.cloud.deploy.v1.Target.require_approval
	RequireApproval *bool `json:"requireApproval,omitempty"`

	// Optional. Information specifying a GKE Cluster.
	// +kcc:proto:field=google.cloud.deploy.v1.Target.gke
	Gke *GkeCluster `json:"gke,omitempty"`

	// Optional. Information specifying an Anthos Cluster.
	// +kcc:proto:field=google.cloud.deploy.v1.Target.anthos_cluster
	AnthosCluster *AnthosCluster `json:"anthosCluster,omitempty"`

	// Optional. Information specifying a Cloud Run deployment target.
	// +kcc:proto:field=google.cloud.deploy.v1.Target.run
	Run *CloudRunLocation `json:"run,omitempty"`

	// Optional. Information specifying a multiTarget.
	// +kcc:proto:field=google.cloud.deploy.v1.Target.multi_target
	MultiTarget *MultiTarget `json:"multiTarget,omitempty"`

	// Optional. Information specifying a Custom Target.
	// +kcc:proto:field=google.cloud.deploy.v1.Target.custom_target
	CustomTarget *CustomTarget `json:"customTarget,omitempty"`

	// Optional. This checksum is computed by the server based on the value of
	//  other fields, and may be sent on update and delete requests to ensure the
	//  client has an up-to-date value before proceeding.
	// +kcc:proto:field=google.cloud.deploy.v1.Target.etag
	Etag *string `json:"etag,omitempty"`

	// Configurations for all execution that relates to this `Target`.
	//  Each `ExecutionEnvironmentUsage` value may only be used in a single
	//  configuration; using the same value multiple times is an error.
	//  When one or more configurations are specified, they must include the
	//  `RENDER` and `DEPLOY` `ExecutionEnvironmentUsage` values.
	//  When no configurations are specified, execution will use the default
	//  specified in `DefaultPool`.
	// +kcc:proto:field=google.cloud.deploy.v1.Target.execution_configs
	ExecutionConfigs []ExecutionConfig `json:"executionConfigs,omitempty"`

	// Optional. The deploy parameters to use for this target.
	// +kcc:proto:field=google.cloud.deploy.v1.Target.deploy_parameters
	DeployParameters map[string]string `json:"deployParameters,omitempty"`

	// TODO ACPANA: missing union field of type:
	// unsupported map type with key string and value message
	// needs controllerbuilder changes or manual add.
}

// +k8s:openapi-gen=true
type TargetStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DeployTarget resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *DeployTargetObservedState `json:"observedState,omitempty"`
}

// DeployTargetObservedState is the state of the DeployTarget resource as most recently observed in GCP.
// +kcc:proto=google.cloud.deploy.v1.Target
type DeployTargetObservedState struct {
	// Output only. Resource id of the `Target`.
	// +kcc:proto:field=google.cloud.deploy.v1.Target.target_id
	TargetID *string `json:"targetID,omitempty"`

	// Output only. Unique identifier of the `Target`.
	// +kcc:proto:field=google.cloud.deploy.v1.Target.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. Time at which the `Target` was created.
	// +kcc:proto:field=google.cloud.deploy.v1.Target.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Most recent time at which the `Target` was updated.
	// +kcc:proto:field=google.cloud.deploy.v1.Target.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.MultiTarget
type MultiTarget struct {
	// Required. The target_ids of this multiTarget.
	// +kcc:proto:field=google.cloud.deploy.v1.MultiTarget.target_ids
	TargetIds []string `json:"targetIDs,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.CustomTarget
type CustomTarget struct {
	// Required. The name of the CustomTargetTypeRef. Format must be
	//  `projects/{project}/locations/{location}/customTargetTypes/{custom_target_type}`.
	// +kcc:proto:field=google.cloud.deploy.v1.CustomTarget.custom_target_type
	CustomTargetTypeRef *CustomTargetTypeRef `json:"customTargetTypeRef,omitempty"`
}

// +kcc:proto=google.cloud.deploy.v1.PrivatePool
type PrivatePool struct {
	// Required. Resource name of the Cloud Build worker pool to use. The format
	//  is `projects/{project}/locations/{location}/workerPools/{pool}`.
	// +kcc:proto:field=google.cloud.deploy.v1.PrivatePool.worker_pool
	WorkerPool *string `json:"workerPool,omitempty"`

	// Optional. Google service account to use for execution. If unspecified,
	//  the project execution service account
	//  (<PROJECT_NUMBER>-compute@developer.gserviceaccount.com) will be used.
	// +kcc:proto:field=google.cloud.deploy.v1.PrivatePool.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// Optional. Cloud Storage location where execution outputs should be stored.
	//  This can either be a bucket ("gs://my-bucket") or a path within a bucket
	//  ("gs://my-bucket/my-dir").
	//  If unspecified, a default bucket located in the same region will be used.
	// +kcc:proto:field=google.cloud.deploy.v1.PrivatePool.artifact_storage
	ArtifactStorage *string `json:"artifactStorage,omitempty"`
}
