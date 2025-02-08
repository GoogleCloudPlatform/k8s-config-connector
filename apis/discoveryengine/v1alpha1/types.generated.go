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


// +kcc:proto=google.cloud.discoveryengine.v1beta.SiteVerificationInfo
type SiteVerificationInfo struct {
	// Site verification state indicating the ownership and validity.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SiteVerificationInfo.site_verification_state
	SiteVerificationState *string `json:"siteVerificationState,omitempty"`

	// Latest site verification time.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.SiteVerificationInfo.verify_time
	VerifyTime *string `json:"verifyTime,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.TargetSite
type TargetSite struct {

	// Required. Input only. The user provided URI pattern from which the
	//  `generated_uri_pattern` is generated.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.TargetSite.provided_uri_pattern
	ProvidedURIPattern *string `json:"providedURIPattern,omitempty"`

	// The type of the target site, e.g., whether the site is to be included or
	//  excluded.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.TargetSite.type
	Type *string `json:"type,omitempty"`

	// Input only. If set to false, a uri_pattern is generated to include all
	//  pages whose address contains the provided_uri_pattern. If set to true, an
	//  uri_pattern is generated to try to be an exact match of the
	//  provided_uri_pattern or just the specific page if the provided_uri_pattern
	//  is a specific one. provided_uri_pattern is always normalized to
	//  generate the URI pattern to be used by the search engine.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.TargetSite.exact_match
	ExactMatch *bool `json:"exactMatch,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.TargetSite.FailureReason
type TargetSite_FailureReason struct {
	// Failed due to insufficient quota.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.TargetSite.FailureReason.quota_failure
	QuotaFailure *TargetSite_FailureReason_QuotaFailure `json:"quotaFailure,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.TargetSite.FailureReason.QuotaFailure
type TargetSite_FailureReason_QuotaFailure struct {
	// This number is an estimation on how much total quota this project needs
	//  to successfully complete indexing.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.TargetSite.FailureReason.QuotaFailure.total_required_quota
	TotalRequiredQuota *int64 `json:"totalRequiredQuota,omitempty"`
}

// +kcc:proto=google.cloud.discoveryengine.v1beta.TargetSite
type TargetSiteObservedState struct {
	// Output only. The fully qualified resource name of the target site.
	//  `projects/{project}/locations/{location}/collections/{collection}/dataStores/{data_store}/siteSearchEngine/targetSites/{target_site}`
	//  The `target_site_id` is system-generated.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.TargetSite.name
	Name *string `json:"name,omitempty"`

	// Output only. This is system-generated based on the provided_uri_pattern.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.TargetSite.generated_uri_pattern
	GeneratedURIPattern *string `json:"generatedURIPattern,omitempty"`

	// Output only. Root domain of the provided_uri_pattern.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.TargetSite.root_domain_uri
	RootDomainURI *string `json:"rootDomainURI,omitempty"`

	// Output only. Site ownership and validity verification status.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.TargetSite.site_verification_info
	SiteVerificationInfo *SiteVerificationInfo `json:"siteVerificationInfo,omitempty"`

	// Output only. Indexing status.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.TargetSite.indexing_status
	IndexingStatus *string `json:"indexingStatus,omitempty"`

	// Output only. The target site's last updated time.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.TargetSite.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Failure reason.
	// +kcc:proto:field=google.cloud.discoveryengine.v1beta.TargetSite.failure_reason
	FailureReason *TargetSite_FailureReason `json:"failureReason,omitempty"`
}
