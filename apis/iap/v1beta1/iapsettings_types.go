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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
)

var IAPSettingsGVK = GroupVersion.WithKind("IAPSettings")

type Parent struct {
	// Organization-level settings
	OrganizationRef *refs.OrganizationRef `json:"organizationRef,omitempty"`

	// Folder-level settings
	FolderRef *refs.FolderRef `json:"folderRef,omitempty"`

	// Project-level settings
	ProjectRef *refs.ProjectRef `json:"projectRef,omitempty"`

	// Project-wide web service settings
	ProjectWebRef *ProjectWebRef `json:"projectWebRef,omitempty"`

	// Project-wide Compute service settings
	ComputeServiceRef *ComputeServiceRef `json:"computeServiceRef,omitempty"`

	// Project-wide App Engine service settings
	AppEngineRef *AppEngineRef `json:"appEngineRef,omitempty"`
}

type ProjectWebRef struct {
	ProjectRef *refs.ProjectRef `json:"projectRef"`
}

type ComputeServiceRef struct {
	ProjectRef *refs.ProjectRef `json:"projectRef"`
	// Optional. If specified, settings apply to the region
	Region *string `json:"region,omitempty"`
	// Optional. If specified, settings apply to the service
	ServiceRef *computev1beta1.ComputeBackendServiceRef `json:"serviceRef,omitempty"`
}

type AppEngineRef struct {
	ProjectRef     *refs.ProjectRef              `json:"projectRef"`
	ApplicationRef *refs.AppEngineApplicationRef `json:"applicationRef"`
	// Optional. If specified, settings apply to the service
	ServiceRef *refs.AppEngineServiceRef `json:"serviceRef,omitempty"`
	// Optional. If specified, settings apply to the version
	VersionRef *refs.AppEngineVersionRef `json:"versionRef,omitempty"`
}

// IAPSettingsSpec defines the desired state of IAPSettings
// +kcc:proto=google.cloud.iap.v1.IapSettings
type IAPSettingsSpec struct {
	// The IAPSettings name.
	ResourceID *string `json:"resourceID,omitempty"`

	// +kubebuilder:validation:XValidation:rule="(has(self.organizationRef) ? 1 : 0) + (has(self.folderRef) ? 1 : 0) + (has(self.projectRef) ? 1 : 0) + (has(self.projectWebRef) ? 1 : 0) + (has(self.computeServiceRef) ? 1 : 0) + (has(self.appEngineRef) ? 1 : 0) == 1",message="Exactly one parent field must be set"
	Parent `json:",inline"`

	// Top level wrapper for all access related setting in IAP
	// +kcc:proto:field=google.cloud.iap.v1.IapSettings.access_settings
	AccessSettings *AccessSettings `json:"accessSettings,omitempty"`

	// Top level wrapper for all application related settings in IAP
	// +kcc:proto:field=google.cloud.iap.v1.IapSettings.application_settings
	ApplicationSettings *ApplicationSettings `json:"applicationSettings,omitempty"`
}

// IAPSettingsStatus defines the config connector machine state of IAPSettings
type IAPSettingsStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the IAPSettings resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	// NOTYET: there is no output only field
	// ObservedState *IAPSettingsObservedState `json:"observedState,omitempty"`
}

// IAPSettingsObservedState is the state of the IAPSettings resource as most recently observed in GCP.
// +kcc:proto=google.cloud.iap.v1.IapSettings
// NOTYET: there is no output only field
// type IAPSettingsObservedState struct {
// }

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpiapsettings
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"
// +kubebuilder:storageversion

// IAPSettings is the Schema for the IAPSettings API
// +k8s:openapi-gen=true
type IAPSettings struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   IAPSettingsSpec   `json:"spec,omitempty"`
	Status IAPSettingsStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// IAPSettingsList contains a list of IAPSettings
type IAPSettingsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IAPSettings `json:"items"`
}

func init() {
	SchemeBuilder.Register(&IAPSettings{}, &IAPSettingsList{})
}
