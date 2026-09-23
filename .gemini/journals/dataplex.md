### [2026-05-26] Implement DataplexDataScan Types
- **Context**: Implementing KRM types for DataplexDataScan (google.cloud.dataplex.v1)
- **Problem**: DataplexDataScan requires manual scaffolding of inner struct fields inside `datascan_types.go` because the `prunetypes` generator removes nested types in `types.generated.go` if they are unreachable. Additionally, `k8s.io/apimachinery/pkg/runtime/schema` must be correctly imported in `datascan_reference.go` instead of `k8s.io/apimachinery/pkg/schema`.
- **Solution**: Scaffolded fields `DataQualitySpec`, `DataProfileSpec`, `DataDiscoverySpec` and their related `ResultObservedState` fields manually in `datascan_types.go`. Fixed import path in `datascan_reference.go`.
- **Impact**: Agents working on direct controllers for `dataplex` or any other nested specification APIs need to remember to correctly instantiate nested structure types in `datascan_types.go` and ensure proper imports for `schema.GroupVersionKind`.

### [2026-06-24] Customizing Sub-structures and References for DataplexMetadataJob
- **Context**: Implementing KRM types and IdentityV2 for `DataplexMetadataJob` under `v1alpha1`.
- **Problem**: The GCP `MetadataJob` API defines relative resource name fields as raw primitive strings/lists (e.g., `entry_groups`, `entry_types`, `aspect_types`). KCC guidelines mandate that all fields referencing other GCP/KCC resources must be proper reference fields (e.g., using `EntryGroupRef`, `EntryTypeRef`, `AspectTypeRef`) following the `Ref`/`Refs` suffix convention.
- **Solution**: Hand-coded customized sub-structure definitions (e.g., `MetadataJob_ImportJobSpec_ImportJobScope`, `MetadataJob_ExportJobSpec_ExportJobScope`) in `metadatajob_types.go` using reference types rather than raw strings. The controller builder tool/generator automatically mapping fields ending in `Refs` to matching repeated fields in proto (e.g., `entryGroupRefs` to `entry_groups` in proto), while also cleanly generating unchanged sub-structures in `types.generated.go`.
- **Impact**: Subsequent developers can leverage this pattern to customize nested sub-structures and lists of references directly in their `_types.go` files without causing duplicate type definitions in `types.generated.go`.

### [2026-06-29] Implement Greenfield DataplexMetadataJob Controller and Mocks
- **Context**: Implementing the Greenfield direct controller, E2E fixtures, and fuzzer for `DataplexMetadataJob` under `v1alpha1`.
- **Problem**: `DataplexMetadataJob` is a completely immutable resource in GCP (there is no `UpdateMetadataJob` or `DeleteMetadataJob` in the GCP API). Thus, it has no `Update` or `Delete` actions and needs to be registered with `registry.CannotBeDeleted()`. In mockgcp, long-running operations must also be correctly simulated so that the status is updated with completion results.
- **Solution**: Implemented `dataplexmetadatajob_controller.go` to reject modifications by returning an update error if any spec difference is found, and implemented a no-op `Delete` method. Registered the model with `registry.CannotBeDeleted()`. Created a comprehensive mock service in mockgcp (`mockdataplex/metadatajob.go`) to simulate `CreateMetadataJob`, `GetMetadataJob`, and `CancelMetadataJob`. Implemented fuzzer with `Unimplemented_NotYetTriaged` for unmapped fields and added timestamp normalization to prevent volatile comparison test failures.
- **Impact**: Demonstrates standard patterns for implementing direct controllers for immutable GCP resources and structuring mocks for custom LRO results.

### [2026-09-16] Implement Greenfield DataplexAspectType Controller, Fixtures and Fuzzer
- **Context**: Implementing the Greenfield direct controller, E2E fixtures, and fuzzer for `DataplexAspectType` under `v1alpha1`.
- **Problem**: Encountered multiple constraints and schema/documentation discrepancies during GCP recording:
  1. Top-level `metadataTemplate` of an `AspectType` must be of type `record`. Setting a primitive type at the root is rejected by GCP: `ValidationException: Top level record can only have type RECORD`.
  2. Documentation states primitive types are `integer` and `boolean`, but the real Dataplex API DataType Java enum maps them to `int` and `bool`. Setting `integer` yields: `ValidationException: No enum constant com.google.cloud.dataplex.template.types.DataType.INTEGER`.
  3. Fuzz tests identified an unmapped GCP-only proto field `.data_classification`.
- **Solution**:
  1. Restructured both minimal and maximal KRM spec fixtures to use a root `type: "record"`.
  2. Standardized the primitive integer fields to use `type: "int"`.
  3. Marked `.data_classification` as `Unimplemented_NotYetTriaged` in the fuzzer.
- **Impact**: Real GCP tests fully verified. Keeps documentation/behavior aligned for other Dataplex template-based resources.

### [2026-09-17] Implement DataplexDataTaxonomy Direct Controller and E2E Tests
- **Context**: Implementing the Greenfield direct controller, E2E fixtures, and fuzzer for `DataplexDataTaxonomy` under `v1alpha1`.
- **Problem**: Dataplex `DataTaxonomy` creation requests sent to the GCP API in regional locations like `us-central1` can consistently fail with gRPC code 13 (`INTERNAL` error: `An internal error has occurred`) even when all API enablements and IAM permissions are verified correct. Multi-region and global locations are rejected as invalid path locations.
- **Solution**: Hand-coded the direct controller and fuzzer following KCC standard practices and verified complete fuzzer field mapping coverage with the targeted test framework (`FOCUS=DataplexDataTaxonomy go test ./pkg/fuzztesting/fuzztests/ -v -run=TestFocusedMappers`). Updated API field presence exceptions golden file.
- **Impact**: Highlights potential GCP service-side limitations or quirks for the `DataTaxonomy` API in test projects, requiring the direct controller logic and schema mappings to be carefully designed and tested offline/via fuzzer when live API execution is obstructed by service-side Internal errors.

### [2026-09-23] Implement Direct KRM Types and Identity for DataplexDataProduct
- **Context**: Implementing KRM types and IdentityV2 for `DataplexDataProduct` under `v1alpha1`.
- **Problem**: `DataplexDataProduct` contains a map field `access_groups` which in protobuf translates to a nested map type (`map<string, AccessGroup>`). The default KCC `controllerbuilder generate-types` tool does not support direct generation of protobuf maps with message/struct values, reporting them as unreachable or unsupported. Furthermore, the Go client library for Dataplex in KCC was pinned to an older version (`v1.34.0`) that lacked the newly added `DataProduct` proto types entirely, leading to compilation issues.
- **Solution**:
  1. Updated `apis/git.versions` to point to a newer `googleapis` commit (`28ba5d15234eff44241492ca72b135e1171fc537`) containing `data_products.proto`.
  2. Upgraded `cloud.google.com/go/dataplex` dependency in `go.mod` and `mockgcp/go.mod` to `v1.36.0` to pull in the compiled proto definitions.
  3. Manually specified the spec and observed state fields in `dataplexdataproduct_types.go`, including defining `AccessGroups` as `map[string]DataProduct_AccessGroup`.
  4. Handwrote custom mapping functions `AccessGroups_FromProto` and `AccessGroups_ToProto` in `pkg/controller/direct/dataplex/dataproduct_mapping.go` to successfully translate between KRM maps and protobuf maps.
  5. Implemented `DataProductIdentity` utilizing `gcpurls.Template` with standard Project and Location references.
- **Impact**: Provides a robust template for implementing KRM types, hand-written map mappers, and custom GCP-to-KRM model mapping for complex direct resources.
