*   Special shout-outs to @acpana, @app/dependabot, @barney-s, @cheftako, @codebot-robot, @eugenenuke, @gemmahou, @justinsb, @katrielt, @ldanielmadariaga, @maqiuyujoyce, @suwandim, @xiaoweim for their contributions to this release.

## New Beta Resources (Direct Reconciler):

*   [`ComputeReservation`](https://cloud.google.com/config-connector/docs/reference/resource-docs/compute/computereservation)
    *   Manage [Compute Engine Reservations](https://cloud.google.com/compute/docs/instances/reserving-zonal-resources).

## New Alpha Resources (Direct Reconciler):

*   `CloudDeployAutomation`
*   `ComputeFutureReservation`
*   `GKEHubMembershipBinding`
*   `GKEHubNamespace`
*   `GKEHubScopeRBACRoleBinding`
*   `NetworkServicesWasmPlugin`
*   `VertexAIDataLabelingJob`

## New Fields:

*   [`MemorystoreInstance`](https://cloud.google.com/config-connector/docs/reference/resource-docs/memorystore/memorystoreinstance)
    *   Added `spec.automatedBackupConfig` field.
    *   Added `spec.crossInstanceReplicationConfig` field.
    *   Added `spec.maintenanceVersion` field.
    *   Added `status.observedState.availableMaintenanceVersions` field.
    *   Added `status.observedState.crossInstanceReplicationConfig` field.
    *   Added `status.observedState.effectiveMaintenanceVersion` field.
    *   Added `status.observedState.pscAttachmentDetails` field.

## Bug Fixes:

*   [Issue 7228](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/7228) Fix `BigQueryDataTransferConfig` resource duplication loop.
*   [Issue 7625](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/7625) Enable projectID to projectNumber transform in fields in Container LROs.
*   [Issue 7907](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/7907) Refactor `DataplexLake`, `DataplexZone`, `DataplexEntryGroup`, and `DataplexEntryType` to use the identity and references pattern.
