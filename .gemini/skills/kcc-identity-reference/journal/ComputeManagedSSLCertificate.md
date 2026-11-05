# ComputeManagedSSLCertificate Identity and Reference Journal

## Observations & Learnings

- **Global Location Format:** `ComputeManagedSSLCertificate` represents Google-managed SSL certificates in GCP Compute Engine, which only exist globally at `projects/{project}/global/sslCertificates/{sslCertificate}` (unlike self-managed `ComputeSSLCertificate`, which can also be regional).
- **CAI Alignment:** The URL format `//compute.googleapis.com/projects/{project}/global/sslCertificates/{sslCertificate}` natively matches the entry in `docs/ai/metadata/cloudassetinventory_names.jsonl` under `compute.googleapis.com/SslCertificate` without requiring an exception in `pkg/gcpurls/registry_test.go`.
- **Status Cross-Check & Legacy Fallback:**
  - `ComputeManagedSSLCertificateStatus` contains `ExternalRef *string` as well as `ObservedState.SelfLink`.
  - In `GetIdentity`, the parsed spec identity is verified against `status.externalRef` or `status.observedState.selfLink` to detect drift.
  - In `Normalize`, `refs.NormalizeWithFallback` checks `status.observedState.selfLink` and `status.selfLink` to ensure the resource is created and ready in GCP before returning the resolved identity string, returning empty if neither status field is populated so that `k8s.NewReferenceNotReadyError` is returned for unready resources.
- **No Schema Changes:** The CRD schema remained completely untouched (`dev/tasks/diff-crds --base master` returned empty diff), maintaining strict backwards compatibility.
- **Verification:**
  - Unit tests in `computemanagedsslcertificate_identity_test.go` verified URL parsing, formatting, validation, `GetIdentity` drift detection, and `Normalize` fallback.
  - CAIS identity test `TestGoldenIdentitiesYamlFiles` passed and updated `_identities.yaml` for the `computemanagedsslcertificate` fixture.
  - E2E tests (`TestAllInSeries/fixtures/computemanagedsslcertificate`) passed with `mockgcp`.
