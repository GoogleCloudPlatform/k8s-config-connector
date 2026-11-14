# Journal: ComputeInstanceGroupManager Direct Takeover Validation

During the direct promotion validation for `ComputeInstanceGroupManager`, we discovered several key reasons for takeover drift:

## Findings and Analysis

1. **Nested Reference Resolution:**
   - The fields `instanceTemplateRef` (top-level), `versions[].instanceTemplateRef`, and `autoHealingPolicies[].healthCheckRef` are defined as legacy structures inside `apis/compute/v1beta1/computeinstancegroupmanager_types.go`.
   - Because these structures did not implement the `refs.Ref` interface, the declarative reference normalizer (`common.NormalizeReferences`) was unable to resolve them.
   - This meant they retained their local Kubernetes names (e.g. `computeinstancetemplate-${uniqueId}`) in the direct controller's converted protobuf message, instead of being populated with their resolved external selfLinks.
   - On the other hand, the real GCP API returned the fully-resolved selfLinks of these resources.
   - We resolved this by implementing the `refs.Ref` interface for `ComputeInstanceTemplateRef`, `VersionsInstanceTemplateRef`, and `ComputeHealthCheckRef` inside a new file: `apis/compute/v1beta1/computeinstancegroupmanager_reference_impls.go`.

2. **API Default Values and Mirroring Drift:**
   - When GCE receives an `instanceTemplate` at the top level (deprecated field), it automatically mirrors it into the `versions` array (e.g. `versions[0].instanceTemplate`).
   - If `versions` is not specified in the desired spec, GCE still returns the mirrored `versions` list in the actual response.
   - Also, GCE automatically populates default values for `updatePolicy.replacementMethod` (defaults to `"SUBSTITUTE"`) and `updatePolicy.type` (defaults to `"OPPORTUNISTIC"`).
   - This caused comparison mismatches (drift) between `desired` and `actual` objects during takeover.
   - We resolved this by updating the `populateDefaults` function inside `compareComputeInstanceGroupManager` (`pkg/controller/direct/compute/computeinstancegroupmanager_controller.go`) to mirror the top-level `instanceTemplate` into `versions` when `versions` is not specified, and to populate GCE's default update policy values.

## Results

With these fixes in place, the migration test suite executed both `zonalcomputeinstancegroupmanager` and `regionalcomputeinstancegroupmanager` tests on live GCP flawlessly:
- **Phase 3 (Direct takeover)** and **Phase 4 (Direct re-reconciliation)** executed as perfect read-only no-ops with zero write requests and zero drift.
- Golden HTTP cassettes were successfully re-recorded and saved under their respective fixture directories.
