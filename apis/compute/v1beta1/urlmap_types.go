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

	// defaultRouteAction takes effect when none of the hostRules match. The load balancer performs advanced routing actions, such as URL rewrites and header transformations, before forwarding the request to the selected backend. Only one of defaultUrlRedirect, defaultService or defaultRouteAction.weightedBackendService can be set. URL maps for classic Application Load Balancers only support the urlRewrite action within defaultRouteAction. defaultRouteAction has no effect when the URL map is bound to a target gRPC proxy that has the validateForProxyless field set to true.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMap.default_route_action
	DefaultRouteAction *HTTPRouteAction `json:"defaultRouteAction,omitempty"`

	// The full or partial URL of the defaultService resource to which traffic is directed if none of the hostRules match. If defaultRouteAction is also specified, advanced routing actions, such as URL rewrites, take effect before sending the request to the backend. Only one of defaultUrlRedirect, defaultService or defaultRouteAction.weightedBackendService can be set. defaultService has no effect when the URL map is bound to a target gRPC proxy that has the validateForProxyless field set to true.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMap.default_service
	DefaultService *ComputeURLMapService `json:"defaultService,omitempty"`

	// When none of the specified hostRules match, the request is redirected to a URL specified by defaultUrlRedirect. Only one of defaultUrlRedirect, defaultService or defaultRouteAction.weightedBackendService can be set. Not supported when the URL map is bound to a target gRPC proxy.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMap.default_url_redirect
	DefaultURLRedirect *HTTPRedirectAction `json:"defaultURLRedirect,omitempty"`

	// An optional description of this resource. Provide this property when you create the resource.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMap.description
	Description *string `json:"description,omitempty"`

	// Specifies changes to request and response headers that need to take effect for the selected backendService. The headerAction specified here take effect after headerAction specified under pathMatcher. headerAction is not supported for load balancers that have their loadBalancingScheme set to EXTERNAL. Not supported when the URL map is bound to a target gRPC proxy that has validateForProxyless field set to true.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMap.header_action
	HeaderAction *HTTPHeaderAction `json:"headerAction,omitempty"`

	// The list of host rules to use against the URL.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMap.host_rules
	HostRule []HostRule `json:"hostRule,omitempty"`

	// The list of named PathMatchers to use against the URL.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMap.path_matchers
	PathMatcher []PathMatcher `json:"pathMatcher,omitempty"`

	// The list of expected URL mapping tests. Request to update this UrlMap will succeed only if all of the test cases pass. You can specify a maximum of 100 tests per UrlMap. Not supported when the URL map is bound to a target gRPC proxy.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMap.tests
	Test []URLMapTest `json:"test,omitempty"`
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

	/* [Output Only] Creation timestamp in RFC3339 text format. */
	// +kcc:proto:field=google.cloud.compute.v1.UrlMap.creation_timestamp
	CreationTimestamp *string `json:"creationTimestamp,omitempty"`

	/* Fingerprint of this resource. This field is used internally during updates of this resource. */
	// +kcc:proto:field=google.cloud.compute.v1.UrlMap.fingerprint
	Fingerprint *string `json:"fingerprint,omitempty"`

	/* [Output Only] The unique identifier for the resource. This identifier is defined by the server. */
	// +kcc:proto:field=google.cloud.compute.v1.UrlMap.id
	MapId *int64 `json:"mapId,omitempty"`

	/* [Output Only] Server-defined URL for this resource. */
	// +kcc:proto:field=google.cloud.compute.v1.UrlMap.self_link
	SelfLink *string `json:"selfLink,omitempty"`
}

// ComputeURLMapObservedState is the state of the ComputeURLMap resource as most recently observed in GCP.
// +kcc:observedstate:proto=google.cloud.compute.v1.UrlMap
type ComputeURLMapObservedState struct {
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputeurlmap;gcpcomputeurlmaps
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=stable"
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

type ComputeURLMapService struct {
	/* The ComputeBackendBucket resource to which traffic is directed. */
	BackendBucketRef *ComputeBackendBucketRef `json:"backendBucketRef,omitempty"`

	/* The ComputeBackendService resource to which traffic is directed. */
	BackendServiceRef *ComputeBackendServiceRef `json:"backendServiceRef,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.PathMatcher
type PathMatcher struct {
	// defaultRouteAction takes effect when none of the pathRules or routeRules match. The load balancer performs advanced routing actions, such as URL rewrites and header transformations, before forwarding the request to the selected backend. Only one of defaultUrlRedirect, defaultService or defaultRouteAction.weightedBackendService can be set. URL maps for classic Application Load Balancers only support the urlRewrite action within a path matcher's defaultRouteAction.
	// +kcc:proto:field=google.cloud.compute.v1.PathMatcher.default_route_action
	DefaultRouteAction *HTTPRouteAction `json:"defaultRouteAction,omitempty"`

	// The full or partial URL to the BackendService resource. This URL is used if none of the pathRules or routeRules defined by this PathMatcher are matched.
	// +kcc:proto:field=google.cloud.compute.v1.PathMatcher.default_service
	DefaultService *ComputeURLMapService `json:"defaultService,omitempty"`

	// When none of the specified pathRules or routeRules match, the request is redirected to a URL specified by defaultUrlRedirect. Only one of defaultUrlRedirect, defaultService or defaultRouteAction.weightedBackendService can be set. Not supported when the URL map is bound to a target gRPC proxy.
	// +kcc:proto:field=google.cloud.compute.v1.PathMatcher.default_url_redirect
	DefaultURLRedirect *HTTPRedirectAction `json:"defaultURLRedirect,omitempty"`

	// An optional description of this resource. Provide this property when you create the resource.
	// +kcc:proto:field=google.cloud.compute.v1.PathMatcher.description
	Description *string `json:"description,omitempty"`

	// Specifies changes to request and response headers that need to take effect for the selected backend service. HeaderAction specified here are applied after the matching HttpRouteRule HeaderAction and before the HeaderAction in the UrlMap HeaderAction is not supported for load balancers that have their loadBalancingScheme set to EXTERNAL. Not supported when the URL map is bound to a target gRPC proxy that has validateForProxyless field set to true.
	// +kcc:proto:field=google.cloud.compute.v1.PathMatcher.header_action
	HeaderAction *HTTPHeaderAction `json:"headerAction,omitempty"`

	// The name to which this PathMatcher is referred by the HostRule.
	// +kcc:proto:field=google.cloud.compute.v1.PathMatcher.name
	Name *string `json:"name,omitempty"`

	// The list of path rules. Use this list instead of routeRules when routing based on simple path matching is all that's required. The order by which path rules are specified does not matter. Matches are always done on the longest-path-first basis. For example: a pathRule with a path /a/b/c/* will match before /a/b/* irrespective of the order in which those paths appear in this list. Within a given pathMatcher, only one of pathRules or routeRules must be set.
	// +kcc:proto:field=google.cloud.compute.v1.PathMatcher.path_rules
	PathRule []PathRule `json:"pathRule,omitempty"`

	// The list of HTTP route rules. Use this list instead of pathRules when advanced route matching and routing actions are desired. routeRules are evaluated in order of priority, from the lowest to highest number. Within a given pathMatcher, you can set only one of pathRules or routeRules.
	// +kcc:proto:field=google.cloud.compute.v1.PathMatcher.route_rules
	RouteRules []HTTPRouteRule `json:"routeRules,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.PathRule
type PathRule struct {
	// The list of path patterns to match. Each must start with / and the only place a * is allowed is at the end following a /. The string fed to the path matcher does not include any text after the first ? or #, and those chars are not allowed here.
	// +kcc:proto:field=google.cloud.compute.v1.PathRule.paths
	Paths []string `json:"paths,omitempty"`

	// In response to a matching path, the load balancer performs advanced routing actions, such as URL rewrites and header transformations, before forwarding the request to the selected backend. Only one of urlRedirect, service or routeAction.weightedBackendService can be set. URL maps for classic Application Load Balancers only support the urlRewrite action within a path rule's routeAction.
	// +kcc:proto:field=google.cloud.compute.v1.PathRule.route_action
	RouteAction *HTTPRouteAction `json:"routeAction,omitempty"`

	// The full or partial URL of the backend service resource to which traffic is directed if this rule is matched. If routeAction is also specified, advanced routing actions, such as URL rewrites, take effect before sending the request to the backend. Only one of urlRedirect, service or routeAction.weightedBackendService can be set.
	// +kcc:proto:field=google.cloud.compute.v1.PathRule.service
	Service *ComputeURLMapService `json:"service,omitempty"`

	// When a path pattern is matched, the request is redirected to a URL specified by urlRedirect. Only one of urlRedirect, service or routeAction.weightedBackendService can be set. Not supported when the URL map is bound to a target gRPC proxy.
	// +kcc:proto:field=google.cloud.compute.v1.PathRule.url_redirect
	URLRedirect *HTTPRedirectAction `json:"urlRedirect,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.HttpRouteRule
type HTTPRouteRule struct {
	// The short description conveying the intent of this routeRule. The description can have a maximum length of 1024 characters.
	// +kcc:proto:field=google.cloud.compute.v1.HttpRouteRule.description
	Description *string `json:"description,omitempty"`

	// Specifies changes to request and response headers that need to take effect for the selected backendService. The headerAction value specified here is applied before the matching pathMatchers[].headerAction and after pathMatchers[].routeRules[].routeAction.weightedBackendService.backendServiceWeightAction[].headerAction HeaderAction is not supported for load balancers that have their loadBalancingScheme set to EXTERNAL. Not supported when the URL map is bound to a target gRPC proxy that has validateForProxyless field set to true.
	// +kcc:proto:field=google.cloud.compute.v1.HttpRouteRule.header_action
	HeaderAction *HTTPHeaderAction `json:"headerAction,omitempty"`

	// The list of criteria for matching attributes of a request to this routeRule. This list has OR semantics: the request matches this routeRule when any of the matchRules are satisfied. However predicates within a given matchRule have AND semantics. All predicates within a matchRule must match for the request to match the rule.
	// +kcc:proto:field=google.cloud.compute.v1.HttpRouteRule.match_rules
	MatchRules []HTTPRouteRuleMatch `json:"matchRules,omitempty"`

	// For routeRules within a given pathMatcher, priority determines the order in which a load balancer interprets routeRules. RouteRules are evaluated in order of priority, from the lowest to highest number. The priority of a rule decreases as its number increases (1, 2, 3, N+1). The first rule that matches the request is applied. You cannot configure two or more routeRules with the same priority. Priority for each rule must be set to a number from 0 to 2147483647 inclusive. Priority numbers can have gaps, which enable you to add or remove rules in the future without affecting the rest of the rules. For example, 1, 2, 3, 4, 5, 9, 12, 16 is a valid series of priority numbers to which you could add rules numbered from 6 to 8, 10 to 11, and 13 to 15 in the future without any impact on existing rules.
	// +kcc:proto:field=google.cloud.compute.v1.HttpRouteRule.priority
	Priority *int32 `json:"priority,omitempty"`

	// In response to a matching matchRule, the load balancer performs advanced routing actions, such as URL rewrites and header transformations, before forwarding the request to the selected backend. Only one of urlRedirect, service or routeAction.weightedBackendService can be set. URL maps for classic Application Load Balancers only support the urlRewrite action within a route rule's routeAction.
	// +kcc:proto:field=google.cloud.compute.v1.HttpRouteRule.route_action
	RouteAction *HTTPRouteAction `json:"routeAction,omitempty"`

	// The full or partial URL of the backend service resource to which traffic is directed if this rule is matched. If routeAction is also specified, advanced routing actions, such as URL rewrites, take effect before sending the request to the backend. Only one of urlRedirect, service or routeAction.weightedBackendService can be set.
	// +kcc:proto:field=google.cloud.compute.v1.HttpRouteRule.service
	Service *ComputeURLMapService `json:"service,omitempty"`

	// When this rule is matched, the request is redirected to a URL specified by urlRedirect. Only one of urlRedirect, service or routeAction.weightedBackendService can be set. Not supported when the URL map is bound to a target gRPC proxy.
	// +kcc:proto:field=google.cloud.compute.v1.HttpRouteRule.url_redirect
	URLRedirect *HTTPRedirectAction `json:"urlRedirect,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.HttpRouteAction
type HTTPRouteAction struct {
	// The specification for allowing client-side cross-origin requests. For more information about the W3C recommendation for cross-origin resource sharing (CORS), see Fetch API Living Standard. Not supported when the URL map is bound to a target gRPC proxy.
	// +kcc:proto:field=google.cloud.compute.v1.HttpRouteAction.cors_policy
	CorsPolicy *CorsPolicy `json:"corsPolicy,omitempty"`

	// The specification for fault injection introduced into traffic to test the resiliency of clients to backend service failure. As part of fault injection, when clients send requests to a backend service, delays can be introduced by a load balancer on a percentage of requests before sending those requests to the backend service. Similarly requests from clients can be aborted by the load balancer for a percentage of requests. timeout and retry_policy is ignored by clients that are configured with a fault_injection_policy if: 1. The traffic is generated by fault injection AND 2. The fault injection is not a delay fault injection. Fault injection is not supported with the classic Application Load Balancer . To see which load balancers support fault injection, see Load balancing: Routing and traffic management features.
	// +kcc:proto:field=google.cloud.compute.v1.HttpRouteAction.fault_injection_policy
	FaultInjectionPolicy *HTTPFaultInjection `json:"faultInjectionPolicy,omitempty"`

	// Specifies the policy on how requests intended for the route's backends are shadowed to a separate mirrored backend service. The load balancer does not wait for responses from the shadow service. Before sending traffic to the shadow service, the host / authority header is suffixed with -shadow. Not supported when the URL map is bound to a target gRPC proxy that has the validateForProxyless field set to true.
	// +kcc:proto:field=google.cloud.compute.v1.HttpRouteAction.request_mirror_policy
	RequestMirrorPolicy *RequestMirrorPolicy `json:"requestMirrorPolicy,omitempty"`

	// Specifies the retry policy associated with this route.
	// +kcc:proto:field=google.cloud.compute.v1.HttpRouteAction.retry_policy
	RetryPolicy *HTTPRetryPolicy `json:"retryPolicy,omitempty"`

	// Specifies the timeout for the selected route. Timeout is computed from the time the request has been fully processed (known as *end-of-stream*) up until the response has been processed. Timeout includes all retries. If not specified, this field uses the largest timeout among all backend services associated with the route. Not supported when the URL map is bound to a target gRPC proxy that has validateForProxyless field set to true.
	// +kcc:proto:field=google.cloud.compute.v1.HttpRouteAction.timeout
	Timeout *Duration `json:"timeout,omitempty"`

	// The spec to modify the URL of the request, before forwarding the request to the matched service. urlRewrite is the only action supported in UrlMaps for classic Application Load Balancers. Not supported when the URL map is bound to a target gRPC proxy that has the validateForProxyless field set to true.
	// +kcc:proto:field=google.cloud.compute.v1.HttpRouteAction.url_rewrite
	URLRewrite *URLRewrite `json:"urlRewrite,omitempty"`

	// A list of weighted backend services to send traffic to when a route match occurs. The weights determine the fraction of traffic that flows to their corresponding backend service. If all traffic needs to go to a single backend service, there must be one weightedBackendService with weight set to a non-zero number. After a backend service is identified and before forwarding the request to the backend service, advanced routing actions such as URL rewrites and header transformations are applied depending on additional settings specified in this HttpRouteAction.
	// +kcc:proto:field=google.cloud.compute.v1.HttpRouteAction.weighted_backend_services
	WeightedBackendServices []WeightedBackendService `json:"weightedBackendServices,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.UrlRewrite
type URLRewrite struct {
	// Before forwarding the request to the selected service, the request's host header is replaced with contents of hostRewrite. The value must be from 1 to 255 characters.
	// +kcc:proto:field=google.cloud.compute.v1.UrlRewrite.host_rewrite
	HostRewrite *string `json:"hostRewrite,omitempty"`

	// Before forwarding the request to the selected backend service, the matching portion of the request's path is replaced by pathPrefixRewrite. The value must be from 1 to 1024 characters.
	// +kcc:proto:field=google.cloud.compute.v1.UrlRewrite.path_prefix_rewrite
	PathPrefixRewrite *string `json:"pathPrefixRewrite,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.RequestMirrorPolicy
type RequestMirrorPolicy struct {
	// The full or partial URL to the BackendService resource being mirrored to.
	// +kcc:proto:field=google.cloud.compute.v1.RequestMirrorPolicy.backend_service
	BackendServiceRef *ComputeBackendServiceRef `json:"backendServiceRef,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.UrlMapTest
type URLMapTest struct {
	// Description of this test case.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMapTest.description
	Description *string `json:"description,omitempty"`

	// The expected output URL evaluated by the load balancer containing the scheme, host, path and query parameters.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMapTest.expected_output_url
	ExpectedOutputURL *string `json:"expectedOutputURL,omitempty"`

	// For rules with urlRedirect, the test passes only if expectedRedirectResponseCode matches the HTTP status code in load balancer's redirect response. expectedRedirectResponseCode cannot be set when service is set.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMapTest.expected_redirect_response_code
	ExpectedRedirectResponseCode *int32 `json:"expectedRedirectResponseCode,omitempty"`

	// HTTP headers for this request. If headers contains a host header, then host must also match the header value.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMapTest.headers
	Header []URLMapTestHeader `json:"header,omitempty"`

	// Host portion of the URL. If headers contains a host header, then host must also match the header value.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMapTest.host
	Host *string `json:"host,omitempty"`

	// Path portion of the URL.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMapTest.path
	Path *string `json:"path,omitempty"`

	// Expected BackendService or BackendBucket resource the given URL should be mapped to. The service field cannot be set if expectedRedirectResponseCode is set.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMapTest.service
	Service *ComputeURLMapService `json:"service,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.UrlMapTestHeader
type URLMapTestHeader struct {
	// Header name.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMapTestHeader.name
	Name *string `json:"name,omitempty"`

	// Header value.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMapTestHeader.value
	Value *string `json:"value,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.WeightedBackendService
type WeightedBackendService struct {
	// The full or partial URL to the default BackendService resource. Before forwarding the request to backendService, the load balancer applies any relevant headerActions specified as part of this backendServiceWeight.
	// +kcc:proto:field=google.cloud.compute.v1.WeightedBackendService.backend_service
	BackendServiceRef *ComputeBackendServiceRef `json:"backendServiceRef,omitempty"`

	// Specifies changes to request and response headers that need to take effect for the selected backendService. headerAction specified here take effect before headerAction in the enclosing HttpRouteRule, PathMatcher and UrlMap. headerAction is not supported for load balancers that have their loadBalancingScheme set to EXTERNAL. Not supported when the URL map is bound to a target gRPC proxy that has validateForProxyless field set to true.
	// +kcc:proto:field=google.cloud.compute.v1.WeightedBackendService.header_action
	HeaderAction *HTTPHeaderAction `json:"headerAction,omitempty"`

	// Specifies the fraction of traffic sent to a backend service, computed as weight / (sum of all weightedBackendService weights in routeAction) . The selection of a backend service is determined only for new traffic. Once a user's request has been directed to a backend service, subsequent requests are sent to the same backend service as determined by the backend service's session affinity policy. Don't configure session affinity if you're using weighted traffic splitting. If you do, the weighted traffic splitting configuration takes precedence. The value must be from 0 to 1000.
	// +kcc:proto:field=google.cloud.compute.v1.WeightedBackendService.weight
	Weight *uint32 `json:"weight,omitempty"`
}
