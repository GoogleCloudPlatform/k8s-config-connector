### [2026-05-23] EdgeContainerMachine Direct Controller and MockGCP Support
- **Context**: Implementing the Phase 2 (Controller and E2E Tests) for `EdgeContainerMachine` [Issue #8603]
- **Problem**: `EdgeContainerMachine` represents physical hardware registered to Distributed Cloud Edge, meaning it is a completely read-only resource in the GCP API. It lacks `CreateMachine`, `UpdateMachine`, and `DeleteMachine` methods in the GCP API and client library, which makes standard KCC CRUD reconciliation challenging and raises questions about how to test it hermetically.
- **Solution**:
  1. Implemented a read-only direct controller that returns a clear descriptive error on `Create` (explaining physical machines cannot be created via the API and must be registered out of band) and validates that `Update` spec matches the GCP resource.
  2. Implemented dynamic auto-creation of mock machines on-the-fly inside MockGCP's `GetMachine` and `ListMachines` if they are not already in mock storage. This perfectly simulates pre-existing registered physical machines in the GCP environment, enabling E2E tests to compile and pass hermetically.
- **Impact**: Provides an elegant and robust template for handling other read-only physical hardware or out-of-band registered resources in KCC, allowing full test coverage without needing external test environment setup.
