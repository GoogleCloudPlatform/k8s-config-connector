# ApiHub Service Journal

### 2026-05-24 Initial Scaffolding and Identity for APIHubAPI
- **Context**: Greenfield implementation of APIHubAPI types, CRD, and IdentityV2 under `apihub.cnrm.cloud.google.com/v1alpha1`.
- **Problem**: Scarce local build cache for `googleapis.pb` caused `generate-types` to fail.
- **Solution**: Executed `./dev/tools/controllerbuilder/generate-proto.sh` to compile googleapis proto definitions into `.build/googleapis.pb` using the pinned Google API SHA in `apis/git.versions`.
- **Impact**: Provides clean scaffolding, deepcopy, CRD, and client generation workflow. Also maps system defined attribute references to KRM `APIHubAttributeValueRef`.

### 2026-06-05 Scaffolding and Identity for ApiHubRuntimeProjectAttachment
- **Context**: Greenfield implementation of ApiHubRuntimeProjectAttachment types, CRD, and IdentityV2 under `apihub.cnrm.cloud.google.com/v1alpha1`.
- **Problem**: `gcpurls.Template` failed with a runtime panic during initialization because the snake_case template variable `{runtime_project_attachment}` did not match standard Go field casing mapping when defined as camelCase `RuntimeProjectAttachment string`.
- **Solution**: Defined the Identity Go struct field name with underscores: `Runtime_project_attachment string`. This correctly matches the `gcpurls.Template` parser's expectations.
- **Impact**: Other direct resource types with snake_case variables in their GCP URL template formats must follow this pattern to avoid initialization panics.

### 2026-06-05 Initial Scaffolding and Identity for ApiHubPlugin
- **Context**: Greenfield implementation of ApiHubPlugin types, CRD, and IdentityV2 under `apihub.cnrm.cloud.google.com/v1alpha1`.
- **Problem**: Need to manually define KRM Spec/ObservedState fields in `plugin_types.go` that align with proto-defined `Plugin` structures.
- **Solution**: Hand-coded `ApiHubPluginSpec` with `DisplayName` (Required, `*string`), `Type` (Required, `*AttributeValues`), and `Description` (Optional, `*string`). Implemented `ApiHubPluginIdentity` utilizing `gcpurls.Template` for the path pattern `projects/{project}/locations/{location}/plugins/{plugin}` and wrote unit tests for `FromExternal`.
- **Impact**: Provides valid schema and identity parsing matching GCP's ApiHub Plugin resource specifications.

### [2026-05-19] APIHubAttribute implementation
- **Context**: Implementing APIHubAttribute (direct).
- **Problem**: 
  1. `generate-fuzzer` attempts to run a `gemini-cli prompt` tool which may fail in certain environments.
  2. `fuzz-roundtrippers` fails if Spec-only fields are lost during `ObservedState` roundtripping.
- **Solution**: 
  1. Wrote the fuzzer manually utilizing `fuzztesting.NewKRMTypedFuzzer`.
  2. Added Spec-only fields to `f.SpecFields` (e.g. `f.SpecFields.Insert(".display_name")`) so that the fuzzer understands they will not roundtrip from ObservedState.
- **Impact**: Agents scaffolding direct resources for ApiHub (or similar) should manually scaffold fuzzers if the generation script fails, and ensure they properly register Spec and Status fields to pass `fuzz-roundtrippers`.
