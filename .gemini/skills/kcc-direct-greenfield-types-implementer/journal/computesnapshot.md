### 2026-06-13 Scaffolding direct KRM types and generate.sh for ComputeSnapshot
- **Context**: Implementing direct KRM types and generate.sh configuration for `ComputeSnapshot` ([Issue #9887](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/9887)).
- **Work Done**:
  1. Registered the `ComputeSnapshot:Snapshot` resource in the direct KRM types generator script (`apis/compute/v1beta1/generate.sh`).
  2. Created the `apis/compute/v1beta1/snapshot_types.go` file.
  3. Defined nested resource schemas like `CustomerEncryptionKey` and proper KCC resource references like `ComputeDiskRef` for `SourceDiskRef` instead of using legacy `v1alpha1.ResourceRef`.
  4. Ran the generator to successfully output KRM types, deepcopies, GVK registration, and mappers.
  5. Regenerated Kubernetes CRDs, static configuration maps, and documentation.
- **Impact**: Sets up the direct types scaffold for ComputeSnapshot enabling transitioning it from the Terraform controller to the Greenfield direct controller.
