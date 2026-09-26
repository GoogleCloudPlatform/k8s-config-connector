# ComputeSecurityPolicy Direct Migration Journal

## Overview
Successfully promoted `ComputeSecurityPolicy` to be direct-by-default, enabled server-side apply (SSA) re-reconciliation verification, and resolved 0-write re-reconciliation differences.

## Learnings & Observations

### 1. Default Reconciler Promotion
- Changed the default controller for `ComputeSecurityPolicy` from `ReconcilerTypeTerraform` to `ReconcilerTypeDirect` in `pkg/controller/resourceconfig/static_config.go`.
- Removed `ComputeSecurityPolicy` from the skipped re-reconciliation list in `tests/e2e/ratcheting.go`.

### 2. 0-write Re-reconciliation & GCP Defaults Alignment
- Discovered that during the re-reconciliation phase, the direct controller detected spurious diffs on the `rules` list, triggering unintended `PATCH` operations.
- This occurred because the GCP-assigned defaults for rules—such as `Kind: "compute#securityPolicyRule"`, `Description: ""` (when not specified), `Preview: false` (when not specified), and nested matcher `ExprOptions` / `Expr` subfields—existed on the `actual` rules from the server but were missing from `desired`.
- Solved this by modifying `assignGCPDefaults` in `computesecuritypolicy_controller.go` to match `desired.Rules` and `actual.Rules` elements by their unique `Priority` and align those read-only and defaulted fields from actual to desired before performing the structured diff. This successfully established true 0-write re-reconciliation.

### 3. Cleaning Up Obsolete Golden Files
- Since the resource is now direct-by-default, the test runner no longer records comparative old-controller execution logs.
- Cleaned up obsolete comparative logs/diffs (`*_old_controller*` and `*.diff`) under the `pkg/test/resourcefixture/testdata/basic/compute/v1beta1/computesecuritypolicy` test cases to ensure the presubmit and e2e checks pass cleanly.
