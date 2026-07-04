When implementing the round-trip KRM fuzzer for MonitoringMetricDescriptor, we observed:
1. `MonitoringMetricDescriptor` has both Spec and Status round-trip mapping functions:
   - `MonitoringMetricDescriptorSpec_FromProto` / `MonitoringMetricDescriptorSpec_ToProto`
   - `MonitoringMetricDescriptorStatus_FromProto` / `MonitoringMetricDescriptorStatus_ToProto`
2. We implemented a full round-trip fuzzer at `pkg/controller/direct/monitoring/metricdescriptor_fuzzer.go` using `fuzztesting.RegisterKRMFuzzer` to cover both Spec and Status fields.
3. The GCP proto field `metadata.time_series_resource_hierarchy_level` was not mapped to KRM, so we marked it as `Unimplemented_NotYetTriaged` in the fuzzer to prevent round-trip data loss errors.
4. All other mapped fields for Spec, Status, and Identity were correctly registered, allowing both Spec and Status fuzz tests to pass flawlessly.
