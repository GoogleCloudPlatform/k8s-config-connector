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


// +kcc:proto=google.cloud.aiplatform.v1.FeatureView
type FeatureView struct {
	// Optional. Configures how data is supposed to be extracted from a BigQuery
	//  source to be loaded onto the FeatureOnlineStore.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureView.big_query_source
	BigQuerySource *FeatureView_BigQuerySource `json:"bigQuerySource,omitempty"`

	// Optional. Configures the features from a Feature Registry source that
	//  need to be loaded onto the FeatureOnlineStore.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureView.feature_registry_source
	FeatureRegistrySource *FeatureView_FeatureRegistrySource `json:"featureRegistrySource,omitempty"`

	// Optional. The Vertex RAG Source that the FeatureView is linked to.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureView.vertex_rag_source
	VertexRagSource *FeatureView_VertexRagSource `json:"vertexRagSource,omitempty"`

	// Identifier. Name of the FeatureView. Format:
	//  `projects/{project}/locations/{location}/featureOnlineStores/{feature_online_store}/featureViews/{feature_view}`
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureView.name
	Name *string `json:"name,omitempty"`

	// Optional. Used to perform consistent read-modify-write updates. If not set,
	//  a blind "overwrite" update happens.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureView.etag
	Etag *string `json:"etag,omitempty"`

	// Optional. The labels with user-defined metadata to organize your
	//  FeatureViews.
	//
	//  Label keys and values can be no longer than 64 characters
	//  (Unicode codepoints), can only contain lowercase letters, numeric
	//  characters, underscores and dashes. International characters are allowed.
	//
	//  See https://goo.gl/xmQnxf for more information on and examples of labels.
	//  No more than 64 user labels can be associated with one
	//  FeatureOnlineStore(System labels are excluded)." System reserved label keys
	//  are prefixed with "aiplatform.googleapis.com/" and are immutable.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureView.labels
	Labels map[string]string `json:"labels,omitempty"`

	// Configures when data is to be synced/updated for this FeatureView. At the
	//  end of the sync the latest featureValues for each entityId of this
	//  FeatureView are made ready for online serving.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureView.sync_config
	SyncConfig *FeatureView_SyncConfig `json:"syncConfig,omitempty"`

	// Optional. Configuration for index preparation for vector search. It
	//  contains the required configurations to create an index from source data,
	//  so that approximate nearest neighbor (a.k.a ANN) algorithms search can be
	//  performed during online serving.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureView.index_config
	IndexConfig *FeatureView_IndexConfig `json:"indexConfig,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.FeatureView.BigQuerySource
type FeatureView_BigQuerySource struct {
	// Required. The BigQuery view URI that will be materialized on each sync
	//  trigger based on FeatureView.SyncConfig.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureView.BigQuerySource.uri
	URI *string `json:"uri,omitempty"`

	// Required. Columns to construct entity_id / row keys.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureView.BigQuerySource.entity_id_columns
	EntityIDColumns []string `json:"entityIDColumns,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.FeatureView.FeatureRegistrySource
type FeatureView_FeatureRegistrySource struct {
	// Required. List of features that need to be synced to Online Store.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureView.FeatureRegistrySource.feature_groups
	FeatureGroups []FeatureView_FeatureRegistrySource_FeatureGroup `json:"featureGroups,omitempty"`

	// Optional. The project number of the parent project of the Feature Groups.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureView.FeatureRegistrySource.project_number
	ProjectNumber *int64 `json:"projectNumber,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.FeatureView.FeatureRegistrySource.FeatureGroup
type FeatureView_FeatureRegistrySource_FeatureGroup struct {
	// Required. Identifier of the feature group.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureView.FeatureRegistrySource.FeatureGroup.feature_group_id
	FeatureGroupID *string `json:"featureGroupID,omitempty"`

	// Required. Identifiers of features under the feature group.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureView.FeatureRegistrySource.FeatureGroup.feature_ids
	FeatureIds []string `json:"featureIds,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.FeatureView.IndexConfig
type FeatureView_IndexConfig struct {
	// Optional. Configuration options for the tree-AH algorithm (Shallow tree
	//  + Asymmetric Hashing). Please refer to this paper for more details:
	//  https://arxiv.org/abs/1908.10396
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureView.IndexConfig.tree_ah_config
	TreeAhConfig *FeatureView_IndexConfig_TreeAHConfig `json:"treeAhConfig,omitempty"`

	// Optional. Configuration options for using brute force search, which
	//  simply implements the standard linear search in the database for each
	//  query. It is primarily meant for benchmarking and to generate the
	//  ground truth for approximate search.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureView.IndexConfig.brute_force_config
	BruteForceConfig *FeatureView_IndexConfig_BruteForceConfig `json:"bruteForceConfig,omitempty"`

	// Optional. Column of embedding. This column contains the source data to
	//  create index for vector search. embedding_column must be set when using
	//  vector search.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureView.IndexConfig.embedding_column
	EmbeddingColumn *string `json:"embeddingColumn,omitempty"`

	// Optional. Columns of features that're used to filter vector search
	//  results.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureView.IndexConfig.filter_columns
	FilterColumns []string `json:"filterColumns,omitempty"`

	// Optional. Column of crowding. This column contains crowding attribute
	//  which is a constraint on a neighbor list produced by
	//  [FeatureOnlineStoreService.SearchNearestEntities][google.cloud.aiplatform.v1.FeatureOnlineStoreService.SearchNearestEntities]
	//  to diversify search results. If
	//  [NearestNeighborQuery.per_crowding_attribute_neighbor_count][google.cloud.aiplatform.v1.NearestNeighborQuery.per_crowding_attribute_neighbor_count]
	//  is set to K in
	//  [SearchNearestEntitiesRequest][google.cloud.aiplatform.v1.SearchNearestEntitiesRequest],
	//  it's guaranteed that no more than K entities of the same crowding
	//  attribute are returned in the response.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureView.IndexConfig.crowding_column
	CrowdingColumn *string `json:"crowdingColumn,omitempty"`

	// Optional. The number of dimensions of the input embedding.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureView.IndexConfig.embedding_dimension
	EmbeddingDimension *int32 `json:"embeddingDimension,omitempty"`

	// Optional. The distance measure used in nearest neighbor search.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureView.IndexConfig.distance_measure_type
	DistanceMeasureType *string `json:"distanceMeasureType,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.FeatureView.IndexConfig.BruteForceConfig
type FeatureView_IndexConfig_BruteForceConfig struct {
}

// +kcc:proto=google.cloud.aiplatform.v1.FeatureView.IndexConfig.TreeAHConfig
type FeatureView_IndexConfig_TreeAHConfig struct {
	// Optional. Number of embeddings on each leaf node. The default value is
	//  1000 if not set.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureView.IndexConfig.TreeAHConfig.leaf_node_embedding_count
	LeafNodeEmbeddingCount *int64 `json:"leafNodeEmbeddingCount,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.FeatureView.SyncConfig
type FeatureView_SyncConfig struct {
	// Cron schedule (https://en.wikipedia.org/wiki/Cron) to launch scheduled
	//  runs. To explicitly set a timezone to the cron tab, apply a prefix in
	//  the cron tab: "CRON_TZ=${IANA_TIME_ZONE}" or "TZ=${IANA_TIME_ZONE}".
	//  The ${IANA_TIME_ZONE} may only be a valid string from IANA time zone
	//  database. For example, "CRON_TZ=America/New_York 1 * * * *", or
	//  "TZ=America/New_York 1 * * * *".
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureView.SyncConfig.cron
	Cron *string `json:"cron,omitempty"`

	// Optional. If true, syncs the FeatureView in a continuous manner to Online
	//  Store.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureView.SyncConfig.continuous
	Continuous *bool `json:"continuous,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.FeatureView.VertexRagSource
type FeatureView_VertexRagSource struct {
	// Required. The BigQuery view/table URI that will be materialized on each
	//  manual sync trigger. The table/view is expected to have the following
	//  columns and types at least:
	//   - `corpus_id` (STRING, NULLABLE/REQUIRED)
	//   - `file_id` (STRING, NULLABLE/REQUIRED)
	//   - `chunk_id` (STRING, NULLABLE/REQUIRED)
	//   - `chunk_data_type` (STRING, NULLABLE/REQUIRED)
	//   - `chunk_data` (STRING, NULLABLE/REQUIRED)
	//   - `embeddings` (FLOAT, REPEATED)
	//   - `file_original_uri` (STRING, NULLABLE/REQUIRED)
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureView.VertexRagSource.uri
	URI *string `json:"uri,omitempty"`

	// Optional. The RAG corpus id corresponding to this FeatureView.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureView.VertexRagSource.rag_corpus_id
	RagCorpusID *int64 `json:"ragCorpusID,omitempty"`
}

// +kcc:proto=google.cloud.aiplatform.v1.FeatureView
type FeatureViewObservedState struct {
	// Output only. Timestamp when this FeatureView was created.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureView.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. Timestamp when this FeatureView was last updated.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureView.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureView.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.aiplatform.v1.FeatureView.satisfies_pzi
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`
}
