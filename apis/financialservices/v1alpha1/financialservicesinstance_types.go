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

var FinancialServicesInstanceGVK = GroupVersion.WithKind("FinancialServicesInstance")

// FinancialServicesInstanceSpec defines the desired state of FinancialServicesInstance
// +kcc:spec:proto=google.cloud.financialservices.v1.Instance
type FinancialServicesInstanceSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The FinancialServicesInstance name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Labels as key value pairs
	// +optional
	Labels map[string]string `json:"labels,omitempty"`

	// Required. The KMS key name used for CMEK (encryption-at-rest).
	//  format:
	//  `projects/{project}/locations/{location}/keyRings/{keyRing}/cryptoKeys/{cryptoKey}`
	//  VPC-SC restrictions apply.
	// +required
	KMSKeyRef *refsv1beta1.KMSCryptoKeyRef `json:"kmsKeyRef"`
}

// FinancialServicesInstanceStatus defines the config connector machine state of FinancialServicesInstance
type FinancialServicesInstanceStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the FinancialServicesInstance resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *FinancialServicesInstanceObservedState `json:"observedState,omitempty"`
}

// FinancialServicesInstanceObservedState is the state of the FinancialServicesInstance resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.financialservices.v1.Instance
type FinancialServicesInstanceObservedState struct {
	// Output only. Timestamp when the Instance was created.
	//  Assigned by the server.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when the Instance was last updated.
	//  Assigned by the server.
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. State of the instance.
	//  Assigned by the server.
	State *string `json:"state,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpfinancialservicesinstance;gcpfinancialservicesinstances
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// FinancialServicesInstance is the Schema for the FinancialServicesInstance API
// +k8s:openapi-gen=true
type FinancialServicesInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   FinancialServicesInstanceSpec   `json:"spec,omitempty"`
	Status FinancialServicesInstanceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// FinancialServicesInstanceList contains a list of FinancialServicesInstance
type FinancialServicesInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FinancialServicesInstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FinancialServicesInstance{}, &FinancialServicesInstanceList{})
}
