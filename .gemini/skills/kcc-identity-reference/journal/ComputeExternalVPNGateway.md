# ComputeExternalVPNGateway Identity and Reference Journal

## Observations & Learnings

- **Global Scope Integration:** `ComputeExternalVPNGateway` is a global resource in GCP, meaning its URL pattern contains `/global/` and lacks a region/location parameter, e.g. `projects/{project}/global/externalVpnGateways/{externalVPNGateway}`. This maps beautifully to the `gcpurls.Template` structure, matching standard compute network global structures.
- **Match CAIS Integration:** The CAIS name format for `ExternalVpnGateway` was already present in `cloudassetinventory_names.jsonl` under `compute.googleapis.com/ExternalVpnGateway`. By matching the exact pattern `projects/{project}/global/externalVpnGateways/{externalVPNGateway}`, the golden CAIS registration test was automatically satisfied once we regenerated the golden identities metadata with `WRITE_GOLDEN_OUTPUT=1`.
- **Backward Compatibility Validation:** No schema changes were made. We preserved the status.selfLink fallback mechanism in the reference `Normalize` method, and in the identity's `GetIdentity` method, ensuring fully safe backward compatibility for TF-based reconciliation and drift detection.
