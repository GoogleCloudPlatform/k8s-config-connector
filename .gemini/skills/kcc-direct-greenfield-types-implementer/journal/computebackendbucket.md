# ComputeBackendBucket Greenfield/Brownfield Types Journal

- **Date**: Saturday, June 13, 2026
- **Resource**: `ComputeBackendBucket`
- **Group/Version**: `compute.cnrm.cloud.google.com/v1beta1`

## Observations and Lessons Learned

1. **Re-enabling "Unreachable" Types**:
   - During code generation, the controller builder comments out types in `types.generated.go` that it determines to be unreachable because they aren't fully resolved yet.
   - These commented-out types (e.g., `BackendBucket`, `BackendBucketCDNPolicy`, `BackendbucketNegativeCachingPolicy`) are highly valuable because they contain the exact field annotations (`+kcc:proto:field`) and doc comments mapping to the GCP protobuf.
   - We utilized these commented-out definitions to handwrite the properties in `backendbucket_types.go` and restore their exact schemas.

2. **Stability Label for Migrated Resources**:
   - Although following the greenfield skill, `ComputeBackendBucket` is already a stable resource in `v1beta1`.
   - Downgrading its label to `stability-level=alpha` would be a schema regression. Therefore, we explicitly kept/assigned `stability-level=stable` and `tf2crd=true` to maintain 100% backward compatibility.

3. **Reference Reuse**:
   - `ComputeBackendBucket` references a StorageBucket (`bucketRef`).
   - Instead of writing a new local reference structure or using raw/legacy types, we imported and reused the existing `*storagev1beta1.StorageBucketRef` from `apis/storage/v1beta1`.
