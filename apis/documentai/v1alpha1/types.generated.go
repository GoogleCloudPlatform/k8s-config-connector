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


// +kcc:proto=google.cloud.documentai.v1beta3.DocumentSchema
type DocumentSchema struct {
	// Display name to show to users.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.DocumentSchema.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Description of the schema.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.DocumentSchema.description
	Description *string `json:"description,omitempty"`

	// Entity types of the schema.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.DocumentSchema.entity_types
	EntityTypes []DocumentSchema_EntityType `json:"entityTypes,omitempty"`

	// Metadata of the schema.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.DocumentSchema.metadata
	Metadata *DocumentSchema_Metadata `json:"metadata,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1beta3.DocumentSchema.EntityType
type DocumentSchema_EntityType struct {
	// If specified, lists all the possible values for this entity.  This
	//  should not be more than a handful of values.  If the number of values
	//  is >10 or could change frequently use the `EntityType.value_ontology`
	//  field and specify a list of all possible values in a value ontology
	//  file.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.DocumentSchema.EntityType.enum_values
	EnumValues *DocumentSchema_EntityType_EnumValues `json:"enumValues,omitempty"`

	// User defined name for the type.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.DocumentSchema.EntityType.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Name of the type. It must be unique within the schema file and
	//  cannot be a "Common Type".  The following naming conventions are used:
	//
	//  - Use `snake_casing`.
	//  - Name matching is case-sensitive.
	//  - Maximum 64 characters.
	//  - Must start with a letter.
	//  - Allowed characters: ASCII letters `[a-z0-9_-]`.  (For backward
	//    compatibility internal infrastructure and tooling can handle any ascii
	//    character.)
	//  - The `/` is sometimes used to denote a property of a type.  For example
	//    `line_item/amount`.  This convention is deprecated, but will still be
	//    honored for backward compatibility.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.DocumentSchema.EntityType.name
	Name *string `json:"name,omitempty"`

	// The description of the entity type. Could be used to provide more
	//  information about the entity type for model calls.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.DocumentSchema.EntityType.description
	Description *string `json:"description,omitempty"`

	// The entity type that this type is derived from.  For now, one and only
	//  one should be set.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.DocumentSchema.EntityType.base_types
	BaseTypes []string `json:"baseTypes,omitempty"`

	// Description the nested structure, or composition of an entity.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.DocumentSchema.EntityType.properties
	Properties []DocumentSchema_EntityType_Property `json:"properties,omitempty"`

	// Metadata for the entity type.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.DocumentSchema.EntityType.entity_type_metadata
	EntityTypeMetadata *EntityTypeMetadata `json:"entityTypeMetadata,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1beta3.DocumentSchema.EntityType.EnumValues
type DocumentSchema_EntityType_EnumValues struct {
	// The individual values that this enum values type can include.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.DocumentSchema.EntityType.EnumValues.values
	Values []string `json:"values,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1beta3.DocumentSchema.EntityType.Property
type DocumentSchema_EntityType_Property struct {
	// The name of the property.  Follows the same guidelines as the
	//  EntityType name.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.DocumentSchema.EntityType.Property.name
	Name *string `json:"name,omitempty"`

	// The description of the property. Could be used to provide more
	//  information about the property for model calls.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.DocumentSchema.EntityType.Property.description
	Description *string `json:"description,omitempty"`

	// User defined name for the property.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.DocumentSchema.EntityType.Property.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// A reference to the value type of the property.  This type is subject
	//  to the same conventions as the `Entity.base_types` field.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.DocumentSchema.EntityType.Property.value_type
	ValueType *string `json:"valueType,omitempty"`

	// Occurrence type limits the number of instances an entity type appears
	//  in the document.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.DocumentSchema.EntityType.Property.occurrence_type
	OccurrenceType *string `json:"occurrenceType,omitempty"`

	// Any additional metadata about the property can be added here.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.DocumentSchema.EntityType.Property.property_metadata
	PropertyMetadata *PropertyMetadata `json:"propertyMetadata,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1beta3.DocumentSchema.Metadata
type DocumentSchema_Metadata struct {
	// If true, a `document` entity type can be applied to subdocument
	//  (splitting). Otherwise, it can only be applied to the entire document
	//  (classification).
	// +kcc:proto:field=google.cloud.documentai.v1beta3.DocumentSchema.Metadata.document_splitter
	DocumentSplitter *bool `json:"documentSplitter,omitempty"`

	// If true, on a given page, there can be multiple `document` annotations
	//  covering it.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.DocumentSchema.Metadata.document_allow_multiple_labels
	DocumentAllowMultipleLabels *bool `json:"documentAllowMultipleLabels,omitempty"`

	// If set, all the nested entities must be prefixed with the parents.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.DocumentSchema.Metadata.prefixed_naming_on_properties
	PrefixedNamingOnProperties *bool `json:"prefixedNamingOnProperties,omitempty"`

	// If set, we will skip the naming format validation in the schema. So the
	//  string values in `DocumentSchema.EntityType.name` and
	//  `DocumentSchema.EntityType.Property.name` will not be checked.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.DocumentSchema.Metadata.skip_naming_validation
	SkipNamingValidation *bool `json:"skipNamingValidation,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1beta3.EntityTypeMetadata
type EntityTypeMetadata struct {
	// Whether the entity type should be considered inactive.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.EntityTypeMetadata.inactive
	Inactive *bool `json:"inactive,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1beta3.Evaluation.Metrics
type Evaluation_Metrics struct {
	// The calculated precision.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.Evaluation.Metrics.precision
	Precision *float32 `json:"precision,omitempty"`

	// The calculated recall.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.Evaluation.Metrics.recall
	Recall *float32 `json:"recall,omitempty"`

	// The calculated f1 score.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.Evaluation.Metrics.f1_score
	F1Score *float32 `json:"f1Score,omitempty"`

	// The amount of occurrences in predicted documents.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.Evaluation.Metrics.predicted_occurrences_count
	PredictedOccurrencesCount *int32 `json:"predictedOccurrencesCount,omitempty"`

	// The amount of occurrences in ground truth documents.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.Evaluation.Metrics.ground_truth_occurrences_count
	GroundTruthOccurrencesCount *int32 `json:"groundTruthOccurrencesCount,omitempty"`

	// The amount of documents with a predicted occurrence.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.Evaluation.Metrics.predicted_document_count
	PredictedDocumentCount *int32 `json:"predictedDocumentCount,omitempty"`

	// The amount of documents with a ground truth occurrence.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.Evaluation.Metrics.ground_truth_document_count
	GroundTruthDocumentCount *int32 `json:"groundTruthDocumentCount,omitempty"`

	// The amount of true positives.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.Evaluation.Metrics.true_positives_count
	TruePositivesCount *int32 `json:"truePositivesCount,omitempty"`

	// The amount of false positives.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.Evaluation.Metrics.false_positives_count
	FalsePositivesCount *int32 `json:"falsePositivesCount,omitempty"`

	// The amount of false negatives.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.Evaluation.Metrics.false_negatives_count
	FalseNegativesCount *int32 `json:"falseNegativesCount,omitempty"`

	// The amount of documents that had an occurrence of this label.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.Evaluation.Metrics.total_documents_count
	TotalDocumentsCount *int32 `json:"totalDocumentsCount,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1beta3.EvaluationReference
type EvaluationReference struct {
	// The resource name of the Long Running Operation for the evaluation.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.EvaluationReference.operation
	Operation *string `json:"operation,omitempty"`

	// The resource name of the evaluation.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.EvaluationReference.evaluation
	Evaluation *string `json:"evaluation,omitempty"`

	// An aggregate of the statistics for the evaluation with fuzzy matching on.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.EvaluationReference.aggregate_metrics
	AggregateMetrics *Evaluation_Metrics `json:"aggregateMetrics,omitempty"`

	// An aggregate of the statistics for the evaluation with fuzzy matching off.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.EvaluationReference.aggregate_metrics_exact
	AggregateMetricsExact *Evaluation_Metrics `json:"aggregateMetricsExact,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1beta3.FieldExtractionMetadata
type FieldExtractionMetadata struct {
	// Summary options config.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.FieldExtractionMetadata.summary_options
	SummaryOptions *SummaryOptions `json:"summaryOptions,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1beta3.ProcessorVersion
type ProcessorVersion struct {
	// Identifier. The resource name of the processor version.
	//  Format:
	//  `projects/{project}/locations/{location}/processors/{processor}/processorVersions/{processor_version}`
	// +kcc:proto:field=google.cloud.documentai.v1beta3.ProcessorVersion.name
	Name *string `json:"name,omitempty"`

	// The display name of the processor version.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.ProcessorVersion.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// The schema of the processor version. Describes the output.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.ProcessorVersion.document_schema
	DocumentSchema *DocumentSchema `json:"documentSchema,omitempty"`

	// The time the processor version was created.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.ProcessorVersion.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// The most recently invoked evaluation for the processor version.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.ProcessorVersion.latest_evaluation
	LatestEvaluation *EvaluationReference `json:"latestEvaluation,omitempty"`

	// The KMS key name used for encryption.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.ProcessorVersion.kms_key_name
	KMSKeyName *string `json:"kmsKeyName,omitempty"`

	// The KMS key version with which data is encrypted.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.ProcessorVersion.kms_key_version_name
	KMSKeyVersionName *string `json:"kmsKeyVersionName,omitempty"`

	// If set, information about the eventual deprecation of this version.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.ProcessorVersion.deprecation_info
	DeprecationInfo *ProcessorVersion_DeprecationInfo `json:"deprecationInfo,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1beta3.ProcessorVersion.DeprecationInfo
type ProcessorVersion_DeprecationInfo struct {
	// The time at which this processor version will be deprecated.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.ProcessorVersion.DeprecationInfo.deprecation_time
	DeprecationTime *string `json:"deprecationTime,omitempty"`

	// If set, the processor version that will be used as a replacement.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.ProcessorVersion.DeprecationInfo.replacement_processor_version
	ReplacementProcessorVersion *string `json:"replacementProcessorVersion,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1beta3.ProcessorVersion.GenAiModelInfo
type ProcessorVersion_GenAiModelInfo struct {
	// Information for a pretrained Google-managed foundation model.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.ProcessorVersion.GenAiModelInfo.foundation_gen_ai_model_info
	FoundationGenAiModelInfo *ProcessorVersion_GenAiModelInfo_FoundationGenAiModelInfo `json:"foundationGenAiModelInfo,omitempty"`

	// Information for a custom Generative AI model created by the user.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.ProcessorVersion.GenAiModelInfo.custom_gen_ai_model_info
	CustomGenAiModelInfo *ProcessorVersion_GenAiModelInfo_CustomGenAiModelInfo `json:"customGenAiModelInfo,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1beta3.ProcessorVersion.GenAiModelInfo.CustomGenAiModelInfo
type ProcessorVersion_GenAiModelInfo_CustomGenAiModelInfo struct {
	// The type of custom model created by the user.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.ProcessorVersion.GenAiModelInfo.CustomGenAiModelInfo.custom_model_type
	CustomModelType *string `json:"customModelType,omitempty"`

	// The base processor version ID for the custom model.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.ProcessorVersion.GenAiModelInfo.CustomGenAiModelInfo.base_processor_version_id
	BaseProcessorVersionID *string `json:"baseProcessorVersionID,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1beta3.ProcessorVersion.GenAiModelInfo.FoundationGenAiModelInfo
type ProcessorVersion_GenAiModelInfo_FoundationGenAiModelInfo struct {
	// Whether finetuning is allowed for this base processor version.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.ProcessorVersion.GenAiModelInfo.FoundationGenAiModelInfo.finetuning_allowed
	FinetuningAllowed *bool `json:"finetuningAllowed,omitempty"`

	// The minimum number of labeled documents in the training dataset
	//  required for finetuning.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.ProcessorVersion.GenAiModelInfo.FoundationGenAiModelInfo.min_train_labeled_documents
	MinTrainLabeledDocuments *int32 `json:"minTrainLabeledDocuments,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1beta3.PropertyMetadata
type PropertyMetadata struct {
	// Whether the property should be considered as "inactive".
	// +kcc:proto:field=google.cloud.documentai.v1beta3.PropertyMetadata.inactive
	Inactive *bool `json:"inactive,omitempty"`

	// Field extraction metadata on the property.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.PropertyMetadata.field_extraction_metadata
	FieldExtractionMetadata *FieldExtractionMetadata `json:"fieldExtractionMetadata,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1beta3.SummaryOptions
type SummaryOptions struct {
	// How long the summary should be.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.SummaryOptions.length
	Length *string `json:"length,omitempty"`

	// The format the summary should be in.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.SummaryOptions.format
	Format *string `json:"format,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1beta3.ProcessorVersion
type ProcessorVersionObservedState struct {
	// Output only. The state of the processor version.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.ProcessorVersion.state
	State *string `json:"state,omitempty"`

	// Output only. Denotes that this `ProcessorVersion` is managed by Google.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.ProcessorVersion.google_managed
	GoogleManaged *bool `json:"googleManaged,omitempty"`

	// Output only. The model type of this processor version.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.ProcessorVersion.model_type
	ModelType *string `json:"modelType,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.ProcessorVersion.satisfies_pzs
	SatisfiesPzs *bool `json:"satisfiesPzs,omitempty"`

	// Output only. Reserved for future use.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.ProcessorVersion.satisfies_pzi
	SatisfiesPzi *bool `json:"satisfiesPzi,omitempty"`

	// Output only. Information about Generative AI model-based processor
	//  versions.
	// +kcc:proto:field=google.cloud.documentai.v1beta3.ProcessorVersion.gen_ai_model_info
	GenAiModelInfo *ProcessorVersion_GenAiModelInfo `json:"genAiModelInfo,omitempty"`
}
