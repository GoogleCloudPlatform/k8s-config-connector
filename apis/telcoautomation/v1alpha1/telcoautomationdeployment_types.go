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

// API sources for TelcoAutomationDeployment, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/telcoautomation/v1/telcoautomation.proto
// +kcc:source:service-docs=https://cloud.google.com/telecom-network-automation
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/telecom-network-automation/docs/reference/rest/v1/projects.locations.orchestrationClusters.deployments

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var TelcoAutomationDeploymentGVK = GroupVersion.WithKind("TelcoAutomationDeployment")

// TelcoAutomationDeploymentSpec defines the desired state of TelcoAutomationDeployment
// +kcc:spec:proto=google.cloud.telcoautomation.v1.Deployment
// +kcc:required-from-proto
type TelcoAutomationDeploymentSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The TelcoAutomationDeployment name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. The blueprint revision from which this deployment was created.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.Deployment.source_blueprint_revision
	// +required
	SourceBlueprintRevision *string `json:"sourceBlueprintRevision,omitempty"`

	// Optional. Human readable name of a Deployment.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.Deployment.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Files present in a deployment.
	//  When invoking UpdateDeployment API, only the modified files should be
	//  included in this. Files that are not included in the update of a deployment
	//  will not be changed.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.Deployment.files
	Files []File `json:"files,omitempty"`

	// Optional. Labels are key-value attributes that can be set on a deployment
	//  resource by the user.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.Deployment.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Immutable. The WorkloadCluster on which to create the Deployment.
	//  This field should only be passed when the deployment_level of the source
	//  blueprint specifies deployments on workload clusters e.g.
	//  WORKLOAD_CLUSTER_DEPLOYMENT.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.Deployment.workload_cluster
	WorkloadCluster *string `json:"workloadCluster,omitempty"`
}

// TelcoAutomationDeploymentStatus defines the config connector machine state of TelcoAutomationDeployment
type TelcoAutomationDeploymentStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the TelcoAutomationDeployment resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *TelcoAutomationDeploymentObservedState `json:"observedState,omitempty"`
}

// TelcoAutomationDeploymentObservedState is the state of the TelcoAutomationDeployment resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.telcoautomation.v1.Deployment
type TelcoAutomationDeploymentObservedState struct {
	// Output only. Immutable. The revision ID of the deployment.
	//  A new revision is committed whenever a change in deployment is applied.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.Deployment.revision_id
	RevisionID *string `json:"revisionID,omitempty"`

	// Output only. The timestamp that the revision was created.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.Deployment.revision_create_time
	RevisionCreateTime *string `json:"revisionCreateTime,omitempty"`

	// Output only. State of the deployment (DRAFT, APPLIED, DELETING).
	// +kcc:proto:field=google.cloud.telcoautomation.v1.Deployment.state
	State *string `json:"state,omitempty"`

	// Output only. Name of the repository where the deployment package files are
	//  stored.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.Deployment.repository
	Repository *string `json:"repository,omitempty"`

	// Output only. Deployment creation time.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.Deployment.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the deployment was updated.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.Deployment.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Source provider is the author of a public blueprint, from
	//  which this deployment is created.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.Deployment.source_provider
	SourceProvider *string `json:"sourceProvider,omitempty"`

	// Output only. Attributes to where the deployment can inflict changes. The
	//  value can only be [SINGLE_DEPLOYMENT, MULTI_DEPLOYMENT].
	// +kcc:proto:field=google.cloud.telcoautomation.v1.Deployment.deployment_level
	DeploymentLevel *string `json:"deploymentLevel,omitempty"`

	// Output only. Indicates if the deployment can be rolled back, exported from
	//  public blueprint.
	// +kcc:proto:field=google.cloud.telcoautomation.v1.Deployment.rollback_support
	RollbackSupport *bool `json:"rollbackSupport,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcptelcoautomationdeployment;gcptelcoautomationdeployments
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// TelcoAutomationDeployment is the Schema for the TelcoAutomationDeployment API
// +k8s:openapi-gen=true
type TelcoAutomationDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   TelcoAutomationDeploymentSpec   `json:"spec,omitempty"`
	Status TelcoAutomationDeploymentStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TelcoAutomationDeploymentList contains a list of TelcoAutomationDeployment
type TelcoAutomationDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TelcoAutomationDeployment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TelcoAutomationDeployment{}, &TelcoAutomationDeploymentList{})
}
