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
   - `TestJudgementQueueIsWellFormed` **fails** if reference entries (`possible-reference*`) for the same `(kind, group, field)` disagree on `status`. Always resolve all contiguous `possible-reference*` signals for a field together.
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
2. Verify that `apis/<service>/generate.sh` includes `${CONTROLLERBUILDER} generate-mapper` with `--prepopulate-resource` (and `--multiversion` if the service has both `v1alpha1` and `v1beta1`). If the resource came from an earlier batch PR that omitted `generate-mapper`, add `generate-mapper --prepopulate-resource` to `apis/<service>/generate.sh` and any needed `go.mod` or stub mapper helpers so `pkg/controller/direct/<service>/...` compiles cleanly.
3. Locate the proto definition in `.build/third_party/googleapis/google/...` matching `+kcc:spec:proto=<full_proto_name>` (run `./dev/tools/controllerbuilder/generate-proto.sh` if `.build/third_party/googleapis` is missing). Inspect the resource message, `option (google.api.resource)`, and `Create<Resource>Request`.

### Phase 2: Resolve Queue Entries by `reason`

- **If `<Kind>` has Parent, Location, Root, or Identity reasons** (`location-or-parent-ref`, `parent-ref-guessed`, `parent-ref-not-modelled`, `parent-multiple-patterns`, `location-parent-unknown`, `root-*`, `identity-*`):
  - Read and follow **[`references/PARENT_AND_IDENTITY.md`](file:///usr/local/google/home/lmadariaga/github/k8s-config-connector/.gemini/skills/kcc-direct-greenfield-types-implementer/judgement/references/PARENT_AND_IDENTITY.md)** using `view_file`.
- **If `<Kind>` has Reference, Output-Only Comment, `required-not-enforced`, or Shared Nested Struct reasons** (`possible-reference*`, `reference-not-representable`, `*-output-only-pattern-in-comment`, `required-not-enforced`):
  - Read and follow **[`references/REFERENCES_AND_NESTED_STRUCTS.md`](file:///usr/local/google/home/lmadariaga/github/k8s-config-connector/.gemini/skills/kcc-direct-greenfield-types-implementer/judgement/references/REFERENCES_AND_NESTED_STRUCTS.md)** using `view_file`.
- **For all other reasons, use this Quick-Resolution Table**:

| `reason` | Resolution Action |
| :--- | :--- |
| `ambiguous-proto-name` | Qualify `--resource <Kind>:<full.proto.Package.ProtoName>` in `apis/<service>/generate.sh` (or `service.yaml`), verify `+kcc:spec:proto=` in `<kind>_types.go`, re-run `generate.sh`, and mark `edited`. |
| `empty-observedstate` | Inspect `<Kind>Spec` against proto comments: if server-populated output-only fields landed in `Spec`, move them to `<Kind>ObservedState` (`edited`). If the resource genuinely has no output fields, comment out `ObservedState *<Kind>ObservedState` in `<Kind>Status` and `type <Kind>ObservedState struct` so `TestCRDObjectTypes` passes (`edited`). |
| `etag-in-spec` | Remove `Etag *string` from `<Kind>Spec` (or the nested Spec struct, isolating it per `references/REFERENCES_AND_NESTED_STRUCTS.md` if shared) and ensure `Etag *string` is in `ObservedState`. Mark `edited`. |
| `server-set-field-placed` / `observedstate-identity-field-omitted` | Confirm against the proto that the field belongs in `ObservedState` (or `status.externalRef` for `name`), remove any `+kcc:guess` comment in `<kind>_types.go`, and mark `accepted` (or `edited` if moved back to `Spec`). |
| `message-map-derived` | Verify the generated `map[string]<Struct>` field and its mapper helpers in `pkg/controller/direct/<service>/mapper.generated.go` compile cleanly. Mark `accepted`. |
| `deprecated-field` / `deprecated-message` / `deprecated-resource` | Top-level deprecated fields are already omitted (`accepted`). For `deprecated-message`, remove the deprecated field from `Spec` (see `references/REFERENCES_AND_NESTED_STRUCTS.md` if inside a nested struct) and mark `edited`. For `deprecated-resource`, confirm intent and mark `accepted`. |
| `verify-resource-docs-link` / `verify-service-docs-link` | Update the doc URL in `<kind>_types.go` to a valid `https://cloud.google.com/...` URL (or remove if no public page exists), remove `// +kcc:guess=source-link`, and mark `edited`. |
| `unsupported-field-type` | Check the `// TODO: unsupported` comment; implement a standard KRM mapping in `<kind>_types.go` (`edited`) or mark `deferred` with a note. |
| `well-known-proto-type` / `sensitive-field` | Leave `status: open` (or mark `status: resolved`, `resolution: deferred` with a note) pending team-wide greenfield conventions. |

### Phase 3: Holistic Proto-vs-KRM Audit (Before Closing `untriaged-bulk-generation`)
1. **Unflagged Subtle References**: Scan every remaining `*string` / `[]string` field in `<Kind>Spec` and reachable structs in `types.generated.go` for unflagged GCP resource references (e.g., `sourceProjectID`, `defaultDatabase`, `kmsKey`, `serviceAccount`, `network`, `repository`) and convert them to `*Ref` fields.
2. **Nested `Etag` / Output-Only Fields & Empty `ObservedState`**: Check nested Spec structs for unannotated `Etag` or output-only fields; if `<Kind>ObservedState` ends up empty, comment out `ObservedState` in `<Kind>Status` and `type <Kind>ObservedState struct`.
3. **Close `untriaged-bulk-generation`**: Mark `<Kind>`'s `untriaged-bulk-generation` entry `status: resolved` (`edited` or `accepted` with a summary `note`).

### Phase 4: Regeneration & Verification
1. Run `./apis/<service>/generate.sh` and confirm `git diff apis/<service>/judgement_queue.yaml` shows no newly opened entries for `<Kind>`.
2. Run verification tests (including `TestCRDBackwardCompatibility` and `TestHandWrittenProtoStructAnnotationsAndNaming`):
   ```bash
   go test -v ./tests/apichecks/... -run "TestMissingRefs|TestJudgementQueue|TestSourceLinks|TestCRDObjectTypes|TestCRDBackwardCompatibility|TestHandWrittenProtoStructAnnotationsAndNaming"
   go test -v ./apis/<service>/... ./pkg/controller/direct/<service>/...
   ```
3. Record any service-specific proto or reference quirks in `.gemini/journals/<service>.md`.
