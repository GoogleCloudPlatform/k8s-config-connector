# Journal: MonitoringGroup Direct KRM Types Implementation

- **Resource**: `MonitoringGroup:Group` under service `google.monitoring.v3`.
- **API Version**: `v1beta1`.
- **Observations & Learnings**:
  - `MonitoringGroup` was migrated as a brownfield resource from a DCL-based schema.
  - Since it is already stable in KCC (`v1beta1`), the stability level in `group_types.go` was configured as `stable` instead of `alpha` to maintain backward compatibility.
  - The default scaffolding output by `generate-types` contained a `Location` field. Since `MonitoringGroup` in GCP Monitoring API is a global/project-level resource with no location parameter (as verified in the `.build/third_party/googleapis/google/monitoring/v3/group.proto` file), the `Location` field was manually removed.
  - Required fields `displayName` and `filter` were kept as pointers (`*string`) to seamlessly generate the mapper and allow for direct assignable matching without needing custom helper mapper functions. Adding `// +required` ensures that they are still marked as required at the OpenAPI/CRD validation level, keeping full compatibility with the existing CRD definition.
  - Added `ParentRef *MonitoringGroupRef json:"parentRef,omitempty"` to reference the parent group (`parent_name`), which was automatically validated by the standard `add-validation-to-crds` script.
