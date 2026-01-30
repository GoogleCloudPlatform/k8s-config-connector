# Release 1.144.0

*   Special shout-outs to @acpana, @anhdle-sso, @cheftako, @codebot-robot, @himanikh, @justinsb, @katrielt, @xiaoweim, @yuwenma for their contributions to this release.

## New Beta Resources (Direct Reconciler):

*   [`TagsLocationTagBinding`](https://cloud.google.com/config-connector/docs/reference/resource-docs/tags/tagslocationtagbinding)
    *   `TagsLocationTagBinding` is promoted to beta and now uses the direct reconciler by default.
    *   Supports tagging of regional resources, including `ArtifactRegistryRepository`, `CloudRun` (`RunJob`, `RunService`), `BigQueryDataset`, `BigQueryTable`, and `StorageBucket`.
    *   `spec.location` should be set to the region of the resource being tagged.

## Reconciliation Improvements

*   [`TagsLocationTagBinding`](https://cloud.google.com/config-connector/docs/reference/resource-docs/tags/tagslocationtagbinding)
    *   Switched to direct reconciliation as the default reconciler.

## Bug Fixes:



*   Fixed spurious diffs in `TagsLocationTagBinding` caused by project number vs. project ID mismatches.
