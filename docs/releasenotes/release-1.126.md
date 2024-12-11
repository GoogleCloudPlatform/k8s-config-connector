# v1.126.0

* Special shout-outs to @acpana, @Camila-B, @cheftako, @ericpang777, @gemmahou, @himanikh, @jasonvigil, @jingyih, @jsoref, @justinsb, @maqiuyujoyce, @nb-goog, @xiaoweim, @yuwenma, @ziyue-101 for their contributions to this release.

## Announcement 

* Config Connector system management CRDs `ControllerReconciler` and `NamespacedControllerReconciler` are promoted to Beta. See how to [configure the Controller manager rate limit](https://cloud.google.com/config-connector/docs/how-to/customize-controller-manager-rate-limit).

## New Beta Resources (Direct Reconciler):

* [`BigQueryTransferConfig`](https://cloud.google.com/bigquery/docs/reference/datatransfer/rest/v1/projects.locations.transferConfigs)

 * Manage the metadata needed to perform a Big Query data transfer.

* [`KMSKeyHandle`](https://cloud.google.com/bigquery/docs/reference/datatransfer/rest/v1/projects.locations.transferConfigs)

 * Manage the provisioning of a CryptoKey.

## New Fields:

* `IAMPolicyMember`

  * Use Service Account from BigQueryConnectionConnection via `spec.memberFrom.bigQueryConnectionConnectionRef`. See [example](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/d2d887f6e416d329930b14d1e7ab99a3b4591ecb/config/samples/resources/iampolicymember/policy-member-with-bigqueryconnection-reference/iam_v1beta1_iampolicymember.yaml#L23) on IAMPolicyMember use BigqueryConectionConnection "cloudSQL"

* `IAMPartialPolicy`

  * Use Service Account from BigQueryConnectionConnection via `spec.memberFrom.bigQueryConnectionConnectionRef`.

## New Alpha Resources (Direct Reconciler):

* Add new resource [`WorkstationConfig`](https://cloud.google.com/workstations/docs/create-configuration)
