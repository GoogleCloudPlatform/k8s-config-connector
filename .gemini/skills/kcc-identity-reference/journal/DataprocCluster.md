# DataprocCluster Identity and Reference Migration

## Observations

- `DataprocCluster` is currently a DCL-based resource and its types are still in `pkg/clients/generated/apis/dataproc/v1beta1/`.
- However, it already had `cluster_identity.go` and `cluster_reference.go` in `apis/dataproc/v1beta1/` using an old pattern.
- I renamed these files to `dataproccluster_identity.go` and `dataproccluster_reference.go` and updated them to the `IdentityV2` and `Ref` pattern.
- Discrepancy found: The old `NormalizedExternal` was looking for `spec.region` in the unstructured object, but the CRD uses `spec.location`. I corrected this to use `refs.GetLocation(obj)`.
- GCP URL format for `DataprocCluster` uses `regions/{{region}}`, while the KRM field is `location`. This mapping is handled in `getIdentityFromDataprocClusterSpec`.
- I had to fix `pkg/controller/direct/bigqueryconnection/connection_controller.go` which was calling the old `NormalizedExternal` method on `DataprocClusterRef`.

## Lessons Learned

- Even if a resource is not yet "direct", it may have identity/reference logic in `apis/` if it's referenced by other direct resources.
- When renaming methods like `NormalizedExternal` to `Normalize`, check for usages in direct controllers of other resources.
- Always verify the KRM field name for location/region against the CRD, as DCL resources can be inconsistent.
