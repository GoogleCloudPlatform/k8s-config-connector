# Assured Workloads Tribal Knowledge

### [2026-09-08] AssuredWorkloadsWorkload Direct Controller Implementation Quirks
- **Context**: Implementing direct controller for `AssuredWorkloadsWorkload` (Phase 2, issue #12005 / #12182).
- **Problem**:
  1. **Regional Endpoint Requirement**: Assured Workloads API requires connecting to regional endpoints (e.g. `us-central1-assuredworkloads.googleapis.com:443`). Without regional endpoint configuration, requests fail or are rejected.
  2. **Child Resource Deletion Requirement**: Calling `DeleteWorkload` on an Assured Workloads workload will immediately fail with `400 FAILED_PRECONDITION: contains projects or other resources that are not deleted` if any child resources (projects or folders) exist under the workload.
  3. **Eventual Consistency on Deletion**: Even after child projects and folders are deleted via Cloud Resource Manager API, there is a delay (typically 1-3 seconds) before the Assured Workloads backend recognizes the deletion.
  4. **Update Mask Path Prefixing**: Assured Workloads `UpdateWorkload` API requires field paths in `update_mask` to be prefixed with `workload.` (e.g., `workload.display_name`, `workload.labels`).
  5. **Re-reconciliation False Diff in Labels**: In KCC, resource labels are on `metadata.labels` rather than `spec.labels`. When translating `maskedActualSpec` to proto, proto labels would be nil, causing false diffs during re-reconciliation against `desired.Labels`.
  6. **Service-Generated Workload ID**: Workload IDs are server-generated and cannot be chosen by the user. Controller must acquire the workload ID from `CreateWorkload` LRO response and save it to `.status.externalRef`.
- **Solution**:
  1. Configured regional endpoint dynamically in `model.client()`: `option.WithEndpoint(fmt.Sprintf("%s-assuredworkloads.googleapis.com:443", location))`.
  2. In `Delete()`, query `a.actual.Resources` and proactively delete child projects (`DeleteProject`) and folders (`DeleteFolder`) before calling `DeleteWorkload`.
  3. Implemented retry polling with exponential backoff in `Delete()` to absorb the eventual consistency delay until `DeleteWorkload` succeeds. Also added a filter in `tests/e2e/httplog.go` `RemoveExtraEvents` to discard transient 400 responses during delete polling.
  4. Explicitly formatted update mask paths with `workload.` prefix in `Update()`.
  5. In `Update()`, explicitly copied `maskedActual.Labels = a.actual.Labels` before running `common.DiffForTopLevelFields` to prevent false diffs.
  6. Implemented service-generated ID handling according to `kcc-direct-service-generated-id` standards.
- **Impact**: Crucial for anyone maintaining `AssuredWorkloadsWorkload` or implementing other Assured Workloads direct resources.
