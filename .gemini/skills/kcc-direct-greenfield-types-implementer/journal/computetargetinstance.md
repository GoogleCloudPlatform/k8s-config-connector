# Journal: ComputeTargetInstance Implementation

## Observations & Findings

1. **Resource Mapping**: `ComputeTargetInstance` maps directly to GCP's `TargetInstance` in the `google.cloud.compute.v1` service API.
2. **Schema Alignment for Zonal Resource**:
   - The auto-generated structure by default introduced a `Location string` field.
   - However, the pre-existing CRD for `ComputeTargetInstance` (originally configured via the Terraform-based controller) uses `zone` as a mandatory spec property, while `location` is not part of the schema.
   - To preserve strict schema backward compatibility, the `Location` field was replaced with a required `Zone string` field, and references such as `InstanceRef`, `ComputeNetworkRef`, and `ComputeSecurityPolicyRef` were added.
3. **Labels**:
   - Standard labels (`cnrm.cloud.google.com/managed-by-kcc`, `cnrm.cloud.google.com/system`, `cnrm.cloud.google.com/tf2crd`) are added, and the stability level is explicitly marked as `stable` to align with the existing `v1beta1` stable CRD.
4. **Code Generation**:
   - Running `./generate.sh` successfully commented out the unreachable type comment block in `types.generated.go`, generated deepcopy methods, and updated CRDs.
   - Running `dev/tasks/generate-resource-report` updated the CRD reports in `docs/reports/` to indicate that `ComputeTargetInstance` now has direct KRM types.
