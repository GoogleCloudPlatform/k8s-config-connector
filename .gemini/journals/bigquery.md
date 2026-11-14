# BigQuery

## Greenfield Notes

### BigQueryRowAccessPolicy
- Proto: `google.cloud.bigquery.v2.RowAccessPolicy` in `google/cloud/bigquery/v2/row_access_policy.proto`.
- Resource has two resource name patterns declared:
  - `projects/{project}/datasets/{data_set}/tables/{table}/rowAccessPolicies/{row_access_policy}`
  - `projects/{projectsid}/datasets/{datasetsid}/tables/{tablesid}/rowAccessPolicies/{rowaccesspoliciesid}`
- Defaulted to `projects/{project}/datasets/{data_set}/tables/{table}/rowAccessPolicies/{row_access_policy}`.
- Added to `BLOCKED_KINDS` in `dev/tasks/build-release-bundle`.
