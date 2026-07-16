---
name: reviewgen-brownfield-feature
description: Comprehensive review guide and criteria for PRs adding or modifying features in brownfield resources (TF patches, DCL patches, CRD changes, MockGCP alignment, and Direct Controller alignment).
---

# Review Guide for Brownfield Resource Feature Changes

This skill outlines the mandatory review criteria for PRs that introduce new fields, update behaviors, or backport patches for **Brownfield resources** (Terraform-based or DCL-based resources).

---

## 1. Terraform (TF) Based Resource Changes

### Patch Verification
*   **Target Files**: Patch files are under `third_party/github.com/hashicorp/terraform-provider-google-beta/google-beta/services/[servicename]/resource_[servicename]_[typename].go` (or extracted schema files like `node_config.go`).
*   **Adding New Fields**:
    1. Verify schema addition in the schema map (`map[string]*schema.Schema`).
    2. Verify expanding logic (KRM/Terraform struct -> GCP SDK) is implemented for new fields.
    3. Verify flattening logic (GCP SDK -> Terraform schema) is implemented for new fields.
*   **Modifying Existing Fields (Backwards-Compatible Behavior Changes)**:
    1. Schema definition updated as required.
    2. Related expanding/flattening logic updated to match the modified schema.

### CRD & Service Mapping Changes
*   **Resource References**: If a new reference field is added, ensure it is properly configured in `config/servicemappings/[servicename].yaml` under `resourceReferences`.
*   **CRD Schema Alignment**: Verify that schema changes are reflected in the corresponding generated CRD file under `config/crds/resources/`.

---

## 2. Declarative Resource Client Library (DCL) Based Resource Changes

> [!IMPORTANT]
> Triple check if DCL changes are strictly necessary before approving. Direct migrations or TF backports are preferred.

### Patch Verification
*   **Target Files**: Patch files are under `third_party/github.com/GoogleCloudPlatform/declarative-resource-client-library/services/google/[servicename]/`.
*   **Multi-Version Requirement**: If multiple versions exist for the type, changes **must** be made across all supported versions.
*   **Adding New Fields**:
    1. Schema file `[typename].yaml` updated with the new fields.
    2. Expanding logic added in `[typename]_internal.go`.
    3. Flattening logic added in `[typename]_internal.go`.
    4. Supporting logic updated in `[typename].go` or helper files.
*   **Modifying Existing Fields**:
    1. Schema file `[typename].yaml` updated.
    2. Implementation logic updated to match.

### CRD Schema Changes
*   Ensure reference fields are configured in `[typename].yaml`.
*   Verify CRD YAML changes are generated in `config/crds/resources/`.

---

## 3. Shared Review Invariants

### A. Test Coverage
*   **New Fields**: Must be covered in `create.yaml` and/or `update.yaml` within a dedicated fixture test directory under `pkg/test/resourcefixture/testdata/basic/`. Multiple test cases are required if fields are mutually exclusive (`oneOf`).
*   **Existing Fields / Behavior Changes**: New behavior or state transitions must be covered in `create.yaml` and/or `update.yaml`.

### B. MockGCP and Real GCP Alignment

1. **Determining Target Log Files & Workflow via `static_config.go`**:
   - Inspect `pkg/controller/resourceconfig/static_config.go` for the resource's `{Group, Kind}` entry under `ControllerConfigStatic`.
   - **Case 1: Unmigrated Dual-Controller Resource** (`DefaultController` is `Terraform` or `DCL`, but `SupportedControllers` includes `Direct`):
     - Target logs to compare for the TF/DCL patch: compare **`_http_old_controller.log`** vs **`_http_old_controller_mock.log`**.
   - **Case 2: Unmigrated Single-Controller Resource** (Only `Terraform` or `DCL` in `SupportedControllers`):
     - Target logs to compare for the TF/DCL patch: compare **`_http.log`** vs **`_http_mock.log`**.
   - **Case 3: Migrated Resource (Yellow Flag ⚠️)** (`DefaultController` is `Direct`):
     - **Warning**: It does not make sense to add or modify features via a legacy TF/DCL patch once the default controller has been cut over to `Direct`.
     - **Review Action**: Flag this in the review summary as a Yellow Flag. Instruct the contributor to follow the Direct Controller feature workflow instead (implementing the new fields/features directly in `pkg/controller/direct/` and `apis/`).

2. **Verification of Real GCP Test Execution (Mandatory First Step)**:
   * Ensure that the real GCP test has actually been executed.
   * **Detection Heuristics**:
     - Look for expected minor/unimportant diffs between real logs and mock logs (e.g., dynamic server timestamps, IDs, or volatile hashes).
     - **Identical Logs Check**: If real and mock logs are 100% identical with no minor/dynamic diffs, flag this in the review to check with the contributor to verify if the real GCP test (`hack/record-gcp`) was actually run (or if mock logs were merely copied).
     - Or confirm if the human contributor explicitly states in PR comments/description that they ran the real GCP test.
   * **Bot-Submitted PRs**: If the PR is submitted by a code bot (e.g., `lovelace-coder-bot`), explicitly ask the human reviewer to verify if the bot has run the real GCP test (providing the time range during which the real test could have been running).

3. **Mock Support**: MockGCP support for the updated brownfield resource must be implemented or verified to run tests hermetically.

4. **HTTP Log Diffs & Verification Rules**:
   * **Single Controller (Unmigrated)**: Compare real GCP HTTP logs (`_http.log`) against mock GCP HTTP logs (`_http_mock.log`).
   * **Dual Controller (Unmigrated)**: Compare legacy real GCP logs (`_http_old_controller.log`) against legacy mock logs (`_http_old_controller_mock.log`).
   * **Strict Rule**: Verify that real vs mock log diffs **do not contain discrepancies in the newly modified/added fields**. Any discrepancy indicates a bug or gap in the MockGCP implementation, but we'll only focus on the discrepancies that impact the target feature in the PR.
   * **`golden_alignment_test.go` Test Coverage Rule**: Automated log comparison is executed by `pkg/test/resourcefixture/golden_alignment_test.go`. If the review identifies unexpected log diffs or discrepancies in the target feature that `golden_alignment_test.go` missed, updating the test file is **out of scope** for the feature PR; instead, **`reviewgen` should create a GitHub issue** (e.g. via `gh issue create`) to capture this test gap and report the issue back in the review summary.

5. **KRM Golden Object File Verification (`_generated_object_<testname>.golden.yaml`)**:
   * **Dual-Controller Object Diffs**: For unmigrated dual-controller resources, check `_final_object_old_controller.golden.yaml` against `_generated_object_<testname>.golden.yaml` (and inspect `_final_object.diff`) to ensure status alignment between the legacy controller and direct controller.

### C. Direct Controller Alignment
1. **Adding New Fields**:
   * **Direct Types, Mappers & Fuzzers**: If the brownfield resource already has a direct type (`apis/`), mapper (`pkg/controller/direct/`), and/or fuzzer, ensure changes are added to those files.
   * **Reference Types**: If the new field requires a reference type that was not already implemented, the reference type must be added.
   * **Direct Controller `Update()` Logic**: If the resource already has a direct controller (`_controller.go`), verify if special handling is needed in `Update()`:
     - *Value Equivalence Diffing*: If the desired value and actual value of the new field are formatted differently but mean the same thing, the diffing logic in `Update()` must be updated to recognize equivalence.
     - *Field Comparison Lists*: When `Update()` compares desired and actual resources based on explicit field names, add the new fields to the comparison list.
2. **Modifying Existing Fields (Backwards-Compatibility Changes)**:
   * **Direct Types**: If the brownfield resource already has direct types in `apis/`, update the type files to reflect the schema changes.
   * **Direct Mappers, Fuzzers & Controllers**: If the brownfield resource already has a direct mapper, fuzzer, or controller, determine if updates are required:
     - *Immutability Check Removal*: If a field becomes mutable, check the controller logic and remove any obsolete immutability validation checks that would prevent in-place updates.
     - *Mapper/Fuzzer Alignment*: Update mappers and fuzzers to ensure proper translation and test coverage for the modified field.
3. **Dual-Controller Traffic Alignment Verification**:
   * For an **Unmigrated Dual-Controller Resource**, after verifying code/type changes, inspect both real logs (**`_http_old_controller.log`** vs **`_http.log`**, `_http.diff`) and mock logs (**`_http_old_controller_mock.log`** vs **`_http_mock.log`**, `_http_mock.diff`) to ensure traffic aligns for the newly added/modified fields.
   * Discrepancies between legacy controller traffic and direct controller traffic for the target feature indicate that the direct mapper/controller was not properly updated or handles the new field differently.

---

## Standard Review Summary Template

```markdown
### Brownfield Feature Review Results

* **Real GCP Test Verification**: [Pass/Pending/Needs-Human-Verification] - (Verified via minor _http.log vs _http_mock.log diffs or contributor claim. For bot PRs: prompt human reviewer with timestamp window)
* **TF / DCL Patch Integrity**: [Pass/Fail] - (Verified schema, expanders, and flatteners)
* **CRD & Service Mapping**: [Pass/Fail] - (CRD and servicemapping updated)
* **Test Fixture Coverage**: [Pass/Fail] - (Verified create.yaml / update.yaml coverage)
* **MockGCP and Real GCP Alignment**: [Pass/Fail] - (No unexpected diffs in new fields in _http.log vs _http_mock.log)
* **Direct Controller Alignment**: [Pass/Fail] - (Updated apis/, mappers, fuzzers, and immutability checks if present)

#### Actionable Findings:
1. [File line number and specific issue]
```
