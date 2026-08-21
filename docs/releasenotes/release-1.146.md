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
*   Added structured reporting diff to numerous direct controllers. Resource diffs between desired and actual state are now directly accessible in controller logs.

## Bug Fixes:

*   [`SQLInstance`](https://cloud.google.com/config-connector/docs/reference/resource-docs/sql/sqlinstance)
    *   Added client-side default for `RetainedBackups` and `RetentionUnit`, and validated the `edition` field.
    *   Added `replicaConfiguration` as an unmanageable field.
    *   Controller now correctly defaults the field `enablePrivatePathForGoogleCloudServices` to `false`.

*   [`CertificateManagerDNSAuthorization`](https://cloud.google.com/config-connector/docs/reference/resource-docs/certificatemanager/certificatemanagerdnsauthorization)
    *   Sanitized Kubernetes labels to avoid 400 errors from invalid characters.
