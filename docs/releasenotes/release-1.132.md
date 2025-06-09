*   Special shout-outs to acpana@, anhdle-sso@, barney-s@, cheftako@, djyau@,
    fqbright@, gemmahou@, googlyrahman@, jingyih@, justinsb@, maqiuyujoyce@,
    xiaoweim@, yuwenma@ for their contributions to this release.

## New Beta Resources (Direct Reconciler):

*   [`SpeechCustomClass`](https://cloud.google.com/config-connector/docs/reference/resource-docs/speech/speechcustomclass)
*   [`SpeechPhraseSet`](https://cloud.google.com/config-connector/docs/reference/resource-docs/speech/speechphraseset)
*   [`SpeechRecognizer`](https://cloud.google.com/config-connector/docs/reference/resource-docs/speech/speechrecognizer)
*   [`VertexAINotebooksInstance`](https://cloud.google.com/config-connector/docs/reference/resource-docs/notebooks/notebookinstance)
*   [`VertexAIMetadataStore`](https://cloud.google.com/config-connector/docs/reference/resource-docs/vertexai/vertexaimetadatastore)

## New Alpha Resources (Direct Reconciler):

*   `OrgPolicyPolicy`
*   `OrgPolicyCustomConstraint`
*   `SpeechRecognizer`
*   `StorageAnywhereCache`

## New Fields:

*   [SpannerInstance](https://cloud.google.com/config-connector/docs/reference/resource-docs/spanner/spannerinstance)
    For opt-in direct controller,
    *   Added `spec.labels` field.
    *   Added `spec.defaultBackupScheduleType` field.
*   [SecretManagerSecret](https://cloud.google.com/config-connector/docs/reference/resource-docs/secretmanager/secretmanagersecret)
    For opt-in direct controller,
    *   Added `spec.labels` field.

## Reconciliation Improvements:

*   [BigtableAppProfile](https://cloud.google.com/config-connector/docs/reference/resource-docs/bigtable/bigtableappprofile)
    *   You can opt-in the direct controller by adding the
        `alpha.cnrm.cloud.google.com/reconciler: direct` annotation to the
        `BigtableAppProfile` resource.
    *   Added support for `spec.dataBoostIsolationReadOnly` field for resources
        reconciled by the direct controller.
*   [CloudIdentityGroup](https://cloud.google.com/config-connector/docs/reference/resource-docs/cloudidentity/cloudidentitygroup) and [CloudIdentityMembership](https://cloud.google.com/config-connector/docs/reference/resource-docs/cloudidentity/cloudidentitymembership)
    *   You can opt-in the direct controller by adding the
        `alpha.cnrm.cloud.google.com/reconciler: direct` annotation to the
        `CloudIdentityGroup` and `CloudIdentityMembership` resources.
    *   With direct reconciliation, creating new resources will no longer write
        back the service-generated ID to `spec.resourceID`. To acquire a
        resource, you can find its resourceID from the last part of
        `status.externalRef` field, or via gcloud command or Cloud Console. The
        `spec.resourceID` field is used for acquisition only, leave the field
        unset when creating a new resource.
