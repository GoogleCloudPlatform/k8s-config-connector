## Reconciliation Improvements

*   [SpannerInstance](https://cloud.google.com/config-connector/docs/reference/resource-docs/spanner/spannerinstance)
    * You can opt-in the direct controller by adding the
        `alpha.cnrm.cloud.google.com/reconciler: direct` annotation to the
        `SpannerInstance` resource`.
    * Direct controller is opt-in if using the following fields:
        * `spec.labels`
        * `spec.defaultBackupScheduleType`
        * `spec.edition`
        * `spec.autoscalingConfig`