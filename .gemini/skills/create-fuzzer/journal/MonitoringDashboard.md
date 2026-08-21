When implementing the round-trip KRM fuzzer for MonitoringDashboard, we found:
1. `MonitoringDashboard` had an existing spec-only fuzzer using `NewKRMTypedSpecFuzzer` and `RegisterKRMSpecFuzzer`.
2. We upgraded it to a full round-trip fuzzer using `NewKRMTypedFuzzer` and `RegisterKRMFuzzer`.
3. We mapped both Spec and Status mappers `MonitoringDashboardStatus_FromProto` and `MonitoringDashboardStatus_ToProto`.
4. We replaced the direct `.UnimplementedFields.Insert` and `.StatusFields.Insert` calls with proper `SpecField`, `StatusField`, `Unimplemented_Identity`, and `Unimplemented_NotYetTriaged` wrapper helpers to respect clean API design guidelines.
