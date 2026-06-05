/*
Copyright 2020 The Kubernetes Authors.

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
	addonv1alpha1 "sigs.k8s.io/kubebuilder-declarative-pattern/pkg/patterns/addon/pkg/apis/v1alpha1"
)

// SimpleTestSpec defines the desired state of Guestbook
type SimpleTestSpec struct {
	addonv1alpha1.CommonSpec `json:",inline"`
	addonv1alpha1.PatchSpec  `json:",inline"`
}

// SimpleTestStatus defines the observed state of Guestbook
type SimpleTestStatus struct {
	addonv1alpha1.CommonStatus     `json:",inline"`
	addonv1alpha1.StatusConditions `json:",inline"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// SimpleTest is the Schema for the simpletest API
type SimpleTest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SimpleTestSpec   `json:"spec,omitempty"`
	Status SimpleTestStatus `json:"status,omitempty"`
}

var _ addonv1alpha1.CommonObject = &SimpleTest{}
var _ addonv1alpha1.Patchable = &SimpleTest{}

func (o *SimpleTest) ComponentName() string {
	return "simpletest"
}

func (o *SimpleTest) CommonSpec() addonv1alpha1.CommonSpec {
	return o.Spec.CommonSpec
}

func (o *SimpleTest) PatchSpec() addonv1alpha1.PatchSpec {
	return o.Spec.PatchSpec
}

func (o *SimpleTest) GetCommonStatus() addonv1alpha1.CommonStatus {
	return o.Status.CommonStatus
}

func (o *SimpleTest) SetCommonStatus(s addonv1alpha1.CommonStatus) {
	o.Status.CommonStatus = s
}

//+kubebuilder:object:root=true

// SimpleTestList contains a list of Guestbook
type SimpleTestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SimpleTest `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SimpleTest{}, &SimpleTestList{})
}
