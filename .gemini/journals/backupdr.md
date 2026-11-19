### [2026-11-16] Greenfield Step 2: Implement direct controller, E2E fixtures, and fuzzer for BackupDRAutoProtectionPolicy
- **Context**: Greenfield Step 2 implementation for `BackupDRAutoProtectionPolicy` (Issue #15802).
- **Details**:
  1. Upgraded `cloud.google.com/go/backupdr` to `v1.20.0` across `go.mod` and `mockgcp/go.mod` to support `AutoProtectionPolicy` client methods and protobuf types.
  2. In `pkg/controller/direct/backupdr/backupdrbackupplanassociation_fuzzer.go`, marked new fields introduced in v1.20.0 (`.alloydb_cluster_backup_plan_association_properties` and `.auto_protection_policy_binding`) as untriaged.
  3. Removed `AutoProtectionPolicy` messages from `--skip-proto` in `apis/backupdr/generate.sh` and regenerated mappers into `pkg/controller/direct/backupdr/mapper.generated.go`.
  4. Generated direct controller (`backupdrautoprotectionpolicy_controller.go`) and roundtrip fuzzer (`backupdrautoprotectionpolicy_fuzzer.generated.go`).
  5. Implemented `ObservedState` mapping on `Create` and `Update` from the LRO response.
  6. Added reference normalization in `Update` via `projectMapper.ReplaceProjectNumberWithIDInLink` for `detail.BackupPlan` to prevent spurious update diffs caused by GCP returning project numbers in resource URLs while KRM spec provides project IDs.
  7. Created minimal and full E2E fixtures (`backupdrautoprotectionpolicy-minimal` and `backupdrautoprotectionpolicy-full`), recording them against real GCP. Discovered that GCP BackupDR MVP restricts `criteria.matchingConditions[].labelCondition.values` to exactly 1 value per condition.
  8. Verified all tests pass: `TestFocusedMappers`, `TestDirectResourceFileNaming`, `TestGenerateScripts`, `TestRegisteredTemplatesMatchCAI`, and real GCP E2E validation checks (7/7 passed).
- **Impact**: Delivers complete direct controller, roundtrip fuzzer, and real GCP verified E2E fixtures for `BackupDRAutoProtectionPolicy`.

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

### [2026-11-12] BackupDRAutoProtectionPolicy Proto Overlay, Prepopulation, and Judgement Queue
- **Context**: Prepopulating direct KRM types, identity, and judgement queue for `BackupDRAutoProtectionPolicy` under `apis/backupdr/v1alpha1`.
- **Problem**: The pinned googleapis commit in `apis/git.versions` does not contain `AutoProtectionPolicy`, which was added upstream in `google/cloud/backupdr/v1/autoprotectionpolicy.proto` and `backupdr.proto`.
- **Solution**:
  1. Created a proto overlay `mockgcp/apis/google/cloud/backupdr/v1/autoprotectionpolicy.proto` and `mockgcp/apis/google/cloud/backupdr/v1/backupdr.proto` that preserves the pinned base for existing BackupDR resources while adding `AutoProtectionPolicy`, its RPCs, and its request/response messages.
  2. Registered the overlay in `dev/tools/controllerbuilder/generate-proto.sh` and in `kccProtoOverrides` in `dev/tools/controllerbuilder/pkg/sourcelinks/sourcelinks.go`.
  3. Configured `generate-mapper` in `apis/backupdr/generate.sh` with `--skip-proto` for `AutoProtectionPolicy` and its nested messages so mapper generation succeeds without requiring an upgrade to `cloud.google.com/go/backupdr` or any `replace` directives in `go.mod`.
  4. Added `BackupDRAutoProtectionPolicy` to `dev/tasks/build-release-bundle` `BLOCKED_KINDS` list in alphabetical order.
  5. Added `generate-types` with `--prepopulate-resource`, `generate-identity`, and updated `generate-mapper` in `apis/backupdr/generate.sh`.
  6. Updated `alpha-missingfields.txt` and `spec_dislike_etag.txt` in `tests/apichecks/testdata/exceptions/`.
- **Impact**: Enables clean mechanical scaffolding of `BackupDRAutoProtectionPolicy` with types, IdentityV2, and judgement queue open for Step 1b triage without polluting `go.mod` or risking dependency drift.
