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

// +generated:types
// krm.group: compute.cnrm.cloud.google.com
// krm.version: v1beta1
// proto.service: google.cloud.compute.v1
// resource: ComputeFirewallPolicyRule:FirewallPolicyRule
// resource: ComputeForwardingRule:ForwardingRule
// resource: ComputeSubnetwork:Subnetwork
// resource: ComputeTargetTcpProxy:TargetTcpProxy

package v1beta1

// +kcc:proto=google.cloud.compute.v1.FirewallPolicyRuleSecureTag
type FirewallPolicyRuleSecureTag struct {
	// Name of the secure tag, created with TagManager's TagValue API.
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRuleSecureTag.name
	Name *string `json:"name,omitempty"`

	// Output only. [Output Only] State of the secure tag, either `EFFECTIVE` or
	//  `INEFFECTIVE`. A secure tag is `INEFFECTIVE` when it is deleted
	//  or its network is deleted.
	//  Check the State enum for the list of possible values.
	// +kcc:proto:field=google.cloud.compute.v1.FirewallPolicyRuleSecureTag.state
	State *string `json:"state,omitempty"`
}

// +kcc:proto=google.cloud.compute.v1.SubnetworkParams
type SubnetworkParams struct {
	// Tag keys/values directly bound to this resource. Tag keys and values have the same definition as resource manager tags. The field is allowed for INSERT only. The keys/values to set on the resource should be specified in either ID { : } or Namespaced format { : }. For example the following are valid inputs: * {"tagKeys/333" : "tagValues/444", "tagKeys/123" : "tagValues/456"} * {"123/environment" : "production", "345/abc" : "xyz"} Note: * Invalid combinations of ID & namespaced format is not supported. For instance: {"123/environment" : "tagValues/444"} is invalid.
	// +kcc:proto:field=google.cloud.compute.v1.SubnetworkParams.resource_manager_tags
	ResourceManagerTags map[string]string `json:"resourceManagerTags,omitempty"`
}
