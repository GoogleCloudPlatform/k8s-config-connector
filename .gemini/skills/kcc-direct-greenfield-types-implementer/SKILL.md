---
name: kcc-direct-greenfield-types-implementer
description: Automate the initial scaffolding of a KCC "direct" resource, including CRD types and generation scripts. Use this when starting a new "direct" implementation for a GCP resource.
---

# KCC Direct Greenfield Types Implementer

This skill guides the initial scaffolding of *new* (greenfield) KCC "direct" resources, ensuring standardized CRD generation and adherence to project-wide validation patterns.

## Prerequisites
You **must** also apply the standards from the base skill: `.gemini/skills/kcc-direct-base-types-implementer/SKILL.md`.

## Inputs
- `service`: The Google API service name (e.g., `google.cloud.aiplatform.v1`).
- `resource`: The mapping of KCC Kind to GCP Resource (e.g., `VertexAIExampleStore:ExampleStore`).
- `api_version`: The KCC API version (default: `v1alpha1`).

## Workflow

### 1. Add to generate.sh
Locate `apis/<service_short>/<api_version>/generate.sh`. If it doesn't exist, create it following the standard KCC template:
```bash
#!/bin/bash
set -e
go run ../../../tooling/main.go generate-types \
  --service <service> \
  --api-version <group>.cnrm.cloud.google.com/<api_version> \
  --resource <resource>
```

### 2. Generate Types
Run the `generate.sh` script.

### 3. Validate and Enhance Output
Apply the baseline validations from `kcc-direct-base-types-implementer`, plus these greenfield-specific rules:

- **Stability Level**: Add `// +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"`.
- **Field Validation**: Manually add or verify kubebuilder tags:
  - Use `// +kubebuilder:validation:Required` for fields that are mandatory in the GCP API.
  - Use `// +kubebuilder:validation:Optional` for all other fields.
- **Enums**: 
  - Use `*string` for the Go type of proto enum fields (do NOT use custom wrapped string types).
  - Use `// +kubebuilder:validation:Enum=VALUE1;VALUE2` to provide validation in the CRD while keeping the Go type simple.
- **Naming Conventions & Acronyms**:
  - Structs in `*_types.go` must be strictly CamelCased (no snake_case allowed).
  - Ensure GCP acronyms are capitalized correctly in resource and field names (e.g., `APIHub`, `SIP`, `JSON`, `PDF`, `DER`, `CEL`).
  - **No Exceptions**: Do **NOT** add new exceptions to `tests/apichecks/testdata/exceptions/acronyms.txt`. Fix the Go field names instead.
- **Pointers**: Validate that optional or nullable fields (like `Location`, booleans, integers) are correctly typed as pointers (e.g., `*string`, `*bool`) rather than raw values.
- **ETag Exclusion**: Remove `etag` fields from the KRM spec; they are generally disallowed. Do **NOT** add exceptions to `spec_dislike_etag.txt`.
- **KCC References (`missingrefs.txt`)**:
  - Ensure fields representing GCP resources are implemented using the KCC Ref pattern (e.g., `ProjectRef`, `NetworkRef`) rather than plain strings.
  - **No Exceptions**: Do **NOT** add exceptions to `tests/apichecks/testdata/exceptions/missingrefs.txt`. Fix the underlying field types instead.
  - Reuse existing refs from `apis/refs/v1beta1/` (e.g., `Compute`, `BigQueryDataset`). If creating a new ref, put it in a service-specific file (e.g., `apis/dlp/v1alpha1/<resource>_reference.go`).
- **Avoiding Scope Creep**:
  - Do **NOT** modify anything under `apis/<service>/v1beta1/` (especially `generate.sh` scripts) when implementing a `v1alpha1` resource.
  - Do **NOT** add CI presubmits/workflows (e.g., `.github/workflows/ci-presubmit.yaml`).
  - Do **NOT** include the controller implementation (`pkg/controller/direct/...`).
- **Package Collisions & Generation**:
  - Instead of copying raw protos into `mockgcp/`, use the `generate.sh` script with the specific upstream SHA to fetch the protos.
  - Ensure `docs/reports/crd_report.md` and `.csv` are correctly updated and committed after running `make generate`.
  - If validations fail, ensure Go clients are regenerated (`./scripts/generate-go-crd-clients/generate-clients.sh`).

### 4. Journaling
Append any quirks about the proto-to-struct mapping (e.g., field name collisions) to `.gemini/journals/<service>.md` using the format described in the `kcc-agentic-journaler` skill.
