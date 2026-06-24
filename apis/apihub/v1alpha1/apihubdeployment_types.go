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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var APIHubDeploymentGVK = GroupVersion.WithKind("APIHubDeployment")

// APIHubDeploymentSpec defines the desired state of APIHubDeployment
// +kcc:spec:proto=google.cloud.apihub.v1.Deployment
type APIHubDeploymentSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location *string `json:"location"`

	// The APIHubDeployment name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// Required. The display name of the deployment.
	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName,omitempty"`

	// Optional. The description of the deployment.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty"`

	// Optional. The documentation of the deployment.
	// +kubebuilder:validation:Optional
	Documentation *Documentation `json:"documentation,omitempty"`

	// Required. The type of deployment.
	//  This maps to the following system defined attribute:
	//  `projects/{project}/locations/{location}/attributes/system-deployment-type`
	//  attribute.
	//  The number of values for this attribute will be based on the
	//  cardinality of the attribute. The same can be retrieved via GetAttribute
	//  API. All values should be from the list of allowed values defined for the
	//  attribute.
	// +kubebuilder:validation:Required
	DeploymentTypeRef *APIHubAttributeValueRef `json:"deploymentTypeRef,omitempty"`

	// Required. A URI to the runtime resource. This URI can be used to manage the
	//  resource. For example, if the runtime resource is of type APIGEE_PROXY,
	//  then this field will contain the URI to the management UI of the proxy.
	// +kubebuilder:validation:Required
	ResourceURI *string `json:"resourceURI,omitempty"`

	// Required. The endpoints at which this deployment resource is listening for
	//  API requests. This could be a list of complete URIs, hostnames or an IP
	//  addresses.
	// +kubebuilder:validation:Required
	Endpoints []string `json:"endpoints,omitempty"`

	// Optional. The SLO for this deployment.
	//  This maps to the following system defined attribute:
	//  `projects/{project}/locations/{location}/attributes/system-slo`
	//  attribute.
	//  The number of values for this attribute will be based on the
	//  cardinality of the attribute. The same can be retrieved via GetAttribute
	//  API. All values should be from the list of allowed values defined for the
	//  attribute.
	// +kubebuilder:validation:Optional
	SloRef *APIHubAttributeValueRef `json:"sloRef,omitempty"`

	// Optional. The environment mapping to this deployment.
	//  This maps to the following system defined attribute:
	//  `projects/{project}/locations/{location}/attributes/system-environment`
	//  attribute.
	//  The number of values for this attribute will be based on the
	//  cardinality of the attribute. The same can be retrieved via GetAttribute
	//  API. All values should be from the list of allowed values defined for the
	//  attribute.
	// +kubebuilder:validation:Optional
	EnvironmentRef *APIHubAttributeValueRef `json:"environmentRef,omitempty"`
}

type APIHubAttributeValueRef struct {
	/* The `id` of an allowed value of an attribute, when not managed by Config Connector. */
	External string `json:"external,omitempty"`
	/* The `name` field of a `APIHubAttribute` resource. */
	Name string `json:"name,omitempty"`
	/* The `namespace` field of a `APIHubAttribute` resource. */
	Namespace string `json:"namespace,omitempty"`
}

// APIHubDeploymentStatus defines the config connector machine state of APIHubDeployment
type APIHubDeploymentStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the APIHubDeployment resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *APIHubDeploymentObservedState `json:"observedState,omitempty"`
}

// APIHubDeploymentObservedState is the state of the APIHubDeployment resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.apihub.v1.Deployment
type APIHubDeploymentObservedState struct {
	// Output only. The API versions linked to this deployment.
	//  Note: A particular deployment could be linked to multiple different API
	//  versions (of same or different APIs).
	APIVersions []string `json:"apiVersions,omitempty"`

	// Output only. The time at which the deployment was created.
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which the deployment was last updated.
	UpdateTime *string `json:"updateTime,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpapihubdeployment;gcpapihubdeployments
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// APIHubDeployment is the Schema for the APIHubDeployment API
// +k8s:openapi-gen=true
type APIHubDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   APIHubDeploymentSpec   `json:"spec,omitempty"`
	Status APIHubDeploymentStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// APIHubDeploymentList contains a list of APIHubDeployment
type APIHubDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []APIHubDeployment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&APIHubDeployment{}, &APIHubDeploymentList{})
}
