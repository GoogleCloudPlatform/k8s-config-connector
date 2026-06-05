# Database Migration Service Journal

### [2026-06-05] Datamigration (CloudDMS) DatabaseMigrationConversionWorkspace Greenfield Scaffolding
- **Context**: Greenfield types scaffolding for `DatabaseMigrationConversionWorkspace`.
- **Problem**: Running `generate-types` with `--service google.cloud.datamigration.v1` fails since the actual underlying Google API and protobuf package is named `google.cloud.clouddms.v1`, despite the service endpoint being `datamigration.googleapis.com`.
- **Solution**: Set `--service` parameter to `google.cloud.clouddms.v1` inside the `generate.sh` script, while keeping the group as `datamigration.cnrm.cloud.google.com` to achieve the target KRM API version and structure.
- **Impact**: Enables successful code generation and alignment for direct resources belonging to Database Migration Service (DMS).

### [2026-06-05] DatabaseMigrationMigrationJob Greenfield Types Implementation
- **Context**: Greenfield implementation of `DatabaseMigrationMigrationJob` type scaffolding, IdentityV2, and CRD generation under the new `datamigration` group (`datamigration.cnrm.cloud.google.com/v1alpha1`).
- **Problem**: The GCP service uses the host `datamigration.googleapis.com` but the proto package name is registered as `google.cloud.clouddms.v1`. When running the `controllerbuilder` generator tool, using `--service google.cloud.datamigration.v1` caused a "failed to find the proto message" error because the proto file structure in googleapis is defined under `clouddms`.
- **Solution**: Set `--service google.cloud.clouddms.v1` in the `generate.sh` CLI invocation, while keeping `--api-version datamigration.cnrm.cloud.google.com/v1alpha1` and resource name mappings matching the new `DatabaseMigrationMigrationJob`. This correctly aligned KCC's codegen to read from the compiled `clouddms` protobuf schemas while scaffolding types under the brand-new KCC `datamigration` API group.
- **Impact**: Enables next-generation `datamigration` resources to be successfully scaffolded using the `clouddms` protos without needing duplicate proto copies or experiencing generator lookup errors.

### 2026-06-05 DatabaseMigrationConnectionProfile Greenfield Implementation
- **Context**: Implementing the first step of the Greenfield resource types for `DatabaseMigrationConnectionProfile` (under KRM API group `datamigration.cnrm.cloud.google.com`).
- **Problem**: The issue description asked for scaffolding of the resource using `google.cloud.datamigration.v1` as the proto service name. However, GCP's underlying proto package for this service in the cloned Google APIs is actually `google.cloud.clouddms.v1` (with host `datamigration.googleapis.com`), and `google.cloud.datamigration.v1` is not present, leading to "proto message not found" errors in the code generator.
- **Solution**: Set the `--service` flag in `generate.sh` to `google.cloud.clouddms.v1` but mapped the output Kubernetes group `--api-version` to `datamigration.cnrm.cloud.google.com/v1alpha1`. This correctly pulled the `ConnectionProfile` proto message from the correct package while outputting the files into the `apis/datamigration/` directory with the desired group.
- **Impact**: Any future implementers of `DatabaseMigrationConnectionProfile` controllers, mappers, or related resources must use the `google.cloud.clouddms.v1` proto package and services for communication with the GCP API.
