# StorageControl Journal

### 2026-05-23 Implement StorageManagedFolder direct controller and E2E mock support
- **Context**: Implementing direct controller and MockGCP support for `StorageManagedFolder` (KCC issue #8555).
- **Problem**: MockGCP had JSON/HTTP endpoints for managed folders in the storage service but lacked GRPC/protobuf implementations for `GetManagedFolder`, `CreateManagedFolder`, and `DeleteManagedFolder` in the `StorageControlService`.
- **Solution**: Implemented `GetManagedFolder`, `CreateManagedFolder`, and `DeleteManagedFolder` on `StorageControlService` under `mockgcp/mockstorage/storagecontrol.go` to store and load GRPC-based `pb.ManagedFolder` records from the generic mock GCS storage backend.
- **Impact**: Unblocked direct controller E2E testing of `StorageManagedFolder` and other Storage Control GRPC-based resources against MockGCP, enabling fast hermetic validations.
