# PR #2: Implement MockGCP Server for NetworkServicesWasmPlugin

### BRIEF Change description
This PR implements the MockGCP server for `NetworkServicesWasmPlugin`, enabling hermetic testing of the resource. It follows the three-PR resource addition workflow (PR #2 of 3).

### Key Changes
- **Mock Server:** Implemented CRUD operations for `WasmPlugin` and `WasmPluginVersion` in `mockgcp/mocknetworkservices/wasmplugin.go`.
- **Fidelity Alignment:** The mock correctly handles the "inlined" version creation pattern seen in the Service Extensions API while matching real GCP's behavior where versions are managed as separate child resources (clearing the `versions` map in GET responses).
- **Normalization:** Added rules in `mockgcp/mocknetworkservices/normalize.go` to stabilize volatile fields like timestamps and LRO metadata (`requestedCancellation`).
- **Test Coverage:** Added a `gcloud` based CRUD script in `mockgcp/mocknetworkservices/testdata/wasmplugin/crud/` verified against real GCP and the mock server.

### Verification Results
- **Record GCP:** Successfully captured real GCP behavior using the sample image `us-docker.pkg.dev/service-extensions-samples/plugins/local-reply:main`.
- **Compare Mock:** Verified that the mock implementation produces logs that align with real GCP (allowing for standard normalization).
- **Test Status:** `compare-mock mocknetworkservices/testdata/wasmplugin/crud` passes.

### Checklist
- [x] Generated mock server code is complete.
- [x] The mock correctly implements basic CRUD operations.
- [x] Golden logs recorded and aligned with real GCP.
- [x] Run `make ready-pr` (pending final step in PR workflow).
