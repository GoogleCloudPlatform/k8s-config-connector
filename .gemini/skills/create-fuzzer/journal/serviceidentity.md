# Journal: ServiceIdentity Fuzzer Implementation

## Observations & Learnings

### 1. Fuzzing a Resource with Parent-Only KRM Spec and Status-Only Proto Fields
`ServiceIdentity` is a unique resource where the KRM Spec contains only structural parent/identifier references (`resourceID` and `projectRef`), while the GCP protobuf representation (`pb.ServiceIdentity`) contains solely status-related metadata (`email` and `unique_id`).
In this scenario:
- There are no Spec fields to fuzz or map back and forth since KCC handles URL path assembly at reconciliation-time, not through protobuf Spec-to-Proto mapping.
- All protobuf fields (`.email`, `.unique_id`) are mapped directly to KRM Status / ObservedState and registered as `f.StatusField` inside the fuzzer.
- The fuzzer compiles and runs successfully, verifying that status fields are properly round-tripped and that spec-related structures are consistently ignored during Status comparisons.
