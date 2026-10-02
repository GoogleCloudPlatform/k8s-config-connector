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

var SaaSServiceMgmtUnitGVK = GroupVersion.WithKind("SaaSServiceMgmtUnit")

// SaaSServiceMgmtUnitSpec defines the desired state of SaaSServiceMgmtUnit
// +kcc:spec:proto=google.cloud.saasplatform.saasservicemgmt.v1beta1.Unit
type SaaSServiceMgmtUnitSpec struct {
	// The project that this resource belongs to.
	// +required
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Location field is immutable"
	// Immutable. The location of this resource.
	// +required
	Location *string `json:"location"`

	// The SaaSServiceMgmtUnit name. If not given, the metadata.name will be used.
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="UnitKind field is immutable"
	// Optional. Reference to the UnitKind this Unit belongs to. Immutable once set.
	// +optional
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.Unit.unit_kind
	UnitKind *string `json:"unitKind,omitempty"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="Tenant field is immutable"
	// Optional. Reference to the Saas Tenant resource this unit belongs to. This
	//  for example informs the maintenance policies to use for scheduling future
	//  updates on a unit. (optional and immutable once created)
	// +optional
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.Unit.tenant
	Tenant *string `json:"tenant,omitempty"`

	// Optional. Captures requested directives for performing future maintenance
	//  on the unit. This includes a request for the unit to skip maintenance for a
	//  period of time and remain pinned to its current release as well as controls
	//  for postponing maintenance scheduled in future.
	// +optional
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.Unit.maintenance
	Maintenance *Unit_MaintenanceSettings `json:"maintenance,omitempty"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="ManagementMode field is immutable"
	// Optional. Immutable. Indicates whether the Unit life cycle is controlled
	//  by the user or by the system.
	//  Immutable once created.
	// +optional
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.Unit.management_mode
	ManagementMode *string `json:"managementMode,omitempty"`
}

// SaaSServiceMgmtUnitStatus defines the config connector machine state of SaaSServiceMgmtUnit
type SaaSServiceMgmtUnitStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the SaaSServiceMgmtUnit resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *SaaSServiceMgmtUnitObservedState `json:"observedState,omitempty"`
}

// SaaSServiceMgmtUnitObservedState is the state of the SaaSServiceMgmtUnit resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.saasplatform.saasservicemgmt.v1beta1.Unit
type SaaSServiceMgmtUnitObservedState struct {
	// Optional. Output only. The current Release object for this Unit.
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.Unit.release
	Release *string `json:"release,omitempty"`

	// Optional. Output only. List of concurrent UnitOperations that are operating
	//  on this Unit.
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.Unit.ongoing_operations
	OngoingOperations []string `json:"ongoingOperations,omitempty"`

	// Optional. Output only. List of pending (wait to be executed) UnitOperations
	//  for this unit.
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.Unit.pending_operations
	PendingOperations []string `json:"pendingOperations,omitempty"`

	// Optional. Output only. List of scheduled UnitOperations for this unit.
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.Unit.scheduled_operations
	ScheduledOperations []string `json:"scheduledOperations,omitempty"`

	// Optional. Output only. List of Units that depend on this unit. Unit can
	//  only be deprovisioned if this list is empty. Maximum 1000.
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.Unit.dependents
	Dependents []UnitDependencyObservedState `json:"dependents,omitempty"`

	// Optional. Output only. Set of dependencies for this unit. Maximum 10.
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.Unit.dependencies
	Dependencies []UnitDependencyObservedState `json:"dependencies,omitempty"`

	// Optional. Output only. Indicates the current input variables deployed by
	//  the unit
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.Unit.input_variables
	InputVariables []UnitVariable `json:"inputVariables,omitempty"`

	// Optional. Output only. Set of key/value pairs corresponding to output
	//  variables from execution of actuation templates. The variables are declared
	//  in actuation configs (e.g in helm chart or terraform) and the values are
	//  fetched and returned by the actuation engine upon completion of execution.
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.Unit.output_variables
	OutputVariables []UnitVariable `json:"outputVariables,omitempty"`

	// Optional. Output only. Current lifecycle state of the resource (e.g. if
	//  it's being created or ready to use).
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.Unit.state
	State *string `json:"state,omitempty"`

	// Optional. Output only. A set of conditions which indicate the various
	//  conditions this resource can have.
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.Unit.conditions
	Conditions []UnitCondition `json:"conditions,omitempty"`

	// Optional. Output only. Indicates the system managed state of the unit.
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.Unit.system_managed_state
	SystemManagedState *string `json:"systemManagedState,omitempty"`

	// Optional. Output only. If set, indicates the time when the system will
	//  start removing the unit.
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.Unit.system_cleanup_at
	SystemCleanupAt *string `json:"systemCleanupAt,omitempty"`

	// Output only. The unique identifier of the resource. UID is unique in the
	//  time and space for this resource within the scope of the service. It is
	//  typically generated by the server on successful creation of a resource
	//  and must not be changed. UID is used to uniquely identify resources
	//  with resource name reuses. This should be a UUID4.
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.Unit.uid
	Uid *string `json:"uid,omitempty"`

	// Output only. An opaque value that uniquely identifies a version or
	//  generation of a resource. It can be used to confirm that the client
	//  and server agree on the ordering of a resource being written.
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.Unit.etag
	Etag *string `json:"etag,omitempty"`

	// Output only. The timestamp when the resource was created.
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.Unit.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the resource was last updated. Any
	//  change to the resource made by users must refresh this value.
	//  Changes to a resource made by the service should refresh this value.
	// +kcc:proto:field=google.cloud.saasplatform.saasservicemgmt.v1beta1.Unit.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpsaasservicemgmtunit;gcpsaasservicemgmtunits
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// SaaSServiceMgmtUnit is the Schema for the SaaSServiceMgmtUnit API
// +k8s:openapi-gen=true
type SaaSServiceMgmtUnit struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   SaaSServiceMgmtUnitSpec   `json:"spec,omitempty"`
	Status SaaSServiceMgmtUnitStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// SaaSServiceMgmtUnitList contains a list of SaaSServiceMgmtUnit
type SaaSServiceMgmtUnitList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SaaSServiceMgmtUnit `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SaaSServiceMgmtUnit{}, &SaaSServiceMgmtUnitList{})
}
