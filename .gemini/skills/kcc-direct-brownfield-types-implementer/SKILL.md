---
name: kcc-direct-brownfield-types-implementer
description: Guides the implementation of KRM types and CRD scaffolding for new "direct" resources.
---

# KCC Direct Types Implementer

## Overview
This skill provides the mandatory standards for creating the initial KRM types (`_types.go`) and generation scripts (`generate.sh`) for a new direct resource in Config Connector.

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

### 2. Standards for <kind>_types.go
After running the generator, verify the `_types.go` file meets these requirements:

- **Copyright**: Must be `// Copyright 2026 Google LLC`.
- **CRD Labels**: Include exactly these labels in the type definition:
  ```go
  // +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
  // +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
  // +kubebuilder:metadata:labels="cnrm.cloud.google.com/stability-level=alpha"
  ```
- **Proto Mapping**: Ensure `+kcc:proto` tags are present on the Spec and ObservedState structs to link them to the GCP API definitions.
- **Status Fields**: `status.observedGeneration` must be an `*int64`.
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
