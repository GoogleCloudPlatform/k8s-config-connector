// Copyright 2026 Google LLC
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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var SaaSServiceMgmtUnitKindGVK = GroupVersion.WithKind("SaaSServiceMgmtUnitKind")

// SaaSServiceMgmtUnitKindSpec defines the desired state of SaaSServiceMgmtUnitKind
// +kcc:spec:proto=google.cloud.saasplatform.saasservicemgmt.v1beta1.UnitKind
type SaaSServiceMgmtUnitKindSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Location field is immutable"
	// Immutable. The location of this resource.
	// +required
	Location *string `json:"location"`

	// The SaaSServiceMgmtUnitKind name. If not given, the metadata.name will be used.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. A reference to the Release object to use as default for creating
	//  new units of this UnitKind (optional).
	//
	//  If not specified, a new unit must explicitly reference which release to use
	//  for its creation.
	// +optional
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.UnitKind.default_release
	DefaultReleaseRef *SaasServiceMgmtReleaseRef `json:"defaultReleaseRef,omitempty"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Dependencies field is immutable"
	// Optional. Immutable. List of other unit kinds that this release will depend
	//  on. Dependencies will be automatically provisioned if not found.
	//  Maximum 10.
	// +optional
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.UnitKind.dependencies
	Dependencies []Dependency `json:"dependencies,omitempty"`

	// Optional. List of inputVariables for this release that will either be
	//  retrieved from a dependency’s outputVariables, or will be passed on to a
	//  dependency’s inputVariables. Maximum 100.
	// +optional
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.UnitKind.input_variable_mappings
	InputVariableMappings []VariableMapping `json:"inputVariableMappings,omitempty"`

	// Optional. List of outputVariables for this unit kind will be passed to this
	//  unit's outputVariables. Maximum 100.
	// +optional
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.UnitKind.output_variable_mappings
	OutputVariableMappings []VariableMapping `json:"outputVariableMappings,omitempty"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="SaasRef field is immutable"
	// Required. Immutable. A reference to the Saas that defines the product
	//  (managed service) that the producer wants to manage with SaaS Runtime. Part
	//  of the SaaS Runtime common data model. Immutable once set.
	// +required
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.UnitKind.saas
	SaasRef *SaaSServiceMgmtSaaSRef `json:"saasRef"`

	// Optional. The labels on the resource, which can be used for categorization.
	//  similar to Kubernetes resource labels.
	// +optional
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.UnitKind.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. Annotations is an unstructured key-value map stored with a
	//  resource that may be set by external tools to store and retrieve arbitrary
	//  metadata. They are not queryable and should be preserved when modifying
	//  objects.
	//
	//  More info: https://kubernetes.io/docs/user-guide/annotations
	// +optional
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.UnitKind.annotations
	Annotations map[string]string `json:"annotations,omitempty"`
}

// +kcc:proto=google.cloud.saasplatform.saasservicemgmt.v1beta1.Dependency
type Dependency struct {
	// Required. Immutable. The unit kind of the dependency.
	// +required
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.Dependency.unit_kind
	UnitKindRef *SaaSServiceMgmtUnitKindRef `json:"unitKindRef"`

	// Required. An alias for the dependency. Used for input variable mapping.
	// +required
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.Dependency.alias
	Alias *string `json:"alias"`
}

// +kcc:proto=google.cloud.saasplatform.saasservicemgmt.v1beta1.VariableMapping
type VariableMapping struct {
	// Optional. Output variables which will get their values from dependencies
	// +optional
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.VariableMapping.from
	From *FromMapping `json:"from,omitempty"`

	// Optional. Input variables whose values will be passed on to dependencies.
	// +optional
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.VariableMapping.to
	To *ToMapping `json:"to,omitempty"`

	// Required. name of the variable
	// +required
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.VariableMapping.variable
	Variable *string `json:"variable"`
}

// +kcc:proto=google.cloud.saasplatform.saasservicemgmt.v1beta1.FromMapping
type FromMapping struct {
	// Required. Alias of the dependency that the outputVariable will pass its
	//  value to
	// +required
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.FromMapping.dependency
	Dependency *string `json:"dependency"`

	// Required. Name of the outputVariable on the dependency
	// +required
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.FromMapping.output_variable
	OutputVariable *string `json:"outputVariable"`
}

// +kcc:proto=google.cloud.saasplatform.saasservicemgmt.v1beta1.ToMapping
type ToMapping struct {
	// Required. Alias of the dependency that the inputVariable will pass its
	//  value to
	// +required
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.ToMapping.dependency
	Dependency *string `json:"dependency"`

	// Required. Name of the inputVariable on the dependency
	// +required
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.ToMapping.input_variable
	InputVariable *string `json:"inputVariable"`

	// Optional. Tells SaaS Runtime if this mapping should be used during lookup
	//  or not
	// +optional
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.ToMapping.ignore_for_lookup
	IgnoreForLookup *bool `json:"ignoreForLookup,omitempty"`
}

// SaaSServiceMgmtUnitKindStatus defines the config connector machine state of SaaSServiceMgmtUnitKind
type SaaSServiceMgmtUnitKindStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the SaaSServiceMgmtUnitKind resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *SaaSServiceMgmtUnitKindObservedState `json:"observedState,omitempty"`
}

// SaaSServiceMgmtUnitKindObservedState is the state of the SaaSServiceMgmtUnitKind resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.saasplatform.saasservicemgmt.v1beta1.UnitKind
type SaaSServiceMgmtUnitKindObservedState struct {
	// Output only. The unique identifier of the resource. UID is unique in the
	//  time and space for this resource within the scope of the service. It is
	//  typically generated by the server on successful creation of a resource
	//  and must not be changed. UID is used to uniquely identify resources
	//  with resource name reuses. This should be a UUID4.
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.UnitKind.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. An opaque value that uniquely identifies a version or
	//  generation of a resource. It can be used to confirm that the client
	//  and server agree on the ordering of a resource being written.
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.UnitKind.etag
	Etag *string `json:"etag,omitempty"`

	// Output only. The timestamp when the resource was created.
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.UnitKind.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the resource was last updated. Any
	//  change to the resource made by users must refresh this value.
	//  Changes to a resource made by the service should refresh this value.
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.UnitKind.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpsaasservicemgmtunitkind;gcpsaasservicemgmtunitkinds
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// SaaSServiceMgmtUnitKind is the Schema for the SaaSServiceMgmtUnitKind API
// +k8s:openapi-gen=true
type SaaSServiceMgmtUnitKind struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   SaaSServiceMgmtUnitKindSpec   `json:"spec,omitempty"`
	Status SaaSServiceMgmtUnitKindStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// SaaSServiceMgmtUnitKindList contains a list of SaaSServiceMgmtUnitKind
type SaaSServiceMgmtUnitKindList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SaaSServiceMgmtUnitKind `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SaaSServiceMgmtUnitKind{}, &SaaSServiceMgmtUnitKindList{})
}
