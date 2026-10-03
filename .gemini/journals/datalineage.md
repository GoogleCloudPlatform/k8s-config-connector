### [2026-05-26] DataLineageProcess Attributes Field
- **Context**: Implementing `DataLineageProcess` KRM types using the `generate-types` scripter.
- **Problem**: The proto contains a field `map<string, google.protobuf.Value> attributes`. The `generate-types` script fails to generate this, marking it as `// TODO: unsupported map type with key string and value message` and creating unreachable types for `ListValue` and `Value`.
- **Solution**: Manually declared the `Attributes` field in `DataLineageProcessSpec` as `map[string]apiextensionsv1.JSON \`json:"attributes,omitempty"\``. Added the other unreachable generated types like `Origin` and `DisplayName` directly to the `process_types.go` file.
- **Impact**: When scaffolding new `datalineage` resources containing `google.protobuf.Value` maps, these fields must be manually ported from the unreachable types in `types.generated.go` over to `_types.go` using `apiextensionsv1.JSON`.

### [2026-10-01] DataLineageProcess MockGCP Log Alignment
- **Context**: Aligning MockGCP output with real GCP logs for `DataLineageProcess`.
- **Findings**: Verified both `datalineageprocess-minimal` and `datalineageprocess-maximal` test fixtures. MockGCP `LineageServer` correctly emulates CRUD operations, OperationMetadata on LRO responses, and normalizations (`Previsit` placeholder replacements for timestamps, project numbers, and unique IDs).
- **Result**: `TestGoldenLogAlignment` and `TestAllInSeries` e2e fixture runs pass hermetically with zero discrepancies against real GCP golden logs.

