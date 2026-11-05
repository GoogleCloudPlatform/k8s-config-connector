# AppEngineServiceSplitTraffic Migration Journal

## Overview
Implemented the direct controller logic for `AppEngineServiceSplitTraffic` (brownfield resource migrating from Terraform).

## Key Observations and Implementation Insights

### 1. Sub-Aspect Semantics of Service Split Traffic
- `AppEngineServiceSplitTraffic` manages the `split` property of an AppEngine `Service` resource (`apps/{app}/services/{service}`).
- In `Find`, the controller retrieves the existing service via `GetService` and checks that the `split` section exists with traffic allocations (`actual.GetSplit() != nil && len(actual.GetSplit().GetAllocations()) > 0`). If present, `Find` returns `true, nil` and sets `a.actual` to the fetched `pb.Service`.
- Diff detection utilizes `common.CompareSpecifiedSpec`, which produces `diffRes.MergedDesired` to preserve server-defaulted subfields (e.g. `split.shardBy` when omitted in KRM) and avoid spurious re-reconciliations.
- `spec.migrateTraffic` is a mutable-but-unreadable (MBUR) field sent as a query parameter on `UpdateServiceRequest`. Its state is tracked using `createOp.SetMBUR`, `updateOp.CompareMBUR`, and `updateOp.SetMBUR` via the `cnrm.cloud.google.com/mutable-but-unreadable-fields-hash` annotation.
- `status.externalRef` is populated on `updateStatus()` and `Export()` with the resource's canonical GCP ID (`apps/{app}/services/{service}`).

### 2. Deletion Semantics
- In GCP, AppEngine services cannot be deleted solely via split traffic manipulation (deleting split traffic has no GCP API equivalent; services retain their traffic allocations).
- The legacy Terraform provider treated delete as removing the resource from state with a warning.
- In the direct controller, the model is registered using `registry.CannotBeDeleted()`, and `Delete` is a safe no-op returning `(true, nil)`.

### 3. Enum & Numeric Allocations Serialization
- The GAPIC REST client adds `%24alt=json%3Benum-encoding%3Dint` to query parameters and encodes enums (such as `TrafficSplit.ShardBy`: `IP` = 2, `COOKIE` = 1, `RANDOM` = 3) and allocations as numeric values in JSON payloads, whereas legacy Terraform serialized string values.
- `pkg/test/resourcefixture/golden_alignment_test.go` normalizes these representation differences so that mock golden alignment succeeds deterministically.

### 4. Controller Registration and Dual-Testing
- `pkg/controller/resourceconfig/static_config.go` was updated to include `k8s.ReconcilerTypeDirect` in `SupportedControllers` for `AppEngineServiceSplitTraffic` while keeping `DefaultController: k8s.ReconcilerTypeTerraform`.
- Golden comparison files (`_http_mock.diff`, `_final_object.diff`, `_http_old_controller_mock.log`, `_final_object_old_controller.golden.yaml`) verify behavioral equivalence between the direct and legacy controllers.
