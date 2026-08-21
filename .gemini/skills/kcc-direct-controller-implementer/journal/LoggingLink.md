# LoggingLink Direct Controller

Journal entries regarding implementation details and architectural updates for LoggingLink.

## [2026-06-14] Refactored LoggingLink Controller to Align with KCC Best Practices

Refactored the LoggingLink direct controller `pkg/controller/direct/logging/link_controller.go` to strictly adhere to KCC expert design patterns:

1. **Proto Format Desired State**: Instead of holding references to raw KRM objects for desired state, we convert the KRM Spec to its Proto representation once in `AdapterForObject` and store it as `desired *loggingpb.Link` in the adapter. This keeps interfaces clean, avoids redundant conversions, and aligns with `actual`.
2. **Handling Non-API / KRM-only Spec Fields**: Client-side options (such as `resourceID` on LoggingLink spec) are parsed in `AdapterForObject` and stored as separate explicitly-named fields (`desiredID string`) on the adapter struct.
3. **Reference Normalization**: Added the standard `common.NormalizeReferences` invocation in `AdapterForObject`.
4. **Clean Client Construction**: Retrieved REST options using `m.config.RESTClientOptions()` and constructed the client with `gcp.NewConfigRESTClient(...)` instead of manually setting authenticated HTTP client.
5. **Immutability Error Reporting**: Since LoggingLink is immutable in GCP and does not support updates, updated the `Update` method to perform `compareLink` check on spec fields and return a descriptive error `fmt.Errorf("LoggingLink is immutable and cannot be updated")` if any diffs are found. This surfaces the immutable constraint on resource status rather than silently ignoring diffs.
6. **E2E Validation**: Validated changes across all 17 logging-related fixtures, ensuring 100% test success under MockGCP.
