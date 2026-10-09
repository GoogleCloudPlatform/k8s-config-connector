# Migration Diff Journal for ArtifactRegistryRepository

- **Issue**: #13870
- **Resource**: `ArtifactRegistryRepository` (`artifactregistry.cnrm.cloud.google.com`)
- **Reported diff fields**: `cleanup_policy_dry_run`, `cleanup_policies`
- **Initial Date/Time**: Fri Oct 9 09:58:52 PDT 2026

## Log of Attempts & Runs

### Run 1: Baseline TestMigrationToDirect on Existing Fixtures
- **Start Time (PDT)**: Fri Oct 9 09:58:52 PDT 2026
- **End Time (PDT)**: Fri Oct 9 10:20:00 PDT 2026
- **Target**: `TestMigrationToDirect/fixtures/artifactregistryrepository`
- **Goal**: Run `TestMigrationToDirect` against real GCP for existing `ArtifactRegistryRepository` fixtures to observe baseline behavior and check whether the reported diffs reproduce on existing fixtures.
- **Results**:
  - `artifactregistryrepository`: 0 migration diffs reported for direct takeover. Only legacy creation diffs.
  - `artifactregistryrepositorycommonrepository`: 0 migration diffs reported for direct takeover. Only legacy creation diffs.
  - Neither existing fixture configured `cleanupPolicies` or `cleanupPolicyDryRun` (one is virtual repository, one is remote repository).
- **Conclusion**: Existing fixtures do not configure `cleanupPolicies` or `cleanupPolicyDryRun`. Following Step 1.3 of the issue description, added a new fixture `artifactregistryrepositorycleanuppolicies` to test migration of standard repositories with cleanup policies.

### Run 2: TestMigrationToDirect on artifactregistryrepositorycleanuppolicies (Explicit tagState)
- **Start Time (PDT)**: Fri Oct 9 10:29:58 PDT 2026
- **Finish Time (PDT)**: Fri Oct 9 10:31:52 PDT 2026
- **Target**: `TestMigrationToDirect/fixtures/artifactregistryrepositorycleanuppolicies`
- **Configuration**: Explicitly configured `tagState: "TAGGED"` on policy condition and `cleanupPolicyDryRun: false`.
- **Results**:
  - Direct controller did not report diffs because all fields were explicitly populated in KRM and matched GCP.

### Run 3: TestMigrationToDirect on artifactregistryrepositorycleanuppolicies (Defaulted tagState / cleanupPolicyDryRun)
- **Start Time (PDT)**: Fri Oct 9 10:41:32 PDT 2026
- **Finish Time (PDT)**: Fri Oct 9 10:43:35 PDT 2026
- **Target**: `TestMigrationToDirect/fixtures/artifactregistryrepositorycleanuppolicies`
- **Configuration**: Omitted `tagState` from condition (testing Terraform's default of `ANY`) and omitted `cleanupPolicyDryRun` (testing optional boolean handling).
- **Results**:
  - **Bug reproduced**:
    - `_migration_diffs.json` recorded direct takeover diff:
      `"controller": "direct", "isNewObject": false`, with diff on `"cleanup_policies"` (`old` tag_state: 3 [ANY], `new` tag_state: nil).
    - `_http_migration_phase3_direct_takeover.log` recorded unexpected non-GET call:
      `PATCH https://artifactregistry.googleapis.com/v1/projects/${projectId}/locations/us-west1/repositories/arrepository-${uniqueId}?%24alt=json%3Benum-encoding%3Dint&updateMask=cleanupPolicies`.
    - Test failed with: `FAIL: unexpected write request during migration reconciliation: PATCH ... updateMask=cleanupPolicies`.
- **Root Cause**:
  1. In the Terraform provider schema, `tag_state` has `Default: "ANY"`. When `tagState` is omitted from KRM, Terraform provisions the policy with `tagState: ANY` (enum value 3) on GCP. The direct controller does not default `tagState` to `ANY`, so upon takeover it detects a mismatch between GCP (`ANY`) and desired KRM (unset/nil), triggering an unnecessary `PATCH` call.
  2. For `cleanup_policy_dry_run`, `cleanupPolicyDryRun` is optional in KRM. In direct controller `ToProto`, `direct.ValueOf(in.CleanupPolicyDryRun)` forces `false` if `nil`, which can differ if omitted or not explicitly managed.
