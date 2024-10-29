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

type Jinja2 struct {
	Template string `json:"template"`
}

// ConfigReference - For BYO Expanders, we can extend it
type ConfigReference struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace,omitempty"`
}

type ExpanderConfig struct {
	// Built in expanders
	Jinja2 *Jinja2 `json:"jinja2,omitempty"`
	// For BYO Expanders use generic template or ref for external config
	Template  string           `json:"template,omitempty"`
	Reference *ConfigReference `json:"configref,omitempty"`
}

type Expander struct {
	Name string `json:"name,omitempty"`

	// Type indicates what expander to use
	//   jinja - jinja2 expander
	//   ...
	// +kubebuilder:default=jinja2
	Type string `json:"type"`
	// +kubebuilder:default=latest
	Version string `json:"version,omitempty"`

	// TODO (barney-s): Make ConfigReference the only way to specify and dont have any inline expander configs
	//  This would make the UX experience uniform.
	ExpanderConfig `json:""`
}

type Sinc struct {
	Name    string `json:"name"`
	Version string `json:"version"`

	// NOTE: Tighten the Composition API to include fields that are used in the controller
	//  As we add features we can uncomment these fields

	//ConfigAPIGroup  string `json:"configAPIGroup,omitempty"`
	//ConfigName      string `json:"configName,omitempty"`
	//ConfigNamespace string `json:"configNamespace,omitempty"`
	//Image           string `json:"image"`
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

// ReadyOn defines ready condition for a GVK
type ReadyOn struct {
	Group     string `json:"group"`
	Version   string `json:"version,omitempty"`
	Kind      string `json:"kind"`
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	Ready     string `json:"readyIf"`
}

// CompositionSpec defines the desired state of Composition
type CompositionSpec struct {
	// NOTE: Tighten the Composition API to include fields that are used in the controller
	//  As we add features we can uncomment these fields
	//Name           string     `json:"name"`
	//Namespace      string     `json:"namespace"`
	//InputName      string     `json:"inputName,omitempty"`
	//InputNamespace string     `json:"inputNamespace,omitempty"`
	//Sinc      Sinc       `json:"sinc,omitempty"`

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

	// Readiness
	Readiness []ReadyOn `json:"readiness,omitempty"`
}

type ValidationStatus string

const (
	// ValidationStatusUnknown is when it is not validated
	ValidationStatusUnknown ValidationStatus = "unknown"
	// ValidationStatusSuccess is when valdiation succeeds
	ValidationStatusSuccess ValidationStatus = "success"
	// ValidationStatusFailed is when valdiation fails
	ValidationStatusFailed ValidationStatus = "failed"
	// ValidationStatusError is when validation was not called
	ValidationStatusError ValidationStatus = "error"
)

// StageStatus captures the status of a stage
type StageValidationStatus struct {
	ValidationStatus ValidationStatus `json:"validationStatus,omitempty"`
	Reason           string           `json:"reason,omitempty"`
	Message          string           `json:"message,omitempty"`
}

// CompositionStatus defines the observed state of Composition
type CompositionStatus struct {
	Generation int64                            `json:"generation,omitempty"`
	Conditions []metav1.Condition               `json:"conditions,omitempty"`
	Stages     map[string]StageValidationStatus `json:"stages,omitempty"`
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
