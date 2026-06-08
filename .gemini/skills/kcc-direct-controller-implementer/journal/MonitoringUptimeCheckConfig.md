# MonitoringUptimeCheckConfig Direct Controller Implementation Journal

### 2026-06-08 MonitoringUptimeCheckConfig Direct Reconciler Migration
- **Context**: Transitioning `MonitoringUptimeCheckConfig` from the DCL reconciler to the modern, direct reconciler.
- **Findings**:
  - The `UptimeCheckService` (`GetUptimeCheckConfig`, `CreateUptimeCheckConfig`, `UpdateUptimeCheckConfig`, and `DeleteUptimeCheckConfig`) operations are synchronous and do not return Long Running Operations (LRO).
  - The standard Go client for Cloud Monitoring Uptime Checks (`cloud.google.com/go/monitoring/apiv3/v2`) uses gRPC under the hood. Therefore, we used `GRPCClientOptions()` instead of `RESTClientOptions()` to construct `NewUptimeCheckClient` in `client.go`, preventing dial option conflicts such as `WithHTTPClient is incompatible with gRPC dial options`.
  - The Spec's `ProjectRef` field is defined using the root `refs` package (`"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs".ProjectRef`), while the E2E helper `refs.ResolveProject` expects `refsv1beta1.ProjectRef` (`"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1".ProjectRef`). We successfully mapped/converted this reference before calling the resolver.
  - The E2E tests under MockGCP passed cleanly and we regenerated the golden logs and objects using `WRITE_GOLDEN_OUTPUT=1`.
