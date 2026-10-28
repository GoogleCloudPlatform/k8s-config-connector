# ComputeNetworkEndpoint Direct Controller Journal

## Implementation Details

- **Direct Controller Implementation**: Implemented `ComputeNetworkEndpointAdapter` and registered model with `registry.RegisterModel(krm.ComputeNetworkEndpointGVK, NewComputeNetworkEndpointModel)`.
- **API Client**: Added `newNetworkEndpointGroupsClient` to `pkg/controller/direct/compute/client.go` using the official GAPIC REST client (`compute.NewNetworkEndpointGroupsRESTClient`).
- **Endpoint Operations**:
  - `Find`: Uses `ListNetworkEndpoints` on the parent `NetworkEndpointGroup` and filters by `ipAddress`, `port`, and `instance` name.
  - `Create`: Uses `AttachNetworkEndpoints` and waits on the returned long-running operation.
  - `Update`: Spec is immutable; checks for diffs using `compareComputeNetworkEndpoint` and returns an error if any diff is found.
  - `Delete`: Uses `DetachNetworkEndpoints` and waits on the returned long-running operation. Gracefully handles not-found errors.
- **Testing**:
  - Validated with `hack/record-gcp` on live GCP for both standard e2e fixtures and `TestMigrationToDirect`.
  - Golden files and mock logs updated and verified with `hack/compare-mock` and `TestGoldenLogAlignment`.
