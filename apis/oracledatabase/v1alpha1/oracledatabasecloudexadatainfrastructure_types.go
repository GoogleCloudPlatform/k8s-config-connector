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

// API sources for OracleDatabaseCloudExadataInfrastructure, recorded by generate-types:
// +kcc:source:proto=https://github.com/googleapis/googleapis/blob/054c87e45c4683ea5430ab609cbd4a445c4c4c71/google/cloud/oracledatabase/v1/exadata_infra.proto
// +kcc:source:service-docs=https://cloud.google.com/oracle/database/docs
// +kcc:source:resource-docs=https://docs.cloud.google.com/oracle/database/docs/reference/rest/v1/projects.locations.cloudExadataInfrastructures

package v1alpha1

import (
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var OracleDatabaseCloudExadataInfrastructureGVK = GroupVersion.WithKind("OracleDatabaseCloudExadataInfrastructure")

// OracleDatabaseCloudExadataInfrastructureSpec defines the desired state of OracleDatabaseCloudExadataInfrastructure
// +kcc:spec:proto=google.cloud.oracledatabase.v1.CloudExadataInfrastructure
// +kcc:required-from-proto
type OracleDatabaseCloudExadataInfrastructureSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The OracleDatabaseCloudExadataInfrastructure name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`
	// Optional. User friendly name for this resource.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudExadataInfrastructure.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. The GCP Oracle zone where Oracle Exadata Infrastructure is
	//  hosted. Example: us-east4-b-r2. If not specified, the system will pick a
	//  zone based on availability.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudExadataInfrastructure.gcp_oracle_zone
	GcpOracleZone *string `json:"gcpOracleZone,omitempty"`

	// Optional. Various properties of the infra.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudExadataInfrastructure.properties
	Properties *CloudExadataInfrastructureProperties `json:"properties,omitempty"`

	// Optional. Labels or tags associated with the resource.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudExadataInfrastructure.labels
	Labels map[string]string `json:"labels,omitempty"`
}

// OracleDatabaseCloudExadataInfrastructureStatus defines the config connector machine state of OracleDatabaseCloudExadataInfrastructure
type OracleDatabaseCloudExadataInfrastructureStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the OracleDatabaseCloudExadataInfrastructure resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *OracleDatabaseCloudExadataInfrastructureObservedState `json:"observedState,omitempty"`
}

// OracleDatabaseCloudExadataInfrastructureObservedState is the state of the OracleDatabaseCloudExadataInfrastructure resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.oracledatabase.v1.CloudExadataInfrastructure
type OracleDatabaseCloudExadataInfrastructureObservedState struct {
	// Output only. Entitlement ID of the private offer against which this
	//  infrastructure resource is provisioned.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudExadataInfrastructure.entitlement_id
	EntitlementID *string `json:"entitlementID,omitempty"`

	// Optional. Various properties of the infra.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudExadataInfrastructure.properties
	Properties *CloudExadataInfrastructurePropertiesObservedState `json:"properties,omitempty"`

	// Output only. The date and time that the Exadata Infrastructure was created.
	// +kcc:proto:field=google.cloud.oracledatabase.v1.CloudExadataInfrastructure.create_time
	CreateTime *string `json:"createTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcporacledatabasecloudexadatainfrastructure;gcporacledatabasecloudexadatainfrastructures
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// OracleDatabaseCloudExadataInfrastructure is the Schema for the OracleDatabaseCloudExadataInfrastructure API
// +k8s:openapi-gen=true
type OracleDatabaseCloudExadataInfrastructure struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   OracleDatabaseCloudExadataInfrastructureSpec   `json:"spec,omitempty"`
	Status OracleDatabaseCloudExadataInfrastructureStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// OracleDatabaseCloudExadataInfrastructureList contains a list of OracleDatabaseCloudExadataInfrastructure
type OracleDatabaseCloudExadataInfrastructureList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OracleDatabaseCloudExadataInfrastructure `json:"items"`
}

func init() {
	SchemeBuilder.Register(&OracleDatabaseCloudExadataInfrastructure{}, &OracleDatabaseCloudExadataInfrastructureList{})
}
