---
name: kcc-bulk-controller-generator-generate
description: Phase 1 of the KCC Bulk Controller & Fuzzer Generator workflow. Adds `generate-controller --prepopulate-resource` and `generate-fuzzer --prepopulate-resource` to `apis/<service>/generate.sh`, runs batch generation, tests compilation and `FuzzAllMappers`, and splits services into Category 1/2/3 branches.
---

# KCC Bulk Controller & Fuzzer Generator — Phase 1: Generation & Category Splitting

Follow this guide to bulk-generate deterministic direct controllers (`<proto_lower>_controller.go`) and round-trip fuzzers (`<kind_lower>_fuzzer.generated.go`) across multiple greenfield resources, either **post-judgement** (after Step 1b) or **pre-judgement** (immediately after Step 1a).

---

## Step 1: Identify Candidate Greenfield Resources

For each target service (`apis/<service>/`):
1. Inspect `apis/<service>/generate.sh` and `apis/<service>/judgement_queue.yaml`.
   - **Mode A (Post-Judgement):** Select `--resource <Kind>:<ProtoMessage>` entries whose `untriaged-bulk-generation` entry in `apis/<service>/judgement_queue.yaml` is `status: resolved` (or finalized greenfield types without an open `untriaged-bulk-generation` entry).
   - **Mode B (Pre-Judgement):** Select `--resource <Kind>:<ProtoMessage>` entries that were prepopulated in Step 1a (`untriaged-bulk-generation` is `status: open`).
2. **Existing Controller & Fuzzer Protection (Automatic + Pre-Check)**:
   - Both `generate-controller --prepopulate-resource` (without `--overwrite`) and `generate-fuzzer --prepopulate-resource` automatically skip any resource that already has a handwritten controller (`RegisterModel` for `<Kind>GVK` without `// +generated:controller`) or handwritten fuzzer.
   - However, only pass the target greenfield `--resource <Kind>:<ProtoMessage>` flags to `generate-controller` and `generate-fuzzer` in `apis/<service>/generate.sh`.

---

## Step 2: Update `apis/<service>/generate.sh`

In `apis/<service>/generate.sh`, immediately after `${CONTROLLERBUILDER} generate-mapper ...`, add `generate-controller` and `generate-fuzzer` invocations with `--prepopulate-resource`:

```bash
${CONTROLLERBUILDER} generate-controller \
  --service <proto.service.package> \
  --api-version <service>.cnrm.cloud.google.com/v1alpha1 \
  --prepopulate-resource \
  --resource <KindA>:<ProtoA> \
  --resource <KindB>:<ProtoB>

${CONTROLLERBUILDER} generate-fuzzer \
  --prepopulate-resource \
  --service <proto.service.package> \
  --api-version <service>.cnrm.cloud.google.com/v1alpha1 \
  --resource <KindA>:<ProtoA> \
  --resource <KindB>:<ProtoB>
```

- **Pinned Proto Descriptor Sets:** If `apis/<service>/generate.sh` passes `--proto-source-path ${REPO_ROOT}/.build/googleapis-<sha>.pb` to `generate-types`, pass the exact same `--proto-source-path` flag to `generate-controller` and `generate-fuzzer`.
- **Single Binary Compilation:** Build `bin/controllerbuilder` once at the start of the batch (`go build -o bin/controllerbuilder ./dev/tools/controllerbuilder`) so every `apis/<service>/generate.sh` reuses the compiled binary.

---

## Step 3: Run Generation, Compile, and Fuzz-Test Each Service

For each target service:
1. Run `SKIP_GENERATE_CRDS=1 ./apis/<service>/generate.sh`.
2. Check whether `pkg/controller/direct/<service>/...` compiles:
   ```bash
   go build ./pkg/controller/direct/<service>/...
   ```
   - If `go build` reports a missing `cloud.google.com/go/<service>/apiv1` package or missing GAPIC types, run `go get cloud.google.com/go/<service>@latest && go mod tidy` (classifying this service as **Category 2: Initial Setup**).
3. Verify the generated fuzzers pass `FuzzAllMappers` and naming checks pass:
   ```bash
   go test -v ./pkg/fuzztesting/fuzztests -run="FuzzAllMappers" -fuzztime=1s
   go test -v ./tests/apichecks -run TestDirectResourceFileNaming
   ```

---

## Step 4: Classify and Split into Category Branches

Split the modified services into up to 3 clean branches off `master`:

1. **Category 1 (`Zero-Touch`):**
   - Services where adding `generate-controller --prepopulate-resource` and `generate-fuzzer --prepopulate-resource` to `apis/<service>/generate.sh` generated `<proto_lower>_controller.go` and `<kind_lower>_fuzzer.generated.go`, compiled cleanly without `go.mod` changes, and passed `FuzzAllMappers`.
2. **Category 2 (`Initial Setup`):**
   - Services that required adding/updating GAPIC client dependencies in `go.mod` / `go.sum`, or adding `<Kind>Spec` / `<Kind>ObservedState` mappers via `generate-mapper --prepopulate-resource`.
3. **Category 3 (`Major Surgery`):**
   - Services where a resource has non-standard CRUD RPCs (e.g., singleton resources without `Create`/`Delete`, custom LRO metadata, or requiring a `<funcName>FilterSpec` / `<funcName>FilterStatus` hook for fuzzer edge cases).

Commit each category to its own branch and proceed to Phase 2 (`verify/SKILL.md`) using an isolated subagent.
