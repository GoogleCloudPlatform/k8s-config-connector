# Journal for Solve Migration Diff Issues for Direct Takeover

## 2026-06-12 - BigQueryTable View Schema Diff
- When migrating resources with derived/computed fields (e.g., a BigQuery view or materialized view's `schema`), the field is typically omitted in the KRM spec but returned by the GCP API.
- Unlike Terraform (which automatically handles this by storing the fetched value in the state file because the field is `Computed`), Kubernetes/KCC is strictly declarative based on the KRM spec. If the spec has a `nil` schema, the controller will attempt to clear/delete the schema on GCP during reconciliation.
- To fix this, you must write a custom comparison function (e.g., `tableSchemaEq`) that returns `true` (no diff) if the `desired` schema is `nil`, regardless of the `actual` schema returned by GCP.
- **Gotcha (Parameter Swap):** In comparison files, parameter names like `a` and `b` are extremely error-prone. They are easily swapped between caller and callee, causing diffs to be reported backwards or checks to be bypassed. Always name them `actual` and `desired` consistently throughout the file.
