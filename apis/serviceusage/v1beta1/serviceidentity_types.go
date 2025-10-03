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

package v1beta1

import (
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ServiceIdentityGVK = GroupVersion.WithKind("ServiceIdentity")

// ServiceIdentitySpec defines the desired state of ServiceIdentity
// +kcc:spec:proto=google.api.serviceusage.v1beta1.ServiceIdentity
type ServiceIdentitySpec struct {
	// The service name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// The project that this service identity belongs to.
	// +required
	ProjectRef *refs.ProjectRef `json:"projectRef,omitempty"`
}

// ServiceIdentityStatus defines the config connector machine state of ServiceIdentity
type ServiceIdentityStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// The email address of the service account that a service producer would use
	//  to access consumer resources.
	// +kcc:proto:field=google.api.serviceusage.v1beta1.ServiceIdentity.email
	Email *string `json:"email,omitempty"`

	/* NOTYET:terraform-compat
	// A unique specifier for the ServiceIdentity resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`
	*/

	/* NOTYET:terraform-compat
	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ServiceIdentityObservedState `json:"observedState,omitempty"`
	*/
}

// ServiceIdentityObservedState is the state of the ServiceIdentity resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.api.serviceusage.v1beta1.ServiceIdentity
type ServiceIdentityObservedState struct {
	// The email address of the service account that a service producer would use
	//  to access consumer resources.
	// +kcc:proto:field=google.api.serviceusage.v1beta1.ServiceIdentity.email
	Email *string `json:"email,omitempty"`

	// The unique and stable id of the service account.
	//  https://cloud.google.com/iam/reference/rest/v1/projects.serviceAccounts#ServiceAccount
	// +kcc:proto:field=google.api.serviceusage.v1beta1.ServiceIdentity.unique_id
	UniqueID *string `json:"uniqueID,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpserviceidentity;gcpserviceidentities
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/stability-level=stable";"cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ServiceIdentity is the Schema for the ServiceIdentity API
// +k8s:openapi-gen=true
type ServiceIdentity struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ServiceIdentitySpec   `json:"spec,omitempty"`
	Status ServiceIdentityStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ServiceIdentityList contains a list of ServiceIdentity
type ServiceIdentityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceIdentity `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ServiceIdentity{}, &ServiceIdentityList{})
}
