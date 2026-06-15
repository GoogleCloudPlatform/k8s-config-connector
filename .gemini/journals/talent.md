### [2026-06-15] CloudTalentSolutionCompany Direct Types Implementation
- **Context**: Implementing initial direct KRM types, CRD, and IdentityV2 for `CloudTalentSolutionCompany` (under `jobs.cnrm.cloud.google.com`).
- **Problem**: The issue requested the use of service `google.cloud.jobs.v1`, but `google.cloud.jobs.v1` did not exist in `googleapis.pb` or `proto-list-final.yaml`. Google Cloud Talent Solution API was renamed from Jobs API to Talent API and is located under `google.cloud.talent.v4`.
- **Solution**: We mapped the generator `--service` flag to `google.cloud.talent.v4` and the `+kcc:proto` annotation to `google.cloud.talent.v4` which successfully resolved the GCP `Company` proto and successfully scaffolded the Go types.
- **Impact**: The service mapping needs to point to `google.cloud.talent.v4` for any subsequent Cloud Talent Solution resource implementations.
