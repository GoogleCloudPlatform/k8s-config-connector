# KCC Direct Migration Journal - IAPBrand

## Observations

- **Multi-version Name Collisions**: Since `apis/iap/` contains both `v1alpha1` and `v1beta1` directories with `*_types.go` definitions, the mapper generator `generate-mapper` will scan all directories and generate mapping functions for both versions. Without the `--multiversion` flag on `generate-mapper`, identical duplicate function names are generated for both versions, leading to Go compilation errors (`redeclared in this block`).
- **Use of `--multiversion`**: Adding the `--multiversion` flag to `generate-mapper` solved the collision by appending `_v1beta1` and `_v1alpha1` suffixes to the generated function signatures.
- **Renaming Hand-coded Mappers**: When `--multiversion` is active, hand-coded mappings must also use version suffixes (e.g. `_v1beta1_ToProto`) so the generator correctly identifies them as existing mappings and skips generating conflicting copies.
- **Mappers File Renaming**: Hand-coded mappers were consolidated into `pkg/controller/direct/iap/mappers.go` instead of `iapsettings_mappings.go` to follow the standard `mappers.go` naming convention.
- **Spec-Only Fuzzer**: Since `orgInternalOnly` is output-only and mapped to status, we used `NewKRMTypedSpecFuzzer` for a spec-only round-trip check and listed `.org_internal_only` under `Unimplemented_NotYetTriaged` in `iapbrand_fuzzer.go`.
