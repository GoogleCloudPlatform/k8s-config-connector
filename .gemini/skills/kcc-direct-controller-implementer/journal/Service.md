# Service Direct Controller Implementation Journal

### 2026-06-06 Service Direct Reconciler Migration
- **Context**: Transitioning the `Service` (ServiceUsage) resource from the Terraform reconciler to the modern, direct reconciler.
- **Findings**:
  - The ServiceUsage Service represents an enabled service in GCP. When disabled, the state returns `DISABLED` (as observed in `mockserviceusage/serviceusagev1.go`).
  - There are no Spec properties besides `ProjectRef` and `ResourceID` on the KRM schema for `Service`.
  - For creation/enabling, we call `EnableService`, and for deletion/disabling we call `DisableService`. Both return a Long-Running Operation (LRO) which we wait on using `lro.Wait(ctx)`.
  - Because `Service` has no other Spec fields, the `Update` operation is a no-op, as any required enabled-state matching is already handled by `Find`.
  - We ran `dev/tasks/generate_static_config.py` which populated the direct reconciler override for `Service` in `pkg/controller/resourceconfig/static_config.go`.
  - We added `Service` to the `forceDirect = true` switch cases inside `tests/e2e/unified_test.go` to ensure both direct and legacy reconcilers are verified in E2E tests.
  - All E2E fixtures and fuzz-roundtrip tests passed perfectly under MockGCP.
