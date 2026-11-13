### [2026-11-13] Greenfield Step 1a: Prepopulate direct KRM types, identity, and judgement queue for BackupDRRestoreVerificationPlan
- **Context**: Prepopulating direct KRM types, identity, reference, and `judgement_queue.yaml` for `BackupDRRestoreVerificationPlan` under `backupdr.cnrm.cloud.google.com/v1alpha1` (Issue #15530).
- **Details**:
  1. Updated `mockgcp/apis/google/cloud/backupdr/v1/` with `restore_verification_plan.proto` and updated `backupdr.proto` (including CRUD RPCs for `RestoreVerificationPlan`), configured `dev/tools/controllerbuilder/generate-proto.sh` to copy them, and re-generated `.build/googleapis.pb`.
  2. Updated `apis/backupdr/generate.sh` to add `generate-types --prepopulate-resource --allow-non-declarative`, `generate-identity`, and updated `generate-mapper` with `--prepopulate-resource` for `BackupDRRestoreVerificationPlan:RestoreVerificationPlan`.
  3. Configured `generate-mapper` in `apis/backupdr/generate.sh` with `--skip-proto RestoreVerificationPlan,RestoreVerificationCleanupRule` so mapper generation succeeds without requiring an incompatible upgrade to `cloud.google.com/go/backupdr` in Step 1a.
  4. Executed `./apis/backupdr/generate.sh` to generate `backupdrrestoreverificationplan_types.go`, `backupdrrestoreverificationplan_identity.generated.go`, `backupdrrestoreverificationplan_reference.generated.go`, `backupdrrestoreverificationplan_identity_generated_test.go`, CRD, and `judgement_queue.yaml`.
  5. Added `BackupDRRestoreVerificationPlan` to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle`.
  6. Verified all unit tests and `apichecks` pass, updating `tests/apichecks/testdata/exceptions/alpha-missingfields.txt`.
- **Impact**: Establishes initial scaffolding and judgement queue for `BackupDRRestoreVerificationPlan` ready for Step 1b triage.
