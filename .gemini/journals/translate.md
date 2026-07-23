# Translate Service Journal

### 2026-07-23 TranslateAdaptiveMtDataset Greenfield Implementation
- **Context**: Greenfield implementation of the direct controller, fuzzer, and E2E fixtures for `TranslateAdaptiveMtDataset` (v1alpha1) under `pkg/controller/direct/translate`.
- **Problem**: The GCP Cloud Translation API (v3) has no `Update` or `Patch` method for `AdaptiveMtDataset` resources. This means the resource is entirely immutable from the GCP perspective.
- **Solution**:
  - Implemented the direct controller where the `Update` method checks for diffs using `compareTranslateAdaptiveMtDataset`. If any diff is detected in the spec, it returns a clear, descriptive error (`fmt.Errorf("TranslateAdaptiveMtDataset resource is immutable and cannot be updated. Field(s) changed: %v", diffs.FieldIDs())`) which surfaces on the KRM status condition.
  - Resolved `a.desired.Name` by manually assigning the full resource name (`a.id.String()`) in `AdapterForObject` before return. This ensures the GCP dataset identifier matches KRM specifications.
  - Successfully added the GVK to the static config registry mapping and imported the translate package in the direct registration package.
  - Created minimal and maximal test fixtures and updated the alpha exceptions list of missing fields.
- **Impact**: Demonstrates a clean pattern for implementing completely immutable resources under KCC's direct controller architecture.
