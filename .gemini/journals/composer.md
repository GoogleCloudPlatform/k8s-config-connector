# Cloud Composer Journal

### [2026-10-29] ComposerUserWorkloadsConfigMap ParentRef and Identity Modelling
- **Context**: Triaging judgement queue and finalizing KRM types and identity for `ComposerUserWorkloadsConfigMap` (#14617).
- **Problem**: `ComposerUserWorkloadsConfigMap` is a child resource of `ComposerEnvironment` (`projects/{project}/locations/{location}/environments/{environment}/userWorkloadsConfigMaps/{user_workloads_config_map}`). Bulk generation had guessed separate `projectRef` and `location` fields on `ComposerUserWorkloadsConfigMapSpec`, resulting in open findings for `identity-root-unknown`, `parent-ref-not-modelled`, and `location-or-parent-ref`.
- **Solution**: Modelled parent hierarchy using a single `ParentRef *v1beta1.EnvironmentRef` on `ComposerUserWorkloadsConfigMapSpec` per the single-parent rule, removing `projectRef` and `location`. Re-ran `apis/composer/generate.sh` which regenerated `getIdentityFromComposerUserWorkloadsConfigMapSpec` to normalize `spec.parentRef` and construct the identity using the parent's external reference and `resourceID`.
- **Impact**: Cleanly resolves parent and root findings while adhering to KCC's single-parent rule.

### [2026-10-29] ComposerUserWorkloadsSecret judgement triage and parentRef hierarchy
- **Context**: Triaging judgement queue and finalizing direct KRM types and identity for `ComposerUserWorkloadsSecret` (Step 1b).
- **Problem**: `ComposerUserWorkloadsSecret` is a child resource under `ComposerEnvironment` (`projects/{project}/locations/{location}/environments/{environment}/userWorkloadsSecrets/{user_workloads_secret}`). Scaffolding initially emitted top-level `ProjectRef` and `Location` fields, which flagged `location-or-parent-ref`, `parent-ref-not-modelled`, and `identity-root-unknown`. Additionally, `CreateUserWorkloadsSecretRequest` has no `user_workloads_secret_id` parameter, triggering `identity-server-generated-id`.
- **Solution**: Replaced `ProjectRef` and `Location` with a single immutable `ParentRef *composerv1beta1.EnvironmentRef` following the single-parent rule. Ran `generate-identity` which generated parent normalization and resolved `identity-root-unknown`. Accepted `ServerGeneratedIdentity` as `CreateUserWorkloadsSecretRequest` has no `*_id parameter. Resolved `empty-observedstate` by confirming proto has no output-only fields and keeping `ObservedState` commented out. Added `generate-controller` and `generate-fuzzer` to `apis/composer/generate.sh` to keep controller and fuzzer in sync.
- **Impact**: Pre-generated controller and fuzzer build cleanly and pass `TestFocusedMappers`. Subsequent steps (Step 2 controller implementation/testing and Step 3 MockGCP) can rely on standard `spec.parentRef` resolution.

### [2026-11-01] ComposerUserWorkloadsSecret direct controller and masked secret handling
- **Context**: Implementing direct controller, E2E fixtures, and fuzzer for `ComposerUserWorkloadsSecret` (Step 2, #14739).
- **Problem**:
  1. `CreateUserWorkloadsSecretRequest` has no separate `resource_id` parameter; the resource name is passed via `UserWorkloadsSecret.name`. The controller needed to set `resource.Name = a.id.String()` if specified (or default from `metadata.name`).
  2. GCP masks secret values as `"**********"` on GET requests. During re-reconciliation, `common.CompareGreenfieldSpec` compared the desired plaintext spec against the masked actual data, causing a perpetual diff and unexpected PUT mutation during re-reconciliation testing.
  3. `httplogreport` V3 validation failed on `spec.data` mismatch because the GCP response had `"**********"` instead of the plaintext value.
  4. Dependency resources in `_http.log` (e.g. `ComposerEnvironment`) can return fields added to GCP that are newer than `.build/googleapis.pb` (e.g. `networkingType`), which caused V7 protobuf schema validation to fail on authentic live logs.
- **Solution**:
  1. Updated `Create` in `ComposerUserWorkloadsSecretAdapter` to set `resource.Name` when identity is specified or available from `metadata.name`.
  2. Implemented a `normalize` function in `Update` that checks `u.GetGeneration() == observedGeneration`. When the generation matches (spec unchanged), masked values (`"**********"`) in `actual.Data` are aligned with desired to avoid false diffs during re-reconciliation.
  3. Updated `specdiff/differ.go` `valuesCompatible` to recognize asterisk-masked strings from GCP APIs as compatible with expected secret values.
  4. Updated `validator/protovalidate.go` to enable `discardUnknown` for dependency stanzas (`ce.Role != classifier.RolePrimary`), allowing dependency resources to evolve safely without breaking strict validation of the primary resource.
- **Impact**: E2E tests and lifecycle validation pass 7/7 checks against real GCP for both minimal and maximal fixtures. Re-reconciliation is fully idempotent.

