/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

type FileContent struct {
	FileName string               `json:"name"`
	Content  runtime.RawExtension `json:"content,omitempty"`
	Template string               `json:"template,omitempty"`
}

// HelmConfigurationSpec defines the desired state of HelmConfiguration
type HelmConfigurationSpec struct {
	// Chart https://helm.sh/docs/topics/charts/#the-chartyaml-file
	Chart         runtime.RawExtension `json:"chart"`
	DefaultValues runtime.RawExtension `json:"defaultValues,omitempty"`
	Templates     []FileContent        `json:"templates"`
	CRDs          []FileContent        `json:"crds,omitempty"`
}

// HelmConfigurationStatus defines the observed state of HelmConfiguration
type HelmConfigurationStatus struct {
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// HelmConfiguration is the Schema for the helmconfigurations API
type HelmConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HelmConfigurationSpec   `json:"spec,omitempty"`
	Status HelmConfigurationStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// HelmConfigurationList contains a list of HelmConfiguration
type HelmConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HelmConfiguration `json:"items"`
}

func init() {
	SchemeBuilder.Register(&HelmConfiguration{}, &HelmConfigurationList{})
}
