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
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AccessContextManagerServicePerimeterResourceGVK = GroupVersion.WithKind("AccessContextManagerServicePerimeterResource")

type ServicePerimeterResourcePerimeterNameRef struct {
	/* Only the `external` field is supported to configure the reference.

	The name of the Service Perimeter to add this resource to.
	Referencing a resource name leads to recursive reference and Config Connector does not support the feature for now. */
	External string `json:"external,omitempty"`

	/* Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names */
	Name string `json:"name,omitempty"`

	/* Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ */
	Namespace string `json:"namespace,omitempty"`
}

type ServicePerimeterResourceResourceRef struct {
	/* Allowed value: string of the format `projects/{{value}}`, where {{value}} is the `number` field of a `Project` resource. */
	External string `json:"external,omitempty"`

	/* Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names */
	Name string `json:"name,omitempty"`

	/* Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ */
	Namespace string `json:"namespace,omitempty"`
}

// AccessContextManagerServicePerimeterResourceSpec defines the desired state of AccessContextManagerServicePerimeterResource
// +kcc:spec:proto=google.identity.accesscontextmanager.v1.ServicePerimeter
type AccessContextManagerServicePerimeterResourceSpec struct {
	PerimeterNameRef ServicePerimeterResourcePerimeterNameRef `json:"perimeterNameRef"`

	/* A GCP resource that is inside of the service perimeter. */
	ResourceRef ServicePerimeterResourceResourceRef `json:"resourceRef"`
}

// AccessContextManagerServicePerimeterResourceStatus defines the config connector machine state of AccessContextManagerServicePerimeterResource
type AccessContextManagerServicePerimeterResourceStatus struct {
	/* Conditions represent the latest available observations of the
	   AccessContextManagerServicePerimeterResource's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* ObservedState is the state of the resource as most recently observed in GCP. */
	// +optional
	ObservedState *AccessContextManagerServicePerimeterResourceObservedState `json:"observedState,omitempty"`
}

// AccessContextManagerServicePerimeterResourceObservedState is the state of the AccessContextManagerServicePerimeterResource resource as most recently observed in GCP.
type AccessContextManagerServicePerimeterResourceObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpaccesscontextmanagerserviceperimeterresource;gcpaccesscontextmanagerserviceperimeterresources
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// +kubebuilder:storageversion
// AccessContextManagerServicePerimeterResource is the Schema for the accesscontextmanager API
// +k8s:openapi-gen=true
type AccessContextManagerServicePerimeterResource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AccessContextManagerServicePerimeterResourceSpec   `json:"spec,omitempty"`
	Status AccessContextManagerServicePerimeterResourceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AccessContextManagerServicePerimeterResourceList contains a list of AccessContextManagerServicePerimeterResource
type AccessContextManagerServicePerimeterResourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AccessContextManagerServicePerimeterResource `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AccessContextManagerServicePerimeterResource{}, &AccessContextManagerServicePerimeterResourceList{})
}
