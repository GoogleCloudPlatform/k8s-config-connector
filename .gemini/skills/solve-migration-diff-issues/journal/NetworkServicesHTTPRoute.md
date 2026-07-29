# Migration Journal: NetworkServicesHTTPRoute

## Takeover Diff Issues Identified and Solved

During the direct migration validation for `NetworkServicesHTTPRoute`, we identified key takeover diff issues between the DCL/legacy controller's recorded GCP state and the Direct controller's mapped state:

1. **`gateways` and `meshes` URL vs. relative path discrepancy:**
   - **Symptom:** During direct takeover, the Direct controller detected differences for the `gateways` and `meshes` fields and attempted to send a PATCH request with an update mask including `gateways,meshes`.
   - **Root Cause:** The legacy DCL controller created gateways/meshes using fully-qualified GCP URLs (e.g., `https://networkservices.googleapis.com/v1alpha1/projects/${projectId}/locations/global/gateways/...`), which is what GCP returned in actual state. However, the Direct controller mapped KRM references to standard relative paths (e.g., `projects/${projectId}/locations/global/gateways/...`). Because the strings did not match exactly, a diff was detected.
   - **Solution:** Implemented normalization helpers (`normalizeGCPResourceName` and `normalizeHTTPRoute`) inside `networkserviceshttproute_controller.go` to trim scheme and domain prefixes, aligning both actual and desired gateway and mesh lists to standard relative paths before comparison.

2. **`serviceName` full selfLink vs. relative path discrepancy in `rules`:**
   - **Symptom:** The Direct controller reported a difference for the `rules` field, prompting an update mask including `rules` during takeover.
   - **Root Cause:** The `serviceRef` references inside `rules[].action.destinations` (and `requestMirrorPolicy.destination`) were normalized by KCC's reference resolver to full Compute BackendService selfLink URLs (e.g., `https://www.googleapis.com/compute/v1/projects/${projectId}/global/backendServices/...`). However, GCP's API returned the relative path name (e.g., `projects/${projectId}/global/backendServices/...`).
   - **Solution:** Enhanced the normalization helpers to trim standard compute domains (e.g., `www.googleapis.com/compute/v1/`, `www.googleapis.com/compute/beta/`, `compute.googleapis.com/compute/v1/`) and API prefixes. This ensures all `ServiceName` fields inside route action destinations and mirror policies are compared as standard relative paths.

With these changes, the `TestMigrationToDirect` E2E test runs cleanly with a perfect **0-write no-op takeover** on both mock and real GCP targets.
