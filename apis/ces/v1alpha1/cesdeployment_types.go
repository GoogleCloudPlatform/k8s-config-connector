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

// API sources for CESDeployment, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/ces/v1beta/deployment.proto
// +kcc:source:service-docs=https://docs.cloud.google.com/customer-engagement-ai/conversational-agents/ps
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/customer-engagement-ai/conversational-agents/ps/docs/reference/rest/v1beta/projects.locations.apps.deployments

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CESDeploymentGVK = GroupVersion.WithKind("CESDeployment")

// CESDeploymentSpec defines the desired state of CESDeployment
// +kcc:spec:proto=google.cloud.ces.v1beta.Deployment
// +kcc:required-from-proto
type CESDeploymentSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The CESDeployment name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. Display name of the deployment.
	// +kcc:proto:field=google.cloud.ces.v1beta.Deployment.display_name
	// +required
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. The resource name of the app version to deploy.
	//  Format:
	//  `projects/{project}/locations/{location}/apps/{app}/versions/{version}`
	//  Use `projects/{project}/locations/{location}/apps/{app}/versions/-` to use
	//  the draft app.
	// +kcc:proto:field=google.cloud.ces.v1beta.Deployment.app_version
	AppVersion *string `json:"appVersion,omitempty"`

	// Required. The channel profile used in the deployment.
	// +kcc:proto:field=google.cloud.ces.v1beta.Deployment.channel_profile
	// +required
	ChannelProfile *ChannelProfile `json:"channelProfile,omitempty"`

	// Optional. Experiment configuration for the deployment.
	// +kcc:proto:field=google.cloud.ces.v1beta.Deployment.experiment_config
	ExperimentConfig *ExperimentConfig `json:"experimentConfig,omitempty"`
}

// CESDeploymentStatus defines the config connector machine state of CESDeployment
type CESDeploymentStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the CESDeployment resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *CESDeploymentObservedState `json:"observedState,omitempty"`
}

// CESDeploymentObservedState is the state of the CESDeployment resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.ces.v1beta.Deployment
type CESDeploymentObservedState struct {
	// Output only. Timestamp when this deployment was created.
	// +kcc:proto:field=google.cloud.ces.v1beta.Deployment.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this deployment was last updated.
	// +kcc:proto:field=google.cloud.ces.v1beta.Deployment.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Etag used to ensure the object hasn't changed during a
	//  read-modify-write operation. If the etag is empty, the update will
	//  overwrite any concurrent changes.
	// +kcc:proto:field=google.cloud.ces.v1beta.Deployment.etag
	Etag *string `json:"etag,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcesdeployment;gcpcesdeployments
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CESDeployment is the Schema for the CESDeployment API
// +k8s:openapi-gen=true
type CESDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CESDeploymentSpec   `json:"spec,omitempty"`
	Status CESDeploymentStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CESDeploymentList contains a list of CESDeployment
type CESDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CESDeployment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CESDeployment{}, &CESDeploymentList{})
}
