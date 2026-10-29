# Migration Journal: ContainerNodePool

## Observations

### Test Fixtures
- Resource: `ContainerNodePool` (`container.cnrm.cloud.google.com/v1beta1`)
- Test fixtures evaluated under `pkg/test/resourcefixture/testdata/basic/container/v1beta1/containernodepool`:
  1. `containernodepool`
  2. `containernodepool-advanced-machine-features`
  3. `containernodepool-containerdconfig`
  4. `containernodepool-datacache`
  5. `containernodepool-enableconfidentialnodes`
  6. `containernodepool-kubeletconfig`
  7. `containernodepool-queuedprovisioning`
  8. `containernodepool-resourcemanagertags`
  9. `containernodepool-subnetworkref`
  10. `containernodepool-swap`
  11. `containernodepool-swap-dedicated-ssd`
  12. `containernodepool-swap-ephemeral-ssd`
  13. `containernodepool-swap-ephemeral-ssd-gib`
  14. `containernodepool-swap-gib`
  15. `containernodepool-windowsnodeconfig`
  16. `containernodepoolremovetaint`

None of these fixtures specify `alpha.cnrm.cloud.google.com/reconciler: direct` in `create.yaml` (they all use the legacy Terraform controller initially).

### MockGCP Migration Validation
- Executed `TestMigrationToDirect` against MockGCP for all 16 `containernodepool` fixtures.
- 15 out of 16 fixtures passed cleanly on MockGCP with 0-write Phase 3 direct takeover and 0-write Phase 4 direct re-reconciliation.
- Fixture `containernodepool-containerdconfig` failed on MockGCP because MockGCP's generated proto definition for `ContainerdConfig` is missing `WritableCgroups` and `RegistryHosts`, causing MockGCP to drop those fields and triggering an update `PUT` request during migration takeover. This is a known MockGCP proto limitation documented in `tests/e2e/ratcheting.go`.

### Real GCP Execution
- Executed `./hack/record-gcp` against live GCP (project `cnrm-barni-4`) for:
  - `./hack/record-gcp "TestMigrationToDirect/fixtures/containernodepool$"`
  - `./hack/record-gcp "TestMigrationToDirect/fixtures/containernodepool-containerdconfig$"` (ensured secret `kcc-test-ca-cert` existed in project `cnrm-barni-4` prior to test run)
- In both fixtures, all four phases executed successfully:
  1. **Phase 1 (Legacy Create):** Successfully provisioned cluster dependencies and the `ContainerNodePool` via the legacy Terraform controller.
  2. **Phase 2 (Legacy Re-reconcile):** Re-reconciliation with the legacy controller succeeded.
  3. **Phase 3 (Direct Takeover):** The Direct controller adopted the resource cleanly with 0 writes (only `GET` requests sent to the Container API: `https://container.googleapis.com/v1/projects/${projectId}/locations/us-central1-a/clusters/.../nodePools/...`).
  4. **Phase 4 (Direct Re-reconcile):** The Direct controller performed re-reconciliation cleanly with 0 writes (only `GET` requests).
- Real GCP execution confirms that real GCP retains all `containerdConfig` fields (`writableCgroups`, `registryHosts`, `privateRegistryAccessConfig`), and the Direct controller performs a clean 0-write takeover on real GCP.

### Takeover Diffs
- Inspection of `_migration_diffs.json` confirmed 0 diffs during Direct takeover and re-reconciliation (no `"isNewObject": false` entries). Full state compatibility is maintained between the legacy Terraform controller and the direct controller.

### Ratcheting Exclusions
- Verified `tests/e2e/ratcheting.go`: `ContainerNodePool` is already active and tested for re-reconciliation in `ShouldTestRereconiliation` (returning `true` for all fixtures, with only `containernodepool-containerdconfig` returning `false` due to MockGCP's proto limitation).
