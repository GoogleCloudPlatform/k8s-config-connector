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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type InstanceConfigRef struct {
	/* The external name of the referenced resource */
	External string `json:"external,omitempty"`
	/* Kind of the referent. */
	Kind string `json:"kind,omitempty"`
	/* Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names */
	Name string `json:"name,omitempty"`
	/* Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ */
	Namespace string `json:"namespace,omitempty"`
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// InstanceConfigSpec defines the desired state of InstanceConfig
type InstanceConfigSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of InstanceConfig. Edit instanceconfig_types.go to remove/update
	// Foo string `json:"foo,omitempty"`
}

// InstanceConfigStatus defines the observed state of InstanceConfig
type InstanceConfigStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// InstanceConfig is the Schema for the instanceconfigs API
type InstanceConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   InstanceConfigSpec   `json:"spec,omitempty"`
	Status InstanceConfigStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// InstanceConfigList contains a list of InstanceConfig
type InstanceConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []InstanceConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&InstanceConfig{}, &InstanceConfigList{})
}
