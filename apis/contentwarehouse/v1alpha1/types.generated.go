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


// +kcc:proto=google.cloud.contentwarehouse.v1.DateTimeArray
type DateTimeArray struct {
	// List of datetime values.
	//  Both OffsetDateTime and ZonedDateTime are supported.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.DateTimeArray.values
	Values []DateTime `json:"values,omitempty"`
}

// +kcc:proto=google.cloud.contentwarehouse.v1.Document
type Document struct {
	// The resource name of the document.
	//  Format:
	//  projects/{project_number}/locations/{location}/documents/{document_id}.
	//
	//  The name is ignored when creating a document.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.Document.name
	Name *string `json:"name,omitempty"`

	// The reference ID set by customers. Must be unique per project and location.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.Document.reference_id
	ReferenceID *string `json:"referenceID,omitempty"`

	// Required. Display name of the document given by the user. This name will be
	//  displayed in the UI. Customer can populate this field with the name of the
	//  document. This differs from the 'title' field as 'title' is optional and
	//  stores the top heading in the document.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.Document.display_name
	DisplayName *string `json:"displayName,omitempty"`

	// Title that describes the document.
	//  This can be the top heading or text that describes the document.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.Document.title
	Title *string `json:"title,omitempty"`

	// Uri to display the document, for example, in the UI.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.Document.display_uri
	DisplayURI *string `json:"displayURI,omitempty"`

	// The Document schema name.
	//  Format:
	//  projects/{project_number}/locations/{location}/documentSchemas/{document_schema_id}.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.Document.document_schema_name
	DocumentSchemaName *string `json:"documentSchemaName,omitempty"`

	// Other document format, such as PPTX, XLXS
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.Document.plain_text
	PlainText *string `json:"plainText,omitempty"`

	// Document AI format to save the structured content, including OCR.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.Document.cloud_ai_document
	CloudAiDocument *Document `json:"cloudAiDocument,omitempty"`

	// A path linked to structured content file.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.Document.structured_content_uri
	StructuredContentURI *string `json:"structuredContentURI,omitempty"`

	// Raw document file in Cloud Storage path.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.Document.raw_document_path
	RawDocumentPath *string `json:"rawDocumentPath,omitempty"`

	// Raw document content.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.Document.inline_raw_document
	InlineRawDocument []byte `json:"inlineRawDocument,omitempty"`

	// List of values that are user supplied metadata.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.Document.properties
	Properties []Property `json:"properties,omitempty"`

	// This is used when DocAI was not used to load the document and parsing/
	//  extracting is needed for the inline_raw_document.  For example, if
	//  inline_raw_document is the byte representation of a PDF file, then
	//  this should be set to: RAW_DOCUMENT_FILE_TYPE_PDF.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.Document.raw_document_file_type
	RawDocumentFileType *string `json:"rawDocumentFileType,omitempty"`

	// If true, makes the document visible to asynchronous policies and rules.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.Document.async_enabled
	AsyncEnabled *bool `json:"asyncEnabled,omitempty"`

	// Indicates the category (image, audio, video etc.) of the original content.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.Document.content_category
	ContentCategory *string `json:"contentCategory,omitempty"`

	// If true, text extraction will not be performed.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.Document.text_extraction_disabled
	TextExtractionDisabled *bool `json:"textExtractionDisabled,omitempty"`

	// If true, text extraction will be performed.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.Document.text_extraction_enabled
	TextExtractionEnabled *bool `json:"textExtractionEnabled,omitempty"`

	// The user who creates the document.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.Document.creator
	Creator *string `json:"creator,omitempty"`

	// The user who lastly updates the document.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.Document.updater
	Updater *string `json:"updater,omitempty"`
}

// +kcc:proto=google.cloud.contentwarehouse.v1.EnumArray
type EnumArray struct {
	// List of enum values.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.EnumArray.values
	Values []string `json:"values,omitempty"`
}

// +kcc:proto=google.cloud.contentwarehouse.v1.EnumValue
type EnumValue struct {
	// String value of the enum field. This must match defined set of enums
	//  in document schema using EnumTypeOptions.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.EnumValue.value
	Value *string `json:"value,omitempty"`
}

// +kcc:proto=google.cloud.contentwarehouse.v1.FloatArray
type FloatArray struct {
	// List of float values.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.FloatArray.values
	Values []float32 `json:"values,omitempty"`
}

// +kcc:proto=google.cloud.contentwarehouse.v1.IntegerArray
type IntegerArray struct {
	// List of integer values.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.IntegerArray.values
	Values []int32 `json:"values,omitempty"`
}

// +kcc:proto=google.cloud.contentwarehouse.v1.MapProperty
type MapProperty struct {

	// TODO: unsupported map type with key string and value message

}

// +kcc:proto=google.cloud.contentwarehouse.v1.Property
type Property struct {
	// Required. Must match the name of a PropertyDefinition in the
	//  DocumentSchema.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.Property.name
	Name *string `json:"name,omitempty"`

	// Integer property values.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.Property.integer_values
	IntegerValues *IntegerArray `json:"integerValues,omitempty"`

	// Float property values.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.Property.float_values
	FloatValues *FloatArray `json:"floatValues,omitempty"`

	// String/text property values.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.Property.text_values
	TextValues *TextArray `json:"textValues,omitempty"`

	// Enum property values.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.Property.enum_values
	EnumValues *EnumArray `json:"enumValues,omitempty"`

	// Nested structured data property values.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.Property.property_values
	PropertyValues *PropertyArray `json:"propertyValues,omitempty"`

	// Date time property values.
	//  It is not supported by CMEK compliant deployment.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.Property.date_time_values
	DateTimeValues *DateTimeArray `json:"dateTimeValues,omitempty"`

	// Map property values.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.Property.map_property
	MapProperty *MapProperty `json:"mapProperty,omitempty"`

	// Timestamp property values.
	//  It is not supported by CMEK compliant deployment.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.Property.timestamp_values
	TimestampValues *TimestampArray `json:"timestampValues,omitempty"`
}

// +kcc:proto=google.cloud.contentwarehouse.v1.PropertyArray
type PropertyArray struct {
	// List of property values.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.PropertyArray.properties
	Properties []Property `json:"properties,omitempty"`
}

// +kcc:proto=google.cloud.contentwarehouse.v1.TextArray
type TextArray struct {
	// List of text values.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.TextArray.values
	Values []string `json:"values,omitempty"`
}

// +kcc:proto=google.cloud.contentwarehouse.v1.TimestampArray
type TimestampArray struct {
	// List of timestamp values.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.TimestampArray.values
	Values []TimestampValue `json:"values,omitempty"`
}

// +kcc:proto=google.cloud.contentwarehouse.v1.TimestampValue
type TimestampValue struct {
	// Timestamp value
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.TimestampValue.timestamp_value
	TimestampValue *string `json:"timestampValue,omitempty"`

	// The string must represent a valid instant in UTC and is parsed using
	//  java.time.format.DateTimeFormatter.ISO_INSTANT.
	//  e.g. "2013-09-29T18:46:19Z"
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.TimestampValue.text_value
	TextValue *string `json:"textValue,omitempty"`
}

// +kcc:proto=google.cloud.contentwarehouse.v1.Value
type Value struct {
	// Represents a float value.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.Value.float_value
	FloatValue *float32 `json:"floatValue,omitempty"`

	// Represents a integer value.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.Value.int_value
	IntValue *int32 `json:"intValue,omitempty"`

	// Represents a string value.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.Value.string_value
	StringValue *string `json:"stringValue,omitempty"`

	// Represents an enum value.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.Value.enum_value
	EnumValue *EnumValue `json:"enumValue,omitempty"`

	// Represents a datetime value.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.Value.datetime_value
	DatetimeValue *DateTime `json:"datetimeValue,omitempty"`

	// Represents a timestamp value.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.Value.timestamp_value
	TimestampValue *TimestampValue `json:"timestampValue,omitempty"`

	// Represents a boolean value.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.Value.boolean_value
	BooleanValue *bool `json:"booleanValue,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Barcode
type Barcode struct {
	// Format of a barcode.
	//  The supported formats are:
	//
	//  - `CODE_128`: Code 128 type.
	//  - `CODE_39`: Code 39 type.
	//  - `CODE_93`: Code 93 type.
	//  - `CODABAR`: Codabar type.
	//  - `DATA_MATRIX`: 2D Data Matrix type.
	//  - `ITF`: ITF type.
	//  - `EAN_13`: EAN-13 type.
	//  - `EAN_8`: EAN-8 type.
	//  - `QR_CODE`: 2D QR code type.
	//  - `UPC_A`: UPC-A type.
	//  - `UPC_E`: UPC-E type.
	//  - `PDF417`: PDF417 type.
	//  - `AZTEC`: 2D Aztec code type.
	//  - `DATABAR`: GS1 DataBar code type.
	// +kcc:proto:field=google.cloud.documentai.v1.Barcode.format
	Format *string `json:"format,omitempty"`

	// Value format describes the format of the value that a barcode
	//  encodes.
	//  The supported formats are:
	//
	//  - `CONTACT_INFO`: Contact information.
	//  - `EMAIL`: Email address.
	//  - `ISBN`: ISBN identifier.
	//  - `PHONE`: Phone number.
	//  - `PRODUCT`: Product.
	//  - `SMS`: SMS message.
	//  - `TEXT`: Text string.
	//  - `URL`: URL address.
	//  - `WIFI`: Wifi information.
	//  - `GEO`: Geo-localization.
	//  - `CALENDAR_EVENT`: Calendar event.
	//  - `DRIVER_LICENSE`: Driver's license.
	// +kcc:proto:field=google.cloud.documentai.v1.Barcode.value_format
	ValueFormat *string `json:"valueFormat,omitempty"`

	// Raw value encoded in the barcode.
	//  For example: `'MEBKM:TITLE:Google;URL:https://www.google.com;;'`.
	// +kcc:proto:field=google.cloud.documentai.v1.Barcode.raw_value
	RawValue *string `json:"rawValue,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.BoundingPoly
type BoundingPoly struct {
	// The bounding polygon vertices.
	// +kcc:proto:field=google.cloud.documentai.v1.BoundingPoly.vertices
	Vertices []Vertex `json:"vertices,omitempty"`

	// The bounding polygon normalized vertices.
	// +kcc:proto:field=google.cloud.documentai.v1.BoundingPoly.normalized_vertices
	NormalizedVertices []NormalizedVertex `json:"normalizedVertices,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Document
type Document struct {
	// Optional. Currently supports Google Cloud Storage URI of the form
	//  `gs://bucket_name/object_name`. Object versioning is not supported.
	//  For more information, refer to [Google Cloud Storage Request
	//  URIs](https://cloud.google.com/storage/docs/reference-uris).
	// +kcc:proto:field=google.cloud.documentai.v1.Document.uri
	URI *string `json:"uri,omitempty"`

	// Optional. Inline document content, represented as a stream of bytes.
	//  Note: As with all `bytes` fields, protobuffers use a pure binary
	//  representation, whereas JSON representations use base64.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.content
	Content []byte `json:"content,omitempty"`

	// An IANA published [media type (MIME
	//  type)](https://www.iana.org/assignments/media-types/media-types.xhtml).
	// +kcc:proto:field=google.cloud.documentai.v1.Document.mime_type
	MimeType *string `json:"mimeType,omitempty"`

	// Optional. UTF-8 encoded text in reading order from the document.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.text
	Text *string `json:"text,omitempty"`

	// Styles for the [Document.text][google.cloud.documentai.v1.Document.text].
	// +kcc:proto:field=google.cloud.documentai.v1.Document.text_styles
	TextStyles []Document_Style `json:"textStyles,omitempty"`

	// Visual page layout for the [Document][google.cloud.documentai.v1.Document].
	// +kcc:proto:field=google.cloud.documentai.v1.Document.pages
	Pages []Document_Page `json:"pages,omitempty"`

	// A list of entities detected on
	//  [Document.text][google.cloud.documentai.v1.Document.text]. For document
	//  shards, entities in this list may cross shard boundaries.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.entities
	Entities []Document_Entity `json:"entities,omitempty"`

	// Placeholder.  Relationship among
	//  [Document.entities][google.cloud.documentai.v1.Document.entities].
	// +kcc:proto:field=google.cloud.documentai.v1.Document.entity_relations
	EntityRelations []Document_EntityRelation `json:"entityRelations,omitempty"`

	// Placeholder.  A list of text corrections made to
	//  [Document.text][google.cloud.documentai.v1.Document.text].  This is usually
	//  used for annotating corrections to OCR mistakes.  Text changes for a given
	//  revision may not overlap with each other.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.text_changes
	TextChanges []Document_TextChange `json:"textChanges,omitempty"`

	// Information about the sharding if this document is sharded part of a larger
	//  document. If the document is not sharded, this message is not specified.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.shard_info
	ShardInfo *Document_ShardInfo `json:"shardInfo,omitempty"`

	// Any error that occurred while processing this document.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.error
	Error *Status `json:"error,omitempty"`

	// Placeholder. Revision history of this document.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.revisions
	Revisions []Document_Revision `json:"revisions,omitempty"`

	// Parsed layout of the document.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.document_layout
	DocumentLayout *Document_DocumentLayout `json:"documentLayout,omitempty"`

	// Document chunked based on chunking config.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.chunked_document
	ChunkedDocument *Document_ChunkedDocument `json:"chunkedDocument,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Document.ChunkedDocument
type Document_ChunkedDocument struct {
	// List of chunks.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.ChunkedDocument.chunks
	Chunks []Document_ChunkedDocument_Chunk `json:"chunks,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Document.ChunkedDocument.Chunk
type Document_ChunkedDocument_Chunk struct {
	// ID of the chunk.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.ChunkedDocument.Chunk.chunk_id
	ChunkID *string `json:"chunkID,omitempty"`

	// Unused.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.ChunkedDocument.Chunk.source_block_ids
	SourceBlockIds []string `json:"sourceBlockIds,omitempty"`

	// Text content of the chunk.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.ChunkedDocument.Chunk.content
	Content *string `json:"content,omitempty"`

	// Page span of the chunk.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.ChunkedDocument.Chunk.page_span
	PageSpan *Document_ChunkedDocument_Chunk_ChunkPageSpan `json:"pageSpan,omitempty"`

	// Page headers associated with the chunk.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.ChunkedDocument.Chunk.page_headers
	PageHeaders []Document_ChunkedDocument_Chunk_ChunkPageHeader `json:"pageHeaders,omitempty"`

	// Page footers associated with the chunk.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.ChunkedDocument.Chunk.page_footers
	PageFooters []Document_ChunkedDocument_Chunk_ChunkPageFooter `json:"pageFooters,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Document.ChunkedDocument.Chunk.ChunkPageFooter
type Document_ChunkedDocument_Chunk_ChunkPageFooter struct {
	// Footer in text format.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.ChunkedDocument.Chunk.ChunkPageFooter.text
	Text *string `json:"text,omitempty"`

	// Page span of the footer.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.ChunkedDocument.Chunk.ChunkPageFooter.page_span
	PageSpan *Document_ChunkedDocument_Chunk_ChunkPageSpan `json:"pageSpan,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Document.ChunkedDocument.Chunk.ChunkPageHeader
type Document_ChunkedDocument_Chunk_ChunkPageHeader struct {
	// Header in text format.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.ChunkedDocument.Chunk.ChunkPageHeader.text
	Text *string `json:"text,omitempty"`

	// Page span of the header.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.ChunkedDocument.Chunk.ChunkPageHeader.page_span
	PageSpan *Document_ChunkedDocument_Chunk_ChunkPageSpan `json:"pageSpan,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Document.ChunkedDocument.Chunk.ChunkPageSpan
type Document_ChunkedDocument_Chunk_ChunkPageSpan struct {
	// Page where chunk starts in the document.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.ChunkedDocument.Chunk.ChunkPageSpan.page_start
	PageStart *int32 `json:"pageStart,omitempty"`

	// Page where chunk ends in the document.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.ChunkedDocument.Chunk.ChunkPageSpan.page_end
	PageEnd *int32 `json:"pageEnd,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Document.DocumentLayout
type Document_DocumentLayout struct {
	// List of blocks in the document.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.DocumentLayout.blocks
	Blocks []Document_DocumentLayout_DocumentLayoutBlock `json:"blocks,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Document.DocumentLayout.DocumentLayoutBlock
type Document_DocumentLayout_DocumentLayoutBlock struct {
	// Block consisting of text content.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.DocumentLayout.DocumentLayoutBlock.text_block
	TextBlock *Document_DocumentLayout_DocumentLayoutBlock_LayoutTextBlock `json:"textBlock,omitempty"`

	// Block consisting of table content/structure.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.DocumentLayout.DocumentLayoutBlock.table_block
	TableBlock *Document_DocumentLayout_DocumentLayoutBlock_LayoutTableBlock `json:"tableBlock,omitempty"`

	// Block consisting of list content/structure.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.DocumentLayout.DocumentLayoutBlock.list_block
	ListBlock *Document_DocumentLayout_DocumentLayoutBlock_LayoutListBlock `json:"listBlock,omitempty"`

	// ID of the block.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.DocumentLayout.DocumentLayoutBlock.block_id
	BlockID *string `json:"blockID,omitempty"`

	// Page span of the block.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.DocumentLayout.DocumentLayoutBlock.page_span
	PageSpan *Document_DocumentLayout_DocumentLayoutBlock_LayoutPageSpan `json:"pageSpan,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Document.DocumentLayout.DocumentLayoutBlock.LayoutListBlock
type Document_DocumentLayout_DocumentLayoutBlock_LayoutListBlock struct {
	// List entries that constitute a list block.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.DocumentLayout.DocumentLayoutBlock.LayoutListBlock.list_entries
	ListEntries []Document_DocumentLayout_DocumentLayoutBlock_LayoutListEntry `json:"listEntries,omitempty"`

	// Type of the list_entries (if exist). Available options are `ordered`
	//  and `unordered`.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.DocumentLayout.DocumentLayoutBlock.LayoutListBlock.type
	Type *string `json:"type,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Document.DocumentLayout.DocumentLayoutBlock.LayoutListEntry
type Document_DocumentLayout_DocumentLayoutBlock_LayoutListEntry struct {
	// A list entry is a list of blocks.
	//  Repeated blocks support further hierarchies and nested blocks.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.DocumentLayout.DocumentLayoutBlock.LayoutListEntry.blocks
	Blocks []Document_DocumentLayout_DocumentLayoutBlock `json:"blocks,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Document.DocumentLayout.DocumentLayoutBlock.LayoutPageSpan
type Document_DocumentLayout_DocumentLayoutBlock_LayoutPageSpan struct {
	// Page where block starts in the document.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.DocumentLayout.DocumentLayoutBlock.LayoutPageSpan.page_start
	PageStart *int32 `json:"pageStart,omitempty"`

	// Page where block ends in the document.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.DocumentLayout.DocumentLayoutBlock.LayoutPageSpan.page_end
	PageEnd *int32 `json:"pageEnd,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Document.DocumentLayout.DocumentLayoutBlock.LayoutTableBlock
type Document_DocumentLayout_DocumentLayoutBlock_LayoutTableBlock struct {
	// Header rows at the top of the table.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.DocumentLayout.DocumentLayoutBlock.LayoutTableBlock.header_rows
	HeaderRows []Document_DocumentLayout_DocumentLayoutBlock_LayoutTableRow `json:"headerRows,omitempty"`

	// Body rows containing main table content.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.DocumentLayout.DocumentLayoutBlock.LayoutTableBlock.body_rows
	BodyRows []Document_DocumentLayout_DocumentLayoutBlock_LayoutTableRow `json:"bodyRows,omitempty"`

	// Table caption/title.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.DocumentLayout.DocumentLayoutBlock.LayoutTableBlock.caption
	Caption *string `json:"caption,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Document.DocumentLayout.DocumentLayoutBlock.LayoutTableCell
type Document_DocumentLayout_DocumentLayoutBlock_LayoutTableCell struct {
	// A table cell is a list of blocks.
	//  Repeated blocks support further hierarchies and nested blocks.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.DocumentLayout.DocumentLayoutBlock.LayoutTableCell.blocks
	Blocks []Document_DocumentLayout_DocumentLayoutBlock `json:"blocks,omitempty"`

	// How many rows this cell spans.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.DocumentLayout.DocumentLayoutBlock.LayoutTableCell.row_span
	RowSpan *int32 `json:"rowSpan,omitempty"`

	// How many columns this cell spans.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.DocumentLayout.DocumentLayoutBlock.LayoutTableCell.col_span
	ColSpan *int32 `json:"colSpan,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Document.DocumentLayout.DocumentLayoutBlock.LayoutTableRow
type Document_DocumentLayout_DocumentLayoutBlock_LayoutTableRow struct {
	// A table row is a list of table cells.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.DocumentLayout.DocumentLayoutBlock.LayoutTableRow.cells
	Cells []Document_DocumentLayout_DocumentLayoutBlock_LayoutTableCell `json:"cells,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Document.DocumentLayout.DocumentLayoutBlock.LayoutTextBlock
type Document_DocumentLayout_DocumentLayoutBlock_LayoutTextBlock struct {
	// Text content stored in the block.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.DocumentLayout.DocumentLayoutBlock.LayoutTextBlock.text
	Text *string `json:"text,omitempty"`

	// Type of the text in the block. Available options are: `paragraph`,
	//  `subtitle`, `heading-1`, `heading-2`, `heading-3`, `heading-4`,
	//  `heading-5`, `header`, `footer`.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.DocumentLayout.DocumentLayoutBlock.LayoutTextBlock.type
	Type *string `json:"type,omitempty"`

	// A text block could further have child blocks.
	//  Repeated blocks support further hierarchies and nested blocks.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.DocumentLayout.DocumentLayoutBlock.LayoutTextBlock.blocks
	Blocks []Document_DocumentLayout_DocumentLayoutBlock `json:"blocks,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Document.Entity
type Document_Entity struct {
	// Optional. Provenance of the entity.
	//  Text anchor indexing into the
	//  [Document.text][google.cloud.documentai.v1.Document.text].
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Entity.text_anchor
	TextAnchor *Document_TextAnchor `json:"textAnchor,omitempty"`

	// Required. Entity type from a schema e.g. `Address`.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Entity.type
	Type *string `json:"type,omitempty"`

	// Optional. Text value of the entity e.g. `1600 Amphitheatre Pkwy`.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Entity.mention_text
	MentionText *string `json:"mentionText,omitempty"`

	// Optional. Deprecated.  Use `id` field instead.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Entity.mention_id
	MentionID *string `json:"mentionID,omitempty"`

	// Optional. Confidence of detected Schema entity. Range `[0, 1]`.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Entity.confidence
	Confidence *float32 `json:"confidence,omitempty"`

	// Optional. Represents the provenance of this entity wrt. the location on
	//  the page where it was found.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Entity.page_anchor
	PageAnchor *Document_PageAnchor `json:"pageAnchor,omitempty"`

	// Optional. Canonical id. This will be a unique value in the entity list
	//  for this document.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Entity.id
	ID *string `json:"id,omitempty"`

	// Optional. Normalized entity value. Absent if the extracted value could
	//  not be converted or the type (e.g. address) is not supported for certain
	//  parsers. This field is also only populated for certain supported document
	//  types.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Entity.normalized_value
	NormalizedValue *Document_Entity_NormalizedValue `json:"normalizedValue,omitempty"`

	// Optional. Entities can be nested to form a hierarchical data structure
	//  representing the content in the document.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Entity.properties
	Properties []Document_Entity `json:"properties,omitempty"`

	// Optional. The history of this annotation.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Entity.provenance
	Provenance *Document_Provenance `json:"provenance,omitempty"`

	// Optional. Whether the entity will be redacted for de-identification
	//  purposes.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Entity.redacted
	Redacted *bool `json:"redacted,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Document.Entity.NormalizedValue
type Document_Entity_NormalizedValue struct {
	// Money value. See also:
	//  https://github.com/googleapis/googleapis/blob/master/google/type/money.proto
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Entity.NormalizedValue.money_value
	MoneyValue *Money `json:"moneyValue,omitempty"`

	// Date value. Includes year, month, day. See also:
	//  https://github.com/googleapis/googleapis/blob/master/google/type/date.proto
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Entity.NormalizedValue.date_value
	DateValue *Date `json:"dateValue,omitempty"`

	// DateTime value. Includes date, time, and timezone. See also:
	//  https://github.com/googleapis/googleapis/blob/master/google/type/datetime.proto
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Entity.NormalizedValue.datetime_value
	DatetimeValue *DateTime `json:"datetimeValue,omitempty"`

	// Postal address. See also:
	//  https://github.com/googleapis/googleapis/blob/master/google/type/postal_address.proto
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Entity.NormalizedValue.address_value
	AddressValue *PostalAddress `json:"addressValue,omitempty"`

	// Boolean value. Can be used for entities with binary values, or for
	//  checkboxes.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Entity.NormalizedValue.boolean_value
	BooleanValue *bool `json:"booleanValue,omitempty"`

	// Integer value.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Entity.NormalizedValue.integer_value
	IntegerValue *int32 `json:"integerValue,omitempty"`

	// Float value.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Entity.NormalizedValue.float_value
	FloatValue *float32 `json:"floatValue,omitempty"`

	// Optional. An optional field to store a normalized string.
	//  For some entity types, one of respective `structured_value` fields may
	//  also be populated. Also not all the types of `structured_value` will be
	//  normalized. For example, some processors may not generate `float`
	//  or `integer` normalized text by default.
	//
	//  Below are sample formats mapped to structured values.
	//
	//  - Money/Currency type (`money_value`) is in the ISO 4217 text format.
	//  - Date type (`date_value`) is in the ISO 8601 text format.
	//  - Datetime type (`datetime_value`) is in the ISO 8601 text format.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Entity.NormalizedValue.text
	Text *string `json:"text,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Document.EntityRelation
type Document_EntityRelation struct {
	// Subject entity id.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.EntityRelation.subject_id
	SubjectID *string `json:"subjectID,omitempty"`

	// Object entity id.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.EntityRelation.object_id
	ObjectID *string `json:"objectID,omitempty"`

	// Relationship description.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.EntityRelation.relation
	Relation *string `json:"relation,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Document.Page
type Document_Page struct {
	// 1-based index for current
	//  [Page][google.cloud.documentai.v1.Document.Page] in a parent
	//  [Document][google.cloud.documentai.v1.Document]. Useful when a page is
	//  taken out of a [Document][google.cloud.documentai.v1.Document] for
	//  individual processing.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.page_number
	PageNumber *int32 `json:"pageNumber,omitempty"`

	// Rendered image for this page. This image is preprocessed to remove any
	//  skew, rotation, and distortions such that the annotation bounding boxes
	//  can be upright and axis-aligned.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.image
	Image *Document_Page_Image `json:"image,omitempty"`

	// Transformation matrices that were applied to the original document image
	//  to produce [Page.image][google.cloud.documentai.v1.Document.Page.image].
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.transforms
	Transforms []Document_Page_Matrix `json:"transforms,omitempty"`

	// Physical dimension of the page.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.dimension
	Dimension *Document_Page_Dimension `json:"dimension,omitempty"`

	// [Layout][google.cloud.documentai.v1.Document.Page.Layout] for the page.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.layout
	Layout *Document_Page_Layout `json:"layout,omitempty"`

	// A list of detected languages together with confidence.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.detected_languages
	DetectedLanguages []Document_Page_DetectedLanguage `json:"detectedLanguages,omitempty"`

	// A list of visually detected text blocks on the page.
	//  A block has a set of lines (collected into paragraphs) that have a common
	//  line-spacing and orientation.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.blocks
	Blocks []Document_Page_Block `json:"blocks,omitempty"`

	// A list of visually detected text paragraphs on the page.
	//  A collection of lines that a human would perceive as a paragraph.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.paragraphs
	Paragraphs []Document_Page_Paragraph `json:"paragraphs,omitempty"`

	// A list of visually detected text lines on the page.
	//  A collection of tokens that a human would perceive as a line.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.lines
	Lines []Document_Page_Line `json:"lines,omitempty"`

	// A list of visually detected tokens on the page.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.tokens
	Tokens []Document_Page_Token `json:"tokens,omitempty"`

	// A list of detected non-text visual elements e.g. checkbox,
	//  signature etc. on the page.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.visual_elements
	VisualElements []Document_Page_VisualElement `json:"visualElements,omitempty"`

	// A list of visually detected tables on the page.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.tables
	Tables []Document_Page_Table `json:"tables,omitempty"`

	// A list of visually detected form fields on the page.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.form_fields
	FormFields []Document_Page_FormField `json:"formFields,omitempty"`

	// A list of visually detected symbols on the page.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.symbols
	Symbols []Document_Page_Symbol `json:"symbols,omitempty"`

	// A list of detected barcodes.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.detected_barcodes
	DetectedBarcodes []Document_Page_DetectedBarcode `json:"detectedBarcodes,omitempty"`

	// Image quality scores.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.image_quality_scores
	ImageQualityScores *Document_Page_ImageQualityScores `json:"imageQualityScores,omitempty"`

	// The history of this page.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.provenance
	Provenance *Document_Provenance `json:"provenance,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Document.Page.Block
type Document_Page_Block struct {
	// [Layout][google.cloud.documentai.v1.Document.Page.Layout] for
	//  [Block][google.cloud.documentai.v1.Document.Page.Block].
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.Block.layout
	Layout *Document_Page_Layout `json:"layout,omitempty"`

	// A list of detected languages together with confidence.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.Block.detected_languages
	DetectedLanguages []Document_Page_DetectedLanguage `json:"detectedLanguages,omitempty"`

	// The history of this annotation.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.Block.provenance
	Provenance *Document_Provenance `json:"provenance,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Document.Page.DetectedBarcode
type Document_Page_DetectedBarcode struct {
	// [Layout][google.cloud.documentai.v1.Document.Page.Layout] for
	//  [DetectedBarcode][google.cloud.documentai.v1.Document.Page.DetectedBarcode].
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.DetectedBarcode.layout
	Layout *Document_Page_Layout `json:"layout,omitempty"`

	// Detailed barcode information of the
	//  [DetectedBarcode][google.cloud.documentai.v1.Document.Page.DetectedBarcode].
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.DetectedBarcode.barcode
	Barcode *Barcode `json:"barcode,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Document.Page.DetectedLanguage
type Document_Page_DetectedLanguage struct {
	// The [BCP-47 language
	//  code](https://www.unicode.org/reports/tr35/#Unicode_locale_identifier),
	//  such as `en-US` or `sr-Latn`.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.DetectedLanguage.language_code
	LanguageCode *string `json:"languageCode,omitempty"`

	// Confidence of detected language. Range `[0, 1]`.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.DetectedLanguage.confidence
	Confidence *float32 `json:"confidence,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Document.Page.Dimension
type Document_Page_Dimension struct {
	// Page width.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.Dimension.width
	Width *float32 `json:"width,omitempty"`

	// Page height.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.Dimension.height
	Height *float32 `json:"height,omitempty"`

	// Dimension unit.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.Dimension.unit
	Unit *string `json:"unit,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Document.Page.FormField
type Document_Page_FormField struct {
	// [Layout][google.cloud.documentai.v1.Document.Page.Layout] for the
	//  [FormField][google.cloud.documentai.v1.Document.Page.FormField] name.
	//  e.g. `Address`, `Email`, `Grand total`, `Phone number`, etc.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.FormField.field_name
	FieldName *Document_Page_Layout `json:"fieldName,omitempty"`

	// [Layout][google.cloud.documentai.v1.Document.Page.Layout] for the
	//  [FormField][google.cloud.documentai.v1.Document.Page.FormField] value.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.FormField.field_value
	FieldValue *Document_Page_Layout `json:"fieldValue,omitempty"`

	// A list of detected languages for name together with confidence.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.FormField.name_detected_languages
	NameDetectedLanguages []Document_Page_DetectedLanguage `json:"nameDetectedLanguages,omitempty"`

	// A list of detected languages for value together with confidence.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.FormField.value_detected_languages
	ValueDetectedLanguages []Document_Page_DetectedLanguage `json:"valueDetectedLanguages,omitempty"`

	// If the value is non-textual, this field represents the type. Current
	//  valid values are:
	//
	//  - blank (this indicates the `field_value` is normal text)
	//  - `unfilled_checkbox`
	//  - `filled_checkbox`
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.FormField.value_type
	ValueType *string `json:"valueType,omitempty"`

	// Created for Labeling UI to export key text.
	//  If corrections were made to the text identified by the
	//  `field_name.text_anchor`, this field will contain the correction.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.FormField.corrected_key_text
	CorrectedKeyText *string `json:"correctedKeyText,omitempty"`

	// Created for Labeling UI to export value text.
	//  If corrections were made to the text identified by the
	//  `field_value.text_anchor`, this field will contain the correction.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.FormField.corrected_value_text
	CorrectedValueText *string `json:"correctedValueText,omitempty"`

	// The history of this annotation.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.FormField.provenance
	Provenance *Document_Provenance `json:"provenance,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Document.Page.Image
type Document_Page_Image struct {
	// Raw byte content of the image.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.Image.content
	Content []byte `json:"content,omitempty"`

	// Encoding [media type (MIME
	//  type)](https://www.iana.org/assignments/media-types/media-types.xhtml)
	//  for the image.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.Image.mime_type
	MimeType *string `json:"mimeType,omitempty"`

	// Width of the image in pixels.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.Image.width
	Width *int32 `json:"width,omitempty"`

	// Height of the image in pixels.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.Image.height
	Height *int32 `json:"height,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Document.Page.ImageQualityScores
type Document_Page_ImageQualityScores struct {
	// The overall quality score. Range `[0, 1]` where `1` is perfect quality.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.ImageQualityScores.quality_score
	QualityScore *float32 `json:"qualityScore,omitempty"`

	// A list of detected defects.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.ImageQualityScores.detected_defects
	DetectedDefects []Document_Page_ImageQualityScores_DetectedDefect `json:"detectedDefects,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Document.Page.ImageQualityScores.DetectedDefect
type Document_Page_ImageQualityScores_DetectedDefect struct {
	// Name of the defect type. Supported values are:
	//
	//  - `quality/defect_blurry`
	//  - `quality/defect_noisy`
	//  - `quality/defect_dark`
	//  - `quality/defect_faint`
	//  - `quality/defect_text_too_small`
	//  - `quality/defect_document_cutoff`
	//  - `quality/defect_text_cutoff`
	//  - `quality/defect_glare`
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.ImageQualityScores.DetectedDefect.type
	Type *string `json:"type,omitempty"`

	// Confidence of detected defect. Range `[0, 1]` where `1` indicates
	//  strong confidence that the defect exists.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.ImageQualityScores.DetectedDefect.confidence
	Confidence *float32 `json:"confidence,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Document.Page.Layout
type Document_Page_Layout struct {
	// Text anchor indexing into the
	//  [Document.text][google.cloud.documentai.v1.Document.text].
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.Layout.text_anchor
	TextAnchor *Document_TextAnchor `json:"textAnchor,omitempty"`

	// Confidence of the current
	//  [Layout][google.cloud.documentai.v1.Document.Page.Layout] within
	//  context of the object this layout is for. e.g. confidence can be for a
	//  single token, a table, a visual element, etc. depending on context.
	//  Range `[0, 1]`.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.Layout.confidence
	Confidence *float32 `json:"confidence,omitempty"`

	// The bounding polygon for the
	//  [Layout][google.cloud.documentai.v1.Document.Page.Layout].
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.Layout.bounding_poly
	BoundingPoly *BoundingPoly `json:"boundingPoly,omitempty"`

	// Detected orientation for the
	//  [Layout][google.cloud.documentai.v1.Document.Page.Layout].
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.Layout.orientation
	Orientation *string `json:"orientation,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Document.Page.Line
type Document_Page_Line struct {
	// [Layout][google.cloud.documentai.v1.Document.Page.Layout] for
	//  [Line][google.cloud.documentai.v1.Document.Page.Line].
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.Line.layout
	Layout *Document_Page_Layout `json:"layout,omitempty"`

	// A list of detected languages together with confidence.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.Line.detected_languages
	DetectedLanguages []Document_Page_DetectedLanguage `json:"detectedLanguages,omitempty"`

	// The  history of this annotation.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.Line.provenance
	Provenance *Document_Provenance `json:"provenance,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Document.Page.Matrix
type Document_Page_Matrix struct {
	// Number of rows in the matrix.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.Matrix.rows
	Rows *int32 `json:"rows,omitempty"`

	// Number of columns in the matrix.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.Matrix.cols
	Cols *int32 `json:"cols,omitempty"`

	// This encodes information about what data type the matrix uses.
	//  For example, 0 (CV_8U) is an unsigned 8-bit image. For the full list
	//  of OpenCV primitive data types, please refer to
	//  https://docs.opencv.org/4.3.0/d1/d1b/group__core__hal__interface.html
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.Matrix.type
	Type *int32 `json:"type,omitempty"`

	// The matrix data.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.Matrix.data
	Data []byte `json:"data,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Document.Page.Paragraph
type Document_Page_Paragraph struct {
	// [Layout][google.cloud.documentai.v1.Document.Page.Layout] for
	//  [Paragraph][google.cloud.documentai.v1.Document.Page.Paragraph].
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.Paragraph.layout
	Layout *Document_Page_Layout `json:"layout,omitempty"`

	// A list of detected languages together with confidence.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.Paragraph.detected_languages
	DetectedLanguages []Document_Page_DetectedLanguage `json:"detectedLanguages,omitempty"`

	// The  history of this annotation.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.Paragraph.provenance
	Provenance *Document_Provenance `json:"provenance,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Document.Page.Symbol
type Document_Page_Symbol struct {
	// [Layout][google.cloud.documentai.v1.Document.Page.Layout] for
	//  [Symbol][google.cloud.documentai.v1.Document.Page.Symbol].
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.Symbol.layout
	Layout *Document_Page_Layout `json:"layout,omitempty"`

	// A list of detected languages together with confidence.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.Symbol.detected_languages
	DetectedLanguages []Document_Page_DetectedLanguage `json:"detectedLanguages,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Document.Page.Table
type Document_Page_Table struct {
	// [Layout][google.cloud.documentai.v1.Document.Page.Layout] for
	//  [Table][google.cloud.documentai.v1.Document.Page.Table].
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.Table.layout
	Layout *Document_Page_Layout `json:"layout,omitempty"`

	// Header rows of the table.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.Table.header_rows
	HeaderRows []Document_Page_Table_TableRow `json:"headerRows,omitempty"`

	// Body rows of the table.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.Table.body_rows
	BodyRows []Document_Page_Table_TableRow `json:"bodyRows,omitempty"`

	// A list of detected languages together with confidence.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.Table.detected_languages
	DetectedLanguages []Document_Page_DetectedLanguage `json:"detectedLanguages,omitempty"`

	// The history of this table.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.Table.provenance
	Provenance *Document_Provenance `json:"provenance,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Document.Page.Table.TableCell
type Document_Page_Table_TableCell struct {
	// [Layout][google.cloud.documentai.v1.Document.Page.Layout] for
	//  [TableCell][google.cloud.documentai.v1.Document.Page.Table.TableCell].
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.Table.TableCell.layout
	Layout *Document_Page_Layout `json:"layout,omitempty"`

	// How many rows this cell spans.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.Table.TableCell.row_span
	RowSpan *int32 `json:"rowSpan,omitempty"`

	// How many columns this cell spans.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.Table.TableCell.col_span
	ColSpan *int32 `json:"colSpan,omitempty"`

	// A list of detected languages together with confidence.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.Table.TableCell.detected_languages
	DetectedLanguages []Document_Page_DetectedLanguage `json:"detectedLanguages,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Document.Page.Table.TableRow
type Document_Page_Table_TableRow struct {
	// Cells that make up this row.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.Table.TableRow.cells
	Cells []Document_Page_Table_TableCell `json:"cells,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Document.Page.Token
type Document_Page_Token struct {
	// [Layout][google.cloud.documentai.v1.Document.Page.Layout] for
	//  [Token][google.cloud.documentai.v1.Document.Page.Token].
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.Token.layout
	Layout *Document_Page_Layout `json:"layout,omitempty"`

	// Detected break at the end of a
	//  [Token][google.cloud.documentai.v1.Document.Page.Token].
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.Token.detected_break
	DetectedBreak *Document_Page_Token_DetectedBreak `json:"detectedBreak,omitempty"`

	// A list of detected languages together with confidence.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.Token.detected_languages
	DetectedLanguages []Document_Page_DetectedLanguage `json:"detectedLanguages,omitempty"`

	// The history of this annotation.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.Token.provenance
	Provenance *Document_Provenance `json:"provenance,omitempty"`

	// Text style attributes.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.Token.style_info
	StyleInfo *Document_Page_Token_StyleInfo `json:"styleInfo,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Document.Page.Token.DetectedBreak
type Document_Page_Token_DetectedBreak struct {
	// Detected break type.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.Token.DetectedBreak.type
	Type *string `json:"type,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Document.Page.Token.StyleInfo
type Document_Page_Token_StyleInfo struct {
	// Font size in points (`1` point is `¹⁄₇₂` inches).
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.Token.StyleInfo.font_size
	FontSize *int32 `json:"fontSize,omitempty"`

	// Font size in pixels, equal to _unrounded
	//  [font_size][google.cloud.documentai.v1.Document.Page.Token.StyleInfo.font_size]_
	//  * _resolution_ ÷ `72.0`.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.Token.StyleInfo.pixel_font_size
	PixelFontSize *float64 `json:"pixelFontSize,omitempty"`

	// Letter spacing in points.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.Token.StyleInfo.letter_spacing
	LetterSpacing *float64 `json:"letterSpacing,omitempty"`

	// Name or style of the font.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.Token.StyleInfo.font_type
	FontType *string `json:"fontType,omitempty"`

	// Whether the text is bold (equivalent to
	//  [font_weight][google.cloud.documentai.v1.Document.Page.Token.StyleInfo.font_weight]
	//  is at least `700`).
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.Token.StyleInfo.bold
	Bold *bool `json:"bold,omitempty"`

	// Whether the text is italic.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.Token.StyleInfo.italic
	Italic *bool `json:"italic,omitempty"`

	// Whether the text is underlined.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.Token.StyleInfo.underlined
	Underlined *bool `json:"underlined,omitempty"`

	// Whether the text is strikethrough. This feature is not supported yet.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.Token.StyleInfo.strikeout
	Strikeout *bool `json:"strikeout,omitempty"`

	// Whether the text is a subscript. This feature is not supported yet.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.Token.StyleInfo.subscript
	Subscript *bool `json:"subscript,omitempty"`

	// Whether the text is a superscript. This feature is not supported yet.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.Token.StyleInfo.superscript
	Superscript *bool `json:"superscript,omitempty"`

	// Whether the text is in small caps. This feature is not supported yet.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.Token.StyleInfo.smallcaps
	Smallcaps *bool `json:"smallcaps,omitempty"`

	// TrueType weight on a scale `100` (thin) to `1000` (ultra-heavy).
	//  Normal is `400`, bold is `700`.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.Token.StyleInfo.font_weight
	FontWeight *int32 `json:"fontWeight,omitempty"`

	// Whether the text is handwritten.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.Token.StyleInfo.handwritten
	Handwritten *bool `json:"handwritten,omitempty"`

	// Color of the text.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.Token.StyleInfo.text_color
	TextColor *Color `json:"textColor,omitempty"`

	// Color of the background.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.Token.StyleInfo.background_color
	BackgroundColor *Color `json:"backgroundColor,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Document.Page.VisualElement
type Document_Page_VisualElement struct {
	// [Layout][google.cloud.documentai.v1.Document.Page.Layout] for
	//  [VisualElement][google.cloud.documentai.v1.Document.Page.VisualElement].
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.VisualElement.layout
	Layout *Document_Page_Layout `json:"layout,omitempty"`

	// Type of the
	//  [VisualElement][google.cloud.documentai.v1.Document.Page.VisualElement].
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.VisualElement.type
	Type *string `json:"type,omitempty"`

	// A list of detected languages together with confidence.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Page.VisualElement.detected_languages
	DetectedLanguages []Document_Page_DetectedLanguage `json:"detectedLanguages,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Document.PageAnchor
type Document_PageAnchor struct {
	// One or more references to visual page elements
	// +kcc:proto:field=google.cloud.documentai.v1.Document.PageAnchor.page_refs
	PageRefs []Document_PageAnchor_PageRef `json:"pageRefs,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Document.PageAnchor.PageRef
type Document_PageAnchor_PageRef struct {
	// Required. Index into the
	//  [Document.pages][google.cloud.documentai.v1.Document.pages] element,
	//  for example using
	//  `[Document.pages][page_refs.page]` to locate the related page element.
	//  This field is skipped when its value is the default `0`. See
	//  https://developers.google.com/protocol-buffers/docs/proto3#json.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.PageAnchor.PageRef.page
	Page *int64 `json:"page,omitempty"`

	// Optional. The type of the layout element that is being referenced if
	//  any.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.PageAnchor.PageRef.layout_type
	LayoutType *string `json:"layoutType,omitempty"`

	// Optional. Deprecated.  Use
	//  [PageRef.bounding_poly][google.cloud.documentai.v1.Document.PageAnchor.PageRef.bounding_poly]
	//  instead.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.PageAnchor.PageRef.layout_id
	LayoutID *string `json:"layoutID,omitempty"`

	// Optional. Identifies the bounding polygon of a layout element on the
	//  page. If `layout_type` is set, the bounding polygon must be exactly the
	//  same to the layout element it's referring to.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.PageAnchor.PageRef.bounding_poly
	BoundingPoly *BoundingPoly `json:"boundingPoly,omitempty"`

	// Optional. Confidence of detected page element, if applicable. Range
	//  `[0, 1]`.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.PageAnchor.PageRef.confidence
	Confidence *float32 `json:"confidence,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Document.Provenance
type Document_Provenance struct {
	// The index of the revision that produced this element.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Provenance.revision
	Revision *int32 `json:"revision,omitempty"`

	// The Id of this operation.  Needs to be unique within the scope of the
	//  revision.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Provenance.id
	ID *int32 `json:"id,omitempty"`

	// References to the original elements that are replaced.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Provenance.parents
	Parents []Document_Provenance_Parent `json:"parents,omitempty"`

	// The type of provenance operation.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Provenance.type
	Type *string `json:"type,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Document.Provenance.Parent
type Document_Provenance_Parent struct {
	// The index of the index into current revision's parent_ids list.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Provenance.Parent.revision
	Revision *int32 `json:"revision,omitempty"`

	// The index of the parent item in the corresponding item list (eg. list
	//  of entities, properties within entities, etc.) in the parent revision.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Provenance.Parent.index
	Index *int32 `json:"index,omitempty"`

	// The id of the parent provenance.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Provenance.Parent.id
	ID *int32 `json:"id,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Document.Revision
type Document_Revision struct {
	// If the change was made by a person specify the name or id of that
	//  person.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Revision.agent
	Agent *string `json:"agent,omitempty"`

	// If the annotation was made by processor identify the processor by its
	//  resource name.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Revision.processor
	Processor *string `json:"processor,omitempty"`

	// Id of the revision, internally generated by doc proto storage.
	//  Unique within the context of the document.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Revision.id
	ID *string `json:"id,omitempty"`

	// The revisions that this revision is based on.  This can include one or
	//  more parent (when documents are merged.)  This field represents the
	//  index into the `revisions` field.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Revision.parent
	Parent []int32 `json:"parent,omitempty"`

	// The revisions that this revision is based on. Must include all the ids
	//  that have anything to do with this revision - eg. there are
	//  `provenance.parent.revision` fields that index into this field.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Revision.parent_ids
	ParentIds []string `json:"parentIds,omitempty"`

	// The time that the revision was created, internally generated by
	//  doc proto storage at the time of create.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Revision.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Human Review information of this revision.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Revision.human_review
	HumanReview *Document_Revision_HumanReview `json:"humanReview,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Document.Revision.HumanReview
type Document_Revision_HumanReview struct {
	// Human review state. e.g. `requested`, `succeeded`, `rejected`.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Revision.HumanReview.state
	State *string `json:"state,omitempty"`

	// A message providing more details about the current state of processing.
	//  For example, the rejection reason when the state is `rejected`.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Revision.HumanReview.state_message
	StateMessage *string `json:"stateMessage,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Document.ShardInfo
type Document_ShardInfo struct {
	// The 0-based index of this shard.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.ShardInfo.shard_index
	ShardIndex *int64 `json:"shardIndex,omitempty"`

	// Total number of shards.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.ShardInfo.shard_count
	ShardCount *int64 `json:"shardCount,omitempty"`

	// The index of the first character in
	//  [Document.text][google.cloud.documentai.v1.Document.text] in the overall
	//  document global text.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.ShardInfo.text_offset
	TextOffset *int64 `json:"textOffset,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Document.Style
type Document_Style struct {
	// Text anchor indexing into the
	//  [Document.text][google.cloud.documentai.v1.Document.text].
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Style.text_anchor
	TextAnchor *Document_TextAnchor `json:"textAnchor,omitempty"`

	// Text color.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Style.color
	Color *Color `json:"color,omitempty"`

	// Text background color.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Style.background_color
	BackgroundColor *Color `json:"backgroundColor,omitempty"`

	// [Font weight](https://www.w3schools.com/cssref/pr_font_weight.asp).
	//  Possible values are `normal`, `bold`, `bolder`, and `lighter`.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Style.font_weight
	FontWeight *string `json:"fontWeight,omitempty"`

	// [Text style](https://www.w3schools.com/cssref/pr_font_font-style.asp).
	//  Possible values are `normal`, `italic`, and `oblique`.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Style.text_style
	TextStyle *string `json:"textStyle,omitempty"`

	// [Text
	//  decoration](https://www.w3schools.com/cssref/pr_text_text-decoration.asp).
	//  Follows CSS standard. <text-decoration-line> <text-decoration-color>
	//  <text-decoration-style>
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Style.text_decoration
	TextDecoration *string `json:"textDecoration,omitempty"`

	// Font size.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Style.font_size
	FontSize *Document_Style_FontSize `json:"fontSize,omitempty"`

	// Font family such as `Arial`, `Times New Roman`.
	//  https://www.w3schools.com/cssref/pr_font_font-family.asp
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Style.font_family
	FontFamily *string `json:"fontFamily,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Document.Style.FontSize
type Document_Style_FontSize struct {
	// Font size for the text.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Style.FontSize.size
	Size *float32 `json:"size,omitempty"`

	// Unit for the font size. Follows CSS naming (such as `in`, `px`, and
	//  `pt`).
	// +kcc:proto:field=google.cloud.documentai.v1.Document.Style.FontSize.unit
	Unit *string `json:"unit,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Document.TextAnchor
type Document_TextAnchor struct {
	// The text segments from the
	//  [Document.text][google.cloud.documentai.v1.Document.text].
	// +kcc:proto:field=google.cloud.documentai.v1.Document.TextAnchor.text_segments
	TextSegments []Document_TextAnchor_TextSegment `json:"textSegments,omitempty"`

	// Contains the content of the text span so that users do
	//  not have to look it up in the text_segments.  It is always
	//  populated for formFields.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.TextAnchor.content
	Content *string `json:"content,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Document.TextAnchor.TextSegment
type Document_TextAnchor_TextSegment struct {
	// [TextSegment][google.cloud.documentai.v1.Document.TextAnchor.TextSegment]
	//  start UTF-8 char index in the
	//  [Document.text][google.cloud.documentai.v1.Document.text].
	// +kcc:proto:field=google.cloud.documentai.v1.Document.TextAnchor.TextSegment.start_index
	StartIndex *int64 `json:"startIndex,omitempty"`

	// [TextSegment][google.cloud.documentai.v1.Document.TextAnchor.TextSegment]
	//  half open end UTF-8 char index in the
	//  [Document.text][google.cloud.documentai.v1.Document.text].
	// +kcc:proto:field=google.cloud.documentai.v1.Document.TextAnchor.TextSegment.end_index
	EndIndex *int64 `json:"endIndex,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Document.TextChange
type Document_TextChange struct {
	// Provenance of the correction.
	//  Text anchor indexing into the
	//  [Document.text][google.cloud.documentai.v1.Document.text].  There can
	//  only be a single `TextAnchor.text_segments` element.  If the start and
	//  end index of the text segment are the same, the text change is inserted
	//  before that index.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.TextChange.text_anchor
	TextAnchor *Document_TextAnchor `json:"textAnchor,omitempty"`

	// The text that replaces the text identified in the `text_anchor`.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.TextChange.changed_text
	ChangedText *string `json:"changedText,omitempty"`

	// The history of this annotation.
	// +kcc:proto:field=google.cloud.documentai.v1.Document.TextChange.provenance
	Provenance []Document_Provenance `json:"provenance,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.NormalizedVertex
type NormalizedVertex struct {
	// X coordinate.
	// +kcc:proto:field=google.cloud.documentai.v1.NormalizedVertex.x
	X *float32 `json:"x,omitempty"`

	// Y coordinate (starts from the top of the image).
	// +kcc:proto:field=google.cloud.documentai.v1.NormalizedVertex.y
	Y *float32 `json:"y,omitempty"`
}

// +kcc:proto=google.cloud.documentai.v1.Vertex
type Vertex struct {
	// X coordinate.
	// +kcc:proto:field=google.cloud.documentai.v1.Vertex.x
	X *int32 `json:"x,omitempty"`

	// Y coordinate (starts from the top of the image).
	// +kcc:proto:field=google.cloud.documentai.v1.Vertex.y
	Y *int32 `json:"y,omitempty"`
}

// +kcc:proto=google.protobuf.Any
type Any struct {
	// A URL/resource name that uniquely identifies the type of the serialized
	//  protocol buffer message. This string must contain at least
	//  one "/" character. The last segment of the URL's path must represent
	//  the fully qualified name of the type (as in
	//  `path/google.protobuf.Duration`). The name should be in a canonical form
	//  (e.g., leading "." is not accepted).
	//
	//  In practice, teams usually precompile into the binary all types that they
	//  expect it to use in the context of Any. However, for URLs which use the
	//  scheme `http`, `https`, or no scheme, one can optionally set up a type
	//  server that maps type URLs to message definitions as follows:
	//
	//  * If no scheme is provided, `https` is assumed.
	//  * An HTTP GET on the URL must yield a [google.protobuf.Type][]
	//    value in binary format, or produce an error.
	//  * Applications are allowed to cache lookup results based on the
	//    URL, or have them precompiled into a binary to avoid any
	//    lookup. Therefore, binary compatibility needs to be preserved
	//    on changes to types. (Use versioned type names to manage
	//    breaking changes.)
	//
	//  Note: this functionality is not currently available in the official
	//  protobuf release, and it is not used for type URLs beginning with
	//  type.googleapis.com.
	//
	//  Schemes other than `http`, `https` (or the empty scheme) might be
	//  used with implementation specific semantics.
	// +kcc:proto:field=google.protobuf.Any.type_url
	TypeURL *string `json:"typeURL,omitempty"`

	// Must be a valid serialized protocol buffer of the above specified type.
	// +kcc:proto:field=google.protobuf.Any.value
	Value []byte `json:"value,omitempty"`
}

// +kcc:proto=google.protobuf.FloatValue
type FloatValue struct {
	// The float value.
	// +kcc:proto:field=google.protobuf.FloatValue.value
	Value *float32 `json:"value,omitempty"`
}

// +kcc:proto=google.rpc.Status
type Status struct {
	// The status code, which should be an enum value of
	//  [google.rpc.Code][google.rpc.Code].
	// +kcc:proto:field=google.rpc.Status.code
	Code *int32 `json:"code,omitempty"`

	// A developer-facing error message, which should be in English. Any
	//  user-facing error message should be localized and sent in the
	//  [google.rpc.Status.details][google.rpc.Status.details] field, or localized
	//  by the client.
	// +kcc:proto:field=google.rpc.Status.message
	Message *string `json:"message,omitempty"`

	// A list of messages that carry the error details.  There is a common set of
	//  message types for APIs to use.
	// +kcc:proto:field=google.rpc.Status.details
	Details []Any `json:"details,omitempty"`
}

// +kcc:proto=google.type.Color
type Color struct {
	// The amount of red in the color as a value in the interval [0, 1].
	// +kcc:proto:field=google.type.Color.red
	Red *float32 `json:"red,omitempty"`

	// The amount of green in the color as a value in the interval [0, 1].
	// +kcc:proto:field=google.type.Color.green
	Green *float32 `json:"green,omitempty"`

	// The amount of blue in the color as a value in the interval [0, 1].
	// +kcc:proto:field=google.type.Color.blue
	Blue *float32 `json:"blue,omitempty"`

	// The fraction of this color that should be applied to the pixel. That is,
	//  the final pixel color is defined by the equation:
	//
	//    `pixel color = alpha * (this color) + (1.0 - alpha) * (background color)`
	//
	//  This means that a value of 1.0 corresponds to a solid color, whereas
	//  a value of 0.0 corresponds to a completely transparent color. This
	//  uses a wrapper message rather than a simple float scalar so that it is
	//  possible to distinguish between a default value and the value being unset.
	//  If omitted, this color object is rendered as a solid color
	//  (as if the alpha value had been explicitly given a value of 1.0).
	// +kcc:proto:field=google.type.Color.alpha
	Alpha *FloatValue `json:"alpha,omitempty"`
}

// +kcc:proto=google.type.Date
type Date struct {
	// Year of the date. Must be from 1 to 9999, or 0 to specify a date without
	//  a year.
	// +kcc:proto:field=google.type.Date.year
	Year *int32 `json:"year,omitempty"`

	// Month of a year. Must be from 1 to 12, or 0 to specify a year without a
	//  month and day.
	// +kcc:proto:field=google.type.Date.month
	Month *int32 `json:"month,omitempty"`

	// Day of a month. Must be from 1 to 31 and valid for the year and month, or 0
	//  to specify a year by itself or a year and month where the day isn't
	//  significant.
	// +kcc:proto:field=google.type.Date.day
	Day *int32 `json:"day,omitempty"`
}

// +kcc:proto=google.type.DateTime
type DateTime struct {
	// Optional. Year of date. Must be from 1 to 9999, or 0 if specifying a
	//  datetime without a year.
	// +kcc:proto:field=google.type.DateTime.year
	Year *int32 `json:"year,omitempty"`

	// Required. Month of year. Must be from 1 to 12.
	// +kcc:proto:field=google.type.DateTime.month
	Month *int32 `json:"month,omitempty"`

	// Required. Day of month. Must be from 1 to 31 and valid for the year and
	//  month.
	// +kcc:proto:field=google.type.DateTime.day
	Day *int32 `json:"day,omitempty"`

	// Required. Hours of day in 24 hour format. Should be from 0 to 23. An API
	//  may choose to allow the value "24:00:00" for scenarios like business
	//  closing time.
	// +kcc:proto:field=google.type.DateTime.hours
	Hours *int32 `json:"hours,omitempty"`

	// Required. Minutes of hour of day. Must be from 0 to 59.
	// +kcc:proto:field=google.type.DateTime.minutes
	Minutes *int32 `json:"minutes,omitempty"`

	// Required. Seconds of minutes of the time. Must normally be from 0 to 59. An
	//  API may allow the value 60 if it allows leap-seconds.
	// +kcc:proto:field=google.type.DateTime.seconds
	Seconds *int32 `json:"seconds,omitempty"`

	// Required. Fractions of seconds in nanoseconds. Must be from 0 to
	//  999,999,999.
	// +kcc:proto:field=google.type.DateTime.nanos
	Nanos *int32 `json:"nanos,omitempty"`

	// UTC offset. Must be whole seconds, between -18 hours and +18 hours.
	//  For example, a UTC offset of -4:00 would be represented as
	//  { seconds: -14400 }.
	// +kcc:proto:field=google.type.DateTime.utc_offset
	UtcOffset *string `json:"utcOffset,omitempty"`

	// Time zone.
	// +kcc:proto:field=google.type.DateTime.time_zone
	TimeZone *TimeZone `json:"timeZone,omitempty"`
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

// +kcc:proto=google.type.PostalAddress
type PostalAddress struct {
	// The schema revision of the `PostalAddress`. This must be set to 0, which is
	//  the latest revision.
	//
	//  All new revisions **must** be backward compatible with old revisions.
	// +kcc:proto:field=google.type.PostalAddress.revision
	Revision *int32 `json:"revision,omitempty"`

	// Required. CLDR region code of the country/region of the address. This
	//  is never inferred and it is up to the user to ensure the value is
	//  correct. See http://cldr.unicode.org/ and
	//  http://www.unicode.org/cldr/charts/30/supplemental/territory_information.html
	//  for details. Example: "CH" for Switzerland.
	// +kcc:proto:field=google.type.PostalAddress.region_code
	RegionCode *string `json:"regionCode,omitempty"`

	// Optional. BCP-47 language code of the contents of this address (if
	//  known). This is often the UI language of the input form or is expected
	//  to match one of the languages used in the address' country/region, or their
	//  transliterated equivalents.
	//  This can affect formatting in certain countries, but is not critical
	//  to the correctness of the data and will never affect any validation or
	//  other non-formatting related operations.
	//
	//  If this value is not known, it should be omitted (rather than specifying a
	//  possibly incorrect default).
	//
	//  Examples: "zh-Hant", "ja", "ja-Latn", "en".
	// +kcc:proto:field=google.type.PostalAddress.language_code
	LanguageCode *string `json:"languageCode,omitempty"`

	// Optional. Postal code of the address. Not all countries use or require
	//  postal codes to be present, but where they are used, they may trigger
	//  additional validation with other parts of the address (e.g. state/zip
	//  validation in the U.S.A.).
	// +kcc:proto:field=google.type.PostalAddress.postal_code
	PostalCode *string `json:"postalCode,omitempty"`

	// Optional. Additional, country-specific, sorting code. This is not used
	//  in most regions. Where it is used, the value is either a string like
	//  "CEDEX", optionally followed by a number (e.g. "CEDEX 7"), or just a number
	//  alone, representing the "sector code" (Jamaica), "delivery area indicator"
	//  (Malawi) or "post office indicator" (e.g. Côte d'Ivoire).
	// +kcc:proto:field=google.type.PostalAddress.sorting_code
	SortingCode *string `json:"sortingCode,omitempty"`

	// Optional. Highest administrative subdivision which is used for postal
	//  addresses of a country or region.
	//  For example, this can be a state, a province, an oblast, or a prefecture.
	//  Specifically, for Spain this is the province and not the autonomous
	//  community (e.g. "Barcelona" and not "Catalonia").
	//  Many countries don't use an administrative area in postal addresses. E.g.
	//  in Switzerland this should be left unpopulated.
	// +kcc:proto:field=google.type.PostalAddress.administrative_area
	AdministrativeArea *string `json:"administrativeArea,omitempty"`

	// Optional. Generally refers to the city/town portion of the address.
	//  Examples: US city, IT comune, UK post town.
	//  In regions of the world where localities are not well defined or do not fit
	//  into this structure well, leave locality empty and use address_lines.
	// +kcc:proto:field=google.type.PostalAddress.locality
	Locality *string `json:"locality,omitempty"`

	// Optional. Sublocality of the address.
	//  For example, this can be neighborhoods, boroughs, districts.
	// +kcc:proto:field=google.type.PostalAddress.sublocality
	Sublocality *string `json:"sublocality,omitempty"`

	// Unstructured address lines describing the lower levels of an address.
	//
	//  Because values in address_lines do not have type information and may
	//  sometimes contain multiple values in a single field (e.g.
	//  "Austin, TX"), it is important that the line order is clear. The order of
	//  address lines should be "envelope order" for the country/region of the
	//  address. In places where this can vary (e.g. Japan), address_language is
	//  used to make it explicit (e.g. "ja" for large-to-small ordering and
	//  "ja-Latn" or "en" for small-to-large). This way, the most specific line of
	//  an address can be selected based on the language.
	//
	//  The minimum permitted structural representation of an address consists
	//  of a region_code with all remaining information placed in the
	//  address_lines. It would be possible to format such an address very
	//  approximately without geocoding, but no semantic reasoning could be
	//  made about any of the address components until it was at least
	//  partially resolved.
	//
	//  Creating an address only containing a region_code and address_lines, and
	//  then geocoding is the recommended way to handle completely unstructured
	//  addresses (as opposed to guessing which parts of the address should be
	//  localities or administrative areas).
	// +kcc:proto:field=google.type.PostalAddress.address_lines
	AddressLines []string `json:"addressLines,omitempty"`

	// Optional. The recipient at the address.
	//  This field may, under certain circumstances, contain multiline information.
	//  For example, it might contain "care of" information.
	// +kcc:proto:field=google.type.PostalAddress.recipients
	Recipients []string `json:"recipients,omitempty"`

	// Optional. The name of the organization at the address.
	// +kcc:proto:field=google.type.PostalAddress.organization
	Organization *string `json:"organization,omitempty"`
}

// +kcc:proto=google.type.TimeZone
type TimeZone struct {
	// IANA Time Zone Database time zone, e.g. "America/New_York".
	// +kcc:proto:field=google.type.TimeZone.id
	ID *string `json:"id,omitempty"`

	// Optional. IANA Time Zone Database version number, e.g. "2019a".
	// +kcc:proto:field=google.type.TimeZone.version
	Version *string `json:"version,omitempty"`
}

// +kcc:proto=google.cloud.contentwarehouse.v1.Document
type DocumentObservedState struct {
	// Output only. The time when the document is last updated.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.Document.update_time
	UpdateTime *string `json:"updateTime,omitempty"`

	// Output only. The time when the document is created.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.Document.create_time
	CreateTime *string `json:"createTime,omitempty"`

	// Output only. If linked to a Collection with RetentionPolicy, the date when
	//  the document becomes mutable.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.Document.disposition_time
	DispositionTime *string `json:"dispositionTime,omitempty"`

	// Output only. Indicates if the document has a legal hold on it.
	// +kcc:proto:field=google.cloud.contentwarehouse.v1.Document.legal_hold
	LegalHold *bool `json:"legalHold,omitempty"`
}
