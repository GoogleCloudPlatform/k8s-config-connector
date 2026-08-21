# ComputeNetwork Direct Controller Implementation Journal

## Observations
1. **Client-Side Behavior (deleteDefaultRoutesOnCreate)**: This is a client-side (KRM-only) behavioral field that is not represented in the GCP network proto schema. If enabled, we query the Routes API using the newly created network's `selfLink` and delete any matching routes (`destRange="0.0.0.0/0"`). This matches the exact behavior of the Terraform-based legacy controller.
2. **Global Scope Identity**: Unlike many compute resources which are regional or zonal, Compute Network is a global resource. It relies on the project-id and network name for its identity representation.
3. **Merge Patch Update (No UpdateMask)**: The Compute Network PATCH endpoint behaves as a merge patch on the resource itself, meaning that `computepb.PatchNetworkRequest` does not have or require an explicit `UpdateMask` field.
