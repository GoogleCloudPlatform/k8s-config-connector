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


// +kcc:proto=google.cloud.asset.v1p2beta1.Feed
type Feed struct {
	// Required. The format will be
	//  projects/{project_number}/feeds/{client-assigned_feed_identifier} or
	//  folders/{folder_number}/feeds/{client-assigned_feed_identifier} or
	//  organizations/{organization_number}/feeds/{client-assigned_feed_identifier}
	//
	//  The client-assigned feed identifier must be unique within the parent
	//  project/folder/organization.
	// +kcc:proto:field=google.cloud.asset.v1p2beta1.Feed.name
	Name *string `json:"name,omitempty"`

	// A list of the full names of the assets to receive updates. You must specify
	//  either or both of asset_names and asset_types. Only asset updates matching
	//  specified asset_names or asset_types are exported to the feed. For
	//  example:
	//  `//compute.googleapis.com/projects/my_project_123/zones/zone1/instances/instance1`.
	//  See [Resource
	//  Names](https://cloud.google.com/apis/design/resource_names#full_resource_name)
	//  for more info.
	// +kcc:proto:field=google.cloud.asset.v1p2beta1.Feed.asset_names
	AssetNames []string `json:"assetNames,omitempty"`

	// A list of types of the assets to receive updates. You must specify either
	//  or both of asset_names and asset_types. Only asset updates matching
	//  specified asset_names or asset_types are exported to the feed.
	//  For example:
	//  "compute.googleapis.com/Disk" See [Introduction to Cloud Asset
	//  Inventory](https://cloud.google.com/resource-manager/docs/cloud-asset-inventory/overview)
	//  for all supported asset types.
	// +kcc:proto:field=google.cloud.asset.v1p2beta1.Feed.asset_types
	AssetTypes []string `json:"assetTypes,omitempty"`

	// Asset content type. If not specified, no content but the asset name and
	//  type will be returned.
	// +kcc:proto:field=google.cloud.asset.v1p2beta1.Feed.content_type
	ContentType *string `json:"contentType,omitempty"`

	// Required. Feed output configuration defining where the asset updates are
	//  published to.
	// +kcc:proto:field=google.cloud.asset.v1p2beta1.Feed.feed_output_config
	FeedOutputConfig *FeedOutputConfig `json:"feedOutputConfig,omitempty"`
}

// +kcc:proto=google.cloud.asset.v1p2beta1.FeedOutputConfig
type FeedOutputConfig struct {
	// Destination on Pub/Sub.
	// +kcc:proto:field=google.cloud.asset.v1p2beta1.FeedOutputConfig.pubsub_destination
	PubsubDestination *PubsubDestination `json:"pubsubDestination,omitempty"`
}

// +kcc:proto=google.cloud.asset.v1p2beta1.PubsubDestination
type PubsubDestination struct {
	// The name of the Pub/Sub topic to publish to.
	//  For example: `projects/PROJECT_ID/topics/TOPIC_ID`.
	// +kcc:proto:field=google.cloud.asset.v1p2beta1.PubsubDestination.topic
	Topic *string `json:"topic,omitempty"`
}
