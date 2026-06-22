# Journal: Match MockGCP with Real GCP for BigtableGCPolicy

During the matching process for `BigtableGCPolicy` (`bigtable.cnrm.cloud.google.com`), several key insights and practices were identified:

## Direct/MockGCP Integration Structure

- `BigtableGCPolicy` is a resource configured on Google Cloud Bigtable Column Families.
- Rather than a separate REST or gRPC service, it maps under the hood to the `ModifyColumnFamilies` method of GCP's `BigtableTableAdmin` service (`google.bigtable.admin.v2.BigtableTableAdmin`).
- The mock implementation for this exists under `mockgcp/mockbigtable/table.go`.

## Service Enablement for Recording

When recording the live GCP baseline with `hack/record-gcp`, we encountered that the following APIs must be enabled on the target GCP project:
- `bigtableadmin.googleapis.com`
- `bigtable.googleapis.com`

We enabled them successfully using `gcloud services enable`:
```bash
gcloud services enable bigtableadmin.googleapis.com
gcloud services enable bigtable.googleapis.com
```

## API Log Differences & Normalization

Real GCP results differed from the previously captured golden logs in two subtle ways, which were captured and updated:
1. **Instance List Return:** Real GCP returned the `name` field for the Bigtable Instance (`"name": "projects/${projectId}/instances/gcpolicy-dep-${uniqueId}"`) which was omitted/different in older logs.
2. **Column Families Empty ValueType:** Real GCP did not return `"valueType": {}` within the `columnFamilies` returned map under `GetTable`, whereas older golden logs contained it.

Updating the golden log (`_http.log`) using `hack/record-gcp` correctly synchronized these and now both real and mock GCP executions pass flawlessly with zero golden-file discrepancies.
