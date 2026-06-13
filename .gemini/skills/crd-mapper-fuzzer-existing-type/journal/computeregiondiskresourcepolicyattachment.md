# ComputeRegionDiskResourcePolicyAttachment Journal

## Overview
We implemented the direct KRM types and configured code generation for `ComputeRegionDiskResourcePolicyAttachment`.

## Details
1. **Existing Types & Migration**:
   - `ComputeRegionDiskResourcePolicyAttachment` is an attachment resource representing the attachment relationship between a `ComputeRegionDisk` and a `ComputeResourcePolicy`.
   - In GCP Compute Engine, this attachment is not a standalone API resource, but is instead managed via sub-resource endpoints on the disks (e.g. `addResourcePolicies` and `removeResourcePolicies` sub-resource actions), and GET actions return the attached resource policies as a nested list on the disk resource itself.
   - For this reason, there is no standalone proto message representing this attachment resource.

2. **Types Hand-coding**:
   - We created `apis/compute/v1alpha1/computeregiondiskresourcepolicyattachment_types.go` and hand-coded the types to exactly match the baseline CRD schema.
   - We used the standard `refs.ProjectRef` from `github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs` to represent the project reference since it does not have a `kind` field.
   - We defined a custom nested structure `ComputeRegionDiskResourcePolicyAttachmentDiskRef` representing the disk reference since the baseline CRD lacked a `kind` field on `diskRef`.
   - To match the baseline CRD schema precisely, we annotated the `Spec` field on the root object with `// +required`.

3. **Verification**:
   - Ran `dev/tasks/diff-crds` which showed 100% strict schema compatibility (absolutely empty output!).
   - Ran fast targets `make manifests generate-go-client ensure fmt`, which successfully compiled the new client types, formatted the codebase, and regenerated the resource report.
