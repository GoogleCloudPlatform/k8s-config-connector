# Cloud Composer Journal

### [2026-10-29] ComposerUserWorkloadsConfigMap ParentRef and Identity Modelling
- **Context**: Triaging judgement queue and finalizing KRM types and identity for `ComposerUserWorkloadsConfigMap` (#14617).
- **Problem**: `ComposerUserWorkloadsConfigMap` is a child resource of `ComposerEnvironment` (`projects/{project}/locations/{location}/environments/{environment}/userWorkloadsConfigMaps/{user_workloads_config_map}`). Bulk generation had guessed separate `projectRef` and `location` fields on `ComposerUserWorkloadsConfigMapSpec`, resulting in open findings for `identity-root-unknown`, `parent-ref-not-modelled`, and `location-or-parent-ref`.
- **Solution**: Modelled parent hierarchy using a single `ParentRef *v1beta1.EnvironmentRef` on `ComposerUserWorkloadsConfigMapSpec` per the single-parent rule, removing `projectRef` and `location`. Re-ran `apis/composer/generate.sh` which regenerated `getIdentityFromComposerUserWorkloadsConfigMapSpec` to normalize `spec.parentRef` and construct the identity using the parent's external reference and `resourceID`.
- **Impact**: Cleanly resolves parent and root findings while adhering to KCC's single-parent rule.

### [2026-10-29] ComposerUserWorkloadsSecret judgement triage and parentRef hierarchy
- **Context**: Triaging judgement queue and finalizing direct KRM types and identity for `ComposerUserWorkloadsSecret` (Step 1b).
- **Problem**: `ComposerUserWorkloadsSecret` is a child resource under `ComposerEnvironment` (`projects/{project}/locations/{location}/environments/{environment}/userWorkloadsSecrets/{user_workloads_secret}`). Scaffolding initially emitted top-level `ProjectRef` and `Location` fields, which flagged `location-or-parent-ref`, `parent-ref-not-modelled`, and `identity-root-unknown`. Additionally, `CreateUserWorkloadsSecretRequest` has no `user_workloads_secret_id` parameter, triggering `identity-server-generated-id`.
- **Solution**: Replaced `ProjectRef` and `Location` with a single immutable `ParentRef *composerv1beta1.EnvironmentRef` following the single-parent rule. Ran `generate-identity` which generated parent normalization and resolved `identity-root-unknown`. Accepted `ServerGeneratedIdentity` as `CreateUserWorkloadsSecretRequest` has no `*_id` parameter. Resolved `empty-observedstate` by confirming proto has no output-only fields and keeping `ObservedState` commented out. Added `generate-controller` and `generate-fuzzer` to `apis/composer/generate.sh` to keep controller and fuzzer in sync.
- **Impact**: Pre-generated controller and fuzzer build cleanly and pass `TestFocusedMappers`. Subsequent steps (Step 2 controller implementation/testing and Step 3 MockGCP) can rely on standard `spec.parentRef` resolution.

### [2026-11-03] ComposerUserWorkloadsSecret direct controller implementation, masked secrets, and E2E recording
- **Context**: Implementing direct controller, fuzzer, and recording real GCP E2E fixtures for `ComposerUserWorkloadsSecret` (#14739).
- **Learnings**:
  - **Client-Specified Resource Name on Create**: Cloud Composer's `CreateUserWorkloadsSecretRequest` has no query parameter for `secret_id`; instead, the full resource name (`projects/{project}/locations/{location}/environments/{environment}/userWorkloadsSecrets/{secret}`) must be set on `req.UserWorkloadsSecret.Name`. If `spec.resourceID` is omitted, the controller defaults the secret name to `metadata.name`.
  - **Masked Secret Values on GET**: GCP masks all values in `data` with `"**********"` on `GetUserWorkloadsSecret`. To prevent spurious diffs during re-reconciliation, `Update` uses a `normalize` function with `CompareGreenfieldSpec`: when `status.ObservedGeneration` matches `u.GetGeneration()` and `actual.Data` values are `"**********"`, they are normalized to the desired spec values.
  - **Cloud Composer 3 Dependency**: `userWorkloadsSecrets` is supported exclusively on Cloud Composer 3 (`composer-3-airflow-*`). E2E fixture `dependencies.yaml` provisions a standard Composer 3 environment (`imageVersion: composer-3-airflow-2.11.1-build.22`) with `roles/composer.worker`.

### [2026-11-03] ComposerUserWorkloadsConfigMap Direct Controller and E2E Recording
- **Context**: Implementing Greenfield Step 2 direct controller, E2E fixtures, and fuzzer verification for `ComposerUserWorkloadsConfigMap` (#14710).
- **Findings & Solutions**:
  - `UserWorkloadsConfigMap` is supported on Cloud Composer 3 environments (`composer-3-*`). Test fixtures specify `dependencies.yaml` with `IAMServiceAccount`, `IAMPolicyMember` (`roles/composer.worker`), and `ComposerEnvironment` with `imageVersion: composer-3-airflow-2.11.1-build.22`.
  - In `AdapterForObject`, ensured `common.NormalizeGreenfieldReferences` is called before `GetIdentity`, and defaulted `id.UserWorkloadsConfigMap` to `obj.GetName()` if `spec.resourceID` was unset (per CRD spec description).
  - In `Create`, ensured `resource.Name = a.id.String()` is set when `a.id.HasIdentitySpecified()`, allowing client-specified configmap naming in the request body.
  - Recorded authentic golden HTTP logs against real GCP (`cnrm-barni-4`), passing all 7/7 automated validation checks for both minimal and maximal fixtures.

