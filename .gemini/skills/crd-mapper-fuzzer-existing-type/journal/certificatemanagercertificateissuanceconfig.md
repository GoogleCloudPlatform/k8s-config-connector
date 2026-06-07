# Journal: CertificateManagerCertificateIssuanceConfig KRM Types Implementation

## Observations

1. **Schema Integrity and Versioning:**
   - The `CertificateManagerCertificateIssuanceConfig` resource is currently only defined in the `v1alpha1` API version.
   - The existing `certificateissuanceconfig_types.go` file contains the manual type definitions, which inline `parent.ProjectAndLocationRef`. This handles both `spec.projectRef` and `spec.location` correctly, maintaining exact alignment with the baseline CRD schema.

2. **Skipped Types Reference:**
   - Updated `apis/certificatemanager/v1alpha1/generate.sh` to include the `--include-skipped-output` flag for both `generate-types` and `generate-mapper` commands.
   - This successfully generated commented-out references in `types.generated.go` for the existing manually defined types, providing excellent reference material for future developers while preserving schema compatibility.

3. **Mapper Code Generation & Compilation:**
   - Running the generation scripts successfully populated `pkg/controller/direct/certificatemanager/mapper.generated.go` with `FromProto` and `ToProto` functions for `CertificateManagerCertificateIssuanceConfig`.
   - The generator correctly updated/adapted the imports and namespaces for `v1alpha1` (e.g. mapping `krm` to `apis/certificatemanager/v1alpha1`), ensuring the entire package compiles successfully.
