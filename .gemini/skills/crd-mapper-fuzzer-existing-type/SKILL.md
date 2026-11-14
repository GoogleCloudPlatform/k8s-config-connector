---
name: crd-mapper-fuzzer-existing-type
description: Standards and workflows for creating direct KRM Go types for an existing CRD while maintaining strict schema compatibility.
---

# CRD Mapper Fuzzer Existing Type

## Overview
This skill outlines standard practices when transitioning an existing KCC resource (e.g., from Terraform or DCL) to a direct controller by generating the initial KRM types (`_types.go`), ensuring **strict schema compatibility** with the baseline CRD. You must NOT change the schema at all (other than descriptions). Do not add any new fields, including `spec.projectRef`, unless it was already part of the baseline CRD. You MUST run `dev/tasks/diff-crds` to check for any schema changes.

## Migration Lifecycle & Step Boundaries

The KCC Brownfield migration lifecycle is strictly partitioned across distinct sequential steps:
1. **Step 1 (`step/types`)**: Direct KRM Go types, `generate.sh` (`generate-types`), and CRD schema equivalence.
2. **Step 2 (`step/identity`)**: Direct Identity and Reference (`_identity.go`, `_reference.go`, `generate-identity`).
3. **Step 3 (`step/fuzzer`)**: Direct Mappers (`generate-mapper`, `mappers.go`) and round-trip KRM Fuzzer (`<kind>_fuzzer.go`).
4. **Step 4 (`step/mockgcp`)**: MockGCP implementation and Real GCP alignment (`_http.log`).
5. **Step 5 (`step/controller`)**: Direct reconciliation controller (`<kind>_controller.go`) and E2E test fixtures.

### Step 1 Scope Boundaries (Strict Rules)
When working on a **Step 1 (`step/types`)** issue or PR:
- **IN SCOPE**:
  - Direct KRM Go types definition (`apis/<service>/<version>/<kind>_types.go`)
  - Auto-generated deepcopy and type files (`zz_generated.deepcopy.go`, `types.generated.go`)
  - API package files if new package (`doc.go`, `groupversion_info.go`)
  - Configuration of `apis/<service>/generate.sh` to run `generate-types` and `dev/tasks/generate-crds`
  - CRD regeneration in `config/crds/resources/`
  - Zero-diff validation via `./dev/tasks/diff-crds --base master --allow-status-external-ref-addition`
- **STRICTLY OUT OF SCOPE (DO NOT INCLUDE IN STEP 1 PRs)**:
  - **NO Mappers**: Do NOT run `generate-mapper` or create mapper files (`pkg/controller/direct/<service>/mapper.generated.go`, `mappers.go`). Mappers belong to **Step 3 (Fuzzer)** / **Step 5 (Controller)**. Running `generate-mapper` prematurely causes cross-service code pollution and unneeded churn.
  - **NO Identity/References**: Do NOT run `generate-identity` or create `_identity.go` / `_reference.go`. Identity belongs to **Step 2 (Identity)**.
  - **NO Controller Logic**: Do NOT implement `<kind>_controller.go` or touch `pkg/controller/direct/`.
  - **NO Committed Journals**: Do NOT commit agent journal markdown files or anything under `.gemini/skills/` to git branches.

## Step 1 Workflow

### 1. Configure generate.sh (Types Only)
Configure `apis/<service>/generate.sh` to include the resource for type generation only:

* **Types Generation**: Run `generate-types` for each version present under `apis/<service>/` (e.g. `v1alpha1` and `v1beta1`).
* **DO NOT run `generate-mapper` or `generate-identity` in Step 1**: Mappers belong to Step 3 and Identity belongs to Step 2.

Sourcing `${REPO_ROOT}/dev/tools/goimports.sh`, cd'ing to `${REPO_ROOT}/dev/tools/controllerbuilder`, and passing `--include-skipped-output` to `generate-types` ensures that any output otherwise skipped is still generated but commented out. This provides an invaluable reference when manual modifications/hand-coding of types are needed.

- **Missing Protos in generate-proto.sh**: If `generate-types` reports `failed to find the proto message ... proto: not found`, verify that the proto package directory (e.g. `${VERSION_DIR}/google/<service>/*/*.proto`) is listed under `PROTO_FILES` in `dev/tools/controllerbuilder/generate-proto.sh`.
- **Keep Type File Names Matching Lowercase Proto Message**: If the KRM Kind name differs from the underlying Proto message name (e.g. Kind `NotebookInstance` but Proto `Instance`), do NOT rename the types file to follow the lowercase KRM Kind name (e.g. `notebookinstance_types.go`). The `generate-types` tool expects the file to be named `<lowercase_proto_message_name>_types.go` (e.g. `instance_types.go`). Renaming it will cause generator panics and duplicate/untracked file generation.

```bash
# Run from dev/tools/controllerbuilder for Step 1
go run . generate-types \
    --service <proto.package> \
    --api-version "<service>.cnrm.cloud.google.com/<version>" \
    --include-skipped-output \
    --resource <Kind>:<ProtoMessage>
```

### 2. Standards for Strict Schema Compatibility
When defining the KRM Go type in `<kind>_types.go`, you must ensure it matches the original CRD schema exactly. You must not add, remove, or modify any fields (including `spec.projectRef`) in a way that alters the KRM schema:

- **Do Not Change the Schema**: You must **not change the schema** when the type already exists. Description changes are OK, but adding/removing/renaming fields (such as adding `projectRef` if the baseline CRD did not have it) is strictly forbidden.
- **Define Nested Structs as Pointers to Avoid Hand-coded Mappers**: If the baseline CRD defines nested structures (even if required), you should define them as pointers in the KRM Go structs rather than value types. This allows the automatic mapping generator `generate-mapper` to traverse and map them automatically, completely eliminating the need for hand-coded mapper functions. If the baseline CRD schema requires those nested fields, mark the pointer fields with `// +required` (or `// +kubebuilder:validation:Required`) to ensure they remain required in the generated CRD schema.
- **Run diff-crds**: You MUST run `dev/tasks/diff-crds --allow-status-external-ref-addition` frequently (and definitely before opening/updating a PR) to identify any schema changes or deviations between the baseline CRD and the generated one. The diff-crds output must be absolutely empty (or contain only minor description reflows if expected).
- **No spec.projectRef addition**: If the baseline CRD did not contain `spec.projectRef`, do not add it to the Spec struct in `<kind>_types.go`.
- **Reference Hand-coding & Manual Edits**: If there are schema mismatches, you must manually copy, edit, or hand-code types (e.g. to change or remove fields) until the schemas match perfectly.
- **Hand-code custom reference types**: If a resource reference structure in the baseline CRD (like `FolderRef`, `OrganizationRef`, or `BillingAccountRef`) lacks a `kind` field or retains specific fields (like `name`/`namespace` in `OrganizationRef`), you **must hand-code custom reference types** locally in `<kind>_types.go`. For a project reference without a `kind` field, you MUST import and use `refs.ProjectRef` from `github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs` instead of defining a local structure.
- **Use Real Reference Types**: You must NOT redefine or duplicate existing canonical reference types (such as `StorageBucketRef` or `BigQueryTableRef`) locally. Instead, import and use the real, canonical reference types from their respective packages under `apis/<service>/<version>` (e.g., `github.com/GoogleCloudPlatform/k8s-config-connector/apis/storage/v1beta1` for `StorageBucketRef`). This ensures they correctly implement `refs.Ref` and successfully resolve with common KCC helper functions.
- **Match Signatures for automatic validation**: The schema builder `scripts/add-validation-to-crds` automatically adds OpenAPI `oneOf` blocks depending on field signatures.
  - If a custom reference type has fields `external`, `name`, and `namespace`, its signature is `"external,name,namespace"`, which triggers the generator to automatically append the `oneOf` reference constraint block.
- **Spec-level validations**: If the baseline CRD contains spec-level validations (such as requiring exactly one of different parent reference fields), ensure those constraints are registered/hardcoded in `scripts/add-validation-to-crds/parse-crds.go` so they are successfully preserved in the generated CRD.
- **Handling Type Mismatches in generated mappers**: If a KRM field type (such as `int64`) does not match the proto's field type (such as `int32`), the generated assignment code will fail compilation. Solve this by writing handcoded mapping functions (e.g., `KindSpec_FromProto` / `ToProto`) in a handcoded `mapper.go` file within the direct controller package. The generator will automatically detect these and skip generating the conflicting versions.
  - **Prefer converting int32 fields to int64 in Go types**: If a KRM field is defined as `int32` in the CRD/Go types but its corresponding field in the GCP Go client library (e.g., `computepb`) is defined as `int64`, you should convert the Go KRM type of that field from `int32` to `int64`. This is backward compatible (as `int64` is a superset of `int32` on the Kubernetes side), allows the generated mapping code to compile cleanly, and completely removes any need for hand-coded mapping functions, lossy conversions, or fuzzer truncation filters.
- **Document Handcoded Deviations**: For all handcoded mapping functions, you MUST include clear comments above the function explaining exactly why it deviates from the default behavior and why it cannot be automatically generated (e.g., secret reference fields, body string-to-byte-slice mapping, or custom reference type logic).
- **Ensure correct +kcc:proto annotation**: Always make sure KRM type structs representing API fields have the correct `// +kcc:proto=<proto_message>` annotation. However, the top-level root KRM struct (e.g., `ComputeMachineImage`) does NOT need the `// +kcc:proto=<proto_message>` annotation on its definition (place it only on `Spec` and `Status` structures). Leaving `+kcc:proto` off the root struct prevents the generator from outputting conflicting root-level mapping functions (such as `FromProto` / `ToProto`), avoiding type assignment errors where the GCP proto status field (such as a string pointer) conflicts with the KRM custom status struct.
- **Match Field Formats**: Ensure date/time strings have `// +kubebuilder:validation:Format=date-time` to match original formats.
- **Keep Legacy Reconciler Labels & API Versions**: You must retain existing labels such as `// +kubebuilder:metadata:labels="cnrm.cloud.google.com/dcl2crd=true"` (or `tf2crd=true` for Terraform-based resources) on the struct definition in `<kind>_types.go`. Removing these labels will prevent the corresponding legacy controllers from being registered correctly for the resource during the transition phase. Additionally, if the baseline CRD contains older versions that must be preserved (e.g., `v1alpha1`), you MUST add `// +kubebuilder:metadata:labels="internal.cloud.google.com/additional-versions=v1alpha1"` to the beta resource structure so that those versions are still generated.
- **Validation & Status ExternalRef**: Always include `ExternalRef` in `<Kind>Status`:
  ```go
  	// A unique specifier for the <Kind> resource in GCP.
  	ExternalRef *string `json:"externalRef,omitempty"`
  ```
  (See `apis/accesscontextmanager/v1beta1/accesscontextmanagerserviceperimeter_types.go:L338` for reference).
  Adding `status.externalRef` is expected and explicitly allowed by `dev/tasks/diff-crds --allow-status-external-ref-addition` (used by CI `dev/ci/presubmits/crd-equivalence-check`).
  No other new fields should be introduced under `spec` or `status`. In particular, do not add `observedState` to `<Kind>Status` if it was not already present in the original baseline CRD.
  - **Handling Baseline CRDs without ObservedState in Status**: If the baseline CRD status does not include `observedState`, do not add `observedState` to `<Kind>Status`. To satisfy `generate-mapper` when GCP proto messages have output-only fields, define an empty `<Kind>ObservedState` struct (with `// +kcc:observedstate:proto=...`), but do NOT reference it as a field inside `<Kind>Status`. This allows `mapper.generated.go` to compile cleanly without polluting the generated CRD schema.
- **Preserve Lack of Top-Level Descriptions (`_no_description`)**: If the baseline CRD does not have top-level `description` fields on `spec:` or `status:`, omit doc comments on the `type <Kind>Spec struct` and `type <Kind>Status struct` definitions. Adding doc comments on those structs causes `controller-gen` to generate descriptions for `spec` and `status`, which triggers `_no_description.spec=true` or `_no_description.status=true` diffs in `dev/tasks/diff-crds`.
- **Preserve OpenAPI integer without format**: If the baseline CRD defined integer fields with `type: integer` without `format: int32` or `format: int64`, using Go `int32` / `int64` types will cause `controller-gen` to emit `format: int32` / `int64`, failing `TestCRDBackwardCompatibility` (`field type changed: ... (integer -> int32)`). Use Go `int` (or `*int` for optional fields) so `controller-gen` emits plain `type: integer` without a format tag, and write handcoded mappers if needed to cast between `int` and proto `int32` / `int64`.
- **Matching Go struct field names to Proto field names**: The mapping generator automatically matches KRM Go struct fields to Proto fields by mapping Go camel-case names to proto snake_case/camel-case field names. If a KRM field is named differently in the CRD (e.g., `filterLabels`), we can rename the Go struct field to match the proto name exactly (e.g., renaming the Go field name to `Labels` while preserving the JSON tag `json:"filterLabels,omitempty"` so that there is absolutely no schema change). This allows us to fully leverage automatic mapper generation and completely avoid writing hand-coded mappers. If the mapped type is defined in a different package (like `google.api.MonitoredResource`), we can also specify multiple services to the `--service` flag of `generate-mapper` (e.g. `--service google.monitoring.v3,google.api`) to enable automatic traversal and generation for that shared type.

### 3. Step 3 (Fuzzer & Mappers) Best Practices
**Fuzzer and Mapper implementation timing**: Implementing mappers and fuzzers is **strictly deferred to Step 3 (`step/fuzzer`)**. They MUST NOT be included in Step 1 PRs.

When implementing Step 3 (`step/fuzzer`):
- **Generate Mappers**: Add `generate-mapper` to `apis/<service>/generate.sh` targeting the direct controller directory `pkg/controller/direct/<service>/`.
- **File Naming**: Ensure the fuzzer is named `<kind>_fuzzer.go` (e.g., `entitlement_fuzzer.go`) in lowercase under `pkg/controller/direct/<service>/`. Do not use generic names like `fuzzers.go`.
- **Use Type-Safe Helpers**: Do not call `f.SpecFields.Insert`, `f.StatusFields.Insert`, or `f.UnimplementedFields.Insert` directly on sets. Instead, use the type-safe helper methods defined on the `KRMTypedFuzzer` struct:
  - Use `f.SpecField(fieldPath)` to mark a field as round-tripping to/from the Spec.
  - Use `f.StatusField(fieldPath)` to mark a field as round-tripping to/from the Status.
  - Use `f.Unimplemented_Identity(fieldPath)` for identity/URL fields like `.name`.
  - Use `f.Unimplemented_Internal(fieldPath)` for internal service-only/hidden implementation details (e.g., resource fields that are duplicates of KRM metadata or parent references).
  - Use `f.Unimplemented_NotYetTriaged(fieldPath)` for fields that are not implemented or under development.
  - Use `f.Unimplemented_LabelsAnnotations(fieldPath)` for labels or annotations.
- **Move Hand-Coded Mappers**: Ensure all hand-coded mapper functions reside in a file called `mappers.go` within the direct controller package to distinguish them from the generated mapper file `mapper.generated.go`.

**No Dedicated Unit Test Needed**: There is no need to add a dedicated/standalone `_fuzzer_test.go` file for the resource. Simply registering the fuzzer using `fuzztesting.RegisterKRMFuzzer()` in your fuzzer implementation file (e.g. `backupvault_fuzzer.go`) is fully sufficient. The existing shared testing framework (such as `pkg/fuzztesting/fuzztests/fuzz_test.go`) will automatically discover and run it.

### 4. Verification & Acceptance Criteria (Step 1 Types)
- Run `dev/tasks/diff-crds --base master --allow-status-external-ref-addition` to verify there are absolutely no unintended schema changes.
- Since we are transitioning an existing type, the primary acceptance criterion is "does it generate the same CRD schema" (with the expected addition of `status.externalRef`).
- Verify no mappers, identity files, or controller files were added to the PR.
- Once the schema is identical, run `make ready-pr` to regenerate Go clients (and compile-check the changes, run custom linters, format the files, and regenerate static configs).
- When creating the PR, include the release note block set to NONE:
  ```release-note
  NONE
  ```

