# Journal Entry - MonitoringGroup

## Learnings & Observations

1. **Self-Reference / Duplicate Declarations**:
   The baseline monitoring API packages already contained local types like `MonitoringGroupRef` within `uptimecheckconfig_types.go`. When adding `monitoringgroup_reference.go`, we had to remove the duplicate struct declaration from `uptimecheckconfig_types.go` to prevent compile errors.

2. **Custom Reference & Ignored Fields mapping**:
   Since the `Group` message in GCP Monitoring only has string fields (like `parent_name`) for references, and the KRM CRD has `parentRef` and `projectRef` without a `kind` field in the original schema, we hand-coded the mappings in `mappers.go` to bridge `MonitoringGroupRef` to `parent_name` and ignored `projectRef` in mapping as the `Group` proto does not have a project field (the project is supplied as the parent in request URLs).

3. **RegisterKRMSpecFuzzer**:
   If a resource doesn't have custom mappings for status fields from proto (and thus status from/to proto mappers are nil), always register the fuzzer using `fuzztesting.RegisterKRMSpecFuzzer(fuzzer)` instead of `fuzztesting.RegisterKRMFuzzer(fuzzer)`. This prevents the shared fuzzer test suite from executing status fuzzing on a nil mapper, preventing segmentation faults.

4. **TestRegisteredTemplatesMatchCAI Exception**:
   Because `monitoring.googleapis.com/Group` is not tracked in Cloud Asset Inventory (CAI), we registered its normalized URL pattern `"//monitoring.googleapis.com/projects/{}/groups/{}"` as an exception in `pkg/gcpurls/registry_test.go` to avoid failing template matching assertions.
