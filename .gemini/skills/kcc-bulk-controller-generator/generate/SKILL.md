---
name: kcc-bulk-controller-generator-generate
description: Phase 1 of the KCC Bulk Controller & Fuzzer Generator workflow. Adds `generate-controller --prepopulate-resource` and `generate-fuzzer --prepopulate-resource` to `apis/<service>/generate.sh`, runs batch generation, tests compilation and `FuzzAllMappers`, and splits services into Category 1/2/3 branches.
---

# KCC Bulk Controller & Fuzzer Generator — Phase 1: Generation & Category Splitting

Follow this guide to bulk-generate deterministic direct controllers (`<proto_lower>_controller.go`) and round-trip fuzzers (`<kind_lower>_fuzzer.generated.go`) across multiple greenfield resources. **Post-judgement (after Step 1b) is the standard mode**; pre-judgement generation (immediately after Step 1a) must only be used in ad-hoc bulk campaigns, never in systemic per-resource workflows.

---

## Step 1: Identify Candidate Greenfield Resources

For each target service (`apis/<service>/`):
1. Inspect `apis/<service>/generate.sh` and `apis/<service>/judgement_queue.yaml`.
   - **Standard Mode (Post-Judgement):** Select `--resource <Kind>:<ProtoMessage>` entries whose `untriaged-bulk-generation` entry in `apis/<service>/judgement_queue.yaml` is `status: resolved` (or finalized greenfield types without an open `untriaged-bulk-generation` entry).
   - **Ad-Hoc Campaign Mode Only (Pre-Judgement):** Select `--resource <Kind>:<ProtoMessage>` entries that were prepopulated in Step 1a (`untriaged-bulk-generation` is `status: open`).
2. **Existing Controller & Fuzzer Protection (Automatic + Pre-Check)**:
   - Both `generate-controller --prepopulate-resource` (without `--overwrite`) and `generate-fuzzer --prepopulate-resource` automatically skip any resource that already has a handwritten controller (`RegisterModel` for `<Kind>GVK` without `// +generated:controller`) or handwritten fuzzer.
   - However, only pass the target greenfield `--resource <Kind>:<ProtoMessage>` flags to `generate-controller` and `generate-fuzzer` in `apis/<service>/generate.sh`.
   - **AIP-122 Resource Check:** Ensure each target `<ProtoMessage>` declares a `(google.api.resource)` option in the service's descriptor set and is not a legacy `tf2crd`/`dcl2crd` CRD; non-AIP-122 messages without `(google.api.resource)` (such as `DataprocJob`) require a handwritten controller.

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

- **Pinned Proto Descriptor Sets (`PROTO_SHA` / `PROTO_OUT`):** If `apis/<service>/generate.sh` passes `--proto-source-path ${REPO_ROOT}/.build/googleapis-<sha>.pb` (e.g., `ces`, `mapmanagement`, `bigqueryanalyticshub`), pass the exact same `--proto-source-path` flag to `generate-mapper`, `generate-controller`, and `generate-fuzzer`, and restore `.build/googleapis.pb` to the repo-pinned SHA in `apis/git.versions` if a service script checks out a custom `googleapis` commit. *(Note: if the GAPIC Go package in `go.mod` is newer than the pinned `PROTO_SHA` and adds new fields to the Go proto struct, pass the current `.build/googleapis-<latest>.pb` to `generate-fuzzer` so all Go proto fields are covered by the fuzzer.)*
- **Single Binary Compilation:** Build `bin/controllerbuilder` once at the start of the batch (`go build -o bin/controllerbuilder ./dev/tools/controllerbuilder`) so every `apis/<service>/generate.sh` reuses the compiled binary.

---

## Step 3: Run Generation, Handle Complex/High-Touch Patterns, and Verify

For each target service:
1. Run `SKIP_GENERATE_CRDS=1 ./apis/<service>/generate.sh`.
2. **Update Static Config (`static_config.go` — MANDATORY):**
   Every newly registered controller (`registry.RegisterModel`) must be added to `pkg/controller/resourceconfig/static_config.go` or `validate-generated-files` will fail:
   ```bash
   python3 dev/tasks/generate_static_config.py
   ```
3. **Handle Complex / High-Touch Resource Patterns:**
   - **Pattern 1 — Missing GAPIC Modules or Multi-Proto-Service Packages (`iam/apiv3`, `dialogflow` spanning `v2` + `cx/v3`, `mapmanagement/apiv2`):**
     - If `go build ./pkg/controller/direct/<service>/...` reports a missing `cloud.google.com/go/<service>/apiv...` package, run `go get cloud.google.com/go/<service>@latest` and then sync **both** `go.mod` and `mockgcp/go.mod`:
       ```bash
       go mod tidy && (cd mockgcp && go mod tidy)
       ```
       *(Or run `./dev/tasks/ensure` so `validate-ensure` passes with zero diff.)*
   - **Pattern 2 — Shared KRM Struct, Multi-Version, or Legacy `mapper.generated.go` Collisions (`dialogflow`, `iam`, `mapmanagement`, `networkconnectivity`, `bigqueryanalyticshub`, `securesourcemanager`):**
     - If `apis/<service>/` has both `v1alpha1` and `v1beta1` types, check whether `--multiversion` on `generate-mapper` resolves duplicate helper names.
     - If `pkg/controller/direct/<service>/mapper.generated.go` is already owned by mockgcp-backed or multi-service resources (so running `generate-mapper` overwrites existing mappers and fails `validate-manifests`, or omitting `generate-mapper` from `generate.sh` fails `TestGenerateScripts`), extract the new resource's mapper functions into `pkg/controller/direct/<service>/<kind_lower>_mapper.go` (without the `// +generated:mapper` header, using the exact lowercase `<kind_lower>_` prefix required by `TestDirectResourceFileNaming`) and keep the existing `mapper.generated.go` intact.
     - If multiple resources in the same `apis/<service>/v1alpha1` package share a nested proto message name with conflicting KRM field edits (as in `securesourcemanager`), isolate the conflicting nested struct in `<kind_lower>_types.go` before generating mappers.
   - **Pattern 3 — Custom LROs & Immutable Resources Without `Update*` RPCs (`DialogflowConversationModel`, `NetworkConnectivityServiceConnectionToken`):**
     - For immutable resources that have `Create*`, `Get*`, and `Delete*` RPCs but no `Update*` RPC, verify that the generated `Update` method performs a no-op status refresh (`a.setStatus(u, actual)`) without calling an `Update*` RPC, and inspect any `// TODO:` comments in `<proto_lower>_controller.go`.
   - **Pattern 4 — Top-Level `Spec.Labels` vs `metadata.labels` (`TestVerifyLabelsTestCasesExist`):**
     - Registering a controller in `pkg/controller/direct/register/register.go` activates `TestVerifyLabelsTestCasesExist`. If a resource still has top-level `Labels map[string]string` in `<Kind>Spec` (from pre-judgement Step 1a), `TestVerifyLabelsTestCasesExist` will fail unless `Spec.Labels` is removed (reconciled via `metadata.labels` and `CompareGreenfieldSpecAndLabels`) or a labels test fixture exists.
   - **Pattern 5 — `TestRegisteredTemplatesMatchCAI` (`pkg/gcpurls/registry_test.go`):**
     - Importing `_ "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/<service>"` in `register.go` links `apis/<service>/` into the main binary, activating `TestRegisteredTemplatesMatchCAI` for all `gcpurls.Register` templates in that package. If a valid GCP resource URL template is absent from `cloudassetinventory_names.jsonl`, add it to `ignoredTemplates` in `pkg/gcpurls/registry_test.go`.
   - **Pattern 6 — MockGCP Discovery Proto vs GAPIC `protoregistry.GlobalTypes` Collision (`mockgcp/common/httpmux/json.go`):**
     - Importing a GAPIC protobuf package (such as `cloud.google.com/go/networkconnectivity/apiv1/networkconnectivitypb`) registers `google.cloud.<service>.*` in `protoregistry.GlobalTypes`. If `mockgcp/mock<service>` uses Discovery-generated protos (`mockgcp.cloud.<service>.*`, where field numbers are ordered alphabetically rather than matching `.proto` definitions) inside `google.protobuf.Any` (e.g., LRO `OperationMetadata`), `mockgcp/common/httpmux/json.go`'s `FindMessageByURL` will pick the `google.` descriptor first and fail with `500 Internal Server Error: proto: google.protobuf.Any: unable to unmarshal ... cannot parse invalid wire-format data`. Add the service prefix exception in `mockgcp/common/httpmux/json.go` (matching `mockgcp/common/httptogrpc/proto_resolver.go`) so `mockgcp.` descriptors are used.
4. **Verify Compilation, 200-Seed Fuzzer Round-Trip, and Repo Checks:**
   - Run **at least 200 seeds locally** on every generated fuzzer (`FuzzSpec` + `FuzzStatus`) so intermittent 1-in-$N$ seed failures are caught before CI's 8 `fuzz-roundtrippers` shards run:
     ```bash
     go build ./pkg/controller/direct/...
     FOCUS=<Kind> go test -v ./pkg/fuzztesting/fuzztests -run TestFocusedMappers
     go test -v ./tests/apichecks -run "TestGenerateScripts|TestDirectResourceFileNaming"
     go test -v ./pkg/gcpurls/... ./pkg/test/resourcefixture/... -run "TestRegisteredTemplatesMatchCAI|TestVerifyLabelsTestCasesExist"
     ```
   - **Consuming Fuzzer Diagnostic Hints (`// HINT:` & `<hint_for_agent>`):** If a fuzzer fails on any seed, inspect both the inline `// HINT:` comments in `pkg/controller/direct/<service>/<kind_lower>_fuzzer.generated.go` and the `<hint_for_agent>...</hint_for_agent>` block in the test output.
   - Run `make fmt` so all generated files pass `validate-fmt` (`gofmt`, `goimports`, and license headers).

---

## Step 4: Classify and Split into Category Branches

Split the modified services into up to 3 clean branches off `master`:

1. **Category 1 (`Zero-Touch`):**
   - Services where adding `generate-controller --prepopulate-resource` and `generate-fuzzer --prepopulate-resource` to `apis/<service>/generate.sh` generated `<proto_lower>_controller.go` and `<kind_lower>_fuzzer.generated.go`, compiled cleanly without `go.mod` changes, and passed 200-seed fuzzing.
2. **Category 2 (`Initial Setup`):**
   - Services that required adding/updating GAPIC client dependencies in `go.mod` / `go.sum` (synced with `mockgcp/go.mod`), or adding `<Kind>Spec` / `<Kind>ObservedState` mappers via `generate-mapper --prepopulate-resource`.
3. **Category 3 (`Major Surgery`):**
   - Higher-touch services requiring isolated `<kind_lower>_mapper.go` files for multi-service/legacy packages, custom LRO/immutable RPC handling, or `FilterSpec` / `FilterStatus` fuzzer hooks.

Commit each category to its own branch and proceed to Phase 2 (`verify/SKILL.md`) using an isolated subagent.

