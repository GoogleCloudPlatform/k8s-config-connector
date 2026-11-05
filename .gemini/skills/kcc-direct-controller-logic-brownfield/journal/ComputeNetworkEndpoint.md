# ComputeNetworkEndpoint Migration Journal

## Overview
Migrated `ComputeNetworkEndpoint` from the legacy Terraform-based reconciler to the `direct` reconciler architecture.

## Key Observations and Decisions
1. **Subresource / Attachment Nature:**
   - A `ComputeNetworkEndpoint` is not an independent top-level GCP REST resource, but rather an endpoint attached to a zonal `NetworkEndpointGroup` (NEG).
   - Existence is verified by querying `NetworkEndpointGroups.ListNetworkEndpoints` and matching on `ipAddress`, `port`, and `instance` name.
   - Creation uses `NetworkEndpointGroups.AttachNetworkEndpoints` and deletion uses `NetworkEndpointGroups.DetachNetworkEndpoints`.
   - The resource is immutable in GCP (no Update endpoint exists; any change requires detach and re-attach). Thus, the `Update` adapter method checks for diffs and surfaces an immutability error if spec changes are detected.

2. **Re-Reconciliation Read-Only Method:**
   - In GCP Compute Engine, listing endpoints (`NetworkEndpointGroups.ListNetworkEndpoints`) uses HTTP method `POST` (`/listNetworkEndpoints`) rather than `GET`.
   - The test framework's re-reconciliation check (`unified_test.go` and `migration_test.go`) previously assumed only `GET` requests were read-only for HTTP. It was updated to also recognize `POST .../listNetworkEndpoints` as a read-only query during steady-state re-reconciliation.

3. **Instance Formatting & Normalization:**
   - The GAPIC client and Terraform provider send the short instance name (e.g., `my-instance`) rather than the full instance URL (`projects/.../zones/.../instances/my-instance`).
   - The adapter extracts the short instance name from the resolved identity to match both live GCP expectations and existing golden logs.

4. **Testing & Golden Logs:**
   - Successfully recorded authentic GCP baseline logs with `hack/record-gcp` against live GCP for both base reconciliation and migration takeover (`TestMigrationToDirect`).
   - Aligned mock behavior and verified golden snapshots with `hack/compare-mock` and `TestGoldenLogAlignment`.
