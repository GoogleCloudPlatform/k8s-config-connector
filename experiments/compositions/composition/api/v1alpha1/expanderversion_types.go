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

	semver "github.com/Masterminds/semver/v3"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.
type ExpanderType string

const (
	// ExpanderTypeJob run expander in a job
	ExpanderTypeJob ExpanderType = "job"
	// ExpanderTypeGRPC expect expander service to be present
	ExpanderTypeGRPC ExpanderType = "grpc"
)

type ExpanderConfigGVK struct {
	Group   string `json:"group"`
	Version string `json:"version"`
	Kind    string `json:"kind"`
}

// ExpanderVersionSpec defines the desired state of ExpanderVersion
type ExpanderVersionSpec struct {
	// ImageRegistry is the designated registry for where to pull the named expander image
	ImageRegistry string `json:"imageRegistry,omitempty"`

	// Image if different from removePrefix(expanderversion.name , "composition-")
	Image string `json:"image,omitempty"`

	// ValidVersions is a list of valid versions of the named expander
	ValidVersions []string `json:"validVersions"`

	// Type indicates what sort of expander:
	//   job - job based expander. ephemeral
	//   grpc - grpc service expander. persistent
	// +kubebuilder:validation:Enum=job;grpc
	// +kubebuilder:default=job
	Type ExpanderType `json:"type"`

	// ExpanderConfig GVK
	Config ExpanderConfigGVK `json:"config,omitempty"`
}

// ExpanderVersionStatus defines the observed state of ExpanderVersion
type ExpanderVersionStatus struct {
	VersionMap map[string]string  `json:"versionMap,omitempty"`
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ExpanderVersion is the Schema for the expanderversions API
type ExpanderVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ExpanderVersionSpec   `json:"spec,omitempty"`
	Status ExpanderVersionStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ExpanderVersionList contains a list of ExpanderVersion
type ExpanderVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ExpanderVersion `json:"items"`
}

// Status helpers
func (s *ExpanderVersionStatus) ClearCondition(condition ConditionType) {
	meta.RemoveStatusCondition(&s.Conditions, string(condition))
}

// Validation
func (ev *ExpanderVersion) Validate() bool {
	// Several of these validations should eventually be CEL rules on the composition CRD
	ev.Status.ClearCondition(ValidationFailed)
	// Validate
	message := ""
	if ev.Spec.ImageRegistry == "" && ev.Spec.Type == ExpanderTypeJob {
		message += "spec.imageRegistry required for type=job; "
	}
	if len(ev.Spec.ValidVersions) == 0 {
		message += "spec.validVersions required; "
	}

	invalidVersions := []string{}
	invalidVersionMessage := ""
	for _, r := range ev.Spec.ValidVersions {
		_, err := semver.NewVersion(r)
		if err != nil {
			invalidVersions = append(invalidVersions, r)
			invalidVersionMessage += fmt.Sprintf("%s: %v\n", r, err)
		}
	}
	if invalidVersionMessage != "" {
		message += fmt.Sprintf("invalid versions: %v\n", invalidVersions)
		message += invalidVersionMessage
	}

	if message != "" {
		ev.Status.Conditions = append(ev.Status.Conditions, metav1.Condition{
			LastTransitionTime: metav1.Now(),
			Message:            message,
			Reason:             "ValidationFailed",
			Type:               string(ValidationFailed),
			Status:             metav1.ConditionTrue,
		})
		return false
	}
	return true
}

func init() {
	SchemeBuilder.Register(&ExpanderVersion{}, &ExpanderVersionList{})
}
