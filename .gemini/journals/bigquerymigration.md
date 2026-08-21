# BigQueryMigration Journal

### [2026-06-03] Resolving reaching transitive types through KRM Spec mapping
- **Context**: Implementing direct types, CRD, and IdentityV2 for `BigQueryMigrationMigrationWorkflow:MigrationWorkflow`.
- **Problem**: The GCP API `MigrationWorkflow` resource contains a `tasks` field which is a `map[string]MigrationTask`. Because KCC's `generate-types` uses strict reachability pruning (`--prune-unused-types=true`), it originally pruned and commented out `MigrationTask` and all of its referenced schemas as unreachable types inside `types.generated.go`.
- **Solution**: By manually defining the `DisplayName` and `Tasks map[string]MigrationTask` fields on the `BigQueryMigrationMigrationWorkflowSpec` struct inside `migrationworkflow_types.go` and using the proper `+kcc:proto:field` tag comments, the controllerbuilder tool successfully detects transitive reachability. On the subsequent run, it automatically uncomments all transitive Go structs in `types.generated.go`.
- **Impact**: When scaffolding new direct resources with complex map or list-of-message fields, always declare the top-level speculative fields first, then regenerate types so the nested structs are automatically unpacked.
