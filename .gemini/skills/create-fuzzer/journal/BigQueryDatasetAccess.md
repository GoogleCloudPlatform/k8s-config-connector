# Journal: BigQueryDatasetAccess Fuzzer Implementation

## Observations & Learnings

### 1. Fuzzing BigQueryDatasetAccess with NoProto (OpenAPI) Mappers
- `BigQueryDatasetAccess` represents a single dataset access rule in BigQuery. While the KRM type is annotated with `// +kcc:spec:proto=google.cloud.bigquery.v2.Access`, the Google Cloud BigQuery v2 API is exposed primarily via the REST / OpenAPI client `google.golang.org/api/bigquery/v2`.
- In `google.golang.org/api/bigquery/v2`, the access item struct is `DatasetAccess`, containing:
  - `Dataset *DatasetAccessEntry`
  - `Domain string`
  - `GroupByEmail string`
  - `IamMember string`
  - `Role string`
  - `SpecialGroup string`
  - `UserByEmail string`
  - `View *TableReference`
  - Untriaged fields: `Condition *Expr`, `Routine *RoutineReference`
  - Bookkeeping fields: `ForceSendFields []string`, `NullFields []string`
- The KRM Spec fields (`dataset`, `domain`, `groupByEmail`, `iamMember`, `role`, `specialGroup`, `userByEmail`, `view`) map directly to these fields in `DatasetAccess`.
- The KRM Spec fields `datasetId`, `projectRef`, and `resourceID` represent hierarchical resource identity in KCC and are not part of the `DatasetAccess` API payload body.
- The KRM Status has no server-reported status fields from `DatasetAccess` (`apiUpdatedMember`, `conditions`, and `observedGeneration` are KCC-internal / client-side status fields). Status mappers return empty structs.
- In `BigQueryDatasetAccessDatasetAccessEntry`, the nested `dataset` field is a value struct (`BigQueryDatasetAccessDataset`), while `api.DatasetAccessEntry.Dataset` is a pointer (`*api.DatasetReference`). In `FilterSpec`, we ensure that if `Dataset` is present on `DatasetAccess`, its nested `Dataset` reference pointer is initialized (`in.Dataset.Dataset = &api.DatasetReference{}`) to avoid round-trip diffs between nil and zero-value structs.
- Implemented `pkg/controller/direct/bigquery/bigquerydatasetaccess_mappings.go` and `pkg/controller/direct/bigquery/bigquerydatasetaccess_fuzzer.go`.
- Registered via `fuzztesting.RegisterKRMFuzzer_NoProto(bigqueryDatasetAccessFuzzer())`.
- Verified round-trip tests pass via `FOCUS=BigQueryDatasetAccess go test -count=1 -v ./pkg/fuzztesting/fuzztests/ -run TestFocusedMappers`.
