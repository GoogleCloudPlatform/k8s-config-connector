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
// +kcc:spec:proto=google.cloud.alloydb.v1beta.Instance
type AlloyDBInstanceSpec struct {

	// The AlloyDBInstance cluster that this resource belongs to.
	// +required
	ClusterRef *refs.AlloyDBClusterRef `json:"clusterRef,omitempty"`

	// Optional. The instanceId of the resource. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Annotations to allow client tools to store small amount
	// of arbitrary data. This is distinct from labels.
	Annotations map[string]string `json:"annotations,omitempty"`

	// Availability type of an Instance. If empty, defaults to REGIONAL for primary instances.
	//
	// For read pools, availabilityType is always UNSPECIFIED. Instances in the
	// read pools are evenly distributed across available zones within the region
	// (i.e. read pools with more than one node will have a node in at least two zones).
	// Possible values: ["AVAILABILITY_TYPE_UNSPECIFIED", "ZONAL", "REGIONAL"].
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
	GCEZone *string `json:"gceZone,omitempty"`

	// Not recommended. We recommend that you use `instanceTypeRef` instead.
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

// +kcc:proto=google.cloud.alloydb.v1beta.Instance.InstanceNetworkConfig
type Instance_InstanceNetworkConfig struct {
	// Optional. A list of external network authorized to
	// access this instance. This field is only allowed to be set when
	// 'enablePublicIp' is set to true.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Instance.InstanceNetworkConfig.authorized_external_networks
	AuthorizedExternalNetworks []Instance_InstanceNetworkConfig_AuthorizedNetwork `json:"authorizedExternalNetworks,omitempty"`

	// Optional. Enabling public ip for the instance. If
	// a user wishes to disable this, please also clear the list of
	// the authorized external networks set on the same instance.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Instance.InstanceNetworkConfig.enable_public_ip
	EnablePublicIP *bool `json:"enablePublicIp,omitempty"`

	// Optional. Enabling an outbound public IP address to support a database
	//  server sending requests out into the internet.
	// +kcc:proto:field=google.cloud.alloydb.v1beta.Instance.InstanceNetworkConfig.enable_outbound_public_ip
	EnableOutboundPublicIP *bool `json:"enableOutboundPublicIp,omitempty"`
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
	IPAddress *string `json:"ipAddress,omitempty"`

	// The name of the instance resource.
	Name *string `json:"name,omitempty"`

	// The outbound public IP addresses for the instance. This is available ONLY when
	// networkConfig.enableOutboundPublicIp is set to true. These IP addresses are used
	// for outbound connections.
	OutboundPublicIPAddresses []string `json:"outboundPublicIpAddresses,omitempty"`

	// The public IP addresses for the Instance. This is available
	// ONLY when networkConfig.enablePublicIp is set to true. This is the
	// connection endpoint for an end-user application.
	PublicIPAddress *string `json:"publicIpAddress,omitempty"`

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
// +kubebuilder:resource:categories=gcp,shortName=gcpalloydbinstance;gcpalloydbinstances
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=stable";"cnrm.cloud.google.com/system=true"
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
