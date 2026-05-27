*   Special shout-outs to @acpana, @anhdle-sso, @barney-s, @codebot-robot, @gemmahou, @himanigulati01, @justinsb, @katrielt, @ldanielmadariaga, @maqiuyujoyce, @suwandim, @xiaoweim for their contributions to this release.

## New Alpha Resources (Direct Reconciler):

*   `BinaryAuthorizationPlatformPolicy`
*   `VertexAIDeploymentResourcePool`

## New Fields:

*   [`ComputeReservation`](https://cloud.google.com/config-connector/docs/reference/resource-docs/compute/computereservation)
    *   Added `spec.shareSettings` field.

*   [`ComputeForwardingRule`](https://cloud.google.com/config-connector/docs/reference/resource-docs/compute/computeforwardingrule)
    *   Added `status.target` field.

## Reconciliation Improvements

We have added support for direct reconciliation to more resources, with opt-in behaviour. The API is unchanged. To use the direct reconciler, add the `cnrm.cloud.google.com/reconciler: direct` annotation to the corresponding Config Connector object.

*   [`ComputeReservation`](https://cloud.google.com/config-connector/docs/reference/resource-docs/compute/computereservation)
*   [`FirestoreIndex`](https://cloud.google.com/config-connector/docs/reference/resource-docs/firestore/firestoreindex)

## Bug Fixes:

*   [SQLInstance](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8025): Fix case sensitivity in SQLInstance `availabilityType`.
*   [Preview Tool](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/7743): Fix crash on typed resources and hang on defaulting in preview mode.
*   [ComputeForwardingRule](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/7371): Fix target field matching in ComputeForwardingRule.
*   [ComputeFutureReservation](https://github.com/GoogleCloudPlatform/k8s-config-connector/pull/8479): Fix future reservation times validation.
