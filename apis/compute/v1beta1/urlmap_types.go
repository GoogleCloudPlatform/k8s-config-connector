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

	// defaultCustomErrorResponsePolicy specifies how the Load Balancer returns error responses when BackendServiceor BackendBucket responds with an error. This policy takes effect at the load balancer level and applies only when no policy has been defined for the error code at lower levels like PathMatcher, RouteRule and PathRule within this UrlMap.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMap.default_custom_error_response_policy
	DefaultCustomErrorResponsePolicy *CustomErrorResponsePolicy `json:"defaultCustomErrorResponsePolicy,omitempty"`

	// defaultRouteAction takes effect when none of the hostRules match. The load balancer performs advanced routing actions, such as URL rewrites and header transformations, before forwarding the request to the selected backend.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMap.default_route_action
	DefaultRouteAction *HTTPRouteAction `json:"defaultRouteAction,omitempty"`

	// The full or partial URL of the defaultService resource to which traffic is directed if none of the hostRules match.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMap.default_service
	DefaultService *ComputeURLMapServiceRef `json:"defaultService,omitempty"`

	// When none of the specified hostRules match, the request is redirected to a URL specified by defaultUrlRedirect.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMap.default_url_redirect
	DefaultURLRedirect *HTTPRedirectAction `json:"defaultURLRedirect,omitempty"`

	// An optional description of this resource. Provide this property when you create the resource.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMap.description
	Description *string `json:"description,omitempty"`

	// Specifies changes to request and response headers that need to take effect for the selected backendService.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMap.header_action
	HeaderAction *HTTPHeaderAction `json:"headerAction,omitempty"`

	// The list of host rules to use against the URL.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMap.host_rules
	HostRules []HostRule `json:"hostRules,omitempty"`

	// The list of named PathMatchers to use against the URL.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMap.path_matchers
	PathMatchers []PathMatcher `json:"pathMatchers,omitempty"`

	// The list of expected URL mapping tests. Request to update the UrlMap succeeds only if all test cases pass.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMap.tests
	Tests []URLMapTest `json:"tests,omitempty"`
}

type ComputeURLMapServiceRef struct {
	BackendBucketRef  *ComputeBackendBucketRef  `json:"backendBucketRef,omitempty"`
	BackendServiceRef *ComputeBackendServiceRef `json:"backendServiceRef,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.PathMatcher
type PathMatcher struct {
	// defaultCustomErrorResponsePolicy specifies how the Load Balancer returns error responses when BackendServiceor BackendBucket responds with an error.
	// +kcc:proto:field=google.cloud.compute.v1.PathMatcher.default_custom_error_response_policy
	DefaultCustomErrorResponsePolicy *CustomErrorResponsePolicy `json:"defaultCustomErrorResponsePolicy,omitempty"`

	// defaultRouteAction takes effect when none of the pathRules or routeRules match.
	// +kcc:proto:field=google.cloud.compute.v1.PathMatcher.default_route_action
	DefaultRouteAction *HTTPRouteAction `json:"defaultRouteAction,omitempty"`

	// The full or partial URL to the BackendService resource.
	// +kcc:proto:field=google.cloud.compute.v1.PathMatcher.default_service
	DefaultService *ComputeURLMapServiceRef `json:"defaultService,omitempty"`

	// When none of the specified pathRules or routeRules match, the request is redirected to a URL specified by defaultUrlRedirect.
	// +kcc:proto:field=google.cloud.compute.v1.PathMatcher.default_url_redirect
	DefaultURLRedirect *HTTPRedirectAction `json:"defaultURLRedirect,omitempty"`

	// An optional description of this resource.
	// +kcc:proto:field=google.cloud.compute.v1.PathMatcher.description
	Description *string `json:"description,omitempty"`

	// Specifies changes to request and response headers that need to take effect for the selected backend service.
	// +kcc:proto:field=google.cloud.compute.v1.PathMatcher.header_action
	HeaderAction *HTTPHeaderAction `json:"headerAction,omitempty"`

	// The name to which this PathMatcher is referred by the HostRule.
	// +kcc:proto:field=google.cloud.compute.v1.PathMatcher.name
	Name *string `json:"name,omitempty"`

	// The list of path rules.
	// +kcc:proto:field=google.cloud.compute.v1.PathMatcher.path_rules
	PathRules []PathRule `json:"pathRules,omitempty"`

	// The list of HTTP route rules.
	// +kcc:proto:field=google.cloud.compute.v1.PathMatcher.route_rules
	RouteRules []HTTPRouteRule `json:"routeRules,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.PathRule
type PathRule struct {
	// customErrorResponsePolicy specifies how the Load Balancer returns error responses when BackendServiceor BackendBucket responds with an error.
	// +kcc:proto:field=google.cloud.compute.v1.PathRule.custom_error_response_policy
	CustomErrorResponsePolicy *CustomErrorResponsePolicy `json:"customErrorResponsePolicy,omitempty"`

	// The list of path patterns to match.
	// +kcc:proto:field=google.cloud.compute.v1.PathRule.paths
	Paths []string `json:"paths,omitempty"`

	// In response to a matching path, the load balancer performs advanced routing actions, such as URL rewrites and header transformations, before forwarding the request to the selected backend.
	// +kcc:proto:field=google.cloud.compute.v1.PathRule.route_action
	RouteAction *HTTPRouteAction `json:"routeAction,omitempty"`

	// The full or partial URL of the backend service resource to which traffic is directed if this rule is matched.
	// +kcc:proto:field=google.cloud.compute.v1.PathRule.service
	Service *ComputeURLMapServiceRef `json:"service,omitempty"`

	// When a path pattern is matched, the request is redirected to a URL specified by urlRedirect.
	// +kcc:proto:field=google.cloud.compute.v1.PathRule.url_redirect
	URLRedirect *HTTPRedirectAction `json:"urlRedirect,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.HttpRouteRule
type HTTPRouteRule struct {
	// customErrorResponsePolicy specifies how the Load Balancer returns error responses when BackendServiceor BackendBucket responds with an error.
	// +kcc:proto:field=google.cloud.compute.v1.HttpRouteRule.custom_error_response_policy
	CustomErrorResponsePolicy *CustomErrorResponsePolicy `json:"customErrorResponsePolicy,omitempty"`

	// The short description conveying the intent of this routeRule.
	// +kcc:proto:field=google.cloud.compute.v1.HttpRouteRule.description
	Description *string `json:"description,omitempty"`

	// Specifies changes to request and response headers that need to take effect for the selected backendService.
	// +kcc:proto:field=google.cloud.compute.v1.HttpRouteRule.header_action
	HeaderAction *HTTPHeaderAction `json:"headerAction,omitempty"`

	// The list of criteria for matching attributes of a request to this routeRule.
	// +kcc:proto:field=google.cloud.compute.v1.HttpRouteRule.match_rules
	MatchRules []HTTPRouteRuleMatch `json:"matchRules,omitempty"`

	// For routeRules within a given pathMatcher, priority determines the order in which a load balancer interprets routeRules.
	// +kcc:proto:field=google.cloud.compute.v1.HttpRouteRule.priority
	Priority *int32 `json:"priority,omitempty"`

	// In response to a matching matchRule, the load balancer performs advanced routing actions, such as URL rewrites and header transformations, before forwarding the request to the selected backend.
	// +kcc:proto:field=google.cloud.compute.v1.HttpRouteRule.route_action
	RouteAction *HTTPRouteAction `json:"routeAction,omitempty"`

	// The full or partial URL of the backend service resource to which traffic is directed if this rule is matched.
	// +kcc:proto:field=google.cloud.compute.v1.HttpRouteRule.service
	Service *ComputeURLMapServiceRef `json:"service,omitempty"`

	// When this rule is matched, the request is redirected to a URL specified by urlRedirect.
	// +kcc:proto:field=google.cloud.compute.v1.HttpRouteRule.url_redirect
	URLRedirect *HTTPRedirectAction `json:"urlRedirect,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.WeightedBackendService
type WeightedBackendService struct {
	// The full or partial URL to the default BackendService resource.
	// +kcc:proto:field=google.cloud.compute.v1.WeightedBackendService.backend_service
	BackendService *ComputeBackendServiceRef `json:"backendServiceRef,omitempty"`

	// Specifies changes to request and response headers that need to take effect for the selected backendService.
	// +kcc:proto:field=google.cloud.compute.v1.WeightedBackendService.header_action
	HeaderAction *HTTPHeaderAction `json:"headerAction,omitempty"`

	// Specifies the fraction of traffic sent to a backend service.
	// +kcc:proto:field=google.cloud.compute.v1.WeightedBackendService.weight
	Weight *uint32 `json:"weight,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.UrlMapTest
type URLMapTest struct {
	// Description of this test case.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMapTest.description
	Description *string `json:"description,omitempty"`

	// The expected output URL evaluated by the load balancer containing the scheme, host, path and query parameters.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMapTest.expected_output_url
	ExpectedOutputURL *string `json:"expectedOutputURL,omitempty"`

	// For rules with urlRedirect, the test passes only if expectedRedirectResponseCode matches the HTTP status code in load balancer's redirect response.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMapTest.expected_redirect_response_code
	ExpectedRedirectResponseCode *int32 `json:"expectedRedirectResponseCode,omitempty"`

	// HTTP headers for this request.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMapTest.headers
	Headers []URLMapTestHeader `json:"headers,omitempty"`

	// Host portion of the URL.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMapTest.host
	Host *string `json:"host,omitempty"`

	// Path portion of the URL.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMapTest.path
	Path *string `json:"path,omitempty"`

	// Expected BackendService or BackendBucket resource the given URL should be mapped to.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMapTest.service
	Service *ComputeURLMapServiceRef `json:"service,omitempty"`
}

// ComputeURLMapStatus defines the config connector machine state of ComputeURLMap
// +kcc:status:proto=google.cloud.compute.v1.UrlMap
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

	// [Output Only] Creation timestamp in RFC3339 text format.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMap.creation_timestamp
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// Fingerprint of this resource. A hash of the contents stored in this object.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMap.fingerprint
	Fingerprint *string `json:"fingerprint,omitempty"`

	// [Output Only] The unique identifier for the resource. This identifier is defined by the server.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMap.id
	ID *uint64 `json:"id,omitempty"`

	// [Output Only] Server-defined URL for the resource.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMap.self_link
	SelfLink *string `json:"selfLink,omitempty"`
}

// ComputeURLMapObservedState is the state of the ComputeURLMap resource as most recently observed in GCP.
type ComputeURLMapObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputeurlmap;gcpcomputeurlmaps
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
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
