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


// +kcc:proto=google.iam.v1beta.WorkloadIdentityPool
type WorkloadIdentityPool struct {

	// A display name for the pool. Cannot exceed 32 characters.
	// +kcc:proto:field=google.iam.v1beta.WorkloadIdentityPool.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// A description of the pool. Cannot exceed 256 characters.
	// +kcc:proto:field=google.iam.v1beta.WorkloadIdentityPool.description
	Description *string `json:"description,omitempty"`

	// Whether the pool is disabled. You cannot use a disabled pool to exchange
	//  tokens, or use existing tokens to access resources. If
	//  the pool is re-enabled, existing tokens grant access again.
	// +kcc:proto:field=google.iam.v1beta.WorkloadIdentityPool.disabled
	Disabled *bool `json:"disabled,omitempty"`
}

// +kcc:proto=google.iam.v1beta.WorkloadIdentityPool
type WorkloadIdentityPoolObservedState struct {
	// Output only. The resource name of the pool.
	// +kcc:proto:field=google.iam.v1beta.WorkloadIdentityPool.name
	Name *string `json:"name,omitempty"`

	// Output only. The state of the pool.
	// +kcc:proto:field=google.iam.v1beta.WorkloadIdentityPool.state
	State *string `json:"state,omitempty"`
}
