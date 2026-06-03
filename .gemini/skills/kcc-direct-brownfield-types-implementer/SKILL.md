---
name: kcc-direct-brownfield-types-implementer
description: Guides the implementation of KRM types and CRD scaffolding for migrating existing resources to "direct" controllers.
---

# KCC Direct Brownfield Types Implementer

This skill provides the mandatory standards for creating the initial KRM types (`_types.go`) and generation scripts (`generate.sh`) when migrating an *existing* (brownfield) resource to the direct controller approach.

## Prerequisites
You **must** also apply the standards from the base skill: `.gemini/skills/kcc-direct-base-types-implementer/SKILL.md`.

## Workflow

### 1. Configure generate.sh
Create or update `apis/<service>/v1alpha1/generate.sh`. Use a versioned proto source path to ensure stability.

```bash
#!/bin/bash
set -o errexit
set -o nounset
set -o pipefail

REPO_ROOT="$(git rev-parse --show-toplevel)"
cd ${REPO_ROOT}/dev/tools/controllerbuilder

# Use the pinned SHA from apis/git.versions or a specific override
PROTO_SHA="<sha>" 
PROTO_OUT="${REPO_ROOT}/.build/googleapis-${PROTO_SHA}.pb"

./generate-proto.sh ${PROTO_SHA} ${PROTO_OUT}

go run . generate-types \
    --service <proto.package> \
    --api-version <service>.cnrm.cloud.google.com/v1alpha1 \
    --resource <Kind>:<ProtoMessage> \
    --proto-source-path ${PROTO_OUT}
```

### 2. Validate and Enhance Output
Apply the baseline validations from `kcc-direct-base-types-implementer`, plus these brownfield-specific rules:

- **Stability Level**: Add the appropriate stability level label (often `alpha` or matching the existing CRD).
  ```go
  // +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
  ```
- **Proto Mapping**: Ensure `+kcc:proto` tags are present on the Spec and ObservedState structs to link them to the GCP API definitions.
- **Use Existing References**: ALWAYS reuse existing resource reference structures that live in `apis/refs/` instead of hand-coding or defining duplicate types.
  * For example, `ProjectRef` (which lives in `apis/refs/v1beta1/project_ref.go`) and other resource reference types should be imported from `github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1` rather than being defined locally in `<kind>_types.go`.
- **Strict Schema Compatibility**: At the initial stage of creating a direct Go type for an existing resource (transitioning from Terraform/DCL), the Go type should be strictly schema-compatible with the existing CRD definition.
  * Do NOT add new fields like `externalRef` or `observedState` under `Status` yet.
  * Run `dev/tasks/diff-crds` to verify schema compatibility and ensure no unintended new fields are introduced.

### 3. Fuzzers
- Create a fuzzer for the mapper to verify that round-trip conversions (FromProto and ToProto) are lossless and correct.
- See the `kcc-direct-controller-implementer` skill for details on implementing fuzzers.

### 4. Registration
- Ensure the new Kind is registered in `apis/<service>/v1alpha1/register.go`.
- Run `dev/tasks/generate-crds` and verify the YAML appears in `config/crds/resources/`.
