When implementing the round-trip KRM fuzzer for MonitoringMonitoredProject, we observed:
1. `MonitoringMonitoredProject` maps its Spec fields (`metricsScope` and `resourceID`) to the single GCP proto field `.name` under the pattern `locations/global/metricsScopes/{metrics_scope}/projects/{project}` in custom, handcoded mapper functions:
   - `MonitoringMonitoredProjectSpec_FromProto` / `MonitoringMonitoredProjectSpec_ToProto`
   - `MonitoringMonitoredProjectStatus_FromProto` / `MonitoringMonitoredProjectStatus_ToProto`
2. We implemented a complete, type-safe generic round-trip KRM fuzzer at `pkg/controller/direct/monitoring/monitoredproject_fuzzer.go` using `fuzztesting.RegisterKRMFuzzer`.
3. The GCP proto status field `.create_time` maps to KRM status `createTime` and was correctly registered as `f.StatusField(".create_time")`.
4. The GCP proto identity field `.name` was correctly registered as `f.Unimplemented_Identity(".name")`, allowing both Spec and Status round-trip fuzz tests to pass flawlessly.
