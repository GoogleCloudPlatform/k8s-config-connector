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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var AlloyDBInstanceGVK = GroupVersion.WithKind("AlloyDBInstance")

// AlloyDBInstanceSpec defines the desired state of AlloyDBInstance
// +kcc:proto=google.cloud.alloydb.v1beta.Instance
type AlloyDBInstanceSpec struct {

	// The AlloyDBInstance cluster that this resource belongs to.
	// +required
	ClusterRef *refs.AlloyDBClusterRef `json:"clusterRef,omitempty"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="ResourceID field is immutable"
	// Immutable.
	// Optional. The instanceId of the resource. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Annotations to allow client tools to store small amount
	// of arbitrary data. This is distinct from labels.
	Annotations map[string]string `json:"annotations,omitempty"`

	// Availability type of an Instance. Defaults to REGIONAL for both primary and read instances.
	//
	// Note that primary and read instances can have different availability types.
	// Only READ_POOL instance supports ZONAL type. Users can't specify the zone for READ_POOL instance.
	// Zone is automatically chosen from the list of zones in the region specified.
	// Read pool of size 1 can only have zonal availability. Read pools with node count of 2 or more
	// can have regional availability (nodes are present in 2 or more zones in a region). Possible values: ["AVAILABILITY_TYPE_UNSPECIFIED", "ZONAL", "REGIONAL"].
	AvailabilityType *string `json:"availabilityType,omitempty"`

	// Database flags. Set at instance level. * They are copied
	// from primary instance on read instance creation. * Read instances
	// can set new or override existing flags that are relevant for reads,
	// e.g. for enabling columnar cache on a read instance. Flags set on
	// read instance may or may not be present on primary.
	DatabaseFlags map[string]string `json:"databaseFlags,omitempty"`

	// User-settable and human-readable display name for the
	// Instance.
	DisplayName *string `json:"displayName,omitempty"`

	// The Compute Engine zone that the instance should serve
	// from, per https://cloud.google.com/compute/docs/regions-zones This
	// can ONLY be specified for ZONAL instances. If present for a REGIONAL
	// instance, an error will be thrown. If this is absent for a ZONAL
	// instance, instance is created in a random zone with available capacity.
	GceZone *string `json:"gceZone,omitempty"`

	// We recommend that you use `instanceTypeRef` instead.
	// The type of the instance. Possible values: [PRIMARY, READ_POOL, SECONDARY]
	InstanceType *string `json:"instanceType,omitempty"`

	// The type of instance.
	// Possible values: ["PRIMARY", "READ_POOL", "SECONDARY"]
	//
	// For PRIMARY and SECONDARY instances, set the value to refer to the name of the associated cluster.
	// This is recommended because the instance type of primary and secondary instances is tied to the cluster type of the associated cluster.
	// If the secondary cluster is promoted to primary cluster, then the associated secondary instance also becomes primary instance.
	// Example:
	// instanceTypeRef:
	//   name: clusterName
	// For instances of type READ_POOL, set the value using external keyword.
	// Example:
	// instanceTypeRef:
	//   external: READ_POOL
	// If the instance type is SECONDARY, the delete instance operation does not delete the secondary instance but abandons it instead.
	// Use deletionPolicy = "FORCE" in the associated secondary cluster and delete the cluster forcefully to delete the secondary cluster as well its associated secondary instance.
	InstanceTypeRef *refs.AlloyDBClusterTypeRef `json:"instanceTypeRef,omitempty"`

	// Configurations for the machines that host the underlying
	// database engine.
	MachineConfig *Instance_MachineConfig `json:"machineConfig,omitempty"`

	// Instance level network configuration.
	NetworkConfig *Instance_InstanceNetworkConfig `json:"networkConfig,omitempty"`

	// Read pool specific config. If the instance type is READ_POOL,
	// this configuration must be provided.
	ReadPoolConfig *Instance_ReadPoolConfig `json:"readPoolConfig,omitempty"`
}

// AlloyDBInstanceStatus defines the config connector machine state of AlloyDBInstance
type AlloyDBInstanceStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the AlloyDBInstance resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	/* NOTYET
	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *AlloyDBInstanceObservedState `json:"observedState,omitempty"`
	*/

	// Time the Instance was created in UTC.
	CreateTime *string `json:"createTime,omitempty"`

	// The IP address for the Instance. This is the connection
	// endpoint for an end-user application.
	IpAddress *string `json:"ipAddress,omitempty"`

	// The name of the instance resource.
	Name *string `json:"name,omitempty"`

	// The outbound public IP addresses for the instance. This is available ONLY when
	// networkConfig.enableOutboundPublicIp is set to true. These IP addresses are used
	// for outbound connections.
	OutboundPublicIpAddresses []string `json:"outboundPublicIpAddresses,omitempty"`

	// The public IP addresses for the Instance. This is available
	// ONLY when networkConfig.enablePublicIp is set to true. This is the
	// connection endpoint for an end-user application.
	PublicIpAddress *string `json:"publicIpAddress,omitempty"`

	// Set to true if the current state of Instance does not
	// match the user's intended state, and the service is actively updating
	// the resource to reconcile them. This can happen due to user-triggered
	// updates or system actions like failover or maintenance.
	Reconciling *bool `json:"reconciling,omitempty"`

	// The current state of the alloydb instance.
	State *string `json:"state,omitempty"`

	// The system-generated UID of the resource.
	Uid *string `json:"uid,omitempty"`

	// Time the Instance was updated in UTC.
	UpdateTime *string `json:"updateTime,omitempty"`
}

/* NOTYET
// AlloyDBInstanceSpec defines the desired state of AlloyDBInstance
// +kcc:proto=google.cloud.alloydb.v1beta.Instance
// AlloyDBInstanceObservedState is the state of the AlloyDBInstance resource as most recently observed in GCP.
type AlloyDBInstanceObservedState struct {
}
*/

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// TODO(user): make sure the pluralizaiton below is correct
// +kubebuilder:resource:categories=gcp,shortName=gcpalloydbinstance;gcpalloydbinstances
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true";"cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=stable";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// AlloyDBInstance is the Schema for the AlloyDBInstance API
// +k8s:openapi-gen=true
// +kubebuilder:storageversion
type AlloyDBInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   AlloyDBInstanceSpec   `json:"spec,omitempty"`
	Status AlloyDBInstanceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// AlloyDBInstanceList contains a list of AlloyDBInstance
type AlloyDBInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AlloyDBInstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AlloyDBInstance{}, &AlloyDBInstanceList{})
}
