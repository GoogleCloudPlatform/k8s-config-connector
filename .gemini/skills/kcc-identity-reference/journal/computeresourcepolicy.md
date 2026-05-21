# ComputeResourcePolicy Identity and Reference Pattern Journal

## Observations

1. **Region vs Location**: `ComputeResourcePolicy` uses `Region` instead of `Location`. This is common for regional Compute resources.
2. **GCP URL format**: The GCP URL for `ResourcePolicy` is `projects/{project}/regions/{region}/resourcePolicies/{resourcePolicy}`. The path segment `resourcePolicies` is camelCase. `gcpurls.Template` handles this literal part correctly.
3. **IdentityV2 Implementation**: The existing implementation already used `gcpurls.Template` but with uppercase template variables (`{Project}`). I updated them to lowercase (`{project}`) to match the canonical example in `SKILL.md`.
4. **GetIdentity Refactoring**: Refactored `GetIdentity` to use `getIdentityFromComputeResourcePolicySpec` to match the canonical pattern and ensure consistency.
5. **Reference Docstring**: Updated the `External` field docstring in `ComputeResourcePolicyRef` to use `{{}}` and lowercase variables, e.g., `projects/{{project}}/regions/{{region}}/resourcePolicies/{{resourcepolicy}}`.
6. **Cross-checking Identity**: `ComputeResourcePolicy` has `status.externalRef`, so I used it for cross-checking in `GetIdentity`.
