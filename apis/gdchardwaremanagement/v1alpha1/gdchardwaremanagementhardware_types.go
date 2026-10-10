// Copyright 2025 Google LLC
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

// API sources for GDCHardwareManagementHardware, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/gdchardwaremanagement/v1alpha/resources.proto
// +kcc:source:service-docs=https://cloud.google.com/distributed-cloud/edge/latest/docs
// +kcc:source:resource-docs=https://docs.cloud.google.com/distributed-cloud/edge/latest/docs/reference/hardware/rest/v1alpha/projects.locations.hardware

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var GDCHardwareManagementHardwareGVK = GroupVersion.WithKind("GDCHardwareManagementHardware")

// GDCHardwareManagementHardwareSpec defines the desired state of GDCHardwareManagementHardware
// +kcc:spec:proto=google.cloud.gdchardwaremanagement.v1alpha.Hardware
// +kcc:required-from-proto
type GDCHardwareManagementHardwareSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The GDCHardwareManagementHardware name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. Display name for this hardware.
	// +kcc:proto:field=google.cloud.gdchardwaremanagement.v1alpha.Hardware.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Labels associated with this hardware as key value pairs.
	//  For more information about labels, see [Create and manage
	//  labels](https://cloud.google.com/resource-manager/docs/creating-managing-labels).
	// +kcc:proto:field=google.cloud.gdchardwaremanagement.v1alpha.Hardware.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. Name of the order that this hardware belongs to.
	//  Format: `projects/{project}/locations/{location}/orders/{order}`
	// +kcc:proto:field=google.cloud.gdchardwaremanagement.v1alpha.Hardware.order
	// +required
	Order *string `json:"order,omitempty"`

	// Required. Name for the site that this hardware belongs to.
	//  Format: `projects/{project}/locations/{location}/sites/{site}`
	// +kcc:proto:field=google.cloud.gdchardwaremanagement.v1alpha.Hardware.site
	// +required
	Site *string `json:"site,omitempty"`

	// Required. Configuration for this hardware.
	// +kcc:proto:field=google.cloud.gdchardwaremanagement.v1alpha.Hardware.config
	// +required
	Config *HardwareConfig `json:"config,omitempty"`

	// Optional. Physical properties of this hardware.
	// +kcc:proto:field=google.cloud.gdchardwaremanagement.v1alpha.Hardware.physical_info
	PhysicalInfo *HardwarePhysicalInfo `json:"physicalInfo,omitempty"`

	// Optional. Information for installation of this hardware.
	// +kcc:proto:field=google.cloud.gdchardwaremanagement.v1alpha.Hardware.installation_info
	InstallationInfo *HardwareInstallationInfo `json:"installationInfo,omitempty"`

	// Required. Name for the zone that this hardware belongs to.
	//  Format: `projects/{project}/locations/{location}/zones/{zone}`
	// +kcc:proto:field=google.cloud.gdchardwaremanagement.v1alpha.Hardware.zone
	// +required
	Zone *string `json:"zone,omitempty"`

	// Optional. Requested installation date for this hardware. If not specified,
	//  this is auto-populated from the order's fulfillment_time upon submission or
	//  from the HardwareGroup's requested_installation_date upon order acceptance.
	// +kcc:proto:field=google.cloud.gdchardwaremanagement.v1alpha.Hardware.requested_installation_date
	RequestedInstallationDate *Date `json:"requestedInstallationDate,omitempty"`
}

// GDCHardwareManagementHardwareStatus defines the config connector machine state of GDCHardwareManagementHardware
type GDCHardwareManagementHardwareStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the GDCHardwareManagementHardware resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *GDCHardwareManagementHardwareObservedState `json:"observedState,omitempty"`
}

// GDCHardwareManagementHardwareObservedState is the state of the GDCHardwareManagementHardware resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.gdchardwaremanagement.v1alpha.Hardware
type GDCHardwareManagementHardwareObservedState struct {
	// Output only. Time when this hardware was created.
	// +kcc:proto:field=google.cloud.gdchardwaremanagement.v1alpha.Hardware.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time when this hardware was last updated.
	// +kcc:proto:field=google.cloud.gdchardwaremanagement.v1alpha.Hardware.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Name for the hardware group that this hardware belongs to.
	//  Format:
	//  `projects/{project}/locations/{location}/orders/{order}/hardwareGroups/{hardware_group}`
	// +kcc:proto:field=google.cloud.gdchardwaremanagement.v1alpha.Hardware.hardware_group
	HardwareGroup *string `json:"hardwareGroup,omitempty"`

	// Output only. Current state for this hardware.
	// +kcc:proto:field=google.cloud.gdchardwaremanagement.v1alpha.Hardware.state
	State *string `json:"state,omitempty"`

	// Output only. Link to the Customer Intake Questionnaire (CIQ) sheet for this
	//  Hardware.
	// +kcc:proto:field=google.cloud.gdchardwaremanagement.v1alpha.Hardware.ciq_uri
	CiqURI *string `json:"ciqURI,omitempty"`

	// Output only. Estimated installation date for this hardware.
	// +kcc:proto:field=google.cloud.gdchardwaremanagement.v1alpha.Hardware.estimated_installation_date
	EstimatedInstallationDate *Date `json:"estimatedInstallationDate,omitempty"`

	// Output only. Actual installation date for this hardware. Filled in by
	//  Google.
	// +kcc:proto:field=google.cloud.gdchardwaremanagement.v1alpha.Hardware.actual_installation_date
	ActualInstallationDate *Date `json:"actualInstallationDate,omitempty"`

	// Output only. Per machine asset information needed for turnup.
	// +kcc:proto:field=google.cloud.gdchardwaremanagement.v1alpha.Hardware.machine_infos
	MachineInfos []Hardware_MachineInfoObservedState `json:"machineInfos,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpgdchardwaremanagementhardware;gcpgdchardwaremanagementhardwares
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// GDCHardwareManagementHardware is the Schema for the GDCHardwareManagementHardware API
// +k8s:openapi-gen=true
type GDCHardwareManagementHardware struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   GDCHardwareManagementHardwareSpec   `json:"spec,omitempty"`
	Status GDCHardwareManagementHardwareStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// GDCHardwareManagementHardwareList contains a list of GDCHardwareManagementHardware
type GDCHardwareManagementHardwareList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GDCHardwareManagementHardware `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GDCHardwareManagementHardware{}, &GDCHardwareManagementHardwareList{})
}
