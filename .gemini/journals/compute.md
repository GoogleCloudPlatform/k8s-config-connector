# Compute Service Journals

### [2026-06-20] Direct Migration of ComputeSecurityPolicy
- **Context**: Implementing the direct controller and test fixtures for `ComputeSecurityPolicy` (issue #10602).
- **Problem**: `ComputeSecurityPolicy` has distinct clients for global vs. regional resources, namely `SecurityPoliciesClient` and `RegionSecurityPoliciesClient`. Both are needed to fully support both global and regional KCC resource configurations without separate controllers.
- **Solution**: Evaluated the resource identity to dynamically detect if `Region` is specified. If `Region` is non-empty, we instantiate and invoke regional client methods (e.g., `InsertRegionSecurityPolicyRequest`, `PatchRegionSecurityPolicyRequest`, `DeleteRegionSecurityPolicyRequest`, and regional `Get`), otherwise we default to the global client methods. Copied GCP fingerprint metadata back to desired model in `Update` to avoid any modification conflicts.
- **Impact**: Highly modular, unified adapter pattern that completely avoids duplicating reconcilers for global and regional variants of Compute API resources.

### [2026-06-23] Direct Migration of ComputeAutoscaler
- **Context**: Implementing the direct controller and test fixtures for `ComputeAutoscaler` (issue #10727).
- **Problem**: `ComputeAutoscaler` has a target reference (`TargetRef`) referring to `ComputeInstanceGroupManager`, which is a legacy DCL-reconciled resource. Since MockGCP does not support mocking the DCL-managed `InstanceGroupManagers` API, any `ComputeAutoscaler` E2E tests containing this dependency will be skipped on MockGCP.
- **Solution**: We implemented the direct controller for `ComputeAutoscaler` and registered `k8s.ReconcilerTypeDirect` in `static_config.go` under `SupportedControllers`. To support resolving the referenced resource, we wrote a helper `resolveComputeAutoscalerRefs` to dynamically fetch the status fields (`status.externalRef` or `status.selfLink`) of the referenced `ComputeInstanceGroupManager`. We verified mappers and fuzzer correctness by running targeted fuzz roundtrip tests with `FOCUS=ComputeAutoscaler`.
- **Impact**: Dynamic testing ensures the direct controller is automatically vetted, and any mappers or fuzzers are fully verified, even if real API dependencies prevent executing the full E2E test against MockGCP.
