# MemorystoreInstance Identity and Reference Migration

## Observations

1.  **Circular Dependencies**: `MemorystoreInstance` is referenced by `ComputeForwardingRule`. `MemorystoreInstance` itself references `ComputeNetwork`. This created a circular dependency when trying to move `MemorystoreInstanceRef` to the `memorystore/v1beta1` package.
2.  **Shared Reference Location**: To resolve the circular dependency, BOTH `MemorystoreInstanceRef` and `MemorystoreInstanceIdentity` were moved to a new unversioned package `apis/memorystore/refs/`. This avoids versioning the identity and prevents circular dependencies while keeping the files scoped to the `memorystore` service instead of cluttering the global `apis/refs` package.
3.  **Identity Implementation**: The `GetIdentity` method (which depends on the resource struct) stays in the versioned `apis/memorystore/v1beta1` package and imports the identity struct from `apis/memorystore/refs/`.

## Shortcomings in SKILL.md

The current skill assumes that BOTH identity and reference files should always be moved to the resource package. However, for resources that are part of a circular dependency (e.g., cross-service references), the reference struct and identity should be moved to a service-specific `refs` package (e.g., `apis/memorystore/refs/`) instead of the versioned resource package.
