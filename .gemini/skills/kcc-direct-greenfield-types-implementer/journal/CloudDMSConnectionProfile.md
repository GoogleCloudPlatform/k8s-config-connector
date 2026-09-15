# CloudDMSConnectionProfile Greenfield Types Implementation Journal

## Observations & Design Choices

1. **Reachable/Unreachable Types Bootstrapping**:
   - In the initial code generation run, `controllerbuilder` automatically commented out several nested structs in `types.generated.go` as "unreachable".
   - When referencing these types in `CloudDMSConnectionProfileObservedState`, we discovered that `OracleConnectionProfile` and `AlloyDbConnectionProfile` have no output-only fields in the GCP DMS proto definition.
   - Consequently, `OracleConnectionProfileObservedState` and `AlloyDbConnectionProfileObservedState` types do not exist. Including them in `CloudDMSConnectionProfileObservedState` caused `controller-gen` and `generate-crds` compile failures.
   - Removing them from the manual types definition cleanly solved the bootstrapping issue and allowed `types.generated.go` to be generated and compiled successfully.

2. **Reference Implementation**:
   - Mapped `alloydb.cluster_id` to `alloydb.clusterRef` using the standard `refsv1beta1.AlloyDBClusterTypeRef` from `apis/refs/v1beta1`.
   - Mapped `oracle.private_connectivity.private_connection` to `oracle.private_connectivity.privateConnectionRef` utilizing the already-supported `PrivateConnectionRef` in the `clouddms` package.
   - For `mysql.cloud_sql_id` and `postgresql.cloud_sql_id`, we kept them as `CloudSQLID *string` since they can represent external instances and are mapped automatically.
