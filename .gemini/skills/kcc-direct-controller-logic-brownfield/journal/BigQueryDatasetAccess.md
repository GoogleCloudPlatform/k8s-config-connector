# BigQueryDatasetAccess Direct Controller Journal

## Overview
Migrated `BigQueryDatasetAccess` from the legacy Terraform-based reconciler to the `direct` reconciler architecture.

## Key Observations and Decisions
1. **Subresource Nature on BigQuery Dataset:**
   - In the BigQuery REST API, dataset access entries are not standalone REST resources; they are represented as elements within the `access` array of the parent `Dataset` resource (`projects/{project}/datasets/{dataset_id}`).
   - `Find`: Retrieves the parent dataset via `Datasets.Get` and iterates through `dataset.Access` using matching criteria based on role (canonicalizing predefined roles like `roles/bigquery.dataViewer` to basic roles like `READER`) and identity (`userByEmail`, `groupByEmail`, `domain`, `specialGroup`, `iamMember`, `view`, or `dataset`).
   - `Create`: Retrieves the parent dataset's existing access list, checks for pre-existence, appends the new access entry, and calls `Datasets.Patch`. Also tracks whether GCP translated `iamMember` to one of the basic member types to set `status.apiUpdatedMember`.
   - `Update`: `BigQueryDatasetAccess` is immutable in BigQuery. The `Update` method verifies that the current entry still matches desired spec and returns an error if any immutable field is modified, ensuring 0 writes during steady-state re-reconciliation.
   - `Delete`: Retrieves the parent dataset, removes the matching access entry from the access list, and calls `Datasets.Patch`.

2. **Dual-Controller Configuration:**
   - Registered `k8s.ReconcilerTypeDirect` in `SupportedControllers` for `BigQueryDatasetAccess` in `pkg/controller/resourceconfig/static_config.go` while keeping `DefaultController` as `k8s.ReconcilerTypeTerraform`.

3. **Diff Analysis & Verification:**
   - Recorded authentic golden HTTP logs against real GCP (`cnrm-barni-4`) using `./hack/record-gcp` and verified MockGCP alignment using `./hack/compare-mock`.
   - In `_http.diff` and `_http_mock.diff`, the direct controller avoids sending the extraneous `datasetId` attribute inside the `DatasetAccess` JSON payload, which Terraform MMv1 previously serialized from internal resource state.
   - In `_final_object.diff`, legacy-specific annotations (`cnrm.cloud.google.com/management-conflict-prevention-policy: none` and `cnrm.cloud.google.com/state-into-spec: absent`) are omitted in direct mode.
