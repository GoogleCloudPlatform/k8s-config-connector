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
	"k8s.io/apimachinery/pkg/types"
)

type Stage struct {
	Manifest string `json:"manifest,omitempty"`
	Values   string `json:"values,omitempty"`
}

// PlanSpec defines the desired state of Plan
type PlanSpec struct {
	Stages map[string]Stage `json:"stages,omitempty"`
}

type HealthType string

const (
	HEALTHY   HealthType = "Healthy"
	UNHEALTHY HealthType = "Unhealthy"
)

type ResourceStatus struct {
	Group     string     `json:"group,omitempty"`
	Version   string     `json:"version,omitempty"`
	Kind      string     `json:"kind"`
	Namespace string     `json:"namespace,omitempty"`
	Name      string     `json:"name,omitempty"`
	Status    string     `json:"status,omitempty"`
	Health    HealthType `json:"health"`
}

// StageStatus captures the status of a stage
type StageStatus struct {
	ResourceCount int              `json:"resourceCount"`
	AppliedCount  int              `json:"appliedCount,omitempty"`
	LastApplied   []ResourceStatus `json:"lastApplied,omitempty"`
}

// PlanStatus defines the observed state of Plan
type PlanStatus struct {
	// Facade's generation last successfully reconciled
	InputGeneration int64 `json:"inputGeneration"`
	// Composition generation last successfully reconciled
	CompositionGeneration int64 `json:"compositionGeneration"`
	// Composition UID
	CompositionUID types.UID `json:"compositionUID,omitempty"`

	// Plan generation we last successfully reconciled
	Generation int64                   `json:"generation,omitempty"`
	Conditions []metav1.Condition      `json:"conditions,omitempty"`
	Stages     map[string]*StageStatus `json:"stages,omitempty"`
	LastPruned []ResourceStatus        `json:"lastPruned,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Plan is the Schema for the plans API
type Plan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PlanSpec   `json:"spec,omitempty"`
	Status PlanStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// PlanList contains a list of Plan
type PlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Plan `json:"items"`
}

// Status helpers
func (s *PlanStatus) ClearCondition(condition ConditionType) {
	meta.RemoveStatusCondition(&s.Conditions, string(condition))
}

func (s *PlanStatus) AppendCondition(t ConditionType, sts metav1.ConditionStatus, m, r string) {
	s.Conditions = append(s.Conditions, metav1.Condition{
		LastTransitionTime: metav1.Now(),
		Message:            m,
		Reason:             r,
		Type:               string(t),
		Status:             sts,
	})
}

func (s *PlanStatus) AppendWaitingCondition(e, m, r string) {
	message := fmt.Sprintf("Expander: %s, Message: %s", e, m)
	s.AppendCondition(Waiting, metav1.ConditionTrue, message, r)
}

func (s *PlanStatus) AppendErrorCondition(e, m, r string) {
	message := fmt.Sprintf("Expander: %s, Message: %s", e, m)
	s.AppendCondition(Error, metav1.ConditionTrue, message, r)
}

func init() {
	SchemeBuilder.Register(&Plan{}, &PlanList{})
}
