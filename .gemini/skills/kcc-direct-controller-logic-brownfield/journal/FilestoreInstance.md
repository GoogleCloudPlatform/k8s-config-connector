# FilestoreInstance Direct Controller Logic (Brownfield) Implementation Journal

## Observations
1. **Dynamic Multi-Controller Testing**: Verified that registering `k8s.ReconcilerTypeDirect` under `SupportedControllers` in `static_config.go` correctly triggers the unified test framework to run and record both legacy DCL and new direct controller execution paths.
2. **Comparing Diff Output**: The generated `_final_object.diff` and `_http_mock.diff` files were examined to confirm semantic equivalence between the controllers. Differences were only in expected areas (such as externalRef presence and observedGeneration, plus direct controller using the newer `/v1` endpoint instead of DCL's `/v1beta1`).

## Lessons Learned & Best Practices from PR Feedback
1. **Timing of Reference Normalization**: `NormalizeReferences()` should only be called at the very beginning of the `Create()` and `Update()` execution paths. It must not be executed inside `AdapterForObject()`.
2. **KRM-Based Desired State (Brownfield)**: The Adapter should preserve the unmodified KRM object as its `desired` state. This makes it possible to inspect which KRM spec fields are set or unset by the user when running reconciliations.
3. **Assigning Default Values for Unspecified Spec Fields**: Unspecified spec fields in the KRM desired state (such as tier, description, nfs export options, or network subfields) should have their values copied from the `actual` state before calling `common.CompareProtoMessage(clonedDesired, actual, common.BasicDiff)`. This prevents the comparator from incorrectly flagging unset user fields that got populated with default values by the GCP server as active drift.
4. **Label Cleanup and Managed Keys**: To clean up KRM-style labels (containing `/`) and automatically attach system-managed labels (such as `managed-by-cnrm: true`), use the standard `label.GCPLabels` helper during label propagation in both creation and update paths.
5. **Reusing LRO Wait Results**: Since Google Cloud client library LRO `Wait(ctx)` operations return the finalized protobuf representation of the created or updated resource, use that returned value directly as the latest state instead of triggering a redundant GET call.
6. **No Impossible Deleted Checks**: Do not add defensive `id.Instance == ""` checks at the beginning of `Delete()` as `Find()` would have already filtered out non-existent objects.
