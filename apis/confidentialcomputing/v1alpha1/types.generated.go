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


// +kcc:proto=google.cloud.confidentialcomputing.v1.Challenge
type Challenge struct {
}

// +kcc:proto=google.cloud.confidentialcomputing.v1.Challenge
type ChallengeObservedState struct {
	// Output only. The resource name for this Challenge in the format
	//  `projects/*/locations/*/challenges/*`
	// +kcc:proto:field=google.cloud.confidentialcomputing.v1.Challenge.name
	Name *string `json:"name,omitempty"`

	// Output only. The time at which this Challenge was created
	// +kcc:proto:field=google.cloud.confidentialcomputing.v1.Challenge.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. The time at which this Challenge will no longer be usable. It
	//  is also the expiration time for any tokens generated from this Challenge.
	// +kcc:proto:field=google.cloud.confidentialcomputing.v1.Challenge.expire_time
	ExpireTime *string `json:"expireTime,omitempty"`

	// Output only. Indicates if this challenge has been used to generate a token.
	// +kcc:proto:field=google.cloud.confidentialcomputing.v1.Challenge.used
	Used *bool `json:"used,omitempty"`

	// Output only. Identical to nonce, but as a string.
	// +kcc:proto:field=google.cloud.confidentialcomputing.v1.Challenge.tpm_nonce
	TpmNonce *string `json:"tpmNonce,omitempty"`
}
