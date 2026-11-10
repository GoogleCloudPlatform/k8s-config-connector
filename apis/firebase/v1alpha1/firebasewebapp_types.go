// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var FirebaseWebAppGVK = GroupVersion.WithKind("FirebaseWebApp")

type FirebaseWebAppSpec struct {
	/* The globally unique, Google-assigned identifier (UID) for the Firebase API key associated with the WebApp.
	If apiKeyId is not set during creation, then Firebase automatically associates an apiKeyId with the WebApp.
	This auto-associated key may be an existing valid key or, if no valid key exists, a new one will be provisioned. */
	// +optional
	ApiKeyId *string `json:"apiKeyId,omitempty"`

	// +optional
	DeletionPolicy *string `json:"deletionPolicy,omitempty"`

	/* The user-assigned display name of the App. */
	// +required
	DisplayName string `json:"displayName"`

	/* Immutable. */
	// +optional
	Project *string `json:"project,omitempty"`

	/* Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

type FirebaseWebAppStatus struct {
	/* Conditions represent the latest available observation of the
	   resource's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	/* A unique specifier for the FirebaseWebApp resource in GCP. */
	// +optional
	ExternalRef *string `json:"externalRef,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* The globally unique, Firebase-assigned identifier of the App.
	This identifier should be treated as an opaque token, as the data format is not specified. */
	// +optional
	AppId *string `json:"appId,omitempty"`

	/* The URLs where the 'WebApp' is hosted. */
	// +optional
	AppUrls []string `json:"appUrls,omitempty"`

	/* The fully qualified resource name of the App, for example:
	projects/projectId/webApps/appId. */
	// +optional
	Name *string `json:"name,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpfirebasewebapp;gcpfirebasewebapps
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",type="date",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",description="The last transition time for the value in 'Status'"

// FirebaseWebApp is the Schema for the firebase API
// +k8s:openapi-gen=true
type FirebaseWebApp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FirebaseWebAppSpec   `json:"spec"`
	Status FirebaseWebAppStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FirebaseWebAppList contains a list of FirebaseWebApp
type FirebaseWebAppList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FirebaseWebApp `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FirebaseWebApp{}, &FirebaseWebAppList{})
}
