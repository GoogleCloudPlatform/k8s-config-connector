# ComputeDiskResourcePolicyAttachment Migration Journal

## Overview
Migrated `ComputeDiskResourcePolicyAttachment` from the legacy Terraform-based reconciler to the `direct` reconciler architecture.

## Key Observations and Decisions
1. **Resource Behavior:**
   - A `ComputeDiskResourcePolicyAttachment` acts as a link/attachment rather than a standard, independent GCP resource.
   - Its presence is verified by retrieving the target `Disk` object using `disksClient.Get()` and checking if the resource policy URI/selfLink is present in the `disk.ResourcePolicies` slice.
   - Reconciliation and deletion use specific Compute API POST methods (`AddResourcePolicies` and `RemoveResourcePolicies`).

2. **Region Inference:**
   - Attached policies are regional, but disks can be zonal. The resource policy URI references a region (e.g., `projects/<project>/regions/<region>/resourcePolicies/<policyName>`).
   - We infer the region from the zone (e.g., `us-central1-a` -> `us-central1`) by splitting by `-` and joining all but the last segment.

3. **Status:**
   - The resource has no mutable spec fields, meaning updates are unsupported.
   - Direct and legacy behaviors align perfectly, passing MockGCP verification tests.
