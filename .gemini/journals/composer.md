# Cloud Composer Journal

### [2026-10-29] ComposerUserWorkloadsConfigMap ParentRef and Identity Modelling
- **Context**: Triaging judgement queue and finalizing KRM types and identity for `ComposerUserWorkloadsConfigMap` (#14617).
- **Problem**: `ComposerUserWorkloadsConfigMap` is a child resource of `ComposerEnvironment` (`projects/{project}/locations/{location}/environments/{environment}/userWorkloadsConfigMaps/{user_workloads_config_map}`). Bulk generation had guessed separate `projectRef` and `location` fields on `ComposerUserWorkloadsConfigMapSpec`, resulting in open findings for `identity-root-unknown`, `parent-ref-not-modelled`, and `location-or-parent-ref`.
- **Solution**: Modelled parent hierarchy using a single `ParentRef *v1beta1.EnvironmentRef` on `ComposerUserWorkloadsConfigMapSpec` per the single-parent rule, removing `projectRef` and `location`. Re-ran `apis/composer/generate.sh` which regenerated `getIdentityFromComposerUserWorkloadsConfigMapSpec` to normalize `spec.parentRef` and construct the identity using the parent's external reference and `resourceID`.
- **Impact**: Cleanly resolves parent and root findings while adhering to KCC's single-parent rule.
