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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	refv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
)

var APIQuotaAdjusterSettingsGVK = GroupVersion.WithKind("APIQuotaAdjusterSettings")

type AdjusterSettingsParent struct {
	// +required
	ProjectRef *refv1beta1.ProjectRef `json:"projectRef"`
}

// APIQuotaAdjusterSettingsSpec defines the desired state of APIQuotaAdjusterSettings
// +kcc:spec:proto=google.api.cloudquotas.v1beta.QuotaAdjusterSettings
type APIQuotaAdjusterSettingsSpec struct {
	AdjusterSettingsParent `json:",inline"`
	// The APIQuotaAdjusterSettings name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Required. The configured value of the enablement at the given resource.
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaAdjusterSettings.enablement
	//+required
	Enablement *string `json:"enablement,omitempty"`
}

// APIQuotaAdjusterSettingsStatus defines the config connector machine state of APIQuotaAdjusterSettings
type APIQuotaAdjusterSettingsStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the APIQuotaAdjusterSettings resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *APIQuotaAdjusterSettingsObservedState `json:"observedState,omitempty"`
}

// APIQuotaAdjusterSettingsObservedState is the state of the APIQuotaAdjusterSettings resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.api.cloudquotas.v1beta.QuotaAdjusterSettings
type APIQuotaAdjusterSettingsObservedState struct {
	// Output only. The timestamp when the QuotaAdjusterSettings was last updated.
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaAdjusterSettings.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
	// Optional. The current etag of the QuotaAdjusterSettings. If an etag is
	//  provided on update and does not match the current server's etag of the
	//  QuotaAdjusterSettings, the request will be blocked and an ABORTED error
	//  will be returned. See https://google.aip.dev/134#etags for more details on
	//  etags.
	// +kcc:proto:field=google.api.cloudquotas.v1beta.QuotaAdjusterSettings.etag
	Etag *string `json:"etag,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TODO(user): make sure the pluralizaiton below is correct
// +kubebuilder:resource:categories=gcp,shortName=gcpapiquotaadjustersetting;gcpapiquotaadjustersettings
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// APIQuotaAdjusterSettings is the Schema for the APIQuotaAdjusterSettings API
// +k8s:openapi-gen=true
type APIQuotaAdjusterSettings struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   APIQuotaAdjusterSettingsSpec   `json:"spec,omitempty"`
	Status APIQuotaAdjusterSettingsStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// APIQuotaAdjusterSettingsList contains a list of APIQuotaAdjusterSettings
type APIQuotaAdjusterSettingsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []APIQuotaAdjusterSettings `json:"items"`
}

func init() {
	SchemeBuilder.Register(&APIQuotaAdjusterSettings{}, &APIQuotaAdjusterSettingsList{})
}
