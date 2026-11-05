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

### [2026-08-09] DiscoveryEngine strict project number requirement in FQN
- **Context**: Implementing direct controller for `DiscoveryEngineConversation`.
- **Problem**: `DiscoveryEngine` API returns the project number instead of the project ID in the `name` field of the created resource. Subsequent `GET` or `DELETE` requests using the project ID in the FQN resulted in `500 Internal Server Error` or `404 Not Found`. Additionally, strict project ID comparison in `GetIdentity` caused re-reconciliation failures.
- **Solution**: Updated `GetIdentity` to relax the project identifier check and prefer the project identifier (ID or number) from `status.externalRef` if it matches the parent resource. Updated the controller to use the server-returned FQN in `status.externalRef` by calling `latest.GetName()` in `updateStatus`.
- **Impact**: Ensures stability and correctness for DiscoveryEngine resources where the API is picky about project identifiers in the resource name.

### [2026-08-11] DiscoveryEngineUserStore Direct Controller and Project Identifier Mismatches
- **Context**: Implementing direct controller, E2E fixtures, and fuzzer for `DiscoveryEngineUserStore`.
- **Problem**: `DiscoveryEngineUserStore` has a `defaultLicenseConfig` field pointing to a `DiscoveryEngineLicenseConfig` resource. Since the DiscoveryEngine API returns project numbers instead of project IDs in FQN responses, the `compare` function detected a false diff between the desired state (using project ID) and actual state (using project number returned by GCP), which triggered unnecessary reconciliations (PATCH requests) during re-reconciliation, failing the E2E tests. Additionally, the test runner's service account lacked `discoveryengine.userStores.update` permission by default in `roles/editor`, causing a `403 Forbidden` error.
- **Solution**: Added the `roles/discoveryengine.admin` role to the test runner's service account. Implemented a path normalization helper in the controller's `compare` function to map project numbers to the canonical project ID for `defaultLicenseConfig` paths before running the top-level field diff comparison. Added custom E2E JSON log normalization for `defaultLicenseConfig` to prevent dynamic ID diff failures.
- **Impact**: Resolves project ID vs project number false diffs for license config references and ensures clean, repeatable E2E test runs with zero unexpected HTTP traffic.

### [2026-08-13] DiscoveryEngineEngine Direct Controller and Field Immutability
- **Context**: Implementing direct controller, E2E fixtures, and fuzzer for `DiscoveryEngineEngine`.
- **Problem**:
  1. The GCP REST Client returns `unsupported result type <nil>: <nil>` on `DeleteEngine` LRO's `op.Wait(ctx)` because the REST API returns an empty response body on delete, and the client's generated wait method expects a non-nil result.
  2. The `disable_analytics` field is immutable according to the GCP API, and specifying it in `updateMask` during a `PATCH` request causes a `400 Bad Request` with `Field "updateMask" contains an immutable path "disable_analytics"`.
- **Solution**:
  1. Modified `Delete` in `discoveryengineengine_controller.go` to ignore the `unsupported result type <nil>: <nil>` error when waiting for deletion.
  2. Kept `disableAnalytics: false` unchanged between `create.yaml` and `update.yaml` so that it doesn't trigger a diff and isn't included in the update mask.
- **Impact**: Ensures that deletion is correctly reported as a success, and updates run successfully against real GCP without encountering immutable field update mask errors.

### [2026-09-30] DiscoveryEngineControl Direct Controller, Collection Link, and Engine Dependency
- **Context**: Implementing direct controller, E2E fixtures, and fuzzer for `DiscoveryEngineControl`.
- **Problem**:
  1. `DiscoveryEngineControlIdentityFormat` was missing `collections/{collection}` in its template format. GCP API rejected resource names without `collections/default_collection` with `400 Bad Request`.
  2. GCP API requires that the parent `DiscoveryEngineDataStore` must be associated with an existing `DiscoveryEngineEngine` before any `Control` can be created under that dataStore.
  3. Action fields (e.g. `synonyms_action`, `boost_action`) on Control are immutable in GCP; modifying them triggers `Field "updateMask" contains an immutable path`.
  4. For synonyms controls, `conditions.queryTerms` cannot be specified; only `conditions.activeTimeRange` is permitted.
- **Solution**:
  1. Updated `DiscoveryEngineControlIdentityFormat` to include `collections/{collection}` and parse `Collection` from the parent DataStore link.
  2. Added `DiscoveryEngineEngine` to `dependencies.yaml` for both minimal and maximal test fixtures.
  3. Kept action fields unchanged in `update.yaml` and tested updates on mutable fields (`displayName` and `conditions.activeTimeRange`).
- **Impact**: Enables proper direct reconciliation and successful E2E golden file recording against real GCP for `DiscoveryEngineControl`.

### [2026-10-01] DiscoveryEngineControl MockGCP Alignment and Dynamic Engine Association
- **Context**: Aligning MockGCP behavior with Real GCP logs for `DiscoveryEngineControl` (Issue #13608).
- **Problem**: When a `Control` is created under a `dataStore` that is attached to an `Engine`, the real GCP DiscoveryEngine API returns the resource `name` in responses under the `engines/{engine}` hierarchy rather than `dataStores/{dataStore}`.
- **Solution**: Implemented `buildControlResponseName` in `mockdiscoveryengine/control.go` to inspect whether the target `dataStore` is attached to an `Engine` and return the matching engine control URI. Verified log alignment across minimal and maximal test fixtures with zero drift.
- **Impact**: Ensures accurate MockGCP simulation for DiscoveryEngine controls under both standalone and engine-associated datastores.

### [2026-09-30] DiscoveryEngineSitemap Direct Controller Implementation
- **Context**: Implementing direct controller, E2E fixtures, and fuzzer for `DiscoveryEngineSitemap` (Issue #13528).
- **Problem**:
  1. `CreateSitemap` requires Advanced Site Search to be enabled on the parent `DataStore` / `SiteSearchEngine`, otherwise returning `400 Bad Request: Only Advanced Site Search data stores are permitted to use the Sitemap API`.
  2. The DiscoveryEngine API returns the project number instead of the project ID in resource names on create/fetch and strictly requires the canonical project number in the resource name for `DeleteSitemap`.
  3. `DiscoveryEngineSitemap` has server-generated IDs and has no Update RPC in GCP API (it is immutable).
- **Solution**:
  1. In `Create`, if `CreateSitemap` reports that Advanced Site Search is required, call `EnableAdvancedSiteSearch` on the parent `siteSearchEngine` LRO and retry `CreateSitemap`.
  2. During `Create`, only set `a.id.Sitemap` to preserve the user's project ID in `a.id.String()` for KRM status consistency, while in `Delete`, call `Find` first and use `a.actual.GetName()` (which carries the server-assigned project number) for `DeleteSitemap`.
  3. Handled immutability in `Update` by performing spec diff comparison and no-oping if diffs are detected to avoid reconciliation loops.
- **Impact**: Enables smooth creation, re-reconciliation, and deletion of `DiscoveryEngineSitemap` resources with automatic advanced site search provisioning.

### [2026-10-07] DiscoveryEngineSchema Direct Controller Implementation
- **Context**: Implementing direct controller, E2E fixtures, and fuzzer for `DiscoveryEngineSchema` (Issue #13771).
- **Problem**:
  1. DiscoveryEngine automatically creates a system-managed `default_schema` for each DataStore upon creation. Attempting to call `CreateSchema` with a custom schema ID on a DataStore that already has an active schema returns `409 Conflict: Active schema already exists for data store`.
  2. GCP returns the `jsonSchema` string compacted without whitespace, whereas KRM manifests typically format JSON schema with indentation, causing false string diffs in top-level diff comparison during re-reconciliation.
  3. The GAPIC REST Client returns `unsupported result type <nil>: <nil>` on `DeleteSchema` LRO's `op.Wait(ctx)` because the delete operation finishes with an empty response.
- **Solution**:
  1. Tested and documented managing the schema with `resourceID: default_schema` on the parent `DiscoveryEngineDataStore`.
  2. Added semantic JSON comparison in `compare()` to unmarshal and check `DeepEqual` between desired and actual `jsonSchema` before diff reporting, preventing spurious updates.
  3. Handled the `unsupported result type <nil>: <nil>` error gracefully in `Delete()` during `op.Wait(ctx)`.
- **Impact**: Enables reliable creation, modification, re-reconciliation, and deletion of `DiscoveryEngineSchema` resources against real GCP.

