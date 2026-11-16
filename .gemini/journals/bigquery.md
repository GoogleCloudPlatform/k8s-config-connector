# BigQuery

## Greenfield Notes

### BigQueryRowAccessPolicy
- Proto: `google.cloud.bigquery.v2.RowAccessPolicy` in `google/cloud/bigquery/v2/row_access_policy.proto`.
- Resource has two resource name patterns declared:
  - `projects/{project}/datasets/{data_set}/tables/{table}/rowAccessPolicies/{row_access_policy}`
  - `projects/{projectsid}/datasets/{datasetsid}/tables/{tablesid}/rowAccessPolicies/{rowaccesspoliciesid}`
- Defaulted to `projects/{project}/datasets/{data_set}/tables/{table}/rowAccessPolicies/{row_access_policy}`.
- Added to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle`.
- Greenfield Step 1b (types-judgement):
  - Parent hierarchy modeled as single required immutable `spec.parentRef` (`*bigqueryv1beta1.BigQueryTableRef`) with CEL immutability rule per `PARENT_AND_IDENTITY.md` Case B.
  - Removed `projectRef`, `location`, and self-referential `rowAccessPolicyReference` struct from `BigQueryRowAccessPolicySpec`.
  - Locked identity pattern via `--pattern BigQueryRowAccessPolicy=projects/{project}/datasets/{dataset}/tables/{table}/rowAccessPolicies/{rowAccessPolicy}` in `apis/bigquery/generate.sh` to match `BigQueryTable`'s `{dataset}` segment naming.
  - Updated resource docs link to `https://cloud.google.com/bigquery/docs/reference/rest/v2/rowAccessPolicies`.
- Greenfield Step 2 (controller):
  - Controller generated using BigQuery REST client (`google.golang.org/api/bigquery/v2`).
  - BigQuery `RowAccessPolicies.Get` does not return `grantees` in the response body; instead, the grantees are returned as IAM policy bindings for role `roles/bigquery.filteredDataViewer`. The controller calls `RowAccessPolicies.GetIamPolicy` in `Find` to populate `actual.Grantees` for drift detection.
  - Deletion requires `.Force(true)` if deleting the last row access policy on a table to avoid GCP 400 `Dropping the last row access policy would make the table accessible to all users... If this is intended, please set force to true.`
  - Fuzzer implemented using `fuzztesting.NewKRMTypedFuzzer_NoProto` with `FilterSpec` and `FilterStatus` clearing `RowAccessPolicyReference` (which is modeled as `parentRef` and `resourceID`).
  - E2E basic fixture verified and recorded against real GCP (`hack/record-gcp`), passing all 7 automated validation checks.

