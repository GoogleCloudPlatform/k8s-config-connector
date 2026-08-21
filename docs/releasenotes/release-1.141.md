** This version is not yet released; this document is gathering release notes for the future release **

*   Special shout-outs to @acpana, @anhdle-sso, @cheftako, @gemmahou, @GraceAtwood, @justinsb, @katrielt, @maqiuyujoyce, @marqc, @nancynh, @xiaoweim, and @yuwenma for their contributions to this release.

## New features:

*   Enabled Vertical Pod Autoscaler (VPA) support for Config Connector controllers.
    *   Added `verticalPodAutoscalerMode` field to `ConfigConnector` and `ConfigConnectorContext` resources.

## New Fields:

*   [`RunJob`](https://cloud.google.com/config-connector/docs/reference/resource-docs/run/runjob)
    *   Added `spec.template.spec.containers[].port` field.

*   [`DataplexTask`](https://cloud.google.com/config-connector/docs/reference/resource-docs/dataplex/dataplextask)
    *   Replaced `project` with `projectRef`.
    *   Replaced `serviceAccount` with `serviceAccountRef`.
    *   Replaced `kmsKey` with `kmsKeyRef`.

## Bug Fixes:

*   Fixed various issues in `observedState` handling for resources with reference fields.
*   Fixed an issue where IAMPolicy and IAMPartialPolicy controllers would alphabetize the members field within the resource spec and write it back. This behavior can conflict with intent-based reconciliation from GitOps systems such as Config Sync, causing a loop of updates and potentially exhausting IAM read quotas.

