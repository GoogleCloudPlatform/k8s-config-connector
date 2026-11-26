### [2026-11-26] Greenfield Step 2: Implement direct controller, E2E fixtures, and fuzzer for BackupDRRestoreTemplate
- **Context**: Greenfield Step 2 implementation for `BackupDRRestoreTemplate` (Issue #16531).
- **Details**:
  1. Generated direct controller (`backupdrrestoretemplate_controller.go`) and round-trip fuzzer (`backupdrrestoretemplate_fuzzer.generated.go`) using `controllerbuilder` one-shot commands (`generate-controller` and `generate-fuzzer --prepopulate-resource`).
  2. Registered `BackupDRRestoreTemplate` with `k8s.ReconcilerTypeDirect` in `pkg/controller/resourceconfig/static_config.go`.
  3. Created minimal and maximal test fixtures under `pkg/test/resourcefixture/testdata/basic/backupdr/v1alpha1/backupdrrestoretemplate/` covering all `BackupDRRestoreTemplateSpec` fields (`projectRef`, `location`, `resourceType`, `labels`, `description`, `backupSelectionConfig`, `preRestoreConfig`, `postRestoreConfig`, `verificationConfig`, and `restoreProperties`).
  4. Updated `tests/apichecks/testdata/exceptions/alpha-missingfields.txt` to remove all 13 fields now covered by fixtures.
  5. Verified 200-seed fuzzer round-trip (`FOCUS=BackupDRRestoreTemplate go test -v ./pkg/fuzztesting/fuzztests -run TestFocusedMappers`), CRD schema tests, and direct controller compilation.
- **Observations & Learnings**:
  - `backupdr.googleapis.com` RestoreTemplate operations (`CreateRestoreTemplate`) require allowlisting for "Automated Restore Verification" (`Project ... is not allowlisted for Automated Restore Verification. This feature is currently available for allowlisted projects only.`). Live recording against real GCP is blocked in non-allowlisted projects. Minimal and maximal fixtures are defined with valid references and configurations ready for mock implementation in Phase 3.

### [2026-11-26] Greenfield Step 1b: Triage judgement queue and finalize direct KRM types and identity for BackupDRRestoreTemplate
- **Context**: Triaged `apis/backupdr/judgement_queue.yaml` and finalized KRM types and identity for `BackupDRRestoreTemplate` (Issue #16512).
- **Details**:
  1. Inlined `*parent.ProjectAndLocationRef` in `BackupDRRestoreTemplateSpec`, replacing `ProjectRef` and `Location`.
  2. Moved nested types `BackupSelectionConfig`, `PreRestoreConfig`, `PostRestoreConfig`, and `VerificationConfig` into `backupdrrestoretemplate_types.go`.
  3. Implemented `BackupDRDataSource` external-only reference and identity (`backupdrdatasource_identity.go`, `backupdrdatasource_reference.go`, and conformance test `backupdrdatasource_identity_test.go`) with CAI pattern `projects/{project}/locations/{location}/backupVaults/{backupVault}/dataSources/{dataSource}`. `GetNamespacedName()` returns empty and `Normalize()` validates `r.External`. Converted `dataSource` in `BackupSelectionConfig` to `dataSourceRef *BackupDRDataSourceRef`.
  4. Converted `cloudRunJob` fields in `PreRestoreConfig`, `PostRestoreConfig`, and `VerificationConfig` to `cloudRunJobRef *runv1beta1.JobRef`.
  5. Implemented handwritten ToProto mappers in `pkg/controller/direct/backupdr/restoretemplate_mappers.go` to handle mapping `dataSourceRef` and `cloudRunJobRef` into proto oneof wrappers (`BackupSelectionConfig.source` and `pre_restore_action`/`post_restore_action`/`verification_action`).
  6. Moved `etag` from `Spec` to `ObservedState`, removing it from `tests/apichecks/testdata/exceptions/spec_dislike_etag.txt`.
  7. Preserved `Spec.Labels` on greenfield direct resource `BackupDRRestoreTemplateSpec` per KCC conventions.
  8. Verified CAI URL template `projects/{project}/locations/{location}/restoreTemplates/{restoreTemplate}` matching proto resource definition and resolved `identity-not-in-cai` as accepted.
  9. Removed unverified REST documentation URL (`+kcc:guess=source-link reason=verify-resource-docs-link` and `+kcc:source:resource-docs=...`).
  10. Triaged and resolved all 9 entries for `BackupDRRestoreTemplate` in `apis/backupdr/judgement_queue.yaml`.
  11. Updated `tests/apichecks/testdata/exceptions/naming_violations.txt` and `alpha-missingfields.txt`.
- **Impact**: Finalized direct KRM types, identity, external reference, and CRDs for `BackupDRRestoreTemplate`.

### [2026-11-26] Prepopulate direct KRM types, identity, judgement queue, and mappers for BackupDRRestoreTemplate
- **Context**: Prepopulating BackupDRRestoreTemplate under `apis/backupdr/v1alpha1` (Greenfield Step 1a, Issue #15531).
- **Protos & Dependencies**: RestoreTemplate proto (`google.cloud.backupdr.v1.RestoreTemplate`) is present in `origin/master` of googleapis. Added the protos under `mockgcp/apis/google/cloud/backupdr/v1/` and wired them in `dev/tools/controllerbuilder/generate-proto.sh` and `sourcelinks.go` (`kccProtoOverrides`). Upgraded `cloud.google.com/go/backupdr` to `v1.20.0` in `go.mod` and `mockgcp/go.mod` without replace directives.
- **Mappers**: Generated direct mappers with `--prepopulate-resource` in `apis/backupdr/generate.sh` into `pkg/controller/direct/backupdr/mapper.generated.go`. In `pkg/controller/direct/backupdr/backupdrbackupplanassociation_fuzzer.go`, marked new fields introduced in v1.20.0 (`.alloydb_cluster_backup_plan_association_properties` and `.auto_protection_policy_binding`) as untriaged per `.gemini/skills/update-client-library-version/SKILL.md`.
- **Identity & References**: Generated `backupdrrestoretemplate_identity.generated.go` and `backupdrrestoretemplate_reference.generated.go` with pattern `projects/{project}/locations/{location}/restoreTemplates/{restore_template}`. Conformance tests pass.
- **Judgement Queue**: Generated `apis/backupdr/judgement_queue.yaml` with open `untriaged-bulk-generation` and related entries for Step 1b triage.
- **Release Bundle**: Added `BackupDRRestoreTemplate` to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle`.
- **Exceptions**: Updated `tests/apichecks/testdata/exceptions/alpha-missingfields.txt` and `spec_dislike_etag.txt` with `backupdrrestoretemplates` entries.

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
