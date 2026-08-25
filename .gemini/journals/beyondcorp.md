### [2026-08-25] BeyondCorpClientConnectorService Direct Controller & Fixtures
- **Context**: Implementing direct controller, E2E fixtures, and fuzzer registration for BeyondCorpClientConnectorService.
- **Problem**: The resource types and controller had been scaffolded but the controller was not registered in the centralized registry, and E2E test fixtures did not exist.
- **Solution**: Imported the package in `pkg/controller/direct/register/register.go` so its `init()` method executes. Created both "minimal" and "maximal" test fixtures (along with an `update.yaml` to cover mutable fields like `displayName` and `destinationRoutes`). Executed the API check tests to verify that every spec field is covered by the test fixtures, updating `alpha-missingfields.txt` to remove `BeyondCorpClientConnectorService` from the missing fields exception list.
- **Impact**: Enables end-to-end reconciliation tests and guarantees complete API schema coverage.

### [2026-05-19] ClientGateway proto has no specifiable fields
- **Context**: Implementing the direct controller for BeyondCorpClientGateway
- **Problem**: The proto definition for ClientGateway only contains output fields (except for name). As a result, the KRM Spec struct generated initially lacked any domain-specific fields, and types.generated.go commented out ClientGatewayObservedState.
- **Solution**: The BeyondCorpClientGatewaySpec was only configured with standard identity fields (ProjectRef, Location, ResourceID). All domain-specific fields (createTime, updateTime, state, id, clientConnectorService) from the proto were manually added to the BeyondCorpClientGatewayObservedState in clientgateway_types.go, so they become available on status.observedState.
- **Impact**: When scaffolding a resource that primarily has output-only fields in GCP, be prepared to manually move those fields into the ObservedState and expect that the resource update method cannot change any domain fields (only identity/status updates happen).
