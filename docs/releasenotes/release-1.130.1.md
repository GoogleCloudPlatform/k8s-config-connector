# Follow `SAMPLE_XXX` format to write the content.
# Delete `SAMPLE_` and its content before publishing the release note.
# Delete the entire header if no updates.

** This version is not yet released; this document is gathering release notes for the future release **


* Special shout-outs to acpana@, anhdle-sso@, barney-s@, cheftako@, ericpang777@, gemmahou@, jasonvigil@, jingyih@, justinsb@, maqiuyujoyce@, xiaoweim@, yuwenma@ for their contributions to this release.
TODO: list contributors with `git log v1.128.0... | grep Merge | grep from | awk '{print $6}' | cut -d '/' -f 1 | sort | uniq`.

## Announcement

<!-- ### SAMPLE_Simplified and More Reliable Resource Development -->

## New Beta Resources (Direct Reconciler):


## New Alpha Resources (Direct Reconciler):

## New Fields:
* [SpannerInstance](https://cloud.google.com/config-connector/docs/reference/resource-docs/spanner/spannerinstance). For opt-in direct controller,
  * Added `spec.labels` field.
  * Added `spec.defaultBackupScheduleType` field.
* [SecretManagerSecret](https://cloud.google.com/config-connector/docs/reference/resource-docs/secretmanager/secretmanagersecret). For opt-in direct controller,
  * Added `spec.labels` field.

<!-- ## Reconciliation Improvements -->

<!-- ## New features: -->

<!-- ## Bug Fixes: -->