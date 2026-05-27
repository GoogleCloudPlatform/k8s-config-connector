### [2026-05-18] mockgcp discoveryengine migration to httptogrpc
- **Context**: Moving mockgcp discoveryengine away from grpc-gateway.
- **Problem**: The service implementation was already using the official `pb` library, but was still using `grpc-gateway` for HTTP multiplexing via `pbhttp`.
- **Solution**: Removed `discoveryengine` protos from `mockgcp/Makefile`. Deleted generated files in `mockgcp/generated/google/cloud/discoveryengine`. Updated `mockgcp/mockdiscoveryengine/service.go` to use `httptogrpc.NewGRPCMux` and `mux.AddOperationsPath`. Removed `mux.RewriteError` as it is not supported by `httptogrpc`.
- **Impact**: Removes legacy grpc-gateway dependencies for discoveryengine in mockgcp.

### [2026-05-27] IdentityMappingStore Missing from CAI
- **Context**: Implementing IdentityV2 for DiscoveryEngineIdentityMappingStore.
- **Problem**: The resource IdentityMappingStore URL template was missing from CAI and the GCP URLs registry, causing the TestRegisteredTemplatesMatchCAI test to fail.
- **Solution**: Added the exception to pkg/gcpurls/registry_test.go.
- **Impact**: Agents must remember to add unregistered URL templates to registry_test.go when the resource is not supported by CAIS to prevent test failures.

### [2026-05-27] Scaffolded DiscoveryEngineSampleQuerySet with mockgcp proto overlay
- **Context**: Implementing `DiscoveryEngineSampleQuerySet` for `v1alpha1` (mapped to GCP `v1`).
- **Problem**: `SampleQuerySet` proto did not exist in GCP `v1` (only in `v1beta` and `v1alpha`). The KRM version `v1alpha1` was already using `v1` for other resources (e.g. `DataStore`). Running `generate-types` multiple times overwrites `types.generated.go`, and it doesn't support multiple service versions in a single run.
- **Solution**: Copied `sample_query_set.proto` from `v1beta` to `mockgcp/apis/google/cloud/discoveryengine/v1/sample_query_set.proto`, changed its package to `google.cloud.discoveryengine.v1`, and added the mockgcp directory to `generate-proto.sh`. This effectively overlays the `v1beta` proto into `v1` during code generation, allowing `generate-types` to generate everything together without losing types.
- **Impact**: When adding a new resource that only exists in a beta API to an existing KRM version that maps to the v1 API, use the `mockgcp/apis/...` overlay technique to trick the generator into finding the proto, instead of splitting KRM versions or manually copying generated structs.