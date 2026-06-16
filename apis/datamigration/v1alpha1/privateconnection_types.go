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
	common "github.com/GoogleCloudPlatform/k8s-config-connector/apis/common"
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var DatabaseMigrationPrivateConnectionGVK = GroupVersion.WithKind("DatabaseMigrationPrivateConnection")

// +kcc:proto=google.cloud.clouddms.v1.VpcPeeringConfig
type VpcPeeringConfig struct {
	// Required. Fully qualified name of the VPC that Database Migration Service
	//  will peer to.
	// +kcc:proto:field=google.cloud.clouddms.v1.VpcPeeringConfig.vpc_name
	VpcNameRef *computev1beta1.ComputeNetworkRef `json:"vpcNameRef,omitempty"`

	// Required. A free subnet for peering. (CIDR of /29)
	// +kcc:proto:field=google.cloud.clouddms.v1.VpcPeeringConfig.subnet
	Subnet *string `json:"subnet,omitempty"`
}

// DatabaseMigrationPrivateConnectionSpec defines the desired state of DatabaseMigrationPrivateConnection
// +kcc:spec:proto=google.cloud.clouddms.v1.PrivateConnection
type DatabaseMigrationPrivateConnectionSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The DatabaseMigrationPrivateConnection name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// The resource labels for private connections to use to annotate any related
	//  underlying resources such as Compute Engine VMs. An object containing a
	//  list of "key": "value" pairs.
	//
	//  Example: `{ "name": "wrench", "mass": "1.3kg", "count": "3" }`.
	// +kcc:proto:field=google.cloud.clouddms.v1.PrivateConnection.labels
	Labels map[string]string `json:"labels,omitempty"`

	// The private connection display name.
	// +kcc:proto:field=google.cloud.clouddms.v1.PrivateConnection.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// VPC peering configuration.
	// +kcc:proto:field=google.cloud.clouddms.v1.PrivateConnection.vpc_peering_config
	VpcPeeringConfig *VpcPeeringConfig `json:"vpcPeeringConfig,omitempty"`
}

// DatabaseMigrationPrivateConnectionStatus defines the config connector machine state of DatabaseMigrationPrivateConnection
type DatabaseMigrationPrivateConnectionStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the DatabaseMigrationPrivateConnection resource in Google Cloud.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in Google Cloud.
	ObservedState *DatabaseMigrationPrivateConnectionObservedState `json:"observedState,omitempty"`
}

// DatabaseMigrationPrivateConnectionObservedState is the state of the DatabaseMigrationPrivateConnection resource as most recently observed in Google Cloud.
// +kcc:observedstate:proto=google.cloud.clouddms.v1.PrivateConnection
type DatabaseMigrationPrivateConnectionObservedState struct {
	// Output only. The create time of the resource.
	// +kcc:proto:field=google.cloud.clouddms.v1.PrivateConnection.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The last update time of the resource.
	// +kcc:proto:field=google.cloud.clouddms.v1.PrivateConnection.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The state of the private connection.
	// +kcc:proto:field=google.cloud.clouddms.v1.PrivateConnection.state
	State *string `json:"state,omitempty"`

	// Output only. The error details in case of state FAILED.
	// +kcc:proto:field=google.cloud.clouddms.v1.PrivateConnection.error
	Error *common.Status `json:"error,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpdatabasemigrationprivateconnection;gcpdatabasemigrationprivateconnections
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// DatabaseMigrationPrivateConnection is the Schema for the DatabaseMigrationPrivateConnection API
// +k8s:openapi-gen=true
type DatabaseMigrationPrivateConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   DatabaseMigrationPrivateConnectionSpec   `json:"spec,omitempty"`
	Status DatabaseMigrationPrivateConnectionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// DatabaseMigrationPrivateConnectionList contains a list of DatabaseMigrationPrivateConnection
type DatabaseMigrationPrivateConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatabaseMigrationPrivateConnection `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DatabaseMigrationPrivateConnection{}, &DatabaseMigrationPrivateConnectionList{})
}
