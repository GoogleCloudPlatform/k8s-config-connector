# AppEngineFlexibleAppVersion Migration Journal

## Overview
Migrated `AppEngineFlexibleAppVersion` from the legacy Terraform-based reconciler to the `direct` reconciler architecture.

## Key Observations and Decisions
1. **GAPIC REST Client and Long-Running Operations (LRO):**
   - The direct controller uses GAPIC's official `api.NewVersionsRESTClient` (`cloud.google.com/go/appengine/apiv1`) and `api.NewServicesRESTClient`.
   - `CreateVersion`, `UpdateVersion`, and `DeleteVersion` return long-running operations (`*api.CreateVersionOperation`, `*api.UpdateVersionOperation`, `*api.DeleteVersionOperation`).
   - For App Engine versions, operations complete with empty response payloads or nil response types upon success. GAPIC's `Wait(ctx)` may return `"unsupported result type <nil>: <nil>"`, which is safely handled as a successful completion matching patterns elsewhere in direct controllers.

2. **Creation vs. Update Semantics:**
   - In App Engine Admin API, flexible app version creation and update both use `CreateVersion` (HTTP POST) with the version ID specified in `Version.Id` and `Env` set to `"flex"`.
   - In `Update`, `common.CompareBrownfieldSpec` compares desired KRM spec against the fetched `actual` `pb.Version`, detecting modifications while respecting server-defaulted fields.

3. **Deletion Semantics:**
   - Supported `spec.noopOnDestroy`: When set to true, version deletion in GCP is skipped.
   - Supported `spec.deleteServiceOnDestroy`: When set to true, the parent App Engine service (`apps/{app}/services/{service}`) is deleted via `servicesClient.DeleteService`.

4. **Brownfield Dual-Testing Alignment:**
   - `pkg/controller/resourceconfig/static_config.go` was updated to add `k8s.ReconcilerTypeDirect` to `SupportedControllers` for `AppEngineFlexibleAppVersion` while maintaining `DefaultController: k8s.ReconcilerTypeTerraform`.
   - Golden comparison files (`_http_mock.diff`, `_final_object.diff`, `_http_old_controller_mock.log`, `_final_object_old_controller.golden.yaml`) verify behavioral equivalence between the direct and legacy controllers.
   - All tests pass hermetically against MockGCP.
