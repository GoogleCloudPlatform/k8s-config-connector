# ComputeRegionPerInstanceConfig MockGCP Alignment Journal

## Overview
This document records observations and learnings from implementing MockGCP support and aligning real GCP behavior for `ComputeRegionPerInstanceConfig`.

## Key Observations
1. **Regional IGM Constraint**:
   - GCP Compute Engine requires `updatePolicy.instanceRedistributionType: NONE` on a regional Instance Group Manager for `createInstances` to be invoked.
2. **Read-only POST Methods in Compute Engine**:
   - `listPerInstanceConfigs` is an HTTP POST method in GCP's Compute Engine REST API despite being a read-only list operation. During re-reconciliation tests in `tests/e2e/unified_test.go`, `/listPerInstanceConfigs` must be recognized as read-only in `case "POST"` alongside `/listNetworkEndpoints`.
3. **Dependency Filtering in Golden Alignment Test**:
   - In `pkg/test/resourcefixture/golden_alignment_test.go`, `getPlaceholdersForKind` must include `ComputeInstanceTemplate` (`${instanceTemplateID}`) and `ComputeInstanceGroupManager` (`${instanceGroupManagerID}`) so that `filterDependencyEvents` accurately identifies and ignores dependency HTTP events when comparing target CRUD requests.
4. **Per-Instance Config Mocking Pattern**:
   - Per-instance configs are stored under their parent instance group manager: `projects/{project}/regions/{region}/instanceGroupManagers/{igm}/perInstanceConfigs/{name}`.
   - Operations on per-instance configs return regional operations with `operationType` matching `compute.instanceGroupManagers.<action>` (e.g., `createInstances`, `updatePerInstanceConfigs`, `deletePerInstanceConfigs`, `applyUpdatesToInstances`).
