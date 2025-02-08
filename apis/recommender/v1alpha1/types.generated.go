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


// +kcc:proto=google.cloud.recommender.v1beta1.Insight
type Insight struct {
	// Name of the insight.
	// +kcc:proto:field=google.cloud.recommender.v1beta1.Insight.name
	Name *string `json:"name,omitempty"`

	// Free-form human readable summary in English. The maximum length is 500
	//  characters.
	// +kcc:proto:field=google.cloud.recommender.v1beta1.Insight.description
	Description *string `json:"description,omitempty"`

	// Fully qualified resource names that this insight is targeting.
	// +kcc:proto:field=google.cloud.recommender.v1beta1.Insight.target_resources
	TargetResources []string `json:"targetResources,omitempty"`

	// Insight subtype. Insight content schema will be stable for a given subtype.
	// +kcc:proto:field=google.cloud.recommender.v1beta1.Insight.insight_subtype
	InsightSubtype *string `json:"insightSubtype,omitempty"`

	// A struct of custom fields to explain the insight.
	//  Example: "grantedPermissionsCount": "1000"
	// +kcc:proto:field=google.cloud.recommender.v1beta1.Insight.content
	Content map[string]string `json:"content,omitempty"`

	// Timestamp of the latest data used to generate the insight.
	// +kcc:proto:field=google.cloud.recommender.v1beta1.Insight.last_refresh_time
	LastRefreshTime *string `json:"lastRefreshTime,omitempty"`

	// Observation period that led to the insight. The source data used to
	//  generate the insight ends at last_refresh_time and begins at
	//  (last_refresh_time - observation_period).
	// +kcc:proto:field=google.cloud.recommender.v1beta1.Insight.observation_period
	ObservationPeriod *string `json:"observationPeriod,omitempty"`

	// Information state and metadata.
	// +kcc:proto:field=google.cloud.recommender.v1beta1.Insight.state_info
	StateInfo *InsightStateInfo `json:"stateInfo,omitempty"`

	// Category being targeted by the insight.
	// +kcc:proto:field=google.cloud.recommender.v1beta1.Insight.category
	Category *string `json:"category,omitempty"`

	// Insight's severity.
	// +kcc:proto:field=google.cloud.recommender.v1beta1.Insight.severity
	Severity *string `json:"severity,omitempty"`

	// Fingerprint of the Insight. Provides optimistic locking when updating
	//  states.
	// +kcc:proto:field=google.cloud.recommender.v1beta1.Insight.etag
	Etag *string `json:"etag,omitempty"`

	// Recommendations derived from this insight.
	// +kcc:proto:field=google.cloud.recommender.v1beta1.Insight.associated_recommendations
	AssociatedRecommendations []Insight_RecommendationReference `json:"associatedRecommendations,omitempty"`
}

// +kcc:proto=google.cloud.recommender.v1beta1.Insight.RecommendationReference
type Insight_RecommendationReference struct {
	// Recommendation resource name, e.g.
	//  projects/[PROJECT_NUMBER]/locations/[LOCATION]/recommenders/[RECOMMENDER_ID]/recommendations/[RECOMMENDATION_ID]
	// +kcc:proto:field=google.cloud.recommender.v1beta1.Insight.RecommendationReference.recommendation
	Recommendation *string `json:"recommendation,omitempty"`
}

// +kcc:proto=google.cloud.recommender.v1beta1.InsightStateInfo
type InsightStateInfo struct {
	// Insight state.
	// +kcc:proto:field=google.cloud.recommender.v1beta1.InsightStateInfo.state
	State *string `json:"state,omitempty"`

	// A map of metadata for the state, provided by user or automations systems.
	// +kcc:proto:field=google.cloud.recommender.v1beta1.InsightStateInfo.state_metadata
	StateMetadata map[string]string `json:"stateMetadata,omitempty"`
}
