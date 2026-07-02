### [2026-05-12] NetworkConnectivityRegionalEndpoint lack of mutable fields
- **Context**: Implementing `NetworkConnectivityRegionalEndpoint` using direct controller pattern.
- **Problem**: The GCP API definition for `RegionalEndpoint` specifies no input fields other than the resource ID and parent URI. `Update`/`Patch` operations are completely omitted from the API surface.
- **Solution**: The `Update()` method in the direct controller adapter was simplified to skip any GCP API `Patch` calls since no such method exists and no fields can be updated. It just updates the status observed state.
- **Impact**: Future agents working on networkconnectivity or similar singleton-like resources should be aware that missing `Patch` or mutable fields in the API definition necessitates an `Update()` method that only performs status updates.

### [2026-07-01] NetworkConnectivityRegionalEndpoint MockGCP Verification
- **Context**: Verified Phase 3 (MockGCP and Alignment) for `NetworkConnectivityRegionalEndpoint`.
- **Validation**: Executed the entire E2E fixtures suite under `networkconnectivity.cnrm.cloud.google.com` API group (covering both `networkconnectivityregionalendpoint-minimal` and `networkconnectivityregionalendpoint-maximal`).
- **Result**: Both minimal and maximal E2E tests pass completely against MockGCP with exit code 0. No HTTP or KRM differences were found, validating that the MockGCP implementation (`mockgcp/mocknetworkconnectivity/regionalendpoints.go`) is already fully aligned, complete, and correct.

