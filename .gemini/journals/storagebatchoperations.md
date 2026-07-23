# Storage Batch Operations Journal

### [2026-07-02] StorageBatchOperationsJob Type Implementation and Reference Mapping
- **Context**: Implementing initial KRM types, CRD, and IdentityV2 for `StorageBatchOperationsJob` under `apis/storagebatchoperations/v1alpha1`.
- **Problem**: The generated structures for `RewriteObject` and `BucketList_Bucket` used raw strings for GCP resource identifiers (`kms_key` and `bucket`). To maintain high KCC standards, we needed to map these to KCC resource references (`refsv1beta1.KMSCryptoKeyRef` and `storagev1beta1.StorageBucketRef`).
- **Solution**: We hand-wrote the types `RewriteObject` and `BucketList_Bucket` inside `storagebatchoperationsjob_types.go` with the correct `// +kcc:proto=...` annotations, replacing the string fields with proper KCC reference structures. The generator automatically skipped creating these types in `types.generated.go` and resolved all dependencies correctly.
- **Impact**: Ensures that `StorageBatchOperationsJob` uses KRM-native references for storage buckets and Cloud KMS crypto-keys without requiring custom mapper adjustments or causing type conflicts during code generation.
