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

// +kcc:proto=google.cloud.rapidmigrationassessment.v1.GuestOsScan
type GuestOSScan struct {
	// reference to the corresponding Guest OS Scan in MC Source.
	// +kcc:proto:field=google.cloud.rapidmigrationassessment.v1.GuestOsScan.core_source
	CoreSource *string `json:"coreSource,omitempty"`
}

// +kcc:proto=google.cloud.rapidmigrationassessment.v1.VSphereScan
type VSphereScan struct {
	// reference to the corresponding VSphere Scan in MC Source.
	// +kcc:proto:field=google.cloud.rapidmigrationassessment.v1.VSphereScan.core_source
	CoreSource *string `json:"coreSource,omitempty"`
}

// RapidMigrationAssessmentCollectorSpec defines the desired state of RapidMigrationAssessmentCollector
// +kcc:spec:proto=google.cloud.rapidmigrationassessment.v1.Collector
type RapidMigrationAssessmentCollectorSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Location field is immutable"
	// Immutable.
	// +required
	Location string `json:"location"`

	// The RapidMigrationAssessmentCollector name. If not given, the metadata.name will be used.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// Labels as key value pairs.
	// +kcc:proto:field=google.cloud.rapidmigrationassessment.v1.Collector.labels
	// +optional
	Labels map[string]string `json:"labels,omitempty"`

	// User specified name of the Collector.
	// +kcc:proto:field=google.cloud.rapidmigrationassessment.v1.Collector.display_name
	// +optional
	DisplayName *string `json:"displayName,omitempty"`

	// User specified description of the Collector.
	// +kcc:proto:field=google.cloud.rapidmigrationassessment.v1.Collector.description
	// +optional
	Description *string `json:"description,omitempty"`

	// Service Account email used to ingest data to this Collector.
	// +kcc:proto:field=google.cloud.rapidmigrationassessment.v1.Collector.service_account
	// +optional
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// User specified expected asset count.
	// +kcc:proto:field=google.cloud.rapidmigrationassessment.v1.Collector.expected_asset_count
	// +optional
	ExpectedAssetCount *int64 `json:"expectedAssetCount,omitempty"`

	// How many days to collect data.
	// +kcc:proto:field=google.cloud.rapidmigrationassessment.v1.Collector.collection_days
	// +optional
	CollectionDays *int32 `json:"collectionDays,omitempty"`

	// Uri for EULA (End User License Agreement) from customer.
	// +kcc:proto:field=google.cloud.rapidmigrationassessment.v1.Collector.eula_uri
	// +optional
	EulaURI *string `json:"eulaURI,omitempty"`
}

// RapidMigrationAssessmentCollectorStatus defines the config connector machine state of RapidMigrationAssessmentCollector
type RapidMigrationAssessmentCollectorStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the RapidMigrationAssessmentCollector resource in GCP.
	// +optional
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	// +optional
	ObservedState *RapidMigrationAssessmentCollectorObservedState `json:"observedState,omitempty"`
}

// RapidMigrationAssessmentCollectorObservedState is the state of the RapidMigrationAssessmentCollector resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.rapidmigrationassessment.v1.Collector
type RapidMigrationAssessmentCollectorObservedState struct {
	// Output only. Create time stamp.
	// +kcc:proto:field=google.cloud.rapidmigrationassessment.v1.Collector.create_time
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Update time stamp.
	// +kcc:proto:field=google.cloud.rapidmigrationassessment.v1.Collector.update_time
	// +optional
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Store cloud storage bucket name (which is a guid) created with
	//  this Collector.
	// +kcc:proto:field=google.cloud.rapidmigrationassessment.v1.Collector.bucket
	// +optional
	Bucket *string `json:"bucket,omitempty"`

	// Output only. State of the Collector.
	// +kcc:proto:field=google.cloud.rapidmigrationassessment.v1.Collector.state
	// +optional
	State *string `json:"state,omitempty"`

	// Output only. Client version.
	// +kcc:proto:field=google.cloud.rapidmigrationassessment.v1.Collector.client_version
	// +optional
	ClientVersion *string `json:"clientVersion,omitempty"`

	// Output only. Reference to MC Source Guest Os Scan.
	// +kcc:proto:field=google.cloud.rapidmigrationassessment.v1.Collector.guest_os_scan
	// +optional
	GuestOSScan *GuestOSScan `json:"guestOSScan,omitempty"`

	// Output only. Reference to MC Source vsphere_scan.
	// +kcc:proto:field=google.cloud.rapidmigrationassessment.v1.Collector.vsphere_scan
	// +optional
	VsphereScan *VSphereScan `json:"vsphereScan,omitempty"`
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
