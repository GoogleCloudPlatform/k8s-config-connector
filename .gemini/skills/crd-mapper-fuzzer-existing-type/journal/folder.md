# Journal: Folder transition to direct KRM types

## Observations & Key Learnings

1. **Reusing Custom References in the same Package**:
   - `OrganizationRef` was already hand-coded in `project_types.go` because the baseline `Project` CRD has `organizationRef` with `external`, `name`, and `namespace`.
   - Since `Folder` resides in the same `v1beta1` resourcemanager package, we were able to directly reference `OrganizationRef` in `FolderSpec` without having to redeclare it or duplicate its helper methods (`GetGVK`, `GetNamespacedName`, etc.).

2. **Handling Field-Level Signature & `oneOf` Spec Validation**:
   - The Folder resource restricts spec fields so that only `folderRef` or `organizationRef` (or neither) may be specified. This is represented by a spec-level `oneOf` block in the baseline CRD.
   - We updated `scripts/add-validation-to-crds/parse-crds.go` to match the exact field signature `displayName,folderRef,organizationRef,resourceID` for kind `Folder`. This replicates the `oneOf` constraint generator logic originally built for `Project`.

3. **Hand-coded Mappers over Auto-generation**:
   - Default auto-generated mappings for fields like `displayName` can result in undefined signature helper functions like `FolderSpec_DisplayName_ToProto` if the generator gets confused by custom types.
   - Hand-coding the full `FolderSpec_FromProto`, `FolderSpec_ToProto`, `FolderStatus_FromProto`, and `FolderStatus_ToProto` mapping functions in `pkg/controller/direct/resourcemanager/mapper.go` is cleaner, more robust, and compiles perfectly with zero warnings.
