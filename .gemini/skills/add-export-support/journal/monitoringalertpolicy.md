# MonitoringAlertPolicy Export Support Journal

## Overview
Validated and verified the existing export support for the `MonitoringAlertPolicy` resource in Config Connector.

## Key Observations and Verification Findings

1. **AdapterForURL Implementation**:
   - The model `alertPolicyModel` in `pkg/controller/direct/monitoring/monitoringalertpolicy_controller.go` correctly implements `AdapterForURL`.
   - It validates the URL prefix, parses the external URI using the `MonitoringAlertPolicyIdentity`'s `FromExternal` parser, and instantiates the adapter with the correct client.

2. **Export Method Structure & Safety**:
   - The adapter's `Export` method maps the retrieved GCP/proto state using `MonitoringAlertPolicySpec_FromProto`.
   - It populates the `ResourceID` on the spec:
     ```go
     obj.Spec.ResourceID = direct.LazyPtr(a.id.AlertPolicy)
     ```
   - It adheres strictly to the critical ordering requirement specified in `SKILL.md`:
     ```go
     u.Object = uObj
     u.SetName(a.id.AlertPolicy)
     u.SetGroupVersionKind(krm.MonitoringAlertPolicyGVK)
     ```
     Assigning `u.Object` first avoids discarding metadata when constructing the `unstructured.Unstructured` object.
   - For project binding, since `MonitoringAlertPolicy` is a non-reference-bound resource (it uses container annotations rather than `spec.projectRef`), it correctly sets the project ID annotation using:
     ```go
     export.SetProjectID(u, a.id.Project)
     ```
   - It applies resource labels using:
     ```go
     export.SetLabels(u, a.actual.UserLabels)
     ```

3. **E2E Export Integration & Testing**:
   - `MonitoringAlertPolicy` is fully registered in `tests/e2e/export.go` using the `resolveCAISURI` helper under the GVK `monitoring.cnrm.cloud.google.com/MonitoringAlertPolicy`.
   - The E2E tests were executed against `mockgcp` using `./dev/tasks/run-e2e`:
     ```bash
     RUN_TESTS="TestAllInSeries/fixtures/monitoringalertpolicy" ./dev/tasks/run-e2e
     ```
   - All 4 test fixtures (`monitoringalertpolicy`, `monitoringalertpolicysqlboolean`, `monitoringalertpolicysqlhourly`, `monitoringalertpolicysqlrowcount`) executed and exported successfully.
   - The exported KRM objects match `_exported.yaml` perfectly without producing any `_exported_object.diff` files.
   - The fallback/old controller correctly logged that the legacy exporter does not support `MonitoringAlertPolicy` CAI URLs (which is expected and handled gracefully by the test harness).
