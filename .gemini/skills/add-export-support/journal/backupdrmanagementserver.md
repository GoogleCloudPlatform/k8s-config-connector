# Export Support Journal for BackupDRManagementServer

## Observations
- `BackupDRManagementServer` is a `v1alpha1` direct KCC resource.
- Its identity format is `projects/{project}/locations/{location}/managementServers/{managementserver}` and is managed by `ManagementServerIdentityFormat` under the `apis/backupdr/v1alpha1` package.
- In `backupdrmanagementserver_controller.go`, we implemented `AdapterForURL` to support parsing external CAI URIs directly.
- In the `Export` method:
  - We updated the unstructured assignment sequence so that `u.Object = uObj` is set prior to name and GVK fields. This prevents unstructured metadata from being silently overwritten and wiped out.
  - We mapped the resource's `spec.projectRef.external = a.id.Project` and correctly avoided `export.SetProjectID(...)` because the resource is reference-bound to projects.
  - We set the `ResourceID` in Spec using the short name `a.id.ManagementServer` and set `u.SetName(...)` using that short name.
  - We called `export.SetLabels(u, a.actual.Labels)` to export resource labels to Kubernetes object metadata labels.
- Added a case to `tests/e2e/export.go` to handle `BackupDRManagementServer` GVK using `resolveCAISURI(h, obj)`.
- Verified and generated the corresponding `_exported.yaml` files under the `backupdrmanagementserver` test fixture data folder.
- Ran tests against mockgcp and envtest successfully. The E2E export tests passed with no diffs.
