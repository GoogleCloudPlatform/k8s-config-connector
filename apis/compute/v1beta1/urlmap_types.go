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
	refsv1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ComputeURLMapGVK = GroupVersion.WithKind("ComputeURLMap")

// ComputeURLMapSpec defines the desired state of ComputeURLMap
// +kcc:spec:proto=google.cloud.compute.v1.UrlMap
type ComputeURLMapSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The ComputeURLMap name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// defaultCustomErrorResponsePolicy specifies how the Load Balancer returns error responses when BackendServiceor BackendBucket responds with an error.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMap.default_custom_error_response_policy
	DefaultCustomErrorResponsePolicy *CustomErrorResponsePolicy `json:"defaultCustomErrorResponsePolicy,omitempty"`

	// defaultRouteAction takes effect when none of the hostRules match.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMap.default_route_action
	DefaultRouteAction *HTTPRouteAction `json:"defaultRouteAction,omitempty"`

	// The full or partial URL to the BackendService resource. This URL is used if none of the hostRules match.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMap.default_service
	DefaultService *string `json:"defaultService,omitempty"`

	// When none of the specified hostRules match, the request is redirected to a URL specified by defaultUrlRedirect.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMap.default_url_redirect
	DefaultURLRedirect *HTTPRedirectAction `json:"defaultURLRedirect,omitempty"`

	// An optional description of this resource.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMap.description
	Description *string `json:"description,omitempty"`

	// Specifies changes to request and response headers that need to take effect for the selected backend service.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMap.header_action
	HeaderAction *HTTPHeaderAction `json:"headerAction,omitempty"`

	// The list of HostRules to use against the URL.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMap.host_rules
	HostRules []HostRule `json:"hostRules,omitempty"`

	// The list of named PathMatchers to use against the URL.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMap.path_matchers
	PathMatchers []PathMatcher `json:"pathMatchers,omitempty"`

	// The list of expected URL mappings.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMap.tests
	Tests []URLMapTest `json:"tests,omitempty"`
}

// ComputeURLMapStatus defines the config connector machine state of ComputeURLMap
type ComputeURLMapStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	// A unique specifier for the ComputeURLMap resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ComputeURLMapObservedState `json:"observedState,omitempty"`
}

// ComputeURLMapObservedState is the state of the ComputeURLMap resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.UrlMap
type ComputeURLMapObservedState struct {
	// Creation timestamp in RFC3339 text format.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMap.creation_timestamp
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// Fingerprint used for optimistic locking.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMap.fingerprint
	Fingerprint *string `json:"fingerprint,omitempty"`

	// The unique identifier for the resource.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMap.id
	URLMapId *uint64 `json:"urlMapId,omitempty"`

	// The self-link for the resource.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMap.self_link
	SelfLink *string `json:"selfLink,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputeurlmap;gcpcomputeurlmaps
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ComputeURLMap is the Schema for the ComputeURLMap API
// +k8s:openapi-gen=true
type ComputeURLMap struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ComputeURLMapSpec   `json:"spec,omitempty"`
	Status ComputeURLMapStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ComputeURLMapList contains a list of ComputeURLMap
type ComputeURLMapList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ComputeURLMap `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ComputeURLMap{}, &ComputeURLMapList{})
}
