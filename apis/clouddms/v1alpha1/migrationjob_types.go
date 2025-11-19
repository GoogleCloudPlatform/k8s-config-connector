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

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/common/parent"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CloudDMSMigrationJobGVK = GroupVersion.WithKind("CloudDMSMigrationJob")

// CloudDMSMigrationJobSpec defines the desired state of CloudDMSMigrationJob
// +kcc:spec:proto=google.cloud.clouddms.v1.MigrationJob
type CloudDMSMigrationJobSpec struct {
	// Required. Defines the parent path of the resource.
	*parent.ProjectAndLocationRef `json:",inline"`

	// The resource labels for migration job to use to annotate any related
	//  underlying resources such as Compute Engine VMs. An object containing a
	//  list of "key": "value" pairs.
	//
	//  Example: `{ "name": "wrench", "mass": "1.3kg", "count": "3" }`.
	// +kcc:proto:field=google.cloud.clouddms.v1.MigrationJob.labels
	// Labels map[string]string `json:"labels,omitempty"`

	// The migration job display name.
	// +kcc:proto:field=google.cloud.clouddms.v1.MigrationJob.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Required. The migration job type.
	// +kcc:proto:field=google.cloud.clouddms.v1.MigrationJob.type
	Type *string `json:"type,omitempty"`

	// The path to the dump file in Google Cloud Storage,
	//  in the format: (gs://[BUCKET_NAME]/[OBJECT_NAME]).
	//  This field and the "dump_flags" field are mutually exclusive.
	// +kcc:proto:field=google.cloud.clouddms.v1.MigrationJob.dump_path
	DumpPath *string `json:"dumpPath,omitempty"`

	// The initial dump flags.
	//  This field and the "dump_path" field are mutually exclusive.
	// +kcc:proto:field=google.cloud.clouddms.v1.MigrationJob.dump_flags
	DumpFlags *MigrationJob_DumpFlags `json:"dumpFlags,omitempty"`

	// Required. The Connection Profile resource of the source connection profile.
	// +kcc:proto:field=google.cloud.clouddms.v1.MigrationJob.source
	SourceRef *CloudDMSConnectionProfileRef `json:"sourceRef,omitempty"`

	// Required. The Connection Profile of the destination connection profile.
	// +kcc:proto:field=google.cloud.clouddms.v1.MigrationJob.destination
	DestinationRef *CloudDMSConnectionProfileRef `json:"destinationRef,omitempty"`

	// The details needed to communicate to the source over Reverse SSH
	//  tunnel connectivity.
	// +kcc:proto:field=google.cloud.clouddms.v1.MigrationJob.reverse_ssh_connectivity
	ReverseSSHConnectivity *ReverseSSHConnectivity `json:"reverseSSHConnectivity,omitempty"`

	// The details of the VPC network that the source database is located in.
	// +kcc:proto:field=google.cloud.clouddms.v1.MigrationJob.vpc_peering_connectivity
	VPCPeeringConnectivity *VPCPeeringConnectivity `json:"vpcPeeringConnectivity,omitempty"`

	// static ip connectivity data (default, no additional details needed).
	// +kcc:proto:field=google.cloud.clouddms.v1.MigrationJob.static_ip_connectivity
	StaticIPConnectivity *StaticIPConnectivity `json:"staticIPConnectivity,omitempty"`

	// The database engine type and provider of the source.
	// +kcc:proto:field=google.cloud.clouddms.v1.MigrationJob.source_database
	SourceDatabase *DatabaseType `json:"sourceDatabase,omitempty"`

	// The database engine type and provider of the destination.
	// +kcc:proto:field=google.cloud.clouddms.v1.MigrationJob.destination_database
	DestinationDatabase *DatabaseType `json:"destinationDatabase,omitempty"`

	// The conversion workspace used by the migration.
	// +kcc:proto:field=google.cloud.clouddms.v1.MigrationJob.conversion_workspace
	ConversionWorkspace *ConversionWorkspaceInfo `json:"conversionWorkspace,omitempty"`

	// This field can be used to select the entities to migrate as part of
	//  the migration job. It uses AIP-160 notation to select a subset of the
	//  entities configured on the associated conversion-workspace. This field
	//  should not be set on migration-jobs that are not associated with a
	//  conversion workspace.
	// +kcc:proto:field=google.cloud.clouddms.v1.MigrationJob.filter
	Filter *string `json:"filter,omitempty"`

	// The CMEK (customer-managed encryption key) fully qualified key name used
	//  for the migration job.
	//  This field supports all migration jobs types except for:
	//  * Mysql to Mysql (use the cmek field in the cloudsql connection profile
	//  instead).
	//  * PostrgeSQL to PostgreSQL (use the cmek field in the cloudsql
	//  connection profile instead).
	//  * PostgreSQL to AlloyDB (use the kms_key_name field in the alloydb
	//  connection profile instead).
	//  Each Cloud CMEK key has the following format:
	//  projects/[PROJECT]/locations/[REGION]/keyRings/[RING]/cryptoKeys/[KEY_NAME]
	// +kcc:proto:field=google.cloud.clouddms.v1.MigrationJob.cmek_key_name
	CmekKeyNameRef *refsv1beta1.KMSCryptoKeyRef `json:"cmekKeyNameRef,omitempty"`

	// Optional. Data dump parallelism settings used by the migration.
	//  Currently applicable only for MySQL to Cloud SQL for MySQL migrations only.
	// +kcc:proto:field=google.cloud.clouddms.v1.MigrationJob.performance_config
	PerformanceConfig *MigrationJob_PerformanceConfig `json:"performanceConfig,omitempty"`

	// The CloudDMSMigrationJob name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
}

// CloudDMSMigrationJobStatus defines the config connector machine state of CloudDMSMigrationJob
type CloudDMSMigrationJobStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the CloudDMSMigrationJob resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *CloudDMSMigrationJobObservedState `json:"observedState,omitempty"`
}

// CloudDMSMigrationJobObservedState is the state of the CloudDMSMigrationJob resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.clouddms.v1.MigrationJob
type CloudDMSMigrationJobObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpclouddmsmigrationjob;gcpclouddmsmigrationjobs
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CloudDMSMigrationJob is the Schema for the CloudDMSMigrationJob API
// +k8s:openapi-gen=true
type CloudDMSMigrationJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CloudDMSMigrationJobSpec   `json:"spec,omitempty"`
	Status CloudDMSMigrationJobStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CloudDMSMigrationJobList contains a list of CloudDMSMigrationJob
type CloudDMSMigrationJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudDMSMigrationJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudDMSMigrationJob{}, &CloudDMSMigrationJobList{})
}
