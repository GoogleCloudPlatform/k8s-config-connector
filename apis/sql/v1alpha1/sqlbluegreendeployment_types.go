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

// API sources for SQLBlueGreenDeployment, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/sql/v1/cloud_sql_blue_green_deployments.proto
// +kcc:source:service-docs=https://cloud.google.com/sql/docs
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/sql/docs/reference/rest/v1/projects.locations.blueGreenDeployments

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var SQLBlueGreenDeploymentGVK = GroupVersion.WithKind("SQLBlueGreenDeployment")

// SQLBlueGreenDeploymentSpec defines the desired state of SQLBlueGreenDeployment
// +kcc:spec:proto=google.cloud.sql.v1.BlueGreenDeployment
// +kcc:required-from-proto
type SQLBlueGreenDeploymentSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The SQLBlueGreenDeployment name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. User-provided description for the deployment.
	//  The description can be up to 255 characters long.
	// +kcc:proto:field=google.cloud.sql.v1.BlueGreenDeployment.description
	Description *string `json:"description,omitempty"`

	// Required. Immutable. Required on create, and immutable. The full resource
	//  name of the source instance (the "blue" instance). Format:
	//  projects/{project}/instances/{instance}
	// +kcc:proto:field=google.cloud.sql.v1.BlueGreenDeployment.source_instance
	// +required
	SourceInstance *string `json:"sourceInstance,omitempty"`

	// Optional. Immutable. Optional on create, and immutable. The configuration
	//  intended for the target instance(s) when the deployment was created.
	// +kcc:proto:field=google.cloud.sql.v1.BlueGreenDeployment.requested_config
	RequestedConfig *BlueGreenDeployment_RequestedConfig `json:"requestedConfig,omitempty"`
}

// SQLBlueGreenDeploymentStatus defines the config connector machine state of SQLBlueGreenDeployment
type SQLBlueGreenDeploymentStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the SQLBlueGreenDeployment resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *SQLBlueGreenDeploymentObservedState `json:"observedState,omitempty"`
}

// SQLBlueGreenDeploymentObservedState is the state of the SQLBlueGreenDeployment resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.sql.v1.BlueGreenDeployment
type SQLBlueGreenDeploymentObservedState struct {
	// Output only. The time when the deployment was created.
	//  Example: `2024-01-01T00:00:00Z`
	// +kcc:proto:field=google.cloud.sql.v1.BlueGreenDeployment.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The current state of the blue-green deployment.
	// +kcc:proto:field=google.cloud.sql.v1.BlueGreenDeployment.state
	State *string `json:"state,omitempty"`

	// Output only. The full resource name of the primary target instance (the
	//  "green" instance) that will be promoted during switchover. This field is
	//  always populated once the deployment is created. Format:
	//  projects/{project}/instances/{instance}
	// +kcc:proto:field=google.cloud.sql.v1.BlueGreenDeployment.switchover_target_instance
	SwitchoverTargetInstance *string `json:"switchoverTargetInstance,omitempty"`

	// Output only. Provides details on why switchover is not possible.
	//  This field is empty unless a switchover attempt failed or the state is
	//  `SWITCHOVER_NOT_READY`.
	//  Example: "The target database version does not match the source instance
	//  database version."
	// +kcc:proto:field=google.cloud.sql.v1.BlueGreenDeployment.error_detail
	ErrorDetail *string `json:"errorDetail,omitempty"`

	// Output only. A list representing the pairs of source and target instances
	//  in the deployment.
	// +kcc:proto:field=google.cloud.sql.v1.BlueGreenDeployment.deployment_mappings
	DeploymentMappings []BlueGreenDeployment_SourceTargetPairedNodeObservedState `json:"deploymentMappings,omitempty"`

	// Output only. Combined list of tasks for all paired nodes.
	// +kcc:proto:field=google.cloud.sql.v1.BlueGreenDeployment.deployment_tasks
	DeploymentTasks *BlueGreenDeployment_DeploymentTasksObservedState `json:"deploymentTasks,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpsqlbluegreendeployment;gcpsqlbluegreendeployments
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// SQLBlueGreenDeployment is the Schema for the SQLBlueGreenDeployment API
// +k8s:openapi-gen=true
type SQLBlueGreenDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   SQLBlueGreenDeploymentSpec   `json:"spec,omitempty"`
	Status SQLBlueGreenDeploymentStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// SQLBlueGreenDeploymentList contains a list of SQLBlueGreenDeployment
type SQLBlueGreenDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SQLBlueGreenDeployment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SQLBlueGreenDeployment{}, &SQLBlueGreenDeploymentList{})
}
