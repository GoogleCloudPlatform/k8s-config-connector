### [2026-06-15] VectorSearchCollection Types and Identity Scaffolding
- **Context**: Implementing Greenfield direct types for `VectorSearchCollection` (`vectorsearch.cnrm.cloud.google.com/v1alpha1`).
- **Problem**: The pinned Google APIs SHA (`731d7f2ab6`) did not contain the standalone `google.cloud.vectorsearch.v1` proto definitions since it was added to the googleapis repository later.
- **Solution**: Updated `apis/git.versions` to use `120a55ddd98884993645c8ceb474dffbf8286595` (which contains the required `vectorsearch` surface).
- **Impact**: Enables `controllerbuilder` and generator tools to compile the `vectorsearch` protos properly.

### [2026-06-15] Robust Fallback for apiextensionsv1 Package Resolution in controllerbuilder
- **Context**: Running code generation pipelines that invoke `generate-mapper`.
- **Problem**: The AST inspector `inspect` in `dev/tools/controllerbuilder/pkg/gocode/ast.go` panics with `could not find import for "apiextensionsv1.JSON"` when visiting structs (e.g. in `dialogflow/v1alpha1`) that reference `apiextensionsv1.JSON` but do not have an explicit import. This crashed the entire code generation flow universally.
- **Solution**: Added a robust fallback in `ast.go` to dynamically resolve `apiextensionsv1` to `"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"` when matching the alias.
- **Impact**: Stabilizes the AST parser and ensures the mapper generator never crashes due to missing imports of standard types across any resource.

### [2026-08-20] VectorSearchCollection Global Endpoints and Schema Constraints
- **Context**: Implementing Greenfield direct controller and E2E fixtures for `VectorSearchCollection` (`vectorsearch.cnrm.cloud.google.com/v1alpha1`).
- **Problem**: Regionalizing endpoints (e.g. `us-central1-vectorsearch.googleapis.com`) results in 404 HTML errors on creation, and creating collections with empty `vector_schema` and `data_schema` returns a 400 Bad Request.
- **Solution**: Explicitly set the client endpoint to the global REST endpoint `https://vectorsearch.googleapis.com` and updated the minimal fixture to configure a simple `vectorSchema`.
- **Impact**: The controller reconciled correctly on real GCP, and recorded golden logs are extremely stable.

### [2026-10-29] DedicatedInfrastructure Optional Enum Presence on Roundtrip
- **Context**: Triaging judgement queue and finalizing types/fuzzer for `VectorsearchIndex` (`vectorsearch.cnrm.cloud.google.com/v1alpha1`).
- **Problem**: Proto field `DedicatedInfrastructure.mode` is an `optional Mode mode = 1` enum. `direct.Enum_FromProto` maps 0 (`MODE_UNSPECIFIED`) to `nil`, losing presence on roundtrip in `TestFocusedMappers`.
- **Solution**: Implemented `DedicatedInfrastructure_FromProto` in `vectorsearchcollection_mapper.go` to preserve pointer presence when `in.Mode != nil` via `direct.LazyPtr(in.GetMode().String())`, ensuring roundtrip fidelity between proto and KRM.
- **Impact**: Roundtrip fuzz testing passes cleanly for `VectorsearchIndex` without dropping zero-value enums.

### [2026-11-02] VectorsearchIndex Parent Reference Normalization and Update Reconciliation
- **Context**: Implementing Greenfield direct controller, E2E fixtures, and fuzzer for `VectorsearchIndex` (`vectorsearch.cnrm.cloud.google.com/v1alpha1`).
- **Problem**: `VectorsearchIndex` is a child resource under `VectorSearchCollection` (`spec.parentRef`). Its identity depends on `spec.parentRef` being normalized first, and static tools (`TestGoldenIdentitiesYamlFiles`, `TestFixtureExternalRefEvidence`) require resolving `{name: ...}` references from dependencies offline. Additionally, index mutations require `UpdateIndex` with LRO polling.
- **Solution**: Called `common.NormalizeGreenfieldReferences` before `obj.GetIdentity` in `AdapterForObject`, hand-maintained `VectorSearchCollectionRef.Normalize` using `refs.NormalizeWithFallback` and `getIdentityFromVectorSearchCollectionSpec`, and wired `UpdateIndex` with `updateMask` and `op.Wait(ctx)`.
- **Impact**: Ensures offline identity resolution in tests and successful end-to-end reconciliation on real GCP.

