# DataprocCluster Direct Controller Journal

## Implementation Details

- **Empty LRO Response Workaround**: The GCP Dataproc `UpdateCluster` LRO response is empty (`google.protobuf.Empty`), and `CreateCluster` may return a resource state before post-creation metrics or instance names are populated. To prevent status fields (like state, metrics, and instance names) from being cleared on updates, we refactored the adapter to fetch the latest fully-populated resource via `GetCluster` immediately after any create or update operation completes.
- **Diff Comparison with `CompareProtoMessageStructuredDiff`**: Replaced custom `compareCluster` and `expandUpdateMask` with standard `common.CompareProtoMessageStructuredDiff`. We also sort the resulting fieldmask paths for determinism in HTTP requests and golden fixtures.
- **Propagate Metadata Labels**: Added explicit labels propagation in both `Create` and `Update` methods of `dataproccluster_controller.go` to preserve system/custom labels (such as `managed-by-cnrm: true`) on GCP.
- **Delete Not Found Handling**: Handled already-deleted resource check in `Delete` using `direct.IsNotFound(err)` to ensure idempotency.
