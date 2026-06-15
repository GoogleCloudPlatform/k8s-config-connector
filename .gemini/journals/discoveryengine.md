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

### [2026-06-05] DiscoveryEngineConversation Type Generation and Mapper Limitation
- **Context**: Implementing initial KRM types and IdentityV2 for `DiscoveryEngineConversation`.
- **Problem**: When generating types, nested protobuf types under skipped messages (e.g. `SearchResponse.Summary` inside `SearchResponse` which has the suffix `Response` and is thus skipped in `generatemappercommand.go`) did not get their mapper functions generated. Additionally, when types are first generated, they are commented out in `types.generated.go` as unreachable until the manual spec/status structs reference them.
- **Solution**: We referenced the nested types like `ConversationMessage` and `Reply` directly inside `DiscoveryEngineConversationSpec` in `conversation_types.go` and reran type generation, which successfully made the generator detect their reachability and automatically uncomment them in `types.generated.go`. To avoid mapper-compilation failures on local tests without committing mapper files (as instructed), we discarded changes to `mapper.generated.go` since no controller is registered to use the mappers yet.
- **Impact**: When adding new resources, make sure to completely specify their spec and status structures with the target protobuf types in the handwritten types file *before* final type generation, so the `prune-types` tool correctly marks them as reachable.

### [2026-06-05] Handling Unreachable Nested Types in Greenfield Generation
- **Context**: Implementing direct types for `DiscoveryEngineControl` under Issue #9237.
- **Problem**: When a new resource is first added to `generate.sh`, its nested types (e.g., `Condition`, `Control_BoostAction`) are initially generated as "unreachable types" and commented out in `types.generated.go`.
- **Solution**: First define the Spec fields referencing these types in `<resource>_types.go` and run `./generate.sh` again. The generator automatically identifies them as reachable, uncomments them, and makes them available.
- **Impact**: This avoids the need to manually copy or define nested proto structs, keeping types.generated.go fully managed by the builder.
