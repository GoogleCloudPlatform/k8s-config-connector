# Migration Journal: ComputeNodeTemplate

## Takeover Diff Issues Identified and Solved

During the direct migration validation for `ComputeNodeTemplate`, we identified two key takeover diff issues between the Terraform/legacy controller's recorded GCP state and the Direct controller's mapped state:

1. **`cpu_overcommit_type` default discrepancy:**
   - **Symptom:** The Direct controller reported a difference of `"cpu_overcommit_type" ("old": "NONE", "new": "")`.
   - **Root Cause:** In the KRM spec, `cpuOvercommitType` is an optional, omitted field, which maps to `nil` in the desired state. However, GCP defaults the value to `"NONE"`. The legacy controller set this default in its state, but the Direct controller comparison logic didn't normalize this omitted optional field to its default.
   - **Solution:** Updated `populateDefaults` inside `compareComputeNodeTemplate` to set `CpuOvercommitType` to `"NONE"` if it is `nil`, ensuring both actual and desired states align perfectly.

2. **`region` URL vs. short name discrepancy:**
   - **Symptom:** The Direct controller reported a difference of `"region" ("old": "https://www.googleapis.com/compute/v1/projects/${projectId}/regions/us-central1", "new": "us-central1")`.
   - **Root Cause:** GCP returns the region as a fully-qualified URL in the API response, while KRM specifies it as a simple short name (e.g., `us-central1`).
   - **Solution:** Reused the existing package-level `lastComponent` helper in `populateDefaults` inside `compareComputeNodeTemplate` to parse and extract the last path segment from the region URL. This normalizes both actual and desired region strings to the short name before comparison.

With these changes, the `TestMigrationToDirect` E2E test runs successfully with a perfect **0-write no-op takeover** on both mock and real GCP targets.
