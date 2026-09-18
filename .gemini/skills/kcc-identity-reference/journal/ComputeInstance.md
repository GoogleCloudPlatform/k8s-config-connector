# ComputeInstance Identity and Reference Journal

## Observations & Learnings

- **Readiness Checks in Fallbacks Prevent Premature Reference Resolution:** We observed that the original `InstanceRef.Normalize` implementation resolved references to `ComputeInstance` from spec fields even if the `ComputeInstance` was not yet ready or created in GCP. By integrating a standard `isResourceReady` check of `status.conditions`, we successfully prevent cascading lookup failures and incorrect dependency resolution during reconciliation of resources depending on `ComputeInstance`.
- **Package-Scoped helper reuse:** We identified that `isResourceReady` is declared at the package level in `computerouternat_reference.go`. Rather than declaring it again in `computeinstance_reference.go` (which causes symbol redeclared compiler errors), we reused the package-level helper to keep the implementation concise.
- **No Schema Changes:** The schema and configuration of `ComputeInstance` was preserved completely untouched, keeping backward compatibility and adhering strictly to safety protocols.
