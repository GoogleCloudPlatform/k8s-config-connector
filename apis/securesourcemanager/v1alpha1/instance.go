// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	"reflect"

	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

var (
	// SchemeBuilder is used to add go types to the GroupVersionKind scheme.
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}

	// AddToScheme is a global function that registers this API group & version to a scheme
	AddToScheme = SchemeBuilder.AddToScheme

	SecureSourceManagerInstanceGVK = schema.GroupVersionKind{
		Group:   SchemeGroupVersion.Group,
		Version: SchemeGroupVersion.Version,
		Kind:    reflect.TypeOf(SecureSourceManagerInstance{}).Name(),
	}
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpsecuresourcemanagerinstance;gcpsecuresourcemanagerinstances
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=alpha";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// SecureSourceManagerInstance is the Schema for the monitoring API
// +k8s:openapi-gen=true
type SecureSourceManagerInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec SecureSourceManagerInstanceSpec `json:"spec,omitempty"`

	Status SecureSourceManagerInstanceStatus `json:"status,omitempty"`
}

// +kcc:proto=google.cloud.securesourcemanager.v1.Instance
type SecureSourceManagerInstanceSpec struct {
	/* Immutable. The Project that this resource belongs to. */
	ProjectRef refs.ProjectRef `json:"projectRef"`

	/* Immutable. Location of the instance. */
	Location string `json:"location"`

	/* Immutable. Optional. The name of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Labels as key value pairs.
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Immutable. Customer-managed encryption key name, in the format
	//  projects/*/locations/*/keyRings/*/cryptoKeys/*.
	KmsKey *string `json:"kmsKey,omitempty"`
}

// +kcc:proto=google.cloud.securesourcemanager.v1.Instance
type SecureSourceManagerInstanceObservedState struct {
	// // Output only. Create timestamp.
	// CreateTime *string `json:"createTime,omitempty"`

	// // Output only. Update timestamp.
	// UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Current state of the instance.
	State *string `json:"state,omitempty"`

	// Output only. An optional field providing information about the current
	//  instance state.
	StateNote *string `json:"stateNote,omitempty"`

	// Output only. A list of hostnames for this instance.
	HostConfig *Instance_HostConfig `json:"hostConfig,omitempty"`
}

type SecureSourceManagerInstanceStatus struct {
	/* Conditions represent the latest available observations of the
	   SecureSourceManagerInstance's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	/* ObservedState is the latest GCP state. */
	ObservedState *SecureSourceManagerInstanceObservedState `json:"observedState,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SecureSourceManagerInstanceList contains a list of SecureSourceManagerInstance
type SecureSourceManagerInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecureSourceManagerInstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SecureSourceManagerInstance{}, &SecureSourceManagerInstanceList{})
}
