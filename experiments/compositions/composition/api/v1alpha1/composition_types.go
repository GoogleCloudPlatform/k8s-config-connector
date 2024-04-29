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

// TODO(barney-s) : Remove proto annotations

// ConditionType defines the type of ManagedConfigSync condition
type ConditionType string

// The valid conditions of Compositions
const (
	Ready ConditionType = "Ready"
	// Error implies the last reconcile attempt failed
	Error ConditionType = "Error"
	// Validation implies the validation failed
	ValidationFailed ConditionType = "ValidationFailed"
	// Waiting - Plan is waiting for values to progress
	Waiting ConditionType = "Waiting"
)

type ResourceRef struct {
	// OPTION 1
	// <Kind>.<group>/<version>/<namespace>/<name>
	// resource: ServiceIdentity.serviceusage.cnrm.cloud.google.com/v1beta1//sqladmin.googleapis.com

	// OPTION 2
	Group      string `json:"group,omitempty" protobuf:"bytes,1,name=group"`
	Version    string `json:"version,omitempty" protobuf:"bytes,2,opt,name=version"`
	Kind       string `json:"kind" protobuf:"bytes,3,name=kind"`
	Name       string `json:"name,omitempty" protobuf:"bytes,4,name=name"`
	NameSuffix string `json:"nameSuffix,omitempty" protobuf:"bytes,2,name=nameSuffix"`
}

type FieldRef struct {
	Path string `json:"path" protobuf:"bytes,1,name=path"`
	As   string `json:"as" protobuf:"bytes,2,name=as"`
}

type ValuesFrom struct {
	Name        string      `json:"name" protobuf:"bytes,1,name=name"`
	ResourceRef ResourceRef `json:"resourceRef" protobuf:"resourceRef,2,name=resourceref"`
	FieldRef    []FieldRef  `json:"fieldRef" protobuf:"fieldRef,3,name=fieldref"`
}

type Expander struct {
	Name string `json:"name,omitempty" protobuf:"bytes,1,name=name"`
	// Type indicates what expander to use
	//   jinja - jinja2 expander
	//   none - No expander
	// +kubebuilder:validation:Enum=jinja2;none
	// +kubebuilder:default=jinja2
	Type string `json:"type" protobuf:"bytes,2,name=name"`
	// +kubebuilder:default=latest
	Version  string `json:"version,omitempty" protobuf:"bytes,3,opt,name=version"`
	Template string `json:"template" protobuf:"bytes,4,name=template"`

	ValuesFrom []ValuesFrom `json:"valuesFrom,omitempty" protobuf:"valuesFrom,5,opt,name=valuesFrom"`
	// NOTE: Tighten the Composition API to include fields that are used in the controller
	//  As we add features we can uncomment these fields

	//ConfigAPIGroup  string `json:"configAPIGroup,omitempty" protobuf:"bytes,3,opt,name=configAPIGroup"`
	//ConfigName      string `json:"configName,omitempty" protobuf:"bytes,4,opt,name=configName"`
	//ConfigNamespace string `json:"configNamespace,omitempty" protobuf:"bytes,5,opt,name=configNamespace"`
	//Image string `json:"image" protobuf:"bytes,6,opt,name=image"`
}

type Sinc struct {
	Name    string `json:"name" protobuf:"bytes,1,opt,name=name"`
	Version string `json:"version" protobuf:"bytes,2,opt,name=version"`

	// NOTE: Tighten the Composition API to include fields that are used in the controller
	//  As we add features we can uncomment these fields

	//ConfigAPIGroup  string `json:"configAPIGroup,omitempty" protobuf:"bytes,3,opt,name=configAPIGroup"`
	//ConfigName      string `json:"configName,omitempty" protobuf:"bytes,4,opt,name=configName"`
	//ConfigNamespace string `json:"configNamespace,omitempty" protobuf:"bytes,5,opt,name=configNamespace"`
	//Image           string `json:"image" protobuf:"bytes,6,opt,name=image"`
}

type NamespaceMode string

const (
	// NamespaceModeNone is when nothing is set, this is the same as Inherit
	NamespaceModeNone NamespaceMode = ""
	// NamespaceModeInherit implies all the objects namespace is replaced with the  input api object's namespace
	NamespaceModeInherit NamespaceMode = "inherit"
	// NamespaceModeExplicit implies the objects in the template must have its namespace set
	NamespaceModeExplicit NamespaceMode = "explicit"
)

// CompositionSpec defines the desired state of Composition
type CompositionSpec struct {
	// NOTE: Tighten the Composition API to include fields that are used in the controller
	//  As we add features we can uncomment these fields
	//Name           string     `json:"name" protobuf:"bytes,1,name=name"`
	//Namespace      string     `json:"namespace" protobuf:"bytes,2,name=namespace"`
	//InputName      string     `json:"inputName,omitempty" protobuf:"bytes,4,name=inputName"`
	//InputNamespace string     `json:"inputNamespace,omitempty" protobuf:"bytes,5,name=inputNamespace"`
	//Sinc      Sinc       `json:"sinc,omitempty" protobuf:"bytes,6,name=sinc"`

	Description string `json:"description,omitempty"`

	// TODO (barney -s) rename to FacadeAPIGroup,facadeAPIGroup

	// Use existing KRM API
	InputAPIGroup string `json:"inputAPIGroup,omitempty"`

	//+kubebuilder:validation:MinItems=1
	Expanders []Expander `json:"expanders"`
	// Namespace mode indicates how compositions set the namespace of the objects from expanders.
	// ""|inherit implies inherit the facade api's namespace. Only namespaced objects are allowed.
	// explicit     implies the objects in the template must have the namespace set.
	// +kubebuilder:validation:Enum=inherit;explicit
	NamespaceMode NamespaceMode `json:"namespaceMode,omitempty"`
}

// CompositionStatus defines the observed state of Composition
type CompositionStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Composition is the Schema for the compositions API
type Composition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CompositionSpec   `json:"spec,omitempty"`
	Status CompositionStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// CompositionList contains a list of Composition
type CompositionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Composition `json:"items"`
}

// Status helpers
func (s *CompositionStatus) ClearCondition(condition ConditionType) {
	meta.RemoveStatusCondition(&s.Conditions, string(condition))
}

// Validation
func (s *Composition) Validate() bool {
	// Several of these validations should br CEL rules on the composition CRD
	// However for now they help me shape the controller.
	s.Status.ClearCondition(ValidationFailed)
	// Validate Expanders
	message := ""
	for expanderIndex, expander := range s.Spec.Expanders {
		if expander.Name == "" {
			message += fmt.Sprintf(".spec.expanders[%d] missing name; ", expanderIndex)
		}
		if expander.ValuesFrom != nil {
			for i, v := range expander.ValuesFrom {
				if v.ResourceRef.Name == "" && v.ResourceRef.NameSuffix == "" {
					message += fmt.Sprintf(".spec.expanders[%d](name:%s).valuesFrom[%d] requires name or nameSuffix; ",
						expanderIndex, expander.Name, i)
				}
			}
		}
	}
	if message != "" {
		s.Status.Conditions = append(s.Status.Conditions, metav1.Condition{
			LastTransitionTime: metav1.Now(),
			Message:            message,
			Reason:             "ExpanderValidationFailed",
			Type:               string(ValidationFailed),
			Status:             metav1.ConditionTrue,
		})
		return false
	}
	return true
}

func init() {
	SchemeBuilder.Register(&Composition{}, &CompositionList{})
}
