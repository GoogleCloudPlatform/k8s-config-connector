# Journal: Match MockGCP with Real GCP for SQLSSLCert

During the mock implementation and matching process for `sql.sslCerts` (`SQLSSLCert`), several key insights and practices were identified:

## Mock Service Implementation
1. **gRPC and HTTP Handlers**: The proto definitions and gRPC-gateway bindings (`SqlSslCertsServiceServer`, `RegisterSqlSslCertsServiceServer`, `RegisterSqlSslCertsServiceHandler`) were already generated in `mockgcp/generated/mockgcp/cloud/sql/v1beta4/`, so only the mock handler implementation in `mockgcp/mocksql/sqlsslcert.go` and registration in `service.go` were required.
2. **Deterministic Certificate Matching**: On creation (`Insert`), Cloud SQL generates client certificates, private keys, serial numbers, and SHA-1 fingerprints. By populating mock certificates deterministically in `mocksql/sqlsslcert.go`, the golden object and HTTP logs match real GCP behavior identically without introducing non-deterministic drift.
3. **LRO Target Linking**: Cloud SQL operations for `SslCert` have target links pointing to the parent instance (`projects/{project}/instances/{instance}`). We handled `*pb.SslCert` in `startLRO` in `mocksql/operations_store.go`.

## Normalization in MockGCP (`normalize.go`)
1. **Nested Operation Fields**: In `SslCertsInsertResponse`, the operation object is nested (`.operation.user`, `.operation.insertTime`). Real GCP populates the service account email and operation insert time, which are normalized in `mocksql/normalize.go`.
2. **Nested Certificate Timestamps**: The response also nests certificate create and expiration timestamps (`.clientCert.certInfo.createTime`, `.clientCert.certInfo.expirationTime`), which are stabilized in `mocksql/normalize.go`.
