### [2026-05-12] VPCSCConfig requires manual Enum field addition in KRM Spec
- **Context**: Implementing direct controller for `ArtifactRegistryVPCSCConfig:VPCSCConfig`.
- **Problem**: The `VPCSCConfig` proto contains only one configurable field `VpcscPolicy`, which is an enum. `generate-types` skips adding it to the KRM `Spec` struct in `vpcscconfig_types.go`.
- **Solution**: Manually added `VpcscPolicy *string` to the `ArtifactRegistryVPCSCConfigSpec` struct and ran `generate-mapper` again. The mapper generator then correctly identified it and generated `direct.Enum_FromProto` / `direct.Enum_ToProto` logic in `mapper.generated.go`.
- **Impact**: When scaffolding new resources with enum-only fields or where `generate-types` skips enums, developers must manually declare them in `_types.go` before generating mappers.
