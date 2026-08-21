# Journal: Transitioning MonitoringAlertPolicy to Direct KRM Types

- **Date**: Wednesday, June 10, 2026
- **Resource**: `MonitoringAlertPolicy` (`monitoring.cnrm.cloud.google.com/v1beta1`)

## Key Observations and Learnings

1. **Custom Reference Signature**:
   `MonitoringNotificationChannelRef` was already defined in the package, referencing notification channels. By utilizing this local type instead of `k8sv1alpha1.ResourceRef`, we completely matched the schema's properties (avoiding the addition of a `kind` field which was not in the original CRD). Moreover, since it contains exactly `external`, `name`, and `namespace` fields, the schema builder automatically recognized its signature and generated the proper OpenAPI `oneOf` reference constraints!

2. **Validation on Struct Fields**:
   The baseline CRD required the `spec` field of the resource itself. Adding `// +required` and removing `,omitempty` on `Spec` in `MonitoringAlertPolicy` successfully generated `required: [spec]` in the schema.

3. **Handcoded Mappers for Empty Durations**:
   The proto of `AlertPolicy` uses `google.protobuf.Duration` for fields like `duration` and `forecast_horizon`. In the KRM, these are mapped to string fields (like `Duration string`). Passing empty strings `""` to `direct.StringDuration_ToProto` results in compilation/fuzz errors. We solved this robustly by checking if the string is empty before invoking the duration mapping helpers in our custom `AlertpolicyConditionAbsent_ToProto`, `AlertpolicyConditionMonitoringQueryLanguage_ToProto`, `AlertpolicyConditionThreshold_ToProto`, and `AlertpolicyForecastOptions_ToProto` functions.

4. **100,000 Fuzz Roundtrips**:
   Configured the KRM fuzzer with type-safe helpers and successfully triaged/unimplemented fields that only exist in the proto (e.g. `validity`, `user_labels`, `condition_sql` etc.) to guarantee 100% roundtrip fidelity.
