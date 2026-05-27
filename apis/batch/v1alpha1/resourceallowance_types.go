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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CloudBatchResourceAllowanceGVK = GroupVersion.WithKind("CloudBatchResourceAllowance")

// CloudBatchResourceAllowanceSpec defines the desired state of CloudBatchResourceAllowance
// +kcc:spec:proto=google.cloud.batch.v1alpha.ResourceAllowance
type CloudBatchResourceAllowanceSpec struct {
	// The project that this resource belongs to.
	// +kubebuilder:validation:Required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	// +kubebuilder:validation:Required
	Location string `json:"location"`

	// The CloudBatchResourceAllowance name. If not given, the metadata.name will be used.
	// +kubebuilder:validation:Optional
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Labels are attributes that can be set and used by both the
	// user and by Batch.
	// +kubebuilder:validation:Optional
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Notification configurations.
	// +kubebuilder:validation:Optional
	Notifications []Notification `json:"notifications,omitempty"`

	// The detail of usage resource allowance.
	// +kubebuilder:validation:Optional
	UsageResourceAllowance *UsageResourceAllowance `json:"usageResourceAllowance,omitempty"`
}

// CloudBatchResourceAllowanceStatus defines the config connector machine state of CloudBatchResourceAllowance
type CloudBatchResourceAllowanceStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the CloudBatchResourceAllowance resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *CloudBatchResourceAllowanceObservedState `json:"observedState,omitempty"`
}

// CloudBatchResourceAllowanceObservedState is the state of the CloudBatchResourceAllowance resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.batch.v1alpha.ResourceAllowance
type CloudBatchResourceAllowanceObservedState struct {
	// Output only. A system generated unique ID (in UUID4 format) for the
	// ResourceAllowance.
	Uid *string `json:"uid,omitempty"`

	// Output only. Time when the ResourceAllowance was created.
	CreateTime *string `json:"createTime,omitempty"`

	// The detail of usage resource allowance.
	UsageResourceAllowance *UsageResourceAllowanceObservedState `json:"usageResourceAllowance,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcloudbatchresourceallowance;gcpcloudbatchresourceallowances
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CloudBatchResourceAllowance is the Schema for the CloudBatchResourceAllowance API
// +k8s:openapi-gen=true
type CloudBatchResourceAllowance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CloudBatchResourceAllowanceSpec   `json:"spec,omitempty"`
	Status CloudBatchResourceAllowanceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CloudBatchResourceAllowanceList contains a list of CloudBatchResourceAllowance
type CloudBatchResourceAllowanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudBatchResourceAllowance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudBatchResourceAllowance{}, &CloudBatchResourceAllowanceList{})
}

// +kcc:observedstate:proto=google.cloud.batch.v1alpha.UsageResourceAllowance
type UsageResourceAllowanceObservedState struct {
	// Output only. Status of a usage ResourceAllowance.
	// +kcc:proto:field=google.cloud.batch.v1alpha.UsageResourceAllowance.status
	Status *UsageResourceAllowanceStatusObservedState `json:"status,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.batch.v1alpha.UsageResourceAllowanceStatus
type UsageResourceAllowanceStatusObservedState struct {
	// Output only. ResourceAllowance state.
	// +kcc:proto:field=google.cloud.batch.v1alpha.UsageResourceAllowanceStatus.state
	State *string `json:"state,omitempty"`

	// Output only. ResourceAllowance consumption status for usage resources.
	// +kcc:proto:field=google.cloud.batch.v1alpha.UsageResourceAllowanceStatus.limit_status
	LimitStatus *UsageResourceAllowanceStatus_LimitStatusObservedState `json:"limitStatus,omitempty"`

	// Output only. The report of ResourceAllowance consumptions in a time period.
	// +kcc:proto:field=google.cloud.batch.v1alpha.UsageResourceAllowanceStatus.report
	Report *UsageResourceAllowanceStatus_ConsumptionReportObservedState `json:"report,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.batch.v1alpha.UsageResourceAllowanceStatus.LimitStatus
type UsageResourceAllowanceStatus_LimitStatusObservedState struct {
	// Output only. The consumption interval.
	// +kcc:proto:field=google.cloud.batch.v1alpha.UsageResourceAllowanceStatus.LimitStatus.consumption_interval
	ConsumptionInterval *Interval `json:"consumptionInterval,omitempty"`

	// Output only. Limit value of a UsageResourceAllowance within its one
	//  duration.
	// +kcc:proto:field=google.cloud.batch.v1alpha.UsageResourceAllowanceStatus.LimitStatus.limit
	Limit *float64 `json:"limit,omitempty"`

	// Output only. Accumulated consumption during `consumption_interval`.
	// +kcc:proto:field=google.cloud.batch.v1alpha.UsageResourceAllowanceStatus.LimitStatus.consumed
	Consumed *float64 `json:"consumed,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.batch.v1alpha.UsageResourceAllowanceStatus.ConsumptionReport
type UsageResourceAllowanceStatus_ConsumptionReportObservedState struct {
	// Output only. ResourceAllowance consumptions in the latest calendar
	// period. Key is the calendar period in string format. Batch currently
	// supports HOUR, DAY, MONTH and YEAR.
	// +kcc:proto:field=google.cloud.batch.v1alpha.UsageResourceAllowanceStatus.ConsumptionReport.latest_period_consumptions
	LatestPeriodConsumptions map[string]UsageResourceAllowanceStatus_PeriodConsumptionObservedState `json:"latestPeriodConsumptions,omitempty"`
}

// +kcc:observedstate:proto=google.cloud.batch.v1alpha.UsageResourceAllowanceStatus.PeriodConsumption
type UsageResourceAllowanceStatus_PeriodConsumptionObservedState struct {
	// Output only. The consumption interval.
	// +kcc:proto:field=google.cloud.batch.v1alpha.UsageResourceAllowanceStatus.PeriodConsumption.consumption_interval
	ConsumptionInterval *Interval `json:"consumptionInterval,omitempty"`

	// Output only. Accumulated consumption during `consumption_interval`.
	// +kcc:proto:field=google.cloud.batch.v1alpha.UsageResourceAllowanceStatus.PeriodConsumption.consumed
	Consumed *float64 `json:"consumed,omitempty"`
}

// +kcc:proto=google.type.Interval
type Interval struct {
	// Optional. Inclusive start of the interval.
	// +kcc:proto:field=google.type.Interval.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Optional. Exclusive end of the interval.
	// +kcc:proto:field=google.type.Interval.end_time
	EndTime *string `json:"endTime,omitempty"`
}

// +kcc:proto=google.cloud.batch.v1alpha.UsageResourceAllowance
type UsageResourceAllowance struct {
	// Required. Spec of a usage ResourceAllowance.
	// +kcc:proto:field=google.cloud.batch.v1alpha.UsageResourceAllowance.spec
	Spec *UsageResourceAllowanceSpec `json:"spec,omitempty"`
}

// +kcc:proto=google.cloud.batch.v1alpha.UsageResourceAllowanceSpec
type UsageResourceAllowanceSpec struct {
	// Required. Spec type is unique for each usage ResourceAllowance.
	// Batch now only supports type as "cpu-core-hours" for CPU usage consumption tracking.
	// +kcc:proto:field=google.cloud.batch.v1alpha.UsageResourceAllowanceSpec.type
	Type *string `json:"type,omitempty"`

	// Required. Threshold of a UsageResourceAllowance limiting how many resources can be consumed for each type.
	// +kcc:proto:field=google.cloud.batch.v1alpha.UsageResourceAllowanceSpec.limit
	Limit *UsageResourceAllowanceSpec_Limit `json:"limit,omitempty"`
}

// +kcc:proto=google.cloud.batch.v1alpha.UsageResourceAllowanceSpec.Limit
type UsageResourceAllowanceSpec_Limit struct {
	// Optional. A CalendarPeriod represents the abstract concept of a time period that has a canonical start.
	// +kcc:proto:field=google.cloud.batch.v1alpha.UsageResourceAllowanceSpec.Limit.calendar_period
	CalendarPeriod *string `json:"calendarPeriod,omitempty"`

	// Required. Limit value of a UsageResourceAllowance within its one duration.
	// +kcc:proto:field=google.cloud.batch.v1alpha.UsageResourceAllowanceSpec.Limit.limit
	Limit *float64 `json:"limit,omitempty"`
}

// +kcc:proto=google.cloud.batch.v1alpha.Notification
type Notification struct {
	// Required. The Pub/Sub topic where notifications like the resource allowance
	// state changes will be published.
	// +kcc:proto:field=google.cloud.batch.v1alpha.Notification.pubsub_topic
	PubsubTopic *string `json:"pubsubTopic,omitempty"`
}
