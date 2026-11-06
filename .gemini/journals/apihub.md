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

### 2026-09-01 Greenfield Implementation of APIHubExternalAPI Controller
- **Context**: Implement direct controller, E2E fixtures, and fuzzer for APIHubExternalAPI.
- **Problem**: E2E harness requires static registration of GroupKind controller type in the resourceconfig static map, or it fails with unconfigured controller config errors.
- **Solution**: Implemented `apihubexternalapi_controller.go` and `apihubexternalapi_fuzzer.go` under package `apihubexternalapi` and registered it in `register.go` and `static_config.go` maps.
- **Impact**: Provides full lifecycle reconciliation, validation, round-trip fuzzing, and live GCP golden recording/verification for `APIHubExternalAPI`.

### 2026-11-05 Greenfield Types and Identity Finalization for APIHubPluginInstance
- **Context**: Triaged `judgement_queue.yaml` and finalized KRM types and identity for APIHubPluginInstance.
- **Problem**: APIHubPluginInstance is a child resource under `projects/{project}/locations/{location}/plugins/{plugin}/instances/{instance}`. Initial scaffolding had separate `projectRef` and `location`, and lacked `APIHubPluginRef`.
- **Solution**: Generated `APIHubPluginRef` and added `APIHubPlugin:Plugin` to `apis/apihub/generate.sh`. Modeled parent as required immutable `spec.parentRef` (`*APIHubPluginRef`). Triaged all `judgement_queue.yaml` entries.
- **Impact**: Clean identity resolution conforming to `IdentityV2` and `refs.Ref`, passing all apichecks and unit tests.

### 2026-11-06 Greenfield Direct Controller Implementation for APIHubPluginInstance
- **Context**: Implemented direct controller, fuzzer, and E2E fixtures for `APIHubPluginInstance`.
- **Problem**: GCP API Hub returns 500 when calling `GetPlugin` on non-existent custom plugin IDs, and requires `pluginCategory` and `actionsConfig` when creating custom user-owned `APIHubPlugin` resources. Furthermore, `UpdatePluginInstance` rejects `actions` in `update_mask`.
- **Solution**: Handled 500 in `APIHubPlugin` controller `Find` as not-found, defaulted required fields (`PluginCategory` and `ActionsConfig`) on `APIHubPlugin` creation, and filtered out `actions` from `updateMask` in `APIHubPluginInstance` `Update`.
- **Impact**: All E2E lifecycle stages (create, update, re-reconcile, delete) pass cleanly against real GCP for both minimal and maximal fixtures.



