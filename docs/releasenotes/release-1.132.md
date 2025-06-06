# Follow `SAMPLE_XXX` format to write the content.
# Delete `SAMPLE_` and its content before publishing the release note.
# Delete the entire header if no updates.

** This version is not yet released; this document is gathering release notes for the future release **


* Special shout-outs to acpana@, anhdle-sso@, barney-s@, cheftako@, ericpang777@, gemmahou@, jingyih@, justinsb@, maqiuyujoyce@, xiaoweim@, yuwenma@ for their contributions to this release.

## Announcement

<!-- ### SAMPLE_Simplified and More Reliable Resource Development -->

## New Beta Resources (Direct Reconciler):
 * [`SpeechCustomClass`](https://cloud.google.com/dotnet/docs/reference/Google.Cloud.Speech.V2/latest/Google.Cloud.Speech.V2.CustomClass)
 * [`SpeechPhraseSet`](https://cloud.google.com/dotnet/docs/reference/Google.Cloud.Speech.V2/latest/Google.Cloud.Speech.V2.PhraseSet)
 * [`SpeechRecognizer`](https://cloud.google.com/speech-to-text/v2/docs/recognizers)
 * `VertexAINotebooksInstance`
 * `VertexAIMetadataStore`


## New Alpha Resources (Direct Reconciler):
 * `OrgPolicyPolicy`
 * `OrgPolicyCustomConstraint`
 * `SpeechRecognizer`
 * `StorageAnywhereCache`

## New Fields:
* [SpannerInstance](https://cloud.google.com/config-connector/docs/reference/resource-docs/spanner/spannerinstance). For opt-in direct controller,
  * Added `spec.labels` field.
  * Added `spec.defaultBackupScheduleType` field.
* [SecretManagerSecret](https://cloud.google.com/config-connector/docs/reference/resource-docs/secretmanager/secretmanagersecret). For opt-in direct controller,
  * Added `spec.labels` field.

## Reconciliation Improvements:
* [BigtableAppProfile](https://cloud.google.com/config-connector/docs/reference/resource-docs/bigtable/bigtableappprofile).
    * You can opt-in the direct controller by adding the `alpha.cnrm.cloud.google.com/reconciler: direct` annotation to the `BigtableAppProfile` resource.
    * Added support for `spec.dataBoostIsolationReadOnly` field for resources reconciled by the direct controller.
* [CloudIdentityGroup](https://cloud.google.com/config-connector/docs/reference/resource-docs/cloudidentity/cloudidentitygroup).
* [CloudIdentityMembership](https://cloud.google.com/config-connector/docs/reference/resource-docs/cloudidentity/cloudidentitymembership)
  * You can opt-in the direct controller by adding the `alpha.cnrm.cloud.google.com/reconciler: direct` annotation to the `CloudIdentityGroup` and `CloudIdentityMembership` resources.

<!-- ## New features: -->

<!-- ## Bug Fixes: -->
