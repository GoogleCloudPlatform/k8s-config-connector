### [2026-06-06] BigLakeCatalog Identity & Refs
- **Context**: Moving BigLakeCatalog to the `identity.IdentityV2` and `refs.Ref` patterns as per issue #9452.
- **Problem**:
    1. BigLakeCatalog was using the older `identity.Identity` with custom string manipulation in `catalog_identity.go`.
    2. The parent resolution and string formatting was not utilizing `gcpurls.Template`.
- **Solution**:
    1. Reimplemented `apis/bigquerybiglake/v1alpha1/catalog_identity.go` to use `identity.IdentityV2` backed by `gcpurls.Template`.
    2. Rewrote `catalog_reference.go` to register and utilize the `BigLakeCatalogIdentity` via `refs.NormalizeWithFallback`.
    3. Updated `database_identity.go` (the child of BigLakeCatalog) to reference `BigLakeCatalogIdentity` instead of `CatalogIdentity`.
    4. Regenerated the generated deepcopy code using `dev/tasks/generate-types-and-mappers` to remove deepcopy methods for the identity struct (since `BigLakeCatalogIdentity` has `// +k8s:deepcopy-gen=false`).
    5. Added BigLakeCatalog format exemption to `pkg/gcpurls/registry_test.go` as it is not present in CAI.
    6. Created a new comprehensive unit test file `biglakecatalog_identity_test.go` to thoroughly test parsing and conversion.
- **Impact**: BigLakeCatalog now perfectly conforms to the modern, type-safe, and robust identity and reference patterns.
