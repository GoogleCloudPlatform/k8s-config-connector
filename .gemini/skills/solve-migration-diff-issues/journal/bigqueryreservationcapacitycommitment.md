# Migration Journal: BigQueryReservationCapacityCommitment

## Context & Objectives
Validating direct migration and re-reconciliation for `BigQueryReservationCapacityCommitment` (`bigqueryreservation.cnrm.cloud.google.com/v1alpha1`).

## Live GCP Recording Attempt (Step 1)
Live GCP recording was attempted using the following commands:
```bash
./hack/record-gcp TestMigrationToDirect/fixtures/bigquerycapacitycommitment$
./hack/record-gcp TestMigrationToDirect/fixtures/bigqueryreservationcapacitycommitmentautogen$
```

### Error Encountered
The test run on live GCP failed due to IAM permission restrictions in project `cnrm-barni-4`:
```text
googleapi: Error 403: Caller does not have required permission to use project cnrm-barni-4. Grant the caller the Owner, Editor, or BigQuery Admin role, or a user-defined role with the bigquery.capacityCommitments.create permission to use this service., forbidden
```

## Hermetic MockGCP Validation
Due to the live GCP permission restriction, validation was conducted hermetically against MockGCP (`E2E_GCP_TARGET=mock`):
1. **Re-reconciliation Testing**:
   - Removed `BigQueryReservationCapacityCommitment` from `ShouldTestRereconiliation` in `tests/e2e/ratcheting.go`.
   - Verified that `TestAllInSeries` runs and passes re-reconciliation for both fixtures:
     - `bigquerycapacitycommitment`
     - `bigqueryreservationcapacitycommitmentautogen`
2. **Direct Migration Testing**:
   - Verified that `TestMigrationToDirect` passes for both fixtures:
     - Phase 1 (Legacy Create): Successfully provisioned with legacy Terraform controller.
     - Phase 2 (Legacy Re-reconcile): No-op update succeeded.
     - Phase 3 (Direct Takeover): Direct controller adopted the resource cleanly with a 0-write no-op reconciliation.
     - Phase 4 (Direct Re-reconcile): Direct controller completed re-reconciliation with zero writes.
     - Cleanup: Resource successfully deleted.

## Policy & Log Preservation
Per `.gemini/skills/solve-migration-diff-issues/SKILL.md` (Step 1), because live GCP recording could not run due to missing permissions, mock-generated `_http_migration_phase*.log` and `_migration_diffs.json` files are omitted and not checked in.

The resource maintains `DefaultController: k8s.ReconcilerTypeTerraform` and dual support (`k8s.ReconcilerTypeDirect`, `k8s.ReconcilerTypeTerraform`) in `pkg/controller/resourceconfig/static_config.go` with dual-controller golden files retained.
