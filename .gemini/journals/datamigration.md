### [2026-06-05] Datamigration (CloudDMS) Greenfield Scaffolding
- **Context**: Greenfield types scaffolding for `DatabaseMigrationConversionWorkspace`.
- **Problem**: Running `generate-types` with `--service google.cloud.datamigration.v1` fails since the actual underlying Google API and protobuf package is named `google.cloud.clouddms.v1`, despite the service endpoint being `datamigration.googleapis.com`.
- **Solution**: Set `--service` parameter to `google.cloud.clouddms.v1` inside the `generate.sh` script, while keeping the group as `datamigration.cnrm.cloud.google.com` to achieve the target KRM API version and structure.
- **Impact**: Enables successful code generation and alignment for direct resources belonging to Database Migration Service (DMS).
