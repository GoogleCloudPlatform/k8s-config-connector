---
name: run-presubmits-and-linters
description: Guides through running and debugging CI presubmit scripts locally and understanding custom repository linters.
---

# Run Presubmits and Linters

This skill guides an automated agent through running and debugging Config Connector CI presubmit scripts locally and understanding custom repository linters.

## Presubmit Scripts Overview

All presubmits run in GitHub Actions, generated from `dev/ci/presubmits/`. Each script in `dev/ci/presubmits/` maps 1:1 to a CI job.

If a specific CI presubmit fails, run the corresponding script locally to diagnose and fix the failure.

### Common Presubmit Scripts

- **`./dev/ci/presubmits/unit-tests`**: Runs unit tests and verifies golden log alignment across packages.
- **`./dev/ci/presubmits/validate-generated-files`**: Validates CRDs and generated mapper/types files (`dev/tasks/generate-types-and-mappers`).
- **`./dev/ci/presubmits/tests-e2e-fixtures-<service>`**: Runs E2E fixture tests against mock GCP for the specified service.
- **`./dev/ci/presubmits/tests-e2e-samples-<service>`**: Runs sample tests for the specified service.

### Fixing Out-of-Date Golden Output

When golden logs or YAML fixtures are out of date, run the failing presubmit script with `WRITE_GOLDEN_OUTPUT=1`:
```bash
WRITE_GOLDEN_OUTPUT=1 ./dev/ci/presubmits/tests-e2e-fixtures-<service>
```
Review the resulting git diff to verify that changes match expected GCP API behavior.

---

## Custom Linters

Custom linters are maintained under `dev/linters/`.

### `jsonunmarshalreuse`

Checks for suboptimal `json.Unmarshal` (and `util.Marshal`) practices where a non-empty variable might be reused:
- **Slices**: Unmarshalling into a non-empty slice causes existing elements to be lost/overwritten (including slices allocated via `make([]T, N)` with `N > 0`).
- **Maps and Structs**: Unmarshalling into non-empty variables results in merging existing elements with unmarshalled values.
- **Ignored Fields**: Struct fields tagged with `json:"-"` are ignored.
- **Tests**: Intentional reuse in test files (e.g., `pkg/k8s/managedfields_test.go`) is expected when explicitly testing merge semantics.

---

## Standard Pre-PR Validation

Before creating or pushing commits, execute:
```bash
make fmt && go vet ./...
./dev/ci/presubmits/unit-tests
```
