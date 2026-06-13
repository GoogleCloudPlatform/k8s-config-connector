# Compute Service Journals

### 2026-06-13 Implement direct KRM types for ComputeBackendBucketSignedURLKey
- **Context**: Implementing direct KRM types and generate.sh configuration for `ComputeBackendBucketSignedURLKey`. (Issue #9819)
- **Problem**: `ComputeBackendBucketSignedURLKey` is a global resource representing a SignedUrlKey of a BackendBucket. The target protobuf resource is `google.cloud.compute.v1.SignedUrlKey` which is simple and contains only the `key_name` and `key_value` fields. The key value is sensitive and was previously implemented in the Terraform controller with a `keyValue` secret structure.
- **Solution**: 
  - Added the mapping `--resource ComputeBackendBucketSignedURLKey:SignedUrlKey` to `generate.sh` of the `compute` API group.
  - Hand-wrote `signedurlkey_types.go` using `secret.Legacy` for the sensitive `keyValue` field and `computev1beta1.ComputeBackendBucketRef` for `backendBucketRef`.
  - Added a new reference helper struct `ComputeBackendBucketRef` in a new file `computebackendbucket_reference.go` under `v1beta1` to handle normalization of references to `ComputeBackendBucket`.
- **Impact**: Enables standard KRM direct type generation, deepcopy logic, CRD building, and mappers for `ComputeBackendBucketSignedURLKey`.
