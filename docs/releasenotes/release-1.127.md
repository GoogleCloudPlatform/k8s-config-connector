# v1.126.0

** This version is not yet released; this document is gathering release notes for the future release **

* ...

* Special shout-outs to ... for their contributions to this release.
TODO: list contributors with `git log v1.126.0... | grep Merge | grep from | awk '{print $6}' | cut -d '/' -f 1 | sort | uniq`

## Announcement 

## New Beta Resources (Direct Reconciler):

* `Placeholder`

## Modified Beta Reconciliation


We have added support for direct reconciliation to more resources, with opt-in behaviour. The API is unchanged. To use the direct reconciler, add the `alpha.cnrm.cloud.google.com/reconciler: direct` annotation to the corresponding Config Connector object. The following resources now have direct reconciliation support (and we list some of the issues that this fixes): 

* [`ComputeTargetTCPProxy`](https://cloud.google.com/compute/docs/reference/rest/v1/targetTcpProxies)

  * Use regional TargetRCPProxy via the `spec.location` configuration.


## New Fields:

* `Placeholder`

## New Alpha Resources (Direct Reconciler):

* `Placeholder`
