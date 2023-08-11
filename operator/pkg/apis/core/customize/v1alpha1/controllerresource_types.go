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

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	addonv1alpha1 "sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/addon/pkg/apis/v1alpha1"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=controllerresources,scope=Cluster

// ControllerResource is the Schema for resource customization API for config connector controllers.
type ControllerResource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ControllerResourceSpec   `json:"spec"`
	Status ControllerResourceStatus `json:"status,omitempty"`
}

// ControllerResourceSpec is the specification of the resource customization for containers of
// a config connector controller.
type ControllerResourceSpec struct {
	// The list of containers whose resource requirements to be customized.
	// Required
	Containers []ContainerResourceSpec `json:"containers"`
	// The number of desired replicas of the config connector controller.
	// This field takes effect only if the controller name is "cnrm-webhook-manager".
	// +optional
	Replicas *int64 `json:"replicas,omitempty"`
}

// ContainerResourceSpec is the specification of the resource customization for a container of
// a config connector controller.
type ContainerResourceSpec struct {
	// The name of the container whose resource requirements will be customized.
	// +kubebuilder:validation:Enum=manager;webhook;deletiondefender;prom-to-sd;recorder
	// Required
	Name string `json:"name"`
	// Resources specifies the resource customization of this container.
	// Required
	Resources corev1.ResourceRequirements `json:"resources"`
}

// ControllerResourceStatus defines the observed state of ControllerResource.
type ControllerResourceStatus struct {
	addonv1alpha1.CommonStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// ControllerResourceList contains a list of ControllerResource
type ControllerResourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ControllerResource `json:"items"`
}

func (c *ControllerResource) SetCommonStatus(s addonv1alpha1.CommonStatus) {
	c.Status.CommonStatus = s
}

func init() {
	SchemeBuilder.Register(&ControllerResource{}, &ControllerResourceList{})
}
