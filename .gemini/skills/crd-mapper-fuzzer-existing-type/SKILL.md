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
- **Run diff-crds**: You MUST run `dev/tasks/diff-crds` frequently (and definitely before opening/updating a PR) to identify any schema changes or deviations between the baseline CRD and the generated one. The diff-crds output must be absolutely empty (or contain only minor description reflows if expected).
- **No spec.projectRef addition**: If the baseline CRD did not contain `spec.projectRef`, do not add it to the Spec struct in `<kind>_types.go`.
- **Reference Hand-coding & Manual Edits**: If there are schema mismatches, you must manually copy, edit, or hand-code types (e.g. to change or remove fields) until the schemas match perfectly.
- **Hand-code custom reference types**: If a resource reference structure (like `ProjectRef`, `FolderRef`, `OrganizationRef`, or `BillingAccountRef`) in the baseline CRD lacks a `kind` field or retains specific fields (like `name`/`namespace` in `OrganizationRef`), you **must hand-code custom reference types** locally in `<kind>_types.go` or import the clean, kindless reference types from `github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs` (e.g. `refs.ProjectRef`).
- **Match Signatures for automatic validation**: The schema builder `scripts/add-validation-to-crds` automatically adds OpenAPI `oneOf` blocks depending on field signatures.
  - If a custom reference type has fields `external`, `name`, and `namespace`, its signature is `"external,name,namespace"`, which triggers the generator to automatically append the `oneOf` reference constraint block.
- **Spec-level validations**: If the baseline CRD contains spec-level validations (such as requiring exactly one of different parent reference fields), ensure those constraints are registered/hardcoded in `scripts/add-validation-to-crds/parse-crds.go` so they are successfully preserved in the generated CRD.
- **Handling Type Mismatches in generated mappers**: If a KRM field type (such as `int64`) does not match the proto's field type (such as `int32`), the generated assignment code will fail compilation. Solve this by writing handcoded mapping functions (e.g., `KindSpec_FromProto` / `ToProto`) in a handcoded `mapper.go` file within the direct controller package. The generator will automatically detect these and skip generating the conflicting versions.
- **Match Field Formats**: Ensure date/time strings have `// +kubebuilder:validation:Format=date-time` to match original formats.
- **Keep Legacy Reconciler Labels**: You must retain existing labels such as `// +kubebuilder:metadata:labels="cnrm.cloud.google.com/dcl2crd=true"` (or `tf2crd=true` for Terraform-based resources) on the struct definition in `<kind>_types.go`. Removing these labels will prevent the corresponding legacy controllers from being registered correctly for the resource during the transition phase.
- **Validation**: No new fields should be introduced under status (e.g. do not add `observedState` or `externalRef` yet if they weren't in the original CRD).

### 3. Verification & Acceptance Criteria
- Run `dev/tasks/diff-crds` to verify there are absolutely no unintended schema changes.
- Since we are transitioning an existing type, we do not need a fuzzer yet. The primary acceptance criterion is "does it generate the same CRD schema".
- Once the schema is identical, run `make ready-pr` to regenerate Go clients (and compile-check the changes, run custom linters, format the files, and regenerate static configs).
- Finally, run `dev/tasks/generate-resource-report` to update the resource reports `docs/reports/crd_report.csv` and `docs/reports/crd_report.md` to reflect the transition.
