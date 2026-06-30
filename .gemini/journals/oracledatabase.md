### [2026-06-15] Scaffolding OracleDatabaseAutonomousDatabase
- **Context**: Implementing direct types, CRD, and IdentityV2 for `OracleDatabaseAutonomousDatabase` (Kind: `OracleDatabaseAutonomousDatabase`).
- **Problem**: 
  1. Compiling direct types via `generate.sh` can fail with `no such file or directory` if `.build/googleapis.pb` is missing.
  2. The `gcpurls.Template` tool can panic if the variable names in the template (e.g. `{autonomous_database}`) do not match the CamelCased struct fields of the identity (e.g. `AutonomousDatabase`).
  3. Staging and validation checks require client-go SDK and CI workflow files to be regenerated and staged.
- **Solution**:
  1. Execute `./dev/tools/controllerbuilder/generate-proto.sh` first to compile the googleapis descriptor.
  2. In the `gcpurls.Template` definition, use `{autonomousDatabase}` matching the struct field name `AutonomousDatabase`.
  3. Run `dev/tasks/generate-ci-cd-jobs`, stage all newly created types, CRD, clients, and then run `dev/tasks/generate-github-actions` to regenerate the workflow definitions.
- **Impact**: Helps future developers/agents avoid template parsing panics, missing build cache errors, and validation mismatches when adding new GCP service types to Config Connector.
