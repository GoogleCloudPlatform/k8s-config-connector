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


// +kcc:proto=google.cloud.bigquery.datapolicies.v1.DataMaskingPolicy
type DataMaskingPolicy struct {
	// A predefined masking expression.
	// +kcc:proto:field=google.cloud.bigquery.datapolicies.v1.DataMaskingPolicy.predefined_expression
	PredefinedExpression *string `json:"predefinedExpression,omitempty"`

	// The name of the BigQuery routine that contains the custom masking
	//  routine, in the format of
	//  `projects/{project_number}/datasets/{dataset_id}/routines/{routine_id}`.
	// +kcc:proto:field=google.cloud.bigquery.datapolicies.v1.DataMaskingPolicy.routine
	Routine *string `json:"routine,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.datapolicies.v1.DataPolicy
type DataPolicy struct {
	// Policy tag resource name, in the format of
	//  `projects/{project_number}/locations/{location_id}/taxonomies/{taxonomy_id}/policyTags/{policyTag_id}`.
	// +kcc:proto:field=google.cloud.bigquery.datapolicies.v1.DataPolicy.policy_tag
	PolicyTag *string `json:"policyTag,omitempty"`

	// The data masking policy that specifies the data masking rule to use.
	// +kcc:proto:field=google.cloud.bigquery.datapolicies.v1.DataPolicy.data_masking_policy
	DataMaskingPolicy *DataMaskingPolicy `json:"dataMaskingPolicy,omitempty"`

	// Type of data policy.
	// +kcc:proto:field=google.cloud.bigquery.datapolicies.v1.DataPolicy.data_policy_type
	DataPolicyType *string `json:"dataPolicyType,omitempty"`

	// User-assigned (human readable) ID of the data policy that needs to be
	//  unique within a project. Used as {data_policy_id} in part of the resource
	//  name.
	// +kcc:proto:field=google.cloud.bigquery.datapolicies.v1.DataPolicy.data_policy_id
	DataPolicyID *string `json:"dataPolicyID,omitempty"`
}

// +kcc:proto=google.cloud.bigquery.datapolicies.v1.DataPolicy
type DataPolicyObservedState struct {
	// Output only. Resource name of this data policy, in the format of
	//  `projects/{project_number}/locations/{location_id}/dataPolicies/{data_policy_id}`.
	// +kcc:proto:field=google.cloud.bigquery.datapolicies.v1.DataPolicy.name
	Name *string `json:"name,omitempty"`
}
