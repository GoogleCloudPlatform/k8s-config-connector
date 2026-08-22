# Journal: AlloyDBInstance Identity & References Migration

## Summary of Changes
Refactored the existing `AlloyDBInstance` identity and reference logic to conform perfectly to the `identity.IdentityV2` / `refs.Ref` canonical pattern.

Specifically:
- Updated `getIdentityFromAlloyDBInstanceSpec` in `apis/alloydb/v1beta1/alloydbinstance_identity.go` to accept a generic `client.Object` instead of a typed `*AlloyDBInstance`.
- Implemented a clean type-switch inside `getIdentityFromAlloyDBInstanceSpec` to handle both typed `*AlloyDBInstance` and `*unstructured.Unstructured` objects using `unstructured.NestedMap` for extracting `spec.clusterRef` on unstructured objects.
- Simplified `Normalize` fallback logic in `apis/alloydb/v1beta1/alloydbinstance_reference.go` to directly pass the unstructured object `u` to `getIdentityFromAlloyDBInstanceSpec` without needing the boilerplate conversion via `runtime.DefaultUnstructuredConverter.FromUnstructured`.
- Cleaned up the unused `k8s.io/apimachinery/pkg/runtime` import in `alloydbinstance_reference.go`.

## Key Learnings
- **Dynamic Field Resolution**: When a resource has dependent references (like `spec.clusterRef` in `AlloyDBInstance`), using a type-switch inside the spec-identity-resolver allows us to safely and efficiently support both typed and unstructured objects without requiring `Normalize` to unpack the whole object if we're dealing with unstructured payloads.
- **Reference Resolution Helpers**: Helpers like `refs.ResolveAlloyDBCluster` expect `client.Object` and a structured pointer. By dynamically parsing unstructured specs into smaller helper structs (like `refs.AlloyDBClusterRef`), we maintain clean separation of concerns and robust error handling.
