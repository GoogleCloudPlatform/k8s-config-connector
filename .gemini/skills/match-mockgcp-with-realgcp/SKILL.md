# Skill: Match Mock behaviour with real GCP api

This skill provides a structured workflow for matching the mock {service}{resource} behaviour with the real GCP.

When the golden tests for K8s Config Connector mock output diverge from real GCP output, we need to inspect the discrepancies and fix either the mock implementation or the normalizers. This ensures that `hack/compare-mock` runs cleanly and accurately represents GCP API contracts.

## Workflow

### Align Mockgcp

1.  Run `hack/record-gcp "fixtures/^<testname>$"` to capture real GCP behavior.
    *   **Troubleshooting Service Not Enabled**: If `hack/record-gcp` fails because a GCP service is not enabled (e.g., error mentions that the API is disabled or has not been used in the project before), enable the service using `gcloud` and try again:
        ```bash
        gcloud services enable <service-name>.googleapis.com
        ```
        *(For example: `gcloud services enable compute.googleapis.com` or `gcloud services enable run.googleapis.com`)*
2.  Run `hack/compare-mock "fixtures/^<testname>$"` to check mock behavior.
3.  Iteratively fix discrepancies in the mock implementation or `normalize.go`.

#### Tips for fixing the discrepancies:

1. Look closely at the `compare-mock` HTTP log differences (typically mock on left `=>` real on right).
2. For missing default values (e.g., `<missing> => REGIONAL`), add a `populateDefaultsFor<Resource>` function to the mock service's file (e.g. `mockgcp/mockcompute/networksv1.go`). Make sure it is called on `Insert` and `Get`.
3. For generated IDs or volatile values (e.g. IPs, resource URLs) where real GCP generates dynamically but mockgcp outputs something static, you need to update the normalizer `mockgcp/mock<service>/normalize.go`.
4. Run `hack/compare-mock "fixtures/^<testname>$"` to see the diff and overwrite `_http.log`.
5. Run `git diff` on the test fixtures to ensure that the golden `_http.log` accurately replaces volatile data with placeholder variables (e.g., `${ipAddress}`).
6. Certain operation metadata values (e.g. `done: <missing> => false`) can be safely ignored as mock operations are generally simpler.

### Important:

* It is important to commit the files modified by running realgcp tests in its own commit.
* This is for the human reviewer to compare the diff in the test artifacts when running real and mockgcp.

