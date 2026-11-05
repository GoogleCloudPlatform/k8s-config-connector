---
name: reviewgen-greenfield-mockgcp
description: Review criteria for KCC Greenfield Step 3 (MockGCP implementation) and Step 4 (MockGCP/RealGCP log alignment) PRs. Verifies the mock service, untouched real-GCP baselines, golden log alignment, normalizer scoping, and scope limits.
---

# Review Guide for KCC Greenfield MockGCP PRs (Steps 3 and 4)

Use this guide for a **Greenfield Step 3** PR (`Greenfield: Implement MockGCP and Alignment for <Kind>`) or a **Greenfield Step 4** PR (`Greenfield: Align MockGCP logs with RealGCP for <Kind>`). The real-GCP logs (`_http.log`, `_audit_probe.log`) were recorded in Step 2; these steps make MockGCP reproduce them.

## Detect the step

- **Step 3**: the PR adds or extends the mock service for `<Kind>` under `mockgcp/`, usually with new `_http_mock.log` files.
- **Step 4**: the mock already exists on `master`; the PR fixes mock behavior or normalizers so the mock logs match the real ones.

## 1. Checks for both steps

*   **Real-GCP baseline**: No `_http.log` or `_audit_probe.log` changed, unless the PR explains a fixture or request change and the new log comes from `hack/record-gcp`. No log or golden file was edited by hand.
*   **Normalizers**: `Previsit` is scoped to `<service_name>.googleapis.com` URLs. API-version and URL-prefix replacements are exact strings for this service only, with no broad wildcards.
*   **Scope**: No legacy (Terraform/DCL) controller changes, no new `missingfields.txt` lines, fixture directory names contain the kind name, and `<Kind>` isn't excluded in `ShouldTestRereconiliation` (`tests/e2e/ratcheting.go`).
*   **`BLOCKED_KINDS`**: `<Kind>` is still in `BLOCKED_KINDS` in `dev/tasks/build-release-bundle`: only the pipeline's final Release step removes it.

## 2. Step 3: MockGCP implementation

*   **Mock service**: `<Kind>` is implemented in `mockgcp/mock<service_name>/`, and a new service is registered in `mockgcp/register.go`. CRUD handles the parent like real GCP, including rejecting a missing parent.
*   **Alignment and CI**: Every fixture has `_http_mock.log` and goldens. `TestGoldenLogAlignment/<kind_lower>` passes with zero errors, and `./dev/ci/presubmits/tests-e2e-fixtures-suite <service>` passes.
*   **Mock behavior**: Defaults are set on both write and read paths, and proto3 zero-value gaps have `TODO`s.
*   **Release note**: `NONE`. The pipeline's final Release step adds the new-resource note.

## 3. Step 4: MockGCP/RealGCP alignment

*   **Alignment**: `TestGoldenLogAlignment/<kind_lower>` passes with zero errors for every fixture, `hack/compare-mock "fixtures/<kind_lower>"` runs cleanly, and `./dev/ci/presubmits/tests-e2e-fixtures-suite <service>` passes.
*   **Fix quality**: Defaults are applied on both write and read paths, conditional defaults are commented, proto3 zero-value gaps and missing proto fields have `TODO`s, and canonicalization happens in the mock's Create.

# Review Comment Template

```markdown
### KCC Greenfield MockGCP (Step 3/4) Auto-Review Results
* **Step**: [3 (MockGCP) / 4 (Alignment)]
* **Real-GCP Baseline Untouched**: [Pass/Fail]
* **Mock Service & CRUD (Step 3) / Fix Quality (Step 4)**: [Pass/Fail]
* **Golden Log Alignment & CI Fixtures**: [Pass/Fail]
* **Normalizer Scoping**: [Pass/Fail]
* **Scope, Exceptions & `BLOCKED_KINDS`**: [Pass/Fail]

#### Detailed Findings / Actions Required:
1. [Specify file, line number, and exact issue]
```
