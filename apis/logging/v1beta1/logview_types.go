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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	addonv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var LoggingLogViewGVK = GroupVersion.WithKind("LoggingLogView")

// LoggingLogViewSpec defines the desired state of LoggingLogView
// +kcc:spec:proto=google.logging.v2.LogView
type LoggingLogViewSpec struct {
	/* Immutable. The BillingAccount that this resource belongs to. Only one of [billingAccountRef, folderRef, organizationRef, projectRef] may be specified. */
	// +optional
	BillingAccountRef *addonv1alpha1.ResourceRef `json:"billingAccountRef,omitempty"`

	/* Immutable. */
	BucketRef LoggingLogBucketRef `json:"bucketRef"`

	/* Describes this view. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* Filter that restricts which log entries in a bucket are visible in this view. Filters are restricted to be a logical AND of ==/!= of any of the following: - originating project/folder/organization/billing account. - resource type - log id For example: SOURCE("projects/myproject") AND resource.type = "gce_instance" AND LOG_ID("stdout") */
	// +optional
	Filter *string `json:"filter,omitempty"`

	/* Immutable. The Folder that this resource belongs to. Only one of [billingAccountRef, folderRef, organizationRef, projectRef] may be specified. */
	// +optional
	FolderRef *refsv1beta1.FolderRef `json:"folderRef,omitempty"`

	/* Immutable. The location of the resource. The supported locations are: global, us-central1, us-east1, us-west1, asia-east1, europe-west1. */
	// +optional
	Location *string `json:"location,omitempty"`

	/* Immutable. The Organization that this resource belongs to. Only one of [billingAccountRef, folderRef, organizationRef, projectRef] may be specified. */
	// +optional
	OrganizationRef *refsv1beta1.OrganizationRef `json:"organizationRef,omitempty"`

	/* Immutable. The Project that this resource belongs to. Only one of [billingAccountRef, folderRef, organizationRef, projectRef] may be specified. */
	// +optional
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef,omitempty"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

// LoggingLogViewStatus defines the config connector machine state of LoggingLogView
type LoggingLogViewStatus struct {
	/* Conditions represent the latest available observations of the
	   LoggingLogView's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	/* Output only. The creation timestamp of the view. */
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* Output only. The last update timestamp of the view. */
	// +optional
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcplogginglogview;gcplogginglogviews
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// LoggingLogView is the Schema for the LoggingLogView API
// +k8s:openapi-gen=true
type LoggingLogView struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   LoggingLogViewSpec   `json:"spec,omitempty"`
	Status LoggingLogViewStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// LoggingLogViewList contains a list of LoggingLogView
type LoggingLogViewList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoggingLogView `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LoggingLogView{}, &LoggingLogViewList{})
}
