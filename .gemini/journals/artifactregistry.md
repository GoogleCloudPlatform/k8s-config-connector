### [2026-05-12] VPCSCConfig requires manual Enum field addition in KRM Spec
- **Context**: Implementing direct controller for `ArtifactRegistryVPCSCConfig:VPCSCConfig`.
- **Problem**: The `VPCSCConfig` proto contains only one configurable field `VpcscPolicy`, which is an enum. `generate-types` skips adding it to the KRM `Spec` struct in `vpcscconfig_types.go`.
- **Solution**: Manually added `VpcscPolicy *string` to the `ArtifactRegistryVPCSCConfigSpec` struct and ran `generate-mapper` again. The mapper generator then correctly identified it and generated `direct.Enum_FromProto` / `direct.Enum_ToProto` logic in `mapper.generated.go`.
- **Impact**: When scaffolding new resources with enum-only fields or where `generate-types` skips enums, developers must manually declare them in `_types.go` before generating mappers.

### [2026-10-27] ArtifactRegistryRule repository rules and CEL expression support
- **Context**: Implementing direct controller, E2E fixtures, and fuzzer for `ArtifactRegistryRule:Rule` (`artifactregistry.cnrm.cloud.google.com/v1alpha1`).
- **Problem**: `Rule` resources represent repository-level or package-level download policies under an `ArtifactRegistryRepository`. The resource schema uses nested `google.type.Expr` for conditions along with enum fields (`action`, `operation`) and packageID.
- **Solution**: The rule condition CEL expression, title, description, location, action, operation, and packageID are fully supported and round-tripped by GCP Artifact Registry v1 API without field loss on create or update. The controller operates via unary RPCs (`CreateRule`, `UpdateRule`, `DeleteRule`) rather than LROs.
- **Impact**: Subsequent steps (MockGCP implementation and promotion) can reference the recorded golden fixtures covering both minimal and maximal configurations with all CEL condition fields.

### [2026-10-29] MockGCP ArtifactRegistryRule implementation and test harness registration
- **Context**: Implementing MockGCP support for `ArtifactRegistryRule` (Step 3).
- **Problem**: `ArtifactRegistryRule` uses unary RPCs (`CreateRule`, `GetRule`, `UpdateRule`, `DeleteRule`, `ListRules`) on `ArtifactRegistryServer`. When running E2E mock tests, `ArtifactRegistryRule` was initially skipped by `harness.go` because its GroupKind was not registered in `MaybeSkip`.
- **Solution**: Implemented unary CRUD methods in `mockartifactregistry/artifactregistryrule.go`, supported `ruleName` parsing in `mockartifactregistry/names.go`, cleaned up child rules on `DeleteRepository`, and registered `schema.GroupKind{Group: "artifactregistry.cnrm.cloud.google.com", Kind: "ArtifactRegistryRule"}` in `config/tests/samples/create/harness.go`.
- **Impact**: Fixture tests for `ArtifactRegistryRule` can now run hermetically in MockGCP during CI presubmit.


