# DataprocCluster Direct Migration Journal

During the direct migration validation for `DataprocCluster`, we encountered and resolved several issues.

## Observations & Root Causes
- **System-Assigned Labels Diff:** Dataproc automatically assigns labels starting with `goog-` (e.g. `goog-dataproc-autozone`, `goog-dataproc-cluster-name`, etc.) during cluster creation. During `Update` (which handles takeover comparison), the `desired` specification built by the Direct controller lacks these labels, causing a mismatch/diff with the `actual` cluster on GCP. Consequently, the Direct controller would issue `PATCH` requests to try and clear these system-assigned labels.
- **Non-deterministic migration JSON diff sorting:** The migration test runner `tests/e2e/migration_test.go` records Raw Diffs as they are discovered by concurrent/asynchronous reconciliations. This causes the array inside `_migration_diffs.json` to have an arbitrary and flaky order depending on which controller reconciles faster.

## Resolutions
- **Preserving system-assigned labels:** We updated `Update` in `pkg/controller/direct/dataproc/dataproccluster_controller.go` to import the `"strings"` package and copy any system-assigned labels (keys starting with `"goog-"`) from `a.actual.Labels` into `cluster.Labels` before performing the comparison. This cleanly prevents any false diffs or unexpected write operations on system-assigned labels.
- **Deterministic migration diff sorting:** We updated `formatDiffsRaw` in `tests/e2e/migration_test.go` to sort the array of raw diffs by `Resource` and `Controller` before marshaling it to JSON. This guarantees deterministic diff output order, preventing flaky test failures.
