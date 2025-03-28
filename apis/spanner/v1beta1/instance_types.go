// Copyright 2024 Google LLC
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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var SpannerInstanceGVK = GroupVersion.WithKind("SpannerInstance")

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SpannerInstanceSpec defines the desired state of SpannerInstance
// +kcc:proto=google.spanner.admin.instance.v1.Instance
type SpannerInstanceSpec struct {
	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Config field is immutable"
	/* Immutable. The name of the instance's configuration (similar but not
	quite the same as a region) which defines the geographic placement and
	replication of your databases in this instance. It determines where your data
	is stored. Values are typically of the form 'regional-europe-west1' , 'us-central' etc.
	In order to obtain a valid list please consult the
	[Configuration section of the docs](https://cloud.google.com/spanner/docs/instances). */
	Config string `json:"config"`

	/* The descriptive name for this instance as it appears in UIs. Must be
	unique per project and between 4 and 30 characters in length. */
	DisplayName string `json:"displayName"`

	// +optional
	NumNodes *int32 `json:"numNodes,omitempty"`

	// +optional
	ProcessingUnits *int32 `json:"processingUnits,omitempty"`

	// Optional. The autoscaling configuration. Autoscaling is enabled if this
	// field is set. When autoscaling is enabled, node_count and processing_units
	// are treated as OUTPUT_ONLY fields and reflect the current compute capacity
	// allocated to the instance.
	// +optional
	AutoscalingConfig *AutoscalingConfig `json:"autoscalingConfig,omitempty"`

	// Optional. The `Edition` of the current instance.
	// Currently accepted values are STANDARD, ENTERPRISE, ENTERPRISE_PLUS.
	// If edition is unspecified, it has automatically upgraded to the lowest edition that matches your usage pattern.
	// +optional
	Edition *string `json:"edition,omitempty"`

	// The SpannerInstance name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
}

// SpannerInstanceStatus defines the config connector machine state of SpannerInstance
type SpannerInstanceStatus struct {
	/* Conditions represent the latest available observations of the
	   SpannerInstance's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the SpannerInstance resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	/* Instance status: 'CREATING' or 'READY'. */
	// +optional
	State *string `json:"state,omitempty"`

	/* ObservedState is the state of the resource as most recently observed in GCP. */
	// +optional
	ObservedState *SpannerInstanceObservedState `json:"observedState,omitempty"`
}

// SpannerInstanceObservedState is the state of the SpannerInstance resource as most recently observed in GCP.
type SpannerInstanceObservedState struct {
	// NumNodes and ProcessUnits is output fields with AutoScaler is set.
	NumNodes        *int32 `json:"numNodes,omitempty"`
	ProcessingUnits *int32 `json:"processingUnits,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpspannerinstance;gcpspannerinstances
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// SpannerInstance is the Schema for the SpannerInstance API
// +k8s:openapi-gen=true
type SpannerInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   SpannerInstanceSpec   `json:"spec,omitempty"`
	Status SpannerInstanceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// SpannerInstanceList contains a list of SpannerInstance
type SpannerInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SpannerInstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SpannerInstance{}, &SpannerInstanceList{})
}
