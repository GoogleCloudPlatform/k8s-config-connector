// Copyright 2026 Google LLC
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
// krm.group: parametermanager.cnrm.cloud.google.com
// krm.version: v1alpha1
// proto.service: google.cloud.parametermanager.v1
// resource: ParameterManagerParameter:Parameter

package v1alpha1

// +kcc:proto=google.iam.v1.ResourcePolicyMember
type ResourcePolicyMember struct {
}

// +kcc:observedstate:proto=google.iam.v1.ResourcePolicyMember
type ResourcePolicyMemberObservedState struct {
	// IAM policy binding member referring to a Google Cloud resource by
	//  user-assigned name (https://google.aip.dev/122). If a resource is deleted
	//  and recreated with the same name, the binding will be applicable to the new
	//  resource.
	//
	//  Example:
	//  `principal://parametermanager.googleapis.com/projects/12345/name/locations/us-central1-a/parameters/my-parameter`
	// +kcc:proto:field=google.iam.v1.ResourcePolicyMember.iam_policy_name_principal
	IAMPolicyNamePrincipal *string `json:"iamPolicyNamePrincipal,omitempty"`

	// IAM policy binding member referring to a Google Cloud resource by
	//  system-assigned unique identifier (https://google.aip.dev/148#uid). If a
	//  resource is deleted and recreated with the same name, the binding will not
	//  be applicable to the new resource
	//
	//  Example:
	//  `principal://parametermanager.googleapis.com/projects/12345/uid/locations/us-central1-a/parameters/a918fed5`
	// +kcc:proto:field=google.iam.v1.ResourcePolicyMember.iam_policy_uid_principal
	IAMPolicyUidPrincipal *string `json:"iamPolicyUidPrincipal,omitempty"`
}
