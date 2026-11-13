### [2026-11-13] Prepopulate direct KRM types, identity, and judgement queue for BackupDRRestoreTemplate
- **Context**: Prepopulating BackupDRRestoreTemplate under `apis/backupdr/v1alpha1` (Step 1a).
- **Protos & Dependencies**: RestoreTemplate proto (`google.cloud.backupdr.v1.RestoreTemplate`) is present in `origin/master` of googleapis. Added the protos under `mockgcp/apis/google/cloud/backupdr/v1/` and wired them in `dev/tools/controllerbuilder/generate-proto.sh` and `sourcelinks.go` (`kccProtoOverrides`). Skipped generating mappers for `RestoreTemplate` via `--skip-proto` in `apis/backupdr/generate.sh` so mappers and client library updates can be handled in dedicated steps without cascading dependency bumps.
- **Identity & References**: Generated `backupdrrestoretemplate_identity.generated.go` and `backupdrrestoretemplate_reference.generated.go` with pattern `projects/{project}/locations/{location}/restoreTemplates/{restore_template}`. Conformance tests pass.
- **Judgement Queue**: Generated `apis/backupdr/judgement_queue.yaml` with open `untriaged-bulk-generation` and related entries for Step 1b triage.
- **Release Bundle**: Added `BackupDRRestoreTemplate` to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle`.
- **Exceptions**: Updated `tests/apichecks/testdata/exceptions/alpha-missingfields.txt` and `spec_dislike_etag.txt` with `backupdrrestoretemplates` entries.
