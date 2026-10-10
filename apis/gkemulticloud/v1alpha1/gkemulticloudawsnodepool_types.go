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

// API sources for GKEMulticloudAWSNodePool, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/gkemulticloud/v1/aws_resources.proto
// +kcc:source:service-docs=https://cloud.google.com/kubernetes-engine/multi-cloud/docs
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/kubernetes-engine/docs/reference/rest/v1/projects.locations.awsClusters.awsNodePools

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var GKEMulticloudAWSNodePoolGVK = GroupVersion.WithKind("GKEMulticloudAWSNodePool")

// GKEMulticloudAWSNodePoolSpec defines the desired state of GKEMulticloudAWSNodePool
// +kcc:spec:proto=google.cloud.gkemulticloud.v1.AwsNodePool
// +kcc:required-from-proto
type GKEMulticloudAWSNodePoolSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The GKEMulticloudAWSNodePool name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. The Kubernetes version to run on this node pool (e.g.
	//  `1.19.10-gke.1000`).
	//
	//  You can list all supported versions on a given Google Cloud region by
	//  calling
	//  [GetAwsServerConfig][google.cloud.gkemulticloud.v1.AwsClusters.GetAwsServerConfig].
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsNodePool.version
	// +required
	Version *string `json:"version,omitempty"`

	// Required. The configuration of the node pool.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsNodePool.config
	// +required
	Config *AwsNodeConfig `json:"config,omitempty"`

	// Required. Autoscaler configuration for this node pool.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsNodePool.autoscaling
	// +required
	Autoscaling *AwsNodePoolAutoscaling `json:"autoscaling,omitempty"`

	// Required. The subnet where the node pool node run.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsNodePool.subnet_id
	// +required
	SubnetID *string `json:"subnetID,omitempty"`

	// Allows clients to perform consistent read-modify-writes
	//  through optimistic concurrency control.
	//
	//  Can be sent on update and delete requests to ensure the
	//  client has an up-to-date value before proceeding.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsNodePool.etag
	Etag *string `json:"etag,omitempty"`

	// Optional. Annotations on the node pool.
	//
	//  This field has the same restrictions as Kubernetes annotations.
	//  The total size of all keys and values combined is limited to 256k.
	//  Key can have 2 segments: prefix (optional) and name (required),
	//  separated by a slash (/).
	//  Prefix must be a DNS subdomain.
	//  Name must be 63 characters or less, begin and end with alphanumerics,
	//  with dashes (-), underscores (_), dots (.), and alphanumerics between.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsNodePool.annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// Required. The constraint on the maximum number of pods that can be run
	//  simultaneously on a node in the node pool.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsNodePool.max_pods_constraint
	// +required
	MaxPodsConstraint *MaxPodsConstraint `json:"maxPodsConstraint,omitempty"`

	// Optional. The Management configuration for this node pool.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsNodePool.management
	Management *AwsNodeManagement `json:"management,omitempty"`

	// Optional. Node kubelet configs.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsNodePool.kubelet_config
	KubeletConfig *NodeKubeletConfig `json:"kubeletConfig,omitempty"`

	// Optional. Update settings control the speed and disruption of the update.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsNodePool.update_settings
	UpdateSettings *UpdateSettings `json:"updateSettings,omitempty"`
}

// GKEMulticloudAWSNodePoolStatus defines the config connector machine state of GKEMulticloudAWSNodePool
type GKEMulticloudAWSNodePoolStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the GKEMulticloudAWSNodePool resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *GKEMulticloudAWSNodePoolObservedState `json:"observedState,omitempty"`
}

// GKEMulticloudAWSNodePoolObservedState is the state of the GKEMulticloudAWSNodePool resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.gkemulticloud.v1.AwsNodePool
type GKEMulticloudAWSNodePoolObservedState struct {
	// Output only. The lifecycle state of the node pool.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsNodePool.state
	State *string `json:"state,omitempty"`

	// Output only. A globally unique identifier for the node pool.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsNodePool.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. If set, there are currently changes in flight to the node
	//  pool.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsNodePool.reconciling
	Reconciling *bool `json:"reconciling,omitempty"`

	// Output only. The time at which this node pool was created.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsNodePool.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which this node pool was last updated.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsNodePool.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. A set of errors found in the node pool.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AwsNodePool.errors
	Errors []AwsNodePoolError `json:"errors,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpgkemulticloudawsnodepool;gcpgkemulticloudawsnodepools
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// GKEMulticloudAWSNodePool is the Schema for the GKEMulticloudAWSNodePool API
// +k8s:openapi-gen=true
type GKEMulticloudAWSNodePool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   GKEMulticloudAWSNodePoolSpec   `json:"spec,omitempty"`
	Status GKEMulticloudAWSNodePoolStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// GKEMulticloudAWSNodePoolList contains a list of GKEMulticloudAWSNodePool
type GKEMulticloudAWSNodePoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GKEMulticloudAWSNodePool `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GKEMulticloudAWSNodePool{}, &GKEMulticloudAWSNodePoolList{})
}
