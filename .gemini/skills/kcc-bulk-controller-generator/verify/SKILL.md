---
name: kcc-bulk-controller-generator-verify
description: Phase 2 of the KCC Bulk Controller & Fuzzer Generator workflow. Executed by an isolated verifier subagent to audit a bulk-generated controller & fuzzer branch against KCC invariants.
---

# KCC Bulk Controller & Fuzzer Generator — Phase 2: Isolated Verification

Run this checklist in an isolated subagent context on a committed bulk-controller branch (`Category 1`, `Category 2`, or `Category 3`).

---

## Verification Checklist

### 1. Existing Handwritten Controllers & Fuzzers Are Untouched
Run `git status` and `git diff origin/master...HEAD --name-status`:
- Verify that **every** newly created or modified `*_controller.go` file in the diff contains the `// +generated:controller` header marker.
- Verify that **every** newly created or modified `*_fuzzer*.go` file in the diff is named `<kind_lower>_fuzzer.generated.go` and contains the `// +generated:fuzzer` header marker.
- Strictly fail verification if any pre-existing handwritten `*_controller.go` or `*_fuzzer.go` file was overwritten or deleted.

### 2. Canonical Greenfield Controller Standards (`// +generated:controller`)
Inspect every generated `pkg/controller/direct/<service>/<proto_lower>_controller.go` in the branch diff:
1. **`// +generated:controller` Marker:** Present at the top of the file.
2. **GAPIC Client Construction:** Uses `gcp.New<Service>RESTClient(ctx, opts...)` with `m.config.RESTClientOptions()` (or `New<Service>Client` with `m.config.GRPCClientOptions()` only when the GAPIC package has no REST constructor).
3. **`common.NormalizeGreenfieldReferences`:** Called at the start of `AdapterForObject` (`common.NormalizeGreenfieldReferences(ctx, reader, obj)`).
4. **Declared-Spec-Only Diffing (`common.CompareGreenfieldSpec` / `common.CompareGreenfieldSpecAndLabels`):** `Update` calls `common.CompareGreenfieldSpec(ctx, a.desired, actual, krm.<Kind>Spec_FromProto)` (or `common.CompareGreenfieldSpecAndLabels(..., a.labels)` when top-level proto `labels` are reconciled via `metadata.labels`) and reports diffs via `structuredreporting.ReportDiff(ctx, report)`.
5. **Status & ExternalRef Updates:** `setStatus(u, resolved)` is called at the end of `Create` and on both no-op and mutating `Update` paths, setting `status.ExternalRef = direct.LazyPtr(a.id.String())`.

### 3. `generate.sh` Idempotency & `static_config.go` Sync
- Re-run `SKIP_GENERATE_CRDS=1 ./apis/<service>/generate.sh` for the services in the branch and verify `git diff` is completely empty (confirming `// +generated:controller` and `// +generated:fuzzer` files regenerate idempotently in place).
- Re-run `python3 dev/tasks/generate_static_config.py` and verify `git diff pkg/controller/resourceconfig/static_config.go` is empty.
- If `go.mod` was modified, verify `go mod tidy && (cd mockgcp && go mod tidy)` produces zero diff.

### 4. Build, 200-Seed Fuzzer Round-Trip, & Repo Unit Tests
Run the following commands and verify they pass with zero errors:
```bash
go build ./pkg/controller/direct/...
go test -v ./pkg/fuzztesting/fuzztests -run="FuzzAllMappers"
go test -v ./tests/apichecks -run "TestGenerateScripts|TestDirectResourceFileNaming"
go test -v ./pkg/gcpurls/... ./pkg/test/resourcefixture/... -run "TestRegisteredTemplatesMatchCAI|TestVerifyLabelsTestCasesExist"
```
*(Ensure every new fuzzer is tested across at least 200 seeds via `FuzzSpec` and `FuzzStatus`.)*


---

## Output Format: Structured Verification Report

Return a concise report with:
- **Branch & Category:** `<branch> (<Category 1|2|3>)`
- **Services & Kinds Count:** `<N services, M controllers, K fuzzers>`
- **Handwritten Protection Check:** `PASS / FAIL`
- **5 Canonical Greenfield Standards Check:** `PASS / FAIL`
- **`generate.sh` Idempotency Check:** `PASS / FAIL`
- **`go build`, `FuzzAllMappers`, & `TestDirectResourceFileNaming`:** `PASS / FAIL`
- **Actionable Findings (if any):** Exact file paths and fixes needed.
