### [2026-05-17] bigquerydatatransfer migration
- **Context**: Moving mockgcp bigquerydatatransfer to httptogrpc (PR 8302).
- **Problem**: Previously used RewriteError for 404s to clear error.Errors.
- **Solution**: Removed as httptogrpc does not support it.
- **Impact**: It is important to remember that RewriteError is not supported in httptogrpc.

### [2026-06-02] BigQueryDataTransferTransferConfig direct types implementation
- **Context**: Implementing initial direct KRM types, CRD, and IdentityV2 for BigQueryDataTransferTransferConfig (Issue #9021).
- **Problem**: The gcpurls.Template parser matches placeholders case-insensitively but is sensitive to underscores. A template with `{transfer_config}` segment expects a struct field named `Transfer_config` (which lowers to `transfer_config`). Using `TransferConfig` (which lowers to `transferconfig`) causes a "field transfer_config not found" panic at runtime init.
- **Solution**: Named the struct field `Transfer_config` (PascalCase with underscore) so that `strings.ToLower(f.Name)` results in `transfer_config`, aligning perfectly with the `{transfer_config}` placeholder.
- **Impact**: When implementing IdentityV2 with gcpurls.Template, any snake_case placeholder `{foo_bar}` must map to a struct field named `Foo_bar` rather than `FooBar`.
