# DocumentAIProcessorDefaultVersion Journal

## Observations & Learnings

### 1. Resource and Proto Mapping
- `DocumentAIProcessorDefaultVersion` is a brownfield resource originally managed by the Terraform provider (`google_document_ai_processor_default_version`).
- In GCP's Document AI API (`google.cloud.documentai.v1`), setting the default processor version is an RPC (`SetDefaultProcessorVersion`) rather than a standalone REST resource. The request message is `SetDefaultProcessorVersionRequest`.
- In `apis/documentai/generate.sh`, we added `--resource DocumentAIProcessorDefaultVersion:SetDefaultProcessorVersionRequest` to `generate-types` under `v1alpha1`.

### 2. Strict Schema Compatibility
- The baseline CRD `documentaiprocessordefaultversions.documentai.cnrm.cloud.google.com` is `v1alpha1` and contains only:
  - `spec.resourceID` (*string, optional)
  - `spec.version` (string, required)
  - `status.conditions` ([]v1alpha1.Condition, optional)
  - `status.observedGeneration` (*int64, optional)
- To avoid breaking schema compatibility, direct-controller standard fields such as `projectRef`, `location`, `observedState`, and `externalRef` were omitted because they were not present in the baseline CRD.
- Top-level doc comments on `DocumentAIProcessorDefaultVersionSpec` and `DocumentAIProcessorDefaultVersionStatus` were omitted to match the baseline CRD, which did not contain top-level `spec` and `status` descriptions. This prevents `_no_description.spec` and `_no_description.status` diffs in `diff-crds`.
- Required kubebuilder labels (`managed-by-kcc=true`, `stability-level=alpha`, `system=true`, `tf2crd=true`) were preserved.
- Running `dev/tasks/diff-crds --base master` verified zero schema differences.
