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
