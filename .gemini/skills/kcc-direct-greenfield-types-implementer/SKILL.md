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
- **Terminology**: Replace any occurrences of "GCP" with "Google Cloud" in all comments, field descriptions, and docstrings.

### 4. Journaling
Append any quirks about the proto-to-struct mapping (e.g., field name collisions) to `.gemini/journals/<service>.md` using the format described in the `kcc-agentic-journaler` skill.
