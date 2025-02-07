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


// +kcc:proto=google.cloud.asset.v1.Feed
type Feed struct {
	// Required. The format will be
	//  projects/{project_number}/feeds/{client-assigned_feed_identifier} or
	//  folders/{folder_number}/feeds/{client-assigned_feed_identifier} or
	//  organizations/{organization_number}/feeds/{client-assigned_feed_identifier}
	//
	//  The client-assigned feed identifier must be unique within the parent
	//  project/folder/organization.
	// +kcc:proto:field=google.cloud.asset.v1.Feed.name
	Name *string `json:"name,omitempty"`

	// A list of the full names of the assets to receive updates. You must specify
	//  either or both of asset_names and asset_types. Only asset updates matching
	//  specified asset_names or asset_types are exported to the feed.
	//  Example:
	//  `//compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1`.
	//  For a list of the full names for supported asset types, see [Resource
	//  name format](/asset-inventory/docs/resource-name-format).
	// +kcc:proto:field=google.cloud.asset.v1.Feed.asset_names
	AssetNames []string `json:"assetNames,omitempty"`

	// A list of types of the assets to receive updates. You must specify either
	//  or both of asset_names and asset_types. Only asset updates matching
	//  specified asset_names or asset_types are exported to the feed.
	//  Example: `"compute.googleapis.com/Disk"`
	//
	//  For a list of all supported asset types, see
	//  [Supported asset types](/asset-inventory/docs/supported-asset-types).
	// +kcc:proto:field=google.cloud.asset.v1.Feed.asset_types
	AssetTypes []string `json:"assetTypes,omitempty"`

	// Asset content type. If not specified, no content but the asset name and
	//  type will be returned.
	// +kcc:proto:field=google.cloud.asset.v1.Feed.content_type
	ContentType *string `json:"contentType,omitempty"`

	// Required. Feed output configuration defining where the asset updates are
	//  published to.
	// +kcc:proto:field=google.cloud.asset.v1.Feed.feed_output_config
	FeedOutputConfig *FeedOutputConfig `json:"feedOutputConfig,omitempty"`

	// A condition which determines whether an asset update should be published.
	//  If specified, an asset will be returned only when the expression evaluates
	//  to true.
	//  When set, `expression` field in the `Expr` must be a valid [CEL expression]
	//  (https://github.com/google/cel-spec) on a TemporalAsset with name
	//  `temporal_asset`. Example: a Feed with expression ("temporal_asset.deleted
	//  == true") will only publish Asset deletions. Other fields of `Expr` are
	//  optional.
	//
	//  See our [user
	//  guide](https://cloud.google.com/asset-inventory/docs/monitoring-asset-changes-with-condition)
	//  for detailed instructions.
	// +kcc:proto:field=google.cloud.asset.v1.Feed.condition
	Condition *Expr `json:"condition,omitempty"`

	// A list of relationship types to output, for example:
	//  `INSTANCE_TO_INSTANCEGROUP`. This field should only be specified if
	//  content_type=RELATIONSHIP.
	//  * If specified:
	//  it outputs specified relationship updates on the [asset_names] or the
	//  [asset_types]. It returns an error if any of the [relationship_types]
	//  doesn't belong to the supported relationship types of the [asset_names] or
	//  [asset_types], or any of the [asset_names] or the [asset_types] doesn't
	//  belong to the source types of the [relationship_types].
	//  * Otherwise:
	//  it outputs the supported relationships of the types of [asset_names] and
	//  [asset_types] or returns an error if any of the [asset_names] or the
	//  [asset_types] has no replationship support.
	//  See [Introduction to Cloud Asset
	//  Inventory](https://cloud.google.com/asset-inventory/docs/overview)
	//  for all supported asset types and relationship types.
	// +kcc:proto:field=google.cloud.asset.v1.Feed.relationship_types
	RelationshipTypes []string `json:"relationshipTypes,omitempty"`
}

// +kcc:proto=google.cloud.asset.v1.FeedOutputConfig
type FeedOutputConfig struct {
	// Destination on Pub/Sub.
	// +kcc:proto:field=google.cloud.asset.v1.FeedOutputConfig.pubsub_destination
	PubsubDestination *PubsubDestination `json:"pubsubDestination,omitempty"`
}

// +kcc:proto=google.cloud.asset.v1.PubsubDestination
type PubsubDestination struct {
	// The name of the Pub/Sub topic to publish to.
	//  Example: `projects/PROJECT_ID/topics/TOPIC_ID`.
	// +kcc:proto:field=google.cloud.asset.v1.PubsubDestination.topic
	Topic *string `json:"topic,omitempty"`
}

// +kcc:proto=google.type.Expr
type Expr struct {
	// Textual representation of an expression in Common Expression Language
	//  syntax.
	// +kcc:proto:field=google.type.Expr.expression
	Expression *string `json:"expression,omitempty"`

	// Optional. Title for the expression, i.e. a short string describing
	//  its purpose. This can be used e.g. in UIs which allow to enter the
	//  expression.
	// +kcc:proto:field=google.type.Expr.title
	Title *string `json:"title,omitempty"`

	// Optional. Description of the expression. This is a longer text which
	//  describes the expression, e.g. when hovered over it in a UI.
	// +kcc:proto:field=google.type.Expr.description
	Description *string `json:"description,omitempty"`

	// Optional. String indicating the location of the expression for error
	//  reporting, e.g. a file name and a position in the file.
	// +kcc:proto:field=google.type.Expr.location
	Location *string `json:"location,omitempty"`
}
