---
name: kcc-direct-greenfield-types-judgement
description: Step 2 of the KCC Greenfield pipeline. Triages and resolves all `judgement_queue.yaml` entries for a bulk-prepopulated greenfield resource `<Kind>`, finalizes KRM types/identity/references/mappers/CRDs, and verifies `apichecks`.
---

# KCC Direct Greenfield Types Implementer — Step 2: Judgement Queue Triage Workflow

Use this workflow for **Greenfield Step 2** when `apis/<service_short>/<version>/<kind>_types.go` and `apis/<service_short>/judgement_queue.yaml` already exist (either from a Step 1 single-resource prepopulate PR or a multi-resource batch prepopulate PR).

## Prerequisites
You **must** also apply the baseline KRM standards in:
- [`.gemini/skills/kcc-direct-base-types-implementer/SKILL.md`](file:///usr/local/google/home/lmadariaga/github/k8s-config-connector/.gemini/skills/kcc-direct-base-types-implementer/SKILL.md)
- [`.gemini/skills/kcc-identity-reference/SKILL.md`](file:///usr/local/google/home/lmadariaga/github/k8s-config-connector/.gemini/skills/kcc-identity-reference/SKILL.md)

---

## Core Rules & Invariants of `judgement_queue.yaml`

1. **Canonical `(Kind, Field, Reason)` Grouping**:
   - `judgement_queue.yaml` groups entries contiguously by `(Group, Kind)` -> `untriaged-bulk-generation` -> Kind-level entries (`field: ""`) -> field-level entries grouped by `field`.
   - Work through `<Kind>`'s contiguous block **field-by-field** so you evaluate all stacked signals for a `(Kind, Field)` together in one pass.
2. **Never Delete Entries**:
   - Never delete an entry from `judgement_queue.yaml`. Transition `status: open` to `status: resolved` with `resolution` (`accepted`, `edited`, `deferred`, or `not-applicable`) and a clear `note` (required for `edited`, `deferred`, and `not-applicable`).
3. **Scope of Entries to Resolve**:
   - Resolve every entry with `kind: <Kind>` (plus any legacy `protoMessage:` entry reachable from `<Kind>Spec` or `<Kind>ObservedState` if present). Leave other Kinds untouched.
   - **Deferred Team-Policy Items**: Leave `well-known-proto-type` and `sensitive-field` as `status: open` (or mark `status: resolved`, `resolution: deferred` with a `note`).
4. **Matching Statuses for Multi-Signal Reference Fields**:
   - `TestJudgementQueueIsWellFormed` **fails** if reference entries (`authoritative-proto-reference`, `possible-reference*`) for the same `(kind, group, field)` disagree on `status`. Always resolve all contiguous reference signals for a field together.
5. **Remove `+kcc:guess` Markers from Handwritten Files**:
   - Remove resolved `// +kcc:guess=...` comments from `<kind>_types.go` and `<kind>_identity*.go` (and from any struct moved out of `types.generated.go`). Do not manually edit `types.generated.go`.
6. **Never Reward-Hack `TestMissingRefs`**:
   - **NEVER** delete, truncate, or alter proto doc comments (such as `"Format: projects/{project}/..."`) in `<kind>_types.go` or `types.generated.go` to silence `TestMissingRefs`, and **NEVER** add entries to `tests/apichecks/testdata/exceptions/missingrefs.txt`.
7. **Resolve `untriaged-bulk-generation` Last**:
   - Keep `untriaged-bulk-generation` `open` while iterating on edits, and mark it `resolved` only after all other actionable entries for `<Kind>` are resolved and Phase 3 is complete.

---

## Workflow

### Phase 1: Inventory, `generate-mapper` Check, & Proto Discovery
1. Read `apis/<service>/judgement_queue.yaml` (locating `<Kind>`'s contiguous block), `apis/<service>/generate.sh`, and `apis/<service>/v1alpha1/<kind>_types.go`.
2. Verify that `apis/<service>/generate.sh` runs in the canonical order:
   - `generate-types --prepopulate-resource`
   - `generate-identity`
   - `generate-mapper --prepopulate-resource` (plus `--multiversion` if the service has both `v1alpha1` and `v1beta1`, and `--proto-source-path ${PROTO_OUT}` if `generate.sh` pins a custom `PROTO_SHA`)
   - `generate-controller --prepopulate-resource` and `generate-fuzzer --prepopulate-resource` (if a pre-generated controller/fuzzer already exists for `<Kind>`)
   - `cd ${REPO_ROOT}`, `dev/tasks/generate-crds`, and `goimports`.
   *(Note for multi-service or legacy packages such as `dialogflow`, `iam`, `mapmanagement`, `networkconnectivity`, or `bigqueryanalyticshub`: if running `generate-mapper` in `generate.sh` would overwrite existing mockgcp/legacy or multi-service mappers in `pkg/controller/direct/<service>/mapper.generated.go`, keep `<Kind>`'s mapper functions isolated in `pkg/controller/direct/<service>/<kind_lower>_mapper.go` using the exact lowercase `<kind_lower>_` prefix required by `TestDirectResourceFileNaming`.)*
3. Locate the proto definition in `.build/third_party/googleapis/google/...` matching `+kcc:spec:proto=<full_proto_name>` (run `./dev/tools/controllerbuilder/generate-proto.sh` if `.build/third_party/googleapis` is missing). Inspect the resource message, `option (google.api.resource)`, and `Create<Resource>Request`.

### Phase 2: Resolve Queue Entries by `reason`

- **If `<Kind>` has Parent, Location, Root, or Identity reasons** (`location-or-parent-ref`, `parent-ref-guessed`, `parent-ref-not-modelled`, `parent-multiple-patterns`, `location-parent-unknown`, `root-*`, `identity-*`):
  - Read and follow **[`references/PARENT_AND_IDENTITY.md`](file:///usr/local/google/home/lmadariaga/github/k8s-config-connector/.gemini/skills/kcc-direct-greenfield-types-implementer/judgement/references/PARENT_AND_IDENTITY.md)** using `view_file`.
- **If `<Kind>` has Reference, Output-Only Comment, `required-not-enforced`, or Shared Nested Struct reasons** (`authoritative-proto-reference`, `possible-reference*`, `reference-not-representable`, `*-output-only-pattern-in-comment`, `required-not-enforced`):
  - Read and follow **[`references/REFERENCES_AND_NESTED_STRUCTS.md`](file:///usr/local/google/home/lmadariaga/github/k8s-config-connector/.gemini/skills/kcc-direct-greenfield-types-implementer/judgement/references/REFERENCES_AND_NESTED_STRUCTS.md)** using `view_file`.
- **For all other reasons, use this Quick-Resolution Table**:

| `reason` | Resolution Action |
| :--- | :--- |
| `ambiguous-proto-name` | Qualify `--resource <Kind>:<full.proto.Package.ProtoName>` in `apis/<service>/generate.sh` (or `service.yaml`), verify `+kcc:spec:proto=` in `<kind>_types.go`, re-run `generate.sh`, and mark `edited`. |
| `empty-observedstate` | Inspect `<Kind>Spec` against proto comments: if server-populated output-only fields landed in `Spec`, move them to `<Kind>ObservedState` (`edited`). If the resource genuinely has no output fields, comment out `ObservedState *<Kind>ObservedState` in `<Kind>Status` and `type <Kind>ObservedState struct` so `TestCRDObjectTypes` passes (`edited`). |
| `etag-in-spec` | Remove `Etag *string` from `<Kind>Spec` (or the nested Spec struct, isolating it per `references/REFERENCES_AND_NESTED_STRUCTS.md` if shared) and ensure `Etag *string` is in `ObservedState`. Mark `edited`. |
| `possible-status-field-in-spec` | Check the proto comment and RPC semantics for the flagged optional field (`state`, `status`, `createTime`, `updateTime`, `deleteTime`, `lastUpdateTime`, `creationTimestamp`, `startTime`, `endTime`, `uid`, `id`, `selfLink`, `selfLinkWithID`, `reconciling`): if it is server-populated output state, move it from `<Kind>Spec` to `<Kind>ObservedState` (`edited`); if `id`/`uid` duplicates `spec.resourceID` or is server-generated, remove it from `<Kind>Spec` (and place in `<Kind>ObservedState` if server-generated, `edited`); if it is genuinely user-configurable desired state (e.g., an admin state toggle or scheduled start/end window), keep it in `<Kind>Spec` (`accepted`). |
| `server-set-field-placed` | Confirm against the proto that the field belongs in `ObservedState`, remove any `+kcc:guess` comment in `<kind>_types.go`, and mark `accepted` (or `edited` if moved back to `Spec`). |
| `message-map-derived` | Verify the generated `map[string]<Struct>` field and its mapper helpers in `pkg/controller/direct/<service>/mapper.generated.go` compile cleanly. Mark `accepted`. |
| `deprecated-field` / `deprecated-message` / `deprecated-resource` | Top-level deprecated fields are already omitted (`accepted`). For `deprecated-message`, remove the deprecated field from `Spec` (see `references/REFERENCES_AND_NESTED_STRUCTS.md` if inside a nested struct) and mark `edited`. For `deprecated-resource`, confirm intent and mark `accepted`. |
| `verify-resource-docs-link` / `verify-service-docs-link` | Update the doc URL in `<kind>_types.go` to a valid `https://cloud.google.com/...` URL (or remove if no public page exists), remove `// +kcc:guess=source-link`, and mark `edited`. |
| `unsupported-field-type` | Check the `// TODO: unsupported` comment; implement a standard KRM mapping in `<kind>_types.go` (`edited`) or mark `deferred` with a note. |
| `well-known-proto-type` / `sensitive-field` | Leave `status: open` (or mark `status: resolved`, `resolution: deferred` with a note) pending team-wide greenfield conventions. |

### Phase 3: Holistic Proto-vs-KRM Audit (Before Closing `untriaged-bulk-generation`)
1. **Unflagged Subtle References**: Scan every remaining `*string` / `[]string` field in `<Kind>Spec` and reachable structs in `types.generated.go` for unflagged GCP resource references (e.g., `sourceProjectID`, `defaultDatabase`, `kmsKey`, `serviceAccount`, `network`, `repository`) and convert them to `*Ref` fields.
2. **Top-Level `Spec.Labels` Removal (`TestVerifyLabelsTestCasesExist`)**: Ensure top-level `Labels map[string]string` is removed from `<Kind>Spec` (top-level proto `labels` are reconciled via `metadata.labels`); leaving `Labels` in `<Kind>Spec` causes `TestVerifyLabelsTestCasesExist` to fail when a controller is registered.
3. **Nested `Etag` / Output-Only Fields & Empty `ObservedState`**: Check nested Spec structs for unannotated `Etag` or output-only fields; if `<Kind>ObservedState` ends up empty, comment out `ObservedState` in `<Kind>Status` and `type <Kind>ObservedState struct`.
4. **Close `untriaged-bulk-generation`**: Mark `<Kind>`'s `untriaged-bulk-generation` entry `status: resolved` (`edited` or `accepted` with a summary `note`).

### Phase 4: Regeneration, Lockstep Controller/Fuzzer Sync, & Verification
1. Run `./apis/<service>/generate.sh` and confirm `git diff apis/<service>/judgement_queue.yaml` shows no newly opened entries for `<Kind>`.
2. **If a Pre-Generated Controller (`*_controller.go`) or Fuzzer (`<kind_lower>_fuzzer.generated.go`) Exists in `pkg/controller/direct/<service>/` (MANDATORY)**:
   - Because closing `judgement_queue.yaml` entries changes KRM fields, `*Ref` types, `ObservedState` presence, or `IdentityV2` vs `ServerGeneratedIdentity`, you **must** regenerate and fix the controller, mapper, and fuzzer in lockstep:
     - Ensure `apis/<service>/generate.sh` (or direct `controllerbuilder generate-controller --prepopulate-resource` / `generate-fuzzer --prepopulate-resource` calls) updates `pkg/controller/direct/<service>/<proto_lower>_controller.go` and `pkg/controller/direct/<service>/<kind_lower>_fuzzer.generated.go`.
     - If `ambiguous-proto-name` changed `<ProtoMessage>`, remove any stale `<old_proto_lower>_controller.go` file and run `python3 dev/tasks/generate_static_config.py`.
     - Verify that `pkg/controller/direct/<service>/...` compiles cleanly and passes 200-seed round-trip fuzzing:
       ```bash
       go build ./pkg/controller/direct/<service>/...
       FOCUS=<Kind> go test -v ./pkg/fuzztesting/fuzztests -run TestFocusedMappers
       ```
3. Confirm `<Kind>` is present in `BLOCKED_KINDS` in `dev/tasks/build-release-bundle` (sorted alphabetically; if missing, add it—do **not** remove `<Kind>` from `BLOCKED_KINDS` during Step 2, as a separate graduation step removes it at the end of the pipeline).
4. If `go.mod` was updated, run `go mod tidy && (cd mockgcp && go mod tidy)` (or `make ensure`) so `validate-ensure` passes.
5. Run verification tests (including `TestCRDBackwardCompatibility`, `TestHandWrittenProtoStructAnnotationsAndNaming`, `TestGenerateScripts`, `TestDirectResourceFileNaming`, and `TestRegisteredTemplatesMatchCAI`):
   ```bash
   go test -v ./tests/apichecks/... -run "TestMissingRefs|TestJudgementQueue|TestSourceLinks|TestCRDObjectTypes|TestCRDBackwardCompatibility|TestHandWrittenProtoStructAnnotationsAndNaming|TestGenerateScripts|TestDirectResourceFileNaming"
   go test -v ./pkg/gcpurls/... ./pkg/test/resourcefixture/... -run "TestRegisteredTemplatesMatchCAI|TestVerifyLabelsTestCasesExist"
   go test -v ./apis/<service>/... ./pkg/controller/direct/<service>/...
   ```
6. Record any service-specific proto or reference quirks in `.gemini/journals/<service>.md`.

