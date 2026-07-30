# Journal for ComputeAutoscaler Direct Takeover Migration

## 2026-07-30 - ComputeAutoscaler Direct Takeover Diffs
- **Issue:** Under direct takeover, the direct controller saw diffs on `target` and `zone` fields.
- **Root Cause:**
  - `actual.Zone` returned by GCE was a fully qualified URL (e.g., `https://www.googleapis.com/compute/v1/projects/.../zones/us-central1-a`), whereas KRM `desired.Zone` was a short name (e.g., `us-central1-a`).
  - `actual.Target` returned by GCE was a fully qualified URL, whereas KRM `desired.Target` was a normalized relative path (e.g., `projects/...`).
- **Fix:**
  - In `compareComputeAutoscaler`, we normalized both `maskedActual` and `clonedDesired` within `populateDefaults`:
    - `*obj.Zone = lastComponent(*obj.Zone)` (where `lastComponent` extracts the portion after the last slash).
    - `*obj.Target = refs.TrimComputeURIPrefix(*obj.Target)` (using the standard GCE URI prefix trimmer).
- **Observations:**
  - `mockgcp` does not currently implement `autoscalers` endpoints. For this reason, the migration test must run against real GCP to record the correct cassettes and golden outputs, as mockgcp runs result in 404s.
