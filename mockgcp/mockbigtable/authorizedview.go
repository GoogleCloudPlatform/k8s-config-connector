in.proto.service.definition: service BigtableTableAdmin {
  option (google.api.default_host) = "bigtableadmin.googleapis.com";
  option (google.api.oauth_scopes) =
      "https://www.googleapis.com/auth/bigtable.admin,"
      "https://www.googleapis.com/auth/bigtable.admin.table,"
      "https://www.googleapis.com/auth/cloud-bigtable.admin,"
      "https://www.googleapis.com/auth/cloud-bigtable.admin.table,"
      "https://www.googleapis.com/auth/cloud-platform,"
      "https://www.googleapis.com/auth/cloud-platform.read-only";

  // Creates a new table in the specified instance.
  // The table can be created with a full set of initial column families,
  // specified in the request.
  rpc CreateTable(CreateTableRequest) returns (Table) {
    option (google.api.http) = {
      post: "/v2/{parent=projects/*/instances/*}/tables"
      body: "table"
    };
    option (google.api.method_signature) = "parent,table_id,table";
  }

  // Creates a new table from the specified snapshot. The target table must
  // not exist. The snapshot and the table must be in the same instance.
  //
  // Note: This is a private alpha release of Cloud Bigtable snapshots. This
  // feature is not currently available to most Cloud Bigtable customers. This
  // feature might be changed in backward-incompatible ways and is not
  // recommended for production use. It is not subject to any SLA or deprecation
  // policy.
  rpc CreateTableFromSnapshot(CreateTableFromSnapshotRequest)
      returns (google.longrunning.Operation) {
    option (google.api.http) = {
      post: "/v2/{parent=projects/*/instances/*}/tables:createFromSnapshot"
      body: "*"
    };
    option (google.longrunning.operation_info) = {
      response_type: "Table"
      metadata_type: "CreateTableFromSnapshotMetadata"
    };
    option (google.api.method_signature) = "parent,table_id,source_snapshot";
  }

  // Lists all tables served from a specified instance.
  rpc ListTables(ListTablesRequest) returns (ListTablesResponse) {
    option (google.api.http) = {
      get: "/v2/{parent=projects/*/instances/*}/tables"
    };
    option (google.api.method_signature) = "parent";
  }

  // Gets metadata information about the specified table.
  rpc GetTable(GetTableRequest) returns (Table) {
    option (google.api.http) = {
      get: "/v2/{name=projects/*/instances/*/tables/*}"
    };
    option (google.api.method_signature) = "name";
  }

  // Updates a specified table.
  rpc UpdateTable(UpdateTableRequest)
      returns (google.longrunning.Operation) {
    option (google.api.http) = {
      patch: "/v2/{table.name=projects/*/instances/*/tables/*}"
      body: "table"
    };
    option (google.longrunning.operation_info) = {
      response_type: "Table"
      metadata_type: "UpdateTableMetadata"
    };
  }

  // Permanently deletes a specified table and all of its data.
  rpc DeleteTable(DeleteTableRequest) returns (google.protobuf.Empty) {
    option (google.api.http) = {
      delete: "/v2/{name=projects/*/instances/*/tables/*}"
    };
    option (google.api.method_signature) = "name";
  }

  // Performs a series of column family modifications on the specified table.
  // Either all or none of the modifications will occur before this method
  // returns, but data requests received prior to that point may see a table
  // where only some modifications have taken effect.
  rpc ModifyColumnFamilies(ModifyColumnFamiliesRequest) returns (Table) {
    option (google.api.http) = {
      post: "/v2/{name=projects/*/instances/*/tables/*}:modifyColumnFamilies"
      body: "*"
    };
    option (google.api.method_signature) = "name,modifications";
  }

  // Permanently drop/delete a row range from a specified table. The request can
  // specify whether to delete all rows in a table, or only those that match a
  // particular prefix.
  rpc DropRowRange(DropRowRangeRequest) returns (google.protobuf.Empty) {
    option (google.api.http) = {
      post: "/v2/{name=projects/*/instances/*/tables/*}:dropRowRange"
      body: "*"
    };
    option (google.api.method_signature) = "name";
  }

  // Generates a series of tokens to be used for API responded that are wrapped
  // in stream.
  rpc GenerateConsistencyToken(GenerateConsistencyTokenRequest)
      returns (GenerateConsistencyTokenResponse) {
    option (google.api.http) = {
      post: "/v2/{name=projects/*/instances/*/tables/*}:generateConsistencyToken"
      body: "*"
    };
    option (google.api.method_signature) = "name";
  }

  // Checks replication consistency based on a consistency token, that is, if
  // replication has caught up based on the conditions specified in the token
  // and the check request.
  rpc CheckConsistency(CheckConsistencyRequest)
      returns (CheckConsistencyResponse) {
    option (google.api.http) = {
      post: "/v2/{name=projects/*/instances/*/tables/*}:checkConsistency"
      body: "*"
    };
    option (google.api.method_signature) = "name,consistency_token";
  }

  // Creates a new snapshot in the specified cluster from the specified
  // source table. The cluster and the table must be in the same instance.
  //
  // Note: This is a private alpha release of Cloud Bigtable snapshots. This
  // feature is not currently available to most Cloud Bigtable customers. This
  // feature might be changed in backward-incompatible ways and is not
  // recommended for production use. It is not subject to any SLA or deprecation
  // policy.
  rpc SnapshotTable(SnapshotTableRequest)
      returns (google.longrunning.Operation) {
    option (google.api.http) = {
      post: "/v2/{name=projects/*/instances/*/tables/*}:snapshot"
      body: "*"
    };
    option (google.longrunning.operation_info) = {
      response_type: "Snapshot"
      metadata_type: "SnapshotTableMetadata"
    };
    option (google.api.method_signature) = "name,cluster,snapshot_id";
  }

  // Gets metadata information about the specified snapshot.
  //
  // Note: This is a private alpha release of Cloud Bigtable snapshots. This
  // feature is not currently available to most Cloud Bigtable customers. This
  // feature might be changed in backward-incompatible ways and is not
  // recommended for production use. It is not subject to any SLA or deprecation
  // policy.
  rpc GetSnapshot(GetSnapshotRequest) returns (Snapshot) {
    option (google.api.http) = {
      get: "/v2/{name=projects/*/instances/*/clusters/*/snapshots/*}"
    };
    option (google.api.method_signature) = "name";
  }

  // Lists all snapshots in a cluster.
  //
  // Note: This is a private alpha release of Cloud Bigtable snapshots. This
  // feature is not currently available to most Cloud Bigtable customers. This
  // feature might be changed in backward-incompatible ways and is not
  // recommended for production use. It is not subject to any SLA or deprecation
  // policy.
  rpc ListSnapshots(ListSnapshotsRequest) returns (ListSnapshotsResponse) {
    option (google.api.http) = {
      get: "/v2/{parent=projects/*/instances/*/clusters/*}/snapshots"
    };
    option (google.api.method_signature) = "parent";
  }

  // Permanently deletes the specified snapshot.
  //
  // Note: This is a private alpha release of Cloud Bigtable snapshots. This
  // feature is not currently available to most Cloud Bigtable customers. This
  // feature might be changed in backward-incompatible ways and is not
  // recommended for production use. It is not subject to any SLA or deprecation
  // policy.
  rpc DeleteSnapshot(DeleteSnapshotRequest) returns (google.protobuf.Empty) {
    option (google.api.http) = {
      delete: "/v2/{name=projects/*/instances/*/clusters/*/snapshots/*}"
    };
    option (google.api.method_signature) = "name";
  }

  // Starts creating a new Cloud Bigtable Backup. The returned backup
  // [long-running operation][google.longrunning.Operation] can be used to
  // track creation of the backup. The
  // [metadata][google.longrunning.Operation.metadata] field type is
  // [CreateBackupMetadata][google.bigtable.admin.v2.CreateBackupMetadata]. The
  // [response][google.longrunning.Operation.response] type is
  // [Backup][google.bigtable.admin.v2.Backup], if successful. Cancelling the
  // long-running operation will cause the new backup to be deleted.
  rpc CreateBackup(CreateBackupRequest) returns (google.longrunning.Operation) {
    option (google.api.http) = {
      post: "/v2/{parent=projects/*/instances/*/clusters/*}/backups"
      body: "backup"
    };
    option (google.api.method_signature) = "parent,backup,backup_id";
    option (google.longrunning.operation_info) = {
      response_type: "Backup"
      metadata_type: "CreateBackupMetadata"
    };
  }

  // Gets metadata on a pending or completed Cloud Bigtable Backup.
  rpc GetBackup(GetBackupRequest) returns (Backup) {
    option (google.api.http) = {
      get: "/v2/{name=projects/*/instances/*/clusters/*/backups/*}"
    };
    option (google.api.method_signature) = "name";
  }

  // Updates a pending or completed Cloud Bigtable Backup.
  rpc UpdateBackup(UpdateBackupRequest) returns (Backup) {
    option (google.api.http) = {
      patch: "/v2/{backup.name=projects/*/instances/*/clusters/*/backups/*}"
      body: "backup"
    };
    option (google.api.method_signature) = "backup,update_mask";
  }

  // Deletes a pending or completed Cloud Bigtable backup.
  rpc DeleteBackup(DeleteBackupRequest) returns (google.protobuf.Empty) {
    option (google.api.http) = {
      delete: "/v2/{name=projects/*/instances/*/clusters/*/backups/*}"
    };
    option (google.api.method_signature) = "name";
  }

  // Lists Cloud Bigtable backups. Returns both completed and pending
  // backups.
  rpc ListBackups(ListBackupsRequest) returns (ListBackupsResponse) {
    option (google.api.http) = {
      get: "/v2/{parent=projects/*/instances/*/clusters/*}/backups"
    };
    option (google.api.method_signature) = "parent";
  }

  // Create a new AuthorizedView in the specified table.
  rpc CreateAuthorizedView(CreateAuthorizedViewRequest)
      returns (AuthorizedView) {
    option (google.api.http) = {
      post: "/v2/{parent=projects/*/instances/*/tables/*}/authorizedViews"
      body: "authorized_view"
    };
    option (google.api.method_signature) = "parent,authorized_view,view_id";
  }

  // Lists all AuthorizedViews to the specified table.
  rpc ListAuthorizedViews(ListAuthorizedViewsRequest)
      returns (ListAuthorizedViewsResponse) {
    option (google.api.http) = {
      get: "/v2/{parent=projects/*/instances/*/tables/*}/authorizedViews"
    };
    option (google.api.method_signature) = "parent";
  }

  // Gets information about the specified authorized view.
  rpc GetAuthorizedView(GetAuthorizedViewRequest) returns (AuthorizedView) {
    option (google.api.http) = {
      get: "/v2/{name=projects/*/instances/*/tables/*/authorizedViews/*}"
    };
    option (google.api.method_signature) = "name";
  }

  // Updates a specified authorized view.
  rpc UpdateAuthorizedView(UpdateAuthorizedViewRequest)
      returns (google.longrunning.Operation) {
    option (google.api.http) = {
      patch: "/v2/{authorized_view.name=projects/*/instances/*/tables/*/authorizedViews/*}"
      body: "authorized_view"
    };
    option (google.api.method_signature) = "authorized_view,update_mask";
    option (google.longrunning.operation_info) = {
      response_type: "AuthorizedView"
      metadata_type: "UpdateAuthorizedViewMetadata"
    };
  }

  // Permanently deletes the specified authorized view.
  rpc DeleteAuthorizedView(DeleteAuthorizedViewRequest)
      returns (google.protobuf.Empty) {
    option (google.api.http) = {
      delete: "/v2/{name=projects/*/instances/*/tables/*/authorizedViews/*}"
    };
    option (google.api.method_signature) = "name";
  }
}

```


