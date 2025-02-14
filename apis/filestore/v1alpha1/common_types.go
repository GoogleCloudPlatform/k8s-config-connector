package v1alpha1

// +kcc:proto=google.cloud.filestore.v1.FileShareConfig
type FileShareConfig struct {
	// Required. The name of the file share. Must use 1-16 characters for the
	//  basic service tier and 1-63 characters for all other service tiers.
	//  Must use lowercase letters, numbers, or underscores `[a-z0-9_]`. Must
	//  start with a letter. Immutable.
	// +kcc:proto:field=google.cloud.filestore.v1.FileShareConfig.name
	Name *string `json:"name,omitempty"`

	// File share capacity in gigabytes (GB).
	//  Filestore defines 1 GB as 1024^3 bytes.
	// +kcc:proto:field=google.cloud.filestore.v1.FileShareConfig.capacity_gb
	CapacityGB *int64 `json:"capacityGB,omitempty"`

	// NOTYET: Is a Ref
	// // The resource name of the backup, in the format
	// //  `projects/{project_number}/locations/{location_id}/backups/{backup_id}`,
	// //  that this file share has been restored from.
	// // +kcc:proto:field=google.cloud.filestore.v1.FileShareConfig.source_backup
	// SourceBackupRef *string `json:"sourceBackup,omitempty"`

	// Nfs Export Options.
	//  There is a limit of 10 export options per file share.
	// +kcc:proto:field=google.cloud.filestore.v1.FileShareConfig.nfs_export_options
	NfsExportOptions []NfsExportOptions `json:"nfsExportOptions,omitempty"`
}
