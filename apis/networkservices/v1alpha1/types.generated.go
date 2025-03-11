// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

// +kcc:proto=google.cloud.networkservices.v1.EndpointMatcher
type EndpointMatcher struct {
	// The matcher is based on node metadata presented by xDS clients.
	// +kcc:proto:field=google.cloud.networkservices.v1.EndpointMatcher.metadata_label_matcher
	MetadataLabelMatcher *EndpointMatcher_MetadataLabelMatcher `json:"metadataLabelMatcher,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.EndpointMatcher.MetadataLabelMatcher
type EndpointMatcher_MetadataLabelMatcher struct {
	// Specifies how matching should be done.
	//
	//  Supported values are:
	//  MATCH_ANY: At least one of the Labels specified in the
	//    matcher should match the metadata presented by xDS client.
	//  MATCH_ALL: The metadata presented by the xDS client should
	//    contain all of the labels specified here.
	//
	//  The selection is determined based on the best match. For
	//  example, suppose there are three EndpointPolicy
	//  resources P1, P2 and P3 and if P1 has a the matcher as
	//  MATCH_ANY <A:1, B:1>, P2 has MATCH_ALL <A:1,B:1>, and P3 has
	//  MATCH_ALL <A:1,B:1,C:1>.
	//
	//  If a client with label <A:1> connects, the config from P1
	//  will be selected.
	//
	//  If a client with label <A:1,B:1> connects, the config from P2
	//  will be selected.
	//
	//  If a client with label <A:1,B:1,C:1> connects, the config
	//  from P3 will be selected.
	//
	//  If there is more than one best match, (for example, if a
	//  config P4 with selector <A:1,D:1> exists and if a client with
	//  label <A:1,B:1,D:1> connects), an error will be thrown.
	// +kcc:proto:field=google.cloud.networkservices.v1.EndpointMatcher.MetadataLabelMatcher.metadata_label_match_criteria
	MetadataLabelMatchCriteria *string `json:"metadataLabelMatchCriteria,omitempty"`

	// The list of label value pairs that must match labels in the
	//  provided metadata based on filterMatchCriteria This list can
	//  have at most 64 entries. The list can be empty if the match
	//  criteria is MATCH_ANY, to specify a wildcard match (i.e this
	//  matches any client).
	// +kcc:proto:field=google.cloud.networkservices.v1.EndpointMatcher.MetadataLabelMatcher.metadata_labels
	MetadataLabels []EndpointMatcher_MetadataLabelMatcher_MetadataLabels `json:"metadataLabels,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.EndpointMatcher.MetadataLabelMatcher.MetadataLabels
type EndpointMatcher_MetadataLabelMatcher_MetadataLabels struct {
	// Required. Label name presented as key in xDS Node Metadata.
	// +kcc:proto:field=google.cloud.networkservices.v1.EndpointMatcher.MetadataLabelMatcher.MetadataLabels.label_name
	LabelName *string `json:"labelName,omitempty"`

	// Required. Label value presented as value corresponding to the above
	//  key, in xDS Node Metadata.
	// +kcc:proto:field=google.cloud.networkservices.v1.EndpointMatcher.MetadataLabelMatcher.MetadataLabels.label_value
	LabelValue *string `json:"labelValue,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.EndpointPolicy
type EndpointPolicy struct {
	// Required. Name of the EndpointPolicy resource. It matches pattern
	//  `projects/{project}/locations/global/endpointPolicies/{endpoint_policy}`.
	// +kcc:proto:field=google.cloud.networkservices.v1.EndpointPolicy.name
	Name *string `json:"name,omitempty"`

	// Optional. Set of label tags associated with the EndpointPolicy resource.
	// +kcc:proto:field=google.cloud.networkservices.v1.EndpointPolicy.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Required. The type of endpoint policy. This is primarily used to validate
	//  the configuration.
	// +kcc:proto:field=google.cloud.networkservices.v1.EndpointPolicy.type
	Type *string `json:"type,omitempty"`

	// Optional. This field specifies the URL of AuthorizationPolicy resource that
	//  applies authorization policies to the inbound traffic at the
	//  matched endpoints. Refer to Authorization. If this field is not
	//  specified, authorization is disabled(no authz checks) for this
	//  endpoint.
	// +kcc:proto:field=google.cloud.networkservices.v1.EndpointPolicy.authorization_policy
	AuthorizationPolicy *string `json:"authorizationPolicy,omitempty"`

	// Required. A matcher that selects endpoints to which the policies should be
	//  applied.
	// +kcc:proto:field=google.cloud.networkservices.v1.EndpointPolicy.endpoint_matcher
	EndpointMatcher *EndpointMatcher `json:"endpointMatcher,omitempty"`

	// Optional. Port selector for the (matched) endpoints. If no port selector is
	//  provided, the matched config is applied to all ports.
	// +kcc:proto:field=google.cloud.networkservices.v1.EndpointPolicy.traffic_port_selector
	TrafficPortSelector *TrafficPortSelector `json:"trafficPortSelector,omitempty"`

	// Optional. A free-text description of the resource. Max length 1024
	//  characters.
	// +kcc:proto:field=google.cloud.networkservices.v1.EndpointPolicy.description
	Description *string `json:"description,omitempty"`

	// Optional. A URL referring to ServerTlsPolicy resource. ServerTlsPolicy is
	//  used to determine the authentication policy to be applied to terminate the
	//  inbound traffic at the identified backends. If this field is not set,
	//  authentication is disabled(open) for this endpoint.
	// +kcc:proto:field=google.cloud.networkservices.v1.EndpointPolicy.server_tls_policy
	ServerTLSPolicy *string `json:"serverTLSPolicy,omitempty"`

	// Optional. A URL referring to a ClientTlsPolicy resource. ClientTlsPolicy
	//  can be set to specify the authentication for traffic from the proxy to the
	//  actual endpoints. More specifically, it is applied to the outgoing traffic
	//  from the proxy to the endpoint. This is typically used for sidecar model
	//  where the proxy identifies itself as endpoint to the control plane, with
	//  the connection between sidecar and endpoint requiring authentication. If
	//  this field is not set, authentication is disabled(open). Applicable only
	//  when EndpointPolicyType is SIDECAR_PROXY.
	// +kcc:proto:field=google.cloud.networkservices.v1.EndpointPolicy.client_tls_policy
	ClientTLSPolicy *string `json:"clientTLSPolicy,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.TrafficPortSelector
type TrafficPortSelector struct {
	// Optional. A list of ports. Can be port numbers or port range
	//  (example, [80-90] specifies all ports from 80 to 90, including
	//  80 and 90) or named ports or * to specify all ports. If the
	//  list is empty, all ports are selected.
	// +kcc:proto:field=google.cloud.networkservices.v1.TrafficPortSelector.ports
	Ports []string `json:"ports,omitempty"`
}

// +kcc:proto=google.cloud.networkservices.v1.EndpointPolicy
type EndpointPolicyObservedState struct {
	// Output only. The timestamp when the resource was created.
	// +kcc:proto:field=google.cloud.networkservices.v1.EndpointPolicy.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The timestamp when the resource was updated.
	// +kcc:proto:field=google.cloud.networkservices.v1.EndpointPolicy.update_time
	UpdateTime *string `json:"updateTime,omitempty"`
}
