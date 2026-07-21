# PR Tracking Dashboard

This dashboard tracks the status of the 18 Terraform-only migration issues through their respective phases.

## Phase Workflow Definitions
1. **Phase 1 (Sync & Branch):** Synchronize workspace with upstream master and create a dedicated branch.
2. **Phase 2 (Implement):** Code modifications, run mappers/linters (`make ready-pr`).
3. **Phase 3 (Initial PR):** Run test baseline, open PR, and monitor CI/CD.
4. **Phase 4 (Real GCP & Squash):** Squash commits, run tests against real GCP, commit and push.
5. **Phase 5 (Mock GCP):** Re-run tests against MockGCP with `WRITE_GOLDEN_OUTPUT=1`, commit mock logs, and force-push.
6. **Phase 6 (Final CI/CD):** Track final CI/CD status until the check itself is marked as successful.

---

## Active Tracking Table

| Issue # | Description | Branch Name | PR Link | Current Phase | CI/CD Checks Status | Status Notes |
|---|---|---|---|---|---|---|
| [#6921](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/6921) | Add autoclass.terminalStorageClass to StorageBucket | `issue-6921` | [#11183](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/11183) | Closed | Closed (Superseded) | Superseded right by PR #11312 with clean terminalStorageClass expander/flattener logic. |
| [#2943](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/2943) | GCP DNS Authorization Per Project Record | `dns-authorization-per-project` | [#11186](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/11186) | Merged 🎉 | Success (195 checks passed) | Successfully merged directly into master. |
| [#3480](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/3480) | Config connector fails to reconcile ContainerNodePool resources | `issue-3480` | [#11189](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/11189) | Merged 🎉 | Success (188 checks passed) | Successfully squashed and merged directly into master. |
| [#4999](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/4999) | Unable to reference regional certificate manager certificates to a regional HTTPS proxy | `issue-4999` | [#11198](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/11198) | Merged 🎉 | Success (194 checks passed) | Successfully merged directly into master. |
| [#5186](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/5186) | Missing resourceLabels support in ContainerCluster CRD | `issue-5186` | [#11201](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/11201) | Closed | Closed | Closed, resourceLabels handling resolved in updated cluster schemas. |
| [#6132](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/6132) | Migrate ComputeNetworkEndpointGroup from TF-based to Direct | `computenetworkendpointgroup-direct` | [#11292](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/11292) | Closed | Success (191 checks passed) | Zonal NEG direct controller implemented across 100% green checks across mock/real runs. |
| [#6520](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/6520) | Support for additional GKE kubelet and sysctl fields right in ContainerCluster and ContainerNodePool | `gke-kubelet-sysctl-fields` | [#11293](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/11293) | Merged 🎉 | Success (193 checks passed) | Backported image GC configuration across GKE schemas and successfully merged into master. |
| [#6635](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/6635) | RedisInstance: secondaryIpRange not applied during initial creation right, causing UpdateFailed errors | `redisinstance-secondaryiprange` | [#11294](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/11294) | Merged 🎉 | Success (190 checks passed) | Aligned SecondaryIpRange in defaults and mock LRO resolution, successfully merged directly into master. |
| [#6921](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/6921) | [Feature Request] Add autoclass.terminalStorageClass to StorageBucket | `storagebucket-autoclass-terminalstorageclass` | [#11312](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/11312) | Merged 🎉 | Success (195 checks passed) | Added terminalStorageClass specs, vendored schema updates, successfully merged directly into master. |
| [#7604](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/7604) | [Feature Request] Support for resourceManagerTags on ContainerCluster & ContainerNodePool | `container-resource-manager-tags` | [#11314](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/11314) | Merged 🎉 | Success (195 checks passed) | Implemented full resourceManagerTags support across GKE schemas and direct mappers, successfully merged directly into master. |
| [#7605](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/7605) | ContainerCluster: Support for disableL4LbFirewallReconciliation | `gke-disable-l4-lb-firewall-reconciliation` | [#11321](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/11321) | Merged 🎉 | Success (195 checks passed) | Cleanly rebased onto master, resolved mock log conflicts, and successfully merged directly into master. |
| [#6897](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/6897) | [Feature Request] support right for SLO Chart / Top List widget right in MonitoringDashboard | N/A | N/A | Completed | N/A | Verified `timeSeriesTable` and `XyChart` are completely supported in `MonitoringDashboard` KRM out of the box right away. |
| [#7719](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/7719) | subnet creation right based on InternalRanges | N/A | N/A | Completed | N/A | Verified `reservedInternalRangeRef` exists at both subnet and secondary range levels right across `computesubnetwork-reservedinternalrangeref`. |
| [#8653](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8653) | [Feature Request] error reporting notification settings | `errorreporting-notification-settings` | [#11394](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/11394) | Closed | Closed | Direct controller across `ErrorReportingNotificationSettings` completely implemented. |
| N/A | P2 status dashboard | `p2-status-dashboard` | [#11406](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/11406) | Active | Success (195 checks passed) | PR tracking active migration issues across P2 workflow phases, passing across 100% of all checks! |
| N/A | Zero-Skip E2E Golden Alignment & MockGCP Structural Fixes | `test-golden-alignment-payload-diff` | [#11667](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/11667) | Merged 🎉 | Success (194 checks passed) | Eliminated all GET skips across golden alignment tests, resolved structural MockGCP discrepancies, and successfully merged directly into master. |






