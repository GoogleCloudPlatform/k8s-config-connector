### [2026-06-15] NetworkManagementVpcFlowLogsConfig Types & Identity Scaffold
- **Context**: Implementing initial types, CRD, and IdentityV2 for `NetworkManagementVpcFlowLogsConfig` under `v1alpha1`.
- **Problem**: 
  1. The code generator writes all package-level reachable types into a single `types.generated.go`. Running `generate-types` separately for multiple resources inside the same package causes the second run to overwrite/prune unreachable types for the first resource, breaking existing types (e.g. `ConnectivityTest`).
  2. The `gcpurls.Template` matching mechanism is strict with CamelCase fields. Standard underscores in URL parameter names (like `{vpc_flow_logs_config}`) will cause panics if the struct uses CamelCase field naming (like `VpcFlowLogsConfig`).
- **Solution**: 
  1. Combined the resource generator calls in `generate.sh` into a single invocation using multiple `--resource` flags, so that the generator correctly processes and retains reachable types for all resources in `types.generated.go`.
  2. Configured the `gcpurls.Template` parameter variable without underscores (e.g. `{vpcflowlogsconfig}`) so that lowercasing matches the CamelCased struct field `VpcFlowLogsConfig`.
  3. Added first-class KCC resource references (`ComputeVPNTunnelRef` and `ComputeInterconnectAttachmentRef`) in `computerefs.go` and mapped them in the Spec to enforce proper KCC reference patterns.
- **Impact**: Ensures package-level type generation does not prune existing package structures, prevents runtime panics at initialization, and maintains strict reference safety in CRDs.
