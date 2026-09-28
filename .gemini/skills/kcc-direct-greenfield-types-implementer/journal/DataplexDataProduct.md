# DataplexDataProduct Journal Entry

### [2026-09-25] Implement Direct KRM Types and Identity for DataplexDataProduct
- **Context**: Implementing KRM types and IdentityV2 for `DataplexDataProduct` under `v1alpha1`.
- **Problem**: `DataplexDataProduct` contains a map field `access_groups` which in protobuf translates to a nested map type (`map<string, AccessGroup>`). The default KCC `controllerbuilder generate-types` tool does not support direct generation of protobuf maps with message/struct values, reporting them as unreachable or unsupported. Furthermore, the Go client library for Dataplex in KCC was pinned to an older version (`v1.34.0`) that lacked the newly added `DataProduct` proto types entirely, leading to compilation issues.
- **Solution**:
  1. Updated `apis/git.versions` to point to a newer `googleapis` commit (`28ba5d15234eff44241492ca72b135e1171fc537`) containing `data_products.proto`.
  2. Upgraded `cloud.google.com/go/dataplex` dependency in `go.mod` and `mockgcp/go.mod` to `v1.36.0` to pull in the compiled proto definitions.
  3. Manually specified the spec and observed state fields in `dataplexdataproduct_types.go`, including defining `AccessGroups` as `map[string]DataProduct_AccessGroup`.
  4. Handwrote custom mapping functions `AccessGroups_FromProto` and `AccessGroups_ToProto` in `pkg/controller/direct/dataplex/dataproduct_mapping.go` to successfully translate between KRM maps and protobuf maps.
  5. Implemented `DataProductIdentity` utilizing `gcpurls.Template` with standard Project and Location references.
- **Impact**: Provides a robust template for implementing KRM types, hand-written map mappers, and custom GCP-to-KRM model mapping for complex direct resources.
