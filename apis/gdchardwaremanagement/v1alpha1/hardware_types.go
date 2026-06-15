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

var GDCHardwareManagementHardwareGVK = GroupVersion.WithKind("GDCHardwareManagementHardware")

type GDCHardwareManagementOrderRef struct {
	// A reference to an externally managed GDCHardwareManagementOrder resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/orders/{{orderID}}".
	External string `json:"external,omitempty"`

	// The name of a GDCHardwareManagementOrder resource.
	Name string `json:"name,omitempty"`

	// The namespace of a GDCHardwareManagementOrder resource.
	Namespace string `json:"namespace,omitempty"`
}

type GDCHardwareManagementSiteRef struct {
	// A reference to an externally managed GDCHardwareManagementSite resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/sites/{{siteID}}".
	External string `json:"external,omitempty"`

	// The name of a GDCHardwareManagementSite resource.
	Name string `json:"name,omitempty"`

	// The namespace of a GDCHardwareManagementSite resource.
	Namespace string `json:"namespace,omitempty"`
}

type GDCHardwareManagementZoneRef struct {
	// A reference to an externally managed GDCHardwareManagementZone resource.
	// Should be in the format "projects/{{projectID}}/locations/{{location}}/zones/{{zoneID}}".
	External string `json:"external,omitempty"`

	// The name of a GDCHardwareManagementZone resource.
	Name string `json:"name,omitempty"`

	// The namespace of a GDCHardwareManagementZone resource.
	Namespace string `json:"namespace,omitempty"`
}

// GDCHardwareManagementHardwareSpec defines the desired state of GDCHardwareManagementHardware
// +kcc:spec:proto=google.cloud.gdchardwaremanagement.v1alpha.Hardware
type GDCHardwareManagementHardwareSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The GDCHardwareManagementHardware name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Optional. Display name for this hardware.
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. Labels associated with this hardware as key value pairs.
	//  For more information about labels, see [Create and manage
	//  labels](https://cloud.google.com/resource-manager/docs/creating-managing-labels).
	Labels map[string]string `json:"labels,omitempty"`

	// Required. Reference to the order that this hardware belongs to.
	OrderRef *GDCHardwareManagementOrderRef `json:"orderRef,omitempty"`

	// Required. Reference to the site that this hardware belongs to.
	SiteRef *GDCHardwareManagementSiteRef `json:"siteRef,omitempty"`

	// Required. Configuration for this hardware.
	Config *HardwareConfig `json:"config,omitempty"`

	// Optional. Physical properties of this hardware.
	PhysicalInfo *HardwarePhysicalInfo `json:"physicalInfo,omitempty"`

	// Optional. Information for installation of this hardware.
	InstallationInfo *HardwareInstallationInfo `json:"installationInfo,omitempty"`

	// Required. Reference to the zone that this hardware belongs to.
	ZoneRef *GDCHardwareManagementZoneRef `json:"zoneRef,omitempty"`

	// Optional. Requested installation date for this hardware. If not specified,
	//  this is auto-populated from the order's fulfillment_time upon submission or
	//  from the HardwareGroup's requested_installation_date upon order acceptance.
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
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Time when this hardware was last updated.
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Name for the hardware group that this hardware belongs to.
	HardwareGroup *string `json:"hardwareGroup,omitempty"`

	// Output only. Current state for this hardware.
	State *string `json:"state,omitempty"`

	// Output only. Link to the Customer Intake Questionnaire (CIQ) sheet for this
	//  Hardware.
	CiqURI *string `json:"ciqURI,omitempty"`

	// Output only. Estimated installation date for this hardware.
	EstimatedInstallationDate *Date `json:"estimatedInstallationDate,omitempty"`

	// Output only. Actual installation date for this hardware. Filled in by
	//  Google.
	ActualInstallationDate *Date `json:"actualInstallationDate,omitempty"`

	// Output only. Per machine asset information needed for turnup.
	MachineInfos []Hardware_MachineInfoObservedState `json:"machineInfos,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpgdchardwaremanagementhardware;gcpgdchardwaremanagementhardwares
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
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
