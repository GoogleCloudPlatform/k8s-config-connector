# v1.119.0


* Special shout-outs to @acpana, @anhdle-sso, @barney-s, @cheftako, @gemmahou, @hankfreund, @jasonvigil, @jingyih, @justinsb, @maqiuyujoyce, @varsharmavs, @xiaoweim, @yuwenma, @zicongmei, @ziyue-101 for their
  contributions to this release.

## New features:

* Allow more customization of resource reconciliation
  * Added a new `ControllerReconciler` CRD (v1alpha1). See [example](https://github.com/GoogleCloudPlatform/k8s-config-connector/blob/master/operator/config/samples/namespaced_controller_reconciler_customization_sample.yaml)
  * This feature allows users to customize the client-side kube-apiserver request rate limit.

* Continue moving towards Direct Actuation as our preferred mechanism.
  * The default reconciler now uses Direct Actuation, if the Config Connector CRD does not have a `cnrm.cloud.google.com/tf2crd: "true"` or `cnrm.cloud.google.com/dcl2crd: "true"` label.

## New Resources:

* `CloudBuildWorkerPool`
  * Added `CloudBuildWorkerPool` (v1alpha1) resource for service `cloudbuild`.
  * This resource uses Direct Actuation.

## New Fields:

* `ComputeForwardingRule`
  * Added the `spec.target.serviceAttachmentRef` field, allowing a `ComputeForwardingRule` to target a `ComputeServiceAttachment`.

* `ContainerCluster`
  * Added previous output-only spec fields to `status.observedState`   
    * Added `status.observedState.masterAuth.clusterCaCertificate`
    * Added `status.observedState.privateClusterConfig.privateEndpoint`
    * Added `status.observedState.privateClusterConfig.publicEndpoint`    
