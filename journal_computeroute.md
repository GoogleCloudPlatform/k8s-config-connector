# Journal: ComputeRoute Diff Discrepancies Migration

## [2026-10-01] Replication of network and next_hop_gateway diff discrepancies
- **Context**: Resolving issue #13581: ComputeRoute: Fix network and next_hop_gateway diff discrepancies between tf and direct controllers.
- **Problem**:
  - In `TestMigrationToDirect`, during Phase 3 (Direct controller takeover) and Phase 4 (Direct controller re-reconciliation), `ComputeRoute` reports diffs on `network` and `next_hop_gateway`.
  - For `network`: `actual` from GCP is `https://www.googleapis.com/compute/v1/projects/${projectId}/global/networks/${networkID}` (or `beta`), whereas `desired` from KRM after reference normalization is `projects/${projectId}/global/networks/${networkID}`.
  - For `next_hop_gateway`: `actual` from GCP is `https://www.googleapis.com/compute/v1/projects/${projectId}/global/gateways/default-internet-gateway`, whereas `desired` in KRM is `default-internet-gateway` (or `global/gateways/default-internet-gateway` / `projects/...`).
  - Because `ComputeRoute` is immutable in GCP, when direct controller finds any diff in `RouteAdapter.Update()`, it attempts an update and fails with `"ComputeRoute is immutable and cannot be updated"`.
- **Replication Details**:
  - Ran `TestMigrationToDirect/fixtures/computeroute$` and verified `network` diff in `_migration_diffs.json`.
  - Created test fixture `computeroutegateway` covering `nextHopGateway` and verified both `network` and `next_hop_gateway` diffs in `_migration_diffs.json`.
- **Fix Plan**:
  - Update `compareComputeRoute` in `pkg/controller/direct/compute/computeroute_controller.go`:
    1. Canonicalize `Network` URL by trimming URI prefix and expanding to canonical `projects/{project}/global/networks/{network}`.
    2. Canonicalize `NextHopGateway` by trimming URI prefix and canonicalizing short name / relative path to `projects/{project}/global/gateways/{gateway}`.
    3. Canonicalize other reference / URL fields (`NextHopIlb`, `NextHopInstance`, `NextHopVpnTunnel`) by trimming URI prefix.
    4. Pass route identity `id *krm.ComputeRouteIdentity` to `compareComputeRoute` to access project ID for expansion.
  - Re-run `TestMigrationToDirect` on all fixtures and verify:
    - No migration diff in `_migration_diffs.json` (no entries with `"isNewObject": false`).
    - Direct controller cleanly reports no diffs and succeeds without attempting invalid updates.
