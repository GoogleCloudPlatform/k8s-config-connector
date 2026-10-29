# ComputeRegionPerInstanceConfig Identity and Reference Journal

## Observations & Learnings

- **Sub-resource under Regional Instance Group Manager:** `ComputeRegionPerInstanceConfig` represents per-instance configuration for regional managed instance groups in Compute Engine. Its canonical GCP URL format follows `projects/{project}/regions/{region}/instanceGroupManagers/{regionInstanceGroupManager}/{name}`.
- **Parent Reference Handling:** In `ComputeRegionPerInstanceConfigSpec`, the parent IGM is referenced through `regionInstanceGroupManagerRef`. This field supports both external references (canonical URLs, full compute URIs, or short names) and Kubernetes object references by `name` (querying `ComputeInstanceGroupManager` via `client.Reader`).
- **Absence from Cloud Asset Inventory (CAI):** Like several other fine-grained Compute Engine sub-resources, `ComputeRegionPerInstanceConfig` is not represented in Cloud Asset Inventory. An exception was registered in `pkg/gcpurls/registry_test.go` under `ignoredTemplates` to satisfy `TestRegisteredTemplatesMatchCAI`.
- **Legacy Fallback Normalization:** As a legacy Terraform-managed resource, `ComputeRegionPerInstanceConfig` does not yet expose `status.externalRef`. `ComputeRegionPerInstanceConfigRef.Normalize` leverages `refs.NormalizeWithFallback`, checking that `status.conditions` indicates `Ready == True` prior to falling back to spec-based identity resolution.
- **Preservation of CRD Schema:** No OpenAPI / CRD schema fields were modified in `computeregionperinstanceconfig_types.go`, guaranteeing strict backward compatibility verified via `dev/tasks/diff-crds`.
