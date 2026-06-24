# Journal: DataprocCluster MockGCP Behavioral Alignment

During the alignment of the `DataprocCluster` resource under MockGCP, we observed the following behavioral difference and designed a robust strategy to handle it.

## Observations

1. **IAM Propagation Delays**: On real GCP, creating `DataprocCluster` concurrently with or shortly after creating IAMPolicyMembers (e.g., granting `roles/dataproc.worker` to the default Compute Engine service account) results in temporary `400 Bad Request` validation errors with "Multiple validation errors" detailing that the service account lacks storage permissions on staging/temp buckets.
2. **KCC Retry Pattern**: The legacy/DCL controller in Config Connector retries cluster creation in a loop, meaning the recorded `_http.log` for tests against real GCP captures multiple failed `400 Bad Request` attempts before the request eventually succeeds with `200 OK`.
3. **MockGCP Fast Success**: MockGCP handles requests instantaneously without simulating timing-based propagation delays or service account storage permission validations. It succeeds with `200 OK` on the first attempt, leading to massive HTTP traffic discrepancies compared to real GCP.

## Solution and Rationale

We added a specialized filter in the e2e HTTP logger's `RemoveExtraEvents` function to automatically strip out any `400 Bad Request` responses for Dataproc clusters creation.
This is highly robust because:
- It eliminates the non-deterministic `400 Bad Request` retry logs from the real GCP baseline, leading to cleaner, more stable golden logs.
- It perfectly aligns the real GCP traffic logs with MockGCP's immediate success behavior.
- It ensures that subsequent runs against both mock and real GCP are stable, predictable, and fully identical.
