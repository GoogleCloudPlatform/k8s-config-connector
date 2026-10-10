### [2026-09-19] Greenfield Implementation for OracleDatabaseExadbVMCluster
- **Context**: Implementing Greenfield direct KRM types, identity, and generate.sh for `OracleDatabaseExadbVMCluster` (under issue #13321).
- **Problem**: The generator produced a mapping for `google.type.TimeZone` field (`TimeZone`), but the corresponding `TimeZone_FromProto` and `TimeZone_ToProto` mapper functions were undefined because `TimeZone` is a common message-based type under `google.golang.org/genproto/googleapis/type/datetime`.
- **Solution**: Created a handwritten mapper file `custom_mappers.go` under `pkg/controller/direct/oracledatabase/` and implemented the mapping functions `TimeZone_FromProto` and `TimeZone_ToProto` using the `google.golang.org/genproto/googleapis/type/datetime` package.
- **Impact**: Unblocks compiling direct controllers that reference common message-based fields (like `google.type.TimeZone`) where standard scalar mapping is insufficient.

### [2026-09-23] Greenfield Implementation for OracleDatabaseAutonomousDatabase
- **Context**: Implementing Greenfield direct KRM types, IdentityV2, Reference, and CRD for `OracleDatabaseAutonomousDatabase` (under issue #10293).
- **Problem**: The GCP proto defines `EncryptionKey` which is reused in both spec (input) and output state (under `properties.encryptionKeyHistoryEntries[]` in observed state). If the spec's `EncryptionKey` uses `KmsKeyRef` to reference a KMS CryptoKey, the status fields will inherit this KCC reference type, failing the `TestNoRefsInStatus` validation test.
- **Solution**: Designed and implemented custom `EncryptionKeyObservedState` and `EncryptionKeyHistoryEntryObservedState` types in `_types.go` specifically for status/observedState. These types represent reference fields (like `kmsKey`) as plain string pointers, isolating reference logic from status representation.
- **Impact**: Ensures that status fields strictly contain plain-string values instead of unresolved KCC references, preventing validation and reconciliation feedback loops.

### [2026-10-08] Greenfield Direct Controller for OracleDatabaseAutonomousDatabase
- **Context**: Implementing Greenfield direct controller, E2E fixtures, and fuzzer for `OracleDatabaseAutonomousDatabase` (under issue #13806).
- **Problem**: When executing `hack/record-gcp` against live GCP APIs, `CreateAutonomousDatabase` returns HTTP 400 `ENTITLEMENT_NOT_FOUND` ("The billing account doesn't have an associated marketplace order. Please confirm or complete a purchasing offer before provisioning the resource.") because Oracle Database@Google Cloud requires an active Marketplace order/entitlement linked to the billing account.
- **Solution**: Implemented the direct controller using GAPIC REST client (`oracledatabase.NewRESTClient`) with `RESTClientOptions()`, proper reference normalization, structured diff comparison using `DiffForTopLevelFields`, and status updates from live GCP responses. Created minimal and maximal test fixtures and registered `OracleDatabaseAutonomousDatabase` in `static_config.go`.
- **Impact**: Full direct controller implementation, fuzzer roundtripping (10,000 iterations verified), and golden file generation against live GCP endpoints.


