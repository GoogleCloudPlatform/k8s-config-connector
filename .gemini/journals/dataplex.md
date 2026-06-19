# Dataplex Journal

### 2026-06-05 Implementing DataplexDataTaxonomy Direct Types
- **Context**: Implementing KRM types, CRD, and IdentityV2 for DataplexDataTaxonomy.
- **Problem**: The proto-to-KRM generator initially marked `DataTaxonomy` and `DataTaxonomyObservedState` as unreachable because the scaffolded `DataplexDataTaxonomySpec` and `DataplexDataTaxonomyObservedState` were empty, causing those fields/types to be commented out as unreachable in `types.generated.go`.
- **Solution**: We defined `DataplexDataTaxonomySpec` to use `parent.ProjectAndLocationRef` and standard fields, and implemented `DataplexDataTaxonomyObservedState` with the proper output-only fields (`Uid`, `CreateTime`, `UpdateTime`, `AttributeCount`, `ClassCount`). Running the generator again successfully resolved all unreachability and correctly generated the deepcopy and CRD files.
- **Impact**: When adding greenfield resources, always ensure Spec and ObservedState map the proto fields before final generation so that types are correctly kept in `types.generated.go`.

### 2026-06-05 DataplexDataProduct Initial Greenfield Implementation
- **Context**: Implementing the greenfield `DataplexDataProduct` resource (initial types, Identity, and References).
- **Problem**: 
  1. The default googleapis commit pin (`1765b559c42386788ff0c6412491277b4791107a`) did not contain the `google/cloud/dataplex/v1/data_products.proto` definitions, which were introduced in later commit `69ca7ae2e` ("feat(dataplex): add DataProductService...").
  2. The `access_groups` field is of type `map<string, AccessGroup>` which the generator flagged as unsupported and skipped.
- **Solution**:
  1. Updated `apis/git.versions` to use the latest googleapis master commit `ee4a3e1ce4e8d16661fcc624322000ad792ffb8a` containing `data_products.proto` and regenerated the types.
  2. Defined `AccessGroups` as a slice of `DataProduct_AccessGroup` in `DataplexDataProductSpec` to natively support it in KRM. Added a `projects/{project}/locations/{location}/dataProducts/{dataproduct}` template format exception to `pkg/gcpurls/registry_test.go` as it is not currently published in CAI names.
- **Impact**: Allows the next agents implementing controller/mapper logic for DataplexDataProduct to have fully compiling type structures and valid IdentityV2 / references.
