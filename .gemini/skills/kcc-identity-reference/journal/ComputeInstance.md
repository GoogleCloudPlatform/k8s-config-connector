# ComputeInstance Migration Journal

## Overview
Migrated `ComputeInstance` (`InstanceRef` / `InstanceIdentity`) to the `identity.IdentityV2` and `refs.Ref` pattern using `gcpurls.Template`.

## Implementation Details
1. **Naming Conventions**:
   - The GCP resource `Instance` has a Kubernetes custom resource kind named `ComputeInstance`.
   - The Go structs were kept as `InstanceIdentity` and `InstanceRef` (matching previous filenames/structs) to avoid breaking API or internal schema references across other service packages (e.g., `backupdr`, `clouddms`, `datamigration`, `networkmanagement`).

2. **Template Mapping**:
   - Format: `projects/{project}/zones/{zone}/instances/{instance}`
   - The existing `InstanceRef` was updated to implement `refs.Ref`.
   - For backwards compatibility, the `NormalizedExternal` method was kept as a wrapper around the new `Normalize` method.
