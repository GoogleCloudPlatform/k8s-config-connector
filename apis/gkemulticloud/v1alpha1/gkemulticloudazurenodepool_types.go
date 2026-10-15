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

// API sources for GKEMulticloudAzureNodePool, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/gkemulticloud/v1/azure_resources.proto
// +kcc:source:service-docs=https://cloud.google.com/kubernetes-engine/multi-cloud/docs
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/kubernetes-engine/docs/reference/rest/v1/projects.locations.azureClusters.azureNodePools

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var GKEMulticloudAzureNodePoolGVK = GroupVersion.WithKind("GKEMulticloudAzureNodePool")

// GKEMulticloudAzureNodePoolSpec defines the desired state of GKEMulticloudAzureNodePool
// +kcc:spec:proto=google.cloud.gkemulticloud.v1.AzureNodePool
// +kcc:required-from-proto
type GKEMulticloudAzureNodePoolSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The GKEMulticloudAzureNodePool name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. The Kubernetes version (e.g. `1.19.10-gke.1000`) running on this
	//  node pool.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureNodePool.version
	// +required
	Version *string `json:"version,omitempty"`

	// Required. The node configuration of the node pool.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureNodePool.config
	// +required
	Config *AzureNodeConfig `json:"config,omitempty"`

	// Required. The ARM ID of the subnet where the node pool VMs run. Make sure
	//  it's a subnet under the virtual network in the cluster configuration.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureNodePool.subnet_id
	// +required
	SubnetID *string `json:"subnetID,omitempty"`

	// Required. Autoscaler configuration for this node pool.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureNodePool.autoscaling
	// +required
	Autoscaling *AzureNodePoolAutoscaling `json:"autoscaling,omitempty"`

	// Allows clients to perform consistent read-modify-writes
	//  through optimistic concurrency control.
	//
	//  Can be sent on update and delete requests to ensure the
	//  client has an up-to-date value before proceeding.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureNodePool.etag
	Etag *string `json:"etag,omitempty"`

	// Optional. Annotations on the node pool.
	//
	//  This field has the same restrictions as Kubernetes annotations.
	//  The total size of all keys and values combined is limited to 256k.
	//  Keys can have 2 segments: prefix (optional) and name (required),
	//  separated by a slash (/).
	//  Prefix must be a DNS subdomain.
	//  Name must be 63 characters or less, begin and end with alphanumerics,
	//  with dashes (-), underscores (_), dots (.), and alphanumerics between.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureNodePool.annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// Required. The constraint on the maximum number of pods that can be run
	//  simultaneously on a node in the node pool.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureNodePool.max_pods_constraint
	// +required
	MaxPodsConstraint *MaxPodsConstraint `json:"maxPodsConstraint,omitempty"`

	// Optional. The Azure availability zone of the nodes in this nodepool.
	//
	//  When unspecified, it defaults to `1`.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureNodePool.azure_availability_zone
	AzureAvailabilityZone *string `json:"azureAvailabilityZone,omitempty"`

	// Optional. The Management configuration for this node pool.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureNodePool.management
	Management *AzureNodeManagement `json:"management,omitempty"`
}

// GKEMulticloudAzureNodePoolStatus defines the config connector machine state of GKEMulticloudAzureNodePool
type GKEMulticloudAzureNodePoolStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the GKEMulticloudAzureNodePool resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *GKEMulticloudAzureNodePoolObservedState `json:"observedState,omitempty"`
}

// GKEMulticloudAzureNodePoolObservedState is the state of the GKEMulticloudAzureNodePool resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.gkemulticloud.v1.AzureNodePool
type GKEMulticloudAzureNodePoolObservedState struct {
	// Output only. The current state of the node pool.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureNodePool.state
	State *string `json:"state,omitempty"`

	// Output only. A globally unique identifier for the node pool.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureNodePool.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. If set, there are currently pending changes to the node
	//  pool.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureNodePool.reconciling
	Reconciling *bool `json:"reconciling,omitempty"`

	// Output only. The time at which this node pool was created.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureNodePool.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which this node pool was last updated.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureNodePool.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. A set of errors found in the node pool.
	// +kcc:proto:field=google.cloud.gkemulticloud.v1.AzureNodePool.errors
	Errors []AzureNodePoolError `json:"errors,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpgkemulticloudazurenodepool;gcpgkemulticloudazurenodepools
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// GKEMulticloudAzureNodePool is the Schema for the GKEMulticloudAzureNodePool API
// +k8s:openapi-gen=true
type GKEMulticloudAzureNodePool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   GKEMulticloudAzureNodePoolSpec   `json:"spec,omitempty"`
	Status GKEMulticloudAzureNodePoolStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// GKEMulticloudAzureNodePoolList contains a list of GKEMulticloudAzureNodePool
type GKEMulticloudAzureNodePoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GKEMulticloudAzureNodePool `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GKEMulticloudAzureNodePool{}, &GKEMulticloudAzureNodePoolList{})
}
