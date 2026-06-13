# Journal: MonitoringServiceLevelObjective

## Observations & Learnings

1. **DCL-specific custom fields**:
   - `MonitoringServiceLevelObjective` was previously managed by a DCL-based controller.
   - The DCL schema included empty fields (like `availability` and `operationAvailability` on `BasicSli`) and fields not present in standard monitoring v3 Protobuf (like `experience` on `BasicSli_LatencyCriteria`/`operationLatency`).
   - By handcoding these structs locally in `servicelevelobjective_types.go`, we forced the code generator to skip generating standard Protobuf definitions and use our custom schema definitions instead.

2. **Kubebuilder preservation of unknown fields**:
   - Empty structs generate as `type: object` with no properties by default in `controller-gen`.
   - The baseline CRD had them as `json` (i.e. `x-kubernetes-preserve-unknown-fields: true`).
   - We resolved this difference beautifully by adding `// +kubebuilder:validation:XPreserveUnknownFields` onto the empty structs `BasicSli_AvailabilityCriteria` and `BasicSli_OperationAvailability`.

3. **Required double fields/Format preservation**:
   - KRM fields of type `float64` that are not pointers lose the `format: double` OpenAPI property in the generated schema by default.
   - We preserved `format: double` strictly by adding `// +kubebuilder:validation:Format=double` annotations on `Goal`, `Range.Min`/`Range.Max`, and `WindowsBasedSli_PerformanceThreshold.Threshold`.

4. **Fuzzer configuration for Spec-only**:
   - Since `ServiceLevelObjective` Status fields (`createTime`, `deleteTime`, and `serviceManagementOwned`) are not mapped to/from the Protobuf in direct controller (or we don't have status mappers), we registered a Spec-only fuzzer using `NewKRMTypedSpecFuzzer` and `RegisterKRMSpecFuzzer`.
   - We resolved unspecified Protobuf enum values (`CALENDAR_PERIOD_UNSPECIFIED`) during round-trip fuzzing by clearing the `Period` oneof in `FilterSpec` if it contains the unspecified calendar period value.
