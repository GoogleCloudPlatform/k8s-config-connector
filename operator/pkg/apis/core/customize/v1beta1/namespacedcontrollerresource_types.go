// Copyright 2023 Google LLC
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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	addonv1alpha1 "sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/addon/pkg/apis/v1alpha1"
)

// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=namespacedcontrollerresources

// NamespacedControllerResource is the Schema for resource customization API for namespaced config connector controllers.
type NamespacedControllerResource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NamespacedControllerResourceSpec   `json:"spec"`
	Status NamespacedControllerResourceStatus `json:"status,omitempty"`
}

// NamespacedControllerResourceSpec is the specification of the resource customization for containers of
// a namespaced config connector controller.
type NamespacedControllerResourceSpec struct {
	// The list of containers whose resource requirements to be customized.
	// Required
	Containers []ContainerResourceSpec `json:"containers"`
}

// NamespacedControllerResourceStatus defines the observed state of NamespacedControllerResource.
type NamespacedControllerResourceStatus struct {
	addonv1alpha1.CommonStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// NamespacedControllerResourceList contains a list of NamespacedControllerResource
type NamespacedControllerResourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NamespacedControllerResource `json:"items"`
}

func (c *NamespacedControllerResource) GetCommonStatus() addonv1alpha1.CommonStatus {
	return c.Status.CommonStatus
}

func (c *NamespacedControllerResource) SetCommonStatus(s addonv1alpha1.CommonStatus) {
	c.Status.CommonStatus = s
}

func init() {
	SchemeBuilder.Register(&NamespacedControllerResource{}, &NamespacedControllerResourceList{})
}
