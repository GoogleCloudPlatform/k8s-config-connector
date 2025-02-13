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

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ApigeeEnvironmentGVK = GroupVersion.WithKind("ApigeeEnvironment")

// ApigeeEnvironmentSpec defines the desired state of ApigeeEnvironment
// +kcc:proto=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Environment
type ApigeeEnvironmentSpec struct {
	// The ApigeeEnvironment name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Reference to parent Apigee Organization.
	// +required
	ApigeeOrganizationRef *ApigeeOrganizationRef `json:"apigeeOrganizationRef,omitempty"`

	/* NOTYET: Add this once direct controller is implemented
	// Optional. API Proxy type supported by the environment. The type can be set when creating the Environment and cannot be changed.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Environment.api_proxy_type
	ApiProxyType *string `json:"apiProxyType,omitempty"`
	*/

	/* NOTYET: Add this once direct controller is implemented
	// Optional. Deployment type supported by the environment. The deployment type can be set when creating the environment and cannot be changed. When you enable archive deployment, you will be **prevented from performing** a [subset of actions](/apigee/docs/api-platform/local-development/overview#prevented-actions) within the environment, including: * Managing the deployment of API proxy or shared flow revisions * Creating, updating, or deleting resource files * Creating, updating, or deleting target servers
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Environment.deployment_type
	DeploymentType *string `json:"deploymentType,omitempty"`
	*/

	// Optional. Description of the environment.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Environment.description
	Description *string `json:"description,omitempty"`

	// Optional. Display name for this environment.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Environment.display_name
	DisplayName *string `json:"displayName,omitempty"`

	/* NOTYET: Add this once direct controller is implemented
	// Optional. URI of the forward proxy to be applied to the runtime instances in this environment. Must be in the format of {scheme}://{hostname}:{port}. Note that the scheme must be one of "http" or "https", and the port must be supplied. To remove a forward proxy setting, update the field to an empty value. Note: At this time, PUT operations to add forwardProxyUri to an existing environment fail if the environment has nodeConfig set up. To successfully add the forwardProxyUri setting in this case, include the NodeConfig details with the request.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Environment.forward_proxy_uri
	ForwardProxyURI *string `json:"forwardProxyURI,omitempty"`
	*/

	/* NOTYET: Add this once direct controller is implemented
	// Optional. NodeConfig of the environment.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Environment.node_config
	NodeConfig *GoogleCloudApigeeV1NodeConfig `json:"nodeConfig,omitempty"`
	*/

	// Optional. Key-value pairs that may be used for customizing the environment.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Environment.properties
	Properties *map[string]string `json:"properties,omitempty"`

	/* NOTYET: Add this once direct controller is implemented
	// Optional. EnvironmentType selected for the environment.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Environment.type
	Type *string `json:"type,omitempty"`
	*/
}

/* NOTYET: Add this once direct controller is implemented
// +kcc:proto=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1NodeConfig
type GoogleCloudApigeeV1NodeConfig struct {
	// Output only. The current total number of gateway nodes that each environment currently has across all instances.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1NodeConfig.current_aggregate_node_count
	CurrentAggregateNodeCount *int64 `json:"currentAggregateNodeCount,omitempty"`

	// Optional. The maximum total number of gateway nodes that the is reserved for all instances that has the specified environment. If not specified, the default is determined by the recommended maximum number of nodes for that gateway.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1NodeConfig.max_node_count
	MaxNodeCount *int64 `json:"maxNodeCount,omitempty"`

	// Optional. The minimum total number of gateway nodes that the is reserved for all instances that has the specified environment. If not specified, the default is determined by the recommended minimum number of nodes for that gateway.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1NodeConfig.min_node_count
	MinNodeCount *int64 `json:"minNodeCount,omitempty"`
}
*/

// ApigeeEnvironmentStatus defines the config connector machine state of ApigeeEnvironment
type ApigeeEnvironmentStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ApigeeEnvironment resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ApigeeEnvironmentObservedState `json:"observedState,omitempty"`

	// Output only. Creation time of this environment as milliseconds since epoch.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Environment.created_at
	CreatedAt *int64 `json:"createdAt,omitempty"`

	/* NOTYET: Perhaps add this once direct controller is implemented. Or, we may only add it to observedState.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Environment.has_attached_flow_hooks
	HasAttachedFlowHooks *bool `json:"hasAttachedFlowHooks,omitempty"`
	*/

	// Output only. Last modification time of this environment as milliseconds since epoch.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Environment.last_modified_at
	LastModifiedAt *int64 `json:"lastModifiedAt,omitempty"`

	// Output only. State of the environment. Values other than ACTIVE means the resource is not ready to use.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Environment.state
	State *string `json:"state,omitempty"`
}

// ApigeeEnvironmentObservedState is the state of the ApigeeEnvironment resource as most recently observed in GCP.
// +kcc:proto=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Environment
type ApigeeEnvironmentObservedState struct {
	/* NOTYET: Add this once direct controller is implemented
	// Output only. Creation time of this environment as milliseconds since epoch.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Environment.created_at
	CreatedAt *int64 `json:"createdAt,omitempty"`
	*/

	/* NOTYET: Add this once direct controller is implemented
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Environment.has_attached_flow_hooks
	HasAttachedFlowHooks *bool `json:"hasAttachedFlowHooks,omitempty"`
	*/

	/* NOTYET: Add this once direct controller is implemented
	// Output only. Last modification time of this environment as milliseconds since epoch.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Environment.last_modified_at
	LastModifiedAt *int64 `json:"lastModifiedAt,omitempty"`
	*/

	/* NOTYET: Add this once direct controller is implemented
	// Output only. State of the environment. Values other than ACTIVE means the resource is not ready to use.
	// +kcc:proto:field=mockgcp.cloud.apigee.v1.GoogleCloudApigeeV1Environment.state
	State *string `json:"state,omitempty"`
	*/
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpapigeeenvironment;gcpapigeeenvironments
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/dcl2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"
// +kubebuilder:storageversion

// ApigeeEnvironment is the Schema for the ApigeeEnvironment API
// +k8s:openapi-gen=true
type ApigeeEnvironment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ApigeeEnvironmentSpec   `json:"spec,omitempty"`
	Status ApigeeEnvironmentStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ApigeeEnvironmentList contains a list of ApigeeEnvironment
type ApigeeEnvironmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApigeeEnvironment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ApigeeEnvironment{}, &ApigeeEnvironmentList{})
}
