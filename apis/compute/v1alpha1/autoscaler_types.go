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
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ComputeRegionAutoscalerGVK = GroupVersion.WithKind("ComputeRegionAutoscaler")

// ComputeRegionAutoscalerSpec defines the desired state of ComputeRegionAutoscaler
// +kcc:spec:proto=google.cloud.compute.v1.Autoscaler
type ComputeRegionAutoscalerSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource (region).
	Location string `json:"location"`

	// The ComputeRegionAutoscaler name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// The configuration parameters for the autoscaling algorithm.
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.Autoscaler.autoscaling_policy
	AutoscalingPolicy *AutoscalingPolicy `json:"autoscalingPolicy"`

	// An optional description of this resource.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.Autoscaler.description
	Description *string `json:"description,omitempty"`

	// Reference to the managed instance group that this autoscaler will scale.
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.Autoscaler.target
	TargetRef *computev1beta1.ComputeInstanceGroupManagerRef `json:"targetRef"`
}

// +kcc:proto=google.cloud.compute.v1.AutoscalingPolicy
type AutoscalingPolicy struct {
	// The number of seconds that your application takes to initialize on a VM instance. This is referred to as the [initialization period](/compute/docs/autoscaler#cool_down_period). Specifying an accurate initialization period improves autoscaler decisions. For example, when scaling out, the autoscaler ignores data from VMs that are still initializing because those VMs might not yet represent normal usage of your application. The default initialization period is 60 seconds. Initialization periods might vary because of numerous factors. We recommend that you test how long your application takes to initialize. To do this, create a VM and time your application's startup process.
	// +kcc:proto:field=google.cloud.compute.v1.AutoscalingPolicy.cool_down_period_sec
	CoolDownPeriodSec *int32 `json:"coolDownPeriodSec,omitempty"`

	// Defines the CPU utilization policy that allows the autoscaler to scale based on the average CPU utilization of a managed instance group.
	// +kcc:proto:field=google.cloud.compute.v1.AutoscalingPolicy.cpu_utilization
	CPUUtilization *AutoscalingPolicyCPUUtilization `json:"cpuUtilization,omitempty"`

	// Configuration parameters of autoscaling based on a custom metric.
	// +kcc:proto:field=google.cloud.compute.v1.AutoscalingPolicy.custom_metric_utilizations
	CustomMetricUtilizations []AutoscalingPolicyCustomMetricUtilization `json:"customMetricUtilizations,omitempty"`

	// Configuration parameters of autoscaling based on load balancer.
	// +kcc:proto:field=google.cloud.compute.v1.AutoscalingPolicy.load_balancing_utilization
	LoadBalancingUtilization *AutoscalingPolicyLoadBalancingUtilization `json:"loadBalancingUtilization,omitempty"`

	// The maximum number of instances that the autoscaler can scale out to. This is required when creating or updating an autoscaler. The maximum number of replicas must not be lower than minimal number of replicas.
	// +kcc:proto:field=google.cloud.compute.v1.AutoscalingPolicy.max_num_replicas
	MaxNumReplicas *int32 `json:"maxNumReplicas,omitempty"`

	// The minimum number of replicas that the autoscaler can scale in to. This cannot be less than 0. If not provided, autoscaler chooses a default value depending on maximum number of instances allowed.
	// +kcc:proto:field=google.cloud.compute.v1.AutoscalingPolicy.min_num_replicas
	MinNumReplicas *int32 `json:"minNumReplicas,omitempty"`

	// Defines the operating mode for this policy. The following modes are available: - OFF: Disables the autoscaler but maintains its configuration. - ONLY_SCALE_OUT: Restricts the autoscaler to add VM instances only. - ON: Enables all autoscaler activities according to its policy. For more information, see "Turning off or restricting an autoscaler"
	//  Check the Mode enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.AutoscalingPolicy.mode
	Mode *string `json:"mode,omitempty"`

	// +kcc:proto:field=google.cloud.compute.v1.AutoscalingPolicy.scale_in_control
	ScaleInControl *AutoscalingPolicyScaleInControl `json:"scaleInControl,omitempty"`

	// Scaling schedules defined for an autoscaler. Multiple schedules can be set on an autoscaler and they can overlap.
	// +kcc:proto:field=google.cloud.compute.v1.AutoscalingPolicy.scaling_schedules
	ScalingSchedules []AutoscalingPolicyScalingSchedule `json:"scalingSchedules,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.AutoscalingPolicyScalingSchedule
type AutoscalingPolicyScalingSchedule struct {
	Name string `json:"name"`

	// A description of a scaling schedule.
	// +kcc:proto:field=google.cloud.compute.v1.AutoscalingPolicyScalingSchedule.description
	Description *string `json:"description,omitempty"`

	// A boolean value that specifies whether a scaling schedule can influence autoscaler recommendations. If set to true, then a scaling schedule has no effect. This field is optional, and its value is false by default.
	// +kcc:proto:field=google.cloud.compute.v1.AutoscalingPolicyScalingSchedule.disabled
	Disabled *bool `json:"disabled,omitempty"`

	// The duration of time intervals, in seconds, for which this scaling schedule is to run. The minimum allowed value is 300. This field is required.
	// +kcc:proto:field=google.cloud.compute.v1.AutoscalingPolicyScalingSchedule.duration_sec
	DurationSec *int32 `json:"durationSec,omitempty"`

	// The minimum number of VM instances that the autoscaler will recommend in time intervals starting according to schedule. This field is required.
	// +kcc:proto:field=google.cloud.compute.v1.AutoscalingPolicyScalingSchedule.min_required_replicas
	MinRequiredReplicas *int32 `json:"minRequiredReplicas,omitempty"`

	// The start timestamps of time intervals when this scaling schedule is to provide a scaling signal. This field uses the extended cron format (with an optional year field). The expression can describe a single timestamp if the optional year is set, in which case the scaling schedule runs once. The schedule is interpreted with respect to time_zone. This field is required. Note: These timestamps only describe when autoscaler starts providing the scaling signal. The VMs need additional time to become serving.
	// +kcc:proto:field=google.cloud.compute.v1.AutoscalingPolicyScalingSchedule.schedule
	Schedule *string `json:"schedule,omitempty"`

	// The time zone to use when interpreting the schedule. The value of this field must be a time zone name from the tz database: https://en.wikipedia.org/wiki/Tz_database. This field is assigned a default value of "UTC" if left empty.
	// +kcc:proto:field=google.cloud.compute.v1.AutoscalingPolicyScalingSchedule.time_zone
	TimeZone *string `json:"timeZone,omitempty"`
}

// ComputeRegionAutoscalerStatus defines the config connector machine state of ComputeRegionAutoscaler
type ComputeRegionAutoscalerStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeRegionAutoscaler resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ComputeRegionAutoscalerObservedState `json:"observedState,omitempty"`
}

// ComputeRegionAutoscalerObservedState is the state of the ComputeRegionAutoscaler resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.Autoscaler
type ComputeRegionAutoscalerObservedState struct {
	// [Output Only] Creation timestamp in RFC3339 text format.
	// +kcc:proto:field=google.cloud.compute.v1.Autoscaler.creation_timestamp
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// [Output Only] The unique identifier for the resource. This identifier is defined by the server.
	// +kcc:proto:field=google.cloud.compute.v1.Autoscaler.id
	ID *uint64 `json:"id,omitempty"`

	// [Output Only] Target recommended MIG size (number of instances) computed by autoscaler.
	// +kcc:proto:field=google.cloud.compute.v1.Autoscaler.recommended_size
	RecommendedSize *int32 `json:"recommendedSize,omitempty"`

	// [Output Only] Server-defined URL for the resource.
	// +kcc:proto:field=google.cloud.compute.v1.Autoscaler.self_link
	SelfLink *string `json:"selfLink,omitempty"`

	// [Output Only] The status of the autoscaler configuration.
	// +kcc:proto:field=google.cloud.compute.v1.Autoscaler.status
	Status *string `json:"status,omitempty"`

	// [Output Only] Human-readable details about the current state of the autoscaler.
	// +kcc:proto:field=google.cloud.compute.v1.Autoscaler.status_details
	StatusDetails []AutoscalerStatusDetails `json:"statusDetails,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputeregionautoscaler;gcpcomputeregionautoscalers
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeRegionAutoscaler is the Schema for the ComputeRegionAutoscaler API
// +k8s:openapi-gen=true
type ComputeRegionAutoscaler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeRegionAutoscalerSpec   `json:"spec,omitempty"`
	Status ComputeRegionAutoscalerStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeRegionAutoscalerList contains a list of ComputeRegionAutoscaler
type ComputeRegionAutoscalerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeRegionAutoscaler `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeRegionAutoscaler{}, &ComputeRegionAutoscalerList{})
}
