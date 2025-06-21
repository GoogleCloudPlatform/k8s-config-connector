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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:scope=Namespaced

// GithubRepo is the Schema for the GithubRepos API
type GithubRepo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GithubRepoSpec   `json:"spec,omitempty"`
	Status GithubRepoStatus `json:"status,omitempty"`
}

// GithubRepoSpec defines the desired state of GithubRepo
type GithubRepoSpec struct {
	Org  string `json:"org,omitempty"`
	Repo string `json:"repo,omitempty"`

	ProtectedTags []ProtectedTag `json:"protectedTags,omitempty"`
}

type ProtectedTag struct {
	Pattern string `json:"pattern,omitempty"`
}

type GithubRepoStatus struct {
}

//+kubebuilder:object:root=true

// GithubRepoList contains a list of GithubRepo
type GithubRepoList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GithubRepo `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GithubRepo{}, &GithubRepoList{})
}
