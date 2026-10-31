# Cloud Composer Journal

### [2026-10-29] ComposerUserWorkloadsConfigMap ParentRef and Identity Modelling
- **Context**: Triaging judgement queue and finalizing KRM types and identity for `ComposerUserWorkloadsConfigMap` (#14617).
- **Problem**: `ComposerUserWorkloadsConfigMap` is a child resource of `ComposerEnvironment` (`projects/{project}/locations/{location}/environments/{environment}/userWorkloadsConfigMaps/{user_workloads_config_map}`). Bulk generation had guessed separate `projectRef` and `location` fields on `ComposerUserWorkloadsConfigMapSpec`, resulting in open findings for `identity-root-unknown`, `parent-ref-not-modelled`, and `location-or-parent-ref`.
- **Solution**: Modelled parent hierarchy using a single `ParentRef *v1beta1.EnvironmentRef` on `ComposerUserWorkloadsConfigMapSpec` per the single-parent rule, removing `projectRef` and `location`. Re-ran `apis/composer/generate.sh` which regenerated `getIdentityFromComposerUserWorkloadsConfigMapSpec` to normalize `spec.parentRef` and construct the identity using the parent's external reference and `resourceID`.
- **Impact**: Cleanly resolves parent and root findings while adhering to KCC's single-parent rule.

### [2026-10-31] ComposerUserWorkloadsConfigMap Direct Controller and ID Specification
- **Context**: Implementing Greenfield direct controller and E2E fixtures for ComposerUserWorkloadsConfigMap (#14710).
- **Problem**: CreateUserWorkloadsConfigMapRequest does not have a separate *_id parameter; instead, the resource name in the body (UserWorkloadsConfigMap.name) must be specified in the format projects/{project}/locations/{location}/environments/{environment}/userWorkloadsConfigMaps/{userWorkloadsConfigMapId}. Omitting name in the request body results in INVALID_ARGUMENT (400). Therefore, the resource ID is client-specified rather than server-generated.
- **Solution**: Set resource.Name = a.id.String() in Create before invoking CreateUserWorkloadsConfigMap. Updated ComposerUserWorkloadsConfigMapIdentity to default to metadata.name when spec.resourceID is empty (not server-generated), and updated judgement_queue.yaml accordingly.
- **Impact**: Ensures config maps are created with their full resource name matching a.id.String(), preventing 400 errors from the GCP API.
