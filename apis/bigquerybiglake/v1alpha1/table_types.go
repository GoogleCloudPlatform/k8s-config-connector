package v1alpha1

type BigLakeTableSpec struct {

	// The BigLakeCatalog that this resource belongs to.
	// +required
	CatalogRef *CatalogRef `json:"catalogRef"`

	// Required. The parent resource where this table will be created.
	// Format:
	// projects/{project_id_or_number}/locations/{location_id}/catalogs/{catalog_id}/databases/{database_id}
	DatabaseRef *refs.BigqueryBigLakeDatabaseRef `json:"databaseRef"`

	// The BigLake Table ID. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	// The table type.
	// +kcc:proto:field=google.cloud.bigquery.biglake.v1.Table.type
	// +optional
	Type *string `json:"type,omitempty"`

	// Options of a Hive table.
	// +kcc:proto:field=google.cloud.bigquery.biglake.v1.Table.hive_options
	// +optional
	HiveOptions *HiveTableOptions `json:"hiveOptions,omitempty"`

	// The checksum of a table object computed by the server based on the value of
	//  other fields. It may be sent on update requests to ensure the client has an
	//  up-to-date value before proceeding. It is only checked for update table
	//  operations.
	// +kcc:proto:field=google.cloud.bigquery.biglake.v1.Table.etag
	// +optional
	Etag *string `json:"etag,omitempty"`
}
