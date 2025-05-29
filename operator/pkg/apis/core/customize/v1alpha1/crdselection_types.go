// Copyright 2025 Google LLC
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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	addonv1alpha1 "sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/addon/pkg/apis/v1alpha1"
)

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=apigroupinstallations,scope=Cluster,shortName=apigroupinstallation;apigroupinstallations

// APIGroupInstallation enables a specific API group of CRDs in Config Connector.
// The name of the resource must match the API group to enable.
// If any APIGroupInstallation resources exist, only the API groups specified by these resources will be installed.
// If no APIGroupInstallation resources exist, all API groups will be installed (default behavior).
type APIGroupInstallation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec could be empty as the presence of the resource indicates
	// that the API group should be installed.
	Spec   APIGroupInstallationSpec   `json:"spec,omitempty"`
	Status APIGroupInstallationStatus `json:"status,omitempty"`
}

// APIGroupInstallationSpec defines the desired state of APIGroupInstallation.
// Currently empty as the presence of the resource indicates intent.
type APIGroupInstallationSpec struct {
	// TODO: add support for specifying CRDs under this API group?

	// TODO: Add a field to enable installation of dependent API groups?
}

// APIGroupInstallationStatus defines the observed state of APIGroupInstallation.
type APIGroupInstallationStatus struct {
	addonv1alpha1.CommonStatus `json:",inline"`

	// TODO: InstalledCRDs is the count of CRDs that were installed from this API group.
	// InstalledCRDs int `json:"installedCRDs,omitempty"`
}

// +kubebuilder:object:root=true

// APIGroupInstallationList contains a list of APIGroupInstallation.
type APIGroupInstallationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []APIGroupInstallation `json:"items"`
}

func (c *APIGroupInstallation) SetCommonStatus(s addonv1alpha1.CommonStatus) {
	c.Status.CommonStatus = s
}

func init() {
	SchemeBuilder.Register(&APIGroupInstallation{}, &APIGroupInstallationList{})
}
