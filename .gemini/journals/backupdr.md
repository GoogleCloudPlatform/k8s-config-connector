### [2026-11-13] Greenfield Step 1b: Triage judgement queue and finalize direct KRM types and identity for BackupDRRestoreVerificationPlan
- **Context**: Triaged `apis/backupdr/judgement_queue.yaml` and finalized KRM types and identity for `BackupDRRestoreVerificationPlan` (Issue #15568).
- **Details**:
  1. Inlined `*parent.ProjectAndLocationRef` in `BackupDRRestoreVerificationPlanSpec`, replacing `ProjectRef` and `Location`.
  2. Removed `Labels map[string]string` from `Spec` in favor of top-level KRM `metadata.labels`.
  3. Confirmed resource ID is client-specified via `restore_verification_plan_id` in `CreateRestoreVerificationPlanRequest`.
  4. Verified CAI URL template `//backupdr.googleapis.com/projects/{project}/locations/{location}/restoreVerificationPlans/{restoreVerificationPlan}` matching proto resource definition.
  5. Removed unverified REST documentation URL (`+kcc:guess=source-link reason=verify-resource-docs-link` and `+kcc:source:resource-docs=...`) as public documentation page does not exist yet.
  6. Resolved all entries in `apis/backupdr/judgement_queue.yaml`.
  7. Ran `./apis/backupdr/generate.sh` to update `zz_generated.deepcopy.go` and CRD definitions, and updated `tests/apichecks/testdata/exceptions/alpha-missingfields.txt`.
- **Impact**: Finalized direct KRM types, identity, and CRDs for `BackupDRRestoreVerificationPlan`.

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
