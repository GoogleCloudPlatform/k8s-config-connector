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
)

// AppTeamSpec defines the desired state of AppTeam
type AppTeamSpec struct {
	Project string `json:"project"`
}

// AppTeamStatus defines the observed state of AppTeam
type AppTeamStatus struct {
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// AppTeam is the Schema for the appteams API
type AppTeam struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AppTeamSpec   `json:"spec,omitempty"`
	Status AppTeamStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// AppTeamList contains a list of AppTeam
type AppTeamList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppTeam `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AppTeam{}, &AppTeamList{})
}
