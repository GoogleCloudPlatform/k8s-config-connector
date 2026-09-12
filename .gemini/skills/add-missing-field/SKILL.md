---
name: add-missing-field
description: Guides adding missing fields already supported in GCP API but not yet supported in KCC CRD to Config Connector direct resources.
---

# Skill: Add Missing Field

## Prerequisites & Pre-checks
1. **Controller Type Check**: Ensure the target resource is managed by a **Direct** controller (for legacy resources managed via Terraform/DCL, use the `update-terraform-fields` skill instead).
2. **SDK Freshness**: Verify that the field exists in `.build/third_party/googleapis/` and the target Go SDK in `go.mod`. If the protobuf definitions or client library are outdated, first update the client library using the `update-client-library-version` skill.
3. **Spec vs. ObservedState Split**: Determine whether the field is **Spec** (user-configurable input) or **ObservedState / Status** (output-only / read-only):
   - **Spec fields**: Require KRM type definition in `Spec`, symmetrical mapping (`ToProto` and `FromProto`), controller update/diff handling, fuzzer `f.SpecField(...)`, and E2E test fixture coverage.
   - **ObservedState fields**: Require KRM type definition in `ObservedState`, `FromProto` mapping only, and fuzzer `f.StatusField(...)`. They do not require `ToProto`, update logic, or equality diffs.

---

## Workflow

### Step 1: KRM API Type Scaffolding
- Check if the field represents a resource reference (URI, resource name, or service-generated ID):
  - **If Reference Type (`*refs.<Kind>Ref`)**:
    - If the parent struct is in the manual types file (`apis/<service>/<version>/<message>_types.go`), add the reference field there.
    - If the parent struct is in `types.generated.go`, move the **entire parent struct definition** from `types.generated.go` to `<message>_types.go` and update the field to use the reference type (e.g., `*refs.<Kind>Ref`).
  - **If Not a Reference Type**:
    - If the parent struct is already in the manual `<message>_types.go` file, add the field manually there.
    - If the parent struct is generated (not in `<message>_types.go`), do not add it manually—proceed to Step 2 to generate it automatically.
- Annotate the field with:
  1. `// +optional`
  2. `// +kcc:proto:field=<GCP_PROTO_PATH>` (e.g., `// +kcc:proto:field=google.cloud.apigateway.v1.Api.state`).

### Step 2: Code & CRD Regeneration
- Run code generation to rebuild CRDs and generate types/mappers:
  ```bash
  dev/tasks/generate-types-and-mappers
  ```
  *(or `./apis/<service>/<version>/generate.sh`)*
- Confirm the new field appears in `types.generated.go` (if generated) and under `config/crds/resources/`.

### Step 3: Controller & Mapper Updates
- **Mappers (`mapper.go` / `_mappings.go`)**:
  - **Manual fields/structs**: If defined in `<message>_types.go` (e.g., using `*refs.<Kind>Ref`), write manual conversion functions (`ToProto` and `FromProto` for Spec; `FromProto` only for ObservedState).
  - **Generated fields**: Automatically mapped in `mapper.generated.go`.
- **Controller Reconciliation (`_controller.go`)**:
  - In `Adapter.Update()`, ensure mutable Spec fields are included in the `UpdateMask` / patch payload when modified.
  - **Preserve live state (`actual`)** during updates when the KRM spec is unspecified.
- **Equality Checks (`_equality.go`)**:
  - Update `Diff` logic. If the newly added field is omitted from the KRM spec, **do not diff it** (treat it as unmanaged).

### Step 4: Fuzzer Support & Focused Tests
- In `<resource>_fuzzer.go` (or `<resource>_legacy_fuzzer.go`), promote the field from `f.Unimplemented_NotYetTriaged` to:
  - `f.SpecField(...)` for spec fields (e.g., `f.SpecField(".my_field")`).
  - `f.StatusField(...)` for status / observedState fields.
  - If introducing a complex struct but deferring nested subfields, mark deferred subfields with `f.Unimplemented_NotYetTriaged(".parent.subfield")`.
- Verify round-trip mapping with focused fuzzer tests:
  ```bash
  FOCUS=MyResourceKind go test -v ./pkg/fuzztesting/fuzztests/... -run TestFocusedMappers
  ```

### Step 5: Test Fixtures & Ratcheting Exclusions
- Add the field to `create.yaml` (and `update.yaml` if mutable) under `pkg/test/resourcefixture/testdata/basic/<service>/...`.
- **Remove from Ratcheting (Mandatory)**: In `tests/e2e/ratcheting.go`, remove the resource's `GroupKind` case from `ShouldTestRereconiliation` to enable re-reconciliation validation.

### Step 6: Real GCP Recording & Mock Alignment
- Delegate to the **`match-mockgcp-with-realgcp`** skill:
  1. Record live traffic: `hack/record-gcp <fixture-path>`
  2. Match mock behavior: `hack/compare-mock <fixture-path>` (implement minimal stubs under `mockgcp/` if needed).

### Step 7: Validation & Presubmit
- Run API checks:
  ```bash
  go test -v ./tests/apichecks/...
  ```
- Run pre-PR check:
  ```bash
  make ready-pr
  ```

### Step 8: Journaling
- Append any GCP SDK nuances, API quirks, or service-specific findings to `.gemini/journals/<service>.md` (per the `kcc-agentic-journaler` skill).

---

## Reference & Deep-Dive Guidelines

### Reference 1: Reference Types & Resolution
- **Prefer Reference Types**: Use `*refs.<Kind>Ref` rather than raw `*string` when a field represents a GCP resource URI or identifier.
- **Acronym Naming**: Capitalize acronyms in field names (e.g., `KMSKeyRef *refs.KMSCryptoKeyRef` rather than `KmsKeyRef`) so the code generator can automatically detect and map the reference.
- **Automatic vs. Manual Resolution**:
  - If `AdapterForObject` calls `common.NormalizeReferences(ctx, reader, obj, nil)`, references are resolved automatically before mapping.
  - Manual `refs.Resolve...` calls in the controller are only needed if `common.NormalizeReferences` is not used.
- **Proto Pointer Binding (`usesPointersInProtoBinding`)**: In proto3, fields declared as `optional string` map to `*string` in Go. If the target service is not in `usesPointersInProtoBinding`, the generator may emit `out.KmsKey = in.KMSKeyRef.External`. Write a handwritten `Spec_ToProto` mapper in `<resource>_mapper.go` to assign the pointer:
  ```go
  out.KmsKey = &in.KMSKeyRef.External
  ```

### Reference 2: Mapper Nuances & Skipped Output
- **Root Type Mapper Bypass**: If the root `Spec` or `ObservedState` has a handwritten mapper (`_FromProto` / `_ToProto`), new top-level fields will not be auto-generated; they must be manually mapped in that function.
- **Cascading Skipped Mappers (`// MISSING: <Type>`)**: When the generator detects `/* found existing non-generated mapping function ... */`, it skips generating child mappers. Handwritten `_FromProto` / `_ToProto` functions must be provided for any nested types marked as `// MISSING: <Type>` in `mapper.generated.go`.
- **Inspection Helper**: Passing `--include-skipped-output` to `generate.sh` prints the skipped generated code directly into `types.generated.go` and `mapper.generated.go`, serving as a helpful template for handwritten mappers.

### Reference 3: Optional Zero-Value Fields (`ForceSendFields` vs. Proto3 Pointers)
- **Proto2 / REST Discovery Clients**: In SDKs where fields are value types with `omitempty` (e.g., Cloud SQL `google.golang.org/api/sql/v1beta4` or Compute), explicitly configured zero-values (`false`, `0`) require adding the field name to `ForceSendFields`:
  ```go
  if in.SomeBoolean != nil {
      out.SomeBoolean = *in.SomeBoolean
      out.ForceSendFields = append(out.ForceSendFields, "SomeBoolean")
  }
  ```
- **Proto3 / gRPC SDKs**: Field presence is natively represented via Go pointers (e.g., `*bool`). Non-nil zero-values (`false`, `0`) are automatically serialized and sent without `ForceSendFields`.

### Reference 4: Backwards Compatibility & Preserving Live State
- **Non-Destructive Reconciliation**: Adding a new field to a CRD must never break existing live GCP resources. Unspecified fields parse as `nil`.
- **Avoid Destructive Static Defaulting**: Do not default omitted fields to a static value if that field already has an active value on GCP.
- **Preserve Live State (`actual`)**: When updating a resource, copy live values from `actual` to `out` if omitted in the KRM spec:
  ```go
  // Preserve live GCP state when KRM spec is unspecified
  if in.MyNewField == nil && actual.MyNewField != nil {
      out.MyNewField = actual.MyNewField
  }
  ```
