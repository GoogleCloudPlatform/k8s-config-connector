# SQLSSLCert Fuzzer Journal

## Observations & Learnings

- **Proto & KRM Mapping Details:**
  - `SQLSSLCert` corresponds to `google.cloud.sql.v1beta4.SslCert`.
  - In KRM, `SQLSSLCertSpec` contains `commonName` (which maps to proto `.common_name`), `instanceRef` (a reference to the Cloud SQL instance, represented as the parent path / instance name in GCP), and `resourceID` (the service-generated SHA1 fingerprint, used for acquisition).
  - In KRM, `SQLSSLCertStatus` contains `cert`, `certSerialNumber`, `createTime`, `expirationTime`, and `sha1Fingerprint` which map 1:1 to `.cert`, `.cert_serial_number`, `.create_time`, `.expiration_time`, and `.sha1_fingerprint` in the proto.
  - `privateKey` and `serverCaCert` in `SQLSSLCertStatus` are returned by Cloud SQL only upon creation inside `SslCertDetail`, and are not part of `SslCert`.
  - Identity / GCP path fields `.instance` and `.self_link` are registered with `f.Unimplemented_Identity()`.
  - Proto metadata field `.kind` is marked with `f.Unimplemented_NotYetTriaged()`.

- **Mappers Implementation:**
  - Implemented `SQLSSLCertSpec_FromProto`, `SQLSSLCertSpec_ToProto`, `SQLSSLCertStatus_FromProto`, and `SQLSSLCertStatus_ToProto` in `pkg/controller/direct/sql/sqlsslcert_mapper.go`.
  - Created the fuzzer in `pkg/controller/direct/sql/sqlsslcert_fuzzer.go` registered with `fuzztesting.RegisterKRMFuzzer`.

- **Verification:**
  - `pkg/controller/direct/sql` is already imported in `pkg/controller/direct/register/register.go`.
  - Verified with `FOCUS=SQLSSLCert go test -count=1 -v ./pkg/fuzztesting/fuzztests/ -run TestFocusedMappers`.
  - Verified with `go test -count=1 -v ./pkg/fuzztesting/fuzztests/ -run TestListTypesWithoutFuzzers`, confirming `SQLSSLCert` is no longer on the missing fuzzers list.
  - Ran package unit tests with `go test -count=1 -v ./pkg/controller/direct/sql/...`.
