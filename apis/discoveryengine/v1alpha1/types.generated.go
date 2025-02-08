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


// +kcc:proto=google.cloud.discoveryengine.v1.SiteVerificationInfo
type SiteVerificationInfo struct {
	// Site verification state indicating the ownership and validity.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.SiteVerificationInfo.site_verification_state
	SiteVerificationState *string `json:"siteVerificationState,omitempty"`

	// Latest site verification time.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.SiteVerificationInfo.verify_time
	VerifyTime *string `json:"verifyTime,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1.TargetSite.FailureReason
type TargetSite_FailureReason struct {
	// Failed due to insufficient quota.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.TargetSite.FailureReason.quota_failure
	QuotaFailure *TargetSite_FailureReason_QuotaFailure `json:"quotaFailure,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1.TargetSite.FailureReason.QuotaFailure
type TargetSite_FailureReason_QuotaFailure struct {
	// This number is an estimation on how much total quota this project needs
	//  to successfully complete indexing.
	// +kcc:proto:field=google.cloud.discoveryengine.v1.TargetSite.FailureReason.QuotaFailure.total_required_quota
	TotalRequiredQuota *int64 `json:"totalRequiredQuota,omitempty"`
}
