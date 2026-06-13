# ComputeDiskResourcePolicyAttachment Journal

## Overview
We transitioned the KRM types for `ComputeDiskResourcePolicyAttachment` under `apis/compute/v1alpha1/` to a direct Go struct representation.

## Details
1. **No Upstream Proto Resource**:
   - `ComputeDiskResourcePolicyAttachment` is not a standalone direct resource in the upstream GCP Compute REST/proto API; instead, attaching/detaching a policy maps to a sub-operation on `ComputeDisk` (`disks.addResourcePolicies` and `disks.removeResourcePolicies`).
   - For this reason, there is no corresponding 1:1 proto message we could map to.
   - Therefore, we hand-wrote the types file completely rather than generating it via `generate-types`.

2. **Strict Schema Compatibility**:
   - The root level `spec` was required in the baseline CRD. We preserved this by annotating the `Spec` field of `ComputeDiskResourcePolicyAttachment` struct with the `// +required` comment.
   - Run `dev/tasks/diff-crds` frequently to confirm zero schema differences.

3. **Custom Reference Types**:
   - Since the original `diskRef` does not have a `kind` field, we hand-coded the custom `ComputeDiskRef` locally in the types file containing only `External`, `Name`, and `Namespace` fields.
   - For `projectRef` which also has no `kind` field, we imported and used the canonical `ProjectRef` from `github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs` to adhere to the standard.

4. **Validation and Verification**:
   - Executed `make ready-pr` which regenerated the strong-typed clients under `pkg/clients/generated/apis/` using the updated CRD.
   - All schema validation diffs are clean.
