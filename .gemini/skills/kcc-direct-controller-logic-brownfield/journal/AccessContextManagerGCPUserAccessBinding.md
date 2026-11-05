# AccessContextManagerGCPUserAccessBinding Migration Journal

## Overview
Migrated `AccessContextManagerGCPUserAccessBinding` from the legacy Terraform-based reconciler to the `direct` reconciler architecture.

## Key Observations and Implementation Insights

### 1. Service-Generated Resource ID Semantics
- `AccessContextManagerGCPUserAccessBinding` uses service-generated identifiers assigned by GCP during resource creation (`organizations/{org}/gcpUserAccessBindings/{bindingID}`).
- In accordance with the `kcc-direct-service-generated-id` rules:
  - In `Find()`, if `a.id.GcpUserAccessBinding == ""`, the method immediately returns `(false, nil)` without issuing a premature GET call to GCP.
  - In `Create()`, `desiredpb.Name` is cleared (`desiredpb.Name = ""`) before sending the `CreateGcpUserAccessBinding` RPC so the GCP service generates the ID.
  - After creation, the binding ID is parsed from `created.GetName()` and populated into `spec.resourceID` via `createOp.SetSpecResourceID` when empty, and recorded in `status.name`.
  - In `Delete()`, if `a.id.GcpUserAccessBinding == ""`, it returns `(true, nil)` immediately.

### 2. Reconciliation and Update Semantics
- The Access Context Manager GCP User Access Binding API only permits mutable updates to `access_levels` via `UpdateGcpUserAccessBinding` with `update_mask { paths: "access_levels" }`. `group_key` and `name` are immutable.
- Diff detection uses `common.CompareBrownfieldSpec`, ensuring zero-write steady-state re-reconciliation.
- If a diff on immutable `groupKey` is attempted, `Update()` returns a clear validation error.

### 3. Client Implementation
- Utilized GAPIC REST client constructor `accesscontextmanager.NewRESTClient` initialized with `m.config.RESTClientOptions()`.
- API operations return long-running operations (`*accesscontextmanager.CreateGcpUserAccessBindingOperation`, `*accesscontextmanager.UpdateGcpUserAccessBindingOperation`, `*accesscontextmanager.DeleteGcpUserAccessBindingOperation`), which are cleanly waited on using `op.Wait(ctx)`.

### 4. Golden Fixture and Static Config Alignment
- `SupportedControllers` in `pkg/controller/resourceconfig/static_config.go` was updated to include `k8s.ReconcilerTypeDirect` while keeping `DefaultController: k8s.ReconcilerTypeTerraform`.
- Golden comparison files (`_final_object.diff`, `_final_object_old_controller.golden.yaml`, `_http_old_controller_mock.log`) confirm parity between direct and legacy reconcilers.
