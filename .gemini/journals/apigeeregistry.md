### [2026-05-26] ApigeeRegistryInstance Scaffolding
- **Context**: Implementing ApigeeRegistryInstance KRM types and IdentityV2.
- **Problem**: Need to replace `cmek_key_name` primitive with a reference in the spec.
- **Solution**: Defined `Instance_Config` in `instance_types.go` overriding the generated type, replacing `CmekKeyName *string` with `CmekKeyNameRef *refsv1beta1.KMSCryptoKeyRef`, successfully hiding the generated `Instance_Config` from `types.generated.go`. Also ensured `Location` is a pointer `*string` to pass the primitive pointer invariants.
- **Impact**: Demonstrates that replacing generated structures in `instance_types.go` properly prunes them from `types.generated.go` upon running `generate-types`.

### [2026-06-05] ApigeeRegistryArtifact Scaffolding and CAI Exceptions
- **Context**: Implementing Greenfield direct types and IdentityV2 for ApigeeRegistryArtifact (Issue #9253).
- **Problem**: 
  1. `generate-types` left the `Artifact` and `ArtifactObservedState` properties in `types.generated.go` as commented-out unreachable types because they were not initially referenced by `ApigeeRegistryArtifactSpec` or `ApigeeRegistryArtifactObservedState`.
  2. The Google API pattern `projects/{project}/locations/{location}/artifacts/{artifact}` (using host `apigeeregistry.googleapis.com`) is not in `docs/ai/metadata/cloudassetinventory_names.jsonl`, which causes `TestRegisteredTemplatesMatchCAI` to fail since it is not registered in CAI.
- **Solution**: 
  1. Manually copied/adapted the fields from `types.generated.go` into `artifact_types.go` under `ApigeeRegistryArtifactSpec` and `ApigeeRegistryArtifactObservedState`. Changed `Location string` to `Location *string` to satisfy the pointer scalar primitive review invariant.
  2. Added `"//apigeeregistry.googleapis.com/projects/{}/locations/{}/artifacts/{}"` to `ignoredTemplates` map inside `pkg/gcpurls/registry_test.go`.
- **Impact**: When adding new Greenfield resources, always check if they are absent from CAI (`cloudassetinventory_names.jsonl`) and add them to `ignoredTemplates` in `pkg/gcpurls/registry_test.go` to keep the registered templates checks green.
