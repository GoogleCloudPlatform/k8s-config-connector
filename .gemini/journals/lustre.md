# Lustre Journal

## Resource
- Kind: `LustreInstance`
- Proto Message: `google.cloud.lustre.v1.Instance`
- Service: `lustre`
- API Version: `lustre.cnrm.cloud.google.com/v1alpha1`

## Notes & Findings
- **Initial Greenfield Prepopulation (Step 1a)**:
  - Added `apis/lustre/generate.sh` with `generate-types`, `generate-identity`, and `generate-mapper` commands using `--prepopulate-resource`.
  - Identity pattern: `projects/{project}/locations/{location}/instances/{instance}`.
  - Generated types, identity, reference, and mappers.
  - Added `LustreInstance` to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle`.
  - Added `cloud.google.com/go/lustre` at `v0.7.0` to `go.mod` without forcing unwanted transitive dependency upgrades.
  - Updated `tests/apichecks/testdata/exceptions/alpha-missingfields.txt` and `operator/pkg/controllers/configconnector/testdata/golden/simple/_expected.yaml`.
  - All unit tests and apichecks pass.
- **Judgement Queue Triage & Finalization (Step 1b)**:
  - Inlined `*parent.ProjectAndLocationRef` replacing separate `ProjectRef` and `Location` fields on `LustreInstanceSpec`.
  - Converted `network` string to `NetworkRef *computerefs.ComputeNetworkRef` referencing `compute.googleapis.com/Network`.
  - Omitted deprecated `gkeSupportEnabled` field from `LustreInstanceSpec`.
  - Removed top-level `Labels` from `LustreInstanceSpec` (managed via `metadata.labels`).
  - Regenerated CRD and `mapper.generated.go`.
  - Re-recorded `tests/apichecks/testdata/exceptions/alpha-missingfields.txt`.
  - All tests and apichecks pass.

