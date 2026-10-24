---
name: reviewgen-greenfield-types-prepopulate
description: Lightweight review checklist for Greenfield Step 1 (Mechanical Prepopulate) PRs. Verifies `--prepopulate-resource` flags, open `judgement_queue.yaml` entries, zero manual edits to generated types, mapper compilation, and `apichecks`.
---

# Review Guide for KCC Greenfield Step 1 PRs — Prepopulate Checklist

Use this review guide when reviewing a **Greenfield Step 1 (Mechanical Prepopulate)** PR (or multi-resource batch prepopulate PR), where new resources are scaffolded with `controllerbuilder` and `apis/<service>/judgement_queue.yaml` has `reason: untriaged-bulk-generation` with `status: open`.

Please enforce all of the following review criteria and invariants.

---

## 1. `apis/<service>/generate.sh` Flags & GA Proto Requirement

*   **Required Commands, Order, & Flags**:
    *   Verify `apis/<service>/generate.sh` runs in this exact order:
        1. `controllerbuilder generate-types` with `--resource <Kind>:<ProtoMessage>` and `--prepopulate-resource`.
        2. `controllerbuilder generate-identity` with `--resource <Kind>:<ProtoMessage>`.
        3. `controllerbuilder generate-mapper` **after** all `generate-types` and `generate-identity` calls, with `--prepopulate-resource` (plus `--multiversion` if the service has both `v1alpha1` and `v1beta1`).
        4. `cd ${REPO_ROOT}`, `dev/tasks/generate-crds`, and `goimports` on `pkg/controller/direct/<service>/`.
*   **No `--allow-non-ga-proto`**:
    *   Reject the PR if `--allow-non-ga-proto` is passed in `apis/<service>/generate.sh` (unless the PR issue explicitly authorized a non-GA proto exception).
*   **API Version & Copyright**:
    *   New greenfield resources must be scaffolded under `apis/<service>/v1alpha1/` (`v1alpha1`).
    *   All new `.go` and `.sh` files must carry a `2026` copyright header.

---

## 2. Judgement Queue Open State (`apis/<service>/judgement_queue.yaml`)

*   **`untriaged-bulk-generation` Must Be `open`**:
    *   Verify `apis/<service>/judgement_queue.yaml` is committed and contains `reason: untriaged-bulk-generation` with `status: open` for each newly prepopulated `<Kind>`.
*   **No Premature Triage in Step 1**:
    *   Reject the PR if any `judgement_queue.yaml` entries for the newly scaffolded `<Kind>` were manually marked `status: resolved` or deleted during Step 1. Triage belongs strictly in Step 2 (`judgement/SKILL.md`).

---

## 3. Pure Mechanical Output (No Manual Edits to Generated KRM Files)

*   **Zero Diff After Re-Running `./apis/<service>/generate.sh`**:
    *   Run `./apis/<service>/generate.sh` and `git diff --exit-code` to verify that:
        *   `apis/<service>/v1alpha1/<kind>_types.go`
        *   `apis/<service>/v1alpha1/types.generated.go`
        *   `apis/<service>/v1alpha1/<kind>_identity.generated.go`
        *   `apis/<service>/v1alpha1/<kind>_reference.generated.go`
        *   `pkg/controller/direct/<service>/mapper.generated.go`
        *   `config/crds/resources/apiextensions.k8s.io_v1_customresourcedefinition_<plural>.<group>.cnrm.cloud.google.com.yaml`
        are 100% unmodified generator output.
    *   Reject any hand-edits to `<kind>_types.go`, `<kind>_identity.generated.go`, or `<kind>_reference.generated.go` in Step 1.
*   **Allowed Compilation Fixes for `pkg/controller/direct/<service>/`**:
    *   Minimal changes required for `pkg/controller/direct/<service>/mapper.generated.go` to compile ARE allowed in Step 1:
        *   Adding the service's Go client library module to `go.mod` / `go.sum`.
        *   Adding external proto packages (e.g., `,google.type`) to `generate-mapper --service` in `apis/<service>/generate.sh`.
        *   Minimal stub/hand-written mapper helper functions in `pkg/controller/direct/<service>/` (e.g., for `*Request` / `*Metadata` messages skipped by `generate-mapper`) or updating existing `v1beta1` controller callers when enabling `--multiversion`.

---

## 4. Release Bundle Blocklist, Compilation, & `apichecks` Verification

*   **`BLOCKED_KINDS` in `dev/tasks/build-release-bundle`**:
    *   Verify that each newly prepopulated `<Kind>` is added (formatted with 4-space indentation and double quotes as `    "<Kind>"`, sorted alphabetically into the existing list without duplicates) to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle` so unverified Greenfield CRDs are excluded from `operator/channels/packages/configconnector/*/crds.yaml` until graduated at the end of the pipeline.
*   **No Individual CI Presubmit Files**:
    *   Verify that no new individual CI presubmit scripts were created under `dev/ci/presubmits/` (e.g., `tests-e2e-fixtures-<kind>`). E2E fixture testing in CI is handled by existing matrix suite runners.
*   **Build, `apichecks`, `envtest`, & Operator Golden Verification**:
    *   Run:
        ```bash
        go build ./apis/<service>/... ./pkg/controller/direct/<service>/...
        go test -v ./apis/<service>/... ./pkg/controller/direct/<service>/...
        go test -v ./tests/apichecks/...
        . dev/tasks/setup-envtest && go test ./pkg/cli/preview/... -run TestPreview
        go test ./operator/pkg/controllers/configconnector/...
        ```
*   **Exception Files**:
    *   In Step 1 (`--prepopulate-resource`), running `WRITE_GOLDEN_OUTPUT=1 go test ./tests/apichecks/...` may legitimately record untriaged items in `alpha-missingfields.txt`, `refs_not_representable.txt`, `spec_dislike_etag.txt`, `sensitive.txt`, `comments.txt`, and `recursivetypes.txt`, or add genuine parameterless proto `oneof` variant structs to `knownInvalidCRDs` in `tests/apichecks/crds_test.go` (annotated with `// GCP known pattern: Empty oneof option (<field path>)`).
    *   Strictly reject any modifications to `tests/apichecks/testdata/exceptions/missingrefs.txt`, `naming_violations.txt`, or `tests/apichecks/testdata/exceptions/multi_version_crd_diff/`.

---

# Review Comment Template

When proposing changes or stating LGTM, format the review description as follows:

```markdown
### KCC Greenfield Step 1 (Prepopulate) Auto-Review Results
* **Trigger criteria matched**: [Yes/No]
* **`generate.sh` & Flags Check**: [Pass/Fail] - (Verify `generate-types --prepopulate-resource`, `generate-identity`, `generate-mapper --prepopulate-resource`, and no `--allow-non-ga-proto`)
* **Judgement Queue Open Check**: [Pass/Fail] - (Confirm `apis/<service>/judgement_queue.yaml` has `untriaged-bulk-generation` with `status: open` and no premature manual resolutions)
* **Mechanical Output & Idempotency**: [Pass/Fail] - (Confirm `./apis/<service>/generate.sh` produces zero diff and `<kind>_types.go` was not hand-edited)
* **Release Bundle Blocklist (`BLOCKED_KINDS`)**: [Pass/Fail] - (Confirm `<Kind>` is added in alphabetical order to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle`)
* **No Individual CI Files Check**: [Pass/Fail] - (Confirm no individual `dev/ci/presubmits/tests-e2e-fixtures-*` script was created)
* **Mapper Compilation & `apichecks`**: [Pass/Fail] - (Confirm `apis/<service>/...` and `pkg/controller/direct/<service>/...` build and `tests/apichecks/...` passes without unauthorized exception edits)

#### Detailed Findings / Actions Required:
1. [Specify file, line number, and exact issue]
```
