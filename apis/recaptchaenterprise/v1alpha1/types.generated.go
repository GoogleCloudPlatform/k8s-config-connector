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


// +kcc:proto=google.cloud.recaptchaenterprise.v1.ChallengeMetrics
type ChallengeMetrics struct {
	// Count of reCAPTCHA checkboxes or badges rendered. This is mostly equivalent
	//  to a count of pageloads for pages that include reCAPTCHA.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.ChallengeMetrics.pageload_count
	PageloadCount *int64 `json:"pageloadCount,omitempty"`

	// Count of nocaptchas (successful verification without a challenge) issued.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.ChallengeMetrics.nocaptcha_count
	NocaptchaCount *int64 `json:"nocaptchaCount,omitempty"`

	// Count of submitted challenge solutions that were incorrect or otherwise
	//  deemed suspicious such that a subsequent challenge was triggered.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.ChallengeMetrics.failed_count
	FailedCount *int64 `json:"failedCount,omitempty"`

	// Count of nocaptchas (successful verification without a challenge) plus
	//  submitted challenge solutions that were correct and resulted in
	//  verification.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.ChallengeMetrics.passed_count
	PassedCount *int64 `json:"passedCount,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.Metrics
type Metrics struct {

	// Inclusive start time aligned to a day (UTC).
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.Metrics.start_time
	StartTime *string `json:"startTime,omitempty"`

	// Metrics are continuous and in order by dates, and in the granularity
	//  of day. All Key types should have score-based data.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.Metrics.score_metrics
	ScoreMetrics []ScoreMetrics `json:"scoreMetrics,omitempty"`

	// Metrics are continuous and in order by dates, and in the granularity
	//  of day. Only challenge-based keys (CHECKBOX, INVISIBLE) have
	//  challenge-based data.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.Metrics.challenge_metrics
	ChallengeMetrics []ChallengeMetrics `json:"challengeMetrics,omitempty"`
}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.ScoreDistribution
type ScoreDistribution struct {

	// TODO: unsupported map type with key int32 and value int64

}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.ScoreMetrics
type ScoreMetrics struct {
	// Aggregated score metrics for all traffic.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.ScoreMetrics.overall_metrics
	OverallMetrics *ScoreDistribution `json:"overallMetrics,omitempty"`

	// TODO: unsupported map type with key string and value message

}

// +kcc:proto=google.cloud.recaptchaenterprise.v1.Metrics
type MetricsObservedState struct {
	// Output only. Identifier. The name of the metrics, in the format
	//  `projects/{project}/keys/{key}/metrics`.
	// +kcc:proto:field=google.cloud.recaptchaenterprise.v1.Metrics.name
	Name *string `json:"name,omitempty"`
}
