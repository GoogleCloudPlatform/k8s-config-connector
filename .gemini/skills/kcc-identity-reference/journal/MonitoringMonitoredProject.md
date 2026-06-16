# MonitoringMonitoredProject Identity and Reference Journal

## Observations

- `MonitoringMonitoredProject` is a metrics-scoped resource under GCP's monitoring API.
- Its GCP resource URL format is `locations/global/metricsScopes/{metricsScope}/projects/{project}`.
- It is registered in the `apis/monitoring/v1beta1` group, and its `Spec` struct has `MetricsScope string` and optional `ResourceID *string` fields.
- Its `Status` struct does not have an `ExternalRef` or `Name` field, so no status crosscheck was performed in `GetIdentity`.
- Since it is not present in CAI's definitions, its URL template has been added to the exclusion list `ignoredTemplates` in `pkg/gcpurls/registry_test.go` to prevent `TestRegisteredTemplatesMatchCAI` from failing.
