# ComputeAutoscaler Identity and Reference Journal

## Observations & Learnings

- **Zonal Resource Support with Custom Spec Field:** `ComputeAutoscaler` is a zonal resource that specifies its zone via `spec.zone` instead of a generic parent location or `spec.location` field. The identity parsing extraction resolves this dynamically from `obj.Spec.Zone`.
- **Backward Compatibility via status.selfLink:** Since `ComputeAutoscaler` is a legacy/Terraform-based resource, its status struct does not have an `externalRef` field but does contain `selfLink`. During normalization, `Normalize` falls back to reading `status.selfLink` as the authoritative resource identifier using `refs.TrimComputeURIPrefix`.
- **Dual-Field Templates in CAI:** The GCP `compute.googleapis.com/Autoscaler` definition on CAI supports both regional and zonal autoscalers, but KCC splits them into separate CRDs: `ComputeAutoscaler` (zonal) and `ComputeRegionAutoscaler` (regional). The `ComputeAutoscaler` identity specifically uses the zonal URL format: `projects/{project}/zones/{zone}/autoscalers/{autoscaler}`.
- **Strict Adherence to Non-breaking Schema Policies:** We did not alter any OpenAPI / CRD schema fields (such as adding `status.externalRef`), ensuring backward compatibility and schema stability.
