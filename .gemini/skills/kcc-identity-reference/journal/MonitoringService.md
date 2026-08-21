# MonitoringService Identity and Reference Journal

## Observations

- `MonitoringService` has its KRM schema defined in `apis/monitoring/v1beta1/service_types.go`, allowing us to implement `identity.Resource` on the typed pointer `*MonitoringService` directly.
- The GCP resource name format for a Monitoring Service is `projects/{project}/services/{service}`.
- The GCP Monitoring v3 API uses a `v3/` path segment in its URLs (e.g., `https://monitoring.googleapis.com/v3/projects/my-project/services/my-service`). To support this versioned URL format seamlessly, we added `v3/` prefix normalization in the `FromExternal()` method in `monitoringservice_identity.go`.
- The resource's `Status` struct does not contain a `Name` or `ExternalRef` field, so we did not perform any status-based cross-checks in `GetIdentity`, strictly adhering to the schema-preservation rule.
- We registered `//monitoring.googleapis.com/projects/{}/services/{}` as an ignored template exception in `pkg/gcpurls/registry_test.go` since it is not documented in CAIS, preventing `TestRegisteredTemplatesMatchCAI` from failing.
