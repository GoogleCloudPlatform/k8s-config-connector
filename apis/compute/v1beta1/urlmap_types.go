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

// ComputeURLMapSpec defines the desired state of ComputeURLMap
// +kcc:spec:proto=google.cloud.compute.v1.UrlMap
type ComputeURLMapSpec struct {
	// The project that this resource belongs to.
	ProjectRef *refsv1beta1.ProjectRef `json:"projectRef"`

	// The location of this resource.
	Location string `json:"location"`

	// The ComputeURLMap name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// defaultCustomErrorResponsePolicy specifies how the Load Balancer returns error responses when BackendService or BackendBucket responds with an error.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.UrlMap.default_custom_error_response_policy
	DefaultCustomErrorResponsePolicy *CustomErrorResponsePolicy `json:"defaultCustomErrorResponsePolicy,omitempty"`

	// defaultRouteAction takes effect when none of the hostRules match.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.UrlMap.default_route_action
	DefaultRouteAction *ComputeURLMapHTTPRouteAction `json:"defaultRouteAction,omitempty"`

	// The defaultService resource to which traffic is directed if none of the hostRules match.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.UrlMap.default_service
	DefaultService *ComputeURLMapDefaultService `json:"defaultService,omitempty"`

	// When none of the specified hostRules match, the request is redirected to a URL specified by defaultUrlRedirect.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.UrlMap.default_url_redirect
	DefaultURLRedirect *HTTPRedirectAction `json:"defaultURLRedirect,omitempty"`

	// An optional description of this resource.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.UrlMap.description
	Description *string `json:"description,omitempty"`

	// Specifies changes to request and response headers that need to take effect for the selected backendService.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.UrlMap.header_action
	HeaderAction *HTTPHeaderAction `json:"headerAction,omitempty"`

	// The list of host rules to use against the URL.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.UrlMap.host_rules
	HostRules []ComputeURLMapHostRule `json:"hostRules,omitempty"`

	// The list of named PathMatchers to use against the URL.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.UrlMap.path_matchers
	PathMatchers []ComputeURLMapPathMatcher `json:"pathMatchers,omitempty"`

	// The list of expected URL mapping tests.
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.UrlMap.tests
	Tests []ComputeURLMapTest `json:"tests,omitempty"`
}

type ComputeURLMapDefaultService struct {
	// +optional
	BackendBucketRef *ComputeBackendBucketRef `json:"backendBucketRef,omitempty"`

	// +optional
	BackendServiceRef *ComputeBackendServiceRef `json:"backendServiceRef,omitempty"`
}

type ComputeURLMapHostRule struct {
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.HostRule.description
	Description *string `json:"description,omitempty"`

	// +required
	// +kcc:proto:field=google.cloud.compute.v1.HostRule.hosts
	Hosts []string `json:"hosts"`

	// +required
	// +kcc:proto:field=google.cloud.compute.v1.HostRule.path_matcher
	PathMatcher *string `json:"pathMatcher"`
}

type ComputeURLMapPathMatcher struct {
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.PathMatcher.default_custom_error_response_policy
	DefaultCustomErrorResponsePolicy *CustomErrorResponsePolicy `json:"defaultCustomErrorResponsePolicy,omitempty"`

	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.PathMatcher.default_route_action
	DefaultRouteAction *ComputeURLMapHTTPRouteAction `json:"defaultRouteAction,omitempty"`

	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.PathMatcher.default_service
	DefaultService *ComputeURLMapDefaultService `json:"defaultService,omitempty"`

	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.PathMatcher.default_url_redirect
	DefaultURLRedirect *HTTPRedirectAction `json:"defaultURLRedirect,omitempty"`

	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.PathMatcher.description
	Description *string `json:"description,omitempty"`

	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.PathMatcher.header_action
	HeaderAction *HTTPHeaderAction `json:"headerAction,omitempty"`

	// +required
	// +kcc:proto:field=google.cloud.compute.v1.PathMatcher.name
	Name *string `json:"name"`

	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.PathMatcher.path_rules
	PathRules []ComputeURLMapPathRule `json:"pathRules,omitempty"`

	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.PathMatcher.route_rules
	RouteRules []ComputeURLMapHTTPRouteRule `json:"routeRules,omitempty"`
}

type ComputeURLMapPathRule struct {
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.PathRule.custom_error_response_policy
	CustomErrorResponsePolicy *CustomErrorResponsePolicy `json:"customErrorResponsePolicy,omitempty"`

	// +required
	// +kcc:proto:field=google.cloud.compute.v1.PathRule.paths
	Paths []string `json:"paths"`

	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.PathRule.route_action
	RouteAction *ComputeURLMapHTTPRouteAction `json:"routeAction,omitempty"`

	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.PathRule.service
	Service *ComputeURLMapDefaultService `json:"service,omitempty"`

	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.PathRule.url_redirect
	URLRedirect *HTTPRedirectAction `json:"urlRedirect,omitempty"`
}

type ComputeURLMapHTTPRouteRule struct {
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.HttpRouteRule.custom_error_response_policy
	CustomErrorResponsePolicy *CustomErrorResponsePolicy `json:"customErrorResponsePolicy,omitempty"`

	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.HttpRouteRule.description
	Description *string `json:"description,omitempty"`

	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.HttpRouteRule.header_action
	HeaderAction *HTTPHeaderAction `json:"headerAction,omitempty"`

	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.HttpRouteRule.match_rules
	MatchRules []HTTPRouteRuleMatch `json:"matchRules,omitempty"`

	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.HttpRouteRule.priority
	Priority *int32 `json:"priority,omitempty"`

	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.HttpRouteRule.route_action
	RouteAction *ComputeURLMapHTTPRouteAction `json:"routeAction,omitempty"`

	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.HttpRouteRule.service
	Service *ComputeURLMapDefaultService `json:"service,omitempty"`

	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.HttpRouteRule.url_redirect
	URLRedirect *HTTPRedirectAction `json:"urlRedirect,omitempty"`
}

type ComputeURLMapHTTPRouteAction struct {
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.HttpRouteAction.cors_policy
	CorsPolicy *CorsPolicy `json:"corsPolicy,omitempty"`

	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.HttpRouteAction.fault_injection_policy
	FaultInjectionPolicy *HTTPFaultInjection `json:"faultInjectionPolicy,omitempty"`

	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.HttpRouteAction.max_stream_duration
	MaxStreamDuration *Duration `json:"maxStreamDuration,omitempty"`

	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.HttpRouteAction.request_mirror_policy
	RequestMirrorPolicy *ComputeURLMapRequestMirrorPolicy `json:"requestMirrorPolicy,omitempty"`

	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.HttpRouteAction.retry_policy
	RetryPolicy *HTTPRetryPolicy `json:"retryPolicy,omitempty"`

	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.HttpRouteAction.timeout
	Timeout *Duration `json:"timeout,omitempty"`

	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.HttpRouteAction.url_rewrite
	URLRewrite *URLRewrite `json:"urlRewrite,omitempty"`

	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.HttpRouteAction.weighted_backend_services
	WeightedBackendServices []ComputeURLMapWeightedBackendService `json:"weightedBackendServices,omitempty"`
}

type ComputeURLMapRequestMirrorPolicy struct {
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.RequestMirrorPolicy.backend_service
	BackendServiceRef *ComputeBackendServiceRef `json:"backendServiceRef"`
}

type ComputeURLMapWeightedBackendService struct {
	// +required
	// +kcc:proto:field=google.cloud.compute.v1.WeightedBackendService.backend_service
	BackendServiceRef *ComputeBackendServiceRef `json:"backendServiceRef"`

	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.WeightedBackendService.header_action
	HeaderAction *HTTPHeaderAction `json:"headerAction,omitempty"`

	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.WeightedBackendService.weight
	Weight *uint32 `json:"weight,omitempty"`
}

type ComputeURLMapTest struct {
	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.UrlMapTest.description
	Description *string `json:"description,omitempty"`

	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.UrlMapTest.expected_output_url
	ExpectedOutputURL *string `json:"expectedOutputURL,omitempty"`

	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.UrlMapTest.expected_redirect_response_code
	ExpectedRedirectResponseCode *int32 `json:"expectedRedirectResponseCode,omitempty"`

	// +optional
	// +kcc:proto:field=google.cloud.compute.v1.UrlMapTest.headers
	Headers []URLMapTestHeader `json:"headers,omitempty"`

	// +required
	// +kcc:proto:field=google.cloud.compute.v1.UrlMapTest.host
	Host *string `json:"host"`

	// +required
	// +kcc:proto:field=google.cloud.compute.v1.UrlMapTest.path
	Path *string `json:"path"`

	// +required
	// +kcc:proto:field=google.cloud.compute.v1.UrlMapTest.service
	Service *ComputeURLMapDefaultService `json:"service"`
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

	// Creation timestamp in RFC3339 text format.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMap.creation_timestamp
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	// The unique identifier for the resource. This identifier is defined by the server.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMap.id
	MapID *uint64 `json:"mapId,omitempty"`

	// Server-defined URL for the resource.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMap.self_link
	SelfLink *string `json:"selfLink,omitempty"`

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ComputeURLMapObservedState `json:"observedState,omitempty"`
}

// ComputeURLMapObservedState is the state of the ComputeURLMap resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.UrlMap
type ComputeURLMapObservedState struct {
	// Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMap.fingerprint
	Fingerprint *string `json:"fingerprint,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputeurlmap;gcpcomputeurlmaps
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
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
