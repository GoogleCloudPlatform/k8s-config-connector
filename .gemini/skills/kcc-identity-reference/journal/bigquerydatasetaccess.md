# BigQueryDatasetAccess Identity and Reference Journal

## Observations

- `BigQueryDatasetAccess` is not a standalone GCP REST resource with its own unique endpoint; rather, it represents a nested entry in a `BigQueryDataset` access list.
- In Terraform and Config Connector, it is managed as an independent resource. Its ID is constructed as `projects/{project}/datasets/{dataset}`.
- Consequently, we implemented its identity format to match the parent Dataset's format: `"projects/{project}/datasets/{dataset}"`.
- Since `"//bigquery.googleapis.com/projects/{}/datasets/{}"` is already defined in `cloudassetinventory_names.jsonl` (associated with `BigQueryDataset`), registering this template does not violate `TestRegisteredTemplatesMatchCAI`, and no manual exception had to be added in `pkg/gcpurls/registry_test.go`.
- The `BigQueryDatasetAccess` resource is located under `apis/bigquery/v1alpha1`, so the identity and reference files were created in that directory.
