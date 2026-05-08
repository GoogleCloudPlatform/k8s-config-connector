---
name: kcc-direct-types-implementer
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

### 3. Registration
- Ensure the new Kind is registered in `apis/<service>/v1alpha1/register.go`.
- Run `dev/tasks/generate-crds` and verify the YAML appears in `config/crds/resources/`.
