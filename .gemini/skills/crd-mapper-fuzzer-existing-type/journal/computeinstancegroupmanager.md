# Journal: ComputeInstanceGroupManager transition to direct KRM types

## Observations & Key Learnings

1. **TargetSize Type Mismatch**:
   - In the baseline CRD, `spec.targetSize` is defined as a simple integer (`format: int64`, `type: integer`).
   - However, in other contexts like `spec.versions[].targetSize`, it is represented as an object with `calculated`, `fixed`, and `percent` fields (`InstancegroupmanagerTargetSize`).
   - For strict schema compatibility, we mapped `spec.targetSize` directly to `int64` in `ComputeInstanceGroupManagerSpec`, while keeping `spec.versions[].targetSize` as `*InstancegroupmanagerTargetSize`.

2. **ProjectRef kind Field Prevention**:
   - The baseline CRD schema for `spec.projectRef` does not contain a `kind` property.
   - Importing and using `ProjectRef` from `"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs"` (instead of `"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"`) ensures that `kind` is not generated under `projectRef`, keeping the generated CRD schema absolutely clean and compatible.

3. **Location Field Cardinality**:
   - The `location` field in the baseline CRD is optional. We defined `location` as `*string` and annotated it with `// +optional` to avoid it being generated as a required property.
