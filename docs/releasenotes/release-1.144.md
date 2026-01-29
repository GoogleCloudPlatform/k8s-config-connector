# Release 1.144.0

*   Special shout-outs to @justinsb, @xiaoweim for their contributions to this release.

## New Beta Resources (Direct Reconciler):

*   [`TagsLocationTagBinding`](https://cloud.google.com/config-connector/docs/reference/resource-docs/tags/tagslocationtagbinding)
    *   `TagsLocationTagBinding` is promoted to beta and now uses the direct reconciler by default.
    *   Supports tagging of regional resources, including `ArtifactRegistryRepository`, `CloudRun` (`RunJob`, `RunService`), `BigQueryDataset`, `BigQueryTable`, and `StorageBucket`.
    *   `spec.location` should be set to the region of the resource being tagged.

## Reconciliation Improvements

*   [`TagsTagBinding`](https://cloud.google.com/config-connector/docs/reference/resource-docs/tags/tagstagbinding)
    *   Moved to direct reconciliation for improved reliability and performance.
*   [`TagsLocationTagBinding`](https://cloud.google.com/config-connector/docs/reference/resource-docs/tags/tagslocationtagbinding)
    *   Switched to direct reconciliation as the default reconciler.

## New features:

*   Added support for forcing the direct controller default via the label `cnrm.cloud.google.com/default-reconciler: direct` on the CRD. This allows opting in to direct reconciliation for all instances of a resource type.

## Bug Fixes:

*   Fixed a race condition in the Config Connector manager that could occur during high-concurrency reconciliation.
*   Fixed spurious diffs in `TagsLocationTagBinding` caused by project number vs. project ID mismatches.
*   Fixed invalid YAML in the `RunService` resource sample.
*   Updated generated file markers in YAML files to prevent accidental manual edits.
