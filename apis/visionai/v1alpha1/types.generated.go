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


// +kcc:proto=google.cloud.visionai.v1.DataSchema
type DataSchema struct {
	// Resource name of the data schema in the form of:
	//  `projects/{project_number}/locations/{location}/corpora/{corpus}/dataSchemas/{data_schema}`
	//  where {data_schema} part should be the same as the `key` field below.
	// +kcc:proto:field=google.cloud.visionai.v1.DataSchema.name
	Name *string `json:"name,omitempty"`

	// Required. The key of this data schema. This key should be matching the key
	//  of user specified annotation and unique inside corpus. This value can be up
	//  to 63 characters, and valid characters are /[a-z][0-9]-/. The first
	//  character must be a letter, the last could be a letter or a number.
	// +kcc:proto:field=google.cloud.visionai.v1.DataSchema.key
	Key *string `json:"key,omitempty"`

	// The schema details mapping to the key.
	// +kcc:proto:field=google.cloud.visionai.v1.DataSchema.schema_details
	SchemaDetails *DataSchemaDetails `json:"schemaDetails,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.DataSchemaDetails
type DataSchemaDetails struct {
	// Type of the annotation.
	// +kcc:proto:field=google.cloud.visionai.v1.DataSchemaDetails.type
	Type *string `json:"type,omitempty"`

	// Config for protobuf any type.
	// +kcc:proto:field=google.cloud.visionai.v1.DataSchemaDetails.proto_any_config
	ProtoAnyConfig *DataSchemaDetails_ProtoAnyConfig `json:"protoAnyConfig,omitempty"`

	// Config for List data type.
	// +kcc:proto:field=google.cloud.visionai.v1.DataSchemaDetails.list_config
	ListConfig *DataSchemaDetails_ListConfig `json:"listConfig,omitempty"`

	// Config for CustomizedStruct data type.
	// +kcc:proto:field=google.cloud.visionai.v1.DataSchemaDetails.customized_struct_config
	CustomizedStructConfig *DataSchemaDetails_CustomizedStructConfig `json:"customizedStructConfig,omitempty"`

	// The granularity associated with this DataSchema.
	// +kcc:proto:field=google.cloud.visionai.v1.DataSchemaDetails.granularity
	Granularity *string `json:"granularity,omitempty"`

	// The search strategy to be applied on the `key` above.
	// +kcc:proto:field=google.cloud.visionai.v1.DataSchemaDetails.search_strategy
	SearchStrategy *DataSchemaDetails_SearchStrategy `json:"searchStrategy,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.DataSchemaDetails.CustomizedStructConfig
type DataSchemaDetails_CustomizedStructConfig struct {

	// TODO: unsupported map type with key string and value message

}

// +kcc:proto=google.cloud.visionai.v1.DataSchemaDetails.ListConfig
type DataSchemaDetails_ListConfig struct {
	// The value's data schema in the list.
	// +kcc:proto:field=google.cloud.visionai.v1.DataSchemaDetails.ListConfig.value_schema
	ValueSchema *DataSchemaDetails `json:"valueSchema,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.DataSchemaDetails.ProtoAnyConfig
type DataSchemaDetails_ProtoAnyConfig struct {
	// The type URI of the proto message.
	// +kcc:proto:field=google.cloud.visionai.v1.DataSchemaDetails.ProtoAnyConfig.type_uri
	TypeURI *string `json:"typeURI,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.DataSchemaDetails.SearchStrategy
type DataSchemaDetails_SearchStrategy struct {
	// The type of search strategy to be applied on the `key` above.
	//  The allowed `search_strategy_type` is different for different data types,
	//  which is documented in the DataSchemaDetails.DataType. Specifying
	//  unsupported `search_strategy_type` for data types will result in
	//  INVALID_ARGUMENT error.
	// +kcc:proto:field=google.cloud.visionai.v1.DataSchemaDetails.SearchStrategy.search_strategy_type
	SearchStrategyType *string `json:"searchStrategyType,omitempty"`

	// Optional. Configs the path to the confidence score, and the threshold.
	//  Only if the score is greater than the threshold, current field will be
	//  built into the index. Only applies to leaf nodes using EXACT_SEARCH or
	//  SMART_SEARCH.
	// +kcc:proto:field=google.cloud.visionai.v1.DataSchemaDetails.SearchStrategy.confidence_score_index_config
	ConfidenceScoreIndexConfig *DataSchemaDetails_SearchStrategy_ConfidenceScoreIndexConfig `json:"confidenceScoreIndexConfig,omitempty"`
}

// +kcc:proto=google.cloud.visionai.v1.DataSchemaDetails.SearchStrategy.ConfidenceScoreIndexConfig
type DataSchemaDetails_SearchStrategy_ConfidenceScoreIndexConfig struct {
	// Required. The path to the confidence score field. It is a string that
	//  concatenates all the data schema keys along the path. See the example
	//  above. If the data schema contains LIST, use '_ENTRIES' to concatenate.
	//  Example data schema contains a list:
	//  "key": "list-name-score",
	//  "schemaDetails": {
	//    "type": "LIST",
	//    "granularity": "GRANULARITY_PARTITION_LEVEL",
	//    "listConfig": {
	//      "valueSchema": {
	//        "type": "CUSTOMIZED_STRUCT",
	//        "granularity": "GRANULARITY_PARTITION_LEVEL",
	//        "customizedStructConfig": {
	//          "fieldSchemas": {
	//            "name": {
	//              "type": "STRING",
	//              "granularity": "GRANULARITY_PARTITION_LEVEL",
	//              "searchStrategy": {
	//                "searchStrategyType": "SMART_SEARCH"
	//                "confidence_score_index_config": {
	//                  "field_path": "list-name-score._ENTRIES.score",
	//                  "threshold": "0.9",
	//                }
	//              }
	//            },
	//            "score": {
	//              "type": "FLOAT",
	//              "granularity": "GRANULARITY_PARTITION_LEVEL",
	//            }
	//          }
	//        }
	//      }
	//    }
	//  }
	// +kcc:proto:field=google.cloud.visionai.v1.DataSchemaDetails.SearchStrategy.ConfidenceScoreIndexConfig.field_path
	FieldPath *string `json:"fieldPath,omitempty"`

	// Required. The threshold.
	// +kcc:proto:field=google.cloud.visionai.v1.DataSchemaDetails.SearchStrategy.ConfidenceScoreIndexConfig.threshold
	Threshold *float32 `json:"threshold,omitempty"`
}
