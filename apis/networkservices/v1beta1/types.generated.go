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

package v1beta1

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
	// +required
	// +required
	// Required. Label name presented as key in xDS Node Metadata.
	// +kcc:proto:field=google.cloud.networkservices.v1.EndpointMatcher.MetadataLabelMatcher.MetadataLabels.label_name
	LabelName *string `json:"labelName,omitempty"`

	// +required
	// +required
	// Required. Label value presented as value corresponding to the above
	//  key, in xDS Node Metadata.
	// +kcc:proto:field=google.cloud.networkservices.v1.EndpointMatcher.MetadataLabelMatcher.MetadataLabels.label_value
	LabelValue *string `json:"labelValue,omitempty"`
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
