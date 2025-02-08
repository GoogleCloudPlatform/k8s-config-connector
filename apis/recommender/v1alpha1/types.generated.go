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


// +kcc:proto=google.cloud.recommender.v1beta1.InsightTypeConfig
type InsightTypeConfig struct {
	// Name of insight type config.
	//  Eg,
	//  projects/[PROJECT_NUMBER]/locations/[LOCATION]/insightTypes/[INSIGHT_TYPE_ID]/config
	// +kcc:proto:field=google.cloud.recommender.v1beta1.InsightTypeConfig.name
	Name *string `json:"name,omitempty"`

	// InsightTypeGenerationConfig which configures the generation of
	//  insights for this insight type.
	// +kcc:proto:field=google.cloud.recommender.v1beta1.InsightTypeConfig.insight_type_generation_config
	InsightTypeGenerationConfig *InsightTypeGenerationConfig `json:"insightTypeGenerationConfig,omitempty"`

	// Fingerprint of the InsightTypeConfig. Provides optimistic locking when
	//  updating.
	// +kcc:proto:field=google.cloud.recommender.v1beta1.InsightTypeConfig.etag
	Etag *string `json:"etag,omitempty"`

	// Last time when the config was updated.
	// +kcc:proto:field=google.cloud.recommender.v1beta1.InsightTypeConfig.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Allows clients to store small amounts of arbitrary data. Annotations must
	//  follow the Kubernetes syntax.
	//  The total size of all keys and values combined is limited to 256k.
	//  Key can have 2 segments: prefix (optional) and name (required),
	//  separated by a slash (/).
	//  Prefix must be a DNS subdomain.
	//  Name must be 63 characters or less, begin and end with alphanumerics,
	//  with dashes (-), underscores (_), dots (.), and alphanumerics between.
	// +kcc:proto:field=google.cloud.recommender.v1beta1.InsightTypeConfig.annotations
	Annotations map[string]string `json:"annotations,omitempty"`

	// A user-settable field to provide a human-readable name to be used in user
	//  interfaces.
	// +kcc:proto:field=google.cloud.recommender.v1beta1.InsightTypeConfig.display_name
	DisplayName *string `json:"displayName,omitempty"`
}

// +kcc:proto=google.cloud.recommender.v1beta1.InsightTypeGenerationConfig
type InsightTypeGenerationConfig struct {
	// Parameters for this InsightTypeGenerationConfig. These configs can be used
	//  by or are applied to all subtypes.
	// +kcc:proto:field=google.cloud.recommender.v1beta1.InsightTypeGenerationConfig.params
	Params map[string]string `json:"params,omitempty"`
}

// +kcc:proto=google.cloud.recommender.v1beta1.InsightTypeConfig
type InsightTypeConfigObservedState struct {
	// Output only. Immutable. The revision ID of the config.
	//  A new revision is committed whenever the config is changed in any way.
	//  The format is an 8-character hexadecimal string.
	// +kcc:proto:field=google.cloud.recommender.v1beta1.InsightTypeConfig.revision_id
	RevisionID *string `json:"revisionID,omitempty"`
}
