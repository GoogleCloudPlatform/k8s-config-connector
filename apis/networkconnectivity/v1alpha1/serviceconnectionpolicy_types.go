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
	computev1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	refs "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	NetworkConnectivityServiceConnectionPolicyGVK = schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: GroupVersion.Version,
		Kind:    "NetworkConnectivityServiceConnectionPolicy",
	}
)

// NetworkConnectivityServiceConnectionPolicySpec defines the desired state of NetworkConnectivityServiceConnectionPolicy
// +kcc:proto=mockgcp.cloud.networkconnectivity.v1.ServiceConnectionPolicy
type NetworkConnectivityServiceConnectionPolicySpec struct {

	/* Immutable. The Project that this resource belongs to. */
	ProjectRef refs.ProjectRef `json:"projectRef"`

	/* Immutable. Location of the resource. */
	Location *string `json:"location"`

	// The NetworkConnectivityServiceConnectionPolicy name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// A description of this resource.
	Description *string `json:"description,omitempty"`

	// // User-defined labels.
	// Labels map[string]string `json:"labels,omitempty"`

	// // Immutable. The name of a ServiceConnectionPolicy. Format: projects/{project}/locations/{location}/serviceConnectionPolicies/{service_connection_policy} See: https://google.aip.dev/122#fields-representing-resource-names
	// Name *string `json:"name,omitempty"`

	// The resource path of the consumer network. Example: - projects/{projectNumOrId}/global/networks/{resourceId}.
	Network *computev1beta1.ComputeNetworkRef `json:"networkRef,omitempty"`

	// Configuration used for Private Service Connect connections. Used when Infrastructure is PSC.
	PscConfig *PscConfig `json:"pscConfig,omitempty"`

	// The service class identifier for which this ServiceConnectionPolicy is for. The service class identifier is a unique, symbolic representation of a ServiceClass. It is provided by the Service Producer. Google services have a prefix of gcp. For example, gcp-cloud-sql. 3rd party services do not. For example, test-service-a3dfcx.
	ServiceClass *string `json:"serviceClass,omitempty"`
}

// +kcc:proto=mockgcp.cloud.networkconnectivity.v1.PscConfig
type PscConfig struct {
	/* TODO:AFTER-MAPPINGS
	// Optional. List of Projects, Folders, or Organizations from where the Producer instance can be within. For example, a network administrator can provide both 'organizations/foo' and 'projects/bar' as allowed_google_producers_resource_hierarchy_levels. This allowlists this network to connect with any Producer instance within the 'foo' organization or the 'bar' project. By default, allowed_google_producers_resource_hierarchy_level is empty. The format for each allowed_google_producers_resource_hierarchy_level is / where is one of 'projects', 'folders', or 'organizations' and is either the ID or the number of the resource type. Format for each allowed_google_producers_resource_hierarchy_level value: 'projects/' or 'folders/' or 'organizations/' Eg. [projects/my-project-id, projects/567, folders/891, organizations/123]
	AllowedGoogleProducersResourceHierarchyLevel []string `json:"allowedGoogleProducersResourceHierarchyLevel,omitempty"`
	*/

	// Optional. Max number of PSC connections for this policy.
	Limit *int64 `json:"limit,omitempty"`

	// Required. ProducerInstanceLocation is used to specify which authorization mechanism to use to determine which projects the Producer instance can be within.
	ProducerInstanceLocation *string `json:"producerInstanceLocation,omitempty"`

	// The resource paths of subnetworks to use for IP address management. Example: projects/{projectNumOrId}/regions/{region}/subnetworks/{resourceId}.
	Subnetworks []computev1beta1.ComputeSubnetworkRef `json:"subnetworkRefs,omitempty"`
}

// NetworkConnectivityServiceConnectionPolicyStatus defines the config connector machine state of NetworkConnectivityServiceConnectionPolicy
type NetworkConnectivityServiceConnectionPolicyStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* A unique specifier for the NetworkConnectivityServiceConnectionPolicy resource in GCP.*/
	// +optional
	ExternalRef *string `json:"externalRef,omitempty"`

	/* ObservedState is the state of the resource as most recently observed in GCP. */
	// +optional
	ObservedState *NetworkConnectivityServiceConnectionPolicyObservedState `json:"observedState,omitempty"`
}

// NetworkConnectivityServiceConnectionPolicySpec defines the desired state of NetworkConnectivityServiceConnectionPolicy
// +kcc:proto=mockgcp.cloud.networkconnectivity.v1.ServiceConnectionPolicy
type NetworkConnectivityServiceConnectionPolicyObservedState struct {
	// Output only. Time when the ServiceConnectionMap was created.
	CreateTime *string `json:"createTime,omitempty"`

	// Optional. The etag is computed by the server, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
	Etag *string `json:"etag,omitempty"`

	// Output only. The type of underlying resources used to create the connection.
	Infrastructure *string `json:"infrastructure,omitempty"`

	// Output only. [Output only] Information about each Private Service Connect connection.
	PscConnections []PscConnection `json:"pscConnections,omitempty"`

	// Output only. Time when the ServiceConnectionMap was updated.
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpnetworkconnectivityserviceconnectionpolicy;gcpnetworkconnectivityserviceconnectionpolicies
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// NetworkConnectivityServiceConnectionPolicy is the Schema for the NetworkConnectivityServiceConnectionPolicy API
// +k8s:openapi-gen=true
type NetworkConnectivityServiceConnectionPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkConnectivityServiceConnectionPolicySpec   `json:"spec,omitempty"`
	Status NetworkConnectivityServiceConnectionPolicyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// NetworkConnectivityServiceConnectionPolicyList contains a list of NetworkConnectivityServiceConnectionPolicy
type NetworkConnectivityServiceConnectionPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkConnectivityServiceConnectionPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NetworkConnectivityServiceConnectionPolicy{}, &NetworkConnectivityServiceConnectionPolicyList{})
}
