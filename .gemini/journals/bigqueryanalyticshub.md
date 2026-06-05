# BigQuery Analytics Hub Journal

### [2026-06-05] BigQueryAnalyticsHubListing Update Mask Quirk
- **Context**: Implementing direct controller for `BigQueryAnalyticsHubListing`.
- **Problem**: When updating a `Listing`, the GCP API (via `grpc-gateway` and `google-cloud-go`) uses camelCase for fields in the `updateMask`. For the `source` oneof, the specific field name must be used in the mask (e.g., `bigqueryDataset`), not the oneof name itself (`source`).
- **Solution**: Updated MockGCP's `UpdateListing` to handle `bigqueryDataset` and other camelCase field names in the `updateMask`.
- **Impact**: Future updates to `Listing` or similar resources with `oneof` fields should be aware that the `updateMask` will contain the specific field name from the `oneof`.
