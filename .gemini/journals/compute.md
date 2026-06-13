# Compute Service Journal

### [2026-06-13] ComputeTargetHTTPSProxy Type Migration to Direct KRM Standards
- **Context**: Migrating ComputeTargetHTTPSProxy to use direct KRM types.
- **Problem**: The existing hand-written types in `apis/compute/v1beta1/targethttpsproxy_types.go` used `*int` for `ObservedGeneration` and `ProxyId`. Direct KRM standards mandate that `ObservedGeneration` must be exactly `*int64`.
- **Solution**: Updated `ObservedGeneration` and `ProxyId` to `*int64` in `targethttpsproxy_types.go` and updated `targethttpsproxy_mapper.go` accordingly. Ran `dev/tasks/generate-types-and-mappers` to regenerate CRDs and `zz_generated.deepcopy.go`.
- **Impact**: Ensures that `ComputeTargetHTTPSProxy` adheres to direct controller requirements and resolves compilation mismatches.
