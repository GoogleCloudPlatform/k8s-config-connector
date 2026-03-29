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

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DataprocAutoscalingPolicyGVK = GroupVersion.WithKind("DataprocAutoscalingPolicy")

// DataprocAutoscalingPolicySpec defines the desired state of DataprocAutoscalingPolicy
// +kcc:spec:proto=google.cloud.dataproc.v1.AutoscalingPolicy
// +kubebuilder:object:generate=true
type DataprocAutoscalingPolicySpec struct {
	// Immutable. Required. YARN autoscaling configuration.
	// +required
	// +kcc:proto:field=google.cloud.dataproc.v1.AutoscalingPolicy.basic_algorithm
	BasicAlgorithm *BasicAutoscalingAlgorithm `json:"basicAlgorithm"`

	// Immutable. Required. Describes how the autoscaler will operate for primary workers.
	// +required
	// +kcc:proto:field=google.cloud.dataproc.v1.AutoscalingPolicy.worker_config
	WorkerConfig *InstanceGroupAutoscalingPolicyConfig `json:"workerConfig"`

	// Immutable. Optional. Describes how the autoscaler will operate for secondary workers.
	// +kcc:proto:field=google.cloud.dataproc.v1.AutoscalingPolicy.secondary_worker_config
	SecondaryWorkerConfig *SecondaryInstanceGroupAutoscalingPolicyConfig `json:"secondaryWorkerConfig,omitempty"`

	// Immutable. The location for the resource
	// +required
	Location string `json:"location"`

	// Immutable. The Project that this resource belongs to.
	ProjectRef *parent.ProjectRef `json:"projectRef,omitempty"`

	// Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default.
	// +kcc:proto:field=google.cloud.dataproc.v1.AutoscalingPolicy.id
	ResourceID *string `json:"resourceID,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.BasicAutoscalingAlgorithm
// +kubebuilder:object:generate=true
type BasicAutoscalingAlgorithm struct {
	// Required. YARN autoscaling configuration.
	// +required
	// +kcc:proto:field=google.cloud.dataproc.v1.BasicAutoscalingAlgorithm.yarn_config
	YarnConfig *BasicYarnAutoscalingConfig `json:"yarnConfig"`

	// Optional. Duration between scaling events. A scaling period starts after
	//  the update operation from the previous event has completed.
	//
	//  Bounds: [2m, 1d]. Default: 2m.
	// +kcc:proto:field=google.cloud.dataproc.v1.BasicAutoscalingAlgorithm.cooldown_period
	CooldownPeriod *string `json:"cooldownPeriod,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.BasicYarnAutoscalingConfig
// +kubebuilder:object:generate=true
type BasicYarnAutoscalingConfig struct {
	// Required. Timeout for YARN graceful decommissioning of Node Managers.
	//  Specifies the duration to wait for jobs to complete before forcefully
	//  removing workers (and potentially interrupting jobs). Only applicable to
	//  downscaling operations.
	//
	//  Bounds: [0s, 1d].
	// +required
	// +kcc:proto:field=google.cloud.dataproc.v1.BasicYarnAutoscalingConfig.graceful_decommission_timeout
	GracefulDecommissionTimeout *string `json:"gracefulDecommissionTimeout"`

	// Required. Fraction of average YARN pending memory in the last cooldown
	//  period for which to add workers. A scale-up factor of 1.0 will result in
	//  scaling up so that there is no pending memory remaining after the update
	//  (more aggressive scaling). A scale-up factor closer to 0 will result in a
	//  smaller magnitude of scaling up (less aggressive scaling). See [How
	//  autoscaling
	//  works](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/autoscaling#how_autoscaling_works)
	//  for more information.
	//
	//  Bounds: [0.0, 1.0].
	// +required
	// +kubebuilder:validation:Format=double
	// +kcc:proto:field=google.cloud.dataproc.v1.BasicYarnAutoscalingConfig.scale_up_factor
	ScaleUpFactor *float64 `json:"scaleUpFactor"`

	// Required. Fraction of average YARN pending memory in the last cooldown
	//  period for which to remove workers. A scale-down factor of 1 will result in
	//  scaling down so that there is no available memory remaining after the
	//  update (more aggressive scaling). A scale-down factor of 0 disables
	//  removing workers, which can be beneficial for autoscaling a single job.
	//  See [How autoscaling
	//  works](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/autoscaling#how_autoscaling_works)
	//  for more information.
	//
	//  Bounds: [0.0, 1.0].
	// +required
	// +kubebuilder:validation:Format=double
	// +kcc:proto:field=google.cloud.dataproc.v1.BasicYarnAutoscalingConfig.scale_down_factor
	ScaleDownFactor *float64 `json:"scaleDownFactor"`

	// Optional. Minimum scale-up threshold as a fraction of total cluster size
	//  before scaling occurs. For example, in a 20-worker cluster, a threshold of
	//  0.1 means the autoscaler must recommend at least a 2-worker scale-up for
	//  the cluster to scale. A threshold of 0 means the autoscaler will scale up
	//  on any recommended change.
	//
	//  Bounds: [0.0, 1.0]. Default: 0.0.
	// +kubebuilder:validation:Format=double
	// +kcc:proto:field=google.cloud.dataproc.v1.BasicYarnAutoscalingConfig.scale_up_min_worker_fraction
	ScaleUpMinWorkerFraction *float64 `json:"scaleUpMinWorkerFraction,omitempty"`

	// Optional. Minimum scale-down threshold as a fraction of total cluster size
	//  before scaling occurs. For example, in a 20-worker cluster, a threshold of
	//  0.1 means the autoscaler must recommend at least a 2 worker scale-down for
	//  the cluster to scale. A threshold of 0 means the autoscaler will scale down
	//  on any recommended change.
	//
	//  Bounds: [0.0, 1.0]. Default: 0.0.
	// +kubebuilder:validation:Format=double
	// +kcc:proto:field=google.cloud.dataproc.v1.BasicYarnAutoscalingConfig.scale_down_min_worker_fraction
	ScaleDownMinWorkerFraction *float64 `json:"scaleDownMinWorkerFraction,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.InstanceGroupAutoscalingPolicyConfig
// +kubebuilder:object:generate=true
type InstanceGroupAutoscalingPolicyConfig struct {
	// Optional. Minimum number of instances for this group.
	//
	//  Primary workers - Bounds: [2, max_instances]. Default: 2.
	//  Secondary workers - Bounds: [0, max_instances]. Default: 0.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupAutoscalingPolicyConfig.min_instances
	MinInstances *int64 `json:"minInstances,omitempty"`

	// Required. Maximum number of instances for this group. Required for primary
	//  workers. Note that by default, clusters will not use secondary workers.
	//  Required for secondary workers if the minimum secondary instances is set.
	//
	//  Primary workers - Bounds: [min_instances, ).
	//  Secondary workers - Bounds: [min_instances, ). Default: 0.
	// +required
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupAutoscalingPolicyConfig.max_instances
	MaxInstances *int64 `json:"maxInstances"`

	// Optional. Weight for the instance group, which is used to determine the
	//  fraction of total workers in the cluster from this instance group.
	//  For example, if primary workers have weight 2, and secondary workers have
	//  weight 1, the cluster will have approximately 2 primary workers for each
	//  secondary worker.
	//
	//  The cluster may not reach the specified balance if constrained
	//  by min/max bounds or other autoscaling settings. For example, if
	//  `max_instances` for secondary workers is 0, then only primary workers will
	//  be added. The cluster can also be out of balance when created.
	//
	//  If weight is not set on any instance group, the cluster will default to
	//  equal weight for all groups: the cluster will attempt to maintain an equal
	//  number of workers in each group within the configured size bounds for each
	//  group. If weight is set for one group only, the cluster will default to
	//  zero weight on the unset group. For example if weight is set only on
	//  primary workers, the cluster will use primary workers only and no
	//  secondary workers.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupAutoscalingPolicyConfig.weight
	Weight *int64 `json:"weight,omitempty"`
}

// +kcc:proto=google.cloud.dataproc.v1.InstanceGroupAutoscalingPolicyConfig
// +kubebuilder:object:generate=true
type SecondaryInstanceGroupAutoscalingPolicyConfig struct {
	// Optional. Minimum number of instances for this group.
	//
	//  Primary workers - Bounds: [2, max_instances]. Default: 2.
	//  Secondary workers - Bounds: [0, max_instances]. Default: 0.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupAutoscalingPolicyConfig.min_instances
	MinInstances *int64 `json:"minInstances,omitempty"`

	// Required. Maximum number of instances for this group. Required for primary
	//  workers. Note that by default, clusters will not use secondary workers.
	//  Required for secondary workers if the minimum secondary instances is set.
	//
	//  Primary workers - Bounds: [min_instances, ).
	//  Secondary workers - Bounds: [min_instances, ). Default: 0.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupAutoscalingPolicyConfig.max_instances
	MaxInstances *int64 `json:"maxInstances,omitempty"`

	// Optional. Weight for the instance group, which is used to determine the
	//  fraction of total workers in the cluster from this instance group.
	//  For example, if primary workers have weight 2, and secondary workers have
	//  weight 1, the cluster will have approximately 2 primary workers for each
	//  secondary worker.
	//
	//  The cluster may not reach the specified balance if constrained
	//  by min/max bounds or other autoscaling settings. For example, if
	//  `max_instances` for secondary workers is 0, then only primary workers will
	//  be added. The cluster can also be out of balance when created.
	//
	//  If weight is not set on any instance group, the cluster will default to
	//  equal weight for all groups: the cluster will attempt to maintain an equal
	//  number of workers in each group within the configured size bounds for each
	//  group. If weight is set for one group only, the cluster will default to
	//  zero weight on the unset group. For example if weight is set only on
	//  primary workers, the cluster will use primary workers only and no
	//  secondary workers.
	// +kcc:proto:field=google.cloud.dataproc.v1.InstanceGroupAutoscalingPolicyConfig.weight
	Weight *int64 `json:"weight,omitempty"`
}

// DataprocAutoscalingPolicyStatus defines the config connector machine state of DataprocAutoscalingPolicy
// +kubebuilder:object:generate=true
type DataprocAutoscalingPolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	// +kubebuilder:validation:Format=""
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DataprocAutoscalingPolicy resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdataprocautoscalingpolicy;gcpdataprocautoscalingpolicies
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/dcl2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DataprocAutoscalingPolicy is the Schema for the DataprocAutoscalingPolicy API
// +k8s:openapi-gen=true
type DataprocAutoscalingPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DataprocAutoscalingPolicySpec   `json:"spec,omitempty"`
	Status DataprocAutoscalingPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DataprocAutoscalingPolicyList contains a list of DataprocAutoscalingPolicy
type DataprocAutoscalingPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataprocAutoscalingPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DataprocAutoscalingPolicy{}, &DataprocAutoscalingPolicyList{})
}
