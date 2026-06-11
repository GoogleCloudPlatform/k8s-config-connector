# Journal: ComputeSSLPolicy Transition to Direct Type

During the implementation of direct types for `ComputeSSLPolicy`, the following insights and notes were collected:

1. **Strict Schema Compatibility Constraints**:
   - `ComputeSSLPolicy` is a global resource and does not have `ProjectRef` or `Location` fields in its original spec. Standard direct generation templates add these automatically, but we had to explicitly exclude them to avoid any CRD schema modifications.
   - The output of `diff-crds` confirmed that our type is perfectly schema compatible (other than standard description updates and making `spec` a required block as expected for direct types).

2. **Capitalization and Near Misses**:
   - In `pb.SslPolicy`, the minimum TLS version field is represented as `min_tls_version`, which converts to `MinTLSVersion` in the proto-generated Go struct.
   - Our original KRM Go field was `MinTlsVersion`. The mapping generator flagged this as a near miss: `"MinTLSVersion" vs "MinTlsVersion"`.
   - By renaming the KRM Go struct field to `MinTLSVersion` while retaining the original JSON tag `json:"minTlsVersion,omitempty"`, we resolved the mismatch without any changes to the CRD schema. This fully automated the mapper generation and avoided the need for a manual mapping implementation.

3. **Status Field Mapping**:
   - We registered the status structure with `// +kcc:status:proto=google.cloud.compute.v1.SslPolicy` on `ComputeSSLPolicyStatus`.
   - This automatically generated mappers for `creationTimestamp`, `enabledFeatures`, `fingerprint`, and `selfLink`.
