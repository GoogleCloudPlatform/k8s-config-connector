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


// +kcc:proto=google.cloud.recommender.v1beta1.CostProjection
type CostProjection struct {
	// An approximate projection on amount saved or amount incurred. Negative cost
	//  units indicate cost savings and positive cost units indicate increase.
	//  See google.type.Money documentation for positive/negative units.
	//
	//  A user's permissions may affect whether the cost is computed using list
	//  prices or custom contract prices.
	// +kcc:proto:field=google.cloud.recommender.v1beta1.CostProjection.cost
	Cost *Money `json:"cost,omitempty"`

	// Duration for which this cost applies.
	// +kcc:proto:field=google.cloud.recommender.v1beta1.CostProjection.duration
	Duration *string `json:"duration,omitempty"`
}

// +kcc:proto=google.cloud.recommender.v1beta1.Impact
type Impact struct {
	// Category that is being targeted.
	// +kcc:proto:field=google.cloud.recommender.v1beta1.Impact.category
	Category *string `json:"category,omitempty"`

	// Use with CategoryType.COST
	// +kcc:proto:field=google.cloud.recommender.v1beta1.Impact.cost_projection
	CostProjection *CostProjection `json:"costProjection,omitempty"`

	// Use with CategoryType.SECURITY
	// +kcc:proto:field=google.cloud.recommender.v1beta1.Impact.security_projection
	SecurityProjection *SecurityProjection `json:"securityProjection,omitempty"`

	// Use with CategoryType.SUSTAINABILITY
	// +kcc:proto:field=google.cloud.recommender.v1beta1.Impact.sustainability_projection
	SustainabilityProjection *SustainabilityProjection `json:"sustainabilityProjection,omitempty"`
}

// +kcc:proto=google.cloud.recommender.v1beta1.Operation
type Operation struct {
	// Type of this operation. Contains one of 'add', 'remove', 'replace', 'move',
	//  'copy', 'test' and 'custom' operations. This field is case-insensitive and
	//  always populated.
	// +kcc:proto:field=google.cloud.recommender.v1beta1.Operation.action
	Action *string `json:"action,omitempty"`

	// Type of GCP resource being modified/tested. This field is always populated.
	//  Example: cloudresourcemanager.googleapis.com/Project,
	//  compute.googleapis.com/Instance
	// +kcc:proto:field=google.cloud.recommender.v1beta1.Operation.resource_type
	ResourceType *string `json:"resourceType,omitempty"`

	// Contains the fully qualified resource name. This field is always populated.
	//  ex: //cloudresourcemanager.googleapis.com/projects/foo.
	// +kcc:proto:field=google.cloud.recommender.v1beta1.Operation.resource
	Resource *string `json:"resource,omitempty"`

	// Path to the target field being operated on. If the operation is at the
	//  resource level, then path should be "/". This field is always populated.
	// +kcc:proto:field=google.cloud.recommender.v1beta1.Operation.path
	Path *string `json:"path,omitempty"`

	// Can be set with action 'copy' to copy resource configuration across
	//  different resources of the same type. Example: A resource clone can be
	//  done via action = 'copy', path = "/", from = "/",
	//  source_resource = <source> and resource_name = <target>.
	//  This field is empty for all other values of `action`.
	// +kcc:proto:field=google.cloud.recommender.v1beta1.Operation.source_resource
	SourceResource *string `json:"sourceResource,omitempty"`

	// Can be set with action 'copy' or 'move' to indicate the source field within
	//  resource or source_resource, ignored if provided for other operation types.
	// +kcc:proto:field=google.cloud.recommender.v1beta1.Operation.source_path
	SourcePath *string `json:"sourcePath,omitempty"`

	// Value for the `path` field. Will be set for actions:'add'/'replace'.
	//  Maybe set for action: 'test'. Either this or `value_matcher` will be set
	//  for 'test' operation. An exact match must be performed.
	// +kcc:proto:field=google.cloud.recommender.v1beta1.Operation.value
	Value *Value `json:"value,omitempty"`

	// Can be set for action 'test' for advanced matching for the value of
	//  'path' field. Either this or `value` will be set for 'test' operation.
	// +kcc:proto:field=google.cloud.recommender.v1beta1.Operation.value_matcher
	ValueMatcher *ValueMatcher `json:"valueMatcher,omitempty"`

	// TODO: unsupported map type with key string and value message


	// TODO: unsupported map type with key string and value message

}

// +kcc:proto=google.cloud.recommender.v1beta1.OperationGroup
type OperationGroup struct {
	// List of operations across one or more resources that belong to this group.
	//  Loosely based on RFC6902 and should be performed in the order they appear.
	// +kcc:proto:field=google.cloud.recommender.v1beta1.OperationGroup.operations
	Operations []Operation `json:"operations,omitempty"`
}

// +kcc:proto=google.cloud.recommender.v1beta1.Recommendation
type Recommendation struct {
	// Name of recommendation.
	// +kcc:proto:field=google.cloud.recommender.v1beta1.Recommendation.name
	Name *string `json:"name,omitempty"`

	// Free-form human readable summary in English. The maximum length is 500
	//  characters.
	// +kcc:proto:field=google.cloud.recommender.v1beta1.Recommendation.description
	Description *string `json:"description,omitempty"`

	// Contains an identifier for a subtype of recommendations produced for the
	//  same recommender. Subtype is a function of content and impact, meaning a
	//  new subtype might be added when significant changes to `content` or
	//  `primary_impact.category` are introduced. See the Recommenders section
	//  to see a list of subtypes for a given Recommender.
	//
	//  Examples:
	//    For recommender = "google.iam.policy.Recommender",
	//    recommender_subtype can be one of "REMOVE_ROLE"/"REPLACE_ROLE"
	// +kcc:proto:field=google.cloud.recommender.v1beta1.Recommendation.recommender_subtype
	RecommenderSubtype *string `json:"recommenderSubtype,omitempty"`

	// Last time this recommendation was refreshed by the system that created it
	//  in the first place.
	// +kcc:proto:field=google.cloud.recommender.v1beta1.Recommendation.last_refresh_time
	LastRefreshTime *string `json:"lastRefreshTime,omitempty"`

	// The primary impact that this recommendation can have while trying to
	//  optimize for one category.
	// +kcc:proto:field=google.cloud.recommender.v1beta1.Recommendation.primary_impact
	PrimaryImpact *Impact `json:"primaryImpact,omitempty"`

	// Optional set of additional impact that this recommendation may have when
	//  trying to optimize for the primary category. These may be positive
	//  or negative.
	// +kcc:proto:field=google.cloud.recommender.v1beta1.Recommendation.additional_impact
	AdditionalImpact []Impact `json:"additionalImpact,omitempty"`

	// Recommendation's priority.
	// +kcc:proto:field=google.cloud.recommender.v1beta1.Recommendation.priority
	Priority *string `json:"priority,omitempty"`

	// Content of the recommendation describing recommended changes to resources.
	// +kcc:proto:field=google.cloud.recommender.v1beta1.Recommendation.content
	Content *RecommendationContent `json:"content,omitempty"`

	// Information for state. Contains state and metadata.
	// +kcc:proto:field=google.cloud.recommender.v1beta1.Recommendation.state_info
	StateInfo *RecommendationStateInfo `json:"stateInfo,omitempty"`

	// Fingerprint of the Recommendation. Provides optimistic locking when
	//  updating states.
	// +kcc:proto:field=google.cloud.recommender.v1beta1.Recommendation.etag
	Etag *string `json:"etag,omitempty"`

	// Insights that led to this recommendation.
	// +kcc:proto:field=google.cloud.recommender.v1beta1.Recommendation.associated_insights
	AssociatedInsights []Recommendation_InsightReference `json:"associatedInsights,omitempty"`

	// Corresponds to a mutually exclusive group ID within a recommender.
	//  A non-empty ID indicates that the recommendation belongs to a mutually
	//  exclusive group. This means that only one recommendation within the group
	//  is suggested to be applied.
	// +kcc:proto:field=google.cloud.recommender.v1beta1.Recommendation.xor_group_id
	XorGroupID *string `json:"xorGroupID,omitempty"`
}

// +kcc:proto=google.cloud.recommender.v1beta1.Recommendation.InsightReference
type Recommendation_InsightReference struct {
	// Insight resource name, e.g.
	//  projects/[PROJECT_NUMBER]/locations/[LOCATION]/insightTypes/[INSIGHT_TYPE_ID]/insights/[INSIGHT_ID]
	// +kcc:proto:field=google.cloud.recommender.v1beta1.Recommendation.InsightReference.insight
	Insight *string `json:"insight,omitempty"`
}

// +kcc:proto=google.cloud.recommender.v1beta1.RecommendationContent
type RecommendationContent struct {
	// Operations to one or more Google Cloud resources grouped in such a way
	//  that, all operations within one group are expected to be performed
	//  atomically and in an order.
	// +kcc:proto:field=google.cloud.recommender.v1beta1.RecommendationContent.operation_groups
	OperationGroups []OperationGroup `json:"operationGroups,omitempty"`

	// Condensed overview information about the recommendation.
	// +kcc:proto:field=google.cloud.recommender.v1beta1.RecommendationContent.overview
	Overview map[string]string `json:"overview,omitempty"`
}

// +kcc:proto=google.cloud.recommender.v1beta1.RecommendationStateInfo
type RecommendationStateInfo struct {
	// The state of the recommendation, Eg ACTIVE, SUCCEEDED, FAILED.
	// +kcc:proto:field=google.cloud.recommender.v1beta1.RecommendationStateInfo.state
	State *string `json:"state,omitempty"`

	// A map of metadata for the state, provided by user or automations systems.
	// +kcc:proto:field=google.cloud.recommender.v1beta1.RecommendationStateInfo.state_metadata
	StateMetadata map[string]string `json:"stateMetadata,omitempty"`
}

// +kcc:proto=google.cloud.recommender.v1beta1.SecurityProjection
type SecurityProjection struct {
	// This field can be used by the recommender to define details specific to
	//  security impact.
	// +kcc:proto:field=google.cloud.recommender.v1beta1.SecurityProjection.details
	Details map[string]string `json:"details,omitempty"`
}

// +kcc:proto=google.cloud.recommender.v1beta1.SustainabilityProjection
type SustainabilityProjection struct {
	// Carbon Footprint generated in kg of CO2 equivalent.
	//  Chose kg_c_o2e so that the name renders correctly in camelCase (kgCO2e).
	// +kcc:proto:field=google.cloud.recommender.v1beta1.SustainabilityProjection.kg_c_o2e
	KgCO2e *float64 `json:"kgCO2e,omitempty"`

	// Duration for which this sustanability applies.
	// +kcc:proto:field=google.cloud.recommender.v1beta1.SustainabilityProjection.duration
	Duration *string `json:"duration,omitempty"`
}

// +kcc:proto=google.cloud.recommender.v1beta1.ValueMatcher
type ValueMatcher struct {
	// To be used for full regex matching. The regular expression is using the
	//  Google RE2 syntax (https://github.com/google/re2/wiki/Syntax), so to be
	//  used with RE2::FullMatch
	// +kcc:proto:field=google.cloud.recommender.v1beta1.ValueMatcher.matches_pattern
	MatchesPattern *string `json:"matchesPattern,omitempty"`
}

// +kcc:proto=google.protobuf.ListValue
type ListValue struct {
	// Repeated field of dynamically typed values.
	// +kcc:proto:field=google.protobuf.ListValue.values
	Values []Value `json:"values,omitempty"`
}

// +kcc:proto=google.protobuf.Value
type Value struct {
	// Represents a null value.
	// +kcc:proto:field=google.protobuf.Value.null_value
	NullValue *string `json:"nullValue,omitempty"`

	// Represents a double value.
	// +kcc:proto:field=google.protobuf.Value.number_value
	NumberValue *float64 `json:"numberValue,omitempty"`

	// Represents a string value.
	// +kcc:proto:field=google.protobuf.Value.string_value
	StringValue *string `json:"stringValue,omitempty"`

	// Represents a boolean value.
	// +kcc:proto:field=google.protobuf.Value.bool_value
	BoolValue *bool `json:"boolValue,omitempty"`

	// Represents a structured value.
	// +kcc:proto:field=google.protobuf.Value.struct_value
	StructValue map[string]string `json:"structValue,omitempty"`

	// Represents a repeated `Value`.
	// +kcc:proto:field=google.protobuf.Value.list_value
	ListValue *ListValue `json:"listValue,omitempty"`
}

// +kcc:proto=google.type.Money
type Money struct {
	// The three-letter currency code defined in ISO 4217.
	// +kcc:proto:field=google.type.Money.currency_code
	CurrencyCode *string `json:"currencyCode,omitempty"`

	// The whole units of the amount.
	//  For example if `currencyCode` is `"USD"`, then 1 unit is one US dollar.
	// +kcc:proto:field=google.type.Money.units
	Units *int64 `json:"units,omitempty"`

	// Number of nano (10^-9) units of the amount.
	//  The value must be between -999,999,999 and +999,999,999 inclusive.
	//  If `units` is positive, `nanos` must be positive or zero.
	//  If `units` is zero, `nanos` can be positive, zero, or negative.
	//  If `units` is negative, `nanos` must be negative or zero.
	//  For example $-1.75 is represented as `units`=-1 and `nanos`=-750,000,000.
	// +kcc:proto:field=google.type.Money.nanos
	Nanos *int32 `json:"nanos,omitempty"`
}
