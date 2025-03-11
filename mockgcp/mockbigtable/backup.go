in.proto.service: google.bigtable.admin.v2.BigtableTableAdmin
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
    deprecated: true;
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
  rpc UpdateTable(UpdateTableRequest) returns (google.longrunning.Operation) {
    option (google.api.http) = {
      patch: "/v2/{table.name=projects/*/instances/*/tables/*}"
      body: "table"
    };
    option (google.api.method_signature) = "table,update_mask";
    option (google.longrunning.operation_info) = {
      response_type: "Table"
      metadata_type: "UpdateTableMetadata"
    };
  }

  // Deletes a specified table and all of its data.
  rpc DeleteTable(DeleteTableRequest) returns (google.protobuf.Empty) {
    option (google.api.http) = {
      delete: "/v2/{name=projects/*/instances/*/tables/*}"
    };
    option (google.api.method_signature) = "name";
  }

  // Performs a series of column family modifications on the specified table.
  // Either all or none of the

