### [2026-05-18] mockgcp discoveryengine migration to httptogrpc
- **Context**: Moving mockgcp discoveryengine away from grpc-gateway.
- **Problem**: The service implementation was already using the official `pb` library, but was still using `grpc-gateway` for HTTP multiplexing via `pbhttp`.
- **Solution**: Removed `discoveryengine` protos from `mockgcp/Makefile`. Deleted generated files in `mockgcp/generated/google/cloud/discoveryengine`. Updated `mockgcp/mockdiscoveryengine/service.go` to use `httptogrpc.NewGRPCMux` and `mux.AddOperationsPath`. Removed `mux.RewriteError` as it is not supported by `httptogrpc`.
- **Impact**: Removes legacy grpc-gateway dependencies for discoveryengine in mockgcp.### [2026-05-27] IdentityMappingStore Missing from CAI
- **Context**: Implementing IdentityV2 for DiscoveryEngineIdentityMappingStore.
- **Problem**: The resource IdentityMappingStore URL template was missing from CAI and the GCP URLs registry, causing the TestRegisteredTemplatesMatchCAI test to fail.
- **Solution**: Added the exception to pkg/gcpurls/registry_test.go.
- **Impact**: Agents must remember to add unregistered URL templates to registry_test.go when the resource is not supported by CAIS to prevent test failures.

