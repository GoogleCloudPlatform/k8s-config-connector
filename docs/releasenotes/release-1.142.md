# v1.142.0

*   Special shout-outs to @acpana, @anhdle-sso, @cheftako, @gemmahou, @himanikh, @justinsb, @maqiuyujoyce, @vkanishk15, @xiaoweim, @yuwenma for their contributions to this release.

## New Beta Resources (Direct Reconciler):

*   `AlloyDBBackup`
*   `AccessContextManagerAccessLevel`

## New Fields:

*   `AlloyDBInstance`: Added `spec.observabilityConfig` and `spec.queryInsightsConfig` fields.
*   `ContainerNodePool`: Added `spec.nodeConfig.enableNestedVirtualization` field.
*   `MonitoringDashboard`: Added support for `spec.charts[].dataSets[].timeSeriesQuery.opsAnalyticsQuery.sqlQueryRef`

## Reconciliation Improvements

We have added support for direct reconciliation to more resources, with opt-in
behaviour. The API is unchanged. To use the direct reconciler, add the
`alpha.cnrm.cloud.google.com/reconciler: direct` annotation to the corresponding
Config Connector object. The following resources now have direct reconciliation
support (and we list some of the issues that this fixes):

*   `TagsLocationTagBinding`: Now supports direct reconciliation.

## New features:

*   IAM: Add support for `iam.cnrm.cloud.google.com/disable-dependent-services` annotation.
*   Added support for Cilium cluster-wide network policy.

## Bug Fixes:

*   `BatchJob`: Fixed a bug where the resource could not be created.
*   `FirewallPolicyRule`: Fixed an issue with updating the resource.
*   `IAMServiceAccountKey`: Fixed a bug that caused re-reconciliation.
*   Fixed a bug where `ComputeBackendService` could not refer to `clientTLSPolicy` due to an invalid format.
*   Fixed a bug where interconnect attachments were not ignored.
*   Fixed a bug in the GitHub MCP server.
*   Fixed a bug in the private cluster endpoint for mockgcp.
