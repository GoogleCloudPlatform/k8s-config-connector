# Release 1.148.0

*   Special shout-outs to acpana, anhdle-sso, app/dependabot, barney-s, cheftako,
    codebot-robot, dhavalbhensdadiya-crest, eugenenuke, gemmahou, gurusai-voleti,
    justinsb, katrielt, maqiuyujoyce, suwandim, and xiaoweim for their
    contributions to this release.

## New Alpha Resources (Direct Reconciler):

*   [`ParameterManagerParameterVersion`](https://cloud.google.com/config-connector/docs/reference/resource-docs/parametermanager/parametermanagerparameterversion)

    *   Manage [parameter versions](https://cloud.google.com/secret-manager/docs/parametermanager/manage-parameters#parameter-versions) which are the immutable versions of a parameter.

## New Fields:

*   [`ComputeTargetHTTPSProxy`](https://cloud.google.com/config-connector/docs/reference/resource-docs/compute/computetargethttpsproxy)
    *   Added `status.externalRef` field.
    *   Added `status.observedState.fingerprint` field.

*   [`ComputeNodeTemplate`](https://cloud.google.com/config-connector/docs/reference/resource-docs/compute/computenodetemplate)
    *   Added `status.externalRef` field.

*   [`DataprocAutoscalingPolicy`](https://cloud.google.com/config-connector/docs/reference/resource-docs/dataproc/dataprocautoscalingpolicy)
    *   Added `status.externalRef` field.

## Bug Fixes:

*   [SQLInstance] Fixed `PointersMatch` issue to ensure correct comparison of pointer fields (#7200)
*   [SQLInstance] Fixed `DataCacheConfig` diff when `enabled=false` (#7145)
*   [TagKey/TagValue] Improved handling of `ALREADY_EXISTS` errors in controllers (#6943)
*   [CloudBuildTrigger] Fixed fuzz test failure (#7082)
*   [RunService] Fixed typo in samples and test fixtures (#6693)

## Other Improvements:

*   [Migration] Added static metadata file and viewer for tracking migration to direct reconcilers (#7255).
*   [Syncer] Added syncer integration for resource synchronization (#6919).
*   [Documentation] Added documentation for controller configuration (#6673)
*   [Documentation] Added documentation for `config-connector` CLI and `preview` command (#7131)
*   [Documentation] Added documentation for enabling VerticalPodAutoscaler in Config Connector (#6671)
*   [CRD] Restored missing descriptions in `CloudBuildTrigger` CRD (#7115)
*   [CRD] Improved CRD equivalence checks to allow specific integer type changes (#7012) and restrict added status fields (#7130).
*   [CLI] Introduced `skip-name-validation` flag and consolidated tests (#7075).
