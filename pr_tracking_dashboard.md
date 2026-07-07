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
| [#6921](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/6921) | Add autoclass.terminalStorageClass to StorageBucket | `issue-6921` | [#11183](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/11183) | Completed | Success (all checks passed) | PR is green and ready for merge. |
| [#2943](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/2943) | GCP DNS Authorization Per Project Record | `dns-authorization-per-project` | [#11186](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/11186) | Phase 6 | Monitoring (running) | Scoped test fields removal to dnsauthorization only in normalize.go to avoid breaking other tests, and force-pushed. |
| [#3480](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/3480) | Config connector fails to reconcile ContainerNodePool resources | `issue-3480` | [#11189](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/11189) | Completed | Success (193/193 passed) | PR is 100% green and ready for squash/merge. |
| [#4999](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/4999) | Unable to reference regional certificate manager certificates to a regional HTTPS proxy | `issue-4999` | [#11198](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/11198) | Phase 6 | Monitoring (running) | Enforced CreateInOrder on regional target HTTPS proxy tests to resolve non-deterministic golden log order, regenerated logs, and force-pushed. |
| [#5186](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/5186) | Missing resourceLabels support in ContainerCluster CRD | `issue-5186` | [#11201](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/11201) | Completed | Success (all checks passed) | PR is green and ready for merge. |
| [#6132](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/6132) | Migrate ComputeNetworkEndpointGroup from TF-based to Direct | `computenetworkendpointgroup-direct` | [#11292](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/11292) | Phase 3 | Monitoring (running) | Zonal NEG direct controller and MockGCP endpoints implemented; E2E mock test verified successfully locally. |
| [#6520](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/6520) | Support for additional GKE kubelet and sysctl fields in ContainerCluster and ContainerNodePool | `gke-kubelet-sysctl-fields` | [#11293](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/11293) | Phase 3 | Monitoring (running) | Backported image GC fields to vendored GKE TF provider; updated GKE KRM types and CRD schemas; verified all 32 container E2E tests pass. |
| [#6635](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/6635) | RedisInstance: secondaryIpRange not applied during initial creation, causing UpdateFailed errors | `redisinstance-secondaryiprange` | [#11294](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/11294) | Phase 3 | Monitoring (running) | Aligned SecondaryIpRange in populateInstanceDefaults to prevent update diff mismatch against server-allocated IP; added mock LRO auto-resolution; verified all 6 Redis E2E tests pass. |
| [#6921](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/6921) | [Feature Request] Add autoclass.terminalStorageClass to StorageBucket | `storagebucket-autoclass-terminalstorageclass` | [#11312](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/11312) | Phase 3 | Monitoring (running) | Added terminalStorageClass to StorageBucket spec, vendored TF provider schema, expander, and flattener; verified all StorageBucket E2E test cases pass cleanly against mock. |
| [#7604](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/7604) | [Feature Request] Support for resourceManagerTags on ContainerCluster & ContainerNodePool | `container-resource-manager-tags` | [#11314](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/11314) | Phase 3 | Monitoring (running) | Added resourceManagerTags to GKE ContainerCluster and ContainerNodePool spec, GKE TF provider schema, and direct fuzzer mappers; implemented mockGCP update handlers; verified E2E tests pass. |
| [#7605](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/7605) | ContainerCluster: Support for disableL4LbFirewallReconciliation | `gke-disable-l4-lb-firewall-reconciliation` | [#11321](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/11321) | Phase 3 | Monitoring (running) | Added disableL4LbFirewallReconciliation to GKE ContainerCluster spec, GKE TF provider schema, and direct expanders/flatteners; implemented mockGCP update handlers; verified E2E tests pass. |
| [#6897](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/6897) | [Feature Request] support for SLO Chart / Top List widget in MonitoringDashboard | N/A | N/A | Completed | N/A | Verified `timeSeriesTable` and `XyChart` are already supported in `MonitoringDashboard` KRM. E2E test `monitoringdashboard-toplist` passes. |
| [#7719](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/7719) | subnet creation based on InternalRanges | N/A | N/A | Completed | N/A | Verified `reservedInternalRangeRef` is already supported at both subnet and secondary range levels. E2E test `computesubnetwork-reservedinternalrangeref` passes. |
| [#8653](https://github.com/GoogleCloudPlatform/k8s-config-connector/issues/8653) | [Feature Request] error reporting notification settings | `errorreporting-notification-settings` | N/A | Phase 5 | Local Pass | Implemented direct controller `ErrorReportingNotificationSettings` and MockGCP service. Verified MockGCP E2E test passes locally. |






