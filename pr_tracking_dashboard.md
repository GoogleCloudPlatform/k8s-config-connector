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

| Issue # | Description | Branch Name | PR Link | PR Status | Current Phase | CI/CD Checks Status | Status Notes |
|---|---|---|---|---|---|---|---|
| [#6921](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/6921) | Add autoclass.terminalStorageClass to StorageBucket | `issue-6921` | [#11183](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/11183) | CLOSED | Completed | Success | PR is closed on upstream. |
| [#2943](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/2943) | GCP DNS Authorization Per Project Record | `dns-authorization-per-project` | [#11186](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/11186) | OPEN | Phase 6 | Success | All GHA checks passed on the PR. |
| [#3480](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/3480) | Config connector fails to reconcile ContainerNodePool resources | `issue-3480` | [#11189](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/11189) | MERGED | Completed | Success | PR has been successfully merged into master. |
| [#4999](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/4999) | Unable to reference regional certificate manager certificates to a regional HTTPS proxy | `issue-4999` | [#11198](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/11198) | OPEN | Phase 6 | Failed (tests-e2e-fixtures-compute failed) | GHA failed on tests-e2e-fixtures-compute. Needs investigation. |
| [#5186](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/5186) | Missing resourceLabels support in ContainerCluster CRD | `issue-5186` | [#11201](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/11201) | CLOSED | Completed | Success | PR is closed on upstream. |
| [#6132](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/6132) | Migrate ComputeNetworkEndpointGroup from TF-based to Direct | `computenetworkendpointgroup-direct` | [#11292](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/11292) | CLOSED | Completed | Success | PR is closed on upstream. |
| [#6520](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/6520) | Support for additional GKE kubelet and sysctl fields in ContainerCluster and ContainerNodePool | `gke-kubelet-sysctl-fields` | [#11293](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/11293) | OPEN | Phase 6 | Success | GHA checks passed on upstream. |
| [#6635](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/6635) | RedisInstance: secondaryIpRange not applied during initial creation, causing UpdateFailed errors | `redisinstance-secondaryiprange` | [#11294](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/11294) | MERGED | Completed | Success | PR has been successfully merged into master. |
| [#6921](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/6921) | [Feature Request] Add autoclass.terminalStorageClass to StorageBucket | `storagebucket-autoclass-terminalstorageclass` | [#11312](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/11312) | OPEN | Phase 5 | Failed (unit-tests & validations failed) | GHA failed on unit-tests and validations. Needs investigation. |
| [#7604](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/7604) | [Feature Request] Support for resourceManagerTags on ContainerCluster & ContainerNodePool | `container-resource-manager-tags` | [#11314](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/11314) | OPEN | Phase 6 | Pending (15 running) | Updated mock logs committed and pushed. Aligned logs verification is currently running on GHA. |
| [#7605](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/7605) | ContainerCluster: Support for disableL4LbFirewallReconciliation | `gke-disable-l4-lb-firewall-reconciliation` | [#11321](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/11321) | OPEN | Phase 6 | Success | GHA checks passed on upstream. |
| [#6897](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/6897) | [Feature Request] support for SLO Chart / Top List widget in MonitoringDashboard | N/A | N/A | N/A | Completed | N/A | Verified `timeSeriesTable` and `XyChart` are already supported in `MonitoringDashboard` KRM. E2E test `monitoringdashboard-toplist` passes. |
| [#7719](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/7719) | subnet creation based on InternalRanges | N/A | N/A | N/A | Completed | N/A | Verified `reservedInternalRangeRef` is already supported at both subnet and secondary range levels. E2E test `computesubnetwork-reservedinternalrangeref` passes. |
| [#8653](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8653) | [Feature Request] error reporting notification settings | `errorreporting-notification-settings` | [#11394](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/11394) | OPEN | Phase 6 | Success | GHA checks passed on upstream. |






