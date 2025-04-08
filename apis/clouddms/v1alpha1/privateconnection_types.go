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
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CloudDMSPrivateConnectionGVK = GroupVersion.WithKind("CloudDMSPrivateConnection")

type Parent struct {
	// Required. The location of the application.
	Location string `json:"location,omitempty"`

	// Required. The host project of the application.
	ProjectRef *refs.ProjectRef `json:"projectRef,omitempty"`
}

// +kcc:proto=google.protobuf.Any
type Any struct {
	// A URL/resource name that uniquely identifies the type of the serialized
	//  protocol buffer message. This string must contain at least
	//  one "/" character. The last segment of the URL's path must represent
	//  the fully qualified name of the type (as in
	//  `path/google.protobuf.Duration`). The name should be in a canonical form
	//  (e.g., leading "." is not accepted).
	//
	//  In practice, teams usually precompile into the binary all types that they
	//  expect it to use in the context of Any. However, for URLs which use the
	//  scheme `http`, `https`, or no scheme, one can optionally set up a type
	//  server that maps type URLs to message definitions as follows:
	//
	//  * If no scheme is provided, `https` is assumed.
	//  * An HTTP GET on the URL must yield a [google.protobuf.Type][]
	//    value in binary format, or produce an error.
	//  * Applications are allowed to cache lookup results based on the
	//    URL, or have them precompiled into a binary to avoid any
	//    lookup. Therefore, binary compatibility needs to be preserved
	//    on changes to types. (Use versioned type names to manage
	//    breaking changes.)
	//
	//  Note: this functionality is not currently available in the official
	//  protobuf release, and it is not used for type URLs beginning with
	//  type.googleapis.com.
	//
	//  Schemes other than `http`, `https` (or the empty scheme) might be
	//  used with implementation specific semantics.
	// +kcc:proto:field=google.protobuf.Any.type_url
	TypeURL *string `json:"typeURL,omitempty"`

	// Must be a valid serialized protocol buffer of the above specified type.
	// +kcc:proto:field=google.protobuf.Any.value
	Value []byte `json:"value,omitempty"`
}

// +kcc:proto=google.rpc.Status
type Status struct {
	// The status code, which should be an enum value of
	//  [google.rpc.Code][google.rpc.Code].
	// +kcc:proto:field=google.rpc.Status.code
	Code *int32 `json:"code,omitempty"`

	// A developer-facing error message, which should be in English. Any
	//  user-facing error message should be localized and sent in the
	//  [google.rpc.Status.details][google.rpc.Status.details] field, or localized
	//  by the client.
	// +kcc:proto:field=google.rpc.Status.message
	Message *string `json:"message,omitempty"`

	// NOT YET
	// A list of messages that carry the error details.  There is a common set of
	//  message types for APIs to use.
	// +kcc:proto:field=google.rpc.Status.details
	// Details []Any `json:"details,omitempty"`
}

// +kcc:proto=google.cloud.clouddms.v1.VpcPeeringConfig
type VpcPeeringConfig struct {
	// Required. Fully qualified name of the VPC that Database Migration Service
	//  will peer to.
	// +kcc:proto:field=google.cloud.clouddms.v1.VpcPeeringConfig.vpc_name
	VpcName *refs.ComputeNetworkRef `json:"vpcName,omitempty"`

	// Required. A free subnet for peering. (CIDR of /29)
	// +kcc:proto:field=google.cloud.clouddms.v1.VpcPeeringConfig.subnet
	Subnet *string `json:"subnet,omitempty"`
}

// CloudDMSPrivateConnectionStatus defines the config connector machine state of CloudDMSPrivateConnection
type CloudDMSPrivateConnectionStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the CloudDMSPrivateConnection resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *CloudDMSPrivateConnectionObservedState `json:"observedState,omitempty"`
}

// CloudDMSPrivateConnectionSpec defines the desired state of CloudDMSPrivateConnection
// +kcc:proto=google.cloud.clouddms.v1.PrivateConnection
type CloudDMSPrivateConnectionSpec struct {
	// The CloudDMSPrivateConnection name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. Defines the parent path of the resource.
	*Parent `json:",inline"`

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

// CloudDMSPrivateConnectionObservedState is the state of the CloudDMSPrivateConnection resource as most recently observed in GCP.
// +kcc:proto=google.cloud.clouddms.v1.PrivateConnection
type CloudDMSPrivateConnectionObservedState struct {
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
	Error *Status `json:"error,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpclouddmsprivateconnection;gcpclouddmsprivateconnections
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// CloudDMSPrivateConnection is the Schema for the CloudDMSPrivateConnection API
// +k8s:openapi-gen=true
type CloudDMSPrivateConnection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   CloudDMSPrivateConnectionSpec   `json:"spec,omitempty"`
	Status CloudDMSPrivateConnectionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// CloudDMSPrivateConnectionList contains a list of CloudDMSPrivateConnection
type CloudDMSPrivateConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudDMSPrivateConnection `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CloudDMSPrivateConnection{}, &CloudDMSPrivateConnectionList{})
}
