// Copyright 2024 Google LLC
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
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=namespacedcontrollerreconcilers

// NamespacedControllerReconciler is the Schema for reconciliation related customization for
// namespaced config connector controllers.
type NamespacedControllerReconciler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NamespacedControllerReconcilerSpec   `json:"spec"`
	Status NamespacedControllerReconcilerStatus `json:"status,omitempty"`
}

// NamespacedControllerReconciler is the specification of NamespacedControllerReconciler.
type NamespacedControllerReconcilerSpec struct {
	// The list of containers whose reconciliation related parameters to be customized.
	// +optional
	Containers []ContainerReconcilerSpec `json:"containers,omitempty"`
}

// ContainerReconcilerSpec is the specification of the reconciliation related customization for
// a container of a config connector controller.
type ContainerReconcilerSpec struct {
	// The name of the container.
	// +kubebuilder:validation:Enum=manager
	// Required
	Name string `json:"name"`
	// RateLimit describes the token bucket rate limit to the kubernetes client.
	// Please note this rate limit is shared among all the Config Connector resources' requests.
	// If not specified, the default will be Token Bucket with qps 20, burst 30.
	// +optional
	RateLimit *RateLimit `json:"rateLimit,omitempty"`
}

type RateLimit struct {
	// The QPS of the token bucket rate limit for all the requests to the kubernetes client.
	// +optional
	QPS int `json:"qps,omitempty"`
	// The burst of the token bucket rate limit for all the requests to the kubernetes client.
	// +optional
	Burst int `json:"burst,omitempty"`
}

// NamespacedControllerReconcilerStatus defines the observed state of NamespacedControllerReconciler.
type NamespacedControllerReconcilerStatus struct {
	addonv1alpha1.CommonStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// NamespacedControllerReconcilerList contains a list of NamespacedControllerReconciler.
type NamespacedControllerReconcilerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NamespacedControllerReconciler `json:"items"`
}

func (c *NamespacedControllerReconciler) SetCommonStatus(s addonv1alpha1.CommonStatus) {
	c.Status.CommonStatus = s
}

func init() {
	SchemeBuilder.Register(&NamespacedControllerReconciler{}, &NamespacedControllerReconcilerList{})
}
