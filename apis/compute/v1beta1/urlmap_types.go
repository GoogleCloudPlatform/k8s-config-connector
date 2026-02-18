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

	// defaultCustomErrorResponsePolicy specifies how the Load Balancer returns error responses when BackendServiceor BackendBucket responds with an error. This policy takes effect at the load balancer level and applies only when no policy has been defined for the error code at lower levels like PathMatcher, RouteRule and PathRule within this UrlMap. For example, consider a UrlMap with the following configuration: - defaultCustomErrorResponsePolicy containing policies for responding to 5xx and 4xx errors - A PathMatcher configured for *.example.com has defaultCustomErrorResponsePolicy for 4xx. If a request for http://www.example.com/ encounters a 404, the policy in pathMatcher.defaultCustomErrorResponsePolicy will be enforced. When the request for http://www.example.com/ encounters a 502, the policy in UrlMap.defaultCustomErrorResponsePolicy will be enforced. When a request that does not match any host in *.example.com such as http://www.myotherexample.com/, encounters a 404, UrlMap.defaultCustomErrorResponsePolicy takes effect. When used in conjunction with defaultRouteAction.retryPolicy, retries take precedence. Only once all retries are exhausted, the defaultCustomErrorResponsePolicy is applied. While attempting a retry, if load balancer is successful in reaching the service, the defaultCustomErrorResponsePolicy is ignored and the response from the service is returned to the client. defaultCustomErrorResponsePolicy is supported only for global external Application Load Balancers.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMap.default_custom_error_response_policy
	DefaultCustomErrorResponsePolicy *CustomErrorResponsePolicy `json:"defaultCustomErrorResponsePolicy,omitempty"`

	// defaultRouteAction takes effect when none of the hostRules match. The load balancer performs advanced routing actions, such as URL rewrites and header transformations, before forwarding the request to the selected backend. Only one of defaultUrlRedirect, defaultService or defaultRouteAction.weightedBackendService can be set. URL maps for classic Application Load Balancers only support the urlRewrite action within defaultRouteAction. defaultRouteAction has no effect when the URL map is bound to a target gRPC proxy that has the validateForProxyless field set to true.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMap.default_route_action
	DefaultRouteAction *HTTPRouteAction `json:"defaultRouteAction,omitempty"`

	// The full or partial URL of the defaultService resource to which traffic is directed if none of the hostRules match. If defaultRouteAction is also specified, advanced routing actions, such as URL rewrites, take effect before sending the request to the backend. Only one of defaultUrlRedirect, defaultService or defaultRouteAction.weightedBackendService can be set. defaultService has no effect when the URL map is bound to a target gRPC proxy that has the validateForProxyless field set to true.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMap.default_service
	DefaultService *string `json:"defaultService,omitempty"`

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
	HostRules []HostRule `json:"hostRules,omitempty"`

	// The list of named PathMatchers to use against the URL.
	// +kcc:proto:field=google.cloud.compute.v1.UrlMap.path_matchers
	PathMatchers []PathMatcher `json:"pathMatchers,omitempty"`

	// The list of expected URL mapping tests. Request to update this UrlMap will succeed only if all of the test cases pass. You can specify a maximum of 100 tests per UrlMap. Not supported when the URL map is bound to a target gRPC proxy.
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
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcomputeurlmap;gcpcomputeurlmaps
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
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
