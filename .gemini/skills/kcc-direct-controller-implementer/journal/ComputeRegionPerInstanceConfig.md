# ComputeRegionPerInstanceConfig Implementation Journal

## Overview
Implemented direct controller for `ComputeRegionPerInstanceConfig` migrating from legacy Terraform-based reconciler.

## Key Observations and Decisions
1. **Subresource Nature on Regional Instance Group Manager (IGM):**
   - `ComputeRegionPerInstanceConfig` manages per-instance configuration for an instance within a regional IGM.
   - In GCP Compute Engine API, there is no standalone `Get` API for per-instance configs; `Find` uses `RegionInstanceGroupManagers.ListPerInstanceConfigs` and iterates through items matching on `item.GetName() == id.Name`.
   - `Create` invokes `RegionInstanceGroupManagers.CreateInstances` with the `PerInstanceConfig` specification.
   - `Update` compares desired and actual preserved state using `common.DiffForTopLevelFields`. When changes are present, it invokes `RegionInstanceGroupManagers.UpdatePerInstanceConfigs`, looks up the managed instance's full URL via `RegionInstanceGroupManagers.ListManagedInstances`, and triggers `RegionInstanceGroupManagers.ApplyUpdatesToInstances` with the configured or default `minimalAction` (`NONE`) and `mostDisruptiveAllowedAction` (`REPLACE`).
   - `Delete` calls `RegionInstanceGroupManagers.DeletePerInstanceConfigs` and optionally triggers `ApplyUpdatesToInstances` if `removeInstanceStateOnDestroy` is true.

2. **Brownfield Migration Strategy:**
   - Kept `DefaultController` as `k8s.ReconcilerTypeTerraform` in `pkg/controller/resourceconfig/static_config.go` and added `k8s.ReconcilerTypeDirect` to `SupportedControllers`.
   - Verified that both legacy and direct controllers run in `TestAllInSeries` against MockGCP, generating expected golden diffs (`_http_mock.diff`, `_final_object.diff`, `_http_old_controller_mock.log`, `_final_object_old_controller.golden.yaml`).
