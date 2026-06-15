### [2026-06-15] StorageBatchOperationsJob Greenfield Type Scaffolding
- **Context**: Greenfield scaffolding of KRM types, CRD, and IdentityV2 for `StorageBatchOperationsJob` (v1alpha1).
- **Problem**: The proto definition contains nested references (e.g., `RewriteObject.kms_key` and `BucketList.Bucket.bucket`) that need to be resolved to proper KRM reference objects instead of standard string fields.
- **Solution**: We defined `BucketList_Bucket` and `RewriteObject` manually in `job_types.go` referencing `storagev1beta1.StorageBucketRef` and `refsv1beta1.KMSCryptoKeyRef` with `// +kcc:proto=...` tags. The controller builder generator successfully skipped generating these types, resolving references automatically in the CRD schema.
- **Impact**: Ensures strict type-safety and correct KCC reference behavior for downstream resources while avoiding schema exclusions or bypasses.
