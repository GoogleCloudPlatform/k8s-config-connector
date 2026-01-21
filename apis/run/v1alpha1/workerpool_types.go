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
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var RunWorkerPoolGVK = GroupVersion.WithKind("RunWorkerPool")

// RunWorkerPoolSpec defines the desired state of RunWorkerPool
// +kcc:spec:proto=google.cloud.run.v2.WorkerPool
type RunWorkerPoolSpec struct {
	// Required. Defines the parent path of the resource.
	*parent.ProjectAndLocationRef `json:",inline"`

	// The RunWorkerPool name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// User-provided description of the WorkerPool. This field currently has a
	//  512-character limit.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPool.description
	Description *string `json:"description,omitempty"`

	// Optional. Unstructured key value map that may be set by external tools to
	//  store and arbitrary metadata. They are not queryable and should be
	//  preserved when modifying objects.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPool.annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// Arbitrary identifier for the API client.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPool.client
	Client *string `json:"client,omitempty"`

	// Arbitrary version identifier for the API client.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPool.client_version
	ClientVersion *string `json:"clientVersion,omitempty"`

	// Optional. The launch stage as defined by Google Cloud Platform Launch Stages.
	//  Cloud Run supports `ALPHA`, `BETA`, and `GA`. If no value is specified, GA is assumed.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPool.launch_stage
	LaunchStage *string `json:"launchStage,omitempty"`

	// Optional. Settings for the Binary Authorization feature.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPool.binary_authorization
	BinaryAuthorization *BinaryAuthorization `json:"binaryAuthorization,omitempty"`

	// Required. The template used to create revisions for this WorkerPool.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPool.template
	Template *WorkerPoolRevisionTemplate `json:"template,omitempty"`

	// Optional. Specifies how to distribute instances over a collection of
	//  Revisions belonging to the WorkerPool.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPool.instance_splits
	InstanceSplits []InstanceSplit `json:"instanceSplits,omitempty"`

	// Optional. Specifies worker-pool-level scaling settings
	// +kcc:proto:field=google.cloud.run.v2.WorkerPool.scaling
	Scaling *WorkerPoolScaling `json:"scaling,omitempty"`

	// One or more custom audiences that you want this worker pool to support.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPool.custom_audiences
	CustomAudiences []string `json:"customAudiences,omitempty"`
}

// RunWorkerPoolStatus defines the config connector machine state of RunWorkerPool
type RunWorkerPoolStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the RunWorkerPool resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *RunWorkerPoolObservedState `json:"observedState,omitempty"`

	// LastModifiedCookie is used for stateful reconciliation to detect changes.
	LastModifiedCookie *string `json:"lastModifiedCookie,omitempty"`
}

// RunWorkerPoolObservedState is the state of the RunWorkerPool resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.run.v2.WorkerPool
type RunWorkerPoolObservedState struct {
	// Output only. Server assigned unique identifier for the trigger.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPool.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. The creation time.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPool.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The last-modified time.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPool.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The deletion time.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPool.delete_time
	DeleteTime *string `json:"deleteTime,omitempty"`

	// Output only. For a deleted resource, the time after which it will be permanently deleted.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPool.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`

	// Output only. Email address of the authenticated creator.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPool.creator
	Creator *string `json:"creator,omitempty"`

	// Output only. Email address of the last authenticated modifier.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPool.last_modifier
	LastModifier *string `json:"lastModifier,omitempty"`

	// Output only. The Condition of this WorkerPool, containing its readiness status.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPool.terminal_condition
	TerminalCondition *Condition `json:"terminalCondition,omitempty"`

	// Output only. The Conditions of all other associated sub-resources.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPool.conditions
	Conditions []Condition `json:"conditions,omitempty"`

	// Output only. Name of the latest revision that is serving traffic.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPool.latest_ready_revision
	LatestReadyRevision *string `json:"latestReadyRevision,omitempty"`

	// Output only. Name of the last created revision.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPool.latest_created_revision
	LatestCreatedRevision *string `json:"latestCreatedRevision,omitempty"`

	// Output only. Detailed status information for corresponding instance splits.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPool.instance_split_statuses
	InstanceSplitStatuses []InstanceSplitStatus `json:"instanceSplitStatuses,omitempty"`

	// Output only. Returns true if the WorkerPool is currently being acted upon by the system.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPool.reconciling
	Reconciling *bool `json:"reconciling,omitempty"`

	// Output only. A system-generated fingerprint for this version of the resource.
	// +kcc:proto:field=google.cloud.run.v2.WorkerPool.etag
	Etag *string `json:"etag,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcprunworkerpool;gcprunworkerpools
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// RunWorkerPool is the Schema for the RunWorkerPool API
// +k8s:openapi-gen=true
type RunWorkerPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   RunWorkerPoolSpec   `json:"spec,omitempty"`
	Status RunWorkerPoolStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// RunWorkerPoolList contains a list of RunWorkerPool
type RunWorkerPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RunWorkerPool `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RunWorkerPool{}, &RunWorkerPoolList{})
}
