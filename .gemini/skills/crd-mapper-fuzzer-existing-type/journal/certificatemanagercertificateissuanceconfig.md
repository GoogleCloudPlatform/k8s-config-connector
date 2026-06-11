# Journal: CertificateManagerCertificateIssuanceConfig KRM Types Implementation

## Observations

1. **Schema Integrity and Versioning:**
   - The `CertificateManagerCertificateIssuanceConfig` resource is currently only defined in the `v1alpha1` API version.
   - The existing `certificateissuanceconfig_types.go` file contains the manual type definitions, which inline `parent.ProjectAndLocationRef`. This handles both `spec.projectRef` and `spec.location` correctly, maintaining exact alignment with the baseline CRD schema.

2. **Skipped Types Reference:**
   - Updated `apis/certificatemanager/v1alpha1/generate.sh` to include the `--include-skipped-output` flag for the `generate-types` command.
   - This successfully generated commented-out references in `types.generated.go` for the existing manually defined types, providing excellent reference material for future developers while preserving schema compatibility.

3. **Mapper Code Generation & Compilation with `--multiversion`:**
   - In accordance with review feedback, we removed the redundant call to `generate-mapper` from `apis/certificatemanager/v1alpha1/generate.sh` and configured `apis/certificatemanager/v1beta1/generate.sh` to call `generate-mapper` with the `--multiversion` flag.
   - Using `--multiversion` appends version suffixes to generated mapper functions (e.g. `_v1beta1` or `_v1alpha1`).
   - We updated the custom handwritten mapper functions in `pkg/controller/direct/certificatemanager/certificate_mapper.go` to match the version suffix pattern (e.g., `_v1beta1_FromProto`), enabling the generator to identify and skip them cleanly.
   - We updated controllers and fuzzers (`dnsauthorization_controller.go`, `certificate_fuzzer.go`, `dnsauthorization_fuzzer.go`, `certificateissuanceconfig_fuzzer.go`) to use the newly versioned function signatures, ensuring the package compiles flawlessly.
