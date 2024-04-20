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
	"reflect"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var ServiceConnectionPolicyGVK = schema.GroupVersionKind{
	Group:   SchemeGroupVersion.Group,
	Version: SchemeGroupVersion.Version,
	Kind:    reflect.TypeOf(NetworkConnectivityServiceConnectionPolicy{}).Name(),
}

type ProjectRef struct {
	/* The external name of the referenced resource */
	External string `json:"external,omitempty"`
	/* Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names */
	Name string `json:"name,omitempty"`
	/* Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ */
	Namespace string `json:"namespace,omitempty"`
}

type ServiceConnectionPolicySpec struct {

	/* The project that this resource belongs to. */
	ProjectRef ProjectRef `json:"projectRef"`

	/* Immutable. The location where the cluster should reside. */
	Location string `json:"location"`

	/* Immutable. Optional. The apiId of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	// Description: A description of this resource.
	Description string `json:"description,omitempty"`

	// Labels: User-defined labels.
	Labels map[string]string `json:"labels,omitempty"`

	// Network: The resource path of the consumer network. Example: -
	// projects/{projectNumOrId}/global/networks/{resourceId}.
	Network string `json:"network,omitempty"`

	// PscConfig: Configuration used for Private Service Connect
	// connections. Used when Infrastructure is PSC.
	PscConfig *PscConfig `json:"pscConfig,omitempty"`

	// ServiceClass: The service class identifier for which this
	// ServiceConnectionPolicy is for. The service class identifier is a
	// unique, symbolic representation of a ServiceClass. It is provided by
	// the Service Producer. Google services have a prefix of gcp. For
	// example, gcp-cloud-sql. 3rd party services do not. For example,
	// test-service-a3dfcx.
	ServiceClass string `json:"serviceClass,omitempty"`
}

type ServiceConnectionPolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   ComputeNetwork's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	ObservedState *ServiceConnectionPolicyObservedState `json:"observedState,omitempty"`
}

type ServiceConnectionPolicyObservedState struct {
	// // CreateTime: Output only. Time when the ServiceConnectionMap was
	// // created.
	// CreateTime string `json:"createTime,omitempty"`

	// // Etag: Optional. The etag is computed by the server, and may be sent
	// // on update and delete requests to ensure the client has an up-to-date
	// // value before proceeding.
	// Etag string `json:"etag,omitempty"`

	// Infrastructure: Output only. The type of underlying resources used to
	// create the connection.
	//
	// Possible values:
	//   "INFRASTRUCTURE_UNSPECIFIED" - An invalid infrastructure as the
	// default case.
	//   "PSC" - Private Service Connect is used for connections.
	Infrastructure string `json:"infrastructure,omitempty"`

	// // Name: Immutable. The name of a ServiceConnectionPolicy. Format:
	// // projects/{project}/locations/{location}/serviceConnectionPolicies/{ser
	// // vice_connection_policy} See:
	// // https://google.aip.dev/122#fields-representing-resource-names
	// Name string `json:"name,omitempty"`

	// PscConnections: Output only. [Output only] Information about each
	// Private Service Connect connection.
	PscConnections []*PscConnection `json:"pscConnections,omitempty"`

	// // UpdateTime: Output only. Time when the ServiceConnectionMap was
	// // updated.
	// UpdateTime string `json:"updateTime,omitempty"`

}

type PscConfig struct {
	// Limit: Optional. Max number of PSC connections for this policy.
	Limit *int64 `json:"limit,omitempty,string"`

	// Subnetworks: The resource paths of subnetworks to use for IP address
	// management. Example:
	// projects/{projectNumOrId}/regions/{region}/subnetworks/{resourceId}.
	Subnetworks []string `json:"subnetworks,omitempty"`
}

type PscConnection struct {
	// ConsumerAddress: The resource reference of the consumer address.
	ConsumerAddress *string `json:"consumerAddress,omitempty"`

	// ConsumerForwardingRule: The resource reference of the PSC Forwarding
	// Rule within the consumer VPC.
	ConsumerForwardingRule *string `json:"consumerForwardingRule,omitempty"`

	// ConsumerTargetProject: The project where the PSC connection is
	// created.
	ConsumerTargetProject *string `json:"consumerTargetProject,omitempty"`

	// // Error: The most recent error during operating this connection.
	// Error *GoogleRpcStatus `json:"error,omitempty"`

	// // ErrorInfo: Output only. The error info for the latest error during
	// // operating this connection.
	// ErrorInfo *GoogleRpcErrorInfo `json:"errorInfo,omitempty"`

	// // ErrorType: The error type indicates whether the error is consumer
	// // facing, producer facing or system internal.
	// //
	// // Possible values:
	// //   "CONNECTION_ERROR_TYPE_UNSPECIFIED" - An invalid error type as the
	// // default case.
	// //   "ERROR_INTERNAL" - The error is due to Service Automation system
	// // internal.
	// //   "ERROR_CONSUMER_SIDE" - The error is due to the setup on consumer
	// // side.
	// //   "ERROR_PRODUCER_SIDE" - The error is due to the setup on producer
	// // side.
	// ErrorType string `json:"errorType,omitempty"`

	// // GceOperation: The last Compute Engine operation to setup PSC
	// // connection.
	// GceOperation string `json:"gceOperation,omitempty"`

	// PscConnectionId: The PSC connection id of the PSC forwarding rule.
	PscConnectionId *string `json:"pscConnectionId,omitempty"`

	// SelectedSubnetwork: Output only. The URI of the subnetwork selected
	// to allocate IP address for this connection.
	SelectedSubnetwork *string `json:"selectedSubnetwork,omitempty"`

	// State: State of the PSC Connection
	//
	// Possible values:
	//   "STATE_UNSPECIFIED" - An invalid state as the default case.
	//   "ACTIVE" - The connection is fully established and ready to use.
	//   "FAILED" - The connection is not functional since some resources on
	// the connection fail to be created.
	//   "CREATING" - The connection is being created.
	//   "DELETING" - The connection is being deleted.
	State *string `json:"state,omitempty"`
}

// NetworkConnectivityServiceConnectionPolicy manages a serviceConnectionPolicy in the networkconnectivity API.
// +k8s:openapi-gen=true
// +kubebuilder:resource:categories=gcp,shortName=gcpserviceconnectionpolicies
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
type NetworkConnectivityServiceConnectionPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServiceConnectionPolicySpec   `json:"spec,omitempty"`
	Status ServiceConnectionPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NetworkConnectivityServiceConnectionPolicyList contains a list of ServiceConnectionPolicy
type NetworkConnectivityServiceConnectionPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkConnectivityServiceConnectionPolicy `json:"items"`
}

// func init() {
// 	SchemeBuilder.Register(&NetworkConnectivityServiceConnectionPolicy{}, &NetworkConnectivityServiceConnectionPolicyList{})
// }
