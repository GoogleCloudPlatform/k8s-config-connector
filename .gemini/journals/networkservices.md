# Network Services Journal

### [2026-06-19] Support for NetworkServicesGateway and MockGCP updates
- **Context**: Migrating `NetworkServicesGateway` to a direct controller, and recording E2E mock test fixtures.
- **Problem**: 
  1. The E2E mock test was originally skipped because it depended on `NetworkSecurityServerTLSPolicy`, which is not implemented in MockGCP.
  2. MockGCP's `UpdateGateway` did not support updating fields other than `labels` and `description`, causing an error `update_mask path "ports" not valid` when the direct controller attempted to update the ports during the E2E test.
- **Solution**: 
  1. Removed `serverTlsPolicyRef` from the `-direct` test fixtures (`create.yaml` and `update.yaml`) since it is an optional field, and deleted `dependencies.yaml`. This allows the direct test to execute fully against MockGCP instead of being skipped.
  2. Added `ports` and `server_tls_policy` support to `UpdateGateway` in `mockgcp/mocknetworkservices/gateway.go`.
- **Impact**: Enables 100% test coverage and validation of the direct `NetworkServicesGateway` controller under the hermetic MockGCP test runner, ensuring future updates do not break it.
