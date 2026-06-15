### 2026-06-15 Greenfield Implementation of ParallelstoreInstance Types
- **Context**: Implementing initial direct KRM types, CRD, and IdentityV2 for `ParallelstoreInstance` under the service `google.cloud.parallelstore.v1`.
- **Problem**: Greenfield scaffolding of types via `controllerbuilder` requires creating package boilerplate (`doc.go`, `groupversion_info.go`, `generate.sh`) and then cleanly separating generated files and identity files.
- **Solution**: 
  1. Scaffolded `apis/parallelstore/v1alpha1/doc.go` and `groupversion_info.go`.
  2. Defined `apis/parallelstore/v1alpha1/generate.sh` which only invokes `--resource ParallelstoreInstance:Instance` to generate types and CRD.
  3. Formatted and validated fields from the generated unreachable structs inside `types.generated.go`.
  4. Implemented `ParallelstoreInstanceIdentity` and `ParallelstoreInstanceRef` utilizing the `gcpurls.Template` design.
- **Impact**: Provides a clean pattern for scaffolding new direct KCC types.
