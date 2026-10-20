---
name: kcc-bulk-resource-generator-verify
description: Phase 2 verifier workflow for checking a bulk-generated greenfield branch in an isolated subagent context. Enforces reviewgen-greenfield-new-types/prepopulate/SKILL.md plus Category 1/2/3 boundary purity and batch idempotency.
---

# KCC Bulk Resource Generator — Phase 2: Isolated Branch Verification

Run this skill inside a **fresh, isolated subagent** to verify a bulk-generated branch (`Category 1`, `Category 2`, or `Category 3`) before opening or publishing the PR.

---

## 1. Baseline Prepopulate Correctness Checklist

First, read `.gemini/skills/reviewgen-greenfield-new-types/prepopulate/SKILL.md` (`IsSkillFile: true`) and enforce all of its review invariants across every service touched in the branch (`git diff --name-only upstream/master`):
1. **`apis/<service>/generate.sh` Flags, Command Ordering, & GA Protos:**
   - Every touched `apis/<service>/generate.sh` uses `${CONTROLLERBUILDER}` (not legacy `go run .`) and runs `generate-types --prepopulate-resource`, `generate-identity`, and `generate-mapper --prepopulate-resource` (plus `--multiversion` where `v1beta1` is present).
   - **Strict Command Ordering:** Verify that **all** `generate-types` and `generate-identity` invocations appear **before** `generate-mapper` in `apis/<service>/generate.sh` (followed by `cd ${REPO_ROOT}`, `dev/tasks/generate-crds`, and `goimports`). Reject any `generate.sh` where a `generate-types` or `generate-identity` block was appended after `generate-mapper`.
   - **Zero** occurrences of `--allow-non-ga-proto`.
   - All new resources are under `v1alpha1` with `2026` copyright headers.
2. **Judgement Queue Open State (`apis/<service>/judgement_queue.yaml`):**
   - Each newly scaffolded `<Kind>` has `reason: untriaged-bulk-generation` with `status: open`.
   - No `judgement_queue.yaml` entries for newly added resources were prematurely marked `status: resolved` or deleted.
3. **Release Bundle Blocklist (`BLOCKED_KINDS` in `dev/tasks/build-release-bundle`):**
   - Verify that every newly generated `<Kind>` in the branch is added (formatted with 4-space indentation and double quotes as `    "<Kind>"`, sorted alphabetically into the existing list without duplicates) to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle` so unverified Greenfield CRDs are excluded from `operator/channels/packages/configconnector/*/crds.yaml` until graduated at the end of the pipeline.
4. **Compilation, `tests/apichecks`, `envtest`, & Operator Goldens:**
   - Run `go build` and `go test` across all touched `apis/<service>/...` and `pkg/controller/direct/<service>/...` packages.
   - Run `go test ./tests/apichecks/...` and confirm `tests/apichecks/testdata/exceptions/missingrefs.txt`, `naming_violations.txt`, and `tests/apichecks/testdata/exceptions/multi_version_crd_diff/` are untouched (`WRITE_GOLDEN_OUTPUT=1` in Step 1 may legitimately record untriaged items in `alpha-missingfields.txt`, `refs_not_representable.txt`, `spec_dislike_etag.txt`, `sensitive.txt`, `comments.txt`, and `recursivetypes.txt`).
   - Run `. dev/tasks/setup-envtest && go test ./pkg/cli/preview/... -run TestPreview` and `go test ./operator/pkg/controllers/configconnector/...` to verify `kube-apiserver` accepts all new CRD OpenAPI schemas and operator RBAC goldens (`_expected.yaml`) are in sync.
   - If `TestDirectResourceFileNaming` fails on a pre-existing unprefixed stub `<resource>_reference.go` file after `<service><resource>_types.go` is added, verify it was renamed to `<service><resource>_reference.go` and that `var <Kind>GVK = GroupVersion.WithKind("<Kind>")` lives in `<service><resource>_types.go`.

---

## 2. Category Boundary & Purity Verification

Verify that the branch has no untracked files (`git status --porcelain` is empty) and strictly conforms to its declared PR category (`Category 1`, `Category 2`, or `Category 3`):

### If Verifying a **Category 1 (`Zero-Touch`)** Branch
- **No `go.mod` / `go.sum` Changes:**
  ```bash
  git diff --exit-code upstream/master -- go.mod go.sum
  ```
- **No New Service Boilerplate:**
  - Verify no new `apis/<service>/v1alpha1/doc.go` or `apis/<service>/v1alpha1/groupversion_info.go` files were added (`git diff --diff-filter=A --name-only upstream/master -- 'apis/**/doc.go' 'apis/**/groupversion_info.go'` must be empty).
- **No Handwritten Mapper or Controller Edits:**
  - Under `pkg/controller/direct/`, the **only** modified or added files may be `mapper.generated.go` (`git diff --name-only upstream/master -- 'pkg/controller/direct/**' ':!pkg/controller/direct/**/mapper.generated.go'` must be empty).

### If Verifying a **Category 2 (`Initial Setup`)** Branch
- **Every Service Requires Setup:**
  - Confirm each service in the branch either adds a new `apis/<service>/v1alpha1/{doc.go,groupversion_info.go}` package, requires a `go.mod`/`go.sum` dependency addition, or adds a shared proto package (such as `,google.type`) to `generate-mapper --service` in `apis/<service>/generate.sh`.
- **No Handwritten Mapper or Type Surgery:**
  - Under `pkg/controller/direct/`, the **only** modified or added files may be `mapper.generated.go`.
  - No manual edits to `<kind>_types.go`, `types.generated.go`, `<kind>_identity.generated.go`, or `<kind>_reference.generated.go`.

### If Verifying a **Category 3 (`Major Surgery`)** Branch
- **Explicit Surgery Inventory:**
  - List every non-mechanical change in the branch (e.g., handwritten `<kind>_mapper.go` helper files in `pkg/controller/direct/<service>/`, `--multiversion` controller caller updates, `TestCRDObjectTypes` empty `oneof` workarounds, recursive proto `+kubebuilder:validation:Schemaless` markers, or multi-stage `generate.sh` scripts).
  - Verify that if any workaround was needed for `TestCRDObjectTypes` (`field ... is an object with no properties` on a parameterless proto `oneof` variant), the CRD entry in `knownInvalidCRDs` in `tests/apichecks/crds_test.go` carries the comment `// GCP known pattern: Empty oneof option (<field path>)` and that the underlying proto struct is genuinely an empty `oneof` option (not a broken/truncated struct generation).

---

## 3. Fast Batch Idempotency Check

Instead of running `./apis/<service>/generate.sh` with full CRD regeneration per service (which takes several minutes per service), verify mechanical idempotency across the entire branch in one batch (`generate-proto.sh` automatically reuses `.build/googleapis-<SHA>.pb`):

```bash
(cd dev/tools/controllerbuilder && go build -o ../../../bin/controllerbuilder . && ./generate-proto.sh)

for script in $(git diff --name-only upstream/master -- 'apis/*/generate.sh'); do
  SKIP_GENERATE_CRDS=1 ./"${script}"
done

./dev/tasks/generate-crds
git diff --exit-code
```

- For **Category 1** and **Category 2** branches, `git diff --exit-code` **must** pass with zero diff.
- For **Category 3** branches, `git diff --exit-code` must pass unless a documented generator limitation required a tracked workaround in the surgery inventory above.

---

## 4. Verification Report Format

Return a structured summary to the parent orchestrator agent:

```markdown
### Bulk Resource Generator Verification Report (`<Branch>` — `<Category>`)
* **Prepopulate Checklist (`reviewgen-greenfield-new-types/prepopulate/SKILL.md`)**: [Pass/Fail]
* **Release Bundle Blocklist (`BLOCKED_KINDS` in `dev/tasks/build-release-bundle`)**: [Pass/Fail]
* **Category Boundary Purity (`<Category>`)**: [Pass/Fail]
* **Batch Idempotency (`generate.sh` + `generate-crds`)**: [Pass/Fail]
* **Compilation & `tests/apichecks`**: [Pass/Fail]

#### Findings / Required Fixes:
1. [File path and description of any violation, or "None — ready for PR"]
```
