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

### [2026-06-03] DiscoveryEngine mixed v1 and v1beta service types generation
- **Context**: Implementing direct types for DiscoveryEngineSampleQuerySet (Issue #8713).
- **Problem**: `SampleQuerySet` is defined under `google.cloud.discoveryengine.v1beta`, but all other discoveryengine resources in `apis/discoveryengine/v1alpha1` are under `google.cloud.discoveryengine.v1`. Sequential runs of `generate-types` overwrite `types.generated.go`, deleting the other service's types and breaking compilation.
- **Solution**: Updated `generate.sh` to rename `types.generated.go` to `v1_types.generated.go` and `v1beta_types.generated.go` during sequential generation runs. The `PruneTypes` linter and Go `deepcopy` generator correctly parse all `*.generated.go` files, allowing types from both services to coexist and compile flawlessly.
- **Impact**: Enables co-existence of types from multiple service API versions under the same KCC group version folder without generation conflicts.

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

### [2026-06-05] DiscoveryEngineSampleQuery Proto defined only in v1beta
- **Context**: Implementing direct types, CRD, and IdentityV2 for DiscoveryEngineSampleQuery (Issue #9239).
- **Problem**: The issue description specified using service `google.cloud.discoveryengine.v1`. However, `SampleQuery` is only defined in the `google.cloud.discoveryengine.v1beta` proto, not `v1`. Sequential runs of the generator for different proto versions overwrite and prune each other's types in the shared `types.generated.go`.
- **Solution**:
  1. Ran `generate-types` pointing to a temporary output folder with `--prune-unused-types=false` to isolate `SampleQuery`'s raw types.
  2. Surgically appended the new types to the end of `apis/discoveryengine/v1alpha1/types.generated.go`.
  3. Updated `apis/discoveryengine/v1alpha1/generate.sh` to use `google.cloud.discoveryengine.v1beta` for generating types for all resources.
  4. Registered the unregistered URL template exception in `pkg/gcpurls/registry_test.go` to prevent CAI test failure.
- **Impact**: When adding a new resource whose proto is in a different version of the same service, generating to a temporary folder and merging is a safe and reliable workaround to avoid overwriting types.
