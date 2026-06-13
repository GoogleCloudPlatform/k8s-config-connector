# ComputeBackendServiceSignedURLKey Journal

## Key Observations and Learnings
- **Strict Schema Compatibility:** The baseline `ComputeBackendServiceSignedURLKey` CRD defines a required field `spec.keyValue` which maps to a secret value/source struct `BackendservicesignedurlkeyKeyValue`. However, the protobuf message `SignedUrlKey` defines `key_value` as a simple string.
- **Type Mismatch Resolution:** To resolve this type mismatch and prevent compilation errors in the auto-generated mapper, we implemented handcoded mapping functions `ComputeBackendServiceSignedURLKeySpec_v1alpha1_FromProto` and `ComputeBackendServiceSignedURLKeySpec_v1alpha1_ToProto` in `computebackendservicesignedurlkey_mappings.go`. The mapper generator automatically skips generating the conflicted versions once the signature is matched.
- **Canonical Reference Types:** `ComputeBackendServiceRef` was imported from the `computev1beta1` package to match KCC standards, while `ProjectRef` was imported from `refsv1beta1`.
- **Validation Constraints (`oneOf`):** For the custom validation blocks like the legacy `value`/`valueFrom` `oneOf` rules, we configured the `scripts/add-validation-to-crds/parse-crds.go` file to inject the `legacyRefRule` for `ComputeBackendServiceSignedURLKey` when parsing the `value,valueFrom` signature.
