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

### 4. Journaling
Append any quirks about the proto-to-struct mapping (e.g., field name collisions) to `.gemini/journals/<service>.md` using the format described in the `kcc-agentic-journaler` skill.

## Greenfield vs. Brownfield (Pre-existing Resources)

While this skill primarily targets new (greenfield) resources, it is also applied when scaffolding direct types for pre-existing (brownfield) resources currently managed by other controllers (e.g., Terraform or DCL). In these cases, you must satisfy additional strict constraints:

### 1. Schema Compatibility (`dev/tasks/diff-crds`)
- When migrating a pre-existing resource, you **MUST NOT** introduce schema breaking changes or unwanted schema drift.
- Use the `dev/tasks/diff-crds` script to verify that the generated CRD is strictly compatible with the baseline CRD.

### 2. Reference Fields Compatibility
- Standard references (e.g., `k8sv1alpha1.ResourceRef`) contain a `kind` property.
- If the pre-existing/baseline CRD did not define a `kind` property for a reference field, introducing one will trigger validation schema mismatches.
- To maintain compatibility, you must declare a custom reference type (e.g., `InstanceResourceRef`) that omits the `kind` property, and use it in place of standard reference types for those fields.

### 3. Hand-coded Mapper Stubs
- Complex custom structures may not automatically map to/from Protobuf types.
- If the mapper generator (`generate-mapper`) fails or raises type mismatch errors, you must hand-code the mapper functions or stubs in the corresponding direct controller package (e.g., `pkg/controller/direct/<service>/<resource>_mapper.go`) to satisfy signature requirements and keep the build clean.

