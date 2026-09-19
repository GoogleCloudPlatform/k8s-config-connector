### [2026-09-19] Greenfield Implementation for OracleDatabaseExadbVMCluster
- **Context**: Implementing Greenfield direct KRM types, identity, and generate.sh for `OracleDatabaseExadbVMCluster` (under issue #13321).
- **Problem**: The generator produced a mapping for `google.type.TimeZone` field (`TimeZone`), but the corresponding `TimeZone_FromProto` and `TimeZone_ToProto` mapper functions were undefined because `TimeZone` is a common message-based type under `google.golang.org/genproto/googleapis/type/datetime`.
- **Solution**: Created a handwritten mapper file `custom_mappers.go` under `pkg/controller/direct/oracledatabase/` and implemented the mapping functions `TimeZone_FromProto` and `TimeZone_ToProto` using the `google.golang.org/genproto/googleapis/type/datetime` package.
- **Impact**: Unblocks compiling direct controllers that reference common message-based fields (like `google.type.TimeZone`) where standard scalar mapping is insufficient.
