# BigQuery Journals

### 2026-06-12 Refactored BigQueryDataset direct controller to use proto-format desired state pattern
- **Context**: Refactoring BigQueryDataset direct controller to meet KCC's latest quality guidelines (Issue #9767).
- **Problem**: The original controller was holding a reference to the raw KRM object in the adapter, leading to multiple conversions, manual reference resolution inside `Create`/`Update`, and complex logic.
- **Solution**:
  1. Converted KRM Spec to the Go library-specific `*bigquery.DatasetMetadata` representation once in `AdapterForObject`.
  2. Resolved the KMS crypto key reference and copied Kube metadata labels during the adapter creation phase.
  3. Stored helper pointers like `isCaseInsensitive *bool` on `Adapter` to cleanly check for presence/omission of optional fields.
  4. Simplified `Create` and `Update` reconciliation methods to operate directly on the pre-converted desired fields.
- **Impact**: Code complexity is significantly reduced, the controller conforms to KCC's latest architecture/expert standards, and subsequent maintenance is much easier.
