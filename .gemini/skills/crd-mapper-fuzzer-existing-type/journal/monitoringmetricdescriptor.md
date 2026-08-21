# Journal: MonitoringMetricDescriptor

## Learnings & Observations

1. **Schema Compatibility with ProjectRef**:
   - `refs.ProjectRef` under `github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1` has a `kind` field.
   - Using this package causes `dev/tasks/diff-crds` to flag a schema deviation (`+ spec.projectRef.kind=string`) when the baseline CRD does not have it.
   - **Resolution**: Use `github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs` which contains a `ProjectRef` representation *without* the `kind` field. This achieves a 100% perfect schema-match.

2. **Custom / Handcoded Mappers**:
   - KRM struct names that match DCL types (e.g., `MetricdescriptorLabels` and `MetricdescriptorMetadata`) may differ in casing (e.g., lower `d` in `descriptor`) compared to what the proto-based types generator produces as unreachable types (e.g., `MetricDescriptor_MetricDescriptorMetadata` with capital `D`).
   - Because of this, the automatic mapper generation can call functions like `MetricdescriptorMetadata_FromProto` that are undefined in the generated output.
   - Required KRM string fields (e.g. `MetricKind string`, `ValueType string`) conflict with the generator's default assumption that they are pointers (`*string`).
   - **Resolution**: Handcoded mapping functions in `pkg/controller/direct/monitoring/metricdescriptor_mappings.go` resolve all these type mismatches and naming conventions. The generator automatically detects and skips generating conflicting duplicates.
