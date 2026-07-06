---
name: crd-mapper-fuzzer-existing-type
description: Standards and workflows for creating direct KRM Go types for an existing CRD while maintaining strict schema compatibility.
---

# CRD Mapper Fuzzer Existing Type

## Overview
This skill outlines standard practices when transitioning an existing KCC resource (e.g., from Terraform or DCL) to a direct controller by generating the initial KRM types (`_types.go`), ensuring **strict schema compatibility** with the baseline CRD. You must NOT change the schema at all (other than descriptions). Do not add any new fields, including `spec.projectRef`, unless it was already part of the baseline CRD. You MUST run `dev/tasks/diff-crds` to check for any schema changes.

## Workflow

### 1. Configure generate.sh
Configure `apis/<service>/<version>/generate.sh` to include the resource and pass the `--include-skipped-output` flag. Passing `--include-skipped-output` to both `generate-types` and `generate-mapper` ensures that any output otherwise skipped is still generated but commented out. This provides an invaluable reference when manual modifications/hand-coding of types are needed.

- **Keep Type File Names Matching Lowercase Proto Message**: If the KRM Kind name differs from the underlying Proto message name (e.g. Kind `NotebookInstance` but Proto `Instance`), do NOT rename the types file to follow the lowercase KRM Kind name (e.g. `notebookinstance_types.go`). The `generate-types` tool expects the file to be named `<lowercase_proto_message_name>_types.go` (e.g. `instance_types.go`). Renaming it will cause generator panics and duplicate/untracked file generation.

```bash
go run . generate-types \
    --service <proto.package> \
    --api-version "<service>.cnrm.cloud.google.com/<version>" \
    --include-skipped-output \
    --resource <Kind>:<ProtoMessage>

go run . generate-mapper \
    --service <proto.package> \
    --api-version "<service>.cnrm.cloud.google.com/<version>" \
    --include-skipped-output
```

### 2. Standards for Strict Schema Compatibility
When defining the KRM Go type in `<kind>_types.go`, you must ensure it matches the original CRD schema exactly. You must not add, remove, or modify any fields (including `spec.projectRef`) in a way that alters the KRM schema:

- **Do Not Change the Schema**: You must **not change the schema** when the type already exists. Description changes are OK, but adding/removing/renaming fields (such as adding `projectRef` if the baseline CRD did not have it) is strictly forbidden.
- **Define Nested Structs as Pointers to Avoid Hand-coded Mappers**: If the baseline CRD defines nested structures (even if required), you should define them as pointers in the KRM Go structs rather than value types. This allows the automatic mapping generator `generate-mapper` to traverse and map them automatically, completely eliminating the need for hand-coded mapper functions. If the baseline CRD schema requires those nested fields, mark the pointer fields with `// +required` (or `// +kubebuilder:validation:Required`) to ensure they remain required in the generated CRD schema.
- **Run diff-crds**: You MUST run `dev/tasks/diff-crds` frequently (and definitely before opening/updating a PR) to identify any schema changes or deviations between the baseline CRD and the generated one. The diff-crds output must be absolutely empty (or contain only minor description reflows if expected).
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
- **Ensure correct +kcc:proto annotation**: Always make sure all KRM types have the correct `// +kcc:proto=<proto_message>` annotation on the struct definition. Even if you think you have to hand-code the mapper, check that the annotation is present. Running the generator produces a commented-out generated mapper inside `mapper.generated.go` if the type is annotated. Review that commented-out mapper first before deciding to hand-code; you might be able to fully automate the mapper simply by renaming the Go struct field to match the proto field name exactly (without changing the json tag).
- **Match Field Formats**: Ensure date/time strings have `// +kubebuilder:validation:Format=date-time` to match original formats.
- **Keep Legacy Reconciler Labels & API Versions**: You must retain existing labels such as `// +kubebuilder:metadata:labels="cnrm.cloud.google.com/dcl2crd=true"` (or `tf2crd=true` for Terraform-based resources) on the struct definition in `<kind>_types.go`. Removing these labels will prevent the corresponding legacy controllers from being registered correctly for the resource during the transition phase. Additionally, if the baseline CRD contains older versions that must be preserved (e.g., `v1alpha1`), you MUST add `// +kubebuilder:metadata:labels="internal.cloud.google.com/additional-versions=v1alpha1"` to the beta resource structure so that those versions are still generated.
- **Validation**: No new fields should be introduced under status. In particular, do not add `observedState` or `externalRef` to the Status struct of the newly generated Go types if they were not already present in the original baseline CRD. Even though `externalRef` is standard for direct controllers, adding it changes the schema and must be avoided for strict schema-compatibility. The output of `dev/tasks/diff-crds` must be absolutely empty.
- **Matching Go struct field names to Proto field names**: The mapping generator automatically matches KRM Go struct fields to Proto fields by mapping Go camel-case names to proto snake_case/camel-case field names. If a KRM field is named differently in the CRD (e.g., `filterLabels`), we can rename the Go struct field to match the proto name exactly (e.g., renaming the Go field name to `Labels` while preserving the JSON tag `json:"filterLabels,omitempty"` so that there is absolutely no schema change). This allows us to fully leverage automatic mapper generation and completely avoid writing hand-coded mappers. If the mapped type is defined in a different package (like `google.api.MonitoredResource`), we can also specify multiple services to the `--service` flag of `generate-mapper` (e.g. `--service google.monitoring.v3,google.api`) to enable automatic traversal and generation for that shared type.

### 3. Fuzzer Best Practices
**Fuzzer implementation timing**: Note that implementing a fuzzer is **NOT required for brownfield resources in the first step** (defining direct types). The fuzzer is implemented and run later in the transition lifecycle (e.g., in Step 3 or 4 when writing and validating the controller reconciliation logic).

When writing a KRM round-trip fuzzer, or if a fuzzer already exists:
- **File Naming**: Ensure the fuzzer is named `<kind>_fuzzer.go` (e.g., `entitlement_fuzzer.go`) in lowercase. Do not use generic names like `fuzzers.go`.
- **Use Type-Safe Helpers**: Do not call `f.SpecFields.Insert`, `f.StatusFields.Insert`, or `f.UnimplementedFields.Insert` directly on sets. Instead, use the type-safe helper methods defined on the `KRMTypedFuzzer` struct:
  - Use `f.SpecField(fieldPath)` to mark a field as round-tripping to/from the Spec.
  - Use `f.StatusField(fieldPath)` to mark a field as round-tripping to/from the Status.
  - Use `f.Unimplemented_Identity(fieldPath)` for identity/URL fields like `.name`.
  - Use `f.Unimplemented_Internal(fieldPath)` for internal service-only/hidden implementation details (e.g., resource fields that are duplicates of KRM metadata or parent references).
  - Use `f.Unimplemented_NotYetTriaged(fieldPath)` for fields that are not implemented or under development.
  - Use `f.Unimplemented_LabelsAnnotations(fieldPath)` for labels or annotations.
- **Move Hand-Coded Mappers**: Ensure all hand-coded mapper functions reside in a file called `mappers.go` within the direct controller package to distinguish them from the generated mapper file `mapper.generated.go`.

**No Dedicated Unit Test Needed**: There is no need to add a dedicated/standalone `_fuzzer_test.go` file for the resource. Simply registering the fuzzer using `fuzztesting.RegisterKRMFuzzer()` in your fuzzer implementation file (e.g. `backupvault_fuzzer.go`) is fully sufficient. The existing shared testing framework (such as `pkg/fuzztesting/fuzztests/fuzz_test.go`) will automatically discover and run it.

### 4. Verification & Acceptance Criteria
- Run `dev/tasks/diff-crds` to verify there are absolutely no unintended schema changes.
- Since we are transitioning an existing type, the primary acceptance criterion is "does it generate the same CRD schema".
- Once the schema is identical, run `make ready-pr` to regenerate Go clients (and compile-check the changes, run custom linters, format the files, and regenerate static configs).
- Finally, run `dev/tasks/generate-resource-report` to update the resource reports `docs/reports/crd_report.csv` and `docs/reports/crd_report.md` to reflect the transition.
