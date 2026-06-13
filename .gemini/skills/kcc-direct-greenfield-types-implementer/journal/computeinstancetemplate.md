# Journal entry for ComputeInstanceTemplate

## Observations & Learnings

- **Type Name Collisions via Proto Mappings:** 
  The GCP proto `google.cloud.compute.v1.AcceleratorConfig` was already mapped to `ReservationGuestAccelerators` in `reservation_types.go` (using the comment `// +kcc:proto=google.cloud.compute.v1.AcceleratorConfig`).
  When generating types for `ComputeInstanceTemplate`, fields in `InstanceProperties` referred to `AcceleratorConfig`.
  Because the type mapper skips generating `AcceleratorConfig` (thinking it already exists), no `AcceleratorConfig` struct is generated in `types.generated.go`.
  Declaring `type AcceleratorConfig = ReservationGuestAccelerators` as a type alias led to duplicate `DeepCopy` and `DeepCopyInto` method generation by `controller-gen`, as it treats type aliases as the same underlying type, leading to syntax errors in `zz_generated.deepcopy.go`.
  
- **Resolution:**
  Declaring `AcceleratorConfig` as a distinct struct explicitly under `instancetemplate_types.go` resolved the duplication and the missing-type compilation issues cleanly.
