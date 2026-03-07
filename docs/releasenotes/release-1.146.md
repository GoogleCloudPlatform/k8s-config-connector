** This version is not yet released; this document is gathering release notes
for the future release **

*   Special shout-outs to @acpana, @anhdle-sso, @cheftako, @dhavalbhensdadiya-crest, @dk-nc, @gemmahou, @himanikh, @justinsb, @katrielt, @maqiuyujoyce, @xiaoweim, and @yuwenma for their contributions to this release.

## New Alpha Resources (Direct Reconciler):

*   [`ParameterManagerParameter`](https://cloud.google.com/config-connector/docs/reference/resource-docs/parametermanager/parametermanagerparameter)
    *   Manage [Parameter Manager Parameters](https://cloud.google.com/secret-manager/parameter-manager/docs/overview).

## New Fields:

*   [`ContainerCluster`](https://cloud.google.com/config-connector/docs/reference/resource-docs/container/containercluster)
    *   Added `spec.controlPlaneEndpointsConfig.dnsEndpointConfig.enableK8sTokensViaDns` field.

## Reconciliation Improvements:

*   [`ContainerCluster`](https://cloud.google.com/config-connector/docs/reference/resource-docs/container/containercluster)
    *   Made `spec.clusterAutoscaling.autoProvisioningDefaults.bootDiskKMSKeyRef` mutable.

*   `NetworkServicesWasmPlugin`
    *   Introduced identity and reference.

*   Added structured reporting diff to numerous direct controllers to enhance diff visibility.

*   Added foundational API and mapper for `CloudDeployTarget` and `NetworkServicesLBRouteExtension`.

## New features:

*   Added a MCP server for CRD change checks.
*   Added the ability to test CRD equivalence call from the CLI.

## Bug Fixes:

*   [`SQLInstance`](https://cloud.google.com/config-connector/docs/reference/resource-docs/sql/sqlinstance)
    *   Added client-side default for `RetainedBackups` and `RetentionUnit`, and validated the `edition` field.
    *   Added `replicaConfiguration` as an unmanageable field.
    *   Controller now correctly defaults the field `enablePrivatePathForGoogleCloudServices` to `false`.

*   [`CertificateManagerDnsAuthorization`](https://cloud.google.com/config-connector/docs/reference/resource-docs/certificatemanager/certificatemanagerdnsauthorization)
    *   Sanitized Kubernetes labels to avoid 400 errors from invalid characters.

*   `ConfigConnector` Core
    *   `preview` now performs an early exit when no resources are found to reconcile.
    *   Fixed CRD field description for shared parent.
    *   Fixed incorrect exit status in lint filter.
    *   Updated `mockgcp` to improve compute regional resource mocks and defaults.
  