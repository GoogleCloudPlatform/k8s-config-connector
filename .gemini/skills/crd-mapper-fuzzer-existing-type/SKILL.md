---
name: crd-mapper-fuzzer-existing-type
description: Standards and workflows for creating direct KRM Go types for an existing CRD while maintaining strict schema compatibility.
---

# CRD Mapper Fuzzer Existing Type

## Overview
This skill outlines standard practices when transitioning an existing KCC resource (e.g., from Terraform or DCL) to a direct controller by generating the initial KRM types (`_types.go`), ensuring **strict schema compatibility** with the baseline CRD.

## Workflow

### 1. Configure generate.sh
Configure `apis/<service>/<version>/generate.sh` to include the resource and use the `--include-skipped-output` flag. Passing `--include-skipped-output` ensures that any output otherwise skipped is still generated but commented out. This provides an invaluable reference when manual modifications/hand-coding are needed.

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
When defining the KRM Go type in `<kind>_types.go`, you must ensure it matches the original CRD schema exactly:

- **Run diff-crds**: Run `dev/tasks/diff-crds` frequently to identify any schema changes or deviations between the baseline CRD and the generated one.
- **Reference Hand-coding**: If a resource reference structure (like `ProjectRef`, `FolderRef`, `OrganizationRef`, or `BillingAccountRef`) in the baseline CRD lacks a `kind` field or retains specific fields (like `name`/`namespace` in `OrganizationRef`), you **must hand-code custom reference types** locally in `<kind>_types.go` instead of importing the standard types from `apis/refs/v1beta1/`.
- **Match Signatures for automatic validation**: The schema builder `scripts/add-validation-to-crds` automatically adds OpenAPI `oneOf` blocks depending on field signatures.
  - If a custom reference type has fields `external`, `name`, and `namespace`, its signature is `"external,name,namespace"`, which triggers the generator to automatically append the `oneOf` reference constraint block.
- **Match Field Formats**: Ensure date/time strings have `// +kubebuilder:validation:Format=date-time` to match original formats.
- **Validation**: No new fields should be introduced under status (e.g. do not add `observedState` or `externalRef` yet if they weren't in the original CRD).

### 3. Verification & Acceptance Criteria
- Run `dev/tasks/diff-crds` to verify there are absolutely no unintended schema changes.
- Since we are transitioning an existing type, we do not need a fuzzer yet. The primary acceptance criterion is "does it generate the same CRD schema".
- Once the schema is identical, run `dev/tasks/generate-resource-report` to update the resource reports `docs/reports/crd_report.csv` and `docs/reports/crd_report.md` to reflect the transition.
