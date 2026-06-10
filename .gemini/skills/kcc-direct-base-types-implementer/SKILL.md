---
name: kcc-direct-base-types-implementer
description: Base skill containing shared standards for all KCC direct resource types (both greenfield and brownfield).
---

# KCC Direct Base Types Implementer

This skill provides the mandatory baseline standards that apply to *all* new KRM types (`_types.go`) for direct resources in Config Connector, regardless of whether they are greenfield or brownfield migrations.

## Shared Standards for <kind>_types.go

After running the generator (via `generate.sh`), you must verify and enforce the following baseline requirements on the resulting `_types.go` file:

- **Copyright**: New files must start with `// Copyright 2026 Google LLC`. Do NOT modify the copyright year on existing files.
- **CRD Labels**: Include at least these two labels in the type definition:
  ```go
  // +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
  // +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
  ```
  *(Note: See greenfield/brownfield skills for the correct `stability-level` label to append.)*
- **Status Fields**: `status.observedGeneration` must be exactly `*int64`.
