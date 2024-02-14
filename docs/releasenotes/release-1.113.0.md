* Initial support for status.observedState in ContainerCluster,
  ContainerNodePool and RedisInstance; we are trying to encourage use of
  `cnrm.cloud.google.com/state-into-spec: absent` and eventually
  make it the default.  Some important resource information (such as the
  certificate for connecting to a GKE cluster) is only currently available
  in spec, and we recommend instead reading it from observedState
  where this is available.  We expect to add more fields to observedState
  in the future.

* Isolate terraform provider into a git subtree so we can more directly fix
  problems.

* Special shout-outs to svwijk@, katrielt@, sofam@, higef@ for their
  contributions to this release.

## New Resources:

* Added support for `ComputeNetworkFirewallPolicy` (v1beta1) resource.
* Added support for `TagsLocationTagBinding` (v1alpha1) resource.

## New Fields:

* RunJob (CloudRun Job)
  * Added `spec.template.vpcAccess.connectorRef` field.