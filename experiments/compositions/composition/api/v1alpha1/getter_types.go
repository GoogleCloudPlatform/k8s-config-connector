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
	"fmt"

	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ResourceRef struct {
	// OPTION 1
	// <Kind>.<group>/<version>/<namespace>/<name>
	// resource: ServiceIdentity.serviceusage.cnrm.cloud.google.com/v1beta1//sqladmin.googleapis.com

	// OPTION 2
	Group    string `json:"group,omitempty"`
	Version  string `json:"version,omitempty"`
	Resource string `json:"resource"`
	Kind     string `json:"kind"`

	// OneOf validation needed for Name and NameSuffix in CRD Definition
	Name       string `json:"name,omitempty"`
	NameSuffix string `json:"nameSuffix,omitempty"`
}

type FieldRef struct {
	Path string `json:"path"`
	As   string `json:"as"`
}

type ValuesFrom struct {
	Name        string      `json:"name"`
	ResourceRef ResourceRef `json:"resourceRef"`
	FieldRef    []FieldRef  `json:"fieldRef"`
}

// GetterSpec defines the desired state of Getter
type GetterSpec struct {
	ValuesFrom []ValuesFrom `json:"valuesFrom,omitempty"`
}

// GetterStatus defines the observed state of Getter
type GetterStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Getter is the Schema for the getters API
type Getter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GetterSpec   `json:"spec,omitempty"`
	Status GetterStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// GetterList contains a list of Getter
type GetterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Getter `json:"items"`
}

// Status helpers
func (g *GetterStatus) ClearCondition(condition ConditionType) {
	meta.RemoveStatusCondition(&g.Conditions, string(condition))
}

// Validation
func (g *Getter) Validate() bool {
	// Several of these validations should br CEL rules on the composition CRD
	// However for now they help me shape the controller.
	g.Status.ClearCondition(ValidationFailed)
	// Validate Expanders
	message := ""
	for index, vf := range g.Spec.ValuesFrom {
		if vf.ResourceRef.Name == "" && vf.ResourceRef.NameSuffix == "" {
			message += fmt.Sprintf(".spec.valuesFrom[%d] requires name or nameSuffix; ", index)
		}
	}
	if message != "" {
		g.Status.Conditions = append(g.Status.Conditions, metav1.Condition{
			LastTransitionTime: metav1.Now(),
			Message:            message,
			Reason:             "GetterValidationFailed",
			Type:               string(ValidationFailed),
			Status:             metav1.ConditionTrue,
		})
		return false
	}
	return true
}

func init() {
	SchemeBuilder.Register(&Getter{}, &GetterList{})
}
