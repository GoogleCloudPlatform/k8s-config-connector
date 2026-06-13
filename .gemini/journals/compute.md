# Compute Service Agentic Journal

### 2026-06-13 Implement direct KRM types for ComputeGlobalNetworkEndpoint
- **Context**: Implementing direct KRM types, `generate.sh` configuration, and references for `ComputeGlobalNetworkEndpoint`.
- **Problem**: 
  1. `ComputeGlobalNetworkEndpoint` is a sub-resource attached to a `ComputeGlobalNetworkEndpointGroup`. The parent `ComputeGlobalNetworkEndpointGroup` is an unmigrated (legacy TF-based) resource, meaning we do not have its direct KRM Go struct definitions in the `apis/` directory.
  2. `ComputeGlobalNetworkEndpoint` does not have a standard standalone REST URL on GCP, nor is it tracked under Cloud Asset Inventory (CAI). Registering its `gcpurls.Template` would trigger failure in `TestRegisteredTemplatesMatchCAI`.
- **Solution**: 
  1. Implemented a custom `Normalize` fallback for `ComputeGlobalNetworkEndpointGroupRef` that dynamically extracts `status.selfLink` from the `unstructured.Unstructured` representation, or resolves the project ID via `refs.ResolveProjectID` and builds the identity from `spec.resourceID` or metadata name.
  2. Added the `"//compute.googleapis.com/projects/{}/global/networkEndpointGroups/{}/networkEndpoints/{}"` pattern to `ignoredTemplates` in `pkg/gcpurls/registry_test.go` to cleanly bypass the CAI match check while maintaining high-fidelity identity template mapping.
- **Impact**: Enables smooth integration and robust referencing of unmigrated resources from direct KRM resources without needing a full prior migration.
