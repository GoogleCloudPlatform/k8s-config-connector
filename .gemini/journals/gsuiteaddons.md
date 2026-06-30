# GSuite Add-ons Journal

### [2026-06-29] Implementing Direct KRM Types for GSuiteAddonsDeployment
- **Context**: Scaffolded and implemented initial KRM types, CRD, and IdentityV2 for `GSuiteAddonsDeployment` under `apis/gsuiteaddons/v1alpha1`. (Issue #10276)
- **Problem**: 
  1. The auto-scaffolded spec initially had a `Location` field, but `GSuiteAddonsDeployment` is a project-scoped resource with GCP URL template `projects/{project}/deployments/{deployment}`, meaning no location field is required in the Spec.
  2. Nested structures like `AddOns` and `OauthScopes` are labeled as "unreachable" and commented out by the generator in `types.generated.go` on initial run.
  3. When nested structs reference `apiextensionsv1.JSON`, compilation fails because `apiextensionsv1` is not imported or used.
- **Solution**: 
  1. Removed `Location` from the Spec, as it's a project-scoped resource.
  2. Defined the mutable spec fields (`OauthScopes` and `AddOns`) directly in `GSuiteAddonsDeploymentSpec` inside `gsuiteaddonsdeployment_types.go` and ran the generator again. Rerunning `generate.sh` successfully detected these as reachable and generated active, uncommented Go types in `types.generated.go`.
  3. Imported `apiextensionsv1` in `gsuiteaddonsdeployment_types.go` and defined the dummy variable `var _ = apiextensionsv1.JSON{}` to allow compiling the package standalone.
- **Impact**: Any future developer implementing GSuite Add-ons controllers or mappers can rely on fully compliant, type-safe, and compilable KRM definitions and a correct project-scoped IdentityV2 implementation.
