# Export Support Journal for Service (ServiceUsage)

## Observations
- The `Service` direct controller (under `serviceusage.cnrm.cloud.google.com`) already had an `Export` method, but its exported `projectRef.external` value prefixed the project ID with `"projects/"`.
- In contrast, the legacy/old controller exported the raw project ID (e.g., `project-${uniqueId}`) without the `"projects/"` prefix.
- This mismatched format resulted in an `_exported_object.diff` file being generated during the test execution.
- Modifying `serviceAdapter.Export` in `pkg/controller/direct/serviceusage/service_controller.go` to output the raw `a.id.Project` (matching the legacy exporter) resolved the mismatch, eliminated the diff file, and achieved perfect alignment between the direct and legacy exporters.
