# Network Management Journal

### [2026-07-02] Scaffolding multiple resources within the same service
- **Context**: [NetworkManagementVpcFlowLogsConfig](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/10291)
- **Problem**: Running `go run . generate-types` for `NetworkManagementVpcFlowLogsConfig` overwrote and pruned the unreachable types of `NetworkManagementConnectivityTest` in `types.generated.go`. This is because the type generator prunes unused types by default and only knows about the single resource specified via the `--resource` flag if run sequentially.
- **Solution**: Pass multiple `--resource` flags to a single invocation of `generate-types` in `apis/networkmanagement/v1alpha1/generate.sh` to allow the generator to reach and preserve types for both resources:
  ```bash
  go run . generate-types \
    --service google.cloud.networkmanagement.v1 \
    --api-version networkmanagement.cnrm.cloud.google.com/v1alpha1  \
    --resource NetworkManagementConnectivityTest:ConnectivityTest \
    --resource NetworkManagementVpcFlowLogsConfig:VpcFlowLogsConfig
  ```
- **Impact**: Ensures type-generation safety and prevents build breaks/regressions when scaffolding a new resource in an existing API group directory.
