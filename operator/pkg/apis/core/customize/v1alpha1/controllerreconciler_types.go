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
// config connector controllers in namespaced mode.
type NamespacedControllerReconciler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NamespacedControllerReconcilerSpec   `json:"spec"`
	Status NamespacedControllerReconcilerStatus `json:"status,omitempty"`
}

// NamespacedControllerReconciler is the specification of NamespacedControllerReconciler.
type NamespacedControllerReconcilerSpec struct {
	// RateLimit configures the token bucket rate limit to the kubernetes client used
	// by the manager container of the config connector namespaced controller manager.
	// Please note this rate limit is shared among all the Config Connector resources' requests.
	// If not specified, the default will be Token Bucket with qps 20, burst 30.
	// +optional
	RateLimit *RateLimit `json:"rateLimit,omitempty"`
	// Configures the debug endpoint on the service.
	// +optional
	Pprof *PprofConfig `json:"pprof,omitempty"`
}

type RateLimit struct {
	// The QPS of the token bucket rate limit for all the requests to the kubernetes client.
	// +optional
	QPS int `json:"qps,omitempty"`
	// The burst of the token bucket rate limit for all the requests to the kubernetes client.
	// +optional
	Burst int `json:"burst,omitempty"`
}

type PprofConfig struct {
	// Control if pprof should be turned on and which types should be enabled.
	// +kubebuilder:validation:Enum=none;all
	// +optional
	Support string `json:"support,omitempty"`
	// The port that the pprof server binds to if enabled
	// +optional
	Port int `json:"port,omitempty"`
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

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=controllerreconcilers,scope=Cluster

// ControllerReconciler is the Schema for reconciliation related customization for
// config connector controllers in cluster mode.
type ControllerReconciler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ControllerReconcilerSpec   `json:"spec"`
	Status ControllerReconcilerStatus `json:"status,omitempty"`
}

// ControllerReconcilerSpec is the specification of ControllerReconciler.
type ControllerReconcilerSpec struct {
	// RateLimit configures the token bucket rate limit to the kubernetes client used
	// by the manager container of the config connector controller manager in cluster mode.
	// Please note this rate limit is shared among all the Config Connector resources' requests.
	// If not specified, the default will be Token Bucket with qps 20, burst 30.
	// +optional
	RateLimit *RateLimit `json:"rateLimit,omitempty"`
	// Configures the debug endpoint on the service.
	// +optional
	Pprof *PprofConfig `json:"pprof,omitempty"`
}

// ControllerReconcilerStatus defines the observed state of ControllerReconciler.
type ControllerReconcilerStatus struct {
	addonv1alpha1.CommonStatus `json:",inline"`
}

func (c *ControllerReconciler) SetCommonStatus(s addonv1alpha1.CommonStatus) {
	c.Status.CommonStatus = s
}

// +kubebuilder:object:root=true

// ControllerReconcilerList contains a list of ControllerReconciler.
type ControllerReconcilerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ControllerReconciler `json:"items"`
}

var ValidRateLimitControllers = []string{
	"cnrm-controller-manager",
}

var SupportedPprofControllers = []string{
	"cnrm-controller-manager",
}

func init() {
	SchemeBuilder.Register(
		&NamespacedControllerReconciler{},
		&NamespacedControllerReconcilerList{},
		&ControllerReconciler{},
		&ControllerReconcilerList{},
	)
}
