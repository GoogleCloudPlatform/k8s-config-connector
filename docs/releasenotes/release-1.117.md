# v1.117.0

* This release improves our support for VertexAI.

* Special shout-outs to @199201shubhamsahu, @acpana, @anhdle-sso, @barney-s,
  @cheftako, @gemmahou, @jingyih, @justinsb, @katrielt, @maqiuyujoyce,
  @nicslatts, @xiaoweim, @yuwenma, @zicongmei and @ziyue-101
  for their contributions to this release.

## Resources promoted from alpha to beta:

* `VertexAIDataSet`
  * Output fields are now in `status.observedState`.
  * The KMS key is now specified using a reference: `spec.encryptionSpec.kmsKeyNameRef`

* `VertexAIIndex`
  * Output fields are now in `status.observedState`.
  * Note that `isCompleteOverwrite` is currently not supported: it is not
    obviously compatible with declarative operation.

* `VertexAIEndpoints`
  * Output fields are now in `status.observedState`.
  * The KMS key is now specified using a reference: `spec.encryptionSpec.kmsKeyNameRef`
  * The network is now specified using a reference: `spec.networkRef`

## New Fields:

* ComputeNetwork
  * The `spec.enableUlaInternalIpv6` field is no longer immutable - it can now
    be changed without recreating the network.

