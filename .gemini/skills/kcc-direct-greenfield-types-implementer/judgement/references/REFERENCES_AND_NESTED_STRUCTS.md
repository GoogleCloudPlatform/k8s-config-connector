# Resource References, Output-Only Fields, & Shared Nested Structs Playbook

Use this reference when resolving any of the following `judgement_queue.yaml` reasons for `<Kind>`:
- **Resource References**: `authoritative-proto-reference`, `possible-reference`, `possible-reference-by-description`, `possible-reference-by-description-loose`, `possible-reference-by-name`, `possible-reference-by-sibling`, `reference-not-representable`
- **Output-Only in Comments**: `well-known-output-only-pattern-in-comment`, `possible-output-only-pattern-in-comment`
- **Required Fields on Shared Structs**: `required-not-enforced`

---

## 1. Customizing Nested Structs in `types.generated.go` (`+kcc:proto=` & Shared Struct Isolation)

Whenever resolving a queue entry requires a **breaking** edit to a nested struct in `apis/<service>/v1alpha1/types.generated.go` (removing fields, converting fields to `*Ref`, or adding `+required`), first check which other Kinds in `apis/<service>/` reference that struct via `grep -rn "<Struct>" apis/<service>/` *(do not rely only on `judgement_queue.yaml`: pre-existing old-scheme Kinds have **0** entries in `judgement_queue.yaml`, already-triaged Kinds have `untriaged-bulk-generation` with `status: resolved`, and untriaged Kinds have `status: open`)*:

1. **Always Preserve `// +kcc:proto=<full.proto.MessageName>`**:
   - Every hand-written nested struct in `<kind>_types.go` **must** carry its `// +kcc:proto=<full.proto.MessageName>` marker comment (and `// +kcc:proto:field=...` on its fields).
   - `controllerbuilder` uses `+kcc:proto=` to generate mapper functions and to distinguish base structs (`type <Struct> struct`) from variant structs; `TestHandWrittenProtoStructAnnotationsAndNaming` in `tests/apichecks` fails if `+kcc:proto=` is missing or if the Go struct name does not contain `<Struct>`.

2. **Non-breaking changes from `generate.sh`** (new optional fields, updated doc comments):
   - Safe in place on `type <Struct> struct` in `types.generated.go`.

3. **Breaking changes when `<Struct>` is ONLY used by `<Kind>` (or only by untriaged Kinds being updated together)**:
   - Move `type <Struct> struct` (keeping its **original Go name** and `// +kcc:proto=<full.proto.MessageName>`) from `types.generated.go` into `<kind>_types.go` and edit it in place.
   - Do **not** move unmodified parent structs that merely reference `<Struct>`; `generate-types` will keep generating the parent structs in `types.generated.go` and automatically point them at your `<Struct>` in `<kind>_types.go`.

4. **Breaking changes when `<Struct>` is SHARED with another Kind** (especially a pre-existing Kind with 0 queue entries, an already-triaged Kind with `status: resolved`, or another untriaged Kind you want to keep out of this PR's diff):
   - **Never** mutate `type <Struct> struct` under its original name in place — doing so breaks other Kinds' CRDs and mappers (caught by `TestCRDBackwardCompatibility`)!
   - Instead, reuse an existing compatible hand-written variant in the package if one matches your schema, or fork `<Struct>` into `<kind>_types.go` under a **distinct struct name that still contains `<Struct>`** (e.g., `<Kind>_<Struct>`, `<Struct>Spec` paired with `<Struct>ObservedState`, or the generated `<Struct>WithRequired`) while preserving `// +kcc:proto=<full.proto.MessageName>`.
   - Because the variant's Go name differs from the base `<Struct>`, `controllerbuilder` automatically retains generator ownership of the base `type <Struct> struct` in `types.generated.go` for other Kinds.
   - *(Note on deep nesting)*: If `<Struct>` sits below an intermediate generated struct (`<Kind>Spec` -> `Outer` -> `<Struct>`) that is *also* shared with another Kind, fork `Outer` in `<kind>_types.go` as well (keeping `+kcc:proto=` and `Outer` in its name) so `<Kind>Spec` points to the customized chain while `types.generated.go` keeps the untouched `Outer` -> `<Struct>` for other Kinds.

---

## 2. Resolving Resource Reference Reasons (`authoritative-proto-reference`, `possible-reference*`, `reference-not-representable`)

- **`authoritative-proto-reference`** (and legacy `possible-reference`): Emitted when the proto field carries an explicit `google.api.resource_reference` (`type` or `child_type`) annotation. When `google.api.resource_reference` is present, heuristic `possible-reference-*` signals on that field are suppressed so `authoritative-proto-reference` is the single finding.
- **`possible-reference-by-description*`, `possible-reference-by-name`, `possible-reference-by-sibling`**: Emitted when `google.api.resource_reference` is absent; a single field may carry 2–3 contiguous heuristic `possible-reference-*` entries that reinforce signal. Read **all** of their `detail:` lines together alongside the proto field comment:

- **Case A — The field references a GCP resource**:
  1. If the field is inside a nested struct in `types.generated.go`, move or fork the struct into `<kind>_types.go` following Section 1 above (keeping `// +kcc:proto=<full.proto.MessageName>`; use `<Struct>` if exclusive to `<Kind>`, or a distinct `<Struct>`-containing name such as `<Kind>_<Struct>` if shared).
  2. Replace the `*string` or `[]string` field with `<fieldName>Ref *<TargetKind>Ref` (or `[]<TargetKind>Ref` for slices, singularizing the plural field name and appending `Refs`, e.g., `datasets` -> `DatasetRefs` with `json:"datasetRefs,omitempty"`), keeping `// +kcc:proto:field=...`.
  3. Remove any `// +kcc:guess=possible-reference` or `// +kcc:guess=sibling-ref` comment above the field in `<kind>_types.go`.
  4. **If `<TargetKind>Ref` does not exist yet in the repository**:
     - If `<TargetKind>` is in the same proto service, generate its identity and reference via `controllerbuilder generate-identity` (or create `<targetkind>_identity.go` and an external-only `<targetkind>_reference.go` in `apis/<target_service>/v1alpha1/` following `.gemini/skills/kcc-direct-base-types-implementer/SKILL.md`, updating `testdata/exceptions/naming_violations.txt` if `TestDirectResourceFileNaming` requires it).
     - **Never** add new files or types under `apis/refs/v1beta1/`.
     - For greenfield reference types that already have a full KCC CRD with `status.externalRef`, delegate `Normalize` strictly to `refs.Normalize` (never `refs.NormalizeWithFallback`).
  5. Mark the `authoritative-proto-reference` entry (or **all** contiguous `possible-reference*` entries) for that field as `status: resolved`, `resolution: edited`, `note: "converted to <fieldName>Ref (*<TargetKind>Ref)"`.

- **Case B — The field does NOT reference a control-plane GCP resource** (e.g., an opaque user-defined ID, a third-party ID, a display string, a CVE ID, a filter expression, a wildcard `*` annotation, or a data-plane object not managed by KCC such as a Cloud Storage object / `StorageBucketObject`):
  1. Remove any `// +kcc:guess=possible-reference` or `// +kcc:guess=sibling-ref` comment if the field is in `<kind>_types.go`.
  2. Mark the `authoritative-proto-reference` entry (or **all** contiguous `possible-reference*` entries) for that field as `status: resolved`, `resolution: not-applicable`, `note: "<concise explanation of what the field holds and why it is not a GCP resource reference>"`.
  3. Run `go test ./tests/apichecks/... -run TestMissingRefs` to verify `TestMissingRefs` passes **without** deleting/modifying proto doc comments or editing `missingrefs.txt`.

- **Case C — `reference-not-representable`**:
  1. Confirm the reason in `detail` (e.g., polymorphic multi-kind resource name or format KCC cannot express as a typed `*Ref` today).
  2. Mark `status: resolved`, `resolution: accepted` (or `edited` if a single canonical KCC reference applies).

---

## 3. Resolving Output-Only Comment Reasons (`well-known-output-only-pattern-in-comment`, `possible-output-only-pattern-in-comment`)

The entry's `detail` specifies where the field was generated in Spec (`.spec.<x>`) and where it would live in ObservedState (`.status.observedState.<x>`). Inspect the proto comment:

- **If genuinely server-populated (`OUTPUT_ONLY` in prose)**:
  1. **Top-level field**: Move the field from `<Kind>Spec` to `<Kind>ObservedState` in `<kind>_types.go`.
  2. **Nested struct field**: Move or fork the containing struct into `<kind>_types.go` following Section 1 (preserving `// +kcc:proto=<full.proto.MessageName>`), remove the output-only field from the Spec struct, and ensure `<Struct>ObservedState` carries the field under `<Kind>ObservedState`.
  3. Mark `status: resolved`, `resolution: edited`, `note: "moved output-only field from <spec_path> to <observed_state_path>"`.
- **If user-settable** (e.g., `"Output only for X mode, required for Y mode"` or a false positive):
  1. Keep the field in Spec and mark `status: resolved`, `resolution: not-applicable`, `note: "<why the field is user-configurable in Spec>"`.

---

## 4. Resolving `required-not-enforced`

Follow the instructions in `detail`:
- If a hand-written struct holds a generated struct `<Struct>` that has a `<Struct>WithRequired` copy, switch the field type to `<Struct>WithRequired` (isolating the parent struct per Section 1 if the parent struct is shared with a non-strict Kind).
- If a shared struct in `types.generated.go` needs `+required` alongside other edits (such as `*Ref` conversions or output-only field removal), fork it in `<kind>_types.go` per Section 1 (preserving `// +kcc:proto=<full.proto.MessageName>` and keeping `<Struct>` in its Go name) and add `// +required` to the required fields. Mark `resolution: edited`.
