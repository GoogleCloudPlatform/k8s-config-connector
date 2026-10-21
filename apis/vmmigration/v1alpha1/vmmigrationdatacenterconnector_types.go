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

// API sources for VMMigrationDatacenterConnector, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/1765b559c42386788ff0c6412491277b4791107a/google/cloud/vmmigration/v1/vmmigration.proto
// +kcc:source:service-docs=https://cloud.google.com/migrate/virtual-machines
// +kcc:guess=source-link reason=verify-resource-docs-link
// +kcc:source:resource-docs=https://docs.cloud.google.com/migrate/virtual-machines/docs/reference/rest/v1/projects.locations.sources.datacenterConnectors

package v1alpha1

import (
	common "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var VMMigrationDatacenterConnectorGVK = GroupVersion.WithKind("VMMigrationDatacenterConnector")

// VMMigrationDatacenterConnectorSpec defines the desired state of VMMigrationDatacenterConnector
// +kcc:spec:proto=google.cloud.vmmigration.v1.DatacenterConnector
// +kcc:required-from-proto
type VMMigrationDatacenterConnectorSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The VMMigrationDatacenterConnector name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Immutable. A unique key for this connector. This key is internal to the OVA
	//  connector and is supplied with its creation during the registration process
	//  and can not be modified.
	// +kcc:proto:field=google.cloud.vmmigration.v1.DatacenterConnector.registration_id
	RegistrationID *string `json:"registrationID,omitempty"`

	// The service account to use in the connector when communicating with the
	//  cloud.
	// +kcc:proto:field=google.cloud.vmmigration.v1.DatacenterConnector.service_account
	ServiceAccount *string `json:"serviceAccount,omitempty"`

	// The version running in the DatacenterConnector. This is supplied by the OVA
	//  connector during the registration process and can not be modified.
	// +kcc:proto:field=google.cloud.vmmigration.v1.DatacenterConnector.version
	Version *string `json:"version,omitempty"`
}

// VMMigrationDatacenterConnectorStatus defines the config connector machine state of VMMigrationDatacenterConnector
type VMMigrationDatacenterConnectorStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the VMMigrationDatacenterConnector resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *VMMigrationDatacenterConnectorObservedState `json:"observedState,omitempty"`
}

// VMMigrationDatacenterConnectorObservedState is the state of the VMMigrationDatacenterConnector resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.vmmigration.v1.DatacenterConnector
type VMMigrationDatacenterConnectorObservedState struct {
	// Output only. The time the connector was created (as an API call, not when
	//  it was actually installed).
	// +kcc:proto:field=google.cloud.vmmigration.v1.DatacenterConnector.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The last time the connector was updated with an API call.
	// +kcc:proto:field=google.cloud.vmmigration.v1.DatacenterConnector.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The communication channel between the datacenter connector and
	//  Google Cloud.
	// +kcc:proto:field=google.cloud.vmmigration.v1.DatacenterConnector.bucket
	Bucket *string `json:"bucket,omitempty"`

	// Output only. State of the DatacenterConnector, as determined by the health
	//  checks.
	// +kcc:proto:field=google.cloud.vmmigration.v1.DatacenterConnector.state
	State *string `json:"state,omitempty"`

	// Output only. The time the state was last set.
	// +kcc:proto:field=google.cloud.vmmigration.v1.DatacenterConnector.state_time
	StateTime *string `json:"stateTime,omitempty"`

	// Output only. Provides details on the state of the Datacenter Connector in
	//  case of an error.
	// +kcc:proto:field=google.cloud.vmmigration.v1.DatacenterConnector.error
	Error *common.Status `json:"error,omitempty"`

	// Output only. Appliance OVA version.
	//  This is the OVA which is manually installed by the user and contains the
	//  infrastructure for the automatically updatable components on the appliance.
	// +kcc:proto:field=google.cloud.vmmigration.v1.DatacenterConnector.appliance_infrastructure_version
	ApplianceInfrastructureVersion *string `json:"applianceInfrastructureVersion,omitempty"`

	// Output only. Appliance last installed update bundle version.
	//  This is the version of the automatically updatable components on the
	//  appliance.
	// +kcc:proto:field=google.cloud.vmmigration.v1.DatacenterConnector.appliance_software_version
	ApplianceSoftwareVersion *string `json:"applianceSoftwareVersion,omitempty"`

	// Output only. The available versions for updating this appliance.
	// +kcc:proto:field=google.cloud.vmmigration.v1.DatacenterConnector.available_versions
	AvailableVersions *AvailableUpdates `json:"availableVersions,omitempty"`

	// Output only. The status of the current / last upgradeAppliance operation.
	// +kcc:proto:field=google.cloud.vmmigration.v1.DatacenterConnector.upgrade_status
	UpgradeStatus *UpgradeStatus `json:"upgradeStatus,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvmmigrationdatacenterconnector;gcpvmmigrationdatacenterconnectors
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VMMigrationDatacenterConnector is the Schema for the VMMigrationDatacenterConnector API
// +k8s:openapi-gen=true
type VMMigrationDatacenterConnector struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   VMMigrationDatacenterConnectorSpec   `json:"spec,omitempty"`
	Status VMMigrationDatacenterConnectorStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// VMMigrationDatacenterConnectorList contains a list of VMMigrationDatacenterConnector
type VMMigrationDatacenterConnectorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VMMigrationDatacenterConnector `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VMMigrationDatacenterConnector{}, &VMMigrationDatacenterConnectorList{})
}
