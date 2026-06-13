# Compute Service Journal

### [2026-06-13] ComputeTargetPool Direct KRM Implementation
- **Context**: Implementing direct KRM types and reference/identity for `ComputeTargetPool` (Issue #9894).
- **Problem**: `ComputeTargetPool` needs to reference legacy KRM resources (`ComputeHTTPHealthCheck` of kind `ComputeHTTPHealthCheck` which is still TF/DCL-based) and other direct resources like `ComputeInstance` and `ComputeTargetPool` itself.
- **Solution**: We defined a custom reference struct `ComputeHTTPHealthCheckRef` and registered it with `refs.Register`. We implemented its `Normalize` method with a fallback function that reads `status.selfLink` from the legacy unstructured object. This ensures perfect backwards compatibility without migrating `ComputeHTTPHealthCheck` itself. We also used `computev1beta1.InstanceRef` for instances.
- **Impact**: Demonstrates how direct resources in `v1alpha1` can seamlessly and cleanly reference legacy TF/DCL-based resources in their specs while keeping Go types clean and backward-compatible.
