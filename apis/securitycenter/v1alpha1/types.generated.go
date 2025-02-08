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


// +kcc:proto=google.cloud.securitycenter.v1beta1.SecurityMarks
type SecurityMarks struct {
	// The relative resource name of the SecurityMarks. See:
	//  https://cloud.google.com/apis/design/resource_names#relative_resource_name
	//  Examples:
	//  "organizations/{organization_id}/assets/{asset_id}/securityMarks"
	//  "organizations/{organization_id}/sources/{source_id}/findings/{finding_id}/securityMarks".
	// +kcc:proto:field=google.cloud.securitycenter.v1beta1.SecurityMarks.name
	Name *string `json:"name,omitempty"`

	// Mutable user specified security marks belonging to the parent resource.
	//  Constraints are as follows:
	//
	//    * Keys and values are treated as case insensitive
	//    * Keys must be between 1 - 256 characters (inclusive)
	//    * Keys must be letters, numbers, underscores, or dashes
	//    * Values have leading and trailing whitespace trimmed, remaining
	//      characters must be between 1 - 4096 characters (inclusive)
	// +kcc:proto:field=google.cloud.securitycenter.v1beta1.SecurityMarks.marks
	Marks map[string]string `json:"marks,omitempty"`
}
