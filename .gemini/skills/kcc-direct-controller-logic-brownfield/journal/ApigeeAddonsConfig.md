# ApigeeAddonsConfig Migration Journal

## Overview
Migrated `ApigeeAddonsConfig` from the legacy Terraform-based reconciler to the `direct` reconciler architecture.

## Key Observations and Decisions
1. **API Client & Endpoints:**
   - Uses `google.golang.org/api/apigee/v1` (`OrganizationsService`).
   - Apigee addons configuration is managed at the organization level via `GET organizations/{org}` and `POST organizations/{org}:setAddons`.
   - `setAddons` replaces the add-on configuration and returns a `GoogleLongrunningOperation`.
   - Used `WaitForApigeeOp` for long-running operation polling.

2. **Lifecycle & Reconciliation Mechanics:**
   - In GCP, `addonsConfig` is a configuration block on the organization rather than an independent resource with dedicated create/delete endpoints.
   - In `Find`: retrieves the organization via `organizationsClient.Get(a.id.ParentString())`. If the organization exists, `Find` succeeds and populates `actual` with `org.AddonsConfig`.
   - In `Create` / `Update`: maps `desired.Spec.AddonsConfig` to `api.GoogleCloudApigeeV1AddonsConfig` and calls `setAddons`. `Update` calculates field diffs across the supported add-on configurations (`AdvancedApiOpsConfig`, `IntegrationConfig`, `MonetizationConfig`, `ApiSecurityConfig`, `ConnectorsPlatformConfig`) and skips the API call when no diffs exist (achieving 0-write steady-state re-reconciliation).
   - In `Delete`: sends an empty `SetAddons` request (`{}`) to clear add-on configurations on the organization.

3. **Status & Export:**
   - `ApigeeAddonsConfigStatus` contains only standard conditions and observedGeneration; no additional GCP status fields are present.
   - Implemented `Export` and `AdapterForURL` adhering to `ApigeeAddonsConfigIdentityFormat`.

4. **Brownfield Dual-Testing Alignment:**
   - Updated `pkg/controller/resourceconfig/static_config.go` to add `k8s.ReconcilerTypeDirect` to `SupportedControllers` while retaining `DefaultController: k8s.ReconcilerTypeTerraform`.
   - Generated dual-controller golden files (`_final_object_old_controller.golden.yaml`, `_http_old_controller_mock.log`, `_http_mock.diff`, `_final_object.diff`).
   - Verified that both legacy and direct controllers pass mock testing (`TestAllInSeries`) and takeover/migration testing (`TestMigrationToDirect`).
