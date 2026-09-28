### 2026-09-10 Greenfield Implementation for LustreInstance (Managed Lustre)
- **Context**: Implementing Greenfield Direct Controller, API Types, MockGCP, and Fixtures for `LustreInstance` under `lustre.cnrm.cloud.google.com/v1alpha1`.
- **Problem**: 
  1. **Capacity Increments**: The upstream protobuf documentation for `capacity_gib` indicates an allowed range of 18,000 to 954,000 GiB in increments of 9,000. In reality, the GCP API validates that `capacity_gib` must be at least 36,000 and a multiple of 36,000 GiB (e.g. 36000, 72000, etc., up to 6,120,000 GiB). Specifying 9,000 or 18,000 causes an `INVALID_ARGUMENT` error from GCP.
  2. **VPC Peering (Private Services Access)**: Lustre instances require private services access peering on the VPC network (`servicenetworking.googleapis.com`). Without a pre-existing peering connection (`ComputeGlobalAddress` + `ServiceNetworkingConnection`), instance creation fails.
  3. **Deletion Ordering**: During E2E test teardown, deleting the VPC network before the Lustre instance is fully deleted causes API errors. In `pkg/test/resourcefixture/testdata/basic/lustre/v1alpha1/lustreinstance/lustreinstance-basic/create.yaml`, the fixture must rely on `opt.DeleteInOrder = true` when testing with shared networks.
  4. **Proto Status Fields**: The GCP REST API returns `gstate` instead of `state` in the instance message, which needs careful mapping or fallback in the direct controller adapter.
- **Solution**:
  1. Updated test fixtures and MockGCP validation to enforce 36,000 GiB multiples.
  2. Included Private Services Access dependencies in the E2E basic fixture.
  3. Registered `lustreinstance` with `DeleteInOrder: true` in `tests/e2e/options.go`.
  4. Implemented direct controller with GAPIC REST client (`google.golang.org/api/lustre/v1`) and structured top-level field diffing.
- **Impact**: Provides exact operational knowledge for developers and agents working with Google Cloud Managed Lustre and avoids provisioning failures in real GCP recording.
