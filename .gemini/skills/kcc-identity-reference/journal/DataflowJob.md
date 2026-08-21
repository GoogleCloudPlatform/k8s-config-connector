# DataflowJob Identity and Reference Implementation Journal

### 2026-06-04 DataflowJob Identity and Refs Implementation
- **Context**: Transitioning `DataflowJob` to the `identity` and `refs` pattern.
- **Observation 1**: The GCP URL format for `DataflowJob` is `//dataflow.googleapis.com/projects/{{PROJECT_ID}}/locations/{{LOCATION}}/jobs/{{JOB}}`. This is registered in CAI under `docs/ai/metadata/cloudassetinventory_names.jsonl` on line 215, meaning we do not need to add any exception to `pkg/gcpurls/registry_test.go`.
- **Observation 2**: The path parameter for region/location in the GCP URL is `{location}`, but the KRM spec defines `spec.region` instead of `spec.location`. We solved this by extracting `obj.Spec.Region` and mapping it to the `Location` field in `DataflowJobIdentity`.
- **Observation 3**: `DataflowJobStatus` does not contain `externalRef` or `name` fields (only `jobId` which is not a full URL). Following the strict schema compatibility mandates, we did not modify the CRD status schema, and therefore did not implement status cross-checks in `GetIdentity`.
- **Outcome**: The implementation successfully compiles, all template registry unit tests pass, and the CRD schemas generated are 100% identical to the baseline.
