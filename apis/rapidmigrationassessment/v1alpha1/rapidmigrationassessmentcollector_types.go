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

var RapidMigrationAssessmentCollectorGVK = GroupVersion.WithKind("RapidMigrationAssessmentCollector")

// RapidMigrationAssessmentCollectorSpec defines the desired state of RapidMigrationAssessmentCollector
// +kcc:spec:proto=google.cloud.rapidmigrationassessment.v1.Collector
type RapidMigrationAssessmentCollectorSpec struct {
	// Immutable. The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Location field is immutable"
	// Immutable. The location of this resource.
	// +required
	Location string `json:"location"`

	// Immutable. The RapidMigrationAssessmentCollector name. If not given, the metadata.name will be used.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// User specified name of the Collector.
	// +optional
	DisplayName *string `json:"displayName,omitempty"`

	// User specified description of the Collector.
	// +optional
	Description *string `json:"description,omitempty"`

	// Service Account email used to ingest data to this Collector.
	// +optional
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// User specified expected asset count.
	// +optional
	ExpectedAssetCount *int64 `json:"expectedAssetCount,omitempty"`

	// How many days to collect data.
	// +optional
	CollectionDays *int32 `json:"collectionDays,omitempty"`

	// Uri for EULA (End User License Agreement) from customer.
	// +optional
	EulaURI *string `json:"eulaURI,omitempty"`

	// Labels as key value pairs.
	// +optional
	Labels map[string]string `json:"labels,omitempty"`
}

// RapidMigrationAssessmentCollectorStatus defines the config connector machine state of RapidMigrationAssessmentCollector
type RapidMigrationAssessmentCollectorStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the RapidMigrationAssessmentCollector resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *RapidMigrationAssessmentCollectorObservedState `json:"observedState,omitempty"`
}

// RapidMigrationAssessmentCollectorObservedState is the state of the RapidMigrationAssessmentCollector resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.rapidmigrationassessment.v1.Collector
type RapidMigrationAssessmentCollectorObservedState struct {
	// Output only. Create time stamp.
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Update time stamp.
	// +optional
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Store cloud storage bucket name (which is a guid) created with
	//  this Collector.
	// +optional
	Bucket *string `json:"bucket,omitempty"`

	// Output only. State of the Collector.
	// +kubebuilder:validation:Enum=STATE_UNSPECIFIED;STATE_INITIALIZING;STATE_READY_TO_USE;STATE_REGISTERED;STATE_ACTIVE;STATE_PAUSED;STATE_DELETING;STATE_DECOMMISSIONED;STATE_ERROR
	// +optional
	State *string `json:"state,omitempty"`

	// Output only. Client version.
	// +optional
	ClientVersion *string `json:"clientVersion,omitempty"`

	// Output only. Reference to MC Source Guest Os Scan.
	// +optional
	GuestOSScan *GuestOSScan `json:"guestOSScan,omitempty"`

	// Output only. Reference to MC Source vsphere_scan.
	// +optional
	VsphereScan *VSphereScan `json:"vsphereScan,omitempty"`
}

type GuestOSScan struct {
	// reference to the corresponding Guest OS Scan in MC Source.
	// +optional
	CoreSource *string `json:"coreSource,omitempty"`
}

type VSphereScan struct {
	// reference to the corresponding VSphere Scan in MC Source.
	// +optional
	CoreSource *string `json:"coreSource,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcprapidmigrationassessmentcollector;gcprapidmigrationassessmentcollectors
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// RapidMigrationAssessmentCollector is the Schema for the RapidMigrationAssessmentCollector API
// +k8s:openapi-gen=true
type RapidMigrationAssessmentCollector struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   RapidMigrationAssessmentCollectorSpec   `json:"spec,omitempty"`
	Status RapidMigrationAssessmentCollectorStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// RapidMigrationAssessmentCollectorList contains a list of RapidMigrationAssessmentCollector
type RapidMigrationAssessmentCollectorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RapidMigrationAssessmentCollector `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RapidMigrationAssessmentCollector{}, &RapidMigrationAssessmentCollectorList{})
}
