# MemorystoreInstance Identity and Reference Migration

## Observations

1.  **Circular Dependencies**: `MemorystoreInstance` is referenced by `ComputeForwardingRule`. `MemorystoreInstance` itself references `ComputeNetwork`. This created a circular dependency when trying to move `MemorystoreInstanceRef` to the `memorystore/v1beta1` package.
2.  **Shared Reference Location**: To resolve the circular dependency, `MemorystoreInstanceRef` was moved to `apis/refs/v1beta1/` (package `v1beta1`) instead of the resource's own package. This follows the pattern of other shared references like `AlloyDBClusterRef` and `SQLInstanceRef`.
3.  **Identity Location**: `MemorystoreInstanceIdentity` and the `GetIdentity` method were successfully moved to `apis/memorystore/v1beta1/memorystoreinstance_identity.go`, following the `IdentityV2` pattern.
4.  **Import Cleanups**: Moving the identity and reference required updating several files in `apis/` and `pkg/controller/direct/`, and removing the old `apis/refs/` (package `refs`) which became empty and invalid.

## Shortcomings in SKILL.md

The current skill assumes that BOTH identity and reference files should always be moved to the resource package. However, for resources that are part of a circular dependency (e.g., cross-service references), the reference struct should stay in or be moved to a shared package like `apis/refs/v1beta1/`.
