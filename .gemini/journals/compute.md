# Compute Service Journals

### [2026-06-20] Direct Migration of ComputeSecurityPolicy
- **Context**: Implementing the direct controller and test fixtures for `ComputeSecurityPolicy` (issue #10602).
- **Problem**: `ComputeSecurityPolicy` has distinct clients for global vs. regional resources, namely `SecurityPoliciesClient` and `RegionSecurityPoliciesClient`. Both are needed to fully support both global and regional KCC resource configurations without separate controllers.
- **Solution**: Evaluated the resource identity to dynamically detect if `Region` is specified. If `Region` is non-empty, we instantiate and invoke regional client methods (e.g., `InsertRegionSecurityPolicyRequest`, `PatchRegionSecurityPolicyRequest`, `DeleteRegionSecurityPolicyRequest`, and regional `Get`), otherwise we default to the global client methods. Copied GCP fingerprint metadata back to desired model in `Update` to avoid any modification conflicts.
- **Impact**: Highly modular, unified adapter pattern that completely avoids duplicating reconcilers for global and regional variants of Compute API resources.
