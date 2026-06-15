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
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var OracleDatabaseAutonomousDatabaseGVK = GroupVersion.WithKind("OracleDatabaseAutonomousDatabase")

// OracleDatabaseAutonomousDatabaseSpec defines the desired state of OracleDatabaseAutonomousDatabase
// +kcc:spec:proto=google.cloud.oracledatabase.v1.AutonomousDatabase
type OracleDatabaseAutonomousDatabaseSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// Optional. The name of the Autonomous Database. The database name must be
	//  unique in the project. The name must begin with a letter and can contain a
	//  maximum of 30 alphanumeric characters.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabase.database
	Database *string `json:"database,omitempty"`

	// Optional. The display name for the Autonomous Database. The name does not
	//  have to be unique within your project.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabase.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. The password for the default ADMIN user.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabase.admin_password
	AdminPassword *string `json:"adminPassword,omitempty"`

	// Optional. The properties of the Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabase.properties
	Properties *AutonomousDatabaseProperties `json:"properties,omitempty"`

	// Optional. The labels or tags associated with the Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabase.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Optional. The name of the VPC network used by the Autonomous Database in
	//  the following format: projects/{project}/global/networks/{network}
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabase.network
	NetworkRef *computev1beta1.ComputeNetworkRef `json:"networkRef,omitempty"`

	// Optional. The subnet CIDR range for the Autonmous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabase.cidr
	CIDR *string `json:"cidr,omitempty"`

	// The OracleDatabaseAutonomousDatabase name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
}

// OracleDatabaseAutonomousDatabaseStatus defines the config connector machine state of OracleDatabaseAutonomousDatabase
type OracleDatabaseAutonomousDatabaseStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the OracleDatabaseAutonomousDatabase resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *OracleDatabaseAutonomousDatabaseObservedState `json:"observedState,omitempty"`
}

// OracleDatabaseAutonomousDatabaseObservedState is the state of the OracleDatabaseAutonomousDatabase resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.oracledatabase.v1.AutonomousDatabase
type OracleDatabaseAutonomousDatabaseObservedState struct {
	// Output only. The ID of the subscription entitlement associated with the
	//  Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabase.entitlement_id
	EntitlementID *string `json:"entitlementID,omitempty"`

	// Optional. The properties of the Autonomous Database.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabase.properties
	Properties *AutonomousDatabasePropertiesObservedState `json:"properties,omitempty"`

	// Output only. The date and time that the Autonomous Database was created.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.AutonomousDatabase.create_time
	CreateTime *string `json:"createTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcporacledatabaseautonomousdatabase;gcporacledatabaseautonomousdatabases
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// OracleDatabaseAutonomousDatabase is the Schema for the OracleDatabaseAutonomousDatabase API
// +k8s:openapi-gen=true
type OracleDatabaseAutonomousDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   OracleDatabaseAutonomousDatabaseSpec   `json:"spec,omitempty"`
	Status OracleDatabaseAutonomousDatabaseStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// OracleDatabaseAutonomousDatabaseList contains a list of OracleDatabaseAutonomousDatabase
type OracleDatabaseAutonomousDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OracleDatabaseAutonomousDatabase `json:"items"`
}

func init() {
	SchemeBuilder.Register(&OracleDatabaseAutonomousDatabase{}, &OracleDatabaseAutonomousDatabaseList{})
}
