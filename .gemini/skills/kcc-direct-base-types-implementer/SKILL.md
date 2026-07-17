---
name: kcc-direct-base-types-implementer
description: Base skill containing shared standards for all KCC direct resource types (both greenfield and brownfield).
---

# KCC Direct Base Types Implementer

This skill provides the mandatory baseline standards that apply to *all* new KRM types (`_types.go`) for direct resources in Config Connector, regardless of whether they are greenfield or brownfield migrations.

## Shared Standards for <kind>_types.go

After running the generator (via `generate.sh`), you must verify and enforce the following baseline requirements on the resulting `_types.go` file:

- **Copyright**: The file must start with `// Copyright 2026 Google LLC`.
- **CRD Labels**: Include at least these two labels in the type definition:
  ```go
  // +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true"
  // +kubebuilder:metadata:labels="cnrm.cloud.google.com/system=true"
  ```
  *(Note: See greenfield/brownfield skills for the correct `stability-level` label to append.)*
- **Status Fields**: `status.observedGeneration` must be exactly `*int64`.
- **Reference Fields**: Ensure that fields referencing other GCP/KCC resources are implemented as proper KCC reference fields (e.g., using `pubsubv1beta1.PubSubTopicRef` or `refsv1beta1.KMSCryptoKeyRef`), following the `Ref` suffix naming convention. You **MUST NOT** add new exceptions to `tests/apichecks/testdata/exceptions/missingrefs.txt`. All reference-like fields must be implemented as proper references.
- **Reference Types Location**: Whenever a reference type (e.g. `<Kind>Ref` implementing `refsv1beta1.Ref`) is needed, it must **always** be defined and implemented in its own separate file named `<kind>_reference.go` (e.g., `filestorebackup_reference.go`) rather than inside `_types.go`. This keeps the main type definitions clean and isolated from reference resolution boilerplate.

- **KCC Proto Annotations**:
  To enable auto-generation of mappers, you must add the correct "kcc:proto" annotations to Go structs in `_types.go`:
  * The Spec struct must be annotated with `// +kcc:spec:proto=<proto_type>` (e.g. `// +kcc:spec:proto=google.cloud.compute.v1.ServiceAttachment`).
  * The ObservedState struct (if present) must be annotated with `// +kcc:observedstate:proto=<proto_type>`.
  * The Status struct (if there is no separate ObservedState struct) must be annotated with `// +kcc:status:proto=<proto_type>`.
  * Nested/referenced helper structs (both in Spec and Status) must be annotated with `// +kcc:proto=<proto_sub_type>` (e.g. `// +kcc:proto=google.cloud.compute.v1.ServiceAttachmentConnectedEndpoint`).

