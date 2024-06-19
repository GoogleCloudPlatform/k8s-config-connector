# v1.119.0


* Special shout-outs to @acpana, @anhdle-sso, @barney-s, @cheftako, @gemmahou, @hankfreund, @jasonvigil, @jingyih, @justinsb, @maqiuyujoyce, @varsharmavs, @xiaoweim, @yuwenma, @zicongmei, @ziyue-101 for their
  contributions to this release.

## New features:

* Customize the ConfigConnector Reconciliation
  * Added a new `ControllerReconciler` CRD (v1alpha1). See [example](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/master/operator/config/samples/namespaced_controller_reconciler_customization_sample.yaml)
  * This feature allows users customizing the client-side token bucket rate limit.
  * This feature releases with the ConfigConnector operator bundle.

* Make Direct Controller the default reconciler.
  * Added Direct Controller registration
  * Set the default reconciler to Direct Controller if the ConfigConnector CRD does not have `cnrm.cloud.google.com/tf2crd: "true"` or `cnrm.cloud.google.com/dcl2crd: "true"` label. 

## New Resources:

* `CloudBuildWorkerPool`
  * Added `CloudBuildWorkerPool` (v1alpha1) resource for service `cloudbuild`.
  * This resource uses Direct Controller.

* `MonitoringDashboard`
  * Added `MonitoringDashboard` (v1beta1) resource for service `monitoring`.
  * This resource uses Direct Controller.

* `ComputeServiceAttachment`
  * Added `ComputeServiceAttachment` (v1beta1) for service `compute`. 
  * Added `ComputeServiceAttachment` as dependency of `ComputeForwardingRule` through `spec.target.serviceAttachmentRef`.

## New Fields:

* `ComputeForwardingRule`
  * Added the `spec.target.serviceAttachmentRef` dependency of `ComputeForwardingRule`

* `ContainerCluster`
  * Added previous output-only spec fields to `status.observedState`   
    * Added `status.observedState.masterAuth.clusterCaCertificate`
    * Added `status.observedState.privateClusterConfig.privateEndpoint`
    * Added `status.observedState.privateClusterConfig.publicEndpoint`    

## Fixes:

* several fixes on MockGCP
* several improvements on releases and presubmit/postsubmit builds.